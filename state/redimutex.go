package state

import (
	"fmt"
	"time"
)

func (m *Manager) LockBlock(blockNumber int64) (bool, error) {
	key := fmt.Sprint(blockNumber)
	return m.RedisLock(key)
}

func (m *Manager) UnlockBlock(blockNumber int64) error {
	key := fmt.Sprint(blockNumber)
	return m.RedisUnlock(key)
}

func (m *Manager) RedisLock(key string) (bool, error) {
	key = fmt.Sprintf("lock:%s", key)
	return m.redis.SetNX(key, true, time.Second*30).Result()
}

func (m *Manager) RedisUnlock(key string) error {
	key = fmt.Sprintf("lock:%s", key)
	return m.redis.Del(key).Err()
}
