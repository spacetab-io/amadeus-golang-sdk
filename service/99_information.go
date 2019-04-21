package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/commandCryptic/v07.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/checkRules/v07.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/checkRules/v07.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/reply/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/displayTST/v07.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/displayTST/v07.1/response"
)

func (s *service) PNRRetrieve(query *PNR_Retrieve_v11_3.Request) (*PNR_Reply_v11_3.Response, *sdk.ResponseSOAPHeader, error) {
	return s.sdk.PNRRetrieveV113(query)
}

func (s *service) TicketDisplayTST(query *Ticket_DisplayTSTRequest_v07_1.Request) (*Ticket_DisplayTSTResponse_v07_1.Response, *sdk.ResponseSOAPHeader, error) {
	return s.sdk.TicketDisplayTSTV071(query)
}

func (s *service) FareCheckRules(query *Fare_CheckRulesRequest_v07_1.Request) (*Fare_CheckRulesResponse_v07_1.Response, *sdk.ResponseSOAPHeader, error) {
	return s.sdk.FareCheckRulesV071(query)
}

func (s *service) CommandCryptic(msg string) (*CommandCryptic_v07_3.Response, *sdk.ResponseSOAPHeader, error) {
	q := &CommandCryptic_v07_3.Request{
		MessageAction: &CommandCryptic_v07_3.MessageAction{
			MessageFunctionDetails: &CommandCryptic_v07_3.MessageFunctionDetails{
				MessageFunction: msg,
			},
		},
	}
	return s.sdk.CommandCrypticV073(q)
}
