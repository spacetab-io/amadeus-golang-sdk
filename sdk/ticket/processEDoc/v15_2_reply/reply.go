package Ticket_ProcessEDocReply_v15_2 // tatres152

import (
	//"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/sdk/formats"
)

type TicketProcessEDocReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TATRES_15_2_1A Ticket_ProcessEDocReply"`

	MsgActionDetails *MessageActionDetailsTypeI `xml:"msgActionDetails"`

	// Travel agent details at level 0
	SysProvider *TicketAgentInfoTypeI `xml:"sysProvider,omitempty"` // minOccurs="0"

	// Reservation details at level 0
	ReferenceInfo *ReservationControlInformationTypeI `xml:"referenceInfo,omitempty"` // minOccurs="0"

	// Monetary information at level 0
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo,omitempty"` // minOccurs="0"

	// Pricing details at level 0
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"` // minOccurs="0"

	// Origin - Destination at Level 0
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"` // minOccurs="0"

	// Tour code at Level 0
	TourInfo *TourInformationTypeI `xml:"tourInfo,omitempty"` // minOccurs="0"

	// Form of payment at Level 0
	Fop *FormOfPaymentTypeI `xml:"fop,omitempty"` // minOccurs="0"

	// Document count
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits,omitempty"` // minOccurs="0"

	// tax details at Level 0
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // minOccurs="0" maxOccurs="6"

	// Error at message level
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"` // minOccurs="0"

	// Coupon details
	CouponInfo *CouponInformationTypeI `xml:"couponInfo,omitempty"` // minOccurs="0"

	// Document details
	DocumentInformation []*DocumentInformationDetailsTypeI `xml:"documentInformation,omitempty"` // minOccurs="0" maxOccurs="99"

	// Freeflow text at Level 0 POD Payment On Demand indicator
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // minOccurs="0" maxOccurs="99"

	// Fare details
	FareDetails *FareInformationTypeI `xml:"fareDetails,omitempty"` // minOccurs="0"

	// Document group at level 1
	GeneralDocGroup []*GeneralDocGroup `xml:"generalDocGroup,omitempty"` // minOccurs="0" maxOccurs="999"

	// Passenger Group
	DocGroup []*DocGroup `xml:"docGroup,omitempty"` // minOccurs="0" maxOccurs="99"
}

type GeneralDocGroup struct {
	// Document details of group 1
	DocInfo *TicketNumberTypeI `xml:"docInfo"`

	// Originator details group 1
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"` // minOccurs="0"

	// error at TKT level
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"` // minOccurs="0"

	// Coupon details of group 1
	CouponInfo []*CouponInformationTypeI `xml:"couponInfo,omitempty"` // minOccurs="0" maxOccurs="4"
}

type DocGroup struct {
	// Passenger Details
	PaxInfo *TravellerInformationTypeI `xml:"paxInfo"`

	EnhancedPaxInfo *EnhancedTravellerInformationType `xml:"enhancedPaxInfo,omitempty"` // minOccurs="0"

	// Travel agent details at level 2
	SysProvider *TicketAgentInfoTypeI `xml:"sysProvider,omitempty"` // minOccurs="0"

	// Reservation details at level 2
	ReferenceInfo *ReservationControlInformationTypeI `xml:"referenceInfo,omitempty"` // minOccurs="0"

	// Monetary information at level 1
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo,omitempty"` // minOccurs="0"

	// Form of payment at level 2
	Fop *FormOfPaymentTypeI `xml:"fop,omitempty"` // minOccurs="0"

	// Pricing details at level 1
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"` // minOccurs="0"

	// Origin - Destination at Level 2
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"` // minOccurs="0"

	// Frequent traveller info at Level 2
	FrequentTravellerInfo *FrequentTravellerInformationTypeI `xml:"frequentTravellerInfo,omitempty"` // minOccurs="0"

	// Tour code L2
	TourInfo *TourInformationTypeI `xml:"tourInfo,omitempty"` // minOccurs="0"

	// Originator details group 2
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"` // minOccurs="0"

	// Coupon count
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits,omitempty"` // minOccurs="0"

	// tax details at Level 2
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // minOccurs="0" maxOccurs="5"

	// Error at Pax level
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"` // minOccurs="0"

	// Document details at Level 2
	DocumentInformation []*DocumentInformationDetailsTypeI `xml:"documentInformation,omitempty"` // minOccurs="0" maxOccurs="99"

	// Freeflow text at Level 1 RA Revenue Attribution indicator
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // minOccurs="0" maxOccurs="99"

	OtherTextInfo *FreeTextInformationType `xml:"otherTextInfo,omitempty"` // minOccurs="0"

	// form of identification at Level 2
	CustomerReference *ConsumerReferenceInformationTypeI `xml:"customerReference,omitempty"` // minOccurs="0"

	// Fare details
	FareDetails *FareInformationTypeI `xml:"fareDetails,omitempty"` // minOccurs="0"

	// contains the RFIC/RFISC
	PricingInfo *PricingTicketingSubsequentTypeI `xml:"pricingInfo,omitempty"` // minOccurs="0"

	// Document Group
	DocDetailsGroup []*DocDetailsGroup `xml:"docDetailsGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	// Structured fare calc group
	FareElementsGroup []*FareElementsGroup `xml:"fareElementsGroup,omitempty"` // minOccurs="0" maxOccurs="2"

	// Reissued Flown Group
	ReissuedFlownDetails []*ReissuedFlownDetails `xml:"reissuedFlownDetails,omitempty"` // minOccurs="0" maxOccurs="99"

	// Staff group
	StaffTravellerGroup *StaffTravellerGroup `xml:"staffTravellerGroup,omitempty"` // minOccurs="0"

	// Exchanged information Group
	OriginalIssuanceGroup []*OriginalIssuanceGroup `xml:"originalIssuanceGroup,omitempty"` // minOccurs="0" maxOccurs="2"

	// Fee group
	CarrierFeeGroup []*CarrierFeeGroup `xml:"carrierFeeGroup,omitempty"` // minOccurs="0" maxOccurs="9"
}

type DocDetailsGroup struct {
	// Document details of group 2 Conjunction Box indicator
	DocInfo *TicketNumberTypeI `xml:"docInfo"`

	// error at level 2
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"` // minOccurs="0"

	// Originator details group 3
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"` // minOccurs="0"

	// Date information at level 3
	ValidityDates *DateAndTimeInformationTypeI `xml:"validityDates,omitempty"` // minOccurs="0"

	// Freeflow text at Level 2
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // minOccurs="0" maxOccurs="2"

	// Pricing details at level 2
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"` // minOccurs="0"

	// form of identification at Level 3
	CustomerReference *ConsumerReferenceInformationTypeI `xml:"customerReference,omitempty"` // minOccurs="0"

	// Coupon Group
	CouponGroup []*CouponGroup `xml:"couponGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	FareComponentDetailsGroup []*FareComponentDetailsGroup `xml:"fareComponentDetailsGroup,omitempty"` // minOccurs="0" maxOccurs="4"
}

type CouponGroup struct {
	// Coupon details of group 4
	CouponInfo *CouponInformationTypeI `xml:"couponInfo"`

	// Travel details at Level 4
	Leg []*TravelProductInformationTypeI_29340S `xml:"leg,omitempty"` // minOccurs="0" maxOccurs="2"

	// Reservation details at level 4
	ReferenceInfo *ReservationControlInformationTypeI `xml:"referenceInfo,omitempty"` // minOccurs="0"

	// Related product information at level 4
	BookingStatus *RelatedProductInformationTypeI `xml:"bookingStatus,omitempty"` // minOccurs="0"

	// Pricing subsequent at L4
	PricingInfo *PricingTicketingSubsequentTypeI `xml:"pricingInfo,omitempty"` // minOccurs="0"

	BaggageInfo *ExcessBaggageTypeI `xml:"baggageInfo,omitempty"` // minOccurs="0"

	// Frequent traveller info at Level 4
	FrequentTravellerInfo *FrequentTravellerInformationTypeI `xml:"frequentTravellerInfo,omitempty"` // minOccurs="0"

	// Date information at level 4
	ValidityDates *DateAndTimeInformationTypeI `xml:"validityDates,omitempty"` // minOccurs="0"

	// Error at coupon level
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"` // minOccurs="0"

	// Originator details group 3
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"` // minOccurs="0"

	// Freeflow text at Level 3
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // minOccurs="0" maxOccurs="99"

	// Pricing details at level 3 Noshow indicator
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"` // minOccurs="0"

	FareBasisDetails *FareQualifierDetailsType_174783S `xml:"fareBasisDetails,omitempty"` // minOccurs="0"
}

