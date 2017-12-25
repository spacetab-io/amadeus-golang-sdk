package pnr_addmultielements

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/formats"
)

type PNRAddMultiElements struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A PNR_AddMultiElements"`

	// To specify a reference to a reservation
	ReservationInfo *ReservationControlInformationTypeI `xml:"reservationInfo,omitempty"`  // minOccurs="0"

	// To specify specific Actions to be processed on PNR
	PnrActions *OptionalPNRActionsType `xml:"pnrActions"`

	TravellerInfo []*TravellerInfo `xml:"travellerInfo,omitempty"`  // minOccurs="0" maxOccurs="100"

	OriginDestinationDetails []*OriginDestinationDetails `xml:"originDestinationDetails,omitempty"`  // minOccurs="0" maxOccurs="50"

	DataElementsMaster *DataElementsMaster `xml:"dataElementsMaster,omitempty"`  // minOccurs="0"
}

type TravellerInfo struct {

	// To specify the PNR segments/elements references and action to apply
	ElementManagementPassenger *ElementManagementSegmentType `xml:"elementManagementPassenger"`

	PassengerData []*PassengerData `xml:"passengerData"`  // maxOccurs="2"
}

type PassengerData struct {

	// To specify a traveler(s) and personal details relating to the traveler(s).  In values with 'X00' (X being any letter), 00 should be replaced by any value between 01 and 99.
	TravellerInformation *TravellerInformationTypeI `xml:"travellerInformation"`

	// Passenger date of birth (DDMMYYYY) If the passenger has an infant, not in a separate TIF, then the date is used for the infant date of birth.
	DateOfBirth *DateAndTimeInformationType `xml:"dateOfBirth,omitempty"`  // minOccurs="0"
}

type OriginDestinationDetails struct {

	// To convey the origin and destination of a journey
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination"`

	ItineraryInfo []*ItineraryInfo `xml:"itineraryInfo,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ItineraryInfo struct {

	// To specify the PNR segments/elements references and action to apply
	ElementManagementItinerary *ElementManagementSegmentType `xml:"elementManagementItinerary"`

	AirAuxItinerary *AirAuxItinerary `xml:"airAuxItinerary,omitempty"`  // minOccurs="0"

	// To provide specific reference identification
	ReferenceForSegment *ReferenceInfoType `xml:"referenceForSegment,omitempty"`  // minOccurs="0"
}

type AirAuxItinerary struct {

	// To specify details related to a product
	TravelProduct *TravelProductInformationType `xml:"travelProduct"`

	// To specify the message type and business function
	MessageAction *MessageActionDetailsTypeI `xml:"messageAction"`

	// To indicate quantity and action required in relation to a product
	RelatedProduct *RelatedProductInformationTypeI `xml:"relatedProduct,omitempty"`  // minOccurs="0"

	// To specify the details for making a selection
	SelectionDetailsAir *SelectionDetailsTypeI `xml:"selectionDetailsAir,omitempty"`  // minOccurs="0"

	// To specify a reference to a reservation
	ReservationInfoSell *ReservationControlInformationTypeI `xml:"reservationInfoSell,omitempty"`  // minOccurs="0"

	// To provide free form or coded long text information.
	FreetextItinerary *LongFreeTextType `xml:"freetextItinerary,omitempty"`  // minOccurs="0"
}

type DataElementsMaster struct {

	// marker
	Marker1 *DummySegmentTypeI `xml:"marker1"`

	DataElementsIndiv []*DataElementsIndiv `xml:"dataElementsIndiv,omitempty"`  // minOccurs="0" maxOccurs="250"
}

type DataElementsIndiv struct {

	// To specify the PNR segments/elements references and action to apply
	ElementManagementData *ElementManagementSegmentType `xml:"elementManagementData"`

	// To specify the Amadeus PNR Individual Security element Note: up to 3 ISI segments may be needed to represent the ES element and its 3 types of receiver
	PnrSecurity *IndividualPnrSecurityInformationType `xml:"pnrSecurity,omitempty"`  // minOccurs="0"

	// specify amadeus accounting informations
	Acccounting *AccountingInformationElementType `xml:"acccounting,omitempty"`  // minOccurs="0"

	// To specify different kinds of remarks
	MiscellaneousRemark *MiscellaneousRemarksType `xml:"miscellaneousRemark,omitempty"`  // minOccurs="0"

	// To specify special requests or services information relating to a traveller.
	ServiceRequest *SpecialRequirementsDetailsTypeI `xml:"serviceRequest,omitempty"`  // minOccurs="0"

	// To provide date and time details relative to flight movements
	DateAndTimeInformation *DateAndTimeInformationTypeI `xml:"dateAndTimeInformation,omitempty"`  // minOccurs="0"

	// To specify the Amadeus PNR Tour Code element
	TourCode *TourCodeType `xml:"tourCode,omitempty"`  // minOccurs="0"

	// To specify an Amadeus PNR Ticket element
	TicketElement *TicketElementType `xml:"ticketElement,omitempty"`  // minOccurs="0"

	// To provide free form or coded long text information.
	FreetextData *LongFreeTextType `xml:"freetextData,omitempty"`  // minOccurs="0"

	// To specify the way data are mapped for structured
	StructuredAddress *StructuredAddressType `xml:"structuredAddress,omitempty"`  // minOccurs="0"

	// To specify an Amadeus PNR Option element
	OptionElement *OptionElementType `xml:"optionElement,omitempty"`  // minOccurs="0"

	// To request a Hard Copy Print
	Printer *PrinterIdentificationType `xml:"printer,omitempty"`  // minOccurs="0"

	// The group handles Seat Request with possibly rail preferences
	SeatGroup *SeatEntityType `xml:"seatGroup,omitempty"`  // minOccurs="0"

	// To specify the Amadeus PNR Fare elements
	FareElement *FareElementsType `xml:"fareElement,omitempty"`  // minOccurs="0"

	// To specify the Amadeus PNR Fare discount element
	FareDiscount *FareDiscountElementType `xml:"fareDiscount,omitempty"`  // minOccurs="0"

	// To specify the Amadeus PNR Manual Document Registration element
	ManualFareDocument *ManualDocumentRegistrationType `xml:"manualFareDocument,omitempty"`  // minOccurs="0"

	// To specify the Amadeus PNR Commission element
	Commission *CommissionElementType `xml:"commission,omitempty"`  // minOccurs="0"

	// To specify the Amadeus PNR Original Issue / Issue in Exchange For element
	OriginalIssue *OriginalIssueType `xml:"originalIssue,omitempty"`  // minOccurs="0"

	// To convey details describing the form of payment
	FormOfPayment *FormOfPaymentTypeI `xml:"formOfPayment,omitempty"`  // minOccurs="0"

	// To convey additional details of the form of payment
	FopExtension []*MarketSpecificDataType `xml:"fopExtension,omitempty"`  // minOccurs="0" maxOccurs="3"

	// To convey the FOP service details
	ServiceDetails []*StatusTypeI `xml:"serviceDetails,omitempty"`  // minOccurs="0" maxOccurs="3"

	// To specify frequent traveler verification
	FrequentTravellerVerification *FrequentTravellerVerificationType `xml:"frequentTravellerVerification,omitempty"`  // minOccurs="0"

	// To specify the Amadeus PNR Ticketing Carrier Designator element
	TicketingCarrier *TicketingCarrierDesignatorType `xml:"ticketingCarrier,omitempty"`  // minOccurs="0"

	// To specify the Amadeus PNR Fare Print Override element
	FarePrintOverride *FarePrintOverrideType `xml:"farePrintOverride,omitempty"`  // minOccurs="0"

	// To convey frequent traveler program information relating to a specific traveller or group.
	FrequentTravellerData *FrequentTravellerInformationTypeU `xml:"frequentTravellerData,omitempty"`  // minOccurs="0"

	// To specify access level of an entity (office) to the element.
	AccessLevel *ExtendedOwnershipSecurityDetailsType `xml:"accessLevel,omitempty"`  // minOccurs="0"

	// To provide specific reference identification
	ReferenceForDataElement *ReferenceInfoType `xml:"referenceForDataElement,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type AccountingElementType struct {

	// Account number
	Number formats.AlphaNumericString_Length1To10 `xml:"number,omitempty"`  // minOccurs="0"

	// Cost Number
	CostNumber formats.AlphaNumericString_Length1To60 `xml:"costNumber,omitempty"`  // minOccurs="0"

	// IATA company number
	CompanyNumber formats.AlphaNumericString_Length1To12 `xml:"companyNumber,omitempty"`  // minOccurs="0"

	// Client Reference Number
	ClientReference formats.AlphaNumericString_Length1To60 `xml:"clientReference,omitempty"`  // minOccurs="0"

	GSTTaxDetails formats.AlphaNumericString_Length1To109 `xml:"gSTTaxDetails,omitempty"`  // minOccurs="0"
}

