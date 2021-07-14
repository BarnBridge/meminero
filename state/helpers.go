package state

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/barnbridge/smartbackend/config"
)

func NewPostgres() (*sql.DB, error) {
	var log = logrus.WithField("module", "state")

	log.Info("connecting to postgres")
	db, err := sql.Open("postgres", viper.GetString("db.connection-string"))
	if err != nil {
		return nil, errors.Wrap(err, "could not init db connection")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "could not ping database")
	}

	err = automigrate(db)
	if err != nil {
		return nil, err
	}

	log.Info("connected to postgres successfuly")

	return db, nil
}

func NewPGX() (*pgxpool.Pool, error) {
	var log = logrus.WithField("module", "state")

	log.Info("connecting pgx")
	dbpool, err := pgxpool.Connect(context.Background(), viper.GetString("db.connection-string"))
	if err != nil {
		return nil, errors.Wrap(err, "could not init db connection")
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "could not ping database")
	}

	log.Info("connected to pgx successfuly")

	return dbpool, nil
}

func automigrate(db *sql.DB) error {
	var log = logrus.WithField("module", "state")

	if !config.Store.Database.Automigrate {
		return nil
	}

	log.Info("attempting automatic execution of migrations")

	err := goose.Up(db, "/")
	if err != nil && err != goose.ErrNoNextVersion {
		return errors.Wrap(err, "could not execute migrations")
	}

	log.Info("database version is up to date")

	return nil
}

func NewRedis() (*redis.Client, error) {
	var log = logrus.WithField("module", "state")

	log.Info("setting up redis connection")
	r := redis.NewClient(&redis.Options{
		Addr:        config.Store.Redis.Server,
		Password:    config.Store.Redis.Password,
		DB:          0,
		ReadTimeout: time.Second * 1,
	})

	err := r.Ping().Err()
	if err != nil {
		return nil, errors.Wrap(err, "could not ping redis")
	}

	log.Info("connected to redis successfully")

	return r, nil
}
