package Fare_InformativePricingWithoutPNR_v12_4 // tipnrq124

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/sdk/formats"
)

type FareInformativePricingWithoutPNR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TIPNRQ_12_4_1A Fare_InformativePricingWithoutPNR"`

	// Contains general information about the message, especially the use case.
	MessageDetails *MessageActionDetailsTypeI `xml:"messageDetails"`

	OriginatorGroup *OriginatorGroup `xml:"originatorGroup,omitempty"` // minOccurs="0"

	// Handles: * Fare conversion option * Fare selection option
	CurrencyOverride *ConversionRateTypeI `xml:"currencyOverride,omitempty"` // minOccurs="0"

	// Used for Unifare pricing
	CorporateFares *CorporateFareInformationType `xml:"corporateFares,omitempty"` // minOccurs="0"

	TaxExemptGroup *TaxExemptGroup `xml:"taxExemptGroup,omitempty"` // minOccurs="0"

	// Form of payment at query level
	GeneralFormOfPayment *FormOfPaymentTypeI `xml:"generalFormOfPayment,omitempty"` // minOccurs="0"

	PassengersGroup []*PassengersGroup `xml:"passengersGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	PricingOptionsGroup []*PricingOptionsGroup `xml:"pricingOptionsGroup,omitempty"` // minOccurs="0" maxOccurs="99"

	// This option allows to target the wished amount (called Instant pricing option)
	InstantPricingOption *MonetaryInformationType `xml:"instantPricingOption,omitempty"` // minOccurs="0"

	TripsGroup *TripsGroup `xml:"tripsGroup"`

	ObFeeRequestGroup *ObFeeRequestGroup `xml:"obFeeRequestGroup,omitempty"` // minOccurs="0"
}

type OriginatorGroup struct {
	// Supplementary protocol or business related information.
	AdditionalBusinessInformation *AdditionalBusinessSourceInformationTypeI `xml:"additionalBusinessInformation,omitempty"` // minOccurs="0"
}

type TaxExemptGroup struct {
	// Handles tax exemption option.
	TaxExempt *TaxTypeI `xml:"taxExempt"`
}

type PassengersGroup struct {
	// Contains: * Number of passengers in the group * Group tattoo
	SegmentRepetitionControl *SegmentRepetitionControlTypeI_69179S `xml:"segmentRepetitionControl"`

	// Passengers' tattoos provided by the carrier in case of LCC pricing with U2. NOT USED FOR FSC.
	TravellersID *SpecificTravellerTypeI `xml:"travellersID,omitempty"` // minOccurs="0"

	PtcGroup []*PtcGroup `xml:"ptcGroup,omitempty"` // minOccurs="0" maxOccurs="3"
}

type PtcGroup struct {
	// PTC/Discount Code
	DiscountPtc *FareInformationTypeI `xml:"discountPtc"`

	// Form of payment at passenger level
	PassengerFormOfPayment *FormOfPaymentTypeI `xml:"passengerFormOfPayment,omitempty"` // minOccurs="0"
}

type PricingOptionsGroup struct {
	// Handles the following pricing options: * Point of Sell override * Point of Ticketing override * Fare list * P/E ticket * Price by fare basis indicator * Validating Carrier * Award Publishing Carrier
	PricingDetails *PricingTicketingDetailsType `xml:"pricingDetails"`

	// * Expanded parameters * Forced PTC flag
	ExtPricingDetails *FareQualifierDetailsTypeI `xml:"extPricingDetails,omitempty"` // minOccurs="0"
}

type TripsGroup struct {
	// This segment can be left empty.
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination"`

	SegmentGroup []*SegmentGroup `xml:"segmentGroup"` // maxOccurs="99"
}

type SegmentGroup struct {
	// Convey the information related to a segment (company, flight number, origin, destination...).
	SegmentInformation *TravelProductInformationTypeI `xml:"segmentInformation"`

	// Used for technical stops, even if it is currently deprecated.
	AdditionnalSegmentDetails *AdditionalProductDetailsTypeI `xml:"additionnalSegmentDetails,omitempty"` // minOccurs="0"

	// Contain Zap Off details
	ZapOffDetails *FareQualifierDetailsTypeI `xml:"zapOffDetails,omitempty"` // minOccurs="0"

	InventoryGroup *InventoryGroup `xml:"inventoryGroup,omitempty"` // minOccurs="0"

	PricePsgByFareBasisInfo []*PricePsgByFareBasisInfo `xml:"pricePsgByFareBasisInfo,omitempty"` // minOccurs="0" maxOccurs="99"

	Trigger *DummySegmentTypeI `xml:"trigger"`

	// Options at passenger level: PTC/Discount Tier Level
	PsgDetailsGroup []*GroupPassengerDetailsType `xml:"psgDetailsGroup,omitempty"` // minOccurs="0" maxOccurs="99"
}

