package repository

import "merrypay/port"

type Model struct {
	Model port.Store
}

func NewModel(model port.Store) *Model {
	return &Model{
		Model: model,
	}
}
