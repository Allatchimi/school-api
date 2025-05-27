package utils

import "time"

func AreDateEquals(date1 time.Time, date2 time.Time) bool {
	return date1.UTC().Compare(date2.UTC()) == 0
}
