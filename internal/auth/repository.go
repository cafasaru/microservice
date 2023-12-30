package auth

import (
	"context"

	postgres "github.com/cafasaru/nats_starter/internal/postgres/sqlc"
)

// PGRespository interface is the interface for interacting with the Postgres database
type PGRespository interface {
	Login(ctx context.Context, username string, password string) (postgres.User, error)
	RefreshToken(ctx context.Context, refreshToken string) (postgres.User, error)
	VerifyToken(ctx context.Context, accessToken string) (postgres.User, error)
}
