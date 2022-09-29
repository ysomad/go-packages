package pagination

import (
	"time"

	"golang.org/x/exp/constraints"
)

// CursorUInt is data transfer object for cursor pagination using Serial PKs.
type CursorUInt[T constraints.Unsigned] struct {
	Limit  uint16
	LastID T
}

func NewCursorUInt[T constraints.Unsigned](lastID T, limit uint16) CursorUInt[T] {
	return CursorUInt[T]{
		Limit:  limit,
		LastID: lastID,
	}
}

// CursorString is data transfer object for cursor pagination using string as PKs, for example UUID.
type CursorString struct {
	LastCreatedAt time.Time
	LastID        string
	Limit         uint16
}

func NewCursorString(lastID string, lastCreatedAt time.Time, limit uint16) CursorString {
	return CursorString{
		LastCreatedAt: lastCreatedAt,
		LastID:        lastID,
		Limit:         limit,
	}
}

// Offset is data transfer object for limit offset pagination.
type Offset struct {
	Limit  uint16
	Offset uint16
}

func NewOffset(limit, offset uint16) Offset {
	return Offset{
		Limit:  limit,
		Offset: offset,
	}
}