type AccountingInformationElementType struct {

	// One of these 4 data elements is mandatory , but non in particular
	Account *AccountingElementType `xml:"account,omitempty"`  // minOccurs="0"

	// Number of units qualifier
	AccountNumberOfUnits formats.AlphaNumericString_Length1To3 `xml:"accountNumberOfUnits,omitempty"`  // minOccurs="0"
}

type ClassDetailsType struct {

	// For the booking class code.
	Code formats.AlphaNumericString_Length1To1 `xml:"code,omitempty"`  // minOccurs="0"

	BookingClass formats.AlphaNumericString_Length2To2 `xml:"bookingClass,omitempty"`  // minOccurs="0"
}

type CommissionElementType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// Commission indicator  M, C, P, CR, PR
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`  // minOccurs="0"

	// Commission
	CommissionInfo *CommissionInformationType `xml:"commissionInfo"`

	// Old commission
	OldCommission *CommissionInformationType_6428C `xml:"oldCommission,omitempty"`  // minOccurs="0"

	// Manual capping Amount (after tag C)
	ManualCapping *formats.NumericDecimal_Length1To10 `xml:"manualCapping,omitempty"`  // minOccurs="0"
}

type CommissionInformationType struct {

	// Percentage (max 2 decimals)
	Percentage *formats.NumericDecimal_Length1To5 `xml:"percentage,omitempty"`  // minOccurs="0"

	// Amount (before tag A)
	Amount *formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`  // minOccurs="0"

	// VAT indicator  V
	VatIndicator formats.AlphaNumericString_Length1To2 `xml:"vatIndicator,omitempty"`  // minOccurs="0"

	// Net remit indicator  N
	RemitIndicator formats.AlphaNumericString_Length1To2 `xml:"remitIndicator,omitempty"`  // minOccurs="0"
}

type CommissionInformationType_6428C struct {

	// Percentage (max 2 decimals)
	Percentage *formats.NumericInteger_Length1To5 `xml:"percentage,omitempty"`  // minOccurs="0"

	// Amount (before tag A)
	Amount *formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`  // minOccurs="0"

	// VAT indicator  V
	VatIndicator formats.AlphaNumericString_Length1To2 `xml:"vatIndicator,omitempty"`  // minOccurs="0"

	// Net remit indicator  N
	RemitIndicator formats.AlphaNumericString_Length1To2 `xml:"remitIndicator,omitempty"`  // minOccurs="0"
}

type CompanyIdentificationType struct {

	// To specify the frequent traveller program company code
	Code formats.AlphaNumericString_Length1To3 `xml:"code"`

	// To specify the frequent traveller program partnership company  code
	PartnerCode formats.AlphaNumericString_Length1To3 `xml:"partnerCode,omitempty"`  // minOccurs="0"

	// To specify the frequent traveller program other partnership company  code
	OtherPartnerCode formats.AlphaNumericString_Length1To3 `xml:"otherPartnerCode,omitempty"`  // minOccurs="0"
}

type CompanyIdentificationTypeI struct {

	// 1. Air segment: Airline code 2. ATX segment: Airline code of the airline to take action 3. CAR segment: Airline code of the airline to take action 4. HTL segment: Airline code of the airline to take action
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification"`

	// To convey a second carrier (e.g. in case of multi airline open segments)
	SecondIdentification formats.AlphaNumericString_Length1To3 `xml:"secondIdentification,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsTypeI struct {

	// Miscellaneous Charge Order element : date.- SEAT ssr : date of cahnge of gauge - gourp seat SSR : date of change of gauge
	FirstDate formats.Date_DDMMYY `xml:"firstDate"`

	// Miscellaneous Charge Order element : ARC carrier code code.
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType"`

	// Miscellaneous Charge Order element : ARC city code.
	LocationIdentification formats.AlphaNumericString_Length1To3 `xml:"locationIdentification"`
}

type DateAndTimeDetailsTypeI_56946C struct {

	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Inf/Child date of birth
	Date formats.AlphaNumericString_Length1To8 `xml:"date,omitempty"`  // minOccurs="0"
}

type DateAndTimeInformationType struct {

	// DATE AND TIME DETAILS
	DateAndTimeDetails *DateAndTimeDetailsTypeI_56946C `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0"
}

type DateAndTimeInformationTypeI struct {

	// date and time details
	DateAndTime *DateAndTimeDetailsTypeI `xml:"dateAndTime"`
}

type DiscountInformationType struct {

	// Discount code
	AdjustmentReason formats.AlphaNumericString_Length1To6 `xml:"adjustmentReason,omitempty"`  // minOccurs="0"

	// Discount percentage
	Percentage *formats.NumericInteger_Length2To2 `xml:"percentage,omitempty"`  // minOccurs="0"

	// Status code
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`  // minOccurs="0"

	// Staff employee number
	StaffNumber formats.AlphaNumericString_Length1To7 `xml:"staffNumber,omitempty"`  // minOccurs="0"

	// Staff employee name
	StaffName formats.AlphaString_Length1To30 `xml:"staffName,omitempty"`  // minOccurs="0"
}

