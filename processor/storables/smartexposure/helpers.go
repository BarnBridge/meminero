package smartexposure

import (
	"errors"

	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

func BuildTokensSliceForSE(state *state.Manager) (map[string]types.Token, error) {
	tokens := make(map[string]types.Token)
	for _, pool := range state.SmartExposure.Pools {
		token, exists := state.Tokens[pool.TokenA.Address]
		if !exists {
			return nil, errors.New("could not find tokenA in state for calculate price")
		}
		tokens[token.Address] = token

		token, exists = state.Tokens[pool.TokenB.Address]
		if !exists {
			return nil, errors.New("could not find tokenB in state for calculate price")
		}
		tokens[token.Address] = token
	}

	return tokens, nil
}
