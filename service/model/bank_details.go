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

func (q *Queries) UpdateBankDetail(ctx context.Context, arg BankDetailParams) (types.BankDetail, error) {
	query := `UPDATE bank_details SET bank_name = $1, account_name = $2, account_number = $3 where owner = $4
						RETURNING id, bank_name, account_number, account_name, owner, created_at`
	row := q.db.QueryRowContext(ctx, query, arg.BankName, arg.AccountName, arg.AccountNumber, arg.Owner)
	var bankDetail types.BankDetail
	err := row.Scan(
		&bankDetail.ID,
		&bankDetail.BankName,
		&bankDetail.AccountNumber,
		&bankDetail.AccountName,
		&bankDetail.Owner,
		&bankDetail.CreatedAt,
	)
	return bankDetail, err
}

func (q *Queries) UpdateBankDetailNoReturn(ctx context.Context, arg BankDetailParams) error {
	query := `UPDATE bank_details SET bank_name = $1, account_name = $2, account_number = $3 where owner = $4`
	_, err := q.db.ExecContext(ctx, query, arg.BankName, arg.AccountName, arg.AccountNumber, arg.Owner)
	if err != nil {
		return err
	}
	return nil
}
