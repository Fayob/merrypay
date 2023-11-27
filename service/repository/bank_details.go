package repository

import (
	"context"
	"fmt"
	"merrypay/types"
)

func (m *Model) AddBankDetails(ctx context.Context, arg types.BankDetailParams) error {
	if arg.AccountName == "" || arg.AccountNumber == "" || arg.BankName == "" || arg.Owner == "" {
		return fmt.Errorf("please fill all required fields")
	}

	return m.Model.SaveBankDetails(ctx, arg)

	// return err
}

func (m *Model) GetBankDetails(ctx context.Context, username string) (types.BankDetail, error) {
	if username == "" {
		return types.BankDetail{}, fmt.Errorf("please supply username")
	}

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
	
	updatedBankDetails, err := m.Model.UpdateBankDetail(ctx, arg)
	if err != nil {
		return types.BankDetail{}, err
	}

	return updatedBankDetails, nil
}
