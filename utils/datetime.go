package utils

import (
	"fmt"
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

func AmadeusDateTimeConvert(dateStr string, timeStr string) time.Time {
	if dateStr == "" {
		return time.Unix(0, 0)
	}
	day, _ := strconv.Atoi(dateStr[0:2])
	month, _ := strconv.Atoi(dateStr[2:4])
	year, _ := strconv.Atoi(dateStr[4:6])
	if year < 25 {
		year += 2000
	} else {
		year += 1900
	}
	hour, minute := 0, 0
	if timeStr != "" {
		hour, _ = strconv.Atoi(timeStr[0:2])
		minute, _ = strconv.Atoi(timeStr[2:4])
	}
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
}

func DateToAmadeusDate(date time.Time) string {
	year := strconv.Itoa(date.Year())
	return fmt.Sprintf("%02d%02d%s", date.Day(), date.Month(), year[2:4])
}