type FareComponentDetailsGroup struct {
	// fare Component identification
	FareComponentID *ItemNumberType `xml:"fareComponentID"`

	// Market information related to fare component
	MarketFareComponent *TravelProductInformationTypeI `xml:"marketFareComponent,omitempty"` // minOccurs="0"

	// Monetary Information
	MonetaryInformation *MonetaryInformationType `xml:"monetaryInformation,omitempty"` // minOccurs="0"

	// Component Class information
	ComponentClassInfo *PricingOrTicketingSubsequentType `xml:"componentClassInfo,omitempty"` // minOccurs="0"

	// Fare Qualifier Detail
	FareQualifiersDetail *FareQualifierDetailsType `xml:"fareQualifiersDetail,omitempty"` // minOccurs="0"

	// Details of the fare family used for this fare component
	FareFamilyDetails *FareFamilyType `xml:"fareFamilyDetails,omitempty"` // minOccurs="0"

	// Carrier owner of the fare family
	FareFamilyOwner *TransportIdentifierType `xml:"fareFamilyOwner,omitempty"` // minOccurs="0"

	// The segments included in Coupon Fare information group (Tax/Fee/Charge Data at Ticket Coupon Level) have been developed for Amadeus Revenue Accounting User only as not mandatory in Reso 722f
	CouponDetailsGroup []*CouponDetailsGroup `xml:"couponDetailsGroup"` // maxOccurs="99"
}

type CouponDetailsGroup struct {
	// Tattoo + type of the product identifying the coupon.
	ProductId *ReferenceInfoType `xml:"productId"`

	// Flight Connection Type
	FlightConnectionType *TravelProductInformationType `xml:"flightConnectionType,omitempty"` // minOccurs="0"

	// Tax Details Airport Tax included
	TaxDetails *TaxType `xml:"taxDetails,omitempty"` // minOccurs="0"
}

type FareElementsGroup struct {
	// Fare component information
	FareComponentInfo *FareComponentInformationTypeI `xml:"fareComponentInfo"`

	// structured Fare Calc Group6
	FareComponentsGroup []*FareComponentsGroup `xml:"fareComponentsGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	// structured Fare Calc Monetary Information 1
	StructuredFareCalcMonInfo1 *MonetaryInformationTypeI `xml:"structuredFareCalcMonInfo1,omitempty"` // minOccurs="0"

	// Structured Fare Calc taxe information
	StructuredFareCalcTaxInfo []*TaxTypeI `xml:"structuredFareCalcTaxInfo,omitempty"` // minOccurs="0" maxOccurs="99"

	// structured Fare Calc Conversion Info
	StructuredFareCalcConversionInfo *ConversionRateTypeI `xml:"structuredFareCalcConversionInfo,omitempty"` // minOccurs="0"
}

type FareComponentsGroup struct {
	// structured Fare Calc quantity
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits"`

	// structured Fare Calc Group6
	PricedFareComponentsGroup []*PricedFareComponentsGroup `xml:"pricedFareComponentsGroup,omitempty"` // minOccurs="0" maxOccurs="99"
}

type PricedFareComponentsGroup struct {
	// structured Fare Calc item
	StructuredFareCalcItem *ItemNumberTypeI `xml:"structuredFareCalcItem"`

	// structured Fare Calc Group8
	FareCouponGroup []*FareCouponGroup `xml:"fareCouponGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	// structured Fare Calc Monetary Information 2
	StructuredFareCalcMonInfo2 *MonetaryInformationTypeI `xml:"structuredFareCalcMonInfo2,omitempty"` // minOccurs="0"

	// structured Fare Calc Pricing subsequent
	StructuredFareCalcPricingInfo *PricingTicketingSubsequentTypeI `xml:"structuredFareCalcPricingInfo,omitempty"` // minOccurs="0"

	// structured Fare Calculation details
	StructuredFareCalcDetails *FareCalculationCodeDetailsTypeI `xml:"structuredFareCalcDetails,omitempty"` // minOccurs="0"

	// structured Fare Calc rate
	StructuredFareCalcRate *FareRulesInformationTypeI `xml:"structuredFareCalcRate,omitempty"` // minOccurs="0"

	// Pricing details
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo"`

	// Fare details
	StructuredFareCalcFareDetails *FareInformationTypeI `xml:"structuredFareCalcFareDetails,omitempty"` // minOccurs="0"
}

type FareCouponGroup struct {
	// structured Fare Calc Action
	StructuredFareCalcAction *ActionDetailsTypeI `xml:"structuredFareCalcAction"`

	StructuredFareCalcLeg *TravelProductInformationTypeI_29340S `xml:"structuredFareCalcLeg,omitempty"` // minOccurs="0"
}

type ReissuedFlownDetails struct {
	// Refund status
	RefundStatus *StatusTypeI `xml:"refundStatus"`

	// Refund traveller details
	RefundLeg []*TravelProductInformationTypeI_29340S `xml:"refundLeg,omitempty"` // minOccurs="0" maxOccurs="2"

	// Pricing details
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"` // minOccurs="0"

	// Refund document details
	RefundDocumentInfo *TicketNumberTypeI `xml:"refundDocumentInfo,omitempty"` // minOccurs="0"

	// Refund coupon details
	CouponInfo []*CouponInformationTypeI `xml:"couponInfo,omitempty"` // minOccurs="0" maxOccurs="9"

	// Refund pricing subsequent
	RefundPricingInfo *PricingTicketingSubsequentTypeI `xml:"refundPricingInfo,omitempty"` // minOccurs="0"

	// Refund routing details
	RefundRoutingDetails []*RoutingInformationTypeI `xml:"refundRoutingDetails,omitempty"` // minOccurs="0" maxOccurs="99"

	// Refund product details
	RefundProductDetails *AdditionalProductDetailsTypeI `xml:"refundProductDetails,omitempty"` // minOccurs="0"
}

