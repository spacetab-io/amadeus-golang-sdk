package pnr_addmultielements

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type PNRAddMultiElements struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A PNR_AddMultiElements"`

	// To specify a reference to a reservation
	ReservationInfo *ReservationControlInformationTypeI `xml:"reservationInfo,omitempty"`

	// To specify specific Actions to be processed on PNR
	PnrActions *OptionalPNRActionsType `xml:"pnrActions,omitempty"`

	TravellerInfo *TravellerInfo `xml:"travellerInfo,omitempty"`

	OriginDestinationDetails *OriginDestinationDetails `xml:"originDestinationDetails,omitempty"`

	DataElementsMaster *DataElementsMaster `xml:"dataElementsMaster,omitempty"`
}

type DataElementsMaster struct {

	// marker
	Marker1 *DummySegmentTypeI `xml:"marker1,omitempty"`

	DataElementsIndiv *DataElementsIndiv `xml:"dataElementsIndiv,omitempty"`
}

type DataElementsIndiv struct {

	// To specify the PNR segments/elements references and action to apply
	ElementManagementData *ElementManagementSegmentType `xml:"elementManagementData,omitempty"`

	// To specify the Amadeus PNR Individual Security element Note: up to 3 ISI segments may be needed to represent the ES element and its 3 types of receiver
	PnrSecurity *IndividualPnrSecurityInformationType `xml:"pnrSecurity,omitempty"`

	// specify amadeus accounting informations
	Acccounting *AccountingInformationElementType `xml:"acccounting,omitempty"`

	// To specify different kinds of remarks
	MiscellaneousRemark *MiscellaneousRemarksType `xml:"miscellaneousRemark,omitempty"`

	// To specify special requests or services information relating to a traveller.
	ServiceRequest *SpecialRequirementsDetailsTypeI `xml:"serviceRequest,omitempty"`

	// To provide date and time details relative to flight movements
	DateAndTimeInformation *DateAndTimeInformationTypeI `xml:"dateAndTimeInformation,omitempty"`

	// To specify the Amadeus PNR Tour Code element
	TourCode *TourCodeType `xml:"tourCode,omitempty"`

	// To specify an Amadeus PNR Ticket element
	TicketElement *TicketElementType `xml:"ticketElement,omitempty"`

	// To provide free form or coded long text information.
	FreetextData *LongFreeTextType `xml:"freetextData,omitempty"`

	// To specify the way data are mapped for structured
	StructuredAddress *StructuredAddressType `xml:"structuredAddress,omitempty"`

	// To specify an Amadeus PNR Option element
	OptionElement *OptionElementType `xml:"optionElement,omitempty"`

	// To request a Hard Copy Print
	Printer *PrinterIdentificationType `xml:"printer,omitempty"`

	// The group handles Seat Request with possibly rail preferences
	SeatGroup *SeatEntityType `xml:"seatGroup,omitempty"`

	// To specify the Amadeus PNR Fare elements
	FareElement *FareElementsType `xml:"fareElement,omitempty"`

	// To specify the Amadeus PNR Fare discount element
	FareDiscount *FareDiscountElementType `xml:"fareDiscount,omitempty"`

	// To specify the Amadeus PNR Manual Document Registration element
	ManualFareDocument *ManualDocumentRegistrationType `xml:"manualFareDocument,omitempty"`

	// To specify the Amadeus PNR Commission element
	Commission *CommissionElementType `xml:"commission,omitempty"`

	// To specify the Amadeus PNR Original Issue / Issue in Exchange For element
	OriginalIssue *OriginalIssueType `xml:"originalIssue,omitempty"`

	// To convey details describing the form of payment
	FormOfPayment *FormOfPaymentTypeI `xml:"formOfPayment,omitempty"`

	// To convey additional details of the form of payment
	FopExtension *MarketSpecificDataType `xml:"fopExtension,omitempty"`

	// To convey the FOP service details
	ServiceDetails *StatusTypeI `xml:"serviceDetails,omitempty"`

	// To specify frequent traveler verification
	FrequentTravellerVerification *FrequentTravellerVerificationType `xml:"frequentTravellerVerification,omitempty"`

	// To specify the Amadeus PNR Ticketing Carrier Designator element
	TicketingCarrier *TicketingCarrierDesignatorType `xml:"ticketingCarrier,omitempty"`

	// To specify the Amadeus PNR Fare Print Override element
	FarePrintOverride *FarePrintOverrideType `xml:"farePrintOverride,omitempty"`

	// To convey frequent traveler program information relating to a specific traveller or group.
	FrequentTravellerData *FrequentTravellerInformationTypeU `xml:"frequentTravellerData,omitempty"`

	// To specify access level of an entity (office) to the element.
	AccessLevel *ExtendedOwnershipSecurityDetailsType `xml:"accessLevel,omitempty"`

	// To provide specific reference identification
	ReferenceForDataElement *ReferenceInfoType `xml:"referenceForDataElement,omitempty"`
}

type OriginDestinationDetails struct {

	// To convey the origin and destination of a journey
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`

	ItineraryInfo *ItineraryInfo `xml:"itineraryInfo,omitempty"`
}

type ItineraryInfo struct {

	// To specify the PNR segments/elements references and action to apply
	ElementManagementItinerary *ElementManagementSegmentType `xml:"elementManagementItinerary,omitempty"`

	AirAuxItinerary *AirAuxItinerary `xml:"airAuxItinerary,omitempty"`

	// To provide specific reference identification
	ReferenceForSegment *ReferenceInfoType `xml:"referenceForSegment,omitempty"`
}

type AirAuxItinerary struct {

	// To specify details related to a product
	TravelProduct *TravelProductInformationType `xml:"travelProduct,omitempty"`

	// To specify the message type and business function
	MessageAction *MessageActionDetailsTypeI `xml:"messageAction,omitempty"`

	// To indicate quantity and action required in relation to a product
	RelatedProduct *RelatedProductInformationTypeI `xml:"relatedProduct,omitempty"`

	// To specify the details for making a selection
	SelectionDetailsAir *SelectionDetailsTypeI `xml:"selectionDetailsAir,omitempty"`

	// To specify a reference to a reservation
	ReservationInfoSell *ReservationControlInformationTypeI `xml:"reservationInfoSell,omitempty"`

	// To provide free form or coded long text information.
	FreetextItinerary *LongFreeTextType `xml:"freetextItinerary,omitempty"`
}

type TravellerInfo struct {

	// To specify the PNR segments/elements references and action to apply
	ElementManagementPassenger *ElementManagementSegmentType `xml:"elementManagementPassenger,omitempty"`

	PassengerData *PassengerData `xml:"passengerData,omitempty"`
}

type PassengerData struct {

	// To specify a traveler(s) and personal details relating to the traveler(s).  In values with 'X00' (X being any letter), 00 should be replaced by any value between 01 and 99.
	TravellerInformation *TravellerInformationTypeI `xml:"travellerInformation,omitempty"`

	// Passenger date of birth (DDMMYYYY) If the passenger has an infant, not in a separate TIF, then the date is used for the infant date of birth.
	DateOfBirth *DateAndTimeInformationType `xml:"dateOfBirth,omitempty"`
}

type AccountingElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A AccountingElementType"`

	// Account number
	Number formats.AlphaNumericString_Length1To10 `xml:"number,omitempty"`

	// Cost Number
	CostNumber formats.AlphaNumericString_Length1To60 `xml:"costNumber,omitempty"`

	// IATA company number
	CompanyNumber formats.AlphaNumericString_Length1To12 `xml:"companyNumber,omitempty"`

	// Client Reference Number
	ClientReference formats.AlphaNumericString_Length1To60 `xml:"clientReference,omitempty"`

	GSTTaxDetails formats.AlphaNumericString_Length1To109 `xml:"gSTTaxDetails,omitempty"`
}

type AccountingInformationElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A AccountingInformationElementType"`

	// One of these 4 data elements is mandatory , but non in particular
	Account *AccountingElementType `xml:"account,omitempty"`

	// Number of units qualifier
	AccountNumberOfUnits formats.AlphaNumericString_Length1To3 `xml:"accountNumberOfUnits,omitempty"`
}

type ClassDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ClassDetailsType"`

	// For the booking class code.
	Code formats.AlphaNumericString_Length1To1 `xml:"code,omitempty"`

	BookingClass formats.AlphaNumericString_Length2To2 `xml:"bookingClass,omitempty"`
}

type CommissionElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A CommissionElementType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// Commission indicator  M, C, P, CR, PR
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`

	// Commission
	CommissionInfo *CommissionInformationType `xml:"commissionInfo,omitempty"`

	// Old commission
	OldCommission *CommissionInformationType_6428C `xml:"oldCommission,omitempty"`

	// Manual capping Amount (after tag C)
	ManualCapping formats.NumericDecimal_Length1To10 `xml:"manualCapping,omitempty"`
}

type CommissionInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A CommissionInformationType"`

	// Percentage (max 2 decimals)
	Percentage formats.NumericDecimal_Length1To5 `xml:"percentage,omitempty"`

	// Amount (before tag A)
	Amount formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`

	// VAT indicator  V
	VatIndicator formats.AlphaNumericString_Length1To2 `xml:"vatIndicator,omitempty"`

	// Net remit indicator  N
	RemitIndicator formats.AlphaNumericString_Length1To2 `xml:"remitIndicator,omitempty"`
}

type CommissionInformationType_6428C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A CommissionInformationType_6428C"`

	// Percentage (max 2 decimals)
	Percentage formats.NumericInteger_Length1To5 `xml:"percentage,omitempty"`

	// Amount (before tag A)
	Amount formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`

	// VAT indicator  V
	VatIndicator formats.AlphaNumericString_Length1To2 `xml:"vatIndicator,omitempty"`

	// Net remit indicator  N
	RemitIndicator formats.AlphaNumericString_Length1To2 `xml:"remitIndicator,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A CompanyIdentificationTypeI"`

	// 1. Air segment: Airline code 2. ATX segment: Airline code of the airline to take action 3. CAR segment: Airline code of the airline to take action 4. HTL segment: Airline code of the airline to take action
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification,omitempty"`

	// To convey a second carrier (e.g. in case of multi airline open segments)
	SecondIdentification formats.AlphaNumericString_Length1To3 `xml:"secondIdentification,omitempty"`
}

type CompanyIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A CompanyIdentificationType"`

	// To specify the frequent traveller program company code
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`

	// To specify the frequent traveller program partnership company  code
	PartnerCode formats.AlphaNumericString_Length1To3 `xml:"partnerCode,omitempty"`

	// To specify the frequent traveller program other partnership company  code
	OtherPartnerCode formats.AlphaNumericString_Length1To3 `xml:"otherPartnerCode,omitempty"`
}

type DateAndTimeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A DateAndTimeDetailsTypeI"`

	// Miscellaneous Charge Order element : date.- SEAT ssr : date of cahnge of gauge - gourp seat SSR : date of change of gauge
	FirstDate formats.Date_DDMMYY `xml:"firstDate,omitempty"`

	// Miscellaneous Charge Order element : ARC carrier code code.
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`

	// Miscellaneous Charge Order element : ARC city code.
	LocationIdentification formats.AlphaNumericString_Length1To3 `xml:"locationIdentification,omitempty"`
}

type DateAndTimeDetailsTypeI_56946C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A DateAndTimeDetailsTypeI_56946C"`

	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Inf/Child date of birth
	Date formats.AlphaNumericString_Length1To8 `xml:"date,omitempty"`
}

type DateAndTimeInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A DateAndTimeInformationTypeI"`

	// date and time details
	DateAndTime *DateAndTimeDetailsTypeI `xml:"dateAndTime,omitempty"`
}

type DateAndTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A DateAndTimeInformationType"`

	// DATE AND TIME DETAILS
	DateAndTimeDetails *DateAndTimeDetailsTypeI_56946C `xml:"dateAndTimeDetails,omitempty"`
}

type DiscountInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A DiscountInformationType"`

	// Discount code
	AdjustmentReason formats.AlphaNumericString_Length1To6 `xml:"adjustmentReason,omitempty"`

	// Discount percentage
	Percentage formats.NumericInteger_Length2To2 `xml:"percentage,omitempty"`

	// Status code
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Staff employee number
	StaffNumber formats.AlphaNumericString_Length1To7 `xml:"staffNumber,omitempty"`

	// Staff employee name
	StaffName formats.AlphaString_Length1To30 `xml:"staffName,omitempty"`
}

type DummySegmentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A DummySegmentTypeI"`
}

type ElementManagementSegmentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ElementManagementSegmentType"`

	// segments/elements references - type and number
	Reference *ReferencingDetailsType `xml:"reference,omitempty"`

	// PNR segment or element name
	SegmentName formats.AlphaNumericString_Length1To3 `xml:"segmentName,omitempty"`
}

type ExtendedOwnershipSecurityDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ExtendedOwnershipSecurityDetailsType"`

	// This composite contains security data between entities or related to one entity.
	SecurityDetails *ExtendedSecurityDetailsType `xml:"securityDetails,omitempty"`
}

type ExtendedSecurityDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ExtendedSecurityDetailsType"`

	// Used to specify which kind of entities is specified. F - for family
	TypeOfEntity formats.AlphaNumericString_Length1To3 `xml:"typeOfEntity,omitempty"`

	// Used to specify the access mode regarding agreement and entities
	AccessMode formats.AlphaNumericString_Length1To1 `xml:"accessMode,omitempty"`

	// Used to specify entity on which the detailed security applied. Mask is specified useing *, for instance, corporate 1A0 is specified as follows : ***1A0***.
	InhouseIdentification formats.AlphaNumericString_Length1To9 `xml:"inhouseIdentification,omitempty"`
}

type FareDiscountElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FareDiscountElementType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// To specify the discount details. Only 1 repetition must be used. The Fare Discount element cannot process multiple discounts. If you wish to enter multiple discounts for a passenger, you should enter several FD elements and associate them to the same passenger.
	Discount *DiscountInformationType `xml:"discount,omitempty"`

	// Date of birth
	BirthDate formats.AlphaNumericString_Length6To6 `xml:"birthDate,omitempty"`

	// number of years or months
	NumberDetail *NumberOfUnitDetailsType `xml:"numberDetail,omitempty"`

	// company ID and reference number (should not be used)
	RpInformation *RpInformationType `xml:"rpInformation,omitempty"`

	// Customer info
	Customer *ItemDetailsType `xml:"customer,omitempty"`

	// Resident Discount applicable for discount code RM,RC,DC,BP or BI
	ResidentDisc *ItemDetailsType_64822C `xml:"residentDisc,omitempty"`
}

type FareElementsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FareElementsType"`

	// E for FE - Endorsements / Restrictions element  K for FK - Shadow AIR office ID element  S for FS - Miscellaneous Ticketing Information element  Z for FZ - Miscellaneous Information element
	GeneralIndicator formats.AlphaNumericString_Length1To1 `xml:"generalIndicator,omitempty"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// 1. FK element: Office identification
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`

	// 1. FE element: Free text 2. FS element: Free text 3. FZ element: Free text
	FreetextLong formats.AlphaNumericString_Length1To126 `xml:"freetextLong,omitempty"`
}

type FarePrintOverrideDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FarePrintOverrideDetailsType"`

	// Base fare override
	BaseFare formats.AlphaNumericString_Length1To11 `xml:"baseFare,omitempty"`

	// Total fare override
	TotalFare formats.AlphaNumericString_Length1To11 `xml:"totalFare,omitempty"`

	// equivalent fare override
	EquivalentFare formats.AlphaNumericString_Length1To11 `xml:"equivalentFare,omitempty"`

	// Tax amount override
	TaxAmount formats.AlphaNumericString_Length1To11 `xml:"taxAmount,omitempty"`
}

type FarePrintOverrideType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FarePrintOverrideType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// fare override detaild
	Override *FarePrintOverrideDetailsType `xml:"override,omitempty"`
}

type FormOfPaymentDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FormOfPaymentDetailsTypeI"`

	// Form(s) of payment   CA for Cash (Amadeus code CASH)  CK for Check (Amadeus code CHECK)  CC for Credit card (Amadeus code CC)  MS for Miscellaneous (Amadeus code MS)
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To9 `xml:"amount,omitempty"`

	// Credit card code
	CreditCardCode formats.AlphaNumericString_Length1To3 `xml:"creditCardCode,omitempty"`

	// Credit card account number
	AccountNumber formats.AlphaNumericString_Length1To35 `xml:"accountNumber,omitempty"`

	// Credit card expiry date
	ExpiryDate formats.Date_MMYY `xml:"expiryDate,omitempty"`

	// Credit card approval code
	ApprovalCode formats.AlphaNumericString_Length1To8 `xml:"approvalCode,omitempty"`

	// Customer account number
	CustomerAccountNumber formats.AlphaNumericString_Length1To10 `xml:"customerAccountNumber,omitempty"`

	// Extended credit payment request  E for extended payment
	PaymentTimeReference formats.AlphaNumericString_Length1To3 `xml:"paymentTimeReference,omitempty"`

	// free text
	Freetext formats.AlphaNumericString_Length1To70 `xml:"freetext,omitempty"`

	// Currency (as there is no place elsewhere)
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`
}

type FormOfPaymentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FormOfPaymentTypeI"`

	// form of payment details
	Fop *FormOfPaymentDetailsTypeI `xml:"fop,omitempty"`
}

type FormatedTourCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FormatedTourCodeType"`

	// Tour type  IT, BT
	ProductId formats.AlphaNumericString_Length1To2 `xml:"productId,omitempty"`

	// Last digit of year in which tour code became effective
	Year formats.NumericInteger_Length1To1 `xml:"year,omitempty"`

	// Airline code of sponsor
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Conference area approval code
	ApprovalCode formats.AlphaNumericString_Length1To1 `xml:"approvalCode,omitempty"`

	// Tour identification
	PartyId formats.AlphaNumericString_Length1To8 `xml:"partyId,omitempty"`
}

type FreeFormatTourCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FreeFormatTourCodeType"`

	// Free format indicator  FF
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To14 `xml:"freetext,omitempty"`
}

type FreeTextQualificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FreeTextQualificationType"`

	// Identifies whether the free text is coded or not coded  3 for Literal text
	SubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"subjectQualifier,omitempty"`

	// 1. AP element: 2. AQ element: 3. OS element
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// Transmittable/non-transmittable indicator (S or X). Codeset list not applicable.
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Airline or system code
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`
}

type FrequentTravellerIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FrequentTravellerIdentificationTypeU"`

	// Airline Code
	CompanyId formats.AlphaNumericString_Length2To3 `xml:"companyId,omitempty"`

	// Frequent Traveller Number
	MembershipNumber formats.AlphaNumericString_Length1To25 `xml:"membershipNumber,omitempty"`
}

type FrequentTravellerInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FrequentTravellerInformationTypeU"`

	// frequent traveller identification
	FrequentTraveller *FrequentTravellerIdentificationTypeU `xml:"frequentTraveller,omitempty"`
}

type FrequentTravellerVerificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A FrequentTravellerVerificationType"`

	// A  Add FT name  N  Name verification  R   Redeem miles
	ActionRequest formats.AlphaNumericString_Length1To3 `xml:"actionRequest,omitempty"`

	// company identification
	Company *CompanyIdentificationType `xml:"company,omitempty"`

	// frequent flyer number
	Account *ProductAccountDetailsType `xml:"account,omitempty"`
}

type IndividualPnrSecurityInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A IndividualPnrSecurityInformationType"`

	// Returned before EOT or after retrieve by RTE
	Security *IndividualSecurityType `xml:"security,omitempty"`

	// Returned when retrieved
	SecurityInfo *SecurityInformationType `xml:"securityInfo,omitempty"`

	// Code as in the display:  G for Amadeus Global Core Office Identification  I for IATA number  P for Pseudo-Office Identification  Default is G.
	Indicator formats.AlphaNumericString_Length1To1 `xml:"indicator,omitempty"`
}

type IndividualSecurityType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A IndividualSecurityType"`

	// Type of receiver G:  Type of receiver I:  Type of receiver P:
	Identification formats.AlphaNumericString_Length1To14 `xml:"identification,omitempty"`

	// R for Read  B for Both read and write  N for None
	AccessMode formats.AlphaString_Length1To1 `xml:"accessMode,omitempty"`
}

type IssueInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A IssueInformationType"`

	// Numeric airline code
	AirlineCode formats.AlphaNumericString_Length3To3 `xml:"airlineCode,omitempty"`

	// {Original issue , exchange for} document number
	DocumentNumber formats.AlphaNumericString_Length1To10 `xml:"documentNumber,omitempty"`

	// {Original issue , exchange for} document check digit
	DocumentCD formats.NumericInteger_Length1To1 `xml:"documentCD,omitempty"`

	// 1st coupon number
	Coupon1 formats.AlphaNumericString_Length1To1 `xml:"coupon1,omitempty"`

	// 2nd coupon number
	Coupon2 formats.AlphaNumericString_Length1To1 `xml:"coupon2,omitempty"`

	// 3rd coupon number
	Coupon3 formats.AlphaNumericString_Length1To1 `xml:"coupon3,omitempty"`

	// 4th coupon number
	Coupon4 formats.AlphaNumericString_Length1To1 `xml:"coupon4,omitempty"`

	// Last 2 digits of the last conjunction document of the {original issue , exchange for} document
	LastConjunction formats.NumericInteger_Length2To2 `xml:"lastConjunction,omitempty"`

	// Last conjunction document of the {original issue , exchange for} document check digit
	ExchangeDocumentCD formats.NumericInteger_Length1To1 `xml:"exchangeDocumentCD,omitempty"`

	// 1st last conjunction document of the {original issue , exchange for} document coupon number
	LastConjunction1 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction1,omitempty"`

	// 2nd last conjunction document of the {original issue , exchange for} document coupon number
	LastConjunction2 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction2,omitempty"`

	// 3rd last conjunction document of the {original issue , exchange for} document 4th coupon number
	LastConjunction3 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction3,omitempty"`

	// 4th last conjunction document of the {original issue , exchange for} document 1st coupon number
	LastConjunction4 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction4,omitempty"`

	// City code of the issue
	CityCode formats.AlphaString_Length3To3 `xml:"cityCode,omitempty"`

	// Date of the {original , new} issue
	DateOfIssue formats.Date_DDMMYY `xml:"dateOfIssue,omitempty"`

	// IATA number
	IataNumber formats.AlphaNumericString_Length1To9 `xml:"iataNumber,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`
}

type ItemDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ItemDetailsType"`

	// Only applicable for some discount codes
	CompanyId formats.AlphaString_Length2To2 `xml:"companyId,omitempty"`

	// Only applicable for some discount codes
	CardType formats.NumericInteger_Length1To1 `xml:"cardType,omitempty"`

	// Only applicable for some discount codes
	CardNumber formats.NumericInteger_Length1To13 `xml:"cardNumber,omitempty"`

	// for PC - alpha value for other - numeric value
	CardCheck formats.AlphaNumericString_Length1To1 `xml:"cardCheck,omitempty"`

	// for PC only
	Owner formats.NumericInteger_Length2To2 `xml:"owner,omitempty"`

	// for PC only
	Version formats.NumericInteger_Length1To1 `xml:"version,omitempty"`
}

type ItemDetailsType_64822C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ItemDetailsType_64822C"`

	// DN, TR,GR,AM,CR,MR
	IdCardCode formats.AlphaString_Length2To2 `xml:"idCardCode,omitempty"`

	// T - used for TR resident discount only
	IdCardType formats.AlphaString_Length1To1 `xml:"idCardType,omitempty"`

	// Card Number
	CardNumber formats.NumericInteger_Length1To13 `xml:"cardNumber,omitempty"`

	// Card alpha check
	AlphaCheck formats.AlphaNumericString_Length1To1 `xml:"alphaCheck,omitempty"`

	// Zip code
	ZipCode formats.AlphaNumericString_Length1To9 `xml:"zipCode,omitempty"`

	// Credential certificate number
	CertificateNumber formats.AlphaNumericString_Length1To20 `xml:"certificateNumber,omitempty"`
}

type LocationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A LocationTypeI"`

	// 1. Air segment: Boarding point 2. ATX segment: Boarding point 3. CAR segment: Pick-up point city 4. HTL segment: Check-in city 5. MIS segment: City code
	CityCode formats.AlphaNumericString_Length1To5 `xml:"cityCode,omitempty"`

	// 1. TRN SNCF segment: Off point city name
	CityName formats.AlphaNumericString_Length1To17 `xml:"cityName,omitempty"`
}

type LongFreeTextType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A LongFreeTextType"`

	// free text details
	FreetextDetail *FreeTextQualificationType `xml:"freetextDetail,omitempty"`

	// Long free text information
	LongFreetext formats.AlphaNumericString_Length1To199 `xml:"longFreetext,omitempty"`
}

type ManualDocumentRegistrationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ManualDocumentRegistrationType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// documentation details
	Document *ManualDocumentType `xml:"document,omitempty"`

	// Free text
	Freeflow formats.AlphaNumericString_Length1To49 `xml:"freeflow,omitempty"`
}

type ManualDocumentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ManualDocumentType"`

	// Numeric airline code
	CompanyId formats.NumericString_Length3To3 `xml:"companyId,omitempty"`

	// Ticket number
	TicketNumber formats.NumericString_Length10To10 `xml:"ticketNumber,omitempty"`

	// Ticket number check digit
	TicketNumberCd formats.NumericInteger_Length1To1 `xml:"ticketNumberCd,omitempty"`

	// Last conjunction ticket number
	LastConjuction formats.NumericInteger_Length2To2 `xml:"lastConjuction,omitempty"`

	// Last conjunction ticket number check digit
	LastConjuctionCD formats.NumericInteger_Length1To1 `xml:"lastConjuctionCD,omitempty"`
}

type MarketSpecificDataDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A MarketSpecificDataDetailsType"`

	// Credit Card Verification data (*CV data)
	CvData formats.AlphaNumericString_Length1To10 `xml:"cvData,omitempty"`

	// Printed and reported freeflow
	PrintedFreeflow formats.AlphaNumericString_Length1To70 `xml:"printedFreeflow,omitempty"`

	// Reported but not printed freeflow
	ReportedFreeflow formats.AlphaNumericString_Length1To70 `xml:"reportedFreeflow,omitempty"`

	// Credit Card ONO data.
	OnoData formats.AlphaNumericString_Length1To15 `xml:"onoData,omitempty"`

	// Credit Card GWT data
	GwtData formats.AlphaNumericString_Length1To15 `xml:"gwtData,omitempty"`

	// Credit Card Holder name.
	ChdData formats.AlphaNumericString_Length1To40 `xml:"chdData,omitempty"`

	// Delegation code.
	DelegationCode formats.AlphaNumericString_Length1To3 `xml:"delegationCode,omitempty"`

	// MCO Fop Document Number
	McoDocNumber formats.AlphaNumericString_Length1To13 `xml:"mcoDocNumber,omitempty"`

	// MCO Fop Coupon Number
	McoCouponNumber formats.AlphaNumericString_Length1To4 `xml:"mcoCouponNumber,omitempty"`

	// MCO Fop Iata Number
	McoIataNumber formats.NumericInteger_Length1To9 `xml:"mcoIataNumber,omitempty"`

	// MCO Fop Place of Issue
	McoPlaceOfIssue formats.AlphaNumericString_Length1To3 `xml:"mcoPlaceOfIssue,omitempty"`

	// MCO Fop date of Issue. DDMMMYY
	McoDateOfIssue formats.AlphaNumericString_Length7To7 `xml:"mcoDateOfIssue,omitempty"`

	// Standard Fop Iata Number
	IataNumber formats.NumericInteger_Length1To9 `xml:"iataNumber,omitempty"`
}

type MarketSpecificDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A MarketSpecificDataType"`

	// FOP detail Sequence Number.
	FopSequenceNumber formats.NumericInteger_Length1To2 `xml:"fopSequenceNumber,omitempty"`

	// Passenger type: PAX or INF.
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// Form of Payment Market Specific Data
	NewFopsDetails *MarketSpecificDataDetailsType `xml:"newFopsDetails,omitempty"`

	// To provide extended FOP details.
	ExtFOP *ReferencingDetailsTypeI `xml:"extFOP,omitempty"`
}

type MessageActionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A MessageActionDetailsTypeI"`

	// type of segment
	Business *MessageFunctionBusinessDetailsTypeI `xml:"business,omitempty"`
}

type MessageFunctionBusinessDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A MessageFunctionBusinessDetailsTypeI"`

	// 1. Air segment:
	Function formats.AlphaNumericString_Length1To3 `xml:"function,omitempty"`
}

type MiscellaneousRemarkType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A MiscellaneousRemarkType"`

	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// This is the 3rd character (x) of the remark title RIx or RMx, or 2 letter code for RMxx, conditional for RM, not applicable for RC and RQ
	Category formats.AlphaNumericString_Length1To2 `xml:"category,omitempty"`

	// Free text and message sequence numbers of the remarks.
	Freetext formats.AlphaNumericString_Length1To127 `xml:"freetext,omitempty"`

	// Provider type (element RIA):  1 for Air provider   2 for Car provider (CCR)  3 for Hotel Provider (HHL)  M for Miscellaneous
	ProviderType formats.AlphaNumericString_Length1To3 `xml:"providerType,omitempty"`

	// MCO element : Currency code
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// MCO element : total fee amount
	Amount formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`

	// Office Id (confidential remark RC)
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`
}

type MiscellaneousRemarksType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A MiscellaneousRemarksType"`

	// miscellaneous remarks
	Remarks *MiscellaneousRemarkType `xml:"remarks,omitempty"`
}

type NetRemitTourCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A NetRemitTourCodeType"`

	// Net remit indicator  N
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`

	// Free flow text of next remit
	Freetext formats.AlphaNumericString_Length1To14 `xml:"freetext,omitempty"`
}

type NumberOfUnitDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A NumberOfUnitDetailsType"`

	// Number of  years or months
	Units formats.NumericInteger_Length1To2 `xml:"units,omitempty"`

	// YRS for year  MTH for month
	UnitsQualifier formats.AlphaNumericString_Length1To3 `xml:"unitsQualifier,omitempty"`
}

type OptionElementInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A OptionElementInformationType"`

	// Office Id
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`

	// Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// Queue number
	Queue formats.NumericInteger_Length1To3 `xml:"queue,omitempty"`

	// Category number
	Category formats.NumericInteger_Length1To3 `xml:"category,omitempty"`

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To200 `xml:"freetext,omitempty"`
}

type OptionElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A OptionElementType"`

	// queuing option details
	OptionDetail *OptionElementInformationType `xml:"optionDetail,omitempty"`
}

type OptionalPNRActionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A OptionalPNRActionsType"`

	// 0       No Special Processing 10 - 49  PNR processing options 50 - 99  PNRACC options 100 - 149  Other Opt 150 - 199  Error Processing 200 - 229  Car Opt 230 - 259  Hotel Opt 260 - 299  Air + AUX Opt 300 - 329  Ticketing Opt
	OptionCode formats.NumericInteger_Length1To3 `xml:"optionCode,omitempty"`
}

type OriginAndDestinationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A OriginAndDestinationDetailsTypeI"`

	// Airport/city code of  Origin In a Client request message, a non-blank ODI is used in an air sell request to advise that the following segments (TVL etc...) are connected. There is a maximum of 6 TVLs following a non-blank ODI.
	Origin formats.AlphaString_Length3To3 `xml:"origin,omitempty"`

	// Airport/city code of  Destination
	Destination formats.AlphaString_Length3To3 `xml:"destination,omitempty"`
}

