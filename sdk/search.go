package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/masterPricerTravelBoardSearch/v14_3_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/masterPricerTravelBoardSearch/v14_3_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"strings"
)

const (
	ActSearchFareMasterPricerTravelBoardSearch = "FMPTBQ_14_3_1A"
	// ActSearchFareMasterPricerTravelBoardSearch = "FMPTBQ_16_3_1A"
)

func (client *AmadeusClient) FareMasterPricerTravelBoardSearchV143(query *Fare_MasterPricerTravelBoardSearch_v14_3.FareMasterPricerTravelBoardSearch) (*Fare_MasterPricerTravelBoardSearchReply_v14_3.FareMasterPricerTravelBoardSearchReply, *ResponseSOAP4Header, error) {
	var soapAction = ActSearchFareMasterPricerTravelBoardSearch
	var reply = new(Fare_MasterPricerTravelBoardSearchReply_v14_3.FareMasterPricerTravelBoardSearchReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
