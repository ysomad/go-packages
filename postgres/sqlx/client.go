package sqlx

import (
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const defaultMaxOpenConns = 1

// Client is implementation of postgres client using pgx as driver,
// sqlx as implementation of database/sql and squirrel as query builder
type Client struct {
	maxOpenConns int
	DB           *sqlx.DB
	Builder      sq.StatementBuilderType
}

func NewClient(connString string, opts ...Option) (*Client, error) {
	c := &Client{
		maxOpenConns: defaultMaxOpenConns,
	}

	for _, opt := range opts {
		opt(c)
	}

	var err error

	c.DB, err = sqlx.Connect("pgx", connString)
	if err != nil {
		return nil, err
	}

	c.DB.SetMaxOpenConns(c.maxOpenConns)
	c.Builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return c, err
}

func (c *Client) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
