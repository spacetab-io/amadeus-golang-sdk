package Air_SellFromRecommendationReply_v05_2 // itares052

//import "encoding/xml"

type AirSellFromRecommendationReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/ITARES_05_2_IA Air_SellFromRecommendationReply"`

	Message *Message `xml:"message,omitempty"` // minOccurs="0"

	ErrorAtMessageLevel []*ErrorAtMessageLevel `xml:"errorAtMessageLevel,omitempty"` // minOccurs="0" maxOccurs="5"

	ItineraryDetails []*ItineraryDetails `xml:"itineraryDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type Message struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	ResponseType string `xml:"responseType,omitempty"` // minOccurs="0"
}

type MessageFunctionDetails struct {
	// Format limitations: an..3
	BusinessFunction string `xml:"businessFunction,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	MessageFunction string `xml:"messageFunction,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	ResponsibleAgency string `xml:"responsibleAgency,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"` // minOccurs="0" maxOccurs="20"
}

type ErrorAtMessageLevel struct {
	ErrorSegment *ErrorSegment `xml:"errorSegment"`

	InformationText *InformationText `xml:"informationText,omitempty"` // minOccurs="0"
}

type ErrorSegment struct {
	ErrorDetails *ErrorDetails `xml:"errorDetails"`
}

type ErrorDetails struct {
	// Format limitations: an..5
	ErrorCode string `xml:"errorCode"`

	// Format limitations: an..3
	ErrorCategory string `xml:"errorCategory"`

	// Format limitations: an..3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"` // minOccurs="0"
}

type InformationText struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"` // minOccurs="0"

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"` // minOccurs="0" maxOccurs="99"
}

type FreeTextQualification struct {
	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..4
	InformationType string `xml:"informationType,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Status string `xml:"status,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	CompanyId string `xml:"companyId,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Language string `xml:"language,omitempty"` // minOccurs="0"
}

type ItineraryDetails struct {
	OriginDestination *OriginDestination `xml:"originDestination"`

	ErrorItinerarylevel []*ErrorItinerarylevel `xml:"errorItinerarylevel,omitempty"` // minOccurs="0" maxOccurs="5"

	SegmentInformation []*SegmentInformation `xml:"segmentInformation,omitempty"` // minOccurs="0" maxOccurs="9"
}

type OriginDestination struct {
	// Format limitations: an..3
	Origin string `xml:"origin,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Destination string `xml:"destination,omitempty"` // minOccurs="0"
}

type ErrorItinerarylevel struct {
	ErrorSegment *ErrorSegment `xml:"errorSegment"`

	InformationText *InformationText `xml:"informationText,omitempty"` // minOccurs="0"
}

type SegmentInformation struct {
	FlightDetails *FlightDetails `xml:"flightDetails"`

	ApdSegment *ApdSegment `xml:"apdSegment,omitempty"` // minOccurs="0"

	ActionDetails *ActionDetails `xml:"actionDetails"`

	InformationText *InformationText `xml:"informationText,omitempty"` // minOccurs="0"

	ErrorAtSegmentLevel []*ErrorAtSegmentLevel `xml:"errorAtSegmentLevel,omitempty"` // minOccurs="0" maxOccurs="5"
}

type FlightDetails struct {
	FlightDate *FlightDate `xml:"flightDate,omitempty"` // minOccurs="0"

	BoardPointDetails *BoardPointDetails `xml:"boardPointDetails,omitempty"` // minOccurs="0"

	OffpointDetails *OffpointDetails `xml:"offpointDetails,omitempty"` // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"` // minOccurs="0"

	FlightIdentification *FlightIdentification `xml:"flightIdentification,omitempty"` // minOccurs="0"

	FlightTypeDetails *FlightTypeDetails `xml:"flightTypeDetails,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	SpecialSegment string `xml:"specialSegment,omitempty"` // minOccurs="0"

	MarriageDetails []*MarriageDetails `xml:"marriageDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type FlightDate struct {
	// Format limitations: an..35
	DepartureDate string `xml:"departureDate,omitempty"` // minOccurs="0"

	// Format limitations: n..4
	DepartureTime *float64 `xml:"departureTime,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	ArrivalDate string `xml:"arrivalDate,omitempty"` // minOccurs="0"

	// Format limitations: n..4
	ArrivalTime *float64 `xml:"arrivalTime,omitempty"` // minOccurs="0"

	// Format limitations: n1
	DateVariation string `xml:"dateVariation,omitempty"` // minOccurs="0"
}

type BoardPointDetails struct {
	// Format limitations: an..25
	TrueLocationId string `xml:"trueLocationId,omitempty"` // minOccurs="0"

	// Format limitations: an..17
	TrueLocation string `xml:"trueLocation,omitempty"` // minOccurs="0"
}

type OffpointDetails struct {
	// Format limitations: an..25
	TrueLocationId string `xml:"trueLocationId,omitempty"` // minOccurs="0"

	// Format limitations: an..17
	TrueLocation string `xml:"trueLocation,omitempty"` // minOccurs="0"
}

type CompanyDetails struct {
	// Format limitations: an..35
	MarketingCompany string `xml:"marketingCompany,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	OperatingCompany string `xml:"operatingCompany,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	OtherCompany string `xml:"otherCompany,omitempty"` // minOccurs="0"
}

