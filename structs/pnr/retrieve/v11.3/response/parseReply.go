package PNR_Retrieve_v11_3_response

import (
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs/pnr/retrieve"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/utils"
)

func (r *Response) ToCommon() *PNR_Information.Response {

	var response PNR_Information.Response

	if r == nil {
		return &response
	}

	for _, travellerInfo := range r.TravellerInfo {
		for _, passengerData := range travellerInfo.PassengerData {
			for _, passenger := range passengerData.TravellerInformation.Passenger {
				response.TravellesInfo = append(response.TravellesInfo, PNR_Information.TravellerInfo{
					FirstName:   passenger.FirstName,
					LastName:    passengerData.TravellerInformation.Traveller.Surname,
					Type:        PNR_Information.PaxType(passenger.Type),
					DateOfBirth: utils.AmadeusDateConvert(passengerData.DateOfBirth.DateAndTimeDetails.Date, ""),
					Quaifier:    travellerInfo.ElementManagementPassenger.Reference.Qualifier,
					Number:      int(travellerInfo.ElementManagementPassenger.Reference.Number),
				})
			}
		}
	}

	var segments []PNR_Information.Segment
	for _, originDestinationDetails := range r.OriginDestinationDetails {
		for _, itineraryInfo := range originDestinationDetails.ItineraryInfo {
			if itineraryInfo.TravelProduct != nil {
				segments = append(segments, PNR_Information.Segment{
					DepartureDate: utils.AmadeusDateTimeConvert(itineraryInfo.TravelProduct.Product.DepDate, itineraryInfo.TravelProduct.Product.DepTime),
					ArrivalDate:   utils.AmadeusDateTimeConvert(itineraryInfo.TravelProduct.Product.ArrDate, itineraryInfo.TravelProduct.Product.ArrTime),
					FlightNumber:  itineraryInfo.TravelProduct.ProductDetails.Identification,
					DepartureLocation: PNR_Information.Location{
						AirportCode: itineraryInfo.TravelProduct.BoardpointDetail.CityCode,
					},
					ArrivalLocation: PNR_Information.Location{
						AirportCode: itineraryInfo.TravelProduct.OffpointDetail.CityCode,
					},
					MarketingAirline: &PNR_Information.Airline{
						CodeEng: itineraryInfo.ItineraryReservationInfo.Reservation.CompanyId,
					},
					Aircraft: &PNR_Information.Aircraft{
						Code: &itineraryInfo.FlightDetail.ProductDetails.Equipment,
					},
				})
			}
		}
	}

	response.Segments = segments

	return &response
}
