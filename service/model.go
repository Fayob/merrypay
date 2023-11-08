package service

import (
	"database/sql"
	"merrypay/db"
)

type Queries struct {
	db db.SQLQueries
}

func NewQuery(db *sql.DB) *Queries {
	return &Queries{
		db: db,
	}
}
