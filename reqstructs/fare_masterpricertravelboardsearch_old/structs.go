package fare_masterpricertravelboardsearch_old

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type FareMasterPricerTravelBoardSearchOld struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBQ_12_3_1A Fare_MasterPricerTravelBoardSearch"`

	// Number of seats , recommendations.
	NumberOfUnit *NumberOfUnitsType `xml:"numberOfUnit,omitempty"`  // minOccurs="0"

	// Global options
	GlobalOptions *AttributeType `xml:"globalOptions,omitempty"`  // minOccurs="0"

	// Traveler Details
	PaxReference []*TravellerReferenceInformationType `xml:"paxReference,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Customer references
	CustomerRef *ConsumerReferenceInformationType `xml:"customerRef,omitempty"`  // minOccurs="0"

	// Fee with different payment forms  by passenger.
	FormOfPaymentByPassenger []*FOPRepresentationType `xml:"formOfPaymentByPassenger,omitempty"`  // minOccurs="0" maxOccurs="60"

	// Solution Family
	SolutionFamily []*FareInformationType `xml:"solutionFamily,omitempty"`  // minOccurs="0" maxOccurs="20"

	FareFamilies []*FareFamilies `xml:"fareFamilies,omitempty"`  // minOccurs="0" maxOccurs="20"

	FareOptions *FareOptions `xml:"fareOptions,omitempty"`  // minOccurs="0"

	// Indicates Price to beat
	PriceToBeat *MonetaryInformationType `xml:"priceToBeat,omitempty"`  // minOccurs="0"

	// Tax Details
	TaxInfo []*TaxType `xml:"taxInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Details of a Flight : Direct, Non stop...
	TravelFlightInfo *TravelFlightInformationType_134790S `xml:"travelFlightInfo,omitempty"`  // minOccurs="0"

	Itinerary []*Itinerary `xml:"itinerary,omitempty"`  // minOccurs="0" maxOccurs="18"

	TicketChangeInfo *TicketChangeInfo `xml:"ticketChangeInfo,omitempty"`  // minOccurs="0"

	CombinationFareFamilies []*CombinationFareFamilies `xml:"combinationFareFamilies,omitempty"`  // minOccurs="0" maxOccurs="2000"

	FeeOption []*FeeOption `xml:"feeOption,omitempty"`  // minOccurs="0" maxOccurs="9"

	OfficeIdDetails []*OfficeIdDetails `xml:"officeIdDetails,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type FareFamilies struct {

	// Segment used to target Commercial Fare Family.
	FamilyInformation *FareFamilyType `xml:"familyInformation"`

	// Description of Fare Family Criteria.
	FamilyCriteria *FareFamilyCriteriaType `xml:"familyCriteria,omitempty"`  // minOccurs="0"

	FareFamilySegment []*FareFamilySegment `xml:"fareFamilySegment,omitempty"`  // minOccurs="0" maxOccurs="6"

	OtherPossibleCriteria []*OtherPossibleCriteria `xml:"otherPossibleCriteria,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type FareFamilySegment struct {

	// Requested segment reference
	ReferenceInfo *ReferenceInfoType `xml:"referenceInfo"`

	// Description of fare family criteria.
	FamilyCriteria *FareFamilyCriteriaType `xml:"familyCriteria,omitempty"`  // minOccurs="0"
}

type OtherPossibleCriteria struct {

	// Logical link with other criteria.
	LogicalLink *BooleanExpressionRuleType `xml:"logicalLink"`

	// Description of fare family criteria.
	FamilyCriteria *FareFamilyCriteriaType `xml:"familyCriteria,omitempty"`  // minOccurs="0"

	FareFamilySegment []*FareFamilySegment `xml:"fareFamilySegment,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type FareOptions struct {

	// Pricing and ticketing details.
	PricingTickInfo *PricingTicketingDetailsType `xml:"pricingTickInfo"`

	// Corporate name/number used to target fares
	Corporate *CorporateIdentificationType `xml:"corporate,omitempty"`  // minOccurs="0"

	// Ticketing price scheme.
	TicketingPriceScheme *TicketingPriceSchemeType `xml:"ticketingPriceScheme,omitempty"`  // minOccurs="0"

	// PSR number
	FeeIdDescription *CodedAttributeType_78503S `xml:"feeIdDescription,omitempty"`  // minOccurs="0"

	// Used to force the currency of pricing
	ConversionRate *ConversionRateType `xml:"conversionRate,omitempty"`  // minOccurs="0"

	// Form of payment information.
	FormOfPayment *FormOfPaymentTypeI `xml:"formOfPayment,omitempty"`  // minOccurs="0"

	// Frequent traveller information
	FrequentTravellerInfo *FrequentTravellerInformationTypeI `xml:"frequentTravellerInfo,omitempty"`  // minOccurs="0"

	// Monetary and cabin information.
	MonetaryCabinInfo *MonetaryAndCabinInformationType `xml:"monetaryCabinInfo,omitempty"`  // minOccurs="0"
}

type Itinerary struct {

	// Indicates reference of the requested segment
	RequestedSegmentRef *OriginAndDestinationRequestType `xml:"requestedSegmentRef"`

	// Specification of the requested departure point
	DepartureLocalization *DepartureLocationType `xml:"departureLocalization,omitempty"`  // minOccurs="0"

	// Specification of the requested arrival point
	ArrivalLocalization *ArrivalLocalizationType `xml:"arrivalLocalization,omitempty"`  // minOccurs="0"

	// Details on requested date and time plus range of date trip duration
	TimeDetails *DateAndTimeInformationType `xml:"timeDetails,omitempty"`  // minOccurs="0"

	// Specify Flight options.
	FlightInfo *TravelFlightInformationType_134790S `xml:"flightInfo,omitempty"`  // minOccurs="0"

	FlightInfoPNR []*FlightInfoPNR `xml:"flightInfoPNR,omitempty"`  // minOccurs="0" maxOccurs="4"

	// Action identification for the requested segment
	RequestedSegmentAction *ActionIdentificationType `xml:"requestedSegmentAction,omitempty"`  // minOccurs="0"

	// Coded attributes
	Attributes *CodedAttributeType_82309S `xml:"attributes,omitempty"`  // minOccurs="0"
}

type FlightInfoPNR struct {

	// Travel Response Details: - Board/Off aiports - Flight number - Part of the journey - Day difference between Board and off  - Departure/Arrival dates - Departure/Arrival times
	TravelResponseDetails *TravelProductInformationTypeI `xml:"travelResponseDetails"`

	// Time Table Effective/Discontinue dates and frequency of operating Travel Solution
	TimeTableDate *StructuredPeriodInformationType `xml:"timeTableDate,omitempty"`  // minOccurs="0"

	// Terminal,Equipment and EFT Details. If a Total EFT of the Travel Solution is requested (TEF in SFLIGQ SDT) then we will have 2 occurences of the APD attached to the first leg TVL: the 1st occurence will give the EFT of the leg, the 2nd occurence will give the total EFT of the Travel Solution.
	TerminalEquipmentDetails []*AdditionalProductDetailsTypeI `xml:"terminalEquipmentDetails,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Codeshare data
	CodeshareData *CommercialAgreementsType `xml:"codeshareData,omitempty"`  // minOccurs="0"

	// Disclosure message from an operating carrier.
	Disclosure *FreeTextInformationType `xml:"disclosure,omitempty"`  // minOccurs="0"

	// Stops Details
	StopDetails *RoutingInformationTypeI `xml:"stopDetails,omitempty"`  // minOccurs="0"

	// Traffic restriction information
	TrafficRestrictionData *TrafficRestrictionTypeI `xml:"trafficRestrictionData,omitempty"`  // minOccurs="0"

	// PNR flight reservation info
	ReservationInfo *PassengerItineraryInformationType `xml:"reservationInfo,omitempty"`  // minOccurs="0"

	IncidentalStopInfo []*IncidentalStopInfo `xml:"incidentalStopInfo,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type IncidentalStopInfo struct {

	// Incidental stop date/time information
	DateTimeInfo *DateAndTimeInformationTypeI `xml:"dateTimeInfo"`
}

type TicketChangeInfo struct {

	// Ticket nb details.
	TicketNumberDetails *TicketNumberTypeI `xml:"ticketNumberDetails"`

	TicketRequestedSegments []*TicketRequestedSegments `xml:"ticketRequestedSegments,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type TicketRequestedSegments struct {

	// Action identification.
	ActionIdentification *ActionIdentificationType `xml:"actionIdentification"`

	// Connected cities in changed ticket requested segment.
	ConnectPointDetails *ConnectionTypeI `xml:"connectPointDetails,omitempty"`  // minOccurs="0"
}

