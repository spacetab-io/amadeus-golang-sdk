package fare_checkrules

import "encoding/xml"

type FareCheckRules struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/FARQNQ_07_1_1A Fare_CheckRules"`

	MsgType *MsgType `xml:"msgType"`

	AvailcabinStatus *AvailcabinStatus `xml:"availcabinStatus,omitempty"`  // minOccurs="0"

	ConversionRate *ConversionRate `xml:"conversionRate,omitempty"`  // minOccurs="0"

	PricingTickInfo *PricingTickInfo `xml:"pricingTickInfo,omitempty"`  // minOccurs="0"

	MultiCorporate *MultiCorporate `xml:"multiCorporate,omitempty"`  // minOccurs="0"

	ItemNumber *ItemNumber `xml:"itemNumber,omitempty"`  // minOccurs="0"

	DateOfFlight *DateOfFlight `xml:"dateOfFlight,omitempty"`  // minOccurs="0"

	FlightQualification []*FlightQualification `xml:"flightQualification,omitempty"`  // minOccurs="0" maxOccurs="99"

	TransportInformation []*TransportInformation `xml:"transportInformation,omitempty"`  // minOccurs="0" maxOccurs="99"

	TripDescription []*TripDescription `xml:"tripDescription,omitempty"`  // minOccurs="0" maxOccurs="99"

	PricingInfo []*PricingInfo `xml:"pricingInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	FareRule *FareRule `xml:"fareRule,omitempty"`  // minOccurs="0"
}

type MsgType struct {

	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`  // minOccurs="0"
}

type MessageFunctionDetails struct {

	MessageFunction string `xml:"messageFunction,omitempty"`  // minOccurs="0"
}

type AvailcabinStatus struct {

	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"`  // minOccurs="0"

	BookingClassDetails []*BookingClassDetails `xml:"bookingClassDetails"`  // maxOccurs="26"
}

type BookingClassDetails struct {

	Designator string `xml:"designator"`

	Option []string `xml:"option,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type ConversionRate struct {

	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`

	OtherConvRateDetails []*OtherConvRateDetails `xml:"otherConvRateDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type ConversionRateDetails struct {

	ConversionType string `xml:"conversionType,omitempty"`  // minOccurs="0"

	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	RateType string `xml:"rateType,omitempty"`  // minOccurs="0"

	PricingAmount *float64 `xml:"pricingAmount,omitempty"`  // minOccurs="0"

	ConvertedValueAmount *float64 `xml:"convertedValueAmount,omitempty"`  // minOccurs="0"

	DutyTaxFeeType string `xml:"dutyTaxFeeType,omitempty"`  // minOccurs="0"

	MeasurementValue *float64 `xml:"measurementValue,omitempty"`  // minOccurs="0"

	MeasurementSignificance string `xml:"measurementSignificance,omitempty"`  // minOccurs="0"
}

type OtherConvRateDetails struct {

	ConversionType string `xml:"conversionType,omitempty"`  // minOccurs="0"

	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	RateType string `xml:"rateType,omitempty"`  // minOccurs="0"

	PricingAmount *float64 `xml:"pricingAmount,omitempty"`  // minOccurs="0"

	ConvertedValueAmount *float64 `xml:"convertedValueAmount,omitempty"`  // minOccurs="0"

	DutyTaxFeeType string `xml:"dutyTaxFeeType,omitempty"`  // minOccurs="0"

	MeasurementValue *float64 `xml:"measurementValue,omitempty"`  // minOccurs="0"

	MeasurementSignificance string `xml:"measurementSignificance,omitempty"`  // minOccurs="0"
}

type PricingTickInfo struct {

	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`  // minOccurs="0"

	LocationDetails *LocationDetails `xml:"locationDetails,omitempty"`  // minOccurs="0"

	OtherLocationDetails *OtherLocationDetails `xml:"otherLocationDetails,omitempty"`  // minOccurs="0"

	IdNumber string `xml:"idNumber,omitempty"`  // minOccurs="0"
}

type ProductDateTimeDetails struct {

	DepartureDate string `xml:"departureDate,omitempty"`  // minOccurs="0"

	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"
}

type LocationDetails struct {

	City string `xml:"city,omitempty"`  // minOccurs="0"

	Country string `xml:"country,omitempty"`  // minOccurs="0"
}