type InventoryGroup struct {
	// To store the flight inventory (opened classes and number of remaining seats) if known.
	Inventory *ProductInformationTypeI `xml:"inventory"`
}

type PricePsgByFareBasisInfo struct {
	// Segment repetition control
	SegmentRepetitionControl *SegmentRepetitionControlType `xml:"segmentRepetitionControl"`

	// Specific traveller details
	SpecificTravellerDetails *SpecificTravellerTypeI_65902S `xml:"specificTravellerDetails,omitempty"` // minOccurs="0"

	RequestedPricingInfo *RequestedPricingInfo `xml:"requestedPricingInfo,omitempty"` // minOccurs="0"
}

type RequestedPricingInfo struct {
	// This group is used in order to identify Pax/Farebasis group
	FareInfo *FareInformationType `xml:"fareInfo"`

	// Date and time information  (for future use)
	DateAndTimeInfo *DateAndTimeInformationTypeI `xml:"dateAndTimeInfo,omitempty"` // minOccurs="0"

	// Pricing by fare basis information
	PricingByFareBasisInfo *FareQualifierDetailsTypeI_82845S `xml:"pricingByFareBasisInfo,omitempty"` // minOccurs="0"

	// For future use!
	RoutingNumberInfo *FareRouteInformationTypeI `xml:"routingNumberInfo,omitempty"` // minOccurs="0"
}

type ObFeeRequestGroup struct {
	// Marker fee options
	MarkersFeeOptions *DummySegmentTypeI `xml:"markersFeeOptions"`

	FeeOptionInfoGroup []*FeeOptionInfoGroup `xml:"feeOptionInfoGroup"` // maxOccurs="5"
}

type FeeOptionInfoGroup struct {
	// Nature of the fee ( OB, OC, ...)
	FeeTypeInfo *SelectionDetailsType `xml:"feeTypeInfo"`

	// Associated tax rate
	RateTaxInfo *MonetaryInformationType_69171S `xml:"rateTaxInfo,omitempty"` // minOccurs="0"

	FeeDetailsInfoGroup []*FeeDetailsInfoGroup `xml:"feeDetailsInfoGroup,omitempty"` // minOccurs="0" maxOccurs="99"
}

type FeeDetailsInfoGroup struct {
	// Fee information
	FeeInfo *SpecificDataInformationTypeI `xml:"feeInfo"`

	// Include, exclude
	FeeProcessingInfo *SelectionDetailsType `xml:"feeProcessingInfo"`

	// Associated amounts: amount to take into account to calculate the fee, ...
	AssociatedAmountsInfo *MonetaryInformationType_69173S `xml:"associatedAmountsInfo,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type AdditionalBusinessSourceInformationTypeI struct {
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails,omitempty"` // minOccurs="0"

	LocationDetails *LocationTypeI `xml:"locationDetails,omitempty"` // minOccurs="0"

	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"` // minOccurs="0"

	SystemCode formats.AlphaNumericString_Length1To35 `xml:"systemCode,omitempty"` // minOccurs="0"
}

type AdditionalFareQualifierDetailsTypeI struct {
	RateClass formats.AlphaNumericString_Length1To35 `xml:"rateClass,omitempty"` // minOccurs="0"

	CommodityCategory formats.AlphaNumericString_Length1To18 `xml:"commodityCategory,omitempty"` // minOccurs="0"

	PricingGroup formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"` // minOccurs="0"

	SecondRateClass []formats.AlphaNumericString_Length1To35 `xml:"secondRateClass,omitempty"` // minOccurs="0" maxOccurs="29"
}

type AdditionalProductDetailsTypeI struct {
	LegDetails *AdditionalProductTypeI `xml:"legDetails,omitempty"` // minOccurs="0"

	// not used
	DepartureStationInfo *StationInformationTypeI `xml:"departureStationInfo,omitempty"` // minOccurs="0"

	// not used
	ArrivalStationInfo *StationInformationTypeI `xml:"arrivalStationInfo,omitempty"` // minOccurs="0"

	// not used
	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"` // minOccurs="0"

	// not used
	TravellerTimeDetails *TravellerTimeDetailsTypeI `xml:"travellerTimeDetails,omitempty"` // minOccurs="0"

	// not used
	FacilitiesInformation []*ProductFacilitiesTypeI `xml:"facilitiesInformation,omitempty"` // minOccurs="0" maxOccurs="10"
}

