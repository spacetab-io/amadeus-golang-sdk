package pnr_list

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type PNRList struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TNLRES_00_1_1A PNR_List"`

	// provide free from or coded text information
	FreeFormText *InteractiveFreeTextTypeI `xml:"freeFormText,omitempty"`  // minOccurs="0"

	Citypair *Citypair `xml:"citypair"`

	// to specify the error or information
	ErrorInformation *ApplicationErrorInformationTypeI `xml:"errorInformation,omitempty"`  // minOccurs="0"
}

type Citypair struct {

	// origin and destination
	OriginDestinationMarker *OriginAndDestinationDetailsTypeI `xml:"originDestinationMarker"`

	TravellerInformationSection []*TravellerInformationSection `xml:"travellerInformationSection"`  // maxOccurs="300"
}

type TravellerInformationSection struct {

	// To specify a traveler and personal details relating to the traveler.
	TravellerInformation *TravellerInformationTypeI `xml:"travellerInformation"`

	// To indicate quantity and action required in relation to a product
	RelatedProduct *RelatedProductInformationTypeI `xml:"relatedProduct,omitempty"`  // minOccurs="0"

	// To specify details related to a product
	TravelProduct *TravelProductInformationTypeI `xml:"travelProduct,omitempty"`  // minOccurs="0"

	// to specify a reference to a reservation
	ReservationInfo *ReservationControlInformationTypeI `xml:"reservationInfo"`

	// to specify details related to availability status or cabin configuration for a product
	ProductInfo *ProductInformationTypeI `xml:"productInfo,omitempty"`  // minOccurs="0"

	// to specify the message type and business function
	MessageAction *MessageActionDetailsTypeI `xml:"messageAction,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type ApplicationErrorDetailTypeI struct {

	// Application error, coded - 366 for name list too long
	ErrorCode formats.AlphaNumericString_Length1To3 `xml:"errorCode"`

	// Code list qualifier - INF for information code , WEC for warning code
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier"`

	// Code list responsible agency, coded
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency"`
}

type ApplicationErrorInformationTypeI struct {

	// APPLICATION ERROR DETAIL
	ErrorDetail *ApplicationErrorDetailTypeI `xml:"errorDetail"`
}

type CompanyIdentificationTypeI struct {

	// Company identification , airline code
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification"`
}

type FreeTextQualificationTypeI struct {

	// Identifies whether the free text is coded or not coded  3 for Literal text
	SubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"subjectQualifier"`

	// Coded text, or specifies type of information in the free text
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`  // minOccurs="0"
}

type InteractiveFreeTextTypeI struct {

	// FREE TEXT QUALIFICATION
	FreetextDetail *FreeTextQualificationTypeI `xml:"freetextDetail,omitempty"`  // minOccurs="0"

	// free text
	Text formats.AlphaNumericString_Length1To70 `xml:"text,omitempty"`  // minOccurs="0"
}

type LocationTypeI struct {

	// Board point  For non air segment, e.g. SUR, can be alphanumeric.  5 chars long Board and Off points for SNCF TRN segment
	CityCode formats.AlphaNumericString_Length1To5 `xml:"cityCode"`
}

type MessageActionDetailsTypeI struct {

	// MESSAGE FUNCTION OR BUSINESS DETAILS
	Business *MessageFunctionBusinessDetailsTypeI `xml:"business"`
}

type MessageFunctionBusinessDetailsTypeI struct {

	// Business function, coded
	Function formats.AlphaNumericString_Length1To3 `xml:"function,omitempty"`  // minOccurs="0"
}

type OriginAndDestinationDetailsTypeI struct {
}

type ProductDateTimeTypeI struct {

	// departure date of the flight
	DepDate formats.Date_DDMMYY `xml:"depDate"`
}

type ProductDetailsTypeI struct {

	// Characteritic identification - class of service
	Identification formats.AlphaString_Length1To1 `xml:"identification"`

	// Item description identification - N for Night class
	Description formats.AlphaNumericString_Length1To1 `xml:"description,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetailsTypeI struct {

	// Product Idenfication -Flight number or  OPEN (ARNK is not a possible value since no SI in similar name list)
	Identification formats.AlphaNumericString_Length1To4 `xml:"identification"`

	// Flight number alpha suffix  A, B, C, D, E
	Subtype formats.AlphaString_Length1To1 `xml:"subtype,omitempty"`  // minOccurs="0"
}

type ProductInformationTypeI struct {

	// PRODUCT DETAILS
	Product *ProductDetailsTypeI `xml:"product,omitempty"`  // minOccurs="0"
}

type RelatedProductInformationTypeI struct {

	// Number in party
	Quantity *formats.NumericInteger_Length1To3 `xml:"quantity,omitempty"`  // minOccurs="0"

	// Status, coded , cancel indicator , XX-cancel Airimp code
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI struct {

	// Company identification - 1A for amadeus
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId"`

	// Reservation control number - amadeus record locator of the requested PNR
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {

	// RESERVATION CONTROL INFORMATION
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`  // minOccurs="0"
}

type TravelProductInformationTypeI struct {

	// PRODUCT DATE OR TIME
	Product *ProductDateTimeTypeI `xml:"product,omitempty"`  // minOccurs="0"

	// LOCATION
	BoardpointDetail *LocationTypeI `xml:"boardpointDetail,omitempty"`  // minOccurs="0"

	// LOCATION
	OffpointDetail *LocationTypeI `xml:"offpointDetail,omitempty"`  // minOccurs="0"

	// COMPANY IDENTIFICATION
	CompanyDetail *CompanyIdentificationTypeI `xml:"companyDetail,omitempty"`  // minOccurs="0"

	// PRODUCT IDENTIFICATION DETAILS
	ProductDetails *ProductIdentificationDetailsTypeI `xml:"productDetails,omitempty"`  // minOccurs="0"
}

type TravellerDetailsTypeI struct {

	// Traveller First Name
	FirstName formats.AlphaNumericString_Length1To56 `xml:"firstName"`
}

type TravellerInformationTypeI struct {

	// TRAVELLER SURNAME INFORMATION
	Traveller *TravellerSurnameInformationTypeI `xml:"traveller"`

	// TRAVELLER DETAILS
	Passenger []*TravellerDetailsTypeI `xml:"passenger,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type TravellerSurnameInformationTypeI struct {

	// Traveller last name
	Surname formats.AlphaNumericString_Length1To57 `xml:"surname"`
}
