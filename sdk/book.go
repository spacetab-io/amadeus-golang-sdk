package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/air/sellFromRecommendation/v05_2_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/air/sellFromRecommendation/v05_2_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/pricePNRWithBookingClass/v14_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/pricePNRWithBookingClass/v14_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/pnr/addMultiElements/v11_3_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/pnr/cancel/v11_3_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/pnr/reply/v11_3"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/createTSTFromPricing/v04_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/createTSTFromPricing/v04_1_reply"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/amadeus-agent-go/support"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/amadeus-agent-go/utils"
	"gitlab.teamc.io/tm-consulting/tmc24/provider-logs/receiver.git"
	"strings"
)

const (
	ActBookAirSellFromRecommendation    = "ITAREQ_05_2_IA"
	ActBookPNRAddMultiElements          = "PNRADD_11_3_1A"
	ActBookFarePricePNRWithBookingClass = "TPCBRQ_14_1_1A"
	ActBookTicketCreateTSTFromPricing   = "TAUTCQ_04_1_1A"
	ActBookPNRCancel                    = "PNRXCL_11_3_1A"
)

func (client *AmadeusClient) AirSellFromRecommendationV052(query *Air_SellFromRecommendation_v05_2.AirSellFromRecommendation, attr *receiver.LogAttributes) (*Air_SellFromRecommendationReply_v05_2.AirSellFromRecommendationReply, *ResponseSOAP4Header, error) {
	var soapAction = ActBookAirSellFromRecommendation
	var reply = new(Air_SellFromRecommendationReply_v05_2.AirSellFromRecommendationReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	//header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) PNRAddMultiElementsV113(query *PNR_AddMultiElements_v11_3.PNRAddMultiElements, attr *receiver.LogAttributes) (*PNR_Reply_v11_3.PNRReply, *ResponseSOAP4Header, error) {
	var soapAction = ActBookPNRAddMultiElements
	var reply = new(PNR_Reply_v11_3.PNRReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) FarePricePNRWithBookingClassV141(query *Fare_PricePNRWithBookingClass_v14_1.FarePricePNRWithBookingClass, attr *receiver.LogAttributes) (*Fare_PricePNRWithBookingClassReply_v14_1.FarePricePNRWithBookingClassReply, *ResponseSOAP4Header, error) {
	var soapAction = ActBookFarePricePNRWithBookingClass
	var reply = new(Fare_PricePNRWithBookingClassReply_v14_1.FarePricePNRWithBookingClassReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) TicketCreateTSTFromPricingV041(query *Ticket_CreateTSTFromPricing_v04_1.TicketCreateTSTFromPricing, attr *receiver.LogAttributes) (*Ticket_CreateTSTFromPricingReply_v04_1.TicketCreateTSTFromPricingReply, *ResponseSOAP4Header, error) {
	var soapAction = ActBookTicketCreateTSTFromPricing
	var reply = new(Ticket_CreateTSTFromPricingReply_v04_1.TicketCreateTSTFromPricingReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) PNRCancelV113(query *PNR_Cancel_v11_3.PNRCancel, attr *receiver.LogAttributes) (*PNR_Reply_v11_3.PNRReply, *ResponseSOAP4Header, error) {
	var soapAction = ActBookPNRCancel
	var reply = new(PNR_Reply_v11_3.PNRReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
