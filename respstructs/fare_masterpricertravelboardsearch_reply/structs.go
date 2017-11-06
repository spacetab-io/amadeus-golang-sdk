package fare_masterpricertravelboardsearch_reply

import (
	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type FareMasterPricerTravelBoardSearchReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A Fare_MasterPricerTravelBoardSearchReply"`

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

	OtherSolutions []*OtherSolutions `xml:"otherSolutions,omitempty"`  // minOccurs="0" maxOccurs="100009"

	WarningInfo []*WarningInfo `xml:"warningInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	GlobalInformation []*GlobalInformation `xml:"globalInformation,omitempty"`  // minOccurs="0" maxOccurs="9"

	ServiceFeesGrp []*ServiceFeesGrp `xml:"serviceFeesGrp,omitempty"`  // minOccurs="0" maxOccurs="3"

	Value []*ValueSearchCriteriaType `xml:"value,omitempty"`  // minOccurs="0" maxOccurs="100009"

	MnrGrp *MnrGrp `xml:"mnrGrp,omitempty"`  // minOccurs="0"
}

type ErrorMessage struct {

	// Application error details.
	ApplicationError *ApplicationErrorInformationType_78543S `xml:"applicationError,omitempty"`

	// Type of error message and free text
	ErrorMessageText *InteractiveFreeTextType_78544S `xml:"errorMessageText,omitempty"`  // minOccurs="0"
}

type AmountInfoForAllPax struct {

	// Itinerary amounts for all passengers
	ItineraryAmounts *MonetaryInformationType_137835S `xml:"itineraryAmounts,omitempty"`

	AmountsPerSgt []*AmountsPerSgt `xml:"amountsPerSgt,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type AmountsPerSgt struct {

	// Requested segment reference
	SgtRef *ReferenceInfoType_133176S `xml:"sgtRef,omitempty"`

	// Amounts : Issue total amount, issue taxes amount, non refundable taxes amount
	Amounts *MonetaryInformationType_137835S `xml:"amounts,omitempty"`  // minOccurs="0"
}

type AmountInfoPerPax struct {

	// Passenger references
	PaxRef *SpecificTravellerType `xml:"paxRef,omitempty"`

	// Passenger attributes : Infant indicator
	PaxAttributes *FareInformationType_80868S `xml:"paxAttributes,omitempty"`  // minOccurs="0"

	// Itinerary amounts information
	ItineraryAmounts *MonetaryInformationType_137835S `xml:"itineraryAmounts,omitempty"`

	AmountsPerSgt []*AmountsPerSgt `xml:"amountsPerSgt,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FeeDetails struct {

	// Fee/Reduction Reference number.
	FeeReference *ItemReferencesAndVersionsType_78564S `xml:"feeReference,omitempty"`

	// Fee/Reduction information.
	FeeInformation *DiscountAndPenaltyInformationType `xml:"feeInformation,omitempty"`  // minOccurs="0"

	// Fee/Reduction parameters.
	FeeParameters *AttributeType_78561S `xml:"feeParameters,omitempty"`  // minOccurs="0"

	// To specify conversion rate details
	ConvertedOrOriginalInfo *ConversionRateTypeI_78562S `xml:"convertedOrOriginalInfo,omitempty"`  // minOccurs="0"
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

	GroupOfFlights []*GroupOfFlights `xml:"groupOfFlights,omitempty"`  // maxOccurs="100000"
}

type GroupOfFlights struct {

	// To indicate parameters for proposed flight group.
	PropFlightGrDetail *ProposedSegmentType `xml:"propFlightGrDetail,omitempty"`

	FlightDetails []*FlightDetails `xml:"flightDetails,omitempty"`  // maxOccurs="4"
}

type FlightDetails struct {

	// Specification of details on the flight and posting availability
	FlightInformation *TravelProductType `xml:"flightInformation,omitempty"`

	// returns booking class and availability context
	AvlInfo []*FlightProductInformationType_141442S `xml:"avlInfo,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Details on Flight date, time and location of technical stop or change of gauge
	TechnicalStop []*DateAndTimeInformationType `xml:"technicalStop,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Code Share Agreement description for current flight.
	CommercialAgreement *CommercialAgreementsType `xml:"commercialAgreement,omitempty"`  // minOccurs="0"

	// Additional Info about flight, such as Reference number, and several options
	AddInfo *HeaderInformationTypeI `xml:"addInfo,omitempty"`  // minOccurs="0"

	// Flight characteristics
	FlightCharacteristics *FlightCharacteristicsType `xml:"flightCharacteristics,omitempty"`  // minOccurs="0"

	// Flight Services by cabin/rbd
	FlightServices []*FlightServicesType `xml:"flightServices,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type Recommendation struct {

	// Specification of the item number
	ItemNumber *ItemNumberType_161497S `xml:"itemNumber,omitempty"`

	// To describe type of recommendation
	WarningMessage []*InteractiveFreeTextType_78544S `xml:"warningMessage,omitempty"`  // minOccurs="0" maxOccurs="4"

	// Indicates the Fare family reference.
	FareFamilyRef *ReferenceInfoType_133176S `xml:"fareFamilyRef,omitempty"`  // minOccurs="0"

	// Recommendation Price and Taxes.
	RecPriceInfo *MonetaryInformationType `xml:"recPriceInfo,omitempty"`

	// Mini rules
	MiniRule []*MiniRulesType_78547S `xml:"miniRule,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Indicates reference of Flight or the fee reference valid for all pax (usage:start with the 1 possible Fee reference, then provide the segments references)
	SegmentFlightRef []*ReferenceInfoType `xml:"segmentFlightRef,omitempty"`  // minOccurs="0" maxOccurs="100009"

	RecommandationSegmentsFareDetails []*RecommandationSegmentsFareDetails `xml:"recommandationSegmentsFareDetails,omitempty"`  // minOccurs="0" maxOccurs="6"

	PaxFareProduct []*PaxFareProduct `xml:"paxFareProduct,omitempty"`  // maxOccurs="10"

	SpecificRecDetails []*SpecificRecDetails `xml:"specificRecDetails,omitempty"`  // minOccurs="0" maxOccurs="100000"
}

type RecommandationSegmentsFareDetails struct {

	// Reference and details about requested segments.
	RecommendationSegRef *OriginAndDestinationRequestType `xml:"recommendationSegRef,omitempty"`

	// Amounts per requested segment.
	SegmentMonetaryInformation *MonetaryInformationType_137835S `xml:"segmentMonetaryInformation,omitempty"`  // minOccurs="0"
}

type PaxFareProduct struct {

	// Passenger Fare Details.
	PaxFareDetail *PricingTicketingSubsequentType_144401S `xml:"paxFareDetail,omitempty"`

	// Indicates Fee references (usage: start with the 1 possible Fee reference, then provide the segments references)
	FeeRef *ReferenceInfoType_134839S `xml:"feeRef,omitempty"`  // minOccurs="0"

	// Passenger Reference
	PaxReference []*TravellerReferenceInformationType `xml:"paxReference,omitempty"`  // maxOccurs="6"

	// add tax details for each passenger of each recommendations.
	PassengerTaxDetails *TaxType `xml:"passengerTaxDetails,omitempty"`  // minOccurs="0"

	Fare []*Fare `xml:"fare,omitempty"`  // minOccurs="0" maxOccurs="7"

	FareDetails []*FareDetails `xml:"fareDetails,omitempty"`  // maxOccurs="6"
}

type Fare struct {

	// Last Date to Ticket, Penalties
	PricingMessage *InteractiveFreeTextType_78559S `xml:"pricingMessage,omitempty"`

	// Amount of penalties, Surcharges...
	MonetaryInformation *MonetaryInformationType_185955S `xml:"monetaryInformation,omitempty"`  // minOccurs="0"
}

type FareDetails struct {

	// Reference of the Requested Segment
	SegmentRef *OriginAndDestinationRequestType `xml:"segmentRef,omitempty"`

	GroupOfFares []*GroupOfFares `xml:"groupOfFares,omitempty"`  // minOccurs="0" maxOccurs="4"

	// Amounts per passenger per requested segment.
	PsgSegMonetaryInformation *MonetaryInformationType_137835S `xml:"psgSegMonetaryInformation,omitempty"`  // minOccurs="0"

	// Majority Cabin Info
	MajCabin []*ProductInformationType `xml:"majCabin,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type GroupOfFares struct {

	// Contains details of Flight and Fare
	ProductInformation *FlightProductInformationType_176659S `xml:"productInformation,omitempty"`

	// Fare calculation code details
	FareCalculationCodeDetails []*FareCalculationCodeDetailsType `xml:"fareCalculationCodeDetails,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Ticket designator, ticket code and fare basis.
	TicketInfos *FareQualifierDetailsType `xml:"ticketInfos,omitempty"`  // minOccurs="0"

	// Reference of Fare Family for each Fare Component
	FareFamiliesRef *ReferenceInfoType_176658S `xml:"fareFamiliesRef,omitempty"`  // minOccurs="0"
}

type SpecificRecDetails struct {

	// Recommendation details
	SpecificRecItem *ItemReferencesAndVersionsType `xml:"specificRecItem,omitempty"`

	SpecificProductDetails []*SpecificProductDetails `xml:"specificProductDetails,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type SpecificProductDetails struct {

	// Product details
	ProductReferences *PricingTicketingSubsequentType `xml:"productReferences,omitempty"`

	FareContextDetails []*FareContextDetails `xml:"fareContextDetails,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type FareContextDetails struct {

	// Reference of requested segment
	RequestedSegmentInfo *OriginAndDestinationRequestType_134833S `xml:"requestedSegmentInfo,omitempty"`

	CnxContextDetails []*CnxContextDetails `xml:"cnxContextDetails,omitempty"`  // minOccurs="0" maxOccurs="4"
}

type CnxContextDetails struct {

	// Fare connection context details
	FareCnxInfo *FlightProductInformationType `xml:"fareCnxInfo,omitempty"`
}

type OtherSolutions struct {

	// Reference to the current solution
	Reference *SequenceDetailsTypeU `xml:"reference,omitempty"`

	AmtGroup []*AmtGroup `xml:"amtGroup,omitempty"`  // minOccurs="0" maxOccurs="10"

	PsgInfo []*PsgInfo `xml:"psgInfo,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type AmtGroup struct {

	// reference to the current amount (per bound, per segment...)
	Ref *ReferenceInfoType_165972S `xml:"ref,omitempty"`

	// Amount Description
	Amount *MonetaryInformationTypeI `xml:"amount,omitempty"`  // minOccurs="0"
}

