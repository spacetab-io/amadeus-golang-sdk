package soap4_0

import (
	flireq051 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/air_flightinfo"
	itareq052 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/air_sellfromrecommendation"
	hsfreq073 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/command_cryptic"
	ttktiq091 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/docissuance_issueticket"
	farqnq071 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_checkrules"
	fcuqcq081 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_convertcurrency"
	tibnrq124 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_informativebestpricingwithoutpnr"
	tipnrq124 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_informativepricingwithoutpnr"
	fmpcaq143 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricercalendar"
	fmpcaq122 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricercalendar_old"
	fmptbq143 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricertravelboardsearch"
	fmptbq163 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricertravelboardsearch_new"
	fmptbq123 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricertravelboardsearch_old"
	tpcbrq141 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_pricepnrwithbookingclass"
	tpcbrq124 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_pricepnrwithbookingclass_old"
	pnradd113 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/pnr_addmultielements"
	pnrxcl113 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/pnr_cancel"
	pnrret113 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/pnr_retrieve"
	qcddrq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_countplanner"
	qcsdrq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_counttotal"
	qdqlrq111 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_list"
	quqmuq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_moveitem"
	quqpcq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_placepnr"
	quqmdq031 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/queue_removeitem"
	vlssoq041 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/security_signout"
	tautcq041 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/ticket_createtstfrompricing"
	ccvrqt061 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/ticket_creditcardcheck"
	ttstdq041 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/ticket_deletetst"
	ttstrq071 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/ticket_displaytst"

	flires051 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/air_flightinfo_reply"
	itares052 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/air_sellfromrecommendation_reply"
	hsfres073 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/command_cryptic_reply"
	ttktir091 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/docissuance_issueticket_reply"
	farqnr071 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_checkrules_reply"
	fcuqcr081 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_convertcurrency_reply"
	tibnrr124 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_informativebestpricingwithoutpnr_reply"
	tipnrr124 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_informativepricingwithoutpnr_reply"
	fmpcar122 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricercalendar_old_reply"
	fmpcar143 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricercalendar_reply"
	fmptbr163 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricertravelboardsearch_new_reply"
	fmptbr123 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricertravelboardsearch_old_reply"
	fmptbr143 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricertravelboardsearch_reply"
	tpcbrr124 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_pricepnrwithbookingclass_old_reply"
	tpcbrr141 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_pricepnrwithbookingclass_reply"
	tnlres001 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/pnr_list"
	pnracc113 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/pnr_reply"
	qcddrr031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_countplanner_reply"
	qcsdrr031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_counttotal_reply"
	qdqlrr111 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_list_reply"
	quqmur031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_moveitem_reply"
	quqpcr031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_placepnr_reply"
	quqmdr031 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/queue_removeitem_reply"
	vlssor041 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/security_signout_reply"
	tautcr041 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/ticket_createtstfrompricing_reply"
	ccvrsp061 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/ticket_creditcardcheck_reply"
	ttstdr041 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/ticket_deletetst_reply"
	ttstrr071 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/ticket_displaytst_reply"
)

var url = "http://webservices.amadeus.com/WSAP/"

