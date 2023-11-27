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
