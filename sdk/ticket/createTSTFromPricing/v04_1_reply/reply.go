package Ticket_CreateTSTFromPricingReply_v04_1 // tautcr041

//import "encoding/xml"

type TicketCreateTSTFromPricingReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A Ticket_CreateTSTFromPricingReply"`

	ApplicationError *ApplicationError `xml:"applicationError,omitempty"` // minOccurs="0"

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"` // minOccurs="0"

	TstList []*TstList `xml:"tstList,omitempty"` // minOccurs="0" maxOccurs="1980"
}

type ApplicationError struct {
	// General error information returned by ticketing application
	ApplicationErrorInfo *ApplicationErrorInformationType `xml:"applicationErrorInfo"`

	// Description in free flow text of the error returned by ticketing application
	ErrorText *InteractiveFreeTextTypeI `xml:"errorText,omitempty"` // minOccurs="0"
}

type TstList struct {
	// TST tattoo number created by the transaction.
	TstReference *ItemReferencesAndVersionsType `xml:"tstReference"`

	// Reference information on passengers.
	PaxInformation *ReferenceInformationTypeI `xml:"paxInformation,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type ApplicationErrorDetailType struct {
	// Code identifying the data validation error condition.
	ApplicationErrorCode string `xml:"applicationErrorCode"`

	// Identification of a code list.
	CodeListQualifier string `xml:"codeListQualifier,omitempty"` // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	CodeListResponsibleAgency string `xml:"codeListResponsibleAgency,omitempty"` // minOccurs="0"
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ApplicationErrorDetail *ApplicationErrorDetailType `xml:"applicationErrorDetail"`
}

type InteractiveFreeTextTypeI struct {
	// Free flow text describing the error
	ErrorFreeText string `xml:"errorFreeText,omitempty"` // minOccurs="0"
}

type ItemReferencesAndVersionsType struct {
	// qualifies the type of the reference used. Code set to define
	ReferenceType string `xml:"referenceType,omitempty"` // minOccurs="0"

	// Tattoo number (It is in fact the Tst Display Number)
	UniqueReference *int32 `xml:"uniqueReference,omitempty"` // minOccurs="0"

	// Gives the TST ID number
	IDDescription *UniqueIdDescriptionType `xml:"iDDescription,omitempty"` // minOccurs="0"
}

type ReferenceInformationTypeI struct {
	// Passenger/segment/TST reference details
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type ReferencingDetailsTypeI struct {
	// Qualifyer of the reference (Pax/Seg/Tst)
	RefQualifier string `xml:"refQualifier,omitempty"` // minOccurs="0"

	// Passenger/segment/TST reference number
	RefNumber *int32 `xml:"refNumber,omitempty"` // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI struct {
	// Record locator.
	ControlNumber string `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {
	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}

type UniqueIdDescriptionType struct {
	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber int32 `xml:"iDSequenceNumber"`
}