type StaffTravellerGroup struct {
	// To indicate the details associated with a traveller's status
	TravellerStatus *TravellerPriorityDetailsTypeI `xml:"travellerStatus"`

	// Staff information
	StaffInfo *SpecificDataInformationTypeI `xml:"staffInfo,omitempty"` // minOccurs="0"
}

type OriginalIssuanceGroup struct {
	// Exchange originator details
	OfficeIdentification *AdditionalBusinessSourceInformationTypeI `xml:"officeIdentification"`

	// Exchange details
	ExchangeDocumentDetails *DocumentInformationDetailsTypeI `xml:"exchangeDocumentDetails,omitempty"` // minOccurs="0"
}

type CarrierFeeGroup struct {
	// Fee type
	SelectionDetails *SelectionDetailsTypeI `xml:"selectionDetails"`

	// fee Sub Group
	CarrierFeeInfo []*CarrierFeeInfo `xml:"carrierFeeInfo"` // maxOccurs="99"
}

type CarrierFeeInfo struct {
	// fee sub type
	SpecialDataInfo *SpecificDataInformationTypeI `xml:"specialDataInfo"`

	// fee monetary Info
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo"`

	// fee form of payment
	Fop *FormOfPaymentTypeI `xml:"fop"`

	// Fee taxes
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // minOccurs="0" maxOccurs="99"
}

//
// Complex structs
//

type ActionDetailsTypeI struct {
	LastItemsDetails []*ReferenceTypeI `xml:"lastItemsDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type AdditionalBusinessSourceInformationTypeI struct {
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	OriginatorDetails *OriginatorIdentificationDetailsTypeI_83809C `xml:"originatorDetails,omitempty"` // minOccurs="0"

	LocationDetails *LocationTypeI `xml:"locationDetails,omitempty"` // minOccurs="0"
}

type AdditionalFareQualifierDetailsType struct {
	// xmlType: AlphaNumericString_Length1To35
	RateClass string `xml:"rateClass,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To18
	CommodityCategory string `xml:"commodityCategory,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	PricingGroup string `xml:"pricingGroup,omitempty"` // minOccurs="0"
}

type AdditionalProductDetailsTypeI struct {
	// Additional details describing a specific means of transport
	LegDetails *AdditionalProductTypeI `xml:"legDetails,omitempty"` // minOccurs="0"
}

type AdditionalProductTypeI struct {
	// UN/IATA code identifying type of aircraft (747, 737, etc.).
	// xmlType: AlphaNumericString_Length1To8
	Equipment string `xml:"equipment,omitempty"` // minOccurs="0"

	// Number of stops enroute made in a journey.
	NumberOfStops *int32 `xml:"numberOfStops,omitempty"` // minOccurs="0"
}

type ApplicationErrorDetailTypeI struct {
	// xmlType: AlphaNumericString_Length1To5
	ErrorCode string `xml:"errorCode"`

	// xmlType: AlphaNumericString_Length1To3
	ErrorCategory string `xml:"errorCategory,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"` // minOccurs="0"
}

type ApplicationErrorInformationTypeI struct {
	ErrorDetails *ApplicationErrorDetailTypeI `xml:"errorDetails"`
}

type BaggageDetailsTypeI struct {
	// Total number of units.
	FreeAllowance *int32 `xml:"freeAllowance,omitempty"` // minOccurs="0"

	// Code to qualify unit as pieces or seats.
	// xmlType: AlphaNumericString_Length1To3
	QuantityCode string `xml:"quantityCode,omitempty"` // minOccurs="0"

	// Code to qualify unit as pounds or kilos.
	// xmlType: AlphaNumericString_Length1To3
	UnitQualifier string `xml:"unitQualifier,omitempty"` // minOccurs="0"
}

type CompanyIdentificationNumbersTypeI struct {
	// xmlType: AlphaNumericString_Length1To15
	Identifier string `xml:"identifier"`

	// xmlType: AlphaNumericString_Length1To15
	OtherIdentifier string `xml:"otherIdentifier,omitempty"` // minOccurs="0"
}

type CompanyIdentificationTypeI struct {
	// Carrier owner fo the fare family
	// xmlType: AlphaNumericString_Length1To35
	OtherCompany string `xml:"otherCompany,omitempty"` // minOccurs="0"
}

type CompanyIdentificationTypeI_51997C struct {
	// xmlType: AlphaNumericString_Length1To35
	MarketingCompany string `xml:"marketingCompany,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	OperatingCompany string `xml:"operatingCompany,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	OtherCompany string `xml:"otherCompany,omitempty"` // minOccurs="0"
}

type CompanyIdentificationTypeI_83773C struct {
	// Pricing System Code
	// xmlType: AlphaNumericString_Length1To35
	PricingSystemCode string `xml:"pricingSystemCode,omitempty"` // minOccurs="0"
}

type ConsumerReferenceIdentificationTypeI struct {
	// To indicate type of passenger check-in identification number (FOID).
	// xmlType: AlphaNumericString_Length1To3
	ReferenceQualifier string `xml:"referenceQualifier"`

	// The actual form of identification number (FOID).
	// xmlType: AlphaNumericString_Length1To35
	ReferenceNumber string `xml:"referenceNumber,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	ReferencePartyName string `xml:"referencePartyName,omitempty"` // minOccurs="0"
}

type ConsumerReferenceInformationTypeI struct {
	CustomerReferences []*ConsumerReferenceIdentificationTypeI `xml:"customerReferences"` // maxOccurs="20"
}

type ConversionRateDetailsTypeI struct {
	// Rate of Exchange
	// xmlType: NumericDecimal_Length1To18
	PricingAmount *float64 `xml:"pricingAmount,omitempty"` // minOccurs="0"
}

type ConversionRateTypeI struct {
	ConversionRateDetails []*ConversionRateDetailsTypeI `xml:"conversionRateDetails"` // maxOccurs="20"
}

type CouponInformationDetailsTypeI struct {
	// Coupon number.
	// xmlType: AlphaNumericString_Length1To6
	CpnNumber string `xml:"cpnNumber,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	CpnStatus string `xml:"cpnStatus,omitempty"` // minOccurs="0"

	// Coupon amount.
	// xmlType: NumericDecimal_Length1To18
	CpnAmount *float64 `xml:"cpnAmount,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	CpnExchangeMedia string `xml:"cpnExchangeMedia,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	SettlementAuthorization string `xml:"settlementAuthorization,omitempty"` // minOccurs="0"

	// Involuntary indicator.
	// xmlType: AlphaNumericString_Length1To3
	VoluntaryIndicator string `xml:"voluntaryIndicator,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	CpnPreviousStatus string `xml:"cpnPreviousStatus,omitempty"` // minOccurs="0"

	// In connection with coupon number.
	// xmlType: AlphaNumericString_Length1To6
	CpnInConnectionWith string `xml:"cpnInConnectionWith,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	CpnSequenceNumber string `xml:"cpnSequenceNumber,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	CpnInConnectionWithQualifier string `xml:"cpnInConnectionWithQualifier,omitempty"` // minOccurs="0"
}

