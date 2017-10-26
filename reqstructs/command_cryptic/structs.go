package command_cryptic

import "encoding/xml"

type CommandCryptic struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/HSFREQ_07_3_1A Command_Cryptic"`

	MessageAction struct {
		MessageFunctionDetails struct {
			BusinessFunction string `xml:"businessFunction,omitempty"`

			MessageFunction string `xml:"messageFunction,omitempty"`

			AdditionalMessageFunction string `xml:"additionalMessageFunction,omitempty"`
		} `xml:"messageFunctionDetails,omitempty"`

		ResponseType string `xml:"responseType,omitempty"`
	} `xml:"messageAction,omitempty"`

	NumberOfUnits struct {
		NumberOfUnitsDetails1 struct {
			Units float64 `xml:"units,omitempty"`

			UnitsQualifier string `xml:"unitsQualifier,omitempty"`
		} `xml:"numberOfUnitsDetails1,omitempty"`

		NumberOfUnitsDetails2 struct {
			Units float64 `xml:"units,omitempty"`

			UnitsQualifier string `xml:"unitsQualifier,omitempty"`
		} `xml:"numberOfUnitsDetails2,omitempty"`
	} `xml:"numberOfUnits,omitempty"`

	IntelligentWorkstationInfo struct {
		CompanyIdentification string `xml:"companyIdentification,omitempty"`
	} `xml:"intelligentWorkstationInfo,omitempty"`

	LongTextString struct {
		TextStringDetails string `xml:"textStringDetails,omitempty"`
	} `xml:"longTextString,omitempty"`
}
