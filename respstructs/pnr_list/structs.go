package pnr_list

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type PNRList struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A PNR_List"`

	// provide free from or coded text information
	FreeFormText *InteractiveFreeTextTypeI `xml:"freeFormText,omitempty"`

	Citypair *Citypair `xml:"citypair,omitempty"`

	// to specify the error or information
	ErrorInformation *ApplicationErrorInformationTypeI `xml:"errorInformation,omitempty"`
}

type ApplicationErrorDetailTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A ApplicationErrorDetailTypeI"`

	// Application error, coded - 366 for name list too long
	ErrorCode formats.AlphaNumericString_Length1To3 `xml:"errorCode,omitempty"`

	// Code list qualifier - INF for information code , WEC for warning code
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Code list responsible agency, coded
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`
}

type ApplicationErrorInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A ApplicationErrorInformationTypeI"`

	// APPLICATION ERROR DETAIL
	ErrorDetail *ApplicationErrorDetailTypeI `xml:"errorDetail,omitempty"`
}

type Citypair struct {
	// origin and destination
	OriginDestinationMarker *OriginAndDestinationDetailsTypeI `xml:"originDestinationMarker,omitempty"`

	TravellerInformationSection *TravellerInformationSection `xml:"travellerInformationSection,omitempty"`
}

type TravellerInformationSection struct {
	// To specify a traveler and personal details relating to the traveler.
	TravellerInformation *TravellerInformationTypeI `xml:"travellerInformation,omitempty"`

	// To indicate quantity and action required in relation to a product
	RelatedProduct *RelatedProductInformationTypeI `xml:"relatedProduct,omitempty"`

	// To specify details related to a product
	TravelProduct *TravelProductInformationTypeI `xml:"travelProduct,omitempty"`

	// to specify a reference to a reservation
	ReservationInfo *ReservationControlInformationTypeI `xml:"reservationInfo,omitempty"`

	// to specify details related to availability status or cabin configuration for a product
	ProductInfo *ProductInformationTypeI `xml:"productInfo,omitempty"`

	// to specify the message type and business function
	MessageAction *MessageActionDetailsTypeI `xml:"messageAction,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A CompanyIdentificationTypeI"`

	// Company identification , airline code
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification,omitempty"`
}

type FreeTextQualificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A FreeTextQualificationTypeI"`

	// Identifies whether the free text is coded or not coded  3 for Literal text
	SubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"subjectQualifier,omitempty"`

	// Coded text, or specifies type of information in the free text
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A InteractiveFreeTextTypeI"`

	// FREE TEXT QUALIFICATION
	FreetextDetail *FreeTextQualificationTypeI `xml:"freetextDetail,omitempty"`

	// free text
	Text formats.AlphaNumericString_Length1To70 `xml:"text,omitempty"`
}

type LocationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A LocationTypeI"`

	// Board point  For non air segment, e.g. SUR, can be alphanumeric.  5 chars long Board and Off points for SNCF TRN segment
	CityCode formats.AlphaNumericString_Length1To5 `xml:"cityCode,omitempty"`
}

type MessageActionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A MessageActionDetailsTypeI"`

	// MESSAGE FUNCTION OR BUSINESS DETAILS
	Business *MessageFunctionBusinessDetailsTypeI `xml:"business,omitempty"`
}

type MessageFunctionBusinessDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A MessageFunctionBusinessDetailsTypeI"`

	// Business function, coded
	Function formats.AlphaNumericString_Length1To3 `xml:"function,omitempty"`
}

type OriginAndDestinationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A OriginAndDestinationDetailsTypeI"`
}

type ProductDateTimeTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A ProductDateTimeTypeI"`

	// departure date of the flight
	DepDate formats.Date_DDMMYY `xml:"depDate,omitempty"`
}

type ProductDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A ProductDetailsTypeI"`

	// Characteritic identification - class of service
	Identification formats.AlphaString_Length1To1 `xml:"identification,omitempty"`

	// Item description identification - N for Night class
	Description formats.AlphaNumericString_Length1To1 `xml:"description,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A ProductIdentificationDetailsTypeI"`

	// Product Idenfication -Flight number or  OPEN (ARNK is not a possible value since no SI in similar name list)
	Identification formats.AlphaNumericString_Length1To4 `xml:"identification,omitempty"`

	// Flight number alpha suffix  A, B, C, D, E
	Subtype formats.AlphaString_Length1To1 `xml:"subtype,omitempty"`
}

type ProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A ProductInformationTypeI"`

	// PRODUCT DETAILS
	Product *ProductDetailsTypeI `xml:"product,omitempty"`
}

type RelatedProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A RelatedProductInformationTypeI"`

	// Number in party
	Quantity formats.NumericInteger_Length1To3 `xml:"quantity,omitempty"`

	// Status, coded , cancel indicator , XX-cancel Airimp code
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A ReservationControlInformationDetailsTypeI"`

	// Company identification - 1A for amadeus
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Reservation control number - amadeus record locator of the requested PNR
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A ReservationControlInformationTypeI"`

	// RESERVATION CONTROL INFORMATION
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}

type TravelProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A TravelProductInformationTypeI"`

	// PRODUCT DATE OR TIME
	Product *ProductDateTimeTypeI `xml:"product,omitempty"`

	// LOCATION
	BoardpointDetail *LocationTypeI `xml:"boardpointDetail,omitempty"`

	// LOCATION
	OffpointDetail *LocationTypeI `xml:"offpointDetail,omitempty"`

	// COMPANY IDENTIFICATION
	CompanyDetail *CompanyIdentificationTypeI `xml:"companyDetail,omitempty"`

	// PRODUCT IDENTIFICATION DETAILS
	ProductDetails *ProductIdentificationDetailsTypeI `xml:"productDetails,omitempty"`
}

type TravellerDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A TravellerDetailsTypeI"`

	// Traveller First Name
	FirstName formats.AlphaNumericString_Length1To56 `xml:"firstName,omitempty"`
}

type TravellerInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A TravellerInformationTypeI"`

	// TRAVELLER SURNAME INFORMATION
	Traveller *TravellerSurnameInformationTypeI `xml:"traveller,omitempty"`

	// TRAVELLER DETAILS
	Passenger *TravellerDetailsTypeI `xml:"passenger,omitempty"`
}

type TravellerSurnameInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A TravellerSurnameInformationTypeI"`

	// Traveller last name
	Surname formats.AlphaNumericString_Length1To57 `xml:"surname,omitempty"`
}
