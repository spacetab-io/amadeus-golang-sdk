package search

//type Response struct {
//	TravellesInfo   []TravellerInfo
//	Segments        []Segment
//	Recommendations []Recommendation
//	RoutesSegments  FlightDetails
//}
//
//type TravellerInfo struct {
//	FirstName   string
//	LastName    string
//	Type        PaxType // ADT, CHD, INF
//	DateOfBirth time.Time
//	Quaifier    string // PT, PI
//	Number      int    // reference number
//}
//
//type PaxType string
//
//type Segment struct {
//	SegmentId         string    `json:"-"`
//	FlightNumber      string    `json:"flight_number"`
//	DepartureLocation Location  `json:"departure"`
//	DepartureDate     time.Time `json:"departure_date"`
//	ArrivalLocation   Location  `json:"arrival"`
//	ArrivalDate       time.Time `json:"arrival_date"`
//	MarketingAirline  *Airline  `json:"marketing_airline"`
//	OperatingAirline  *Airline  `json:"operating_airline"`
//	ValidatingAirline *Airline  `json:"validating_airline,omitempty"`
//	Aircraft          *Aircraft `json:"aircraft,omitempty"`
//}
//
//type Location struct {
//	AirportCode string `json:"airport_code"`
//	CountryCode string `json:"country_code"`
//	CityCode    string `json:"city_code"`
//	Terminal    string `json:"terminal"`
//	Type        string `json:"-"`
//}
//
//type Airline struct {
//	CodeEng string `json:"code_eng" bson:"codeEng"`
//	CodeRus string `json:"code_rus,omitempty" bson:"codeRus,omitempty"`
//	NameEng string `json:"name_eng,omitempty" bson:"nameEng,omitempty"`
//	NameRus string `json:"name_rus,omitempty" bson:"nameRus,omitempty"`
//}
//
//type Aircraft struct {
//	Code *string           `json:"code" db:"code" bson:"code"`
//	Name map[string]string `json:"name,omitempty" db:"name" bson:"name,omitempty"`
//}

// Recommendation structure
//type Recommendation struct {
//	ID                string        `json:"id,omitempty"`
//	Provider          string        `json:"provider"`
//	Class             []string      `json:"class"`
//	Price             int           `json:"price"`
//	ItinerarySegments FlightDetails `json:"itinerary_segments"`
//	ValidatingAirline Airline       `json:"validating_airline"`
//	Interline         bool          `json:"interline"`
//	LastTicketDate    string        `json:"last_date,omitempty"`
//	Refundable        string        `json:"refundable,omitempty"`
//}
//
//// ItineraryID itinerary type
//// groupOfFlights.propFlightGrDetail.flightProposal.ref
//type ItineraryID int
//
//// ItineraryID variant type
//// groupOfFlights.flightDetails array key
//type VariantID int
//
//// FlightDetails map of FlightDetails
//type FlightDetails map[ItineraryID]map[VariantID]Variant
//
//type Variant struct {
//	FlightDetails []FlightDetail
//}
//
//// RouteSegments Route Segments structure
//type RouteSegments struct {
//	GroupOfSegmentsID int                 `json:"-"`
//	ItinerarySegments []*ItinerarySegment `json:"segments"`
//	TravelTime        int                 `json:"travel_time"`
//}
//
//// ItinerarySegment segments of Itinerary
//type ItinerarySegment struct {
//	SegmentID string `json:"segment_id"`
//	CodeShare bool   `json:"codeshare"`
//	*FlightDetail
//}
//
//// FlightDetail structure
//type FlightDetail struct {
//	Class         string         `json:"class"` // @TODO Deprecate it!
//	Cabin         string         `json:"cabin"`
//	FlightClass   string         `json:"flight_class"` // @TODO rename it to Subclass !
//	FareClass     string         `json:"fare_class"`
//	BookingClass  string         `json:"rbd"`
//	FareBasisCode string         `json:"fare_basis_code"`
//	Availability  int            `json:"availability"`
//	//Baggage       *t.BaggageType `json:"baggage,omitempty"`
//
//	DepartureTime     time.Time
//	ArrivalTime       time.Time
//	DepartureLocation Location
//	ArrivalLocation   Location
//}
//
//// Equal compare flight details
//func (fd *FlightDetail) Equal(test *FlightDetail) bool {
//	if fd.BookingClass == test.BookingClass &&
//		fd.FareClass == test.FareClass &&
//		fd.FareBasisCode == test.FareBasisCode &&
//		fd.Availability == test.Availability {
//		return true
//	}
//	return false
//}
