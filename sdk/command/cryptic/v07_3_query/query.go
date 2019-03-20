package Command_Cryptic_v07_3 // hsfreq073

import "encoding/xml"

type CommandCryptic struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/HSFREQ_07_3_1A Command_Cryptic"`

	MessageAction *MessageAction `xml:"messageAction"`

	NumberOfUnits *NumberOfUnits `xml:"numberOfUnits,omitempty"` // minOccurs="0"

	IntelligentWorkstationInfo *IntelligentWorkstationInfo `xml:"intelligentWorkstationInfo,omitempty"` // minOccurs="0"

	LongTextString *LongTextString `xml:"longTextString"`
}

type MessageAction struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails"`

	ResponseType string `xml:"responseType,omitempty"` // minOccurs="0"
}

type MessageFunctionDetails struct {
	BusinessFunction string `xml:"businessFunction,omitempty"` // minOccurs="0"

	MessageFunction string `xml:"messageFunction"`

	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"` // minOccurs="0" maxOccurs="20"
}

type NumberOfUnits struct {
	NumberOfUnitsDetails1 *NumberOfUnitsDetails1 `xml:"numberOfUnitsDetails1"`

	NumberOfUnitsDetails2 []*NumberOfUnitsDetails2 `xml:"numberOfUnitsDetails2,omitempty"` // minOccurs="0" maxOccurs="8"
}

type NumberOfUnitsDetails1 struct {
	Units float64 `xml:"units"`

	UnitsQualifier string `xml:"unitsQualifier,omitempty"` // minOccurs="0"
}

type NumberOfUnitsDetails2 struct {
	Units *float64 `xml:"units,omitempty"` // minOccurs="0"

	UnitsQualifier string `xml:"unitsQualifier,omitempty"` // minOccurs="0"
}

type IntelligentWorkstationInfo struct {
	CompanyIdentification string `xml:"companyIdentification,omitempty"` // minOccurs="0"
}

type LongTextString struct {
	TextStringDetails string `xml:"textStringDetails"`
}
