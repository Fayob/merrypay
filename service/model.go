package service

import (
	"database/sql"
	"merrypay/db"
	"time"
)

type Queries struct {
	db db.SQLQueries
}

func NewQuery(db *sql.DB) *Queries {
	return &Queries{
		db: db,
	}
}

type User struct {
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Membership        string    `json:"membership"`
	Password          string    `json:"password"`
	UpdatedPasswordAt time.Time `json:"updated_password_at"`
	CreatedAt         time.Time `json:"created_at"`
}
