package amadeus

import "encoding/xml"

type SecurityAuthenticate struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A Security_Authenticate"`

	// It contains conversation properties between the SI and the JFE.
	ConversationClt *ConversationIDType `xml:"conversationClt,omitempty"`

	// This segment is dedicated to specify a user identifier. The first one  for the main logon and the second one to the delegate user if any. Either a SignId or a UserId can be used, depending on DE 9972 value (Z for Sign Id, U for User Id). The value itself is specified in DE 9904.
	UserIdentifier UserIdentificationType `xml:"userIdentifier,omitempty"`

	// This segment can be used to specify an AMADEUS Duty Code for the TPF sign in. DUT must be specified in the qualifier.
	DutyCode ReferenceInformationTypeI `xml:"dutyCode,omitempty"`

	// This segment is used to specify the system which is used by the user. The workstation Id (DE 3148) can be specified to determine the location in terms of office of the request (if no workstation Id, the source office must be specified in the first UID segment). Moreover, if a User Id is used instead of a Sign Id, the organization must be specified in DE 9906. WorkstationId is actually the Atid
	SystemDetails SystemDetailsInfoType `xml:"systemDetails,omitempty"`

	// These segments contain the password information. Two segments can be used in case of a New password is required.
	PasswordInfo BinaryDataType `xml:"passwordInfo,omitempty"`

	FullLocation *FullLocationType `xml:"fullLocation,omitempty"`

	// Conatins Baleb and INdex of application. ex: JFE 1
	ApplicationId *ApplicationType `xml:"applicationId,omitempty"`
}

type SecurityAuthenticateReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A Security_AuthenticateReply"`

	ErrorSection *ErrorSection `xml:"errorSection,omitempty"`

	// This segment is only used if process is OK. In that case P is specified.
	ProcessStatus *ResponseAnalysisDetailsType `xml:"processStatus,omitempty"`

	// This segment is used to specify organization details associated with the user.
	OrganizationInfo *OrganizationType `xml:"organizationInfo,omitempty"`

	// Identifier of a group of conversation, shared by sevreal mono-signed conversations.
	ConversationGrp *BusinessProcessIdType `xml:"conversationGrp,omitempty"`
}

type SecuritySignOut struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOQ_04_1_1A Security_SignOut"`

	// It contains conversation properties between the SI and the JFE.
	ConversationClt *ConversationIDType `xml:"conversationClt,omitempty"`
}

type SecuritySignOutReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A Security_SignOutReply"`

	ErrorSection struct {
		// Application Error
		ApplicationError *ApplicationErrorInformationType `xml:"applicationError,omitempty"`

		// Supplementary Info on the Error.
		InteractiveFreeText *InteractiveFreeTextTypeI `xml:"interactiveFreeText,omitempty"`
	} `xml:"errorSection,omitempty"`

	// This segment is only used if process is OK. In that case P is specified.
	ProcessStatus *ResponseAnalysisDetailsType `xml:"processStatus,omitempty"`
}

type CommandCryptic struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/HSFREQ_07_3_1A Command_Cryptic"`

	MessageAction *MessageActionType `xml:"messageAction,omitempty"`

	NumberOfUnits *NumberOfUnitsType `xml:"numberOfUnits,omitempty"`

	IntelligentWorkstationInfo *IntelligentWorkstationInfoType `xml:"intelligentWorkstationInfo,omitempty"`

	LongTextString *LongTextStringType `xml:"longTextString,omitempty"`
}

type CommandCrypticReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/HSFRES_07_3_1A Command_CrypticReply"`

	MessageActionDetails struct {
		MessageFunctionDetails struct {
			BusinessFunction string `xml:"businessFunction,omitempty"`

			MessageFunction string `xml:"messageFunction,omitempty"`
		} `xml:"messageFunctionDetails,omitempty"`

		ResponseType string `xml:"responseType,omitempty"`
	} `xml:"messageActionDetails,omitempty"`

	LongTextString struct {
		TextStringDetails string `xml:"textStringDetails,omitempty"`
	} `xml:"longTextString,omitempty"`
}
