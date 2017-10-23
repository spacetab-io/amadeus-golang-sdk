package amadeus

import "encoding/xml"

type ApplicationIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A ApplicationIdentificationType"`

	// Application Label. Exemple : NGD. Label is the first part of the ApplicationId. Label is considered as an Internal Id, as it is the key of the application.
	InternalId *AlphaNumericStringLength1To35 `xml:"internalId,omitempty"`

	// Application Index. Index is part of the ApplicationId.
	SeqNumber *AlphaNumericStringLength1To6 `xml:"seqNumber,omitempty"`
}

type ApplicationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A applicationId"`

	// Identification of the application.
	ApplicationDetails *ApplicationIdentificationType `xml:"applicationDetails,omitempty"`
}

type BinaryDataType struct {
	// Length of the data element 114Z. The unit is given in number of binary characters (bytes).
	DataLength NumericIntegerLength1To15 `xml:"dataLength,omitempty"`

	// type of the data. When E is specified, this means that the password is not crypted. When B specified, this means that the password is crypted.
	DataType AlphaNumericStringLength1To1 `xml:"dataType,omitempty"`

	// used to store binary data
	BinaryData AlphaNumericStringLength1To99999 `xml:"binaryData,omitempty"`
}

type ConversationIDType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A conversationClt"`

	// Sender identification
	SenderIdentification AlphaNumericStringLength1To35 `xml:"senderIdentification,omitempty"`

	// Recipient identification
	RecipientIdentification AlphaNumericStringLength1To35 `xml:"recipientIdentification,omitempty"`

	// Sender's interchange control reference
	SenderInterchangeControlReference AlphaNumericStringLength1To14 `xml:"senderInterchangeControlReference,omitempty"`

	// Recipient's interchange control reference
	RecipientInterchangeControlReference AlphaNumericStringLength1To14 `xml:"recipientInterchangeControlReference,omitempty"`
}

type FacilityInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A FacilityInformationType"`

	// Type of the Facility, coded. E.g.: - Check-In Desk - Gate... The codeset is not yet defined.
	Type *AlphaNumericStringLength1To3 `xml:"type,omitempty"`

	// Unique Reference to a Facility of a given Type in a terminal. This can be several kind of values. E.g.: - 25 (means Gate 25 when associated  Facility Type= Gate, or Check-in Desk 25 when associated Facility Type= Check-In Desk) - BAEXC: means Ba-Executive Club Lounge when associated to
	Identifier *AlphaNumericStringLength1To5 `xml:"identifier,omitempty"`
}

type LocationIdentificationBatchTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A LocationIdentificationBatchTypeU"`

	// Airport(139) or City (227) code
	Code *AlphaNumericStringLength1To35 `xml:"code,omitempty"`

	// Discriminator between airport or city.
	Qualifier *AlphaNumericStringLength1To17 `xml:"qualifier,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// AMADEUS Office Id the workstation belongs to. Must be empty if a workstation Id is specified in the SYS segment.
	SourceOffice AlphaNumericStringLength1To9 `xml:"sourceOffice,omitempty"`
}

type PlaceLocationIdentificationTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A workstationPos"`

	// Type of location
	LocationType *AlphaNumericStringLength1To3 `xml:"locationType,omitempty"`

	// Description of the location
	LocationDescription *LocationIdentificationBatchTypeU `xml:"locationDescription,omitempty"`

	// Details on the location
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU `xml:"firstLocationDetails,omitempty"`
}

type ReferenceInformationTypeI struct {
	// This composite  duty code information. For Duty code info, DUT must be specified in the Qualifier.
	DutyCodeDetails ReferencingDetailsTypeI `xml:"dutyCodeDetails,omitempty"`
}

type ReferencingDetailsTypeI struct {
	// Code determining if the data is a tree identifier, a category identifier or a parent category identifier.
	ReferenceQualifier AlphaNumericStringLength1To3 `xml:"referenceQualifier,omitempty"`

	// Data value.
	ReferenceIdentifier AlphaNumericStringLength1To35 `xml:"referenceIdentifier,omitempty"`
}

type RelatedLocationOneIdentificationTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A RelatedLocationOneIdentificationTypeU"`

	// Terminal (180) or Building(300) code
	Code *AlphaNumericStringLength1To25 `xml:"code,omitempty"`

	// Discriminator between airport or city.
	Qualifier *AlphaNumericStringLength1To17 `xml:"qualifier,omitempty"`
}

type SystemDetailsInfoType struct {
	// This field contains a workstation Identifier. It is used to retrieve the physical origin of the request (mainly for printing purposes) .
	WorkstationId *AlphaNumericStringLength1To35 `xml:"workstationId,omitempty"`

	// Used to specify an organization when User Id logon is used.
	OrganizationDetails SystemDetailsTypeI `xml:"organizationDetails,omitempty"`

	// Explain what type of workstation ID is stored in data element 3148
	IdQualifier *AlphaNumericStringLength1To1 `xml:"idQualifier,omitempty"`
}

type SystemDetailsTypeI struct {
	// This DE is used to specify an organization Id such as BA, 1A or QF.
	OrganizationId AlphaNumericStringLength1To35 `xml:"organizationId,omitempty"`
}

type TerminalLocationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A locationInfo"`

	// To convey information related to a specific Facility.
	FacilityDetails *FacilityInformationType `xml:"facilityDetails,omitempty"`
}

type UserIdentificationType struct {
	// To specify the source office the workstation belongs to. Not used in the second repetition of the segment (if any, it will not be taken into account). Used when no workstation Id is specified in SYS segment.
	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification,omitempty"`

	// Used to specify which kind of info is given in DE 9900.
	OriginatorTypeCode AlphaNumericStringLength1To1 `xml:"originatorTypeCode,omitempty"`

	// Authority code of the requester If Sign Id : (Numeric Sine)+(Agent Initials) ex : 0001XV). If user Id : Logon User Id.
	Originator AlphaNumericStringLength1To30 `xml:"originator,omitempty"`
}

type ApplicationErrorDetailType struct {
	XMLName xml.Name `xml:"errorDetails"`

	// Code identifying the data validation error condition.
	ErrorCode *AlphaNumericStringLength1To5 `xml:"errorCode,omitempty"`

	// Identification of a code list.
	ErrorCategory *AlphaNumericStringLength1To3 `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner *AlphaNumericStringLength1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationType struct {
	XMLName xml.Name `xml:"applicationError"`

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails,omitempty"`
}

type BusinessProcessIdType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A conversationGrp"`

	// A unique reference to identify the process/business
	ProcessIdentifier *AlphaNumericStringLength1To10 `xml:"processIdentifier,omitempty"`
}

type OrganizationIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A OrganizationIdentificationType"`

	// Organization label (Company Id).
	Label *AlphaNumericStringLength1To10 `xml:"label,omitempty"`
}

type OrganizationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A organizationInfo"`

	// This composite is used to specify an organization details
	OrganizationDetails *OrganizationIdentificationType `xml:"organizationDetails,omitempty"`
}

type ResponseAnalysisDetailsType struct {
	XMLName xml.Name `xml:"processStatus"`

	// P must be specified when status of the process is OK.
	StatusCode *AlphaStringLength1To6 `xml:"statusCode,omitempty"`
}

type FreeTextQualificationTypeI struct {
	XMLName xml.Name `xml:"FreeTextQualificationTypeI"`

	// Subject
	TextSubjectQualifier *AlphaNumericStringLength1To3 `xml:"textSubjectQualifier,omitempty"`

	// Info Type
	InformationType *AlphaNumericStringLength1To4 `xml:"informationType,omitempty"`

	// Language
	Language *AlphaNumericStringLength1To3 `xml:"language,omitempty"`
}

type ErrorFreeTextQualificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A freeTextQualif"`

	// Subject
	Subject *AlphaNumericStringLength1To3 `xml:"subject,omitempty"`

	// Info Type
	InfoType *AlphaNumericStringLength1To4 `xml:"infoType,omitempty"`

	// Language
	Language *AlphaNumericStringLength1To3 `xml:"language,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	XMLName xml.Name `xml:"interactiveFreeText"`

	// Free Text Qualifier
	FreeTextQualifier *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`

	// Free Text
	FreeText *AlphaNumericStringLength1To70 `xml:"freeText,omitempty"`
}

