package Fare_MasterPricerTravelBoardSearchResponse_v14_3

import (
	search "github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch"

	structsCommon "github.com/tmconsulting/amadeus-golang-sdk/structs"
)

// GroupOfFlights Group Of Flights
type FilterGroupOfFlights struct {
	ItineraryID       int
	GroupOfSegmentsID int
	Flights           []*structsCommon.Flight
	MajorityCarrier   string
	//ElapseFlyingTime		string
	Duration int
}

// FilterRules Filter Rules
type FilterRules struct {
	Itineraries       []*search.Itinerary
	OnlyNonStopFlight bool
	Airlines          []string
	//ExcludedAirlines		[]string
}

// CheckRulesCompanies returns true if validating company's rules
func (filterRules *FilterRules) CheckRulesCompanies(flight *structsCommon.Flight) bool {

	//for _, excludedAirline := range filterRules.ExcludedAirlines {
	//	if excludedAirline == flight.OperatingAirline {
	//		return false
	//	}
	//	if excludedAirline == flight.MarketingAirline {
	//		return false
	//	}
	//}
	if len(filterRules.Airlines) > 0 {
		var flag = false
		for _, airline := range filterRules.Airlines {
			if airline == flight.OperatingAirline.CodeEng {
				flag = true
			}
			if airline == flight.MarketingAirline.CodeEng {
				flag = true
			}
		}
		return flag
	}

	return true
}

// CheckRulesTimes Check Rules Times
func (filterRules *FilterRules) CheckRulesTimes(itineraryIndex int, isDeparture, isArrival bool, segment *structsCommon.Flight) bool {

	//var itinerary = filterRules.Itineraries[itineraryIndex]
	//
	//if isDeparture {
	//	var departureDateStart = itinerary.DepartureDate.Date
	//	if !departureDateStart.Equal(segment.DepartureDate) && !departureDateStart.Before(segment.DepartureDate) {
	//		return false
	//	}
	//
	//	if itinerary.DepartureDateTill.Error == nil {
	//		var departureDateEnd = itinerary.DepartureDateTill.Date
	//		if !departureDateEnd.Equal(segment.DepartureDate) && !departureDateEnd.After(segment.DepartureDate) {
	//			return false
	//		}
	//	}
	//}
	//
	//if isArrival && itinerary.ArrivalDate.Error == nil {
	//	var arrivalDate = itinerary.ArrivalDate.Date
	//	if !arrivalDate.Equal(segment.ArrivalDate) && !arrivalDate.After(segment.ArrivalDate) {
	//		return false
	//	}
	//}

	return true
}
