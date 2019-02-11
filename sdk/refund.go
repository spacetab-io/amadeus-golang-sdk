package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ama/ticketIgnoreRefund/v03_0"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ama/ticketInitRefund/v03_0"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ama/ticketProcessRefund/v03_0"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/processEDoc/v15_2_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/ticket/processEDoc/v15_2_reply"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/amadeus-agent-go/support"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/amadeus-agent-go/utils"
	"gitlab.teamc.io/tm-consulting/tmc24/provider-logs/receiver.git"
	"strings"
)

const (
	ActRefundAMATicketIgnoreRefund  = "Ticket_IgnoreRefund_3.0"
	ActRefundAMATicketInitRefund    = "Ticket_InitRefund_3.0"
	ActRefundAMATicketProcessRefund = "Ticket_ProcessRefund_3.0"
	ActRefundTicketProcessEDoc      = "TATREQ_15_2_1A"
)

func (client *AmadeusClient) AMATicketIgnoreRefundV030(query *AMA_TicketIgnoreRefund_v03_0.AMATicketIgnoreRefundRQ, attr *receiver.LogAttributes) (*AMA_TicketIgnoreRefund_v03_0.AMATicketIgnoreRefundRS, *ResponseSOAP4Header, error) {
	var url = "http://webservices.amadeus.com/"
	var soapAction = ActRefundAMATicketIgnoreRefund
	var reply = new(AMA_TicketIgnoreRefund_v03_0.AMATicketIgnoreRefundRS)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(url+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) AMATicketInitRefundV030(query *AMA_TicketInitRefund_v03_0.AMATicketInitRefundRQ, attr *receiver.LogAttributes) (*AMA_TicketInitRefund_v03_0.AMATicketInitRefundRS, *ResponseSOAP4Header, error) {
	var url = "http://webservices.amadeus.com/"
	var soapAction = ActRefundAMATicketInitRefund
	var reply = new(AMA_TicketInitRefund_v03_0.AMATicketInitRefundRS)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(url+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) AMATicketProcessRefundV030(query *AMA_TicketProcessRefund_v03_0.AMATicketProcessRefundRQ, attr *receiver.LogAttributes) (*AMA_TicketProcessRefund_v03_0.AMATicketProcessRefundRS, *ResponseSOAP4Header, error) {
	var url = "http://webservices.amadeus.com/"
	var soapAction = ActRefundAMATicketProcessRefund
	var reply = new(AMA_TicketProcessRefund_v03_0.AMATicketProcessRefundRS)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(url+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) TicketProcessEDocV152(query *Ticket_ProcessEDoc_v15_2.TicketProcessEDoc, attr *receiver.LogAttributes) (*Ticket_ProcessEDocReply_v15_2.TicketProcessEDocReply, *ResponseSOAP4Header, error) {
	var soapAction = ActRefundTicketProcessEDoc
	var reply = new(Ticket_ProcessEDocReply_v15_2.TicketProcessEDocReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
