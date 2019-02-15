package DocIssuance_IssueTicketReply_v09_1 // ttktir091

//import "encoding/xml"

type DocIssuanceIssueTicketReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A DocIssuance_IssueTicketReply"`

	// The type of the answer : warning, aknowledgement, PNR display...
	ProcessingStatus *ResponseAnalysisDetailsType `xml:"processingStatus"`

	// Contains warning and errors.
	ErrorGroup *ErrorGroupType `xml:"errorGroup,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type ApplicationErrorDetailType struct {
	// Code identifying the data validation error condition.
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	ErrorCategory string `xml:"errorCategory,omitempty"` // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"` // minOccurs="0"
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type ErrorGroupType struct {
	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType `xml:"errorOrWarningCodeDetails"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"` // minOccurs="0"
}

type FreeTextDetailsType struct {
	// textSubjectQualifier
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// source
	Source string `xml:"source"`

	// encoding
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {
	// warning/error message details.
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText string `xml:"freeText"`
}

type ResponseAnalysisDetailsType struct {
	// type of the answer
	StatusCode string `xml:"statusCode"`
}
