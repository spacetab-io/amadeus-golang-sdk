package queue_removeitem_reply

//import "encoding/xml"

type QueueRemoveItemReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMDR_03_1_1A Queue_RemoveItemReply"`

	ErrorReturn *ErrorReturn `xml:"errorReturn,omitempty"`  // minOccurs="0"

	// good response to the message
	GoodResponse *ResponseAnalysisDetailsTypeI `xml:"goodResponse,omitempty"`  // minOccurs="0"
}

type ErrorReturn struct {

	// returns the error code
	ErrorDefinition *ApplicationErrorInformationTypeI `xml:"errorDefinition"`

	// contains the text of the error
	ErrorText *FreeTextInformationType `xml:"errorText,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

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

type FreeTextDetailsType struct {

	// qualifier of the following text
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Source Details
	Source string `xml:"source,omitempty"`  // minOccurs="0"

	// Encoding Informations
	Encoding string `xml:"encoding,omitempty"`  // minOccurs="0"
}

type FreeTextInformationType struct {

	// contains only the qualifier.
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`  // minOccurs="0"

	// Free text
	FreeText []string `xml:"freeText"`  // maxOccurs="99"
}

type ResponseAnalysisDetailsTypeI struct {

	// type of response
	ResponseType string `xml:"responseType"`

	// will only be used in a good response
	StatusCode string `xml:"statusCode"`
}
