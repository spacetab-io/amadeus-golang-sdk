package ticket_display_tst_reply

import (
	"encoding/xml"

	"github.com/kidem/amadeus-ws-go/formats"
)

type TicketDisplayTSTReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A Ticket_DisplayTSTReply"`

	// Scrolling information used for long messages. C673.1050 represents the number of TSTs remaining. 0 means that there is no more TST. C673.1154 represents the last TST reference in the list retrieved. For the last reply value in C673.1154 of the first query is placed.
	ScrollingInformation *ActionDetailsTypeI `xml:"scrollingInformation,omitempty"`

	ApplicationError *ApplicationError `xml:"applicationError,omitempty"`

	FareList *FareList `xml:"fareList,omitempty"`
}

type ApplicationError struct {
	// Error information returned by ticketing application
	ApplicationErrorInfo *ApplicationErrorInformationType `xml:"applicationErrorInfo,omitempty"`

	// Description in free flow text of the error returned by ticketing application
	ErrorText *InteractiveFreeTextTypeI `xml:"errorText,omitempty"`
}

type FareList struct {
	// Information on TST creation method such as pricing rules and sales indicator.
	PricingInformation *PricingTicketingSubsequentTypeI `xml:"pricingInformation,omitempty"`

	// TST reference number. Ordering information is not relevant here.
	FareReference *ItemReferencesAndVersionsType_14908S `xml:"fareReference,omitempty"`

	// - Last date to ticket the fare - Creation date - Last modification date
	LastTktDate *StructuredDateTimeInformationType_14907S `xml:"lastTktDate,omitempty"`

	// Validating carrier identification.
	ValidatingCarrier *TransportIdentifierType `xml:"validatingCarrier,omitempty"`

	// Passenger/segment association of the TST.
	PaxSegReference *ReferenceInformationTypeI `xml:"paxSegReference,omitempty"`

	// Specify all fare data information (base fare, equivalent fare...)
	FareDataInformation *MonetaryInformationTypeI `xml:"fareDataInformation,omitempty"`

	TaxInformation *TaxInformation `xml:"taxInformation,omitempty"`

	// Banker's rates are used to convert amounts of the TST (base/equivalent).
	BankerRates *ConversionRateTypeI `xml:"bankerRates,omitempty"`

	PassengerInformation *PassengerInformation `xml:"passengerInformation,omitempty"`

	// Origin and destination of the fare. 1st C3225 occurence : origin city. 2nd C3225 occurence : destination city
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`

	SegmentInformation *SegmentInformation `xml:"segmentInformation,omitempty"`

	// Other fare information : - fare calculation  - payment restrictions. - mileage breakdown freeflow
	OtherPricingInfo *CodedAttributeType `xml:"otherPricingInfo,omitempty"`

	// TST status information such as TST confidentiality.
	StatusInformation *StatusTypeI `xml:"statusInformation,omitempty"`

	// Mentions the TST creation office ID and the agent sign in.
	OfficeDetails *UserIdentificationType `xml:"officeDetails,omitempty"`

	WarningInformation *WarningInformation `xml:"warningInformation,omitempty"`

	AutomaticReissueInfo *AutomaticReissueInfo `xml:"automaticReissueInfo,omitempty"`

	CarrierFeesGroup *CarrierFeesGroup `xml:"carrierFeesGroup,omitempty"`

	// contextual form of payment
	ContextualFop *FormOfPaymentTypeI `xml:"contextualFop,omitempty"`

	// Office Id where the pricing has been made. Implemented with Airline Ticketing Fees
	ContextualPointofSale *UserIdentificationType `xml:"contextualPointofSale,omitempty"`

	// convey the mileage information
	Mileage *AdditionalProductDetailsTypeI `xml:"mileage,omitempty"`
}

type TaxInformation struct {
	// Tax details
	TaxDetails *DutyTaxFeeDetailsTypeU `xml:"taxDetails,omitempty"`

	// Amount details. If the tax is a passenger facility charge (PFC) the detail of the airports related taxes is given here.
	AmountDetails *MonetaryInformationTypeI `xml:"amountDetails,omitempty"`
}

type PassengerInformation struct {
	// Penalty details specified in the request.
	PenDisInformation *DiscountAndPenaltyInformationTypeI_6128S `xml:"penDisInformation,omitempty"`

	// Reference of passengers that have a type code.
	PassengerReference *ReferenceInformationTypeI `xml:"passengerReference,omitempty"`
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

	// Base Fare information
	BaseFareInfo *MonetaryInformationTypeI_20897S `xml:"baseFareInfo,omitempty"`

	FirstDpiGroup *FirstDpiGroup `xml:"firstDpiGroup,omitempty"`

	SecondDpiGroup *SecondDpiGroup `xml:"secondDpiGroup,omitempty"`
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

	// Reissue Information
	ReissueInfo *MonetaryInformationTypeI_20897S `xml:"reissueInfo,omitempty"`

	// Old Tax Information
	OldTaxInfo *MonetaryInformationTypeI_20897S `xml:"oldTaxInfo,omitempty"`

	// Balance Reissue Information
	ReissueBalanceInfo *MonetaryInformationTypeI_20897S `xml:"reissueBalanceInfo,omitempty"`
}

type SecondDpiGroup struct {
	// Discount and penalty info.
	Penalty *DiscountAndPenaltyInformationTypeI `xml:"penalty,omitempty"`

	// Residual Value Information
	ResidualValueInfo *MonetaryInformationTypeI_20897S `xml:"residualValueInfo,omitempty"`

	// Old Tax Information
	OldTaxInfo *MonetaryInformationTypeI_20897S `xml:"oldTaxInfo,omitempty"`

	// Balance Issue Information
	IssueBalanceInfo *MonetaryInformationTypeI_20897S `xml:"issueBalanceInfo,omitempty"`
}

type CarrierFeesGroup struct {
	// Airline Ticketing Fees : OB
	CarrierFeeType *SelectionDetailsTypeI `xml:"carrierFeeType,omitempty"`

	CarrierFeeInfo *CarrierFeeInfo `xml:"carrierFeeInfo,omitempty"`
}

type CarrierFeeInfo struct {
	// contains the Carrier Fee subcode and the properties of the carrier fee:
	CarrierFeeSubcode *SpecificDataInformationTypeI `xml:"carrierFeeSubcode,omitempty"`

	// convey the commercial name of the fee
	CommercialName *InteractiveFreeTextTypeI_37809S `xml:"commercialName,omitempty"`

	// amount of the fee, taxes included
	FeeAmount *MonetaryInformationTypeI_37810S `xml:"feeAmount,omitempty"`

	// tax on the fee
	FeeTax *TaxTypeI `xml:"feeTax,omitempty"`
}

type SegmentInformation struct {
	// Connection information.
	ConnexInformation *ConnectionTypeI `xml:"connexInformation,omitempty"`

	// Details on open segments added to the price calculation. These open segments exist only in the fare calculated, they have no equivalent in the PNR itinerary. This segment gives also information on booking class for best buy transactions.
	SegDetails *TravelProductInformationTypeI `xml:"segDetails,omitempty"`

	// Fare basis information
	FareQualifier *FareQualifierDetailsTypeI `xml:"fareQualifier,omitempty"`

	// Validity information for the fare
	ValidityInformation *StructuredDateTimeInformationType `xml:"validityInformation,omitempty"`

	// baggage allowance information
	BagAllowanceInformation *ExcessBaggageTypeI `xml:"bagAllowanceInformation,omitempty"`

	// Reference of the segment associated to the group.
	SegmentReference *ReferenceInformationTypeI `xml:"segmentReference,omitempty"`

	// The segment order in the pricing response can be different than the one of the PNR itinerary (segments are reordered at price calculation time). This order information is conveyed by the sequence number. If this order information is not present then the order is by default the one of the PNR.
	SequenceInformation *ItemReferencesAndVersionsType `xml:"sequenceInformation,omitempty"`
}

type ActionDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ActionDetailsTypeI"`

	// Information on next list of TSTs to return.
	NextListInformation *ReferenceTypeI `xml:"nextListInformation,omitempty"`
}

type AdditionalFareQualifierDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A AdditionalFareQualifierDetailsTypeI"`

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
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A AdditionalProductDetailsTypeI"`

	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`
}

type ApplicationErrorDetailType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ApplicationErrorDetailType"`

	// Code identifying the data validation error condition.
	ApplicationErrorCode formats.AlphaNumericString_Length1To5 `xml:"applicationErrorCode,omitempty"`

	// Identification of a code list.
	CodeListQualifier formats.AlphaNumericString_Length1To3 `xml:"codeListQualifier,omitempty"`

	// Code identifying the agency responsible for a code list.
	CodeListResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"codeListResponsibleAgency,omitempty"`
}

type ApplicationErrorInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ApplicationErrorInformationType"`

	// Application error details.
	ApplicationErrorDetail *ApplicationErrorDetailType `xml:"applicationErrorDetail,omitempty"`
}

type BaggageDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A BaggageDetailsTypeI"`

	// Baggage allowance quantity (piece concept)
	BaggageQuantity formats.NumericInteger_Length1To1 `xml:"baggageQuantity,omitempty"`

	// Baggage allowance weight
	BaggageWeight formats.NumericDecimal_Length1To2 `xml:"baggageWeight,omitempty"`

	// Baggage allowance type (weight/number)
	BaggageType formats.AlphaNumericString_Length1To3 `xml:"baggageType,omitempty"`

	// Measurement unit for weighing baggage allowance
	MeasureUnit formats.AlphaNumericString_Length1To1 `xml:"measureUnit,omitempty"`
}

type CodedAttributeInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A CodedAttributeInformationType"`

	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType,omitempty"`

	// provides a description for the attribute
	AttributeDescription formats.AlphaNumericString_Length1To547 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A CodedAttributeType"`

	// provides details for the Attribute fare calculation or payment restriction or mileage breakdown freeflow.
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DummyNET struct{}

type CompanyIdentificationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A CompanyIdentificationTypeI"`

	// Carrier code
	CarrierCode formats.AlphaNumericString_Length1To3 `xml:"carrierCode,omitempty"`
}

type CompanyIdentificationTypeI_27116C struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A CompanyIdentificationTypeI_27116C"`

	// Code of the airline.
	CarrierCode formats.AlphaNumericString_Length1To3 `xml:"carrierCode,omitempty"`
}

type ConnectionDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ConnectionDetailsTypeI"`

	// Specify ARNK and surface segments  not included in the fare routing.
	RoutingInformation formats.AlphaNumericString_Length1To4 `xml:"routingInformation,omitempty"`

	// Type of connection for the flight
	ConnexType formats.AlphaNumericString_Length1To1 `xml:"connexType,omitempty"`
}

type ConnectionTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ConnectionTypeI"`

	// Connection details
	ConnecDetails *ConnectionDetailsTypeI `xml:"connecDetails,omitempty"`
}

type ConversionRateDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ConversionRateDetailsTypeI"`

	// Currency of the rate
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// Amount/percentage
	Amount formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`
}

type ConversionRateTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ConversionRateTypeI"`

	// First rate detail.
	FirstRateDetail *ConversionRateDetailsTypeI `xml:"firstRateDetail,omitempty"`

	// Second rate detail.
	SecondRateDetail *ConversionRateDetailsTypeI `xml:"secondRateDetail,omitempty"`
}

type CouponInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A CouponInformationDetailsTypeI"`

	// Coupon number
	CpnNumber formats.AlphaNumericString_Length1To6 `xml:"cpnNumber,omitempty"`
}

type CouponInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A CouponInformationTypeI"`

	// Details on coupon
	CouponDetails *CouponInformationDetailsTypeI `xml:"couponDetails,omitempty"`

	// Details on coupon
	OtherCouponDetails *CouponInformationDetailsTypeI `xml:"otherCouponDetails,omitempty"`
}

type DataInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DataInformationTypeI"`

	// fee attribute
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type DataTypeInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DataTypeInformationTypeI"`

	// fee subcode
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type DiscountAndPenaltyInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DiscountAndPenaltyInformationTypeI"`

	// Used to specify penalty information
	PenDisData *DiscountPenaltyMonetaryInformationTypeI_29792C `xml:"penDisData,omitempty"`
}

type DiscountAndPenaltyInformationTypeI_6128S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DiscountAndPenaltyInformationTypeI_6128S"`

	// Qualify the type of information.  Penalties are not passenger associated and are pure monetary information. Discount are passenger associated but only discount code is specified.
	InfoQualifier formats.AlphaNumericString_Length1To3 `xml:"infoQualifier,omitempty"`

	// Used to specify penalty information.
	PenDisData *DiscountPenaltyMonetaryInformationTypeI `xml:"penDisData,omitempty"`
}

type DiscountPenaltyInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DiscountPenaltyInformationTypeI"`

	// Discount off type.
	ZapOffType formats.AlphaNumericString_Length1To3 `xml:"zapOffType,omitempty"`

	// Discount amount
	ZapOffAmount formats.NumericDecimal_Length1To11 `xml:"zapOffAmount,omitempty"`

	// Discount percentage.
	ZapOffPercentage formats.NumericInteger_Length1To8 `xml:"zapOffPercentage,omitempty"`
}

type DiscountPenaltyMonetaryInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DiscountPenaltyMonetaryInformationTypeI"`

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
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DiscountPenaltyMonetaryInformationTypeI_29792C"`

	// The amount Type can be a percentage or an amount
	PenaltyQualifier formats.AlphaNumericString_Length1To3 `xml:"penaltyQualifier,omitempty"`

	// specify the value
	PenaltyAmount formats.NumericDecimal_Length1To18 `xml:"penaltyAmount,omitempty"`

	// penalty currency code
	PenaltyCurrency formats.AlphaNumericString_Length1To3 `xml:"penaltyCurrency,omitempty"`
}

type DutyTaxFeeAccountDetailTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DutyTaxFeeAccountDetailTypeU"`

	// Iso country of the tax
	IsoCountry formats.AlphaNumericString_Length1To3 `xml:"isoCountry,omitempty"`
}

type DutyTaxFeeDetailsTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DutyTaxFeeDetailsTypeU"`

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
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A DutyTaxFeeTypeDetailsTypeU"`

	// Tax type identifier
	TaxIdentifier formats.AlphaNumericString_Length1To3 `xml:"taxIdentifier,omitempty"`
}

type ExcessBaggageTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ExcessBaggageTypeI"`

	// Baggage allowance information details
	BagAllowanceDetails *BaggageDetailsTypeI `xml:"bagAllowanceDetails,omitempty"`
}

type FareQualifierDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A FareQualifierDetailsTypeI"`

	// Type of movement for this segment to take into account by Fare Quote to calculate the fare.
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`

	// Fare basis detail
	FareBasisDetails *AdditionalFareQualifierDetailsTypeI `xml:"fareBasisDetails,omitempty"`

	// Discount data for zap off to apply to price calculation.
	ZapOffDetails *DiscountPenaltyInformationTypeI `xml:"zapOffDetails,omitempty"`
}

type FormOfPaymentDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A FormOfPaymentDetailsTypeI"`

	// Type of Form Of Payment (Credit card, cash, check...).
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// amount to be charged on this form
	ChargedAmount formats.NumericDecimal_Length1To12 `xml:"chargedAmount,omitempty"`

	// It is mandatory if the Form of Payment at pricing was a credit card with a bin number. It may only occur for Credit Card Types.  It is composed of the first 6 bin numbers of the credit card.  Wildcards are not possible.
	CreditCardNumber formats.AlphaNumericString_Length1To20 `xml:"creditCardNumber,omitempty"`
}

type FormOfPaymentTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A FormOfPaymentTypeI"`

	// FORM OF PAYMENT DETAILS
	FormOfPayment *FormOfPaymentDetailsTypeI `xml:"formOfPayment,omitempty"`
}

type FreeTextQualificationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A FreeTextQualificationTypeI"`

	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A InteractiveFreeTextTypeI"`

	// Free flow text describing the error
	ErrorFreeText formats.AlphaNumericString_Length1To70 `xml:"errorFreeText,omitempty"`
}

type InteractiveFreeTextTypeI_37809S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A InteractiveFreeTextTypeI_37809S"`

	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`

	// commercial name of the fee suncode. 10 an
	FreeText formats.AlphaNumericString_Length1To10 `xml:"freeText,omitempty"`
}

type InteractiveFreeTextTypeI_6759S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A InteractiveFreeTextTypeI_6759S"`

	// Free flow text describing the error
	ErrorFreeText formats.AlphaNumericString_Length1To70 `xml:"errorFreeText,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ItemReferencesAndVersionsType"`

	// Identification details : order number
	SequenceSection *UniqueIdDescriptionType `xml:"sequenceSection,omitempty"`
}