type FlightIdentification struct {
	// Format limitations: an..4
	FlightNumber string `xml:"flightNumber"`

	// Format limitations: an..2
	BookingClass string `xml:"bookingClass,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	OperationalSuffix string `xml:"operationalSuffix,omitempty"` // minOccurs="0"

	// Format limitations: an..7
	Modifier []string `xml:"modifier,omitempty"` // minOccurs="0" maxOccurs="3"
}

type FlightTypeDetails struct {
	// Format limitations: an..6
	FlightIndicator []string `xml:"flightIndicator"` // maxOccurs="9"
}

type MarriageDetails struct {
	// Format limitations: an..3
	Relation string `xml:"relation,omitempty"` // minOccurs="0"

	// Format limitations: n..10
	MarriageIdentifier *float64 `xml:"marriageIdentifier,omitempty"` // minOccurs="0"

	// Format limitations: n..6
	LineNumber *float64 `xml:"lineNumber,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	OtherRelation string `xml:"otherRelation,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	CarrierCode string `xml:"carrierCode,omitempty"` // minOccurs="0"
}

type ApdSegment struct {
	LegDetails *LegDetails `xml:"legDetails,omitempty"` // minOccurs="0"

	DepartureStationInfo *DepartureStationInfo `xml:"departureStationInfo,omitempty"` // minOccurs="0"

	ArrivalStationInfo *ArrivalStationInfo `xml:"arrivalStationInfo,omitempty"` // minOccurs="0"

	FacilitiesInformation []*FacilitiesInformation `xml:"facilitiesInformation,omitempty"` // minOccurs="0" maxOccurs="10"
}

type LegDetails struct {
	// Format limitations: an..8
	Equipment string `xml:"equipment,omitempty"` // minOccurs="0"

	// Format limitations: n..3
	NumberOfStops *float64 `xml:"numberOfStops,omitempty"` // minOccurs="0"

	// Format limitations: n..6
	Duration *float64 `xml:"duration,omitempty"` // minOccurs="0"

	// Format limitations: n..8
	Percentage *float64 `xml:"percentage,omitempty"` // minOccurs="0"

	// Format limitations: an..7
	DaysOfOperation string `xml:"daysOfOperation,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	DateTimePeriod string `xml:"dateTimePeriod,omitempty"` // minOccurs="0"

	// Format limitations: n1
	ComplexingFlightIndicator string `xml:"complexingFlightIndicator,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	Locations []string `xml:"locations,omitempty"` // minOccurs="0" maxOccurs="3"
}

type DepartureStationInfo struct {
	// Format limitations: an..6
	GateDescription string `xml:"gateDescription,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	Terminal string `xml:"terminal,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	Concourse string `xml:"concourse,omitempty"` // minOccurs="0"
}

type ArrivalStationInfo struct {
	// Format limitations: an..6
	GateDescription string `xml:"gateDescription,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	Terminal string `xml:"terminal,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	Concourse string `xml:"concourse,omitempty"` // minOccurs="0"
}

type FacilitiesInformation struct {
	// Format limitations: an..3
	Code string `xml:"code,omitempty"` // minOccurs="0"

	// Format limitations: an..70
	Description string `xml:"description,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"` // minOccurs="0"

	// Format limitations: an..17
	ExtensionCode []string `xml:"extensionCode,omitempty"` // minOccurs="0" maxOccurs="26"
}

type ActionDetails struct {
	// Format limitations: n..15
	Quantity float64 `xml:"quantity"`

	// Format limitations: an..3
	StatusCode []string `xml:"statusCode"` // maxOccurs="10"
}

type ErrorAtSegmentLevel struct {
	ErrorSegment *ErrorSegment `xml:"errorSegment"`

	InformationText *InformationText `xml:"informationText,omitempty"` // minOccurs="0"
}
