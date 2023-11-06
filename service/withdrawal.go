package service

import (
	"context"
	"time"
)

type Withdrawal struct {
	ID          int       `json:"id"`
	Amount      int       `json:"amount"`
	WithdrawBy  string    `json:"withdraw_by"`
	Status      string       `json:"status"`
	InitiatedAt time.Time `json:"initiated_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func (q *Queries) InitiateWithdrawal(ctx context.Context, amount int, initiated_by string) (Withdrawal, error) {
	query := `INSERT INTO withdrawal(amount, withdraw_by) VALUES($1, $2) 
						RETURNING id, amount, withdraw_by, status, initiated_at, completed_at`
	row := q.db.QueryRowContext(ctx, query, amount, initiated_by)
	var withdrawal Withdrawal
	err := row.Scan(
		&withdrawal.ID,
		&withdrawal.Amount,
		&withdrawal.WithdrawBy,
		&withdrawal.Status,
		&withdrawal.InitiatedAt,
		&withdrawal.CompletedAt,
	)

	return withdrawal, err
}