type AdditionalProductTypeI struct {
	Equipment formats.AlphaNumericString_Length1To8 `xml:"equipment,omitempty"` // minOccurs="0"

	NumberOfStops *formats.NumericInteger_Length1To3 `xml:"numberOfStops,omitempty"` // minOccurs="0"

	// not used
	Duration *formats.NumericInteger_Length1To6 `xml:"duration,omitempty"` // minOccurs="0"

	// not used
	Percentage *formats.NumericInteger_Length1To8 `xml:"percentage,omitempty"` // minOccurs="0"

	// not used
	DaysOfOperation formats.AlphaNumericString_Length1To7 `xml:"daysOfOperation,omitempty"` // minOccurs="0"

	// not used
	DateTimePeriod formats.AlphaNumericString_Length1To35 `xml:"dateTimePeriod,omitempty"` // minOccurs="0"

	// not used
	ComplexingFlightIndicator formats.AlphaNumericString_Length1To1 `xml:"complexingFlightIndicator,omitempty"` // minOccurs="0"

	// not used
	Locations []formats.AlphaNumericString_Length1To25 `xml:"locations,omitempty"` // minOccurs="0" maxOccurs="3"
}

type CompanyIdentificationNumbersTypeI struct {
	// Controlling Carrier
	Identifier formats.AlphaNumericString_Length1To15 `xml:"identifier"`

	OtherIdentifier formats.AlphaNumericString_Length1To15 `xml:"otherIdentifier,omitempty"` // minOccurs="0"
}

type CompanyIdentificationTypeI struct {
	MarketingCompany formats.AlphaNumericString_Length1To35 `xml:"marketingCompany,omitempty"` // minOccurs="0"

	OperatingCompany formats.AlphaNumericString_Length1To35 `xml:"operatingCompany,omitempty"` // minOccurs="0"

	OtherCompany formats.AlphaNumericString_Length1To35 `xml:"otherCompany,omitempty"` // minOccurs="0"
}

type ConversionRateDetailsTypeI struct {
	// NOT USED
	ConversionType formats.AlphaNumericString_Length1To3 `xml:"conversionType,omitempty"` // minOccurs="0"

	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"` // minOccurs="0"

	RateType formats.AlphaNumericString_Length1To3 `xml:"rateType,omitempty"` // minOccurs="0"

	PricingAmount *formats.NumericDecimal_Length1To18 `xml:"pricingAmount,omitempty"` // minOccurs="0"

	// NOT USED
	ConvertedValueAmount *formats.NumericDecimal_Length1To18 `xml:"convertedValueAmount,omitempty"` // minOccurs="0"

	// NOT USED
	DutyTaxFeeType formats.AlphaNumericString_Length1To3 `xml:"dutyTaxFeeType,omitempty"` // minOccurs="0"

	// NOT USED
	MeasurementValue *formats.NumericDecimal_Length1To18 `xml:"measurementValue,omitempty"` // minOccurs="0"

	// NOT USED
	MeasurementSignificance formats.AlphaNumericString_Length1To3 `xml:"measurementSignificance,omitempty"` // minOccurs="0"
}

type ConversionRateTypeI struct {
	// Fare Conversion option (/R,FC-xxx): 6345 = currencty  Fare Selection option (/R,FS-xxx): 9875 = 700 6345 = currencty
	ConversionRateDetails *ConversionRateDetailsTypeI `xml:"conversionRateDetails"`

	// Fare Conversion option (/R,FC-xxx): 6345 = currencty  Fare Selection option (/R,FS-xxx): 9875 = 700 6345 = currencty
	OtherConvRateDetails []*ConversionRateDetailsTypeI `xml:"otherConvRateDetails,omitempty"` // minOccurs="0" maxOccurs="19"
}

type CorporateFareIdentifiersTypeI struct {
	// Fare type (e.g. U for unifares)
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier,omitempty"` // minOccurs="0"

	// Corporate ID
	CorporateID []formats.AlphaNumericString_Length1To35 `xml:"corporateID,omitempty"` // minOccurs="0" maxOccurs="20"
}

type CorporateFareInformationType struct {
	CorporateFareIdentifiers []*CorporateFareIdentifiersTypeI `xml:"corporateFareIdentifiers"` // maxOccurs="20"
}

type DataTypeInformationTypeI struct {
	// Carrier fee code
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// Status event
	StatusEvent formats.AlphaNumericString_Length1To3 `xml:"statusEvent,omitempty"` // minOccurs="0"
}

type DateAndTimeDetailsTypeI struct {
	// A for not valid after, B for not valide before
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"` // minOccurs="0"

	// Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"` // minOccurs="0"
}

type DateAndTimeInformationTypeI struct {
	// Date and time details
	DateAndTimeDetails *DateAndTimeDetailsTypeI `xml:"dateAndTimeDetails,omitempty"` // minOccurs="0"
}

