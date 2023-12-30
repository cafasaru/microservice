package server

import (
	"net"

	"github.com/cafasaru/microservice/config"
	"github.com/cafasaru/microservice/gen/v1/authpb"
	authgRPC "github.com/cafasaru/microservice/internal/auth/transport/grpc"
	"github.com/cafasaru/microservice/internal/auth/usecase"
	postgres "github.com/cafasaru/microservice/internal/postgres/sqlc"
	"github.com/cafasaru/microservice/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	certFile        = "ssl/server.crt"
	keyFile         = "ssl/server.pem"
	maxHeaderBytes  = 1 << 20
	gzipLevel       = 5
	stackSize       = 1 << 10 // 1 KB
	csrfTokenHeader = "X-CSRF-Token"
	bodyLimit       = "2M"
)

type server struct {
	log      logger.Logger
	cfg      *config.Config
	natsConn *nats.Conn
	store    *postgres.Store
	tracer   opentracing.Tracer
	echo     *echo.Echo
	redis    *redis.Client
}

// NewServer constructor
func NewServer(
	log logger.Logger,
	cfg *config.Config,
	natsConn *nats.Conn,
	store *postgres.Store,
	tracer opentracing.Tracer,
	redis *redis.Client,
) *server {
	return &server{log: log, cfg: cfg, natsConn: natsConn, store: store, tracer: tracer, redis: redis, echo: echo.New()}
}

func (s *server) Run() error {

	validate := validator.New()

	uc := usecase.NewUsecase()
	service := authgRPC.NewgRPCService(uc, s.log, validate)

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, service)

	s.log.Infof("starting gRPC server on %s", s.cfg.GRPC.Port)
	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return errors.Wrap(err, "net.Listen")
	}

	defer l.Close()

	if s.cfg.HTTP.Development {
		s.log.Info("reflection enabled")
		reflection.Register(grpcServer)
	}

	s.log.Fatal(grpcServer.Serve(l))

	return nil
}
