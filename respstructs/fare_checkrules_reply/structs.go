package fare_checkrules_reply

//import "encoding/xml"

type FareCheckRulesReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FARQNR_07_1_1A Fare_CheckRulesReply"`

	TransactionType *TransactionType `xml:"transactionType"`

	StatusInfo *StatusInfo `xml:"statusInfo,omitempty"`  // minOccurs="0"

	FareRouteInfo *FareRouteInfo `xml:"fareRouteInfo,omitempty"`  // minOccurs="0"

	InfoText []*InfoText `xml:"infoText,omitempty"`  // minOccurs="0" maxOccurs="999"

	ErrorInfo *ErrorInfo `xml:"errorInfo,omitempty"`  // minOccurs="0"

	TariffInfo []*TariffInfo `xml:"tariffInfo,omitempty"`  // minOccurs="0" maxOccurs="999"

	FlightDetails []*FlightDetails `xml:"flightDetails,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type TransactionType struct {

	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`  // minOccurs="0"
}

type MessageFunctionDetails struct {

	// Format limitations: an..3
	MessageFunction string `xml:"messageFunction,omitempty"`  // minOccurs="0"
}

type StatusInfo struct {

	StatusDetails *StatusDetails `xml:"statusDetails"`

	OtherDetails []*OtherDetails `xml:"otherDetails,omitempty"`  // minOccurs="0" maxOccurs="98"
}

type StatusDetails struct {

	// Format limitations: an..3
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"
}

type OtherDetails struct {

	// Format limitations: an..3
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"
}

type FareRouteInfo struct {

	// Format limitations: an..7
	DayOfWeek string `xml:"dayOfWeek,omitempty"`  // minOccurs="0"

	FareQualifierDetails *FareQualifierDetails `xml:"fareQualifierDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	IdentificationNumber string `xml:"identificationNumber,omitempty"`  // minOccurs="0"

	ValidityPeriod *ValidityPeriod `xml:"validityPeriod,omitempty"`  // minOccurs="0"
}

type FareQualifierDetails struct {

	// Format limitations: an..3
	FareQualifier []string `xml:"fareQualifier,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type ValidityPeriod struct {

	// Format limitations: n..6
	FirstDate *float64 `xml:"firstDate,omitempty"`  // minOccurs="0"

	// Format limitations: n..6
	SecondDate *float64 `xml:"secondDate,omitempty"`  // minOccurs="0"
}

type InfoText struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FreeTextQualification struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..4
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"
}

type ErrorInfo struct {

	RejectErrorCode *RejectErrorCode `xml:"rejectErrorCode"`

	ErrorFreeText *ErrorFreeText `xml:"errorFreeText,omitempty"`  // minOccurs="0"
}

type RejectErrorCode struct {

	ErrorDetails *ErrorDetails `xml:"errorDetails"`
}

type ErrorDetails struct {

	// Format limitations: an..3
	ErrorCode string `xml:"errorCode"`
}

type ErrorFreeText struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type TariffInfo struct {

	FareRuleInfo *FareRuleInfo `xml:"fareRuleInfo"`

	FareRuleText []*FareRuleText `xml:"fareRuleText,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type FareRuleInfo struct {

	// Format limitations: an..9
	RuleSectionLocalId string `xml:"ruleSectionLocalId,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..7
	RuleCategoryCode string `xml:"ruleCategoryCode,omitempty"`  // minOccurs="0"
}

type CompanyDetails struct {

	// Format limitations: an..3
	MarketingCompany string `xml:"marketingCompany,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Operatingcompany string `xml:"operatingcompany,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	OtherCompany string `xml:"otherCompany,omitempty"`  // minOccurs="0"
}

