package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	service "merrypay/service/model"
	"merrypay/types"
	"time"
)

func (s *Model) RegisterUser(ctx context.Context, arg types.CreateUserParams) (types.User, error) {
	if arg.Username == "" || arg.FirstName == "" || arg.LastName == "" || arg.Email == "" || arg.Password == "" || arg.Coupon == "" {
		return types.User{}, fmt.Errorf("please fill all required fields")
	}

	if arg.Referral == "" {
		return types.User{}, fmt.Errorf("no referral found, please kindly contact admin to get a referral")
	}

	arg.Username = strings.ToLower(arg.Username)
	arg.Email = strings.ToLower(arg.Email)

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

	referralEarning, err := s.Model.GetEarning(ctx, arg.Referral)
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
	
	user, err := s.Model.CreateUser(ctx, arg)
	if err != nil {
		return types.User{}, err
	}

	err = s.Model.RegisterWithCoupon(ctx, arg.Coupon, arg.Username)
	if err != nil {
		return types.User{}, err
	}

	userEarning, err := s.Model.CreateEarning(ctx, arg.Username)
	if err != nil {
		return types.User{}, err
	}
	
	// create Transaction

	if err := s.updateUserAccount(ctx, "referral", referralEarning); err != nil {
		return types.User{}, err
	}

	if err := s.updateUserAccount(ctx, "media", userEarning); err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (s *Model) LogInUser(ctx context.Context, identifier, password string) (types.User, error) {
	if identifier == "" || password == "" {
		return types.User{}, fmt.Errorf("please fill all required fields")
	}

	identifier = strings.ToLower(identifier)

	user, err := s.Model.FindUser(ctx, identifier)
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

func (s *Model) checkUserByUsername(ctx context.Context, username string) error {
	user, err := s.Model.FindUser(ctx, username)
	if err != sql.ErrNoRows || user.Username == username {
		fmt.Println(err)
		return fmt.Errorf("%v already in use", username)
	}

	return nil
}

func (s *Model) checkUserByEmail(ctx context.Context, email string) error {
	user, err := s.Model.FindUser(ctx, email)
	if err != sql.ErrNoRows || user.Email == email {
		return fmt.Errorf("%v already in use", email)
	}

	return nil
}


func (m *Model) validateCoupon(ctx context.Context, coupon string) error {
	c, err := m.Model.GetCoupon(ctx, coupon)
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

func (m *Model) updateUserAccount(ctx context.Context, kind string, arg types.Earning) error {
	referrals := arg.Referrals + 1
	// referralBalance := arg.ReferralBalance + 6
	// referralTotalBal := arg.ReferralTotalEarning + 6

	// Create Transaction
	transactArg := service.CreateTransaction{
		Kind: kind,
		Amount: 6,
		TransactBy: arg.Owner,
	}
	err := m.Model.CreateTransaction(ctx, transactArg)
	if err != nil {
		return err
	}

	bal, err := m.Model.GetBalanceFromTransaction(ctx, arg.Owner, transactArg.Kind)
	if err != nil {
		return err
	}

	var updatedArg types.UpdateEarningParams
	if kind == "" {
		return fmt.Errorf("kind parameter should not be empty")
	} else if kind == "referral" {
		updatedArg = types.UpdateEarningParams{
			Referrals: referrals,
			ReferralBalance: bal.Balance,
			ReferralTotalEarning: bal.TotalEarning,
			ReferralTotalWithdrawal: bal.TotalWithdrawal,
			MediaBalance: arg.MediaBalance,
			MediaTotalEarning: arg.MediaTotalEarning,
			MediaTotalWithdrawal: arg.MediaTotalWithdrawal,
			Owner: arg.Owner,
		}
		} else {
		updatedArg = types.UpdateEarningParams{
			Referrals: arg.Referrals,
			ReferralBalance: arg.ReferralBalance,
			ReferralTotalEarning: arg.ReferralTotalEarning,
			ReferralTotalWithdrawal: arg.ReferralTotalWithdrawal,
			MediaBalance: bal.Balance,
			MediaTotalEarning: bal.TotalEarning,
			MediaTotalWithdrawal: bal.TotalWithdrawal,
			Owner: arg.Owner,
		}
	}

	_, err = m.Model.UpdateEarning(ctx, updatedArg)
	if err != nil {
		return err
	}
	
	return nil
}
