package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cafasaru/nats_starter/config"
	postgres "github.com/cafasaru/nats_starter/internal/postgres/sqlc"
	"github.com/cafasaru/nats_starter/internal/server"
	"github.com/cafasaru/nats_starter/pkg/jaeger"
	"github.com/cafasaru/nats_starter/pkg/logger"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/opentracing/opentracing-go"
)

// @title starter microservice
// @version 1.0
// @description starter microservice
// @termsOfService http://swagger.io/terms/

// @contact.name FirstName LastName
// @contact.url https://github.com/AleksK1NG
// @contact.email noreply@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /api/v1
func main() {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Info("Starting emails microservice")
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, DevelopmentMode: %t",
		cfg.AppVersion,
		cfg.Logger.Level,
		cfg.HTTP.Development,
	)
	appLogger.Infof("Success loaded config: %+v", cfg.AppVersion)

	tracer, closer, err := jaeger.InitJaeger(cfg)
	if err != nil {
		appLogger.Fatal("cannot create tracer", err)
	}
	appLogger.Info("Jaeger connected")

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	appLogger.Info("Opentracing connected")

	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:5432/%s?sslmode=%s",
		cfg.PostgreSQL.PostgresqlUser,
		cfg.PostgreSQL.PostgresqlPassword,
		cfg.PostgreSQL.PostgresqlHost,
		cfg.PostgreSQL.PostgresqlDBName,
		cfg.PostgreSQL.PostgresqlSSLMode,
	)

	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Connect to nats synadia cloud with creds
	nc, err := nats.Connect("tls://connect.ngs.global", nats.UserCredentials("./ssl/nats.creds"))
	if err != nil {
		panic(err)
	}

	log.Println("Connected to nats server")

	defer nc.Close()

	store := postgres.NewStore(conn)

	s := server.NewServer(appLogger, cfg, nc, store, tracer, nil)

	log.Fatal(s.Run())

}
