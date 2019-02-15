package ticket_displaytst_reply

//import "encoding/xml"

type TicketDisplayTSTReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTRR_07_1_1A Ticket_DisplayTSTReply"`

	// Scrolling information used for long messages. C673.1050 represents the number of TSTs remaining. 0 means that there is no more TST. C673.1154 represents the last TST reference in the list retrieved. For the last reply value in C673.1154 of the first query is placed.
	ScrollingInformation *ActionDetailsTypeI `xml:"scrollingInformation,omitempty"`  // minOccurs="0"

	ApplicationError *ApplicationError `xml:"applicationError,omitempty"`  // minOccurs="0"

	FareList []*FareList `xml:"fareList,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type ApplicationError struct {

	// Error information returned by ticketing application
	ApplicationErrorInfo *ApplicationErrorInformationType `xml:"applicationErrorInfo"`

	// Description in free flow text of the error returned by ticketing application
	ErrorText *InteractiveFreeTextTypeI `xml:"errorText,omitempty"`  // minOccurs="0"
}

type FareList struct {

	// Information on TST creation method such as pricing rules and sales indicator.
	PricingInformation *PricingTicketingSubsequentTypeI `xml:"pricingInformation"`

	// TST reference number. Ordering information is not relevant here.
	FareReference *ItemReferencesAndVersionsType_14908S `xml:"fareReference"`

	// - Last date to ticket the fare - Creation date - Last modification date
	LastTktDate []*StructuredDateTimeInformationType_14907S `xml:"lastTktDate,omitempty"`  // minOccurs="0" maxOccurs="3"

	// Validating carrier identification.
	ValidatingCarrier *TransportIdentifierType `xml:"validatingCarrier,omitempty"`  // minOccurs="0"

	// Passenger/segment association of the TST.
	PaxSegReference *ReferenceInformationTypeI `xml:"paxSegReference"`

	// Specify all fare data information (base fare, equivalent fare...)
	FareDataInformation *MonetaryInformationTypeI `xml:"fareDataInformation,omitempty"`  // minOccurs="0"

	TaxInformation []*TaxInformation `xml:"taxInformation,omitempty"`  // minOccurs="0" maxOccurs="120"

	// Banker's rates are used to convert amounts of the TST (base/equivalent).
	BankerRates *ConversionRateTypeI `xml:"bankerRates,omitempty"`  // minOccurs="0"

	PassengerInformation []*PassengerInformation `xml:"passengerInformation,omitempty"`  // minOccurs="0" maxOccurs="99"

	// Origin and destination of the fare. 1st C3225 occurence : origin city. 2nd C3225 occurence : destination city
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`  // minOccurs="0"

	SegmentInformation []*SegmentInformation `xml:"segmentInformation,omitempty"`  // minOccurs="0" maxOccurs="96"

	// Other fare information : - fare calculation  - payment restrictions. - mileage breakdown freeflow
	OtherPricingInfo []*CodedAttributeType `xml:"otherPricingInfo,omitempty"`  // minOccurs="0" maxOccurs="2"

	// TST status information such as TST confidentiality.
	StatusInformation *StatusTypeI `xml:"statusInformation,omitempty"`  // minOccurs="0"

	// Mentions the TST creation office ID and the agent sign in.
	OfficeDetails *UserIdentificationType `xml:"officeDetails,omitempty"`  // minOccurs="0"

	WarningInformation []*WarningInformation `xml:"warningInformation,omitempty"`  // minOccurs="0" maxOccurs="99"

	AutomaticReissueInfo *AutomaticReissueInfo `xml:"automaticReissueInfo,omitempty"`  // minOccurs="0"

	CarrierFeesGroup []*CarrierFeesGroup `xml:"carrierFeesGroup,omitempty"`  // minOccurs="0" maxOccurs="99"

	// contextual form of payment
	ContextualFop *FormOfPaymentTypeI `xml:"contextualFop,omitempty"`  // minOccurs="0"

	// Office Id where the pricing has been made. Implemented with Airline Ticketing Fees
	ContextualPointofSale *UserIdentificationType `xml:"contextualPointofSale,omitempty"`  // minOccurs="0"

	// convey the mileage information
	Mileage *AdditionalProductDetailsTypeI `xml:"mileage,omitempty"`  // minOccurs="0"
}

