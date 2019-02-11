package Fare_InformativePricingWithoutPNRReply_v12_4 // tipnrr124

//import "encoding/xml"

type FareInformativePricingWithoutPNRReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TIPNRR_12_4_1A Fare_InformativePricingWithoutPNRReply"`

	// Contains general information about the message, especially the use case. Tells if the request was correctly performed of not.
	MessageDetails *MessageActionDetailsTypeI `xml:"messageDetails"`

	ErrorGroup *ErrorGroupType `xml:"errorGroup,omitempty"` // minOccurs="0"

	MainGroup *MainGroup `xml:"mainGroup,omitempty"` // minOccurs="0"
}

type MainGroup struct {
	// A useless separator.
	DummySegment *DummySegmentTypeI `xml:"dummySegment"`

	// Convertion rates and currency information.
	ConvertionRate *ConversionRateTypeI `xml:"convertionRate,omitempty"` // minOccurs="0"

	GeneralIndicatorsGroup *GeneralIndicatorsGroup `xml:"generalIndicatorsGroup,omitempty"` // minOccurs="0"

	PricingGroupLevelGroup []*PricingGroupLevelGroup `xml:"pricingGroupLevelGroup"` // maxOccurs="99"
}

type GeneralIndicatorsGroup struct {
	// Contains pricing indicators such as SITI, SOTO... and International flag.
	GeneralIndicators *PricingTicketingDetailsTypeI `xml:"generalIndicators"`
}

type PricingGroupLevelGroup struct {
	// Number of pax in this fare group
	NumberOfPax *SegmentRepetitionControlTypeI `xml:"numberOfPax"`

	// * IDs of the passengers (same as in the request) * Carrier-related tattoo for LCC pricing (NOT IMPLEMENTED)
	PassengersID []*SpecificTravellerTypeI `xml:"passengersID,omitempty"` // minOccurs="0" maxOccurs="99"

	FareInfoGroup *FareInfoGroup `xml:"fareInfoGroup"`
}

type FareInfoGroup struct {
	// Fare information
	EmptySegment *FareInformationTypeI `xml:"emptySegment"`

	// Contains the following pieces of information: * Ticket designator * NVA date * NVB date
	PricingIndicators *PricingTicketingDetailsTypeI `xml:"pricingIndicators,omitempty"` // minOccurs="0"

	FareAmount *MonetaryInformationType_157205S `xml:"fareAmount,omitempty"` // minOccurs="0"

	// Used to store text data such as: * Horizontal fare calulation line * Mileage fare calculation line * Endorsement information ...
	TextData []*InteractiveFreeTextTypeI `xml:"textData,omitempty"` // minOccurs="0" maxOccurs="99"

	SurchargesGroup *SurchargesGroup `xml:"surchargesGroup,omitempty"` // minOccurs="0"

	CorporateGroup []*CorporateGroup `xml:"corporateGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	NegoFareGroup []*NegoFareGroup `xml:"negoFareGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	SegmentLevelGroup []*SegmentLevelGroup `xml:"segmentLevelGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	StructuredFareCalcGroup *StructuredFareCalcGroup `xml:"structuredFareCalcGroup,omitempty"` // minOccurs="0"

	CarrierFeeGroup []*CarrierFeeGroup `xml:"carrierFeeGroup,omitempty"` // minOccurs="0" maxOccurs="9"

	FareComponentDetailsGroup []*FareComponentDetailsType `xml:"fareComponentDetailsGroup,omitempty"` // minOccurs="0" maxOccurs="99"
}

type SurchargesGroup struct {
	// Stores the data related to taxes.
	TaxesAmount *TaxTypeI `xml:"taxesAmount"`

	// Stores the information related to the penalties: amount or rate, currency...
	PenaltiesAmount *DiscountAndPenaltyInformationTypeI `xml:"penaltiesAmount,omitempty"` // minOccurs="0"

	// To store the Passenger Facility Charges.
	PfcAmount []*MonetaryInformationTypeI `xml:"pfcAmount,omitempty"` // minOccurs="0" maxOccurs="5"
}

type CorporateGroup struct {
	// Stores data qualifying a corporate fare: * Type of fare (Nego/Unifare) * Associated contract number or company name
	CorporateData *FareCalculationCodeDetailsTypeI `xml:"corporateData"`
}

