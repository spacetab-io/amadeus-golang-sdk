package docissuance_issueticket

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type DocIssuanceIssueTicket struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A DocIssuance_IssueTicket"`

	// agent code
	AgentInfo *TicketAgentInfoTypeI `xml:"agentInfo,omitempty"`

	// override past date TST
	OverrideDate *StructuredDateTimeInformationType `xml:"overrideDate,omitempty"`

	// Segment, TST, Coupon or line selection
	Selection *ReferenceInfoType `xml:"selection,omitempty"`

	// Passenger Selection
	PaxSelection *ReferenceInformationType `xml:"paxSelection,omitempty"`

	// stock reference
	Stock *StockInformationType `xml:"stock,omitempty"`

	OptionGroup *OptionGroup `xml:"optionGroup,omitempty"`

	// Infant or Adult passenger
	InfantOrAdultAssociation *TravellerInformationType `xml:"infantOrAdultAssociation,omitempty"`

	// Contains other options that do take parameters.
	OtherCompoundOptions *CodedAttributeType `xml:"otherCompoundOptions,omitempty"`
}

type OptionGroup struct {
	// contains a list of switches, i.e. options that do not take any parameters.
	Switches *StatusTypeI `xml:"switches,omitempty"`

	// Contains sub-options that do take parameters.
	SubCompoundOptions *AttributeType `xml:"subCompoundOptions,omitempty"`

	// To override Alternative date for Invoice printing
	OverrideAlternativeDate *StructuredDateTimeInformationType `xml:"overrideAlternativeDate,omitempty"`
}

type AttributeInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A AttributeInformationTypeU"`

	// provides the sub-option name: EMPR Email to the specified address EMPRA Email to the address specified in the APE element EMPRN Email to addresses associated to nickname FAXR Fax to Specified Number FAXA Fax to Number specified in AP element of the PNR FAXN Fax to number associated to nickname LP Issuance langage PALT Printer Options / Alternate printer PCIH Printer Options / Customize Itinerary Header PDCT Printer Options / Document Type PDPR Printer Options / Distribution Profile PDUP Printer Options / Duplicate Itinerary PMNE Printer Options / Mnemonic PNUM Printer Options / Number POFP Printer Options / Office profile
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType,omitempty"`

	// provides the sub-option value
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type AttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A AttributeType"`

	// Details for the message criteria (name, value).
	CriteriaDetails *AttributeInformationTypeU `xml:"criteriaDetails,omitempty"`
}

type CodedAttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A CodedAttributeInformationType"`

	// provides the option name
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// provides a description for the attribute (option value)
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A CodedAttributeType"`

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`
}

type InternalIDDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A InternalIDDetailsTypeI"`

	// agent code or agent sine
	InhouseId formats.AlphaNumericString_Length1To9 `xml:"inhouseId,omitempty"`

	// Agent qualifier
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type ReferenceInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A ReferenceInfoType"`

	// Reference list (segment tattoos or PNR lines or TST number or Coupon number)
	ReferenceDetails *ReferencingDetailsType `xml:"referenceDetails,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DummyNET struct{}

type ReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A ReferenceInformationType"`

	// Used to convey the passenger tatoo or display number.
	PassengerReference *ReferencingDetailsType_108978C `xml:"passengerReference,omitempty"`
}

type ReferencingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A ReferencingDetailsType"`

	// Reference qualifier
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Reference value
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type ReferencingDetailsType_108978C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A ReferencingDetailsType_108978C"`

	// Reference qualifier
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Passenger tattoo.
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type StatusDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A StatusDetailsTypeI"`

	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type StatusTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A StatusTypeI"`

	StatusDetails *StatusDetailsTypeI `xml:"statusDetails,omitempty"`
}

type StockInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A StockInformationType"`

	// To convey Stock type
	StockTicketNumberDetails *StockTicketNumberDetailsType `xml:"stockTicketNumberDetails,omitempty"`
}

type StockTicketNumberDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A StockTicketNumberDetailsType"`

	// Stock Qualifier: 'S' stock control number
	Qualifier formats.AlphaNumericString_Length1To1 `xml:"qualifier,omitempty"`

	// Stock control number
	ControlNumber formats.AlphaNumericString_Length1To35 `xml:"controlNumber,omitempty"`
}

type StructuredDateTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A StructuredDateTimeInformationType"`

	// Date format
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A StructuredDateTimeType"`

	// Year number.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type TicketAgentInfoTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A TicketAgentInfoTypeI"`

	// agent ID details
	InternalIdDetails *InternalIDDetailsTypeI `xml:"internalIdDetails,omitempty"`
}

type TravellerInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A TravellerInformationType"`

	// Detail if the discount is adult only or infant only
	PaxDetails *TravellerSurnameInformationType `xml:"paxDetails,omitempty"`
}

type TravellerSurnameInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A TravellerSurnameInformationType"`

	// Customer type: A=adult IN = infant
	Type formats.AlphaNumericString_Length1To2 `xml:"type,omitempty"`
}
