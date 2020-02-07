package amadeus_sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/hooks"
	"time"
)

var soapUrl = "http://webservices.amadeus.com/WSAP/"

func CreateAmadeusClient(url, originator, passwordRaw, officeId string, timeout time.Duration) *AmadeusClient {
	var client = new(AmadeusClient)
	client.service = CreateWebServicePTSOAP4Header(url, originator, passwordRaw, officeId, true, timeout)
	//client.service = soap4.NewAmadeusWebServicesPTSOAP4Header(url, originator, passwordRaw, officeId, true)
	client.session = CreateSession()
	client.Hooks = make(hooks.LevelHooks)
	return client
}

type AmadeusClient struct {
	service *WebServicePT
	//service *soap4_0.WebServicesPTSOAP4Header
	session *Session
	//session *soap4_0.Session_v3
	// messageIds		[]string
	Hooks hooks.LevelHooks
}

func (client *AmadeusClient) Close() error {
	if client == nil || client.service == nil {
		return nil
	}
	if client.session != nil && client.session.TransactionStatusCode != TransactionStatusCode[End] {
		client.session.TransactionStatusCode = TransactionStatusCode[End]
		if _, _, err := client.SecuritySignOutV041(); err != nil {
			return err
		}
		client.session = nil
	}
	return nil
}
