package eth

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"

	"github.com/barnbridge/smartbackend/utils"
)

func DecodeString(output string) ([]byte, error) {
	output = utils.Trim0x(output)

	data, err := hex.DecodeString(output)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode log data")
	}

	return data, nil
}

func DecodeFunctionOutputToInterface(a abi.ABI, methodName string, output string, result interface{}) error {
	data, err := DecodeString(output)
	if err != nil {
		return err
	}

	err = a.UnpackIntoInterface(result, methodName, data)
	if err != nil {
		return errors.Wrap(err, "could not unpack log data")
	}

	return nil
}

func ABIGenerateInput(a abi.ABI, methodName string, args ...interface{}) (string, error) {
	input, err := a.Pack(methodName, args...)
	if err != nil {
		return "", nil
	}

	return "0x" + hex.EncodeToString(input), nil
}
