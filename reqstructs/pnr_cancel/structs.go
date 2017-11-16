package pnr_cancel

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type PNRCancel struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/PNRXCL_11_3_1A PNR_Cancel"`

	// reservation control informations
	ReservationInfo *ReservationControlInformationType `xml:"reservationInfo,omitempty"`  // minOccurs="0"

	// specific actions to be processed on PNR
	PnrActions *OptionalPNRActionsType `xml:"pnrActions"`

	// CANCEL PNR ELEMENT
	CancelElements []*CancelPNRElementType `xml:"cancelElements,omitempty"`  // minOccurs="0" maxOccurs="4"
}

//
// Complex structs
//

type CancelPNRElementType struct {

	// E : XE    I : XI D : XD G : Name integration P : Priority line
	EntryType formats.AlphaNumericString_Length1To1 `xml:"entryType"`

	// Element data
	Element []*ElementIdentificationType `xml:"element,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type ElementIdentificationType struct {

	// ELEMENT DATA
	Identifier formats.AlphaNumericString_Length1To3 `xml:"identifier,omitempty"`  // minOccurs="0"

	// element number
	Number *formats.NumericInteger_Length1To5 `xml:"number,omitempty"`  // minOccurs="0"

	// sub element number
	SubElement *formats.NumericInteger_Length1To5 `xml:"subElement,omitempty"`  // minOccurs="0"
}

type OptionalPNRActionsType struct {

	// Processing options. Only the option 0, 10, 11 and 20 are supported.
	OptionCode []formats.NumericInteger_Length1To3 `xml:"optionCode"`  // maxOccurs="40"
}

type ReservationControlInformationDetailsTypeI struct {

	// profile or PNR record locator
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber"`
}

type ReservationControlInformationType struct {

	// record information
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation"`
}
