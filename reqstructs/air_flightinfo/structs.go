package air_flightinfo

import "encoding/xml"

type AirFlightInfo struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/FLIREQ_05_1_1A Air_FlightInfo"`

	GeneralFlightInfo *GeneralFlightInfo `xml:"generalFlightInfo"`
}

type GeneralFlightInfo struct {

	FlightDate *FlightDate `xml:"flightDate,omitempty"`  // minOccurs="0"

	BoardPointDetails *BoardPointDetails `xml:"boardPointDetails,omitempty"`  // minOccurs="0"

	OffPointDetails *OffPointDetails `xml:"offPointDetails,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	FlightIdentification *FlightIdentification `xml:"flightIdentification,omitempty"`  // minOccurs="0"

	FlightTypeDetails *FlightTypeDetails `xml:"flightTypeDetails,omitempty"`  // minOccurs="0"

	MarriageDetails []*MarriageDetails `xml:"marriageDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FlightDate struct {

	DepartureDate string `xml:"departureDate,omitempty"`  // minOccurs="0"

	DepartureTime string `xml:"departureTime,omitempty"`  // minOccurs="0"

	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"

	ArrivalTime string `xml:"arrivalTime,omitempty"`  // minOccurs="0"

	DateVariation string `xml:"dateVariation,omitempty"`  // minOccurs="0"
}

type BoardPointDetails struct {

	TrueLocationId string `xml:"trueLocationId,omitempty"`  // minOccurs="0"
}

type OffPointDetails struct {

	TrueLocationId string `xml:"trueLocationId,omitempty"`  // minOccurs="0"
}

type CompanyDetails struct {

	MarketingCompany string `xml:"marketingCompany"`

	OperatingCompany string `xml:"operatingCompany,omitempty"`  // minOccurs="0"
}

type FlightIdentification struct {

	FlightNumber string `xml:"flightNumber"`

	OperationalSuffix string `xml:"operationalSuffix,omitempty"`  // minOccurs="0"
}

type FlightTypeDetails struct {

	FlightIndicator []string `xml:"flightIndicator,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type MarriageDetails struct {

	Relation string `xml:"relation,omitempty"`  // minOccurs="0"

	MarriageIdentifier *float64 `xml:"marriageIdentifier,omitempty"`  // minOccurs="0"

	LineNumber *float64 `xml:"lineNumber,omitempty"`  // minOccurs="0"

	OtherRelation string `xml:"otherRelation,omitempty"`  // minOccurs="0"

	CarrierCode string `xml:"carrierCode,omitempty"`  // minOccurs="0"
}