type PsgInfo struct {

	// passenger reference
	Ref *SegmentRepetitionControlTypeI `xml:"ref,omitempty"`

	// Passenger Description Info
	Description *FareInformationTypeI `xml:"description,omitempty"`  // minOccurs="0"

	// Passenger frequent traveler info
	FreqTraveller *FrequentTravellerIdentificationCodeType `xml:"freqTraveller,omitempty"`  // minOccurs="0"

	// amount per passenger or group of passenger
	Amount *MonetaryInformationTypeI `xml:"amount,omitempty"`  // minOccurs="0"

	// Fare description
	Fare *FlightProductInformationType_161491S `xml:"fare,omitempty"`  // minOccurs="0"

	// Additional Information
	Attribute *AttributeTypeU `xml:"attribute,omitempty"`  // minOccurs="0" maxOccurs="10"
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

	ServiceFeeRefGrp []*ServiceFeeRefGrp `xml:"serviceFeeRefGrp,omitempty"`  // minOccurs="0" maxOccurs="100000"

	ServiceCoverageInfoGrp []*ServiceCoverageInfoGrp `xml:"serviceCoverageInfoGrp,omitempty"`  // minOccurs="0" maxOccurs="100000"

	// Globalmessage marker
	GlobalMessageMarker *DummySegmentTypeI `xml:"globalMessageMarker,omitempty"`

	ServiceFeeInfoGrp []*ServiceFeeInfoGrp `xml:"serviceFeeInfoGrp,omitempty"`  // minOccurs="0" maxOccurs="100000"

	ServiceDetailsGrp []*ServiceDetailsGrpType `xml:"serviceDetailsGrp,omitempty"`  // minOccurs="0" maxOccurs="200"

	FreeBagAllowanceGrp []*FreeBagAllowanceGrp `xml:"freeBagAllowanceGrp,omitempty"`  // minOccurs="0" maxOccurs="100000"
}

type ServiceFeeRefGrp struct {

	// Reference of service fee global information
	RefInfo *ReferenceInfoType `xml:"refInfo,omitempty"`
}

type ServiceCoverageInfoGrp struct {

	// Item reference number for service coverage details
	ItemNumberInfo *ItemNumberType `xml:"itemNumberInfo,omitempty"`

	ServiceCovInfoGrp []*ServiceCovInfoGrp `xml:"serviceCovInfoGrp,omitempty"`  // minOccurs="0" maxOccurs="200"
}

type ServiceCovInfoGrp struct {

	// Passenger reference number
	PaxRefInfo *SpecificTravellerType `xml:"paxRefInfo,omitempty"`

	// Service coverage information at flight level Matched seat characteristics
	CoveragePerFlightsInfo []*ActionDetailsType `xml:"coveragePerFlightsInfo,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Carrier information
	CarrierInfo *TransportIdentifierType `xml:"carrierInfo,omitempty"`  // minOccurs="0"

	// Service reference number
	RefInfo *ReferenceInfoType_134840S `xml:"refInfo,omitempty"`  // minOccurs="0"
}

type ServiceFeeInfoGrp struct {

	// Item number details
	ItemNumberInfo *ItemNumberType `xml:"itemNumberInfo,omitempty"`

	ServiceDetailsGrp []*ServiceDetailsGrp `xml:"serviceDetailsGrp,omitempty"`  // minOccurs="0" maxOccurs="200"
}

type ServiceDetailsGrp struct {

	// Service reference number
	RefInfo *ReferenceInfoType_134840S `xml:"refInfo,omitempty"`

	ServiceMatchedInfoGroup []*ServiceMatchedInfoGroup `xml:"serviceMatchedInfoGroup,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ServiceMatchedInfoGroup struct {

	// Reference on pax number
	PaxRefInfo *SpecificTravellerType `xml:"paxRefInfo,omitempty"`

	// Pricing oriented service matched information
	PricingInfo *FareInformationType_80868S `xml:"pricingInfo,omitempty"`  // minOccurs="0"

	// Informative Service amount
	AmountInfo *MonetaryInformationType `xml:"amountInfo,omitempty"`  // minOccurs="0"
}

type ServiceDetailsGrpType struct {

	// Service sub-code and options (exclusion,inclusion, mode pushed,polled)
	ServiceOptionInfo *SpecificDataInformationType `xml:"serviceOptionInfo,omitempty"`

	FeeDescriptionGrp *FeeDescriptionGrp `xml:"feeDescriptionGrp,omitempty"`  // minOccurs="0"
}

type FeeDescriptionGrp struct {

	// Specification of the item number
	ItemNumberInfo *ItemNumberType_80866S `xml:"itemNumberInfo,omitempty"`

	// Attributes  (SSR code EMD, RFIC, SSIM)
	ServiceAttributesInfo *AttributeType `xml:"serviceAttributesInfo,omitempty"`  // minOccurs="0"

	// Other service information (service description, ...)
	ServiceDescriptionInfo *SpecialRequirementsDetailsType `xml:"serviceDescriptionInfo,omitempty"`  // minOccurs="0"

	// Commercial name
	CommercialName *InteractiveFreeTextType `xml:"commercialName,omitempty"`  // minOccurs="0"
}

type FreeBagAllowanceGrp struct {

	// Free baggage allownce information
	FreeBagAllownceInfo *ExcessBaggageType `xml:"freeBagAllownceInfo,omitempty"`

	// Item number information
	ItemNumberInfo *ItemNumberType_166130S `xml:"itemNumberInfo,omitempty"`  // minOccurs="0"
}

type MnrGrp struct {
	Mnr *MiniRulesType `xml:"mnr,omitempty"`

	MnrDetails []*MnrDetails `xml:"mnrDetails,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type MnrDetails struct {
	MnrRef *ItemNumberType_176648S `xml:"mnrRef,omitempty"`

	DateInfo []*DateAndTimeInformationType_182345S `xml:"dateInfo,omitempty"`  // minOccurs="0" maxOccurs="16"

	CatGrp []*CatGrp `xml:"catGrp,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type CatGrp struct {

	// Category information
	CatInfo *CategDescrType `xml:"catInfo,omitempty"`

	// Monetary information
	MonInfo *MonetaryInformationType_174241S `xml:"monInfo,omitempty"`  // minOccurs="0"

	// Status information
	StatusInfo *StatusType_182386S `xml:"statusInfo,omitempty"`  // minOccurs="0"
}

type ActionDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ActionDetailsType"`

	// Number of items details
	NumberOfItemsDetails *ProcessingInformationType `xml:"numberOfItemsDetails,omitempty"`  // minOccurs="0"

	// Range of segments
	LastItemsDetails []*ReferenceType `xml:"lastItemsDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type AdditionalFareQualifierDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AdditionalFareQualifierDetailsTypeI"`

	// Rate class
	RateClass formats.AlphaNumericString_Length1To35 `xml:"rateClass,omitempty"`  // minOccurs="0"

	// Ticket designator.
	TicketDesignator formats.AlphaNumericString_Length1To18 `xml:"ticketDesignator,omitempty"`  // minOccurs="0"

	// Pricing group
	PricingGroup formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"`  // minOccurs="0"

	// Second rate class
	SecondRateClass []*formats.AlphaNumericString_Length1To35 `xml:"secondRateClass,omitempty"`  // minOccurs="0" maxOccurs="29"
}

type AdditionalProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AdditionalProductDetailsType"`

	// Type of aircraft
	EquipmentType formats.AlphaNumericString_Length1To3 `xml:"equipmentType,omitempty"`  // minOccurs="0"

	// Day number of the week
	OperatingDay formats.AlphaNumericString_Length1To7 `xml:"operatingDay,omitempty"`  // minOccurs="0"

	// Number of stops made in a journey if different from 0
	TechStopNumber *formats.NumericInteger_Length1To2 `xml:"techStopNumber,omitempty"`  // minOccurs="0"

	// Location places of the stops
	LocationId []*formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`  // minOccurs="0" maxOccurs="3"
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
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type AttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeInformationType"`

	// Type of parameter.
	FeeParameterType formats.AlphaNumericString_Length3To3 `xml:"feeParameterType,omitempty"`  // minOccurs="0"

	// Reference to company Id.
	FeeParameterDescription formats.AlphaNumericString_Length1To15 `xml:"feeParameterDescription,omitempty"`  // minOccurs="0"
}

type AttributeInformationType_97181C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeInformationType_97181C"`

	// Attribute type
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType,omitempty"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type AttributeTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeTypeU"`

	// provides the function of the attribute
	AttributeFunction formats.AlphaNumericString_Length1To3 `xml:"attributeFunction,omitempty"`  // minOccurs="0"

	// provides details for the Attribute
	AttributeDetails *AttributeInformationTypeU `xml:"attributeDetails,omitempty"`
}

type AttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeType"`

	// Criteria Set Type
	AttributeQualifier formats.AlphaNumericString_Length1To3 `xml:"attributeQualifier,omitempty"`  // minOccurs="0"

	// Criteria details
	AttributeDetails []*AttributeInformationType_97181C `xml:"attributeDetails,omitempty"`  // maxOccurs="99"
}

type AttributeType_78561S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A AttributeType_78561S"`

	// Fee/reduction parameters.
	FeeParameter []*AttributeInformationType `xml:"feeParameter,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type BaggageDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A BaggageDetailsType"`

	// Number of pieces or weight
	FreeAllowance *formats.NumericInteger_Length1To15 `xml:"freeAllowance,omitempty"`  // minOccurs="0"

	// Nature of the free allowance ( Number of pieces or weight)
	QuantityCode formats.AMA_EDICodesetType_Length1to3 `xml:"quantityCode,omitempty"`  // minOccurs="0"

	// Unit qualifier
	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type BagtagDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A BagtagDetailsType"`

	// Identifier
	Identifier formats.AlphaNumericString_Length1To35 `xml:"identifier,omitempty"`  // minOccurs="0"

	// Number
	Number *formats.NumericInteger_Length1To15 `xml:"number,omitempty"`  // minOccurs="0"
}

type CabinInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinInformationType"`

	// Identify the features associated to the cabin/class
	Service formats.AlphaNumericString_Length1To5 `xml:"service,omitempty"`

	// Cabin code designator
	Cabin []*formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type CabinProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinProductDetailsType"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`

	// Reservation Booking Modifier
	BookingModifier formats.AlphaNumericString_Length0To1 `xml:"bookingModifier,omitempty"`  // minOccurs="0"

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaNumericString_Length1To1 `xml:"cabin,omitempty"`  // minOccurs="0"

	// Availibility status : posting level
	AvlStatus formats.AlphaNumericString_Length0To3 `xml:"avlStatus,omitempty"`  // minOccurs="0"
}

type CabinProductDetailsType_195516C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinProductDetailsType_195516C"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`  // minOccurs="0"

	// Reservation Booking Modifier
	BookingModifier formats.AlphaNumericString_Length0To1 `xml:"bookingModifier,omitempty"`  // minOccurs="0"

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaNumericString_Length1To1 `xml:"cabin,omitempty"`  // minOccurs="0"

	// Availibility status : posting level
	AvlStatus formats.AlphaNumericString_Length0To3 `xml:"avlStatus,omitempty"`
}

type CabinProductDetailsType_205138C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinProductDetailsType_205138C"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`

	// Reservation Booking Modifier
	BookingModifier formats.AMA_EDICodesetType_Length1 `xml:"bookingModifier,omitempty"`  // minOccurs="0"

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`  // minOccurs="0"

	// Availibility status : posting level
	AvlStatus formats.AMA_EDICodesetType_Length1to3 `xml:"avlStatus,omitempty"`  // minOccurs="0"
}

type CabinProductDetailsType_229142C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CabinProductDetailsType_229142C"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`  // minOccurs="0"

	// Availibility status : posting level
	AvlStatus formats.AMA_EDICodesetType_Length1to3 `xml:"avlStatus,omitempty"`  // minOccurs="0"
}

type CategDescrType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CategDescrType"`

	// Category description information
	DescriptionInfo *CategoryDescriptionType `xml:"descriptionInfo,omitempty"`

	// Category processing indicator
	ProcessIndicator formats.AlphaNumericString_Length1To3 `xml:"processIndicator,omitempty"`  // minOccurs="0"
}

type CategoryDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CategoryDescriptionType"`

	// Category number from ATPCO naming conventions (C05 for Advance Purchase restrictions, C06 for Minimun stay ...)
	Number formats.NumericInteger_Length1To3 `xml:"number,omitempty"`

	// Category Code (ATPCO component code, e.g ADV for Advance purchase, STP for stopover restrictions, ELG for eligibility restrictions...)
	Code formats.AlphaString_Length1To3 `xml:"code,omitempty"`  // minOccurs="0"
}

type ClassInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ClassInformationType"`

	// Identify the features associated to the cabin/class
	Service formats.AlphaNumericString_Length1To5 `xml:"service,omitempty"`

	// Class designator
	Rbd []*formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`  // minOccurs="0" maxOccurs="26"
}

type CodedAttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CodedAttributeInformationType"`

	// Type of fee/reduction
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// Fee Id Number
	AttributeDescription formats.AlphaNumericString_Length1To50 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeInformationType_260992C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CodedAttributeInformationType_260992C"`

	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To10 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CodedAttributeType"`

	// Fee/reduction Id
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails,omitempty"`  // maxOccurs="9"
}

type CommercialAgreementsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CommercialAgreementsType"`

	// Codeshare Details
	CodeshareDetails *CompanyRoleIdentificationType `xml:"codeshareDetails,omitempty"`  // minOccurs="0"

	// Other codeshare details
	OtherCodeshareDetails []*CompanyRoleIdentificationType `xml:"otherCodeshareDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type CompanyIdentificationTextType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyIdentificationTextType"`

	// Company Id Text reference.
	TextRefNumber *formats.NumericInteger_Length0To4 `xml:"textRefNumber,omitempty"`  // minOccurs="0"

	// Company id free text.
	CompanyText formats.AlphaNumericString_Length0To70 `xml:"companyText,omitempty"`  // minOccurs="0"
}

type CompanyIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyIdentificationTypeI"`

	// Company
	MarketingCompany formats.AlphaNumericString_Length2To3 `xml:"marketingCompany,omitempty"`  // minOccurs="0"

	// Company
	OperatingCompany formats.AlphaNumericString_Length2To3 `xml:"operatingCompany,omitempty"`  // minOccurs="0"

	// Company
	OtherCompany formats.AlphaNumericString_Length2To3 `xml:"otherCompany,omitempty"`  // minOccurs="0"
}

type CompanyIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyIdentificationType"`

	// Marketing carrier
	MarketingCarrier formats.AlphaNumericString_Length2To3 `xml:"marketingCarrier,omitempty"`

	// Operating carrier
	OperatingCarrier formats.AlphaNumericString_Length2To3 `xml:"operatingCarrier,omitempty"`  // minOccurs="0"

	// airline alliance code
	Alliance formats.AlphaNumericString_Length1To2 `xml:"alliance,omitempty"`  // minOccurs="0"
}

type CompanyRoleIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyRoleIdentificationType"`

	// Type of code share agreement.
	CodeShareType formats.AlphaString_Length1To1 `xml:"codeShareType,omitempty"`  // minOccurs="0"

	// company identification
	AirlineDesignator formats.AlphaNumericString_Length2To3 `xml:"airlineDesignator,omitempty"`  // minOccurs="0"

	// flight number
	FlightNumber *formats.NumericInteger_Length1To4 `xml:"flightNumber,omitempty"`  // minOccurs="0"
}

type CompanyRoleIdentificationType_120771C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CompanyRoleIdentificationType_120771C"`

	// Type of code share agreement.
	TransportStageQualifier formats.AlphaNumericString_Length1To3 `xml:"transportStageQualifier,omitempty"`  // minOccurs="0"

	// company identification
	Company formats.AlphaNumericString_Length2To3 `xml:"company,omitempty"`  // minOccurs="0"
}

type ConversionRateDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ConversionRateDetailsTypeI"`

	// Conversion type
	ConversionType formats.AlphaNumericString_Length1To3 `xml:"conversionType,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// amount
	Amount formats.AlphaNumericString_Length0To12 `xml:"amount,omitempty"`  // minOccurs="0"
}

type ConversionRateDetailsTypeI_179848C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ConversionRateDetailsTypeI_179848C"`

	// Conversion type
	ConversionType formats.AlphaNumericString_Length1To3 `xml:"conversionType,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaString_Length1To3 `xml:"currency,omitempty"`

	// Conversion rate for pricing
	Rate formats.AlphaNumericString_Length0To18 `xml:"rate,omitempty"`  // minOccurs="0"

	// Converted value amount
	ConvertedAmountLink formats.AlphaNumericString_Length0To18 `xml:"convertedAmountLink,omitempty"`  // minOccurs="0"

	// Applicable ISO country code or Tax designator code.
	TaxQualifier formats.AlphaNumericString_Length0To3 `xml:"taxQualifier,omitempty"`  // minOccurs="0"
}

type ConversionRateTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ConversionRateTypeI"`

	// Detail of conversion rate of First Monetary Unit.
	ConversionRateDetail []*ConversionRateDetailsTypeI_179848C `xml:"conversionRateDetail,omitempty"`  // maxOccurs="9"
}

type ConversionRateTypeI_78562S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ConversionRateTypeI_78562S"`

	// Details of conversion
	ConversionRateDetail []*ConversionRateDetailsTypeI `xml:"conversionRateDetail,omitempty"`  // maxOccurs="9"
}

type CriteriaiDetaislType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A CriteriaiDetaislType"`

	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	Value formats.AlphaNumericString_Length1To18 `xml:"value,omitempty"`  // minOccurs="0"
}

type DataInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DataInformationType"`

	// Ancillary services options
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`  // minOccurs="0"
}

type DataTypeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DataTypeInformationType"`

	// service group/sub-group/sub-code information
	SubType formats.AlphaNumericString_Length1To3 `xml:"subType,omitempty"`

	// Status (automated, manually added, exempted). Default is automated
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateAndTimeDetailsType"`

	// Date time period qualifier
	DateQualifier formats.AlphaNumericString_Length1To3 `xml:"dateQualifier,omitempty"`  // minOccurs="0"

	// First Date
	Date *formats.Date_DDMMYY `xml:"date,omitempty"`  // minOccurs="0"

	// First Time
	FirstTime *formats.Time24_HHMM `xml:"firstTime,omitempty"`  // minOccurs="0"

	// Movement type.
	EquipementType formats.AlphaNumericString_Length1To3 `xml:"equipementType,omitempty"`  // minOccurs="0"

	// Place/location identification.
	LocationId formats.AlphaNumericString_Length3To5 `xml:"locationId,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsType_256192C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateAndTimeDetailsType_256192C"`

	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	Date formats.AlphaNumericString_Length1To35 `xml:"date,omitempty"`  // minOccurs="0"

	// Time
	Time *formats.Time24_HHMM `xml:"time,omitempty"`  // minOccurs="0"

	// Location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`  // minOccurs="0"
}

type DateAndTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateAndTimeInformationType"`

	// Details on date and time
	StopDetails []*DateAndTimeDetailsType `xml:"stopDetails,omitempty"`  // maxOccurs="2"

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DummyNET struct{}

type DateAndTimeInformationType_182345S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateAndTimeInformationType_182345S"`

	// DATE AND TIME DETAILS
	DateAndTimeDetails []*DateAndTimeDetailsType_256192C `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="400"

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DateTimePeriodDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DateTimePeriodDetailsTypeI"`

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Value
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`  // minOccurs="0"
}

type DiscountAndPenaltyInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DiscountAndPenaltyInformationType"`

	// Used to specify airline collected fee or agent collected fee.
	FeeIdentification formats.AlphaNumericString_Length1To3 `xml:"feeIdentification,omitempty"`  // minOccurs="0"

	// Used to specify penalty information
	FeeInformation *DiscountPenaltyMonetaryInformationType `xml:"feeInformation,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DiscountPenaltyInformationType"`

	// Discounted fare,...
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier,omitempty"`

	// Dicount code,...
	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"`  // minOccurs="0"

	// Amount
	Amount *formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`  // minOccurs="0"

	// Percentage
	Percentage *formats.NumericDecimal_Length1To8 `xml:"percentage,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyMonetaryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DiscountPenaltyMonetaryInformationType"`

	// Type of discount/penalty
	FeeType formats.AlphaNumericString_Length1To3 `xml:"feeType,omitempty"`  // minOccurs="0"

	// The amount Type can be a percentage or an amount
	FeeAmountType formats.AlphaNumericString_Length1To3 `xml:"feeAmountType,omitempty"`  // minOccurs="0"

	// specify the value
	FeeAmount *formats.NumericDecimal_Length1To18 `xml:"feeAmount,omitempty"`  // minOccurs="0"

	// Fee currency code.
	FeeCurrency formats.AlphaNumericString_Length1To3 `xml:"feeCurrency,omitempty"`  // minOccurs="0"
}

type DummySegmentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A DummySegmentTypeI"`
}

type ExcessBaggageType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ExcessBaggageType"`

	// Baggage details
	BaggageDetails *BaggageDetailsType `xml:"baggageDetails,omitempty"`  // minOccurs="0"

	// Free baggage allowance details
	BagTagDetails []*BagtagDetailsType `xml:"bagTagDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FareCalculationCodeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareCalculationCodeDetailsType"`

	// Qualifier of the amout or rate
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Amount
	Amount *formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`  // minOccurs="0"

	// Location code
	LocationCode formats.AlphaNumericString_Length1To3 `xml:"locationCode,omitempty"`  // minOccurs="0"

	// Other location code
	OtherLocationCode formats.AlphaNumericString_Length1To3 `xml:"otherLocationCode,omitempty"`  // minOccurs="0"

	// Rate
	Rate *formats.NumericDecimal_Length1To8 `xml:"rate,omitempty"`  // minOccurs="0"
}

type FareCategoryCodesTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareCategoryCodesTypeI"`

	// Fare type
	FareType formats.AlphaNumericString_Length1To20 `xml:"fareType,omitempty"`

	// Other fare type
	OtherFareType []*formats.AlphaNumericString_Length1To20 `xml:"otherFareType,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type FareDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareDetailsTypeI"`

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Rate
	Rate formats.NumericDecimal_Length1To8 `xml:"rate,omitempty"`  // minOccurs="0"

	// Country
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`  // minOccurs="0"

	// Fare category
	FareCategory formats.AlphaNumericString_Length1To3 `xml:"fareCategory,omitempty"`  // minOccurs="0"
}

type FareDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareDetailsType"`

	// Passenger Type qualifier
	PassengerTypeQualifier formats.AlphaNumericString_Length1To3 `xml:"passengerTypeQualifier,omitempty"`  // minOccurs="0"
}