type ErrorInteractiveFreeTextType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A interactiveFreeText"`

	// Free Text Qualifier
	FreeTextQualif *ErrorFreeTextQualificationType `xml:"freeTextQualif,omitempty"`

	// Free Text
	FreeText *AlphaNumericStringLength1To70 `xml:"freeText,omitempty"`
}

type FullLocationType struct {
	// Stores the location of the workstation.
	WorkstationPos *PlaceLocationIdentificationTypeU `xml:"workstationPos,omitempty"`

	// Contains terminal and facility or worstation id.
	LocationInfo *TerminalLocationType `xml:"locationInfo,omitempty"`
}

type Session struct {
	//XMLName xml.Name `xml:"http://xml.amadeus.com/ws/2009/01/WBS_Session-2.0.xsd Session"`
	XMLName xml.Name `xml:"Session"`

	// This element defines the identifier part of the SessionId.
	SessionId string `xml:"SessionId"` //,omitempty

	// This element defines the sequence number of the SessionId.
	SequenceNumber string `xml:"SequenceNumber"` //,omitempty

	// This element defines the SecurityToken of the SessionId.
	SecurityToken string `xml:"SecurityToken"` //,omitempty
}

type ErrorSection struct {
	// Application Error
	ApplicationError *ApplicationErrorInformationType `xml:"applicationError,omitempty"`

	// Supplementary Info on the Error.
	InteractiveFreeText *ErrorInteractiveFreeTextType `xml:"interactiveFreeText,omitempty"`
}

type MessageFunctionDetailsType struct {
	BusinessFunction string `xml:"businessFunction,omitempty"`

	MessageFunction string `xml:"messageFunction,omitempty"`

	AdditionalMessageFunction string `xml:"additionalMessageFunction,omitempty"`
}

type MessageActionType struct {
	MessageFunctionDetails *MessageFunctionDetailsType `xml:"messageFunctionDetails,omitempty"`

	ResponseType string `xml:"responseType,omitempty"`
}

type LongTextStringType struct {
	TextStringDetails string `xml:"textStringDetails,omitempty"`
}

type NumberOfUnitsType struct {
	NumberOfUnitsDetails1 NumberOfUnitsDetailsType `xml:"numberOfUnitsDetails1,omitempty"`

	NumberOfUnitsDetails2 NumberOfUnitsDetailsType `xml:"numberOfUnitsDetails2,omitempty"`
}

type NumberOfUnitsDetailsType struct {
	Units float64 `xml:"units,omitempty"`

	UnitsQualifier string `xml:"unitsQualifier,omitempty"`
}

type IntelligentWorkstationInfoType struct {
	CompanyIdentification string `xml:"companyIdentification,omitempty"`
}

type ActionDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ActionDetailsTypeI"`

	// used for scrollling purposes
	NumberOfItemsDetails *ProcessingInformationTypeI `xml:"numberOfItemsDetails,omitempty"`
}

type ProcessingInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ProcessingInformationTypeI"`

	// determine if move up or move down required
	ActionQualifier *AlphaNumericStringLength1To3 `xml:"actionQualifier,omitempty"`
}

type AdditionalBusinessSourceInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A AdditionalBusinessSourceInformationTypeI"`

	// the office we are targetting
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType,omitempty"`

	// contains the office ID
	OriginatorDetails *OriginatorIdentificationDetailsType `xml:"originatorDetails,omitempty"`
}

type AccountingElementType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A AccountingElementType"`

	// Account number
	Number *AlphaNumericStringLength1To10 `xml:"number,omitempty"`
}

type AccountingInformationElementType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A AccountingInformationElementType"`

	// One of these 4 data elements is mandatory , but non in particular
	Account *AccountingElementType `xml:"account,omitempty"`
}

type CompanyIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A CompanyIdentificationTypeI"`

	// Marketing company.
	MarketingCompany *AlphaNumericStringLength1To3 `xml:"marketingCompany,omitempty"`
}

type DummySegmentTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A DummySegmentTypeI"`
}

