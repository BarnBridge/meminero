package state

import (
	"context"
	"sync"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/state/smartalpha"
	"github.com/barnbridge/meminero/state/smartexposure"

	"github.com/barnbridge/meminero/state/smartyield"
	"github.com/barnbridge/meminero/types"
)

type Manager struct {
	redis  *redis.Client
	logger *logrus.Entry
	db     *pgxpool.Pool
	mu     *sync.Mutex

	Tokens            map[string]types.Token
	monitoredAccounts map[string]bool
	monitoredERC20    map[string]bool

	SmartExposure *smartexposure.SmartExposure
	SmartYield    *smartyield.SmartYield
	SmartAlpha    *smartalpha.SmartAlpha
}

// NewManager instantiates a new task manager and also takes care of the redis connection management
// it subscribes to the best block tracker for new blocks which it'll add to the redis queue automatically
func NewManager(db *pgxpool.Pool) (*Manager, error) {
	m := &Manager{
		db:            db,
		logger:        logrus.WithField("module", "state"),
		mu:            new(sync.Mutex),
		SmartYield:    smartyield.New(),
		SmartExposure: smartexposure.New(),
		SmartAlpha:    smartalpha.New(db),
	}

	var err error
	m.redis, err = NewRedis()
	if err != nil {
		return nil, errors.Wrap(err, "could not setup redis connection")
	}

	return m, nil
}

func (m *Manager) RefreshCache(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	err := m.loadAllAccounts(ctx)
	if err != nil {
		return errors.Wrap(err, "could not fetch monitored accounts")
	}

	err = m.loadAllTokens(ctx)
	if err != nil {
		return errors.Wrap(err, "could not fetch tokens")
	}

	err = m.loadAllERC20(ctx)
	if err != nil {
		return errors.Wrap(err, "could not fetch monitored erc20")
	}

	err = m.SmartExposure.LoadPools(ctx, m.db)
	if err != nil {
		return errors.Wrap(err, "could not fetch smart exposure pools")
	}

	err = m.SmartExposure.LoadTranches(ctx, m.db)
	if err != nil {
		return errors.Wrap(err, "could not fetch smart exposure tranches")
	}

	err = m.SmartYield.LoadPools(ctx, m.db)
	if err != nil {
		return errors.Wrap(err, "could not fetch smart yield pools")
	}

	err = m.SmartYield.LoadRewardPools(ctx, m.db)
	if err != nil {
		return errors.Wrap(err, "could not fetch smart yield reward pools")
	}

	err = m.SmartAlpha.Load(ctx, m.db)
	if err != nil {
		return errors.Wrap(err, "could not load smart alpha state")
	}

	err = m.refreshDBCache(ctx)
	if err != nil {
		return errors.Wrap(err, "could not refresh db cache")
	}

	return nil
}

func (m *Manager) Close() error {
	return m.redis.Close()
}
