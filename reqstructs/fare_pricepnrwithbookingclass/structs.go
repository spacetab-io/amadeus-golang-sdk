package fare_pricepnrwithbookingclass

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type FarePricePNRWithBookingClass struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A Fare_PricePNRWithBookingClass"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`

	// Passenger segment association of the transaction. Several passengers and/or segments can be priced in the same transaction.
	PaxSegReference *ReferenceInformationTypeI_94605S `xml:"paxSegReference,omitempty"`

	// Information on the pricing override request
	OverrideInformation *CodedAttributeType `xml:"overrideInformation,omitempty"`

	// Booking date override information.
	DateOverride *StructuredDateTimeInformationType `xml:"dateOverride,omitempty"`

	// Validating carrier override information.
	ValidatingCarrier *TransportIdentifierType `xml:"validatingCarrier,omitempty"`

	// Point of sale and point of ticketing override.
	CityOverride *PointOfSaleInformationTypeU `xml:"cityOverride,omitempty"`

	// Currency override information. 1st C661 : currency to use for pricing. 2nd C661 : fare selection currency.
	CurrencyOverride *ConversionRateTypeI `xml:"currencyOverride,omitempty"`

	// Details about taxes specified in the request. The number of repetition is 60, as we can have 60 taxes.
	TaxDetails *DutyTaxFeeDetailsTypeU `xml:"taxDetails,omitempty"`

	DiscountInformation *DiscountInformation `xml:"discountInformation,omitempty"`

	PricingFareBase *PricingFareBase `xml:"pricingFareBase,omitempty"`

	FlightInformation *FlightInformation `xml:"flightInformation,omitempty"`

	OpenSegmentsInformation *OpenSegmentsInformation `xml:"openSegmentsInformation,omitempty"`

	BookingClassSelection *BookingClassSelection `xml:"bookingClassSelection,omitempty"`

	FopInformation *FopInformation `xml:"fopInformation,omitempty"`

	// Commercial Agreements with specific Carrier, Ex: Award publishing Carrier.
	CarrierAgreements *CommercialAgreementsTypeI `xml:"carrierAgreements,omitempty"`

	FrequentFlyerInformation *FrequentFlyerInformation `xml:"frequentFlyerInformation,omitempty"`

	InstantPricingOption *MonetaryInformationType `xml:"instantPricingOption,omitempty"`
}

type DiscountInformation struct {
	// Penalty/discount/Fees details specified in the request.
	PenDisInformation *DiscountAndPenaltyInformationTypeI `xml:"penDisInformation,omitempty"`

	// Passenger or segments qualifiers
	ReferenceQualifier *ReferenceInformationTypeI_94606S `xml:"referenceQualifier,omitempty"`
}

type PricingFareBase struct {
	// Segment details for pricing by fare basis request.
	FareBasisOptions *FareQualifierDetailsTypeI `xml:"fareBasisOptions,omitempty"`

	// Reference of segments associated to a fare basis.  S- indicator for segment
	FareBasisSegReference *ReferenceInformationTypeI_94606S `xml:"fareBasisSegReference,omitempty"`

	// provides the dates for fare basis : Not Valid Before and Not Valid After dates.
	FareBasisDates *StructuredDateTimeInformationType_94602S `xml:"fareBasisDates,omitempty"`

	// In this segment you can specify the passengers
	PassengerTattoo *ReferenceInformationType_94604S `xml:"passengerTattoo,omitempty"`
}

type FlightInformation struct {
	// Override options for itinerary specified in the request (breakpoint, turnaround, ...).
	ItineraryOptions *TravelItineraryInformationTypeI `xml:"itineraryOptions,omitempty"`

	// Reference of segments (tattoo number) associated to the override option : point of turnaround, connexion... 'S' is the only reference qualifier relevant here.
	ItinerarySegReference *ReferenceInformationTypeI_94610S `xml:"itinerarySegReference,omitempty"`
}

type OpenSegmentsInformation struct {
	// Extra segment specified in the request. This segment is not part of the PNR.
	ExtendedItinerary *TravelProductInformationTypeI `xml:"extendedItinerary,omitempty"`

	// Reference of segments (tattoo number) associated to extended itinerary option.  'S' is the only reference qualifier relevant here.
	ExtendedItinerarySegReference *ReferenceInformationTypeI_94610S `xml:"extendedItinerarySegReference,omitempty"`
}

type BookingClassSelection struct {
	// Contains the list of booking classes requested per segment.  This option is valid only for Best buy messages.
	BookingClassList *ProductInformationTypeI `xml:"bookingClassList,omitempty"`

	// The REF segment is used to indicate the segment associated to the booking codes option selected.
	BookCodeSegAsso *ReferenceInformationTypeI `xml:"bookCodeSegAsso,omitempty"`
}

