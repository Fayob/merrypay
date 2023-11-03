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

func (q *Queries) RegisterWithCoupon(ctx context.Context, coupon, username string) error {
	query := `UPDATE coupon SET used_by = $1 where digit = $2`
	_, err := q.db.ExecContext(ctx, query, username, coupon)
	if err != nil {
		return err
	}

	return nil
}

type Coupon struct {
	UsedBy interface{} `json:"used_by"`
}

func (q *Queries) CouponUsedBy(ctx context.Context, coupon string) (interface{}, error) {
	query := `SELECT used_by FROM coupon where digit = $1`
	row := q.db.QueryRowContext(ctx, query, coupon)
	var c Coupon
	if err := row.Scan(
		&c.UsedBy,
	); err != nil {
		return nil, err
	}

	return c.UsedBy, nil
}