package nats

import (
	"time"

	"github.com/cafasaru/nats_starter/config"
	"github.com/cafasaru/nats_starter/pkg/logger"
	"github.com/nats-io/nats.go"
)

const (
	connectWait   = time.Second * 30
	interval      = 10
	maxOut        = 5
	MaxReconnects = -1 // infinite
)

func NewNatsConnection(cfg *config.Config, log logger.Logger, tls nats.Option) (*nats.Conn, error) {
	nc, err := nats.Connect(cfg.Nats.URL,
		tls,
		nats.Name(cfg.Nats.ClusterID),
		nats.Timeout(connectWait),
		nats.PingInterval(interval),
		nats.MaxPingsOutstanding(maxOut),
		nats.MaxReconnects(MaxReconnects),
		nats.ReconnectWait(connectWait),
		nats.DisconnectHandler(func(nc *nats.Conn) {
			log.Fatal("Disconnected from NATS")
		}),
	)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
