package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/vatCalculator/CalculateDetailVAT"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/vatCalculator/CalculateVAT"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"strings"
)

func (client *AmadeusClient) CalculateVAT(soapUrl string, query *calculatevat.CalculateVAT) (*calculatevat.CalculateVATResponse, *ResponseSOAP4Header, error) {
	var soapAction = "CalculateVAT"
	var reply = new(calculatevat.CalculateVATResponse)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func (client *AmadeusClient) CalculateDetailVAT(soapUrl string, query *calculatedetailvat.CalculateDetailVAT) (*calculatedetailvat.CalculateDetailVATResponse, *ResponseSOAP4Header, error) {
	var soapAction = "CalculateDetailVAT"
	var reply = new(calculatedetailvat.CalculateDetailVATResponse)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl, soapAction, messageId, query, reply, client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}
