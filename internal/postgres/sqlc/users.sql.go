// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, email, username, password_hash, verification_code, verification_expiration, verified, active, created_at, updated_at
FROM users 
WHERE user_id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, userID uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.VerificationCode,
		&i.VerificationExpiration,
		&i.Verified,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT user_id, email, username, password_hash, verification_code, verification_expiration, verified, active, created_at, updated_at
FROM users 
WHERE UPPER(username) = UPPER($1) LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, upper interface{}) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, upper)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.VerificationCode,
		&i.VerificationExpiration,
		&i.Verified,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const saveUser = `-- name: SaveUser :one

INSERT INTO users (
    user_id,
    email,
    username,
    password_hash,
    verification_code,
    verification_expiration,
    verified,
    active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING user_id, email, username, password_hash, verification_code, verification_expiration, verified, active, created_at, updated_at
`

type SaveUserParams struct {
	UserID                 uuid.UUID      `json:"user_id"`
	Email                  string         `json:"email"`
	Username               string         `json:"username"`
	PasswordHash           []byte         `json:"password_hash"`
	VerificationCode       sql.NullString `json:"verification_code"`
	VerificationExpiration sql.NullTime   `json:"verification_expiration"`
	Verified               bool           `json:"verified"`
	Active                 bool           `json:"active"`
}

// ------- USERS SAVE ------------------
func (q *Queries) SaveUser(ctx context.Context, arg SaveUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, saveUser,
		arg.UserID,
		arg.Email,
		arg.Username,
		arg.PasswordHash,
		arg.VerificationCode,
		arg.VerificationExpiration,
		arg.Verified,
		arg.Active,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.VerificationCode,
		&i.VerificationExpiration,
		&i.Verified,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}