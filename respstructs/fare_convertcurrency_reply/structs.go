package fare_convertcurrency_reply

//import "encoding/xml"

type FareConvertCurrencyReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FCUQCR_08_1_1A Fare_ConvertCurrencyReply"`

	Message *Message `xml:"message"`

	StatusDetails *StatusDetails `xml:"statusDetails,omitempty"`  // minOccurs="0"

	FreeTextInfo []*FreeTextInfo `xml:"freeTextInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	ErrorInfo []*ErrorInfo `xml:"errorInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	ConversionRoundingInfo []*ConversionRoundingInfo `xml:"conversionRoundingInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	InvolvedCurrency []*InvolvedCurrency `xml:"involvedCurrency,omitempty"`  // minOccurs="0" maxOccurs="9"

	ConversionDetails []*ConversionDetails `xml:"conversionDetails,omitempty"`  // minOccurs="0" maxOccurs="999"

	CountryDetails []*CountryDetails `xml:"countryDetails,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type Message struct {

	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`  // minOccurs="0"
}

type MessageFunctionDetails struct {

	MessageFunction string `xml:"messageFunction,omitempty"`  // minOccurs="0"
}

type StatusDetails struct {

	StatusInformation []*StatusInformation `xml:"statusInformation"`  // maxOccurs="99"
}

type StatusInformation struct {

	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"

	Action string `xml:"action,omitempty"`  // minOccurs="0"

	Type string `xml:"type,omitempty"`  // minOccurs="0"

	Description string `xml:"description,omitempty"`  // minOccurs="0"
}

type FreeTextInfo struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FreeTextQualification struct {

	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	Status string `xml:"status,omitempty"`  // minOccurs="0"

	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	Language string `xml:"language,omitempty"`  // minOccurs="0"
}

type ErrorInfo struct {

	ErrorIdentification *ErrorIdentification `xml:"errorIdentification"`

	ErrorText *ErrorText `xml:"errorText,omitempty"`  // minOccurs="0"
}

type ErrorIdentification struct {

	ErrorDetails *ErrorDetails `xml:"errorDetails"`
}

type ErrorDetails struct {

	ErrorCode string `xml:"errorCode"`

	ErrorCategory string `xml:"errorCategory,omitempty"`  // minOccurs="0"

	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`  // minOccurs="0"
}

type ErrorText struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ConversionRoundingInfo struct {

	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`
}

type ConversionRateDetails struct {

	ConversionType string `xml:"conversionType"`

	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	RoundingUnit *RoundingUnit `xml:"roundingUnit,omitempty"`  // minOccurs="0"

	RoundingMethod string `xml:"roundingMethod,omitempty"`  // minOccurs="0"
}

type RoundingUnit struct {
}

type InvolvedCurrency struct {

	ConversionDirection *ConversionDirection `xml:"conversionDirection"`

	CurrencyDetails *CurrencyDetails `xml:"currencyDetails,omitempty"`  // minOccurs="0"
}

type ConversionDirection struct {

	SelectionDetails []*SelectionDetails `xml:"selectionDetails"`  // maxOccurs="99"
}