type CouponInformationTypeI struct {
	CouponDetails []*CouponInformationDetailsTypeI `xml:"couponDetails"` // maxOccurs="4"
}

type DataInformationTypeI struct {
	// The carrier fee application code
	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"` // minOccurs="0"
}

type DataTypeInformationTypeI struct {
	// To specify the type of data, i.e. staff reservation entitlement, status, type of travel, or ticket type
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type"`

	// Reservation entitlement/status, type of travel, ticket type
	// xmlType: AlphaNumericString_Length1To3
	StatusEvent string `xml:"statusEvent,omitempty"` // minOccurs="0"
}

type DateAndTimeDetailsTypeI struct {
	// To identify type of date to follow.
	// xmlType: AlphaNumericString_Length1To3
	Qualifier string `xml:"qualifier,omitempty"` // minOccurs="0"

	// Valid date (ddmmyy).
	// xmlType: AlphaNumericString_Length1To35
	Date string `xml:"date,omitempty"` // minOccurs="0"

	// A time (hhmm).
	// xmlType: AlphaNumericString_Length1To4
	Time string `xml:"time,omitempty"` // minOccurs="0"
}

type DateAndTimeInformationTypeI struct {
	DateAndTimeDetails []*DateAndTimeDetailsTypeI `xml:"dateAndTimeDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type DiscountPenaltyInformationType struct {
	FareQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"fareQualifier,omitempty"` // minOccurs="0"
}

type DocumentDetailsTypeI struct {
	// Certificate number or original issue/ticket document number.
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"` // minOccurs="0"

	// Original issue date (ddmmmyy)
	// xmlType: AlphaNumericString_Length1To35
	Date string `xml:"date,omitempty"` // minOccurs="0"
}

type DocumentInformationDetailsTypeI struct {
	DocumentDetails *DocumentDetailsTypeI `xml:"documentDetails"`
}

type EnhancedTravellerInformationType struct {
	// Name attributes unique for one passenger.
	TravellerNameInfo *TravellerNameInfoType `xml:"travellerNameInfo,omitempty"` // minOccurs="0"

	// 5 possible types of names, for 1 passenger.
	OtherPaxNamesDetails []*TravellerNameDetailsType `xml:"otherPaxNamesDetails"` // maxOccurs="5"
}

type ExcessBaggageDetailsTypeI struct {
	// Currency of the excess baggage charge.
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"

	// Excess baggage rate per unit.
	// xmlType: NumericDecimal_Length1To18
	Amount *float64 `xml:"amount,omitempty"` // minOccurs="0"
}

type ExcessBaggageTypeI struct {
	ExcessBaggageDetails *ExcessBaggageDetailsTypeI `xml:"excessBaggageDetails,omitempty"` // minOccurs="0"

	BaggageDetails []*BaggageDetailsTypeI `xml:"baggageDetails,omitempty"` // minOccurs="0" maxOccurs="3"
}

type FareCalculationCodeDetailsTypeI struct {
	// Fare calculation descriptive indicators
	// xmlType: AlphaNumericString_Length1To3
	ChargeCategory string `xml:"chargeCategory,omitempty"` // minOccurs="0"

	// Charge or fare calculation amount
	// xmlType: NumericDecimal_Length1To18
	Amount *float64 `xml:"amount,omitempty"` // minOccurs="0"

	// May apply to a specific airport/city
	// xmlType: AlphaNumericString_Length1To25
	LocationCode []string `xml:"locationCode,omitempty"` // minOccurs="0" maxOccurs="2"

	// Percentage amount
	// xmlType: NumericDecimal_Length1To8
	Rate *float64 `xml:"rate,omitempty"` // minOccurs="0"
}

type FareComponentDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	DataType string `xml:"dataType,omitempty"` // minOccurs="0"

	// Fare component count
	Count *int32 `xml:"count,omitempty"` // minOccurs="0"

	// Price quote date
	// xmlType: AlphaNumericString_Length1To35
	PricingDate string `xml:"pricingDate,omitempty"` // minOccurs="0"

	// Account code
	// xmlType: AlphaNumericString_Length1To35
	AccountCode string `xml:"accountCode,omitempty"` // minOccurs="0"

	// Input designator
	// xmlType: AlphaNumericString_Length1To35
	InputDesignator string `xml:"inputDesignator,omitempty"` // minOccurs="0"
}

type FareComponentInformationTypeI struct {
	FareComponentDetails *FareComponentDetailsTypeI `xml:"fareComponentDetails,omitempty"` // minOccurs="0"

	// Ticket document number
	// xmlType: AlphaNumericString_Length1To35
	TicketNumber string `xml:"ticketNumber,omitempty"` // minOccurs="0"
}

type FareFamilyDetailsType struct {
	// Commercial fare Family Short name
	// xmlType: AlphaNumericString_Length1To30
	CommercialFamily string `xml:"commercialFamily"`
}

type FareFamilyType struct {
	// Fare Family Short Name
	// xmlType: AlphaNumericString_Length1To30
	FareFamilyname string `xml:"fareFamilyname,omitempty"` // minOccurs="0"

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy *int32 `xml:"hierarchy,omitempty"` // minOccurs="0"

	// Indicates Commercial Fare Family Short names
	CommercialFamilyDetails []*FareFamilyDetailsType `xml:"commercialFamilyDetails,omitempty"` // minOccurs="0" maxOccurs="20"
}

type FareInformationTypeI struct {
	// Specifies an industry defined fare component priced passenger type code (PTC).
	// xmlType: AlphaNumericString_Length1To3
	ValueQualifier string `xml:"valueQualifier,omitempty"` // minOccurs="0"

	FareTypeGrouping *FareTypeGroupingInformationTypeI `xml:"fareTypeGrouping,omitempty"` // minOccurs="0"
}

