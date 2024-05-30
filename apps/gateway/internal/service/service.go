package service

import (
	"time"

	"github.com/nats-io/nats.go"
)

// Config ...
type Config struct {
}

// QueueInterface ...
type QueueInterface interface {
	Request(string, []byte, time.Duration) ([]byte, error)
	PublishRequest(string, string, []byte) error
	GetConn() *nats.Conn
}

// Service ...
type Service struct {
	config *Config
	queue  QueueInterface
}

// New ...
func New(config *Config, queue QueueInterface) (*Service, error) {
	service := &Service{
		config: config,
		queue:  queue,
	}

	return service, nil
}
