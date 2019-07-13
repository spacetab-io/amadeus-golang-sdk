package client

import (
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/checkRules/v07.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/checkRules/v07.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve/v11.3/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve/v11.3/response"
	request2 "github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve/v19.1/request"
	response2 "github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve/v19.1/response"
)

func (client *AmadeusClient) PNRRetrieveV113(query *request.Request) (*response.Response, *ResponseSOAPHeader, error) {
	var reply response.Response
	header, err := client.service.Call(soapUrl, "PNRRET_11_3_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) PNRRetrieveV191(query *request2.Request) (*response2.Response, *ResponseSOAPHeader, error) {
	var reply response2.Response
	header, err := client.service.Call(soapUrl, "PNRRET_19_1_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) FareInformativePricingWithoutPNRV124(query *Fare_InformativePricingWithoutPNR_v12_4.Request) (*Fare_InformativePricingWithoutPNRReply_v12_4.Response, *ResponseSOAPHeader, error) {
	var reply Fare_InformativePricingWithoutPNRReply_v12_4.Response
	header, err := client.service.Call(soapUrl, "TIPNRQ_12_4_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) FareCheckRulesV071(query *Fare_CheckRulesRequest_v07_1.Request) (*Fare_CheckRulesResponse_v07_1.Response, *ResponseSOAPHeader, error) {
	var reply Fare_CheckRulesResponse_v07_1.Response
	header, err := client.service.Call(soapUrl, "FARQNQ_07_1_1A", query, &reply, client)
	return &reply, header, err
}
