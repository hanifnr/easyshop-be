package utils

import "time"

func CurrentTime() time.Time {
	return time.Now()
}

func FormatTimeToDate(time time.Time) string {
	return time.Format("02/01/2006")
}

func FormatTimeToyyyymmdd(time *time.Time) string {
	if time == nil {
		return ""
	}
	return time.Format("2006010")
}

func FormatTimeDetail(time time.Time) string {
	return time.Format("20060102150405")
}
