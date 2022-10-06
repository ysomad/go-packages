package pagination

import (
	"fmt"
	"testing"
	"time"
)

// cursorListItem implements Result
type cursorListItem struct {
	id        string
	createdAt time.Time
}

func (obj cursorListItem) ID() string           { return obj.id }
func (obj cursorListItem) CreatedAt() time.Time { return obj.createdAt }

func generateTestItems(length int) []cursorListItem {
	res := make([]cursorListItem, length)
	now := time.Now()
	for i := 0; i < length; i++ {
		res[i] = cursorListItem{
			id:        fmt.Sprint(i),
			createdAt: now,
		}
	}
	return res
}

func TestNewCursorList(t *testing.T) {
	type args[T cursorListItem] struct {
		objects  []cursorListItem
		pageSize uint32
	}

	type test[T cursorListItem] struct {
		name           string
		args           args[cursorListItem]
		nextPageFound  bool
		wantItemsCount uint32
	}

	tests := []test[cursorListItem]{
		{
			name: "found 51, requested 50, next page exists",
			args: args[cursorListItem]{
				objects:  generateTestItems(51),
				pageSize: 50,
			},
			nextPageFound:  true,
			wantItemsCount: 50,
		},
		{
			name: "found 138, requested 137, next page exists",
			args: args[cursorListItem]{
				objects:  generateTestItems(138),
				pageSize: 137,
			},
			nextPageFound:  true,
			wantItemsCount: 137,
		},
		{
			name: "found 2, requested 1, next page exists",
			args: args[cursorListItem]{
				objects:  generateTestItems(2),
				pageSize: 1,
			},
			nextPageFound:  true,
			wantItemsCount: 1,
		},
		{
			name: "found 35, requested 100, next page doesnt exist",
			args: args[cursorListItem]{
				objects:  generateTestItems(35),
				pageSize: 100,
			},
			nextPageFound:  false,
			wantItemsCount: 35,
		},
		{
			name: "found 0, requested 50, next page doesnt exist",
			args: args[cursorListItem]{
				objects:  []cursorListItem{},
				pageSize: 50,
			},
			nextPageFound:  false,
			wantItemsCount: 0,
		},
		{
			name: "found 0, requested 0, next page doesnt exist",
			args: args[cursorListItem]{
				objects:  []cursorListItem{},
				pageSize: 0,
			},
			nextPageFound:  false,
			wantItemsCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCursorList(tt.args.objects, tt.args.pageSize)

			if got.ItemsCount != tt.wantItemsCount {
				t.Errorf(
					"NewCursorList() got.ItemsCount = %d, wantItemsCount = %d",
					got.ItemsCount,
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
