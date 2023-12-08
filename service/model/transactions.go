package service

import (
	"context"
	"merrypay/types"
)

type CreateTransaction struct {
	Kind       string
	Amount     int
	TransactBy string
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransaction) error {
	query := `INSERT INTO transactions(kind, amount, transact_by) VALUES($1, $2, $3)`

	_, err := q.db.ExecContext(ctx, query, arg.Kind, arg.Amount, arg.TransactBy)
	return err
}

// func (q *Queries) CreateDebitTransaction(ctx context.Context, arg CreateTransaction) error {
// 	query := `INSERT INTO transactions(kind, debit, transact_by) VALUES($1, $2, $3)`

// 	_, err := q.db.ExecContext(ctx, query, arg.Kind, -arg.Amount, arg.TransactBy)
// 	return err
// }

func (q *Queries) GetTransactionByID(ctx context.Context, id int) (types.Transaction, error) {
	query := `SELECT id, amount, kind, transact_by, created_at FROM transactions WHERE id=$1`
	row := q.db.QueryRowContext(ctx, query, id)
	var transaction types.Transaction
	err := row.Scan(
		&transaction.ID,
		&transaction.Amount,
		&transaction.Kind,
		&transaction.TransactBy,
		&transaction.CreatedAt,
	)

	return transaction, err
}

func (q *Queries) GetTransactionsByUsername(ctx context.Context, username string) ([]types.Transaction, error) {
	query := `SELECT id, amount, kind, transact_by, created_at FROM transactions WHERE id=$1`
	rows, err := q.db.QueryContext(ctx, query, username)
	if err != nil {
		return nil, err
	}
	var transactions []types.Transaction
	for rows.Next() {
		var transaction types.Transaction
		if err = rows.Scan(
			&transaction.ID,
			&transaction.Amount,
			&transaction.Kind,
			&transaction.TransactBy,
			&transaction.CreatedAt,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, err
}

type Balance struct {
	Balance         int
	TotalEarning    int
	TotalWithdrawal int
}

func (q *Queries) GetBalanceFromTransaction(ctx context.Context, username, kind string) (Balance, error) {
	var balance Balance
	totalEarningQuery := `SELECT COALESCE(SUM(amount), 0) FROM transactions where transact_by=$1 AND kind=$2 AND amount > 0`
	totalEarning := q.db.QueryRowContext(ctx, totalEarningQuery, username, kind)
	if err := totalEarning.Scan(
		&balance.TotalEarning,
	); err != nil {
		return Balance{}, err
	}

	totalWithdrawalQuery := `SELECT COALESCE(SUM(amount), 0) FROM transactions where transact_by=$1 AND kind=$2 AND amount < 0`
	totalWithdrawal := q.db.QueryRowContext(ctx, totalWithdrawalQuery, username, kind)
	if err := totalWithdrawal.Scan(
		&balance.TotalWithdrawal,
	); err != nil {
		return Balance{}, err
	}

	accBalanceQuery := `SELECT COALESCE(SUM(amount), 0) FROM transactions where transact_by=$1 AND kind=$2`
	accBalance := q.db.QueryRowContext(ctx, accBalanceQuery, username, kind)
	if err := accBalance.Scan(
		&balance.Balance,
	); err != nil {
		return Balance{}, err
	}

	return balance, nil
}
