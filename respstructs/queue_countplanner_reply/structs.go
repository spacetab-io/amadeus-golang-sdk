package queue_countplanner_reply

//import "encoding/xml"

type QueueCountPlannerReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/QCDDRR_03_1_1A Queue_CountPlannerReply"`

	ErrorReturn *ErrorReturn `xml:"errorReturn,omitempty"`  // minOccurs="0"

	DisplayTarget *DisplayTarget `xml:"displayTarget,omitempty"`  // minOccurs="0"
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

type DisplayTarget struct {

	LocalDateAndTime *LocalDateAndTime `xml:"localDateAndTime"`

	TargetOffice *TargetOffice `xml:"targetOffice"`

	DisplayType *DisplayType `xml:"displayType,omitempty"`  // minOccurs="0"

	HourlyOrDailyCount []*HourlyOrDailyCount `xml:"hourlyOrDailyCount"`  // maxOccurs="33"
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

type TargetOffice struct {

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

type DisplayType struct {

	SelectionDetails *SelectionDetails `xml:"selectionDetails"`
}

type SelectionDetails struct {

	// Format limitations: an..3
	Option string `xml:"option"`
}

type HourlyOrDailyCount struct {

	DateOrTimeRange *DateOrTimeRange `xml:"dateOrTimeRange"`

	Count []*Count `xml:"count"`  // maxOccurs="4"
}

type DateOrTimeRange struct {

	// Format limitations: an..3
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	BeginDateTime *BeginDateTime `xml:"beginDateTime,omitempty"`  // minOccurs="0"

	EndDateTime *EndDateTime `xml:"endDateTime,omitempty"`  // minOccurs="0"
}

type BeginDateTime struct {

	// Format limitations: n..6
	Year *float64 `xml:"year,omitempty"`  // minOccurs="0"

	// Format limitations: n..2
	Month *float64 `xml:"month,omitempty"`  // minOccurs="0"

	// Format limitations: n..2
	Day *float64 `xml:"day,omitempty"`  // minOccurs="0"

	// Format limitations: n..6
	Hour *float64 `xml:"hour,omitempty"`  // minOccurs="0"
}

type EndDateTime struct {

	// Format limitations: n..6
	Year *float64 `xml:"year,omitempty"`  // minOccurs="0"

	// Format limitations: n..2
	Month *float64 `xml:"month,omitempty"`  // minOccurs="0"

	// Format limitations: n..2
	Day *float64 `xml:"day,omitempty"`  // minOccurs="0"

	// Format limitations: n..6
	Hour *float64 `xml:"hour,omitempty"`  // minOccurs="0"
}

type Count struct {

	QueueDetails *QueueDetails `xml:"queueDetails"`
}

type QueueDetails struct {

	// Format limitations: n..15
	NumberOfItems float64 `xml:"numberOfItems"`

	// Format limitations: an..3
	Status string `xml:"status,omitempty"`  // minOccurs="0"
}
