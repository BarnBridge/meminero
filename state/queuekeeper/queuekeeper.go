package queuekeeper

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/state"
	"github.com/lacasian/ethwheels/bestblock"
)

type Keeper struct {
	tracker *bestblock.Tracker
	logger  *logrus.Entry
	state   *state.Manager

	lastBlockAdded int64
}

func New(tracker *bestblock.Tracker, state *state.Manager) (*Keeper, error) {
	w := &Keeper{
		tracker: tracker,
		logger:  logrus.WithField("module", "queuekeeper"),
		state:   state,
	}

	return w, nil
}

// Run subscribes to the best block tracker for new blocks and adds tasks to state
func (m *Keeper) Run(ctx context.Context) {
	skipBlocks := config.Store.Feature.QueueKeeper.Lag
	var lastBlock int64
	var started bool

	blocks := m.tracker.Subscribe()
	m.lastBlockAdded = m.tracker.BestBlock()

	for {
		select {
		case b := <-blocks:
			log := m.logger.WithField("block", b)

			if !started || b-config.Store.Feature.QueueKeeper.Lag <= m.lastBlockAdded {
				started = true
				m.lastBlockAdded = b - config.Store.Feature.QueueKeeper.Lag - 1
			}

			if skipBlocks > 0 {
				if b > lastBlock {
					lastBlock = b
					skipBlocks--
				}
				log.Infof("postponing block because lag feature is enabled (%d to go)", skipBlocks)
				continue
			}

			log.Trace("got new block")

			for i := m.lastBlockAdded + 1; i <= b-config.Store.Feature.QueueKeeper.Lag; i++ {
				err := m.state.AddTaskToQueue(i)
				if err != nil {
					log.Error(err)
				} else {
					m.lastBlockAdded = i
				}
			}

			log.Trace("done adding block to todo")
		case <-ctx.Done():
			m.tracker.Unsubscribe(blocks)
		}
	}
}
