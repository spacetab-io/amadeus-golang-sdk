package queue_list_reply

import (
	"encoding/xml"

	"github.com/kidem/amadeus-ws-go/formats"
)

type QueueListReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A Queue_ListReply"`

	ErrorReturn struct {

		// returns the error code
		ErrorDefinition *ApplicationErrorInformationTypeI `xml:"errorDefinition,omitempty"`

		// free text for the error
		ErrorText *FreeTextInformationType `xml:"errorText,omitempty"`
	} `xml:"errorReturn,omitempty"`

	QueueView struct {

		// details of who queue placed the PNR
		Agent *AdditionalBusinessSourceInformationType `xml:"agent,omitempty"`

		// queue being displayed
		QueueNumber *QueueInformationTypeI `xml:"queueNumber,omitempty"`

		// category and date range
		CategoryDetails *SubQueueInformationTypeI `xml:"categoryDetails,omitempty"`

		// date range
		Date *StructuredDateTimeInformationType `xml:"date,omitempty"`

		// all 3 occurences are mandatory and show in order number of  1)PNRs on queue 2)PNRs found 3)PNRs scanned
		PnrCount *NumberOfUnitsType `xml:"pnrCount,omitempty"`

		Item struct {

			// surname of the passenger only or GROUP for a group PNR
			PaxName *TravellerInformationTypeI `xml:"paxName,omitempty"`

			// record locator
			RecLoc *ReservationControlInformationTypeI `xml:"recLoc,omitempty"`

			// 1st segment found in the PNR - if there is data to send
			Segment *TravelProductInformationTypeI `xml:"segment,omitempty"`

			// details of who queue placed the PNR
			Agent *AdditionalBusinessSourceInformationType `xml:"agent,omitempty"`

			// contains 1,2 or 3 of the following Queue placement date/time Ticketing date PNR creation date
			Pnrdates *StructuredDateTimeInformationType_181906S `xml:"pnrdates,omitempty"`
		} `xml:"item,omitempty"`
	} `xml:"queueView,omitempty"`
}

type AdditionalBusinessSourceInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A AdditionalBusinessSourceInformationType"`

	// ORIGINATOR DETAILS
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails,omitempty"`
}

type ApplicationErrorDetailTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ApplicationErrorDetailTypeI"`

	// error code
	ErrorCode *formats.AlphaNumericString_Length1To3 `xml:"errorCode,omitempty"`

	// error category
	ErrorCategory *formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`

	// error code owner
	ErrorCodeOwner *formats.AlphaNumericString_Length1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ApplicationErrorInformationTypeI"`

	// error details
	ErrorDetails *ApplicationErrorDetailTypeI `xml:"errorDetails,omitempty"`
}

type CompanyIdentificationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A CompanyIdentificationTypeI"`

	// carrier code or defined codeset
	MarketingCompany *formats.AlphaNumericString_Length1To3 `xml:"marketingCompany,omitempty"`
}

type FreeTextDetailsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A FreeTextDetailsType"`

	// qualifier
	TextSubjectQualifier *formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// source
	Source *formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// encoding
	Encoding *formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A FreeTextInformationType"`

	// free text
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText *formats.AlphaNumericString_Length1To199 `xml:"freeText,omitempty"`
}

type LocationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A LocationTypeI"`

	// board or off point
	TrueLocation *formats.AlphaNumericString_Length1To3 `xml:"trueLocation,omitempty"`
}

type NumberOfUnitDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A NumberOfUnitDetailsTypeI"`

	// numbers of items
	NumberOfUnit *formats.NumericInteger_Length1To5 `xml:"numberOfUnit,omitempty"`
}

type NumberOfUnitsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A NumberOfUnitsType"`

	// Number of Unit Details
	QuantityDetails *NumberOfUnitDetailsTypeI `xml:"quantityDetails,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A OriginatorIdentificationDetailsTypeI"`

	// office ID of the agent who queue placed the PNR
	InHouseIdentification1 *formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification1,omitempty"`

	// agent sign
	InHouseIdentification2 *formats.AlphaNumericString_Length1To2 `xml:"inHouseIdentification2,omitempty"`
}

type ProductDateTimeTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ProductDateTimeTypeI"`

	// departure date
	DepartureDate *formats.AlphaNumericString_Length1To35 `xml:"departureDate,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ProductIdentificationDetailsTypeI"`

	// flight number
	FlightNumber *formats.AlphaNumericString_Length1To35 `xml:"flightNumber,omitempty"`

	// operational suffix
	OperationalSuffix *formats.AlphaNumericString_Length1To3 `xml:"operationalSuffix,omitempty"`
}

type QueueInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A QueueInformationDetailsTypeI"`

	// queue number
	Number *formats.NumericInteger_Length1To2 `xml:"number,omitempty"`
}

type QueueInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A QueueInformationTypeI"`

	// queue identification
	QueueDetails *QueueInformationDetailsTypeI `xml:"queueDetails,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ReservationControlInformationDetailsTypeI"`

	// contains the record locator to be queue placed
	ControlNumber *formats.AlphaNumericString_Length1To8 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ReservationControlInformationTypeI"`

	// contains the record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}

type StructuredDateTimeInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A StructuredDateTimeInformationType"`

	// used for date range only The date ranges are defined on central system as 1,2,3,4 The actual values of the ranges are set in the office profile
	TimeMode *formats.NumericInteger_Length1To3 `xml:"timeMode,omitempty"`
}

type StructuredDateTimeInformationType_181906S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A StructuredDateTimeInformationType_181906S"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	TimeMode *formats.AlphaNumericString_Length1To3 `xml:"timeMode,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A StructuredDateTimeType"`

	// Year number.
	Year *formats.Year_YYYY `xml:"year,omitempty"`

	// Month number of the year (1 being first month, 0 being NULL data)
	Month *formats.NumericInteger_Length1To2 `xml:"month,omitempty"`

	// day number of the month (1 being first day of the month, 0 is null data)
	Day *formats.NumericInteger_Length1To2 `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour *formats.Hour_hH `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes *formats.Minute_mM `xml:"minutes,omitempty"`
}

type SubQueueInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A SubQueueInformationDetailsTypeI"`

	// E for every category    A for cats with items to be worked C for category number N for nickname CN for both category number and nickname numeric for date range
	IdentificationType *formats.AlphaNumericString_Length1To3 `xml:"identificationType,omitempty"`

	// category number
	ItemNumber *formats.AlphaNumericString_Length1To3 `xml:"itemNumber,omitempty"`

	// used for nickname on inbound used for category name on outbound
	ItemDescription *formats.AlphaNumericString_Length1To35 `xml:"itemDescription,omitempty"`
}

type SubQueueInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A SubQueueInformationTypeI"`

	// identifies the category or categories.
	SubQueueInfoDetails *SubQueueInformationDetailsTypeI `xml:"subQueueInfoDetails,omitempty"`
}

type TravelProductInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A TravelProductInformationTypeI"`

	// departure date
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`

	// board point
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	// off point if present
	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`

	// carrier code or segment type
	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`

	// flight number and suffix
	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`
}

type TravellerInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A TravellerInformationTypeI"`

	PaxDetails *TravellerSurnameInformationTypeI `xml:"paxDetails,omitempty"`
}

type TravellerSurnameInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A TravellerSurnameInformationTypeI"`

	// surname of the passenger or GROUP for a group PNR
	Surname *formats.AlphaNumericString_Length1To70 `xml:"surname,omitempty"`

	Type_ *formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	Quantity *formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"`
}
