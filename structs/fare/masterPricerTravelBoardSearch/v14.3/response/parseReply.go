package Fare_MasterPricerTravelBoardSearchResponse_v14_3

import (
	"fmt"
	search "github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch"
	"strings"
)

func (r *Response) CheckErrorReply() error {

	if r.ErrorMessage != nil {
		return fmt.Errorf(strings.Join(r.ErrorMessage.ErrorMessageText.Description, "\n"))
	}
	return nil
}

func (r *Response) ToCommon() *search.Response {

	var response search.Response

	for itineraryId, flightIndex := range r.FlightIndex {
		for variantId, groupOfFLight := range flightIndex.GroupOfFlights {
			var flightDetails []search.FlightDetail
			for _, groupOfFLight := range groupOfFLight.FlightDetails {

				flightDetails = append(flightDetails, search.FlightDetail{})

				response.RoutesSegments[search.ItineraryID(itineraryId)][search.VariantID(variantId)] = search.Variant{
					flightDetails,
				}
			}
		}
	}

	//for _, recommendation := range r.Recommendation {
	//
	//}

	return &response
}
