package eth

import (
	"fmt"

	"github.com/alethio/web3-go/etherr"
	"github.com/alethio/web3-go/ethrpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

func CallContractFunction(a abi.ABI, addr string, methodName string, methodArgs []interface{}, result interface{}, opts ...interface{}) func() error {
	return func() error {
		input, err := ABIGenerateInput(a, methodName, methodArgs...)
		if err != nil {
			return errors.Wrap(err, "could not generate input for contract call")
		}
		var data string

		if len(opts) > 0 {
			data, err = CallRawAtBlock(addr, input, opts[0].(int64))
		} else {
			data, err = CallRaw(addr, input)
		}

		if err != nil {
			return errors.Wrap(err, "could not execute contract call")
		}

		err = DecodeFunctionOutputToInterface(a, methodName, data, result)
		if err != nil {
			return errors.Wrap(err, "could not decode contract call output")
		}

		return nil
	}
}

func CallRawAtBlock(address string, fnc string, block int64) (string, error) {
	var result string

	obj := make(map[string]string)
	obj["to"] = address
	obj["data"] = fnc
	obj["gas"] = ethrpc.DefaultCallGas

	err := instance.ethrpc.MakeRequest(&result, ethrpc.ETHCall, obj, fmt.Sprintf("0x%x", block))
	if err != nil {
		return "", errors.Wrapf(err, "could not make rpc request (%s.%s)", address, fnc)
	}

	if result == "0x" {
		return "", etherr.Empty
	}

	return result, nil
}

func CallRaw(address string, fnc string) (string, error) {
	var result string

	obj := make(map[string]string)
	obj["to"] = address
	obj["data"] = fnc
	obj["gas"] = ethrpc.DefaultCallGas

	err := instance.ethrpc.MakeRequest(&result, ethrpc.ETHCall, obj, "latest")
	if err != nil {
		return "", errors.Wrapf(err, "could not make rpc request (%s.%s)", address, fnc)
	}

	if result == "0x" {
		return "", etherr.Empty
	}

	return result, nil
}
