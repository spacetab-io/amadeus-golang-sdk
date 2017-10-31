package fare_masterpricertravelboardsearch

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type FareMasterPricerTravelBoardSearch struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A Fare_MasterPricerTravelBoardSearch"`

	// Number of seats , recommendations.
	NumberOfUnit *NumberOfUnitsType `xml:"numberOfUnit,omitempty"`

	// Global options
	GlobalOptions *AttributeType `xml:"globalOptions,omitempty"`

	// Traveler Details
	PaxReference *TravellerReferenceInformationType `xml:"paxReference,omitempty"`

	// Customer references
	CustomerRef *ConsumerReferenceInformationType `xml:"customerRef,omitempty"`

	// Fee with different payment forms  by passenger.
	FormOfPaymentByPassenger *FOPRepresentationType `xml:"formOfPaymentByPassenger,omitempty"`

	// Solution Family
	SolutionFamily *FareInformationType `xml:"solutionFamily,omitempty"`

	// Passenger info group  (9 ADT + 9 IN)
	PassengerInfoGrp *GroupPassengerDetailsType `xml:"passengerInfoGrp,omitempty"`

	FareFamilies *FareFamilies `xml:"fareFamilies,omitempty"`

	FareOptions *FareOptions `xml:"fareOptions,omitempty"`

	// Indicates Price to beat
	PriceToBeat *MonetaryInformationType `xml:"priceToBeat,omitempty"`

	// Tax Details
	TaxInfo *TaxType `xml:"taxInfo,omitempty"`

	// Details of a Flight : Direct, Non stop...
	TravelFlightInfo *TravelFlightInformationType_185853S `xml:"travelFlightInfo,omitempty"`

	ValueSearch *ValueSearchCriteriaType `xml:"valueSearch,omitempty"`

	Itinerary *Itinerary `xml:"itinerary,omitempty"`

	TicketChangeInfo *TicketChangeInfo `xml:"ticketChangeInfo,omitempty"`

	CombinationFareFamilies *CombinationFareFamilies `xml:"combinationFareFamilies,omitempty"`

	FeeOption *FeeOption `xml:"feeOption,omitempty"`

	OfficeIdDetails *OfficeIdDetails `xml:"officeIdDetails,omitempty"`
}

type FareFamilies struct {

	// Segment used to target Commercial Fare Family.
	FamilyInformation *FareFamilyType `xml:"familyInformation,omitempty"`

	// Description of Fare Family Criteria.
	FamilyCriteria *FareFamilyCriteriaType `xml:"familyCriteria,omitempty"`

	FareFamilySegment *FareFamilySegment `xml:"fareFamilySegment,omitempty"`

	OtherPossibleCriteria *OtherPossibleCriteria `xml:"otherPossibleCriteria,omitempty"`
}

type FareFamilySegment struct {

	// Requested segment reference
	ReferenceInfo *ReferenceInfoType `xml:"referenceInfo,omitempty"`

	// Description of fare family criteria.
	FamilyCriteria *FareFamilyCriteriaType `xml:"familyCriteria,omitempty"`
}

type OtherPossibleCriteria struct {

	// Logical link with other criteria.
	LogicalLink *BooleanExpressionRuleType `xml:"logicalLink,omitempty"`

	// Description of fare family criteria.
	FamilyCriteria *FareFamilyCriteriaType `xml:"familyCriteria,omitempty"`

	FareFamilySegment *FareFamilySegment `xml:"fareFamilySegment,omitempty"`
}

type FareOptions struct {

	// Pricing and ticketing details.
	PricingTickInfo *PricingTicketingDetailsType `xml:"pricingTickInfo,omitempty"`

	// Corporate name/number used to target fares
	Corporate *CorporateIdentificationType `xml:"corporate,omitempty"`

	// Ticketing price scheme.
	TicketingPriceScheme *TicketingPriceSchemeType `xml:"ticketingPriceScheme,omitempty"`

	// PSR number
	FeeIdDescription *CodedAttributeType `xml:"feeIdDescription,omitempty"`

	// Used to force the currency of pricing
	ConversionRate *ConversionRateType `xml:"conversionRate,omitempty"`

	// Form of payment information.
	FormOfPayment *FormOfPaymentTypeI `xml:"formOfPayment,omitempty"`

	// Frequent traveller information
	FrequentTravellerInfo *FrequentTravellerIdentificationCodeType_177150S `xml:"frequentTravellerInfo,omitempty"`

	// Monetary and cabin information.
	MonetaryCabinInfo *MonetaryAndCabinInformationType `xml:"monetaryCabinInfo,omitempty"`
}

type Itinerary struct {

	// Indicates reference of the requested segment
	RequestedSegmentRef *OriginAndDestinationRequestType `xml:"requestedSegmentRef,omitempty"`

	// Specification of the requested departure point
	DepartureLocalization *DepartureLocationType `xml:"departureLocalization,omitempty"`

	// Specification of the requested arrival point
	ArrivalLocalization *ArrivalLocalizationType `xml:"arrivalLocalization,omitempty"`

	// Details on requested date and time plus range of date trip duration
	TimeDetails *DateAndTimeInformationType_181295S `xml:"timeDetails,omitempty"`

	// Specify Flight options.
	FlightInfo *TravelFlightInformationType_165053S `xml:"flightInfo,omitempty"`

	ValueSearch *ValueSearchCriteriaType `xml:"valueSearch,omitempty"`

	GroupOfFlights *GroupOfFlights `xml:"groupOfFlights,omitempty"`

	FlightInfoPNR *FlightInfoPNR `xml:"flightInfoPNR,omitempty"`

	// Action identification for the requested segment
	RequestedSegmentAction *ActionIdentificationType `xml:"requestedSegmentAction,omitempty"`

	// Coded attributes
	Attributes *CodedAttributeType_181239S `xml:"attributes,omitempty"`
}

type GroupOfFlights struct {

	// To indicate parameters for proposed flight group.
	PropFlightGrDetail *ProposedSegmentType `xml:"propFlightGrDetail,omitempty"`

	// Indicates Price to beat
	PriceToBeat *MonetaryInformationType `xml:"priceToBeat,omitempty"`

	FlightDetails *FlightDetails `xml:"flightDetails,omitempty"`
}

type FlightDetails struct {

	// Specification of details on the flight and posting availability
	FlightInformation *TravelProductType `xml:"flightInformation,omitempty"`

	// returns booking class and availability context
	AvlInfo *FlightProductInformationType `xml:"avlInfo,omitempty"`

	// Details on Flight date, time and location of technical stop or change of gauge
	TechnicalStop *DateAndTimeInformationType `xml:"technicalStop,omitempty"`

	// Code Share Agreement description for current flight.
	CommercialAgreement *CommercialAgreementsType_78540S `xml:"commercialAgreement,omitempty"`

	// Additional Info about flight, such as Reference number, and several options
	AddInfo *HeaderInformationTypeI `xml:"addInfo,omitempty"`

	// Terminal, Equipment and EFT Details. If a Total EFT of the Travel Solution is requested then we will have 2 occurrences of the terminalEquipmentDetails attached to the first leg flightInformation: the 1st occurence will give the EFT of the leg, the 2nd occurrence will give the total EFT of the Travel Solution.
	TerminalEquipmentDetails *AdditionalProductDetailsTypeI `xml:"terminalEquipmentDetails,omitempty"`

	// PNR flight reservation info
	ReservationInfo *PassengerItineraryInformationType `xml:"reservationInfo,omitempty"`

	// Indicates Price to beat
	PriceToBeat *MonetaryInformationType `xml:"priceToBeat,omitempty"`
}

