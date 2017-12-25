package fare_pricepnrwithbookingclass

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/formats"
)

type FarePricePNRWithBookingClass struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_14_1_1A Fare_PricePNRWithBookingClass"`

	PricingOptionGroup []*PricingOptionGroup `xml:"pricingOptionGroup"`  // maxOccurs="999"
}

type PricingOptionGroup struct {

	PricingOptionKey *PricingOptionKey `xml:"pricingOptionKey"`

	OptionDetail *AttributeType `xml:"optionDetail,omitempty"`  // minOccurs="0"

	CarrierInformation *TransportIdentifierType `xml:"carrierInformation,omitempty"`  // minOccurs="0"

	Currency *CurrenciesType `xml:"currency,omitempty"`  // minOccurs="0"

	PenDisInformation *DiscountAndPenaltyInformationType `xml:"penDisInformation,omitempty"`  // minOccurs="0"

	MonetaryInformation *MonetaryInformationType `xml:"monetaryInformation,omitempty"`  // minOccurs="0"

	TaxInformation []*DutyTaxFeeDetailsType `xml:"taxInformation,omitempty"`  // minOccurs="0" maxOccurs="99"

	DateInformation []*StructuredDateTimeInformationType `xml:"dateInformation,omitempty"`  // minOccurs="0" maxOccurs="2"

	FrequentFlyerInformation *FrequentTravellerIdentificationCodeType `xml:"frequentFlyerInformation,omitempty"`  // minOccurs="0"

	FormOfPaymentInformation *FormOfPaymentType `xml:"formOfPaymentInformation,omitempty"`  // minOccurs="0"

	LocationInformation *PlaceLocationIdentificationType `xml:"locationInformation,omitempty"`  // minOccurs="0"

	PaxSegTstReference *ReferenceInfoType `xml:"paxSegTstReference,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type AttributeInformationTypeU struct {

	// Used for attribute value rather than attributeType
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType"`

	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type AttributeType struct {

	// Details for the message criteria (name, value).
	CriteriaDetails []*AttributeInformationTypeU `xml:"criteriaDetails"`  // maxOccurs="99"
}

type CompanyIdentificationTypeI struct {

	OtherCompany formats.AlphaNumericString_Length1To35 `xml:"otherCompany,omitempty"`  // minOccurs="0"
}

type CurrenciesType struct {

	FirstCurrencyDetails *CurrencyDetailsTypeU `xml:"firstCurrencyDetails,omitempty"`  // minOccurs="0"
}

type CurrencyDetailsTypeU struct {

	CurrencyQualifier formats.AlphaNumericString_Length1To3 `xml:"currencyQualifier"`

	CurrencyIsoCode formats.AlphaNumericString_Length1To3 `xml:"currencyIsoCode,omitempty"`  // minOccurs="0"
}

type DiscountAndPenaltyInformationType struct {

	DiscountPenaltyQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"discountPenaltyQualifier,omitempty"`  // minOccurs="0"

	DiscountPenaltyDetails []*DiscountPenaltyMonetaryInformationType `xml:"discountPenaltyDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type DiscountPenaltyMonetaryInformationType struct {

	Function formats.AMA_EDICodesetType_Length1to3 `xml:"function,omitempty"`  // minOccurs="0"

	AmountType formats.AMA_EDICodesetType_Length1to3 `xml:"amountType,omitempty"`  // minOccurs="0"

	Amount formats.AlphaNumericString_Length1To18 `xml:"amount,omitempty"`  // minOccurs="0"

	Rate formats.AlphaNumericString_Length1To35 `xml:"rate,omitempty"`  // minOccurs="0"

	Currency formats.AMA_EDICodesetType_Length1to3 `xml:"currency,omitempty"`  // minOccurs="0"
}

type DutyTaxFeeAccountDetailType struct {

	IsoCountry formats.AlphaNumericString_Length1To6 `xml:"isoCountry"`
}

type DutyTaxFeeDetailType struct {

	TaxRate formats.AlphaNumericString_Length1To17 `xml:"taxRate,omitempty"`  // minOccurs="0"

	TaxValueQualifier formats.AlphaNumericString_Length1To1 `xml:"taxValueQualifier,omitempty"`  // minOccurs="0"
}

type DutyTaxFeeDetailsType struct {

	TaxQualifier formats.AlphaNumericString_Length1To3 `xml:"taxQualifier"`

	TaxType *DutyTaxFeeAccountDetailType `xml:"taxType,omitempty"`  // minOccurs="0"

	TaxNature formats.AlphaNumericString_Length1To15 `xml:"taxNature,omitempty"`  // minOccurs="0"

	TaxData *DutyTaxFeeDetailType `xml:"taxData,omitempty"`  // minOccurs="0"
}

type FormOfPaymentDetailsType struct {

	Type formats.AlphaNumericString_Length1To10 `xml:"type"`

	Amount *formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`  // minOccurs="0"

	CreditCardNumber formats.AlphaNumericString_Length1To35 `xml:"creditCardNumber,omitempty"`  // minOccurs="0"
}

type FormOfPaymentType struct {

	// Details on the form of payment
	FormOfPayment *FormOfPaymentDetailsType `xml:"formOfPayment"`

	OtherFormOfPayment []*FormOfPaymentDetailsType `xml:"otherFormOfPayment,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type FrequentTravellerIdentificationCodeType struct {

	// Frequent Traveller Info
	FrequentTravellerDetails []*FrequentTravellerIdentificationType `xml:"frequentTravellerDetails"`  // maxOccurs="99"
}

type FrequentTravellerIdentificationType struct {

	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"`  // minOccurs="0"

	Number formats.AlphaNumericString_Length1To28 `xml:"number,omitempty"`  // minOccurs="0"

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`  // minOccurs="0"

	// For example : priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsType struct {

	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount *formats.NumericDecimal_Length1To35 `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"

	// location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`  // minOccurs="0"
}

type MonetaryInformationType struct {

	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type PlaceLocationIdentificationType struct {

	LocationType formats.AlphaNumericString_Length1To3 `xml:"locationType"`

	FirstLocationDetails *RelatedLocationOneIdentificationType `xml:"firstLocationDetails,omitempty"`  // minOccurs="0"

	SecondLocationDetails *RelatedLocationTwoIdentificationType `xml:"secondLocationDetails,omitempty"`  // minOccurs="0"
}

type PricingOptionKey struct {

	PricingOptionKey formats.AlphaNumericString_Length1To3 `xml:"pricingOptionKey"`
}

type ReferenceInfoType struct {

	ReferenceDetails []*ReferencingDetailsType `xml:"referenceDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferencingDetailsType struct {

	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`  // minOccurs="0"

	Value formats.AlphaNumericString_Length1To60 `xml:"value,omitempty"`  // minOccurs="0"
}

type RelatedLocationOneIdentificationType struct {

	Code formats.AlphaNumericString_Length1To25 `xml:"code,omitempty"`  // minOccurs="0"
}

type RelatedLocationTwoIdentificationType struct {

	Code formats.AlphaNumericString_Length1To25 `xml:"code,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeInformationType struct {

	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType struct {

	// Year number.
	Year formats.Year_YYYY `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`  // minOccurs="0"
}

type TransportIdentifierType struct {

	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`  // minOccurs="0"
}