type CombinationFareFamilies struct {

	// Specification of the item number
	ItemFFCNumber *ItemNumberType `xml:"itemFFCNumber"`

	// Number of units.
	NbOfUnits *NumberOfUnitsType_80154S `xml:"nbOfUnits,omitempty"`  // minOccurs="0"

	// Requested segment reference
	ReferenceInfo []*ReferenceInfoType `xml:"referenceInfo,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type FeeOption struct {

	// Nature of the fee (OB,OC,..)
	FeeTypeInfo *SelectionDetailsType `xml:"feeTypeInfo"`

	// Associated rate tax.
	RateTax *MonetaryInformationType_80162S `xml:"rateTax,omitempty"`  // minOccurs="0"

	FeeDetails []*FeeDetails `xml:"feeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FeeDetails struct {

	// Fee information
	FeeInfo *SpecificDataInformationType `xml:"feeInfo"`

	// Associated amounts : amounts to take into account to calculate fee.
	AssociatedAmounts *MonetaryInformationTypeI `xml:"associatedAmounts,omitempty"`  // minOccurs="0"

	FeeDescriptionGrp *FeeDescriptionGrp `xml:"feeDescriptionGrp,omitempty"`  // minOccurs="0"
}

type FeeDescriptionGrp struct {

	// Specification of the item number
	ItemNumberInfo *ItemNumberType_80866S `xml:"itemNumberInfo"`

	// Attributes  (SSR code EMD, RFIC, SSIM)
	ServiceAttributesInfo *AttributeType_61377S `xml:"serviceAttributesInfo,omitempty"`  // minOccurs="0"

	// Other service information (service description, ...)
	ServiceDescriptionInfo *SpecialRequirementsDetailsType `xml:"serviceDescriptionInfo,omitempty"`  // minOccurs="0"
}

type OfficeIdDetails struct {

	// Office Id Information
	OfficeIdInformation *UserIdentificationType `xml:"officeIdInformation"`

	// Number of units.
	NbOfUnits *NumberOfUnitsType_80154S `xml:"nbOfUnits,omitempty"`  // minOccurs="0"

	// UID option
	UidOption *CodedAttributeType `xml:"uidOption,omitempty"`  // minOccurs="0"

	// Pricing and ticketing details.
	PricingTickInfo *PricingTicketingDetailsType `xml:"pricingTickInfo,omitempty"`  // minOccurs="0"

	// Corporate fare information
	CorporateFareInfo *CorporateFareInformationType `xml:"corporateFareInfo,omitempty"`  // minOccurs="0"

	// Details of a Flight : Direct, Non stop...
	TravelFlightInfo *TravelFlightInformationType `xml:"travelFlightInfo,omitempty"`  // minOccurs="0"

	AirlineDistributionDetails []*AirlineDistributionDetails `xml:"airlineDistributionDetails,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type AirlineDistributionDetails struct {

	// Indicates reference of the requested segment
	RequestedSegmentRef *OriginAndDestinationRequestType `xml:"requestedSegmentRef"`

	// Specify Flight options.
	FlightInfo *TravelFlightInformationType `xml:"flightInfo,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type ActionIdentificationType struct {

	// Action request code
	ActionRequestCode formats.AlphaNumericString_Length1To3 `xml:"actionRequestCode"`

	ProductDetails *ProductIdentificationDetailsTypeI_50878C `xml:"productDetails,omitempty"`  // minOccurs="0"
}

