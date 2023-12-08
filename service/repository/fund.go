package repository

import (
	"context"
	"fmt"
	service "merrypay/service/model"
	"merrypay/types"
	"strings"
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

// type updateWithdrawal struct {
// 	ID int
// 	amount int
// 	kind string
// 	withdrawBy string
// }

func (m *Model) CompleteWithdrawal(ctx context.Context, arg types.CompleteWithdrawalParams) error {
	arg.Kind = strings.ToLower(arg.Kind)
	withdrawal, err := m.Model.GetWithdrawalByID(ctx, arg.ID)
	if err != nil {
		return err
	}

	withdrawalEarning, err := m.Model.GetEarning(ctx, withdrawal.WithdrawBy)
	if err != nil {
		return err
	}
	// withdrawalEarning.TotalWithdrawal += amount

	transactArg := service.CreateTransaction{
		Kind:       arg.Kind,
		Amount:     -arg.Amount,
		TransactBy: withdrawal.WithdrawBy,
	}
	err = m.Model.CreateTransaction(ctx, transactArg)
	if err != nil {
		return err
	}
	bal, err := m.Model.GetBalanceFromTransaction(ctx, withdrawal.WithdrawBy, arg.Kind)
	if err != nil {
		return err
	}
	var earningArg types.UpdateEarningParams
	if arg.Kind == "referral" {
		earningArg = types.UpdateEarningParams{
			Referrals:               withdrawalEarning.Referrals,
			ReferralBalance:         bal.Balance,
			ReferralTotalWithdrawal: bal.TotalWithdrawal,
			ReferralTotalEarning:    bal.TotalEarning,
			Owner:                   withdrawalEarning.Owner,
			MediaBalance:            withdrawalEarning.MediaBalance,
			MediaTotalEarning:       withdrawalEarning.MediaTotalEarning,
			MediaTotalWithdrawal:    withdrawalEarning.MediaTotalWithdrawal,
		}
	} else {
		earningArg = types.UpdateEarningParams{
			Referrals:               withdrawalEarning.Referrals,
			ReferralBalance:         withdrawalEarning.ReferralBalance,
			ReferralTotalWithdrawal: withdrawalEarning.ReferralTotalWithdrawal,
			ReferralTotalEarning:    withdrawalEarning.ReferralTotalEarning,
			Owner:                   withdrawalEarning.Owner,
			MediaBalance:            bal.Balance,
			MediaTotalEarning:       bal.TotalEarning,
			MediaTotalWithdrawal:    bal.TotalWithdrawal,
		}
	}

	_, err = m.Model.UpdateEarning(ctx, earningArg)
	if err != nil {
		return err
	}

	_, err = m.Model.UpdateWithdrawal(ctx, arg.ID, "successful")
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) CancelWithdrawal(ctx context.Context, id int) error {
	withdrawal, err := m.Model.UpdateWithdrawal(ctx, id, "failed")
	if err != nil {
		return err
	}

	earning, err := m.Model.GetEarning(ctx, withdrawal.WithdrawBy)
	if err != nil {
		return err
	}

	earning.ReferralBalance += withdrawal.Amount

	arg := types.UpdateEarningParams{
		Referrals: earning.Referrals,
		ReferralBalance: earning.ReferralBalance,
		ReferralTotalEarning: earning.ReferralTotalEarning,
		ReferralTotalWithdrawal: earning.ReferralTotalWithdrawal,
		MediaBalance: earning.MediaBalance,
		MediaTotalEarning: earning.MediaTotalEarning,
		MediaTotalWithdrawal: earning.MediaTotalWithdrawal,
	}

	_, err = m.Model.UpdateEarning(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) GetWithdrawReceiptByID(ctx context.Context, id int) (types.Withdrawal, error) {
	withdrawal, err := m.Model.GetWithdrawalByID(ctx, id)
	if err != nil {
		return types.Withdrawal{}, err
	}

	return withdrawal, nil
}
