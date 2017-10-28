package security_signout_reply

import (
	"encoding/xml"

	"github.com/kidem/amadeus-ws-go/formats"
)

type SecuritySignOutReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A Security_SignOutReply"`

	ErrorSection *ErrorSection `xml:"errorSection,omitempty"`

	// This segment is only used if process is OK. In that case P is specified.
	ProcessStatus *ResponseAnalysisDetailsType `xml:"processStatus,omitempty"`
}

type ErrorSection struct {
	// Application Error
	ApplicationError *ApplicationErrorInformationType `xml:"applicationError,omitempty"`

	// Supplementary Info on the Error.
	InteractiveFreeText *InteractiveFreeTextTypeI `xml:"interactiveFreeText,omitempty"`
}

type ApplicationErrorDetailType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A ApplicationErrorDetailType"`

	// Code identifying the data validation error condition.
	ErrorCode formats.AlphaNumericString_Length1To5 `xml:"errorCode,omitempty"`

	// Identification of a code list.
	ErrorCategory formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner formats.AlphaNumericString_Length1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A ApplicationErrorInformationType"`

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails,omitempty"`
}

type FreeTextQualificationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A FreeTextQualificationTypeI"`

	// Subject
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Info Type
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// Language
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A InteractiveFreeTextTypeI"`

	// Free Text Qualifier
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`

	// Free Text
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type ResponseAnalysisDetailsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A ResponseAnalysisDetailsType"`

	// P must be specified when status of the process is OK.
	StatusCode formats.AlphaString_Length1To6 `xml:"statusCode,omitempty"`
}
