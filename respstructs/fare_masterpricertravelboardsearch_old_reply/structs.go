package fare_masterpricertravelboardsearch_old_reply

//import "encoding/xml"

type FareMasterPricerTravelBoardSearchOldReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_12_3_1A Fare_MasterPricerTravelBoardSearchReply"`

	// Gives status about reply : type of process, region , CPU etc..
	ReplyStatus *StatusType `xml:"replyStatus,omitempty"`  // minOccurs="0"

	ErrorMessage *ErrorMessage `xml:"errorMessage,omitempty"`  // minOccurs="0"

	// Specifies the currency used for pricing.
	ConversionRate *ConversionRateTypeI `xml:"conversionRate,omitempty"`  // minOccurs="0"

	// Solution Family
	SolutionFamily []*FareInformationType `xml:"solutionFamily,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Details of the fare families processed
	FamilyInformation []*FareFamilyType `xml:"familyInformation,omitempty"`  // minOccurs="0" maxOccurs="200"

	AmountInfoForAllPax *AmountInfoForAllPax `xml:"amountInfoForAllPax,omitempty"`  // minOccurs="0"

	AmountInfoPerPax []*AmountInfoPerPax `xml:"amountInfoPerPax,omitempty"`  // minOccurs="0" maxOccurs="20"

	FeeDetails []*FeeDetails `xml:"feeDetails,omitempty"`  // minOccurs="0" maxOccurs="2099"

	// Free text identifying an airline in a code share.
	CompanyIdText []*CompanyIdentificationTextType `xml:"companyIdText,omitempty"`  // minOccurs="0" maxOccurs="5000"

	OfficeIdDetails []*OfficeIdDetails `xml:"officeIdDetails,omitempty"`  // minOccurs="0" maxOccurs="20"

	FlightIndex []*FlightIndex `xml:"flightIndex,omitempty"`  // minOccurs="0" maxOccurs="6"

	Recommendation []*Recommendation `xml:"recommendation,omitempty"`  // minOccurs="0" maxOccurs="100000"

	WarningInfo []*WarningInfo `xml:"warningInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	GlobalInformation []*GlobalInformation `xml:"globalInformation,omitempty"`  // minOccurs="0" maxOccurs="9"

	ServiceFeesGrp []*ServiceFeesGrp `xml:"serviceFeesGrp,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type ErrorMessage struct {

	// Application error details.
	ApplicationError *ApplicationErrorInformationType_78543S `xml:"applicationError"`

	// Type of error message and free text
	ErrorMessageText *InteractiveFreeTextType_78544S `xml:"errorMessageText,omitempty"`  // minOccurs="0"
}

type AmountInfoForAllPax struct {

	// Itinerary amounts for all passengers
	ItineraryAmounts *MonetaryInformationType_137835S `xml:"itineraryAmounts"`

	AmountsPerSgt []*AmountsPerSgt `xml:"amountsPerSgt,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type AmountsPerSgt struct {

	// Requested segment reference
	SgtRef *ReferenceInfoType_133176S `xml:"sgtRef"`

	// Amounts : Issue total amount, issue taxes amount, non refundable taxes amount
	Amounts *MonetaryInformationType_137835S `xml:"amounts,omitempty"`  // minOccurs="0"
}

type AmountInfoPerPax struct {

	// Passenger references
	PaxRef *SpecificTravellerType `xml:"paxRef"`

	// Passenger attributes : Infant indicator
	PaxAttributes *FareInformationType_80868S `xml:"paxAttributes,omitempty"`  // minOccurs="0"

	// Itinerary amounts information
	ItineraryAmounts *MonetaryInformationType_137835S `xml:"itineraryAmounts"`

	AmountsPerSgt []*AmountsPerSgt `xml:"amountsPerSgt,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FeeDetails struct {

	// Fee/Reduction Reference number.
	FeeReference *ItemReferencesAndVersionsType_78564S `xml:"feeReference"`

	// Fee/Reduction information.
	FeeInformation *DiscountAndPenaltyInformationType `xml:"feeInformation,omitempty"`  // minOccurs="0"

	// Fee/Reduction parameters.
	FeeParameters *AttributeType_78561S `xml:"feeParameters,omitempty"`  // minOccurs="0"

	// To specify conversion rate details
	ConvertedOrOriginalInfo *ConversionRateTypeI_78562S `xml:"convertedOrOriginalInfo,omitempty"`  // minOccurs="0"
}

type OfficeIdDetails struct {

	// Office Id Information
	OfficeIdInformation *UserIdentificationType `xml:"officeIdInformation"`

	// Office Id Reference Number
	OfficeIdReference *ItemReferencesAndVersionsType_78536S `xml:"officeIdReference"`
}

type FlightIndex struct {

	// Indicates references and details about requested segments
	RequestedSegmentRef *OriginAndDestinationRequestType `xml:"requestedSegmentRef"`

	GroupOfFlights []*GroupOfFlights `xml:"groupOfFlights"`  // maxOccurs="100000"
}

type GroupOfFlights struct {

	// To indicate parameters for proposed flight group.
	PropFlightGrDetail *ProposedSegmentType `xml:"propFlightGrDetail"`

	FlightDetails []*FlightDetails `xml:"flightDetails"`  // maxOccurs="4"
}

type FlightDetails struct {

	// Specification of details on the flight and posting availability
	FlightInformation *TravelProductType `xml:"flightInformation"`

	// Details on Flight date, time and location of technical stop or change of gauge
	TechnicalStop []*DateAndTimeInformationType `xml:"technicalStop,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Code Share Agreement description for current flight.
	CommercialAgreement *CommercialAgreementsType `xml:"commercialAgreement,omitempty"`  // minOccurs="0"
}

