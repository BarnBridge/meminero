package governance

import (
"database/sql"
"strconv"
	"time"

	"github.com/alethio/web3-go/ethrpc"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
"github.com/sirupsen/logrus"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type GovStorable struct {
	config Config
	Raw    *types.Block
	govAbi abi.ABI

	ethConn *ethclient.Client
	ethRPC *ethrpc.ETH
	logger              *logrus.Entry

	Preprocessed struct {
		BlockTimestamp int64
		BlockNumber    int64
	}
}

func New(config Config, raw *types.Block, govAbi abi.ABI,  ethConn *ethclient.Client) *GovStorable {
	return &GovStorable{
		config:  config,
		Raw:     raw,
		ethConn: ethConn,
		govAbi: govAbi,
		logger: logrus.WithField("module", "storable(governance)"),
	}
}

func (g GovStorable) Execute(tx *sql.Tx) error {
	g.logger.Trace("executing")
	start := time.Now()
	defer func() {
		g.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	governanceDecoder :=ethtypes.NewGovernanceDecoder()
	var govLogs []gethtypes.Log
	for _, data := range g.Raw.Txs {
		for _, log := range data.LogEntries {
			if utils.CleanUpHex(log.LoggedBy) == utils.CleanUpHex(g.config.GovernanceAddress) {
				govLogs = append(govLogs, log)
			}
		}
	}

	if len(govLogs) == 0 {
		log.Debug("no events found")
		return nil
	}

	err := g.handleProposals(govLogs, tx,governanceDecoder)
	if err != nil {
		return err
	}

	err = g.handleEvents(govLogs, tx)
	if err != nil {
		return err
	}

	err = g.handleVotes(govLogs, tx)
	if err != nil {
		return err
	}

	err := g.handleAbrogationProposal(govLogs, tx,governanceDecoder)
	if err != nil {
		return err
	}

	err = g.handleAbrogationProposalVotes(govLogs, tx)
	if err != nil {
		return err
	}

	return nil
}