type DummySegmentTypeI struct {
}

type ElementManagementSegmentType struct {

	// segments/elements references - type and number
	Reference *ReferencingDetailsType `xml:"reference,omitempty"`  // minOccurs="0"

	// PNR segment or element name
	SegmentName formats.AlphaNumericString_Length1To3 `xml:"segmentName"`
}

type ExtendedOwnershipSecurityDetailsType struct {

	// This composite contains security data between entities or related to one entity.
	SecurityDetails []*ExtendedSecurityDetailsType `xml:"securityDetails"`  // maxOccurs="5"
}

type ExtendedSecurityDetailsType struct {

	// Used to specify which kind of entities is specified. F - for family
	TypeOfEntity formats.AlphaNumericString_Length1To3 `xml:"typeOfEntity,omitempty"`  // minOccurs="0"

	// Used to specify the access mode regarding agreement and entities
	AccessMode formats.AlphaNumericString_Length1To1 `xml:"accessMode"`

	// Used to specify entity on which the detailed security applied. Mask is specified useing *, for instance, corporate 1A0 is specified as follows : ***1A0***.
	InhouseIdentification formats.AlphaNumericString_Length1To9 `xml:"inhouseIdentification"`
}

type FareDiscountElementType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// To specify the discount details. Only 1 repetition must be used. The Fare Discount element cannot process multiple discounts. If you wish to enter multiple discounts for a passenger, you should enter several FD elements and associate them to the same passenger.
	Discount []*DiscountInformationType `xml:"discount,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Date of birth
	BirthDate formats.AlphaNumericString_Length6To6 `xml:"birthDate,omitempty"`  // minOccurs="0"

	// number of years or months
	NumberDetail *NumberOfUnitDetailsType `xml:"numberDetail,omitempty"`  // minOccurs="0"

	// company ID and reference number (should not be used)
	RpInformation *RpInformationType `xml:"rpInformation,omitempty"`  // minOccurs="0"

	// Customer info
	Customer *ItemDetailsType `xml:"customer,omitempty"`  // minOccurs="0"

	// Resident Discount applicable for discount code RM,RC,DC,BP or BI
	ResidentDisc *ItemDetailsType_64822C `xml:"residentDisc,omitempty"`  // minOccurs="0"
}

type FareElementsType struct {

	// E for FE - Endorsements / Restrictions element  K for FK - Shadow AIR office ID element  S for FS - Miscellaneous Ticketing Information element  Z for FZ - Miscellaneous Information element
	GeneralIndicator formats.AlphaNumericString_Length1To1 `xml:"generalIndicator,omitempty"`  // minOccurs="0"

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// 1. FK element: Office identification
	OfficeId []formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`  // minOccurs="0" maxOccurs="5"

	// 1. FE element: Free text 2. FS element: Free text 3. FZ element: Free text
	FreetextLong formats.AlphaNumericString_Length1To126 `xml:"freetextLong,omitempty"`  // minOccurs="0"
}

type FarePrintOverrideDetailsType struct {

	// Base fare override
	BaseFare formats.AlphaNumericString_Length1To11 `xml:"baseFare"`

	// Total fare override
	TotalFare formats.AlphaNumericString_Length1To11 `xml:"totalFare,omitempty"`  // minOccurs="0"

	// equivalent fare override
	EquivalentFare formats.AlphaNumericString_Length1To11 `xml:"equivalentFare,omitempty"`  // minOccurs="0"

	// Tax amount override
	TaxAmount []formats.AlphaNumericString_Length1To11 `xml:"taxAmount,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type FarePrintOverrideType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// fare override detaild
	Override *FarePrintOverrideDetailsType `xml:"override,omitempty"`  // minOccurs="0"
}

type FormOfPaymentDetailsTypeI struct {

	// Form(s) of payment   CA for Cash (Amadeus code CASH)  CK for Check (Amadeus code CHECK)  CC for Credit card (Amadeus code CC)  MS for Miscellaneous (Amadeus code MS)
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification"`

	// Amount
	Amount *formats.NumericDecimal_Length1To9 `xml:"amount,omitempty"`  // minOccurs="0"

	// Credit card code
	CreditCardCode formats.AlphaNumericString_Length1To3 `xml:"creditCardCode,omitempty"`  // minOccurs="0"

	// Credit card account number
	AccountNumber formats.AlphaNumericString_Length1To35 `xml:"accountNumber,omitempty"`  // minOccurs="0"

	// Credit card expiry date
	ExpiryDate formats.Date_MMYY `xml:"expiryDate,omitempty"`  // minOccurs="0"

	// Credit card approval code
	ApprovalCode formats.AlphaNumericString_Length1To8 `xml:"approvalCode,omitempty"`  // minOccurs="0"

	// Customer account number
	CustomerAccountNumber formats.AlphaNumericString_Length1To10 `xml:"customerAccountNumber,omitempty"`  // minOccurs="0"

	// Extended credit payment request  E for extended payment
	PaymentTimeReference formats.AlphaNumericString_Length1To3 `xml:"paymentTimeReference,omitempty"`  // minOccurs="0"

	// free text
	Freetext formats.AlphaNumericString_Length1To70 `xml:"freetext,omitempty"`  // minOccurs="0"

	// Currency (as there is no place elsewhere)
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`  // minOccurs="0"
}

type FormOfPaymentTypeI struct {

	// form of payment details
	Fop []*FormOfPaymentDetailsTypeI `xml:"fop"`  // maxOccurs="3"
}

type FormatedTourCodeType struct {

	// Tour type  IT, BT
	ProductId formats.AlphaNumericString_Length1To2 `xml:"productId,omitempty"`  // minOccurs="0"

	// Last digit of year in which tour code became effective
	Year *formats.NumericInteger_Length1To1 `xml:"year,omitempty"`  // minOccurs="0"

	// Airline code of sponsor
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`  // minOccurs="0"

	// Conference area approval code
	ApprovalCode formats.AlphaNumericString_Length1To1 `xml:"approvalCode,omitempty"`  // minOccurs="0"

	// Tour identification
	PartyId formats.AlphaNumericString_Length1To8 `xml:"partyId,omitempty"`  // minOccurs="0"
}

