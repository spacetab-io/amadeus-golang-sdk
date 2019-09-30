package search

import (
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs"
)

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
