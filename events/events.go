package events

import (
	"fmt"
	"practice1/order_user_api_gateway/config"
	"practice1/order_user_api_gateway/pkg/logger"
	"practice1/order_user_api_gateway/pkg/pubsub"
)

// PubsubServer ...
type PubsubServer struct {
	cfg config.Config
	log logger.LoggerI
	RMQ *pubsub.RMQ
}

// New ...
func New(cfg config.Config, log logger.LoggerI) (*PubsubServer, error) {
	rmq, err := pubsub.NewRMQ(
		fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.RmqUsername, cfg.RmqPassword, cfg.RmqUriHOST, cfg.RmqUriPORT),
		log,
	)
	if err != nil {
		return nil, err
	}
	log.Info("---rabbit connected--->", logger.Any("url", fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.RmqUsername, cfg.RmqPassword, cfg.RmqUriHOST, cfg.RmqUriPORT)))

	rmq.AddPublisher("v1.fiscal") // one publisher is enough for application service

	return &PubsubServer{
		cfg: cfg,
		log: log,
		RMQ: rmq,
	}, nil
}
