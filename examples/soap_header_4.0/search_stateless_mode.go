package main

import (
	"log"
	"strings"
	"time"

	"github.com/tmconsulting/amadeus-golang-sdk/examples/soap_header_4.0/helpers"
	fmptbr143 "github.com/tmconsulting/amadeus-golang-sdk/respstructs/fare_masterpricertravelboardsearch_reply"
	"github.com/tmconsulting/amadeus-golang-sdk/utils"
	"gopkg.in/metakeule/fmtdate.v1"
)

func ActionSearchUsingStatelessMode(class, departure, arrival, date string) (*fmptbr143.FareMasterPricerTravelBoardSearchReply, error) {

	messageId := strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	query := helpers.MakeSearchQuery(class, departure, arrival, date)
	services := helpers.NewWebServices()
	response, _, err := services.FareMasterPricerTravelBoardSearch(nil, messageId, query)
	return response, err
}

func main () {
	var cabinClass = "Y"
	var departureLocationId = "DEL"
	var arrivalLocationId = "FRA"
	date := time.Now().AddDate(0, 0, 7)
	var dateStr = fmtdate.Format("DDMMYY", date)
	_, err := ActionSearchUsingStatelessMode(cabinClass, departureLocationId, arrivalLocationId, dateStr)
	if err != nil {
		log.Fatal(err)
	}
}