type DiscountPenaltyInformationTypeI struct {
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier"`

	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"` // minOccurs="0"

	Amount *formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"` // minOccurs="0"

	Percentage *formats.NumericInteger_Length1To8 `xml:"percentage,omitempty"` // minOccurs="0"
}

type DiscountPenaltyInformationTypeI_127142C struct {
	// Fare qualifier: PFB pricing for fare basis  (DOD-ODD not yet used. They will specify directionality)
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier"`

	// fareBasisCode
	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"` // minOccurs="0"
}

type DummySegmentTypeI struct {
}

type FareCategoryCodesTypeI struct {
	FareType formats.AlphaNumericString_Length1To20 `xml:"fareType"`

	OtherFareType []formats.AlphaNumericString_Length1To20 `xml:"otherFareType,omitempty"` // minOccurs="0" maxOccurs="8"
}

type FareDetailsTypeI struct {
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"` // minOccurs="0"

	Rate *formats.NumericInteger_Length1To8 `xml:"rate,omitempty"` // minOccurs="0"

	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"` // minOccurs="0"

	FareCategory formats.AlphaNumericString_Length1To3 `xml:"fareCategory,omitempty"` // minOccurs="0"
}

type FareDetailsTypeI_193637C struct {
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"` // minOccurs="0"

	Rate *formats.NumericDecimal_Length1To8 `xml:"rate,omitempty"` // minOccurs="0"

	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"` // minOccurs="0"

	FareCategory formats.AlphaNumericString_Length1To3 `xml:"fareCategory,omitempty"` // minOccurs="0"
}

type FareInformationType struct {
	// identify pax/fare basis group
	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber"`
}

type FareInformationTypeI struct {
	// PTC or fare discount
	ValueQualifier formats.AlphaNumericString_Length1To3 `xml:"valueQualifier,omitempty"` // minOccurs="0"

	Value *formats.NumericInteger_Length1To15 `xml:"value,omitempty"` // minOccurs="0"

	FareDetails *FareDetailsTypeI `xml:"fareDetails,omitempty"` // minOccurs="0"

	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber,omitempty"` // minOccurs="0"

	FareTypeGrouping *FareTypeGroupingInformationTypeI `xml:"fareTypeGrouping,omitempty"` // minOccurs="0"

	RateCategory []formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"` // minOccurs="0" maxOccurs="9"
}

type FareInformationTypeI_133428S struct {
	ValueQualifier formats.AlphaNumericString_Length1To3 `xml:"valueQualifier,omitempty"` // minOccurs="0"

	Value *formats.NumericInteger_Length1To15 `xml:"value,omitempty"` // minOccurs="0"

	FareDetails *FareDetailsTypeI_193637C `xml:"fareDetails,omitempty"` // minOccurs="0"

	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber,omitempty"` // minOccurs="0"

	FareTypeGrouping *FareTypeGroupingInformationTypeI_193636C `xml:"fareTypeGrouping,omitempty"` // minOccurs="0"

	RateCategory []formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"` // minOccurs="0" maxOccurs="9"
}

type FareQualifierDetailsTypeI struct {
	// Global indicator
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"` // minOccurs="0"

	FareCategories *FareCategoryCodesTypeI `xml:"fareCategories,omitempty"` // minOccurs="0"

	FareDetails *FareDetailsTypeI `xml:"fareDetails,omitempty"` // minOccurs="0"

	AdditionalFareDetails *AdditionalFareQualifierDetailsTypeI `xml:"additionalFareDetails,omitempty"` // minOccurs="0"

	// Zap off option: 9910: base / total fare indicator 5242: ticket designator 5004: amount 5482: percentage
	DiscountDetails []*DiscountPenaltyInformationTypeI `xml:"discountDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type FareQualifierDetailsTypeI_82845S struct {
	DiscountDetails *DiscountPenaltyInformationTypeI_127142C `xml:"discountDetails,omitempty"` // minOccurs="0"
}

type FareRouteInformationTypeI struct {
	// Corresponds to the Routing Number For future use !
	IdentificationNumber formats.AlphaNumericString_Length1To35 `xml:"identificationNumber,omitempty"` // minOccurs="0"
}

type FareTypeGroupingInformationTypeI struct {
	PricingGroup []formats.AlphaNumericString_Length1To6 `xml:"pricingGroup,omitempty"` // minOccurs="0" maxOccurs="5"
}

type FareTypeGroupingInformationTypeI_193636C struct {
	PricingGroup []formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"` // minOccurs="0" maxOccurs="5"
}

