package sqlx

type Option func(*Client)

func MaxOpenConns(size int) Option {
	return func(c *Client) {
		c.maxOpenConns = size
	}
}