type Recommendation struct {

	// Specification of the item number
	ItemNumber *ItemNumberType_133453S `xml:"itemNumber"`

	// To describe type of recommendation
	WarningMessage []*InteractiveFreeTextType_78544S `xml:"warningMessage,omitempty"`  // minOccurs="0" maxOccurs="4"

	// Indicates the Fare family reference.
	FareFamilyRef *ReferenceInfoType_133176S `xml:"fareFamilyRef,omitempty"`  // minOccurs="0"

	// Recommendation Price and Taxes.
	RecPriceInfo *MonetaryInformationType_134806S `xml:"recPriceInfo"`

	// Mini rules
	MiniRule []*MiniRulesType `xml:"miniRule,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Indicates reference of Flight or the fee reference valid for all pax (usage:start with the 1 possible Fee reference, then provide the segments references)
	SegmentFlightRef []*ReferenceInfoType `xml:"segmentFlightRef,omitempty"`  // minOccurs="0" maxOccurs="100009"

	RecommandationSegmentsFareDetails []*RecommandationSegmentsFareDetails `xml:"recommandationSegmentsFareDetails,omitempty"`  // minOccurs="0" maxOccurs="6"

	PaxFareProduct []*PaxFareProduct `xml:"paxFareProduct"`  // maxOccurs="10"

	SpecificRecDetails []*SpecificRecDetails `xml:"specificRecDetails,omitempty"`  // minOccurs="0" maxOccurs="100000"
}

type RecommandationSegmentsFareDetails struct {

	// Reference and details about requested segments.
	RecommendationSegRef *OriginAndDestinationRequestType `xml:"recommendationSegRef"`

	// Amounts per requested segment.
	SegmentMonetaryInformation *MonetaryInformationType_137835S `xml:"segmentMonetaryInformation,omitempty"`  // minOccurs="0"
}

type PaxFareProduct struct {

	// Passenger Fare Details.
	PaxFareDetail *PricingTicketingSubsequentType_134837S `xml:"paxFareDetail"`

	// Indicates Fee references (usage: start with the 1 possible Fee reference, then provide the segments references)
	FeeRef *ReferenceInfoType_134839S `xml:"feeRef,omitempty"`  // minOccurs="0"

	// Passenger Reference
	PaxReference []*TravellerReferenceInformationType `xml:"paxReference"`  // maxOccurs="6"

	// add tax details for each passenger of each recommendations.
	PassengerTaxDetails *TaxType `xml:"passengerTaxDetails,omitempty"`  // minOccurs="0"

	Fare []*Fare `xml:"fare,omitempty"`  // minOccurs="0" maxOccurs="7"

	FareDetails []*FareDetails `xml:"fareDetails"`  // maxOccurs="6"
}

type Fare struct {

	// Last Date to Ticket, Penalties
	PricingMessage *InteractiveFreeTextType_78559S `xml:"pricingMessage"`

	// Amount of penalties, Surcharges...
	MonetaryInformation *MonetaryInformationType `xml:"monetaryInformation,omitempty"`  // minOccurs="0"
}

type FareDetails struct {

	// Reference of the Requested Segment
	SegmentRef *OriginAndDestinationRequestType `xml:"segmentRef"`

	GroupOfFares []*GroupOfFares `xml:"groupOfFares,omitempty"`  // minOccurs="0" maxOccurs="4"

	// Amounts per passenger per requested segment.
	PsgSegMonetaryInformation *MonetaryInformationType_137835S `xml:"psgSegMonetaryInformation,omitempty"`  // minOccurs="0"

	// Majority Cabin Info
	MajCabin []*ProductInformationType `xml:"majCabin,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type GroupOfFares struct {

	// Contains details of Flight and Fare
	ProductInformation *FlightProductInformationType `xml:"productInformation"`

	// Fare calculation code details
	FareCalculationCodeDetails []*FareCalculationCodeDetailsType `xml:"fareCalculationCodeDetails,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Ticket designator, ticket code and fare basis.
	TicketInfos *FareQualifierDetailsType `xml:"ticketInfos,omitempty"`  // minOccurs="0"

	// Reference of Fare Family for each Fare Component
	FareFamiliesRef *ReferenceInfoType_134838S `xml:"fareFamiliesRef,omitempty"`  // minOccurs="0"
}

type SpecificRecDetails struct {

	// Recommendation details
	SpecificRecItem *ItemReferencesAndVersionsType `xml:"specificRecItem"`

	SpecificProductDetails []*SpecificProductDetails `xml:"specificProductDetails,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type SpecificProductDetails struct {

	// Product details
	ProductReferences *PricingTicketingSubsequentType `xml:"productReferences"`

	FareContextDetails []*FareContextDetails `xml:"fareContextDetails,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type FareContextDetails struct {

	// Reference of requested segment
	RequestedSegmentInfo *OriginAndDestinationRequestType_134833S `xml:"requestedSegmentInfo"`

	CnxContextDetails []*CnxContextDetails `xml:"cnxContextDetails,omitempty"`  // minOccurs="0" maxOccurs="4"
}

type CnxContextDetails struct {

	// Fare connection context details
	FareCnxInfo *FlightProductInformationType_134801S `xml:"fareCnxInfo"`
}

type WarningInfo struct {

	// Dummy Segment
	GlobalMessageMarker *DummySegmentTypeI `xml:"globalMessageMarker"`

	// Informative free text information
	GlobalMessage *InteractiveFreeTextType `xml:"globalMessage"`
}

type GlobalInformation struct {

	// Coded attributes
	Attributes *CodedAttributeType `xml:"attributes"`
}

type ServiceFeesGrp struct {

	// Service fee type (OC)
	ServiceTypeInfo *SelectionDetailsType `xml:"serviceTypeInfo"`

	ServiceFeeRefGrp []*ServiceFeeRefGrp `xml:"serviceFeeRefGrp,omitempty"`  // minOccurs="0" maxOccurs="100000"

	ServiceCoverageInfoGrp []*ServiceCoverageInfoGrp `xml:"serviceCoverageInfoGrp,omitempty"`  // minOccurs="0" maxOccurs="100000"

	// Globalmessage marker
	GlobalMessageMarker *DummySegmentTypeI `xml:"globalMessageMarker"`

	ServiceFeeInfoGrp []*ServiceFeeInfoGrp `xml:"serviceFeeInfoGrp,omitempty"`  // minOccurs="0" maxOccurs="100000"

	ServiceDetailsGrp []*ServiceDetailsGrp1 `xml:"serviceDetailsGrp,omitempty"`  // minOccurs="0" maxOccurs="200"
}

type ServiceFeeRefGrp struct {

	// Reference of service fee global information
	RefInfo *ReferenceInfoType `xml:"refInfo"`
}

type ServiceCoverageInfoGrp struct {

	// Item reference number for service coverage details
	ItemNumberInfo *ItemNumberType `xml:"itemNumberInfo"`

	ServiceCovInfoGrp []*ServiceCovInfoGrp `xml:"serviceCovInfoGrp,omitempty"`  // minOccurs="0" maxOccurs="200"
}

type ServiceCovInfoGrp struct {

	// Passenger reference number
	PaxRefInfo *SpecificTravellerType `xml:"paxRefInfo"`

	// Service coverage information at flight level Matched seat characteristics
	CoveragePerFlightsInfo []*ActionDetailsType `xml:"coveragePerFlightsInfo,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Service reference number
	RefInfo *ReferenceInfoType_134840S `xml:"refInfo,omitempty"`  // minOccurs="0"
}

