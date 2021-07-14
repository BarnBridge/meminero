package utils

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

func DecodeString(output string) ([]byte, error) {
	output = Trim0x(output)

	data, err := hex.DecodeString(output)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode log data")
	}

	return data, nil
}

func DecodeFunctionOutput(a abi.ABI, methodName string, output string) (map[string]interface{}, error) {
	data, err := DecodeString(output)
	if err != nil {
		return nil, err
	}

	var decoded = make(map[string]interface{})
	err = a.UnpackIntoMap(decoded, methodName, data)
	if err != nil {
		return nil, errors.Wrap(err, "could not unpack log data")
	}

	return decoded, nil
}
func DecodeFunctionOutputToStruct(a abi.ABI, methodName string, output string, result interface{}) error {
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
