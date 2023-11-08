package port

import (
	"context"
	"merrypay/service"
	"merrypay/types"
)

type Store interface {
	CreateUser(context.Context, types.CreateUserParams) (types.User, error)
	FindUser(context.Context, string) (types.User, error)
	FindAllUsers(context.Context) ([]types.User, error)
	UpdateUser(context.Context, service.UpdateUserParams) (types.User, error)
	UpdatePassword(context.Context, string, string) (string, error)
	DeleteUser(context.Context, string) (string, error)
	SaveCoupon(context.Context, string, string) (string, error)
	RegisterWithCoupon(context.Context, string, string) error
	GetCoupon(context.Context, string) (types.Coupon, error)
	CreateEarning(ctx context.Context, owner string) error
	GetEarning(ctx context.Context, owner string) (types.Earning, error)
	UpdateEarning(ctx context.Context, arg service.UpdateEarningParams) (types.Earning, error)
	SaveBankDetails(context.Context, service.BankDetailParams) error
	FetchBankDetail(context.Context, string) (types.BankDetail, error)
	UpdateBankDetail(context.Context, service.BankDetailParams) (types.BankDetail, error)
	InitiateWithdrawal(context.Context, int, string) (types.Withdrawal, error)
	CompleteWithdrawal(context.Context, int) (types.Withdrawal, error)
	GetWithdrawalByID(ctx context.Context, id int) (types.Withdrawal, error)
	GetUserWithdrawal(ctx context.Context, username string) ([]types.Withdrawal, error)
}

type Handler interface {

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