type TaxInformation struct {

	// Tax details
	TaxDetails *DutyTaxFeeDetailsTypeU `xml:"taxDetails"`

	// Amount details. If the tax is a passenger facility charge (PFC) the detail of the airports related taxes is given here.
	AmountDetails *MonetaryInformationTypeI `xml:"amountDetails,omitempty"`  // minOccurs="0"
}

type PassengerInformation struct {

	// Penalty details specified in the request.
	PenDisInformation *DiscountAndPenaltyInformationTypeI_6128S `xml:"penDisInformation"`

	// Reference of passengers that have a type code.
	PassengerReference *ReferenceInformationTypeI `xml:"passengerReference,omitempty"`  // minOccurs="0"
}

type SegmentInformation struct {

	// Connection information.
	ConnexInformation *ConnectionTypeI `xml:"connexInformation"`

	// Details on open segments added to the price calculation. These open segments exist only in the fare calculated, they have no equivalent in the PNR itinerary. This segment gives also information on booking class for best buy transactions.
	SegDetails *TravelProductInformationTypeI `xml:"segDetails,omitempty"`  // minOccurs="0"

	// Fare basis information
	FareQualifier *FareQualifierDetailsTypeI `xml:"fareQualifier,omitempty"`  // minOccurs="0"

	// Validity information for the fare
	ValidityInformation []*StructuredDateTimeInformationType `xml:"validityInformation,omitempty"`  // minOccurs="0" maxOccurs="2"

	// baggage allowance information
	BagAllowanceInformation *ExcessBaggageTypeI `xml:"bagAllowanceInformation,omitempty"`  // minOccurs="0"

	// Reference of the segment associated to the group.
	SegmentReference *ReferenceInformationTypeI `xml:"segmentReference,omitempty"`  // minOccurs="0"

	// The segment order in the pricing response can be different than the one of the PNR itinerary (segments are reordered at price calculation time). This order information is conveyed by the sequence number. If this order information is not present then the order is by default the one of the PNR.
	SequenceInformation *ItemReferencesAndVersionsType `xml:"sequenceInformation,omitempty"`  // minOccurs="0"
}

type WarningInformation struct {

	// Fare warning information code.
	WarningCode *ApplicationErrorInformationType `xml:"warningCode"`

	// Description in free flow text of the warning concerning the fare.
	WarningText *InteractiveFreeTextTypeI_6759S `xml:"warningText,omitempty"`  // minOccurs="0"
}

type AutomaticReissueInfo struct {

	// This segment contains the original ticket number.
	TicketInfo *TicketNumberTypeI `xml:"ticketInfo"`

	// This segment contains the coupon number (in absolute) corresponding to the first coupon for use from the last flawn segment.
	CouponInfo *CouponInformationTypeI `xml:"couponInfo"`

	PaperCouponRange *PaperCouponRange `xml:"paperCouponRange,omitempty"`  // minOccurs="0"

	// Base Fare information
	BaseFareInfo *MonetaryInformationTypeI_20897S `xml:"baseFareInfo"`

	FirstDpiGroup *FirstDpiGroup `xml:"firstDpiGroup"`

	SecondDpiGroup *SecondDpiGroup `xml:"secondDpiGroup"`
}

type PaperCouponRange struct {

	// This segment contains the original ticket number.
	TicketInfo *TicketNumberTypeI `xml:"ticketInfo"`

	// This segment contains the coupon number (in absolute) corresponding to the first coupon for use from the last flawn segment.
	CouponInfo *CouponInformationTypeI `xml:"couponInfo"`
}

