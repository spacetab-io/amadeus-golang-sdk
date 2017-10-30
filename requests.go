package amadeus

import (
	cc "github.com/tmconsulting/amadeus-ws-go/reqstructs/command_cryptic"
	diit "github.com/tmconsulting/amadeus-ws-go/reqstructs/docissuance_issueticket"
	fppwbc "github.com/tmconsulting/amadeus-ws-go/reqstructs/fare_pricepnrwithbookingclass"
	pr "github.com/tmconsulting/amadeus-ws-go/reqstructs/pnr_retrieve"
	ql "github.com/tmconsulting/amadeus-ws-go/reqstructs/queue_list"
	sa "github.com/tmconsulting/amadeus-ws-go/reqstructs/security_authenticate"
	ss "github.com/tmconsulting/amadeus-ws-go/reqstructs/security_signout"
	tdt "github.com/tmconsulting/amadeus-ws-go/reqstructs/ticket_display_tst"

	ccr "github.com/tmconsulting/amadeus-ws-go/respstructs/command_cryptic_reply"
	diitr "github.com/tmconsulting/amadeus-ws-go/respstructs/docissuance_issueticket_reply"
	fppwbcr "github.com/tmconsulting/amadeus-ws-go/respstructs/fare_pricepnrwithbookingclass_reply"
	pl "github.com/tmconsulting/amadeus-ws-go/respstructs/pnr_list"
	qlr "github.com/tmconsulting/amadeus-ws-go/respstructs/queue_list_reply"
	sar "github.com/tmconsulting/amadeus-ws-go/respstructs/security_authenticate_reply"
	ssr "github.com/tmconsulting/amadeus-ws-go/respstructs/security_signout_reply"
	tdtr "github.com/tmconsulting/amadeus-ws-go/respstructs/ticket_display_tst_reply"
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

func (service *WebServicesPT) SecuritySignOut(request *ss.SecuritySignOut) (*ssr.SecuritySignOutReply, error) {
	soapAction := "VLSSOQ_04_1_1A"

	response := new(ssr.SecuritySignOutReply)
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

func (service *WebServicesPT) DocIssuanceIssueTicket(request *diit.DocIssuanceIssueTicket) (*diitr.DocIssuanceIssueTicketReply, error) {
	soapAction := "TTKTIQ_09_1_1A"

	response := new(diitr.DocIssuanceIssueTicketReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *WebServicesPT) FarePricePNRWithBookingClass(request *fppwbc.FarePricePNRWithBookingClass) (*fppwbcr.FarePricePNRWithBookingClassReply, error) {
	soapAction := "TPCBRQ_12_4_1A"

	response := new(fppwbcr.FarePricePNRWithBookingClassReply)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *WebServicesPT) PNRRetrieve2(request *pr.PNRRetrieve) (*pl.PNRList, error) {
	soapAction := "PNRRET_11_3_1A"

	response := new(pl.PNRList)
	session := new(Session)
	err := service.client.Call(url+"/"+service.wsap+"/"+soapAction, request, response, session)
	if err != nil {
		return nil, err
	}

	return response, nil
}
