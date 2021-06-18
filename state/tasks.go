package state

import (
	"context"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/barnbridge/smartbackend/config"
)

func (m *Manager) NextTask(ctx context.Context) (int64, error) {
	m.logger.Trace("fetching task from state")

	var task int64

	errChan := make(chan error)

	go func() {
		taskResult, err := m.redis.BZPopMin(0, config.Store.Redis.List).Result()
		if err != nil {
			errChan <- err
			return
		}

		taskInt, err := strconv.ParseInt(taskResult.Member.(string), 10, 64)
		if err != nil {
			errChan <- err
			return
		}

		task = taskInt
		close(errChan)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return 0, errors.Wrap(err, "could not read task from redis")
		}

		m.logger.Trace("done fetching task")

		return task, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// AddTaskToQueue inserts a block into the redis sorted set used for queue management using a ZADD command
func (m *Manager) AddTaskToQueue(block int64) error {
	m.logger.WithField("block", block).Trace("adding block to todo")
	return m.redis.ZAdd(config.Store.Redis.List, redis.Z{
		Score:  float64(block),
		Member: block,
	}).Err()
}