type FirstDpiGroup struct {

	// Penalty amount in reissue currency
	ReIssuePenalty *DiscountAndPenaltyInformationTypeI `xml:"reIssuePenalty"`

	// Reissue Information
	ReissueInfo *MonetaryInformationTypeI_20897S `xml:"reissueInfo"`

	// Old Tax Information
	OldTaxInfo *MonetaryInformationTypeI_20897S `xml:"oldTaxInfo"`

	// Balance Reissue Information
	ReissueBalanceInfo *MonetaryInformationTypeI_20897S `xml:"reissueBalanceInfo"`
}

type SecondDpiGroup struct {

	// Discount and penalty info.
	Penalty *DiscountAndPenaltyInformationTypeI `xml:"penalty"`

	// Residual Value Information
	ResidualValueInfo *MonetaryInformationTypeI_20897S `xml:"residualValueInfo"`

	// Old Tax Information
	OldTaxInfo *MonetaryInformationTypeI_20897S `xml:"oldTaxInfo"`

	// Balance Issue Information
	IssueBalanceInfo *MonetaryInformationTypeI_20897S `xml:"issueBalanceInfo"`
}

type CarrierFeesGroup struct {

	// Airline Ticketing Fees : OB
	CarrierFeeType *SelectionDetailsTypeI `xml:"carrierFeeType"`

	CarrierFeeInfo []*CarrierFeeInfo `xml:"carrierFeeInfo,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type CarrierFeeInfo struct {

	// contains the Carrier Fee subcode and the properties of the carrier fee:
	CarrierFeeSubcode *SpecificDataInformationTypeI `xml:"carrierFeeSubcode"`

	// convey the commercial name of the fee
	CommercialName *InteractiveFreeTextTypeI_37809S `xml:"commercialName"`

	// amount of the fee, taxes included
	FeeAmount *MonetaryInformationTypeI_37810S `xml:"feeAmount"`

	// tax on the fee
	FeeTax []*TaxTypeI `xml:"feeTax,omitempty"`  // minOccurs="0" maxOccurs="99"
}

//
// Complex structs
//

type ActionDetailsTypeI struct {

	// Information on next list of TSTs to return.
	NextListInformation *ReferenceTypeI `xml:"nextListInformation"`
}

type AdditionalFareQualifierDetailsTypeI struct {

	// Primary code of the fare basis. This is not a codeset but a free flow text field.
	PrimaryCode string `xml:"primaryCode,omitempty"`  // minOccurs="0"

	// Fare basis code of the fare basis. This is not a codeset but a free flow text field.
	FareBasisCode string `xml:"fareBasisCode,omitempty"`  // minOccurs="0"

	// Ticket designator of the fare basis
	TicketDesignator string `xml:"ticketDesignator,omitempty"`  // minOccurs="0"

	// For any query : discount ticket designator to be assigned by Fare Quote server. For any response : priced PTCs
	DiscTktDesignator string `xml:"discTktDesignator,omitempty"`  // minOccurs="0"
}

type AdditionalProductDetailsTypeI struct {

	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails"`
}

type ApplicationErrorDetailType struct {

	// Code identifying the data validation error condition.
	ApplicationErrorCode string `xml:"applicationErrorCode"`

	// Identification of a code list.
	CodeListQualifier string `xml:"codeListQualifier,omitempty"`  // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	CodeListResponsibleAgency string `xml:"codeListResponsibleAgency,omitempty"`  // minOccurs="0"
}

type ApplicationErrorInformationType struct {

	// Application error details.
	ApplicationErrorDetail *ApplicationErrorDetailType `xml:"applicationErrorDetail"`
}

type BaggageDetailsTypeI struct {

	// Baggage allowance quantity (piece concept)
	BaggageQuantity *int32 `xml:"baggageQuantity,omitempty"`  // minOccurs="0"

	// Baggage allowance weight
	BaggageWeight *float64 `xml:"baggageWeight,omitempty"`  // minOccurs="0"

	// Baggage allowance type (weight/number)
	BaggageType string `xml:"baggageType,omitempty"`  // minOccurs="0"

	// Measurement unit for weighing baggage allowance
	MeasureUnit string `xml:"measureUnit,omitempty"`  // minOccurs="0"
}

