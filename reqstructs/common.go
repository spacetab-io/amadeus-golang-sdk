package reqstructs

import (
	"encoding/xml"

	"github.com/kidem/amadeus-ws-go/formats"
)

type AccountingElementType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A AccountingElementType"`

	// Account number
	Number *formats.AlphaNumericStringLength1To10 `xml:"number,omitempty"`
}

type AccountingInformationElementType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A AccountingInformationElementType"`

	// One of these 4 data elements is mandatory , but non in particular
	Account *AccountingElementType `xml:"account,omitempty"`
}

type ActionDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ActionDetailsTypeI"`

	// used for scrollling purposes
	NumberOfItemsDetails *ProcessingInformationTypeI `xml:"numberOfItemsDetails,omitempty"`
}

type AdditionalBusinessSourceInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A AdditionalBusinessSourceInformationTypeI"`

	// the office we are targetting
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType,omitempty"`

	// contains the office ID
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails,omitempty"`
}

type CompanyIdentificationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A CompanyIdentificationTypeI"`

	// Marketing company.
	MarketingCompany *formats.AlphaNumericStringLength1To3 `xml:"marketingCompany,omitempty"`
}

type DummySegmentTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A DummySegmentTypeI"`
}

type FrequentTravellerIdentificationCodeType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A FrequentTravellerIdentificationCodeType"`

	// Frequent Traveller Info. Repetition 2 is used only in the case we provide a customer value range (only one is accepted).
	FrequentTravellerDetails *FrequentTravellerIdentificationType `xml:"frequentTravellerDetails,omitempty"`

	DummyNET struct {
	} `xml:"Dummy.NET,omitempty"`
}

type FrequentTravellerIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A FrequentTravellerIdentificationType"`

	// This field specifies the Tier Level.   This is a 4 letter string indicating the airline's ranking of frequent flyers. It is not to be confused with Alliance tier.  If tierLevel is filled in a given FTI segment, customerValue must not be filled.
	TierLevel *formats.AlphaNumericStringLength1To4 `xml:"tierLevel,omitempty"`

	// This field specifies the Customer value.   This is a 4 letter string indicating the airline's ranking of frequent flyers. It is not to be confused with Alliance tier.  If customerValue is filled in a given FTI segment, tierLevel field must not be filled.
	CustomerValue *formats.NumericIntegerLength1To4 `xml:"customerValue,omitempty"`
}

type LocationTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A LocationTypeU"`

	// Office identification. It can contain wildcards.
	Name *formats.AlphaNumericStringLength1To9 `xml:"name,omitempty"`
}

type OriginAndDestinationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A OriginAndDestinationDetailsTypeI"`

	// Board point
	Origin *formats.AlphaNumericStringLength3To3 `xml:"origin,omitempty"`

	// Off point
	Destination *formats.AlphaNumericStringLength3To3 `xml:"destination,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A OriginatorIdentificationDetailsTypeI"`

	// the office that is being targetted
	InHouseIdentification1 *formats.AlphaNumericStringLength1To9 `xml:"inHouseIdentification1,omitempty"`
}

type PartyIdentifierTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A PartyIdentifierTypeU"`

	// GDS identifier: 1A, 1S, 1G.
	PartyIdentifier *formats.AlphaNumericStringLength1To3 `xml:"partyIdentifier,omitempty"`
}

type PointOfSaleInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A PointOfSaleInformationType"`

	// Party identification.
	PointOfSale *PartyIdentifierTypeU `xml:"pointOfSale,omitempty"`

	// Office id in case the party identifier is 1A.
	LocationDetails *LocationTypeU `xml:"locationDetails,omitempty"`
}

type ProcessingInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ProcessingInformationTypeI"`

	// determine if move up or move down required
	ActionQualifier *formats.AlphaNumericStringLength1To3 `xml:"actionQualifier,omitempty"`
}

type ProductDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ProductDetailsTypeI"`

	// Class designator.
	Designator *formats.AlphaNumericStringLength1To1 `xml:"designator,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ProductIdentificationDetailsTypeI"`

	// Flight number.
	FlightNumber *formats.AlphaNumericStringLength1To4 `xml:"flightNumber,omitempty"`
}

type ProductInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ProductInformationTypeI"`

	// Booking class details.
	BookingClassDetails *ProductDetailsTypeI `xml:"bookingClassDetails,omitempty"`
}

type QueueInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A QueueInformationDetailsTypeI"`

	// queue number
	Number *formats.NumericIntegerLength1To2 `xml:"number,omitempty"`
}

type QueueInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A QueueInformationTypeI"`

	// queue identification
	QueueDetails *QueueInformationDetailsTypeI `xml:"queueDetails,omitempty"`
}

type RangeDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A RangeDetailsTypeI"`

	// define is a range or not
	RangeQualifier *formats.AlphaNumericStringLength1To3 `xml:"rangeQualifier,omitempty"`

	// define the start and possible end point of the scan
	RangeDetails *RangeTypeI `xml:"rangeDetails,omitempty"`
}

type RangeTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A RangeTypeI"`

	// starting point of the scan
	Min *formats.NumericIntegerLength1To18 `xml:"min,omitempty"`

	// ending point of the scan
	Max *formats.NumericIntegerLength1To18 `xml:"max,omitempty"`
}

type RelatedProductInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A RelatedProductInformationTypeI"`

	// Status code
	StatusCode *formats.AlphaNumericStringLength2To2 `xml:"statusCode,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SelectionDetailsInformationTypeI"`

	// used to determine if a new start or a continuation Also used for search and sort criteria on the ticketing, departure and creation dates
	Option *formats.AlphaNumericStringLength1To3 `xml:"option,omitempty"`
}

type SelectionDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SelectionDetailsTypeI"`

	// used for search and sort criteria
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails,omitempty"`
}

type SourceTypeDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SourceTypeDetailsTypeI"`

	// not needed - but mandatory field So just stick a 4 in it !!
	SourceQualifier1 *formats.AlphaNumericStringLength1To3 `xml:"sourceQualifier1,omitempty"`
}

type StatusDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StatusDetailsTypeI"`

	// Indicator showing what flight information will be transported.
	Indicator *formats.AlphaNumericStringLength1To3 `xml:"indicator,omitempty"`
}

type StatusTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StatusTypeI"`

	// Flight status details.
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails,omitempty"`
}

type StructuredDateTimeInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StructuredDateTimeInformationType"`

	// used for date range only The date ranges are defined on central system as 1,2,3,4 The actual values of the ranges are set in the office profile
	TimeMode *formats.NumericIntegerLength1To3 `xml:"timeMode,omitempty"`
}

type StructuredDateTimeType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StructuredDateTimeType"`

	// Year number.
	Year *formats.YearYYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month *formats.MonthMM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day *formats.DayNN `xml:"day,omitempty"`
}

type StructuredPeriodInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StructuredPeriodInformationType"`

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType `xml:"endDateTime,omitempty"`
}

type SubQueueInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SubQueueInformationDetailsTypeI"`

	// E for every category    A for cats with items to be worked C for category number N for nickname CN for both category number and nickname numeric for date range
	IdentificationType *formats.AlphaNumericStringLength1To3 `xml:"identificationType,omitempty"`

	// category number
	ItemNumber *formats.AlphaNumericStringLength1To3 `xml:"itemNumber,omitempty"`

	// used for nickname on inbound used for category name on outbound
	ItemDescription *formats.AlphaNumericStringLength1To35 `xml:"itemDescription,omitempty"`
}

type SubQueueInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SubQueueInformationTypeI"`

	// identifies the category or categories.
	SubQueueInfoDetails *SubQueueInformationDetailsTypeI `xml:"subQueueInfoDetails,omitempty"`
}

type TransportIdentifierType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A TransportIdentifierType"`

	// Company identification.
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`

	// Flight details.
	FlightDetails *ProductIdentificationDetailsTypeI `xml:"flightDetails,omitempty"`
}

type SearchCriteria struct {
	// used to specify if ticketing, departure or creation options
	SearchOption *SelectionDetailsTypeI `xml:"searchOption,omitempty"`

	// used to specify the dates to be searched on
	Dates *StructuredPeriodInformationType `xml:"dates,omitempty"`
}

type FlightInformation struct {
	// It transport the type of flight information that will follow.
	FlightInformationType *StatusTypeI `xml:"flightInformationType,omitempty"`

	// Board point or Off Point.
	BoardPointOrOffPoint *OriginAndDestinationDetailsTypeI `xml:"boardPointOrOffPoint,omitempty"`

	// Airline code or Flight Number (in fact, airline + flight number)
	AirlineCodeOrFlightNumber *TransportIdentifierType `xml:"airlineCodeOrFlightNumber,omitempty"`

	// Booking class.
	ClassOfService *ProductInformationTypeI `xml:"classOfService,omitempty"`

	// Segment status code.
	SegmentStatus *RelatedProductInformationTypeI `xml:"segmentStatus,omitempty"`
}

type TravellerInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A TravellerInformationTypeI"`

	// Traveller surname information.
	PaxDetails *TravellerSurnameInformationTypeI `xml:"paxDetails,omitempty"`
}

type TravellerSurnameInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A TravellerSurnameInformationTypeI"`

	// Passenger surname.
	Surname *formats.AlphaNumericStringLength1To70 `xml:"surname,omitempty"`
}

type UserIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A UserIdentificationType"`

	// The last 2 characters of the sine of the agent who placed the PNR on queue.
	Originator *formats.AlphaNumericStringLength1To2 `xml:"originator,omitempty"`
}

type SortCriteria struct {
	// dummy for SDT clash
	Dumbo *DummySegmentTypeI `xml:"dumbo,omitempty"`

	// Determine the order of the display.
	SortOption *SelectionDetailsTypeI `xml:"sortOption,omitempty"`
}
