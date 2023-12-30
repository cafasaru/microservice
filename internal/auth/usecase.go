package auth

import (
	"context"

	postgres "github.com/cafasaru/nats_starter/internal/postgres/sqlc"
)

type Usecase interface {
	SignIn(ctx context.Context, username string, password string) (postgres.User, error)
}
