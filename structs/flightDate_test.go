package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRequest(t *testing.T) {
	t.Run("ParseRequest", func(t *testing.T) {
		request := FlightDateTimeUniversal{
			DateStr: "2015-10-20",
		}

		err := request.ParseRequest()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})

	t.Run("ParseRequest", func(t *testing.T) {
		request := FlightDateTimeUniversal{
			DateStr: "2015-10-20T18:40:02",
		}

		err := request.ParseRequest()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})

	t.Run("ParseRequest", func(t *testing.T) {
		request := FlightDateTimeUniversal{
			DateStr: "2015-10-20T18:40",
		}

		err := request.ParseRequest()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})

	t.Run("ParseRequest", func(t *testing.T) {
		request := FlightDateTimeUniversal{
			DateStr: "18:40",
		}

		err := request.ParseRequest()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
}
