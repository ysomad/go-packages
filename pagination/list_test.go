package pagination

import (
	"fmt"
	"testing"
	"time"
)

// cursorListResult implements Result
type cursorListResult struct {
	id        string
	createdAt time.Time
}

func (obj cursorListResult) ID() string           { return obj.id }
func (obj cursorListResult) CreatedAt() time.Time { return obj.createdAt }

func generateTestResults(length int) []cursorListResult {
	res := make([]cursorListResult, length)
	now := time.Now()
	for i := 0; i < length; i++ {
		res[i] = cursorListResult{
			id:        fmt.Sprint(i),
			createdAt: now,
		}
	}
	return res
}

func TestNewCursorList(t *testing.T) {
	t.Parallel()

	type args[T Result] struct {
		objects  []T
		pageSize uint32
	}

	type test[T Result] struct {
		name             string
		args             args[T]
		nextPageFound    bool
		wantResultLength int
	}

	tests := []test[cursorListResult]{
		{
			name: "found 51, requested 50, next page exists",
			args: args[cursorListResult]{
				objects:  generateTestResults(51),
				pageSize: 50,
			},
			nextPageFound:    true,
			wantResultLength: 50,
		},
		{
			name: "found 138, requested 137, next page exists",
			args: args[cursorListResult]{
				objects:  generateTestResults(138),
				pageSize: 137,
			},
			nextPageFound:    true,
			wantResultLength: 137,
		},
		{
			name: "found 2, requested 1, next page exists",
			args: args[cursorListResult]{
				objects:  generateTestResults(2),
				pageSize: 1,
			},
			nextPageFound:    true,
			wantResultLength: 1,
		},
		{
			name: "found 35, requested 100, next page doesnt exist",
			args: args[cursorListResult]{
				objects:  generateTestResults(35),
				pageSize: 100,
			},
			nextPageFound:    false,
			wantResultLength: 35,
		},
		{
			name: "found 0, requested 50, next page doesnt exist",
			args: args[cursorListResult]{
				objects:  []cursorListResult{},
				pageSize: 50,
			},
			nextPageFound:    false,
			wantResultLength: 0,
		},
		{
			name: "found 0, requested 0, next page doesnt exist",
			args: args[cursorListResult]{
				objects:  []cursorListResult{},
				pageSize: 0,
			},
			nextPageFound:    false,
			wantResultLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCursorList(tt.args.objects, tt.args.pageSize)

			resLen := len(got.Results)
			if resLen != tt.wantResultLength {
				t.Errorf(
					"NewCursorList() got results length = %d, want results length = %d",
					resLen,
					tt.wantResultLength,
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