type FreeFormatTourCodeType struct {

	// Free format indicator  FF
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`  // minOccurs="0"

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To14 `xml:"freetext,omitempty"`  // minOccurs="0"
}

type FreeTextQualificationType struct {

	// Identifies whether the free text is coded or not coded  3 for Literal text
	SubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"subjectQualifier"`

	// 1. AP element: 2. AQ element: 3. OS element
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`  // minOccurs="0"

	// Transmittable/non-transmittable indicator (S or X). Codeset list not applicable.
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`  // minOccurs="0"

	// Airline or system code
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`  // minOccurs="0"
}

type FrequentTravellerIdentificationTypeU struct {

	// Airline Code
	CompanyId formats.AlphaNumericString_Length2To3 `xml:"companyId"`

	// Frequent Traveller Number
	MembershipNumber formats.AlphaNumericString_Length1To25 `xml:"membershipNumber"`
}

type FrequentTravellerInformationTypeU struct {

	// frequent traveller identification
	FrequentTraveller *FrequentTravellerIdentificationTypeU `xml:"frequentTraveller"`
}

type FrequentTravellerVerificationType struct {

	// A  Add FT name  N  Name verification  R   Redeem miles
	ActionRequest formats.AlphaNumericString_Length1To3 `xml:"actionRequest"`

	// company identification
	Company *CompanyIdentificationType `xml:"company,omitempty"`  // minOccurs="0"

	// frequent flyer number
	Account *ProductAccountDetailsType `xml:"account"`
}

type IndividualPnrSecurityInformationType struct {

	// Returned before EOT or after retrieve by RTE
	Security []*IndividualSecurityType `xml:"security,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Returned when retrieved
	SecurityInfo *SecurityInformationType `xml:"securityInfo,omitempty"`  // minOccurs="0"

	// Code as in the display:  G for Amadeus Global Core Office Identification  I for IATA number  P for Pseudo-Office Identification  Default is G.
	Indicator formats.AlphaNumericString_Length1To1 `xml:"indicator,omitempty"`  // minOccurs="0"
}

type IndividualSecurityType struct {

	// Type of receiver G:  Type of receiver I:  Type of receiver P:
	Identification formats.AlphaNumericString_Length1To14 `xml:"identification"`

	// R for Read  B for Both read and write  N for None
	AccessMode formats.AlphaString_Length1To1 `xml:"accessMode"`
}

type IssueInformationType struct {

	// Numeric airline code
	AirlineCode formats.AlphaNumericString_Length3To3 `xml:"airlineCode"`

	// {Original issue , exchange for} document number
	DocumentNumber formats.AlphaNumericString_Length1To10 `xml:"documentNumber"`

	// {Original issue , exchange for} document check digit
	DocumentCD *formats.NumericInteger_Length1To1 `xml:"documentCD,omitempty"`  // minOccurs="0"

	// 1st coupon number
	Coupon1 formats.AlphaNumericString_Length1To1 `xml:"coupon1,omitempty"`  // minOccurs="0"

	// 2nd coupon number
	Coupon2 formats.AlphaNumericString_Length1To1 `xml:"coupon2,omitempty"`  // minOccurs="0"

	// 3rd coupon number
	Coupon3 formats.AlphaNumericString_Length1To1 `xml:"coupon3,omitempty"`  // minOccurs="0"

	// 4th coupon number
	Coupon4 formats.AlphaNumericString_Length1To1 `xml:"coupon4,omitempty"`  // minOccurs="0"

	// Last 2 digits of the last conjunction document of the {original issue , exchange for} document
	LastConjunction *formats.NumericInteger_Length2To2 `xml:"lastConjunction,omitempty"`  // minOccurs="0"

	// Last conjunction document of the {original issue , exchange for} document check digit
	ExchangeDocumentCD *formats.NumericInteger_Length1To1 `xml:"exchangeDocumentCD,omitempty"`  // minOccurs="0"

	// 1st last conjunction document of the {original issue , exchange for} document coupon number
	LastConjunction1 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction1,omitempty"`  // minOccurs="0"

	// 2nd last conjunction document of the {original issue , exchange for} document coupon number
	LastConjunction2 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction2,omitempty"`  // minOccurs="0"

	// 3rd last conjunction document of the {original issue , exchange for} document 4th coupon number
	LastConjunction3 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction3,omitempty"`  // minOccurs="0"

	// 4th last conjunction document of the {original issue , exchange for} document 1st coupon number
	LastConjunction4 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction4,omitempty"`  // minOccurs="0"

	// City code of the issue
	CityCode formats.AlphaString_Length3To3 `xml:"cityCode"`

	// Date of the {original , new} issue
	DateOfIssue formats.Date_DDMMYY `xml:"dateOfIssue"`

	// IATA number
	IataNumber formats.AlphaNumericString_Length1To9 `xml:"iataNumber,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// Amount
	Amount *formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`  // minOccurs="0"
}

type ItemDetailsType struct {

	// Only applicable for some discount codes
	CompanyId formats.AlphaString_Length2To2 `xml:"companyId,omitempty"`  // minOccurs="0"

	// Only applicable for some discount codes
	CardType *formats.NumericInteger_Length1To1 `xml:"cardType,omitempty"`  // minOccurs="0"

	// Only applicable for some discount codes
	CardNumber *formats.NumericInteger_Length1To13 `xml:"cardNumber,omitempty"`  // minOccurs="0"

	// for PC - alpha value for other - numeric value
	CardCheck formats.AlphaNumericString_Length1To1 `xml:"cardCheck,omitempty"`  // minOccurs="0"

	// for PC only
	Owner *formats.NumericInteger_Length2To2 `xml:"owner,omitempty"`  // minOccurs="0"

	// for PC only
	Version *formats.NumericInteger_Length1To1 `xml:"version,omitempty"`  // minOccurs="0"
}

type ItemDetailsType_64822C struct {

	// DN, TR,GR,AM,CR,MR
	IdCardCode formats.AlphaString_Length2To2 `xml:"idCardCode,omitempty"`  // minOccurs="0"

	// T - used for TR resident discount only
	IdCardType formats.AlphaString_Length1To1 `xml:"idCardType,omitempty"`  // minOccurs="0"

	// Card Number
	CardNumber *formats.NumericInteger_Length1To13 `xml:"cardNumber,omitempty"`  // minOccurs="0"

	// Card alpha check
	AlphaCheck formats.AlphaNumericString_Length1To1 `xml:"alphaCheck,omitempty"`  // minOccurs="0"

	// Zip code
	ZipCode formats.AlphaNumericString_Length1To9 `xml:"zipCode,omitempty"`  // minOccurs="0"

	// Credential certificate number
	CertificateNumber formats.AlphaNumericString_Length1To20 `xml:"certificateNumber,omitempty"`  // minOccurs="0"
}

type LocationTypeI struct {

	// 1. Air segment: Boarding point 2. ATX segment: Boarding point 3. CAR segment: Pick-up point city 4. HTL segment: Check-in city 5. MIS segment: City code
	CityCode formats.AlphaNumericString_Length1To5 `xml:"cityCode"`

	// 1. TRN SNCF segment: Off point city name
	CityName formats.AlphaNumericString_Length1To17 `xml:"cityName,omitempty"`  // minOccurs="0"
}

type LongFreeTextType struct {

	// free text details
	FreetextDetail *FreeTextQualificationType `xml:"freetextDetail,omitempty"`  // minOccurs="0"

	// Long free text information
	LongFreetext formats.AlphaNumericString_Length1To199 `xml:"longFreetext,omitempty"`  // minOccurs="0"
}