type FormOfPaymentDetailsTypeI struct {
	// PTCs/ Discount-Codes
	Type formats.AlphaNumericString_Length1To10 `xml:"type"`

	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"` // minOccurs="0"

	// Amount charged on this FOP
	Amount *formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"` // minOccurs="0"

	VendorCode formats.AlphaNumericString_Length1To35 `xml:"vendorCode,omitempty"` // minOccurs="0"

	// Used: The 6 first digits of the credit Card, if any
	CreditCardNumber formats.AlphaNumericString_Length1To35 `xml:"creditCardNumber,omitempty"` // minOccurs="0"

	ExpiryDate formats.AlphaNumericString_Length1To35 `xml:"expiryDate,omitempty"` // minOccurs="0"

	ApprovalCode formats.AlphaNumericString_Length1To17 `xml:"approvalCode,omitempty"` // minOccurs="0"

	SourceOfApproval formats.AlphaNumericString_Length1To3 `xml:"sourceOfApproval,omitempty"` // minOccurs="0"

	AuthorisedAmount *formats.NumericDecimal_Length1To18 `xml:"authorisedAmount,omitempty"` // minOccurs="0"

	AddressVerification formats.AlphaNumericString_Length1To3 `xml:"addressVerification,omitempty"` // minOccurs="0"

	CustomerAccount formats.AlphaNumericString_Length1To35 `xml:"customerAccount,omitempty"` // minOccurs="0"

	ExtendedPayment formats.AlphaNumericString_Length1To3 `xml:"extendedPayment,omitempty"` // minOccurs="0"

	FopFreeText formats.AlphaNumericString_Length1To70 `xml:"fopFreeText,omitempty"` // minOccurs="0"

	MembershipStatus formats.AlphaNumericString_Length1To3 `xml:"membershipStatus,omitempty"` // minOccurs="0"

	TransactionInfo formats.AlphaNumericString_Length1To35 `xml:"transactionInfo,omitempty"` // minOccurs="0"
}

type FormOfPaymentTypeI struct {
	// 9888: type of form of payment 5004: amount (amount charged on the current FOP in case of multi FOP) 1154: Bin Number, i.e. the first six digits of any credit card
	FormOfPayment *FormOfPaymentDetailsTypeI `xml:"formOfPayment"`

	// 9888: type of form of payment 5004: amount (amount charged on the current FOP in case of multi FOP) 1154: Bin Number, i.e. the first six digits of any credit card
	OtherFormOfPayment []*FormOfPaymentDetailsTypeI `xml:"otherFormOfPayment,omitempty"` // minOccurs="0" maxOccurs="98"
}

type FrequentTravellerIdentificationCodeType struct {
	// Frequent Traveller Info
	FrequentTravellerDetails []*FrequentTravellerIdentificationType `xml:"frequentTravellerDetails"` // maxOccurs="99"
}

type FrequentTravellerIdentificationType struct {
	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"` // minOccurs="0"

	Number formats.AlphaNumericString_Length1To28 `xml:"number,omitempty"` // minOccurs="0"

	// Specifies which traveller in the TIF segment the frequent traveller number applies (same as 9944 in TIF).
	CustomerReference formats.AlphaNumericString_Length1To10 `xml:"customerReference,omitempty"` // minOccurs="0"

	// status code: 'OK' if the frequent flyer card has been validated
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"` // minOccurs="0"

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"` // minOccurs="0"

	// For example : priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"` // minOccurs="0"

	// For example : Level description
	TierDescription formats.AlphaNumericString_Length1To35 `xml:"tierDescription,omitempty"` // minOccurs="0"

	// For example : Company code of alliance
	CompanyCode formats.AlphaNumericString_Length1To35 `xml:"companyCode,omitempty"` // minOccurs="0"

	CustomerValue *formats.NumericInteger_Length1To4 `xml:"customerValue,omitempty"` // minOccurs="0"
}

type GroupPassengerDetailsType struct {
	// Trigger
	PassengerReference *SegmentRepetitionControlTypeI `xml:"passengerReference"`

	// Passengers' tattoos
	TravellersID *SpecificTravellerTypeI `xml:"travellersID,omitempty"` // minOccurs="0"

	// pricing option at passenger level
	PsgDetailsInfo *PsgDetailsInfo `xml:"psgDetailsInfo"`
}

type PsgDetailsInfo struct {
	// PTC/Discount Code
	DiscountPtc *FareInformationTypeI_133428S `xml:"discountPtc"`

	// Tier level information
	FlequentFlyerDetails *FrequentTravellerIdentificationCodeType `xml:"flequentFlyerDetails,omitempty"` // minOccurs="0"
}

type LocationDetailsTypeI struct {
	// * Point of Sale Override (LocationDetails) * Point of Ticketing Override (OtherLocationDetails)
	City formats.AlphaNumericString_Length1To25 `xml:"city,omitempty"` // minOccurs="0"

	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"` // minOccurs="0"
}

