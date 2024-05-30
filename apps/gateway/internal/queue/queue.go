package queue

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// Config ...
type Config struct {
	NatsAddr   string `yaml:"nats_addr" env:"NATS_ADDR" env-default:"nats://nats:4222"`
	NatsSecret string `yaml:"nats_token" env:"NATS_SECRET" env-default:""`
}

// Queue ...
type Queue struct {
	config *Config
	conn   *nats.Conn
}

// New ...
func New(config *Config) (*Queue, error) {
	fmt.Printf("%#v\n", config)
	conn, err := nats.Connect(config.NatsAddr, nats.Token(config.NatsSecret))
	if err != nil {
		log.Printf("nats: connect error: %s", err.Error())
		return nil, err
	}

	queue := &Queue{
		config: config,
		conn:   conn,
	}

	return queue, nil
}