type NegoFareGroup struct {
	NegoFareIndicators *PricingTicketingSubsequentTypeI `xml:"negoFareIndicators"`

	ExtNegoFareIndicators *FareQualifierDetailsTypeI `xml:"extNegoFareIndicators,omitempty"` // minOccurs="0"

	// Amount of the negociated fares
	NegoFareAmount *DiscountAndPenaltyInformationTypeI `xml:"negoFareAmount,omitempty"` // minOccurs="0"

	// Text informations : indicates to ticketing what to print in Fare and Total boxes.
	NegoFareText *InteractiveFreeTextTypeI `xml:"negoFareText,omitempty"` // minOccurs="0"
}

type SegmentLevelGroup struct {
	// Information about a segment: dates, carrier, board/off point...
	SegmentInformation *TravelProductInformationTypeI_69238S `xml:"segmentInformation"`

	// Contains the following pieces of information: * Ticket designator * NVA date * NVB date
	AdditionalInformation *PricingTicketingDetailsTypeI `xml:"additionalInformation,omitempty"` // minOccurs="0"

	// Contains the following pieces of information: * Transportation class * Fare Basis * Fare by Rule flag
	FareBasis *FareQualifierDetailsTypeI `xml:"fareBasis,omitempty"` // minOccurs="0"

	CabinGroup []*CabinGroup `xml:"cabinGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	// Baggage allowance in a given measurement unit.
	BaggageAllowance *ExcessBaggageTypeI `xml:"baggageAllowance,omitempty"` // minOccurs="0"

	// PTC associated to this segment in this fare group.
	PtcSegment *NumberOfUnitsTypeI `xml:"ptcSegment,omitempty"` // minOccurs="0"

	// Coupon value/TPM for national fares
	CouponInformation *QuantityTypeI `xml:"couponInformation,omitempty"` // minOccurs="0"
}

type CabinGroup struct {
	// This segment gives the cabin information
	CabinSegment *ProductInformationTypeI `xml:"cabinSegment"`
}

type StructuredFareCalcGroup struct {
	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructureFareCalcRoot *FareComponentInformationTypeI `xml:"structureFareCalcRoot"`

	Group27 []*Group27 `xml:"group27,omitempty"` // minOccurs="0" maxOccurs="99"
}

type Group27 struct {
	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG27EQN *NumberOfUnitsTypeI `xml:"structuredFareCalcG27EQN"`

	Group28 []*Group28 `xml:"group28,omitempty"` // minOccurs="0" maxOccurs="99"

	// Dummy segment to differentiate MON in group 29 and 28
	DummySegmentGroup27 *DummySegmentTypeI `xml:"dummySegmentGroup27"`

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG27MON *MonetaryInformationTypeI `xml:"structuredFareCalcG27MON,omitempty"` // minOccurs="0"

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG27TXD []*TaxTypeI `xml:"structuredFareCalcG27TXD,omitempty"` // minOccurs="0" maxOccurs="99"

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG27CVR *ConversionRateTypeI `xml:"structuredFareCalcG27CVR,omitempty"` // minOccurs="0"
}

type Group28 struct {
	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG28ITM *ItemNumberTypeI `xml:"structuredFareCalcG28ITM"`

	Group29 []*Group29 `xml:"group29,omitempty"` // minOccurs="0" maxOccurs="99"

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG28MON *MonetaryInformationTypeI `xml:"structuredFareCalcG28MON,omitempty"` // minOccurs="0"

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG28PTS *PricingTicketingSubsequentTypeI `xml:"structuredFareCalcG28PTS,omitempty"` // minOccurs="0"

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG28FCC *FareCalculationCodeDetailsTypeI `xml:"structuredFareCalcG28FCC,omitempty"` // minOccurs="0"

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG28PTK *PricingTicketingDetailsTypeI `xml:"structuredFareCalcG28PTK,omitempty"` // minOccurs="0"

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG28FRU *FareRulesInformationTypeI `xml:"structuredFareCalcG28FRU,omitempty"` // minOccurs="0"
}

type Group29 struct {
	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG28ADT *ActionDetailsTypeI `xml:"structuredFareCalcG28ADT"`

	// Reserved for a future use to store a structured fare calculation line. NOT IMPLEMENTED.
	StructuredFareCalcG28TVL *TravelProductInformationTypeI_69238S `xml:"structuredFareCalcG28TVL,omitempty"` // minOccurs="0"
}

type CarrierFeeGroup struct {
	// Nature of the fee (OB, OC, ...)
	FeeType *SelectionDetailsTypeI `xml:"feeType"`

	FeeDetails []*FeeDetails `xml:"feeDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type FeeDetails struct {
	// Fee information
	FeeInfo *SpecificDataInformationTypeI `xml:"feeInfo"`

	// Fee associated amounts: amount with/without tax, total tax amount
	FeeAmounts *MonetaryInformationTypeI `xml:"feeAmounts,omitempty"` // minOccurs="0"

	// taxes related to this fee
	FeeTaxes []*TaxTypeI `xml:"feeTaxes,omitempty"` // minOccurs="0" maxOccurs="99"

	// Attributes of this fee (commercial description, ...)
	FeeDescription *InteractiveFreeTextTypeI `xml:"feeDescription,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type ActionDetailsTypeI struct {
	NumberOfItemsDetails *ProcessingInformationTypeI `xml:"numberOfItemsDetails,omitempty"` // minOccurs="0"

	LastItemsDetails []*ReferenceTypeI `xml:"lastItemsDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type AdditionalFareQualifierDetailsTypeI struct {
	// Format limitations: an..35
	RateClass string `xml:"rateClass,omitempty"` // minOccurs="0"

	// Format limitations: an..18
	CommodityCategory string `xml:"commodityCategory,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	PricingGroup string `xml:"pricingGroup,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	SecondRateClass []string `xml:"secondRateClass,omitempty"` // minOccurs="0" maxOccurs="29"
}

