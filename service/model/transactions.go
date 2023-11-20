package service

import (
	"context"
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

