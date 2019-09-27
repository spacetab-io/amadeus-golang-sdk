package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tmconsulting/amadeus-golang-sdk/structs"
	"strings"
	"time"
)

/*
{
	"request_id": "hashcode",
	"client_data": {
		"amadeus_office_id": ""  // officeId, для которого нужно запрашивать сессию amadeus
	},
	"passengers_count": {
		"ADT": 1,
		"CHD": 0,
		"INF": 0
	},
	"allow_mixing": true,  //allow the mixing of classes in segments
	"search_only_s7": false,
	"base_сlass": [],  // ["E", "B", "F"] где E - Economic, B - Business, F - First
	"changes": true,  // допустимы ли рейсы с пересадками
	"airline": [],
	"excluded_airlines": [],
	"with_baggage": false,  // eсли оно true - в выдаче должны быть только те рекомендации, где присутствует возможность провоза багажа
	"is_domestic": false,  // признак домашнего перелёта (Вылетаем ли за пределы страны, домашней для аэропорта вылета)
	"travel_type": "OW",  // ["OW", "RT", "CT"] где OW - One Way, RT - Round Trip, CT - Complex Trip
	"itineraries": {
		"{itinerary_id}": {
			"departure_location": {
				"airport_code": "MSQ",
				"country_code": "BY",
				"city_code": "MSQ"
			},
			"arrival_location": {
				"airport_code": "JFK",
				"country_code": "US",
				"city_code": "NYC"
			},
			"departure_date":  "YYYY-MM-DDThh:mm:ss",
			"departure_date_till": "YYYY-MM-DDThh:mm:ss",
			"arrival_date": "YYYY-MM-DDThh:mm:ss"
		}
	},
	"currency": "RUB"
}
*/

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
	var officeID = request.ClientData.OfficeID
	if officeID == "" {
		err = errors.New("empty OfficeID")
		fmt.Println(err)
		return
	}

	var routeCount = len(request.Itineraries)
	if routeCount == 0 {
		err = errors.New("empty Itineraries")
		fmt.Println(err)
		return
	}

	var travelType = strings.ToUpper(request.TravelType)
	if travelType == "" {
		err = errors.New("empty TravelType")
		fmt.Println(err)
		return
	}

	if travelType == "OW" && routeCount != 1 {
		err = errors.New("number of Itineraries for 'OW' should be equal to one")
		fmt.Println(err)
		return
	}

	if travelType == "RT" && routeCount != 2 {
		err = errors.New("number of Itineraries for 'RT' should be equal to two")
		fmt.Println(err)
		return
	}

	if travelType == "CT" && routeCount < 2 {
		err = errors.New("use TravelType == 'OW'")
		fmt.Println(err)
		return
	}

	request.TravelType = travelType

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

/*
{
	"segments": {
		"amadeus_1": {
			"departure": {
				"airport_code": "DME",
				"terminal": ""
			},
			"arrival": {
				"airport_code": "AER",
				"terminal": ""
			},
			"departure_date": "2018-05-21T09:40",
			"arrival_date": "2018-05-21T12:00",
			"operating_airline_flight_number": "1045",
			//"marketing_airline_flight_number": "1045",
			"aircraft": "738",
			//"aircraft_full_name": "Boeing 737-800 Passenger",
			"marketing_airline": {
				"code_eng": "S7"
			},
			"operating_airline": {
				"code_eng": "GH",
			},
			"codeshare": true,
			"travel_time": 140
		}
	},
	"recommendations": [
		{
			"provider": "amadeus",
			"class": ["E"],
			"price": {
				"total_amount": {
					"amount": 0,
					"currency": "RUB"
				},
				"fare_amount": {
					"amount": 0,
					"currency": "RUB"
				},
				"tax_amount": {
					"amount": 0,
					"currency": "RUB"
				},
				"fee_amount": {
					"amount": 0,
					"currency": "RUB"
				},
				"pax_fare": {
					"ADT": {
						"total_amount": {
							"amount": 0,
							"currency": "RUB"
						},
						"fare_amount": {
							"amount": 0,
							"currency": "RUB"
						},
						"tax_amount": {
							"amount": 0,
							"currency": "RUB"
						},
						"fee_amount": {
							"amount": 0,
							"currency": "RUB"
						}
					},
					"CHD": null, //same as ADT
					"INF": null  //same as ADT
				},
				"taxes": [
					{
						"costs": {
							"currency": "RUB",
							"amount": 0
						},
						"code": ""
					}
				]
			},
			"itineraries_segments": {  //routes_segments по-русски.
				"{itinerary_id}": {  // ссылка на itinerary_id из request
					"{index}": {
						"segments": [
							{
								"segment_id": "amadeus_1",  // ссылка на id из segments
								"codeshare": true,

								// flightDetails["ADT"]
								"class": "E",
								"flight_class": "M", // (Amadeus)
								"booking_code": "Q",
								"fare_basis_code": "QBSOW",
								"availability": 9,
								"baggage": {
									"value": 0,
									"unit": ""
								},

								"pax_details": {  // только для случаев, если "fare_basis_code" отличается от "ADT"
									"CHD": {
										//"class": "E",
										"flight_class": "M", // (Amadeus)
										"booking_code": "Q",
										"fare_basis_code": "QBSOW",
										"availability": 9
										"baggage": {  // только для случаев, если "baggage" отличается от "ADT"
											"value": 0,
											"unit": ""
										}
									},
									"INF": null  //same as CHD
								}
							}
						],
						"travel_time": 140,
					}
				}
			},
			"validating_airline": {
				"code_eng": "UT",
			},
			"interline": true,
			"last_date": "2018-07-15T00:00:00",
			"refundable": "N"
		}
	]
}
*/