type LocationTypeI struct {
	TrueLocationId formats.AlphaNumericString_Length1To25 `xml:"trueLocationId,omitempty"` // minOccurs="0"

	TrueLocation formats.AlphaNumericString_Length1To17 `xml:"trueLocation,omitempty"` // minOccurs="0"
}

type MarriageControlDetailsTypeI struct {
	// not used
	Relation formats.AlphaNumericString_Length1To3 `xml:"relation,omitempty"` // minOccurs="0"

	// not used
	MarriageIdentifier *formats.NumericInteger_Length1To10 `xml:"marriageIdentifier,omitempty"` // minOccurs="0"

	// not used
	LineNumber *formats.NumericInteger_Length1To6 `xml:"lineNumber,omitempty"` // minOccurs="0"

	// not used
	OtherRelation formats.AlphaNumericString_Length1To3 `xml:"otherRelation,omitempty"` // minOccurs="0"

	// not used
	CarrierCode formats.AlphaNumericString_Length1To35 `xml:"carrierCode,omitempty"` // minOccurs="0"
}

type MessageActionDetailsTypeI struct {
	MessageFunctionDetails *MessageFunctionBusinessDetailsTypeI `xml:"messageFunctionDetails,omitempty"` // minOccurs="0"

	ResponseType formats.AlphaNumericString_Length1To3 `xml:"responseType,omitempty"` // minOccurs="0"
}

type MessageFunctionBusinessDetailsTypeI struct {
	BusinessFunction formats.AlphaNumericString_Length1To3 `xml:"businessFunction,omitempty"` // minOccurs="0"

	MessageFunction formats.AlphaNumericString_Length1To3 `xml:"messageFunction,omitempty"` // minOccurs="0"

	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"` // minOccurs="0"

	// not used
	AdditionalMessageFunction []formats.AlphaNumericString_Length1To3 `xml:"additionalMessageFunction,omitempty"` // minOccurs="0" maxOccurs="20"
}

type MileageTimeDetailsTypeI struct {
	FlightLegMileage *formats.NumericInteger_Length1To18 `xml:"flightLegMileage,omitempty"` // minOccurs="0"

	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"` // minOccurs="0"

	ElapsedGroundTime formats.Time24_hhmM `xml:"elapsedGroundTime,omitempty"` // minOccurs="0"
}

type MonetaryInformationDetailsType struct {
	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"` // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"` // minOccurs="0"
}

type MonetaryInformationDetailsType_108001C struct {
	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"` // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"` // minOccurs="0"

	// Location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"` // minOccurs="0"
}

type MonetaryInformationDetailsType_209983C struct {
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount *formats.NumericDecimal_Length1To35 `xml:"amount,omitempty"` // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"` // minOccurs="0"
}

type MonetaryInformationType struct {
	// Indicate  - Action type: IPA - Amount - Currency
	MonetaryDetails *MonetaryInformationDetailsType_209983C `xml:"monetaryDetails"`
}

type MonetaryInformationType_69171S struct {
	// Monetary info
	MonetaryDetails []*MonetaryInformationDetailsType `xml:"monetaryDetails"` // maxOccurs="20"
}

type MonetaryInformationType_69173S struct {
	// Monetary info
	MonetaryDetails []*MonetaryInformationDetailsType_108001C `xml:"monetaryDetails"` // maxOccurs="20"
}

type OriginAndDestinationDetailsTypeI struct {
	Origin formats.AlphaNumericString_Length1To25 `xml:"origin,omitempty"` // minOccurs="0"

	Destination formats.AlphaNumericString_Length1To25 `xml:"destination,omitempty"` // minOccurs="0"
}

type OriginatorIdentificationDetailsTypeI struct {
	OriginatorId *formats.NumericInteger_Length1To9 `xml:"originatorId,omitempty"` // minOccurs="0"

	InHouseIdentification1 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification1,omitempty"` // minOccurs="0"

	InHouseIdentification2 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification2,omitempty"` // minOccurs="0"

	InHouseIdentification3 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification3,omitempty"` // minOccurs="0"
}

