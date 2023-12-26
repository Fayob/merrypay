package port

import (
	"context"
	"merrypay/service/model"
	"merrypay/types"
)

type Store interface {
	CreateUser(context.Context, types.CreateUserParams) (types.User, error)
	FindUser(context.Context, string) (types.User, error)
	FindAllUsers(context.Context) ([]types.User, error)
	UpdateUser(context.Context, types.UpdateUserParams) (types.User, error)
	UpdatePassword(ctx context.Context, password, identifier string) (string, error)
	DeleteUser(context.Context, string) (string, error)
	ReferralHistory(ctx context.Context, username string) ([]types.RefHisResponse, error)
	SaveCoupon(context.Context, string, string) (string, error)
	RegisterWithCoupon(context.Context, string, string) error
	GetCoupon(context.Context, string) (types.Coupon, error)
	CreateEarning(ctx context.Context, owner string) (types.Earning, error)
	GetEarning(ctx context.Context, owner string) (types.Earning, error)
	UpdateEarning(ctx context.Context, arg types.UpdateEarningParams) (types.Earning, error)
	SaveBankDetails(ctx context.Context, arg types.BankDetailParams) error
	FetchBankDetail(ctx context.Context, owner string) (types.BankDetail, error)
	UpdateBankDetail(ctx context.Context, arg types.BankDetailParams) (types.BankDetail, error)
	InitiateWithdrawal(ctx context.Context, arg service.InitiateWithdrawalParams) (types.Withdrawal, error)
	UpdateWithdrawal(ctx context.Context, id int, status string) (types.Withdrawal, error)
	GetWithdrawalByID(ctx context.Context, id int) (types.Withdrawal, error)
	GetUserWithdrawals(ctx context.Context, username string) ([]types.Withdrawal, error)
	GetWithdrawalsByStatus(ctx context.Context, status string) ([]types.Withdrawal, error)
	CreateTransaction(ctx context.Context, arg service.CreateTransaction) error
	GetBalanceFromTransaction(ctx context.Context, username, kind string) (service.Balance, error)
}

// type UserService interface {
	// CreateUser(context.Context, service.CreateUserParams) (service.User, error)
	// FindUser(context.Context, string) (service.User, error)
	// FindAllUsers(context.Context) ([]service.User, error)
	// UpdateUser(context.Context, service.UpdateUserParams) (service.User, error)
	// UpdatePassword(context.Context, string, string) (string, error)
	// DeleteUser(context.Context, string) (string, error)
// }

// type CouponService interface {
	// SaveCoupon(context.Context, string, string) (string, error)
	// RegisterWithCoupon(context.Context, string, string) error
	// CouponUsedBy(context.Context, string) (interface{}, error)
// }

// type EarningService interface {
	// CreateEarning(context.Context, string) error
	// GetEarning(context.Context, string) (service.Earning, error)
	// UpdateEarning(context.Context, service.UpdateEarningParams) (service.Earning, error)
// }

// type BankDetailsService interface {
	// SaveBankDetails(context.Context, service.BankDetailParams) error
	// FetchBankDetail(context.Context, string) (service.BankDetail, error)
	// UpdateBankDetail(context.Context, service.BankDetailParams) (service.BankDetail, error)
// }

// type WithdrawalService interface{
	// InitiateWithdrawal(context.Context, int, string) (service.Withdrawal, error)
	// CompleteWithdrawal(context.Context, int) (service.Withdrawal, error)
// }