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
	ItemsCount     uint32
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

	list.ItemsCount = uint32(len(list.Items))

	return list
}

// DecodeNextPageCursor is a helper method for decodeCursor.
func (l *CursorList[T]) DecodeNextPageCursor() (uuid string, t time.Time, err error) {
	return decodeCursor(l.NextPageCursor)
}

type OffsetList[T any] struct {
	Items      []T
	ItemsCount uint32
}

func NewOffsetList[T any](items []T) *OffsetList[T] {
	return &OffsetList[T]{
		Items:      items,
		ItemsCount: uint32(len(items)),
	}
}
