package fare_masterpricertravelboardsearch_reply

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type FareMasterPricerTravelBoardSearchReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A Fare_MasterPricerTravelBoardSearchReply"`

	// Gives status about reply : type of process, region , CPU etc..
	ReplyStatus *StatusType `xml:"replyStatus,omitempty"`

	ErrorMessage *ErrorMessage `xml:"errorMessage,omitempty"`

	// Specifies the currency used for pricing.
	ConversionRate *ConversionRateTypeI `xml:"conversionRate,omitempty"`

	// Solution Family
	SolutionFamily *FareInformationType `xml:"solutionFamily,omitempty"`

	// Details of the fare families processed
	FamilyInformation *FareFamilyType `xml:"familyInformation,omitempty"`

	AmountInfoForAllPax *AmountInfoForAllPax `xml:"amountInfoForAllPax,omitempty"`

	AmountInfoPerPax *AmountInfoPerPax `xml:"amountInfoPerPax,omitempty"`

	FeeDetails *FeeDetails `xml:"feeDetails,omitempty"`

	// Free text identifying an airline in a code share.
	CompanyIdText *CompanyIdentificationTextType `xml:"companyIdText,omitempty"`

	OfficeIdDetails *OfficeIdDetails `xml:"officeIdDetails,omitempty"`

	FlightIndex *FlightIndex `xml:"flightIndex,omitempty"`

	Recommendation *Recommendation `xml:"recommendation,omitempty"`

	OtherSolutions *OtherSolutions `xml:"otherSolutions,omitempty"`

	WarningInfo *WarningInfo `xml:"warningInfo,omitempty"`

	GlobalInformation *GlobalInformation `xml:"globalInformation,omitempty"`

	ServiceFeesGrp *ServiceFeesGrp `xml:"serviceFeesGrp,omitempty"`

	Value *ValueSearchCriteriaType `xml:"value,omitempty"`

	MnrGrp *MnrGrp `xml:"mnrGrp,omitempty"`
}

type ErrorMessage struct {

	// Application error details.
	ApplicationError *ApplicationErrorInformationType_78543S `xml:"applicationError,omitempty"`

	// Type of error message and free text
	ErrorMessageText *InteractiveFreeTextType_78544S `xml:"errorMessageText,omitempty"`
}

type AmountInfoForAllPax struct {

	// Itinerary amounts for all passengers
	ItineraryAmounts *MonetaryInformationType_137835S `xml:"itineraryAmounts,omitempty"`

	AmountsPerSgt *AmountsPerSgt `xml:"amountsPerSgt,omitempty"`
}

type AmountsPerSgt struct {

	// Requested segment reference
	SgtRef *ReferenceInfoType_133176S `xml:"sgtRef,omitempty"`

	// Amounts : Issue total amount, issue taxes amount, non refundable taxes amount
	Amounts *MonetaryInformationType_137835S `xml:"amounts,omitempty"`
}

type AmountInfoPerPax struct {

	// Passenger references
	PaxRef *SpecificTravellerType `xml:"paxRef,omitempty"`

	// Passenger attributes : Infant indicator
	PaxAttributes *FareInformationType_80868S `xml:"paxAttributes,omitempty"`

	// Itinerary amounts information
	ItineraryAmounts *MonetaryInformationType_137835S `xml:"itineraryAmounts,omitempty"`

	AmountsPerSgt *AmountsPerSgt `xml:"amountsPerSgt,omitempty"`
}

type FeeDetails struct {

	// Fee/Reduction Reference number.
	FeeReference *ItemReferencesAndVersionsType_78564S `xml:"feeReference,omitempty"`

	// Fee/Reduction information.
	FeeInformation *DiscountAndPenaltyInformationType `xml:"feeInformation,omitempty"`

	// Fee/Reduction parameters.
	FeeParameters *AttributeType_78561S `xml:"feeParameters,omitempty"`

	// To specify conversion rate details
	ConvertedOrOriginalInfo *ConversionRateTypeI_78562S `xml:"convertedOrOriginalInfo,omitempty"`
}

type OfficeIdDetails struct {

	// Office Id Information
	OfficeIdInformation *UserIdentificationType `xml:"officeIdInformation,omitempty"`

	// Office Id Reference Number
	OfficeIdReference *ItemReferencesAndVersionsType_78536S `xml:"officeIdReference,omitempty"`
}

type FlightIndex struct {

	// Indicates references and details about requested segments
	RequestedSegmentRef *OriginAndDestinationRequestType `xml:"requestedSegmentRef,omitempty"`

	GroupOfFlights *GroupOfFlights `xml:"groupOfFlights,omitempty"`
}

type GroupOfFlights struct {

	// To indicate parameters for proposed flight group.
	PropFlightGrDetail *ProposedSegmentType `xml:"propFlightGrDetail,omitempty"`

	FlightDetails *FlightDetails `xml:"flightDetails,omitempty"`
}

type FlightDetails struct {

	// Specification of details on the flight and posting availability
	FlightInformation *TravelProductType `xml:"flightInformation,omitempty"`

	// returns booking class and availability context
	AvlInfo *FlightProductInformationType_141442S `xml:"avlInfo,omitempty"`

	// Details on Flight date, time and location of technical stop or change of gauge
	TechnicalStop *DateAndTimeInformationType `xml:"technicalStop,omitempty"`

	// Code Share Agreement description for current flight.
	CommercialAgreement *CommercialAgreementsType `xml:"commercialAgreement,omitempty"`

	// Additional Info about flight, such as Reference number, and several options
	AddInfo *HeaderInformationTypeI `xml:"addInfo,omitempty"`

	// Flight characteristics
	FlightCharacteristics *FlightCharacteristicsType `xml:"flightCharacteristics,omitempty"`

	// Flight Services by cabin/rbd
	FlightServices *FlightServicesType `xml:"flightServices,omitempty"`
}

type Recommendation struct {

	// Specification of the item number
	ItemNumber *ItemNumberType_161497S `xml:"itemNumber,omitempty"`

	// To describe type of recommendation
	WarningMessage *InteractiveFreeTextType_78544S `xml:"warningMessage,omitempty"`

	// Indicates the Fare family reference.
	FareFamilyRef *ReferenceInfoType_133176S `xml:"fareFamilyRef,omitempty"`

	// Recommendation Price and Taxes.
	RecPriceInfo *MonetaryInformationType `xml:"recPriceInfo,omitempty"`

	// Mini rules
	MiniRule *MiniRulesType_78547S `xml:"miniRule,omitempty"`

	// Indicates reference of Flight or the fee reference valid for all pax (usage:start with the 1 possible Fee reference, then provide the segments references)
	SegmentFlightRef *ReferenceInfoType `xml:"segmentFlightRef,omitempty"`

	RecommandationSegmentsFareDetails *RecommandationSegmentsFareDetails `xml:"recommandationSegmentsFareDetails,omitempty"`

	PaxFareProduct *PaxFareProduct `xml:"paxFareProduct,omitempty"`

	SpecificRecDetails *SpecificRecDetails `xml:"specificRecDetails,omitempty"`
}

type RecommandationSegmentsFareDetails struct {

	// Reference and details about requested segments.
	RecommendationSegRef *OriginAndDestinationRequestType `xml:"recommendationSegRef,omitempty"`

	// Amounts per requested segment.
	SegmentMonetaryInformation *MonetaryInformationType_137835S `xml:"segmentMonetaryInformation,omitempty"`
}