type FareDetailsType_193037C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareDetailsType_193037C"`

	// Qualifier
	Qualifier formats.AMA_EDICodesetType_Length1to3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Rate
	Rate *formats.NumericInteger_Length1To8 `xml:"rate,omitempty"`  // minOccurs="0"

	// Country
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`  // minOccurs="0"

	// Fare Category
	FareCategory formats.AMA_EDICodesetType_Length1to3 `xml:"fareCategory,omitempty"`  // minOccurs="0"
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
	FareFamilyname formats.AlphaNumericString_Length1To10 `xml:"fareFamilyname,omitempty"`  // minOccurs="0"

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy *formats.NumericInteger_Length1To4 `xml:"hierarchy,omitempty"`  // minOccurs="0"

	// CABIN USED FOR FARE FAMILY
	Cabin formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`  // minOccurs="0"

	// Indicates Commercial Fare Family Short names
	CommercialFamilyDetails []*FareFamilyDetailsType `xml:"commercialFamilyDetails,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Short description of the fare family
	Description formats.AlphaNumericString_Length1To100 `xml:"description,omitempty"`  // minOccurs="0"

	// Carrier code
	Carrier formats.AlphaNumericString_Length2To3 `xml:"carrier,omitempty"`  // minOccurs="0"

	// Reference to the services details
	Services []*ServicesReferences `xml:"services,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type FareInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareInformationTypeI"`

	// Value qualifier
	ValueQualifier formats.AlphaNumericString_Length1To3 `xml:"valueQualifier,omitempty"`  // minOccurs="0"

	// Value
	Value *formats.NumericInteger_Length1To15 `xml:"value,omitempty"`  // minOccurs="0"
}

type FareInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareInformationType"`

	// Value Qualifier
	ValueQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"valueQualifier,omitempty"`  // minOccurs="0"

	// Value
	Value *formats.NumericInteger_Length1To15 `xml:"value,omitempty"`  // minOccurs="0"

	// Fare Details
	FareDetails *FareDetailsType_193037C `xml:"fareDetails,omitempty"`  // minOccurs="0"

	// Identity Number
	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber,omitempty"`  // minOccurs="0"

	// Fare Type Grouping
	FareTypeGrouping *FareTypeGroupingInformationType `xml:"fareTypeGrouping,omitempty"`  // minOccurs="0"

	// Rate Category
	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"`  // minOccurs="0"
}

type FareInformationType_80868S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareInformationType_80868S"`

	// Fare details
	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`  // minOccurs="0"
}

type FareProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareProductDetailsType"`

	// Fare basis code
	FareBasis formats.AlphaNumericString_Length1To18 `xml:"fareBasis,omitempty"`  // minOccurs="0"
}

type FareProductDetailsType_248552C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareProductDetailsType_248552C"`

	// Fare basis code
	FareBasis formats.AlphaNumericString_Length0To18 `xml:"fareBasis,omitempty"`  // minOccurs="0"

	// PTC priced
	PassengerType formats.AlphaNumericString_Length1To6 `xml:"passengerType,omitempty"`  // minOccurs="0"

	// Type of fare
	FareType []*formats.AlphaNumericString_Length0To3 `xml:"fareType,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FareQualifierDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareQualifierDetailsType"`

	// Route Code
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`  // minOccurs="0"

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
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FareTypeGroupingInformationType"`

	// Pricing Group
	PricingGroup formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"`  // minOccurs="0"
}

type FlightCharacteristicsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightCharacteristicsType"`

	// On-Time Performance
	OnTimePerformance *OnTimePerformanceType `xml:"onTimePerformance,omitempty"`  // minOccurs="0"

	// In flight services
	InFlightSrv formats.AlphaNumericString_Length1To3 `xml:"inFlightSrv,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FlightProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightProductInformationType"`

	// Indicates flight cabin details
	CabinProduct []*CabinProductDetailsType_195516C `xml:"cabinProduct,omitempty"`  // minOccurs="0" maxOccurs="6"

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType `xml:"contextDetails,omitempty"`  // minOccurs="0"
}

type FlightProductInformationType_141442S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightProductInformationType_141442S"`

	// Indicates flight cabin details
	CabinProduct []*CabinProductDetailsType_205138C `xml:"cabinProduct,omitempty"`  // minOccurs="0" maxOccurs="26"

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType_205137C `xml:"contextDetails,omitempty"`  // minOccurs="0"
}

type FlightProductInformationType_161491S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightProductInformationType_161491S"`

	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType_229142C `xml:"cabinProduct,omitempty"`  // minOccurs="0"

	// Fare product details
	FareProductDetail *FareProductDetailsType `xml:"fareProductDetail,omitempty"`  // minOccurs="0"
}

type FlightProductInformationType_176659S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightProductInformationType_176659S"`

	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType `xml:"cabinProduct,omitempty"`  // minOccurs="0"

	// Fare product details
	FareProductDetail *FareProductDetailsType_248552C `xml:"fareProductDetail,omitempty"`  // minOccurs="0"

	// Corporate number or name and number
	CorporateId []*formats.AlphaNumericString_Length1To20 `xml:"corporateId,omitempty"`  // minOccurs="0" maxOccurs="2"

	// To determine if Fare Breaks at this segment
	BreakPoint formats.AlphaString_Length1To1 `xml:"breakPoint,omitempty"`  // minOccurs="0"

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType `xml:"contextDetails,omitempty"`  // minOccurs="0"
}

type FlightServicesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FlightServicesType"`

	// Type of service used
	ServiceType formats.AlphaNumericString_Length1To3 `xml:"serviceType,omitempty"`

	CabinInfo []*CabinInformationType `xml:"cabinInfo,omitempty"`  // minOccurs="0" maxOccurs="99"

	ClassInfo []*ClassInformationType `xml:"classInfo,omitempty"`  // minOccurs="0" maxOccurs="99"
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
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`  // minOccurs="0"
}

type FreeTextQualificationType_120769C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FreeTextQualificationType_120769C"`

	// Type of message
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Coded Text or type of information in 4440 (e.g. type of OSI or free text, canned message value)
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`  // minOccurs="0"

	// ISO code for language of free text (default is English)
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`  // minOccurs="0"
}

type FrequentTravellerIdentificationCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FrequentTravellerIdentificationCodeType"`

	// Frequent Traveller Info
	FrequentTravellerDetails []*FrequentTravellerIdentificationType `xml:"frequentTravellerDetails,omitempty"`  // maxOccurs="99"
}

type FrequentTravellerIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A FrequentTravellerIdentificationType"`

	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"`  // minOccurs="0"

	// Number
	Number formats.AlphaNumericString_Length1To28 `xml:"number,omitempty"`  // minOccurs="0"

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`  // minOccurs="0"

	// For example : priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"`  // minOccurs="0"
}

type HeaderInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A HeaderInformationTypeI"`

	// Status
	Status []*formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Date and Time info
	DateTimePeriodDetails *DateTimePeriodDetailsTypeI `xml:"dateTimePeriodDetails,omitempty"`  // minOccurs="0"

	// Reference number
	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber,omitempty"`  // minOccurs="0"

	// Contains product identification such as UIC code...
	ProductIdentification []*formats.AlphaNumericString_Length1To35 `xml:"productIdentification,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type InteractiveFreeTextType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A InteractiveFreeTextType"`

	// Free text qualification
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Free text
	FreeText formats.AlphaNumericString_Length1To50 `xml:"freeText,omitempty"`  // minOccurs="0"
}

type InteractiveFreeTextType_78534S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A InteractiveFreeTextType_78534S"`

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Free text
	Description []*formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type InteractiveFreeTextType_78544S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A InteractiveFreeTextType_78544S"`

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType_120769C `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Free text
	Description []*formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type InteractiveFreeTextType_78559S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A InteractiveFreeTextType_78559S"`

	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType_120769C `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Free text
	Description []*formats.AlphaNumericString_Length1To500 `xml:"description,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type ItemNumberIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType"`

	// Ancillary Service number
	Number formats.AlphaNumericString_Length1To4 `xml:"number,omitempty"`  // minOccurs="0"

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Responsible agency
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`  // minOccurs="0"
}

