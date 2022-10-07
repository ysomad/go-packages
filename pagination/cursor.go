package pagination

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

// encodeCursor encodes uuid and t into base64 string separated with ",".
func encodeCursor(oid string, t time.Time) string {
	s := fmt.Sprintf("%s,%s", oid, t.Format(time.RFC3339Nano))
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// decodeCursor decodes page cursor into string uuid and created at time,
// cursor must be a string splitted with "," and encoded into base64,
// not decoded token example: `ef6e33ec-5b1d-4ade-8dcc-b508262ee859,2006-01-02T15:04:05.999999999Z07:00`
func decodeCursor(cursor string) (string, time.Time, error) {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", time.Time{}, err
	}

	parts := strings.Split(string(b), ",")
	if len(parts) != 2 {
		return "", time.Time{}, errors.New("invalid cursor")
	}

	t, err := time.Parse(time.RFC3339Nano, parts[1])
	if err != nil {
		return "", time.Time{}, err
	}

	return parts[0], t, nil
}
