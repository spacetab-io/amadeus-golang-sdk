package command_cryptic

import "encoding/xml"

type CommandCryptic struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/HSFREQ_07_3_1A Command_Cryptic"`

	MessageAction *MessageAction `xml:"messageAction,omitempty"`

	NumberOfUnits *NumberOfUnits `xml:"numberOfUnits,omitempty"`

	IntelligentWorkstationInfo *IntelligentWorkstationInfo `xml:"intelligentWorkstationInfo,omitempty"`

	LongTextString *LongTextString `xml:"longTextString,omitempty"`
}

type MessageAction struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`

	ResponseType string `xml:"responseType,omitempty"`
}

type MessageFunctionDetails struct {
	BusinessFunction string `xml:"businessFunction,omitempty"`

	MessageFunction string `xml:"messageFunction,omitempty"`

	AdditionalMessageFunction string `xml:"additionalMessageFunction,omitempty"`
}

type NumberOfUnits struct {
	NumberOfUnitsDetails1 *NumberOfUnitsDetails1 `xml:"numberOfUnitsDetails1,omitempty"`

	NumberOfUnitsDetails2 *NumberOfUnitsDetails2 `xml:"numberOfUnitsDetails2,omitempty"`
}

type NumberOfUnitsDetails1 struct {
	Units float64 `xml:"units,omitempty"`

	UnitsQualifier string `xml:"unitsQualifier,omitempty"`
}

type NumberOfUnitsDetails2 struct {
	Units float64 `xml:"units,omitempty"`

	UnitsQualifier string `xml:"unitsQualifier,omitempty"`
}

type IntelligentWorkstationInfo struct {
	CompanyIdentification string `xml:"companyIdentification,omitempty"`
}

type LongTextString struct {
	TextStringDetails string `xml:"textStringDetails,omitempty"`
}
