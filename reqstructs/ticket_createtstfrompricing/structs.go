package ticket_createtstfrompricing

import (
	"encoding/xml"

	"github.com/kidem/amadeus-ws-go/formats"
)

type TicketCreateTSTFromPricing struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCQ_04_1_1A Ticket_CreateTSTFromPricing"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`

	PsaList *PsaList `xml:"psaList,omitempty"`
}

type PsaList struct {

	// Reference of the fare selected. A fare may have been calculated by Fare Quote for several passengers but there is still the possibility to create a TST only for a part of these passengers.
	ItemReference *ItemReferencesAndVersionsType `xml:"itemReference,omitempty"`

	// Reference information on passengers.
	PaxReference *ReferenceInformationTypeI `xml:"paxReference,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCQ_04_1_1A ItemReferencesAndVersionsType"`

	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// Tattoo number : It is in fact the TST display number.
	UniqueReference formats.NumericInteger_Length1To5 `xml:"uniqueReference,omitempty"`
}

type ReferenceInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCQ_04_1_1A ReferenceInformationTypeI"`

	// Passenger/segment/TST/fare reference details
	RefDetails *ReferencingDetailsTypeI `xml:"refDetails,omitempty"`
}

type ReferencingDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCQ_04_1_1A ReferencingDetailsTypeI"`

	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCQ_04_1_1A ReservationControlInformationDetailsTypeI"`

	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCQ_04_1_1A ReservationControlInformationTypeI"`

	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation,omitempty"`
}
