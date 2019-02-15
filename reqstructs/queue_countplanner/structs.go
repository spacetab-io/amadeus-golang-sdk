package queue_countplanner

import "encoding/xml"

type QueueCountPlanner struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QCDDRQ_03_1_1A Queue_CountPlanner"`

	DisplayType *DisplayType `xml:"displayType"`

	TargetOffice *TargetOffice `xml:"targetOffice,omitempty"`  // minOccurs="0"

	Date *Date `xml:"date,omitempty"`  // minOccurs="0"

	QueueNumber *QueueNumber `xml:"queueNumber,omitempty"`  // minOccurs="0"

	CategoryDetails *CategoryDetails `xml:"categoryDetails,omitempty"`  // minOccurs="0"
}

type DisplayType struct {

	SelectionDetails *SelectionDetails `xml:"selectionDetails"`
}

type SelectionDetails struct {

	Option string `xml:"option"`
}

type TargetOffice struct {

	SourceType *SourceType `xml:"sourceType"`

	OriginatorDetails *OriginatorDetails `xml:"originatorDetails"`
}

type SourceType struct {

	SourceQualifier1 string `xml:"sourceQualifier1"`
}

type OriginatorDetails struct {

	InHouseIdentification1 string `xml:"inHouseIdentification1"`
}

type Date struct {

	TimeMode string `xml:"timeMode,omitempty"`  // minOccurs="0"

	DateTime *DateTime `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type DateTime struct {

	Year float64 `xml:"year"`

	Month float64 `xml:"month"`

	Day float64 `xml:"day"`
}

type QueueNumber struct {

	QueueDetails *QueueDetails `xml:"queueDetails"`
}

type QueueDetails struct {

	Number float64 `xml:"number"`
}

type CategoryDetails struct {

	SubQueueInfoDetails *SubQueueInfoDetails `xml:"subQueueInfoDetails"`
}

type SubQueueInfoDetails struct {

	IdentificationType string `xml:"identificationType"`

	ItemNumber string `xml:"itemNumber,omitempty"`  // minOccurs="0"

	ItemDescription string `xml:"itemDescription,omitempty"`  // minOccurs="0"
}
