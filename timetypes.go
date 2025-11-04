package realtimetrains

import (
	"encoding/json"
	"fmt"
	"time"
)

// TimeHHMM represents a clock time (HHmm) without a date.
// Internally stored as time.Time with date zero (year 0, Jan 1) in UTC.
type TimeHHMM struct{ time.Time }

// TimeHHMMSS represents a clock time (HHmmss) without a date.
type TimeHHMMSS struct{ time.Time }

func parseHHMM(s string) (time.Time, error) {
	if len(s) != 4 {
		return time.Time{}, fmt.Errorf("invalid HHmm: %q", s)
	}
	h := (int(s[0]-'0')*10 + int(s[1]-'0'))
	m := (int(s[2]-'0')*10 + int(s[3]-'0'))
	if h > 23 || m > 59 {
		return time.Time{}, fmt.Errorf("invalid HHmm: %q", s)
	}
	return time.Date(0, 1, 1, h, m, 0, 0, time.UTC), nil
}
func parseHHMMSS(s string) (time.Time, error) {
	if len(s) != 6 {
		return time.Time{}, fmt.Errorf("invalid HHmmss: %q", s)
	}
	h := (int(s[0]-'0')*10 + int(s[1]-'0'))
	m := (int(s[2]-'0')*10 + int(s[3]-'0'))
	sec := (int(s[4]-'0')*10 + int(s[5]-'0'))
	if h > 23 || m > 59 || sec > 59 {
		return time.Time{}, fmt.Errorf("invalid HHmmss: %q", s)
	}
	return time.Date(0, 1, 1, h, m, sec, 0, time.UTC), nil
}

// UnmarshalJSON implements json.Unmarshaler for TimeHHMM.
func (t *TimeHHMM) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	tt, err := parseHHMM(s)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t TimeHHMM) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return fmt.Sprintf("%02d%02d", t.Hour(), t.Minute())
}

// UnmarshalJSON implements json.Unmarshaler for TimeHHMMSS.
func (t *TimeHHMMSS) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	tt, err := parseHHMMSS(s)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t TimeHHMMSS) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return fmt.Sprintf("%02d%02d%02d", t.Hour(), t.Minute(), t.Second())
}
