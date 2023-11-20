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
