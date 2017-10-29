package ticket_display_tst

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type TicketDisplayTST struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A Ticket_DisplayTST"`

	// It can be: - 'ALL' - all TSTs are displayed; - 'SEL' - only the information corresponding to the TST/passenger selection is displayed.
	DisplayMode *CodedAttributeType `xml:"displayMode,omitempty"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`

	// Scrolling information.
	ScrollingInformation *ActionDetailsTypeI `xml:"scrollingInformation,omitempty"`

	// Used to get TST information for selected TST references. As we can have 10 TST per Pax, 99 passenger per PNR, and a TST split with the Infant , the max number of TST is 1980.
	TstReference *ItemReferencesAndVersionsType `xml:"tstReference,omitempty"`

	// Passenger/segment references information about TST(s) to retrieve. - Passenger reference specified : all the TSTs concerning this passenger reference are returned. - Passenger/segment reference : only the TST having this passenger/segment association is returned.
	PsaInformation *ReferenceInformationTypeI `xml:"psaInformation,omitempty"`
}

type ActionDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A ActionDetailsTypeI"`

	// Information on next list of TSTs to return.
	NextListInformation *ReferenceTypeI `xml:"nextListInformation,omitempty"`
}

type CodedAttributeInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A CodedAttributeInformationType"`

	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType,omitempty"`
}

type CodedAttributeType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A CodedAttributeType"`

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A ItemReferencesAndVersionsType"`

	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// Tattoo number (It is in fact the Tst Display Number)
	UniqueReference formats.NumericInteger_Length1To5 `xml:"uniqueReference,omitempty"`

	// Gives the TST ID number
	IDDescription *UniqueIdDescriptionType `xml:"iDDescription,omitempty"`
}

type ReferenceInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A ReferenceInformationTypeI"`

	// Passenger/segment/TST reference details
	RefDetails *ReferencingDetailsTypeI `xml:"refDetails,omitempty"`
}

type ReferenceTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A ReferenceTypeI"`

	// In case of query specifies the number of TSTs to get in reply. In case of response specifies the number of TSTs remaining.
	RemainingInformation formats.NumericInteger_Length1To5 `xml:"remainingInformation,omitempty"`

	// In case of first query specifies the value of  this field in the last reply. In case of other queries specifies the last reference returned in the previous list. In case of reply specifies the last TST reference of the list. In case of last reply the value of this field set in the first query is sent.
	RemainingReference formats.AlphaNumericString_Length1To5 `xml:"remainingReference,omitempty"`
}

type ReferencingDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A ReferencingDetailsTypeI"`

	// Qualifyer of the reference (Pax/Seg/Tst)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Passenger/segment/TST reference number
	RefNumber formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A ReservationControlInformationDetailsTypeI"`

	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A ReservationControlInformationTypeI"`

	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation,omitempty"`
}

type UniqueIdDescriptionType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A UniqueIdDescriptionType"`

	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber formats.NumericInteger_Length1To11 `xml:"iDSequenceNumber,omitempty"`
}
