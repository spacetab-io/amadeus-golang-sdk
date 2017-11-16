package docissuance_issueticket

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type DocIssuanceIssueTicket struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A DocIssuance_IssueTicket"`

	// agent code
	AgentInfo *TicketAgentInfoTypeI `xml:"agentInfo,omitempty"`  // minOccurs="0"

	// override past date TST
	OverrideDate *StructuredDateTimeInformationType `xml:"overrideDate,omitempty"`  // minOccurs="0"

	// Segment, TST, Coupon or line selection
	Selection []*ReferenceInfoType `xml:"selection,omitempty"`  // minOccurs="0" maxOccurs="3"

	// Passenger Selection
	PaxSelection []*ReferenceInformationType `xml:"paxSelection,omitempty"`  // minOccurs="0" maxOccurs="99"

	// stock reference
	Stock *StockInformationType `xml:"stock,omitempty"`  // minOccurs="0"

	OptionGroup []*OptionGroup `xml:"optionGroup,omitempty"`  // minOccurs="0" maxOccurs="99"

	// Infant or Adult passenger
	InfantOrAdultAssociation *TravellerInformationType `xml:"infantOrAdultAssociation,omitempty"`  // minOccurs="0"

	// Contains other options that do take parameters.
	OtherCompoundOptions []*CodedAttributeType `xml:"otherCompoundOptions,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type OptionGroup struct {

	// contains a list of switches, i.e. options that do not take any parameters.
	Switches *StatusTypeI `xml:"switches"`

	// Contains sub-options that do take parameters.
	SubCompoundOptions []*AttributeType `xml:"subCompoundOptions,omitempty"`  // minOccurs="0" maxOccurs="99"

	// To override Alternative date for Invoice printing
	OverrideAlternativeDate *StructuredDateTimeInformationType `xml:"overrideAlternativeDate,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type AttributeInformationTypeU struct {

	// provides the sub-option name: EMPR Email to the specified address EMPRA Email to the address specified in the APE element EMPRN Email to addresses associated to nickname FAXR Fax to Specified Number FAXA Fax to Number specified in AP element of the PNR FAXN Fax to number associated to nickname LP Issuance langage PALT Printer Options / Alternate printer PCIH Printer Options / Customize Itinerary Header PDCT Printer Options / Document Type PDPR Printer Options / Distribution Profile PDUP Printer Options / Duplicate Itinerary PMNE Printer Options / Mnemonic PNUM Printer Options / Number POFP Printer Options / Office profile
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType"`

	// provides the sub-option value
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type AttributeType struct {

	// Details for the message criteria (name, value).
	CriteriaDetails *AttributeInformationTypeU `xml:"criteriaDetails"`
}

type CodedAttributeInformationType struct {

	// provides the option name
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType"`

	// provides a description for the attribute (option value)
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeType struct {

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails"`
}

type InternalIDDetailsTypeI struct {

	// agent code or agent sine
	InhouseId formats.AlphaNumericString_Length1To9 `xml:"inhouseId"`

	// Agent qualifier
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`
}

type ReferenceInfoType struct {

	// Reference list (segment tattoos or PNR lines or TST number or Coupon number)
	ReferenceDetails []*ReferencingDetailsType `xml:"referenceDetails"`  // maxOccurs="99"
}

type ReferenceInformationType struct {

	// Used to convey the passenger tatoo or display number.
	PassengerReference *ReferencingDetailsType_108978C `xml:"passengerReference"`
}

type ReferencingDetailsType struct {

	// Reference qualifier
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// Reference value
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`  // minOccurs="0"
}

type ReferencingDetailsType_108978C struct {

	// Reference qualifier
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	// Passenger tattoo.
	Value formats.AlphaNumericString_Length1To35 `xml:"value"`
}

type StatusDetailsTypeI struct {

	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator"`
}

type StatusTypeI struct {

	StatusDetails *StatusDetailsTypeI `xml:"statusDetails"`
}

type StockInformationType struct {

	// To convey Stock type
	StockTicketNumberDetails *StockTicketNumberDetailsType `xml:"stockTicketNumberDetails"`
}

type StockTicketNumberDetailsType struct {

	// Stock Qualifier: 'S' stock control number
	Qualifier formats.AlphaNumericString_Length1To1 `xml:"qualifier"`

	// Stock control number
	ControlNumber formats.AlphaNumericString_Length1To35 `xml:"controlNumber,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeInformationType struct {

	// Date format
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType struct {

	// Year number.
	Year formats.Year_YYYY `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day"`
}

type TicketAgentInfoTypeI struct {

	// agent ID details
	InternalIdDetails *InternalIDDetailsTypeI `xml:"internalIdDetails"`
}

type TravellerInformationType struct {

	// Detail if the discount is adult only or infant only
	PaxDetails *TravellerSurnameInformationType `xml:"paxDetails"`
}

type TravellerSurnameInformationType struct {

	// Customer type: A=adult IN = infant
	Type formats.AlphaNumericString_Length1To2 `xml:"type"`
}
