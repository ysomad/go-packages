package pagination

// Cursor is data transfer object for seek pagination using string as PKs, for example UUID.
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
