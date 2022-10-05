package pagination

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

// encodeCursor encodes uuid and t into base64 string separated with ",".
func encodeCursor(uuid string, t time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s,%s", uuid, t.Format(time.RFC3339Nano))))
}

// decodeCursor decodes page cursor into string uuid and created at time,
// cursor must be a string splitted with "," and encoded into base64,
// not decoded token example: `ef6e33ec-5b1d-4ade-8dcc-b508262ee859,`
func decodeCursor(cursor string) (string, time.Time, error) {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", time.Time{}, err
	}

	items := strings.Split(string(b), ",")
	if len(items) != 2 {
		return "", time.Time{}, errors.New("invalid cursor")
	}

	t, err := time.Parse(time.RFC3339Nano, items[1])
	if err != nil {
		return "", time.Time{}, err
	}

	return items[0], t, nil
}
