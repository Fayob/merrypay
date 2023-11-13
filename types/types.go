package types

import "time"

type User struct {
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Membership        string    `json:"membership"`
	Password          string    `json:"password"`
	UpdatedPasswordAt time.Time `json:"updated_password_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type Coupon struct {
	Digit     string      `json:"digit"`
	UsedBy    interface{} `json:"used_by"`
	CreatedAt time.Time   `json:"created_at"`
}

type BankDetail struct {
	ID            int    `json:"id"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Owner         string `json:"owner"`
	CreatedAt     string `json:"created_at"`
}

type Earning struct {
	ID                   int    `json:"id"`
	Referrals            int    `json:"referrals"`
	ReferralBalance      int    `json:"referral_balance"`
	ReferralTotalEarning int    `json:"referral_total_earning"`
	TotalWithdrawal      int    `json:"total_withdrawal"`
	MediaEarning         int    `json:"media_earning"`
	Owner                string `json:"owner"`
}

type Withdrawal struct {
	ID          int       `json:"id"`
	Amount      int       `json:"amount"`
	WithdrawBy  string    `json:"withdraw_by"`
	Status      string    `json:"status"`
	InitiatedAt time.Time `json:"initiated_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type Transaction struct {
	ID         int       `json:"id"`
	Amount     int       `json:"amount"`
	Kind       string    `json:"kind"`
	TransactBy string    `json:"transact_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateUserParams struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Coupon    string `json:"coupon"`
	Referral  string `json:"referral"`
}

type UpdateEarningParams struct {
	Referrals            int    `json:"referrals"`
	ReferralBalance      int    `json:"referral_balance"`
	ReferralTotalEarning int    `json:"referral_total_earning"`
	TotalWithdrawal      int    `json:"total_withdrawal"`
	MediaEarning         int    `json:"media_earning"`
	Owner                string `json:"owner"`
}
