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
	query := `SELECT id, referrals, referral_balance, referral_total_earning, referral_total_withdrawal, 
				media_balance, media_total_earning, media_total_withdrawal, owner FROM earnings WHERE owner = $1`

	row := q.db.QueryRowContext(ctx, query, owner)
	var earning types.Earning
	err := row.Scan(
		&earning.ID,
		&earning.Referrals,
		&earning.ReferralBalance,
		&earning.ReferralTotalEarning,
		&earning.ReferralTotalWithdrawal,
		&earning.MediaBalance,
		&earning.MediaTotalEarning,
		&earning.MediaTotalWithdrawal,
		&earning.Owner,
	)

	return earning, err
}

func (q *Queries) UpdateEarning(ctx context.Context, arg types.UpdateEarningParams) (types.Earning, error) {
	query := `UPDATE earnings SET referrals = $2, referral_balance = $3, referral_total_earning = $4,
						referral_total_withdrawal = $5, media_balance = $6, media_total_earning = $7,
						media_total_withdrawal = $8 where owner = $1 RETURNING id, referrals, referral_balance, 
						referral_total_earning, referral_total_withdrawal, media_balance, media_total_earning, 
						media_total_withdrawal, owner`

	row := q.db.QueryRowContext(
		ctx, query, arg.Owner, arg.Referrals, arg.ReferralBalance, 
		arg.ReferralTotalEarning, arg.ReferralTotalWithdrawal, arg.MediaBalance,
		arg.MediaTotalEarning, arg.MediaTotalWithdrawal,
	)
	var earning types.Earning
	err := row.Scan(
		&earning.ID,
		&earning.Referrals,
		&earning.ReferralBalance,
		&earning.ReferralTotalEarning,
		&earning.ReferralTotalWithdrawal,
		&earning.MediaBalance,
		&earning.MediaTotalEarning,
		&earning.MediaTotalWithdrawal,
		&earning.Owner,
	)

	return earning, err
}