type CodedAttributeInformationType struct {

	// provides the attribute Type
	AttributeType string `xml:"attributeType"`

	// provides a description for the attribute
	AttributeDescription string `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeType struct {

	// provides details for the Attribute fare calculation or payment restriction or mileage breakdown freeflow.
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"`  // maxOccurs="3"
}

type CompanyIdentificationTypeI struct {

	// Carrier code
	CarrierCode string `xml:"carrierCode"`
}

type CompanyIdentificationTypeI_27116C struct {

	// Code of the airline.
	CarrierCode string `xml:"carrierCode,omitempty"`  // minOccurs="0"
}

type ConnectionDetailsTypeI struct {

	// Specify ARNK and surface segments  not included in the fare routing.
	RoutingInformation string `xml:"routingInformation,omitempty"`  // minOccurs="0"

	// Type of connection for the flight
	ConnexType string `xml:"connexType,omitempty"`  // minOccurs="0"
}

type ConnectionTypeI struct {

	// Connection details
	ConnecDetails *ConnectionDetailsTypeI `xml:"connecDetails"`
}

type ConversionRateDetailsTypeI struct {

	// Currency of the rate
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// Amount/percentage
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"
}

type ConversionRateTypeI struct {

	// First rate detail.
	FirstRateDetail *ConversionRateDetailsTypeI `xml:"firstRateDetail"`

	// Second rate detail.
	SecondRateDetail *ConversionRateDetailsTypeI `xml:"secondRateDetail,omitempty"`  // minOccurs="0"
}

type CouponInformationDetailsTypeI struct {

	// Coupon number
	CpnNumber string `xml:"cpnNumber"`
}

type CouponInformationTypeI struct {

	// Details on coupon
	CouponDetails *CouponInformationDetailsTypeI `xml:"couponDetails"`

	// Details on coupon
	OtherCouponDetails []*CouponInformationDetailsTypeI `xml:"otherCouponDetails,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type DataInformationTypeI struct {

	// fee attribute
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"
}

type DataTypeInformationTypeI struct {

	// fee subcode
	Type string `xml:"type"`
}

type DiscountAndPenaltyInformationTypeI struct {

	// Used to specify penalty information
	PenDisData *DiscountPenaltyMonetaryInformationTypeI_29792C `xml:"penDisData,omitempty"`  // minOccurs="0"
}

type DiscountAndPenaltyInformationTypeI_6128S struct {

	// Qualify the type of information.  Penalties are not passenger associated and are pure monetary information. Discount are passenger associated but only discount code is specified.
	InfoQualifier string `xml:"infoQualifier,omitempty"`  // minOccurs="0"

	// Used to specify penalty information.
	PenDisData []*DiscountPenaltyMonetaryInformationTypeI `xml:"penDisData,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type DiscountPenaltyInformationTypeI struct {

	// Discount off type.
	ZapOffType string `xml:"zapOffType"`

	// Discount amount
	ZapOffAmount *float64 `xml:"zapOffAmount,omitempty"`  // minOccurs="0"

	// Discount percentage.
	ZapOffPercentage *int32 `xml:"zapOffPercentage,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyMonetaryInformationTypeI struct {

	// Type of penalty.
	PenaltyType string `xml:"penaltyType,omitempty"`  // minOccurs="0"

	// The penalty amount can be described differently: amount/percentage.
	PenaltyQualifier string `xml:"penaltyQualifier,omitempty"`  // minOccurs="0"

	// Amount of the penalty.
	PenaltyAmount *float64 `xml:"penaltyAmount,omitempty"`  // minOccurs="0"

	// This discount code is defined by the airlines. This cannot be coded as airlines might apply any combination of letters for their discounts.
	DiscountCode string `xml:"discountCode,omitempty"`  // minOccurs="0"

	// Penalty currency code.
	PenaltyCurrency string `xml:"penaltyCurrency,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyMonetaryInformationTypeI_29792C struct {

	// The amount Type can be a percentage or an amount
	PenaltyQualifier string `xml:"penaltyQualifier,omitempty"`  // minOccurs="0"

	// specify the value
	PenaltyAmount *float64 `xml:"penaltyAmount,omitempty"`  // minOccurs="0"

	// penalty currency code
	PenaltyCurrency string `xml:"penaltyCurrency,omitempty"`  // minOccurs="0"
}

