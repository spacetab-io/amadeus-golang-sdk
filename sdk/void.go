package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/salesReports/displayQueryReport/v10_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/salesReports/displayQueryReport/v10_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/cancelDocument/v11_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/cancelDocument/v11_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/deleteTST/v04_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/deleteTST/v04_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/displayTST/v07_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/displayTST/v07_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"strings"
)

const (
	ActVoidSalesReportsDisplayQueryReport = "TSRQRQ_10_1_1A"
	ActVoidTicketCancelDocument           = "TRCANQ_11_1_1A"
)

func (client *AmadeusClient) SalesReportsDisplayQueryReportV101(query *SalesReports_DisplayQueryReport_v10_1.SalesReportsDisplayQueryReport) (*SalesReports_DisplayQueryReportReply_v10_1.SalesReportsDisplayQueryReportReply, *ResponseSOAP4Header, error) {
	var soapAction = ActVoidSalesReportsDisplayQueryReport
	var reply = new(SalesReports_DisplayQueryReportReply_v10_1.SalesReportsDisplayQueryReportReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) TicketCancelDocumentV111(query *Ticket_CancelDocument_v11_1.TicketCancelDocument) (*Ticket_CancelDocumentReply_v11_1.TicketCancelDocumentReply, *ResponseSOAP4Header, error) {
	var soapAction = ActVoidTicketCancelDocument
	var reply = new(Ticket_CancelDocumentReply_v11_1.TicketCancelDocumentReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) TicketDeleteTSTv041(query *Ticket_DeleteTST_v04_1.TicketDeleteTST) (*Ticket_DeleteTSTReply_v04_1.TicketDeleteTSTReply, *ResponseSOAP4Header, error) {
	var soapAction = "TTSTDQ_04_1_1A"
	var reply = new(Ticket_DeleteTSTReply_v04_1.TicketDeleteTSTReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) TicketDisplayTSTv071(query *Ticket_DisplayTST_v07_1.TicketDisplayTST) (*Ticket_DisplayTSTReply_v07_1.TicketDisplayTSTReply, *ResponseSOAP4Header, error) {
	var soapAction = "TTSTRQ_07_1_1A" //TTSTRQ_15_1_1A
	var reply = new(Ticket_DisplayTSTReply_v07_1.TicketDisplayTSTReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
