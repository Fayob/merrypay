package repository

import (
	"context"
	"database/sql"
	"fmt"
	"merrypay/types"
)

func (m *Model) GetUser(ctx context.Context, username string) (types.User, error) {
	if username == "" {
		return types.User{}, fmt.Errorf("username cannot be empty")
	}

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

	user, err := m.Model.FindUser(ctx, arg.Username)
	if err != nil {
		return err
	}

	if user.Email != arg.Email {
		_, err := m.Model.FindUser(ctx, arg.Email)
		if err != sql.ErrNoRows {
			fmt.Println(err)
			return fmt.Errorf("email already in use please choose another email address")
		}
	}

	_, err = m.Model.UpdateUser(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}