type ServiceFeeInfoGrp struct {

	// Item number details
	ItemNumberInfo *ItemNumberType `xml:"itemNumberInfo"`

	ServiceDetailsGrp []*ServiceDetailsGrp `xml:"serviceDetailsGrp,omitempty"`  // minOccurs="0" maxOccurs="200"
}

type ServiceDetailsGrp struct {

	// Service reference number
	RefInfo *ReferenceInfoType_134840S `xml:"refInfo"`

	ServiceMatchedInfoGroup []*ServiceMatchedInfoGroup `xml:"serviceMatchedInfoGroup,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ServiceMatchedInfoGroup struct {

	// Reference on pax number
	PaxRefInfo *SpecificTravellerType `xml:"paxRefInfo"`

	// Pricing oriented service matched information
	PricingInfo *FareInformationType_80868S `xml:"pricingInfo,omitempty"`  // minOccurs="0"

	// Informative Service amount
	AmountInfo *MonetaryInformationType_134806S `xml:"amountInfo,omitempty"`  // minOccurs="0"
}

type ServiceDetailsGrp1 struct {

	// Service sub-code and options (exclusion,inclusion, mode pushed,polled)
	ServiceOptionInfo *SpecificDataInformationType `xml:"serviceOptionInfo"`

	FeeDescriptionGrp *FeeDescriptionGrp `xml:"feeDescriptionGrp,omitempty"`  // minOccurs="0"
}

type FeeDescriptionGrp struct {

	// Specification of the item number
	ItemNumberInfo *ItemNumberType_80866S `xml:"itemNumberInfo"`

	// Attributes  (SSR code EMD, RFIC, SSIM)
	ServiceAttributesInfo *AttributeType `xml:"serviceAttributesInfo,omitempty"`  // minOccurs="0"

	// Other service information (service description, ...)
	ServiceDescriptionInfo *SpecialRequirementsDetailsType `xml:"serviceDescriptionInfo,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type ActionDetailsType struct {

	// Number of items details
	NumberOfItemsDetails *ProcessingInformationType `xml:"numberOfItemsDetails,omitempty"`  // minOccurs="0"

	// Range of segments
	LastItemsDetails []*ReferenceType `xml:"lastItemsDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type AdditionalFareQualifierDetailsTypeI struct {

	// Rate class
	RateClass string `xml:"rateClass,omitempty"`  // minOccurs="0"

	// Ticket designator.
	TicketDesignator string `xml:"ticketDesignator,omitempty"`  // minOccurs="0"

	// Pricing group
	PricingGroup string `xml:"pricingGroup,omitempty"`  // minOccurs="0"

	// Second rate class
	SecondRateClass []string `xml:"secondRateClass,omitempty"`  // minOccurs="0" maxOccurs="29"
}

type AdditionalProductDetailsType struct {

	// Type of aircraft
	EquipmentType string `xml:"equipmentType,omitempty"`  // minOccurs="0"

	// Day number of the week
	OperatingDay string `xml:"operatingDay,omitempty"`  // minOccurs="0"

	// Number of stops made in a journey if different from 0
	TechStopNumber *int32 `xml:"techStopNumber,omitempty"`  // minOccurs="0"

	// Location places of the stops
	LocationId []string `xml:"locationId,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type ApplicationErrorInformationType struct {

	// The code assigned by the receiver of a message for identification of a data validation error condition.
	Error string `xml:"error"`
}

type ApplicationErrorInformationType_78543S struct {

	// Details on application error.
	ApplicationErrorDetail *ApplicationErrorInformationType `xml:"applicationErrorDetail"`
}

type AttributeInformationType struct {

	// Type of parameter.
	FeeParameterType string `xml:"feeParameterType,omitempty"`  // minOccurs="0"

	// Reference to company Id.
	FeeParameterDescription string `xml:"feeParameterDescription,omitempty"`  // minOccurs="0"
}

