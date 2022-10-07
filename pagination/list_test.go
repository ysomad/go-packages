package pagination

import (
	"fmt"
	"testing"
	"time"
)

// testCursorListItem implements Result
type testCursorListItem struct {
	id        string
	createdAt time.Time
}

func (obj testCursorListItem) ID() string           { return obj.id }
func (obj testCursorListItem) CreatedAt() time.Time { return obj.createdAt }

func generateTestCursorListItems(length int) []testCursorListItem {
	res := make([]testCursorListItem, length)
	now := time.Now()
	for i := 0; i < length; i++ {
		res[i] = testCursorListItem{
			id:        fmt.Sprint(i),
			createdAt: now,
		}
	}
	return res
}

func TestNewCursorList(t *testing.T) {
	type args[T testCursorListItem] struct {
		objects  []testCursorListItem
		pageSize uint32
	}

	type test[T testCursorListItem] struct {
		name           string
		args           args[testCursorListItem]
		nextPageFound  bool
		wantItemsCount int
	}

	tests := []test[testCursorListItem]{
		{
			name: "found 51, requested 50, next page exists",
			args: args[testCursorListItem]{
				objects:  generateTestCursorListItems(51),
				pageSize: 50,
			},
			nextPageFound:  true,
			wantItemsCount: 50,
		},
		{
			name: "found 138, requested 137, next page exists",
			args: args[testCursorListItem]{
				objects:  generateTestCursorListItems(138),
				pageSize: 137,
			},
			nextPageFound:  true,
			wantItemsCount: 137,
		},
		{
			name: "found 2, requested 1, next page exists",
			args: args[testCursorListItem]{
				objects:  generateTestCursorListItems(2),
				pageSize: 1,
			},
			nextPageFound:  true,
			wantItemsCount: 1,
		},
		{
			name: "found 35, requested 100, next page doesnt exist",
			args: args[testCursorListItem]{
				objects:  generateTestCursorListItems(35),
				pageSize: 100,
			},
			nextPageFound:  false,
			wantItemsCount: 35,
		},
		{
			name: "found 0, requested 50, next page doesnt exist",
			args: args[testCursorListItem]{
				objects:  []testCursorListItem{},
				pageSize: 50,
			},
			nextPageFound:  false,
			wantItemsCount: 0,
		},
		{
			name: "found 0, requested 0, next page doesnt exist",
			args: args[testCursorListItem]{
				objects:  []testCursorListItem{},
				pageSize: 0,
			},
			nextPageFound:  false,
			wantItemsCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCursorList(tt.args.objects, tt.args.pageSize)

			resLen := len(got.Objects)
			if resLen != tt.wantItemsCount {
				t.Errorf(
					"NewCursorList() got items count = %d, want items count = %d",
					resLen,
					tt.wantItemsCount,
				)
			}

			if (got.NextPageCursor != "") != tt.nextPageFound {
				t.Errorf(
					"NewCursorList() got next page cursor = %s, want next page cursor = %t",
					got.NextPageCursor,
					tt.nextPageFound,
				)
			}
		})
	}
}

type testSeekListItem struct{}

func generateTestSeekListItems(length int) []testSeekListItem {
	res := make([]testSeekListItem, length)
	for i := 0; i < length; i++ {
		res[i] = testSeekListItem{}
	}
	return res
}

func TestNewSeekList(t *testing.T) {
	type args struct {
		items    []testSeekListItem
		pageSize uint32
	}

	type test struct {
		name           string
		args           args
		wantItemsCount int
		wantHasNext    bool
	}

	tests := []test{
		{
			// TODO: write tests
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeekList(tt.args.items, tt.args.pageSize)
			// TODO: handle tests
		})
	}
}