type FopInformation struct {
	// Form of payment
	Fop *FormOfPaymentTypeI `xml:"fop,omitempty"`
}

type FrequentFlyerInformation struct {
	// Used for tier level override
	FlequentFlyerdetails *FrequentTravellerIdentificationCodeType `xml:"flequentFlyerdetails,omitempty"`

	// Pax
	PassengerReference *ReferenceInformationType `xml:"passengerReference,omitempty"`

	// Segment
	FlightReference *ReferenceInfoType `xml:"flightReference,omitempty"`
}

type AdditionalFareQualifierDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A AdditionalFareQualifierDetailsTypeI"`

	// Primary code of the fare basis. This is not a codeset but a free flow text field.
	PrimaryCode formats.AlphaNumericString_Length1To3 `xml:"primaryCode,omitempty"`

	// Fare basis code of the fare basis. This is not a codeset but a free flow text field.
	FareBasisCode formats.AlphaNumericString_Length1To6 `xml:"fareBasisCode,omitempty"`

	// Ticket designator of the fare basis
	TicketDesignator formats.AlphaNumericString_Length1To6 `xml:"ticketDesignator,omitempty"`

	// provides the routing Number.
	DiscTktDesignator formats.AlphaNumericString_Length1To4 `xml:"discTktDesignator,omitempty"`
}

type CodedAttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A CodedAttributeInformationType"`

	// provides the attribute type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType,omitempty"`

	// provides a description for the attribute
	AttributeDescription formats.AlphaNumericString_Length1To100 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A CodedAttributeType"`

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`
}

type CommercialAgreementsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A CommercialAgreementsTypeI"`

	// Company options
	CompanyDetails *CompanyRoleIdentificationTypeI `xml:"companyDetails,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A CompanyIdentificationTypeI"`

	// Code of the airline
	CarrierCode formats.AlphaNumericString_Length1To3 `xml:"carrierCode,omitempty"`
}

type CompanyRoleIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A CompanyRoleIdentificationTypeI"`

	// Option qualifier.
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// company identification
	Company formats.AlphaNumericString_Length2To3 `xml:"company,omitempty"`
}

type ConversionRateDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ConversionRateDetailsTypeI"`

	// Currency of the rate
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	RateType formats.AlphaNumericString_Length1To3 `xml:"rateType,omitempty"`

	// Amount/percentage
	Amount formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`
}

type ConversionRateTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ConversionRateTypeI"`

	// Used for fare override, e.g. /FC-USD
	FirstRateDetail *ConversionRateDetailsTypeI `xml:"firstRateDetail,omitempty"`

	// Used for fare selection, e.g. /FS-USD
	SecondRateDetail *ConversionRateDetailsTypeI `xml:"secondRateDetail,omitempty"`
}

type DiscountAndPenaltyInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A DiscountAndPenaltyInformationTypeI"`

	// Qualify the type of information.  -Penalties are not passenger associated and are pure monetary information. Discount are passenger associated but only discount code is specified. -The type can be for Zapp off discount which is segment associated. -The type of Fee: OB are neither passenger associated nor segment associated.
	InfoQualifier formats.AlphaNumericString_Length1To3 `xml:"infoQualifier,omitempty"`

	// Used to specify penalty information
	PenDisData *DiscountPenaltyMonetaryInformationTypeI `xml:"penDisData,omitempty"`
}

type DiscountPenaltyMonetaryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A DiscountPenaltyMonetaryInformationTypeI"`

	// Type of discount/penalty/fee
	PenaltyType formats.AlphaNumericString_Length1To3 `xml:"penaltyType,omitempty"`

	// The amount Type can be a percentage or an amount
	PenaltyQualifier formats.AlphaNumericString_Length1To3 `xml:"penaltyQualifier,omitempty"`

	// specify the value
	PenaltyAmount formats.NumericDecimal_Length1To11 `xml:"penaltyAmount,omitempty"`

	// provides the discount or Fee code
	DiscountCode formats.AlphaNumericString_Length1To11 `xml:"discountCode,omitempty"`

	// penalty currency code
	PenaltyCurrency formats.AlphaNumericString_Length1To3 `xml:"penaltyCurrency,omitempty"`
}

type DutyTaxFeeAccountDetailTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A DutyTaxFeeAccountDetailTypeU"`

	// Iso country of the tax
	IsoCountry formats.AlphaNumericString_Length1To3 `xml:"isoCountry,omitempty"`
}

type DutyTaxFeeDetailTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A DutyTaxFeeDetailTypeU"`

	// Tax override value
	TaxRate formats.NumericDecimal_Length1To9 `xml:"taxRate,omitempty"`

	// Tax value qualifier will be either amount or percentage.
	TaxValueQualifier formats.AlphaNumericString_Length1To1 `xml:"taxValueQualifier,omitempty"`
}

type DutyTaxFeeDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A DutyTaxFeeDetailsTypeU"`

	// Tax data qualifier
	TaxQualifier formats.AlphaNumericString_Length1To3 `xml:"taxQualifier,omitempty"`

	// Tax type identification
	TaxIdentification *DutyTaxFeeTypeDetailsTypeU `xml:"taxIdentification,omitempty"`

	// Type of the tax
	TaxType *DutyTaxFeeAccountDetailTypeU `xml:"taxType,omitempty"`

	// Nature of the tax
	TaxNature formats.AlphaNumericString_Length1To3 `xml:"taxNature,omitempty"`

	// Tax fare data information
	TaxData *DutyTaxFeeDetailTypeU `xml:"taxData,omitempty"`
}

type DutyTaxFeeTypeDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A DutyTaxFeeTypeDetailsTypeU"`

	// Tax type identifier
	TaxIdentifier formats.AlphaNumericString_Length1To3 `xml:"taxIdentifier,omitempty"`
}

type FareDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A FareDetailsTypeI"`
}

type FareQualifierDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A FareQualifierDetailsTypeI"`

	// Type of movement for this segment to take into account by Fare Quote to calculate the fare.
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`

	// This composite contains the block ID called Pax-fare basis association
	BlockPaxSegmentDetails *FareDetailsTypeI `xml:"blockPaxSegmentDetails,omitempty"`

	// Fare basis detail
	FareBasisDetails *AdditionalFareQualifierDetailsTypeI `xml:"fareBasisDetails,omitempty"`
}

type FormOfPaymentDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A FormOfPaymentDetailsTypeI"`

	// Type of Form Of Payment (Credit card, cash, check...).
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// amount to be charged on this form
	ChargedAmount formats.NumericDecimal_Length1To12 `xml:"chargedAmount,omitempty"`

	// It is mandatory if the Form of Payment at pricing was a credit card with a bin number. It may only occur for Credit Card Types.  It is composed of the first 6 bin numbers of the credit card.  Wildcards are not possible.
	CreditCardNumber formats.AlphaNumericString_Length1To20 `xml:"creditCardNumber,omitempty"`
}

type FormOfPaymentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A FormOfPaymentTypeI"`

	// FORM OF PAYMENT DETAILS
	FormOfPayment *FormOfPaymentDetailsTypeI `xml:"formOfPayment,omitempty"`
}

type FrequentTravellerIdentificationCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A FrequentTravellerIdentificationCodeType"`

	// Frequent Traveller Info
	FrequentTravellerDetails *FrequentTravellerIdentificationType `xml:"frequentTravellerDetails,omitempty"`
}

type FrequentTravellerIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A FrequentTravellerIdentificationType"`

	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"`

	// To specify a Tier level
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`
}

type LocationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A LocationTypeI"`

	// code of the city
	CityCode formats.AlphaNumericString_Length1To3 `xml:"cityCode,omitempty"`
}

type LocationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A LocationTypeU"`

	// Code of the city
	CityCode formats.AlphaNumericString_Length1To3 `xml:"cityCode,omitempty"`

	// Qualifier of the city
	CityQualifier formats.AlphaNumericString_Length1To3 `xml:"cityQualifier,omitempty"`
}

type MonetaryInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A MonetaryInformationDetailsType"`

	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To35 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A MonetaryInformationType"`

	// Indicate  - Action type: IPA - Amount - Currency
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails,omitempty"`
}

type PointOfSaleInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A PointOfSaleInformationTypeU"`

	// Detail on city to override
	CityDetail *LocationTypeU `xml:"cityDetail,omitempty"`
}

type ProductDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ProductDetailsTypeI"`

	Designator formats.AlphaNumericString_Length1To17 `xml:"designator,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ProductIdentificationDetailsTypeI"`

	// OPEN or AIR are the two identifications accepted. OPEN means the segment described here is an open segment. AIR means that it is a valid AIR segment
	Identification formats.AlphaNumericString_Length1To6 `xml:"identification,omitempty"`

	// Class of service to use in order to price the extra segment.
	ClassOfService formats.AlphaString_Length1To1 `xml:"classOfService,omitempty"`
}

type ProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ProductInformationTypeI"`

	BookingClassDetails *ProductDetailsTypeI `xml:"bookingClassDetails,omitempty"`
}

type ProductTypeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ProductTypeDetailsTypeI"`

	// Different types of flight
	FlightType formats.AlphaNumericString_Length1To6 `xml:"flightType,omitempty"`

	// Other types of flights as combination might happen
	OtherFlightTypes formats.AlphaNumericString_Length1To6 `xml:"otherFlightTypes,omitempty"`
}

type ReferenceInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferenceInfoType"`

	ReferenceDetails *ReferencingDetailsType_191491C `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferenceInformationTypeI"`

	// The segment reference details
	ReferenceDetails *ReferencingDetailsTypeI_47428C `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationTypeI_94605S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferenceInformationTypeI_94605S"`

	// Passenger/segment/TST/fare reference details
	RefDetails *ReferencingDetailsTypeI_142222C `xml:"refDetails,omitempty"`
}

type ReferenceInformationTypeI_94606S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferenceInformationTypeI_94606S"`

	// Passenger/segment/TST/fare reference details
	RefDetails *ReferencingDetailsTypeI_142223C `xml:"refDetails,omitempty"`
}

type ReferenceInformationTypeI_94610S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferenceInformationTypeI_94610S"`

	// segment reference details
	RefDetails *ReferencingDetailsTypeI_142223C `xml:"refDetails,omitempty"`
}

type ReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferenceInformationType"`

	// Used to convey the passenger tatoo or display number.
	PassengerReference *ReferencingDetailsType `xml:"passengerReference,omitempty"`
}

type ReferenceInformationType_94604S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferenceInformationType_94604S"`

	// Used to convey the passenger tatoo or display number.
	PassengerReference *ReferencingDetailsTypeI `xml:"passengerReference,omitempty"`
}

type ReferencingDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferencingDetailsTypeI"`

	// Passenger tattoo type P.  It is possible to choose Adult passenger or Infant Passenger
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Passenger tattoo
	Value formats.NumericInteger_Length1To5 `xml:"value,omitempty"`
}

type ReferencingDetailsTypeI_142222C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferencingDetailsTypeI_142222C"`

	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsTypeI_142223C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferencingDetailsTypeI_142223C"`

	// Qualifyer of the reference (Pax/Seg)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`

	// Passenger/segment reference number
	RefNumber formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`
}

type ReferencingDetailsTypeI_47428C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferencingDetailsTypeI_47428C"`

	// Only segment reference type selection is accepted.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// segment tattoo number
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type ReferencingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferencingDetailsType"`

	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	Value formats.AlphaNumericString_Length1To60 `xml:"value,omitempty"`
}

type ReferencingDetailsType_191491C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReferencingDetailsType_191491C"`

	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	Value formats.AlphaNumericString_Length1To60 `xml:"value,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReservationControlInformationDetailsTypeI"`

	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A ReservationControlInformationTypeI"`

	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation,omitempty"`
}

type StructuredDateTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A StructuredDateTimeInformationType"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_94602S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A StructuredDateTimeInformationType_94602S"`

	// provides the qualifier for the fare basis dates : Not Valid Before and Not Valid After.
	FareBasisDateQualifier formats.AlphaNumericString_Length1To3 `xml:"fareBasisDateQualifier,omitempty"`

	// provides the Not Valid Before or the Not Valid After Date.
	FareBasisDate *StructuredDateTimeType `xml:"fareBasisDate,omitempty"`
}

type StructuredDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A StructuredDateTimeType"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.NumericInteger_Length1To6 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.NumericInteger_Length1To2 `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.NumericInteger_Length1To2 `xml:"day,omitempty"`
}

type TransportIdentifierType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A TransportIdentifierType"`

	// Information related to validating carrier.
	CarrierInformation *CompanyIdentificationTypeI `xml:"carrierInformation,omitempty"`
}

type TravelItineraryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A TravelItineraryInformationTypeI"`

	// Global routing data for the type of movement.
	GlobalRoute formats.AlphaNumericString_Length1To3 `xml:"globalRoute,omitempty"`

	// Booking class override.
	BookingClass formats.AlphaString_Length1To1 `xml:"bookingClass,omitempty"`

	// Extra information about the flight
	FlightDetails *ProductTypeDetailsTypeI `xml:"flightDetails,omitempty"`
}

type TravelProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A TravelProductInformationTypeI"`

	// departure city code
	DepartureCity *LocationTypeI `xml:"departureCity,omitempty"`

	// arrival city code
	ArrivalCity *LocationTypeI `xml:"arrivalCity,omitempty"`

	// Airline detail information of the extra segment
	AirlineDetail *CompanyIdentificationTypeI `xml:"airlineDetail,omitempty"`

	// Segment detail information
	SegmentDetail *ProductIdentificationDetailsTypeI `xml:"segmentDetail,omitempty"`

	// Ticketing status for this segment. Relevant only in case of reply
	TicketingStatus formats.AlphaNumericString_Length1To2 `xml:"ticketingStatus,omitempty"`
}
