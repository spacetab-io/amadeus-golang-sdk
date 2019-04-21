package sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/structs/commandCryptic/v07.3"
)

func (client *AmadeusClient) CommandCrypticV073(query *CommandCryptic_v07_3.Request) (*CommandCryptic_v07_3.Response, *ResponseSOAPHeader, error) {
	var reply CommandCryptic_v07_3.Response
	header, err := client.service.Call(soapUrl, "HSFREQ_07_3_1A", query, &reply, client)
	return &reply, header, err
}
