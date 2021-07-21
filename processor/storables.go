package processor

import (
	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/processor/storables/accountERC20Transfers"
	"github.com/barnbridge/smartbackend/processor/storables/governance"
	"github.com/barnbridge/smartbackend/processor/storables/yieldfarming"
)

// registerStorables instantiates all the storables defined via code with the requested raw data
// Only the storables that are registered will be executed when the Store function is called
func (p *Processor) registerStorables() {
	if config.Store.Storable.AccountERC20Transfers.Enabled {
		p.storables = append(p.storables, accountERC20Transfers.New(p.Block, p.state))
	}

	if config.Store.Storable.Governance.Enabled {
		p.storables = append(p.storables, governance.New(p.Block))
	}

	if config.Store.Storable.YieldFarming.Enabled {
		p.storables = append(p.storables, yieldfarming.New(p.Block))
	}
}
