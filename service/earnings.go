package service

import (
	"context"
	"merrypay/types"
)

func (q *Queries) CreateEarning(ctx context.Context, owner string) error {
	query := `INSERT INTO earnings(owner) VALUES($1)`
	_, err := q.db.ExecContext(ctx, query, owner)
	if err != nil {
		return err
	}
	return nil
}

func (q *Queries) GetEarning(ctx context.Context, owner string) (types.Earning, error) {
	query := `SELECT id, referrals, referral_balance, referral_total_earning, total_withdrawal, media_earning,
						owner FROM earnings WHERE owner = $1`

	row := q.db.QueryRowContext(ctx, query, owner)
	var earning types.Earning
	err := row.Scan(
		&earning.ID,
		&earning.Referrals,
		&earning.ReferralBalance,
		&earning.ReferralTotalEarning,
		&earning.TotalWithdrawal,
		&earning.MediaEarning,
		&earning.Owner,
	)

	return earning, err
}
