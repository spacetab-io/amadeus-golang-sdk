package structs

import "encoding/json"

type Itinerary struct {
	DepartureLocation Location `json:"departure_location"`
	ArrivalLocation   Location `json:"arrival_location"`
}

type Location struct {
	AirportCode string `json:"airport_code"`
	CountryCode string `json:"country_code"`
	CityCode    string `json:"city_code"`
	Terminal    string `json:"terminal"`
	Type        string `json:"-"`
}

func (loc *Location) MarshalJSON() ([]byte, error) {
	if loc.Type == "P" {
		return json.Marshal(struct {
			AirportCode string `json:"airport_code"`
			Terminal    string `json:"terminal"`
		}{
			AirportCode: loc.AirportCode,
			Terminal:    loc.Terminal,
		})
	}
	return json.Marshal(struct {
		AirportCode string `json:"airport_code"`
		CountryCode string `json:"country_code"`
		CityCode    string `json:"city_code"`
		Terminal    string `json:"terminal"`
	}{
		AirportCode: loc.AirportCode,
		CountryCode: loc.CountryCode,
		CityCode:    loc.CityCode,
		Terminal:    loc.Terminal,
	})
}