type PaxFareProduct struct {

	// Passenger Fare Details.
	PaxFareDetail *PricingTicketingSubsequentType_144401S `xml:"paxFareDetail,omitempty"`

	// Indicates Fee references (usage: start with the 1 possible Fee reference, then provide the segments references)
	FeeRef *ReferenceInfoType_134839S `xml:"feeRef,omitempty"`

	// Passenger Reference
	PaxReference *TravellerReferenceInformationType `xml:"paxReference,omitempty"`

	// add tax details for each passenger of each recommendations.
	PassengerTaxDetails *TaxType `xml:"passengerTaxDetails,omitempty"`

	Fare *Fare `xml:"fare,omitempty"`

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`
}

type Fare struct {

	// Last Date to Ticket, Penalties
	PricingMessage *InteractiveFreeTextType_78559S `xml:"pricingMessage,omitempty"`

	// Amount of penalties, Surcharges...
	MonetaryInformation *MonetaryInformationType_185955S `xml:"monetaryInformation,omitempty"`
}

type FareDetails struct {

	// Reference of the Requested Segment
	SegmentRef *OriginAndDestinationRequestType `xml:"segmentRef,omitempty"`

	GroupOfFares *GroupOfFares `xml:"groupOfFares,omitempty"`

	// Amounts per passenger per requested segment.
	PsgSegMonetaryInformation *MonetaryInformationType_137835S `xml:"psgSegMonetaryInformation,omitempty"`

	// Majority Cabin Info
	MajCabin *ProductInformationType `xml:"majCabin,omitempty"`
}

type GroupOfFares struct {

	// Contains details of Flight and Fare
	ProductInformation *FlightProductInformationType_176659S `xml:"productInformation,omitempty"`

	// Fare calculation code details
	FareCalculationCodeDetails *FareCalculationCodeDetailsType `xml:"fareCalculationCodeDetails,omitempty"`

	// Ticket designator, ticket code and fare basis.
	TicketInfos *FareQualifierDetailsType `xml:"ticketInfos,omitempty"`

	// Reference of Fare Family for each Fare Component
	FareFamiliesRef *ReferenceInfoType_176658S `xml:"fareFamiliesRef,omitempty"`
}

type SpecificRecDetails struct {

	// Recommendation details
	SpecificRecItem *ItemReferencesAndVersionsType `xml:"specificRecItem,omitempty"`

	SpecificProductDetails *SpecificProductDetails `xml:"specificProductDetails,omitempty"`
}

type SpecificProductDetails struct {

	// Product details
	ProductReferences *PricingTicketingSubsequentType `xml:"productReferences,omitempty"`

	FareContextDetails *FareContextDetails `xml:"fareContextDetails,omitempty"`
}

type FareContextDetails struct {

	// Reference of requested segment
	RequestedSegmentInfo *OriginAndDestinationRequestType_134833S `xml:"requestedSegmentInfo,omitempty"`

	CnxContextDetails *CnxContextDetails `xml:"cnxContextDetails,omitempty"`
}

type CnxContextDetails struct {

	// Fare connection context details
	FareCnxInfo *FlightProductInformationType `xml:"fareCnxInfo,omitempty"`
}

type OtherSolutions struct {

	// Reference to the current solution
	Reference *SequenceDetailsTypeU `xml:"reference,omitempty"`

	AmtGroup *AmtGroup `xml:"amtGroup,omitempty"`

	PsgInfo *PsgInfo `xml:"psgInfo,omitempty"`
}

type AmtGroup struct {

	// reference to the current amount (per bound, per segment...)
	Ref *ReferenceInfoType_165972S `xml:"ref,omitempty"`

	// Amount Description
	Amount *MonetaryInformationTypeI `xml:"amount,omitempty"`
}

type PsgInfo struct {

	// passenger reference
	Ref *SegmentRepetitionControlTypeI `xml:"ref,omitempty"`

	// Passenger Description Info
	Description *FareInformationTypeI `xml:"description,omitempty"`

	// Passenger frequent traveler info
	FreqTraveller *FrequentTravellerIdentificationCodeType `xml:"freqTraveller,omitempty"`

	// amount per passenger or group of passenger
	Amount *MonetaryInformationTypeI `xml:"amount,omitempty"`

	// Fare description
	Fare *FlightProductInformationType_161491S `xml:"fare,omitempty"`

	// Additional Information
	Attribute *AttributeTypeU `xml:"attribute,omitempty"`
}

type WarningInfo struct {

	// Dummy Segment
	GlobalMessageMarker *DummySegmentTypeI `xml:"globalMessageMarker,omitempty"`

	// Informative free text information
	GlobalMessage *InteractiveFreeTextType_78534S `xml:"globalMessage,omitempty"`
}

type GlobalInformation struct {

	// Coded attributes
	Attributes *CodedAttributeType `xml:"attributes,omitempty"`
}

type ServiceFeesGrp struct {

	// Service fee type (OC)
	ServiceTypeInfo *SelectionDetailsType `xml:"serviceTypeInfo,omitempty"`

	ServiceFeeRefGrp *ServiceFeeRefGrp `xml:"serviceFeeRefGrp,omitempty"`

	ServiceCoverageInfoGrp *ServiceCoverageInfoGrp `xml:"serviceCoverageInfoGrp,omitempty"`

	// Globalmessage marker
	GlobalMessageMarker *DummySegmentTypeI `xml:"globalMessageMarker,omitempty"`

	ServiceFeeInfoGrp *ServiceFeeInfoGrp `xml:"serviceFeeInfoGrp,omitempty"`

	ServiceDetailsGrp *ServiceDetailsGrpType `xml:"serviceDetailsGrp,omitempty"`

	FreeBagAllowanceGrp *FreeBagAllowanceGrp `xml:"freeBagAllowanceGrp,omitempty"`
}

type ServiceFeeRefGrp struct {

	// Reference of service fee global information
	RefInfo *ReferenceInfoType `xml:"refInfo,omitempty"`
}

type ServiceCoverageInfoGrp struct {

	// Item reference number for service coverage details
	ItemNumberInfo *ItemNumberType `xml:"itemNumberInfo,omitempty"`

	ServiceCovInfoGrp *ServiceCovInfoGrp `xml:"serviceCovInfoGrp,omitempty"`
}

type ServiceCovInfoGrp struct {

	// Passenger reference number
	PaxRefInfo *SpecificTravellerType `xml:"paxRefInfo,omitempty"`

	// Service coverage information at flight level Matched seat characteristics
	CoveragePerFlightsInfo *ActionDetailsType `xml:"coveragePerFlightsInfo,omitempty"`

	// Carrier information
	CarrierInfo *TransportIdentifierType `xml:"carrierInfo,omitempty"`

	// Service reference number
	RefInfo *ReferenceInfoType_134840S `xml:"refInfo,omitempty"`
}

type ServiceFeeInfoGrp struct {

	// Item number details
	ItemNumberInfo *ItemNumberType `xml:"itemNumberInfo,omitempty"`

	ServiceDetailsGrp *ServiceDetailsGrp `xml:"serviceDetailsGrp,omitempty"`
}

type ServiceDetailsGrp struct {

	// Service reference number
	RefInfo *ReferenceInfoType_134840S `xml:"refInfo,omitempty"`

	ServiceMatchedInfoGroup *ServiceMatchedInfoGroup `xml:"serviceMatchedInfoGroup,omitempty"`
}

type ServiceMatchedInfoGroup struct {

	// Reference on pax number
	PaxRefInfo *SpecificTravellerType `xml:"paxRefInfo,omitempty"`

	// Pricing oriented service matched information
	PricingInfo *FareInformationType_80868S `xml:"pricingInfo,omitempty"`

	// Informative Service amount
	AmountInfo *MonetaryInformationType `xml:"amountInfo,omitempty"`
}

type ServiceDetailsGrpType struct {

	// Service sub-code and options (exclusion,inclusion, mode pushed,polled)
	ServiceOptionInfo *SpecificDataInformationType `xml:"serviceOptionInfo,omitempty"`

	FeeDescriptionGrp *FeeDescriptionGrp `xml:"feeDescriptionGrp,omitempty"`
}

type FeeDescriptionGrp struct {

	// Specification of the item number
	ItemNumberInfo *ItemNumberType_80866S `xml:"itemNumberInfo,omitempty"`

	// Attributes  (SSR code EMD, RFIC, SSIM)
	ServiceAttributesInfo *AttributeType `xml:"serviceAttributesInfo,omitempty"`

	// Other service information (service description, ...)
	ServiceDescriptionInfo *SpecialRequirementsDetailsType `xml:"serviceDescriptionInfo,omitempty"`

	// Commercial name
	CommercialName *InteractiveFreeTextType `xml:"commercialName,omitempty"`
}

type FreeBagAllowanceGrp struct {

	// Free baggage allownce information
	FreeBagAllownceInfo *ExcessBaggageType `xml:"freeBagAllownceInfo,omitempty"`

	// Item number information
	ItemNumberInfo *ItemNumberType_166130S `xml:"itemNumberInfo,omitempty"`
}

type MnrGrp struct {
	Mnr *MiniRulesType `xml:"mnr,omitempty"`

	MnrDetails *MnrDetails `xml:"mnrDetails,omitempty"`
}

type MnrDetails struct {
	MnrRef *ItemNumberType_176648S `xml:"mnrRef,omitempty"`

	DateInfo *DateAndTimeInformationType_182345S `xml:"dateInfo,omitempty"`

	CatGrp *CatGrp `xml:"catGrp,omitempty"`
}

type CatGrp struct {

	// Category information
	CatInfo *CategDescrType `xml:"catInfo,omitempty"`

	// Monetary information
	MonInfo *MonetaryInformationType_174241S `xml:"monInfo,omitempty"`

	// Status information
	StatusInfo *StatusType_182386S `xml:"statusInfo,omitempty"`
}

type ActionDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ActionDetailsType"`

	// Number of items details
	NumberOfItemsDetails *ProcessingInformationType `xml:"numberOfItemsDetails,omitempty"`

	// Range of segments
	LastItemsDetails *ReferenceType `xml:"lastItemsDetails,omitempty"`
}

type AdditionalFareQualifierDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AdditionalFareQualifierDetailsTypeI"`

	// Rate class
	RateClass formats.AlphaNumericString_Length1To35 `xml:"rateClass,omitempty"`

	// Ticket designator.
	TicketDesignator formats.AlphaNumericString_Length1To18 `xml:"ticketDesignator,omitempty"`

	// Pricing group
	PricingGroup formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"`

	// Second rate class
	SecondRateClass formats.AlphaNumericString_Length1To35 `xml:"secondRateClass,omitempty"`
}

type AdditionalProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AdditionalProductDetailsType"`

	// Type of aircraft
	EquipmentType formats.AlphaNumericString_Length1To3 `xml:"equipmentType,omitempty"`

	// Day number of the week
	OperatingDay formats.AlphaNumericString_Length1To7 `xml:"operatingDay,omitempty"`

	// Number of stops made in a journey if different from 0
	TechStopNumber formats.NumericInteger_Length1To2 `xml:"techStopNumber,omitempty"`

	// Location places of the stops
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`
}

type ApplicationErrorInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ApplicationErrorInformationType"`

	// The code assigned by the receiver of a message for identification of a data validation error condition.
	Error formats.AlphaNumericString_Length1To4 `xml:"error,omitempty"`
}

type ApplicationErrorInformationType_78543S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ApplicationErrorInformationType_78543S"`

	// Details on application error.
	ApplicationErrorDetail *ApplicationErrorInformationType `xml:"applicationErrorDetail,omitempty"`
}

type AttributeInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeInformationTypeU"`

	// Attribute type
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType,omitempty"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type AttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeInformationType"`

	// Type of parameter.
	FeeParameterType formats.AlphaNumericString_Length3To3 `xml:"feeParameterType,omitempty"`

	// Reference to company Id.
	FeeParameterDescription formats.AlphaNumericString_Length1To15 `xml:"feeParameterDescription,omitempty"`
}

type AttributeInformationType_97181C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeInformationType_97181C"`

	// Attribute type
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType,omitempty"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type AttributeTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeTypeU"`

	// provides the function of the attribute
	AttributeFunction formats.AlphaNumericString_Length1To3 `xml:"attributeFunction,omitempty"`

	// provides details for the Attribute
	AttributeDetails *AttributeInformationTypeU `xml:"attributeDetails,omitempty"`
}

type AttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeType"`

	// Criteria Set Type
	AttributeQualifier formats.AlphaNumericString_Length1To3 `xml:"attributeQualifier,omitempty"`

	// Criteria details
	AttributeDetails *AttributeInformationType_97181C `xml:"attributeDetails,omitempty"`
}

type AttributeType_78561S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeType_78561S"`

	// Fee/reduction parameters.
	FeeParameter *AttributeInformationType `xml:"feeParameter,omitempty"`
}

type BaggageDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A BaggageDetailsType"`

	// Number of pieces or weight
	FreeAllowance formats.NumericInteger_Length1To15 `xml:"freeAllowance,omitempty"`

	// Nature of the free allowance ( Number of pieces or weight)
	QuantityCode formats.AMA_EDICodesetType_Length1to3 `xml:"quantityCode,omitempty"`

	// Unit qualifier
	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`
}

type BagtagDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A BagtagDetailsType"`

	// Identifier
	Identifier formats.AlphaNumericString_Length1To35 `xml:"identifier,omitempty"`

	// Number
	Number formats.NumericInteger_Length1To15 `xml:"number,omitempty"`
}

type CabinInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinInformationType"`

	// Identify the features associated to the cabin/class
	Service formats.AlphaNumericString_Length1To5 `xml:"service,omitempty"`

	// Cabin code designator
	Cabin formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`
}

type CabinProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinProductDetailsType"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`

	// Reservation Booking Modifier
	BookingModifier formats.AlphaNumericString_Length0To1 `xml:"bookingModifier,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaNumericString_Length1To1 `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus formats.AlphaNumericString_Length0To3 `xml:"avlStatus,omitempty"`
}

type CabinProductDetailsType_195516C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinProductDetailsType_195516C"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`

	// Reservation Booking Modifier
	BookingModifier formats.AlphaNumericString_Length0To1 `xml:"bookingModifier,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaNumericString_Length1To1 `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus formats.AlphaNumericString_Length0To3 `xml:"avlStatus,omitempty"`
}

type CabinProductDetailsType_205138C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinProductDetailsType_205138C"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`

	// Reservation Booking Modifier
	BookingModifier formats.AMA_EDICodesetType_Length1 `xml:"bookingModifier,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus formats.AMA_EDICodesetType_Length1to3 `xml:"avlStatus,omitempty"`
}

type CabinProductDetailsType_229142C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinProductDetailsType_229142C"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus formats.AMA_EDICodesetType_Length1to3 `xml:"avlStatus,omitempty"`
}

type CategDescrType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CategDescrType"`

	// Category description information
	DescriptionInfo *CategoryDescriptionType `xml:"descriptionInfo,omitempty"`

	// Category processing indicator
	ProcessIndicator formats.AlphaNumericString_Length1To3 `xml:"processIndicator,omitempty"`
}

type CategoryDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CategoryDescriptionType"`

	// Category number from ATPCO naming conventions (C05 for Advance Purchase restrictions, C06 for Minimun stay ...)
	Number formats.NumericInteger_Length1To3 `xml:"number,omitempty"`

	// Category Code (ATPCO component code, e.g ADV for Advance purchase, STP for stopover restrictions, ELG for eligibility restrictions...)
	Code formats.AlphaString_Length1To3 `xml:"code,omitempty"`
}

type ClassInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ClassInformationType"`

	// Identify the features associated to the cabin/class
	Service formats.AlphaNumericString_Length1To5 `xml:"service,omitempty"`

	// Class designator
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`
}

type CodedAttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CodedAttributeInformationType"`

	// Type of fee/reduction
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// Fee Id Number
	AttributeDescription formats.AlphaNumericString_Length1To50 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeInformationType_260992C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CodedAttributeInformationType_260992C"`

	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To10 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CodedAttributeType"`

	// Fee/reduction Id
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`
}

type CommercialAgreementsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CommercialAgreementsType"`

	// Codeshare Details
	CodeshareDetails *CompanyRoleIdentificationType `xml:"codeshareDetails,omitempty"`

	// Other codeshare details
	OtherCodeshareDetails *CompanyRoleIdentificationType `xml:"otherCodeshareDetails,omitempty"`
}

