package service

import "context"

func (q *Queries) SaveCoupon(ctx context.Context, coupon, username string) (string, error) {
	query := `INSERT INTO coupon(digit, created_by) VALUES($1, $2)`

	_, err := q.db.ExecContext(ctx, query, coupon, username)
	if err != nil {
		return "", err
	}

	return "Saved", nil
}