type ManualDocumentRegistrationType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// documentation details
	Document *ManualDocumentType `xml:"document,omitempty"`  // minOccurs="0"

	// Free text
	Freeflow formats.AlphaNumericString_Length1To49 `xml:"freeflow,omitempty"`  // minOccurs="0"
}

type ManualDocumentType struct {

	// Numeric airline code
	CompanyId formats.NumericString_Length3To3 `xml:"companyId"`

	// Ticket number
	TicketNumber formats.NumericString_Length10To10 `xml:"ticketNumber"`

	// Ticket number check digit
	TicketNumberCd *formats.NumericInteger_Length1To1 `xml:"ticketNumberCd,omitempty"`  // minOccurs="0"

	// Last conjunction ticket number
	LastConjuction *formats.NumericInteger_Length2To2 `xml:"lastConjuction,omitempty"`  // minOccurs="0"

	// Last conjunction ticket number check digit
	LastConjuctionCD *formats.NumericInteger_Length1To1 `xml:"lastConjuctionCD,omitempty"`  // minOccurs="0"
}

type MarketSpecificDataDetailsType struct {

	// Credit Card Verification data (*CV data)
	CvData formats.AlphaNumericString_Length1To10 `xml:"cvData,omitempty"`  // minOccurs="0"

	// Printed and reported freeflow
	PrintedFreeflow formats.AlphaNumericString_Length1To70 `xml:"printedFreeflow,omitempty"`  // minOccurs="0"

	// Reported but not printed freeflow
	ReportedFreeflow formats.AlphaNumericString_Length1To70 `xml:"reportedFreeflow,omitempty"`  // minOccurs="0"

	// Credit Card ONO data.
	OnoData formats.AlphaNumericString_Length1To15 `xml:"onoData,omitempty"`  // minOccurs="0"

	// Credit Card GWT data
	GwtData formats.AlphaNumericString_Length1To15 `xml:"gwtData,omitempty"`  // minOccurs="0"

	// Credit Card Holder name.
	ChdData formats.AlphaNumericString_Length1To40 `xml:"chdData,omitempty"`  // minOccurs="0"

	// Delegation code.
	DelegationCode formats.AlphaNumericString_Length1To3 `xml:"delegationCode,omitempty"`  // minOccurs="0"

	// MCO Fop Document Number
	McoDocNumber formats.AlphaNumericString_Length1To13 `xml:"mcoDocNumber,omitempty"`  // minOccurs="0"

	// MCO Fop Coupon Number
	McoCouponNumber formats.AlphaNumericString_Length1To4 `xml:"mcoCouponNumber,omitempty"`  // minOccurs="0"

	// MCO Fop Iata Number
	McoIataNumber *formats.NumericInteger_Length1To9 `xml:"mcoIataNumber,omitempty"`  // minOccurs="0"

	// MCO Fop Place of Issue
	McoPlaceOfIssue formats.AlphaNumericString_Length1To3 `xml:"mcoPlaceOfIssue,omitempty"`  // minOccurs="0"

	// MCO Fop date of Issue. DDMMMYY
	McoDateOfIssue formats.AlphaNumericString_Length7To7 `xml:"mcoDateOfIssue,omitempty"`  // minOccurs="0"

	// Standard Fop Iata Number
	IataNumber *formats.NumericInteger_Length1To9 `xml:"iataNumber,omitempty"`  // minOccurs="0"
}

type MarketSpecificDataType struct {

	// FOP detail Sequence Number.
	FopSequenceNumber formats.NumericInteger_Length1To2 `xml:"fopSequenceNumber"`

	// Passenger type: PAX or INF.
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// Form of Payment Market Specific Data
	NewFopsDetails *MarketSpecificDataDetailsType `xml:"newFopsDetails,omitempty"`  // minOccurs="0"

	// To provide extended FOP details.
	ExtFOP []*ReferencingDetailsTypeI `xml:"extFOP,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type MessageActionDetailsTypeI struct {

	// type of segment
	Business *MessageFunctionBusinessDetailsTypeI `xml:"business"`
}

type MessageFunctionBusinessDetailsTypeI struct {

	// 1. Air segment:
	Function formats.AlphaNumericString_Length1To3 `xml:"function"`
}

type MiscellaneousRemarkType struct {

	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// This is the 3rd character (x) of the remark title RIx or RMx, or 2 letter code for RMxx, conditional for RM, not applicable for RC and RQ
	Category formats.AlphaNumericString_Length1To2 `xml:"category,omitempty"`  // minOccurs="0"

	// Free text and message sequence numbers of the remarks.
	Freetext formats.AlphaNumericString_Length1To127 `xml:"freetext,omitempty"`  // minOccurs="0"

	// Provider type (element RIA):  1 for Air provider   2 for Car provider (CCR)  3 for Hotel Provider (HHL)  M for Miscellaneous
	ProviderType formats.AlphaNumericString_Length1To3 `xml:"providerType,omitempty"`  // minOccurs="0"

	// MCO element : Currency code
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// MCO element : total fee amount
	Amount *formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`  // minOccurs="0"

	// Office Id (confidential remark RC)
	OfficeId []formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type MiscellaneousRemarksType struct {

	// miscellaneous remarks
	Remarks *MiscellaneousRemarkType `xml:"remarks,omitempty"`  // minOccurs="0"
}

type NetRemitTourCodeType struct {

	// Net remit indicator  N
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`  // minOccurs="0"

	// Free flow text of next remit
	Freetext formats.AlphaNumericString_Length1To14 `xml:"freetext,omitempty"`  // minOccurs="0"
}

type NumberOfUnitDetailsType struct {

	// Number of  years or months
	Units formats.NumericInteger_Length1To2 `xml:"units"`

	// YRS for year  MTH for month
	UnitsQualifier formats.AlphaNumericString_Length1To3 `xml:"unitsQualifier"`
}

type OptionElementInformationType struct {

	// Office Id
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`  // minOccurs="0"

	// Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`  // minOccurs="0"

	// Queue number
	Queue *formats.NumericInteger_Length1To3 `xml:"queue,omitempty"`  // minOccurs="0"

	// Category number
	Category *formats.NumericInteger_Length1To3 `xml:"category,omitempty"`  // minOccurs="0"

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To200 `xml:"freetext,omitempty"`  // minOccurs="0"
}

type OptionElementType struct {

	// queuing option details
	OptionDetail *OptionElementInformationType `xml:"optionDetail,omitempty"`  // minOccurs="0"
}

type OptionalPNRActionsType struct {

	// 0       No Special Processing 10 - 49  PNR processing options 50 - 99  PNRACC options 100 - 149  Other Opt 150 - 199  Error Processing 200 - 229  Car Opt 230 - 259  Hotel Opt 260 - 299  Air + AUX Opt 300 - 329  Ticketing Opt
	OptionCode []formats.NumericInteger_Length1To3 `xml:"optionCode"`  // maxOccurs="40"
}

type OriginAndDestinationDetailsTypeI struct {

	// Airport/city code of  Origin In a Client request message, a non-blank ODI is used in an air sell request to advise that the following segments (TVL etc...) are connected. There is a maximum of 6 TVLs following a non-blank ODI.
	Origin formats.AlphaString_Length3To3 `xml:"origin,omitempty"`  // minOccurs="0"

	// Airport/city code of  Destination
	Destination formats.AlphaString_Length3To3 `xml:"destination,omitempty"`  // minOccurs="0"
}

type OriginalIssueType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// 8 for Voucher/Certificate indicator   RN for RN indicator
	VoucherIndicator formats.AlphaNumericString_Length1To2 `xml:"voucherIndicator,omitempty"`  // minOccurs="0"

	// 1st occurrence for original issue 2 occurrences for issues in exchange
	Issue []*IssueInformationType `xml:"issue"`  // maxOccurs="3"

	// Base fare amount
	BaseFare *formats.NumericDecimal_Length1To10 `xml:"baseFare,omitempty"`  // minOccurs="0"

	// Total tax amount
	TotalTax *formats.NumericDecimal_Length1To10 `xml:"totalTax,omitempty"`  // minOccurs="0"

	// Penalty amount
	Penalty *formats.NumericDecimal_Length1To10 `xml:"penalty,omitempty"`  // minOccurs="0"

	// Free flow text
	Freeflow formats.AlphaNumericString_Length1To126 `xml:"freeflow,omitempty"`  // minOccurs="0"
}