type AttributeInformationType_97181C struct {

	// Attribute type
	AttributeType string `xml:"attributeType"`

	// Attribute description
	AttributeDescription string `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type AttributeType struct {

	// Criteria Set Type
	AttributeQualifier string `xml:"attributeQualifier,omitempty"`  // minOccurs="0"

	// Criteria details
	AttributeDetails []*AttributeInformationType_97181C `xml:"attributeDetails"`  // maxOccurs="99"
}

type AttributeType_78561S struct {

	// Fee/reduction parameters.
	FeeParameter []*AttributeInformationType `xml:"feeParameter,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type CabinProductDetailsType struct {

	// Reservation booking designator - RBD
	Rbd string `xml:"rbd"`

	// Reservation Booking Modifier
	BookingModifier string `xml:"bookingModifier,omitempty"`  // minOccurs="0"

	// Indicates the cabin related to the Booking code
	Cabin string `xml:"cabin,omitempty"`  // minOccurs="0"

	// Availibility status : posting level
	AvlStatus string `xml:"avlStatus,omitempty"`  // minOccurs="0"
}

type CabinProductDetailsType_195516C struct {

	// Reservation booking designator - RBD
	Rbd string `xml:"rbd,omitempty"`  // minOccurs="0"

	// Reservation Booking Modifier
	BookingModifier string `xml:"bookingModifier,omitempty"`  // minOccurs="0"

	// Indicates the cabin related to the Booking code
	Cabin string `xml:"cabin,omitempty"`  // minOccurs="0"

	// Availibility status : posting level
	AvlStatus string `xml:"avlStatus"`
}

type CodedAttributeInformationType struct {

	// Type of fee/reduction
	AttributeType string `xml:"attributeType"`

	// Fee Id Number
	AttributeDescription string `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeType struct {

	// Fee/reduction Id
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"`  // maxOccurs="9"
}

type CommercialAgreementsType struct {

	// Codeshare Details
	CodeshareDetails *CompanyRoleIdentificationType `xml:"codeshareDetails,omitempty"`  // minOccurs="0"

	// Other codeshare details
	OtherCodeshareDetails []*CompanyRoleIdentificationType `xml:"otherCodeshareDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type CompanyIdentificationTextType struct {

	// Company Id Text reference.
	TextRefNumber *int32 `xml:"textRefNumber,omitempty"`  // minOccurs="0"

	// Company id free text.
	CompanyText string `xml:"companyText,omitempty"`  // minOccurs="0"
}

type CompanyIdentificationType struct {

	// Marketing carrier
	MarketingCarrier string `xml:"marketingCarrier"`

	// Operating carrier
	OperatingCarrier string `xml:"operatingCarrier,omitempty"`  // minOccurs="0"
}

type CompanyRoleIdentificationType struct {

	// Type of code share agreement.
	CodeShareType string `xml:"codeShareType,omitempty"`  // minOccurs="0"

	// company identification
	AirlineDesignator string `xml:"airlineDesignator,omitempty"`  // minOccurs="0"

	// flight number
	FlightNumber *int32 `xml:"flightNumber,omitempty"`  // minOccurs="0"
}

type CompanyRoleIdentificationType_120771C struct {

	// Type of code share agreement.
	TransportStageQualifier string `xml:"transportStageQualifier,omitempty"`  // minOccurs="0"

	// company identification
	Company string `xml:"company,omitempty"`  // minOccurs="0"
}

type ConversionRateDetailsTypeI struct {

	// Conversion type
	ConversionType string `xml:"conversionType,omitempty"`  // minOccurs="0"

	// Currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// amount
	Amount string `xml:"amount,omitempty"`  // minOccurs="0"
}

type ConversionRateDetailsTypeI_179848C struct {

	// Conversion type
	ConversionType string `xml:"conversionType,omitempty"`  // minOccurs="0"

	// Currency
	Currency string `xml:"currency"`

	// Conversion rate for pricing
	Rate string `xml:"rate,omitempty"`  // minOccurs="0"

	// Converted value amount
	ConvertedAmountLink string `xml:"convertedAmountLink,omitempty"`  // minOccurs="0"

	// Applicable ISO country code or Tax designator code.
	TaxQualifier string `xml:"taxQualifier,omitempty"`  // minOccurs="0"
}

type ConversionRateTypeI struct {

	// Detail of conversion rate of First Monetary Unit.
	ConversionRateDetail []*ConversionRateDetailsTypeI_179848C `xml:"conversionRateDetail"`  // maxOccurs="9"
}

type ConversionRateTypeI_78562S struct {

	// Details of conversion
	ConversionRateDetail []*ConversionRateDetailsTypeI `xml:"conversionRateDetail"`  // maxOccurs="9"
}

type DataInformationType struct {

	// Ancillary services options
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"
}

type DataTypeInformationType struct {

	// service group/sub-group/sub-code information
	SubType string `xml:"subType"`

	// Status (automated, manually added, exempted). Default is automated
	Option string `xml:"option,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsType struct {

	// Date time period qualifier
	DateQualifier string `xml:"dateQualifier,omitempty"`  // minOccurs="0"

	// First Date
	Date string `xml:"date,omitempty"`  // minOccurs="0"

	// First Time
	FirstTime string `xml:"firstTime,omitempty"`  // minOccurs="0"

	// Movement type.
	EquipementType string `xml:"equipementType,omitempty"`  // minOccurs="0"

	// Place/location identification.
	LocationId string `xml:"locationId,omitempty"`  // minOccurs="0"
}

type DateAndTimeInformationType struct {

	// Details on date and time
	StopDetails []*DateAndTimeDetailsType `xml:"stopDetails"`  // maxOccurs="2"
}

