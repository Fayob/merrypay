package service

import (
	"context"
	"merrypay/types"
)

type InitiateWithdrawalParams struct {
	Amount int
	InitiatedBy string
	Kind string
}

func (q *Queries) InitiateWithdrawal(ctx context.Context, arg InitiateWithdrawalParams) (types.Withdrawal, error) {
	query := `INSERT INTO withdrawal(amount, withdraw_by, kind) VALUES($1, $2, $3) 
						RETURNING id, amount, withdraw_by, kind, status, initiated_at, completed_at`
	row := q.db.QueryRowContext(ctx, query, arg.Amount, arg.InitiatedBy, arg.Kind)
	var withdrawal types.Withdrawal
	err := row.Scan(
		&withdrawal.ID,
		&withdrawal.Amount,
		&withdrawal.WithdrawBy,
		&withdrawal.Kind,
		&withdrawal.Status,
		&withdrawal.InitiatedAt,
		&withdrawal.CompletedAt,
	)

	return withdrawal, err
}

func (q *Queries) UpdateWithdrawal(ctx context.Context, id int, status string) (types.Withdrawal, error) {
	query := `UPDATE withdrawal SET status = $2 where id = $1 
						RETURNING id, amount, withdraw_by, kind, status, initiated_at, completed_at`
	row := q.db.QueryRowContext(ctx, query, id, status)
	var withdrawal types.Withdrawal
	err := row.Scan(
		&withdrawal.ID,
		&withdrawal.Amount,
		&withdrawal.WithdrawBy,
		&withdrawal.Kind,
		&withdrawal.Status,
		&withdrawal.InitiatedAt,
		&withdrawal.CompletedAt,
	)

	return withdrawal, err
}

func (q *Queries) GetWithdrawalByID(ctx context.Context, id int) (types.Withdrawal, error) {
	query := `SELECT id, amount, withdraw_by, kind, status, initiated_at, completed_at
						FROM withdrawal where id = $1`

	row := q.db.QueryRowContext(ctx, query, id)
	var withdrawal types.Withdrawal
	err := row.Scan(
		&withdrawal.ID,
		&withdrawal.Amount,
		&withdrawal.WithdrawBy,
		&withdrawal.Kind,
		&withdrawal.Status,
		&withdrawal.InitiatedAt,
		&withdrawal.CompletedAt,
	)

	return withdrawal, err
}

func (q *Queries) GetUserWithdrawals(ctx context.Context, username string) ([]types.Withdrawal, error) {
	query := `SELECT id, amount, withdraw_by, kind, status, initiated_at, completed_at FROM withdrawal
						where withdraw_by = $1`

	rows, err := q.db.QueryContext(ctx, query, username)
	if err != nil {
		return nil, err
	}
	var withdrawals []types.Withdrawal
	for rows.Next() {
		var withdrawal types.Withdrawal
		err := rows.Scan(
			&withdrawal.ID,
			&withdrawal.Amount,
			&withdrawal.WithdrawBy,
			&withdrawal.Kind,
			&withdrawal.Status,
			&withdrawal.InitiatedAt,
			&withdrawal.CompletedAt,
		)
		if err != nil {
			return nil, err
		}

		withdrawals = append(withdrawals, withdrawal)
	}

	return withdrawals, nil
}