type FlightInfoPNR struct {

	// Travel Response Details: - Board/Off aiports - Flight number - Part of the journey - Day difference between Board and off  - Departure/Arrival dates - Departure/Arrival times
	TravelResponseDetails *TravelProductInformationTypeI `xml:"travelResponseDetails,omitempty"`

	// Time Table Effective/Discontinue dates and frequency of operating Travel Solution
	TimeTableDate *StructuredPeriodInformationType `xml:"timeTableDate,omitempty"`

	// Terminal, Equipment and EFT Details. If a Total EFT of the Travel Solution is requested then we will have 2 occurrences of the terminalEquipmentDetails attached to the first leg travelResponseDetails: the 1st occurrence will give the EFT of the leg, the 2nd occurrence will give the total EFT of the Travel Solution.
	TerminalEquipmentDetails *AdditionalProductDetailsTypeI `xml:"terminalEquipmentDetails,omitempty"`

	// Codeshare data
	CodeshareData *CommercialAgreementsType `xml:"codeshareData,omitempty"`

	// Disclosure message from an operating carrier.
	Disclosure *FreeTextInformationType `xml:"disclosure,omitempty"`

	// Stops Details
	StopDetails *RoutingInformationTypeI `xml:"stopDetails,omitempty"`

	// Traffic restriction information
	TrafficRestrictionData *TrafficRestrictionTypeI `xml:"trafficRestrictionData,omitempty"`

	// PNR flight reservation info
	ReservationInfo *PassengerItineraryInformationType `xml:"reservationInfo,omitempty"`

	IncidentalStopInfo *IncidentalStopInfo `xml:"incidentalStopInfo,omitempty"`
}

type IncidentalStopInfo struct {

	// Incidental stop date/time information
	DateTimeInfo *DateAndTimeInformationTypeI `xml:"dateTimeInfo,omitempty"`
}

type TicketChangeInfo struct {

	// Ticket nb details.
	TicketNumberDetails *TicketNumberTypeI `xml:"ticketNumberDetails,omitempty"`

	TicketRequestedSegments *TicketRequestedSegments `xml:"ticketRequestedSegments,omitempty"`
}

type TicketRequestedSegments struct {

	// Action identification.
	ActionIdentification *ActionIdentificationType `xml:"actionIdentification,omitempty"`

	// Connected cities in changed ticket requested segment.
	ConnectPointDetails *ConnectionTypeI `xml:"connectPointDetails,omitempty"`
}

type CombinationFareFamilies struct {

	// Specification of the item number
	ItemFFCNumber *ItemNumberType `xml:"itemFFCNumber,omitempty"`

	// Number of units.
	NbOfUnits *NumberOfUnitsType_80154S `xml:"nbOfUnits,omitempty"`

	// Requested segment reference
	ReferenceInfo *ReferenceInfoType `xml:"referenceInfo,omitempty"`
}

type FeeOption struct {

	// Nature of the fee (OB,OC,..)
	FeeTypeInfo *SelectionDetailsType `xml:"feeTypeInfo,omitempty"`

	// Associated rate tax.
	RateTax *MonetaryInformationType_80162S `xml:"rateTax,omitempty"`

	FeeDetails *FeeDetails `xml:"feeDetails,omitempty"`
}

type FeeDetails struct {

	// Fee information
	FeeInfo *SpecificDataInformationType `xml:"feeInfo,omitempty"`

	// Associated amounts : amounts to take into account to calculate fee.
	AssociatedAmounts *MonetaryInformationTypeI `xml:"associatedAmounts,omitempty"`

	FeeDescriptionGrp *FeeDescriptionGrp `xml:"feeDescriptionGrp,omitempty"`
}

type FeeDescriptionGrp struct {

	// Specification of the item number
	ItemNumberInfo *ItemNumberType_80866S `xml:"itemNumberInfo,omitempty"`

	// Attributes  (SSR code EMD, RFIC, SSIM)
	ServiceAttributesInfo *AttributeType_61377S `xml:"serviceAttributesInfo,omitempty"`

	// Other service information (service description, ...)
	ServiceDescriptionInfo *SpecialRequirementsDetailsType `xml:"serviceDescriptionInfo,omitempty"`
}

type OfficeIdDetails struct {

	// Office Id Information
	OfficeIdInformation *UserIdentificationType `xml:"officeIdInformation,omitempty"`

	// Number of units.
	NbOfUnits *NumberOfUnitsType_80154S `xml:"nbOfUnits,omitempty"`

	// UID option
	UidOption *CodedAttributeType_78500S `xml:"uidOption,omitempty"`

	// Pricing and ticketing details.
	PricingTickInfo *PricingTicketingDetailsType `xml:"pricingTickInfo,omitempty"`

	// Corporate fare information
	CorporateFareInfo *CorporateFareInformationType `xml:"corporateFareInfo,omitempty"`

	// Details of a Flight : Direct, Non stop...
	TravelFlightInfo *TravelFlightInformationType `xml:"travelFlightInfo,omitempty"`

	AirlineDistributionDetails *AirlineDistributionDetails `xml:"airlineDistributionDetails,omitempty"`
}

type AirlineDistributionDetails struct {

	// Indicates reference of the requested segment
	RequestedSegmentRef *OriginAndDestinationRequestType `xml:"requestedSegmentRef,omitempty"`

	// Specify Flight options.
	FlightInfo *TravelFlightInformationType `xml:"flightInfo,omitempty"`
}

type ActionIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ActionIdentificationType"`

	// Action request code
	ActionRequestCode formats.AlphaNumericString_Length1To3 `xml:"actionRequestCode,omitempty"`

	ProductDetails *ProductIdentificationDetailsTypeI_50878C `xml:"productDetails,omitempty"`
}

type AdditionalProductDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A AdditionalProductDetailsTypeI"`

	// Flight details
	LegDetails *AdditionalProductTypeI `xml:"legDetails,omitempty"`

	// Departure station
	DepartureStationInfo *StationInformationTypeI `xml:"departureStationInfo,omitempty"`

	ArrivalStationInfo *StationInformationTypeI `xml:"arrivalStationInfo,omitempty"`

	// Ground Time Details when connection
	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`
}

type AdditionalProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A AdditionalProductDetailsType"`

	// Type of aircraft
	EquipmentType formats.AlphaNumericString_Length1To3 `xml:"equipmentType,omitempty"`

	// Day number of the week
	OperatingDay formats.AlphaNumericString_Length1To7 `xml:"operatingDay,omitempty"`

	// Number of stops made in a journey if different from 0
	TechStopNumber formats.NumericInteger_Length1To2 `xml:"techStopNumber,omitempty"`

	// Location places of the stops
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`
}

type AdditionalProductTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A AdditionalProductTypeI"`

	// Equipment type
	Equipment formats.AlphaNumericString_Length1To3 `xml:"equipment,omitempty"`

	// Elaps flying time of the leg
	Duration formats.NumericInteger_Length1To6 `xml:"duration,omitempty"`

	// COG indicator
	ComplexingFlightIndicator formats.AlphaNumericString_Length1To1 `xml:"complexingFlightIndicator,omitempty"`
}

type AgentIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A AgentIdentificationType"`

	// Contains ARC number
	ArcNumber formats.AlphaNumericString_Length1To12 `xml:"arcNumber,omitempty"`

	// ERSP number of the Office or Customer
	ErspNumber formats.AlphaNumericString_Length1To12 `xml:"erspNumber,omitempty"`

	// IATA Number
	IataNumber formats.AlphaNumericString_Length1To12 `xml:"iataNumber,omitempty"`
}

type ArithmeticEvaluationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ArithmeticEvaluationType"`

	// Boolean operator
	CodeOperator formats.AlphaNumericString_Length1To3 `xml:"codeOperator,omitempty"`
}

type ArrivalLocalizationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ArrivalLocalizationType"`

	// Details on the location of the arrival point
	ArrivalPointDetails *ArrivalLocationDetailsType `xml:"arrivalPointDetails,omitempty"`

	// Arrival multi city option
	ArrivalMultiCity *MultiCityOptionType `xml:"arrivalMultiCity,omitempty"`

	// Attribute details
	AttributeDetails *CodedAttributeInformationType_139508C `xml:"attributeDetails,omitempty"`
}

type ArrivalLocationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ArrivalLocationDetailsType"`

	// For Radius: This is the length of the requested radius around the location of destination.
	Distance formats.NumericInteger_Length1To3 `xml:"distance,omitempty"`

	// Distance unit qualifier for radius
	DistanceUnit formats.AlphaNumericString_Length0To3 `xml:"distanceUnit,omitempty"`

	// ATA/IATA airport/city code of arrival  In case of SP request from a PNR, this field is empty.
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Request from no PNR A = consider the locationID specified as an airport (used only when ambiguous) C = consider the locationId as a city  (used only when ambiguous) A and C are not used in case of a request from PNR  Request from a PNR with Radius option: O = radius applies to the origin of the PNR segments D = radius applies to the destination of the PNR segments
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`

	// Latitude in degrees (decimal format) with hemisphere N=north S=south
	Latitude formats.AlphaNumericString_Length6To6 `xml:"latitude,omitempty"`

	// Longitude in degrees (decimal format) with hemisphere E=east, W=west
	Longitude formats.AlphaNumericString_Length6To6 `xml:"longitude,omitempty"`
}

type ArrivalLocationDetailsType_120834C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ArrivalLocationDetailsType_120834C"`

	// For Radius: This is the length of the requested radius around the location of origin.
	Distance formats.NumericInteger_Length1To3 `xml:"distance,omitempty"`

	// Distance unit qualifier for radius
	DistanceUnit formats.AlphaNumericString_Length1To3 `xml:"distanceUnit,omitempty"`

	// ATA/IATA airport/city code of arrival  In case of SP request from a PNR, this field is empty.
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Request from no PNR A = consider the locationID specified as an airport (used only when ambiguous) C = consider the locationId as a city  (used only when ambiguous) A and C are not used in case of a request from PNR  Request from a PNR with Radius option: O = radius applies to the origin of the PNR segments D = radius applies to the destination of the PNR segments
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`

	// Latitude in degrees (decimal format) with hemisphere N=north S=south
	Latitude formats.AlphaNumericString_Length6To6 `xml:"latitude,omitempty"`

	// Longitude in degrees (decimal format) with hemisphere E=east, W=west
	Longitude formats.AlphaNumericString_Length6To6 `xml:"longitude,omitempty"`
}

type AttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A AttributeInformationType"`

	// Code of options
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`

	// Mandatory when alternate date option is used (ALT), must be set to plus (P) or minus (M) a number of days around the original PNR segment dates. E.g.: M1 (for minus 1 day) or P100 (for plus 100 days)
	OptionInformation formats.AlphaNumericString_Length1To35 `xml:"optionInformation,omitempty"`
}

type AttributeInformationType_97181C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A AttributeInformationType_97181C"`

	// Attribute type
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType,omitempty"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type AttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A AttributeType"`

	// Option parameters
	SelectionDetails *AttributeInformationType `xml:"selectionDetails,omitempty"`
}

type AttributeType_61377S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A AttributeType_61377S"`

	// Criteria Set Type
	AttributeQualifier formats.AlphaNumericString_Length1To3 `xml:"attributeQualifier,omitempty"`

	// Criteria details
	AttributeDetails *AttributeInformationType_97181C `xml:"attributeDetails,omitempty"`
}

type BooleanExpressionRuleType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A BooleanExpressionRuleType"`

	// Boolean expression associated to the decision rule.
	BooleanExpression *ArithmeticEvaluationType `xml:"booleanExpression,omitempty"`
}

type CabinClassDesignationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CabinClassDesignationType"`

	// Cabin designator.
	CabinDesignator formats.AlphaNumericString_Length1To1 `xml:"cabinDesignator,omitempty"`
}

type CabinIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CabinIdentificationType"`

	// Cabin qualifier. For Star Pricer,MD stands for Mandatory Cabin qualifier.  For other products no qualifier stands for Mandatory Cabin.
	CabinQualifier formats.AlphaNumericString_Length1To2 `xml:"cabinQualifier,omitempty"`

	// Cabin
	Cabin formats.AlphaString_Length0To1 `xml:"cabin,omitempty"`
}

type CabinIdentificationType_233500C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CabinIdentificationType_233500C"`

	// Cabin qualifier. For Star Pricer,MD stands for Mandatory Cabin qualifier.  For other products no qualifier stands for Mandatory Cabin.
	CabinQualifier formats.AlphaNumericString_Length1To2 `xml:"cabinQualifier,omitempty"`

	// Cabin
	Cabin formats.AlphaString_Length0To1 `xml:"cabin,omitempty"`
}

type CabinProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CabinProductDetailsType"`

	// Reservation booking designator - RBD
	Rbd formats.AlphaString_Length1To1 `xml:"rbd,omitempty"`

	// Reservation Booking Modifier
	BookingModifier formats.AMA_EDICodesetType_Length1 `xml:"bookingModifier,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin formats.AlphaString_Length1To1 `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus formats.AMA_EDICodesetType_Length1to3 `xml:"avlStatus,omitempty"`
}

type CodedAttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeInformationType"`

	// Attribute type identification
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// Attribute Description
	AttributeDescription formats.AlphaNumericString_Length1To20 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeInformationType_120742C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeInformationType_120742C"`

	// Name.
	Name formats.AlphaNumericString_Length1To5 `xml:"name,omitempty"`

	// Value.
	Value formats.AlphaNumericString_Length1To20 `xml:"value,omitempty"`
}

type CodedAttributeInformationType_139508C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeInformationType_139508C"`

	// Attribute type
	Type formats.AlphaNumericString_Length1To5 `xml:"type,omitempty"`

	// Value.
	Value formats.AlphaNumericString_Length1To20 `xml:"value,omitempty"`
}

type CodedAttributeInformationType_247828C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeInformationType_247828C"`

	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To10 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeInformationType_247829C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeInformationType_247829C"`

	// Type of fee/reduction
	FeeType formats.AlphaNumericString_Length1To5 `xml:"feeType,omitempty"`

	// Fee Id Number
	FeeIdNumber formats.AlphaNumericString_Length1To50 `xml:"feeIdNumber,omitempty"`
}

type CodedAttributeInformationType_254574C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeInformationType_254574C"`

	// Attribute type
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To50 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeType"`

	// Fee/reduction Id
	FeeId *CodedAttributeInformationType_247829C `xml:"feeId,omitempty"`
}

type CodedAttributeType_181239S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeType_181239S"`

	// Attribute details
	AttributeDetails *CodedAttributeInformationType_254574C `xml:"attributeDetails,omitempty"`
}

type CodedAttributeType_78500S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CodedAttributeType_78500S"`

	// Atrribute details.
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`
}

type CommercialAgreementsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CommercialAgreementsType"`

	// Codeshare Details
	CodeshareDetails *CompanyRoleIdentificationType `xml:"codeshareDetails,omitempty"`

	// Other codeshare details
	OtherCodeshareDetails *CompanyRoleIdentificationType `xml:"otherCodeshareDetails,omitempty"`
}

type CommercialAgreementsType_78540S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CommercialAgreementsType_78540S"`

	// Codeshare Details
	CodeshareDetails *CompanyRoleIdentificationType_120761C `xml:"codeshareDetails,omitempty"`

	// Other codeshare details
	OtherCodeshareDetails *CompanyRoleIdentificationType_120761C `xml:"otherCodeshareDetails,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CompanyIdentificationTypeI"`

	// Carrier code
	MarketingCompany formats.AlphaNumericString_Length2To3 `xml:"marketingCompany,omitempty"`
}

type CompanyIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CompanyIdentificationType"`
}

type CompanyIdentificationType_120719C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CompanyIdentificationType_120719C"`

	// Item description identification.
	CarrierQualifier formats.AlphaNumericString_Length0To1 `xml:"carrierQualifier,omitempty"`

	// carrier id
	CarrierId formats.AlphaNumericString_Length2To3 `xml:"carrierId,omitempty"`
}

type CompanyIdentificationType_195544C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CompanyIdentificationType_195544C"`

	// Marketing carrier
	MarketingCarrier formats.AlphaNumericString_Length2To3 `xml:"marketingCarrier,omitempty"`

	// Operating carrier
	OperatingCarrier formats.AlphaNumericString_Length2To3 `xml:"operatingCarrier,omitempty"`
}

type CompanyIdentificationType_233548C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CompanyIdentificationType_233548C"`

	// Item description identification.
	CarrierQualifier formats.AlphaNumericString_Length0To1 `xml:"carrierQualifier,omitempty"`

	// carrier id
	CarrierId formats.AlphaNumericString_Length2To3 `xml:"carrierId,omitempty"`
}

type CompanyRoleIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CompanyRoleIdentificationType"`

	// Codeshare qualifier
	TransportStageQualifier formats.AlphaString_Length1To1 `xml:"transportStageQualifier,omitempty"`

	// company identification
	AirlineDesignator formats.AlphaNumericString_Length2To3 `xml:"airlineDesignator,omitempty"`

	// flight number
	FlightNumber formats.NumericInteger_Length1To4 `xml:"flightNumber,omitempty"`

	// suffix
	OperationalSuffix formats.AlphaString_Length1To1 `xml:"operationalSuffix,omitempty"`
}

