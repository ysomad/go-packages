package pagination

import (
	"time"
)

// Result is a interface which must be implemented by item from objects passed into NewCursorList.
type Result interface {
	ID() string
	CreatedAt() time.Time
}

type CursorList[T Result] struct {
	Results        []T    `json:"results"`
	NextPageCursor string `json:"next_page_cursor"`
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
	length := len(objects)
	list := &CursorList[T]{Results: objects}

	if uint32(length) == pageSize+1 {
		lastObj := objects[length-1]
		list.NextPageCursor = encodeCursor(lastObj.ID(), lastObj.CreatedAt())
		list.Results = list.Results[:length-1]
	}

	return list
}

// DecodeNextPageCursor is a helper method for decodeCursor.
func (l *CursorList[T]) DecodeNextPageCursor() (uuid string, t time.Time, err error) {
	return decodeCursor(l.NextPageCursor)
}
