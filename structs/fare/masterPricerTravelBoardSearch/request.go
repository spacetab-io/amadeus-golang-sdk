package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs"
	"strings"
	"time"
)

// SearchRequest structure of Search request
type SearchRequest struct {
	RequestID        string             `json:"request_id"`
	ClientData       structs.ClientInfo `json:"client_data"`
	Passengers       structs.Travellers `json:"passengers_count"`
	AllowMixing      *bool              `json:"allow_mixing,omitempty"`
	ResultOnlyS7     bool               `json:"search_only_s7,omitempty"`
	BaseClass        []string           `json:"base_class,omitempty"` //array_enum("C", "Y", "F")
	Changes          bool               `json:"changes,omitempty"`
	Airlines         []string           `json:"airlines,omitempty"` // Массив IATA кодов авиакомпаний по которым нужно фильтровать результаты поиска
	ExcludedAirlines []string           `json:"excluded_airlines,omitempty"`
	WithBaggage      bool               `json:"with_baggage,omitempty"`
	IsDomestic       bool               `json:"is_domestic,omitempty"`
	TravelType       string             `json:"travel_type"` //enum("OW", "RT", "CT") где "OW" - One Way, "RT" - Round Trip, "CT" - Complex Trip
	Itineraries      mapItineraries     `json:"itineraries"`
	Currency         string             `json:"currency"`
}

type mapItineraries map[int]*Itinerary

// Check validate search request
func (request *SearchRequest) Check() (err error) {

	err = request.ValidateParams()
	if err != nil {
		return
	}

	err = request.ValidateRoutes()
	if err != nil {
		return
	}

	err = request.ValidatePassengers()
	if err != nil {
		return
	}

	err = request.PrepareItineraries()
	if err != nil {
		return
	}

	return
}

// Itinerary of search request
type Itinerary struct {
	*structs.Itinerary
	//ItineraryID		string				`json:"-"`
	DepartureDate     structs.FlightDateTime
	DepartureDateTill structs.FlightDateTimeUniversal
	ArrivalDate       structs.FlightDateTime
}

// UnmarshalJSON overrides UnmarshalJSON
func (i *Itinerary) UnmarshalJSON(data []byte) error {
	type ItineraryRequest struct {
		*structs.Itinerary
		//ItineraryID			string	`json:"itinerary_id,omitempty"`
		DepartureDate     string `json:"departure_date"`
		DepartureDateTill string `json:"departure_date_till,omitempty"`
		ArrivalDate       string `json:"arrival_date,omitempty"`
	}
	var request ItineraryRequest
	if err := json.Unmarshal(data, &request); err != nil {
		//logs.Log.WithError(err).Error("Error unmarshal json") //todo log
		return err
	}

	i.Itinerary = request.Itinerary
	//i.ItineraryID               = request.ItineraryID
	i.DepartureDate.DateStr = request.DepartureDate
	errDep := i.DepartureDate.ParseRequest()
	i.DepartureDateTill.DateStr = request.DepartureDateTill
	errDepTill := i.DepartureDateTill.ParseRequest()
	i.ArrivalDate.DateStr = request.ArrivalDate
	errArrival := i.ArrivalDate.ParseRequest()
	if errDep != nil && (errArrival != nil || errDepTill == nil) {
		return errors.New("empty 'departure_date'")
	}
	return nil
}

func (request *SearchRequest) ValidateParams() (err error) {

	if request.ClientData.OfficeID == "" {
		err = errors.New("empty OfficeID")
		return
	}

	if len(request.Itineraries) == 0 {
		err = errors.New("empty Itineraries")
		return
	}

	if request.TravelType == "" {
		err = errors.New("empty TravelType")
		return
	}

	return nil
}

func (request *SearchRequest) ValidateRoutes() (err error) {

	var travelType = strings.ToUpper(request.TravelType)
	var routeCount = len(request.Itineraries)

	request.TravelType = travelType

	if travelType == "OW" && routeCount != 1 {
		err = errors.New("number of Itineraries for 'OW' should be equal to one")
		return
	}

	if travelType == "RT" && routeCount != 2 {
		err = errors.New("number of Itineraries for 'RT' should be equal to two")
		return
	}

	if travelType == "CT" && routeCount < 2 {
		err = errors.New("use TravelType == 'OW'")
		return
	}

	return nil
}

func (request *SearchRequest) ValidatePassengers() (err error) {
	if (request.Passengers.ADT + request.Passengers.CHD) <= 0 {
		err = errors.New("invalid count passengers")
		return
	}
	if request.Passengers.INF < 0 {
		request.Passengers.INF = 0
	}
	if request.Passengers.ADT < request.Passengers.INF {
		err = errors.New("number of infants exceeds the number of adults")
		return
	}

	return nil
}

