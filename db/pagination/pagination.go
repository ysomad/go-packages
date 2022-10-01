package pagination

import (
	"time"

	"golang.org/x/exp/constraints"
)

// SeekUInt is data transfer object for seek pagination using Serial PKs.
type SeekUInt[T constraints.Unsigned] struct {
	Limit  uint16
	LastID T
}

func NewSeekUInt[T constraints.Unsigned](lastID T, limit uint16) SeekUInt[T] {
	return SeekUInt[T]{
		Limit:  limit,
		LastID: lastID,
	}
}

// SeekString is data transfer object for seek pagination using string as PKs, for example UUID.
type SeekString struct {
	LastCreatedAt time.Time
	LastID        string
	Limit         uint16
}

func NewSeekString(lastID string, lastCreatedAt time.Time, limit uint16) SeekString {
	return SeekString{
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
