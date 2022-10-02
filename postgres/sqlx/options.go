package sqlx

type Option func(*Client)

func WithMaxOpenConns(size int) Option {
	return func(c *Client) {
		c.maxOpenConns = size
	}
}
