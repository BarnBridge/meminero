package governance

import (
	"context"
	"time"

	"github.com/alethio/web3-go/ethrpc"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type GovStorable struct {
	config Config
	block  *types.Block
	govAbi abi.ABI

	ethConn *ethclient.Client
	ethRPC *ethrpc.ETH
	logger              *logrus.Entry

	Processed struct {
		proposals []Proposal
		proposalsActions []ProposalActions
		abrProposals []ethtypes.GovernanceAbrogationProposalStartedEvent
		abrProposalsDescription map[string]string
	}
}

func New(config Config, block *types.Block, govAbi abi.ABI,  ethConn *ethclient.Client) *GovStorable {
	return &GovStorable{
		config:  config,
		block:   block,
		ethConn: ethConn,
		govAbi:  govAbi,
		logger:  logrus.WithField("module", "storable(governance)"),
	}
}

func (g GovStorable) Execute(ctx context.Context) error {
	g.logger.Trace("executing")
	start := time.Now()
	defer func() {
		g.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	var govLogs []gethtypes.Log
	for _, data := range g.block.Txs {
		for _, log := range data.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(g.config.GovernanceAddress) {
				govLogs = append(govLogs, log)
			}
		}
	}

	if len(govLogs) == 0 {
		log.Debug("no events found")
		return nil
	}

	err := g.handleProposals(ctx,govLogs)
	if err != nil {
		return err
	}


	err = g.handleAbrogationProposal(ctx,govLogs)
	if err != nil {
		return err
	}

/*	err = g.handleEvents(govLogs, tx)
	if err != nil {
		return err
	}*/

/*	err = g.handleVotes(govLogs, tx)
	if err != nil {
		return err
	}


	err = g.handleAbrogationProposalVotes(govLogs, tx)
	if err != nil {
		return err
	}*/

	return nil
}

func (g *GovStorable) Rollback(ctx context.Context,tx pgx.Tx) error {
	_, err := tx.Exec(ctx, `delete from account_erc20_transfers where included_in_block = $1`, g.block.Number)

	return err
}

func (g *GovStorable) SaveToDatabase(ctx context.Context,tx pgx.Tx) error {
	err := g.storeProposals(ctx,tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposals")
	}

	err = g.storeAbrogrationProposals(ctx,tx)
	if err != nil {
		return errors.Wrap(err,"could not store abrogration proposals")
	}
	return nil
}

func (g *GovStorable) Result() interface{} {
	return g.Processed
}