type PrinterIdentificationDetailsType struct {

	// name of the printer
	Name formats.AlphaNumericString_Length5To6 `xml:"name"`

	// netwrok ID of the printer
	Network formats.AlphaNumericString_Length2To2 `xml:"network,omitempty"`  // minOccurs="0"
}

type PrinterIdentificationType struct {

	// identification details
	IdentifierDetail *PrinterIdentificationDetailsType `xml:"identifierDetail,omitempty"`  // minOccurs="0"

	// 1A office id
	Office formats.AlphaNumericString_Length9To9 `xml:"office,omitempty"`  // minOccurs="0"

	// IATA teletype address
	TeletypeAddress formats.AlphaNumericString_Length7To7 `xml:"teletypeAddress,omitempty"`  // minOccurs="0"
}

type ProductAccountDetailsType struct {

	// To specify the product/account number qualifier.
	NumberQualifier formats.AlphaNumericString_Length1To3 `xml:"numberQualifier,omitempty"`  // minOccurs="0"

	// A code to identify a frequent traveller (e.g. a frequent traveller number)
	Number formats.AlphaNumericString_Length1To27 `xml:"number"`
}

type ProductDateTimeTypeI struct {

	// 1. Air segment: Departure date 2. ATX segment: Requested date 3. CAR segment: Pick-up date 4. HTL segment: Check-in date 5. MIS segment: Date for service required
	DepDate formats.Date_DDMMYY `xml:"depDate"`

	// 1. Air segment Departure time 2. SUR segment: Pick-up time 3. TRN Amtrack segment: Departure time 4. TRN SNCF segment: Departure time
	DepTime formats.Time24_HHMM `xml:"depTime,omitempty"`  // minOccurs="0"

	// 1. Air segment  Arrival date (not in the display) 2. CAR segment Drop-off date 3. HTL segment: Check-out date 4. TTO segment: Return date of the Tour
	ArrDate formats.Date_DDMMYY `xml:"arrDate,omitempty"`  // minOccurs="0"

	// 1. Air segment  Arrival time 2. TRN Amtrack segment: Arrival time 3. TRN SNCF segment: Arrival time
	ArrTime formats.Time24_HHMM `xml:"arrTime,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetailsTypeI struct {

	// 1. Air segment: Flight number or  OPEN  ARNK - air segment arrival unknown 2. CAR se 3. segment: Car type 4. SUR segment: Transportation type  (refer to VGTVD transaction)
	Identification formats.AlphaNumericString_Length1To6 `xml:"identification"`

	// 1. Air segment: Class of service 2. TRN Amtrack segment: Class of service 3. TRN SNCF segment: Class of service
	ClassOfService formats.AlphaString_Length1To1 `xml:"classOfService,omitempty"`  // minOccurs="0"

	// 1. Air segment  Flight number alpha suffix  A, B, C, D, E 2. SUR segment: Departure code  A or D 3. TRN SNCF segment: Train type
	Subtype formats.AlphaString_Length1To1 `xml:"subtype,omitempty"`  // minOccurs="0"

	// 1. Air segment:  N for Night class 2. TRN Amtrack segment:  N for Night class 3. TRN SNCF segment:  N for Night class
	Description formats.AlphaNumericString_Length1To1 `xml:"description,omitempty"`  // minOccurs="0"
}

type ProductTypeDetailsType struct {

	// Used to convey availibility context.
	FlightIndicator formats.AlphaNumericString_Length1To2 `xml:"flightIndicator"`
}

type RailSeatConfigurationType struct {

	// Seat space.
	SeatSpace formats.AlphaNumericString_Length2To2 `xml:"seatSpace,omitempty"`  // minOccurs="0"

	// Coach type.
	CoachType formats.AlphaNumericString_Length2To2 `xml:"coachType,omitempty"`  // minOccurs="0"

	// Seat equipment.
	SeatEquipment formats.AlphaNumericString_Length2To2 `xml:"seatEquipment,omitempty"`  // minOccurs="0"

	// Seat position.
	SeatPosition formats.AlphaNumericString_Length1To1 `xml:"seatPosition,omitempty"`  // minOccurs="0"

	// Seat direction.
	SeatDirection formats.AlphaNumericString_Length1To1 `xml:"seatDirection,omitempty"`  // minOccurs="0"

	// Seat deck.
	SeatDeck formats.AlphaNumericString_Length1To1 `xml:"seatDeck,omitempty"`  // minOccurs="0"

	// Special passenger information.
	SpecialPassengerType []formats.AlphaNumericString_Length1To1 `xml:"specialPassengerType,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type RailSeatPreferencesType struct {

	// Selection of the type of seat request.
	SeatRequestFunction formats.AlphaNumericString_Length1To1 `xml:"seatRequestFunction,omitempty"`  // minOccurs="0"

	// Seat smoking zone indicator.
	SmokingIndicator formats.AlphaString_Length1To1 `xml:"smokingIndicator,omitempty"`  // minOccurs="0"

	// Seat class details.
	ClassDetails *ClassDetailsType `xml:"classDetails,omitempty"`  // minOccurs="0"

	// Seat configuration details.
	SeatConfiguration *RailSeatConfigurationType `xml:"seatConfiguration,omitempty"`  // minOccurs="0"
}

type RailSeatReferenceInformationType struct {

	// Rail seat reference information.
	RailSeatReferenceDetails *SeatReferenceInformationType `xml:"railSeatReferenceDetails,omitempty"`  // minOccurs="0"
}

type ReferenceInfoType struct {

	// This composite is used to transmit association information
	Reference []*ReferencingDetailsType `xml:"reference,omitempty"`  // minOccurs="0" maxOccurs="198"
}

