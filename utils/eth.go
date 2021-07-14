package utils

import (
	"fmt"

	"github.com/alethio/web3-go/etherr"
	"github.com/alethio/web3-go/ethrpc"
	"github.com/pkg/errors"
)

func CallAtBlock(e *ethrpc.ETH, address string, fnc string, block int64) (string, error) {
	var result string

	obj := make(map[string]string)
	obj["to"] = address
	obj["data"] = fnc
	obj["gas"] = ethrpc.DefaultCallGas

	err := e.MakeRequest(&result, ethrpc.ETHCall, obj, fmt.Sprintf("0x%x", block))
	if err != nil {
		return "", errors.Wrapf(err, "could not make rpc request (%s.%s)", address, fnc)
	}

	if result == "0x" {
		return "", etherr.Empty
	}

	return result, nil
}