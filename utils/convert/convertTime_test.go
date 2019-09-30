package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAmadeusTimeToMinutes(t *testing.T) {
	t.Run("AmadeusTimeToMinutes", func(t *testing.T) {
		minutes := AmadeusTimeToMinutes("0040")
		assert.Equal(t, 40, minutes)
	})

	t.Run("AmadeusTimeToMinutes", func(t *testing.T) {
		minutes := AmadeusTimeToMinutes("0120")
		assert.Equal(t, 80, minutes)
	})

	t.Run("AmadeusTimeToMinutes", func(t *testing.T) {
		minutes := AmadeusTimeToMinutes("1020")
		assert.Equal(t, 620, minutes)
	})
}

func TestAmadeusDateConvert(t *testing.T) {
	t.Run("AmadeusDateConvert", func(t *testing.T) {
		convertedTime, err := AmadeusDateConvert("01JAN95")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		expectedTime := time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC)

		assert.Equal(t, expectedTime, convertedTime)
	})
}

func TestAmadeusDateTimeConvert(t *testing.T) {
	t.Run("AmadeusDateConvert", func(t *testing.T) {
		convertedTime := AmadeusDateTimeConvert("201095", "1840")

		expectedTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)

		assert.Equal(t, expectedTime, convertedTime)
	})
}

func TestDateTimeToInt(t *testing.T) {
	t.Run("AmadeusDateConvert", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate, convertedTime := DateTimeToInt(etalonTime)

		assert.Equal(t, 201095, convertedDate)
		assert.Equal(t, 1840, convertedTime)
	})
}

func TestDateToAmadeusDate(t *testing.T) {
	t.Run("DateToAmadeusDate", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate := DateToAmadeusDate(etalonTime)

		assert.Equal(t, "201095", convertedDate)
	})
}

func TestDateToFormatDDMMMYY(t *testing.T) {
	t.Run("DateToFormatDDMMMYY", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate := DateToFormatDDMMMYY(etalonTime)

		assert.Equal(t, "20OCT95", convertedDate)
	})
}

func TestParseFormatDDMMYY(t *testing.T) {
	t.Run("ParseFormatDDMMYY", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 0, 0, 0, 0, time.UTC)
		convertedDate, err := ParseFormatDDMMYY("201095")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})
}

func TestUpdateTimeFormatHHMM(t *testing.T) {
	t.Run("UpdateTimeFormatHHMM", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 0, 0, 0, 0, time.UTC)
		expectedTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate, err := UpdateTimeFormatHHMM(etalonTime, "1840")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, expectedTime, convertedDate)
	})
}

func TestDateToAmadeusTime(t *testing.T) {
	t.Run("DateToAmadeusTime", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate := DateToAmadeusTime(etalonTime)

		assert.Equal(t, "1840", convertedDate)
	})
}

func TestDateToAmadeusTimeInt32(t *testing.T) {
	t.Run("DateToAmadeusTimeInt32", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate := DateToAmadeusTimeInt32(etalonTime)

		assert.Equal(t, int32(1840), convertedDate)
	})
}

func TestTimeToToAmadeusTimeFloat64(t *testing.T) {
	t.Run("TimeToToAmadeusTimeFloat64", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate := TimeToToAmadeusTimeFloat64(etalonTime)

		assert.Equal(t, float64(1840), convertedDate)
	})
}

func TestParseDateFormatISO8601(t *testing.T) {
	t.Run("ParseDateFormatISO8601", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate, err := ParseDateFormatISO8601("1995-10-20T18:40")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})
}

func TestParseDateFormatDDMMYYYY(t *testing.T) {
	t.Run("ParseDateFormatDDMMYYYY", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 0, 0, 0, 0, time.UTC)
		convertedDate, err := ParseDateFormatDDMMYYYY("20.10.1995")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})
}

func TestParseDateFormatDDMMYYYY2(t *testing.T) {
	t.Run("ParseDateFormatDDMMYYYY2", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 0, 0, 0, 0, time.UTC)
		convertedDate, err := ParseDateFormatDDMMYYYY2("20101995")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})
}

func TestParseDateFormatDDMMYY(t *testing.T) {
	t.Run("ParseDateFormatDDMMYY", func(t *testing.T) {
		etalonTime := time.Date(1995, 10, 20, 0, 0, 0, 0, time.UTC)
		convertedDate, err := ParseDateFormatDDMMYY("201095")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})
}

func TestDateStrToTime(t *testing.T) {
	t.Run("DateStrToTime", func(t *testing.T) {
		etalonTime := time.Date(2015, 10, 20, 0, 0, 0, 0, time.UTC)
		convertedDate, err := DateStrToTime("2015-10-20")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})

	t.Run("DateStrToTime", func(t *testing.T) {
		etalonTime := time.Date(2015, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate, err := DateStrToTime("2015-10-2018:40")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})

	t.Run("DateStrToTime", func(t *testing.T) {
		etalonTime := time.Date(2015, 10, 20, 18, 40, 0, 0, time.UTC)
		convertedDate, err := DateStrToTime("2015-10-20T18:40")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})

	t.Run("DateStrToTime", func(t *testing.T) {
		etalonTime := time.Date(2015, 10, 20, 18, 40, 23, 0, time.UTC)
		convertedDate, err := DateStrToTime("2015-10-20T18:40:23")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, etalonTime, convertedDate)
	})
}
