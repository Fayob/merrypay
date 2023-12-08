package repository

import (
	"context"
	"database/sql"
	"fmt"
	"merrypay/types"
	"strings"
)

func (m *Model) GetUser(ctx context.Context, username string) (types.User, error) {
	if username == "" {
		return types.User{}, fmt.Errorf("username cannot be empty")
	}

	username = strings.ToLower(username)

	user, err := m.Model.FindUser(ctx, username)
	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (m *Model) UpdateUser(ctx context.Context, arg types.UpdateUserParams) error {
	if arg.Username == "" || arg.Email == "" || arg.FirstName == "" || arg.LastName == "" {
		return fmt.Errorf("all field must be filled")
	}

	arg.Username = strings.ToLower(arg.Username)
	arg.Email = strings.ToLower(arg.Email)

	user, err := m.Model.FindUser(ctx, arg.Username)
	if err != nil {
		return err
	}

	if user.Email != arg.Email {
		_, err := m.Model.FindUser(ctx, arg.Email)
		if err != sql.ErrNoRows {
			fmt.Println(err)
			return fmt.Errorf("email already in use kindly choose another email address")
		}
	}

	_, err = m.Model.UpdateUser(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) UpdateMembership(ctx context.Context, arg types.MembershipUpdateParams) error {
	if arg.AccessorUsername == "" || arg.AccOwnerUsername == "" {
		return fmt.Errorf("all field must be filled")
	}

	arg.AccOwnerUsername = strings.ToLower(arg.AccOwnerUsername)
	arg.AccessorUsername = strings.ToLower(arg.AccessorUsername)

	accessorUser, err := m.Model.FindUser(ctx, arg.AccessorUsername)
	if err != nil {
		return err
	}
	if accessorUser.Membership != "admin" {
		return fmt.Errorf("unauthorized routes")
	}

	user, err := m.Model.FindUser(ctx, arg.AccOwnerUsername)
	if err != nil {
		return err
	}

	user.Membership = arg.Membership

	updatedArg := types.UpdateUserParams{
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Membership: user.Membership,
	}

	_, err = m.Model.UpdateUser(ctx, updatedArg)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) UpdateUserPassword(ctx context.Context, arg types.UpdatePasswordParams) error {
	arg.Username = strings.ToLower(arg.Username)

	user, err := m.Model.FindUser(ctx, arg.Username)
	if err != nil {
		return err
	}

	if !checkHashPassword(user.Password, arg.OldPassword) {
		return fmt.Errorf("wrong password provided")
	}

	hashedPassword, err := hashPassword(arg.NewPassword)
	if err != nil {
		return err
	}

	_, err = m.Model.UpdatePassword(ctx, hashedPassword, arg.Username)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) DeleteUser(ctx context.Context, username string) error {
	if username == "" {
		return fmt.Errorf("please fill all required fields")
	}

	username = strings.ToLower(username)

	_, err := m.Model.DeleteUser(ctx, username)

	return err
}

func (m *Model) UserReferred(ctx context.Context, username string) ([]types.RefHisResponse, error) {
	if username == "" {
		return nil, fmt.Errorf("please fill all required fields")
	}

	username = strings.ToLower(username)

	userReferred, err := m.Model.ReferralHistory(ctx, username)
	if err != nil {
		return nil, err
	}

	return userReferred, nil
}