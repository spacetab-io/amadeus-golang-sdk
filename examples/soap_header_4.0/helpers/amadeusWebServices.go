package helpers

import (
	soap4 "github.com/tmconsulting/amadeus-golang-sdk/soap4.0"
)

const (
	wsap = "WSAPId"
	url = "https://test.webservices.amadeus.com/" + wsap
	originator = "Originator"
	clearPassword = "AMADEUS"
	officeId = "OfficeId"
)

func NewWebServices() *soap4.WebServicesPTSOAP4Header {

	return soap4.NewAmadeusWebServicesPTSOAP4Header(url, originator, clearPassword, officeId, true)
}