type OtherLocationDetails struct {

	City string `xml:"city,omitempty"`  // minOccurs="0"

	Country string `xml:"country,omitempty"`  // minOccurs="0"
}

type MultiCorporate struct {

	CorporateId []*CorporateId `xml:"corporateId"`  // maxOccurs="20"
}

type CorporateId struct {

	CorporateQualifier string `xml:"corporateQualifier"`

	Identity []string `xml:"identity"`  // maxOccurs="9"
}

type ItemNumber struct {

	ItemNumberDetails []*ItemNumberDetails `xml:"itemNumberDetails"`  // maxOccurs="99"
}

type ItemNumberDetails struct {

	Number string `xml:"number,omitempty"`  // minOccurs="0"

	Type string `xml:"type,omitempty"`  // minOccurs="0"
}

type DateOfFlight struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type DateAndTimeDetails struct {

	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	Date *float64 `xml:"date,omitempty"`  // minOccurs="0"
}

type FlightQualification struct {

	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FareCategories struct {

	FareType string `xml:"fareType"`

	OtherFareType []string `xml:"otherFareType,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type FareDetails struct {

	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	Rate *float64 `xml:"rate,omitempty"`  // minOccurs="0"

	Country string `xml:"country,omitempty"`  // minOccurs="0"

	FareCategory string `xml:"fareCategory,omitempty"`  // minOccurs="0"
}

type AdditionalFareDetails struct {

	RateClass string `xml:"rateClass,omitempty"`  // minOccurs="0"

	CommodityCategory string `xml:"commodityCategory,omitempty"`  // minOccurs="0"

	PricingGroup string `xml:"pricingGroup,omitempty"`  // minOccurs="0"

	SecondRateClass []string `xml:"secondRateClass,omitempty"`  // minOccurs="0" maxOccurs="29"
}

type DiscountDetails struct {

	FareQualifier string `xml:"fareQualifier"`

	RateCategory string `xml:"rateCategory,omitempty"`  // minOccurs="0"

	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	Percentage *float64 `xml:"percentage,omitempty"`  // minOccurs="0"
}

type TransportInformation struct {

	TransportService *TransportService `xml:"transportService"`

	AvailCabinConf *AvailCabinConf `xml:"availCabinConf,omitempty"`  // minOccurs="0"

	RoutingInfo *RoutingInfo `xml:"routingInfo,omitempty"`  // minOccurs="0"

	SelectionDetail *SelectionDetail `xml:"selectionDetail,omitempty"`  // minOccurs="0"
}

type TransportService struct {

	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type CompanyIdentification struct {

	MarketingCompany string `xml:"marketingCompany,omitempty"`  // minOccurs="0"

	Operatingcompany string `xml:"operatingcompany,omitempty"`  // minOccurs="0"

	OtherCompany string `xml:"otherCompany,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetails struct {

	FlightNumber string `xml:"flightNumber"`

	OperationalSuffix string `xml:"operationalSuffix,omitempty"`  // minOccurs="0"
}

type AvailCabinConf struct {

	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"`  // minOccurs="0"

	BookingClassDetails []*BookingClassDetails `xml:"bookingClassDetails"`  // maxOccurs="26"
}

type RoutingInfo struct {

	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type RoutingDetails struct {

	Station string `xml:"station,omitempty"`  // minOccurs="0"

	OtherStation string `xml:"otherStation,omitempty"`  // minOccurs="0"

	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"
}

type SelectionDetail struct {

	SelectionDetails *SelectionDetails `xml:"selectionDetails"`

	SelectionDetailsTwo []*SelectionDetailsTwo `xml:"selectionDetailsTwo,omitempty"`  // minOccurs="0" maxOccurs="98"
}

type SelectionDetails struct {

	Option string `xml:"option"`

	OptionInformation string `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type SelectionDetailsTwo struct {

	Option string `xml:"option"`

	OptionInformation string `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type TripDescription struct {

	OrigDest *OrigDest `xml:"origDest"`

	DateFlightMovement *DateFlightMovement `xml:"dateFlightMovement,omitempty"`  // minOccurs="0"

	Routing []*Routing `xml:"routing,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type OrigDest struct {

	Origin string `xml:"origin,omitempty"`  // minOccurs="0"

	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type DateFlightMovement struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type Routing struct {

	RoutingInfo *RoutingInfo `xml:"routingInfo"`

	TransportService *TransportService `xml:"transportService,omitempty"`  // minOccurs="0"

	SegFareDetails *SegFareDetails `xml:"segFareDetails,omitempty"`  // minOccurs="0"

	PertinentQuantity *PertinentQuantity `xml:"pertinentQuantity,omitempty"`  // minOccurs="0"

	SelectionMakingDetails *SelectionMakingDetails `xml:"selectionMakingDetails,omitempty"`  // minOccurs="0"
}

type SegFareDetails struct {

	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type PertinentQuantity struct {

	QuantityDetails *QuantityDetails `xml:"quantityDetails"`

	OtherQuantityDetails []*OtherQuantityDetails `xml:"otherQuantityDetails,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type QuantityDetails struct {

	Qualifier string `xml:"qualifier"`

	Value float64 `xml:"value"`

	Unit string `xml:"unit,omitempty"`  // minOccurs="0"
}

type OtherQuantityDetails struct {

	Qualifier string `xml:"qualifier"`

	Value float64 `xml:"value"`

	Unit string `xml:"unit,omitempty"`  // minOccurs="0"
}

type SelectionMakingDetails struct {

	SelectionDetails *SelectionDetails `xml:"selectionDetails"`

	SelectionDetailsTwo []*SelectionDetailsTwo `xml:"selectionDetailsTwo,omitempty"`  // minOccurs="0" maxOccurs="98"
}

type PricingInfo struct {

	NumberOfUnits *NumberOfUnits `xml:"numberOfUnits"`

	TicketPricingDate []*TicketPricingDate `xml:"ticketPricingDate,omitempty"`  // minOccurs="0" maxOccurs="3"

	Fare []*Fare `xml:"fare,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type NumberOfUnits struct {

	QuantityDetails *QuantityDetails1 `xml:"quantityDetails"`

	OtherQuantityDetails []*OtherQuantityDetails1 `xml:"otherQuantityDetails,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type QuantityDetails1 struct {

	NumberOfUnit *float64 `xml:"numberOfUnit,omitempty"`  // minOccurs="0"

	UnitQualifier string `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type OtherQuantityDetails1 struct {

	NumberOfUnit *float64 `xml:"numberOfUnit,omitempty"`  // minOccurs="0"

	UnitQualifier string `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type TicketPricingDate struct {

	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`  // minOccurs="0"

	LocationDetails *LocationDetails `xml:"locationDetails,omitempty"`  // minOccurs="0"

	OtherLocationDetails *OtherLocationDetails `xml:"otherLocationDetails,omitempty"`  // minOccurs="0"

	IdNumber string `xml:"idNumber,omitempty"`  // minOccurs="0"
}

type Fare struct {

	DetailsOfFare *DetailsOfFare `xml:"detailsOfFare"`

	FareQualificationDetails []*FareQualificationDetails `xml:"fareQualificationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type DetailsOfFare struct {

	FareTypeGrouping *FareTypeGrouping `xml:"fareTypeGrouping,omitempty"`  // minOccurs="0"
}

type FareTypeGrouping struct {

	PricingGroup []string `xml:"pricingGroup,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type FareQualificationDetails struct {

	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FareRule struct {

	TarifFareRule *TarifFareRule `xml:"tarifFareRule"`

	TravellerIdentification *TravellerIdentification `xml:"travellerIdentification,omitempty"`  // minOccurs="0"

	TravellerDate *TravellerDate `xml:"travellerDate,omitempty"`  // minOccurs="0"
}

type TarifFareRule struct {

	TariffClassId string `xml:"tariffClassId,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	RuleSectionId []string `xml:"ruleSectionId,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type CompanyDetails struct {

	MarketingCompany string `xml:"marketingCompany,omitempty"`  // minOccurs="0"

	Operatingcompany string `xml:"operatingcompany,omitempty"`  // minOccurs="0"

	OtherCompany string `xml:"otherCompany,omitempty"`  // minOccurs="0"
}

type TravellerIdentification struct {

	ReferenceDetails []*ReferenceDetails `xml:"referenceDetails"`  // maxOccurs="99"
}

type ReferenceDetails struct {

	Type string `xml:"type,omitempty"`  // minOccurs="0"

	Value string `xml:"value,omitempty"`  // minOccurs="0"
}

type TravellerDate struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}
