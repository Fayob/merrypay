package port

import (
	"context"
	"merrypay/types"
)

type Server interface {
	RegisterUser(ctx context.Context, arg types.CreateUserParams) (types.User, error)
	LogInUser(ctx context.Context, identifier, password string) (types.User, error)
	GetUser(ctx context.Context, username string) (types.User, error)
	UpdateUser(ctx context.Context, arg types.UpdateUserParams) error
	UpdateMembership(ctx context.Context, arg types.MembershipUpdateParams) error
	UpdateUserPassword(ctx context.Context, arg types.UpdatePasswordParams) error
	UserReferred(ctx context.Context, username string) ([]types.RefHisResponse, error)
	DeleteUser(ctx context.Context, username string) error
	GenerateCoupon(ctx context.Context, username string) (string, error)
	WithdrawFund(ctx context.Context, arg types.WithdrawalParam) (types.Withdrawal, error)
	CompleteWithdrawal(ctx context.Context, id int) error
	CancelWithdrawal(ctx context.Context, id int) error
	GetWithdrawReceiptByID(ctx context.Context, id int) (types.Withdrawal, error)
	GetWithdrawalByStatus(ctx context.Context, status string) ([]types.Withdrawal, error)
	AddBankDetails(ctx context.Context, arg types.BankDetailParams) error 
	GetBankDetails(ctx context.Context, username string) (types.BankDetail, error)
	UpdateBankDetail(ctx context.Context, arg types.BankDetailParams) (types.BankDetail, error)
	GetEarning(ctx context.Context, username string) (types.Earning, error)
	Jackpot(ctx context.Context, guess [5]int, username string) ([5]int, error)
}