type FareRuleText struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FlightDetails struct {

	NbOfSegments *NbOfSegments `xml:"nbOfSegments"`

	AmountConversion *AmountConversion `xml:"amountConversion,omitempty"`  // minOccurs="0"

	QuantityValue *QuantityValue `xml:"quantityValue,omitempty"`  // minOccurs="0"

	PricingAndDateInfo *PricingAndDateInfo `xml:"pricingAndDateInfo,omitempty"`  // minOccurs="0"

	QualificationFareDetails []*QualificationFareDetails `xml:"qualificationFareDetails,omitempty"`  // minOccurs="0" maxOccurs="99"

	TransportService []*TransportService `xml:"transportService,omitempty"`  // minOccurs="0" maxOccurs="4"

	FlightErrorCode []*FlightErrorCode `xml:"flightErrorCode,omitempty"`  // minOccurs="0" maxOccurs="999"

	ProductInfo []*ProductInfo `xml:"productInfo,omitempty"`  // minOccurs="0" maxOccurs="99"

	PriceInfo []*PriceInfo `xml:"priceInfo,omitempty"`  // minOccurs="0" maxOccurs="99"

	FareDetailInfo []*FareDetailInfo `xml:"fareDetailInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	OdiGrp []*OdiGrp `xml:"odiGrp,omitempty"`  // minOccurs="0" maxOccurs="999"

	TravellerGrp []*TravellerGrp `xml:"travellerGrp,omitempty"`  // minOccurs="0" maxOccurs="99"

	FareRouteGrp []*FareRouteGrp `xml:"fareRouteGrp,omitempty"`  // minOccurs="0" maxOccurs="99"

	ItemGrp []*ItemGrp `xml:"itemGrp,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type NbOfSegments struct {

	SegmentControlDetails []*SegmentControlDetails `xml:"segmentControlDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type SegmentControlDetails struct {

	// Format limitations: n..15
	Quantity *float64 `xml:"quantity,omitempty"`  // minOccurs="0"

	// Format limitations: n..15
	NumberOfUnits *float64 `xml:"numberOfUnits,omitempty"`  // minOccurs="0"

	// Format limitations: n..15
	TotalNumberOfItems *float64 `xml:"totalNumberOfItems,omitempty"`  // minOccurs="0"
}

type AmountConversion struct {

	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`

	OtherConversionRateDetails []*OtherConversionRateDetails `xml:"otherConversionRateDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type ConversionRateDetails struct {

	// Format limitations: an..3
	ConversionType string `xml:"conversionType,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	RateType string `xml:"rateType,omitempty"`  // minOccurs="0"

	// Format limitations: n..18
	PricingAmount *float64 `xml:"pricingAmount,omitempty"`  // minOccurs="0"

	// Format limitations: n..18
	MeasurementValue *float64 `xml:"measurementValue,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	MeasurementSignificance string `xml:"measurementSignificance,omitempty"`  // minOccurs="0"
}

type OtherConversionRateDetails struct {

	// Format limitations: an..3
	ConversionType string `xml:"conversionType,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	RateType string `xml:"rateType,omitempty"`  // minOccurs="0"

	// Format limitations: n..18
	PricingAmount *float64 `xml:"pricingAmount,omitempty"`  // minOccurs="0"

	// Format limitations: n..18
	MeasurementValue *float64 `xml:"measurementValue,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	MeasurementSignificance string `xml:"measurementSignificance,omitempty"`  // minOccurs="0"
}

type QuantityValue struct {

	QuantityDetails []*QuantityDetails `xml:"quantityDetails"`  // maxOccurs="9"
}

type QuantityDetails struct {

	// Format limitations: an..3
	Qualifier string `xml:"qualifier"`

	// Format limitations: n..5
	Value float64 `xml:"value"`

	// Format limitations: an..3
	Unit string `xml:"unit,omitempty"`  // minOccurs="0"
}

type PricingAndDateInfo struct {

	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`  // minOccurs="0"

	LocationDetails *LocationDetails `xml:"locationDetails,omitempty"`  // minOccurs="0"

	OtherLocationDetails *OtherLocationDetails `xml:"otherLocationDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	IdNumber string `xml:"idNumber,omitempty"`  // minOccurs="0"
}

type ProductDateTimeDetails struct {

	// Format limitations: n6
	DepartureDate string `xml:"departureDate,omitempty"`  // minOccurs="0"

	// Format limitations: n6
	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"
}

type LocationDetails struct {

	// Format limitations: an3..5
	City string `xml:"city,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Country string `xml:"country,omitempty"`  // minOccurs="0"
}

