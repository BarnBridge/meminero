package state

import (
	"github.com/barnbridge/smartbackend/types"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Manager struct {
	redis  *redis.Client
	logger *logrus.Entry
	db      *pgxpool.Pool

	Tokens        map[string]types.Token
	monitoredAccounts []string
}

var instance *Manager

func Init(db  *pgxpool.Pool) error {
	if instance != nil {
		return nil
	}

	instance = &Manager{db: db}

	return nil
}

func Refresh() error {
	return nil
}


// NewManager instantiates a new task manager and also takes care of the redis connection management
// it subscribes to the best block tracker for new blocks which it'll add to the redis queue automatically
func NewManager() (*Manager, error) {
	m := &Manager{
		logger: logrus.WithField("module", "state"),
	}

	var err error
	m.redis, err = NewRedis()
	if err != nil {
		return nil, errors.Wrap(err, "could not setup redis connection")
	}

	return m, nil
}

func (m *Manager) Close() error {
	return m.redis.Close()
}