// SearchResponse structure of Search response
type SearchResponse struct {
	Segments        MapFlights          `json:"segments,omitempty"`
	Recommendations []*Recommendation   `json:"recommendations,omitempty"`
	Error           *structs.ErrorReply `json:"error,omitempty"`
}

// MapFlights Map of Flights
type MapFlights map[string]*structs.Flight

// Recommendation structure
type Recommendation struct {
	ID                string          `json:"id,omitempty"`
	Provider          string          `json:"provider"`
	Class             []string        `json:"class"`
	Price             structs.Price   `json:"price"`
	ItinerarySegments RoutesSegments  `json:"itinerary_segments"`
	ValidatingAirline structs.Airline `json:"validating_airline"`
	Interline         bool            `json:"interline"`
	LastTicketDate    string          `json:"last_date,omitempty"`
	Refundable        string          `json:"refundable,omitempty"`
}

// Combine compare recomendations
func (recommendation *Recommendation) Combine(recommendation2 *Recommendation) bool {
	if !recommendation.Price.Equal(&recommendation2.Price) {
		return false
	}
	if recommendation.ValidatingAirline.CodeEng != recommendation2.ValidatingAirline.CodeEng {
		return false
	}
	for itineraryID, variants := range recommendation.ItinerarySegments {
		var variantNumber = len(variants) + 1
		for _, segments := range recommendation2.ItinerarySegments[itineraryID] {
			variants[variantNumber] = segments
			variantNumber++
		}
	}
	return true
}

// RoutesSegments map of RoutesSegments
type RoutesSegments map[int]map[int]*RouteSegments

// RouteSegments Route Segments structure
type RouteSegments struct {
	GroupOfSegmentsID int                 `json:"-"`
	ItinerarySegments []*ItinerarySegment `json:"segments"`
	TravelTime        int                 `json:"travel_time"`
}

// ItinerarySegment segments of Itinerary
type ItinerarySegment struct {
	SegmentID string `json:"segment_id"`
	CodeShare bool   `json:"codeshare"`
	*FlightDetail
	PaxDetails mapPaxDetails `json:"pax_details,omitempty"`
}

type mapPaxDetails map[string]*FlightDetail

// FlightDetail Flight Detail structure
type FlightDetail struct {
	Class         string               `json:"class"` // @TODO Deprecate it!
	Cabin         string               `json:"cabin"`
	FareClass     string               `json:"fare_class"`
	BookingClass  string               `json:"rbd"`
	FareBasisCode string               `json:"fare_basis_code"`
	Availability  int                  `json:"availability"`
	Baggage       *structs.BaggageType `json:"baggage,omitempty"`
}

// Equal compare flight details
func (fd *FlightDetail) Equal(test *FlightDetail) bool {
	if fd.BookingClass == test.BookingClass &&
		fd.FareClass == test.FareClass &&
		fd.FareBasisCode == test.FareBasisCode &&
		fd.Availability == test.Availability {
		return true
	}
	return false
}
