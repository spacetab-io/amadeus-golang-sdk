package security_authenticate_reply

import (
	"encoding/xml"

	"github.com/kidem/amadeus-ws-go/formats"
)

type SecurityAuthenticateReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A Security_AuthenticateReply"`

	ErrorSection *ErrorSection `xml:"errorSection,omitempty"`

	// This segment is only used if process is OK. In that case P is specified.
	ProcessStatus *ResponseAnalysisDetailsType `xml:"processStatus,omitempty"`

	// This segment is used to specify organization details associated with the user.
	OrganizationInfo *OrganizationType `xml:"organizationInfo,omitempty"`

	// Identifier of a group of conversation, shared by sevreal mono-signed conversations.
	ConversationGrp *BusinessProcessIdType `xml:"conversationGrp,omitempty"`
}

type ErrorSection struct {
	// Application Error
	ApplicationError *ApplicationErrorInformationType `xml:"applicationError,omitempty"`

	// Supplementary Info on the Error.
	InteractiveFreeText *InteractiveFreeTextTypeI `xml:"interactiveFreeText,omitempty"`
}

type ApplicationErrorDetailType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A ApplicationErrorDetailType"`

	// Code identifying the data validation error condition.
	ErrorCode *formats.AlphaNumericString_Length1To5 `xml:"errorCode,omitempty"`

	// Identification of a code list.
	ErrorCategory *formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner *formats.AlphaNumericString_Length1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A ApplicationErrorInformationType"`

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails,omitempty"`
}

type BusinessProcessIdType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A BusinessProcessIdType"`

	// A unique reference to identify the process/business
	ProcessIdentifier *formats.AlphaNumericString_Length1To10 `xml:"processIdentifier,omitempty"`
}

type FreeTextQualificationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A FreeTextQualificationTypeI"`

	// Subject
	Subject *formats.AlphaNumericString_Length1To3 `xml:"subject,omitempty"`

	// Info Type
	InfoType *formats.AlphaNumericString_Length1To4 `xml:"infoType,omitempty"`

	// Language
	Language *formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A InteractiveFreeTextTypeI"`

	// Free Text Qualifier
	FreeTextQualif *FreeTextQualificationTypeI `xml:"freeTextQualif,omitempty"`

	// Free Text
	FreeText *formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type OrganizationIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A OrganizationIdentificationType"`

	// Organization label (Company Id).
	Label *formats.AlphaNumericString_Length1To10 `xml:"label,omitempty"`
}

type OrganizationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A OrganizationType"`

	// This composite is used to specify an organization details
	OrganizationDetails *OrganizationIdentificationType `xml:"organizationDetails,omitempty"`
}

type ResponseAnalysisDetailsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A ResponseAnalysisDetailsType"`

	// P must be specified when status of the process is OK.
	StatusCode *formats.AlphaString_Length1To6 `xml:"statusCode,omitempty"`
}
