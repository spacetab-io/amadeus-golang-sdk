package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/client"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/request"
	//"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/response"
)

func (s *service) FareMasterPricerTravelBoardSearch(query *search.Request) (*search.Response, *client.ResponseSOAPHeader, error) {
	switch s.mm[FareMasterPricerTravelBoardSearch] {
	case FareMasterPricerTravelBoardSearchV143:
		query := Fare_MasterPricerTravelBoardSearchRequest_v14_3.MakeRequest(query)
		response, header, err := s.sdk.FareMasterPricerTravelBoardSearchV143(query)
		if response == nil {
			return nil, header, err
		}

		errResponse := response.CheckErrorReply()
		if errResponse != nil {
			return nil, header, errResponse
		}

		return response.ToCommon(), header, err
	}
	return nil, nil, nil
}

func (s *service) FareInformativeBestPricingWithout(query *Fare_InformativeBestPricingWithoutPNRRequest_v12_4.Request) (*Fare_InformativeBestPricingWithoutPNRResponse_v12_4.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.FareInformativeBestPricingWithoutPNRV124(query)
}

func (s *service) FareInformativePricingWithoutPNR(query *Fare_InformativePricingWithoutPNR_v12_4.Request) (*Fare_InformativePricingWithoutPNRReply_v12_4.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.FareInformativePricingWithoutPNRV124(query)
}
