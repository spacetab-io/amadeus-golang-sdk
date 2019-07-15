package utils

import (
	"strconv"
	"time"
)

//dateStr 01021980
func AmadeusDateConvert(dateStr string, timeStr string) time.Time {
	if dateStr == "" {
		return time.Unix(0, 0)
	}
	day, _ := strconv.Atoi(dateStr[0:2])
	month, _ := strconv.Atoi(dateStr[2:4])
	year, _ := strconv.Atoi(dateStr[4:8])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
