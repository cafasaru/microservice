package usecase

import (
	"context"

	postgres "github.com/cafasaru/microservice/internal/postgres/sqlc"
)

type authUsecase struct{}

func NewUsecase() *authUsecase {
	return &authUsecase{}
}

func (a *authUsecase) SignIn(ctx context.Context, username string, password string) (postgres.User, error) {
	return postgres.User{}, nil

}