type OtherLocationDetails struct {

	// Format limitations: an3..5
	City string `xml:"city,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Country string `xml:"country,omitempty"`  // minOccurs="0"
}

type QualificationFareDetails struct {

	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FareCategories struct {

	// Format limitations: an..3
	FareType []string `xml:"fareType"`  // maxOccurs="9"
}

type FareDetails struct {

	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Format limitations: n..3
	Rate *float64 `xml:"rate,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Country string `xml:"country,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	FareCategory string `xml:"fareCategory,omitempty"`  // minOccurs="0"
}

type AdditionalFareDetails struct {

	// Format limitations: an..10
	RateClass string `xml:"rateClass,omitempty"`  // minOccurs="0"

	// Format limitations: an..13
	CommodityCategory string `xml:"commodityCategory,omitempty"`  // minOccurs="0"

	// Format limitations: an..10
	FareClass []string `xml:"fareClass,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type DiscountDetails struct {

	// Format limitations: an..3
	FareQualifier string `xml:"fareQualifier"`

	// Format limitations: an..35
	RateCategory string `xml:"rateCategory,omitempty"`  // minOccurs="0"

	// Format limitations: n..3
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Format limitations: n..3
	Percentage *float64 `xml:"percentage,omitempty"`  // minOccurs="0"
}

type TransportService struct {

	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type CompanyIdentification struct {

	// Format limitations: an..3
	MarketingCompany string `xml:"marketingCompany,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Operatingcompany string `xml:"operatingcompany,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	OtherCompany string `xml:"otherCompany,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetails struct {

	// Format limitations: an..4
	FlightNumber string `xml:"flightNumber"`

	// Format limitations: an1
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`  // minOccurs="0"
}

type FlightErrorCode struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ProductInfo struct {

	ProductDetails *ProductDetails `xml:"productDetails"`

	ProductErrorCode []*ProductErrorCode `xml:"productErrorCode,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ProductDetails struct {

	// Format limitations: an..3
	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"`  // minOccurs="0"

	BookingClassDetails []*BookingClassDetails `xml:"bookingClassDetails"`  // maxOccurs="26"
}

