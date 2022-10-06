package pagination

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	testUUID   string
	testTime   time.Time
	testCursor string
)

func generateTestCursor(uuid string, t time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s,%s", uuid, t.Format(time.RFC3339Nano))))
}

func TestMain(m *testing.M) {
	testUUID = "80862cb4-947a-4d64-8dbe-858fea7d84f2"
	testTime = time.Now()
	testCursor = generateTestCursor(testUUID, testTime)

	os.Exit(m.Run())
}

func Test_encodeCursor(t *testing.T) {
	type args struct {
		uuid string
		t    time.Time
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				uuid: testUUID,
				t:    testTime,
			},
			want: testCursor,
		},
		{
			name: "empty uuid",
			args: args{
				uuid: "",
				t:    testTime,
			},
			want: generateTestCursor("", testTime),
		},
		{
			name: "empty time",
			args: args{
				uuid: testUUID,
				t:    time.Time{},
			},
			want: generateTestCursor(testUUID, time.Time{}),
		},
		{
			name: "empty uuid and time",
			args: args{
				uuid: "",
				t:    time.Time{},
			},
			want: generateTestCursor("", time.Time{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeCursor(tt.args.uuid, tt.args.t)
			if got != tt.want {
				t.Errorf("encodeCursor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeCursor(t *testing.T) {
	tests := []struct {
		name          string
		encodedCursor string
		wantUUID      string
		wantTime      time.Time
		wantErr       bool
	}{
		{
			name:          "invalid cursor",
			encodedCursor: "invalid cursor",
			wantUUID:      "",
			wantTime:      time.Time{},
			wantErr:       true,
		},
		{
			name:          "success",
			encodedCursor: testCursor,
			wantUUID:      testUUID,
			wantTime:      testTime,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUUID, gotTime, err := decodeCursor(tt.encodedCursor)

			if (err != nil) != tt.wantErr {
				t.Errorf("decodeCursor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotUUID != tt.wantUUID {
				t.Errorf("decodeCursor() gotUUID = %v wantUUID %v", gotUUID, tt.wantUUID)
			}

			if !gotTime.Equal(tt.wantTime) {
				t.Errorf("decodeCursor() gotTime = %v, wantTime %v", gotTime, tt.wantTime)
			}
		})
	}
}
