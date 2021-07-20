package notifications

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "notifs")

type Worker struct {
	db *sql.DB
}

func (w *Worker) Run(ctx context.Context) {
	// poll for new jobs
	go func() {
		for {
			select {
			case <-time.After(time.Second):
				tx, err := w.db.BeginTx(ctx, nil)
				if err != nil {
					log.Fatalf("start worker tx: %s", err)
				}

				jobs, err := w.jobs(ctx, tx)
				if err != nil {
					log.Fatalf("failed to get jobs: %s", err)
				}

				err = ExecuteJobsWithTx(ctx, tx, jobs...)
				if err != nil {
					log.Fatalf("failed to execute jobs: %s", err)
				}

				err = SoftDeleteJobsWithTx(ctx, tx, jobs...)
				if err != nil {
					log.Fatalf("failed to cleanup jobs: %s", err)
				}

				err = tx.Commit()
				if err != nil {
					log.Fatalf("failed to commit jobs: %s", err)
				}

			case <-ctx.Done():
				log.Info("received exit signal, stopping")
				return
			}
		}
	}()
}

func (w *Worker) jobs(ctx context.Context, tx *sql.Tx) ([]*Job, error) {
	var jobs []*Job
	sel := `
		SELECT
			"id",
			"type",
			"execute_on",
			"metadata",
			"included_in_block"
		FROM
			"notification_jobs"
		WHERE
			"execute_on" < EXTRACT(EPOCH FROM NOW())::bigint
			AND deleted = FALSE
		LIMIT 1000
		;
	`
	rows, err := tx.QueryContext(ctx, sel)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "get jobs")
	}

	for rows.Next() {
		var j Job
		err = rows.Scan(&j.Id, &j.JobType, &j.ExecuteOn, &j.JobData, &j.IncludedInBlock)
		if err != nil {
			return nil, errors.Wrap(err, "scan job row")
		}
		jobs = append(jobs, &j)
	}

	return jobs, nil
}

func NewWorker(config Config) (*Worker, error) {
	log.Info("connecting to postgres")
	db, err := sql.Open("postgres", config.PostgresConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "open postgres connection")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping postgres connection")
	}

	return &Worker{
		db: db,
	}, nil
}
