package PNR_Information

import (
	"time"
)

type Response struct {
	TravellesInfo []TravellerInfo
	Segments      []Segment
}

type TravellerInfo struct {
	FirstName   string
	LastName    string
	Type        PaxType // ADT, CHD, INF
	DateOfBirth time.Time
	Quaifier    string // PT, PI
	Number      int    // reference number
}

type PaxType string

type Segment struct {
	SegmentId         string    `json:"-"`
	FlightNumber      string    `json:"flight_number"`
	DepartureLocation Location  `json:"departure"`
	DepartureDate     time.Time `json:"departure_date"`
	ArrivalLocation   Location  `json:"arrival"`
	ArrivalDate       time.Time `json:"arrival_date"`
	MarketingAirline  *Airline  `json:"marketing_airline"`
	OperatingAirline  *Airline  `json:"operating_airline"`
	ValidatingAirline *Airline  `json:"validating_airline,omitempty"`
	Aircraft          *Aircraft `json:"aircraft,omitempty"`
}

type Location struct {
	AirportCode string `json:"airport_code"`
	CountryCode string `json:"country_code"`
	CityCode    string `json:"city_code"`
	Terminal    string `json:"terminal"`
	Type        string `json:"-"`
}

type Airline struct {
	CodeEng string `json:"code_eng" bson:"codeEng"`
	CodeRus string `json:"code_rus,omitempty" bson:"codeRus,omitempty"`
	NameEng string `json:"name_eng,omitempty" bson:"nameEng,omitempty"`
	NameRus string `json:"name_rus,omitempty" bson:"nameRus,omitempty"`
}

type Aircraft struct {
	Code *string           `json:"code" db:"code" bson:"code"`
	Name map[string]string `json:"name,omitempty" db:"name" bson:"name,omitempty"`
}
