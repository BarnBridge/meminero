package state

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/config"
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

func (m *Manager) AddBatchToQueue(blocks []int64) error {
	start := time.Now()

	var members []redis.Z
	for _, i := range blocks {
		members = append(members, redis.Z{
			Score:  float64(i),
			Member: i,
		})
	}

	const batchSize = 500

	batches := len(blocks)/batchSize + 1

	for i := 0; i < batches; i++ {
		end := batchSize * (i + 1)
		if end > len(members) {
			end = len(members)
		}
		m.logger.Infof("queueing batch [%d, %d]", members[batchSize*i].Member, members[end-1].Member)

		err := m.redis.ZAdd(config.Store.Redis.List, members[batchSize*i:end]...).Err()
		if err != nil && err != redis.Nil {
			return err
		}
	}

	m.logger.WithField("duration", time.Since(start)).Info("queued all blocks")

	return nil
}
