package events

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) decodePoolEvent(log gethtypes.Log) error {
	sy := ethtypes.SmartYield

	if sy.IsSmartYieldBuyTokensEvent(&log) {
		e, err := sy.SmartYieldBuyTokensEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuyTokens event")
		}

		s.processed.JuniorEntryEvents = append(s.processed.JuniorEntryEvents, e)
	}

	if sy.IsSmartYieldSellTokensEvent(&log) {
		e, err := sy.SmartYieldSellTokensEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode SellTokens event")
		}

		s.processed.JuniorInstantWithdrawEvents = append(s.processed.JuniorInstantWithdrawEvents, e)
	}

	if sy.IsSmartYieldBuyJuniorBondEvent(&log) {
		e, err := sy.SmartYieldBuyJuniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuyJuniorBond event")
		}

		s.processed.Junior2StepWithdrawEvents = append(s.processed.Junior2StepWithdrawEvents, e)
	}

	if sy.IsSmartYieldRedeemJuniorBondEvent(&log) {
		e, err := sy.SmartYieldRedeemJuniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode RedeemJuniorBond event")
		}

		s.processed.Junior2StepRedeemEvents = append(s.processed.Junior2StepRedeemEvents, e)
	}

	if sy.IsSmartYieldBuySeniorBondEvent(&log) {
		e, err := sy.SmartYieldBuySeniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode BuySeniorBond event")
		}

		s.processed.SeniorEntryEvents = append(s.processed.SeniorEntryEvents, e)
	}

	if sy.IsSmartYieldRedeemSeniorBondEvent(&log) {
		e, err := sy.SmartYieldRedeemSeniorBondEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode RedeemSeniorBond event")
		}

		s.processed.SeniorRedeemEvents = append(s.processed.SeniorRedeemEvents, e)
	}

	if sy.IsSmartYieldTransferEvent(&log) {
		e, err := sy.SmartYieldTransferEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode Transfer event")
		}

		s.processed.Transfers = append(s.processed.Transfers, e)
	}

	return nil
}
