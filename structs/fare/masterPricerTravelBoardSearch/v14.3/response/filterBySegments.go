package Fare_MasterPricerTravelBoardSearchResponse_v14_3

import (
	structsCommon "github.com/tmconsulting/amadeus-golang-sdk/structs"
)

// InternalGroupOfFlights Group Of Flights
type InternalGroupOfFlights struct {
	ItineraryID       int
	GroupOfSegmentsID int
	Flights           []*structsCommon.Flight
	MajorityCarrier   string
	//ElapseFlyingTime		string
	Duration int
}