type BookingClassDetails struct {

	// Format limitations: an1
	Designator string `xml:"designator"`

	// Format limitations: an1
	Option []string `xml:"option,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type ProductErrorCode struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type PriceInfo struct {

	MonetaryRates *MonetaryRates `xml:"monetaryRates"`

	TaxAmount *TaxAmount `xml:"taxAmount,omitempty"`  // minOccurs="0"

	FareTypeInfo []*FareTypeInfo `xml:"fareTypeInfo,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type MonetaryRates struct {

	MonetaryDetails *MonetaryDetails `xml:"monetaryDetails"`

	AmountTwo []*AmountTwo `xml:"amountTwo,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type MonetaryDetails struct {

	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Format limitations: an..18
	Amount string `xml:"amount"`

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type AmountTwo struct {

	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Format limitations: an..18
	Amount string `xml:"amount"`

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type TaxAmount struct {

	TaxDetails []*TaxDetails `xml:"taxDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type TaxDetails struct {

	// Format limitations: an..17
	Rate string `xml:"rate,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	CountryCode string `xml:"countryCode,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Type []string `xml:"type,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FareTypeInfo struct {

	FareDetailQualif *FareDetailQualif `xml:"fareDetailQualif"`

	FlightMovementDate *FlightMovementDate `xml:"flightMovementDate,omitempty"`  // minOccurs="0"

	FaraRulesInfo *FaraRulesInfo `xml:"faraRulesInfo,omitempty"`  // minOccurs="0"

	SelectionMakingDetails *SelectionMakingDetails `xml:"selectionMakingDetails,omitempty"`  // minOccurs="0"

	AmountConvDetails *AmountConvDetails `xml:"amountConvDetails,omitempty"`  // minOccurs="0"
}

type FareDetailQualif struct {

	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FlightMovementDate struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type DateAndTimeDetails struct {

	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Format limitations: an..6
	Date string `xml:"date,omitempty"`  // minOccurs="0"
}

type FaraRulesInfo struct {

	// Format limitations: an..9
	RuleSectionLocalId string `xml:"ruleSectionLocalId,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..7
	RuleCategoryCode string `xml:"ruleCategoryCode,omitempty"`  // minOccurs="0"
}

type SelectionMakingDetails struct {

	SelectionDetails *SelectionDetails `xml:"selectionDetails"`

	SelectionDetailsTwo []*SelectionDetailsTwo `xml:"selectionDetailsTwo,omitempty"`  // minOccurs="0" maxOccurs="98"
}

type SelectionDetails struct {

	// Format limitations: an..3
	Option string `xml:"option"`

	// Format limitations: an..3
	OptionInformation string `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type SelectionDetailsTwo struct {

	// Format limitations: an..3
	Option string `xml:"option"`

	// Format limitations: an..3
	OptionInformation string `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type AmountConvDetails struct {

	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`

	OtherConversionRateDetails []*OtherConversionRateDetails `xml:"otherConversionRateDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type FareDetailInfo struct {

	NbOfUnits *NbOfUnits `xml:"nbOfUnits"`

	PricingPlusDateInfo *PricingPlusDateInfo `xml:"pricingPlusDateInfo,omitempty"`  // minOccurs="0"

	FareDeatilInfo *FareDeatilInfo `xml:"fareDeatilInfo,omitempty"`  // minOccurs="0"
}

type NbOfUnits struct {

	QuantityDetails []*QuantityDetails1 `xml:"quantityDetails"`  // maxOccurs="9"
}

type QuantityDetails1 struct {

	// Format limitations: n..15
	NumberOfUnit *float64 `xml:"numberOfUnit,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	UnitQualifier string `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type PricingPlusDateInfo struct {

	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`  // minOccurs="0"

	LocationDetails *LocationDetails `xml:"locationDetails,omitempty"`  // minOccurs="0"

	OtherLocationDetails *OtherLocationDetails `xml:"otherLocationDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	IdNumber string `xml:"idNumber,omitempty"`  // minOccurs="0"
}

type FareDeatilInfo struct {

	FareTypeGrouping *FareTypeGrouping `xml:"fareTypeGrouping,omitempty"`  // minOccurs="0"
}

type FareTypeGrouping struct {

	// Format limitations: an..6
	PricingGroup []string `xml:"pricingGroup,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type OdiGrp struct {

	OriginDestination *OriginDestination `xml:"originDestination"`

	FlightDateAndTime []*FlightDateAndTime `xml:"flightDateAndTime,omitempty"`  // minOccurs="0" maxOccurs="5"

	FlightErrorText *FlightErrorText `xml:"flightErrorText,omitempty"`  // minOccurs="0"

	MonGrp []*MonGrp `xml:"monGrp,omitempty"`  // minOccurs="0" maxOccurs="99"

	RoutingGrp []*RoutingGrp `xml:"routingGrp,omitempty"`  // minOccurs="0" maxOccurs="5"

	TravelProductGrp []*TravelProductGrp `xml:"travelProductGrp,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type OriginDestination struct {

	// Format limitations: an3..5
	Origin string `xml:"origin,omitempty"`  // minOccurs="0"

	// Format limitations: an3..5
	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type FlightDateAndTime struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FlightErrorText struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type MonGrp struct {

	MonetaryValues *MonetaryValues `xml:"monetaryValues"`

	FareDetailGrp []*FareDetailGrp `xml:"fareDetailGrp,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type MonetaryValues struct {

	MonetaryDetails *MonetaryDetails `xml:"monetaryDetails"`

	AmountTwo []*AmountTwo `xml:"amountTwo,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type FareDetailGrp struct {

	FareQualif *FareQualif `xml:"fareQualif"`

	AmountCvtRate *AmountCvtRate `xml:"amountCvtRate,omitempty"`  // minOccurs="0"
}

type FareQualif struct {

	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type AmountCvtRate struct {

	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`

	OtherConversionRateDetails []*OtherConversionRateDetails `xml:"otherConversionRateDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type RoutingGrp struct {

	RoutingInfo *RoutingInfo `xml:"routingInfo"`

	ServiceTransport *ServiceTransport `xml:"serviceTransport,omitempty"`  // minOccurs="0"

	QualificationOfFare *QualificationOfFare `xml:"qualificationOfFare,omitempty"`  // minOccurs="0"

	PertinentQuantity *PertinentQuantity `xml:"pertinentQuantity,omitempty"`  // minOccurs="0"
}

type RoutingInfo struct {

	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type RoutingDetails struct {

	// Format limitations: an3..5
	Station string `xml:"station,omitempty"`  // minOccurs="0"

	// Format limitations: an3..5
	OtherStation string `xml:"otherStation,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"
}

type ServiceTransport struct {

	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type QualificationOfFare struct {

	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type PertinentQuantity struct {

	QuantityDetails []*QuantityDetails `xml:"quantityDetails"`  // maxOccurs="9"
}

type TravelProductGrp struct {

	TravelProductInfo *TravelProductInfo `xml:"travelProductInfo"`

	RoutingGrp []*RoutingGrp1 `xml:"routingGrp,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type TravelProductInfo struct {

	FlightDate *FlightDate `xml:"flightDate,omitempty"`  // minOccurs="0"

	BoardPointDetails *BoardPointDetails `xml:"boardPointDetails,omitempty"`  // minOccurs="0"

	OffpointDetails *OffpointDetails `xml:"offpointDetails,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	FlightIdentification *FlightIdentification `xml:"flightIdentification,omitempty"`  // minOccurs="0"

	FlightTypeDetails *FlightTypeDetails `xml:"flightTypeDetails,omitempty"`  // minOccurs="0"

	// Format limitations: n..6
	ItemNumber *float64 `xml:"itemNumber,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	SpecialSegment string `xml:"specialSegment,omitempty"`  // minOccurs="0"

	MarriageDetails *MarriageDetails `xml:"marriageDetails,omitempty"`  // minOccurs="0"
}

type FlightDate struct {

	// Format limitations: n6
	DepartureDate string `xml:"departureDate,omitempty"`  // minOccurs="0"

	// Format limitations: n6
	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"
}

type BoardPointDetails struct {

	// Format limitations: an3..5
	TrueLocationId string `xml:"trueLocationId,omitempty"`  // minOccurs="0"

	// Format limitations: an3..5
	TrueLocation string `xml:"trueLocation,omitempty"`  // minOccurs="0"
}

type OffpointDetails struct {

	// Format limitations: an3..5
	TrueLocationId string `xml:"trueLocationId,omitempty"`  // minOccurs="0"

	// Format limitations: an3..5
	TrueLocation string `xml:"trueLocation,omitempty"`  // minOccurs="0"
}

type FlightIdentification struct {

	// Format limitations: an..4
	FlightNumber string `xml:"flightNumber"`

	// Format limitations: an1
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`  // minOccurs="0"
}

type FlightTypeDetails struct {

	// Format limitations: an..6
	FlightIndicator string `xml:"flightIndicator"`

	// Format limitations: an..6
	SecondSequenceNb []string `xml:"secondSequenceNb,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type MarriageDetails struct {

	// Format limitations: an..3
	Relation string `xml:"relation,omitempty"`  // minOccurs="0"

	// Format limitations: n..10
	MarriageIdentifier *float64 `xml:"marriageIdentifier,omitempty"`  // minOccurs="0"

	// Format limitations: n..6
	LineNumber *float64 `xml:"lineNumber,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	OtherRelation string `xml:"otherRelation,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	CarrierCode string `xml:"carrierCode,omitempty"`  // minOccurs="0"
}

type RoutingGrp1 struct {

	RoutingInfo *RoutingInfo `xml:"routingInfo"`

	TransportServiceChange []*TransportServiceChange `xml:"transportServiceChange,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type TransportServiceChange struct {

	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type TravellerGrp struct {

	TravellerIdentRef *TravellerIdentRef `xml:"travellerIdentRef"`

	FareRulesDetails *FareRulesDetails `xml:"fareRulesDetails,omitempty"`  // minOccurs="0"

	FlightMovementDateInfo *FlightMovementDateInfo `xml:"flightMovementDateInfo,omitempty"`  // minOccurs="0"
}

type TravellerIdentRef struct {

	ReferenceDetails []*ReferenceDetails `xml:"referenceDetails"`  // maxOccurs="99"
}

type ReferenceDetails struct {

	// Format limitations: an..3
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Format limitations: an..10
	Value string `xml:"value,omitempty"`  // minOccurs="0"
}

type FareRulesDetails struct {

	// Format limitations: an..9
	TariffClassId string `xml:"tariffClassId,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..7
	RuleSectionId []string `xml:"ruleSectionId,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FlightMovementDateInfo struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FareRouteGrp struct {

	FareRouteInfo *FareRouteInfo `xml:"fareRouteInfo"`

	JourneyGrp []*JourneyGrp `xml:"journeyGrp,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type JourneyGrp struct {

	JourneyOriginAndDestination *JourneyOriginAndDestination `xml:"journeyOriginAndDestination"`

	JourneyProductGrp []*JourneyProductGrp `xml:"journeyProductGrp,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type JourneyOriginAndDestination struct {

	// Format limitations: an3..5
	Origin string `xml:"origin,omitempty"`  // minOccurs="0"

	// Format limitations: an3..5
	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type JourneyProductGrp struct {

	JourneyProductInfo *JourneyProductInfo `xml:"journeyProductInfo"`

	JourneyRoutingGrp []*JourneyRoutingGrp `xml:"journeyRoutingGrp,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type JourneyProductInfo struct {

	FlightDate *FlightDate `xml:"flightDate,omitempty"`  // minOccurs="0"

	BoardPointDetails *BoardPointDetails `xml:"boardPointDetails,omitempty"`  // minOccurs="0"

	OffpointDetails *OffpointDetails `xml:"offpointDetails,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	FlightIdentification *FlightIdentification `xml:"flightIdentification,omitempty"`  // minOccurs="0"

	FlightTypeDetails *FlightTypeDetails `xml:"flightTypeDetails,omitempty"`  // minOccurs="0"

	// Format limitations: n..6
	ItemNumber *float64 `xml:"itemNumber,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	SpecialSegment string `xml:"specialSegment,omitempty"`  // minOccurs="0"

	MarriageDetails *MarriageDetails `xml:"marriageDetails,omitempty"`  // minOccurs="0"
}

type JourneyRoutingGrp struct {

	JourneyRoutingInfo *JourneyRoutingInfo `xml:"journeyRoutingInfo"`

	JourneyTransportService *JourneyTransportService `xml:"journeyTransportService,omitempty"`  // minOccurs="0"
}

type JourneyRoutingInfo struct {

	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type JourneyTransportService struct {

	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ItemGrp struct {

	ItemNb *ItemNb `xml:"itemNb"`

	ProductAvailabilityStatus *ProductAvailabilityStatus `xml:"productAvailabilityStatus,omitempty"`  // minOccurs="0"

	QuantityItem *QuantityItem `xml:"quantityItem,omitempty"`  // minOccurs="0"

	TransportServiceItem []*TransportServiceItem `xml:"transportServiceItem,omitempty"`  // minOccurs="0" maxOccurs="4"

	FreeTextItem []*FreeTextItem `xml:"freeTextItem,omitempty"`  // minOccurs="0" maxOccurs="99"

	FareQualifItem []*FareQualifItem `xml:"fareQualifItem,omitempty"`  // minOccurs="0" maxOccurs="99"

	OriginDestinationGrp []*OriginDestinationGrp `xml:"originDestinationGrp,omitempty"`  // minOccurs="0" maxOccurs="99"

	UnitGrp []*UnitGrp `xml:"unitGrp,omitempty"`  // minOccurs="0" maxOccurs="9"

	MonetaryGrp []*MonetaryGrp `xml:"monetaryGrp,omitempty"`  // minOccurs="0" maxOccurs="99"

	FarerouteGrp []*FarerouteGrp `xml:"farerouteGrp,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ItemNb struct {

	ItemNumberDetails []*ItemNumberDetails `xml:"itemNumberDetails"`  // maxOccurs="99"
}

type ItemNumberDetails struct {

	// Format limitations: an..3
	Number string `xml:"number,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Type string `xml:"type,omitempty"`  // minOccurs="0"
}

type ProductAvailabilityStatus struct {

	// Format limitations: an..3
	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"`  // minOccurs="0"

	BookingClassDetails []*BookingClassDetails `xml:"bookingClassDetails"`  // maxOccurs="26"
}

type QuantityItem struct {

	QuantityDetails []*QuantityDetails `xml:"quantityDetails"`  // maxOccurs="9"
}

type TransportServiceItem struct {

	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FreeTextItem struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FareQualifItem struct {

	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type OriginDestinationGrp struct {

	OriginDestOfJourney *OriginDestOfJourney `xml:"originDestOfJourney"`

	DateForMovements *DateForMovements `xml:"dateForMovements,omitempty"`  // minOccurs="0"

	RoutingForJourney *RoutingForJourney `xml:"routingForJourney,omitempty"`  // minOccurs="0"
}

type OriginDestOfJourney struct {

	// Format limitations: an3..5
	Origin string `xml:"origin,omitempty"`  // minOccurs="0"

	// Format limitations: an3..5
	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type DateForMovements struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type RoutingForJourney struct {

	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type UnitGrp struct {

	NbOfUnits *NbOfUnits `xml:"nbOfUnits"`

	UnitPricingAndDateInfo *UnitPricingAndDateInfo `xml:"unitPricingAndDateInfo,omitempty"`  // minOccurs="0"

	UnitFareDetails *UnitFareDetails `xml:"unitFareDetails,omitempty"`  // minOccurs="0"
}

type UnitPricingAndDateInfo struct {

	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`  // minOccurs="0"

	LocationDetails *LocationDetails `xml:"locationDetails,omitempty"`  // minOccurs="0"

	OtherLocationDetails *OtherLocationDetails `xml:"otherLocationDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	IdNumber string `xml:"idNumber,omitempty"`  // minOccurs="0"
}

type UnitFareDetails struct {

	FareTypeGrouping *FareTypeGrouping `xml:"fareTypeGrouping,omitempty"`  // minOccurs="0"
}

type MonetaryGrp struct {

	MonetaryValues *MonetaryValues `xml:"monetaryValues"`

	MonetFareRuleValues *MonetFareRuleValues `xml:"monetFareRuleValues,omitempty"`  // minOccurs="0"

	MonetTravellerRef *MonetTravellerRef `xml:"monetTravellerRef,omitempty"`  // minOccurs="0"

	MonetTicketPriceAndDate *MonetTicketPriceAndDate `xml:"monetTicketPriceAndDate,omitempty"`  // minOccurs="0"

	MonetTaxValues *MonetTaxValues `xml:"monetTaxValues,omitempty"`  // minOccurs="0"

	QualifGrp []*QualifGrp `xml:"qualifGrp,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type MonetFareRuleValues struct {

	// Format limitations: an..9
	RuleSectionLocalId string `xml:"ruleSectionLocalId,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..7
	RuleCategoryCode string `xml:"ruleCategoryCode,omitempty"`  // minOccurs="0"
}

type MonetTravellerRef struct {

	ReferenceDetails []*ReferenceDetails `xml:"referenceDetails"`  // maxOccurs="99"
}

type MonetTicketPriceAndDate struct {

	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`  // minOccurs="0"

	LocationDetails *LocationDetails `xml:"locationDetails,omitempty"`  // minOccurs="0"

	OtherLocationDetails *OtherLocationDetails `xml:"otherLocationDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	IdNumber string `xml:"idNumber,omitempty"`  // minOccurs="0"
}

type MonetTaxValues struct {

	TaxDetails []*TaxDetails `xml:"taxDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type QualifGrp struct {

	QualificationFare *QualificationFare `xml:"qualificationFare"`

	QualifSelection *QualifSelection `xml:"qualifSelection,omitempty"`  // minOccurs="0"

	QualifDateFlightMovement *QualifDateFlightMovement `xml:"qualifDateFlightMovement,omitempty"`  // minOccurs="0"

	QualifConversionRate *QualifConversionRate `xml:"qualifConversionRate,omitempty"`  // minOccurs="0"
}

type QualificationFare struct {

	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`  // minOccurs="0"

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`  // minOccurs="0"

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`  // minOccurs="0"

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type QualifSelection struct {

	SelectionDetails *SelectionDetails `xml:"selectionDetails"`

	SelectionDetailsTwo []*SelectionDetailsTwo `xml:"selectionDetailsTwo,omitempty"`  // minOccurs="0" maxOccurs="98"
}

type QualifDateFlightMovement struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type QualifConversionRate struct {

	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`

	OtherConversionRateDetails []*OtherConversionRateDetails `xml:"otherConversionRateDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type FarerouteGrp struct {

	InfoForFareRoute *InfoForFareRoute `xml:"infoForFareRoute"`

	FarerouteTransportService []*FarerouteTransportService `xml:"farerouteTransportService,omitempty"`  // minOccurs="0" maxOccurs="99"

	FinalOdiGrp []*FinalOdiGrp `xml:"finalOdiGrp,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type InfoForFareRoute struct {

	// Format limitations: an..7
	DayOfWeek string `xml:"dayOfWeek,omitempty"`  // minOccurs="0"

	FareQualifierDetails *FareQualifierDetails `xml:"fareQualifierDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	IdentificationNumber string `xml:"identificationNumber,omitempty"`  // minOccurs="0"

	ValidityPeriod *ValidityPeriod `xml:"validityPeriod,omitempty"`  // minOccurs="0"
}

type FarerouteTransportService struct {

	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FinalOdiGrp struct {

	FinalOriginDestination *FinalOriginDestination `xml:"finalOriginDestination"`

	LastOdiRoutingInfo *LastOdiRoutingInfo `xml:"lastOdiRoutingInfo,omitempty"`  // minOccurs="0"

	LastOdiDateFlightMovement *LastOdiDateFlightMovement `xml:"lastOdiDateFlightMovement,omitempty"`  // minOccurs="0"
}

type FinalOriginDestination struct {

	// Format limitations: an3..5
	Origin string `xml:"origin,omitempty"`  // minOccurs="0"

	// Format limitations: an3..5
	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type LastOdiRoutingInfo struct {

	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type LastOdiDateFlightMovement struct {

	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}
