package service

import (
	"context"
	"merrypay/types"
)

type BankDetailParams struct {
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Owner         string `json:"owner"`
}

func (q *Queries) SaveBankDetails(ctx context.Context, arg BankDetailParams) error {
	query := `INSERT INTO bank_details(bank_name, account_name, account_number, owner)
						VALUES($1, $2, $3, $4)`

	_, err := q.db.ExecContext(ctx, query, arg.BankName, arg.AccountName, arg.AccountNumber, arg.Owner)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) FetchBankDetail(ctx context.Context, owner string) (types.BankDetail, error) {
	query := `SELECT id, bank_name, account_name, account_number, owner FROM bank_details where owner = $1`
	row := q.db.QueryRowContext(ctx, query, owner)
	var bankDetail types.BankDetail
	err := row.Scan(
		&bankDetail.ID,
		&bankDetail.BankName,
		&bankDetail.AccountName,
		&bankDetail.AccountNumber,
		&bankDetail.Owner,
	)

	return bankDetail, err
}