type AdditionalProductDetailsTypeI struct {

	// Flight details
	LegDetails *AdditionalProductTypeI `xml:"legDetails,omitempty"`  // minOccurs="0"

	// Departure station
	DepartureStationInfo *StationInformationTypeI `xml:"departureStationInfo,omitempty"`  // minOccurs="0"

	ArrivalStationInfo *StationInformationTypeI `xml:"arrivalStationInfo,omitempty"`  // minOccurs="0"

	// Ground Time Details when connection
	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`  // minOccurs="0"
}

type AdditionalProductTypeI struct {

	// Equipment type
	Equipment formats.AlphaNumericString_Length1To3 `xml:"equipment,omitempty"`  // minOccurs="0"

	// Elaps flying time of the leg
	Duration *formats.NumericInteger_Length1To6 `xml:"duration,omitempty"`  // minOccurs="0"

	// COG indicator
	ComplexingFlightIndicator formats.AlphaNumericString_Length1To1 `xml:"complexingFlightIndicator,omitempty"`  // minOccurs="0"
}

type AgentIdentificationType struct {

	// Contains ARC number
	ArcNumber formats.AlphaNumericString_Length1To12 `xml:"arcNumber,omitempty"`  // minOccurs="0"

	// ERSP number of the Office or Customer
	ErspNumber formats.AlphaNumericString_Length1To12 `xml:"erspNumber,omitempty"`  // minOccurs="0"

	// IATA Number
	IataNumber formats.AlphaNumericString_Length1To12 `xml:"iataNumber,omitempty"`  // minOccurs="0"
}

type ArithmeticEvaluationType struct {

	// Boolean operator
	CodeOperator formats.AlphaNumericString_Length1To3 `xml:"codeOperator,omitempty"`  // minOccurs="0"
}

type ArrivalLocalizationType struct {

	// Details on the location of the arrival point
	ArrivalPointDetails *ArrivalLocationDetailsType `xml:"arrivalPointDetails,omitempty"`  // minOccurs="0"

	// Arrival multi city option
	ArrivalMultiCity []*MultiCityOptionType `xml:"arrivalMultiCity,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Attribute details
	AttributeDetails []*CodedAttributeInformationType_139508C `xml:"attributeDetails,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type ArrivalLocationDetailsType struct {

	// For Radius: This is the length of the requested radius around the location of destination.
	Distance *formats.NumericInteger_Length1To3 `xml:"distance,omitempty"`  // minOccurs="0"

	// Distance unit qualifier for radius
	DistanceUnit formats.AlphaNumericString_Length0To3 `xml:"distanceUnit,omitempty"`  // minOccurs="0"

	// ATA/IATA airport/city code of arrival  In case of SP request from a PNR, this field is empty.
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`  // minOccurs="0"

	// Request from no PNR A = consider the locationID specified as an airport (used only when ambiguous) C = consider the locationId as a city  (used only when ambiguous) A and C are not used in case of a request from PNR  Request from a PNR with Radius option: O = radius applies to the origin of the PNR segments D = radius applies to the destination of the PNR segments
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`  // minOccurs="0"

	// Latitude in degrees (decimal format) with hemisphere N=north S=south
	Latitude formats.AlphaNumericString_Length6To6 `xml:"latitude,omitempty"`  // minOccurs="0"

	// Longitude in degrees (decimal format) with hemisphere E=east, W=west
	Longitude formats.AlphaNumericString_Length6To6 `xml:"longitude,omitempty"`  // minOccurs="0"
}

type ArrivalLocationDetailsType_120834C struct {

	// For Radius: This is the length of the requested radius around the location of origin.
	Distance *formats.NumericInteger_Length1To3 `xml:"distance,omitempty"`  // minOccurs="0"

	// Distance unit qualifier for radius
	DistanceUnit formats.AlphaNumericString_Length1To3 `xml:"distanceUnit,omitempty"`  // minOccurs="0"

	// ATA/IATA airport/city code of arrival  In case of SP request from a PNR, this field is empty.
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`  // minOccurs="0"

	// Request from no PNR A = consider the locationID specified as an airport (used only when ambiguous) C = consider the locationId as a city  (used only when ambiguous) A and C are not used in case of a request from PNR  Request from a PNR with Radius option: O = radius applies to the origin of the PNR segments D = radius applies to the destination of the PNR segments
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`  // minOccurs="0"

	// Latitude in degrees (decimal format) with hemisphere N=north S=south
	Latitude formats.AlphaNumericString_Length6To6 `xml:"latitude,omitempty"`  // minOccurs="0"

	// Longitude in degrees (decimal format) with hemisphere E=east, W=west
	Longitude formats.AlphaNumericString_Length6To6 `xml:"longitude,omitempty"`  // minOccurs="0"
}

type AttributeInformationType struct {

	// Code of options
	Option formats.AlphaNumericString_Length1To3 `xml:"option"`

	// Mandatory when alternate date option is used (ALT), must be set to plus (P) or minus (M) a number of days around the original PNR segment dates. E.g.: M1 (for minus 1 day) or P100 (for plus 100 days)
	OptionInformation formats.AlphaNumericString_Length1To35 `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type AttributeInformationType_97181C struct {

	// Attribute type
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type AttributeType struct {

	// Option parameters
	SelectionDetails []*AttributeInformationType `xml:"selectionDetails"`  // maxOccurs="10"
}

type AttributeType_61377S struct {

	// Criteria Set Type
	AttributeQualifier formats.AlphaNumericString_Length1To3 `xml:"attributeQualifier,omitempty"`  // minOccurs="0"

	// Criteria details
	AttributeDetails []*AttributeInformationType_97181C `xml:"attributeDetails"`  // maxOccurs="99"
}

type BooleanExpressionRuleType struct {

	// Boolean expression associated to the decision rule.
	BooleanExpression *ArithmeticEvaluationType `xml:"booleanExpression"`
}

type CabinClassDesignationType struct {

	// Cabin designator.
	CabinDesignator formats.AlphaNumericString_Length1To1 `xml:"cabinDesignator"`
}

type CabinIdentificationType struct {

	// Cabin qualifier. For Star Pricer,MD stands for Mandatory Cabin qualifier.  For other products no qualifier stands for Mandatory Cabin.
	CabinQualifier formats.AlphaNumericString_Length1To2 `xml:"cabinQualifier,omitempty"`  // minOccurs="0"

	// Cabin
	Cabin []formats.AlphaString_Length0To1 `xml:"cabin"`  // maxOccurs="3"
}

type CodedAttributeInformationType struct {

	// Attribute type identification
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType"`

	// Attribute Description
	AttributeDescription formats.AlphaNumericString_Length1To20 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeInformationType_120700C struct {

	// Type of fee/reduction
	FeeType formats.AlphaNumericString_Length1To5 `xml:"feeType"`

	// Fee Id Number
	FeeIdNumber formats.AlphaNumericString_Length1To50 `xml:"feeIdNumber"`
}

type CodedAttributeInformationType_120742C struct {

	// Name.
	Name formats.AlphaNumericString_Length1To5 `xml:"name"`

	// Value.
	Value []formats.AlphaNumericString_Length1To20 `xml:"value,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type CodedAttributeInformationType_125859C struct {

	// Attribute type
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType"`

	// Attribute description
	AttributeDescription formats.AlphaNumericString_Length1To50 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeInformationType_139508C struct {

	// Attribute type
	Type formats.AlphaNumericString_Length1To5 `xml:"type,omitempty"`  // minOccurs="0"

	// Value.
	Value []formats.AlphaNumericString_Length1To20 `xml:"value,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type CodedAttributeType struct {

	// Atrribute details.
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"`  // maxOccurs="20"
}

type CodedAttributeType_78503S struct {

	// Fee/reduction Id
	FeeId []*CodedAttributeInformationType_120700C `xml:"feeId,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type CodedAttributeType_82309S struct {

	// Attribute details
	AttributeDetails []*CodedAttributeInformationType_125859C `xml:"attributeDetails"`  // maxOccurs="9"
}

type CommercialAgreementsType struct {

	// Codeshare Details
	CodeshareDetails *CompanyRoleIdentificationType `xml:"codeshareDetails"`

	// Other codeshare details
	OtherCodeshareDetails []*CompanyRoleIdentificationType `xml:"otherCodeshareDetails,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type CompanyIdentificationType struct {}

type CompanyIdentificationTypeI struct {

	// Carrier code
	MarketingCompany formats.AlphaNumericString_Length2To3 `xml:"marketingCompany"`
}

type CompanyIdentificationType_120719C struct {

	// Item description identification.
	CarrierQualifier formats.AlphaNumericString_Length0To1 `xml:"carrierQualifier"`

	// carrier id
	CarrierId []formats.AlphaNumericString_Length2To3 `xml:"carrierId"`  // maxOccurs="99"
}

type CompanyRoleIdentificationType struct {

	// Codeshare qualifier
	TransportStageQualifier formats.AlphaString_Length1To1 `xml:"transportStageQualifier"`

	// company identification
	AirlineDesignator formats.AlphaNumericString_Length2To3 `xml:"airlineDesignator,omitempty"`  // minOccurs="0"

	// flight number
	FlightNumber *formats.NumericInteger_Length1To4 `xml:"flightNumber,omitempty"`  // minOccurs="0"

	// suffix
	OperationalSuffix formats.AlphaString_Length1To1 `xml:"operationalSuffix,omitempty"`  // minOccurs="0"
}

type ConnectPointDetailsType struct {

	// Exclusion identification
	ExclusionIdentifier formats.AlphaString_Length0To1 `xml:"exclusionIdentifier"`

	// Place or Location identification
	LocationId formats.AlphaString_Length3To5 `xml:"locationId"`

	// Airport/city qualifier
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`  // minOccurs="0"
}

type ConnectPointDetailsType_195492C struct {

	// Inclusion identification
	InclusionIdentifier formats.AlphaNumericString_Length0To1 `xml:"inclusionIdentifier"`

	// Place or Location identification
	LocationId formats.AlphaString_Length3To5 `xml:"locationId"`

	// Airport/city qualifier
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`  // minOccurs="0"
}

type ConnectionDetailsTypeI struct {

	// Location
	Location formats.AlphaNumericString_Length1To3 `xml:"location"`
}

type ConnectionTypeI struct {

	// Connection details.
	ConnectionDetails []*ConnectionDetailsTypeI `xml:"connectionDetails"`  // maxOccurs="17"
}

type ConsumerReferenceIdentificationTypeI struct {

	ReferenceQualifier formats.AlphaNumericString_Length1To3 `xml:"referenceQualifier"`

	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber,omitempty"`  // minOccurs="0"

	ReferencePartyName formats.AlphaNumericString_Length1To35 `xml:"referencePartyName,omitempty"`  // minOccurs="0"

	TravellerReferenceNbr formats.AlphaNumericString_Length1To10 `xml:"travellerReferenceNbr,omitempty"`  // minOccurs="0"
}

type ConsumerReferenceInformationType struct {

	// Customer references
	CustomerReferences []*ConsumerReferenceIdentificationTypeI `xml:"customerReferences"`  // maxOccurs="20"
}

type ConversionRateDetailsType struct {

	// Conversion type
	ConversionType formats.AlphaNumericString_Length1To3 `xml:"conversionType,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaString_Length1To3 `xml:"currency"`
}

type ConversionRateType struct {

	// Detail of conversion rate of First Monetary Unit
	ConversionRateDetail []*ConversionRateDetailsType `xml:"conversionRateDetail"`  // maxOccurs="2"
}

type CorporateFareIdentifiersType struct {

	// Indicates the type of corporate fares requested
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier,omitempty"`  // minOccurs="0"

	// Corporate contract number or name
	IdentifyNumber []formats.AlphaNumericString_Length1To35 `xml:"identifyNumber,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type CorporateFareInformationType struct {

	// Corporate fare identifiers
	CorporateFareIdentifiers *CorporateFareIdentifiersType `xml:"corporateFareIdentifiers,omitempty"`  // minOccurs="0"
}

type CorporateIdentificationType struct {

	// Corporate identity
	CorporateId []*CorporateIdentityType `xml:"corporateId,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type CorporateIdentityType struct {

	// Indicates if 1A corporate (RC) or Unifare corporate (RW) requested.
	CorporateQualifier formats.AlphaNumericString_Length0To3 `xml:"corporateQualifier"`

	// Corporate contract number or name
	Identity []formats.AlphaNumericString_Length1To20 `xml:"identity"`  // maxOccurs="9"
}

type DataInformationType struct {

	// Ancillary services options
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`  // minOccurs="0"
}

type DataTypeInformationType struct {

	// Carrier fee code
	SubType formats.AlphaNumericString_Length1To3 `xml:"subType"`

	// Status (automated, manually added, exempted). Default is automated
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsType struct {

	// Trip Duration type(Plus,Minus,Combined)
	FlexibilityQualifier formats.AlphaNumericString_Length1To3 `xml:"flexibilityQualifier,omitempty"`  // minOccurs="0"

	// Number of days added or/and retrieved to the trip duration
	TripInterval *formats.NumericInteger_Length1To6 `xml:"tripInterval,omitempty"`  // minOccurs="0"

	// Period between date of departure and date of arrival
	TripDuration *formats.NumericInteger_Length1To4 `xml:"tripDuration,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsTypeI struct {

	// Toidentify type of time Arrival/Departure
	TimeQualifier formats.AlphaNumericString_Length1To3 `xml:"timeQualifier,omitempty"`  // minOccurs="0"

	// Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`  // minOccurs="0"

	// Time
	Time formats.Time24_HHMM `xml:"time,omitempty"`  // minOccurs="0"

	// Time window size in hours
	TimeWindow formats.AlphaNumericString_Length1To3 `xml:"timeWindow,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsTypeI_120740C struct {

	// Date time period qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// First Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`  // minOccurs="0"

	// First Time
	Time formats.Time24_HHMM `xml:"time,omitempty"`  // minOccurs="0"

	// .
	Qualifier2 formats.AlphaNumericString_Length1To3 `xml:"qualifier2,omitempty"`  // minOccurs="0"

	Reserved1 formats.AlphaNumericString_Length1To3 `xml:"reserved1,omitempty"`  // minOccurs="0"

	Reserved2 formats.AlphaNumericString_Length3To5 `xml:"reserved2,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsType_120843C struct {

	// To identify type of range (Plus,Minus or combined)
	RangeQualifier formats.AlphaNumericString_Length1To3 `xml:"rangeQualifier,omitempty"`  // minOccurs="0"

	// Range of dates : Number of Days preceding or/and Following the request departure date
	DayInterval *formats.NumericInteger_Length1To6 `xml:"dayInterval,omitempty"`  // minOccurs="0"

	// Time at destination (local)
	TimeAtdestination *formats.NumericInteger_Length4To4 `xml:"timeAtdestination,omitempty"`  // minOccurs="0"
}

type DateAndTimeInformationType struct {

	// Details on date and Time
	FirstDateTimeDetail *DateAndTimeDetailsTypeI `xml:"firstDateTimeDetail"`

	// Details of the Requested Range of Dates
	RangeOfDate *DateAndTimeDetailsType_120843C `xml:"rangeOfDate,omitempty"`  // minOccurs="0"

	// Details of the trip duration
	TripDetails *DateAndTimeDetailsType `xml:"tripDetails,omitempty"`  // minOccurs="0"
}

type DateAndTimeInformationTypeI struct {

	// DATE AND TIME DETAILS.
	DateTimeDetails []*DateAndTimeDetailsTypeI_120740C `xml:"dateTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type DepartureLocationType struct {

	// Details on localization of the departure point
	DeparturePoint *ArrivalLocationDetailsType_120834C `xml:"departurePoint,omitempty"`  // minOccurs="0"

	// Departure multi city option
	DepMultiCity []*MultiCityOptionType `xml:"depMultiCity,omitempty"`  // minOccurs="0" maxOccurs="20"

	// To specify a series or a range of PNR segments
	FirstPnrSegmentRef *PNRSegmentReferenceType `xml:"firstPnrSegmentRef,omitempty"`  // minOccurs="0"

	// Attribute details
	AttributeDetails []*CodedAttributeInformationType_139508C `xml:"attributeDetails,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type FOPRepresentationType struct {

	// Form of payment information.
	FormOfPaymentDetails *FormOfPaymentTypeI `xml:"formOfPaymentDetails"`

	// Fee/Reduction Passenger reference.
	PassengerFeeReference *ItemReferencesAndVersionsType `xml:"passengerFeeReference,omitempty"`  // minOccurs="0"
}

type FareDetailsType struct {

	// Qualifier
	Qualifier formats.AMA_EDICodesetType_Length1to3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Rate
	Rate *formats.NumericInteger_Length1To8 `xml:"rate,omitempty"`  // minOccurs="0"

	// Country
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`  // minOccurs="0"

	// Fare Category
	FareCategory formats.AMA_EDICodesetType_Length1to3 `xml:"fareCategory,omitempty"`  // minOccurs="0"
}

type FareFamilyCriteriaType struct {

	// Fare publishing carrier.
	CarrierId []formats.AlphaNumericString_Length1To3 `xml:"carrierId,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Reservation booking designator.
	Rdb []formats.AlphaString_Length1To2 `xml:"rdb,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Fare family info.
	FareFamilyInfo *FareQualifierInformationType `xml:"fareFamilyInfo,omitempty"`  // minOccurs="0"

	// Fare product detail.
	FareProductDetail []*FareProductDetailsType `xml:"fareProductDetail,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Corporate information.
	CorporateInfo []*MultipleIdentificationNumbersTypeI `xml:"corporateInfo,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Indicates flight cabin details.
	CabinProduct []*CabinClassDesignationType `xml:"cabinProduct,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Cabin processing option.
	CabinProcessingIdentifier formats.AlphaNumericString_Length1To3 `xml:"cabinProcessingIdentifier,omitempty"`  // minOccurs="0"

	// Product date or time.
	DateTimeDetails []*ProductDateTimeTypeI_194583C `xml:"dateTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Other criteria.
	OtherCriteria []*CodedAttributeInformationType_120742C `xml:"otherCriteria,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type FareFamilyDetailsType struct {

	// Commercial fare Family Short name
	CommercialFamily formats.AlphaNumericString_Length1To10 `xml:"commercialFamily"`
}

type FareFamilyType struct {

	// Fare Family Reference Number
	RefNumber *formats.NumericInteger_Length1To3 `xml:"refNumber,omitempty"`  // minOccurs="0"

	// Fare Family Short Name
	FareFamilyname formats.AlphaNumericString_Length1To10 `xml:"fareFamilyname,omitempty"`  // minOccurs="0"

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy *formats.NumericInteger_Length1To4 `xml:"hierarchy,omitempty"`  // minOccurs="0"

	// Indicates Commercial Fare Family Short names
	CommercialFamilyDetails []*FareFamilyDetailsType `xml:"commercialFamilyDetails,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type FareInformationType struct {

	// Value Qualifier
	ValueQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"valueQualifier,omitempty"`  // minOccurs="0"

	// Value
	Value *formats.NumericInteger_Length1To15 `xml:"value,omitempty"`  // minOccurs="0"

	// Fare Details
	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`  // minOccurs="0"

	// Identity Number
	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber,omitempty"`  // minOccurs="0"

	// Fare Type Grouping
	FareTypeGrouping *FareTypeGroupingInformationType `xml:"fareTypeGrouping,omitempty"`  // minOccurs="0"

	// Rate Category
	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"`  // minOccurs="0"
}

type FareProductDetailsType struct {

	// Fare basis code
	FareBasis formats.AlphaNumericString_Length0To18 `xml:"fareBasis,omitempty"`  // minOccurs="0"

	// Type of fare
	FareType []formats.AlphaNumericString_Length0To3 `xml:"fareType,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type FareQualifierInformationType struct {

	// Fare family combinability.
	FareFamilyQual []formats.AlphaNumericString_Length0To3 `xml:"fareFamilyQual"`  // maxOccurs="9"
}

type FareTypeGroupingInformationType struct {

	// Pricing Group
	PricingGroup formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"`  // minOccurs="0"
}

type FormOfPaymentDetailsTypeI struct {

	// Form of payment identification
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// amount to be charged on this form
	ChargedAmount *formats.NumericDecimal_Length1To12 `xml:"chargedAmount,omitempty"`  // minOccurs="0"

	// Reference number
	CreditCardNumber formats.AlphaNumericString_Length1To20 `xml:"creditCardNumber,omitempty"`  // minOccurs="0"
}

type FormOfPaymentTypeI struct {

	// FORM OF PAYMENT DETAILS
	FormOfPaymentDetails []*FormOfPaymentDetailsTypeI `xml:"formOfPaymentDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FreeTextDetailsType struct {

	// Qualifier of the free text. Always literal in our usage.
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier"`

	// Type of the free text. Always 7 for our usage.
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`  // minOccurs="0"

	// Source of the information.
	Source formats.AlphaNumericString_Length1To3 `xml:"source"`

	// Encoding method used.
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding"`
}

type FreeTextInformationType struct {

	// Details for the free text message
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text corresponding to the DEI 127 data.
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText"`
}

type FrequencyType struct {

	// Indicate if the sequence number represents days of the week or days of the month.
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier"`

	// Used to represent days of the week. 1 is monday and 7 is sunday.
	Value []*formats.NumericInteger_Length1To1 `xml:"value,omitempty"`  // minOccurs="0" maxOccurs="7"
}

type FrequentTravellerIdentificationTypeI struct {

	// carrier
	Carrier formats.AlphaNumericString_Length1To3 `xml:"carrier"`

	// Frequent traveller number
	Number formats.AlphaNumericString_Length1To25 `xml:"number,omitempty"`  // minOccurs="0"

	// Customer reference
	CustomerReference formats.AlphaNumericString_Length1To10 `xml:"customerReference,omitempty"`  // minOccurs="0"

	// Tier level
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`  // minOccurs="0"

	// Priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"`  // minOccurs="0"

	// Tier description
	TierDescription formats.AlphaNumericString_Length1To35 `xml:"tierDescription,omitempty"`  // minOccurs="0"
}

type FrequentTravellerInformationTypeI struct {

	// Frequent traveller details
	FrequentTravellerDetails []*FrequentTravellerIdentificationTypeI `xml:"frequentTravellerDetails"`  // maxOccurs="9"
}

type ItemNumberIdentificationType struct {

	// Fare family combination number
	Number formats.AlphaNumericString_Length1To4 `xml:"number,omitempty"`  // minOccurs="0"

	// Type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	// Qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Responsible agency
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`  // minOccurs="0"
}

type ItemNumberType struct {

	// Indicates the fare family combination number
	ItemNumberId *ItemNumberIdentificationType `xml:"itemNumberId"`
}

type ItemNumberType_80866S struct {

	// Item number details
	ItemNumberDetails *ItemNumberIdentificationType `xml:"itemNumberDetails"`
}

type ItemReferencesAndVersionsType struct {

	// Reference Qualifier.
	PassengerFeeRefType formats.AlphaNumericString_Length1To3 `xml:"passengerFeeRefType,omitempty"`  // minOccurs="0"

	// Reference number.
	PassengerFeeRefNumber *formats.NumericInteger_Length1To3 `xml:"passengerFeeRefNumber,omitempty"`  // minOccurs="0"

	// Unique id description.
	OtherCharacteristics *UniqueIdDescriptionType `xml:"otherCharacteristics,omitempty"`  // minOccurs="0"
}

type ItineraryDetailsType struct {

	// Airport/City Qualifier: the passenger wants to depart/arrive from/to the same airport or city in the specified requested segment
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier"`

	// Requested segment number
	SegmentNumber formats.NumericInteger_Length1To3 `xml:"segmentNumber"`
}

type LocationDetailsTypeI struct {

	// Place or Location identification
	LocationId formats.AlphaString_Length3To5 `xml:"locationId"`

	// Country, coded
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`  // minOccurs="0"
}

type LocationTypeI struct {

	// Departure or Arrival IATA airport code
	TrueLocationId formats.AlphaString_Length3To3 `xml:"trueLocationId"`
}

type MileageTimeDetailsTypeI struct {

	// Ground Time in minutes at Board point (connection with incoming flight)
	ElapsedGroundTime *formats.NumericInteger_Length1To4 `xml:"elapsedGroundTime,omitempty"`  // minOccurs="0"
}

type MonetaryAndCabinInformationDetailsType struct {

	// Amount qualifier.
	AmountType formats.AlphaNumericString_Length0To3 `xml:"amountType,omitempty"`  // minOccurs="0"

	// Amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount"`

	// ISO currency code
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// Airport/city code
	LocationId formats.AlphaString_Length3To5 `xml:"locationId,omitempty"`  // minOccurs="0"

	// Cabin class designator
	CabinClassDesignator []formats.AlphaString_Length1To1 `xml:"cabinClassDesignator,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type MonetaryAndCabinInformationType struct {

	// Monetary and cabin information.
	MoneyAndCabinInfo []*MonetaryAndCabinInformationDetailsType `xml:"moneyAndCabinInfo,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type MonetaryInformationDetailsTypeI struct {

	// Monetary amount type qualifier, coded
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Allowance or charge number
	Amount formats.NumericInteger_Length1To18 `xml:"amount"`

	// Currency, coded
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_194597C struct {

	// Monetary amount type qualifier, coded
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Allowance or charge number
	Amount formats.NumericInteger_Length1To18 `xml:"amount"`

	// Currency, coded
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// Place or Location identification
	LocationId formats.AlphaString_Length3To3 `xml:"locationId,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_65140C struct {

	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_65141C struct {

	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// Location
	Location formats.AlphaNumericString_Length1To3 `xml:"location,omitempty"`  // minOccurs="0"
}

type MonetaryInformationType struct {

	// MONEY INFORMATION
	MoneyInfo *MonetaryInformationDetailsTypeI `xml:"moneyInfo"`

	// MONEY INFORMATION
	AdditionalMoneyInfo []*MonetaryInformationDetailsTypeI_194597C `xml:"additionalMoneyInfo,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type MonetaryInformationTypeI struct {

	// Monetary info
	MonetaryDetails []*MonetaryInformationDetailsTypeI_65141C `xml:"monetaryDetails"`  // maxOccurs="20"
}

type MonetaryInformationType_80162S struct {

	// Monetary info
	MonetaryDetails []*MonetaryInformationDetailsTypeI_65140C `xml:"monetaryDetails"`  // maxOccurs="20"
}

type MultiCityOptionType struct {

	// ATA/IATA airport/city code of arrival multi city option enable to define until 20 airports/cities
	LocationId formats.AlphaString_Length3To5 `xml:"locationId"`

	// Requested arrival point is an airport or a city (default is city and used only when ambiguity)
	AirportCityQualifier formats.AlphaString_Length1To1 `xml:"airportCityQualifier,omitempty"`  // minOccurs="0"
}

type MultipleIdentificationNumbersTypeI struct {

	// Corporate number or ALL.
	CorporateNumberIdentifier formats.AlphaNumericString_Length1To12 `xml:"corporateNumberIdentifier,omitempty"`  // minOccurs="0"

	// Corporate name.
	CorporateName formats.AlphaNumericString_Length1To20 `xml:"corporateName,omitempty"`  // minOccurs="0"
}

type NumberOfUnitDetailsType struct {

	// Number of Units
	NumberOfUnits *formats.NumericInteger_Length1To4 `xml:"numberOfUnits,omitempty"`  // minOccurs="0"

	// Number of unit qualifier
	TypeOfUnit formats.AlphaNumericString_Length1To3 `xml:"typeOfUnit"`
}

type NumberOfUnitDetailsTypeI struct {

	// Number of Units
	NumberOfUnits formats.NumericInteger_Length1To3 `xml:"numberOfUnits"`

	// Number of unit qualifier
	TypeOfUnit formats.AlphaNumericString_Length1To3 `xml:"typeOfUnit"`
}

type NumberOfUnitDetailsType_191580C struct {

	// Number of Units
	NumberOfUnits formats.NumericInteger_Length1To6 `xml:"numberOfUnits"`

	// Number of unit qualifier
	TypeOfUnit formats.AlphaNumericString_Length1To3 `xml:"typeOfUnit"`
}

type NumberOfUnitsType struct {

	// NUMBER OF UNIT DETAILS
	UnitNumberDetail []*NumberOfUnitDetailsType_191580C `xml:"unitNumberDetail"`  // maxOccurs="20"
}

type NumberOfUnitsType_80154S struct {

	// NUMBER OF UNIT DETAILS
	UnitNumberDetail []*NumberOfUnitDetailsType `xml:"unitNumberDetail"`  // maxOccurs="20"
}

type OriginAndDestinationRequestType struct {

	// Requested segment number
	SegRef formats.NumericInteger_Length1To2 `xml:"segRef"`

	// Forces arrival or departure, from/to the same airport/city
	LocationForcing []*ItineraryDetailsType `xml:"locationForcing,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type OriginatorIdentificationDetailsTypeI struct {

	// Office Name.
	OfficeName *formats.NumericInteger_Length1To9 `xml:"officeName,omitempty"`  // minOccurs="0"

	// Agent Sign In .
	AgentSignin formats.AlphaNumericString_Length1To9 `xml:"agentSignin,omitempty"`  // minOccurs="0"

	// Confidential Office Name.
	ConfidentialOffice formats.AlphaNumericString_Length1To9 `xml:"confidentialOffice,omitempty"`  // minOccurs="0"

	// Other Office Name
	OtherOffice formats.AlphaNumericString_Length1To9 `xml:"otherOffice,omitempty"`  // minOccurs="0"
}

type PNRSegmentReferenceType struct {

	// For a request from PNR:  this is the reference number of a PNR air segment. In case a range of PNR segments is specified (eg. segments 2-5), then it is the 1st of the range, the last being in ARR.
	PnrSegmentTattoo *formats.NumericInteger_Length0To35 `xml:"pnrSegmentTattoo,omitempty"`  // minOccurs="0"

	PnrSegmentQualifier formats.AlphaString_Length1To1 `xml:"pnrSegmentQualifier,omitempty"`  // minOccurs="0"
}

type PassengerItineraryInformationType struct {

	// .
	Booking formats.AlphaString_Length1To1 `xml:"booking,omitempty"`  // minOccurs="0"

	// .
	Identifier formats.AlphaNumericString_Length1To1 `xml:"identifier,omitempty"`  // minOccurs="0"

	// .
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`  // minOccurs="0"

	// .
	ItemNumber *formats.NumericInteger_Length1To3 `xml:"itemNumber,omitempty"`  // minOccurs="0"

	// .
	DateTimeDetails *ProductDateTimeType `xml:"dateTimeDetails,omitempty"`  // minOccurs="0"

	// .
	Designator formats.AlphaString_Length1To1 `xml:"designator,omitempty"`  // minOccurs="0"

	// .
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`  // minOccurs="0"

	// .
	ProductTypeDetails *ProductTypeDetailsType `xml:"productTypeDetails,omitempty"`  // minOccurs="0"
}

type PricingTicketingDetailsType struct {

	// Pricing ticketing Details.
	PricingTicketing *PricingTicketingInformationType `xml:"pricingTicketing,omitempty"`  // minOccurs="0"

	// PRODUCT DATE OR TIME
	TicketingDate *ProductDateTimeTypeI_194598C `xml:"ticketingDate,omitempty"`  // minOccurs="0"

	// COMPANY IDENTIFICATION
	CompanyId *CompanyIdentificationType `xml:"companyId,omitempty"`  // minOccurs="0"

	// LOCATION DETAILS
	SellingPoint *LocationDetailsTypeI `xml:"sellingPoint,omitempty"`  // minOccurs="0"

	// LOCATION DETAILS
	TicketingPoint *LocationDetailsTypeI `xml:"ticketingPoint,omitempty"`  // minOccurs="0"

	// Used to Target Transborder Fares
	JourneyOriginPoint *LocationDetailsTypeI `xml:"journeyOriginPoint,omitempty"`  // minOccurs="0"

	// Contains the ARC,IATA and ERSP numbers
	CorporateId *AgentIdentificationType `xml:"corporateId,omitempty"`  // minOccurs="0"
}

type PricingTicketingInformationType struct {

	// Price type qualifier
	PriceType []formats.AlphaNumericString_Length0To3 `xml:"priceType"`  // maxOccurs="50"
}

type ProductDateTimeType struct {

	// .
	Date formats.Date_DDMMYY `xml:"date"`

	// .
	Time formats.Time24_HHMM `xml:"time,omitempty"`  // minOccurs="0"
}

type ProductDateTimeTypeI struct {

	// Departure date in YYYYMMDD format
	DepartureDate formats.Date_YYYYMMDD `xml:"departureDate,omitempty"`  // minOccurs="0"

	// Departure time
	DepartureTime formats.Time24_HHMM `xml:"departureTime,omitempty"`  // minOccurs="0"

	// Arrival date
	ArrivalDate formats.Date_YYYYMMDD `xml:"arrivalDate,omitempty"`  // minOccurs="0"

	// Arrival time
	ArrivalTime formats.Time24_HHMM `xml:"arrivalTime,omitempty"`  // minOccurs="0"

	// Day difference between Departure date of the leg and date of reference (Departure or Arrival date specified in the SDI)
	DateVariation *formats.NumericInteger_Length1To1 `xml:"dateVariation,omitempty"`  // minOccurs="0"
}

type ProductDateTimeTypeI_194583C struct {

	// Ticketing Purchase Date
	Date formats.Date_DDMMYY `xml:"date"`

	// Ticketing purchase date
	OtherDate formats.Date_DDMMYY `xml:"otherDate,omitempty"`  // minOccurs="0"
}

type ProductDateTimeTypeI_194598C struct {

	// First date
	Date formats.Date_DDMMYY `xml:"date"`

	// Half round trip combination.
	RtcDate formats.Date_DDMMYY `xml:"rtcDate,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetailsTypeI struct {

	// Flight number
	FlightNumber formats.NumericInteger_Length1To4 `xml:"flightNumber"`

	// Flight suffix
	OperationalSuffix formats.AlphaString_Length1To1 `xml:"operationalSuffix,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetailsTypeI_50878C struct {

	FlightNumber formats.AlphaNumericString_Length1To5 `xml:"flightNumber"`

	BookingClass formats.AlphaNumericString_Length1To2 `xml:"bookingClass,omitempty"`  // minOccurs="0"

	OperationalSuffix formats.AlphaNumericString_Length1To3 `xml:"operationalSuffix,omitempty"`  // minOccurs="0"

	Modifier []formats.AlphaNumericString_Length1To7 `xml:"modifier,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type ProductLocationDetailsTypeI struct {

	// airport
	Station formats.AlphaString_Length3To3 `xml:"station,omitempty"`  // minOccurs="0"
}

type ProductTypeDetailsType struct {

	// .
	SequenceNumber formats.AlphaNumericString_Length1To6 `xml:"sequenceNumber,omitempty"`  // minOccurs="0"

	// PNR availability context
	AvailabilityContext formats.AlphaNumericString_Length1To6 `xml:"availabilityContext,omitempty"`  // minOccurs="0"
}

type ProductTypeDetailsTypeI struct {

	// Part of the journey (C,E,S), Codeshare service (A), Technical stop at off point in a Direct (TSD), Technical stop at off point in a COG (TSC), E-Ticket candidate (ET), Prohibited Countries (RPC, WPC)
	FlightIndicator []formats.AlphaString_Length1To3 `xml:"flightIndicator"`  // maxOccurs="5"
}

type ProductTypeDetailsType_120801C struct {

	// Type of flight
	FlightType []formats.AlphaNumericString_Length1To2 `xml:"flightType,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type ReferenceInfoType struct {

	// Referencing details
	ReferencingDetail []*ReferencingDetailsType `xml:"referencingDetail,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type ReferencingDetailsType struct {

	// Segment reference qualifier
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Flight or flight group reference
	RefNumber formats.NumericInteger_Length0To3 `xml:"refNumber"`
}

type RoutingInformationTypeI struct {

	// Stops details
	RoutingDetails []*ProductLocationDetailsTypeI `xml:"routingDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type SelectionDetailsInformationType struct {

	// Carrier fee type
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// Carrier fee status
	OptionInformation formats.AlphaNumericString_Length1To3 `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type SelectionDetailsInformationTypeI struct {

	Option formats.AlphaNumericString_Length1To3 `xml:"option"`

	OptionInformation formats.AlphaNumericString_Length1To35 `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type SelectionDetailsType struct {

	// Carrier fees options
	CarrierFeeDetails *SelectionDetailsInformationType `xml:"carrierFeeDetails"`

	OtherSelectionDetails []*SelectionDetailsInformationTypeI `xml:"otherSelectionDetails,omitempty"`  // minOccurs="0" maxOccurs="98"
}

type SpecialRequirementsDataDetailsType struct {

	// SSR seat characteristic
	SeatCharacteristics []formats.AlphaNumericString_Length1To2 `xml:"seatCharacteristics,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type SpecialRequirementsDetailsType struct {

	// To specify the Service Requirement of the customer
	ServiceRequirementsInfo *SpecialRequirementsTypeDetailsType `xml:"serviceRequirementsInfo"`

	// Seat details
	SeatDetails []*SpecialRequirementsDataDetailsType `xml:"seatDetails,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type SpecialRequirementsTypeDetailsType struct {

	// To specify the Service Classification of the Service Requirement.
	ServiceClassification formats.AlphaNumericString_Length1To4 `xml:"serviceClassification"`

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
	ServiceFreeText []formats.AlphaNumericString_Length1To70 `xml:"serviceFreeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type SpecificDataInformationType struct {

	// Carrier fee description
	DataTypeInformation *DataTypeInformationType `xml:"dataTypeInformation"`

	// Data information
	DataInformation []*DataInformationType `xml:"dataInformation,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type StationInformationTypeI struct {

	// Departure terminal
	Terminal formats.AlphaNumericString_Length1To3 `xml:"terminal,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year *formats.NumericInteger_Length4To4 `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month *formats.NumericInteger_Length1To2 `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day *formats.NumericInteger_Length1To2 `xml:"day,omitempty"`  // minOccurs="0"
}

type StructuredPeriodInformationType struct {

	// Effective date of period of operation
	BeginDateTime *StructuredDateTimeType `xml:"beginDateTime,omitempty"`  // minOccurs="0"

	// Discontinue date of period of operation
	EndDateTime *StructuredDateTimeType `xml:"endDateTime,omitempty"`  // minOccurs="0"

	// It is used with a period to give a restriction for days impacted. It permits for example to indicate on which days, a flight operates.
	Frequency *FrequencyType `xml:"frequency,omitempty"`  // minOccurs="0"
}

type TaxDetailsTypeI struct {

	// Duty/tax/fee rate
	Rate formats.AlphaNumericString_Length1To18 `xml:"rate,omitempty"`  // minOccurs="0"

	// Country, coded
	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`  // minOccurs="0"

	// Currency, coded
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// Duty/Tax fee type, coded
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	// Amount type qualifier, coded
	AmountQualifier formats.AlphaNumericString_Length1To3 `xml:"amountQualifier,omitempty"`  // minOccurs="0"
}

type TaxType struct {

	// Duty tax fee category, coded
	WithholdTaxSurcharge formats.AlphaNumericString_Length1To3 `xml:"withholdTaxSurcharge,omitempty"`  // minOccurs="0"

	// TAX DETAILS
	TaxDetail []*TaxDetailsTypeI `xml:"taxDetail,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type TicketNumberDetailsTypeI struct {

	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`  // minOccurs="0"
}

type TicketNumberTypeI struct {

	DocumentDetails []*TicketNumberDetailsTypeI `xml:"documentDetails"`  // maxOccurs="99"
}

type TicketingPriceSchemeType struct {

	// PSR (Price Scheme Reference): unique reference of the price scheme as a 8 digit number.
	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber"`

	// Price Scheme Name
	Name formats.AlphaNumericString_Length1To35 `xml:"name,omitempty"`  // minOccurs="0"

	// Price Scheme Status. Is the price scheme valid for service fee calculation ?
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`  // minOccurs="0"

	// free flow description of the price scheme
	Description formats.AlphaNumericString_Length1To250 `xml:"description,omitempty"`  // minOccurs="0"
}

type TrafficRestrictionDetailsTypeI struct {

	// Traffic Restriction code
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`  // minOccurs="0"
}

type TrafficRestrictionTypeI struct {

	// Traffic Restriction Details
	TrafficRestrictionDetails []*TrafficRestrictionDetailsTypeI `xml:"trafficRestrictionDetails,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type TravelFlightInformationType struct {

	// Cabin identification
	CabinId *CabinIdentificationType `xml:"cabinId,omitempty"`  // minOccurs="0"

	// Company Identification
	CompanyIdentity []*CompanyIdentificationType_120719C `xml:"companyIdentity,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Type of flight details
	FlightDetail *ProductTypeDetailsType_120801C `xml:"flightDetail,omitempty"`  // minOccurs="0"

	// Details of included  connecting points
	InclusionDetail []*ConnectPointDetailsType_195492C `xml:"inclusionDetail,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Further connection details
	ExclusionDetail []*ConnectPointDetailsType `xml:"exclusionDetail,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Nb of connections for each requested segment of the journey.
	UnitNumberDetail []*NumberOfUnitDetailsTypeI `xml:"unitNumberDetail,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type TravelFlightInformationType_134790S struct {

	// Cabin identification
	CabinId *CabinIdentificationType `xml:"cabinId,omitempty"`  // minOccurs="0"

	// Company Identification
	CompanyIdentity []*CompanyIdentificationType_120719C `xml:"companyIdentity,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Type of flight details
	FlightDetail *ProductTypeDetailsType_120801C `xml:"flightDetail,omitempty"`  // minOccurs="0"

	// Details of included connect point
	InclusionDetail []*ConnectPointDetailsType_195492C `xml:"inclusionDetail,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Further connection details
	ExclusionDetail []*ConnectPointDetailsType `xml:"exclusionDetail,omitempty"`  // minOccurs="0" maxOccurs="20"

	// Nb of connections allowed at requested segment level.
	UnitNumberDetail []*NumberOfUnitDetailsTypeI `xml:"unitNumberDetail,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type TravelProductInformationTypeI struct {

	// Flight Date
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`  // minOccurs="0"

	// Board point
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails"`

	// Off point
	OffpointDetails *LocationTypeI `xml:"offpointDetails"`

	// Flight Carrier
	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails"`

	// Flight identification
	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`  // minOccurs="0"

	// Identify flight part of the journey
	FlightTypeDetails *ProductTypeDetailsTypeI `xml:"flightTypeDetails,omitempty"`  // minOccurs="0"
}

type TravellerDetailsType struct {

	// Direct reference of passenger assigned by requesting system.
	Ref formats.NumericInteger_Length1To3 `xml:"ref"`

	// Traveller is an infant
	InfantIndicator *formats.NumericInteger_Length1To1 `xml:"infantIndicator,omitempty"`  // minOccurs="0"
}

type TravellerReferenceInformationType struct {

	// Requested passenger type
	Ptc []formats.AlphaNumericString_Length1To6 `xml:"ptc,omitempty"`  // minOccurs="0" maxOccurs="3"

	// Traveller details
	Traveller []*TravellerDetailsType `xml:"traveller"`  // maxOccurs="9"
}

type UniqueIdDescriptionType struct {

	// Reference qualifier.
	PassengerFeeRefQualif formats.AlphaNumericString_Length1To3 `xml:"passengerFeeRefQualif,omitempty"`  // minOccurs="0"
}

type UserIdentificationType struct {

	// Originator Identification Details
	OfficeIdentification *OriginatorIdentificationDetailsTypeI `xml:"officeIdentification,omitempty"`  // minOccurs="0"

	// Used to specify which kind of info is given in DE 9900.
	OfficeType formats.AlphaNumericString_Length1To1 `xml:"officeType,omitempty"`  // minOccurs="0"

	// The code given to an agent by the originating reservation system.
	OfficeCode formats.AlphaNumericString_Length1To30 `xml:"officeCode,omitempty"`  // minOccurs="0"
}