type FareQualifierDetailsType struct {
	DiscountDetails []*DiscountPenaltyInformationType `xml:"discountDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type FareQualifierDetailsType_174783S struct {
	AdditionalFareDetails *AdditionalFareQualifierDetailsType `xml:"additionalFareDetails,omitempty"` // minOccurs="0"
}

type FareRulesInformationTypeI struct {
	// Tariff identification
	// xmlType: AlphaNumericString_Length1To9
	TariffClassId string `xml:"tariffClassId,omitempty"` // minOccurs="0"

	CompanyDetails *CompanyIdentificationTypeI_83773C `xml:"companyDetails,omitempty"` // minOccurs="0"

	// Fare rule paragraph number code
	// xmlType: AlphaNumericString_Length1To7
	RuleSectionId []string `xml:"ruleSectionId,omitempty"` // minOccurs="0" maxOccurs="99"
}

type FareTypeGroupingInformationTypeI struct {
	// Account code
	// xmlType: AlphaNumericString_Length1To35
	PricingGroup string `xml:"pricingGroup,omitempty"` // minOccurs="0"
}

type FormOfPaymentDetailsTypeI struct {
	// Form of payment type.
	// xmlType: AlphaNumericString_Length1To10
	Type string `xml:"type"`

	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"` // minOccurs="0"

	// Form of payment amount.
	// xmlType: NumericDecimal_Length1To18
	Amount *float64 `xml:"amount,omitempty"` // minOccurs="0"

	// Vendor code (CC).
	// xmlType: AlphaNumericString_Length1To35
	VendorCode string `xml:"vendorCode,omitempty"` // minOccurs="0"

	// Account number (CC/GR/SGR).
	// xmlType: AlphaNumericString_Length1To35
	CreditCardNumber string `xml:"creditCardNumber,omitempty"` // minOccurs="0"

	// Expiration date (CC) (mmyy).
	// xmlType: AlphaNumericString_Length1To35
	ExpiryDate string `xml:"expiryDate,omitempty"` // minOccurs="0"

	// Approval code (CC).
	// xmlType: AlphaNumericString_Length1To17
	ApprovalCode string `xml:"approvalCode,omitempty"` // minOccurs="0"

	// Source of approval code (CC).
	// xmlType: AlphaNumericString_Length1To3
	SourceOfApproval string `xml:"sourceOfApproval,omitempty"` // minOccurs="0"

	// Maximum authorized amount (CC).
	// xmlType: NumericDecimal_Length1To18
	AuthorisedAmount *float64 `xml:"authorisedAmount,omitempty"` // minOccurs="0"

	// Address verification code (CC).
	// xmlType: AlphaNumericString_Length1To3
	AddressVerification string `xml:"addressVerification,omitempty"` // minOccurs="0"

	// Customer file reference.
	// xmlType: AlphaNumericString_Length1To70
	CustomerAccount string `xml:"customerAccount,omitempty"` // minOccurs="0"

	// Extended payment code (CC).
	// xmlType: AlphaNumericString_Length1To3
	ExtendedPayment string `xml:"extendedPayment,omitempty"` // minOccurs="0"

	// Specifies the EMD document number that is being used for payment.
	// xmlType: AlphaNumericString_Length1To70
	FopFreeText string `xml:"fopFreeText,omitempty"` // minOccurs="0"

	// Credit card corporate contract.
	// xmlType: AlphaNumericString_Length1To3
	MembershipStatus string `xml:"membershipStatus,omitempty"` // minOccurs="0"

	// Credit card transaction information.
	// xmlType: AlphaNumericString_Length1To35
	TransactionInfo string `xml:"transactionInfo,omitempty"` // minOccurs="0"
}

type FormOfPaymentTypeI struct {
	FormOfPayment []*FormOfPaymentDetailsTypeI `xml:"formOfPayment"` // maxOccurs="99"
}

type FreeTextDetailsType struct {
	// xmlType: AlphaNumericString_Length1To3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// xmlType: AlphaNumericString_Length1To4
	InformationType string `xml:"informationType,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	Language string `xml:"language,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	Source string `xml:"source"`

	// xmlType: AlphaNumericString_Length1To3
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	// xmlType: AlphaNumericString_Length1To320
	FreeText []string `xml:"freeText"` // maxOccurs="99"
}

type FreeTextQualificationTypeI struct {
	// Function qualifier.
	// xmlType: AlphaNumericString_Length1To3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// A code describing data in 4440.
	// xmlType: AlphaNumericString_Length1To4
	InformationType string `xml:"informationType,omitempty"` // minOccurs="0"

	// Fare calculation reporting indicator or pricing indicator.
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId,omitempty"` // minOccurs="0"
}

type FrequentTravellerIdentificationTypeI struct {
	// Airline designator, coded.
	// xmlType: AlphaNumericString_Length1To35
	Carrier string `xml:"carrier"`

	// A code to identify a frequent traveller - the frequent traveller number.
	// xmlType: AlphaNumericString_Length1To25
	Number string `xml:"number"`
}

type FrequentTravellerInformationTypeI struct {
	FrequentTravellerDetails []*FrequentTravellerIdentificationTypeI `xml:"frequentTravellerDetails"` // maxOccurs="9"
}

type InteractiveFreeTextTypeI struct {
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"` // minOccurs="0"

	// Free text message.
	// xmlType: AlphaNumericString_Length1To70
	FreeText []string `xml:"freeText,omitempty"` // minOccurs="0" maxOccurs="99"
}

type InternalIDDetailsTypeI struct {
	// Reference number/authority code assigned to the requester as in the booking agent's initials or logon id
	// xmlType: AlphaNumericString_Length1To9
	InhouseId string `xml:"inhouseId"`

	// A code to identify type of agent identification.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"
}

type ItemNumberIdentificationType struct {
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"` // minOccurs="0"
}

type ItemNumberIdentificationTypeI struct {
	// Reference number or Fare component number
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"` // minOccurs="0"
}

type ItemNumberType struct {
	ItemNumberDetails []*ItemNumberIdentificationType `xml:"itemNumberDetails"` // maxOccurs="99"
}

type ItemNumberTypeI struct {
	ItemNumberDetails []*ItemNumberIdentificationTypeI `xml:"itemNumberDetails"` // maxOccurs="99"
}

type LocationDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To25
	City string `xml:"city,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	Country string `xml:"country,omitempty"` // minOccurs="0"
}

type LocationTypeI struct {
	// xmlType: AlphaNumericString_Length1To25
	TrueLocationId string `xml:"trueLocationId,omitempty"` // minOccurs="0"
}

type LocationTypeI_52002C struct {
	// xmlType: AlphaNumericString_Length1To25
	TrueLocationId string `xml:"trueLocationId,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To17
	TrueLocation string `xml:"trueLocation,omitempty"` // minOccurs="0"
}

type MarriageControlDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	Relation string `xml:"relation,omitempty"` // minOccurs="0"

	MarriageIdentifier *int32 `xml:"marriageIdentifier,omitempty"` // minOccurs="0"

	LineNumber *int32 `xml:"lineNumber,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	OtherRelation string `xml:"otherRelation,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	CarrierCode string `xml:"carrierCode,omitempty"` // minOccurs="0"
}

type MessageActionDetailsTypeI struct {
	MessageFunctionDetails *MessageFunctionBusinessDetailsTypeI `xml:"messageFunctionDetails"`

	// Indicates wether request was processed successfully.
	// xmlType: AlphaNumericString_Length1To3
	ResponseType string `xml:"responseType,omitempty"` // minOccurs="0"
}

type MessageFunctionBusinessDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	MessageFunction string `xml:"messageFunction,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"` // minOccurs="0" maxOccurs="20"
}

type MonetaryInformationDetailsType struct {
	// xmlType: AlphaNumericString_Length1To3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	// xmlType: AlphaNumericString_Length1To35
	Amount string `xml:"amount,omitempty"` // minOccurs="0"

	// Currency
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"
}

type MonetaryInformationDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount, 'FREE', 'BULK'
	// xmlType: AlphaNumericString_Length1To35
	Amount string `xml:"amount"`

	// ISO currency code.
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"

	// The From city for the carrier fee
	// xmlType: AlphaNumericString_Length1To25
	LocationFrom string `xml:"locationFrom,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To25
	LocationTo string `xml:"locationTo,omitempty"` // minOccurs="0"
}

type MonetaryInformationType struct {
	// Monetary information per fare component
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`

	// Other monetary information per fare component
	OtherMonetaryDetails []*MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"` // minOccurs="0" maxOccurs="19"
}

