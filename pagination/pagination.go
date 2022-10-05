package pagination

import "time"

// Cursor is data transfer object for cursoro pagination using string as PKs, for example UUID.
type Cursor struct {
	NextPageCursor string
	PageSize       uint32
}

func NewCursor(pageSize uint32, cur string) Cursor {
	return Cursor{
		NextPageCursor: cur,
		PageSize:       pageSize,
	}
}

// DecodeNextPageCursor is a helper method for decodeCursor.
func (c Cursor) DecodeNextPageCursor() (uuid string, t time.Time, err error) {
	return decodeCursor(c.NextPageCursor)
}
