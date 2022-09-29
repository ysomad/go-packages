package pagination

import (
	"time"

	"golang.org/x/exp/constraints"
)

// SerialParams is generic data transfer object for cursor pagination using Serial PKs.
type SerialParams[T constraints.Unsigned] struct {
	Limit  uint16 `json:"limit"`
	LastID T      `json:"last_id"`
}

func NewSerialParams[T constraints.Unsigned](lastID T, limit uint16) SerialParams[T] {
	return SerialParams[T]{
		Limit:  limit,
		LastID: lastID,
	}
}

// UUIDParams is data transfer object for cursor pagination using UUID as PKs.
type UUIDParams struct {
	LastCreatedAt time.Time `json:"last_created_at"`
	LastID        string    `json:"last_id"`
	Limit         uint16    `json:"limit"`
}

func NewUUIDParams(lastID string, limit uint16, lastCreatedAt time.Time) UUIDParams {
	return UUIDParams{
		LastCreatedAt: lastCreatedAt,
		LastID:        lastID,
		Limit:         limit,
	}
}

// OffsetParams is data transfer object for limit offset pagination.
type OffsetParams struct {
	Limit  uint16 `json:"limit"`
	Offset uint16 `json:"offset"`
}

func NewOffsetParams(limit, offset uint16) OffsetParams {
	return OffsetParams{
		Limit:  limit,
		Offset: offset,
	}
}

// List is generic object for pagination implementations.
// Usage: `func findRecords(p Params) List[YourModel]`.
type List[T any] struct {
	Result  []T  `json:"result"`
	HasNext bool `json:"has_next"`
}

// NewList creates new List with result and has next indicator which is
// indicates is there more records in db.
// It's assumed that List will be used when querying limit + 1 items from db,
// last item will be removed from `List.Result`.
// Whether there are more records in db will be determined by length of objList.
func NewList[T any](objList []T, limit uint16) List[T] {
	objLen := uint16(len(objList))
	hasNext := objLen == limit+1
	if hasNext {
		objList = objList[:objLen-1]
	}
	return List[T]{
		Result:  objList,
		HasNext: hasNext,
	}
}
