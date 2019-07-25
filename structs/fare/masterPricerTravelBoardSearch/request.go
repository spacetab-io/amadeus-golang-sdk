package search

import (
	"time"
)

type Request struct {
	RequestId        string                `json:"request_id"`
	ClientData       ClientInfo            `json:"client_data"`
	Passengers       Travellers            `json:"passengers_count"`
	AllowMixing      *bool                 `json:"allow_mixing,omitempty"`
	BaseClass        []string              `json:"base_class,omitempty"` //array_enum("", "e", "b", "f")
	Changes          *bool                 `json:"changes,omitempty"`
	Airlines         []string              `json:"airlines,omitempty"` // Массив IATA кодов авиакомпаний по которым нужно фильтровать результаты поиска
	ExcludedAirlines []string              `json:"excluded_airlines,omitempty"`
	WithBaggage      bool                  `json:"with_baggage,omitempty"`
	IsDomestic       bool                  `json:"is_domestic,omitempty"`
	TravelType       string                `json:"travel_type"` //enum("OW", "RT", "CT") где "OW" - One Way, "RT" - Round Trip, "CT" - Complex Trip
	Itineraries      map[string]*Itinerary `json:"itineraries"`
	Currency         string                `json:"currency"`
}

type Itinerary struct {
	DepartureLocation string
	ArrivalLocation   string
	DepartureDate     time.Time
	DepartureDateTill *time.Time
}

type ClientInfo struct {
	OfficeId string `json:"office_id"`
	Timezone string `json:"timezone"`
}

type Travellers struct {
	ADT int `json:"ADT,omitempty"`
	CHD int `json:"CHD,omitempty"`
	INF int `json:"INF,omitempty"`
}