type DutyTaxFeeAccountDetailTypeU struct {

	// Iso country of the tax
	IsoCountry string `xml:"isoCountry"`
}

type DutyTaxFeeDetailsTypeU struct {

	// Tax data qualifier
	TaxQualifier string `xml:"taxQualifier"`

	// Tax type identification
	TaxIdentification *DutyTaxFeeTypeDetailsTypeU `xml:"taxIdentification"`

	// Type of the tax
	TaxType *DutyTaxFeeAccountDetailTypeU `xml:"taxType,omitempty"`  // minOccurs="0"

	// Nature of the tax
	TaxNature string `xml:"taxNature,omitempty"`  // minOccurs="0"

	// Exempt tax indicator. If an tax is Exempted no amount is provided for this tax.
	TaxExempt string `xml:"taxExempt,omitempty"`  // minOccurs="0"
}

type DutyTaxFeeTypeDetailsTypeU struct {

	// Tax type identifier
	TaxIdentifier string `xml:"taxIdentifier"`
}

type ExcessBaggageTypeI struct {

	// Baggage allowance information details
	BagAllowanceDetails *BaggageDetailsTypeI `xml:"bagAllowanceDetails"`
}

type FareQualifierDetailsTypeI struct {

	// Type of movement for this segment to take into account by Fare Quote to calculate the fare.
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	// Fare basis detail
	FareBasisDetails *AdditionalFareQualifierDetailsTypeI `xml:"fareBasisDetails,omitempty"`  // minOccurs="0"

	// Discount data for zap off to apply to price calculation.
	ZapOffDetails *DiscountPenaltyInformationTypeI `xml:"zapOffDetails,omitempty"`  // minOccurs="0"
}

type FormOfPaymentDetailsTypeI struct {

	// Type of Form Of Payment (Credit card, cash, check...).
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// amount to be charged on this form
	ChargedAmount *float64 `xml:"chargedAmount,omitempty"`  // minOccurs="0"

	// It is mandatory if the Form of Payment at pricing was a credit card with a bin number. It may only occur for Credit Card Types.  It is composed of the first 6 bin numbers of the credit card.  Wildcards are not possible.
	CreditCardNumber string `xml:"creditCardNumber,omitempty"`  // minOccurs="0"
}

type FormOfPaymentTypeI struct {

	// FORM OF PAYMENT DETAILS
	FormOfPayment []*FormOfPaymentDetailsTypeI `xml:"formOfPayment,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type FreeTextQualificationTypeI struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`
}

type InteractiveFreeTextTypeI struct {

	// Free flow text describing the error
	ErrorFreeText string `xml:"errorFreeText"`
}

type InteractiveFreeTextTypeI_37809S struct {

	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// commercial name of the fee suncode. 10 an
	FreeText string `xml:"freeText"`
}

type InteractiveFreeTextTypeI_6759S struct {

	// Free flow text describing the error
	ErrorFreeText string `xml:"errorFreeText,omitempty"`  // minOccurs="0"
}

type ItemReferencesAndVersionsType struct {

	// Identification details : order number
	SequenceSection *UniqueIdDescriptionType `xml:"sequenceSection,omitempty"`  // minOccurs="0"
}

