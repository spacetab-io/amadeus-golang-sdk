package amadeus_sdk

import (
	calculatevat "github.com/tmconsulting/amadeus-golang-sdk/sdk/vatCalculator/CalculateVAT"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"strings"
)

func (client *AmadeusClient) CalculateVAT(query *calculatevat.CalculateVAT) (*calculatevat.CalculateVATResponse, *ResponseSOAP4Header, error) {
	soapUrl := "http://amadeus.ru/"
	var soapAction = "CalculateVAT"
	var reply = new(calculatevat.CalculateVATResponse)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