type CompanyIdentificationTextType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyIdentificationTextType"`

	// Company Id Text reference.
	TextRefNumber formats.NumericInteger_Length0To4 `xml:"textRefNumber,omitempty"`

	// Company id free text.
	CompanyText formats.AlphaNumericString_Length0To70 `xml:"companyText,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyIdentificationTypeI"`

	// Company
	MarketingCompany formats.AlphaNumericString_Length2To3 `xml:"marketingCompany,omitempty"`

	// Company
	OperatingCompany formats.AlphaNumericString_Length2To3 `xml:"operatingCompany,omitempty"`

	// Company
	OtherCompany formats.AlphaNumericString_Length2To3 `xml:"otherCompany,omitempty"`
}

type CompanyIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyIdentificationType"`

	// Marketing carrier
	MarketingCarrier formats.AlphaNumericString_Length2To3 `xml:"marketingCarrier,omitempty"`

	// Operating carrier
	OperatingCarrier formats.AlphaNumericString_Length2To3 `xml:"operatingCarrier,omitempty"`

	// airline alliance code
	Alliance formats.AlphaNumericString_Length1To2 `xml:"alliance,omitempty"`
}

type CompanyRoleIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyRoleIdentificationType"`

	// Type of code share agreement.
	CodeShareType formats.AlphaString_Length1To1 `xml:"codeShareType,omitempty"`

	// company identification
	AirlineDesignator formats.AlphaNumericString_Length2To3 `xml:"airlineDesignator,omitempty"`

	// flight number
	FlightNumber formats.NumericInteger_Length1To4 `xml:"flightNumber,omitempty"`
}

type CompanyRoleIdentificationType_120771C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyRoleIdentificationType_120771C"`

	// Type of code share agreement.
	TransportStageQualifier formats.AlphaNumericString_Length1To3 `xml:"transportStageQualifier,omitempty"`

	// company identification
	Company formats.AlphaNumericString_Length2To3 `xml:"company,omitempty"`
}

type ConversionRateDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ConversionRateDetailsTypeI"`

	// Conversion type
	ConversionType formats.AlphaNumericString_Length1To3 `xml:"conversionType,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// amount
	Amount formats.AlphaNumericString_Length0To12 `xml:"amount,omitempty"`
}

type ConversionRateDetailsTypeI_179848C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ConversionRateDetailsTypeI_179848C"`

	// Conversion type
	ConversionType formats.AlphaNumericString_Length1To3 `xml:"conversionType,omitempty"`

	// Currency
	Currency formats.AlphaString_Length1To3 `xml:"currency,omitempty"`

	// Conversion rate for pricing
	Rate formats.AlphaNumericString_Length0To18 `xml:"rate,omitempty"`

	// Converted value amount
	ConvertedAmountLink formats.AlphaNumericString_Length0To18 `xml:"convertedAmountLink,omitempty"`

	// Applicable ISO country code or Tax designator code.
	TaxQualifier formats.AlphaNumericString_Length0To3 `xml:"taxQualifier,omitempty"`
}

type ConversionRateTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ConversionRateTypeI"`

	// Detail of conversion rate of First Monetary Unit.
	ConversionRateDetail *ConversionRateDetailsTypeI_179848C `xml:"conversionRateDetail,omitempty"`
}

type ConversionRateTypeI_78562S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ConversionRateTypeI_78562S"`

	// Details of conversion
	ConversionRateDetail *ConversionRateDetailsTypeI `xml:"conversionRateDetail,omitempty"`
}

type CriteriaiDetaislType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CriteriaiDetaislType"`

	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	Value formats.AlphaNumericString_Length1To18 `xml:"value,omitempty"`
}

type DataInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DataInformationType"`

	// Ancillary services options
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type DataTypeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DataTypeInformationType"`

	// service group/sub-group/sub-code information
	SubType formats.AlphaNumericString_Length1To3 `xml:"subType,omitempty"`

	// Status (automated, manually added, exempted). Default is automated
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`
}

type DateAndTimeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateAndTimeDetailsType"`

	// Date time period qualifier
	DateQualifier formats.AlphaNumericString_Length1To3 `xml:"dateQualifier,omitempty"`

	// First Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// First Time
	FirstTime formats.Time24_HHMM `xml:"firstTime,omitempty"`

	// Movement type.
	EquipementType formats.AlphaNumericString_Length1To3 `xml:"equipementType,omitempty"`

	// Place/location identification.
	LocationId formats.AlphaNumericString_Length3To5 `xml:"locationId,omitempty"`
}

type DateAndTimeDetailsType_256192C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateAndTimeDetailsType_256192C"`

	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	Date formats.AlphaNumericString_Length1To35 `xml:"date,omitempty"`

	// Time
	Time formats.Time24_HHMM `xml:"time,omitempty"`

	// Location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`
}

type DateAndTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateAndTimeInformationType"`

	// Details on date and time
	StopDetails *DateAndTimeDetailsType `xml:"stopDetails,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DummyNET struct{}

type DateAndTimeInformationType_182345S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateAndTimeInformationType_182345S"`

	// DATE AND TIME DETAILS
	DateAndTimeDetails *DateAndTimeDetailsType_256192C `xml:"dateAndTimeDetails,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DateTimePeriodDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateTimePeriodDetailsTypeI"`

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Value
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type DiscountAndPenaltyInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DiscountAndPenaltyInformationType"`

	// Used to specify airline collected fee or agent collected fee.
	FeeIdentification formats.AlphaNumericString_Length1To3 `xml:"feeIdentification,omitempty"`

	// Used to specify penalty information
	FeeInformation *DiscountPenaltyMonetaryInformationType `xml:"feeInformation,omitempty"`
}

type DiscountPenaltyInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DiscountPenaltyInformationType"`

	// Discounted fare,...
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier,omitempty"`

	// Dicount code,...
	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// Percentage
	Percentage formats.NumericDecimal_Length1To8 `xml:"percentage,omitempty"`
}

type DiscountPenaltyMonetaryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DiscountPenaltyMonetaryInformationType"`

	// Type of discount/penalty
	FeeType formats.AlphaNumericString_Length1To3 `xml:"feeType,omitempty"`

	// The amount Type can be a percentage or an amount
	FeeAmountType formats.AlphaNumericString_Length1To3 `xml:"feeAmountType,omitempty"`

	// specify the value
	FeeAmount formats.NumericDecimal_Length1To18 `xml:"feeAmount,omitempty"`

	// Fee currency code.
	FeeCurrency formats.AlphaNumericString_Length1To3 `xml:"feeCurrency,omitempty"`
}

type DummySegmentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DummySegmentTypeI"`
}

type ExcessBaggageType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ExcessBaggageType"`

	// Baggage details
	BaggageDetails *BaggageDetailsType `xml:"baggageDetails,omitempty"`

	// Free baggage allowance details
	BagTagDetails *BagtagDetailsType `xml:"bagTagDetails,omitempty"`
}

type FareCalculationCodeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareCalculationCodeDetailsType"`

	// Qualifier of the amout or rate
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`

	// Location code
	LocationCode formats.AlphaNumericString_Length1To3 `xml:"locationCode,omitempty"`

	// Other location code
	OtherLocationCode formats.AlphaNumericString_Length1To3 `xml:"otherLocationCode,omitempty"`

	// Rate
	Rate formats.NumericDecimal_Length1To8 `xml:"rate,omitempty"`
}

type FareCategoryCodesTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareCategoryCodesTypeI"`

	// Fare type
	FareType formats.AlphaNumericString_Length1To20 `xml:"fareType,omitempty"`

	// Other fare type
	OtherFareType formats.AlphaNumericString_Length1To20 `xml:"otherFareType,omitempty"`
}

type FareDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareDetailsTypeI"`

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Rate
	Rate formats.NumericDecimal_Length1To8 `xml:"rate,omitempty"`

	// Country
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`

	// Fare category
	FareCategory formats.AlphaNumericString_Length1To3 `xml:"fareCategory,omitempty"`
}

type FareDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareDetailsType"`

	// Passenger Type qualifier
	PassengerTypeQualifier formats.AlphaNumericString_Length1To3 `xml:"passengerTypeQualifier,omitempty"`
}

type FareDetailsType_193037C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareDetailsType_193037C"`

	// Qualifier
	Qualifier formats.AMA_EDICodesetType_Length1to3 `xml:"qualifier,omitempty"`

	// Rate
	Rate formats.NumericInteger_Length1To8 `xml:"rate,omitempty"`

	// Country
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`

	// Fare Category
	FareCategory formats.AMA_EDICodesetType_Length1to3 `xml:"fareCategory,omitempty"`
}

type FareFamilyDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareFamilyDetailsType"`

	// Commercial fare Family Short name
	CommercialFamily formats.AlphaNumericString_Length1To10 `xml:"commercialFamily,omitempty"`
}

type FareFamilyType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareFamilyType"`

	// Fare Family Reference Number
	RefNumber formats.NumericInteger_Length1To3 `xml:"refNumber,omitempty"`

	// Fare Family Short Name
	FareFamilyname formats.AlphaNumericString_Length1To10 `xml:"fareFamilyname,omitempty"`

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy formats.NumericInteger_Length1To4 `xml:"hierarchy,omitempty"`

	// CABIN USED FOR FARE FAMILY
	Cabin formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`

	// Indicates Commercial Fare Family Short names
	CommercialFamilyDetails *FareFamilyDetailsType `xml:"commercialFamilyDetails,omitempty"`

	// Short description of the fare family
	Description formats.AlphaNumericString_Length1To100 `xml:"description,omitempty"`

	// Carrier code
	Carrier formats.AlphaNumericString_Length2To3 `xml:"carrier,omitempty"`

	// Reference to the services details
	Services *ServicesReferences `xml:"services,omitempty"`
}

type FareInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareInformationTypeI"`

	// Value qualifier
	ValueQualifier formats.AlphaNumericString_Length1To3 `xml:"valueQualifier,omitempty"`

	// Value
	Value formats.NumericInteger_Length1To15 `xml:"value,omitempty"`
}

type FareInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareInformationType"`

	// Value Qualifier
	ValueQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"valueQualifier,omitempty"`

	// Value
	Value formats.NumericInteger_Length1To15 `xml:"value,omitempty"`

	// Fare Details
	FareDetails *FareDetailsType_193037C `xml:"fareDetails,omitempty"`

	// Identity Number
	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber,omitempty"`

	// Fare Type Grouping
	FareTypeGrouping *FareTypeGroupingInformationType `xml:"fareTypeGrouping,omitempty"`

	// Rate Category
	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"`
}

type FareInformationType_80868S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareInformationType_80868S"`

	// Fare details
	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`
}

type FareProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareProductDetailsType"`

	// Fare basis code
	FareBasis formats.AlphaNumericString_Length1To18 `xml:"fareBasis,omitempty"`
}

type FareProductDetailsType_248552C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareProductDetailsType_248552C"`

	// Fare basis code
	FareBasis formats.AlphaNumericString_Length0To18 `xml:"fareBasis,omitempty"`

	// PTC priced
	PassengerType formats.AlphaNumericString_Length1To6 `xml:"passengerType,omitempty"`

	// Type of fare
	FareType formats.AlphaNumericString_Length0To3 `xml:"fareType,omitempty"`
}

type FareQualifierDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareQualifierDetailsType"`

	// Route Code
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`

	// Fare categories
	FareCategories *FareCategoryCodesTypeI `xml:"fareCategories,omitempty"`

	// Fare details
	FareDetails *FareDetailsTypeI `xml:"fareDetails,omitempty"`

	// Additional fare details
	AdditionalFareDetails *AdditionalFareQualifierDetailsTypeI `xml:"additionalFareDetails,omitempty"`

	// Discount details
	DiscountDetails *DiscountPenaltyInformationType `xml:"discountDetails,omitempty"`
}

type FareTypeGroupingInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareTypeGroupingInformationType"`

	// Pricing Group
	PricingGroup formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"`
}

type FlightCharacteristicsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightCharacteristicsType"`

	// On-Time Performance
	OnTimePerformance *OnTimePerformanceType `xml:"onTimePerformance,omitempty"`

	// In flight services
	InFlightSrv formats.AlphaNumericString_Length1To3 `xml:"inFlightSrv,omitempty"`
}

type FlightProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightProductInformationType"`

	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType_195516C `xml:"cabinProduct,omitempty"`

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType `xml:"contextDetails,omitempty"`
}

type FlightProductInformationType_141442S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightProductInformationType_141442S"`

	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType_205138C `xml:"cabinProduct,omitempty"`

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType_205137C `xml:"contextDetails,omitempty"`
}

type FlightProductInformationType_161491S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightProductInformationType_161491S"`

	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType_229142C `xml:"cabinProduct,omitempty"`

	// Fare product details
	FareProductDetail *FareProductDetailsType `xml:"fareProductDetail,omitempty"`
}

type FlightProductInformationType_176659S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightProductInformationType_176659S"`

	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType `xml:"cabinProduct,omitempty"`

	// Fare product details
	FareProductDetail *FareProductDetailsType_248552C `xml:"fareProductDetail,omitempty"`

	// Corporate number or name and number
	CorporateId formats.AlphaNumericString_Length1To20 `xml:"corporateId,omitempty"`

	// To determine if Fare Breaks at this segment
	BreakPoint formats.AlphaString_Length1To1 `xml:"breakPoint,omitempty"`

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType `xml:"contextDetails,omitempty"`
}

type FlightServicesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightServicesType"`

	// Type of service used
	ServiceType formats.AlphaNumericString_Length1To3 `xml:"serviceType,omitempty"`

	CabinInfo *CabinInformationType `xml:"cabinInfo,omitempty"`

	ClassInfo *ClassInformationType `xml:"classInfo,omitempty"`
}

type FreeTextQualificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FreeTextQualificationTypeI"`

	// Text subject qualifier
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`
}

type FreeTextQualificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FreeTextQualificationType"`

	// Type of message
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Coded Text or type of information in 4440 (e.g. type of OSI or free text, canned message value)
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`
}

type FreeTextQualificationType_120769C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FreeTextQualificationType_120769C"`

	// Type of message
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Coded Text or type of information in 4440 (e.g. type of OSI or free text, canned message value)
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// ISO code for language of free text (default is English)
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`
}

type FrequentTravellerIdentificationCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FrequentTravellerIdentificationCodeType"`

	// Frequent Traveller Info
	FrequentTravellerDetails *FrequentTravellerIdentificationType `xml:"frequentTravellerDetails,omitempty"`
}

type FrequentTravellerIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FrequentTravellerIdentificationType"`

	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"`

	// Number
	Number formats.AlphaNumericString_Length1To28 `xml:"number,omitempty"`

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`

	// For example : priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"`
}

type HeaderInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A HeaderInformationTypeI"`

	// Status
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Date and Time info
	DateTimePeriodDetails *DateTimePeriodDetailsTypeI `xml:"dateTimePeriodDetails,omitempty"`

	// Reference number
	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber,omitempty"`

	// Contains product identification such as UIC code...
	ProductIdentification formats.AlphaNumericString_Length1To35 `xml:"productIdentification,omitempty"`
}

type InteractiveFreeTextType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A InteractiveFreeTextType"`

	// Free text qualification
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`

	// Free text
	FreeText formats.AlphaNumericString_Length1To50 `xml:"freeText,omitempty"`
}

type InteractiveFreeTextType_78534S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A InteractiveFreeTextType_78534S"`

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType `xml:"freeTextQualification,omitempty"`

	// Free text
	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`
}

type InteractiveFreeTextType_78544S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A InteractiveFreeTextType_78544S"`

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType_120769C `xml:"freeTextQualification,omitempty"`

	// Free text
	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`
}

type InteractiveFreeTextType_78559S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A InteractiveFreeTextType_78559S"`

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType_120769C `xml:"freeTextQualification,omitempty"`

	// Free text
	Description formats.AlphaNumericString_Length1To500 `xml:"description,omitempty"`
}

type ItemNumberIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType"`

	// Ancillary Service number
	Number formats.AlphaNumericString_Length1To4 `xml:"number,omitempty"`

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Responsible agency
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`
}

type ItemNumberIdentificationType_191597C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType_191597C"`

	// Item number.
	Number formats.AlphaNumericString_Length1To6 `xml:"number,omitempty"`

	// Indicates the item type .
	NumberType formats.AlphaNumericString_Length0To3 `xml:"numberType,omitempty"`
}

type ItemNumberIdentificationType_192331C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType_192331C"`

	// Service coverage number
	Number formats.AlphaNumericString_Length1To6 `xml:"number,omitempty"`

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Responsible agency
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`
}

type ItemNumberIdentificationType_234878C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType_234878C"`

	// Number
	Number formats.NumericInteger_Length1To6 `xml:"number,omitempty"`

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type ItemNumberIdentificationType_248537C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType_248537C"`

	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`
}

type ItemNumberType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType"`

	// Item number details
	ItemNumber *ItemNumberIdentificationType_192331C `xml:"itemNumber,omitempty"`
}

type ItemNumberType_161497S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType_161497S"`

	// Indicates the recommendation number.
	ItemNumberId *ItemNumberIdentificationType_191597C `xml:"itemNumberId,omitempty"`

	// Code share details.
	CodeShareDetails *CompanyRoleIdentificationType_120771C `xml:"codeShareDetails,omitempty"`

	// Pricing ticketind details.
	PriceTicketing *PricingTicketingInformationType `xml:"priceTicketing,omitempty"`
}

type ItemNumberType_166130S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType_166130S"`

	// Item number details
	ItemNumberDetails *ItemNumberIdentificationType_234878C `xml:"itemNumberDetails,omitempty"`
}

type ItemNumberType_176648S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType_176648S"`

	ItemNumberDetails *ItemNumberIdentificationType_248537C `xml:"itemNumberDetails,omitempty"`
}

type ItemNumberType_80866S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType_80866S"`

	// Item number details
	ItemNumberDetails *ItemNumberIdentificationType `xml:"itemNumberDetails,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemReferencesAndVersionsType"`

	// Qualifies the type of the reference used.
	ReferenceType formats.AlphaNumericString_Length1To6 `xml:"referenceType,omitempty"`

	// Unique fee reference.
	RefNumber formats.NumericInteger_Length1To3 `xml:"refNumber,omitempty"`
}

type ItemReferencesAndVersionsType_78536S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemReferencesAndVersionsType_78536S"`

	// Qualifies the type of the reference used.
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// Unique fee reference.
	RefNumber formats.NumericInteger_Length1To3 `xml:"refNumber,omitempty"`
}

type ItemReferencesAndVersionsType_78564S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemReferencesAndVersionsType_78564S"`

	// Qualifies the type of the reference used.
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// Unique fee reference.
	FeeRefNumber formats.NumericInteger_Length1To3 `xml:"feeRefNumber,omitempty"`
}

type ItineraryDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItineraryDetailsType"`

	// Airport/City Qualifier: the passenger wants to depart/arrive from/to the same airport or city in the specified requested segment
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`

	// Requested segment number
	SegmentNumber formats.NumericInteger_Length1To3 `xml:"segmentNumber,omitempty"`
}

type LocationIdentificationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A LocationIdentificationDetailsType"`

	// 3 characters ATA/IATA airport/city code
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Airport/city qualifier: the requested point is an airport when ambiguity exists (e.g. HOU)
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`

	// Terminal information
	Terminal formats.AlphaNumericString_Length1To5 `xml:"terminal,omitempty"`
}

type MiniRulesDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MiniRulesDetailsType"`

	// Coded text (period or day)
	Interpretation formats.AlphaString_Length0To9 `xml:"interpretation,omitempty"`

	// Data type coded or value of interpretation
	Value formats.AlphaNumericString_Length0To5 `xml:"value,omitempty"`
}

type MiniRulesIndicatorType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MiniRulesIndicatorType"`

	// See rule indicator and free form text indicator
	RuleIndicator formats.AlphaString_Length1To1 `xml:"ruleIndicator,omitempty"`
}

type MiniRulesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MiniRulesType"`

	// Categoty of restriction: PTC, Max Adv Pur, Days, ...
	Category formats.AlphaString_Length1To3 `xml:"category,omitempty"`
}

type MiniRulesType_78547S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MiniRulesType_78547S"`

	// Type of restriction: PTC, Max Adv Res, Max Ticketing After Res, ...
	RestrictionType formats.AlphaNumericString_Length0To6 `xml:"restrictionType,omitempty"`

	// Categoty of restriction: PTC, Max Adv Pur, Days, ...
	Category formats.AlphaString_Length0To3 `xml:"category,omitempty"`

	// Indicators
	Indicator *MiniRulesIndicatorType `xml:"indicator,omitempty"`

	// Mini rules
	MiniRules *MiniRulesDetailsType `xml:"miniRules,omitempty"`
}

type MonetaryInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationDetailsTypeI"`

	// type Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To6 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount formats.AlphaNumericString_Length1To35 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationDetailsType"`

	// To specify amount and percentage.
	AmountType formats.AlphaNumericString_Length0To3 `xml:"amountType,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// ISO currency code
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsType_245528C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationDetailsType_245528C"`

	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To35 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`
}

type MonetaryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationTypeI"`

	// Monetary details
	MonetaryDetails *MonetaryInformationDetailsTypeI `xml:"monetaryDetails,omitempty"`
}

type MonetaryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationType"`

	// Monetary information.
	MonetaryDetail *MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"`
}

type MonetaryInformationType_137835S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationType_137835S"`

	// Monetary information.
	MonetaryDetail *MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"`
}

type MonetaryInformationType_174241S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationType_174241S"`

	MonetaryDetails *MonetaryInformationDetailsType_245528C `xml:"monetaryDetails,omitempty"`

	OtherMonetaryDetails *MonetaryInformationDetailsType_245528C `xml:"otherMonetaryDetails,omitempty"`
}

type MonetaryInformationType_185955S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationType_185955S"`

	// Monetary information
	MonetaryDetail *MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"`
}

type OnTimePerformanceType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A OnTimePerformanceType"`

	// Date time period
	DateTimePeriod formats.AlphaNumericString_Length1To35 `xml:"dateTimePeriod,omitempty"`

	// Percentage
	Percentage formats.NumericDecimal_Length1To8 `xml:"percentage,omitempty"`

	// Accuracy
	Accuracy formats.AlphaNumericString_Length1To3 `xml:"accuracy,omitempty"`
}

type OriginAndDestinationRequestType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A OriginAndDestinationRequestType"`

	// Requested segment number
	SegRef formats.NumericInteger_Length1To2 `xml:"segRef,omitempty"`

	// Forces arrival or departure, from/to the same airport/city
	LocationForcing *ItineraryDetailsType `xml:"locationForcing,omitempty"`
}

type OriginAndDestinationRequestType_134833S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A OriginAndDestinationRequestType_134833S"`

	// Requested segment number
	SegRef formats.NumericInteger_Length1To2 `xml:"segRef,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A OriginatorIdentificationDetailsTypeI"`

	// Office Name.
	OfficeName formats.NumericInteger_Length1To9 `xml:"officeName,omitempty"`

	// Agent Sign In .
	AgentSignin formats.AlphaNumericString_Length1To9 `xml:"agentSignin,omitempty"`

	// Confidential Office Name.
	ConfidentialOffice formats.AlphaNumericString_Length1To9 `xml:"confidentialOffice,omitempty"`

	// Other Office Name
	OtherOffice formats.AlphaNumericString_Length1To9 `xml:"otherOffice,omitempty"`
}

type PricingTicketingInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A PricingTicketingInformationType"`

	// Price type qualifier
	PriceType formats.AlphaNumericString_Length0To3 `xml:"priceType,omitempty"`
}

type PricingTicketingSubsequentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A PricingTicketingSubsequentType"`

	// Passenger fare product number
	PaxFareNum formats.AlphaNumericString_Length1To3 `xml:"paxFareNum,omitempty"`
}

type PricingTicketingSubsequentType_144401S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A PricingTicketingSubsequentType_144401S"`

	// Passenger fare product number
	PaxFareNum formats.AlphaNumericString_Length1To3 `xml:"paxFareNum,omitempty"`

	// Total fare amount
	TotalFareAmount formats.NumericDecimal_Length1To18 `xml:"totalFareAmount,omitempty"`

	// Total tax amount
	TotalTaxAmount formats.NumericDecimal_Length1To18 `xml:"totalTaxAmount,omitempty"`

	// Code share details.
	CodeShareDetails *CompanyRoleIdentificationType_120771C `xml:"codeShareDetails,omitempty"`

	// Monetary information.
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails,omitempty"`

	// Pricing ticketing details.
	PricingTicketing *PricingTicketingInformationType `xml:"pricingTicketing,omitempty"`
}

type ProcessingInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProcessingInformationType"`

	// Action qualifier
	ActionQualifier formats.AlphaNumericString_Length1To3 `xml:"actionQualifier,omitempty"`

	// Reference qualifier
	ReferenceQualifier formats.AlphaNumericString_Length1To3 `xml:"referenceQualifier,omitempty"`

	// Reference number
	RefNum formats.AlphaNumericString_Length1To6 `xml:"refNum,omitempty"`
}

type ProductDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductDateTimeType"`

	// Departure date
	DateOfDeparture formats.Date_DDMMYY `xml:"dateOfDeparture,omitempty"`

	// Departure time
	TimeOfDeparture formats.Time24_HHMM `xml:"timeOfDeparture,omitempty"`

	// Arrival date
	DateOfArrival formats.Date_DDMMYY `xml:"dateOfArrival,omitempty"`

	// Arrival time
	TimeOfArrival formats.Time24_HHMM `xml:"timeOfArrival,omitempty"`

	// Arrival date compared to departure date, only if different from 0
	DateVariation formats.NumericInteger_Length1To1 `xml:"dateVariation,omitempty"`
}

type ProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductDetailsType"`

	Designator formats.AlphaNumericString_Length1To17 `xml:"designator,omitempty"`

	AvailabilityStatus formats.AlphaNumericString_Length1To3 `xml:"availabilityStatus,omitempty"`

	SpecialService formats.AlphaNumericString_Length1To3 `xml:"specialService,omitempty"`

	Option formats.AlphaNumericString_Length1To7 `xml:"option,omitempty"`
}

type ProductFacilitiesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductFacilitiesType"`

	// Yes-No indicator whether Last Seat Available
	LastSeatAvailable formats.AlphaString_Length1To1 `xml:"lastSeatAvailable,omitempty"`

	// Level of access
	LevelOfAccess formats.AlphaNumericString_Length1To3 `xml:"levelOfAccess,omitempty"`

	// Yes-No indicator whether electronic ticketing
	ElectronicTicketing formats.AlphaString_Length1To1 `xml:"electronicTicketing,omitempty"`

	// Product identification suffix
	OperationalSuffix formats.AlphaString_Length1To1 `xml:"operationalSuffix,omitempty"`

	// Define whether a flight has been polled or not
	ProductDetailQualifier formats.AlphaNumericString_Length1To3 `xml:"productDetailQualifier,omitempty"`

	// Add some flight restrictions (See code set list)
	FlightCharacteristic formats.AlphaNumericString_Length1To3 `xml:"flightCharacteristic,omitempty"`
}

type ProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductInformationType"`

	// value of the Qualifier: INT for International DOM for Domestic EUR for European  otherwise CM#10569 INVALID INTERNATIONAL INDICATOR is returned.
	ProductDetailsQualifier formats.AlphaNumericString_Length1To3 `xml:"productDetailsQualifier,omitempty"`

	BookingClassDetails *ProductDetailsType `xml:"bookingClassDetails,omitempty"`
}

type ProductTypeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductTypeDetailsType"`

	// Availability connection type.
	AvailabilityCnxType formats.AlphaNumericString_Length1To3 `xml:"availabilityCnxType,omitempty"`
}

type ProductTypeDetailsType_205137C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductTypeDetailsType_205137C"`

	// indicates whether the flight is domestic or international
	Avl formats.AlphaNumericString_Length1To6 `xml:"avl,omitempty"`
}

type ProposedSegmentDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProposedSegmentDetailsType"`

	// Flight proposal reference
	Ref formats.AlphaNumericString_Length1To6 `xml:"ref,omitempty"`

	// Elapse Flying Time
	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`
}

type ProposedSegmentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProposedSegmentType"`

	// Parameters for proposed flight group
	FlightProposal *ProposedSegmentDetailsType `xml:"flightProposal,omitempty"`

	// Flight characteristics.
	FlightCharacteristic formats.AlphaNumericString_Length0To3 `xml:"flightCharacteristic,omitempty"`

	// Majority cabin
	MajCabin formats.AlphaString_Length1To1 `xml:"majCabin,omitempty"`
}

type ReferenceInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType"`

	// Referencing details
	ReferencingDetail *ReferencingDetailsType_191583C `xml:"referencingDetail,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type ReferenceInfoType_133176S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_133176S"`

	// Referencing details
	ReferencingDetail *ReferencingDetailsType `xml:"referencingDetail,omitempty"`
}

type ReferenceInfoType_134839S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_134839S"`

	// Referencing details
	ReferencingDetail *ReferencingDetailsType_195561C `xml:"referencingDetail,omitempty"`
}

type ReferenceInfoType_134840S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_134840S"`

	// Referencing details
	ReferencingDetail *ReferencingDetailsType_195561C `xml:"referencingDetail,omitempty"`
}

type ReferenceInfoType_165972S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_165972S"`

	// Reference details
	ReferenceDetails *ReferencingDetailsType_234704C `xml:"referenceDetails,omitempty"`
}

type ReferenceInfoType_176658S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_176658S"`

	// Referencing details
	ReferencingDetail *ReferencingDetailsType `xml:"referencingDetail,omitempty"`
}

type ReferenceType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceType"`

	// Reference  of leg
	RefOfLeg formats.AlphaNumericString_Length1To6 `xml:"refOfLeg,omitempty"`

	// Reference of segment starting range
	FirstItemIdentifier formats.NumericInteger_Length1To3 `xml:"firstItemIdentifier,omitempty"`

	// Reference of segment ending range
	LastItemIdentifier formats.NumericInteger_Length1To3 `xml:"lastItemIdentifier,omitempty"`
}

type ReferencingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferencingDetailsType"`

	// Reference qualifier
	RefQualifier formats.AlphaNumericString_Length0To3 `xml:"refQualifier,omitempty"`

	// Requested segment reference
	RefNumber formats.NumericInteger_Length0To3 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsType_191583C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferencingDetailsType_191583C"`

	// Service reference qualifier
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Service reference
	RefNumber formats.NumericInteger_Length0To6 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsType_195561C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferencingDetailsType_195561C"`

	// Segment reference qualifier
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Flight or flight group reference
	RefNumber formats.NumericInteger_Length0To3 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsType_234704C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferencingDetailsType_234704C"`

	// Type
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	// Value
	Value formats.AlphaNumericString_Length1To60 `xml:"value,omitempty"`
}

type SegmentRepetitionControlDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SegmentRepetitionControlDetailsTypeI"`

	// traveller number
	Quantity formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"`

	// range of traveller
	NumberOfUnits formats.NumericInteger_Length1To15 `xml:"numberOfUnits,omitempty"`
}

type SegmentRepetitionControlTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SegmentRepetitionControlTypeI"`

	// Segment control details
	SegmentControlDetails *SegmentRepetitionControlDetailsTypeI `xml:"segmentControlDetails,omitempty"`
}

type SelectionDetailsInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SelectionDetailsInformationType"`

	// Carrier fee type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Carrier fee status
	OptionInformation formats.AlphaNumericString_Length1To3 `xml:"optionInformation,omitempty"`
}

type SelectionDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SelectionDetailsType"`

	// Carrier fees options
	CarrierFeeDetails *SelectionDetailsInformationType `xml:"carrierFeeDetails,omitempty"`
}

type SequenceDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SequenceDetailsTypeU"`

	// Sequence details
	SequenceDetails *SequenceInformationTypeU `xml:"sequenceDetails,omitempty"`
}

type SequenceInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SequenceInformationTypeU"`

	// Number
	Number formats.AlphaNumericString_Length1To10 `xml:"number,omitempty"`

	// Identification code
	IdentificationCode formats.AlphaNumericString_Length1To17 `xml:"identificationCode,omitempty"`
}

type ServicesReferences struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ServicesReferences"`

	// Reference of the service
	Reference formats.AlphaNumericString_Length1To4 `xml:"reference,omitempty"`

	// Status of the service
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`
}

type SpecialRequirementsDataDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecialRequirementsDataDetailsType"`

	// SSR seat characteristic
	SeatCharacteristics formats.AlphaNumericString_Length1To2 `xml:"seatCharacteristics,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type SpecialRequirementsDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecialRequirementsDetailsType"`

	// To specify the Service Requirement of the customer
	ServiceRequirementsInfo *SpecialRequirementsTypeDetailsType `xml:"serviceRequirementsInfo,omitempty"`

	// Seat details
	SeatDetails *SpecialRequirementsDataDetailsType `xml:"seatDetails,omitempty"`
}

type SpecialRequirementsTypeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecialRequirementsTypeDetailsType"`

	// To specify the Service Classification of the Service Requirement.
	ServiceClassification formats.AlphaNumericString_Length1To4 `xml:"serviceClassification,omitempty"`

	// Status
	ServiceStatus formats.AlphaNumericString_Length1To3 `xml:"serviceStatus,omitempty"`

	// To specify the number of items involved
	ServiceNumberOfInstances formats.NumericInteger_Length1To15 `xml:"serviceNumberOfInstances,omitempty"`

	// To specify to which marketing carrier the service applies
	ServiceMarketingCarrier formats.AlphaNumericString_Length1To3 `xml:"serviceMarketingCarrier,omitempty"`

	// Specify the Service group
	ServiceGroup formats.AlphaNumericString_Length1To3 `xml:"serviceGroup,omitempty"`

	// Specify the Service Sub-Group
	ServiceSubGroup formats.AlphaNumericString_Length1To3 `xml:"serviceSubGroup,omitempty"`

	// Free Text attached to the Service.
	ServiceFreeText formats.AlphaNumericString_Length1To70 `xml:"serviceFreeText,omitempty"`
}

type SpecificDataInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecificDataInformationType"`

	// Carrier fee description
	DataTypeInformation *DataTypeInformationType `xml:"dataTypeInformation,omitempty"`

	// Data information
	DataInformation *DataInformationType `xml:"dataInformation,omitempty"`
}

type SpecificTravellerDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecificTravellerDetailsType"`

	// Reference number
	ReferenceNumber formats.AlphaNumericString_Length1To10 `xml:"referenceNumber,omitempty"`
}

type SpecificTravellerType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecificTravellerType"`

	// Traveller details
	TravellerDetails *SpecificTravellerDetailsType `xml:"travellerDetails,omitempty"`
}

type StatusDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A StatusDetailsType"`

	// Advisory type information, Fare Server
	AdvisoryTypeInfo formats.AlphaNumericString_Length1To3 `xml:"advisoryTypeInfo,omitempty"`

	// CPU time, user type
	Notification formats.AlphaNumericString_Length1To3 `xml:"notification,omitempty"`

	// CPU time,user type
	Notification2 formats.AlphaNumericString_Length1To3 `xml:"notification2,omitempty"`

	// Capture and trace information
	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`
}

type StatusDetailsType_256255C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A StatusDetailsType_256255C"`

	// list of status/qualifiers Either His for Historical or     Crt for Current
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	Action formats.AlphaNumericString_Length1To3 `xml:"action,omitempty"`
}

type StatusType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A StatusType"`

	// Status details
	Status *StatusDetailsType `xml:"status,omitempty"`
}

type StatusType_182386S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A StatusType_182386S"`

	// STATUS DETAILS
	StatusInformation *StatusDetailsType_256255C `xml:"statusInformation,omitempty"`
}

type TaxDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TaxDetailsType"`

	// Amount
	Rate formats.AlphaNumericString_Length1To12 `xml:"rate,omitempty"`

	// Country code
	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"`

	// Currency code
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Indicator
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

// TaxType ...
type TaxType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TaxType"`

	// Tax category
	TaxCategory formats.AlphaNumericString_Length1To3 `xml:"taxCategory,omitempty"`

	// Tax details
	TaxDetails *TaxDetailsType `xml:"taxDetails,omitempty"`
}

// TransportIdentifierType ...
type TransportIdentifierType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TransportIdentifierType"`

	// Company identification
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`
}

// TravelProductType ..
type TravelProductType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TravelProductType"`

	// Date and time of departure and arrival
	ProductDateTime *ProductDateTimeType `xml:"productDateTime,omitempty"`

	// Location of departure and arrival
	Location *LocationIdentificationDetailsType `xml:"location,omitempty"`

	CompanyID *CompanyIdentificationType `xml:"companyId,omitempty"`

	// Flight number or trainNumber
	FlightOrtrainNumber formats.AlphaNumericString_Length1To8 `xml:"flightOrtrainNumber,omitempty"`

	// Product details
	ProductDetail *AdditionalProductDetailsType `xml:"productDetail,omitempty"`

	// Additional product details
	AddProductDetail *ProductFacilitiesType `xml:"addProductDetail,omitempty"`

	// Attribute details
	AttributeDetails *CodedAttributeInformationType_260992C `xml:"attributeDetails,omitempty"`
}

// TravellerDetailsType ...
type TravellerDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TravellerDetailsType"`

	// Direct reference of passenger assigned by requesting system.
	Ref formats.NumericInteger_Length1To3 `xml:"ref,omitempty"`

	// Traveller is an infant
	InfantIndicator formats.NumericInteger_Length1To1 `xml:"infantIndicator,omitempty"`
}

// TravellerReferenceInformationType ...
type TravellerReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TravellerReferenceInformationType"`

	// Requested passenger type
	Ptc formats.AlphaNumericString_Length1To6 `xml:"ptc,omitempty"`

	// Traveller details
	Traveller *TravellerDetailsType `xml:"traveller,omitempty"`
}

// UserIdentificationType ...
type UserIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A UserIdentificationType"`

	// Originator Identification Details
	OfficeIdentification *OriginatorIdentificationDetailsTypeI `xml:"officeIdentification,omitempty"`

	// Used to specify which kind of info is given in DE 9900.
	OfficeType formats.AlphaNumericString_Length1To1 `xml:"officeType,omitempty"`

	// The code given to an agent by the originating reservation system.
	OfficeCode formats.AlphaNumericString_Length1To30 `xml:"officeCode,omitempty"`
}

type ValueSearchCriteriaType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ValueSearchCriteriaType"`

	Ref formats.AlphaNumericString_Length1To35 `xml:"ref,omitempty"`

	Value formats.AlphaNumericString_Length1To18 `xml:"value,omitempty"`

	CriteriaDetails *CriteriaiDetaislType `xml:"criteriaDetails,omitempty"`
}