type MonetaryInformationTypeI struct {
	MonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"monetaryDetails"` // maxOccurs="99"
}

type NumberOfUnitDetailsTypeI struct {
	// Total number of ticket/document numbers.
	NumberOfUnit int32 `xml:"numberOfUnit"`

	// xmlType: AlphaNumericString_Length1To3
	UnitQualifier string `xml:"unitQualifier,omitempty"` // minOccurs="0"
}

type NumberOfUnitsTypeI struct {
	QuantityDetails []*NumberOfUnitDetailsTypeI `xml:"quantityDetails"` // maxOccurs="9"
}

type OriginAndDestinationDetailsTypeI struct {
	// Origin.
	// xmlType: AlphaNumericString_Length1To25
	Origin string `xml:"origin,omitempty"` // minOccurs="0"

	// Destination
	// xmlType: AlphaNumericString_Length1To25
	Destination string `xml:"destination,omitempty"` // minOccurs="0"
}

type OriginatorDetailsTypeI struct {
	// ISO Country code of the agent issuing the ticket.
	// xmlType: AlphaNumericString_Length1To3
	CodedCountry string `xml:"codedCountry,omitempty"` // minOccurs="0"

	// ISO currency code for currency of originator country.
	// xmlType: AlphaNumericString_Length1To3
	CodedCurrency string `xml:"codedCurrency,omitempty"` // minOccurs="0"

	// ISO code of language.
	// xmlType: AlphaNumericString_Length1To3
	CodedLanguage string `xml:"codedLanguage,omitempty"` // minOccurs="0"
}

type OriginatorIdentificationDetailsTypeI struct {
	// ATA/IATA ID number or pseudo IATA number.
	// xmlType: AlphaNumericString_Length1To9
	OriginatorId string `xml:"originatorId"`

	// Amadeus office identification (AMID)
	// xmlType: AlphaNumericString_Length1To9
	InHouseIdentification1 string `xml:"inHouseIdentification1,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To9
	InHouseIdentification2 string `xml:"inHouseIdentification2,omitempty"` // minOccurs="0"
}

type OriginatorIdentificationDetailsTypeI_83809C struct {
	// Original issue agent numeric code (IATA number)
	OriginatorId *int32 `xml:"originatorId,omitempty"` // minOccurs="0"
}

type OriginatorOfRequestDetailsTypeI struct {
	DeliveringSystem *SystemDetailsTypeI `xml:"deliveringSystem"`

	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification,omitempty"` // minOccurs="0"

	LocationDetails *LocationTypeI `xml:"locationDetails,omitempty"` // minOccurs="0"

	CascadingSystem *SystemDetailsTypeI_83771C `xml:"cascadingSystem,omitempty"` // minOccurs="0"

	// 1 character code for airline agent, travel agent, etc.
	// xmlType: AlphaNumericString_Length1To1
	OriginatorTypeCode string `xml:"originatorTypeCode,omitempty"` // minOccurs="0"

	OriginDetails *OriginatorDetailsTypeI `xml:"originDetails,omitempty"` // minOccurs="0"

	// A code identifying the issuing agent or if the transaction is being initiated by a computer programme then the word 'system'.
	// xmlType: AlphaNumericString_Length1To9
	Originator string `xml:"originator"`
}

type PricingOrTicketingSubsequentType struct {
	// RATE OR TARIFF CLASS INFORMATION
	FareBasisDetails *RateTariffClassInformationType `xml:"fareBasisDetails,omitempty"` // minOccurs="0"
}

type PricingTicketingDetailsTypeI struct {
	PriceTicketDetails *PricingTicketingInformationTypeI `xml:"priceTicketDetails,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	PriceTariffType string `xml:"priceTariffType,omitempty"` // minOccurs="0"

	ProductDateTimeDetails *ProductDateTimeTypeI_52035C `xml:"productDateTimeDetails,omitempty"` // minOccurs="0"

	CompanyDetails *CompanyIdentificationTypeI_51997C `xml:"companyDetails,omitempty"` // minOccurs="0"

	CompanyNumberDetails *CompanyIdentificationNumbersTypeI `xml:"companyNumberDetails,omitempty"` // minOccurs="0"

	LocationDetails *LocationDetailsTypeI `xml:"locationDetails,omitempty"` // minOccurs="0"

	OtherLocationDetails *LocationDetailsTypeI `xml:"otherLocationDetails,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	IdNumber string `xml:"idNumber,omitempty"` // minOccurs="0"

	// xmlType: NumericDecimal_Length1To18
	MonetaryAmount *float64 `xml:"monetaryAmount,omitempty"` // minOccurs="0"
}

type PricingTicketingInformationTypeI struct {
	// Ticketing Mode Indicator
	// xmlType: AlphaNumericString_Length1To3
	TicketingModeIndicator string `xml:"ticketingModeIndicator,omitempty"` // minOccurs="0"

	// International or Domestic Sales Indicator
	// xmlType: AlphaNumericString_Length1To3
	InternationalDomSalesIndicator string `xml:"internationalDomSalesIndicator,omitempty"` // minOccurs="0"

	// Statistical Code
	// xmlType: AlphaNumericString_Length1To3
	StatisticalCode string `xml:"statisticalCode,omitempty"` // minOccurs="0"

	// Self Sale Indicator
	// xmlType: AlphaNumericString_Length1To3
	SelfSaleIndicator string `xml:"selfSaleIndicator,omitempty"` // minOccurs="0"

	// Net Reporting Indicator
	// xmlType: AlphaNumericString_Length1To3
	NetReportingIndicator string `xml:"netReportingIndicator,omitempty"` // minOccurs="0"

	// Tax on Commission
	// xmlType: AlphaNumericString_Length1To3
	TaxOnCommissionIndicator string `xml:"taxOnCommissionIndicator,omitempty"` // minOccurs="0"

	// Non Endorsable indicator
	// xmlType: AlphaNumericString_Length1To3
	NonEndorsableIndicator string `xml:"nonEndorsableIndicator,omitempty"` // minOccurs="0"

	// Non Refundable indicator
	// xmlType: AlphaNumericString_Length1To3
	NonRefundableIndicator string `xml:"nonRefundableIndicator,omitempty"` // minOccurs="0"

	// Penalty Restriction
	// xmlType: AlphaNumericString_Length1To3
	PenaltyRestrictionIndicator string `xml:"penaltyRestrictionIndicator,omitempty"` // minOccurs="0"

	// Present Credit Card indicator
	// xmlType: AlphaNumericString_Length1To3
	PresentCreditCardIndicator string `xml:"presentCreditCardIndicator,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	EmergencySet string `xml:"emergencySet,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	EmergencyClear string `xml:"emergencyClear,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	NonInterlineableIndicator string `xml:"nonInterlineableIndicator,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	NonCommissionable string `xml:"nonCommissionable,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	PresentDebitCardIndicator string `xml:"presentDebitCardIndicator,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	NonReissuableIndicator string `xml:"nonReissuableIndicator,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	CarrierFeeReportingIndicator string `xml:"carrierFeeReportingIndicator,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	RefundSystemComputerCalculated string `xml:"refundSystemComputerCalculated,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	RefundManuallyCalculated string `xml:"refundManuallyCalculated,omitempty"` // minOccurs="0"

	// Noshow indicator
	// xmlType: AlphaNumericString_Length1To3
	Indicators []string `xml:"indicators,omitempty"` // minOccurs="0" maxOccurs="11"
}

