package queue_counttotal

import "encoding/xml"

type QueueCountTotal struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QCSDRQ_03_1_1A Queue_CountTotal"`

	QueueingOptions *QueueingOptions `xml:"queueingOptions"`

	TargetOffice *TargetOffice `xml:"targetOffice,omitempty"`  // minOccurs="0"

	QueueNumber *QueueNumber `xml:"queueNumber,omitempty"`  // minOccurs="0"

	CategorySelection *CategorySelection `xml:"categorySelection,omitempty"`  // minOccurs="0"
}

type QueueingOptions struct {

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

type QueueNumber struct {

	QueueDetails *QueueDetails `xml:"queueDetails"`
}

type QueueDetails struct {

	Number float64 `xml:"number"`
}

type CategorySelection struct {

	SubQueueInfoDetails *SubQueueInfoDetails `xml:"subQueueInfoDetails"`
}

type SubQueueInfoDetails struct {

	IdentificationType string `xml:"identificationType"`

	ItemNumber string `xml:"itemNumber,omitempty"`  // minOccurs="0"

	ItemDescription string `xml:"itemDescription,omitempty"`  // minOccurs="0"
}