type ApplicationErrorDetailType struct {
	// Code identifying the data validation error condition.
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	ErrorCategory string `xml:"errorCategory,omitempty"` // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"` // minOccurs="0"
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type BaggageDetailsTypeI struct {
	// Format limitations: n..15
	FreeAllowance *int32 `xml:"freeAllowance,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	Measurement *int32 `xml:"measurement,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	QuantityCode string `xml:"quantityCode,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	UnitQualifier string `xml:"unitQualifier,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	ProcessIndicator string `xml:"processIndicator,omitempty"` // minOccurs="0"
}

type BagtagDetailsTypeI struct {
	// Format limitations: an..35
	Company string `xml:"company,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	Identifier string `xml:"identifier,omitempty"` // minOccurs="0"

	// Format limitations: n..15
	Number *int32 `xml:"number,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	Location string `xml:"location,omitempty"` // minOccurs="0"

	// Format limitations: an..15
	CompanyNumber string `xml:"companyNumber,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Indicator string `xml:"indicator,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Characteristic string `xml:"characteristic,omitempty"` // minOccurs="0"

	// Format limitations: an..4
	SpecialRequirement string `xml:"specialRequirement,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	Measurement *int32 `xml:"measurement,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	UnitQualifier string `xml:"unitQualifier,omitempty"` // minOccurs="0"

	// Format limitations: an..70
	Description string `xml:"description,omitempty"` // minOccurs="0"
}

type CompanyIdentificationNumbersTypeI struct {
	// Format limitations: an..15
	Identifier string `xml:"identifier"`

	// Format limitations: an..15
	OtherIdentifier string `xml:"otherIdentifier,omitempty"` // minOccurs="0"
}

type CompanyIdentificationTypeI struct {
	// Format limitations: an..35
	MarketingCompany string `xml:"marketingCompany,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	OperatingCompany string `xml:"operatingCompany,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	OtherCompany string `xml:"otherCompany,omitempty"` // minOccurs="0"
}

type ConversionRateDetailsTypeI struct {
	// Format limitations: an..3
	ConversionType string `xml:"conversionType,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	RateType string `xml:"rateType,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	PricingAmount *float64 `xml:"pricingAmount,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	ConvertedValueAmount *int32 `xml:"convertedValueAmount,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	DutyTaxFeeType string `xml:"dutyTaxFeeType,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	MeasurementValue *int32 `xml:"measurementValue,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	MeasurementSignificance string `xml:"measurementSignificance,omitempty"` // minOccurs="0"
}

type ConversionRateTypeI struct {
	ConversionRateDetails *ConversionRateDetailsTypeI `xml:"conversionRateDetails"`

	OtherConvRateDetails []*ConversionRateDetailsTypeI `xml:"otherConvRateDetails,omitempty"` // minOccurs="0" maxOccurs="19"
}

type CouponDetailsType struct {
	// Tattoo + type of the product identifying the coupon.
	ProductId *ReferenceInfoType `xml:"productId"`

	// Flight Connection Type
	FlightConnectionType *TravelProductInformationType `xml:"flightConnectionType,omitempty"` // minOccurs="0"
}

type DataInformationTypeI struct {
	// Format limitations: an..3
	Indicator string `xml:"indicator,omitempty"` // minOccurs="0"

	// Format limitations: n..15
	Value *int32 `xml:"value,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Unit string `xml:"unit,omitempty"` // minOccurs="0"
}

type DataTypeInformationTypeI struct {
	// Format limitations: an..3
	Type string `xml:"type"`

	// Format limitations: an..3
	StatusEvent string `xml:"statusEvent,omitempty"` // minOccurs="0"
}

type DiscountAndPenaltyInformationTypeI struct {
	// Format limitations: an..3
	DiscountPenaltyQualifier string `xml:"discountPenaltyQualifier,omitempty"` // minOccurs="0"

	DiscountPenaltyDetails []*DiscountPenaltyMonetaryInformationTypeI `xml:"discountPenaltyDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type DiscountPenaltyInformationType struct {
	// Used for codes in the AMADEUS code tables. Code Length is three alphanumeric characters.
	FareQualifier string `xml:"fareQualifier,omitempty"` // minOccurs="0"
}

type DiscountPenaltyInformationTypeI struct {
	// Format limitations: an..3
	FareQualifier string `xml:"fareQualifier"`

	// Format limitations: an..35
	RateCategory string `xml:"rateCategory,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	Amount *float64 `xml:"amount,omitempty"` // minOccurs="0"

	// Format limitations: n..8
	Percentage *int32 `xml:"percentage,omitempty"` // minOccurs="0"
}

type DiscountPenaltyMonetaryInformationTypeI struct {
	// Format limitations: an..3
	Function string `xml:"function,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	AmountType string `xml:"amountType,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	Amount *float64 `xml:"amount,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	Rate string `xml:"rate,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"
}

type DummySegmentTypeI struct {
}

type ErrorGroupType struct {
	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType `xml:"errorOrWarningCodeDetails"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"` // minOccurs="0"
}

type ExcessBaggageDetailsTypeI struct {
	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	Amount *float64 `xml:"amount,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	ProcessIndicator string `xml:"processIndicator,omitempty"` // minOccurs="0"
}

type ExcessBaggageTypeI struct {
	ExcessBaggageDetails *ExcessBaggageDetailsTypeI `xml:"excessBaggageDetails,omitempty"` // minOccurs="0"

	BaggageDetails *BaggageDetailsTypeI `xml:"baggageDetails,omitempty"` // minOccurs="0"

	OtherBaggageDetails *BaggageDetailsTypeI `xml:"otherBaggageDetails,omitempty"` // minOccurs="0"

	ExtraBaggageDetails *BaggageDetailsTypeI `xml:"extraBaggageDetails,omitempty"` // minOccurs="0"

	BagTagDetails []*BagtagDetailsTypeI `xml:"bagTagDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type FareCalculationCodeDetailsTypeI struct {
	// Format limitations: an..3
	ChargeCategory string `xml:"chargeCategory,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	Amount *float64 `xml:"amount,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	LocationCode string `xml:"locationCode,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	OtherLocationCode string `xml:"otherLocationCode,omitempty"` // minOccurs="0"

	// Format limitations: n..8
	Rate *float64 `xml:"rate,omitempty"` // minOccurs="0"
}

type FareCategoryCodesTypeI struct {
	// Format limitations: an..20
	FareType string `xml:"fareType"`

	// Format limitations: an..20
	OtherFareType []string `xml:"otherFareType,omitempty"` // minOccurs="0" maxOccurs="8"
}

type FareComponentDetailsType struct {
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

	CouponDetailsGroup []*CouponDetailsType `xml:"couponDetailsGroup"` // maxOccurs="99"
}

type FareComponentDetailsTypeI struct {
	// Type of data _ fare calc or exchanged residual fare data
	DataType string `xml:"dataType,omitempty"` // minOccurs="0"

	// Fare component count
	Count *int32 `xml:"count,omitempty"` // minOccurs="0"

	// Price quote date
	PricingDate string `xml:"pricingDate,omitempty"` // minOccurs="0"

	// Account code
	AccountCode string `xml:"accountCode,omitempty"` // minOccurs="0"

	// Input designator
	InputDesignator string `xml:"inputDesignator,omitempty"` // minOccurs="0"
}

type FareComponentInformationTypeI struct {
	FareComponentDetails *FareComponentDetailsTypeI `xml:"fareComponentDetails,omitempty"` // minOccurs="0"

	// Ticket document number
	TicketNumber string `xml:"ticketNumber,omitempty"` // minOccurs="0"
}

type FareDetailsTypeI struct {
	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"` // minOccurs="0"

	// Format limitations: n..8
	Rate *float64 `xml:"rate,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Country string `xml:"country,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	FareCategory string `xml:"fareCategory,omitempty"` // minOccurs="0"
}

type FareInformationTypeI struct {
	// Not used
	ValueQualifier string `xml:"valueQualifier,omitempty"` // minOccurs="0"

	// Not used
	Value *int32 `xml:"value,omitempty"` // minOccurs="0"

	// Fare information
	FareDetails *FareDetailsTypeI `xml:"fareDetails,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	IdentityNumber string `xml:"identityNumber,omitempty"` // minOccurs="0"

	// Not used
	FareTypeGrouping *FareTypeGroupingInformationTypeI `xml:"fareTypeGrouping,omitempty"` // minOccurs="0"

	// Not used
	RateCategory []string `xml:"rateCategory,omitempty"` // minOccurs="0" maxOccurs="9"
}

type FareQualifierDetailsType struct {
	DiscountDetails []*DiscountPenaltyInformationType `xml:"discountDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type FareQualifierDetailsTypeI struct {
	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"` // minOccurs="0"

	FareCategories *FareCategoryCodesTypeI `xml:"fareCategories,omitempty"` // minOccurs="0"

	FareDetails *FareDetailsTypeI `xml:"fareDetails,omitempty"` // minOccurs="0"

	AdditionalFareDetails *AdditionalFareQualifierDetailsTypeI `xml:"additionalFareDetails,omitempty"` // minOccurs="0"

	DiscountDetails []*DiscountPenaltyInformationTypeI `xml:"discountDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type FareRulesInformationTypeI struct {
	// Format limitations: an..9
	TariffClassId string `xml:"tariffClassId,omitempty"` // minOccurs="0"

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"` // minOccurs="0"

	// Format limitations: an..7
	RuleSectionId []string `xml:"ruleSectionId,omitempty"` // minOccurs="0" maxOccurs="99"
}

type FareTypeGroupingInformationTypeI struct {
	// Format limitations: an..35
	PricingGroup []string `xml:"pricingGroup,omitempty"` // minOccurs="0" maxOccurs="5"
}

type FreeTextDetailsType struct {
	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..4
	InformationType string `xml:"informationType,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Status string `xml:"status,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	CompanyId string `xml:"companyId,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Language string `xml:"language,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Source string `xml:"source"`

	// Format limitations: an..3
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText []string `xml:"freeText"` // maxOccurs="99"
}

type FreeTextQualificationTypeI struct {
	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..4
	InformationType string `xml:"informationType,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Status string `xml:"status,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	CompanyId string `xml:"companyId,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Language string `xml:"language,omitempty"` // minOccurs="0"
}

type InteractiveFreeTextTypeI struct {
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"` // minOccurs="0"

	// Format limitations: an..250
	FreeText []string `xml:"freeText,omitempty"` // minOccurs="0" maxOccurs="99"
}

type ItemNumberIdentificationType struct {
	// Format limitations: an..35
	Number string `xml:"number,omitempty"` // minOccurs="0"
}

type ItemNumberIdentificationTypeI struct {
	// Format limitations: an..35
	Number string `xml:"number,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Type string `xml:"type,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	ResponsibleAgency string `xml:"responsibleAgency,omitempty"` // minOccurs="0"
}

type ItemNumberType struct {
	ItemNumberDetails []*ItemNumberIdentificationType `xml:"itemNumberDetails"` // maxOccurs="99"
}

type ItemNumberTypeI struct {
	ItemNumberDetails []*ItemNumberIdentificationTypeI `xml:"itemNumberDetails"` // maxOccurs="99"
}

type LocationDetailsTypeI struct {
	// Format limitations: an..25
	City string `xml:"city,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Country string `xml:"country,omitempty"` // minOccurs="0"
}

type LocationTypeI struct {
	// Format limitations: an..25
	TrueLocationId string `xml:"trueLocationId,omitempty"` // minOccurs="0"

	// Format limitations: an..17
	TrueLocation string `xml:"trueLocation,omitempty"` // minOccurs="0"
}

type LocationTypeI_208252C struct {
	// Format limitations: an..25
	TrueLocationId string `xml:"trueLocationId,omitempty"` // minOccurs="0"
}

type MarriageControlDetailsTypeI struct {
	// Format limitations: an..3
	Relation string `xml:"relation,omitempty"` // minOccurs="0"

	// Format limitations: n..10
	MarriageIdentifier *int32 `xml:"marriageIdentifier,omitempty"` // minOccurs="0"

	// Format limitations: n..6
	LineNumber *int32 `xml:"lineNumber,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	OtherRelation string `xml:"otherRelation,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	CarrierCode string `xml:"carrierCode,omitempty"` // minOccurs="0"
}

type MessageActionDetailsTypeI struct {
	MessageFunctionDetails *MessageFunctionBusinessDetailsTypeI `xml:"messageFunctionDetails,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	ResponseType string `xml:"responseType,omitempty"` // minOccurs="0"
}

type MessageFunctionBusinessDetailsTypeI struct {
	// Format limitations: an..3
	BusinessFunction string `xml:"businessFunction,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	MessageFunction string `xml:"messageFunction,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	ResponsibleAgency string `xml:"responsibleAgency,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"` // minOccurs="0" maxOccurs="20"
}

type MonetaryInformationDetailsType struct {
	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	Amount string `xml:"amount,omitempty"` // minOccurs="0"

	// Currency
	Currency string `xml:"currency,omitempty"` // minOccurs="0"
}

type MonetaryInformationDetailsTypeI struct {
	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Format limitations: an..35
	Amount string `xml:"amount,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"

	// Format limitations: an..25
	Location string `xml:"location,omitempty"` // minOccurs="0"
}

type MonetaryInformationDetailsType_223848C struct {
	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	Amount string `xml:"amount,omitempty"` // minOccurs="0"

	// Currency
	Currency string `xml:"currency,omitempty"` // minOccurs="0"

	// location
	Location string `xml:"location,omitempty"` // minOccurs="0"
}

type MonetaryInformationType struct {
	// Monetary information per fare component
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`

	// Other monetary information per fare component
	OtherMonetaryDetails []*MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"` // minOccurs="0" maxOccurs="19"
}

type MonetaryInformationTypeI struct {
	MonetaryDetails *MonetaryInformationDetailsTypeI `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"otherMonetaryDetails,omitempty"` // minOccurs="0" maxOccurs="19"
}

type MonetaryInformationType_157205S struct {
	MonetaryDetails *MonetaryInformationDetailsType_223848C `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsType_223848C `xml:"otherMonetaryDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type NumberOfUnitDetailsTypeI struct {
	// Format limitations: n..15
	NumberOfUnit *int32 `xml:"numberOfUnit,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	UnitQualifier string `xml:"unitQualifier,omitempty"` // minOccurs="0"
}

type NumberOfUnitsTypeI struct {
	QuantityDetails *NumberOfUnitDetailsTypeI `xml:"quantityDetails"`

	OtherQuantityDetails []*NumberOfUnitDetailsTypeI `xml:"otherQuantityDetails,omitempty"` // minOccurs="0" maxOccurs="8"
}

type PricingOrTicketingSubsequentType struct {
	// RATE OR TARIFF CLASS INFORMATION
	FareBasisDetails *RateTariffClassInformationType `xml:"fareBasisDetails,omitempty"` // minOccurs="0"
}

type PricingTicketingDetailsTypeI struct {
	PriceTicketDetails *PricingTicketingInformationTypeI `xml:"priceTicketDetails,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	PriceTariffType string `xml:"priceTariffType,omitempty"` // minOccurs="0"

	ProductDateTimeDetails *ProductDateTimeTypeI `xml:"productDateTimeDetails,omitempty"` // minOccurs="0"

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"` // minOccurs="0"

	CompanyNumberDetails *CompanyIdentificationNumbersTypeI `xml:"companyNumberDetails,omitempty"` // minOccurs="0"

	LocationDetails *LocationDetailsTypeI `xml:"locationDetails,omitempty"` // minOccurs="0"

	OtherLocationDetails *LocationDetailsTypeI `xml:"otherLocationDetails,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	IdNumber string `xml:"idNumber,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	MonetaryAmount *float64 `xml:"monetaryAmount,omitempty"` // minOccurs="0"
}

type PricingTicketingInformationTypeI struct {
	// Format limitations: an..3
	Indicators []string `xml:"indicators,omitempty"` // minOccurs="0" maxOccurs="20"
}

type PricingTicketingSubsequentTypeI struct {
	// Format limitations: an..35
	ItemNumber string `xml:"itemNumber,omitempty"` // minOccurs="0"

	FareBasisDetails *RateTariffClassInformationTypeI `xml:"fareBasisDetails,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	FareValue *int32 `xml:"fareValue,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	PriceType string `xml:"priceType,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	SpecialCondition string `xml:"specialCondition,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	OtherSpecialCondition string `xml:"otherSpecialCondition,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	AdditionalSpecialCondition string `xml:"additionalSpecialCondition,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	TaxCategory []string `xml:"taxCategory,omitempty"` // minOccurs="0" maxOccurs="2"
}

type ProcessingInformationTypeI struct {
	// Format limitations: an..3
	ActionQualifier string `xml:"actionQualifier,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	ReferenceQualifier string `xml:"referenceQualifier,omitempty"` // minOccurs="0"

	// Format limitations: an..6
	NumberOfItems string `xml:"numberOfItems,omitempty"` // minOccurs="0"
}

type ProductDateTimeTypeI struct {
	// Format limitations: an..35
	DepartureDate string `xml:"departureDate,omitempty"` // minOccurs="0"

	// Format limitations: n..4
	DepartureTime *int32 `xml:"departureTime,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	ArrivalDate string `xml:"arrivalDate,omitempty"` // minOccurs="0"

	// Format limitations: n..4
	ArrivalTime *int32 `xml:"arrivalTime,omitempty"` // minOccurs="0"

	// Format limitations: n1
	DateVariation *int32 `xml:"dateVariation,omitempty"` // minOccurs="0"
}

type ProductDetailsTypeI struct {
	// Format limitations: an..17
	Designator string `xml:"designator"`

	// Format limitations: an..3
	AvailabilityStatus string `xml:"availabilityStatus,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	SpecialService string `xml:"specialService,omitempty"` // minOccurs="0"

	// Format limitations: an..7
	Option []string `xml:"option,omitempty"` // minOccurs="0" maxOccurs="3"
}

type ProductIdentificationDetailsTypeI struct {
	// Format limitations: an..35
	FlightNumber string `xml:"flightNumber"`

	// Format limitations: an..17
	BookingClass string `xml:"bookingClass,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	OperationalSuffix string `xml:"operationalSuffix,omitempty"` // minOccurs="0"

	// Format limitations: an..7
	Modifier []string `xml:"modifier,omitempty"` // minOccurs="0" maxOccurs="3"
}

type ProductInformationTypeI struct {
	// Format limitations: an..3
	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"` // minOccurs="0"

	BookingClassDetails []*ProductDetailsTypeI `xml:"bookingClassDetails,omitempty"` // minOccurs="0" maxOccurs="26"
}

type ProductTypeDetailsType struct {
	// TST Connection Type
	FlightIndicator string `xml:"flightIndicator"`
}

type ProductTypeDetailsTypeI struct {
	// Format limitations: an..6
	FlightIndicator []string `xml:"flightIndicator"` // maxOccurs="9"
}

type QuantityDetailsTypeI struct {
	// Format limitations: an..3
	Qualifier string `xml:"qualifier"`

	// Format limitations: n..15
	Value int32 `xml:"value"`

	// Format limitations: an..3
	Unit string `xml:"unit,omitempty"` // minOccurs="0"
}

type QuantityTypeI struct {
	QuantityDetails *QuantityDetailsTypeI `xml:"quantityDetails"`

	OtherquantityDetails []*QuantityDetailsTypeI `xml:"otherquantityDetails,omitempty"` // minOccurs="0" maxOccurs="8"
}

type RateTariffClassInformationType struct {
	// Fare Basis Code
	RateTariffClass string `xml:"rateTariffClass,omitempty"` // minOccurs="0"

	// Ticket Designator
	OtherRateTariffClass string `xml:"otherRateTariffClass,omitempty"` // minOccurs="0"
}

type RateTariffClassInformationTypeI struct {
	// Format limitations: an..35
	RateTariffClass string `xml:"rateTariffClass,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	RateTariffIndicator string `xml:"rateTariffIndicator,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	OtherRateTariffClass string `xml:"otherRateTariffClass,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	OtherRateTariffIndicator string `xml:"otherRateTariffIndicator,omitempty"` // minOccurs="0"
}

type ReferenceInfoType struct {
	ReferenceDetails *ReferencingDetailsType `xml:"referenceDetails"`
}

type ReferenceTypeI struct {
	// Format limitations: an..6
	NumberOfItems string `xml:"numberOfItems,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	LastItemIdentifier []string `xml:"lastItemIdentifier,omitempty"` // minOccurs="0" maxOccurs="99"
}

type ReferencingDetailsType struct {
	// Format limitations: an..10
	Type string `xml:"type"`

	// Format limitations: an..60
	Value string `xml:"value"`
}

type SegmentRepetitionControlDetailsTypeI struct {
	// Format limitations: n..15
	Quantity *int32 `xml:"quantity,omitempty"` // minOccurs="0"

	// Format limitations: n..15
	NumberOfUnits *int32 `xml:"numberOfUnits,omitempty"` // minOccurs="0"

	// Format limitations: n..15
	TotalNumberOfItems *int32 `xml:"totalNumberOfItems,omitempty"` // minOccurs="0"
}

type SegmentRepetitionControlTypeI struct {
	SegmentControlDetails []*SegmentRepetitionControlDetailsTypeI `xml:"segmentControlDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type SelectionDetailsInformationTypeI struct {
	// Format limitations: an..3
	Option string `xml:"option"`

	// Format limitations: an..35
	OptionInformation string `xml:"optionInformation,omitempty"` // minOccurs="0"
}

type SelectionDetailsTypeI struct {
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails"`

	OtherSelectionDetails []*SelectionDetailsInformationTypeI `xml:"otherSelectionDetails,omitempty"` // minOccurs="0" maxOccurs="98"
}

type SpecificDataInformationTypeI struct {
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`

	DataInformation []*DataInformationTypeI `xml:"dataInformation,omitempty"` // minOccurs="0" maxOccurs="99"
}

type SpecificTravellerDetailsTypeI struct {
	// Format limitations: an..10
	ReferenceNumber string `xml:"referenceNumber,omitempty"` // minOccurs="0"

	// Format limitations: n..18
	MeasurementValue *int32 `xml:"measurementValue,omitempty"` // minOccurs="0"

	// Format limitations: an..35
	FirstDate string `xml:"firstDate,omitempty"` // minOccurs="0"

	// Format limitations: an..70
	Surname string `xml:"surname,omitempty"` // minOccurs="0"

	// Format limitations: an..70
	FirstName string `xml:"firstName,omitempty"` // minOccurs="0"
}

type SpecificTravellerTypeI struct {
	TravellerDetails []*SpecificTravellerDetailsTypeI `xml:"travellerDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TaxDetailsTypeI struct {
	// Format limitations: an..17
	Rate string `xml:"rate,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	CountryCode string `xml:"countryCode,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	CurrencyCode string `xml:"currencyCode,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	Type []string `xml:"type,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TaxTypeI struct {
	// Format limitations: an..3
	TaxCategory string `xml:"taxCategory,omitempty"` // minOccurs="0"

	TaxDetails []*TaxDetailsTypeI `xml:"taxDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TravelProductInformationType struct {
	BoardPointDetails *LocationTypeI_208252C `xml:"boardPointDetails,omitempty"` // minOccurs="0"

	OffpointDetails *LocationTypeI_208252C `xml:"offpointDetails,omitempty"` // minOccurs="0"

	// TST Connection Type
	FlightTypeDetails *ProductTypeDetailsType `xml:"flightTypeDetails,omitempty"` // minOccurs="0"
}

type TravelProductInformationTypeI struct {
	BoardPointDetails *LocationTypeI_208252C `xml:"boardPointDetails,omitempty"` // minOccurs="0"

	OffpointDetails *LocationTypeI_208252C `xml:"offpointDetails,omitempty"` // minOccurs="0"
}

type TravelProductInformationTypeI_69238S struct {
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"` // minOccurs="0"

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"` // minOccurs="0"

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"` // minOccurs="0"

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"` // minOccurs="0"

	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"` // minOccurs="0"

	FlightTypeDetails *ProductTypeDetailsTypeI `xml:"flightTypeDetails,omitempty"` // minOccurs="0"

	// Format limitations: n..6
	ItemNumber *int32 `xml:"itemNumber,omitempty"` // minOccurs="0"

	// Format limitations: an..3
	SpecialSegment string `xml:"specialSegment,omitempty"` // minOccurs="0"

	MarriageDetails []*MarriageControlDetailsTypeI `xml:"marriageDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}