type DiscountAndPenaltyInformationType struct {

	// Used to specify airline collected fee or agent collected fee.
	FeeIdentification string `xml:"feeIdentification,omitempty"`  // minOccurs="0"

	// Used to specify penalty information
	FeeInformation *DiscountPenaltyMonetaryInformationType `xml:"feeInformation,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyInformationType struct {

	// Discounted fare,...
	FareQualifier string `xml:"fareQualifier"`

	// Dicount code,...
	RateCategory string `xml:"rateCategory,omitempty"`  // minOccurs="0"

	// Amount
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Percentage
	Percentage *float64 `xml:"percentage,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyMonetaryInformationType struct {

	// Type of discount/penalty
	FeeType string `xml:"feeType,omitempty"`  // minOccurs="0"

	// The amount Type can be a percentage or an amount
	FeeAmountType string `xml:"feeAmountType,omitempty"`  // minOccurs="0"

	// specify the value
	FeeAmount *float64 `xml:"feeAmount,omitempty"`  // minOccurs="0"

	// Fee currency code.
	FeeCurrency string `xml:"feeCurrency,omitempty"`  // minOccurs="0"
}

type DummySegmentTypeI struct {
}

type FareCalculationCodeDetailsType struct {

	// Qualifier of the amout or rate
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Amount
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Location code
	LocationCode string `xml:"locationCode,omitempty"`  // minOccurs="0"

	// Other location code
	OtherLocationCode string `xml:"otherLocationCode,omitempty"`  // minOccurs="0"

	// Rate
	Rate *float64 `xml:"rate,omitempty"`  // minOccurs="0"
}

type FareCategoryCodesTypeI struct {

	// Fare type
	FareType string `xml:"fareType"`

	// Other fare type
	OtherFareType []string `xml:"otherFareType,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type FareDetailsType struct {

	// Passenger Type qualifier
	PassengerTypeQualifier string `xml:"passengerTypeQualifier,omitempty"`  // minOccurs="0"
}

type FareDetailsTypeI struct {

	// Qualifier
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Rate
	Rate *float64 `xml:"rate,omitempty"`  // minOccurs="0"

	// Country
	Country string `xml:"country,omitempty"`  // minOccurs="0"

	// Fare category
	FareCategory string `xml:"fareCategory,omitempty"`  // minOccurs="0"
}

type FareDetailsType_193037C struct {

	// Qualifier
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Rate
	Rate *int32 `xml:"rate,omitempty"`  // minOccurs="0"

	// Country
	Country string `xml:"country,omitempty"`  // minOccurs="0"

	// Fare Category
	FareCategory string `xml:"fareCategory,omitempty"`  // minOccurs="0"
}

type FareFamilyType struct {

	// Fare Family Reference Number
	RefNumber int32 `xml:"refNumber"`

	// Fare Family Short Name
	FareFamilyname string `xml:"fareFamilyname,omitempty"`  // minOccurs="0"

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy *int32 `xml:"hierarchy,omitempty"`  // minOccurs="0"

	// CABIN USED FOR FARE FAMILY
	Cabin string `xml:"cabin,omitempty"`  // minOccurs="0"
}

type FareInformationType struct {

	// Value Qualifier
	ValueQualifier string `xml:"valueQualifier,omitempty"`  // minOccurs="0"

	// Value
	Value *int32 `xml:"value,omitempty"`  // minOccurs="0"

	// Fare Details
	FareDetails *FareDetailsType_193037C `xml:"fareDetails,omitempty"`  // minOccurs="0"

	// Identity Number
	IdentityNumber string `xml:"identityNumber,omitempty"`  // minOccurs="0"

	// Fare Type Grouping
	FareTypeGrouping *FareTypeGroupingInformationType `xml:"fareTypeGrouping,omitempty"`  // minOccurs="0"

	// Rate Category
	RateCategory string `xml:"rateCategory,omitempty"`  // minOccurs="0"
}

type FareInformationType_80868S struct {

	// Fare details
	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`  // minOccurs="0"
}

type FareProductDetailsType struct {

	// Fare basis code
	FareBasis string `xml:"fareBasis,omitempty"`  // minOccurs="0"

	// PTC priced
	PassengerType string `xml:"passengerType,omitempty"`  // minOccurs="0"

	// Type of fare
	FareType []string `xml:"fareType,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FareQualifierDetailsType struct {

	// Route Code
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	// Fare categories
	FareCategories *FareCategoryCodesTypeI `xml:"fareCategories,omitempty"`  // minOccurs="0"

	// Fare details
	FareDetails *FareDetailsTypeI `xml:"fareDetails,omitempty"`  // minOccurs="0"

	// Additional fare details
	AdditionalFareDetails *AdditionalFareQualifierDetailsTypeI `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	// Discount details
	DiscountDetails []*DiscountPenaltyInformationType `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FareTypeGroupingInformationType struct {

	// Pricing Group
	PricingGroup string `xml:"pricingGroup,omitempty"`  // minOccurs="0"
}

type FlightProductInformationType struct {

	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType `xml:"cabinProduct,omitempty"`  // minOccurs="0"

	// Fare product details
	FareProductDetail *FareProductDetailsType `xml:"fareProductDetail,omitempty"`  // minOccurs="0"

	// Corporate number or name and number
	CorporateId []string `xml:"corporateId,omitempty"`  // minOccurs="0" maxOccurs="2"

	// To determine if Fare Breaks at this segment
	BreakPoint string `xml:"breakPoint,omitempty"`  // minOccurs="0"

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType `xml:"contextDetails,omitempty"`  // minOccurs="0"
}

type FlightProductInformationType_134801S struct {

	// Indicates flight cabin details
	CabinProduct []*CabinProductDetailsType_195516C `xml:"cabinProduct,omitempty"`  // minOccurs="0" maxOccurs="6"

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType `xml:"contextDetails,omitempty"`  // minOccurs="0"
}

type FreeTextQualificationType struct {

	// Type of message
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Coded Text or type of information in 4440 (e.g. type of OSI or free text, canned message value)
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"
}

type FreeTextQualificationType_120769C struct {

	// Type of message
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Coded Text or type of information in 4440 (e.g. type of OSI or free text, canned message value)
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// ISO code for language of free text (default is English)
	Language string `xml:"language,omitempty"`  // minOccurs="0"
}

type InteractiveFreeTextType struct {

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Free text
	Description []string `xml:"description,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type InteractiveFreeTextType_78544S struct {

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType_120769C `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Free text
	Description []string `xml:"description,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type InteractiveFreeTextType_78559S struct {

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType_120769C `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Free text
	Description []string `xml:"description,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type ItemNumberIdentificationType struct {

	// Ancillary Service number
	Number string `xml:"number,omitempty"`  // minOccurs="0"

	// Type
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Qualifier
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Responsible agency
	ResponsibleAgency string `xml:"responsibleAgency,omitempty"`  // minOccurs="0"
}