type FrequentTravellerIdentificationCodeType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A FrequentTravellerIdentificationCodeType"`

	// Frequent Traveller Info. Repetition 2 is used only in the case we provide a customer value range (only one is accepted).
	FrequentTravellerDetails *FrequentTravellerIdentificationType `xml:"frequentTravellerDetails,omitempty"`

	DummyNET struct {
	} `xml:"Dummy.NET,omitempty"`
}

type FrequentTravellerIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A FrequentTravellerIdentificationType"`

	// This field specifies the Tier Level.   This is a 4 letter string indicating the airline's ranking of frequent flyers. It is not to be confused with Alliance tier.  If tierLevel is filled in a given FTI segment, customerValue must not be filled.
	TierLevel *AlphaNumericStringLength1To4 `xml:"tierLevel,omitempty"`

	// This field specifies the Customer value.   This is a 4 letter string indicating the airline's ranking of frequent flyers. It is not to be confused with Alliance tier.  If customerValue is filled in a given FTI segment, tierLevel field must not be filled.
	CustomerValue *NumericIntegerLength1To4 `xml:"customerValue,omitempty"`
}

type LocationTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A LocationTypeU"`

	// Office identification. It can contain wildcards.
	Name *AlphaNumericStringLength1To9 `xml:"name,omitempty"`
}

type OriginAndDestinationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A OriginAndDestinationDetailsTypeI"`

	// Board point
	Origin *AlphaNumericStringLength3To3 `xml:"origin,omitempty"`

	// Off point
	Destination *AlphaNumericStringLength3To3 `xml:"destination,omitempty"`
}

type OriginatorIdentificationDetailsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A OriginatorIdentificationDetailsTypeI"`

	// the office that is being targetted
	InHouseIdentification1 *AlphaNumericStringLength1To9 `xml:"inHouseIdentification1,omitempty"`
}

type PartyIdentifierTypeU struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A PartyIdentifierTypeU"`

	// GDS identifier: 1A, 1S, 1G.
	PartyIdentifier *AlphaNumericStringLength1To3 `xml:"partyIdentifier,omitempty"`
}

type PointOfSaleInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A PointOfSaleInformationType"`

	// Party identification.
	PointOfSale *PartyIdentifierTypeU `xml:"pointOfSale,omitempty"`

	// Office id in case the party identifier is 1A.
	LocationDetails *LocationTypeU `xml:"locationDetails,omitempty"`
}

type ProductDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ProductDetailsTypeI"`

	// Class designator.
	Designator *AlphaNumericStringLength1To1 `xml:"designator,omitempty"`
}

type ProductIdentificationDetailsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ProductIdentificationDetailsTypeI"`

	// Flight number.
	FlightNumber *AlphaNumericStringLength1To4 `xml:"flightNumber,omitempty"`
}

type ProductInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A ProductInformationTypeI"`

	// Booking class details.
	BookingClassDetails *ProductDetailsTypeI `xml:"bookingClassDetails,omitempty"`
}

type QueueInformationDetailsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A QueueInformationDetailsTypeI"`

	// queue number
	Number *NumericIntegerLength1To2 `xml:"number,omitempty"`
}

type QueueInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A QueueInformationTypeI"`

	// queue identification
	QueueDetails *QueueInformationDetailsType `xml:"queueDetails,omitempty"`
}

type RangeDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A RangeDetailsTypeI"`

	// define is a range or not
	RangeQualifier *AlphaNumericStringLength1To3 `xml:"rangeQualifier,omitempty"`

	// define the start and possible end point of the scan
	RangeDetails *RangeTypeI `xml:"rangeDetails,omitempty"`
}

type RangeTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A RangeTypeI"`

	// starting point of the scan
	Min *NumericIntegerLength1To18 `xml:"min,omitempty"`

	// ending point of the scan
	Max *NumericIntegerLength1To18 `xml:"max,omitempty"`
}

type RelatedProductInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A RelatedProductInformationTypeI"`

	// Status code
	StatusCode *AlphaNumericStringLength2To2 `xml:"statusCode,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SelectionDetailsInformationTypeI"`

	// used to determine if a new start or a continuation Also used for search and sort criteria on the ticketing, departure and creation dates
	Option *AlphaNumericStringLength1To3 `xml:"option,omitempty"`
}

type SelectionDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SelectionDetailsTypeI"`

	// used for search and sort criteria
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails,omitempty"`
}

type SourceTypeDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SourceTypeDetailsTypeI"`

	// not needed - but mandatory field So just stick a 4 in it !!
	SourceQualifier1 *AlphaNumericStringLength1To3 `xml:"sourceQualifier1,omitempty"`
}

type StatusDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StatusDetailsTypeI"`

	// Indicator showing what flight information will be transported.
	Indicator *AlphaNumericStringLength1To3 `xml:"indicator,omitempty"`
}

type StatusTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StatusTypeI"`

	// Flight status details.
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails,omitempty"`
}

type StructuredDateTimeInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StructuredDateTimeInformationType"`

	// used for date range only The date ranges are defined on central system as 1,2,3,4 The actual values of the ranges are set in the office profile
	TimeMode *NumericIntegerLength1To3 `xml:"timeMode,omitempty"`
}

type StructuredDateTimeType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StructuredDateTimeType"`

	// Year number.
	Year *YearYYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month *MonthMM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day *DayDD `xml:"day,omitempty"`
}

type StructuredPeriodInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A StructuredPeriodInformationType"`

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType `xml:"endDateTime,omitempty"`
}

type SubQueueInformationDetailsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SubQueueInformationDetailsTypeI"`

	// E for every category    A for cats with items to be worked C for category number N for nickname CN for both category number and nickname numeric for date range
	IdentificationType *AlphaNumericStringLength1To3 `xml:"identificationType,omitempty"`

	// category number
	ItemNumber *AlphaNumericStringLength1To3 `xml:"itemNumber,omitempty"`

	// used for nickname on inbound used for category name on outbound
	ItemDescription *AlphaNumericStringLength1To35 `xml:"itemDescription,omitempty"`
}

type SubQueueInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A SubQueueInformationTypeI"`

	// identifies the category or categories.
	SubQueueInfoDetails *SubQueueInformationDetailsType `xml:"subQueueInfoDetails,omitempty"`
}

type TransportIdentifierType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A TransportIdentifierType"`

	// Company identification.
	CompanyIdentification *CompanyIdentificationType `xml:"companyIdentification,omitempty"`

	// Flight details.
	FlightDetails *ProductIdentificationDetailsType `xml:"flightDetails,omitempty"`
}

type TravellerInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A TravellerInformationTypeI"`

	// Traveller surname information.
	PaxDetails *TravellerSurnameInformationType `xml:"paxDetails,omitempty"`
}

type TravellerSurnameInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A TravellerSurnameInformationTypeI"`

	// Passenger surname.
	Surname *AlphaNumericStringLength1To70 `xml:"surname,omitempty"`
}

type UserIdentificationTypeQL struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A UserIdentificationType"`

	// The last 2 characters of the sine of the agent who placed the PNR on queue.
	Originator *AlphaNumericStringLength1To2 `xml:"originator,omitempty"`
}

type SearchCriteriaType struct {
	// used to specify if ticketing, departure or creation options
	SearchOption *SelectionDetailsTypeI `xml:"searchOption,omitempty"`

	// used to specify the dates to be searched on
	Dates *StructuredPeriodInformationType `xml:"dates,omitempty"`
}

type FlightInformation struct {
	// It transport the type of flight information that will follow.
	FlightInformationType *StatusTypeI `xml:"flightInformationType,omitempty"`

	// Board point or Off Point.
	BoardPointOrOffPoint *OriginAndDestinationDetailsTypeI `xml:"boardPointOrOffPoint,omitempty"`

	// Airline code or Flight Number (in fact, airline + flight number)
	AirlineCodeOrFlightNumber *TransportIdentifierType `xml:"airlineCodeOrFlightNumber,omitempty"`

	// Booking class.
	ClassOfService *ProductInformationTypeI `xml:"classOfService,omitempty"`

	// Segment status code.
	SegmentStatus *RelatedProductInformationTypeI `xml:"segmentStatus,omitempty"`
}

type SortCriteriaType struct {
	// dummy for SDT clash
	Dumbo *DummySegmentTypeI `xml:"dumbo,omitempty"`

	// Determine the order of the display.
	SortOption *SelectionDetailsTypeI `xml:"sortOption,omitempty"`
}