type PricingTicketingDetailsType struct {
	PriceTicketDetails *PricingTicketingInformationType `xml:"priceTicketDetails,omitempty"` // minOccurs="0"

	PriceTariffType formats.AlphaNumericString_Length1To3 `xml:"priceTariffType,omitempty"` // minOccurs="0"

	// Not used
	ProductDateTimeDetails *ProductDateTimeTypeI `xml:"productDateTimeDetails,omitempty"` // minOccurs="0"

	// * Validating Carrier : 9906 : MarketingCompany  * Award Publishing Carrier :  9906 : OtherCompany
	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"` // minOccurs="0"

	// Controlling Carrier
	CompanyNumberDetails *CompanyIdentificationNumbersTypeI `xml:"companyNumberDetails,omitempty"` // minOccurs="0"

	// Point Of Sale Override Information (3225 : City)
	LocationDetails *LocationDetailsTypeI `xml:"locationDetails,omitempty"` // minOccurs="0"

	// Point Of Ticketing Override Information (3225 : City)
	OtherLocationDetails *LocationDetailsTypeI `xml:"otherLocationDetails,omitempty"` // minOccurs="0"

	// PSR Number
	IdNumber formats.AlphaNumericString_Length1To35 `xml:"idNumber,omitempty"` // minOccurs="0"

	MonetaryAmount *formats.NumericDecimal_Length1To18 `xml:"monetaryAmount,omitempty"` // minOccurs="0"
}

type PricingTicketingInformationType struct {
	Indicators []formats.AlphaNumericString_Length1To3 `xml:"indicators,omitempty"` // minOccurs="0" maxOccurs="20"
}

type ProductDateTimeTypeI struct {
	DepartureDate formats.AlphaNumericString_Length1To35 `xml:"departureDate,omitempty"` // minOccurs="0"

	DepartureTime formats.Time24_hhmM `xml:"departureTime,omitempty"` // minOccurs="0"

	ArrivalDate formats.AlphaNumericString_Length1To35 `xml:"arrivalDate,omitempty"` // minOccurs="0"

	ArrivalTime formats.Time24_hhmM `xml:"arrivalTime,omitempty"` // minOccurs="0"

	DateVariation *formats.NumericInteger_Length1To1 `xml:"dateVariation,omitempty"` // minOccurs="0"
}

type ProductDetailsTypeI struct {
	Designator formats.AlphaNumericString_Length1To17 `xml:"designator"`

	AvailabilityStatus formats.AlphaNumericString_Length1To3 `xml:"availabilityStatus,omitempty"` // minOccurs="0"

	SpecialService formats.AlphaNumericString_Length1To3 `xml:"specialService,omitempty"` // minOccurs="0"

	Option []formats.AlphaNumericString_Length1To7 `xml:"option,omitempty"` // minOccurs="0" maxOccurs="3"
}

type ProductFacilitiesTypeI struct {
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"` // minOccurs="0"

	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"` // minOccurs="0"

	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"` // minOccurs="0"

	ExtensionCode []formats.AlphaNumericString_Length1To17 `xml:"extensionCode,omitempty"` // minOccurs="0" maxOccurs="26"
}

type ProductIdentificationDetailsTypeI struct {
	FlightNumber formats.AlphaNumericString_Length1To35 `xml:"flightNumber"`

	BookingClass formats.AlphaNumericString_Length1To17 `xml:"bookingClass,omitempty"` // minOccurs="0"

	OperationalSuffix formats.AlphaNumericString_Length1To3 `xml:"operationalSuffix,omitempty"` // minOccurs="0"

	Modifier []formats.AlphaNumericString_Length1To7 `xml:"modifier,omitempty"` // minOccurs="0" maxOccurs="3"
}

type ProductInformationTypeI struct {
	ProductDetailsQualifier formats.AlphaNumericString_Length1To3 `xml:"productDetailsQualifier,omitempty"` // minOccurs="0"

	BookingClassDetails []*ProductDetailsTypeI `xml:"bookingClassDetails,omitempty"` // minOccurs="0" maxOccurs="26"
}

type ProductTypeDetailsTypeI struct {
	FlightIndicator []formats.AlphaNumericString_Length1To6 `xml:"flightIndicator"` // maxOccurs="9"
}

type SegmentRepetitionControlDetailsType struct {
	// (Optional, information not taken into account at present)
	ProductReference *formats.NumericInteger_Length1To3 `xml:"productReference,omitempty"` // minOccurs="0"

	// (Optional, information not taken into account at present)
	NumberOfUnits *formats.NumericInteger_Length1To15 `xml:"numberOfUnits,omitempty"` // minOccurs="0"
}

type SegmentRepetitionControlDetailsTypeI struct {
	// NOT USED AT TVL LEVEL!
	Quantity *formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"` // minOccurs="0"

	// NOT USED AT TVL LEVEL!
	NumberOfUnits *formats.NumericInteger_Length1To15 `xml:"numberOfUnits,omitempty"` // minOccurs="0"
}

type SegmentRepetitionControlDetailsTypeI_193634C struct {
	Quantity *formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"` // minOccurs="0"

	NumberOfUnits *formats.NumericInteger_Length1To15 `xml:"numberOfUnits,omitempty"` // minOccurs="0"

	TotalNumberOfItems *formats.NumericInteger_Length1To15 `xml:"totalNumberOfItems,omitempty"` // minOccurs="0"
}