type PricingTicketingSubsequentTypeI struct {
	FareBasisDetails *RateTariffClassInformationTypeI `xml:"fareBasisDetails,omitempty"` // minOccurs="0"

	// Reason for issuance code.
	// xmlType: AlphaNumericString_Length1To3
	SpecialCondition string `xml:"specialCondition,omitempty"` // minOccurs="0"

	// Reason for issuance sub code.
	// xmlType: AlphaNumericString_Length1To3
	OtherSpecialCondition string `xml:"otherSpecialCondition,omitempty"` // minOccurs="0"
}

type ProductDateTimeTypeI struct {
	// xmlType: AlphaNumericString_Length1To35
	DepartureDate string `xml:"departureDate,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To4
	DepartureTime string `xml:"departureTime,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	ArrivalDate string `xml:"arrivalDate,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To4
	ArrivalTime string `xml:"arrivalTime,omitempty"` // minOccurs="0"

	DateVariation *int32 `xml:"dateVariation,omitempty"` // minOccurs="0"
}

type ProductDateTimeTypeI_52035C struct {
	// xmlType: AlphaNumericString_Length1To35
	DepartureDate string `xml:"departureDate,omitempty"` // minOccurs="0"

	//formats: Time24_HHMM
	DepartureTime string `xml:"departureTime,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To35
	ArrivalDate string `xml:"arrivalDate,omitempty"` // minOccurs="0"

	ArrivalTime *int32 `xml:"arrivalTime,omitempty"` // minOccurs="0"

	DateVariation *int32 `xml:"dateVariation,omitempty"` // minOccurs="0"
}

type ProductIdentificationDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To35
	FlightNumber string `xml:"flightNumber"`

	// xmlType: AlphaNumericString_Length1To17
	BookingClass string `xml:"bookingClass,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	OperationalSuffix string `xml:"operationalSuffix,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To7
	Modifier []string `xml:"modifier,omitempty"` // minOccurs="0" maxOccurs="3"
}

type ProductLocationDetailsTypeI struct {
	// Enroute city/airport
	// xmlType: AlphaNumericString_Length1To25
	Station string `xml:"station,omitempty"` // minOccurs="0"
}

type ProductTypeDetailsType struct {
	// TST Connection Type
	// xmlType: AlphaNumericString_Length1To1
	FlightIndicator string `xml:"flightIndicator"`
}

type ProductTypeDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To6
	FlightIndicator []string `xml:"flightIndicator"` // maxOccurs="9"
}

type RateTariffClassInformationType struct {
	// Fare Basis Code
	// xmlType: AlphaNumericString_Length1To35
	RateTariffClass string `xml:"rateTariffClass,omitempty"` // minOccurs="0"

	// Ticket Designator
	// xmlType: AlphaNumericString_Length1To35
	OtherRateTariffClass string `xml:"otherRateTariffClass,omitempty"` // minOccurs="0"
}

type RateTariffClassInformationTypeI struct {
	// Fare basis/ticket designator.
	// xmlType: AlphaNumericString_Length1To35
	RateTariffClass string `xml:"rateTariffClass,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	RateTariffIndicator string `xml:"rateTariffIndicator,omitempty"` // minOccurs="0"
}

type ReferenceInfoType struct {
	ReferenceDetails *ReferencingDetailsType `xml:"referenceDetails"`
}

type ReferenceTypeI struct {
	// The coupon sequence number (1of3, 2of5)
	// xmlType: AlphaNumericString_Length1To6
	NumberOfItems string `xml:"numberOfItems,omitempty"` // minOccurs="0"

	// The coupon/itinerary sequence number.
	// xmlType: AlphaNumericString_Length1To35
	LastItemIdentifier []string `xml:"lastItemIdentifier,omitempty"` // minOccurs="0" maxOccurs="99"
}

type ReferencingDetailsType struct {
	// xmlType: AlphaNumericString_Length1To10
	Type string `xml:"type"`

	// xmlType: AlphaNumericString_Length1To60
	Value string `xml:"value"`
}

type RelatedProductInformationTypeI struct {
	// Sold reservation status code
	// xmlType: AlphaNumericString_Length1To3
	StatusCode []string `xml:"statusCode,omitempty"` // minOccurs="0" maxOccurs="10"
}

type ReservationControlInformationDetailsTypeI struct {
	// A 2-3 character airline/CRS code of the following record reference.
	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId,omitempty"` // minOccurs="0"

	// A reference to a record.
	// xmlType: AlphaNumericString_Length1To20
	ControlNumber string `xml:"controlNumber,omitempty"` // minOccurs="0"

	// A code to identify type of record reference.
	// xmlType: AlphaNumericString_Length1To1
	ControlType string `xml:"controlType,omitempty"` // minOccurs="0"
}

type ReservationControlInformationTypeI struct {
	Reservation []*ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"` // minOccurs="0" maxOccurs="9"
}

type RoutingInformationTypeI struct {
	RoutingDetails []*ProductLocationDetailsTypeI `xml:"routingDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type SelectionDetailsInformationTypeI struct {
	// To specify the type of carrier fee.
	// xmlType: AlphaNumericString_Length1To3
	Option string `xml:"option"`
}

type SelectionDetailsTypeI struct {
	SelectionDetails []*SelectionDetailsInformationTypeI `xml:"selectionDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type SourceTypeDetailsTypeI struct {
	// To specify this is the original issuer of the ticket.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier1 []string `xml:"sourceQualifier1"` // maxOccurs="3"
}

type SpecificDataInformationTypeI struct {
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`

	DataInformation []*DataInformationTypeI `xml:"dataInformation,omitempty"` // minOccurs="0" maxOccurs="99"
}

type StatusDetailsTypeI struct {
	// To specify that this is reissued flown flight coupon data.
	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"` // minOccurs="0"
}

type StatusTypeI struct {
	StatusDetails []*StatusDetailsTypeI `xml:"statusDetails"` // maxOccurs="99"
}

type SystemDetailsTypeI struct {
	// A 2-3 character airline/CRS code, or bilaterally agreed code, of the system that delivers the message.
	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId"`

	// 3 character ATA/IATA airport/city code of the delivering system physical location.
	// xmlType: AlphaNumericString_Length1To25
	LocationId string `xml:"locationId,omitempty"` // minOccurs="0"
}

type SystemDetailsTypeI_83771C struct {
	// A 2-3 character airline/CRS code, or bilaterally agree code, of the system originating the message, WHEN DIFFERENT FROM THE DELIVERING SYSTEM.
	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId,omitempty"` // minOccurs="0"

	// 3 character ATA/IATA airport/city code of the system that originates the message.
	// xmlType: AlphaNumericString_Length1To25
	LocationId string `xml:"locationId,omitempty"` // minOccurs="0"
}

