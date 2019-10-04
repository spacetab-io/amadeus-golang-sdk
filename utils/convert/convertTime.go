package convert

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/metakeule/fmtdate"
)

// AmadeusTimeToMinutes converts amadeus time to minutes
// 1504 to minutes int
func AmadeusTimeToMinutes(timeStr string) int {
	if timeStr == "" {
		return 0
	}
	hours, err := strconv.Atoi(timeStr[0:2])
	if err != nil {
		hours = 0
	}
	minutes, err := strconv.Atoi(timeStr[2:])
	if err != nil {
		minutes = 0
	}
	return hours*60 + minutes
}

// AmadeusDateConvert convert 01JAN95 to time
func AmadeusDateConvert(timeStr string) (time.Time, error) {
	result := time.Unix(0, 0)
	if timeStr == "" {
		return result, errors.New("string empty")
	}
	day, err := strconv.Atoi(timeStr[0:2])
	if err != nil {
		return result, err
	}
	if day < 1 || day > 31 {
		return result, errors.New("invalid day format")
	}
	monthStr := timeStr[2:5]
	month := time.Month(0)
	switch monthStr {
	case "JAN":
		month = time.January
	case "FEB":
		month = time.February
	case "MAR":
		month = time.March
	case "APR":
		month = time.April
	case "MAY":
		month = time.May
	case "JUN":
		month = time.June
	case "JUL":
		month = time.July
	case "AUG":
		month = time.August
	case "SEP":
		month = time.September
	case "OCT":
		month = time.October
	case "NOV":
		month = time.November
	case "DEC":
		month = time.December
	default:
		return result, errors.New("invalid month format")
	}
	year, err := strconv.Atoi(timeStr[5:7])
	if err != nil {
		return result, err
	}
	if year < 25 {
		year += 2000
	} else {
		year += 1900
	}
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), nil
}

// AmadeusDateTimeConvert converts 010219 to time
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

// DateTimeToInt convert time to int
func DateTimeToInt(date time.Time) (d, t int) {
	var hours = date.Hour() * 100
	var minutes = date.Minute()
	var year = strconv.Itoa(date.Year())
	var str = fmt.Sprintf("%02d%02d%s", date.Day(), date.Month(), year[2:4])
	if i, err := strconv.Atoi(str); err == nil {
		d = i
	}
	t = hours + minutes
	return
}

// DateToAmadeusDate converts time to string ddmmyy
func DateToAmadeusDate(date time.Time) string {
	year := strconv.Itoa(date.Year())
	return fmt.Sprintf("%02d%02d%s", date.Day(), date.Month(), year[2:4])
}

// DateToFormatDDMMMYY convert time to 01JAN96
func DateToFormatDDMMMYY(date time.Time) string {
	month, year := "", strconv.Itoa(date.Year())
	switch date.Month() {
	case time.January:
		month = "JAN"
	case time.February:
		month = "FEB"
	case time.March:
		month = "MAR"
	case time.April:
		month = "APR"
	case time.May:
		month = "MAY"
	case time.June:
		month = "JUN"
	case time.July:
		month = "JUL"
	case time.August:
		month = "AUG"
	case time.September:
		month = "SEP"
	case time.October:
		month = "OCT"
	case time.November:
		month = "NOV"
	case time.December:
		month = "DEC"
	}
	return fmt.Sprintf("%02d%s%s", date.Day(), month, year[2:4])
}

// ParseFormatDDMMYY converts time to string
func ParseFormatDDMMYY(str string) (time.Time, error) {
	if str == "" {
		return time.Time{}, errors.New("empty string")
	}
	return fmtdate.Parse("DDMMYY", str)
}

// UpdateTimeFormatHHMM set time with str
func UpdateTimeFormatHHMM(date time.Time, str string) (time.Time, error) {
	if str == "" {
		return date, errors.New("empty string")
	}
	result, err := fmtdate.Parse("hhmm", str)
	if err != nil {
		return date, err
	}
	return time.Date(date.Year(), date.Month(), date.Day(), result.Hour(), result.Minute(), 0, 0, time.UTC), nil
}

// DateToAmadeusTime returns time in 1504 format
func DateToAmadeusTime(date time.Time) string {
	return fmt.Sprintf("%02d%02d", date.Hour(), date.Minute())
}

// DateToAmadeusTimeInt32 converts time to 1504 int32
func DateToAmadeusTimeInt32(date time.Time) int32 {
	var hours = int32(date.Hour()) * 100
	var minutes = int32(date.Minute())
	return hours + minutes
}

// TimeToToAmadeusTimeFloat64 converts time to 1504 float64
func TimeToToAmadeusTimeFloat64(date time.Time) float64 {
	var hours = int32(date.Hour()) * 100
	var minutes = int32(date.Minute())
	return float64(hours + minutes)
}

// ParseDateFormatISO8601 returns time from ISO8601 string
func ParseDateFormatISO8601(str string) (time.Time, error) {
	switch len(str) {
	case 16:
		return time.Parse("2006-01-02T15:04", str)
	case 19:
		return time.Parse("2006-01-02T15:04:05", str)
	}
	var date time.Time
	return date, errors.New("invalid format date")
}

// ParseDateFormatDDMMYYYY returns time from "DD.MM.YYYY" format
func ParseDateFormatDDMMYYYY(str string) (time.Time, error) {
	return fmtdate.Parse("DD.MM.YYYY", str)
}

// ParseDateFormatDDMMYYYY2 returns time from "DDMMYYYY" format
func ParseDateFormatDDMMYYYY2(str string) (time.Time, error) {
	return fmtdate.Parse("DDMMYYYY", str)
}

// ParseDateFormatDDMMYY returns time from "DDMMYY" format
func ParseDateFormatDDMMYY(str string) (time.Time, error) {
	return fmtdate.Parse("DDMMYY", str)
}

// DateStrToTime attempts convert string to time in different formates
func DateStrToTime(dateStr string) (time.Time, error) {
	var result time.Time
	switch len(dateStr) {
	case 10:
		date, err := fmtdate.Parse("YYYY-MM-DD", dateStr)
		if err != nil {
			date, err = fmtdate.Parse("YYYY.MM.DD", dateStr)
			if err != nil {
				date, err = time.Parse("02.01.2006", dateStr)
				if err != nil {
					return result, err
				}
			}
		}
		return date, nil
	case 15:
		date, err := fmtdate.Parse("YYYY-MM-DDhh:mm", dateStr)
		if err != nil {
			date, err = fmtdate.Parse("YYYY.MM.DDhh:mm", dateStr)
			if err != nil {
				return result, err
			}
		}
		return date, nil
	case 16:
		date, err := fmtdate.Parse("YYYY-MM-DDThh:mm", dateStr)
		if err != nil {
			date, err = fmtdate.Parse("YYYY.MM.DDThh:mm", dateStr)
			if err != nil {
				return result, err
			}
		}
		return date, nil
	case 19:
		date, err := fmtdate.Parse("YYYY-MM-DDThh:mm:ss", dateStr)
		if err != nil {
			date, err = fmtdate.Parse("YYYY.MM.DDThh:mm:ss", dateStr)
			if err != nil {
				return result, err
			}
		}
		return date, nil
	default:
		return result, errors.New("unknown date format")
	}
}