type SegmentRepetitionControlType struct {
	// Segment control details
	SegmentControlDetails []*SegmentRepetitionControlDetailsType `xml:"segmentControlDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type SegmentRepetitionControlTypeI struct {
	SegmentControlDetails []*SegmentRepetitionControlDetailsTypeI_193634C `xml:"segmentControlDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type SegmentRepetitionControlTypeI_69179S struct {
	SegmentControlDetails []*SegmentRepetitionControlDetailsTypeI `xml:"segmentControlDetails,omitempty"` // minOccurs="0" maxOccurs="9"
}

type SelectionDetailsInformationType struct {
	// Carrier fee type
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// Carrier fee status
	OptionInformation formats.AlphaNumericString_Length1To3 `xml:"optionInformation,omitempty"` // minOccurs="0"
}

type SelectionDetailsType struct {
	// Carrier fees options
	CarrierFeeDetails *SelectionDetailsInformationType `xml:"carrierFeeDetails"`
}

type SourceTypeDetailsTypeI struct {
	SourceQualifier1 formats.AlphaNumericString_Length1To3 `xml:"sourceQualifier1"`

	SourceQualifier2 formats.AlphaNumericString_Length1To3 `xml:"sourceQualifier2,omitempty"` // minOccurs="0"
}

type SpecificDataInformationTypeI struct {
	// Carrier fee information
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`
}

type SpecificTravellerDetailsTypeI struct {
	// Passenger Tattoos
	MeasurementValue *formats.NumericInteger_Length1To18 `xml:"measurementValue,omitempty"` // minOccurs="0"
}

type SpecificTravellerDetailsTypeI_108010C struct {
	ReferenceNumber formats.AlphaNumericString_Length1To10 `xml:"referenceNumber,omitempty"` // minOccurs="0"

	// Passenger Tattoos
	MeasurementValue *formats.NumericInteger_Length1To18 `xml:"measurementValue,omitempty"` // minOccurs="0"

	FirstDate formats.AlphaNumericString_Length1To35 `xml:"firstDate,omitempty"` // minOccurs="0"

	Surname formats.AlphaNumericString_Length1To70 `xml:"surname,omitempty"` // minOccurs="0"

	FirstName formats.AlphaNumericString_Length1To70 `xml:"firstName,omitempty"` // minOccurs="0"
}

type SpecificTravellerTypeI struct {
	TravellerDetails []*SpecificTravellerDetailsTypeI_108010C `xml:"travellerDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type SpecificTravellerTypeI_65902S struct {
	// Traveller details
	TravellerDetails []*SpecificTravellerDetailsTypeI `xml:"travellerDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type StationInformationTypeI struct {
	GateDescription formats.AlphaNumericString_Length1To6 `xml:"gateDescription,omitempty"` // minOccurs="0"

	Terminal formats.AlphaNumericString_Length1To25 `xml:"terminal,omitempty"` // minOccurs="0"

	Concourse formats.AlphaNumericString_Length1To25 `xml:"concourse,omitempty"` // minOccurs="0"
}

type TaxDetailsTypeI struct {
	Rate formats.AlphaNumericString_Length1To17 `xml:"rate,omitempty"` // minOccurs="0"

	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"` // minOccurs="0"

	// Not used
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"` // minOccurs="0"

	Type []formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TaxTypeI struct {
	TaxCategory formats.AlphaNumericString_Length1To3 `xml:"taxCategory,omitempty"` // minOccurs="0"

	TaxDetails []*TaxDetailsTypeI `xml:"taxDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TravelProductInformationTypeI struct {
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"` // minOccurs="0"

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"` // minOccurs="0"

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"` // minOccurs="0"

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"` // minOccurs="0"

	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"` // minOccurs="0"

	FlightTypeDetails *ProductTypeDetailsTypeI `xml:"flightTypeDetails,omitempty"` // minOccurs="0"

	// Connection indicators: order in the group.
	ItemNumber *formats.NumericInteger_Length1To6 `xml:"itemNumber,omitempty"` // minOccurs="0"

	SpecialSegment formats.AlphaNumericString_Length1To3 `xml:"specialSegment,omitempty"` // minOccurs="0"

	// not used
	MarriageDetails []*MarriageControlDetailsTypeI `xml:"marriageDetails,omitempty"` // minOccurs="0" maxOccurs="99"
}

type TravellerTimeDetailsTypeI struct {
	DepartureTime formats.Time24_hhmM `xml:"departureTime,omitempty"` // minOccurs="0"

	ArrivalTime formats.Time24_hhmM `xml:"arrivalTime,omitempty"` // minOccurs="0"

	CheckInDateTime formats.AlphaNumericString_Length1To10 `xml:"checkInDateTime,omitempty"` // minOccurs="0"
}
