package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/docIssuance/issueTicket/v09_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/docIssuance/issueTicket/v09_1_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"strings"
)

const (
	ActIssueDocIssuanceIssueTicket = "TTKTIQ_09_1_1A"
)

func (client *AmadeusClient) DocIssuanceIssueTicketV091(query *DocIssuance_IssueTicket_v09_1.DocIssuanceIssueTicket) (*DocIssuance_IssueTicketReply_v09_1.DocIssuanceIssueTicketReply, *ResponseSOAP4Header, error) {
	var soapAction = ActIssueDocIssuanceIssueTicket
	var reply = new(DocIssuance_IssueTicketReply_v09_1.DocIssuanceIssueTicketReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
