package pnr_cancel

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type PNRCancel struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/PNRXCL_11_3_1A PNR_Cancel"`

	// reservation control informations
	ReservationInfo *ReservationControlInformationType `xml:"reservationInfo,omitempty"`

	// specific actions to be processed on PNR
	PnrActions *OptionalPNRActionsType `xml:"pnrActions,omitempty"`

	// CANCEL PNR ELEMENT
	CancelElements *CancelPNRElementType `xml:"cancelElements,omitempty"`
}

type CancelPNRElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRXCL_11_3_1A CancelPNRElementType"`

	// E : XE    I : XI D : XD G : Name integration P : Priority line
	EntryType formats.AlphaNumericString_Length1To1 `xml:"entryType,omitempty"`

	// Element data
	Element *ElementIdentificationType `xml:"element,omitempty"`
}

type ElementIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRXCL_11_3_1A ElementIdentificationType"`

	// ELEMENT DATA
	Identifier formats.AlphaNumericString_Length1To3 `xml:"identifier,omitempty"`

	// element number
	Number formats.NumericInteger_Length1To5 `xml:"number,omitempty"`

	// sub element number
	SubElement formats.NumericInteger_Length1To5 `xml:"subElement,omitempty"`
}

type OptionalPNRActionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRXCL_11_3_1A OptionalPNRActionsType"`

	// Processing options. Only the option 0, 10, 11 and 20 are supported.
	OptionCode formats.NumericInteger_Length1To3 `xml:"optionCode,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRXCL_11_3_1A ReservationControlInformationDetailsTypeI"`

	// profile or PNR record locator
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRXCL_11_3_1A ReservationControlInformationType"`

	// record information
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}
