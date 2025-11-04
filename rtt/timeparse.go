package rtt

import (
    "fmt"
    "strconv"
    "time"
)

// ParseHHMM parses an HHmm string (e.g. 1503) into time components (hour, min).
func ParseHHMM(s string) (h, m int, err error) {
    if len(s) != 4 { return 0, 0, fmt.Errorf("invalid HHmm length: %q", s) }
    h, err = strconv.Atoi(s[0:2]); if err != nil { return }
    m, err = strconv.Atoi(s[2:4]); if err != nil { return }
    if h < 0 || h > 23 || m < 0 || m > 59 { return 0, 0, fmt.Errorf("invalid time %q", s) }
    return
}

// ParseHHMMSS parses an HHmmss string (Working Timetable time) into h,m,s.
func ParseHHMMSS(s string) (h, m, sec int, err error) {
    if len(s) != 6 { return 0, 0, 0, fmt.Errorf("invalid HHmmss length: %q", s) }
    h, err = strconv.Atoi(s[0:2]); if err != nil { return }
    m, err = strconv.Atoi(s[2:4]); if err != nil { return }
    sec, err = strconv.Atoi(s[4:6]); if err != nil { return }
    if h < 0 || h > 23 || m < 0 || m > 59 || sec < 0 || sec > 59 { return 0,0,0, fmt.Errorf("invalid time %q", s) }
    return
}

// TimeToday converts an HHmm string to a time.Time on the provided date (local location).
func TimeToday(date time.Time, hhmm string) (time.Time, error) {
    h, m, err := ParseHHMM(hhmm); if err != nil { return time.Time{}, err }
    return time.Date(date.Year(), date.Month(), date.Day(), h, m, 0, 0, date.Location()), nil
}
