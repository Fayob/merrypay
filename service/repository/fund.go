package repository

import (
	"context"
	"fmt"
	service "merrypay/service/model"
	"merrypay/types"
)

func (m *Model) WithdrawFund(ctx context.Context, arg types.WithdrawalParam) (types.Withdrawal, error) {
	withdrawalEarnings, err := m.Model.GetEarning(ctx, arg.WithdrawBy)
	if err != nil {
		return types.Withdrawal{}, err
	}
	if withdrawalEarnings.ReferralBalance < arg.Amount {
		return types.Withdrawal{}, fmt.Errorf("insufficient fund")
	}

	withdrawalEarnings.ReferralBalance -= arg.Amount
	updateArg := types.UpdateEarningParams{
		Referrals:               withdrawalEarnings.Referrals,
		ReferralBalance:         withdrawalEarnings.ReferralBalance,
		ReferralTotalEarning:    withdrawalEarnings.ReferralTotalEarning,
		ReferralTotalWithdrawal: withdrawalEarnings.ReferralTotalWithdrawal,
		MediaBalance:            withdrawalEarnings.MediaBalance,
		MediaTotalEarning:       withdrawalEarnings.MediaTotalEarning,
		MediaTotalWithdrawal:    withdrawalEarnings.MediaTotalWithdrawal,
		Owner:                   withdrawalEarnings.Owner,
	}
	_, err = m.Model.UpdateEarning(ctx, updateArg)
	if err != nil {
		return types.Withdrawal{}, err
	}

	initWithdrawal := service.InitiateWithdrawalParams{
		Amount:      arg.Amount,
		Kind:        arg.Kind,
		InitiatedBy: arg.WithdrawBy,
	}

	withdrawal, err := m.Model.InitiateWithdrawal(ctx, initWithdrawal)
	if err != nil {
		return types.Withdrawal{}, err
	}

	return withdrawal, nil
}
