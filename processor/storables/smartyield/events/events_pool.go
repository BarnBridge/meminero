package events

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) decodePoolEvent(log gethtypes.Log) error {
	sy := ethtypes.SmartYield

	if sy.IsBuyTokensEvent(&log) {
		e, err := sy.BuyTokensEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuyTokens event")
		}

		s.processed.JuniorEntryEvents = append(s.processed.JuniorEntryEvents, e)
	}

	if sy.IsSellTokensEvent(&log) {
		e, err := sy.SellTokensEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode SellTokens event")
		}

		s.processed.JuniorInstantWithdrawEvents = append(s.processed.JuniorInstantWithdrawEvents, e)
	}

	if sy.IsBuyJuniorBondEvent(&log) {
		e, err := sy.BuyJuniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuyJuniorBond event")
		}

		s.processed.Junior2StepWithdrawEvents = append(s.processed.Junior2StepWithdrawEvents, e)
	}

	if sy.IsRedeemJuniorBondEvent(&log) {
		e, err := sy.RedeemJuniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode RedeemJuniorBond event")
		}

		s.processed.Junior2StepRedeemEvents = append(s.processed.Junior2StepRedeemEvents, e)
	}

	if sy.IsBuySeniorBondEvent(&log) {
		e, err := sy.BuySeniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuySeniorBond event")
		}

		s.processed.SeniorEntryEvents = append(s.processed.SeniorEntryEvents, e)
	}

	if sy.IsRedeemSeniorBondEvent(&log) {
		e, err := sy.RedeemSeniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode RedeemSeniorBond event")
		}

		s.processed.SeniorRedeemEvents = append(s.processed.SeniorRedeemEvents, e)
	}

	if sy.IsTransferEvent(&log) {
		e, err := sy.TransferEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode Transfer event")
		}

		s.processed.Transfers = append(s.processed.Transfers, e)
	}

	return nil
}
