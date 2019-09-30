package search

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmconsulting/amadeus-golang-sdk/structs"
	"testing"
	"time"
)

func TestValidateParams(t *testing.T) {
	t.Run("Check Validate", func(t *testing.T) {
		request := SearchRequest{
			ClientData: structs.ClientInfo{
				OfficeID: "",
			},
		}

		err := request.ValidateParams()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})

	t.Run("Check Validate", func(t *testing.T) {
		request := SearchRequest{
			ClientData: structs.ClientInfo{
				OfficeID: "MOWR228FG",
			},
		}

		err := request.ValidateParams()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})

	t.Run("Check Validate", func(t *testing.T) {
		itinerary := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "SVO",
				CityCode:    "MOW",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "LED",
				CityCode:    "LED",
				CountryCode: "RU",
				Type:        "city",
			},
		}

		request := SearchRequest{
			ClientData: structs.ClientInfo{
				OfficeID: "MOWR228FG",
			},
			Itineraries: map[int]*Itinerary{
				1: {
					Itinerary: &itinerary,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(10 * 24 * time.Hour), // add 10 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(11 * 24 * time.Hour).String(), // add 11 days
					},
				},
			},
		}

		err := request.ValidateParams()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})

	t.Run("Check Validate", func(t *testing.T) {
		itinerary := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "SVO",
				CityCode:    "MOW",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "LED",
				CityCode:    "LED",
				CountryCode: "RU",
				Type:        "city",
			},
		}

		request := SearchRequest{
			ClientData: structs.ClientInfo{
				OfficeID: "MOWR228FG",
			},
			Itineraries: map[int]*Itinerary{
				1: {
					Itinerary: &itinerary,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(10 * 24 * time.Hour), // add 10 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(11 * 24 * time.Hour).String(), // add 10 days
					},
				},
			},
			TravelType: "OW",
		}

		err := request.ValidateParams()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
}

func TestValidateRoutes(t *testing.T) {
	t.Run("Validate Routes", func(t *testing.T) {
		itinerary := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "SVO",
				CityCode:    "MOW",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "LED",
				CityCode:    "LED",
				CountryCode: "RU",
				Type:        "city",
			},
		}

		request := SearchRequest{
			Itineraries: map[int]*Itinerary{
				1: {
					Itinerary: &itinerary,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(10 * 24 * time.Hour), // add 10 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(11 * 24 * time.Hour).String(), // add 10 days
					},
				},
			},
			TravelType: "OW",
		}

		err := request.ValidateRoutes()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})

	t.Run("Validate Routes", func(t *testing.T) {
		itinerary := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "SVO",
				CityCode:    "MOW",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "LED",
				CityCode:    "LED",
				CountryCode: "RU",
				Type:        "city",
			},
		}

		itinerary2 := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "LED",
				CityCode:    "LED",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "MOW",
				CityCode:    "MOW",
				CountryCode: "RU",
				Type:        "city",
			},
		}

		request := SearchRequest{
			Itineraries: map[int]*Itinerary{
				1: {
					Itinerary: &itinerary,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(10 * 24 * time.Hour), // add 10 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(11 * 24 * time.Hour).String(), // add 10 days
					},
				},
				2: {
					Itinerary: &itinerary2,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(20 * 24 * time.Hour), // add 20 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(21 * 24 * time.Hour).String(), // add 21 days
					},
				},
			},
			TravelType: "RT",
		}

		err := request.ValidateRoutes()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})

	t.Run("Validate Routes", func(t *testing.T) {
		itinerary := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "SVO",
				CityCode:    "MOW",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "LED",
				CityCode:    "LED",
				CountryCode: "RU",
				Type:        "city",
			},
		}

		itinerary2 := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "LED",
				CityCode:    "LED",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "GOJ",
				CityCode:    "GOJ",
				CountryCode: "RU",
				Type:        "city",
			},
		}

		itinerary3 := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "GOJ",
				CityCode:    "GOJ",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "MOW",
				CityCode:    "MOW",
				CountryCode: "RU",
				Type:        "city",
			},
		}

		request := SearchRequest{
			Itineraries: map[int]*Itinerary{
				1: {
					Itinerary: &itinerary,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(10 * 24 * time.Hour), // add 10 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(11 * 24 * time.Hour).String(), // add 10 days
					},
				},
				2: {
					Itinerary: &itinerary2,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(20 * 24 * time.Hour), // add 20 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(21 * 24 * time.Hour).String(), // add 21 days
					},
				},
				3: {
					Itinerary: &itinerary3,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(20 * 24 * time.Hour), // add 20 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(21 * 24 * time.Hour).String(), // add 21 days
					},
				},
			},
			TravelType: "CT",
		}

		err := request.ValidateRoutes()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
}

func TestValidatePassengers(t *testing.T) {
	t.Run("Validate Passengers", func(t *testing.T) {
		request := SearchRequest{
			Passengers: structs.Travellers{ADT: 2, CHD: 2, INF: 2},
		}

		err := request.ValidatePassengers()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})

	t.Run("Validate Passengers", func(t *testing.T) {
		request := SearchRequest{
			Passengers: structs.Travellers{ADT: 1, CHD: 0, INF: 2},
		}

		err := request.ValidatePassengers()
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}
