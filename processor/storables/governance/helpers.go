package governance

import (
	"fmt"

	"github.com/barnbridge/smartbackend/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

func (g GovStorable) callSimpleFunctionWithInput(a abi.ABI, contract string, name string, input string) (string, error) {
	data, err := utils.CallAtBlock(g.ethRPC, contract, input, g.Preprocessed.BlockNumber)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("could not call %g.%g", contract, name))
	}

	decoded, err := utils.DecodeFunctionOutput(a, name, data)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("could not decode output from %g.%g", contract, name))
	}

	return decoded["Description"].(string), nil
}