func (service *WebServicesPTSOAP4Header) AirFlightInfo(session *Session_v3, messageId string, query *flireq051.AirFlightInfo) (*flires051.AirFlightInfoReply, *ResponseSOAP4Header, error) {
	soapAction := "FLIREQ_05_1_1A"

	reply := new(flires051.AirFlightInfoReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) AirSellFromRecommendation(session *Session_v3, messageId string, query *itareq052.AirSellFromRecommendation) (*itares052.AirSellFromRecommendationReply, *ResponseSOAP4Header, error) {
	soapAction := "ITAREQ_05_2_IA"

	reply := new(itares052.AirSellFromRecommendationReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) CommandCryptic(session *Session_v3, messageId string, query *hsfreq073.CommandCryptic) (*hsfres073.CommandCrypticReply, *ResponseSOAP4Header, error) {
	soapAction := "HSFREQ_07_3_1A"

	reply := new(hsfres073.CommandCrypticReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) DocIssuanceIssueTicket(session *Session_v3, messageId string, query *ttktiq091.DocIssuanceIssueTicket) (*ttktir091.DocIssuanceIssueTicketReply, *ResponseSOAP4Header, error) {
	soapAction := "TTKTIQ_09_1_1A"

	reply := new(ttktir091.DocIssuanceIssueTicketReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareCheckRules(session *Session_v3, messageId string, query *farqnq071.FareCheckRules) (*farqnr071.FareCheckRulesReply, *ResponseSOAP4Header, error) {
	soapAction := "FARQNQ_07_1_1A"

	reply := new(farqnr071.FareCheckRulesReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareConvertCurrency(session *Session_v3, messageId string, query *fcuqcq081.FareConvertCurrency) (*fcuqcr081.FareConvertCurrencyReply, *ResponseSOAP4Header, error) {
	soapAction := "FCUQCQ_08_1_1A"

	reply := new(fcuqcr081.FareConvertCurrencyReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareInformativeBestPricingWithoutPNR(session *Session_v3, messageId string, query *tibnrq124.FareInformativeBestPricingWithoutPNR) (*tibnrr124.FareInformativeBestPricingWithoutPNRReply, *ResponseSOAP4Header, error) {
	soapAction := "TIBNRQ_12_4_1A"

	reply := new(tibnrr124.FareInformativeBestPricingWithoutPNRReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareInformativePricingWithoutPNR(session *Session_v3, messageId string, query *tipnrq124.FareInformativePricingWithoutPNR) (*tipnrr124.FareInformativePricingWithoutPNRReply, *ResponseSOAP4Header, error) {
	soapAction := "TIPNRQ_12_4_1A"

	reply := new(tipnrr124.FareInformativePricingWithoutPNRReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareMasterPricerCalendar(session *Session_v3, messageId string, query *fmpcaq143.FareMasterPricerCalendar) (*fmpcar143.FareMasterPricerCalendarReply, *ResponseSOAP4Header, error) {
	soapAction := "FMPCAQ_14_3_1A"

	reply := new(fmpcar143.FareMasterPricerCalendarReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareMasterPricerCalendarOld(session *Session_v3, messageId string, query *fmpcaq122.FareMasterPricerCalendarOld) (*fmpcar122.FareMasterPricerCalendarOldReply, *ResponseSOAP4Header, error) {
	soapAction := "FMPCAQ_12_2_1A"

	reply := new(fmpcar122.FareMasterPricerCalendarOldReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareMasterPricerTravelBoardSearchNew(session *Session_v3, messageId string, query *fmptbq163.FareMasterPricerTravelBoardSearch) (*fmptbr163.FareMasterPricerTravelBoardSearchReply, *ResponseSOAP4Header, error) {
	soapAction := "FMPTBQ_16_3_1A"

	reply := new(fmptbr163.FareMasterPricerTravelBoardSearchReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareMasterPricerTravelBoardSearch(session *Session_v3, messageId string, query *fmptbq143.FareMasterPricerTravelBoardSearch) (*fmptbr143.FareMasterPricerTravelBoardSearchReply, *ResponseSOAP4Header, error) {
	soapAction := "FMPTBQ_14_3_1A"

	reply := new(fmptbr143.FareMasterPricerTravelBoardSearchReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FareMasterPricerTravelBoardSearchOld(session *Session_v3, messageId string, query *fmptbq123.FareMasterPricerTravelBoardSearchOld) (*fmptbr123.FareMasterPricerTravelBoardSearchOldReply, *ResponseSOAP4Header, error) {
	soapAction := "FMPTBQ_12_3_1A"

	reply := new(fmptbr123.FareMasterPricerTravelBoardSearchOldReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FarePricePNRWithBookingClass(session *Session_v3, messageId string, query *tpcbrq141.FarePricePNRWithBookingClass) (*tpcbrr141.FarePricePNRWithBookingClassReply, *ResponseSOAP4Header, error) {
	soapAction := "TPCBRQ_14_1_1A"

	reply := new(tpcbrr141.FarePricePNRWithBookingClassReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) FarePricePNRWithBookingClassOld(session *Session_v3, messageId string, query *tpcbrq124.FarePricePNRWithBookingClassOld) (*tpcbrr124.FarePricePNRWithBookingClassOldReply, *ResponseSOAP4Header, error) {
	soapAction := "TPCBRQ_12_4_1A"

	reply := new(tpcbrr124.FarePricePNRWithBookingClassOldReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) PNRAddMultiElements(session *Session_v3, messageId string, query *pnradd113.PNRAddMultiElements) (*pnracc113.PNRReply, *ResponseSOAP4Header, error) {
	soapAction := "PNRADD_11_3_1A"

	reply := new(pnracc113.PNRReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) PNRCancel(session *Session_v3, messageId string, query *pnrxcl113.PNRCancel) (*pnracc113.PNRReply, *ResponseSOAP4Header, error) {
	soapAction := "PNRXCL_11_3_1A"

	reply := new(pnracc113.PNRReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) PNRRetrieve(session *Session_v3, messageId string, query *pnrret113.PNRRetrieve) (*pnracc113.PNRReply, *ResponseSOAP4Header, error) {
	soapAction := "PNRRET_11_3_1A"

	reply := new(pnracc113.PNRReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) PNRRetrieve2(session *Session_v3, messageId string, query *pnrret113.PNRRetrieve) (*tnlres001.PNRList, *ResponseSOAP4Header, error) {
	soapAction := "PNRRET_11_3_1A"

	reply := new(tnlres001.PNRList)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) QueueCountPlanner(session *Session_v3, messageId string, query *qcddrq031.QueueCountPlanner) (*qcddrr031.QueueCountPlannerReply, *ResponseSOAP4Header, error) {
	soapAction := "QCDDRQ_03_1_1A"

	reply := new(qcddrr031.QueueCountPlannerReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) QueueCountTotal(session *Session_v3, messageId string, query *qcsdrq031.QueueCountTotal) (*qcsdrr031.QueueCountTotalReply, *ResponseSOAP4Header, error) {
	soapAction := "QCSDRQ_03_1_1A"

	reply := new(qcsdrr031.QueueCountTotalReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) QueueList(session *Session_v3, messageId string, query *qdqlrq111.QueueList) (*qdqlrr111.QueueListReply, *ResponseSOAP4Header, error) {
	soapAction := "QDQLRQ_11_1_1A"

	reply := new(qdqlrr111.QueueListReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) QueueMoveItem(session *Session_v3, messageId string, query *quqmuq031.QueueMoveItem) (*quqmur031.QueueMoveItemReply, *ResponseSOAP4Header, error) {
	soapAction := "QUQMUQ_03_1_1A"

	reply := new(quqmur031.QueueMoveItemReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) QueuePlacePNR(session *Session_v3, messageId string, query *quqpcq031.QueuePlacePNR) (*quqpcr031.QueuePlacePNRReply, *ResponseSOAP4Header, error) {
	soapAction := "QUQPCQ_03_1_1A"

	reply := new(quqpcr031.QueuePlacePNRReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) QueueRemoveItem(session *Session_v3, messageId string, query *quqmdq031.QueueRemoveItem) (*quqmdr031.QueueRemoveItemReply, *ResponseSOAP4Header, error) {
	soapAction := "QUQMDQ_03_1_1A"

	reply := new(quqmdr031.QueueRemoveItemReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) SecuritySignOut(session *Session_v3, messageId string, query *vlssoq041.SecuritySignOut) (*vlssor041.SecuritySignOutReply, *ResponseSOAP4Header, error) {
	soapAction := "VLSSOQ_04_1_1A"

	reply := new(vlssor041.SecuritySignOutReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) TicketCreateTSTFromPricing(session *Session_v3, messageId string, query *tautcq041.TicketCreateTSTFromPricing) (*tautcr041.TicketCreateTSTFromPricingReply, *ResponseSOAP4Header, error) {
	soapAction := "TAUTCQ_04_1_1A"

	reply := new(tautcr041.TicketCreateTSTFromPricingReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) TicketCreditCardCheck(session *Session_v3, messageId string, query *ccvrqt061.TicketCreditCardCheck) (*ccvrsp061.TicketCreditCardCheckReply, *ResponseSOAP4Header, error) {
	soapAction := "CCVRQT_06_1_1A"

	reply := new(ccvrsp061.TicketCreditCardCheckReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) TicketDeleteTST(session *Session_v3, messageId string, query *ttstdq041.TicketDeleteTST) (*ttstdr041.TicketDeleteTSTReply, *ResponseSOAP4Header, error) {
	soapAction := "TTSTDQ_04_1_1A"

	reply := new(ttstdr041.TicketDeleteTSTReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}

func (service *WebServicesPTSOAP4Header) TicketDisplayTST(session *Session_v3, messageId string, query *ttstrq071.TicketDisplayTST) (*ttstrr071.TicketDisplayTSTReply, *ResponseSOAP4Header, error) {
	soapAction := "TTSTRQ_07_1_1A"

	reply := new(ttstrr071.TicketDisplayTSTReply)
	header, err := service.client.Call(url + soapAction, messageId, query, reply, session)
	if err != nil {
		return nil, header, err
	}

	return reply, header, nil
}
