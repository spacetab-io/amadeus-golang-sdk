package Ticket_DisplayTST_v07_1 // ttstrq071

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/sdk/formats"
)

type TicketDisplayTST struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A Ticket_DisplayTST"`

	// It can be: - 'ALL' - all TSTs are displayed; - 'SEL' - only the information corresponding to the TST/passenger selection is displayed.
	DisplayMode *CodedAttributeType `xml:"displayMode"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"` // minOccurs="0"

	// Scrolling information.
	ScrollingInformation *ActionDetailsTypeI `xml:"scrollingInformation,omitempty"` // minOccurs="0"

	// Used to get TST information for selected TST references. As we can have 10 TST per Pax, 99 passenger per PNR, and a TST split with the Infant , the max number of TST is 1980.
	TstReference []*ItemReferencesAndVersionsType `xml:"tstReference,omitempty"` // minOccurs="0" maxOccurs="1980"

	// Passenger/segment references information about TST(s) to retrieve. - Passenger reference specified : all the TSTs concerning this passenger reference are returned. - Passenger/segment reference : only the TST having this passenger/segment association is returned.
	PsaInformation *ReferenceInformationTypeI `xml:"psaInformation,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type ActionDetailsTypeI struct {
	// Information on next list of TSTs to return.
	NextListInformation *ReferenceTypeI `xml:"nextListInformation,omitempty"` // minOccurs="0"
}

type CodedAttributeInformationType struct {
	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType"`
}

type CodedAttributeType struct {
	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails"`
}

type ItemReferencesAndVersionsType struct {
	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"` // minOccurs="0"

	// Tattoo number (It is in fact the Tst Display Number)
	UniqueReference *formats.NumericInteger_Length1To5 `xml:"uniqueReference,omitempty"` // minOccurs="0"

	// Gives the TST ID number
	IDDescription *UniqueIdDescriptionType `xml:"iDDescription,omitempty"` // minOccurs="0"
}

type ReferenceInformationTypeI struct {
	// Passenger/segment/TST reference details
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type ReferenceTypeI struct {
	// In case of query specifies the number of TSTs to get in reply. In case of response specifies the number of TSTs remaining.
	RemainingInformation *formats.NumericInteger_Length1To5 `xml:"remainingInformation,omitempty"` // minOccurs="0"

	// In case of first query specifies the value of  this field in the last reply. In case of other queries specifies the last reference returned in the previous list. In case of reply specifies the last TST reference of the list. In case of last reply the value of this field set in the first query is sent.
	RemainingReference formats.AlphaNumericString_Length1To5 `xml:"remainingReference,omitempty"` // minOccurs="0"
}

type ReferencingDetailsTypeI struct {
	// Qualifyer of the reference (Pax/Seg/Tst)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"` // minOccurs="0"

	// Passenger/segment/TST reference number
	RefNumber *formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"` // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI struct {
	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {
	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}

type UniqueIdDescriptionType struct {
	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber formats.NumericInteger_Length1To11 `xml:"iDSequenceNumber"`
}