func (request *SearchRequest) PrepareItineraries() (err error) {
	//for id, itinerary := range request.Itineraries {
	for i := 1; i <= len(request.Itineraries); i++ {
		itinerary, exists := request.Itineraries[i]
		if !exists {
			err = fmt.Errorf("not found 'itinerary_id' equal to '%d'", i)
			return
		}
		errDep := itinerary.DepartureDate.ParseRequest()
		errDepTill := itinerary.DepartureDateTill.ParseRequest()
		errArrival := itinerary.ArrivalDate.ParseRequest()
		if errDep != nil {
			if errDepTill == nil {
				if itinerary.DepartureDateTill.DateFlag {
					departureDateStr := itinerary.DepartureDateTill.Date.Format("2006-01-02") + "T00:00"
					if departureDate, err := time.Parse("2006-01-02T15:04", departureDateStr); err == nil {
						itinerary.DepartureDate.DateStr = departureDateStr
						itinerary.DepartureDate.Date = departureDate
						itinerary.DepartureDate.TimeFlag = true
						itinerary.DepartureDate.Error = nil
						errDep = nil
					}
				}
			}
			if errDep != nil {
				err = fmt.Errorf("no valid travel date for itinerary ('itinerary_id' equal to '%d')", i)
				return
			}
		}

		if !itinerary.DepartureDate.TimeFlag {
			itinerary.DepartureDate.DateStr = itinerary.DepartureDate.Date.Format("2006-01-02") + "T00:00"
			itinerary.DepartureDate.TimeFlag = true
		}

		if errDepTill == nil {
			if !itinerary.DepartureDateTill.DateFlag {
				departureDateTillStr := itinerary.DepartureDate.Date.Format("2006-01-02") + "T" + itinerary.DepartureDateTill.Date.Format("15:04")
				if departureDateTill, err := time.Parse("2006-01-02T15:04", departureDateTillStr); err == nil {
					itinerary.DepartureDateTill.DateStr = departureDateTillStr
					itinerary.DepartureDateTill.Date = departureDateTill
					itinerary.DepartureDateTill.DateFlag = true
				} else {
					//itinerary.DepartureDateTill.DateFlag = false
					itinerary.DepartureDateTill.TimeFlag = false
					itinerary.DepartureDateTill.Error = err
				}
			} else if !itinerary.DepartureDateTill.TimeFlag {
				departureDateTillStr := itinerary.DepartureDateTill.Date.Format("2006-01-02") + "T23:59"
				if departureDateTill, err := time.Parse("2006-01-02T15:04", departureDateTillStr); err == nil {
					itinerary.DepartureDateTill.DateStr = departureDateTillStr
					itinerary.DepartureDateTill.Date = departureDateTill
					itinerary.DepartureDateTill.TimeFlag = true
				} else {
					itinerary.DepartureDateTill.DateFlag = false
					//itinerary.DepartureDateTill.TimeFlag = false
					itinerary.DepartureDateTill.Error = err
				}
			}
		} // else {
		//	departureDateTillStr := itinerary.DepartureDate.Date.Format("2006-01-02") + "T23:59"
		//	if departureDateTill, err := time.Parse("2006-01-02T15:04", departureDateTillStr); err == nil {
		//		itinerary.DepartureDateTill.DateStr = departureDateTillStr
		//		itinerary.DepartureDateTill.Date = departureDateTill
		//		itinerary.DepartureDateTill.DateFlag = true
		//	}
		//}

		if errArrival == nil {
			if !itinerary.ArrivalDate.TimeFlag {
				arrivalDateStr := itinerary.ArrivalDate.Date.Format("2006-01-02") + "T23:59"
				if arrivalDate, err := time.Parse("2006-01-02T15:04", arrivalDateStr); err == nil {
					itinerary.ArrivalDate.DateStr = arrivalDateStr
					itinerary.ArrivalDate.Date = arrivalDate
					itinerary.ArrivalDate.TimeFlag = true
				} else {
					itinerary.ArrivalDate.TimeFlag = false
					itinerary.ArrivalDate.Error = err
				}
			}
		} else {
			arrivalDateStr := itinerary.DepartureDate.Date.Format("2006-01-02")
			if arrivalDate, err := time.Parse("2006-01-02", arrivalDateStr); err == nil {
				itinerary.ArrivalDate.DateStr = arrivalDateStr
				itinerary.ArrivalDate.Date = arrivalDate
				itinerary.ArrivalDate.TimeFlag = false
			} else {
				itinerary.ArrivalDate.TimeFlag = false
				itinerary.ArrivalDate.Error = err
			}
		}
	}

	return nil
}
