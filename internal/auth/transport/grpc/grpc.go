package grpc

import (
	"context"
	"log"

	"github.com/cafasaru/microservice/gen/v1/authpb"
	"github.com/cafasaru/microservice/internal/auth"
	"github.com/cafasaru/microservice/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
)

type gRPCService struct {
	authUsecase auth.Usecase
	log         logger.Logger
	validator   *validator.Validate
	authpb.UnimplementedAuthServiceServer
}

// NewAuthService constructor
func NewgRPCService(authUsecase auth.Usecase, log logger.Logger, validator *validator.Validate) *gRPCService {
	return &gRPCService{authUsecase: authUsecase, log: log, validator: validator}
}

func (g *gRPCService) SignIn(ctx context.Context, req *authpb.SignInRequest) (*authpb.SignInResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc.auth.SignIn")
	defer span.Finish()

	if err := g.validator.Struct(req); err != nil {
		return nil, err
	}

	user, err := g.authUsecase.SignIn(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	log.Println(user)

	return &authpb.SignInResponse{
		Success:      true,
		AccessToken:  "testing",
		RefreshToken: "testing",
		Message:      "Testing Message",
	}, nil

}
