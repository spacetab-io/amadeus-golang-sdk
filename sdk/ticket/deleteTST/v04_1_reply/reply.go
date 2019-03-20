package Ticket_DeleteTSTReply_v04_1 // ttstdr041

//import "encoding/xml"

type TicketDeleteTSTReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTDR_04_1_1A Ticket_DeleteTSTReply"`

	// Process status after the TST delete
	ProcessStatus *ResponseAnalysisDetailsTypeI `xml:"processStatus,omitempty"` // minOccurs="0"

	ApplicationError *ApplicationError `xml:"applicationError,omitempty"` // minOccurs="0"

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"` // minOccurs="0"
}

type ApplicationError struct {
	// General error information returned by ticketing application
	ApplicationErrorInfo *ApplicationErrorInformationType `xml:"applicationErrorInfo"`

	// Description in free flow text of the error returned by ticketing application
	ErrorText *InteractiveFreeTextTypeI `xml:"errorText,omitempty"` // minOccurs="0"
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

type ReservationControlInformationDetailsTypeI struct {
	// Record locator.
	ControlNumber string `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {
	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}

type ResponseAnalysisDetailsTypeI struct {
	// Type of the response (update/cancel request)
	ResponseType string `xml:"responseType"`

	// Status of the process ran.
	ProcessingStatus string `xml:"processingStatus"`
}
