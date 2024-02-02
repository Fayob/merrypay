package types

import "time"

type User struct {
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Membership        string    `json:"membership"`
	WonJackpot        bool      `json:"won_jackpot"`
	ReferredBy        string    `json:"referred_by"`
	Password          string    `json:"password"`
	UpdatedPasswordAt time.Time `json:"updated_password_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type RefHisResponse struct {
	Username  string
	Email     string
	FirstName string
	LastName  string
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
	ID                      int    `json:"id"`
	Referrals               int    `json:"referrals"`
	ReferralBalance         int    `json:"referral_balance"`
	ReferralTotalEarning    int    `json:"referral_total_earning"`
	ReferralTotalWithdrawal int    `json:"referral_total_withdrawal"`
	MediaBalance            int    `json:"media_balance"`
	MediaTotalEarning       int    `json:"media_total_earning"`
	MediaTotalWithdrawal    int    `json:"media_total_withdrawal"`
	Owner                   string `json:"owner"`
}

type Withdrawal struct {
	ID          int       `json:"id"`
	Amount      int       `json:"amount"`
	WithdrawBy  string    `json:"withdraw_by"`
	Kind        string    `json:"kind"`
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

type UserResponse struct {
	Username   string
	Email      string
	FirstName  string
	LastName   string
	Membership string
	ReferredBy string
	WonJackpot bool
	CreatedAt  time.Time
}

type CreateUserParams struct {
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	Coupon    string `json:"coupon" binding:"required"`
	Referral  string `json:"referral" binding:"required"`
}

type UpdateUserParams struct {
	Username   string `json:"username" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Membership string `json:"membership" binding:"required"`
	WonJackpot bool   `json:"won_jackpot" binding:"required"`
}

type MembershipUpdateParams struct {
	AccessorUsername string `json:"accessor_username"`
	AccOwnerUsername string `json:"acc_owner_username" binding:"required"`
	Membership       string `json:"membership" binding:"required"`
}

type UpdatePasswordParams struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type UpdateEarningParams struct {
	Referrals               int    `json:"referrals" binding:"required"`
	ReferralBalance         int    `json:"referral_balance" binding:"required"`
	ReferralTotalEarning    int    `json:"referral_total_earning" binding:"required"`
	ReferralTotalWithdrawal int    `json:"referral_total_withdrawal" binding:"required"`
	MediaBalance            int    `json:"media_balance" binding:"required"`
	MediaTotalEarning       int    `json:"media_total_earning" binding:"required"`
	MediaTotalWithdrawal    int    `json:"media_total_withdrawal" binding:"required"`
	Owner                   string `json:"owner" binding:"required"`
}

type WithdrawalParam struct {
	Kind       string `json:"kind" binding:"required"`
	Amount     int    `json:"amount" binding:"required"`
	WithdrawBy string `json:"withdraw_by"`
}

type CompleteWithdrawalParams struct {
	ID     int    `json:"id" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
	Kind   string `json:"kind" binding:"required"`
}

type BankDetailParams struct {
	BankName      string `json:"bank_name" binding:"required"`
	AccountName   string `json:"account_name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	Owner         string `json:"owner"`
}

type JackpotParam struct {
	Guess    [5]int `json:"guess" binding:"required"`
	Username string	`json:"username"`
}
