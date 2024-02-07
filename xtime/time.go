package xtime

import (
	"fmt"
	"time"
)

func StringTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// Time function with time zone parameter
func Time(s, timeZone string) time.Time {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		// Handle error, possibly fallback to a default time zone
		fmt.Println("Error loading location:", err)
		loc = time.UTC
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	if err != nil {
		// Handle error
		fmt.Println("Error parsing time:", err)
	}

	return t
}
