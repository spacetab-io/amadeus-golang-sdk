package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/command/cryptic/v07_3_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/command/cryptic/v07_3_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"strings"
)

func (client *AmadeusClient) CommandCrypticV073(query *Command_Cryptic_v07_3.CommandCryptic) (*Command_CrypticReply_v07_3.CommandCrypticReply, *ResponseSOAP4Header, error) {
	var soapAction = "HSFREQ_07_3_1A"
	var reply = new(Command_CrypticReply_v07_3.CommandCrypticReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