type ErrorReturnType struct {
	// returns the error code
	ErrorDefinition *ApplicationErrorInformationTypeI `xml:"errorDefinition,omitempty"`

	// free text for the error
	ErrorText *FreeTextInformationType `xml:"errorText,omitempty"`
}

type QueueViewType struct {

	// details of who queue placed the PNR
	Agent *AdditionalBusinessSourceInformationType `xml:"agent,omitempty"`

	// queue being displayed
	QueueNumber *QueueInformationTypeR `xml:"queueNumber,omitempty"`

	// category and date range
	CategoryDetails *SubQueueInformationTypeR `xml:"categoryDetails,omitempty"`

	// date range
	Date *StructuredDateTimeInformationTypeR `xml:"date,omitempty"`

	// all 3 occurences are mandatory and show in order number of  1)PNRs on queue 2)PNRs found 3)PNRs scanned
	PnrCount *NumberOfUnitsTypeQLR `xml:"pnrCount,omitempty"`

	Item *ItemType `xml:"item,omitempty"`
}

type ItemType struct {
	// surname of the passenger only or GROUP for a group PNR
	PaxName *TravellerInformationTypeR `xml:"paxName,omitempty"`

	// record locator
	RecLoc *ReservationControlInformationTypeI `xml:"recLoc,omitempty"`

	// 1st segment found in the PNR - if there is data to send
	Segment *TravelProductInformationTypeI `xml:"segment,omitempty"`

	// details of who queue placed the PNR
	Agent *AdditionalBusinessSourceInformationType `xml:"agent,omitempty"`

	// contains 1,2 or 3 of the following Queue placement date/time Ticketing date PNR creation date
	Pnrdates *StructuredDateTimeInformationType_181906S `xml:"pnrdates,omitempty"`
}

type AdditionalBusinessSourceInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A AdditionalBusinessSourceInformationType"`

	// ORIGINATOR DETAILS
	OriginatorDetails *OriginatorIdentificationDetailsTypeR `xml:"originatorDetails,omitempty"`
}

type ApplicationErrorDetailTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ApplicationErrorDetailTypeI"`

	// error code
	ErrorCode *AlphaNumericStringLength1To3 `xml:"errorCode,omitempty"`

	// error category
	ErrorCategory *AlphaNumericStringLength1To3 `xml:"errorCategory,omitempty"`

	// error code owner
	ErrorCodeOwner *AlphaNumericStringLength1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ApplicationErrorInformationTypeI"`

	// error details
	ErrorDetails *ApplicationErrorDetailTypeI `xml:"errorDetails,omitempty"`
}

type CompanyIdentificationTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A CompanyIdentificationTypeI"`

	// carrier code or defined codeset
	MarketingCompany *AlphaNumericStringLength1To3 `xml:"marketingCompany,omitempty"`
}

type FreeTextDetailsType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A FreeTextDetailsType"`

	// qualifier
	TextSubjectQualifier *AlphaNumericStringLength1To3 `xml:"textSubjectQualifier,omitempty"`

	// source
	Source *AlphaNumericStringLength1To3 `xml:"source,omitempty"`

	// encoding
	Encoding *AlphaNumericStringLength1To3 `xml:"encoding,omitempty"`
}

type FreeTextInformationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A FreeTextInformationType"`

	// free text
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText *AlphaNumericStringLength1To199 `xml:"freeText,omitempty"`
}

type LocationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A LocationTypeI"`

	// board or off point
	TrueLocation *AlphaNumericStringLength1To3 `xml:"trueLocation,omitempty"`
}

type NumberOfUnitDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A NumberOfUnitDetailsTypeI"`

	// numbers of items
	NumberOfUnit *NumericIntegerLength1To5 `xml:"numberOfUnit,omitempty"`
}

type NumberOfUnitsTypeQLR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A NumberOfUnitsType"`

	// Number of Unit Details
	QuantityDetails *NumberOfUnitDetailsTypeI `xml:"quantityDetails,omitempty"`
}

type OriginatorIdentificationDetailsTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A OriginatorIdentificationDetailsTypeI"`

	// office ID of the agent who queue placed the PNR
	InHouseIdentification1 *AlphaNumericStringLength1To9 `xml:"inHouseIdentification1,omitempty"`

	// agent sign
	InHouseIdentification2 *AlphaNumericStringLength1To2 `xml:"inHouseIdentification2,omitempty"`
}

type ProductDateTimeTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ProductDateTimeTypeI"`

	// departure date
	DepartureDate *AlphaNumericStringLength1To35 `xml:"departureDate,omitempty"`
}

type ProductIdentificationDetailsTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ProductIdentificationDetailsTypeI"`

	// flight number
	FlightNumber *AlphaNumericStringLength1To35 `xml:"flightNumber,omitempty"`

	// operational suffix
	OperationalSuffix *AlphaNumericStringLength1To3 `xml:"operationalSuffix,omitempty"`
}

type QueueInformationDetailsTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A QueueInformationDetailsTypeI"`

	// queue number
	Number *NumericIntegerLength1To2 `xml:"number,omitempty"`
}

type QueueInformationTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A QueueInformationTypeI"`

	// queue identification
	QueueDetails *QueueInformationDetailsTypeR `xml:"queueDetails,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ReservationControlInformationDetailsTypeI"`

	// contains the record locator to be queue placed
	ControlNumber *AlphaNumericStringLength1To8 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A ReservationControlInformationTypeI"`

	// contains the record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}

type StructuredDateTimeInformationTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A StructuredDateTimeInformationType"`

	// used for date range only The date ranges are defined on central system as 1,2,3,4 The actual values of the ranges are set in the office profile
	TimeMode *NumericIntegerLength1To3 `xml:"timeMode,omitempty"`
}

type StructuredDateTimeInformationType_181906S struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A StructuredDateTimeInformationType_181906S"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	TimeMode *AlphaNumericStringLength1To3 `xml:"timeMode,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeTypeR `xml:"dateTime,omitempty"`
}

type StructuredDateTimeTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A StructuredDateTimeType"`

	// Year number.
	Year *YearYYYY `xml:"year,omitempty"`

	// Month number of the year (1 being first month, 0 being NULL data)
	Month *NumericIntegerLength1To2 `xml:"month,omitempty"`

	// day number of the month (1 being first day of the month, 0 is null data)
	Day *NumericIntegerLength1To2 `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour *HourHH `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes *MinuteMM `xml:"minutes,omitempty"`
}

type SubQueueInformationDetailsTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A SubQueueInformationDetailsTypeI"`

	// E for every category    A for cats with items to be worked C for category number N for nickname CN for both category number and nickname numeric for date range
	IdentificationType *AlphaNumericStringLength1To3 `xml:"identificationType,omitempty"`

	// category number
	ItemNumber *AlphaNumericStringLength1To3 `xml:"itemNumber,omitempty"`

	// used for nickname on inbound used for category name on outbound
	ItemDescription *AlphaNumericStringLength1To35 `xml:"itemDescription,omitempty"`
}

type SubQueueInformationTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A SubQueueInformationTypeI"`

	// identifies the category or categories.
	SubQueueInfoDetails *SubQueueInformationDetailsTypeR `xml:"subQueueInfoDetails,omitempty"`
}

type TravelProductInformationTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A TravelProductInformationTypeI"`

	// departure date
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`

	// board point
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	// off point if present
	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`

	// carrier code or segment type
	CompanyDetails *CompanyIdentificationTypeR `xml:"companyDetails,omitempty"`

	// flight number and suffix
	FlightIdentification *ProductIdentificationDetailsTypeR `xml:"flightIdentification,omitempty"`
}

type TravellerInformationTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A TravellerInformationTypeI"`

	PaxDetails *TravellerSurnameInformationTypeR `xml:"paxDetails,omitempty"`
}

type TravellerSurnameInformationTypeR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A TravellerSurnameInformationTypeI"`

	// surname of the passenger or GROUP for a group PNR
	Surname *AlphaNumericStringLength1To70 `xml:"surname,omitempty"`

	Type_ *AlphaNumericStringLength1To3 `xml:"type,omitempty"`

	Quantity *NumericIntegerLength1To15 `xml:"quantity,omitempty"`
}
