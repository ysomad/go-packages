package pagination

import (
	"time"
)

// Item is a interface which must be implemented by item from items passed into NewCursorList.
type Item interface {
	ID() string
	CreatedAt() time.Time
}

type CursorList[T Item] struct {
	Items          []T
	NextPageCursor string
}

// NewCursorList is a constructor for CursorList.
//
// items is a slice of records received from db;
// pageSize is amount of records requested from client;
// pageSize+1 is amount of records requested from db - MUST BE used for SQL LIMIT queries.
//
// NextPageCursor calculates from id and timestamp of the last record in objects with "," separator but only
// if amount of items equals to amount of records requested from db.
func NewCursorList[T Item](items []T, pageSize uint32) *CursorList[T] {
	length := len(items)
	list := &CursorList[T]{Items: items}

	if uint32(length) == pageSize+1 {
		lastItem := items[length-1]
		list.NextPageCursor = encodeCursor(lastItem.ID(), lastItem.CreatedAt())
		list.Items = list.Items[:length-1]
	}

	return list
}

// DecodeNextPageCursor is a helper method for decodeCursor.
func (l *CursorList[T]) DecodeNextPageCursor() (oid string, createdAt time.Time, err error) {
	return decodeCursor(l.NextPageCursor)
}

// SeekList is a result of seek pagination.
type SeekList[T any] struct {
	Items   []T
	HasNext bool
}

func NewSeekList[T any](items []T, pageSize uint32) *SeekList[T] {
	length := uint32(len(items))
	hasNext := length == pageSize+1

	if hasNext {
		items = items[:length-1]
	}

	return &SeekList[T]{
		Items:   items,
		HasNext: hasNext,
	}
}
