package util

import (
	"fmt"
	"time"
)

//JsonTime extend time
type JsonTime time.Time

// UnmarshalJSON sets *t to a copy of data.
func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	parsed, err := time.Parse(`"2006-01-02 00:00:00"`, string(data))
	if err != nil {
		return err
	}
	*t = JsonTime(parsed)
	return
}

// MarshalJSON returns t as the JSON encoding of t.
func (t JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 00:00:00"))
	return []byte(stamp), nil
}

// String return t as the string of t.
func (t JsonTime) String() string {
	return time.Time(t).Format(time.RFC3339Nano)
}

// Time return t as the time of t.
func (t JsonTime) Time() time.Time {
	return time.Time(t)
}
