package repository

import (
	"context"
	"database/sql"
	"fmt"

	service "merrypay/service/model"
	"merrypay/types"
	"time"
)

func (s *Server) RegisterUser(ctx context.Context, arg types.CreateUserParams) (types.User, error) {
	if arg.Username == "" || arg.FirstName == "" || arg.LastName == "" || arg.Email == "" || arg.Password == "" || arg.Coupon == "" {
		return types.User{}, fmt.Errorf("please fill all required fields")
	}

	if arg.Referral == "" {
		return types.User{}, fmt.Errorf("no referral found, please kindly contact admin to get a referral")
	}

	if err := s.checkUserByUsername(ctx, arg.Username); err != nil {
		return types.User{}, err
	}
	if err := s.checkUserByEmail(ctx, arg.Email); err != nil {
		return types.User{}, err
	}
	if !checkLength(arg.FirstName) {
		return types.User{}, fmt.Errorf("first name must be at least two characters")
	}
	if !checkLength(arg.LastName) {
		return types.User{}, fmt.Errorf("last name must be at least two characters")
	}

	if arg.Username == arg.Referral {
		return types.User{}, fmt.Errorf("sorry! You can't refer yourself, kindly contact the admin for a referral")
	}

	err := s.validateCoupon(ctx, arg.Coupon)
	if err != nil {
		return types.User{}, err
	}

	referralEarning, err := s.Server.GetEarning(ctx, arg.Referral)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, fmt.Errorf("unknown referral")
		}
		return types.User{}, err
	}

	hashedPassword, err := hashPassword(arg.Password)
	if err != nil {
		return types.User{}, err
	}

	arg.Password = hashedPassword
	
	user, err := s.Server.CreateUser(ctx, arg)
	if err != nil {
		return types.User{}, err
	}

	err = s.Server.RegisterWithCoupon(ctx, arg.Coupon, arg.Username)
	if err != nil {
		return types.User{}, err
	}

	err = s.Server.CreateEarning(ctx, arg.Username)
	if err != nil {
		return types.User{}, err
	}
	
	// create Transaction

	if err := s.updateReferralAccount(ctx, referralEarning); err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (s *Server) LogInUser(ctx context.Context, identifier, password string) (types.User, error) {
	if identifier == "" || password == "" {
		return types.User{}, fmt.Errorf("please fill all required fields")
	}

	user, err := s.Server.FindUser(ctx, identifier)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, fmt.Errorf("%v is not a registered user", identifier)
		}
		return types.User{}, fmt.Errorf("invalid credentials")
	}

	// Hash Password Functionality
	if !checkHashPassword(user.Password, password) {
		return types.User{}, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

func (s *Server) checkUserByUsername(ctx context.Context, username string) error {
	user, err := s.Server.FindUser(ctx, username)
	if err != sql.ErrNoRows || user.Username == username {
		return fmt.Errorf("%v already in use", username)
	}

	return nil
}

func (s *Server) checkUserByEmail(ctx context.Context, email string) error {
	user, err := s.Server.FindUser(ctx, email)
	if err != sql.ErrNoRows || user.Email == email {
		return fmt.Errorf("%v already in use", email)
	}

	return nil
}


func (s *Server) validateCoupon(ctx context.Context, coupon string) error {
	c, err := s.Server.GetCoupon(ctx, coupon)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid coupon")
		}
		return err
	}
	if c.UsedBy != nil {
		return fmt.Errorf("coupon has been used by another user")
	}

	if time.Since(c.CreatedAt) > (time.Minute * 30) {
		return fmt.Errorf("coupon has expired")
	}

	return nil
}

func (s *Server) updateReferralAccount(ctx context.Context, arg types.Earning) error {
	referrals := arg.Referrals + 1
	// referralBalance := arg.ReferralBalance + 6
	// referralTotalBal := arg.ReferralTotalEarning + 6

	// Create Transaction
	transactArg := service.CreateTransaction{
		Kind: "referral",
		Amount: 6,
		TransactBy: arg.Owner,
	}
	err := s.Server.CreateTransaction(ctx, transactArg)
	if err != nil {
		return err
	}

	bal, err := s.Server.GetBalanceFromTransaction(ctx, arg.Owner, transactArg.Kind)
	if err != nil {
		return err
	}

	updatedArg := types.UpdateEarningParams{
		Referrals: referrals,
		ReferralBalance: bal.Balance,
		ReferralTotalEarning: bal.TotalEarning,
		ReferralTotalWithdrawal: bal.TotalWithdrawal,
		MediaBalance: arg.MediaBalance,
		MediaTotalEarning: arg.MediaTotalEarning,
		MediaTotalWithdrawal: arg.MediaTotalWithdrawal,
		Owner: arg.Owner,
	}

	_, err = s.Server.UpdateEarning(ctx, updatedArg)
	if err != nil {
		return err
	}
	
	return nil
}
