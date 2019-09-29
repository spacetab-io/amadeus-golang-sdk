package structs

import (
	"encoding/json"
	"time"
)

// Flight structure
type Flight struct {
	SegmentID         string    `json:"-"`
	Aircraft          *Aircraft `json:"aircraft,omitempty"`
	FlightNumber      string    `json:"flight_number"`
	DepartureLocation Location  `json:"departure"`
	DepartureDate     time.Time `json:"departure_date"`
	ArrivalLocation   Location  `json:"arrival"`
	ArrivalDate       time.Time `json:"arrival_date"`
	MarketingAirline  *Airline  `json:"marketing_airline"`
	OperatingAirline  *Airline  `json:"operating_airline"`
	ValidatingAirline *Airline  `json:"validating_airline,omitempty"`
	TechStops         *int32    `json:"tech_stops,omitempty"`

	//TravelTime			int			`json:"travel_time,omitempty"`
	//ETicket			bool		`json:"eTicket"`
}

// MarshalJSON overrides MarshalJSON
func (flight *Flight) MarshalJSON() ([]byte, error) {
	//@TODO refactor: перевести на общую структуру
	return json.Marshal(struct {
		Aircraft *Aircraft `json:"aircraft"`
		//AircraftFullName	string		`json:"aircraft_full_name,omitempty"`
		FlightNumber      string    `json:"marketing_airline_flight_number"` //@TODO Найти и передать номер рейса оперирующей компании
		DepartureLocation *Location `json:"departure"`
		DepartureDate     string    `json:"departure_date"`
		ArrivalLocation   *Location `json:"arrival"`
		ArrivalDate       string    `json:"arrival_date"`
		MarketingAirline  *Airline  `json:"marketing_airline"`
		OperatingAirline  *Airline  `json:"operating_airline"`
		//ValidatingAirline	string		`json:"validating_airline"`
		//Codeshare			bool		`json:"codeshare"`
		//TravelTime			int			`json:"travel_time"`
	}{
		Aircraft: flight.Aircraft,
		//AircraftFullName:  flight.AircraftFullName,
		FlightNumber:      flight.FlightNumber,
		DepartureLocation: &flight.DepartureLocation,
		DepartureDate:     flight.DepartureDate.Format(TimeFormat),
		ArrivalLocation:   &flight.ArrivalLocation,
		ArrivalDate:       flight.ArrivalDate.Format(TimeFormat),
		MarketingAirline:  flight.MarketingAirline,
		OperatingAirline:  flight.OperatingAirline,
		//Codeshare:         flight.MarketingAirline != flight.OperatingAirline,
		//ValidatingAirline: flight.ValidatingAirline,
		//TravelTime:        flight.TravelTime,
	})
}

// UnmarshalJSON overrides UnmarshalJSON
func (flight *Flight) UnmarshalJSON(data []byte) error {
	type Flight2 struct {
		Aircraft          *Aircraft `json:"aircraft,omitempty"`
		FlightNumber      string    `json:"marketing_airline_flight_number"` // marketing_airline_flight_number должно быть такое же как MarshalJSON.FlightNumber
		DepartureLocation Location  `json:"departure"`
		DepartureDate     string    `json:"departure_date"`
		ArrivalLocation   Location  `json:"arrival"`
		ArrivalDate       string    `json:"arrival_date"`
		MarketingAirline  Airline   `json:"marketing_airline"`
		OperatingAirline  Airline   `json:"operating_airline"`
		ValidatingAirline Airline   `json:"validating_airline,omitempty"`
		//TravelTime			int			`json:"travel_time,omitempty"`
	}
	var flight2 Flight2
	if err := json.Unmarshal(data, &flight2); err != nil {
		//logs.Log.WithError(err).Error("Error unmarshal json") //todo log
		return err
	}
	flight.Aircraft = flight2.Aircraft
	flight.FlightNumber = flight2.FlightNumber
	flight.DepartureLocation = flight2.DepartureLocation
	departureDate, err := time.Parse(TimeFormat, flight2.DepartureDate)
	if err != nil {
		return err
	}

	flight.DepartureDate = departureDate
	flight.ArrivalLocation = flight2.ArrivalLocation
	arrivalDate, err := time.Parse(TimeFormat, flight2.ArrivalDate)
	if err != nil {
		return err
	}

	flight.ArrivalDate = arrivalDate
	flight.MarketingAirline = &flight2.MarketingAirline
	flight.OperatingAirline = &flight2.OperatingAirline
	flight.ValidatingAirline = &flight2.ValidatingAirline
	//flight.TravelTime = flight2.TravelTime
	return nil
}

// CodeShare returns sign of CodeShare flight
func (flight *Flight) CodeShare() bool {
	return flight.MarketingAirline != flight.OperatingAirline
}
