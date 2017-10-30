package docissuance_issueticket_reply

import (
	"encoding/xml"

	"github.com/kidem/amadeus-ws-go/formats"
)

type DocIssuanceIssueTicketReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A DocIssuance_IssueTicketReply"`

	// The type of the answer : warning, aknowledgement, PNR display...
	ProcessingStatus *ResponseAnalysisDetailsType `xml:"processingStatus,omitempty"`

	// Contains warning and errors.
	ErrorGroup *ErrorGroupType `xml:"errorGroup,omitempty"`
}

type ApplicationErrorDetailType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A ApplicationErrorDetailType"`

	// Code identifying the data validation error condition.
	ErrorCode formats.AlphaNumericString_Length1To5 `xml:"errorCode,omitempty"`

	// Identification of a code list.
	ErrorCategory formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner formats.AlphaNumericString_Length1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A ApplicationErrorInformationType"`

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails,omitempty"`
}

type ErrorGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A ErrorGroupType"`

	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType `xml:"errorOrWarningCodeDetails,omitempty"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"`
}

type FreeTextDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A FreeTextDetailsType"`

	// textSubjectQualifier
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// source
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// encoding
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A FreeTextInformationType"`

	// warning/error message details.
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText formats.AlphaNumericString_Length1To199 `xml:"freeText,omitempty"`
}

type ResponseAnalysisDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A ResponseAnalysisDetailsType"`

	// type of the answer
	StatusCode formats.AlphaString_Length1To6 `xml:"statusCode,omitempty"`
}
