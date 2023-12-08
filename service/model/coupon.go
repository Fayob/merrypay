package service

import (
	"context"
	"merrypay/types"
)

func (q *Queries) SaveCoupon(ctx context.Context, coupon, username string) (string, error) {
	query := `INSERT INTO coupon(digit, created_by) VALUES($1, $2)`

	_, err := q.db.ExecContext(ctx, query, coupon, username)
	if err != nil {
		return "", err
	}

	return "Saved", nil
}

func (q *Queries) RegisterWithCoupon(ctx context.Context, coupon, username string) error {
	query := `UPDATE coupon SET used_by = $1 where digit = $2`
	_, err := q.db.ExecContext(ctx, query, username, coupon)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) GetCoupon(ctx context.Context, coupon string) (types.Coupon, error) {
	query := `SELECT digit, used_by, created_at FROM coupon where digit = $1`
	row := q.db.QueryRowContext(ctx, query, coupon)
	var c types.Coupon
	if err := row.Scan(
		&c.Digit,
		&c.UsedBy,
		&c.CreatedAt,
	); err != nil {
		return types.Coupon{}, err
	}

	return c, nil
}
