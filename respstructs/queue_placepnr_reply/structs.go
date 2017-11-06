package queue_placepnr_reply

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type QueuePlacePNRReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QUQPCR_03_1_1A Queue_PlacePNRReply"`

	ErrorReturn *ErrorReturn `xml:"errorReturn,omitempty"`

	// record locator
	RecordLocator *ReservationControlInformationTypeI `xml:"recordLocator,omitempty"`
}

type ErrorReturn struct {

	// returns the error code
	ErrorDefinition *ApplicationErrorInformationTypeI `xml:"errorDefinition,omitempty"`

	// contains the text of the error
	ErrorText *FreeTextInformationType `xml:"errorText,omitempty"`
}

type ApplicationErrorDetailTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQPCR_03_1_1A ApplicationErrorDetailTypeI"`

	// error code
	ErrorCode formats.AlphaNumericString_Length1To3 `xml:"errorCode,omitempty"`

	// error category
	ErrorCategory formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`

	// error code owner
	ErrorCodeOwner formats.AlphaNumericString_Length1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQPCR_03_1_1A ApplicationErrorInformationTypeI"`

	// error details
	ErrorDetails *ApplicationErrorDetailTypeI `xml:"errorDetails,omitempty"`
}

type FreeTextDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQPCR_03_1_1A FreeTextDetailsType"`

	// qualifier of the following text
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Source Details
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// Encoding Informations
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQPCR_03_1_1A FreeTextInformationType"`

	// contains only the qualifier.
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text
	FreeText formats.AlphaNumericString_Length1To199 `xml:"freeText,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQPCR_03_1_1A ReservationControlInformationDetailsTypeI"`

	// contains the record locator to be queue placed
	ControlNumber formats.AlphaNumericString_Length1To8 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQPCR_03_1_1A ReservationControlInformationTypeI"`

	// contains the record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}
