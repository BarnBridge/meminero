package tokenprices

import (
	"context"
)

func (s *Storable) Execute(ctx context.Context) error {
	var err error
	s.processed.prices, err = GetTokensPrices(ctx, s.state.Tokens, s.block.Number)
	if err != nil {
		return err
	}

	return nil
}
