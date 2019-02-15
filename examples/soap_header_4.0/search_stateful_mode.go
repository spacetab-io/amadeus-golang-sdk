package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tmconsulting/amadeus-golang-sdk/examples/soap_header_4.0/helpers"
	fmptbr143 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricertravelboardsearch_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/soap4.0"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"gopkg.in/metakeule/fmtdate.v1"
)

func ActionSearchUsingStatefulMode(session *soap4_0.Session_v3, class, departure, arrival, date string) (*fmptbr143.FareMasterPricerTravelBoardSearchReply, *soap4_0.ResponseSOAP4Header, error) {
	messageId := strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	query := helpers.MakeSearchQuery(class, departure, arrival, date)
	services := helpers.NewWebServices()
	reply, header, err := services.FareMasterPricerTravelBoardSearch(session, messageId, query)
	return reply, header, err
}

func main () {
	var cabinClass = "Y"
	var departureLocationId = "DEL"
	var arrivalLocationId = "FRA"
	date := time.Now().AddDate(0, 0, 7)
	var dateStr = fmtdate.Format("DDMMYY", date)
	session := &soap4_0.Session_v3{
		TransactionStatusCode: soap4_0.TransactionStatusCode[soap4_0.Start],
	}
	_, header, err := ActionSearchUsingStatefulMode(session, cabinClass, departureLocationId, arrivalLocationId, dateStr)
	if err != nil {
		log.Fatal(err)
	} else {
		session = &header.Session
		sequenceNumber, _ := strconv.Atoi(session.SequenceNumber)
		sequenceNumber++
		session.SequenceNumber = strconv.Itoa(sequenceNumber)
		session.TransactionStatusCode = soap4_0.TransactionStatusCode[soap4_0.End]
		dateStr = fmtdate.Format("DDMMYY", date.AddDate(0, 0, 1))
		_, _, err := ActionSearchUsingStatefulMode(session, cabinClass, departureLocationId, arrivalLocationId, dateStr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
