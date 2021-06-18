package state

import (
	"github.com/barnbridge/smartbackend/config"
)

func (m *Manager) Reset() error {
	err := m.redis.Del(config.Store.Redis.List).Err()
	if err != nil {
		return err
	}

	return nil
}
