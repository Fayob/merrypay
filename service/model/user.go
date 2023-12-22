package service

import (
	"context"
	"fmt"
	"merrypay/types"
	"time"
)

func (q *Queries) CreateUser(ctx context.Context, arg types.CreateUserParams) (types.User, error) {
	query := `INSERT INTO users(username, first_name, last_name, email, password, referred_by) 
						VALUES($1, $2, $3, $4, $5, $6) RETURNING username, first_name, last_name, email, membership, 
						won_jackpot, referred_by, password, updated_password_at, created_at`

	row := q.db.QueryRowContext(ctx, query, arg.Username, arg.FirstName, arg.LastName, arg.Email, arg.Password, arg.Referral)
	var user types.User
	err := row.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Membership,
		&user.WonJackpot,
		&user.ReferredBy,
		&user.Password,
		&user.UpdatedPasswordAt,
		&user.CreatedAt,
	)
	// 	if pqErr, ok := err.(*pq.Error); ok {
	// 		if pqErr.Constraint == "users_email_key" {
	// 			fmt.Println(pqErr.Code)
	// 			return "", fmt.Errorf("email already in use")
	// 		}
	// 	}

	return user, err
}

func (q *Queries) FindUser(ctx context.Context, arg string) (types.User, error) {
	var user types.User
	query := `SELECT username, first_name, last_name, email, membership, won_jackpot, referred_by, password, updated_password_at, created_at
						FROM users where username = $1 or email = $1`
	res := q.db.QueryRowContext(ctx, query, arg)
	err := res.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Membership,
		&user.WonJackpot,
		&user.ReferredBy,
		&user.Password,
		&user.UpdatedPasswordAt,
		&user.CreatedAt,
	)
	return user, err
}

func (q *Queries) FindAllUsers(ctx context.Context) ([]types.User, error) {
	query := `SELECT username, first_name, last_name, email, membership, won_jackpot, 
						referred_by, password, updated_password_at, created_at FROM users`

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var users []types.User

	for rows.Next() {
		var user types.User
		if err := rows.Scan(
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Membership,
			&user.WonJackpot,
			&user.ReferredBy,
			&user.Password,
			&user.UpdatedPasswordAt,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (q *Queries) UpdateUser(ctx context.Context, arg types.UpdateUserParams) (types.User, error) {
	query := `UPDATE users SET first_name = $2, last_name = $3, email = $4, membership = $5, won_jackpot = $6 where username = $1 
						RETURNING username, first_name, last_name, email, membership, won_jackpot, referred_by, created_at`

	row := q.db.QueryRowContext(ctx, query, arg.Username, arg.FirstName, arg.LastName, arg.Email, arg.Membership, arg.WonJackpot)

	var updatedUser types.User
	err := row.Scan(
		&updatedUser.Username,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Email,
		&updatedUser.Membership,
		&updatedUser.WonJackpot,
		&updatedUser.ReferredBy,
		&updatedUser.CreatedAt,
	)

	return updatedUser, err
}

func (q *Queries) UpdatePassword(ctx context.Context, password, identifier string) (string, error) {
	query := `UPDATE users SET password = $1, Updated_password_at = $2 where username = $3 or email = $3`

	_, err := q.db.ExecContext(ctx, query, password, time.Now(), identifier)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s's password updated successfully", identifier), nil
}

func (q *Queries) DeleteUser(ctx context.Context, arg string) (string, error) {
	_, err := q.db.ExecContext(ctx, `DELETE FROM users where username = $1 or email = $1`, arg)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s deleted successfully", arg), nil
}

func (q *Queries) ReferralHistory(ctx context.Context, username string) ([]types.RefHisResponse, error) {
	query := `SELECT username, email, first_name, last_name FROM users where referred_by = $1`

	rows, err := q.db.QueryContext(ctx, query, username)
	if err != nil {
		return nil, err
	}
	var referrals []types.RefHisResponse
	for rows.Next() {
		var referral types.RefHisResponse
		if err := rows.Scan(
			&referral.Username,
			&referral.Email,
			&referral.FirstName,
			&referral.LastName,
		); err != nil {
			return nil, err
		}
		referrals = append(referrals, referral)
	}

	return referrals, nil
}