type SelectionDetails struct {

	Option string `xml:"option"`

	OptionInformation string `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type CurrencyDetails struct {

	ConversionRateDetails *ConversionRateDetails1 `xml:"conversionRateDetails"`
}

type ConversionRateDetails1 struct {

	Currency string `xml:"currency"`
}

type ConversionDetails struct {

	DateInfo *DateInfo `xml:"dateInfo"`

	ConversionRate *ConversionRate `xml:"conversionRate,omitempty"`  // minOccurs="0"

	ConvertedAmount []*ConvertedAmount `xml:"convertedAmount,omitempty"`  // minOccurs="0" maxOccurs="18"
}

type DateInfo struct {

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

	ConversionRateDetails *ConversionRateDetails2 `xml:"conversionRateDetails"`

	OtherConvRateDetails *OtherConvRateDetails `xml:"otherConvRateDetails,omitempty"`  // minOccurs="0"
}

type ConversionRateDetails2 struct {

	OriginCurrency string `xml:"originCurrency,omitempty"`  // minOccurs="0"

	RateType string `xml:"rateType"`

	Rate *Rate `xml:"rate,omitempty"`  // minOccurs="0"
}

type Rate struct {
}

type OtherConvRateDetails struct {

	DestinationCurrency string `xml:"destinationCurrency,omitempty"`  // minOccurs="0"
}

type ConvertedAmount struct {

	MonetaryInfo *MonetaryInfo `xml:"monetaryInfo"`

	AdditionalConversionDetails *AdditionalConversionDetails `xml:"additionalConversionDetails,omitempty"`  // minOccurs="0"
}

type MonetaryInfo struct {

	MonetaryDetails *MonetaryDetails `xml:"monetaryDetails"`
}

type MonetaryDetails struct {

	TypeQualifier string `xml:"typeQualifier"`

	Amount string `xml:"amount,omitempty"`  // minOccurs="0"

	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type AdditionalConversionDetails struct {

	ConversionRateDetails *ConversionRateDetails3 `xml:"conversionRateDetails"`

	OtherConvRateDetails []*OtherConvRateDetails1 `xml:"otherConvRateDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type ConversionRateDetails3 struct {

	ConversionType string `xml:"conversionType,omitempty"`  // minOccurs="0"

	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	RateType string `xml:"rateType,omitempty"`  // minOccurs="0"

	PricingAmount *PricingAmount `xml:"pricingAmount,omitempty"`  // minOccurs="0"

	ConvertedValueAmount *ConvertedValueAmount `xml:"convertedValueAmount,omitempty"`  // minOccurs="0"

	DutyTaxFeeType string `xml:"dutyTaxFeeType,omitempty"`  // minOccurs="0"

	MeasurementValue *MeasurementValue `xml:"measurementValue,omitempty"`  // minOccurs="0"

	MeasurementSignificance string `xml:"measurementSignificance,omitempty"`  // minOccurs="0"
}

type PricingAmount struct {
}

type ConvertedValueAmount struct {
}

type MeasurementValue struct {
}

type OtherConvRateDetails1 struct {

	ConversionType string `xml:"conversionType,omitempty"`  // minOccurs="0"

	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	RateType string `xml:"rateType,omitempty"`  // minOccurs="0"

	PricingAmount *PricingAmount `xml:"pricingAmount,omitempty"`  // minOccurs="0"

	ConvertedValueAmount *ConvertedValueAmount `xml:"convertedValueAmount,omitempty"`  // minOccurs="0"

	DutyTaxFeeType string `xml:"dutyTaxFeeType,omitempty"`  // minOccurs="0"

	MeasurementValue *MeasurementValue `xml:"measurementValue,omitempty"`  // minOccurs="0"

	MeasurementSignificance string `xml:"measurementSignificance,omitempty"`  // minOccurs="0"
}

type CountryDetails struct {

	CountryInfo *CountryInfo `xml:"countryInfo"`

	AdditionalCountryInfo []*AdditionalCountryInfo `xml:"additionalCountryInfo,omitempty"`  // minOccurs="0" maxOccurs="99"

	MarkerEndCountryDetails *MarkerEndCountryDetails `xml:"markerEndCountryDetails"`
}

type CountryInfo struct {

	LocationType string `xml:"locationType"`

	LocationDescription *LocationDescription `xml:"locationDescription,omitempty"`  // minOccurs="0"

	FirstLocationDetails *FirstLocationDetails `xml:"firstLocationDetails,omitempty"`  // minOccurs="0"

	SecondLocationDetails *SecondLocationDetails `xml:"secondLocationDetails,omitempty"`  // minOccurs="0"
}

type LocationDescription struct {

	Code string `xml:"code,omitempty"`  // minOccurs="0"

	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type FirstLocationDetails struct {

	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type SecondLocationDetails struct {

	Code string `xml:"code,omitempty"`  // minOccurs="0"
}

type AdditionalCountryInfo struct {

	LocationType string `xml:"locationType"`

	LocationDescription *LocationDescription1 `xml:"locationDescription,omitempty"`  // minOccurs="0"
}

type LocationDescription1 struct {

	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type MarkerEndCountryDetails struct {
}
