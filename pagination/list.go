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
	objects        []T
	nextPageCursor string
}

// NewCursorList is a constructor for CursorList.
//
// objects is a slice of records received from db;
// pageSize is amount of records requested from client;
// pageSize+1 is amount of records requested from db - MUST BE used for SQL LIMIT queries.
//
// NextPageCursor calculates from id and timestamp of the last record in objects with "," separator but only
// if amount of objects equals to amount of records requested from db.
func NewCursorList[T Object](objects []T, pageSize uint32) *CursorList[T] {
	length := len(objects)
	list := &CursorList[T]{objects: objects}

	if uint32(length) == pageSize+1 {
		lastObj := objects[length-1]
		list.nextPageCursor = encodeCursor(lastObj.ID(), lastObj.CreatedAt())
		list.objects = list.objects[:length-1]
	}

	return list
}

func (l *CursorList[T]) Objects() []T           { return l.objects }
func (l *CursorList[T]) NextPageCursor() string { return l.nextPageCursor }

// DecodeNextPageCursor is a helper method for decodeCursor.
func (l *CursorList[T]) DecodeNextPageCursor() (oid string, createdAt time.Time, err error) {
	return decodeCursor(l.nextPageCursor)
}

// SeekList is a result of seek pagination.
type SeekList[T any] struct {
	objects []T
	hasNext bool
}

func NewSeekList[T any](objects []T, pageSize uint32) *SeekList[T] {
	length := uint32(len(objects))
	hasNext := length == pageSize+1

	if hasNext {
		objects = objects[:length-1]
	}

	return &SeekList[T]{
		objects: objects,
		hasNext: hasNext,
	}
}

func (l *SeekList[T]) Objects() []T  { return l.objects }
func (l *SeekList[T]) HasNext() bool { return l.hasNext }
