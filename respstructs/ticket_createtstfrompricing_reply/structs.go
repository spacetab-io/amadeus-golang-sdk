package ticket_createtstfrompricing_reply

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type TicketCreateTSTFromPricingReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A Ticket_CreateTSTFromPricingReply"`

	ApplicationError *ApplicationError `xml:"applicationError,omitempty"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`

	TstList *TstList `xml:"tstList,omitempty"`
}

type TstList struct {

	// TST tattoo number created by the transaction.
	TstReference *ItemReferencesAndVersionsType `xml:"tstReference,omitempty"`

	// Reference information on passengers.
	PaxInformation *ReferenceInformationTypeI `xml:"paxInformation,omitempty"`
}

type ApplicationError struct {

	// General error information returned by ticketing application
	ApplicationErrorInfo *ApplicationErrorInformationType `xml:"applicationErrorInfo,omitempty"`

	// Description in free flow text of the error returned by ticketing application
	ErrorText *InteractiveFreeTextTypeI `xml:"errorText,omitempty"`
}

type ApplicationErrorDetailType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A ApplicationErrorDetailType"`

	// Code identifying the data validation error condition.
	ApplicationErrorCode formats.AlphaNumericString_Length1To5 `xml:"applicationErrorCode,omitempty"`

	// Identification of a code list.
	CodeListQualifier formats.AlphaNumericString_Length1To3 `xml:"codeListQualifier,omitempty"`

	// Code identifying the agency responsible for a code list.
	CodeListResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"codeListResponsibleAgency,omitempty"`
}

type ApplicationErrorInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A ApplicationErrorInformationType"`

	// Application error details.
	ApplicationErrorDetail *ApplicationErrorDetailType `xml:"applicationErrorDetail,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A InteractiveFreeTextTypeI"`

	// Free flow text describing the error
	ErrorFreeText formats.AlphaNumericString_Length1To70 `xml:"errorFreeText,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A ItemReferencesAndVersionsType"`

	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// Tattoo number (It is in fact the Tst Display Number)
	UniqueReference formats.NumericInteger_Length1To5 `xml:"uniqueReference,omitempty"`

	// Gives the TST ID number
	IDDescription *UniqueIdDescriptionType `xml:"iDDescription,omitempty"`
}

type ReferenceInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A ReferenceInformationTypeI"`

	// Passenger/segment/TST reference details
	RefDetails *ReferencingDetailsTypeI `xml:"refDetails,omitempty"`
}

type ReferencingDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A ReferencingDetailsTypeI"`

	// Qualifyer of the reference (Pax/Seg/Tst)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Passenger/segment/TST reference number
	RefNumber formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A ReservationControlInformationDetailsTypeI"`

	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A ReservationControlInformationTypeI"`

	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation,omitempty"`
}

type UniqueIdDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A UniqueIdDescriptionType"`

	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber formats.NumericInteger_Length1To11 `xml:"iDSequenceNumber,omitempty"`
}
