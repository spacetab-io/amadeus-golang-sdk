package queue_moveitem

import "encoding/xml"

type QueueMoveItem struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QUQMUQ_03_1_1A Queue_MoveItem"`

	PlacementOption *PlacementOption `xml:"placementOption"`

	TargetDetails []*TargetDetails `xml:"targetDetails"`  // maxOccurs="10"

	MessageText *MessageText `xml:"messageText,omitempty"`  // minOccurs="0"

	RecordLocator *RecordLocator `xml:"recordLocator,omitempty"`  // minOccurs="0"

	NumberOfPNRs *NumberOfPNRs `xml:"numberOfPNRs,omitempty"`  // minOccurs="0"
}

type PlacementOption struct {

	SelectionDetails *SelectionDetails `xml:"selectionDetails"`
}

type SelectionDetails struct {

	Option string `xml:"option"`
}

type TargetDetails struct {

	TargetOffice *TargetOffice `xml:"targetOffice"`

	QueueNumber *QueueNumber `xml:"queueNumber,omitempty"`  // minOccurs="0"

	CategoryDetails *CategoryDetails `xml:"categoryDetails,omitempty"`  // minOccurs="0"

	PlacementDate *PlacementDate `xml:"placementDate,omitempty"`  // minOccurs="0"
}

type TargetOffice struct {

	SourceType *SourceType `xml:"sourceType"`

	OriginatorDetails *OriginatorDetails `xml:"originatorDetails,omitempty"`  // minOccurs="0"
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

type CategoryDetails struct {

	SubQueueInfoDetails *SubQueueInfoDetails `xml:"subQueueInfoDetails"`
}

type SubQueueInfoDetails struct {

	IdentificationType string `xml:"identificationType"`

	ItemNumber string `xml:"itemNumber,omitempty"`  // minOccurs="0"

	ItemDescription string `xml:"itemDescription,omitempty"`  // minOccurs="0"
}

type PlacementDate struct {

	TimeMode string `xml:"timeMode,omitempty"`  // minOccurs="0"

	DateTime *DateTime `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type DateTime struct {

	Year string `xml:"year"`

	Month float64 `xml:"month"`

	Day float64 `xml:"day"`

	Hour *float64 `xml:"hour,omitempty"`  // minOccurs="0"
}

type MessageText struct {

	FreeTextDetails *FreeTextDetails `xml:"freeTextDetails,omitempty"`  // minOccurs="0"

	FreeText []string `xml:"freeText"`  // maxOccurs="99"
}

type FreeTextDetails struct {

	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	Source string `xml:"source,omitempty"`  // minOccurs="0"

	Encoding string `xml:"encoding,omitempty"`  // minOccurs="0"
}

type RecordLocator struct {

	Reservation *Reservation `xml:"reservation"`
}

type Reservation struct {

	ControlNumber string `xml:"controlNumber"`
}

type NumberOfPNRs struct {

	QuantityDetails *QuantityDetails `xml:"quantityDetails"`
}

type QuantityDetails struct {

	NumberOfUnit float64 `xml:"numberOfUnit"`
}
