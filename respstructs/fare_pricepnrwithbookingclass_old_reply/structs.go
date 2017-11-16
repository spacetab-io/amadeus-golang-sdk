package fare_pricepnrwithbookingclass_reply

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type FarePricePNRWithBookingClassReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A Fare_PricePNRWithBookingClassReply"`

	ApplicationError *ErrorGroupType `xml:"applicationError,omitempty"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`

	FareList *FareList `xml:"fareList,omitempty"`
}

type FareList struct {
	// Pricing information such as pricing rule and sales indicator.
	PricingInformation *PricingTicketingSubsequentTypeI `xml:"pricingInformation,omitempty"`

	// Fare reference number. Ordering information is not relevant here.
	FareReference *ItemReferencesAndVersionsType_94584S `xml:"fareReference,omitempty"`

	// Fare Indicators
	FareIndicators *FareInformationType `xml:"fareIndicators,omitempty"`

	// Last date to ticket the fare.
	LastTktDate *StructuredDateTimeInformationType `xml:"lastTktDate,omitempty"`

	// Validating carrier identification.
	ValidatingCarrier *TransportIdentifierType `xml:"validatingCarrier,omitempty"`

	// Passenger/segment association of the fare is specified here.
	PaxSegReference *ReferenceInformationTypeI `xml:"paxSegReference,omitempty"`

	FareDataInformation *MonetaryInformationType `xml:"fareDataInformation,omitempty"`

	TaxInformation *TaxInformation `xml:"taxInformation,omitempty"`

	// Banker's rates are used to convert amounts of the TST (converts base fare to equivalent fare) 1st C661 : 1st bankers' rate which is a percentage (no currency) 2nd C661 : 2nd bankers' rate which is currency+amount.
	BankerRates *ConversionRateTypeI `xml:"bankerRates,omitempty"`

	PassengerInformation *PassengerInformation `xml:"passengerInformation,omitempty"`
	// Origin and destination of the fare. 1st C3225 occurence : origin city. 2nd C3225 occurence : destination city
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`

	SegmentInformation *SegmentInformation `xml:"segmentInformation,omitempty"`

	// Other pricing information such as endorsement, tour name...
	OtherPricingInfo *CodedAttributeType_39223S `xml:"otherPricingInfo,omitempty"`

	WarningInformation *WarningInformation `xml:"warningInformation,omitempty"`

	AutomaticReissueInfo *AutomaticReissueInfo `xml:"automaticReissueInfo,omitempty"`

	// Corporate number
	CorporateInfo *CorporateFareInformationType `xml:"corporateInfo,omitempty"`

	FeeBreakdown *FeeBreakdown `xml:"feeBreakdown,omitempty"`

	// convey the mileage information
	Mileage *AdditionalProductDetailsTypeI `xml:"mileage,omitempty"`

	FareComponentDetailsGroup *FareComponentDetailsType `xml:"fareComponentDetailsGroup,omitempty"`

	EndFareList *DummySegmentTypeI `xml:"endFareList,omitempty"`
}

type TaxInformation struct {
	// Tax details
	TaxDetails *DutyTaxFeeDetailsTypeU `xml:"taxDetails,omitempty"`

	// Amount details. If the tax is a passenger facility charge (PFC) the detail of the airports related taxes is given here.
	AmountDetails *MonetaryInformationTypeI `xml:"amountDetails,omitempty"`
}

type PassengerInformation struct {
	// Penalty/discount details specified in the request.
	PenDisInformation *DiscountAndPenaltyInformationTypeI_6128S `xml:"penDisInformation,omitempty"`

	// Reference of passengers that have a type code.
	PassengerReference *ReferenceInformationTypeI `xml:"passengerReference,omitempty"`
}

type SegmentInformation struct {
	// Connection information.
	ConnexInformation *ConnectionTypeI `xml:"connexInformation,omitempty"`

	// Details on open segments added to the price calculation. These open segments exist only in the fare calculated, they have no equivalent in the PNR itinerary. This segment gives also information on booking class for best buy transactions.
	SegDetails *TravelProductInformationTypeI_26322S `xml:"segDetails,omitempty"`

	// Fare basis information
	FareQualifier *FareQualifierDetailsTypeI `xml:"fareQualifier,omitempty"`

	// Validity information for this fare
	ValidityInformation *StructuredDateTimeInformationType `xml:"validityInformation,omitempty"`

	// Baggage allowance information
	BagAllowanceInformation *ExcessBaggageTypeI `xml:"bagAllowanceInformation,omitempty"`

	// Reference of the segment associated to the group.
	SegmentReference *ReferenceInformationTypeI `xml:"segmentReference,omitempty"`

	// The segment order in the pricing response can be different than the one of the PNR itinerary (segments are reordered at price calculation time). This order inform,ation is conveyed by the sequence number. If this order information is not present then the order is by default the one of the PNR.
	SequenceInformation *ItemReferencesAndVersionsType `xml:"sequenceInformation,omitempty"`
}

type WarningInformation struct {
	// Fare warning information code.
	WarningCode *ApplicationErrorInformationType `xml:"warningCode,omitempty"`

	// Description in free flow text of the warning concerning the fare.
	WarningText *InteractiveFreeTextTypeI_6759S `xml:"warningText,omitempty"`
}

type AutomaticReissueInfo struct {
	// This segment contains the original ticket number.
	TicketInfo *TicketNumberTypeI `xml:"ticketInfo,omitempty"`

	// This segment contains the coupon number (in absolute) corresponding to the first coupon for use from the last flawn segment.
	CouponInfo *CouponInformationTypeI `xml:"couponInfo,omitempty"`

	PaperCouponRange *PaperCouponRange `xml:"paperCouponRange,omitempty"`

	// Base fare Information
	BaseFareInfo *MonetaryInformationTypeI_20897S `xml:"baseFareInfo,omitempty"`

	FirstDpiGroup *FirstDpiGroup `xml:"firstDpiGroup,omitempty"`

	SecondDpiGroup *SecondDpiGroup `xml:"secondDpiGroup,omitempty"`

	// this segment conveys specific reissue attributes like Revalidation flag.
	ReissueAttributes *CodedAttributeType `xml:"reissueAttributes,omitempty"`
}

type PaperCouponRange struct {
	// This segment contains the original ticket number.
	TicketInfo *TicketNumberTypeI `xml:"ticketInfo,omitempty"`

	// This segment contains the coupon number (in absolute) corresponding to the first coupon for use from the last flawn segment.
	CouponInfo *CouponInformationTypeI `xml:"couponInfo,omitempty"`
}

type FirstDpiGroup struct {
	// Penalty amount in reissue currency
	ReIssuePenalty *DiscountAndPenaltyInformationTypeI `xml:"reIssuePenalty,omitempty"`

	// Reissue Informations
	ReissueInfo *MonetaryInformationTypeI_20897S `xml:"reissueInfo,omitempty"`

	// Old Tax informations
	OldTaxInfo *MonetaryInformationTypeI_20897S `xml:"oldTaxInfo,omitempty"`

	// Balance Reissue Informations
	ReissueBalanceInfo *MonetaryInformationTypeI_20897S `xml:"reissueBalanceInfo,omitempty"`
}

type SecondDpiGroup struct {
	// Discount and penalty info.
	Penalty *DiscountAndPenaltyInformationTypeI `xml:"penalty,omitempty"`

	// Residual Value information
	ResidualValueInfo *MonetaryInformationTypeI_20897S `xml:"residualValueInfo,omitempty"`

	// Old Tax informations
	OldTaxInfo *MonetaryInformationTypeI_20897S `xml:"oldTaxInfo,omitempty"`

	// Balance issue Informations
	IssueBalanceInfo *MonetaryInformationTypeI_20897S `xml:"issueBalanceInfo,omitempty"`
}

type FeeBreakdown struct {
	// Nature of the fee (OB, OC)
	FeeType *SelectionDetailsTypeI `xml:"feeType,omitempty"`

	FeeDetails *FeeDetails `xml:"feeDetails,omitempty"`
}

type FeeDetails struct {
	// Fee information
	FeeInfo *SpecificDataInformationTypeI `xml:"feeInfo,omitempty"`

	// Attributes of this fee (commercial description)
	FeeDescription *InteractiveFreeTextTypeI `xml:"feeDescription,omitempty"`

	// Fee associated amounts: amount with/without tax, total tax amount
	FeeAmounts *MonetaryInformationTypeI_39230S `xml:"feeAmounts,omitempty"`

	// taxes related to this fee
	FeeTaxes *TaxTypeI `xml:"feeTaxes,omitempty"`
}

type AdditionalFareQualifierDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A AdditionalFareQualifierDetailsTypeI"`

	// Primary code of the fare basis. This is not a codeset but a free flow text field.
	PrimaryCode formats.AlphaNumericString_Length1To3 `xml:"primaryCode,omitempty"`

	// Fare basis code of the fare basis. This is not a codeset but a free flow text field.
	FareBasisCode formats.AlphaNumericString_Length1To6 `xml:"fareBasisCode,omitempty"`

	// Ticket designator of the fare basis
	TicketDesignator formats.AlphaNumericString_Length1To6 `xml:"ticketDesignator,omitempty"`

	// For any query : discount ticket designator to be assigned by Fare Quote server. For any response : priced PTCs
	DiscTktDesignator formats.AlphaNumericString_Length1To11 `xml:"discTktDesignator,omitempty"`
}

type AdditionalProductDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A AdditionalProductDetailsTypeI"`

	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`
}

type ApplicationErrorDetailType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ApplicationErrorDetailType"`

	// Code identifying the data validation error condition.
	ErrorCode formats.AlphaNumericString_Length1To5 `xml:"errorCode,omitempty"`

	// Identification of a code list.
	ErrorCategory formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner formats.AlphaNumericString_Length1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorDetailType_48648C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ApplicationErrorDetailType_48648C"`

	// Code identifying the data validation error condition.
	ApplicationErrorCode formats.AlphaNumericString_Length1To5 `xml:"applicationErrorCode,omitempty"`

	// Identification of a code list.
	CodeListQualifier formats.AlphaNumericString_Length1To3 `xml:"codeListQualifier,omitempty"`

	// Code identifying the agency responsible for a code list.
	CodeListResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"codeListResponsibleAgency,omitempty"`
}

type ApplicationErrorInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ApplicationErrorInformationType"`

	// Application error details.
	ApplicationErrorDetail *ApplicationErrorDetailType_48648C `xml:"applicationErrorDetail,omitempty"`
}

type ApplicationErrorInformationType_84497S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ApplicationErrorInformationType_84497S"`

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails,omitempty"`
}

type BaggageDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A BaggageDetailsTypeI"`

	// Baggage allowance quantity (piece concept)
	BaggageQuantity formats.NumericInteger_Length1To1 `xml:"baggageQuantity,omitempty"`

	// Baggage allowance weight
	BaggageWeight formats.NumericInteger_Length1To2 `xml:"baggageWeight,omitempty"`

	// Baggage allowance type (weight/number)
	BaggageType formats.AlphaNumericString_Length1To3 `xml:"baggageType,omitempty"`

	// Measurement unit for weighing baggage allowance
	MeasureUnit formats.AlphaNumericString_Length1To1 `xml:"measureUnit,omitempty"`
}

type CodedAttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CodedAttributeInformationType"`

	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// provides a description for the attribute
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeInformationType_66047C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CodedAttributeInformationType_66047C"`

	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType,omitempty"`

	// provides a description for the attribute
	AttributeDescription formats.AlphaNumericString_Length1To500 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CodedAttributeType"`

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`
}

type CodedAttributeType_39223S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CodedAttributeType_39223S"`

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType_66047C `xml:"attributeDetails,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DummyNET struct{}

type CompanyIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CompanyIdentificationTypeI"`

	// Carrier code
	CarrierCode formats.AlphaNumericString_Length1To3 `xml:"carrierCode,omitempty"`
}

type ConnectionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ConnectionDetailsTypeI"`

	// Specify ARNK and surface segments  not included in the fare routing.
	RoutingInformation formats.AlphaNumericString_Length1To4 `xml:"routingInformation,omitempty"`

	// Type of connection for the flight
	ConnexType formats.AlphaNumericString_Length1To1 `xml:"connexType,omitempty"`
}

type ConnectionTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ConnectionTypeI"`

	// Connection details
	ConnecDetails *ConnectionDetailsTypeI `xml:"connecDetails,omitempty"`
}

type ConversionRateDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ConversionRateDetailsTypeI"`

	// Currency of the rate
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// Amount/percentage
	Amount formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`
}

type ConversionRateTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ConversionRateTypeI"`

	// First rate detail.
	FirstRateDetail *ConversionRateDetailsTypeI `xml:"firstRateDetail,omitempty"`

	// Second rate detail.
	SecondRateDetail *ConversionRateDetailsTypeI `xml:"secondRateDetail,omitempty"`
}

type CorporateFareIdentifiersTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CorporateFareIdentifiersTypeI"`

	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier,omitempty"`

	CorporateID formats.AlphaNumericString_Length1To35 `xml:"corporateID,omitempty"`
}

type CorporateFareInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CorporateFareInformationType"`

	CorporateFareIdentifiers *CorporateFareIdentifiersTypeI `xml:"corporateFareIdentifiers,omitempty"`
}

type CouponDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CouponDetailsType"`

	// Tattoo + type of the product identifying the coupon.
	ProductId *ReferenceInfoType `xml:"productId,omitempty"`

	// Flight Connection Type
	FlightConnectionType *TravelProductInformationType `xml:"flightConnectionType,omitempty"`
}

type CouponInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CouponInformationDetailsTypeI"`

	// Coupon number
	CpnNumber formats.AlphaNumericString_Length1To6 `xml:"cpnNumber,omitempty"`
}

type CouponInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A CouponInformationTypeI"`

	// Details on coupon
	CouponDetails *CouponInformationDetailsTypeI `xml:"couponDetails,omitempty"`

	// Details on coupon
	OtherCouponDetails *CouponInformationDetailsTypeI `xml:"otherCouponDetails,omitempty"`
}

type DataInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DataInformationTypeI"`

	// fee attribute
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type DataTypeInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DataTypeInformationTypeI"`

	// fee subcode
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type DiscountAndPenaltyInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DiscountAndPenaltyInformationTypeI"`

	// Used to specify penalty information
	PenDisData *DiscountPenaltyMonetaryInformationTypeI_29792C `xml:"penDisData,omitempty"`
}

type DiscountAndPenaltyInformationTypeI_6128S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DiscountAndPenaltyInformationTypeI_6128S"`

	// Qualify the type of information.  Penalties are not passenger associated and are pure monetary information. Discount are passenger associated but only discount code is specified.
	InfoQualifier formats.AlphaNumericString_Length1To3 `xml:"infoQualifier,omitempty"`

	// Used to specify penalty information.
	PenDisData *DiscountPenaltyMonetaryInformationTypeI `xml:"penDisData,omitempty"`
}

type DiscountPenaltyInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DiscountPenaltyInformationTypeI"`

	// Discount off type.
	ZapOffType formats.AlphaNumericString_Length1To3 `xml:"zapOffType,omitempty"`

	// Discount amount
	ZapOffAmount formats.NumericDecimal_Length1To11 `xml:"zapOffAmount,omitempty"`

	// Discount percentage.
	ZapOffPercentage formats.NumericInteger_Length1To8 `xml:"zapOffPercentage,omitempty"`
}

type DiscountPenaltyInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DiscountPenaltyInformationType"`

	FareQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"fareQualifier,omitempty"`
}

type DiscountPenaltyMonetaryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DiscountPenaltyMonetaryInformationTypeI"`

	// Type of penalty.
	PenaltyType formats.AlphaNumericString_Length1To3 `xml:"penaltyType,omitempty"`

	// The penalty amount can be described differently: amount/percentage.
	PenaltyQualifier formats.AlphaNumericString_Length1To3 `xml:"penaltyQualifier,omitempty"`

	// Amount of the penalty.
	PenaltyAmount formats.NumericDecimal_Length1To11 `xml:"penaltyAmount,omitempty"`

	// This discount code is defined by the airlines. This cannot be coded as airlines might apply any combination of letters for their discounts.
	DiscountCode formats.AlphaNumericString_Length1To6 `xml:"discountCode,omitempty"`

	// Penalty currency code.
	PenaltyCurrency formats.AlphaNumericString_Length1To3 `xml:"penaltyCurrency,omitempty"`
}

type DiscountPenaltyMonetaryInformationTypeI_29792C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DiscountPenaltyMonetaryInformationTypeI_29792C"`

	// The amount Type can be a percentage or an amount
	PenaltyQualifier formats.AlphaNumericString_Length1To3 `xml:"penaltyQualifier,omitempty"`

	// specify the value
	PenaltyAmount formats.NumericDecimal_Length1To18 `xml:"penaltyAmount,omitempty"`

	// penalty currency code
	PenaltyCurrency formats.AlphaNumericString_Length1To3 `xml:"penaltyCurrency,omitempty"`
}

type DummySegmentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DummySegmentTypeI"`
}

type DutyTaxFeeAccountDetailTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DutyTaxFeeAccountDetailTypeU"`

	// Iso country of the tax
	IsoCountry formats.AlphaNumericString_Length1To3 `xml:"isoCountry,omitempty"`
}

type DutyTaxFeeDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DutyTaxFeeDetailsTypeU"`

	// Tax data qualifier
	TaxQualifier formats.AlphaNumericString_Length1To1 `xml:"taxQualifier,omitempty"`

	// Tax type identification
	TaxIdentification *DutyTaxFeeTypeDetailsTypeU `xml:"taxIdentification,omitempty"`

	// Type of the tax
	TaxType *DutyTaxFeeAccountDetailTypeU `xml:"taxType,omitempty"`

	// Nature of the tax
	TaxNature formats.AlphaNumericString_Length1To3 `xml:"taxNature,omitempty"`

	// Exempt tax indicator. If an tax is Exempted no amount is provided for this tax.
	TaxExempt formats.AlphaNumericString_Length1To3 `xml:"taxExempt,omitempty"`
}

type DutyTaxFeeTypeDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A DutyTaxFeeTypeDetailsTypeU"`

	// Tax type identifier
	TaxIdentifier formats.AlphaNumericString_Length1To3 `xml:"taxIdentifier,omitempty"`
}

type ErrorGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ErrorGroupType"`

	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType_84497S `xml:"errorOrWarningCodeDetails,omitempty"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"`
}

type ExcessBaggageTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ExcessBaggageTypeI"`

	// Baggage allowance information details
	BagAllowanceDetails *BaggageDetailsTypeI `xml:"bagAllowanceDetails,omitempty"`
}

type FareComponentDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A FareComponentDetailsType"`

	// fare Component identification
	FareComponentID *ItemNumberType `xml:"fareComponentID,omitempty"`

	// Market information related to fare component
	MarketFareComponent *TravelProductInformationTypeI `xml:"marketFareComponent,omitempty"`

	// Monetary Information
	MonetaryInformation *MonetaryInformationType_157196S `xml:"monetaryInformation,omitempty"`

	// Component Class information
	ComponentClassInfo *PricingOrTicketingSubsequentType `xml:"componentClassInfo,omitempty"`

	// Fare Qualifier Detail
	FareQualifiersDetail *FareQualifierDetailsType `xml:"fareQualifiersDetail,omitempty"`

	CouponDetailsGroup *CouponDetailsType `xml:"couponDetailsGroup,omitempty"`
}

type FareDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A FareDetailsType"`

	// fare indicators
	FareCategory formats.AMA_EDICodesetType_Length1to3 `xml:"fareCategory,omitempty"`
}

type FareInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A FareInformationType"`

	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`
}

type FareQualifierDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A FareQualifierDetailsTypeI"`

	// Type of movement for this segment to take into account by Fare Quote to calculate the fare.
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`

	// Fare basis detail
	FareBasisDetails *AdditionalFareQualifierDetailsTypeI `xml:"fareBasisDetails,omitempty"`

	// Discount data for zap off to apply to price calculation.
	ZapOffDetails *DiscountPenaltyInformationTypeI `xml:"zapOffDetails,omitempty"`
}

type FareQualifierDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A FareQualifierDetailsType"`

	DiscountDetails *DiscountPenaltyInformationType `xml:"discountDetails,omitempty"`
}

type FreeTextDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A FreeTextDetailsType"`

	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	CompanyId formats.AlphaNumericString_Length1To35 `xml:"companyId,omitempty"`

	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`

	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A FreeTextInformationType"`

	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText formats.AlphaNumericString_Length1To199 `xml:"freeText,omitempty"`
}

type FreeTextQualificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A FreeTextQualificationTypeI"`

	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A InteractiveFreeTextTypeI"`

	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`

	FreeText formats.AlphaNumericString_Length1To10 `xml:"freeText,omitempty"`
}

type InteractiveFreeTextTypeI_6759S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A InteractiveFreeTextTypeI_6759S"`

	// Free flow text describing the error
	ErrorFreeText formats.AlphaNumericString_Length1To70 `xml:"errorFreeText,omitempty"`
}

type ItemNumberIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ItemNumberIdentificationType"`

	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`
}

type ItemNumberType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ItemNumberType"`

	ItemNumberDetails *ItemNumberIdentificationType `xml:"itemNumberDetails,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ItemReferencesAndVersionsType"`

	// Identification details : order number
	SequenceSection *UniqueIdDescriptionType `xml:"sequenceSection,omitempty"`
}

type ItemReferencesAndVersionsType_94584S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ItemReferencesAndVersionsType_94584S"`

	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// Tattoo number
	UniqueReference formats.NumericInteger_Length1To5 `xml:"uniqueReference,omitempty"`
}

type LocationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A LocationTypeI"`

	TrueLocationId formats.AlphaNumericString_Length1To25 `xml:"trueLocationId,omitempty"`
}

type LocationTypeI_47688C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A LocationTypeI_47688C"`

	// Code of the city.
	CityCode formats.AlphaString_Length1To3 `xml:"cityCode,omitempty"`
}

type MileageTimeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MileageTimeDetailsTypeI"`

	// mileage total associated to the TST
	TotalMileage formats.NumericInteger_Length1To8 `xml:"totalMileage,omitempty"`
}

type MonetaryInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationDetailsTypeI"`

	// Qualify the type of fare defined in this composite
	FareDataQualifier formats.AlphaNumericString_Length1To3 `xml:"fareDataQualifier,omitempty"`

	// Fare data amount
	FareAmount formats.AlphaNumericString_Length1To11 `xml:"fareAmount,omitempty"`

	// Fare data currency code
	FareCurrency formats.AlphaNumericString_Length1To3 `xml:"fareCurrency,omitempty"`

	// Location of the fare data (PFCs specific)
	FareLocation formats.AlphaNumericString_Length3To3 `xml:"fareLocation,omitempty"`
}

type MonetaryInformationDetailsTypeI_37257C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationDetailsTypeI_37257C"`

	// Type qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// amount
	Amount formats.AlphaNumericString_Length1To35 `xml:"amount,omitempty"`

	// currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`
}

type MonetaryInformationDetailsTypeI_63727C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationDetailsTypeI_63727C"`

	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Location
	Location formats.AlphaNumericString_Length1To3 `xml:"location,omitempty"`
}

type MonetaryInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationDetailsType"`

	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount formats.AlphaNumericString_Length1To35 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsType_223826C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationDetailsType_223826C"`

	FareDataQualifier formats.AlphaNumericString_Length1To3 `xml:"fareDataQualifier,omitempty"`

	// Amount
	FareAmount formats.AlphaNumericString_Length1To11 `xml:"fareAmount,omitempty"`

	// Currency
	FareCurrency formats.AlphaNumericString_Length1To3 `xml:"fareCurrency,omitempty"`

	// location
	FareLocation formats.AlphaNumericString_Length3To3 `xml:"fareLocation,omitempty"`
}

type MonetaryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationTypeI"`

	// Main fare data infomation, can b thee base or  the total fare information which are mandatory  anyway
	FareDataMainInformation *MonetaryInformationDetailsTypeI `xml:"fareDataMainInformation,omitempty"`

	// Supplementary fare data information
	FareDataSupInformation *MonetaryInformationDetailsTypeI `xml:"fareDataSupInformation,omitempty"`
}

type MonetaryInformationTypeI_20897S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationTypeI_20897S"`

	// monetaryDetails
	MonetaryDetails *MonetaryInformationDetailsTypeI_37257C `xml:"monetaryDetails,omitempty"`

	OtherMonetaryDetails *MonetaryInformationDetailsTypeI_37257C `xml:"otherMonetaryDetails,omitempty"`
}

type MonetaryInformationTypeI_39230S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationTypeI_39230S"`

	// Monetary info
	MonetaryDetails *MonetaryInformationDetailsTypeI_63727C `xml:"monetaryDetails,omitempty"`
}

type MonetaryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationType"`

	FareDataMainInformation *MonetaryInformationDetailsType_223826C `xml:"fareDataMainInformation,omitempty"`

	FareDataSupInformation *MonetaryInformationDetailsType_223826C `xml:"fareDataSupInformation,omitempty"`
}

type MonetaryInformationType_157196S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A MonetaryInformationType_157196S"`

	// Monetary information per fare component
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails,omitempty"`

	// Other monetary information per fare component
	OtherMonetaryDetails *MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"`
}

type OriginAndDestinationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A OriginAndDestinationDetailsTypeI"`

	// Code of the city.
	CityCode formats.AlphaNumericString_Length1To4 `xml:"cityCode,omitempty"`
}

type PricingOrTicketingSubsequentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A PricingOrTicketingSubsequentType"`

	// RATE OR TARIFF CLASS INFORMATION
	FareBasisDetails *RateTariffClassInformationType `xml:"fareBasisDetails,omitempty"`
}

type PricingTicketingSubsequentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A PricingTicketingSubsequentTypeI"`

	// Information on TST type.
	TstInformation *RateTariffClassInformationTypeI `xml:"tstInformation,omitempty"`

	// International sales indicator
	SalesIndicator formats.AlphaString_Length1To2 `xml:"salesIndicator,omitempty"`

	// Fare calculation mode indicator. This indicator specifies the type fare.
	Fcmi formats.AlphaNumericString_Length1To1 `xml:"fcmi,omitempty"`

	// Information of original fare used to create TST. The TST is created from Best Fare ( possible or available).
	BestFareType formats.AlphaNumericString_Length1To3 `xml:"bestFareType,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ProductIdentificationDetailsTypeI"`

	// OPEN or AIR are the two identifications accepted.  OPEN means the segment described here is an open segment. AIR means that it is a valid AIR segment.
	Identification formats.AlphaNumericString_Length1To6 `xml:"identification,omitempty"`

	// to describe the transportation class.
	BookingClass formats.AlphaNumericString_Length1To17 `xml:"bookingClass,omitempty"`

	// Class of service to use in order to price the extra segment.
	ClassOfService formats.AlphaString_Length1To1 `xml:"classOfService,omitempty"`
}

type ProductTypeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ProductTypeDetailsType"`

	// TST Connection Type
	FlightIndicator formats.AlphaNumericString_Length1To1 `xml:"flightIndicator,omitempty"`
}

type RateTariffClassInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A RateTariffClassInformationTypeI"`

	// Indicator qualifying the type of TST (basically manual or automatic)
	TstIndicator formats.AlphaNumericString_Length1To1 `xml:"tstIndicator,omitempty"`
}

type RateTariffClassInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A RateTariffClassInformationType"`

	// Fare Basis Code
	RateTariffClass formats.AlphaNumericString_Length1To35 `xml:"rateTariffClass,omitempty"`

	// Ticket Designator
	OtherRateTariffClass formats.AlphaNumericString_Length1To35 `xml:"otherRateTariffClass,omitempty"`
}

type ReferenceInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ReferenceInfoType"`

	ReferenceDetails *ReferencingDetailsType `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ReferenceInformationTypeI"`

	// Passenger/segment/TST/fare reference details
	RefDetails *ReferencingDetailsTypeI `xml:"refDetails,omitempty"`
}

type ReferencingDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ReferencingDetailsTypeI"`

	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ReferencingDetailsType"`

	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	Value formats.AlphaNumericString_Length1To60 `xml:"value,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ReservationControlInformationDetailsTypeI"`

	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A ReservationControlInformationTypeI"`

	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A SelectionDetailsInformationTypeI"`

	Option formats.AlphaNumericString_Length1To2 `xml:"option,omitempty"`
}

type SelectionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A SelectionDetailsTypeI"`

	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails,omitempty"`
}

type SpecificDataInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A SpecificDataInformationTypeI"`

	// Carrier fee code
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation,omitempty"`

	// Carrier fee application code  (NI, NR, CM, NC)
	DataInformation *DataInformationTypeI `xml:"dataInformation,omitempty"`
}

type StructuredDateTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A StructuredDateTimeInformationType"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A StructuredDateTimeType"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.NumericInteger_Length1To6 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type TaxDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A TaxDetailsTypeI"`

	// Tax Amount
	Rate formats.AlphaNumericString_Length1To17 `xml:"rate,omitempty"`

	// ISO code identifying Country
	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"`

	// ISO code identifying currency
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// Tax designator code
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// tax designator code.
	SecondType formats.AlphaNumericString_Length1To3 `xml:"secondType,omitempty"`
}

type TaxTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A TaxTypeI"`

	// Tax details
	TaxDetails *TaxDetailsTypeI `xml:"taxDetails,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type TicketNumberDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A TicketNumberDetailsTypeI"`

	// Ticket number
	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`

	// ticket type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type TicketNumberTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A TicketNumberTypeI"`

	// Details on the document
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails,omitempty"`
}

type TransportIdentifierType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A TransportIdentifierType"`

	// Information related to validating carrier.
	CarrierInformation *CompanyIdentificationTypeI `xml:"carrierInformation,omitempty"`
}

type TravelProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A TravelProductInformationTypeI"`

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`
}

type TravelProductInformationTypeI_26322S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A TravelProductInformationTypeI_26322S"`

	// City of departure for this extra segment.
	DepartureCity *LocationTypeI_47688C `xml:"departureCity,omitempty"`

	// City of arrival for this extra segment.
	ArrivalCity *LocationTypeI_47688C `xml:"arrivalCity,omitempty"`

	// Airline detail information of the extra segment.
	AirlineDetail *CompanyIdentificationTypeI `xml:"airlineDetail,omitempty"`

	// Segment detail information.
	SegmentDetail *ProductIdentificationDetailsTypeI `xml:"segmentDetail,omitempty"`

	// Ticketing status for this segment. Relevant only in case of reply.
	TicketingStatus formats.AlphaNumericString_Length1To2 `xml:"ticketingStatus,omitempty"`
}

type TravelProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A TravelProductInformationType"`

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`

	// TST Connection Type
	FlightTypeDetails *ProductTypeDetailsType `xml:"flightTypeDetails,omitempty"`
}

type UniqueIdDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A UniqueIdDescriptionType"`

	// Number specifying the ordering information of the item described within a group.
	SequenceNumber formats.NumericInteger_Length1To3 `xml:"sequenceNumber,omitempty"`
}
