package pagination

import (
	"time"
)

// Object is a interface which must be implemented by object from objects arg passed into NewCursorList.
type Object interface {
	ID() string
	CreatedAt() time.Time
}

type CursorList[T Object] struct {
	Objects        []T
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
func NewCursorList[T Object](objects []T, pageSize uint32) *CursorList[T] {
	length := len(objects)
	list := &CursorList[T]{Objects: objects}

	if uint32(length) == pageSize+1 {
		lastObj := objects[length-1]
		list.NextPageCursor = encodeCursor(lastObj.ID(), lastObj.CreatedAt())
		list.Objects = list.Objects[:length-1]
	}

	return list
}

// DecodeNextPageCursor is a helper method for decodeCursor.
func (l *CursorList[T]) DecodeNextPageCursor() (oid string, createdAt time.Time, err error) {
	return decodeCursor(l.NextPageCursor)
}

// SeekList is a result of seek pagination.
type SeekList[T any] struct {
	Objects []T
	HasNext bool
}

func NewSeekList[T any](objects []T, pageSize uint32) *SeekList[T] {
	length := uint32(len(objects))
	hasNext := length == pageSize+1

	if hasNext {
		objects = objects[:length-1]
	}

	return &SeekList[T]{
		Objects: objects,
		HasNext: hasNext,
	}
}
