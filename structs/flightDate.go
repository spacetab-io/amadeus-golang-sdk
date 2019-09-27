package structs

import (
	"errors"
	"time"
)

// FlightDateTime structure
type FlightDateTime struct {
	DateStr  string
	Date     time.Time
	TimeFlag bool
	Error    error
}

const (
	formatDateTimeHHMMSS = "2006-01-02T15:04:05"
	formatDateTimeHHMM   = "2006-01-02T15:04"
	formatDate           = "2006-01-02"
	formatTimeHHMMSS     = "15:04:05"
	formatTimeHHMM       = "15:04"

	errMsgEmptyValue    = "empty value"
	errMsgInvalidFormat = "invalid format date"
)

// ParseRequest parse FlightDateTime request
func (datetime *FlightDateTime) ParseRequest() error {
	datetime.TimeFlag = false
	if datetime.DateStr == "" {
		datetime.Error = errors.New(errMsgEmptyValue)
		return datetime.Error
	}
	switch len(datetime.DateStr) {
	case 25:
		result, err := time.Parse(time.RFC3339, datetime.DateStr)
		if err != nil {
			datetime.Error = err
			return datetime.Error
		}

		datetime.Date = result
		datetime.TimeFlag = true
		return nil

	case 19:
		result, err := time.Parse(formatDateTimeHHMMSS, datetime.DateStr)
		if err != nil {
			datetime.Error = err
			return datetime.Error
		}

		datetime.Date = result
		datetime.TimeFlag = true
		return nil

	case 16:
		result, err := time.Parse(formatDateTimeHHMM, datetime.DateStr)
		if err != nil {
			datetime.Error = err
			return datetime.Error
		}

		datetime.Date = result
		datetime.TimeFlag = true
		return nil

	case 10:
		result, err := time.Parse(formatDate, datetime.DateStr)
		if err != nil {
			datetime.Error = err
			return datetime.Error
		}

		datetime.Date = result
		//datetime.TimeFlag = false
		return nil

	}
	datetime.Error = errors.New(errMsgInvalidFormat)
	return datetime.Error
}

// FlightDateTimeUniversal structure
type FlightDateTimeUniversal struct {
	DateStr  string
	Date     time.Time
	DateFlag bool
	TimeFlag bool
	Error    error
}

// ParseRequest parse FlightDateTimeUniversal
func (date *FlightDateTimeUniversal) ParseRequest() error {
	date.DateFlag = false
	date.TimeFlag = false
	if date.DateStr == "" {
		date.Error = errors.New(errMsgEmptyValue)
		return date.Error
	}
	switch len(date.DateStr) {
	case 25:
		result, err := time.Parse(time.RFC3339, date.DateStr)
		if err != nil {
			date.Error = err
			return date.Error
		}

		date.Date = result
		date.DateFlag = true
		date.TimeFlag = true
		return nil

	case 19:
		result, err := time.Parse(formatDateTimeHHMMSS, date.DateStr)
		if err != nil {
			date.Error = err
			return date.Error
		}

		date.Date = result
		date.DateFlag = true
		date.TimeFlag = true
		return nil

	case 16:
		result, err := time.Parse(formatDateTimeHHMM, date.DateStr)
		if err != nil {
			date.Error = err
			return date.Error
		}

		date.Date = result
		date.DateFlag = true
		date.TimeFlag = true
		return nil

	case 10:
		result, err := time.Parse(formatDate, date.DateStr)
		if err == nil {
			date.Error = err
			return date.Error
		}

		date.Date = result
		date.DateFlag = true
		//date.TimeFlag = false
		return nil

	case 8:
		result, err := time.Parse(formatTimeHHMMSS, date.DateStr)
		if err != nil {
			date.Error = err
			return date.Error
		}

		date.Date = result
		//date.DateFlag = false
		date.TimeFlag = true
		return nil

	case 5:
		result, err := time.Parse(formatTimeHHMM, date.DateStr)
		if err != nil {
			date.Error = err
			return date.Error
		}

		date.Date = result
		//date.DateFlag = false
		date.TimeFlag = true
		return nil

	}
	date.Error = errors.New(errMsgInvalidFormat)
	return date.Error
}