type OriginalIssueType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A OriginalIssueType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// 8 for Voucher/Certificate indicator   RN for RN indicator
	VoucherIndicator formats.AlphaNumericString_Length1To2 `xml:"voucherIndicator,omitempty"`

	// 1st occurrence for original issue 2 occurrences for issues in exchange
	Issue *IssueInformationType `xml:"issue,omitempty"`

	// Base fare amount
	BaseFare formats.NumericDecimal_Length1To10 `xml:"baseFare,omitempty"`

	// Total tax amount
	TotalTax formats.NumericDecimal_Length1To10 `xml:"totalTax,omitempty"`

	// Penalty amount
	Penalty formats.NumericDecimal_Length1To10 `xml:"penalty,omitempty"`

	// Free flow text
	Freeflow formats.AlphaNumericString_Length1To126 `xml:"freeflow,omitempty"`
}

type PrinterIdentificationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A PrinterIdentificationDetailsType"`

	// name of the printer
	Name formats.AlphaNumericString_Length5To6 `xml:"name,omitempty"`

	// netwrok ID of the printer
	Network formats.AlphaNumericString_Length2To2 `xml:"network,omitempty"`
}

type PrinterIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A PrinterIdentificationType"`

	// identification details
	IdentifierDetail *PrinterIdentificationDetailsType `xml:"identifierDetail,omitempty"`

	// 1A office id
	Office formats.AlphaNumericString_Length9To9 `xml:"office,omitempty"`

	// IATA teletype address
	TeletypeAddress formats.AlphaNumericString_Length7To7 `xml:"teletypeAddress,omitempty"`
}

type ProductAccountDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ProductAccountDetailsType"`

	// To specify the product/account number qualifier.
	NumberQualifier formats.AlphaNumericString_Length1To3 `xml:"numberQualifier,omitempty"`

	// A code to identify a frequent traveller (e.g. a frequent traveller number)
	Number formats.AlphaNumericString_Length1To27 `xml:"number,omitempty"`
}

type ProductDateTimeTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ProductDateTimeTypeI"`

	// 1. Air segment: Departure date 2. ATX segment: Requested date 3. CAR segment: Pick-up date 4. HTL segment: Check-in date 5. MIS segment: Date for service required
	DepDate formats.Date_DDMMYY `xml:"depDate,omitempty"`

	// 1. Air segment Departure time 2. SUR segment: Pick-up time 3. TRN Amtrack segment: Departure time 4. TRN SNCF segment: Departure time
	DepTime formats.Time24_HHMM `xml:"depTime,omitempty"`

	// 1. Air segment  Arrival date (not in the display) 2. CAR segment Drop-off date 3. HTL segment: Check-out date 4. TTO segment: Return date of the Tour
	ArrDate formats.Date_DDMMYY `xml:"arrDate,omitempty"`

	// 1. Air segment  Arrival time 2. TRN Amtrack segment: Arrival time 3. TRN SNCF segment: Arrival time
	ArrTime formats.Time24_HHMM `xml:"arrTime,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ProductIdentificationDetailsTypeI"`

	// 1. Air segment: Flight number or  OPEN  ARNK - air segment arrival unknown 2. CAR se 3. segment: Car type 4. SUR segment: Transportation type  (refer to VGTVD transaction)
	Identification formats.AlphaNumericString_Length1To6 `xml:"identification,omitempty"`

	// 1. Air segment: Class of service 2. TRN Amtrack segment: Class of service 3. TRN SNCF segment: Class of service
	ClassOfService formats.AlphaString_Length1To1 `xml:"classOfService,omitempty"`

	// 1. Air segment  Flight number alpha suffix  A, B, C, D, E 2. SUR segment: Departure code  A or D 3. TRN SNCF segment: Train type
	Subtype formats.AlphaString_Length1To1 `xml:"subtype,omitempty"`

	// 1. Air segment:  N for Night class 2. TRN Amtrack segment:  N for Night class 3. TRN SNCF segment:  N for Night class
	Description formats.AlphaNumericString_Length1To1 `xml:"description,omitempty"`
}

type ProductTypeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ProductTypeDetailsType"`

	// Used to convey availibility context.
	FlightIndicator formats.AlphaNumericString_Length1To2 `xml:"flightIndicator,omitempty"`
}

type RailSeatConfigurationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A RailSeatConfigurationType"`

	// Seat space.
	SeatSpace formats.AlphaNumericString_Length2To2 `xml:"seatSpace,omitempty"`

	// Coach type.
	CoachType formats.AlphaNumericString_Length2To2 `xml:"coachType,omitempty"`

	// Seat equipment.
	SeatEquipment formats.AlphaNumericString_Length2To2 `xml:"seatEquipment,omitempty"`

	// Seat position.
	SeatPosition formats.AlphaNumericString_Length1To1 `xml:"seatPosition,omitempty"`

	// Seat direction.
	SeatDirection formats.AlphaNumericString_Length1To1 `xml:"seatDirection,omitempty"`

	// Seat deck.
	SeatDeck formats.AlphaNumericString_Length1To1 `xml:"seatDeck,omitempty"`

	// Special passenger information.
	SpecialPassengerType formats.AlphaNumericString_Length1To1 `xml:"specialPassengerType,omitempty"`
}

type RailSeatPreferencesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A RailSeatPreferencesType"`

	// Selection of the type of seat request.
	SeatRequestFunction formats.AlphaNumericString_Length1To1 `xml:"seatRequestFunction,omitempty"`

	// Seat smoking zone indicator.
	SmokingIndicator formats.AlphaString_Length1To1 `xml:"smokingIndicator,omitempty"`

	// Seat class details.
	ClassDetails *ClassDetailsType `xml:"classDetails,omitempty"`

	// Seat configuration details.
	SeatConfiguration *RailSeatConfigurationType `xml:"seatConfiguration,omitempty"`
}

type RailSeatReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A RailSeatReferenceInformationType"`

	// Rail seat reference information.
	RailSeatReferenceDetails *SeatReferenceInformationType `xml:"railSeatReferenceDetails,omitempty"`
}

type ReferenceInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ReferenceInfoType"`

	// This composite is used to transmit association information
	Reference *ReferencingDetailsType `xml:"reference,omitempty"`
}

type ReferencingDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ReferencingDetailsTypeI"`

	// Reference qualifier
	ReferenceQualifier formats.AlphaNumericString_Length1To10 `xml:"referenceQualifier,omitempty"`

	// Reference value.
	DataValue formats.AlphaNumericString_Length1To35 `xml:"dataValue,omitempty"`
}

type ReferencingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ReferencingDetailsType"`

	// Amadeus codes are used here.  PT for Passenger Tatoo reference number  PR for Passenger Client- request-message-defined reference number  ST for Segment Tatoo reference number  SR for Segment Client- request-message-defined reference number
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// refers to a PNR segment/element that had this number in its related EMS segment in the same message (qualifier PT, ST)
	Number formats.AlphaNumericString_Length1To5 `xml:"number,omitempty"`
}

type RelatedProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A RelatedProductInformationTypeI"`

	// 1. Air segment: Number of passengers 2. ATX segment: Number of passengers 3. CAR segment: Number of cars 4. CCR segment: Number of cars 5. HHL segment: Number of  rooms 6. HTL segment: Number of  rooms
	Quantity formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`

	// status
	Status formats.AlphaString_Length1To2 `xml:"status,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ReservationControlInformationDetailsTypeI"`

	// 1A
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// 1. RR element: Record locator of the original PNR  2. Record locator information: Airline record locator 3. Profile record locator information: Profile record locator 4. Air segment: Passive segment airline record locator  Due to technical limitations, RCI for air segment is truncated to 7 characters.
	ControlNumber formats.AlphaNumericString_Length1To19 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A ReservationControlInformationTypeI"`

	// reservation control information - i.e. record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}

type RpInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A RpInformationType"`

	// Airline code (should not be used)
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Numeric value (should not be used)
	ReferenceNumber formats.NumericInteger_Length1To12 `xml:"referenceNumber,omitempty"`
}

type SeatEntityType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SeatEntityType"`

	// To make Seat requests on flights within the PNR
	SeatRequest *SeatRequestType `xml:"seatRequest,omitempty"`

	// Used to convey specific seat details relative to Train for a specific request or the "near-to" seat details for a "next-to" request.
	RailSeatReferenceInformation *RailSeatReferenceInformationType `xml:"railSeatReferenceInformation,omitempty"`

	// Rail Seat Preferences
	RailSeatPreferences *RailSeatPreferencesType `xml:"railSeatPreferences,omitempty"`
}

type SeatReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SeatReferenceInformationType"`

	// Coach number.
	CoachNumber formats.AlphaNumericString_Length1To3 `xml:"coachNumber,omitempty"`

	// Deck number.
	DeckNumber formats.AlphaNumericString_Length1To3 `xml:"deckNumber,omitempty"`

	// Seat number.
	SeatNumber formats.AlphaNumericString_Length1To3 `xml:"seatNumber,omitempty"`
}

type SeatRequestType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SeatRequestType"`

	// seat requirement details
	Seat *SeatRequierementsType `xml:"seat,omitempty"`

	// Maximum 99 for Group Seat request
	Special *SeatRequierementsDataType `xml:"special,omitempty"`
}

type SeatRequierementsDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SeatRequierementsDataType"`

	// Seat number + row (seat request)  Number of seats (Group seat request)
	Data formats.AlphaNumericString_Length1To4 `xml:"data,omitempty"`

	// 3 occurrences may be used for in Amadeus seat request to indicate:  1. smoking/no smoking  2. 1st area preference  2nd area preference or passenger type
	SeatType formats.AlphaNumericString_Length1To2 `xml:"seatType,omitempty"`
}

type SeatRequierementsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SeatRequierementsType"`

	// G for group
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Type of Seat requested. S - Smoking (SMST) N - No Smoking (NSST) RQST
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// Board point
	Boardpoint formats.AlphaString_Length3To3 `xml:"boardpoint,omitempty"`

	// Off point
	Offpoint formats.AlphaString_Length3To3 `xml:"offpoint,omitempty"`
}

type SecurityInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SecurityInformationType"`

	// Date of creation
	CreationDate formats.AlphaNumericString_Length6To6 `xml:"creationDate,omitempty"`

	// Agent initials and duty code as in ORG (eg: AASU)
	AgentCode formats.AlphaNumericString_Length4To4 `xml:"agentCode,omitempty"`

	// Office Id of creation/update
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SelectionDetailsInformationTypeI"`

	// See comment below
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`
}

type SelectionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SelectionDetailsTypeI"`

	// level of sell to be processed
	Selection *SelectionDetailsInformationTypeI `xml:"selection,omitempty"`
}

type SpecialRequirementsDataDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SpecialRequirementsDataDetailsTypeI"`

	// Seat number + row (seat SSR)  Number of seats (Group seat SSR)
	Data formats.AlphaNumericString_Length1To4 `xml:"data,omitempty"`

	// 3 occurrences may be used for in Amadeus seat SSR to indicate: 1. smoking/no smoking  2. 1st area preference 3. 2nd area preference or passenger type
	SeatType formats.AlphaNumericString_Length1To2 `xml:"seatType,omitempty"`
}

type SpecialRequirementsDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SpecialRequirementsDetailsTypeI"`

	// special requirement type details
	Ssr *SpecialRequirementsTypeDetailsTypeI `xml:"ssr,omitempty"`

	// Group seat SSR cannot ask for specific seats but only smoking and/or non-smoking (see Group seat SSR). Therefore the maximum repetitions here is 9 seats (1 per passenger of non-group PNR).
	Ssrb *SpecialRequirementsDataDetailsTypeI `xml:"ssrb,omitempty"`
}

type SpecialRequirementsTypeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A SpecialRequirementsTypeDetailsTypeI"`

	// - ATA/IATA defined Special Service Requirement code.  (refer to IATA AIRIMP documentation) - SK element : Keyword
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// ATA/IATA status code Codeset list not applicable.
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Number of services requested
	Quantity formats.NumericInteger_Length1To3 `xml:"quantity,omitempty"`

	// Airline code or YY
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Processing indicator, coded. - Normal SSR   P01 request for SSR explosion at EOT ...
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Board point
	Boardpoint formats.AlphaString_Length3To3 `xml:"boardpoint,omitempty"`

	// Off point
	Offpoint formats.AlphaString_Length3To3 `xml:"offpoint,omitempty"`

	// free text data
	Freetext formats.AlphaNumericString_Length1To70 `xml:"freetext,omitempty"`
}

type StatusDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A StatusDetailsTypeI"`

	// Service indicator SV : Service Fee
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type StatusTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A StatusTypeI"`

	// To convey the status details
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails,omitempty"`
}

type StructuredAddressInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A StructuredAddressInformationType"`

	// A1 for Address line 1
	OptionA1 formats.AlphaNumericString_Length1To2 `xml:"optionA1,omitempty"`

	// A1 50 char
	OptionTextA1 formats.AlphaNumericString_Length1To50 `xml:"optionTextA1,omitempty"`
}

type StructuredAddressInformationType_5063C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A StructuredAddressInformationType_5063C"`

	// CY for company - NA for name - A2 for addr line - PO for P.O box-ZP for postacl code - CI for city - ST for state-CO for country
	Option formats.AlphaNumericString_Length1To2 `xml:"option,omitempty"`

	// CY-NA-CI -  30char : A2 - 50 char: ST-CT- 25 char : PO 8 char - ZP 9 char.
	OptionText formats.AlphaNumericString_Length1To50 `xml:"optionText,omitempty"`
}

type StructuredAddressType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A StructuredAddressType"`

	// Information type, coded  2  for billing address  P08  for general mailing address  P19  for miscellaneous mailing address  P24  for home mailing address  P25  for delivery mailing address
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// address line
	Address *StructuredAddressInformationType `xml:"address,omitempty"`

	// upto 8 possible address options
	OptionalData *StructuredAddressInformationType_5063C `xml:"optionalData,omitempty"`
}

type TicketElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TicketElementType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// general ticketing information
	Ticket *TicketInformationType `xml:"ticket,omitempty"`

	// Print options (//print options after double slash)
	PrintOptions formats.AlphaNumericString_Length1To127 `xml:"printOptions,omitempty"`
}

type TicketInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TicketInformationType"`

	// Ticketing type  TL, OK, DO, IN, MA, TR, AT, PT, XL, ST, SS
	Indicator formats.AlphaString_Length2To2 `xml:"indicator,omitempty"`

	// Ticketing date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// Ticketing time
	Time formats.Time24_HHMM `xml:"time,omitempty"`

	// Office Id
	OfficeId formats.AlphaNumericString_Length1To9 `xml:"officeId,omitempty"`

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To15 `xml:"freetext,omitempty"`

	// Airline code
	AirlineCode formats.AlphaNumericString_Length1To3 `xml:"airlineCode,omitempty"`

	// Queue number
	QueueNumber formats.AlphaNumericString_Length1To3 `xml:"queueNumber,omitempty"`

	// Category number
	QueueCategory formats.AlphaNumericString_Length1To3 `xml:"queueCategory,omitempty"`

	// SITA Addresses
	SitaAddress formats.AlphaNumericString_Length7To7 `xml:"sitaAddress,omitempty"`
}

type TicketingCarrierDesignatorType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TicketingCarrierDesignatorType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// airline code and printer details
	Carrier *TicketingCarrierType `xml:"carrier,omitempty"`
}

type TicketingCarrierType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TicketingCarrierType"`

	// Airline code of sponsor
	AirlineCode formats.AlphaNumericString_Length1To3 `xml:"airlineCode,omitempty"`

	// Print itinerary option   IBP, IEP, IBPJ, IEPJ
	OptionInfo formats.AlphaNumericString_Length1To4 `xml:"optionInfo,omitempty"`

	// Printer number
	PrinterNumber formats.AlphaNumericString_Length1To8 `xml:"printerNumber,omitempty"`

	// ISO code 639 - 1988
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`
}

type TourCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TourCodeType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// Formatted tour code
	FormatedTour *FormatedTourCodeType `xml:"formatedTour,omitempty"`

	// Net remit
	NetRemit *NetRemitTourCodeType `xml:"netRemit,omitempty"`

	// Freeformat Tour information
	FreeFormatTour *FreeFormatTourCodeType `xml:"freeFormatTour,omitempty"`
}

type TravelProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TravelProductInformationType"`

	// date and time details
	Product *ProductDateTimeTypeI `xml:"product,omitempty"`

	// boardpoint details
	BoardpointDetail *LocationTypeI `xml:"boardpointDetail,omitempty"`

	// offpoint details
	OffpointDetail *LocationTypeI `xml:"offpointDetail,omitempty"`

	// airline or system code
	Company *CompanyIdentificationTypeI `xml:"company,omitempty"`

	// product details - number and class
	ProductDetails *ProductIdentificationDetailsTypeI `xml:"productDetails,omitempty"`

	// Product type details. Here: availibility context.
	FlightTypeDetails *ProductTypeDetailsType `xml:"flightTypeDetails,omitempty"`

	// 1. Air segment  To indicate an Informational Air segment:  N for No action required
	ProcessingIndicator formats.AlphaNumericString_Length1To3 `xml:"processingIndicator,omitempty"`
}

type TravellerDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TravellerDetailsTypeI"`

	// Traveler First Name + titel
	FirstName formats.AlphaNumericString_Length1To56 `xml:"firstName,omitempty"`

	// Traveler Type using  Amadeus codification. In values with 'X00' (X being any letter), 00 should be replaced by any value between 01 and 99.
	Type_ formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// 1. Infant (INF) No more info in Edifact. 2. Infant given name only (INF/BILL) Infant given name will be placed in a 2nd occurence of C324 of this (adult) passenger TIF. The 2nd C324/6353 element will contain INF. 3. Infant given and last name (INFGATES/BILL) Infant is treated as a separate TIF following immediately this (adult) passenger TIF. This following TIF C324/6353 element will contain INF.
	InfantIndicator formats.AlphaNumericString_Length1To1 `xml:"infantIndicator,omitempty"`

	// Identification code, 2 cases: ID < 1 to 51 char free text ) or CR < 1 to 40 char free text )
	IdentificationCode formats.AlphaNumericString_Length1To70 `xml:"identificationCode,omitempty"`
}

type TravellerInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TravellerInformationTypeI"`

	// traveller surname,type and quantity
	Traveller *TravellerSurnameInformationTypeI `xml:"traveller,omitempty"`

	// Occurrence one relates to the traveler.  Occurrence 2 relates only to an infant accompanying the traveler for whom only the given name is present.
	Passenger *TravellerDetailsTypeI `xml:"passenger,omitempty"`
}

type TravellerSurnameInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRADD_11_3_1A TravellerSurnameInformationTypeI"`

	// Traveler Last Name or Group name
	Surname formats.AlphaNumericString_Length1To57 `xml:"surname,omitempty"`

	// G for a group. (The traveler type is in C324/6353)
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// - 1 : only one traveler defined by TIFwith exceptions below. - 2 : the traveler is accompanied by  an infant for whom only the given name is present.
	Quantity formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`
}