type CompanyRoleIdentificationType_120761C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CompanyRoleIdentificationType_120761C"`

	// Type of code share agreement.
	CodeShareType formats.AlphaString_Length1To1 `xml:"codeShareType,omitempty"`

	// company identification
	AirlineDesignator formats.AlphaNumericString_Length2To3 `xml:"airlineDesignator,omitempty"`

	// flight number
	FlightNumber formats.NumericInteger_Length1To4 `xml:"flightNumber,omitempty"`
}

type ConnectPointDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ConnectPointDetailsType"`

	// Exclusion identification
	ExclusionIdentifier formats.AlphaString_Length0To1 `xml:"exclusionIdentifier,omitempty"`

	// Place or Location identification
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Airport/city qualifier
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`
}

type ConnectPointDetailsType_195492C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ConnectPointDetailsType_195492C"`

	// Inclusion identification
	InclusionIdentifier formats.AlphaNumericString_Length0To1 `xml:"inclusionIdentifier,omitempty"`

	// Place or Location identification
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Airport/city qualifier
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`
}

type ConnectionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ConnectionDetailsTypeI"`

	// Location
	Location formats.AlphaNumericString_Length1To3 `xml:"location,omitempty"`
}

type ConnectionTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ConnectionTypeI"`

	// Connection details.
	ConnectionDetails *ConnectionDetailsTypeI `xml:"connectionDetails,omitempty"`
}

type ConsumerReferenceIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ConsumerReferenceIdentificationTypeI"`

	ReferenceQualifier formats.AlphaNumericString_Length1To3 `xml:"referenceQualifier,omitempty"`

	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber,omitempty"`

	ReferencePartyName formats.AlphaNumericString_Length1To35 `xml:"referencePartyName,omitempty"`

	TravellerReferenceNbr formats.AlphaNumericString_Length1To10 `xml:"travellerReferenceNbr,omitempty"`
}

type ConsumerReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ConsumerReferenceInformationType"`

	// Customer references
	CustomerReferences *ConsumerReferenceIdentificationTypeI `xml:"customerReferences,omitempty"`
}

type ConversionRateDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ConversionRateDetailsType"`

	// Conversion type
	ConversionType formats.AlphaNumericString_Length1To3 `xml:"conversionType,omitempty"`

	// Currency
	Currency formats.AlphaString_Length1To3 `xml:"currency,omitempty"`
}

type ConversionRateType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ConversionRateType"`

	// Detail of conversion rate of First Monetary Unit
	ConversionRateDetail *ConversionRateDetailsType `xml:"conversionRateDetail,omitempty"`
}

type CorporateFareIdentifiersType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CorporateFareIdentifiersType"`

	// Indicates the type of corporate fares requested
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier,omitempty"`

	// Corporate contract number or name
	IdentifyNumber formats.AlphaNumericString_Length1To35 `xml:"identifyNumber,omitempty"`
}

type CorporateFareInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CorporateFareInformationType"`

	// Corporate fare identifiers
	CorporateFareIdentifiers *CorporateFareIdentifiersType `xml:"corporateFareIdentifiers,omitempty"`
}

type CorporateIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CorporateIdentificationType"`

	// Corporate identity
	CorporateId *CorporateIdentityType `xml:"corporateId,omitempty"`
}

type CorporateIdentityType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CorporateIdentityType"`

	// Indicates if 1A corporate (RC) or Unifare corporate (RW) requested.
	CorporateQualifier formats.AlphaNumericString_Length0To3 `xml:"corporateQualifier,omitempty"`

	// Corporate contract number or name
	Identity formats.AlphaNumericString_Length1To20 `xml:"identity,omitempty"`
}

type CriteriaiDetaislType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A CriteriaiDetaislType"`

	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	Value formats.AlphaNumericString_Length1To18 `xml:"value,omitempty"`

	Attribute formats.AlphaNumericString_Length1To9 `xml:"attribute,omitempty"`
}

type DataInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DataInformationType"`

	// Ancillary services options
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type DataTypeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DataTypeInformationType"`

	// Carrier fee code
	SubType formats.AlphaNumericString_Length1To3 `xml:"subType,omitempty"`

	// Status (automated, manually added, exempted). Default is automated
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`
}

type DateAndTimeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateAndTimeDetailsTypeI"`

	// Toidentify type of time Arrival/Departure
	TimeQualifier formats.AlphaNumericString_Length1To3 `xml:"timeQualifier,omitempty"`

	// Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// Time
	Time formats.Time24_HHMM `xml:"time,omitempty"`

	// Time window size in hours
	TimeWindow formats.AlphaNumericString_Length1To3 `xml:"timeWindow,omitempty"`
}

type DateAndTimeDetailsTypeI_120740C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateAndTimeDetailsTypeI_120740C"`

	// Date time period qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// First Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// First Time
	Time formats.Time24_HHMM `xml:"time,omitempty"`

	// .
	Qualifier2 formats.AlphaNumericString_Length1To3 `xml:"qualifier2,omitempty"`

	Reserved1 formats.AlphaNumericString_Length1To3 `xml:"reserved1,omitempty"`

	Reserved2 formats.AlphaNumericString_Length3To5 `xml:"reserved2,omitempty"`
}

type DateAndTimeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateAndTimeDetailsType"`

	// Trip Duration type(Plus,Minus,Combined)
	FlexibilityQualifier formats.AlphaNumericString_Length1To3 `xml:"flexibilityQualifier,omitempty"`

	// Number of days added or/and retrieved to the trip duration
	TripInterval formats.NumericInteger_Length1To6 `xml:"tripInterval,omitempty"`

	// Period between date of departure and date of arrival
	TripDuration formats.NumericInteger_Length1To4 `xml:"tripDuration,omitempty"`
}

type DateAndTimeDetailsType_120762C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateAndTimeDetailsType_120762C"`

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

type DateAndTimeDetailsType_254619C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateAndTimeDetailsType_254619C"`

	// To identify type of range (Plus,Minus or combined)
	RangeQualifier formats.AlphaNumericString_Length1To3 `xml:"rangeQualifier,omitempty"`

	// Range of dates : Number of Days preceding or/and Following the request departure date
	DayInterval formats.NumericInteger_Length1To6 `xml:"dayInterval,omitempty"`

	// Time at destination (local)
	TimeAtdestination formats.Time24_HHMM `xml:"timeAtdestination,omitempty"`
}

type DateAndTimeInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateAndTimeInformationTypeI"`

	// DATE AND TIME DETAILS.
	DateTimeDetails *DateAndTimeDetailsTypeI_120740C `xml:"dateTimeDetails,omitempty"`
}

type DateAndTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateAndTimeInformationType"`

	// Details on date and time
	StopDetails *DateAndTimeDetailsType_120762C `xml:"stopDetails,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DummyNET struct{}

type DateAndTimeInformationType_181295S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateAndTimeInformationType_181295S"`

	// Details on date and Time
	FirstDateTimeDetail *DateAndTimeDetailsTypeI `xml:"firstDateTimeDetail,omitempty"`

	// Details of the Requested Range of Dates
	RangeOfDate *DateAndTimeDetailsType_254619C `xml:"rangeOfDate,omitempty"`

	// Details of the trip duration
	TripDetails *DateAndTimeDetailsType `xml:"tripDetails,omitempty"`
}

type DateTimePeriodDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DateTimePeriodDetailsTypeI"`

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Value
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type DepartureLocationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A DepartureLocationType"`

	// Details on localization of the departure point
	DeparturePoint *ArrivalLocationDetailsType_120834C `xml:"departurePoint,omitempty"`

	// Departure multi city option
	DepMultiCity *MultiCityOptionType `xml:"depMultiCity,omitempty"`

	// To specify a series or a range of PNR segments
	FirstPnrSegmentRef *PNRSegmentReferenceType `xml:"firstPnrSegmentRef,omitempty"`

	// Attribute details
	AttributeDetails *CodedAttributeInformationType_139508C `xml:"attributeDetails,omitempty"`
}

type FOPRepresentationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FOPRepresentationType"`

	// Form of payment information.
	FormOfPaymentDetails *FormOfPaymentTypeI `xml:"formOfPaymentDetails,omitempty"`

	// Fee/Reduction Passenger reference.
	PassengerFeeReference *ItemReferencesAndVersionsType `xml:"passengerFeeReference,omitempty"`
}

type FareDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareDetailsType"`

	// Qualifier
	Qualifier formats.AMA_EDICodesetType_Length1to3 `xml:"qualifier,omitempty"`

	// Rate
	Rate formats.NumericInteger_Length1To8 `xml:"rate,omitempty"`

	// Country
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`

	// Fare Category
	FareCategory formats.AMA_EDICodesetType_Length1to3 `xml:"fareCategory,omitempty"`
}

type FareFamilyCriteriaType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareFamilyCriteriaType"`

	// Fare publishing carrier.
	CarrierId formats.AlphaNumericString_Length1To3 `xml:"carrierId,omitempty"`

	// Reservation booking designator.
	Rdb formats.AlphaString_Length1To2 `xml:"rdb,omitempty"`

	// Fare family info.
	FareFamilyInfo *FareQualifierInformationType `xml:"fareFamilyInfo,omitempty"`

	// Fare product detail.
	FareProductDetail *FareProductDetailsType `xml:"fareProductDetail,omitempty"`

	// Corporate information.
	CorporateInfo *MultipleIdentificationNumbersTypeI `xml:"corporateInfo,omitempty"`

	// Indicates flight cabin details.
	CabinProduct *CabinClassDesignationType `xml:"cabinProduct,omitempty"`

	// Cabin processing option.
	CabinProcessingIdentifier formats.AlphaNumericString_Length1To3 `xml:"cabinProcessingIdentifier,omitempty"`

	// Product date or time.
	DateTimeDetails *ProductDateTimeTypeI_194583C `xml:"dateTimeDetails,omitempty"`

	// Other criteria.
	OtherCriteria *CodedAttributeInformationType_120742C `xml:"otherCriteria,omitempty"`
}

type FareFamilyDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareFamilyDetailsType"`

	// Commercial fare Family Short name
	CommercialFamily formats.AlphaNumericString_Length1To10 `xml:"commercialFamily,omitempty"`
}

type FareFamilyType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareFamilyType"`

	// Fare Family Reference Number
	RefNumber formats.NumericInteger_Length1To3 `xml:"refNumber,omitempty"`

	// Fare Family Short Name
	FareFamilyname formats.AlphaNumericString_Length1To10 `xml:"fareFamilyname,omitempty"`

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy formats.NumericInteger_Length1To4 `xml:"hierarchy,omitempty"`

	// Indicates Commercial Fare Family Short names
	CommercialFamilyDetails *FareFamilyDetailsType `xml:"commercialFamilyDetails,omitempty"`
}

type FareInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareInformationTypeI"`

	// PTC (Full Codeset List described in ATPCo Documentation B11109 Appendix A)
	ValueQualifier formats.AlphaNumericString_Length1To3 `xml:"valueQualifier,omitempty"`

	// age
	Value formats.NumericInteger_Length1To15 `xml:"value,omitempty"`
}

type FareInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareInformationType"`

	// Value Qualifier
	ValueQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"valueQualifier,omitempty"`

	// Value
	Value formats.NumericInteger_Length1To15 `xml:"value,omitempty"`

	// Fare Details
	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`

	// Identity Number
	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber,omitempty"`

	// Fare Type Grouping
	FareTypeGrouping *FareTypeGroupingInformationType `xml:"fareTypeGrouping,omitempty"`

	// Rate Category
	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"`
}

type FareProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareProductDetailsType"`

	// Fare basis code
	FareBasis formats.AlphaNumericString_Length0To18 `xml:"fareBasis,omitempty"`

	// Type of fare
	FareType formats.AlphaNumericString_Length0To3 `xml:"fareType,omitempty"`
}

type FareQualifierInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareQualifierInformationType"`

	// Fare family combinability.
	FareFamilyQual formats.AlphaNumericString_Length0To3 `xml:"fareFamilyQual,omitempty"`
}

type FareTypeGroupingInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FareTypeGroupingInformationType"`

	// Pricing Group
	PricingGroup formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"`
}

type FlightProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FlightProductInformationType"`

	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType `xml:"cabinProduct,omitempty"`

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType_205137C `xml:"contextDetails,omitempty"`
}

type FormOfPaymentDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FormOfPaymentDetailsTypeI"`

	// Form of payment identification
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// amount to be charged on this form
	ChargedAmount formats.NumericDecimal_Length1To12 `xml:"chargedAmount,omitempty"`

	// Reference number
	CreditCardNumber formats.AlphaNumericString_Length1To20 `xml:"creditCardNumber,omitempty"`
}

type FormOfPaymentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FormOfPaymentTypeI"`

	// FORM OF PAYMENT DETAILS
	FormOfPaymentDetails *FormOfPaymentDetailsTypeI `xml:"formOfPaymentDetails,omitempty"`
}

type FreeTextDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FreeTextDetailsType"`

	// Qualifier of the free text. Always literal in our usage.
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Type of the free text. Always 7 for our usage.
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// Source of the information.
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// Encoding method used.
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FreeTextInformationType"`

	// Details for the free text message
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text corresponding to the DEI 127 data.
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type FrequencyType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FrequencyType"`

	// Indicate if the sequence number represents days of the week or days of the month.
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Used to represent days of the week. 1 is monday and 7 is sunday.
	Value formats.NumericInteger_Length1To1 `xml:"value,omitempty"`
}

type FrequentTravellerIdentificationCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FrequentTravellerIdentificationCodeType"`

	// Frequent Traveller Info
	FrequentTravellerDetails *FrequentTravellerIdentificationType `xml:"frequentTravellerDetails,omitempty"`
}

type FrequentTravellerIdentificationCodeType_177150S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FrequentTravellerIdentificationCodeType_177150S"`

	// Frequent traveller details
	FrequentTravellerDetails *FrequentTravellerIdentificationType_249074C `xml:"frequentTravellerDetails,omitempty"`
}

type FrequentTravellerIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FrequentTravellerIdentificationType"`

	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"`

	Number formats.AlphaNumericString_Length1To28 `xml:"number,omitempty"`

	// Specifies which traveller in the TIF segment the frequent traveller number applies (same as 9944 in TIF).
	CustomerReference formats.AlphaNumericString_Length1To10 `xml:"customerReference,omitempty"`

	// status code: 'OK' if the frequent flyer card has been validated
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`

	// For example : priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"`

	// For example : Level description
	TierDescription formats.AlphaNumericString_Length1To35 `xml:"tierDescription,omitempty"`

	// For example : Company code of alliance
	CompanyCode formats.AlphaNumericString_Length1To35 `xml:"companyCode,omitempty"`

	CustomerValue formats.NumericInteger_Length1To4 `xml:"customerValue,omitempty"`

	// To specify the product/account number qualifier. (e.g. 2=Corporate Frequent Flyer).
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type FrequentTravellerIdentificationType_249074C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A FrequentTravellerIdentificationType_249074C"`

	// carrier
	Carrier formats.AlphaNumericString_Length1To3 `xml:"carrier,omitempty"`

	// Frequent traveller number
	Number formats.AlphaNumericString_Length1To25 `xml:"number,omitempty"`

	// Customer reference
	CustomerReference formats.AlphaNumericString_Length1To10 `xml:"customerReference,omitempty"`

	// Tier level
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`

	// Priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"`

	// Tier description
	TierDescription formats.AlphaNumericString_Length1To35 `xml:"tierDescription,omitempty"`

	// To specify the product/account number qualifier. (e.g. 2=Corporate Frequent Flyer).
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type GroupPassengerDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A GroupPassengerDetailsType"`

	// Trigger
	PassengerReference *SegmentRepetitionControlTypeI `xml:"passengerReference,omitempty"`

	PsgDetailsInfo *PsgDetailsInfo `xml:"psgDetailsInfo,omitempty"`
}

type PsgDetailsInfo struct {

	// PTC/Discount Code age
	DiscountPtc *FareInformationTypeI `xml:"discountPtc,omitempty"`

	// Tier level information
	FlequentFlyerDetails *FrequentTravellerIdentificationCodeType `xml:"flequentFlyerDetails,omitempty"`
}

type HeaderInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A HeaderInformationTypeI"`

	// Status
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Date and Time info
	DateTimePeriodDetails *DateTimePeriodDetailsTypeI `xml:"dateTimePeriodDetails,omitempty"`

	// Reference number
	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber,omitempty"`

	// Contains product identification such as UIC code...
	ProductIdentification formats.AlphaNumericString_Length1To35 `xml:"productIdentification,omitempty"`
}

type ItemNumberIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ItemNumberIdentificationType"`

	// Fare family combination number
	Number formats.AlphaNumericString_Length1To4 `xml:"number,omitempty"`

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Responsible agency
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`
}

type ItemNumberType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ItemNumberType"`

	// Indicates the fare family combination number
	ItemNumberId *ItemNumberIdentificationType `xml:"itemNumberId,omitempty"`
}

type ItemNumberType_80866S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ItemNumberType_80866S"`

	// Item number details
	ItemNumberDetails *ItemNumberIdentificationType `xml:"itemNumberDetails,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ItemReferencesAndVersionsType"`

	// Reference Qualifier.
	PassengerFeeRefType formats.AlphaNumericString_Length1To3 `xml:"passengerFeeRefType,omitempty"`

	// Reference number.
	PassengerFeeRefNumber formats.NumericInteger_Length1To3 `xml:"passengerFeeRefNumber,omitempty"`

	// Unique id description.
	OtherCharacteristics *UniqueIdDescriptionType `xml:"otherCharacteristics,omitempty"`
}

type ItineraryDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ItineraryDetailsType"`

	// Airport/City Qualifier: the passenger wants to depart/arrive from/to the same airport or city in the specified requested segment
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`

	// Requested segment number
	SegmentNumber formats.NumericInteger_Length1To3 `xml:"segmentNumber,omitempty"`
}

type LocationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A LocationDetailsTypeI"`

	// Place or Location identification
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Country, coded
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`
}

type LocationIdentificationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A LocationIdentificationDetailsType"`

	// 3 characters ATA/IATA airport/city code
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Airport/city qualifier: the requested point is an airport when ambiguity exists (e.g. HOU)
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`

	// Terminal information
	Terminal formats.AlphaNumericString_Length1To5 `xml:"terminal,omitempty"`
}

type LocationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A LocationTypeI"`

	// Departure or Arrival IATA airport code
	TrueLocationId formats.AlphaString_Length3To3 `xml:"trueLocationId,omitempty"`
}

type MileageTimeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MileageTimeDetailsTypeI"`

	// Ground Time in minutes at Board point (connection with incoming flight)
	ElapsedGroundTime formats.NumericInteger_Length1To4 `xml:"elapsedGroundTime,omitempty"`
}

type MonetaryAndCabinInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryAndCabinInformationDetailsType"`

	// Amount qualifier.
	AmountType formats.AlphaNumericString_Length0To3 `xml:"amountType,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// ISO currency code
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Airport/city code
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Cabin class designator
	CabinClassDesignator formats.AlphaString_Length1To1 `xml:"cabinClassDesignator,omitempty"`
}

type MonetaryAndCabinInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryAndCabinInformationType"`

	// Monetary and cabin information.
	MoneyAndCabinInfo *MonetaryAndCabinInformationDetailsType `xml:"moneyAndCabinInfo,omitempty"`
}

type MonetaryInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryInformationDetailsTypeI"`

	// Monetary amount type qualifier, coded
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Allowance or charge number
	Amount formats.NumericInteger_Length1To18 `xml:"amount,omitempty"`

	// Currency, coded
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsTypeI_194597C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryInformationDetailsTypeI_194597C"`

	// Monetary amount type qualifier, coded
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Allowance or charge number
	Amount formats.NumericInteger_Length1To18 `xml:"amount,omitempty"`

	// Currency, coded
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Place or Location identification
	LocationId formats.AlphaString_Length3To3 `xml:"locationId,omitempty"`
}

type MonetaryInformationDetailsTypeI_65140C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryInformationDetailsTypeI_65140C"`

	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsTypeI_65141C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryInformationDetailsTypeI_65141C"`

	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Location
	Location formats.AlphaNumericString_Length1To3 `xml:"location,omitempty"`
}

type MonetaryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryInformationTypeI"`

	// Monetary info
	MonetaryDetails *MonetaryInformationDetailsTypeI_65141C `xml:"monetaryDetails,omitempty"`
}

type MonetaryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryInformationType"`

	// MONEY INFORMATION
	MoneyInfo *MonetaryInformationDetailsTypeI `xml:"moneyInfo,omitempty"`

	// MONEY INFORMATION
	AdditionalMoneyInfo *MonetaryInformationDetailsTypeI_194597C `xml:"additionalMoneyInfo,omitempty"`
}

type MonetaryInformationType_80162S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MonetaryInformationType_80162S"`

	// Monetary info
	MonetaryDetails *MonetaryInformationDetailsTypeI_65140C `xml:"monetaryDetails,omitempty"`
}

type MultiCityOptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MultiCityOptionType"`

	// ATA/IATA airport/city code of arrival multi city option enable to define until 20 airports/cities
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`

	// Requested arrival point is an airport or a city (default is city and used only when ambiguity)
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`
}

type MultipleIdentificationNumbersTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A MultipleIdentificationNumbersTypeI"`

	// Corporate number or ALL.
	CorporateNumberIdentifier formats.AlphaNumericString_Length1To12 `xml:"corporateNumberIdentifier,omitempty"`

	// Corporate name.
	CorporateName formats.AlphaNumericString_Length1To20 `xml:"corporateName,omitempty"`
}

type NumberOfUnitDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A NumberOfUnitDetailsTypeI"`

	// Number of Units
	NumberOfUnits formats.NumericInteger_Length1To3 `xml:"numberOfUnits,omitempty"`

	// Number of unit qualifier
	TypeOfUnit formats.AlphaNumericString_Length1To3 `xml:"typeOfUnit,omitempty"`
}

type NumberOfUnitDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A NumberOfUnitDetailsType"`

	// Number of Units
	NumberOfUnits formats.NumericInteger_Length1To4 `xml:"numberOfUnits,omitempty"`

	// Number of unit qualifier
	TypeOfUnit formats.AlphaNumericString_Length1To3 `xml:"typeOfUnit,omitempty"`
}

type NumberOfUnitDetailsType_260583C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A NumberOfUnitDetailsType_260583C"`

	// Number of Units
	NumberOfUnits formats.NumericInteger_Length1To6 `xml:"numberOfUnits,omitempty"`

	// Number of unit qualifier
	TypeOfUnit formats.AlphaNumericString_Length1To3 `xml:"typeOfUnit,omitempty"`
}

type NumberOfUnitsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A NumberOfUnitsType"`

	// NUMBER OF UNIT DETAILS
	UnitNumberDetail *NumberOfUnitDetailsType_260583C `xml:"unitNumberDetail,omitempty"`
}

type NumberOfUnitsType_80154S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A NumberOfUnitsType_80154S"`

	// NUMBER OF UNIT DETAILS
	UnitNumberDetail *NumberOfUnitDetailsType `xml:"unitNumberDetail,omitempty"`
}

type OriginAndDestinationRequestType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A OriginAndDestinationRequestType"`

	// Requested segment number
	SegRef formats.NumericInteger_Length1To2 `xml:"segRef,omitempty"`

	// Forces arrival or departure, from/to the same airport/city
	LocationForcing *ItineraryDetailsType `xml:"locationForcing,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A OriginatorIdentificationDetailsTypeI"`

	// Office Name.
	OfficeName formats.NumericInteger_Length1To9 `xml:"officeName,omitempty"`

	// Agent Sign In .
	AgentSignin formats.AlphaNumericString_Length1To9 `xml:"agentSignin,omitempty"`

	// Confidential Office Name.
	ConfidentialOffice formats.AlphaNumericString_Length1To9 `xml:"confidentialOffice,omitempty"`

	// Other Office Name
	OtherOffice formats.AlphaNumericString_Length1To9 `xml:"otherOffice,omitempty"`
}

type PNRSegmentReferenceType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A PNRSegmentReferenceType"`

	// For a request from PNR:  this is the reference number of a PNR air segment. In case a range of PNR segments is specified (eg. segments 2-5), then it is the 1st of the range, the last being in ARR.
	PnrSegmentTattoo formats.NumericInteger_Length0To35 `xml:"pnrSegmentTattoo,omitempty"`

	PnrSegmentQualifier formats.AlphaString_Length1To1 `xml:"pnrSegmentQualifier,omitempty"`
}

type PassengerItineraryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A PassengerItineraryInformationType"`

	// .
	Booking formats.AlphaString_Length1To1 `xml:"booking,omitempty"`

	// .
	Identifier formats.AlphaNumericString_Length1To1 `xml:"identifier,omitempty"`

	// .
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// .
	ItemNumber formats.NumericInteger_Length1To3 `xml:"itemNumber,omitempty"`

	// .
	DateTimeDetails *ProductDateTimeType `xml:"dateTimeDetails,omitempty"`

	// .
	Designator formats.AlphaString_Length1To1 `xml:"designator,omitempty"`

	// .
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`

	// .
	ProductTypeDetails *ProductTypeDetailsType `xml:"productTypeDetails,omitempty"`
}

type PricingTicketingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A PricingTicketingDetailsType"`

	// Pricing ticketing Details.
	PricingTicketing *PricingTicketingInformationType `xml:"pricingTicketing,omitempty"`

	// PRODUCT DATE OR TIME
	TicketingDate *ProductDateTimeTypeI_194598C `xml:"ticketingDate,omitempty"`

	// COMPANY IDENTIFICATION
	CompanyId *CompanyIdentificationType `xml:"companyId,omitempty"`

	// LOCATION DETAILS
	SellingPoint *LocationDetailsTypeI `xml:"sellingPoint,omitempty"`

	// LOCATION DETAILS
	TicketingPoint *LocationDetailsTypeI `xml:"ticketingPoint,omitempty"`

	// Used to Target Transborder Fares
	JourneyOriginPoint *LocationDetailsTypeI `xml:"journeyOriginPoint,omitempty"`

	// Contains the ARC,IATA and ERSP numbers
	CorporateId *AgentIdentificationType `xml:"corporateId,omitempty"`
}

type PricingTicketingInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A PricingTicketingInformationType"`

	// Price type qualifier
	PriceType formats.AlphaNumericString_Length0To3 `xml:"priceType,omitempty"`
}

type ProductDateTimeTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductDateTimeTypeI"`

	// Departure date in YYYYMMDD format
	DepartureDate formats.Date_YYYYMMDD `xml:"departureDate,omitempty"`

	// Departure time
	DepartureTime formats.Time24_HHMM `xml:"departureTime,omitempty"`

	// Arrival date
	ArrivalDate formats.Date_YYYYMMDD `xml:"arrivalDate,omitempty"`

	// Arrival time
	ArrivalTime formats.Time24_HHMM `xml:"arrivalTime,omitempty"`

	// Day difference between Departure date of the leg and date of reference (Departure or Arrival date specified in the SDI)
	DateVariation formats.NumericInteger_Length1To1 `xml:"dateVariation,omitempty"`
}

type ProductDateTimeTypeI_194583C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductDateTimeTypeI_194583C"`

	// Ticketing Purchase Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// Ticketing purchase date
	OtherDate formats.Date_DDMMYY `xml:"otherDate,omitempty"`
}

type ProductDateTimeTypeI_194598C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductDateTimeTypeI_194598C"`

	// First date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// Half round trip combination.
	RtcDate formats.Date_DDMMYY `xml:"rtcDate,omitempty"`
}

type ProductDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductDateTimeType"`

	// .
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// .
	Time formats.Time24_HHMM `xml:"time,omitempty"`
}

type ProductDateTimeType_195546C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductDateTimeType_195546C"`

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

type ProductFacilitiesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductFacilitiesType"`

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

type ProductIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductIdentificationDetailsTypeI"`

	// Flight number
	FlightNumber formats.NumericInteger_Length1To4 `xml:"flightNumber,omitempty"`

	// Flight suffix
	OperationalSuffix formats.AlphaString_Length1To1 `xml:"operationalSuffix,omitempty"`
}

type ProductIdentificationDetailsTypeI_50878C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductIdentificationDetailsTypeI_50878C"`

	FlightNumber formats.AlphaNumericString_Length1To5 `xml:"flightNumber,omitempty"`

	BookingClass formats.AlphaNumericString_Length1To2 `xml:"bookingClass,omitempty"`

	OperationalSuffix formats.AlphaNumericString_Length1To3 `xml:"operationalSuffix,omitempty"`

	Modifier formats.AlphaNumericString_Length1To7 `xml:"modifier,omitempty"`
}

type ProductLocationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductLocationDetailsTypeI"`

	// airport
	Station formats.AlphaString_Length3To3 `xml:"station,omitempty"`
}

type ProductTypeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductTypeDetailsTypeI"`

	// Part of the journey (C,E,S), Codeshare service (A), Technical stop at off point in a Direct (TSD), Technical stop at off point in a COG (TSC), E-Ticket candidate (ET), Prohibited Countries (RPC, WPC)
	FlightIndicator formats.AlphaString_Length1To3 `xml:"flightIndicator,omitempty"`
}

type ProductTypeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductTypeDetailsType"`

	// .
	SequenceNumber formats.AlphaNumericString_Length1To6 `xml:"sequenceNumber,omitempty"`

	// PNR availability context
	AvailabilityContext formats.AlphaNumericString_Length1To6 `xml:"availabilityContext,omitempty"`
}

type ProductTypeDetailsType_120801C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductTypeDetailsType_120801C"`

	// Type of flight
	FlightType formats.AlphaNumericString_Length1To2 `xml:"flightType,omitempty"`
}

type ProductTypeDetailsType_205137C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProductTypeDetailsType_205137C"`

	// indicates whether the flight is domestic or international
	Avl formats.AlphaNumericString_Length1To6 `xml:"avl,omitempty"`
}

type ProposedSegmentDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProposedSegmentDetailsType"`

	// Flight proposal reference
	Ref formats.AlphaNumericString_Length1To6 `xml:"ref,omitempty"`

	// Elapse Flying Time
	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`
}

type ProposedSegmentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ProposedSegmentType"`

	// Parameters for proposed flight group
	FlightProposal *ProposedSegmentDetailsType `xml:"flightProposal,omitempty"`

	// Flight characteristics.
	FlightCharacteristic formats.AlphaNumericString_Length0To3 `xml:"flightCharacteristic,omitempty"`

	// Majority cabin
	MajCabin formats.AlphaString_Length1To1 `xml:"majCabin,omitempty"`
}

type ReferenceInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ReferenceInfoType"`

	// Referencing details
	ReferencingDetail *ReferencingDetailsType `xml:"referencingDetail,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type ReferencingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ReferencingDetailsType"`

	// Segment reference qualifier
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Flight or flight group reference
	RefNumber formats.NumericInteger_Length0To3 `xml:"refNumber,omitempty"`
}

type RoutingInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A RoutingInformationTypeI"`

	// Stops details
	RoutingDetails *ProductLocationDetailsTypeI `xml:"routingDetails,omitempty"`
}

type SegmentRepetitionControlDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SegmentRepetitionControlDetailsTypeI"`

	// traveller number
	Quantity formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"`
}

type SegmentRepetitionControlTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SegmentRepetitionControlTypeI"`

	// Segment control details
	SegmentControlDetails *SegmentRepetitionControlDetailsTypeI `xml:"segmentControlDetails,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SelectionDetailsInformationTypeI"`

	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`

	OptionInformation formats.AlphaNumericString_Length1To35 `xml:"optionInformation,omitempty"`
}

type SelectionDetailsInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SelectionDetailsInformationType"`

	// Carrier fee type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Carrier fee status
	OptionInformation formats.AlphaNumericString_Length1To3 `xml:"optionInformation,omitempty"`
}

type SelectionDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SelectionDetailsType"`

	// Carrier fees options
	CarrierFeeDetails *SelectionDetailsInformationType `xml:"carrierFeeDetails,omitempty"`

	OtherSelectionDetails *SelectionDetailsInformationTypeI `xml:"otherSelectionDetails,omitempty"`
}

type SpecialRequirementsDataDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SpecialRequirementsDataDetailsType"`

	// SSR seat characteristic
	SeatCharacteristics formats.AlphaNumericString_Length1To2 `xml:"seatCharacteristics,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type SpecialRequirementsDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SpecialRequirementsDetailsType"`

	// To specify the Service Requirement of the customer
	ServiceRequirementsInfo *SpecialRequirementsTypeDetailsType `xml:"serviceRequirementsInfo,omitempty"`

	// Seat details
	SeatDetails *SpecialRequirementsDataDetailsType `xml:"seatDetails,omitempty"`
}

type SpecialRequirementsTypeDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SpecialRequirementsTypeDetailsType"`

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
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A SpecificDataInformationType"`

	// Carrier fee description
	DataTypeInformation *DataTypeInformationType `xml:"dataTypeInformation,omitempty"`

	// Data information
	DataInformation *DataInformationType `xml:"dataInformation,omitempty"`
}

type StationInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A StationInformationTypeI"`

	// Departure terminal
	Terminal formats.AlphaNumericString_Length1To3 `xml:"terminal,omitempty"`
}

type StructuredDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A StructuredDateTimeType"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.NumericInteger_Length4To4 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.NumericInteger_Length1To2 `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.NumericInteger_Length1To2 `xml:"day,omitempty"`
}

type StructuredPeriodInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A StructuredPeriodInformationType"`

	// Effective date of period of operation
	BeginDateTime *StructuredDateTimeType `xml:"beginDateTime,omitempty"`

	// Discontinue date of period of operation
	EndDateTime *StructuredDateTimeType `xml:"endDateTime,omitempty"`

	// It is used with a period to give a restriction for days impacted. It permits for example to indicate on which days, a flight operates.
	Frequency *FrequencyType `xml:"frequency,omitempty"`
}

type TaxDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TaxDetailsTypeI"`

	// Duty/tax/fee rate
	Rate formats.AlphaNumericString_Length1To18 `xml:"rate,omitempty"`

	// Country, coded
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`

	// Currency, coded
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Duty/Tax fee type, coded
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Amount type qualifier, coded
	AmountQualifier formats.AlphaNumericString_Length1To3 `xml:"amountQualifier,omitempty"`
}

type TaxType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TaxType"`

	// Duty tax fee category, coded
	WithholdTaxSurcharge formats.AlphaNumericString_Length1To3 `xml:"withholdTaxSurcharge,omitempty"`

	// TAX DETAILS
	TaxDetail *TaxDetailsTypeI `xml:"taxDetail,omitempty"`
}

type TicketNumberDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TicketNumberDetailsTypeI"`

	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`
}

type TicketNumberTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TicketNumberTypeI"`

	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails,omitempty"`
}

type TicketingPriceSchemeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TicketingPriceSchemeType"`

	// PSR (Price Scheme Reference): unique reference of the price scheme as a 8 digit number.
	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber,omitempty"`

	// Price Scheme Name
	Name formats.AlphaNumericString_Length1To35 `xml:"name,omitempty"`

	// Price Scheme Status. Is the price scheme valid for service fee calculation ?
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// free flow description of the price scheme
	Description formats.AlphaNumericString_Length1To250 `xml:"description,omitempty"`
}

type TrafficRestrictionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TrafficRestrictionDetailsTypeI"`

	// Traffic Restriction code
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`
}

type TrafficRestrictionTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TrafficRestrictionTypeI"`

	// Traffic Restriction Details
	TrafficRestrictionDetails *TrafficRestrictionDetailsTypeI `xml:"trafficRestrictionDetails,omitempty"`
}

type TravelFlightInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TravelFlightInformationType"`

	// Cabin identification
	CabinId *CabinIdentificationType `xml:"cabinId,omitempty"`

	// Company Identification
	CompanyIdentity *CompanyIdentificationType_120719C `xml:"companyIdentity,omitempty"`

	// Type of flight details
	FlightDetail *ProductTypeDetailsType_120801C `xml:"flightDetail,omitempty"`

	// Details of included  connecting points
	InclusionDetail *ConnectPointDetailsType_195492C `xml:"inclusionDetail,omitempty"`

	// Further connection details
	ExclusionDetail *ConnectPointDetailsType `xml:"exclusionDetail,omitempty"`

	// Nb of connections for each requested segment of the journey.
	UnitNumberDetail *NumberOfUnitDetailsTypeI `xml:"unitNumberDetail,omitempty"`
}

type TravelFlightInformationType_165053S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TravelFlightInformationType_165053S"`

	// Cabin identification
	CabinId *CabinIdentificationType_233500C `xml:"cabinId,omitempty"`

	// Company Identification
	CompanyIdentity *CompanyIdentificationType_120719C `xml:"companyIdentity,omitempty"`

	// Type of flight details
	FlightDetail *ProductTypeDetailsType_120801C `xml:"flightDetail,omitempty"`

	// Details of included connect point
	InclusionDetail *ConnectPointDetailsType_195492C `xml:"inclusionDetail,omitempty"`

	// Further connection details
	ExclusionDetail *ConnectPointDetailsType `xml:"exclusionDetail,omitempty"`

	// Nb of connections allowed at requested segment level.
	UnitNumberDetail *NumberOfUnitDetailsTypeI `xml:"unitNumberDetail,omitempty"`
}

type TravelFlightInformationType_185853S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TravelFlightInformationType_185853S"`

	// Cabin identification
	CabinId *CabinIdentificationType_233500C `xml:"cabinId,omitempty"`

	// Company Identification
	CompanyIdentity *CompanyIdentificationType_233548C `xml:"companyIdentity,omitempty"`

	// Type of flight details
	FlightDetail *ProductTypeDetailsType_120801C `xml:"flightDetail,omitempty"`

	// Details of included connect point
	InclusionDetail *ConnectPointDetailsType_195492C `xml:"inclusionDetail,omitempty"`

	// Further connection details
	ExclusionDetail *ConnectPointDetailsType `xml:"exclusionDetail,omitempty"`

	// Added departed flights flag
	UnitNumberDetail *NumberOfUnitDetailsTypeI `xml:"unitNumberDetail,omitempty"`
}

type TravelProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TravelProductInformationTypeI"`

	// Flight Date
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`

	// Board point
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	// Off point
	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`

	// Flight Carrier
	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`

	// Flight identification
	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`

	// Identify flight part of the journey
	FlightTypeDetails *ProductTypeDetailsTypeI `xml:"flightTypeDetails,omitempty"`
}

type TravelProductType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TravelProductType"`

	// Date and time of departure and arrival
	ProductDateTime *ProductDateTimeType_195546C `xml:"productDateTime,omitempty"`

	// Location of departure and arrival
	Location *LocationIdentificationDetailsType `xml:"location,omitempty"`

	// Company identification
	CompanyId *CompanyIdentificationType_195544C `xml:"companyId,omitempty"`

	// Flight number or trainNumber
	FlightOrtrainNumber formats.AlphaNumericString_Length1To8 `xml:"flightOrtrainNumber,omitempty"`

	// Product details
	ProductDetail *AdditionalProductDetailsType `xml:"productDetail,omitempty"`

	// Additional product details
	AddProductDetail *ProductFacilitiesType `xml:"addProductDetail,omitempty"`

	// Attribute details
	AttributeDetails *CodedAttributeInformationType_247828C `xml:"attributeDetails,omitempty"`
}

type TravellerDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TravellerDetailsType"`

	// Direct reference of passenger assigned by requesting system.
	Ref formats.NumericInteger_Length1To3 `xml:"ref,omitempty"`

	// Traveller is an infant
	InfantIndicator formats.NumericInteger_Length1To1 `xml:"infantIndicator,omitempty"`
}

type TravellerReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A TravellerReferenceInformationType"`

	// Requested passenger type
	Ptc formats.AlphaNumericString_Length1To6 `xml:"ptc,omitempty"`

	// Traveller details
	Traveller *TravellerDetailsType `xml:"traveller,omitempty"`
}

type UniqueIdDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A UniqueIdDescriptionType"`

	// Reference qualifier.
	PassengerFeeRefQualif formats.AlphaNumericString_Length1To3 `xml:"passengerFeeRefQualif,omitempty"`
}

// UserIdentificationType ...
type UserIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A UserIdentificationType"`

	// Originator Identification Details
	OfficeIdentification *OriginatorIdentificationDetailsTypeI `xml:"officeIdentification,omitempty"`

	// Used to specify which kind of info is given in DE 9900.
	OfficeType formats.AlphaNumericString_Length1To1 `xml:"officeType,omitempty"`

	// The code given to an agent by the originating reservation system.
	OfficeCode formats.AlphaNumericString_Length1To30 `xml:"officeCode,omitempty"`
}

type ValueSearchCriteriaType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_14_3_1A ValueSearchCriteriaType"`

	CriteriaName formats.AlphaNumericString_Length1To50 `xml:"criteriaName,omitempty"`

	CriteriaCode formats.AlphaNumericString_Length1To3 `xml:"criteriaCode,omitempty"`

	Value formats.AlphaNumericString_Length1To18 `xml:"value,omitempty"`

	CriteriaDetails *CriteriaiDetaislType `xml:"criteriaDetails,omitempty"`
}
