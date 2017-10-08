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
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A passwordInfo"`

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
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A originIdentification"`

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
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A dutyCode"`

	// This composite  duty code information. For Duty code info, DUT must be specified in the Qualifier.
	DutyCodeDetails ReferencingDetailsTypeI `xml:"dutyCodeDetails,omitempty"`
}

type ReferencingDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A dutyCodeDetails"`

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
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A systemDetails"`

	// This field contains a workstation Identifier. It is used to retrieve the physical origin of the request (mainly for printing purposes) .
	WorkstationId *AlphaNumericStringLength1To35 `xml:"workstationId,omitempty"`

	// Used to specify an organization when User Id logon is used.
	OrganizationDetails SystemDetailsTypeI `xml:"organizationDetails,omitempty"`

	// Explain what type of workstation ID is stored in data element 3148
	IdQualifier *AlphaNumericStringLength1To1 `xml:"idQualifier,omitempty"`
}

type SystemDetailsTypeI struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A organizationDetails"`

	// This DE is used to specify an organization Id such as BA, 1A or QF.
	OrganizationId AlphaNumericStringLength1To35 `xml:"organizationId,omitempty"`
}

type TerminalLocationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A locationInfo"`

	// To convey information related to a specific Facility.
	FacilityDetails *FacilityInformationType `xml:"facilityDetails,omitempty"`
}

type UserIdentificationType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A userIdentifier"`

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
	XMLName xml.Name `xml:"http://xml.amadeus.com/ws/2009/01/WBS_Session-2.0.xsd Session"`

	// This element defines the identifier part of the SessionId.
	SessionId string `xml:"SessionId,omitempty"`

	// This element defines the sequence number of the SessionId.
	SequenceNumber int `xml:"SequenceNumber,omitempty"`

	// This element defines the SecurityToken of the SessionId.
	SecurityToken string `xml:"SecurityToken,omitempty"`
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
