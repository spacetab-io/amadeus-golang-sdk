package queue_counttotal_reply

//import "encoding/xml"

type QueueCountTotalReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QCSDRR_03_1_1A Queue_CountTotalReply"`

	ErrorReturn *ErrorReturn `xml:"errorReturn,omitempty"`  // minOccurs="0"

	QueueCountDisplay *QueueCountDisplay `xml:"queueCountDisplay,omitempty"`  // minOccurs="0"
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

type QueueCountDisplay struct {

	LocalDateAndTime *LocalDateAndTime `xml:"localDateAndTime"`

	OfficeBeingDisplayed *OfficeBeingDisplayed `xml:"officeBeingDisplayed"`

	DateRanges []*DateRanges `xml:"dateRanges,omitempty"`  // minOccurs="0" maxOccurs="4"

	DelayedCount []*DelayedCount `xml:"delayedCount,omitempty"`  // minOccurs="0" maxOccurs="2"

	StandardQueueCountDisplay []*StandardQueueCountDisplay `xml:"standardQueueCountDisplay"`  // maxOccurs="60"
}

type LocalDateAndTime struct {

	DateTime *DateTime `xml:"dateTime"`
}

type DateTime struct {

	// Format limitations: n..6
	Year float64 `xml:"year"`

	// Format limitations: n..2
	Month float64 `xml:"month"`

	// Format limitations: n..2
	Day float64 `xml:"day"`

	// Format limitations: n..6
	Hour float64 `xml:"hour"`

	// Format limitations: n..6
	Minutes float64 `xml:"minutes"`
}

type OfficeBeingDisplayed struct {

	SourceType *SourceType `xml:"sourceType"`

	OriginatorDetails *OriginatorDetails `xml:"originatorDetails"`
}

type SourceType struct {

	// Format limitations: an..3
	SourceQualifier1 string `xml:"sourceQualifier1"`
}

type OriginatorDetails struct {

	// Format limitations: an..9
	InHouseIdentification1 string `xml:"inHouseIdentification1"`
}

type DateRanges struct {

	BeginDateTime *BeginDateTime `xml:"beginDateTime"`

	EndDateTime *EndDateTime `xml:"endDateTime"`
}

type BeginDateTime struct {

	// Format limitations: n..6
	Year float64 `xml:"year"`

	// Format limitations: n..2
	Month float64 `xml:"month"`

	// Format limitations: n..2
	Day float64 `xml:"day"`
}

type EndDateTime struct {

	// Format limitations: n..6
	Year float64 `xml:"year"`

	// Format limitations: n..2
	Month float64 `xml:"month"`

	// Format limitations: n..2
	Day float64 `xml:"day"`
}

type DelayedCount struct {

	QueueDetails *QueueDetails `xml:"queueDetails"`
}

type QueueDetails struct {

	// Format limitations: n..15
	NumberOfItems float64 `xml:"numberOfItems"`

	// Format limitations: an..3
	Status string `xml:"status,omitempty"`  // minOccurs="0"
}

type StandardQueueCountDisplay struct {

	QueueName *QueueName `xml:"queueName"`

	QueueNumber *QueueNumber `xml:"queueNumber"`

	CategoryAndCount []*CategoryAndCount `xml:"categoryAndCount"`  // maxOccurs="1020"
}

type QueueName struct {

	ReferenceDetails *ReferenceDetails `xml:"referenceDetails"`
}

type ReferenceDetails struct {

	// Format limitations: an..35
	Value string `xml:"value"`
}

type QueueNumber struct {

	QueueDetails *QueueDetails1 `xml:"queueDetails"`
}

type QueueDetails1 struct {

	// Format limitations: n..2
	Number float64 `xml:"number"`
}

type CategoryAndCount struct {

	CategoryAndDateRange *CategoryAndDateRange `xml:"categoryAndDateRange"`

	QueueCount []*QueueCount `xml:"queueCount"`  // maxOccurs="4"
}

type CategoryAndDateRange struct {

	SubQueueInfoDetails *SubQueueInfoDetails `xml:"subQueueInfoDetails"`
}

type SubQueueInfoDetails struct {

	// Format limitations: an..3
	IdentificationType string `xml:"identificationType"`

	// Format limitations: an..3
	ItemNumber string `xml:"itemNumber,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	ItemDescription string `xml:"itemDescription,omitempty"`  // minOccurs="0"
}

type QueueCount struct {

	QueueDetails *QueueDetails `xml:"queueDetails"`
}
