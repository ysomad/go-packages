package pagination

import "time"

// Cursor is data transfer object for cursor pagination using string as PKs.
type Cursor struct {
	NextPageCursor string
	PageSize       uint32
}

// DecodeNextPageCursor is a helper method for decodeCursor.
func (c Cursor) DecodeNextPageCursor() (uuid string, t time.Time, err error) {
	return decodeCursor(c.NextPageCursor)
}

type Offset struct {
	Limit  uint32
	Offset uint32
}
