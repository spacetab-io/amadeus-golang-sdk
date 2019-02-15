package SalesReports_DisplayTransactionReport_v13_2 // tsrtrq132

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/sdk/formats"
)

type SalesReportsDisplayTransactionReport struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TSRTRQ_13_2_1A SalesReports_DisplayTransactionReport"`

	// Conveys an optional document number
	DocumentNumber *TicketNumberTypeI `xml:"documentNumber,omitempty"` // minOccurs="0"

	// Conveys an optional sequence/item number to target this given transaction.
	SequenceNumber *ItemNumberTypeI `xml:"sequenceNumber,omitempty"` // minOccurs="0"

	// Conveys optional agency reference.
	AgencyDetails *AdditionalBusinessSourceInformationTypeI `xml:"agencyDetails,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type AdditionalBusinessSourceInformationTypeI struct {
	// Source type
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	// Details (office/ID, IATA number) of the target office
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails,omitempty"` // minOccurs="0"
}

type ItemNumberIdentificationTypeI struct {
	// Conveys a document sequence number within the sales report.
	// xmlType: AlphaNumericString_Length1To6
	Number string `xml:"number,omitempty"` // minOccurs="0"

	// Type of the number.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"
}

type ItemNumberTypeI struct {
	// Conveys the sequence number of a transaction within a sales report display.
	ItemNumberDetails *ItemNumberIdentificationTypeI `xml:"itemNumberDetails"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// IATA number of the agency.
	OriginatorId *formats.NumericInteger_Length1To9 `xml:"originatorId,omitempty"` // minOccurs="0"

	// Office ID of the agency.
	// xmlType: AlphaNumericString_Length1To9
	InHouseIdentification1 string `xml:"inHouseIdentification1,omitempty"` // minOccurs="0"
}

type SourceTypeDetailsTypeI struct {
	// Source qualifier: reporting office.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier1 string `xml:"sourceQualifier1"`

	// Indicates that data are associated to all agencies sharing the same IATA number.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier2 string `xml:"sourceQualifier2,omitempty"` // minOccurs="0"
}

type TicketNumberDetailsTypeI struct {
	// Document number.
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"` // minOccurs="0"
}

type TicketNumberTypeI struct {
	// Details on the document transaction.
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`
}
