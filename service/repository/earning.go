package repository

import (
	"context"
	"merrypay/types"
	"strings"
)

func (m *Model) GetEarning(ctx context.Context, username string) (types.Earning, error) {
	username = strings.ToLower(username)
	earning, err := m.Model.GetEarning(ctx, username)
	if err != nil {
		return types.Earning{}, err
	}
	return earning, nil
}