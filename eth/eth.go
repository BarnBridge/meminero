package eth

import (
	"time"

	"github.com/alethio/web3-go/ethrpc"
	"github.com/alethio/web3-go/ethrpc/provider/httprpc"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/config"
)

type conn struct {
	ethrpc *ethrpc.ETH
}

var instance *conn

func Init() error {
	if instance != nil {
		return nil
	}

	batchLoader, err := httprpc.NewBatchLoader(config.Store.ETH.MaxBatch, 4*time.Millisecond)
	if err != nil {
		return errors.Wrap(err, "could not init batch loader")
	}

	provider, err := httprpc.NewWithLoader(config.Store.ETH.HTTP, batchLoader)
	if err != nil {
		return errors.Wrap(err, "could not init ethprc http provider")
	}

	eth, err := ethrpc.New(provider)
	if err != nil {
		return errors.Wrap(err, "could not create ethrpc")
	}

	instance = &conn{
		ethrpc: eth,
	}

	return nil
}
