package service

import (
	"context"
	"fmt"
	"time"
)

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
		// 	if pqErr, ok := err.(*pq.Error); ok {
		// 		if pqErr.Constraint == "users_email_key" {
		// 			fmt.Println(pqErr.Code)
		// 			return "", fmt.Errorf("email already in use")
		// 		}
		// 	}
		return "", err
	}

	fmt.Println("User Created Successfully")
	return "User Created Successfully", err
}

func (q *Queries) FindUser(ctx context.Context, arg string) (User, error) {
	var user User
	query := `SELECT username, first_name, last_name, email, membership, password, updated_password_at, created_at
						FROM users where username = $1 or email = $1`
	res := q.db.QueryRowContext(ctx, query, arg)
	err := res.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Membership,
		&user.Password,
		&user.UpdatedPasswordAt,
		&user.CreatedAt,
	)
	return user, err
}

func (q *Queries) FindAllUsers(ctx context.Context) ([]User, error) {
	query := `SELECT username, first_name, last_name, email, membership, 
							password, updated_password_at, created_at FROM users`

	rows, err := q.db.QueryContext(ctx, query)
	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Membership,
			&user.Password,
			&user.UpdatedPasswordAt,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

type UpdateUserParams struct {
	Username          string `json:"username"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Membership        string `json:"membership"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (string, error) {
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, membership = $4 
						where username = $5`
	
	_, err := q.db.ExecContext(ctx, query, arg.FirstName, arg.LastName, arg.Email, arg.Membership, arg.Username)
	// var user User
	if err != nil {
		fmt.Println("scanning error")
		return "", err
	}
	return fmt.Sprintf("%s's profile updated successfully", arg.Username), nil
}
