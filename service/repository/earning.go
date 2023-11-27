package repository

import (
	"context"
	"merrypay/types"
)

func (m *Model) GetEarning(ctx context.Context, username string) (types.Earning, error) {
	earning, err := m.Model.GetEarning(ctx, username)
	if err != nil {
		return types.Earning{}, err
	}
	return earning, nil
}