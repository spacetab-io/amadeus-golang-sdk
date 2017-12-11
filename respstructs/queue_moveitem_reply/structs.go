package queue_moveitem_reply

//import "encoding/xml"

type QueueMoveItemReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMUR_03_1_1A Queue_MoveItemReply"`

	ErrorReturn *ErrorReturn `xml:"errorReturn,omitempty"`  // minOccurs="0"

	GoodResponse *GoodResponse `xml:"goodResponse,omitempty"`  // minOccurs="0"
}

type ErrorReturn struct {

	ErrorDefinition *ErrorDefinition `xml:"errorDefinition"`

	ErrorText *ErrorText `xml:"errorText,omitempty"`  // minOccurs="0"
}

type ErrorDefinition struct {

	ErrorDetails *ErrorDetails `xml:"errorDetails"`
}

type ErrorDetails struct {

	// Format limitations: an..3
	ErrorCode string `xml:"errorCode"`

	// Format limitations: an..3
	ErrorCategory string `xml:"errorCategory,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`  // minOccurs="0"
}

type ErrorText struct {

	FreeTextDetails *FreeTextDetails `xml:"freeTextDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..199
	FreeText []string `xml:"freeText"`  // maxOccurs="99"
}

type FreeTextDetails struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..3
	Source string `xml:"source,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Encoding string `xml:"encoding,omitempty"`  // minOccurs="0"
}

type GoodResponse struct {

	// Format limitations: a1
	ResponseType string `xml:"responseType"`

	// Format limitations: a1
	StatusCode string `xml:"statusCode"`
}