type ItemNumberIdentificationType_191597C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType_191597C"`

	// Item number.
	Number formats.AlphaNumericString_Length1To6 `xml:"number,omitempty"`  // minOccurs="0"

	// Indicates the item type .
	NumberType formats.AlphaNumericString_Length0To3 `xml:"numberType,omitempty"`  // minOccurs="0"
}

type ItemNumberIdentificationType_192331C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType_192331C"`

	// Service coverage number
	Number formats.AlphaNumericString_Length1To6 `xml:"number,omitempty"`  // minOccurs="0"

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Responsible agency
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`  // minOccurs="0"
}

type ItemNumberIdentificationType_234878C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType_234878C"`

	// Number
	Number *formats.NumericInteger_Length1To6 `xml:"number,omitempty"`  // minOccurs="0"

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"
}

type ItemNumberIdentificationType_248537C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberIdentificationType_248537C"`

	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`  // minOccurs="0"
}

type ItemNumberType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType"`

	// Item number details
	ItemNumber *ItemNumberIdentificationType_192331C `xml:"itemNumber,omitempty"`
}

type ItemNumberType_161497S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType_161497S"`

	// Indicates the recommendation number.
	ItemNumberId *ItemNumberIdentificationType_191597C `xml:"itemNumberId,omitempty"`  // minOccurs="0"

	// Code share details.
	CodeShareDetails []*CompanyRoleIdentificationType_120771C `xml:"codeShareDetails,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Pricing ticketind details.
	PriceTicketing *PricingTicketingInformationType `xml:"priceTicketing,omitempty"`  // minOccurs="0"
}

type ItemNumberType_166130S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType_166130S"`

	// Item number details
	ItemNumberDetails []*ItemNumberIdentificationType_234878C `xml:"itemNumberDetails,omitempty"`  // maxOccurs="99"
}

type ItemNumberType_176648S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType_176648S"`

	ItemNumberDetails []*ItemNumberIdentificationType_248537C `xml:"itemNumberDetails,omitempty"`  // maxOccurs="99"
}

type ItemNumberType_80866S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemNumberType_80866S"`

	// Item number details
	ItemNumberDetails *ItemNumberIdentificationType `xml:"itemNumberDetails,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemReferencesAndVersionsType"`

	// Qualifies the type of the reference used.
	ReferenceType formats.AlphaNumericString_Length1To6 `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Unique fee reference.
	RefNumber *formats.NumericInteger_Length1To3 `xml:"refNumber,omitempty"`  // minOccurs="0"
}

type ItemReferencesAndVersionsType_78536S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemReferencesAndVersionsType_78536S"`

	// Qualifies the type of the reference used.
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Unique fee reference.
	RefNumber *formats.NumericInteger_Length1To3 `xml:"refNumber,omitempty"`  // minOccurs="0"
}

type ItemReferencesAndVersionsType_78564S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ItemReferencesAndVersionsType_78564S"`

	// Qualifies the type of the reference used.
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Unique fee reference.
	FeeRefNumber *formats.NumericInteger_Length1To3 `xml:"feeRefNumber,omitempty"`  // minOccurs="0"
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
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`  // minOccurs="0"

	// Terminal information
	Terminal formats.AlphaNumericString_Length1To5 `xml:"terminal,omitempty"`  // minOccurs="0"
}

type MiniRulesDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MiniRulesDetailsType"`

	// Coded text (period or day)
	Interpretation formats.AlphaString_Length0To9 `xml:"interpretation,omitempty"`  // minOccurs="0"

	// Data type coded or value of interpretation
	Value []*formats.AlphaNumericString_Length0To5 `xml:"value,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type MiniRulesIndicatorType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MiniRulesIndicatorType"`

	// See rule indicator and free form text indicator
	RuleIndicator []*formats.AlphaString_Length1To1 `xml:"ruleIndicator,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type MiniRulesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MiniRulesType"`

	// Categoty of restriction: PTC, Max Adv Pur, Days, ...
	Category formats.AlphaString_Length1To3 `xml:"category,omitempty"`
}

type MiniRulesType_78547S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MiniRulesType_78547S"`

	// Type of restriction: PTC, Max Adv Res, Max Ticketing After Res, ...
	RestrictionType formats.AlphaNumericString_Length0To6 `xml:"restrictionType,omitempty"`  // minOccurs="0"

	// Categoty of restriction: PTC, Max Adv Pur, Days, ...
	Category formats.AlphaString_Length0To3 `xml:"category,omitempty"`

	// Indicators
	Indicator *MiniRulesIndicatorType `xml:"indicator,omitempty"`  // minOccurs="0"

	// Mini rules
	MiniRules []*MiniRulesDetailsType `xml:"miniRules,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type MonetaryInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationDetailsTypeI"`

	// type Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To6 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount formats.AlphaNumericString_Length1To35 `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationDetailsType"`

	// To specify amount and percentage.
	AmountType formats.AlphaNumericString_Length0To3 `xml:"amountType,omitempty"`  // minOccurs="0"

	// Amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// ISO currency code
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsType_245528C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationDetailsType_245528C"`

	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount *formats.NumericDecimal_Length1To35 `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`  // minOccurs="0"
}

type MonetaryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationTypeI"`

	// Monetary details
	MonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"monetaryDetails,omitempty"`  // maxOccurs="99"
}

type MonetaryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationType"`

	// Monetary information.
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"`  // minOccurs="0" maxOccurs="30"
}

type MonetaryInformationType_137835S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationType_137835S"`

	// Monetary information.
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type MonetaryInformationType_174241S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationType_174241S"`

	MonetaryDetails *MonetaryInformationDetailsType_245528C `xml:"monetaryDetails,omitempty"`

	OtherMonetaryDetails []*MonetaryInformationDetailsType_245528C `xml:"otherMonetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type MonetaryInformationType_185955S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A MonetaryInformationType_185955S"`

	// Monetary information
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"`  // maxOccurs="2"
}

type OnTimePerformanceType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A OnTimePerformanceType"`

	// Date time period
	DateTimePeriod formats.AlphaNumericString_Length1To35 `xml:"dateTimePeriod,omitempty"`  // minOccurs="0"

	// Percentage
	Percentage *formats.NumericDecimal_Length1To8 `xml:"percentage,omitempty"`  // minOccurs="0"

	// Accuracy
	Accuracy formats.AlphaNumericString_Length1To3 `xml:"accuracy,omitempty"`  // minOccurs="0"
}

type OriginAndDestinationRequestType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A OriginAndDestinationRequestType"`

	// Requested segment number
	SegRef formats.NumericInteger_Length1To2 `xml:"segRef,omitempty"`

	// Forces arrival or departure, from/to the same airport/city
	LocationForcing []*ItineraryDetailsType `xml:"locationForcing,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type OriginAndDestinationRequestType_134833S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A OriginAndDestinationRequestType_134833S"`

	// Requested segment number
	SegRef formats.NumericInteger_Length1To2 `xml:"segRef,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A OriginatorIdentificationDetailsTypeI"`

	// Office Name.
	OfficeName *formats.NumericInteger_Length1To9 `xml:"officeName,omitempty"`  // minOccurs="0"

	// Agent Sign In .
	AgentSignin formats.AlphaNumericString_Length1To9 `xml:"agentSignin,omitempty"`  // minOccurs="0"

	// Confidential Office Name.
	ConfidentialOffice formats.AlphaNumericString_Length1To9 `xml:"confidentialOffice,omitempty"`  // minOccurs="0"

	// Other Office Name
	OtherOffice formats.AlphaNumericString_Length1To9 `xml:"otherOffice,omitempty"`  // minOccurs="0"
}

type PricingTicketingInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A PricingTicketingInformationType"`

	// Price type qualifier
	PriceType []*formats.AlphaNumericString_Length0To3 `xml:"priceType,omitempty"`  // maxOccurs="20"
}

type PricingTicketingSubsequentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A PricingTicketingSubsequentType"`

	// Passenger fare product number
	PaxFareNum []*formats.AlphaNumericString_Length1To3 `xml:"paxFareNum,omitempty"`  // maxOccurs="10"
}

