package security_signout_reply

//import "encoding/xml"

type SecuritySignOutReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A Security_SignOutReply"`

	ErrorSection *ErrorSection `xml:"errorSection,omitempty"`  // minOccurs="0"

	// This segment is only used if process is OK. In that case P is specified.
	ProcessStatus *ResponseAnalysisDetailsType `xml:"processStatus,omitempty"`  // minOccurs="0"
}

type ErrorSection struct {

	// Application Error
	ApplicationError *ApplicationErrorInformationType `xml:"applicationError"`

	// Supplementary Info on the Error.
	InteractiveFreeText *InteractiveFreeTextTypeI `xml:"interactiveFreeText,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type ApplicationErrorDetailType struct {

	// Code identifying the data validation error condition.
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	ErrorCategory string `xml:"errorCategory,omitempty"`  // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`  // minOccurs="0"
}

type ApplicationErrorInformationType struct {

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type FreeTextQualificationTypeI struct {

	// Subject
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Info Type
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// Language
	Language string `xml:"language,omitempty"`  // minOccurs="0"
}

type InteractiveFreeTextTypeI struct {

	// Free Text Qualifier
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Free Text
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ResponseAnalysisDetailsType struct {

	// P must be specified when status of the process is OK.
	StatusCode string `xml:"statusCode"`
}
