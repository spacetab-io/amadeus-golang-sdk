package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/checkRules/v07_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/checkRules/v07_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/informativePricingWithoutPNR/v12_4_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/informativePricingWithoutPNR/v12_4_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/informativePricingWithoutPNR/v19_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/informativePricingWithoutPNR/v19_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/pnr/reply/v11_3"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/pnr/retrieve/v11_3_query"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"strings"
)

const (
	ActInfoPNRRetrieve                      = "PNRRET_11_3_1A"
	ActInfoFareInformativePricingWithoutPNR = "TIPNRQ_12_4_1A"
	ActInfoFareCheckRules                   = "FARQNQ_07_1_1A"
)

func (client *AmadeusClient) PNRRetrieveV113(query *PNR_Retrieve_v11_3.PNRRetrieve) (*PNR_Reply_v11_3.PNRReply, *ResponseSOAP4Header, error) {
	var soapAction = ActInfoPNRRetrieve
	var reply = new(PNR_Reply_v11_3.PNRReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) FareInformativePricingWithoutPNRv124(query *Fare_InformativePricingWithoutPNR_v12_4.FareInformativePricingWithoutPNR) (*Fare_InformativePricingWithoutPNRReply_v12_4.FareInformativePricingWithoutPNRReply, *ResponseSOAP4Header, error) {
	var soapAction = ActInfoFareInformativePricingWithoutPNR
	var reply = new(Fare_InformativePricingWithoutPNRReply_v12_4.FareInformativePricingWithoutPNRReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) FareInformativePricingWithoutPNRv191(query *Fare_InformativePricingWithoutPNR_v19_1.FareInformativePricingWithoutPNR) (*Fare_InformativePricingWithoutPNRReply_v19_1.FareInformativePricingWithoutPNRReply, *ResponseSOAP4Header, error) {
	var soapAction = "TIPNRQ_19_1"
	var reply = new(Fare_InformativePricingWithoutPNRReply_v19_1.FareInformativePricingWithoutPNRReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) FareCheckRulesV071(query *Fare_CheckRules_v07_1.FareCheckRules) (*Fare_CheckRulesReply_v07_1.FareCheckRulesReply, *ResponseSOAP4Header, error) {
	var soapAction = ActInfoFareCheckRules
	var reply = new(Fare_CheckRulesReply_v07_1.FareCheckRulesReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
