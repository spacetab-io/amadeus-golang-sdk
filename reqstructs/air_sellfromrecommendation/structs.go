package air_sellfromrecommendation

import "encoding/xml"

type AirSellFromRecommendation struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/ITAREQ_05_2_IA Air_SellFromRecommendation"`

	MessageActionDetails *MessageActionDetails `xml:"messageActionDetails,omitempty"`  // minOccurs="0"

	RecordLocator *RecordLocator `xml:"recordLocator,omitempty"`  // minOccurs="0"

	ItineraryDetails []*ItineraryDetails `xml:"itineraryDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type MessageActionDetails struct {

	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`  // minOccurs="0"
}

type MessageFunctionDetails struct {

	MessageFunction string `xml:"messageFunction,omitempty"`  // minOccurs="0"

	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type RecordLocator struct {

	Reservation []*Reservation `xml:"reservation,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type Reservation struct {

	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	ControlNumber string `xml:"controlNumber,omitempty"`  // minOccurs="0"

	ControlType string `xml:"controlType,omitempty"`  // minOccurs="0"

	Date string `xml:"date,omitempty"`  // minOccurs="0"

	Time *float64 `xml:"time,omitempty"`  // minOccurs="0"
}

type ItineraryDetails struct {

	OriginDestinationDetails *OriginDestinationDetails `xml:"originDestinationDetails"`

	Message *Message `xml:"message,omitempty"`  // minOccurs="0"

	SegmentInformation []*SegmentInformation `xml:"segmentInformation"`  // maxOccurs="9"
}

type OriginDestinationDetails struct {

	Origin string `xml:"origin,omitempty"`  // minOccurs="0"

	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type Message struct {

	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`  // minOccurs="0"
}

type SegmentInformation struct {

	TravelProductInformation *TravelProductInformation `xml:"travelProductInformation"`

	RelatedproductInformation *RelatedproductInformation `xml:"relatedproductInformation"`
}

type TravelProductInformation struct {

	FlightDate *FlightDate `xml:"flightDate"`

	BoardPointDetails *BoardPointDetails `xml:"boardPointDetails"`

	OffpointDetails *OffpointDetails `xml:"offpointDetails"`

	CompanyDetails *CompanyDetails `xml:"companyDetails"`

	FlightIdentification *FlightIdentification `xml:"flightIdentification"`

	FlightTypeDetails *FlightTypeDetails `xml:"flightTypeDetails,omitempty"`  // minOccurs="0"

	SpecialSegment string `xml:"specialSegment,omitempty"`  // minOccurs="0"

	MarriageDetails []*MarriageDetails `xml:"marriageDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FlightDate struct {

	DepartureDate string `xml:"departureDate,omitempty"`  // minOccurs="0"

	DepartureTime *float64 `xml:"departureTime,omitempty"`  // minOccurs="0"

	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"

	ArrivalTime *float64 `xml:"arrivalTime,omitempty"`  // minOccurs="0"

	DateVariation string `xml:"dateVariation,omitempty"`  // minOccurs="0"
}

type BoardPointDetails struct {

	TrueLocationId string `xml:"trueLocationId"`
}

type OffpointDetails struct {

	TrueLocationId string `xml:"trueLocationId"`
}

type CompanyDetails struct {

	MarketingCompany string `xml:"marketingCompany"`
}

type FlightIdentification struct {

	FlightNumber string `xml:"flightNumber"`

	BookingClass string `xml:"bookingClass"`

	OperationalSuffix string `xml:"operationalSuffix,omitempty"`  // minOccurs="0"
}

type FlightTypeDetails struct {

	FlightIndicator []string `xml:"flightIndicator"`  // maxOccurs="9"
}

type MarriageDetails struct {

	Relation string `xml:"relation,omitempty"`  // minOccurs="0"

	MarriageIdentifier *float64 `xml:"marriageIdentifier,omitempty"`  // minOccurs="0"

	LineNumber *float64 `xml:"lineNumber,omitempty"`  // minOccurs="0"
}

type RelatedproductInformation struct {

	Quantity float64 `xml:"quantity"`

	StatusCode []string `xml:"statusCode,omitempty"`  // minOccurs="0" maxOccurs="10"
}