type ItemReferencesAndVersionsType_14908S struct {

	// qualifies the type of the reference used. Code set to define
	ReferenceType string `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Tattoo number (It is in fact the Tst Display Number)
	UniqueReference *int32 `xml:"uniqueReference,omitempty"`  // minOccurs="0"

	// Gives the TST ID number
	IDDescription *UniqueIdDescriptionType_26102C `xml:"iDDescription,omitempty"`  // minOccurs="0"
}

type LocationTypeI struct {

	// Code of the city.
	CityCode string `xml:"cityCode,omitempty"`  // minOccurs="0"
}

type MileageTimeDetailsTypeI struct {

	// mileage total associated to the TST
	TotalMileage int32 `xml:"totalMileage"`
}

type MonetaryInformationDetailsTypeI struct {

	// Qualify the type of fare defined in this composite
	FareDataQualifier string `xml:"fareDataQualifier"`

	// Fare data amount
	FareAmount string `xml:"fareAmount,omitempty"`  // minOccurs="0"

	// Fare data currency code
	FareCurrency string `xml:"fareCurrency,omitempty"`  // minOccurs="0"

	// Location of the fare data (PFCs specific)
	FareLocation string `xml:"fareLocation,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_37257C struct {

	// Type qualifier
	TypeQualifier string `xml:"typeQualifier"`

	// amount
	Amount string `xml:"amount"`

	// currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// location
	Location string `xml:"location,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_64230C struct {

	// qualifier
	TypeQualifier string `xml:"typeQualifier"`

	// In case of exempted Fee, this data element does not contain an amount but a text: EXEMPTED.
	Amount string `xml:"amount"`

	// currency of the fee amount
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationTypeI struct {

	// Main fare data infomation, can b thee base or  the total fare information which are mandatory  anyway
	FareDataMainInformation *MonetaryInformationDetailsTypeI `xml:"fareDataMainInformation"`

	// Supplementary fare data information
	FareDataSupInformation []*MonetaryInformationDetailsTypeI `xml:"fareDataSupInformation,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type MonetaryInformationTypeI_20897S struct {

	// monetaryDetails
	MonetaryDetails *MonetaryInformationDetailsTypeI_37257C `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsTypeI_37257C `xml:"otherMonetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type MonetaryInformationTypeI_37810S struct {

	// MON used for a single fee
	MonetaryDetails *MonetaryInformationDetailsTypeI_64230C `xml:"monetaryDetails"`
}

type OriginAndDestinationDetailsTypeI struct {

	// Code of the city.
	CityCode []string `xml:"cityCode"`  // maxOccurs="2"
}

type OriginatorIdentificationDetailsTypeI struct {

	// Agent Sign In
	InHouseIdentification1 string `xml:"inHouseIdentification1,omitempty"`  // minOccurs="0"

	// Tst office ID : It's the TST creation office ID.
	InHouseIdentification2 string `xml:"inHouseIdentification2,omitempty"`  // minOccurs="0"
}

type PricingTicketingSubsequentTypeI struct {

	// Information on TST type.
	TstInformation *RateTariffClassInformationTypeI `xml:"tstInformation"`

	// International sales indicator
	SalesIndicator string `xml:"salesIndicator,omitempty"`  // minOccurs="0"

	// Fare calculation mode indicator. This indicator specifies the type fare.
	Fcmi string `xml:"fcmi"`
}

type ProductIdentificationDetailsTypeI struct {

	// OPEN or AIR are the two identifications accepted.  OPEN means the segment described here is an open segment. AIR means that it is a valid AIR segment.
	Identification string `xml:"identification"`

	// Class of service to use in order to price the extra segment.
	ClassOfService string `xml:"classOfService,omitempty"`  // minOccurs="0"
}

type RateTariffClassInformationTypeI struct {

	// Indicator qualifying the type of TST (basically manual or automatic)
	TstIndicator string `xml:"tstIndicator"`
}

