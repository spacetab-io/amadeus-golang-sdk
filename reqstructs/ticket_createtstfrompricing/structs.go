package ticket_createtstfrompricing

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/formats"
)

type TicketCreateTSTFromPricing struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCQ_04_1_1A Ticket_CreateTSTFromPricing"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`  // minOccurs="0"

	PsaList []*PsaList `xml:"psaList"`  // maxOccurs="1980"
}

type PsaList struct {

	// Reference of the fare selected. A fare may have been calculated by Fare Quote for several passengers but there is still the possibility to create a TST only for a part of these passengers.
	ItemReference *ItemReferencesAndVersionsType `xml:"itemReference"`

	// Reference information on passengers.
	PaxReference *ReferenceInformationTypeI `xml:"paxReference,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type ItemReferencesAndVersionsType struct {

	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Tattoo number : It is in fact the TST display number.
	UniqueReference *formats.NumericInteger_Length1To5 `xml:"uniqueReference,omitempty"`  // minOccurs="0"
}

type ReferenceInformationTypeI struct {

	// Passenger/segment/TST/fare reference details
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferencingDetailsTypeI struct {

	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber *formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI struct {

	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {

	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}