type ItemNumberIdentificationType_191597C struct {

	// Item number.
	Number string `xml:"number,omitempty"`  // minOccurs="0"

	// Indicates the item type .
	NumberType string `xml:"numberType,omitempty"`  // minOccurs="0"
}

type ItemNumberIdentificationType_192331C struct {

	// Service coverage number
	Number string `xml:"number,omitempty"`  // minOccurs="0"

	// Type
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Qualifier
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Responsible agency
	ResponsibleAgency string `xml:"responsibleAgency,omitempty"`  // minOccurs="0"
}

type ItemNumberType struct {

	// Item number details
	ItemNumber *ItemNumberIdentificationType_192331C `xml:"itemNumber"`
}

type ItemNumberType_133453S struct {

	// Indicates the recommendation number.
	ItemNumberId *ItemNumberIdentificationType_191597C `xml:"itemNumberId,omitempty"`  // minOccurs="0"

	// Code share details.
	CodeShareDetails []*CompanyRoleIdentificationType_120771C `xml:"codeShareDetails,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Pricing ticketind details.
	PriceTicketing *PricingTicketingInformationType `xml:"priceTicketing,omitempty"`  // minOccurs="0"
}

type ItemNumberType_80866S struct {

	// Item number details
	ItemNumberDetails *ItemNumberIdentificationType `xml:"itemNumberDetails"`
}

type ItemReferencesAndVersionsType struct {

	// Qualifies the type of the reference used.
	ReferenceType string `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Unique fee reference.
	RefNumber *int32 `xml:"refNumber,omitempty"`  // minOccurs="0"
}

type ItemReferencesAndVersionsType_78536S struct {

	// Qualifies the type of the reference used.
	ReferenceType string `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Unique fee reference.
	RefNumber *int32 `xml:"refNumber,omitempty"`  // minOccurs="0"
}

type ItemReferencesAndVersionsType_78564S struct {

	// Qualifies the type of the reference used.
	ReferenceType string `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Unique fee reference.
	FeeRefNumber *int32 `xml:"feeRefNumber,omitempty"`  // minOccurs="0"
}

type ItineraryDetailsType struct {

	// Airport/City Qualifier: the passenger wants to depart/arrive from/to the same airport or city in the specified requested segment
	AirportCityQualifier string `xml:"airportCityQualifier"`

	// Requested segment number
	SegmentNumber int32 `xml:"segmentNumber"`
}

type LocationIdentificationDetailsType struct {

	// 3 characters ATA/IATA airport/city code
	LocationId string `xml:"locationId"`

	// Airport/city qualifier: the requested point is an airport when ambiguity exists (e.g. HOU)
	AirportCityQualifier string `xml:"airportCityQualifier,omitempty"`  // minOccurs="0"

	// Terminal information
	Terminal string `xml:"terminal,omitempty"`  // minOccurs="0"
}

type MiniRulesDetailsType struct {

	// Coded text (period or day)
	Interpretation string `xml:"interpretation,omitempty"`  // minOccurs="0"

	// Data type coded or value of interpretation
	Value []string `xml:"value,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type MiniRulesIndicatorType struct {

	// See rule indicator and free form text indicator
	RuleIndicator []string `xml:"ruleIndicator,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type MiniRulesType struct {

	// Type of restriction: PTC, Max Adv Res, Max Ticketing After Res, ...
	RestrictionType string `xml:"restrictionType,omitempty"`  // minOccurs="0"

	// Categoty of restriction: PTC, Max Adv Pur, Days, ...
	Category string `xml:"category"`

	// Indicators
	Indicator *MiniRulesIndicatorType `xml:"indicator,omitempty"`  // minOccurs="0"

	// Mini rules
	MiniRules []*MiniRulesDetailsType `xml:"miniRules,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type MonetaryInformationDetailsType struct {

	// To specify amount and percentage.
	AmountType string `xml:"amountType,omitempty"`  // minOccurs="0"

	// Amount
	Amount float64 `xml:"amount"`

	// ISO currency code
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationType struct {

	// Monetary information
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail"`  // maxOccurs="2"
}

type MonetaryInformationType_134806S struct {

	// Monetary information.
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"`  // minOccurs="0" maxOccurs="30"
}

