package repository

import (
	"context"
	"fmt"
	"merrypay/types"
	"strings"
)

func (m *Model) AddBankDetails(ctx context.Context, arg types.BankDetailParams) error {
	if arg.AccountName == "" || arg.AccountNumber == "" || arg.BankName == "" || arg.Owner == "" {
		return fmt.Errorf("please fill all required fields")
	}

	arg.Owner = strings.ToLower(arg.Owner)

	return m.Model.SaveBankDetails(ctx, arg)
}

func (m *Model) GetBankDetails(ctx context.Context, username string) (types.BankDetail, error) {
	if username == "" {
		return types.BankDetail{}, fmt.Errorf("please supply username")
	}

	username = strings.ToLower(username)

	bankDetail, err := m.Model.FetchBankDetail(ctx, username)
	if err != nil {
		return types.BankDetail{}, err
	}

	return bankDetail, nil
}

func (m *Model) UpdateBankDetail(ctx context.Context, arg types.BankDetailParams) (types.BankDetail, error) {
	if arg.AccountName == "" || arg.AccountNumber == "" || arg.BankName == "" || arg.Owner == "" {
		return types.BankDetail{}, fmt.Errorf("please fill all required fields")
	}

	arg.Owner = strings.ToLower(arg.Owner)
	
	updatedBankDetails, err := m.Model.UpdateBankDetail(ctx, arg)
	if err != nil {
		return types.BankDetail{}, err
	}

	return updatedBankDetails, nil
}
