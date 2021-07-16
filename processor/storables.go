package processor

import (
	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/processor/storables/accountERC20Transfers"
)

// registerStorables instantiates all the storables defined via code with the requested raw data
// Only the storables that are registered will be executed when the Store function is called
func (p *Processor) registerStorables() {
	if config.Store.Storable.AccountERC20Transfers.Enabled {
		p.storables = append(p.storables, accountERC20Transfers.New(p.ctx, p.Block, p.state))
	}
}
