package state

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
)

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