type ReferencingDetailsType struct {

	// Amadeus codes are used here.  PT for Passenger Tatoo reference number  PR for Passenger Client- request-message-defined reference number  ST for Segment Tatoo reference number  SR for Segment Client- request-message-defined reference number
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier"`

	// refers to a PNR segment/element that had this number in its related EMS segment in the same message (qualifier PT, ST)
	Number formats.AlphaNumericString_Length1To5 `xml:"number"`
}

type ReferencingDetailsTypeI struct {

	// Reference qualifier
	ReferenceQualifier formats.AlphaNumericString_Length1To10 `xml:"referenceQualifier"`

	// Reference value.
	DataValue formats.AlphaNumericString_Length1To35 `xml:"dataValue,omitempty"`  // minOccurs="0"
}

type RelatedProductInformationTypeI struct {

	// 1. Air segment: Number of passengers 2. ATX segment: Number of passengers 3. CAR segment: Number of cars 4. CCR segment: Number of cars 5. HHL segment: Number of  rooms 6. HTL segment: Number of  rooms
	Quantity *formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`  // minOccurs="0"

	// status
	Status formats.AlphaString_Length1To2 `xml:"status"`
}

type ReservationControlInformationDetailsTypeI struct {

	// 1A
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`  // minOccurs="0"

	// 1. RR element: Record locator of the original PNR  2. Record locator information: Airline record locator 3. Profile record locator information: Profile record locator 4. Air segment: Passive segment airline record locator  Due to technical limitations, RCI for air segment is truncated to 7 characters.
	ControlNumber formats.AlphaNumericString_Length1To19 `xml:"controlNumber,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationTypeI struct {

	// reservation control information - i.e. record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`  // minOccurs="0"
}

type RpInformationType struct {

	// Airline code (should not be used)
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`  // minOccurs="0"

	// Numeric value (should not be used)
	ReferenceNumber *formats.NumericInteger_Length1To12 `xml:"referenceNumber,omitempty"`  // minOccurs="0"
}

type SeatEntityType struct {

	// To make Seat requests on flights within the PNR
	SeatRequest *SeatRequestType `xml:"seatRequest"`

	// Used to convey specific seat details relative to Train for a specific request or the "near-to" seat details for a "next-to" request.
	RailSeatReferenceInformation []*RailSeatReferenceInformationType `xml:"railSeatReferenceInformation,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Rail Seat Preferences
	RailSeatPreferences *RailSeatPreferencesType `xml:"railSeatPreferences,omitempty"`  // minOccurs="0"
}

type SeatReferenceInformationType struct {

	// Coach number.
	CoachNumber formats.AlphaNumericString_Length1To3 `xml:"coachNumber,omitempty"`  // minOccurs="0"

	// Deck number.
	DeckNumber formats.AlphaNumericString_Length1To3 `xml:"deckNumber,omitempty"`  // minOccurs="0"

	// Seat number.
	SeatNumber formats.AlphaNumericString_Length1To3 `xml:"seatNumber,omitempty"`  // minOccurs="0"
}

type SeatRequestType struct {

	// seat requirement details
	Seat *SeatRequierementsType `xml:"seat,omitempty"`  // minOccurs="0"

	// Maximum 99 for Group Seat request
	Special []*SeatRequierementsDataType `xml:"special,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type SeatRequierementsDataType struct {

	// Seat number + row (seat request)  Number of seats (Group seat request)
	Data formats.AlphaNumericString_Length1To4 `xml:"data,omitempty"`  // minOccurs="0"

	// 3 occurrences may be used for in Amadeus seat request to indicate:  1. smoking/no smoking  2. 1st area preference  2nd area preference or passenger type
	SeatType []formats.AlphaNumericString_Length1To2 `xml:"seatType,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type SeatRequierementsType struct {

	// G for group
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Type of Seat requested. S - Smoking (SMST) N - No Smoking (NSST) RQST
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`  // minOccurs="0"

	// Board point
	Boardpoint formats.AlphaString_Length3To3 `xml:"boardpoint,omitempty"`  // minOccurs="0"

	// Off point
	Offpoint formats.AlphaString_Length3To3 `xml:"offpoint,omitempty"`  // minOccurs="0"
}

type SecurityInformationType struct {

	// Date of creation
	CreationDate formats.AlphaNumericString_Length6To6 `xml:"creationDate"`

	// Agent initials and duty code as in ORG (eg: AASU)
	AgentCode formats.AlphaNumericString_Length4To4 `xml:"agentCode"`

	// Office Id of creation/update
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`  // minOccurs="0"
}

type SelectionDetailsInformationTypeI struct {

	// See comment below
	Option formats.AlphaNumericString_Length1To3 `xml:"option"`
}

type SelectionDetailsTypeI struct {

	// level of sell to be processed
	Selection []*SelectionDetailsInformationTypeI `xml:"selection"`  // maxOccurs="10"
}

type SpecialRequirementsDataDetailsTypeI struct {

	// Seat number + row (seat SSR)  Number of seats (Group seat SSR)
	Data formats.AlphaNumericString_Length1To4 `xml:"data,omitempty"`  // minOccurs="0"

	// 3 occurrences may be used for in Amadeus seat SSR to indicate: 1. smoking/no smoking  2. 1st area preference 3. 2nd area preference or passenger type
	SeatType []formats.AlphaNumericString_Length1To2 `xml:"seatType,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type SpecialRequirementsDetailsTypeI struct {

	// special requirement type details
	Ssr *SpecialRequirementsTypeDetailsTypeI `xml:"ssr"`

	// Group seat SSR cannot ask for specific seats but only smoking and/or non-smoking (see Group seat SSR). Therefore the maximum repetitions here is 9 seats (1 per passenger of non-group PNR).
	Ssrb []*SpecialRequirementsDataDetailsTypeI `xml:"ssrb,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type SpecialRequirementsTypeDetailsTypeI struct {

	// - ATA/IATA defined Special Service Requirement code.  (refer to IATA AIRIMP documentation) - SK element : Keyword
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`  // minOccurs="0"

	// ATA/IATA status code Codeset list not applicable.
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`  // minOccurs="0"

	// Number of services requested
	Quantity *formats.NumericInteger_Length1To3 `xml:"quantity,omitempty"`  // minOccurs="0"

	// Airline code or YY
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`  // minOccurs="0"

	// Processing indicator, coded. - Normal SSR   P01 request for SSR explosion at EOT ...
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`  // minOccurs="0"

	// Board point
	Boardpoint formats.AlphaString_Length3To3 `xml:"boardpoint,omitempty"`  // minOccurs="0"

	// Off point
	Offpoint formats.AlphaString_Length3To3 `xml:"offpoint,omitempty"`  // minOccurs="0"

	// free text data
	Freetext []formats.AlphaNumericString_Length1To70 `xml:"freetext,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type StatusDetailsTypeI struct {

	// Service indicator SV : Service Fee
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`  // minOccurs="0"
}