type ItemReferencesAndVersionsType_14908S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ItemReferencesAndVersionsType_14908S"`

	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// Tattoo number (It is in fact the Tst Display Number)
	UniqueReference formats.NumericInteger_Length1To5 `xml:"uniqueReference,omitempty"`

	// Gives the TST ID number
	IDDescription *UniqueIdDescriptionType_26102C `xml:"iDDescription,omitempty"`
}

type LocationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A LocationTypeI"`

	// Code of the city.
	CityCode formats.AlphaNumericString_Length1To3 `xml:"cityCode,omitempty"`
}

type MileageTimeDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A MileageTimeDetailsTypeI"`

	// mileage total associated to the TST
	TotalMileage formats.NumericInteger_Length1To8 `xml:"totalMileage,omitempty"`
}

type MonetaryInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A MonetaryInformationDetailsTypeI"`

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
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A MonetaryInformationDetailsTypeI_37257C"`

	// Type qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// amount
	Amount formats.AlphaNumericString_Length1To35 `xml:"amount,omitempty"`

	// currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`
}

type MonetaryInformationDetailsTypeI_64230C struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A MonetaryInformationDetailsTypeI_64230C"`

	// qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// In case of exempted Fee, this data element does not contain an amount but a text: EXEMPTED.
	Amount formats.AlphaNumericString_Length1To35 `xml:"amount,omitempty"`

	// currency of the fee amount
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A MonetaryInformationTypeI"`

	// Main fare data infomation, can b thee base or  the total fare information which are mandatory  anyway
	FareDataMainInformation *MonetaryInformationDetailsTypeI `xml:"fareDataMainInformation,omitempty"`

	// Supplementary fare data information
	FareDataSupInformation *MonetaryInformationDetailsTypeI `xml:"fareDataSupInformation,omitempty"`
}

type MonetaryInformationTypeI_20897S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A MonetaryInformationTypeI_20897S"`

	// monetaryDetails
	MonetaryDetails *MonetaryInformationDetailsTypeI_37257C `xml:"monetaryDetails,omitempty"`

	OtherMonetaryDetails *MonetaryInformationDetailsTypeI_37257C `xml:"otherMonetaryDetails,omitempty"`
}

type MonetaryInformationTypeI_37810S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A MonetaryInformationTypeI_37810S"`

	// MON used for a single fee
	MonetaryDetails *MonetaryInformationDetailsTypeI_64230C `xml:"monetaryDetails,omitempty"`
}

type OriginAndDestinationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A OriginAndDestinationDetailsTypeI"`

	// Code of the city.
	CityCode formats.AlphaNumericString_Length1To4 `xml:"cityCode,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A OriginatorIdentificationDetailsTypeI"`

	// Agent Sign In
	InHouseIdentification1 formats.AlphaNumericString_Length1To2 `xml:"inHouseIdentification1,omitempty"`

	// Tst office ID : It's the TST creation office ID.
	InHouseIdentification2 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification2,omitempty"`
}

type PricingTicketingSubsequentTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A PricingTicketingSubsequentTypeI"`

	// Information on TST type.
	TstInformation *RateTariffClassInformationTypeI `xml:"tstInformation,omitempty"`

	// International sales indicator
	SalesIndicator formats.AlphaString_Length1To2 `xml:"salesIndicator,omitempty"`

	// Fare calculation mode indicator. This indicator specifies the type fare.
	Fcmi formats.AlphaNumericString_Length1To1 `xml:"fcmi,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ProductIdentificationDetailsTypeI"`

	// OPEN or AIR are the two identifications accepted.  OPEN means the segment described here is an open segment. AIR means that it is a valid AIR segment.
	Identification formats.AlphaNumericString_Length1To6 `xml:"identification,omitempty"`

	// Class of service to use in order to price the extra segment.
	ClassOfService formats.AlphaString_Length1To1 `xml:"classOfService,omitempty"`
}

type RateTariffClassInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A RateTariffClassInformationTypeI"`

	// Indicator qualifying the type of TST (basically manual or automatic)
	TstIndicator formats.AlphaNumericString_Length1To1 `xml:"tstIndicator,omitempty"`
}

type ReferenceInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ReferenceInformationTypeI"`

	// Passenger/segment/TST reference details
	RefDetails *ReferencingDetailsTypeI `xml:"refDetails,omitempty"`
}

type ReferenceTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ReferenceTypeI"`

	// In case of query specifies the number of TSTs to get in reply. In case of response specifies the number of TSTs remaining.
	RemainingInformation formats.NumericInteger_Length1To5 `xml:"remainingInformation,omitempty"`

	// In case of first query specifies the value of  this field in the last reply. In case of other queries specifies the last reference returned in the previous list. In case of reply specifies the last TST reference of the list. In case of last reply the value of this field set in the first query is sent.
	RemainingReference formats.AlphaNumericString_Length1To5 `xml:"remainingReference,omitempty"`
}

type ReferencingDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A ReferencingDetailsTypeI"`

	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A SelectionDetailsInformationTypeI"`

	Option formats.AlphaNumericString_Length1To2 `xml:"option,omitempty"`
}

type SelectionDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A SelectionDetailsTypeI"`

	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails,omitempty"`
}

type SpecificDataInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A SpecificDataInformationTypeI"`

	// Carrier fee code
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation,omitempty"`

	// Carrier fee application code  (NI, NR, CM, NC)
	DataInformation *DataInformationTypeI `xml:"dataInformation,omitempty"`
}

type StatusDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A StatusDetailsTypeI"`

	// Information on TST.
	TstFlag formats.AlphaNumericString_Length1To3 `xml:"tstFlag,omitempty"`
}

type StatusTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A StatusTypeI"`

	// Detail on the status of the TST.
	FirstStatusDetails *StatusDetailsTypeI `xml:"firstStatusDetails,omitempty"`

	// Other details on the status of the TST.
	OtherStatusDetails *StatusDetailsTypeI `xml:"otherStatusDetails,omitempty"`
}

type StructuredDateTimeInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A StructuredDateTimeInformationType"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_14907S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A StructuredDateTimeInformationType_14907S"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_26100C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A StructuredDateTimeType"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.NumericInteger_Length1To6 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type StructuredDateTimeType_26100C struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A StructuredDateTimeType_26100C"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.NumericInteger_Length1To6 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Minute_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type TaxDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A TaxDetailsTypeI"`

	Rate formats.AlphaNumericString_Length1To17 `xml:"rate,omitempty"`

	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// Details applicable to the tax applied to the Carrier Fee.  TX = Tax applies to fee
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type TaxTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A TaxTypeI"`

	// type of tax
	TaxCategory formats.AlphaNumericString_Length1To3 `xml:"taxCategory,omitempty"`

	FeeTaxDetails *TaxDetailsTypeI `xml:"feeTaxDetails,omitempty"`
}

type TicketNumberDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A TicketNumberDetailsTypeI"`

	// Ticket number
	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`

	// ticket type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type TicketNumberTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A TicketNumberTypeI"`

	// Details on the document
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails,omitempty"`
}

type TransportIdentifierType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A TransportIdentifierType"`

	// Information related to validating carrier.
	CarrierInformation *CompanyIdentificationTypeI `xml:"carrierInformation,omitempty"`
}

type TravelProductInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A TravelProductInformationTypeI"`

	// departure city code
	DepartureCity *LocationTypeI `xml:"departureCity,omitempty"`

	// arrival city code
	ArrivalCity *LocationTypeI `xml:"arrivalCity,omitempty"`

	// Airline detail information of the extra segment.
	AirlineDetail *CompanyIdentificationTypeI_27116C `xml:"airlineDetail,omitempty"`

	// Segment detail information.
	SegmentDetail *ProductIdentificationDetailsTypeI `xml:"segmentDetail,omitempty"`

	// Ticketing status for this segment. Relevant only in case of reply.
	TicketingStatus formats.AlphaNumericString_Length1To2 `xml:"ticketingStatus,omitempty"`
}

type UniqueIdDescriptionType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A UniqueIdDescriptionType"`

	// Number specifying the ordering information of the item described within a group.
	SequenceNumber formats.NumericInteger_Length1To3 `xml:"sequenceNumber,omitempty"`
}

type UniqueIdDescriptionType_26102C struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A UniqueIdDescriptionType_26102C"`

	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber formats.NumericInteger_Length1To11 `xml:"iDSequenceNumber,omitempty"`
}

type UserIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A UserIdentificationType"`

	// Originator Identification Details
	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification,omitempty"`
}