type ReferenceInformationTypeI struct {

	// Passenger/segment/TST reference details
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferenceTypeI struct {

	// In case of query specifies the number of TSTs to get in reply. In case of response specifies the number of TSTs remaining.
	RemainingInformation *int32 `xml:"remainingInformation,omitempty"`  // minOccurs="0"

	// In case of first query specifies the value of  this field in the last reply. In case of other queries specifies the last reference returned in the previous list. In case of reply specifies the last TST reference of the list. In case of last reply the value of this field set in the first query is sent.
	RemainingReference string `xml:"remainingReference,omitempty"`  // minOccurs="0"
}

type ReferencingDetailsTypeI struct {

	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier string `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber *int32 `xml:"refNumber,omitempty"`  // minOccurs="0"
}

type SelectionDetailsInformationTypeI struct {

	// Format limitations: an..2
	Option string `xml:"option"`
}

type SelectionDetailsTypeI struct {

	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails"`
}

type SpecificDataInformationTypeI struct {

	// Carrier fee code
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`

	// Carrier fee application code  (NI, NR, CM, NC)
	DataInformation []*DataInformationTypeI `xml:"dataInformation,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type StatusDetailsTypeI struct {

	// Information on TST.
	TstFlag string `xml:"tstFlag"`
}

type StatusTypeI struct {

	// Detail on the status of the TST.
	FirstStatusDetails *StatusDetailsTypeI `xml:"firstStatusDetails"`

	// Other details on the status of the TST.
	OtherStatusDetails []*StatusDetailsTypeI `xml:"otherStatusDetails,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type StructuredDateTimeInformationType struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeInformationType_14907S struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic string `xml:"businessSemantic"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_26100C `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year *int32 `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month string `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day string `xml:"day,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_26100C struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year *int32 `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month *int32 `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day string `xml:"day,omitempty"`  // minOccurs="0"
}

type TaxDetailsTypeI struct {

	// Format limitations: an..17
	Rate string `xml:"rate"`

	// Format limitations: an..3
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// Details applicable to the tax applied to the Carrier Fee.  TX = Tax applies to fee
	Type []string `xml:"type"`  // maxOccurs="99"
}

type TaxTypeI struct {

	// type of tax
	TaxCategory string `xml:"taxCategory,omitempty"`  // minOccurs="0"

	FeeTaxDetails []*TaxDetailsTypeI `xml:"feeTaxDetails"`  // maxOccurs="99"
}

type TicketNumberDetailsTypeI struct {

	// Ticket number
	Number string `xml:"number"`

	// ticket type
	Type string `xml:"type,omitempty"`  // minOccurs="0"
}

type TicketNumberTypeI struct {

	// Details on the document
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`
}

type TransportIdentifierType struct {

	// Information related to validating carrier.
	CarrierInformation *CompanyIdentificationTypeI `xml:"carrierInformation"`
}

type TravelProductInformationTypeI struct {

	// departure city code
	DepartureCity *LocationTypeI `xml:"departureCity,omitempty"`  // minOccurs="0"

	// arrival city code
	ArrivalCity *LocationTypeI `xml:"arrivalCity,omitempty"`  // minOccurs="0"

	// Airline detail information of the extra segment.
	AirlineDetail *CompanyIdentificationTypeI_27116C `xml:"airlineDetail,omitempty"`  // minOccurs="0"

	// Segment detail information.
	SegmentDetail *ProductIdentificationDetailsTypeI `xml:"segmentDetail,omitempty"`  // minOccurs="0"

	// Ticketing status for this segment. Relevant only in case of reply.
	TicketingStatus string `xml:"ticketingStatus,omitempty"`  // minOccurs="0"
}

type UniqueIdDescriptionType struct {

	// Number specifying the ordering information of the item described within a group.
	SequenceNumber *int32 `xml:"sequenceNumber,omitempty"`  // minOccurs="0"
}

type UniqueIdDescriptionType_26102C struct {

	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber int32 `xml:"iDSequenceNumber"`
}

type UserIdentificationType struct {

	// Originator Identification Details
	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification"`
}
