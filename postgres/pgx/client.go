package pgx

import (
	"context"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	defaultMaxPoolSize  = 1
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

// Client is implementation of postgres client using pgx
// and squirrel as query builder
type Client struct {
	maxPoolSize          int
	connAttempts         int
	connTimeout          time.Duration
	preferSimpleProtocol bool

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func NewClient(connString string, opts ...Option) (*Client, error) {
	c := &Client{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(c.maxPoolSize)
	poolConfig.ConnConfig.PreferSimpleProtocol = c.preferSimpleProtocol

	for c.connAttempts > 0 {
		c.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}
		defer c.Close()

		log.Printf("trying connecting to postgres, attempts left: %d", c.connAttempts)

		time.Sleep(c.connTimeout)

		c.connAttempts--
	}

	if err != nil {
		return nil, err
	}

	c.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return c, nil
}

func (p *Client) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
