package queue

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Request ...
func (q *Queue) Request(method string, data []byte, timeout time.Duration) ([]byte, error) {
	reply, err := q.conn.Request(method, data, timeout)
	if err != nil {
		log.Printf("queue: request error: %e", err)
		return nil, err
	}

	return reply.Data, nil
}

func (q *Queue) PublishRequest(method string, reply string, data []byte) error {
	err := q.conn.PublishRequest(method, reply, data)
	if err != nil {
		log.Printf("queue: publish request error: %e", err)
		return err
	}
	return nil
}

func (q *Queue) GetConn() *nats.Conn {
	return q.conn
}
