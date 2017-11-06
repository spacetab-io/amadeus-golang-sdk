package queue_removeitem

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type QueueRemoveItem struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A Queue_RemoveItem"`

	// determine the target of the removal ordinary queue Ticketing queue option queue
	RemovalOption *SelectionDetailsTypeI `xml:"removalOption,omitempty"`

	TargetDetails *TargetDetails `xml:"targetDetails,omitempty"`
}

type TargetDetails struct {

	// used to specify the target office for which the queue removal is to be done
	TargetOffice *AdditionalBusinessSourceInformationType `xml:"targetOffice,omitempty"`

	// used to specify the queue if required
	QueueNumber *QueueInformationTypeI `xml:"queueNumber,omitempty"`

	// used to select the category
	CategoryDetails *SubQueueInformationTypeI `xml:"categoryDetails,omitempty"`

	// used to pass the target date/time if not default
	PlacementDate *StructuredDateTimeInformationType `xml:"placementDate,omitempty"`

	// contains the record locators to be removed from the queue being targeted
	RecordLocator *ReservationControlInformationTypeI `xml:"recordLocator,omitempty"`
}

type AdditionalBusinessSourceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A AdditionalBusinessSourceInformationType"`

	// SOURCE TYPE
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType,omitempty"`

	// ORIGINATOR DETAILS
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A OriginatorIdentificationDetailsTypeI"`

	// the office that is being targetted
	InHouseIdentification1 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification1,omitempty"`
}

type QueueInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A QueueInformationDetailsTypeI"`

	// queue number
	Number formats.NumericInteger_Length1To2 `xml:"number,omitempty"`
}

type QueueInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A QueueInformationTypeI"`

	// queue identification
	QueueDetails *QueueInformationDetailsTypeI `xml:"queueDetails,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A ReservationControlInformationDetailsTypeI"`

	// contains the record locator to be queue placed
	ControlNumber formats.AlphaNumericString_Length1To8 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A ReservationControlInformationTypeI"`

	// contains the record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A SelectionDetailsInformationTypeI"`

	// removal option
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`
}

type SelectionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A SelectionDetailsTypeI"`

	// specify the option for removal
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails,omitempty"`
}

type SourceTypeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A SourceTypeDetailsTypeI"`

	// to define if own office or different office being targetted
	SourceQualifier1 formats.AlphaNumericString_Length1To3 `xml:"sourceQualifier1,omitempty"`
}

type StructuredDateTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A StructuredDateTimeInformationType"`

	// used to select the date range
	TimeMode formats.NumericInteger_Length1To1 `xml:"timeMode,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A StructuredDateTimeType"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour formats.Hour_hH `xml:"hour,omitempty"`
}

type SubQueueInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A SubQueueInformationDetailsTypeI"`

	// E for every category    A for cats with items to be worked C for category number N for nickname CN for both category number and nickname numeric for date range
	IdentificationType formats.AlphaNumericString_Length1To3 `xml:"identificationType,omitempty"`

	// category number
	ItemNumber formats.AlphaNumericString_Length1To3 `xml:"itemNumber,omitempty"`

	// used for nickname on inbound used for category name on outbound
	ItemDescription formats.AlphaNumericString_Length1To35 `xml:"itemDescription,omitempty"`
}

type SubQueueInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDQ_03_1_1A SubQueueInformationTypeI"`

	// identifies the category or categories.
	SubQueueInfoDetails *SubQueueInformationDetailsTypeI `xml:"subQueueInfoDetails,omitempty"`
}
