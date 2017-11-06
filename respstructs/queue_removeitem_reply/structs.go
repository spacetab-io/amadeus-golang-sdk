package queue_removeitem_reply

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type QueueRemoveItemReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDR_03_1_1A Queue_RemoveItemReply"`

	ErrorReturn *ErrorReturn `xml:"errorReturn,omitempty"`

	// good response to the message
	GoodResponse *ResponseAnalysisDetailsTypeI `xml:"goodResponse,omitempty"`
}

type ErrorReturn struct {

	// returns the error code
	ErrorDefinition *ApplicationErrorInformationTypeI `xml:"errorDefinition,omitempty"`

	// contains the text of the error
	ErrorText *FreeTextInformationType `xml:"errorText,omitempty"`
}

type ApplicationErrorDetailTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDR_03_1_1A ApplicationErrorDetailTypeI"`

	// error code
	ErrorCode formats.AlphaNumericString_Length1To3 `xml:"errorCode,omitempty"`

	// error category
	ErrorCategory formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`

	// error code owner
	ErrorCodeOwner formats.AlphaNumericString_Length1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDR_03_1_1A ApplicationErrorInformationTypeI"`

	// error details
	ErrorDetails *ApplicationErrorDetailTypeI `xml:"errorDetails,omitempty"`
}

type FreeTextDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDR_03_1_1A FreeTextDetailsType"`

	// qualifier of the following text
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Source Details
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// Encoding Informations
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDR_03_1_1A FreeTextInformationType"`

	// contains only the qualifier.
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text
	FreeText formats.AlphaNumericString_Length1To199 `xml:"freeText,omitempty"`
}

type ResponseAnalysisDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDR_03_1_1A ResponseAnalysisDetailsTypeI"`

	// type of response
	ResponseType formats.AlphaString_Length1To1 `xml:"responseType,omitempty"`

	// will only be used in a good response
	StatusCode formats.AlphaString_Length1To1 `xml:"statusCode,omitempty"`
}
