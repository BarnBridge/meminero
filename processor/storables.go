package processor

import (
	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/processor/storables/accounterc20transfers"
	"github.com/barnbridge/smartbackend/processor/storables/erc20transfers"
	"github.com/barnbridge/smartbackend/processor/storables/governance"
)

// registerStorables instantiates all the storables defined via code with the requested raw data
// Only the storables that are registered will be executed when the Store function is called
func (p *Processor) registerStorables() {
	if config.Store.Storable.AccountERC20Transfers.Enabled {
		p.storables = append(p.storables, accounterc20transfers.New(p.Block, p.state))
	}

	if config.Store.Storable.Governance.Enabled {
		p.storables = append(p.storables, governance.New(p.Block))
	}

	if config.Store.Storable.Erc20Transfers.Enabled {
		p.storables = append(p.storables, erc20transfers.New(p.Block, p.state))
	}
}
