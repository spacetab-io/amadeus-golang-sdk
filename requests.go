package amadeus

import (
	cc "github.com/kidem/amadeus-ws-go/reqstructs/command_cryptic"
	ql "github.com/kidem/amadeus-ws-go/reqstructs/queue_list"
	sa "github.com/kidem/amadeus-ws-go/reqstructs/security_authenticate"
	tdt "github.com/kidem/amadeus-ws-go/reqstructs/ticket_display_tst"

	ccr "github.com/kidem/amadeus-ws-go/respstructs/command_cryptic_reply"
	qlr "github.com/kidem/amadeus-ws-go/respstructs/queue_list_reply"
	sar "github.com/kidem/amadeus-ws-go/respstructs/security_authenticate_reply"
	tdtr "github.com/kidem/amadeus-ws-go/respstructs/ticket_display_tst_reply"
)

var url string = "http://webservices.amadeus.com"

func (service *WebServicesPT) SecurityAuthenticate(request *sa.SecurityAuthenticate) (*sar.SecurityAuthenticateReply, *Session, error) {
	soapAction := "VLSSLQ_06_1_1A"

	response := new(sar.SecurityAuthenticateReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, nil, err
	}

	return response, session, nil
}

func (service *WebServicesPT) SecuritySignOut(request *SecuritySignOut) (*SecuritySignOutReply, error) {
	soapAction := "VLSSOQ_04_1_1A"

	response := new(SecuritySignOutReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *WebServicesPT) CommandCryptic(request *cc.CommandCryptic) (*ccr.CommandCrypticReply, *Session, error) {
	soapAction := "HSFREQ_07_3_1A"

	response := new(ccr.CommandCrypticReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, nil, err
	}

	return response, session, nil
}

func (service *WebServicesPT) QueueList(request *ql.QueueList) (*qlr.QueueListReply, error) {
	soapAction := "QDQLRQ_11_1_1A"

	response := new(qlr.QueueListReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *WebServicesPT) TicketDisplayTST(request *tdt.TicketDisplayTST) (*tdtr.TicketDisplayTSTReply, error) {
	soapAction := "TTSTRQ_07_1_1A"

	response := new(tdtr.TicketDisplayTSTReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, err
	}

	return response, nil
}
