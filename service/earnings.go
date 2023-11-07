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

type UpdateEarningParams struct {
	Referrals            int    `json:"referrals"`
	ReferralBalance      int    `json:"referral_balance"`
	ReferralTotalEarning int    `json:"referral_total_earning"`
	TotalWithdrawal      int    `json:"total_withdrawal"`
	MediaEarning         int    `json:"media_earning"`
	Owner                string `json:"owner"`
}

func (q *Queries) UpdateEarning(ctx context.Context, arg UpdateEarningParams) (types.Earning, error) {
	query := `UPDATE earnings SET referrals = $2, referral_balance = $3, referral_total_earning = $4,
						total_withdrawal = $5, media_earning = $6 where owner = $1 RETURNING id, referrals, 
						referral_balance, referral_total_earning, total_withdrawal, media_earning, owner`

	row := q.db.QueryRowContext(ctx, query, arg.Owner, arg.Referrals, arg.ReferralBalance, arg.ReferralTotalEarning, arg.TotalWithdrawal, arg.MediaEarning)
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
