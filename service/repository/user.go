package repository

import (
	"context"
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