type TaxDetailsType struct {
	// xmlType: AlphaNumericString_Length1To17
	Rate string `xml:"rate,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	CountryCode string `xml:"countryCode,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	CurrencyCode string `xml:"currencyCode,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	Type []string `xml:"type,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TaxDetailsTypeI struct {
	// Tax/fee/charge amount.
	// xmlType: AlphaNumericString_Length1To17
	Rate string `xml:"rate,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	CountryCode string `xml:"countryCode,omitempty"` // minOccurs="0"

	// ISO code identifying currency.
	// xmlType: AlphaNumericString_Length1To3
	CurrencyCode string `xml:"currencyCode,omitempty"` // minOccurs="0"

	// Tax/fee/charge type.
	// xmlType: AlphaNumericString_Length1To3
	Type []string `xml:"type,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TaxType struct {
	// xmlType: AlphaNumericString_Length1To3
	TaxCategory string `xml:"taxCategory,omitempty"` // minOccurs="0"

	TaxDetails []*TaxDetailsType `xml:"taxDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TaxTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	TaxCategory string `xml:"taxCategory,omitempty"` // minOccurs="0"

	TaxDetails []*TaxDetailsTypeI `xml:"taxDetails"` // maxOccurs="99"
}

type TicketAgentInfoTypeI struct {
	// Service airline/system provider identifier e.g. the airline passenger accounting and modulus 7
	// xmlType: AlphaNumericString_Length1To15
	CompanyIdNumber string `xml:"companyIdNumber,omitempty"` // minOccurs="0"

	InternalIdDetails []*InternalIDDetailsTypeI `xml:"internalIdDetails,omitempty"` // minOccurs="0" maxOccurs="5"

	// Booking agent IATA number
	BookingIataNumber *int32 `xml:"bookingIataNumber,omitempty"` // minOccurs="0"
}

type TicketNumberDetailsTypeI struct {
	// Document number.
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"` // minOccurs="0"

	// Document type.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"

	// Total number of booklets issued.
	NumberOfBooklets *int32 `xml:"numberOfBooklets,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	DataIndicator string `xml:"dataIndicator,omitempty"` // minOccurs="0"

	// In connection with document number
	// xmlType: AlphaNumericString_Length1To35
	InConnectionWith string `xml:"inConnectionWith,omitempty"` // minOccurs="0"
}

type TicketNumberTypeI struct {
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`

	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"` // minOccurs="0"
}

type TourDetailsTypeI struct {
	// Tour code.
	// xmlType: AlphaNumericString_Length1To35
	TourCode string `xml:"tourCode,omitempty"` // minOccurs="0"
}

type TourInformationTypeI struct {
	TourInformationDetails *TourDetailsTypeI `xml:"tourInformationDetails,omitempty"` // minOccurs="0"
}

type TransportIdentifierType struct {
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"` // minOccurs="0"
}

type TravelProductInformationType struct {
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"` // minOccurs="0"

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"` // minOccurs="0"

	// TST Connection Type
	FlightTypeDetails *ProductTypeDetailsType `xml:"flightTypeDetails,omitempty"` // minOccurs="0"
}

type TravelProductInformationTypeI struct {
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"` // minOccurs="0"

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"` // minOccurs="0"
}

type TravelProductInformationTypeI_29340S struct {
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"` // minOccurs="0"

	BoardPointDetails *LocationTypeI_52002C `xml:"boardPointDetails,omitempty"` // minOccurs="0"

	OffpointDetails *LocationTypeI_52002C `xml:"offpointDetails,omitempty"` // minOccurs="0"

	CompanyDetails *CompanyIdentificationTypeI_51997C `xml:"companyDetails,omitempty"` // minOccurs="0"

	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"` // minOccurs="0"

	FlightTypeDetails *ProductTypeDetailsTypeI `xml:"flightTypeDetails,omitempty"` // minOccurs="0"

	ItemNumber *int32 `xml:"itemNumber,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	SpecialSegment string `xml:"specialSegment,omitempty"` // minOccurs="0"

	MarriageDetails []*MarriageControlDetailsTypeI `xml:"marriageDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TravellerDetailsTypeI struct {
	// Specifies passenger given name and title.
	// xmlType: AlphaNumericString_Length1To70
	GivenName string `xml:"givenName"`
}

type TravellerInformationTypeI struct {
	PaxDetails *TravellerSurnameInformationTypeI `xml:"paxDetails"`

	OtherPaxDetails []*TravellerDetailsTypeI `xml:"otherPaxDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TravellerNameDetailsType struct {
	// xmlType: AlphaNumericString_Length1To5
	NameType string `xml:"nameType,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To1
	ReferenceName string `xml:"referenceName,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To1
	DisplayedName string `xml:"displayedName,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To4
	RomanizationMethod string `xml:"romanizationMethod,omitempty"` // minOccurs="0"

	// Passenger surname
	// xmlType: AlphaNumericString_Length1To70
	Surname string `xml:"surname"`

	// Passenger firstname
	// xmlType: AlphaNumericString_Length1To70
	GivenName string `xml:"givenName,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To70
	Title []string `xml:"title,omitempty"` // minOccurs="0" maxOccurs="2"
}

type TravellerNameInfoType struct {
	// PAX = PAX IN = Infant
	Qualifier formats.AMA_EDICodesetType_Length1to3 `xml:"qualifier,omitempty"` // minOccurs="0"

	Quantity *int32 `xml:"quantity,omitempty"` // minOccurs="0"

	// Passenger type (PTC).
	Type formats.AMA_EDICodesetType_Length1to3 `xml:"type,omitempty"` // minOccurs="0"

	// Passenger type (PTC).
	OtherType formats.AMA_EDICodesetType_Length1to3 `xml:"otherType,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To1
	InfantIndicator string `xml:"infantIndicator,omitempty"` // minOccurs="0"

	// Identification code, 2 cases: ID<1 to 51 char free text) or CR<1 to 40 char free text).
	// xmlType: AlphaNumericString_Length1To70
	TravellerIdentificationCode string `xml:"travellerIdentificationCode,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	Gender string `xml:"gender,omitempty"` // minOccurs="0"

	Age *int32 `xml:"age,omitempty"` // minOccurs="0"
}

type TravellerPriorityDetailsTypeI struct {
	// Staff airline of employment
	// xmlType: AlphaNumericString_Length1To35
	Company string `xml:"company,omitempty"` // minOccurs="0"

	// Staff date of joining (ddmmmyy)
	// xmlType: AlphaNumericString_Length1To35
	DateOfJoining string `xml:"dateOfJoining,omitempty"` // minOccurs="0"

	// Staff id number
	// xmlType: AlphaNumericString_Length1To10
	TravellerReference string `xml:"travellerReference,omitempty"` // minOccurs="0"
}

type TravellerSurnameInformationTypeI struct {
	// Passenger name.
	// xmlType: AlphaNumericString_Length1To70
	Surname string `xml:"surname"`

	// Specifies passenger type (adult, infant, military, etc.).
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"

	// Specify age of unaccompanied minor.
	Quantity *int32 `xml:"quantity,omitempty"` // minOccurs="0"
}