type PricingTicketingSubsequentType_144401S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A PricingTicketingSubsequentType_144401S"`

	// Passenger fare product number
	PaxFareNum formats.AlphaNumericString_Length1To3 `xml:"paxFareNum,omitempty"`

	// Total fare amount
	TotalFareAmount formats.NumericDecimal_Length1To18 `xml:"totalFareAmount,omitempty"`

	// Total tax amount
	TotalTaxAmount formats.NumericDecimal_Length1To18 `xml:"totalTaxAmount,omitempty"`  // minOccurs="0"

	// Code share details.
	CodeShareDetails *CompanyRoleIdentificationType_120771C `xml:"codeShareDetails,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Monetary information.
	MonetaryDetails []*MonetaryInformationDetailsType `xml:"monetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Pricing ticketing details.
	PricingTicketing *PricingTicketingInformationType `xml:"pricingTicketing,omitempty"`  // minOccurs="0"
}

type ProcessingInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProcessingInformationType"`

	// Action qualifier
	ActionQualifier formats.AlphaNumericString_Length1To3 `xml:"actionQualifier,omitempty"`  // minOccurs="0"

	// Reference qualifier
	ReferenceQualifier formats.AlphaNumericString_Length1To3 `xml:"referenceQualifier,omitempty"`  // minOccurs="0"

	// Reference number
	RefNum formats.AlphaNumericString_Length1To6 `xml:"refNum,omitempty"`  // minOccurs="0"
}

type ProductDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductDateTimeType"`

	// Departure date
	DateOfDeparture formats.Date_DDMMYY `xml:"dateOfDeparture,omitempty"`

	// Departure time
	TimeOfDeparture *formats.Time24_HHMM `xml:"timeOfDeparture,omitempty"`  // minOccurs="0"

	// Arrival date
	DateOfArrival *formats.Date_DDMMYY `xml:"dateOfArrival,omitempty"`  // minOccurs="0"

	// Arrival time
	TimeOfArrival *formats.Time24_HHMM `xml:"timeOfArrival,omitempty"`  // minOccurs="0"

	// Arrival date compared to departure date, only if different from 0
	DateVariation *formats.NumericInteger_Length1To1 `xml:"dateVariation,omitempty"`  // minOccurs="0"
}

type ProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductDetailsType"`

	Designator formats.AlphaNumericString_Length1To17 `xml:"designator,omitempty"`

	AvailabilityStatus formats.AlphaNumericString_Length1To3 `xml:"availabilityStatus,omitempty"`  // minOccurs="0"

	SpecialService formats.AlphaNumericString_Length1To3 `xml:"specialService,omitempty"`  // minOccurs="0"

	Option []*formats.AlphaNumericString_Length1To7 `xml:"option,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type ProductFacilitiesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductFacilitiesType"`

	// Yes-No indicator whether Last Seat Available
	LastSeatAvailable formats.AlphaString_Length1To1 `xml:"lastSeatAvailable,omitempty"`  // minOccurs="0"

	// Level of access
	LevelOfAccess formats.AlphaNumericString_Length1To3 `xml:"levelOfAccess,omitempty"`  // minOccurs="0"

	// Yes-No indicator whether electronic ticketing
	ElectronicTicketing formats.AlphaString_Length1To1 `xml:"electronicTicketing,omitempty"`  // minOccurs="0"

	// Product identification suffix
	OperationalSuffix formats.AlphaString_Length1To1 `xml:"operationalSuffix,omitempty"`  // minOccurs="0"

	// Define whether a flight has been polled or not
	ProductDetailQualifier formats.AlphaNumericString_Length1To3 `xml:"productDetailQualifier,omitempty"`  // minOccurs="0"

	// Add some flight restrictions (See code set list)
	FlightCharacteristic []*formats.AlphaNumericString_Length1To3 `xml:"flightCharacteristic,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type ProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductInformationType"`

	// value of the Qualifier: INT for International DOM for Domestic EUR for European  otherwise CM#10569 INVALID INTERNATIONAL INDICATOR is returned.
	ProductDetailsQualifier formats.AlphaNumericString_Length1To3 `xml:"productDetailsQualifier,omitempty"`  // minOccurs="0"

	BookingClassDetails []*ProductDetailsType `xml:"bookingClassDetails,omitempty"`  // minOccurs="0" maxOccurs="26"
}

type ProductTypeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductTypeDetailsType"`

	// Availability connection type.
	AvailabilityCnxType []*formats.AlphaNumericString_Length1To3 `xml:"availabilityCnxType,omitempty"`  // maxOccurs="9"
}

type ProductTypeDetailsType_205137C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProductTypeDetailsType_205137C"`

	// indicates whether the flight is domestic or international
	Avl []*formats.AlphaNumericString_Length1To6 `xml:"avl,omitempty"`  // maxOccurs="9"
}

type ProposedSegmentDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProposedSegmentDetailsType"`

	// Flight proposal reference
	Ref formats.AlphaNumericString_Length1To6 `xml:"ref,omitempty"`  // minOccurs="0"

	// Elapse Flying Time
	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type ProposedSegmentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ProposedSegmentType"`

	// Parameters for proposed flight group
	FlightProposal []*ProposedSegmentDetailsType `xml:"flightProposal,omitempty"`  // maxOccurs="9"

	// Flight characteristics.
	FlightCharacteristic formats.AlphaNumericString_Length0To3 `xml:"flightCharacteristic,omitempty"`  // minOccurs="0"

	// Majority cabin
	MajCabin formats.AlphaString_Length1To1 `xml:"majCabin,omitempty"`  // minOccurs="0"
}

type ReferenceInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType"`

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_191583C `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="200"

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type ReferenceInfoType_133176S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_133176S"`

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferenceInfoType_134839S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_134839S"`

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_195561C `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferenceInfoType_134840S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_134840S"`

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_195561C `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="200"
}

type ReferenceInfoType_165972S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_165972S"`

	// Reference details
	ReferenceDetails []*ReferencingDetailsType_234704C `xml:"referenceDetails,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type ReferenceInfoType_176658S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceInfoType_176658S"`

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type ReferenceType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferenceType"`

	// Reference  of leg
	RefOfLeg formats.AlphaNumericString_Length1To6 `xml:"refOfLeg,omitempty"`  // minOccurs="0"

	// Reference of segment starting range
	FirstItemIdentifier *formats.NumericInteger_Length1To3 `xml:"firstItemIdentifier,omitempty"`  // minOccurs="0"

	// Reference of segment ending range
	LastItemIdentifier *formats.NumericInteger_Length1To3 `xml:"lastItemIdentifier,omitempty"`  // minOccurs="0"
}

type ReferencingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferencingDetailsType"`

	// Reference qualifier
	RefQualifier formats.AlphaNumericString_Length0To3 `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Requested segment reference
	RefNumber formats.NumericInteger_Length0To3 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsType_191583C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferencingDetailsType_191583C"`

	// Service reference qualifier
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Service reference
	RefNumber formats.NumericInteger_Length0To6 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsType_195561C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferencingDetailsType_195561C"`

	// Segment reference qualifier
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Flight or flight group reference
	RefNumber formats.NumericInteger_Length0To3 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsType_234704C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ReferencingDetailsType_234704C"`

	// Type
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`  // minOccurs="0"

	// Value
	Value formats.AlphaNumericString_Length1To60 `xml:"value,omitempty"`  // minOccurs="0"
}

type SegmentRepetitionControlDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SegmentRepetitionControlDetailsTypeI"`

	// traveller number
	Quantity *formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"`  // minOccurs="0"

	// range of traveller
	NumberOfUnits *formats.NumericInteger_Length1To15 `xml:"numberOfUnits,omitempty"`  // minOccurs="0"
}

type SegmentRepetitionControlTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SegmentRepetitionControlTypeI"`

	// Segment control details
	SegmentControlDetails []*SegmentRepetitionControlDetailsTypeI `xml:"segmentControlDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type SelectionDetailsInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SelectionDetailsInformationType"`

	// Carrier fee type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Carrier fee status
	OptionInformation formats.AlphaNumericString_Length1To3 `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type SelectionDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SelectionDetailsType"`

	// Carrier fees options
	CarrierFeeDetails *SelectionDetailsInformationType `xml:"carrierFeeDetails,omitempty"`
}

type SequenceDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SequenceDetailsTypeU"`

	// Sequence details
	SequenceDetails *SequenceInformationTypeU `xml:"sequenceDetails,omitempty"`  // minOccurs="0"
}

type SequenceInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SequenceInformationTypeU"`

	// Number
	Number formats.AlphaNumericString_Length1To10 `xml:"number,omitempty"`

	// Identification code
	IdentificationCode formats.AlphaNumericString_Length1To17 `xml:"identificationCode,omitempty"`  // minOccurs="0"
}

type ServicesReferences struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ServicesReferences"`

	// Reference of the service
	Reference formats.AlphaNumericString_Length1To4 `xml:"reference,omitempty"`  // minOccurs="0"

	// Status of the service
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`  // minOccurs="0"
}

type SpecialRequirementsDataDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecialRequirementsDataDetailsType"`

	// SSR seat characteristic
	SeatCharacteristics []*formats.AlphaNumericString_Length1To2 `xml:"seatCharacteristics,omitempty"`  // minOccurs="0" maxOccurs="5"

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type SpecialRequirementsDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecialRequirementsDetailsType"`

	// To specify the Service Requirement of the customer
	ServiceRequirementsInfo *SpecialRequirementsTypeDetailsType `xml:"serviceRequirementsInfo,omitempty"`

	// Seat details
	SeatDetails []*SpecialRequirementsDataDetailsType `xml:"seatDetails,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type SpecialRequirementsTypeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecialRequirementsTypeDetailsType"`

	// To specify the Service Classification of the Service Requirement.
	ServiceClassification formats.AlphaNumericString_Length1To4 `xml:"serviceClassification,omitempty"`

	// Status
	ServiceStatus formats.AlphaNumericString_Length1To3 `xml:"serviceStatus,omitempty"`  // minOccurs="0"

	// To specify the number of items involved
	ServiceNumberOfInstances *formats.NumericInteger_Length1To15 `xml:"serviceNumberOfInstances,omitempty"`  // minOccurs="0"

	// To specify to which marketing carrier the service applies
	ServiceMarketingCarrier formats.AlphaNumericString_Length1To3 `xml:"serviceMarketingCarrier,omitempty"`  // minOccurs="0"

	// Specify the Service group
	ServiceGroup formats.AlphaNumericString_Length1To3 `xml:"serviceGroup,omitempty"`  // minOccurs="0"

	// Specify the Service Sub-Group
	ServiceSubGroup formats.AlphaNumericString_Length1To3 `xml:"serviceSubGroup,omitempty"`  // minOccurs="0"

	// Free Text attached to the Service.
	ServiceFreeText []*formats.AlphaNumericString_Length1To70 `xml:"serviceFreeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type SpecificDataInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecificDataInformationType"`

	// Carrier fee description
	DataTypeInformation *DataTypeInformationType `xml:"dataTypeInformation,omitempty"`

	// Data information
	DataInformation []*DataInformationType `xml:"dataInformation,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type SpecificTravellerDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecificTravellerDetailsType"`

	// Reference number
	ReferenceNumber formats.AlphaNumericString_Length1To10 `xml:"referenceNumber,omitempty"`  // minOccurs="0"
}

type SpecificTravellerType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A SpecificTravellerType"`

	// Traveller details
	TravellerDetails []*SpecificTravellerDetailsType `xml:"travellerDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type StatusDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A StatusDetailsType"`

	// Advisory type information, Fare Server
	AdvisoryTypeInfo formats.AlphaNumericString_Length1To3 `xml:"advisoryTypeInfo,omitempty"`  // minOccurs="0"

	// CPU time, user type
	Notification formats.AlphaNumericString_Length1To3 `xml:"notification,omitempty"`  // minOccurs="0"

	// CPU time,user type
	Notification2 formats.AlphaNumericString_Length1To3 `xml:"notification2,omitempty"`  // minOccurs="0"

	// Capture and trace information
	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`  // minOccurs="0"
}

type StatusDetailsType_256255C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A StatusDetailsType_256255C"`

	// list of status/qualifiers Either His for Historical or     Crt for Current
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`  // minOccurs="0"

	Action formats.AlphaNumericString_Length1To3 `xml:"action,omitempty"`  // minOccurs="0"
}

type StatusType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A StatusType"`

	// Status details
	Status []*StatusDetailsType `xml:"status,omitempty"`  // maxOccurs="10"
}

type StatusType_182386S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A StatusType_182386S"`

	// STATUS DETAILS
	StatusInformation []*StatusDetailsType_256255C `xml:"statusInformation,omitempty"`  // maxOccurs="99"
}

type TaxDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TaxDetailsType"`

	// Amount
	Rate formats.AlphaNumericString_Length1To12 `xml:"rate,omitempty"`  // minOccurs="0"

	// Country code
	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"`  // minOccurs="0"

	// Currency code
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	// Indicator
	Indicator []*formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`  // minOccurs="0" maxOccurs="98"
}

// TaxType ...
type TaxType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TaxType"`

	// Tax category
	TaxCategory formats.AlphaNumericString_Length1To3 `xml:"taxCategory,omitempty"`  // minOccurs="0"

	// Tax details
	TaxDetails []*TaxDetailsType `xml:"taxDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

// TransportIdentifierType ...
type TransportIdentifierType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TransportIdentifierType"`

	// Company identification
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`  // minOccurs="0"
}

// TravelProductType ..
type TravelProductType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TravelProductType"`

	// Date and time of departure and arrival
	ProductDateTime *ProductDateTimeType `xml:"productDateTime,omitempty"`

	// Location of departure and arrival
	Location []*LocationIdentificationDetailsType `xml:"location,omitempty"`  // maxOccurs="2"

	CompanyID *CompanyIdentificationType `xml:"companyId,omitempty"`  // minOccurs="0"

	// Flight number or trainNumber
	FlightOrtrainNumber formats.AlphaNumericString_Length1To8 `xml:"flightOrtrainNumber,omitempty"`  // minOccurs="0"

	// Product details
	ProductDetail *AdditionalProductDetailsType `xml:"productDetail,omitempty"`  // minOccurs="0"

	// Additional product details
	AddProductDetail *ProductFacilitiesType `xml:"addProductDetail,omitempty"`  // minOccurs="0"

	// Attribute details
	AttributeDetails *CodedAttributeInformationType_260992C `xml:"attributeDetails,omitempty"`  // minOccurs="0" maxOccurs="20"
}

// TravellerDetailsType ...
type TravellerDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TravellerDetailsType"`

	// Direct reference of passenger assigned by requesting system.
	Ref *formats.NumericInteger_Length1To3 `xml:"ref,omitempty"`  // minOccurs="0"

	// Traveller is an infant
	InfantIndicator *formats.NumericInteger_Length1To1 `xml:"infantIndicator,omitempty"`  // minOccurs="0"
}

// TravellerReferenceInformationType ...
type TravellerReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A TravellerReferenceInformationType"`

	// Requested passenger type
	Ptc []*formats.AlphaNumericString_Length1To6 `xml:"ptc,omitempty"`  // minOccurs="0" maxOccurs="3"

	// Traveller details
	Traveller []*TravellerDetailsType `xml:"traveller,omitempty"`  // minOccurs="0" maxOccurs="9"
}

// UserIdentificationType ...
type UserIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A UserIdentificationType"`

	// Originator Identification Details
	OfficeIdentification *OriginatorIdentificationDetailsTypeI `xml:"officeIdentification,omitempty"`  // minOccurs="0"

	// Used to specify which kind of info is given in DE 9900.
	OfficeType formats.AlphaNumericString_Length1To1 `xml:"officeType,omitempty"`  // minOccurs="0"

	// The code given to an agent by the originating reservation system.
	OfficeCode formats.AlphaNumericString_Length1To30 `xml:"officeCode,omitempty"`  // minOccurs="0"
}

type ValueSearchCriteriaType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_14_3_1A ValueSearchCriteriaType"`

	Ref formats.AlphaNumericString_Length1To35 `xml:"ref,omitempty"`  // minOccurs="0"

	Value formats.AlphaNumericString_Length1To18 `xml:"value,omitempty"`  // minOccurs="0"

	CriteriaDetails []*CriteriaiDetaislType `xml:"criteriaDetails,omitempty"`  // minOccurs="0" maxOccurs="10"
}
