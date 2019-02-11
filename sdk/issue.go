package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/docIssuance/issueTicket/v09_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/docIssuance/issueTicket/v09_1_reply"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/amadeus-agent-go/support"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/amadeus-agent-go/utils"
	"gitlab.teamc.io/tm-consulting/tmc24/provider-logs/receiver.git"
	"strings"
)

const (
	ActIssueDocIssuanceIssueTicket = "TTKTIQ_09_1_1A"
)

func (client *AmadeusClient) DocIssuanceIssueTicketV091(query *DocIssuance_IssueTicket_v09_1.DocIssuanceIssueTicket, attr *receiver.LogAttributes) (*DocIssuance_IssueTicketReply_v09_1.DocIssuanceIssueTicketReply, *ResponseSOAP4Header, error) {
	var soapAction = ActIssueDocIssuanceIssueTicket
	var reply = new(DocIssuance_IssueTicketReply_v09_1.DocIssuanceIssueTicketReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
