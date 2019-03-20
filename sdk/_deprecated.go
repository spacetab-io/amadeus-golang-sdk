package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/informativeBestPricingWithoutPNR/v12_4_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/fare/informativeBestPricingWithoutPNR/v12_4_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/pnr/ignore/v04_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/pnr/ignore/v04_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/salesReports/displayTransactionReport/v13_2_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/salesReports/displayTransactionReport/v13_2_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/deleteTST/v04_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/deleteTST/v04_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/displayTST/v07_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/displayTST/v07_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"strings"
)

func (client *AmadeusClient) FareMasterPricerTravelBoardSearchV163(query *Fare_MasterPricerTravelBoardSearch_v16_3.FareMasterPricerTravelBoardSearch, attr *receiver.LogAttributes) (*Fare_MasterPricerTravelBoardSearchReply_v16_3.FareMasterPricerTravelBoardSearchReply, *ResponseSOAP4Header, error) {
	var soapAction = "FMPTBQ_16_3_1A"
	var reply = new(Fare_MasterPricerTravelBoardSearchReply_v16_3.FareMasterPricerTravelBoardSearchReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

// deprecated?
func (client *AmadeusClient) TicketDisplayTSTv071(query *Ticket_DisplayTST_v07_1.TicketDisplayTST) (*Ticket_DisplayTSTReply_v07_1.TicketDisplayTSTReply, *ResponseSOAP4Header, error) {
	var soapAction = "TTSTRQ_07_1_1A"
	var reply = new(Ticket_DisplayTSTReply_v07_1.TicketDisplayTSTReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client.session)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

// deprecated?
func (client *AmadeusClient) TicketDeleteTSTv041(query *Ticket_DeleteTST_v04_1.TicketDeleteTST) (*Ticket_DeleteTSTReply_v04_1.TicketDeleteTSTReply, *ResponseSOAP4Header, error) {
	var soapAction = "TTSTDQ_04_1_1A"
	var reply = new(Ticket_DeleteTSTReply_v04_1.TicketDeleteTSTReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client.session)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

// deprecated?
func (client *AmadeusClient) SalesReportsDisplayTransactionReportV132(query *SalesReports_DisplayTransactionReport_v13_2.SalesReportsDisplayTransactionReport) (*SalesReports_DisplayTransactionReportReply_v13_2.SalesReportsDisplayTransactionReportReply, *ResponseSOAP4Header, error) {
	var soapAction = "TSRTRQ_13_2_1A"
	var reply = new(SalesReports_DisplayTransactionReportReply_v13_2.SalesReportsDisplayTransactionReportReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client.session)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

// deprecated?
func (client *AmadeusClient) PNRIgnoreV041(query *PNR_Ignore_v04_1.PNRIgnore) (*PNR_IgnoreReply_v04_1.PNRIgnoreReply, *ResponseSOAP4Header, error) {
	var soapAction = "CLTREQ_04_1_IA"
	var reply = new(PNR_IgnoreReply_v04_1.PNRIgnoreReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client.session)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

// deprecated?
func (client *AmadeusClient) FareInformativeBestPricingWithoutPNRv124(query *Fare_InformativeBestPricingWithoutPNR_v12_4.FareInformativeBestPricingWithoutPNR) (*Fare_InformativeBestPricingWithoutPNRReply_v12_4.FareInformativeBestPricingWithoutPNRReply, *ResponseSOAP4Header, error) {
	var soapAction = "TIBNRQ_12_4_1A"
	var reply = new(Fare_InformativeBestPricingWithoutPNRReply_v12_4.FareInformativeBestPricingWithoutPNRReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client.session)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