type StatusTypeI struct {

	// To convey the status details
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails"`
}

type StructuredAddressInformationType struct {

	// A1 for Address line 1
	OptionA1 formats.AlphaNumericString_Length1To2 `xml:"optionA1"`

	// A1 50 char
	OptionTextA1 formats.AlphaNumericString_Length1To50 `xml:"optionTextA1"`
}

type StructuredAddressInformationType_5063C struct {

	// CY for company - NA for name - A2 for addr line - PO for P.O box-ZP for postacl code - CI for city - ST for state-CO for country
	Option formats.AlphaNumericString_Length1To2 `xml:"option"`

	// CY-NA-CI -  30char : A2 - 50 char: ST-CT- 25 char : PO 8 char - ZP 9 char.
	OptionText formats.AlphaNumericString_Length1To50 `xml:"optionText"`
}

type StructuredAddressType struct {

	// Information type, coded  2  for billing address  P08  for general mailing address  P19  for miscellaneous mailing address  P24  for home mailing address  P25  for delivery mailing address
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`  // minOccurs="0"

	// address line
	Address *StructuredAddressInformationType `xml:"address"`

	// upto 8 possible address options
	OptionalData []*StructuredAddressInformationType_5063C `xml:"optionalData,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type TicketElementType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// general ticketing information
	Ticket *TicketInformationType `xml:"ticket"`

	// Print options (//print options after double slash)
	PrintOptions formats.AlphaNumericString_Length1To127 `xml:"printOptions,omitempty"`  // minOccurs="0"
}

type TicketInformationType struct {

	// Ticketing type  TL, OK, DO, IN, MA, TR, AT, PT, XL, ST, SS
	Indicator formats.AlphaString_Length2To2 `xml:"indicator"`

	// Ticketing date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`  // minOccurs="0"

	// Ticketing time
	Time formats.Time24_HHMM `xml:"time,omitempty"`  // minOccurs="0"

	// Office Id
	OfficeId formats.AlphaNumericString_Length1To9 `xml:"officeId,omitempty"`  // minOccurs="0"

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To15 `xml:"freetext,omitempty"`  // minOccurs="0"

	// Airline code
	AirlineCode formats.AlphaNumericString_Length1To3 `xml:"airlineCode,omitempty"`  // minOccurs="0"

	// Queue number
	QueueNumber formats.AlphaNumericString_Length1To3 `xml:"queueNumber,omitempty"`  // minOccurs="0"

	// Category number
	QueueCategory formats.AlphaNumericString_Length1To3 `xml:"queueCategory,omitempty"`  // minOccurs="0"

	// SITA Addresses
	SitaAddress []formats.AlphaNumericString_Length7To7 `xml:"sitaAddress,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type TicketingCarrierDesignatorType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// airline code and printer details
	Carrier *TicketingCarrierType `xml:"carrier,omitempty"`  // minOccurs="0"
}

type TicketingCarrierType struct {

	// Airline code of sponsor
	AirlineCode formats.AlphaNumericString_Length1To3 `xml:"airlineCode"`

	// Print itinerary option   IBP, IEP, IBPJ, IEPJ
	OptionInfo formats.AlphaNumericString_Length1To4 `xml:"optionInfo,omitempty"`  // minOccurs="0"

	// Printer number
	PrinterNumber formats.AlphaNumericString_Length1To8 `xml:"printerNumber,omitempty"`  // minOccurs="0"

	// ISO code 639 - 1988
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`  // minOccurs="0"
}

type TourCodeType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// Formatted tour code
	FormatedTour *FormatedTourCodeType `xml:"formatedTour,omitempty"`  // minOccurs="0"

	// Net remit
	NetRemit *NetRemitTourCodeType `xml:"netRemit,omitempty"`  // minOccurs="0"

	// Freeformat Tour information
	FreeFormatTour *FreeFormatTourCodeType `xml:"freeFormatTour,omitempty"`  // minOccurs="0"
}

type TravelProductInformationType struct {

	// date and time details
	Product *ProductDateTimeTypeI `xml:"product,omitempty"`  // minOccurs="0"

	// boardpoint details
	BoardpointDetail *LocationTypeI `xml:"boardpointDetail,omitempty"`  // minOccurs="0"

	// offpoint details
	OffpointDetail *LocationTypeI `xml:"offpointDetail,omitempty"`  // minOccurs="0"

	// airline or system code
	Company *CompanyIdentificationTypeI `xml:"company,omitempty"`  // minOccurs="0"

	// product details - number and class
	ProductDetails *ProductIdentificationDetailsTypeI `xml:"productDetails,omitempty"`  // minOccurs="0"

	// Product type details. Here: availibility context.
	FlightTypeDetails *ProductTypeDetailsType `xml:"flightTypeDetails,omitempty"`  // minOccurs="0"

	// 1. Air segment  To indicate an Informational Air segment:  N for No action required
	ProcessingIndicator formats.AlphaNumericString_Length1To3 `xml:"processingIndicator,omitempty"`  // minOccurs="0"
}

type TravellerDetailsTypeI struct {

	// Traveler First Name + titel
	FirstName formats.AlphaNumericString_Length1To56 `xml:"firstName,omitempty"`  // minOccurs="0"

	// Traveler Type using  Amadeus codification. In values with 'X00' (X being any letter), 00 should be replaced by any value between 01 and 99.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	// 1. Infant (INF) No more info in Edifact. 2. Infant given name only (INF/BILL) Infant given name will be placed in a 2nd occurence of C324 of this (adult) passenger TIF. The 2nd C324/6353 element will contain INF. 3. Infant given and last name (INFGATES/BILL) Infant is treated as a separate TIF following immediately this (adult) passenger TIF. This following TIF C324/6353 element will contain INF.
	InfantIndicator formats.AlphaNumericString_Length1To1 `xml:"infantIndicator,omitempty"`  // minOccurs="0"

	// Identification code, 2 cases: ID < 1 to 51 char free text ) or CR < 1 to 40 char free text )
	IdentificationCode formats.AlphaNumericString_Length1To70 `xml:"identificationCode,omitempty"`  // minOccurs="0"
}

type TravellerInformationTypeI struct {

	// traveller surname,type and quantity
	Traveller *TravellerSurnameInformationTypeI `xml:"traveller"`

	// Occurrence one relates to the traveler.  Occurrence 2 relates only to an infant accompanying the traveler for whom only the given name is present.
	Passenger []*TravellerDetailsTypeI `xml:"passenger,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type TravellerSurnameInformationTypeI struct {

	// Traveler Last Name or Group name
	Surname formats.AlphaNumericString_Length1To57 `xml:"surname"`

	// G for a group. (The traveler type is in C324/6353)
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// - 1 : only one traveler defined by TIFwith exceptions below. - 2 : the traveler is accompanied by  an infant for whom only the given name is present.
	Quantity *formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`  // minOccurs="0"
}
