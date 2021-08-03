package events

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) Execute(ctx context.Context) error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.SmartYield.PoolByAddress(log.Address.String()) != nil {
				err := s.decodePoolEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode pool event")
				}
			}

			if s.state.SmartYield.PoolByControllerAddress(log.Address.String()) != nil {
				err := s.decodeControllerEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode controller event")
				}
			}

			if s.state.SmartYield.PoolByProviderAddress(log.Address.String()) != nil {
				err := s.decodeProviderEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode provider event")
				}
			}
		}
	}

	return nil
}

func (s *Storable) decodePoolEvent(log gethtypes.Log) error {
	sy := ethtypes.SmartYield

	if sy.IsBuyTokensEvent(&log) {
		e, err := sy.BuyTokensEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuyTokens event")
		}

		s.processed.JuniorEntryEvents = append(s.processed.JuniorEntryEvents, e)

		return nil
	}

	if sy.IsSellTokensEvent(&log) {
		e, err := sy.SellTokensEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode SellTokens event")
		}

		s.processed.JuniorInstantWithdrawEvents = append(s.processed.JuniorInstantWithdrawEvents, e)

		return nil
	}

	if sy.IsBuyJuniorBondEvent(&log) {
		e, err := sy.BuyJuniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuyJuniorBond event")
		}

		s.processed.Junior2StepWithdrawEvents = append(s.processed.Junior2StepWithdrawEvents, e)

		return nil
	}

	if sy.IsRedeemJuniorBondEvent(&log) {
		e, err := sy.RedeemJuniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode RedeemJuniorBond event")
		}

		s.processed.Junior2StepRedeemEvents = append(s.processed.Junior2StepRedeemEvents, e)

		return nil
	}

	if sy.IsBuySeniorBondEvent(&log) {
		e, err := sy.BuySeniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuySeniorBond event")
		}

		s.processed.SeniorEntryEvents = append(s.processed.SeniorEntryEvents, e)

		return nil
	}

	if sy.IsRedeemSeniorBondEvent(&log) {
		e, err := sy.RedeemSeniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode RedeemSeniorBond event")
		}

		s.processed.SeniorRedeemEvents = append(s.processed.SeniorRedeemEvents, e)

		return nil
	}

	if sy.IsTransferEvent(&log) {
		e, err := sy.TransferEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode Transfer event")
		}

		s.processed.Transfers = append(s.processed.Transfers, e)

		return nil
	}

	return nil
}

func (s *Storable) decodeControllerEvent(log gethtypes.Log) error {
	c := ethtypes.SmartYieldCompoundController

	if c.IsHarvestEvent(&log) {
		e, err := c.HarvestEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode Harvest event")
		}

		s.processed.ControllerHarvests = append(s.processed.ControllerHarvests, e)
	}

	return nil
}

func (s *Storable) decodeProviderEvent(log gethtypes.Log) error {
	p := ethtypes.SmartYieldCompoundProvider

	if p.IsTransferFeesEvent(&log) {
		e, err := p.TransferFeesEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode TransferFees event")
		}

		s.processed.ProviderTransferFees = append(s.processed.ProviderTransferFees, e)
	}

	return nil
}
