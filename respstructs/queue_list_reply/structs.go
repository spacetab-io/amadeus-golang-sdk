package queue_list_reply

//import "encoding/xml"

type QueueListReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A Queue_ListReply"`

	ErrorReturn *ErrorReturn `xml:"errorReturn,omitempty"`  // minOccurs="0"

	QueueView *QueueView `xml:"queueView,omitempty"`  // minOccurs="0"
}

type ErrorReturn struct {

	// returns the error code
	ErrorDefinition *ApplicationErrorInformationTypeI `xml:"errorDefinition"`

	// free text for the error
	ErrorText *FreeTextInformationType `xml:"errorText,omitempty"`  // minOccurs="0"
}

type QueueView struct {

	// details of who queue placed the PNR
	Agent *AdditionalBusinessSourceInformationType `xml:"agent"`

	// queue being displayed
	QueueNumber *QueueInformationTypeI `xml:"queueNumber"`

	// category and date range
	CategoryDetails *SubQueueInformationTypeI `xml:"categoryDetails"`

	// date range
	Date *StructuredDateTimeInformationType `xml:"date,omitempty"`  // minOccurs="0"

	// all 3 occurences are mandatory and show in order number of  1)PNRs on queue 2)PNRs found 3)PNRs scanned
	PnrCount []*NumberOfUnitsType `xml:"pnrCount"`  // maxOccurs="3"

	Item []*Item `xml:"item"`  // maxOccurs="1000"
}

type Item struct {

	// surname of the passenger only or GROUP for a group PNR
	PaxName *TravellerInformationTypeI `xml:"paxName"`

	// record locator
	RecLoc *ReservationControlInformationTypeI `xml:"recLoc"`

	// 1st segment found in the PNR - if there is data to send
	Segment *TravelProductInformationTypeI `xml:"segment,omitempty"`  // minOccurs="0"

	// details of who queue placed the PNR
	Agent *AdditionalBusinessSourceInformationType `xml:"agent"`

	// contains 1,2 or 3 of the following Queue placement date/time Ticketing date PNR creation date
	Pnrdates []*StructuredDateTimeInformationType_181906S `xml:"pnrdates"`  // maxOccurs="3"
}

//
// Complex structs
//

type AdditionalBusinessSourceInformationType struct {

	// ORIGINATOR DETAILS
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails"`
}

type ApplicationErrorDetailTypeI struct {

	// error code
	ErrorCode string `xml:"errorCode"`

	// error category
	ErrorCategory string `xml:"errorCategory,omitempty"`  // minOccurs="0"

	// error code owner
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`  // minOccurs="0"
}

type ApplicationErrorInformationTypeI struct {

	// error details
	ErrorDetails *ApplicationErrorDetailTypeI `xml:"errorDetails"`
}

type CompanyIdentificationTypeI struct {

	// carrier code or defined codeset
	MarketingCompany string `xml:"marketingCompany"`
}

type FreeTextDetailsType struct {

	// qualifier
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// source
	Source string `xml:"source"`

	// encoding
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {

	// free text
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText []string `xml:"freeText"`  // maxOccurs="99"
}

type LocationTypeI struct {

	// board or off point
	TrueLocation string `xml:"trueLocation"`
}

type NumberOfUnitDetailsTypeI struct {

	// numbers of items
	NumberOfUnit int32 `xml:"numberOfUnit"`
}

type NumberOfUnitsType struct {

	// Number of Unit Details
	QuantityDetails *NumberOfUnitDetailsTypeI `xml:"quantityDetails"`
}

type OriginatorIdentificationDetailsTypeI struct {

	// office ID of the agent who queue placed the PNR
	InHouseIdentification1 string `xml:"inHouseIdentification1"`

	// agent sign
	InHouseIdentification2 string `xml:"inHouseIdentification2,omitempty"`  // minOccurs="0"
}

type ProductDateTimeTypeI struct {

	// departure date
	DepartureDate string `xml:"departureDate,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetailsTypeI struct {

	// flight number
	FlightNumber string `xml:"flightNumber"`

	// operational suffix
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`  // minOccurs="0"
}

type QueueInformationDetailsTypeI struct {

	// queue number
	Number int32 `xml:"number"`
}

type QueueInformationTypeI struct {

	// queue identification
	QueueDetails *QueueInformationDetailsTypeI `xml:"queueDetails"`
}

type ReservationControlInformationDetailsTypeI struct {

	// contains the record locator to be queue placed
	ControlNumber string `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {

	// contains the record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation"`
}

type StructuredDateTimeInformationType struct {

	// used for date range only The date ranges are defined on central system as 1,2,3,4 The actual values of the ranges are set in the office profile
	TimeMode int32 `xml:"timeMode"`
}

type StructuredDateTimeInformationType_181906S struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	TimeMode string `xml:"timeMode"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime"`
}

type StructuredDateTimeType struct {

	// Year number.
	Year string `xml:"year"`

	// Month number of the year (1 being first month, 0 being NULL data)
	Month int32 `xml:"month"`

	// day number of the month (1 being first day of the month, 0 is null data)
	Day int32 `xml:"day"`

	// Hour between 0 and 23
	Hour string `xml:"hour,omitempty"`  // minOccurs="0"

	// Minutes between 0 and 59
	Minutes string `xml:"minutes,omitempty"`  // minOccurs="0"
}

type SubQueueInformationDetailsTypeI struct {

	// E for every category    A for cats with items to be worked C for category number N for nickname CN for both category number and nickname numeric for date range
	IdentificationType string `xml:"identificationType"`

	// category number
	ItemNumber string `xml:"itemNumber,omitempty"`  // minOccurs="0"

	// used for nickname on inbound used for category name on outbound
	ItemDescription string `xml:"itemDescription,omitempty"`  // minOccurs="0"
}

type SubQueueInformationTypeI struct {

	// identifies the category or categories.
	SubQueueInfoDetails *SubQueueInformationDetailsTypeI `xml:"subQueueInfoDetails"`
}

type TravelProductInformationTypeI struct {

	// departure date
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`  // minOccurs="0"

	// board point
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`  // minOccurs="0"

	// off point if present
	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`  // minOccurs="0"

	// carrier code or segment type
	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`  // minOccurs="0"

	// flight number and suffix
	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`  // minOccurs="0"
}

type TravellerInformationTypeI struct {

	PaxDetails *TravellerSurnameInformationTypeI `xml:"paxDetails"`
}

type TravellerSurnameInformationTypeI struct {

	// surname of the passenger or GROUP for a group PNR
	Surname string `xml:"surname"`

	// Format limitations: an..3
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Format limitations: n..15
	Quantity *int32 `xml:"quantity,omitempty"`  // minOccurs="0"
}
