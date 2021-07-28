package processor

import (
	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/processor/storables/accounterc20transfers"
	"github.com/barnbridge/meminero/processor/storables/dao/barn"
	"github.com/barnbridge/meminero/processor/storables/dao/governance"
	"github.com/barnbridge/meminero/processor/storables/erc20transfers"
	sePools "github.com/barnbridge/meminero/processor/storables/smartexposure/pool_state"
	seScrape "github.com/barnbridge/meminero/processor/storables/smartexposure/scrape"
	seTranches "github.com/barnbridge/meminero/processor/storables/smartexposure/tranche_state"
	syERC721 "github.com/barnbridge/meminero/processor/storables/smartyield/erc721"
	syEvents "github.com/barnbridge/meminero/processor/storables/smartyield/events"
	syRewards "github.com/barnbridge/meminero/processor/storables/smartyield/rewards"
	syState "github.com/barnbridge/meminero/processor/storables/smartyield/state"
	"github.com/barnbridge/meminero/processor/storables/tokenprices"
	"github.com/barnbridge/meminero/processor/storables/yieldfarming"
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

	if config.Store.Storable.Barn.Enabled {
		p.storables = append(p.storables, barn.New(p.Block))
	}

	if config.Store.Storable.Erc20Transfers.Enabled {
		p.storables = append(p.storables, erc20transfers.New(p.Block, p.state))
	}

	if config.Store.Storable.YieldFarming.Enabled {
		p.storables = append(p.storables, yieldfarming.New(p.Block))
	}

	if config.Store.Storable.SmartExposure.Enabled {
		p.storables = append(p.storables, seScrape.New(p.Block, p.state))
		p.storables = append(p.storables, seTranches.New(p.Block, p.state))
		p.storables = append(p.storables, sePools.New(p.Block, p.state))
	}

	p.registerSmartYield()

	if config.Store.Storable.TokenPrices.Enabled {
		p.storables = append(p.storables, tokenprices.New(p.Block, p.state))
	}
}

func (p *Processor) registerSmartYield() {
	if config.Store.Storable.SmartYield.Enabled {
		p.storables = append(p.storables, syEvents.New(p.Block, p.state))
		p.storables = append(p.storables, syERC721.New(p.Block, p.state))
		p.storables = append(p.storables, syRewards.New(p.Block, p.state))
		p.storables = append(p.storables, syState.New(p.Block, p.state))
	}
}
