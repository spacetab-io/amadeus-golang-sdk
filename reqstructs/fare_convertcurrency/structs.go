package fare_convertcurrency

import "encoding/xml"

type FareConvertCurrency struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/FCUQCQ_08_1_1A Fare_ConvertCurrency"`

	Message *Message `xml:"message"`

	ConversionDate *ConversionDate `xml:"conversionDate,omitempty"`  // minOccurs="0"

	ConversionRate *ConversionRate `xml:"conversionRate,omitempty"`  // minOccurs="0"

	ConversionDetails []*ConversionDetails `xml:"conversionDetails,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type Message struct {

	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`  // minOccurs="0"
}

type MessageFunctionDetails struct {

	MessageFunction string `xml:"messageFunction,omitempty"`  // minOccurs="0"
}

type ConversionDate struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type DateAndTimeDetails struct {

	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	Date string `xml:"date,omitempty"`  // minOccurs="0"

	Time *Time `xml:"time,omitempty"`  // minOccurs="0"
}

type Time struct {
}

type ConversionRate struct {

	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`
}

type ConversionRateDetails struct {

	RateType string `xml:"rateType"`

	Rate *Rate `xml:"rate,omitempty"`  // minOccurs="0"
}

type Rate struct {
}

type ConversionDetails struct {

	ConversionDirection *ConversionDirection `xml:"conversionDirection"`

	CurrencyInfo *CurrencyInfo `xml:"currencyInfo,omitempty"`  // minOccurs="0"

	AmountInfo *AmountInfo `xml:"amountInfo,omitempty"`  // minOccurs="0"

	LocationInfo *LocationInfo `xml:"locationInfo,omitempty"`  // minOccurs="0"
}

type ConversionDirection struct {

	SelectionDetails []*SelectionDetails `xml:"selectionDetails"`  // maxOccurs="99"
}

type SelectionDetails struct {

	Option string `xml:"option"`

	OptionInformation string `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type CurrencyInfo struct {

	ConversionRateDetails *ConversionRateDetails1 `xml:"conversionRateDetails"`
}

type ConversionRateDetails1 struct {

	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type AmountInfo struct {

	MonetaryDetails []*MonetaryDetails `xml:"monetaryDetails"`  // maxOccurs="20"
}

type MonetaryDetails struct {

	TypeQualifier string `xml:"typeQualifier"`

	Amount string `xml:"amount,omitempty"`  // minOccurs="0"
}

type LocationInfo struct {

	LocationType string `xml:"locationType"`

	LocationDescription *LocationDescription `xml:"locationDescription,omitempty"`  // minOccurs="0"

	FirstLocationDetails *FirstLocationDetails `xml:"firstLocationDetails,omitempty"`  // minOccurs="0"
}

type LocationDescription struct {

	Code string `xml:"code,omitempty"`  // minOccurs="0"

	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type FirstLocationDetails struct {

	Code string `xml:"code,omitempty"`  // minOccurs="0"

	Name string `xml:"name,omitempty"`  // minOccurs="0"
}