type MonetaryInformationType_137835S struct {

	// Monetary information.
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type OriginAndDestinationRequestType struct {

	// Requested segment number
	SegRef int32 `xml:"segRef"`

	// Forces arrival or departure, from/to the same airport/city
	LocationForcing []*ItineraryDetailsType `xml:"locationForcing,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type OriginAndDestinationRequestType_134833S struct {

	// Requested segment number
	SegRef int32 `xml:"segRef"`
}

type OriginatorIdentificationDetailsTypeI struct {

	// Office Name.
	OfficeName *int32 `xml:"officeName,omitempty"`  // minOccurs="0"

	// Agent Sign In .
	AgentSignin string `xml:"agentSignin,omitempty"`  // minOccurs="0"

	// Confidential Office Name.
	ConfidentialOffice string `xml:"confidentialOffice,omitempty"`  // minOccurs="0"

	// Other Office Name
	OtherOffice string `xml:"otherOffice,omitempty"`  // minOccurs="0"
}

type PricingTicketingInformationType struct {

	// Price type qualifier
	PriceType []string `xml:"priceType"`  // maxOccurs="20"
}

type PricingTicketingSubsequentType struct {

	// Passenger fare product number
	PaxFareNum []string `xml:"paxFareNum"`  // maxOccurs="10"
}

type PricingTicketingSubsequentType_134837S struct {

	// Passenger fare product number
	PaxFareNum string `xml:"paxFareNum"`

	// Total fare amount
	TotalFareAmount float64 `xml:"totalFareAmount"`

	// Total tax amount
	TotalTaxAmount *float64 `xml:"totalTaxAmount,omitempty"`  // minOccurs="0"

	// Code share details.
	CodeShareDetails []*CompanyRoleIdentificationType_120771C `xml:"codeShareDetails,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Monetary information.
	MonetaryDetails []*MonetaryInformationDetailsType `xml:"monetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Pricing ticketing details.
	PricingTicketing *PricingTicketingInformationType `xml:"pricingTicketing,omitempty"`  // minOccurs="0"
}

type ProcessingInformationType struct {

	// Action qualifier
	ActionQualifier string `xml:"actionQualifier,omitempty"`  // minOccurs="0"

	// Reference qualifier
	ReferenceQualifier string `xml:"referenceQualifier,omitempty"`  // minOccurs="0"

	// Reference number
	RefNum string `xml:"refNum,omitempty"`  // minOccurs="0"
}

type ProductDateTimeType struct {

	// Departure date
	DateOfDeparture string `xml:"dateOfDeparture"`

	// Departure time
	TimeOfDeparture string `xml:"timeOfDeparture,omitempty"`  // minOccurs="0"

	// Arrival date
	DateOfArrival string `xml:"dateOfArrival,omitempty"`  // minOccurs="0"

	// Arrival time
	TimeOfArrival string `xml:"timeOfArrival,omitempty"`  // minOccurs="0"

	// Arrival date compared to departure date, only if different from 0
	DateVariation *int32 `xml:"dateVariation,omitempty"`  // minOccurs="0"
}

type ProductDetailsType struct {

	// Format limitations: an..17
	Designator string `xml:"designator"`

	// Format limitations: an..3
	AvailabilityStatus string `xml:"availabilityStatus,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	SpecialService string `xml:"specialService,omitempty"`  // minOccurs="0"

	// Format limitations: an..7
	Option []string `xml:"option,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type ProductFacilitiesType struct {

	// Yes-No indicator whether Last Seat Available
	LastSeatAvailable string `xml:"lastSeatAvailable,omitempty"`  // minOccurs="0"

	// Level of access
	LevelOfAccess string `xml:"levelOfAccess,omitempty"`  // minOccurs="0"

	// Yes-No indicator whether electronic ticketing
	ElectronicTicketing string `xml:"electronicTicketing,omitempty"`  // minOccurs="0"

	// Product identification suffix
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`  // minOccurs="0"

	// Define whether a flight has been polled or not
	ProductDetailQualifier string `xml:"productDetailQualifier,omitempty"`  // minOccurs="0"

	// Add some flight restrictions (See code set list)
	FlightCharacteristic []string `xml:"flightCharacteristic,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type ProductInformationType struct {

	// value of the Qualifier: INT for International DOM for Domestic EUR for European  otherwise CM#10569 INVALID INTERNATIONAL INDICATOR is returned.
	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"`  // minOccurs="0"

	BookingClassDetails []*ProductDetailsType `xml:"bookingClassDetails,omitempty"`  // minOccurs="0" maxOccurs="26"
}

type ProductTypeDetailsType struct {

	// Availability connection type.
	AvailabilityCnxType []string `xml:"availabilityCnxType"`  // maxOccurs="9"
}

type ProposedSegmentDetailsType struct {

	// Flight proposal reference
	Ref string `xml:"ref,omitempty"`  // minOccurs="0"

	// Elapse Flying Time
	UnitQualifier string `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type ProposedSegmentType struct {

	// Parameters for proposed flight group
	FlightProposal []*ProposedSegmentDetailsType `xml:"flightProposal"`  // maxOccurs="9"

	// Flight characteristics.
	FlightCharacteristic string `xml:"flightCharacteristic,omitempty"`  // minOccurs="0"

	// Majority cabin
	MajCabin string `xml:"majCabin,omitempty"`  // minOccurs="0"
}

type ReferenceInfoType struct {

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_191583C `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="200"
}

type ReferenceInfoType_133176S struct {

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferenceInfoType_134838S struct {

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type ReferenceInfoType_134839S struct {

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_195561C `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferenceInfoType_134840S struct {

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_195561C `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="200"
}

type ReferenceType struct {

	// Reference  of leg
	RefOfLeg string `xml:"refOfLeg,omitempty"`  // minOccurs="0"

	// Reference of segment starting range
	FirstItemIdentifier *int32 `xml:"firstItemIdentifier,omitempty"`  // minOccurs="0"

	// Reference of segment ending range
	LastItemIdentifier *int32 `xml:"lastItemIdentifier,omitempty"`  // minOccurs="0"
}

type ReferencingDetailsType struct {

	// Reference qualifier
	RefQualifier string `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Requested segment reference
	RefNumber int32 `xml:"refNumber"`
}

type ReferencingDetailsType_191583C struct {

	// Service reference qualifier
	RefQualifier string `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Service reference
	RefNumber int32 `xml:"refNumber"`
}

type ReferencingDetailsType_195561C struct {

	// Segment reference qualifier
	RefQualifier string `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Flight or flight group reference
	RefNumber int32 `xml:"refNumber"`
}

type SelectionDetailsInformationType struct {

	// Carrier fee type
	Type string `xml:"type"`

	// Carrier fee status
	OptionInformation string `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type SelectionDetailsType struct {

	// Carrier fees options
	CarrierFeeDetails *SelectionDetailsInformationType `xml:"carrierFeeDetails"`
}

type SpecialRequirementsDataDetailsType struct {

	// SSR seat characteristic
	SeatCharacteristics []string `xml:"seatCharacteristics,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type SpecialRequirementsDetailsType struct {

	// To specify the Service Requirement of the customer
	ServiceRequirementsInfo *SpecialRequirementsTypeDetailsType `xml:"serviceRequirementsInfo"`

	// Seat details
	SeatDetails []*SpecialRequirementsDataDetailsType `xml:"seatDetails,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type SpecialRequirementsTypeDetailsType struct {

	// To specify the Service Classification of the Service Requirement.
	ServiceClassification string `xml:"serviceClassification"`

	// Status
	ServiceStatus string `xml:"serviceStatus,omitempty"`  // minOccurs="0"

	// To specify the number of items involved
	ServiceNumberOfInstances *int32 `xml:"serviceNumberOfInstances,omitempty"`  // minOccurs="0"

	// To specify to which marketing carrier the service applies
	ServiceMarketingCarrier string `xml:"serviceMarketingCarrier,omitempty"`  // minOccurs="0"

	// Specify the Service group
	ServiceGroup string `xml:"serviceGroup,omitempty"`  // minOccurs="0"

	// Specify the Service Sub-Group
	ServiceSubGroup string `xml:"serviceSubGroup,omitempty"`  // minOccurs="0"

	// Free Text attached to the Service.
	ServiceFreeText []string `xml:"serviceFreeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type SpecificDataInformationType struct {

	// Carrier fee description
	DataTypeInformation *DataTypeInformationType `xml:"dataTypeInformation"`

	// Data information
	DataInformation []*DataInformationType `xml:"dataInformation,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type SpecificTravellerDetailsType struct {

	// Reference number
	ReferenceNumber string `xml:"referenceNumber,omitempty"`  // minOccurs="0"
}

type SpecificTravellerType struct {

	// Traveller details
	TravellerDetails []*SpecificTravellerDetailsType `xml:"travellerDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type StatusDetailsType struct {

	// Advisory type information, Fare Server
	AdvisoryTypeInfo string `xml:"advisoryTypeInfo,omitempty"`  // minOccurs="0"

	// CPU time, user type
	Notification string `xml:"notification,omitempty"`  // minOccurs="0"

	// CPU time,user type
	Notification2 string `xml:"notification2,omitempty"`  // minOccurs="0"

	// Capture and trace information
	Description string `xml:"description,omitempty"`  // minOccurs="0"
}

type StatusType struct {

	// Status details
	Status []*StatusDetailsType `xml:"status"`  // maxOccurs="10"
}

type TaxDetailsType struct {

	// Amount
	Rate string `xml:"rate,omitempty"`  // minOccurs="0"

	// Country code
	CountryCode string `xml:"countryCode,omitempty"`  // minOccurs="0"

	// Currency code
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// Type
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Indicator
	Indicator []string `xml:"indicator,omitempty"`  // minOccurs="0" maxOccurs="98"
}

type TaxType struct {

	// Tax category
	TaxCategory string `xml:"taxCategory,omitempty"`  // minOccurs="0"

	// Tax details
	TaxDetails []*TaxDetailsType `xml:"taxDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type TravelProductType struct {

	// Date and time of departure and arrival
	ProductDateTime *ProductDateTimeType `xml:"productDateTime"`

	// Location of departure and arrival
	Location []*LocationIdentificationDetailsType `xml:"location"`  // maxOccurs="2"

	// Company identification
	CompanyId *CompanyIdentificationType `xml:"companyId,omitempty"`  // minOccurs="0"

	// Flight number or trainNumber
	FlightOrtrainNumber string `xml:"flightOrtrainNumber,omitempty"`  // minOccurs="0"

	// Product details
	ProductDetail *AdditionalProductDetailsType `xml:"productDetail,omitempty"`  // minOccurs="0"

	// Additional product details
	AddProductDetail *ProductFacilitiesType `xml:"addProductDetail,omitempty"`  // minOccurs="0"
}

type TravellerDetailsType struct {

	// Direct reference of passenger assigned by requesting system.
	Ref *int32 `xml:"ref,omitempty"`  // minOccurs="0"

	// Traveller is an infant
	InfantIndicator *int32 `xml:"infantIndicator,omitempty"`  // minOccurs="0"
}

type TravellerReferenceInformationType struct {

	// Requested passenger type
	Ptc []string `xml:"ptc,omitempty"`  // minOccurs="0" maxOccurs="3"

	// Traveller details
	Traveller []*TravellerDetailsType `xml:"traveller,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type UserIdentificationType struct {

	// Originator Identification Details
	OfficeIdentification *OriginatorIdentificationDetailsTypeI `xml:"officeIdentification,omitempty"`  // minOccurs="0"

	// Used to specify which kind of info is given in DE 9900.
	OfficeType string `xml:"officeType,omitempty"`  // minOccurs="0"

	// The code given to an agent by the originating reservation system.
	OfficeCode string `xml:"officeCode,omitempty"`  // minOccurs="0"
}
