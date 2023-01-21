package utils

import "time"

func CurrentTime() time.Time {
	return time.Now()
}

func FormatTimeToDate(time time.Time) string {
	return time.Format("02/01/2006")
}
