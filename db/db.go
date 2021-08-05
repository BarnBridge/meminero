package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
)

type DB struct {
	pool   *pgxpool.Pool
	logger *logrus.Entry
}

func New() (*DB, error) {
	db := &DB{
		logger: logrus.WithField("module", "db"),
	}

	pgxCfg, err := db.pgxPoolConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not build pgx config")
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), pgxCfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to pgx pool")
	}

	db.pool = pool

	return db, nil
}

func (db *DB) Connection() *pgxpool.Pool {
	return db.pool
}

func (db *DB) Migrate(ctx context.Context) error {
	if !config.Store.Database.Automigrate {
		db.logger.Infof("automigration is disabled")
		return nil
	}

	files, err := ioutil.ReadDir(config.Store.Database.MigrationsPath)
	if err != nil {
		return errors.Wrap(err, "reading migration packages")
	}

	err = db.MigratePackage(ctx, "public")
	if err != nil {
		return errors.Wrap(err, "could not migrate package 'public'")
	}

	for _, f := range files {
		db.logger.Debugf("processing file: %s", f.Name())

		if !f.IsDir() || f.Name() == "public" {
			db.logger.Warnf("file ignored: %s", f.Name())
			continue
		}

		err := db.MigratePackage(ctx, f.Name())
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("could not migrate package '%s'", f.Name()))
		}
	}

	return nil
}

func (db *DB) MigratePackage(ctx context.Context, packageName string) error {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return errors.Wrap(err, "acquire db connection")
	}

	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), db.getMigrationsVersionTable(packageName))
	if err != nil {
		return errors.Wrap(err, "creating migrator")
	}

	err = migrator.LoadMigrations(db.getMigrationsPath(packageName))
	_, ok := err.(migrate.NoMigrationsFoundError)
	if err != nil && !ok {
		return errors.Wrap(err, "loading migrations")
	}

	if len(migrator.Migrations) == 0 {
		db.logger.WithField("package", packageName).Info("no migrations to run")
		return nil
	}

	migrator.OnStart = func(sequence int32, name, direction, sql string) {
		db.logger.Infof("(%s) %s executing %s %s\n", packageName, time.Now().Format(time.ANSIC), name, direction)
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		e := errors.Wrap(err, "running migrations")
		if err, ok := err.(migrate.MigrationPgError); ok {
			if err.Detail != "" {
				db.logger.Errorf("DETAIL: %s", err.Detail)
			}
			if err.Position != 0 {
				ele, err := ExtractErrorLine(err.Sql, int(err.Position))
				if err != nil {
					return errors.Wrap(err, "extract error line")
				}
				prefix := fmt.Sprintf("LINE %d: ", ele.LineNum)
				db.logger.Errorf("%s%s\n", prefix, ele.Text)
				padding := strings.Repeat(" ", len(prefix)+ele.ColumnNum-1)
				db.logger.Errorf("%s^\n", padding)
			}
		}
		return e
	}

	db.logger.Infof("(%s) package is up to date", packageName)

	return nil
}

func (db *DB) getMigrationsVersionTable(name string) string {
	return name + ".migration_version"
}

func (db *DB) getMigrationsPath(name string) string {
	return config.Store.Database.MigrationsPath + "/" + name
}

// func (db *DB) WaitMigrationVersion(ctx context.Context, version int32) error {
// 	conn, err := db.pool.Acquire(ctx)
// 	if err != nil {
// 		return errors.Wrap(err, "acquire db connection")
// 	}
// 	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), db.config.MigrationsVersionTable)
// 	if err != nil {
// 		return errors.Wrap(err, "creating migrator")
// 	}
// 	shown := false
// 	for {
// 		if utils.ContextIsDone(ctx) {
// 			return nil
// 		}
// 		v, err := migrator.GetCurrentVersion(ctx)
// 		if err != nil {
// 			return errors.Wrap(err, "get migration version")
// 		}
// 		if v >= version {
// 			log.Infof("minimum migration version met, wanted: %d actual: %d", version, v)
// 			break
// 		}
// 		if !shown {
// 			log.Warnf("minimum migration version below treshold, wanted: %d actual: %d", version, v)
// 			shown = true
// 		}
// 		time.Sleep(time.Second * 15)
// 	}
// 	return nil
// }
