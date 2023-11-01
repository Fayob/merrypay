package service

import (
	"context"
)

type CreateUserParams struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (string, error) {
	_, err := q.db.ExecContext(ctx, `INSERT INTO users(username, first_name, last_name, email, password)
	VALUES($1, $2, $3, $4, $5)`, arg.Username, arg.FirstName, arg.LastName, arg.Email, arg.Password)


	if err != nil {
		return "", err
	}

	return "User Created Successfully", err
}

