package pagination

import (
	"time"
)

// Result is a interface which must be implemented by item from objects passed into `NewCursorList`.
type Result interface {
	ID() string
	CreatedAt() time.Time
}

type CursorList[T Result] struct {
	Results        []Result `json:"results"`
	NextPageCursor string   `json:"next_page_cursor"`
}

// NewCursorList is a constructor for CursorList.
//
// objects is a slice of records received from db;
// pageSize is amount of records requested from client;
// pageSize+1 is amount of records requested from db - MUST BE used for SQL LIMIT queries.
//
// NextPageCursor calculates from id and timestamp of the last record in objects with "," separator but only
// if length of objects equals to amount of records requested from db.
func NewCursorList[T Result](objects []T, pageSize uint32) *CursorList[T] {
	list := new(CursorList[T])
	length := len(objects)

	for i, obj := range objects {
		if uint32(length) == pageSize+1 && i == length-1 {
			list.NextPageCursor = encodeCursor(obj.ID(), obj.CreatedAt())
			break
		}

		list.Results = append(list.Results, obj)
	}

	return list
}

// DecodeNextPageCursor is a helper method for `decodeCursor`.
func (l *CursorList[T]) DecodeNextPageCursor() (uuid string, t time.Time, err error) {
	return decodeCursor(l.NextPageCursor)
}
