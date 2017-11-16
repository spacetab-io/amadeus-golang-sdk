package fare_pricepnrwithbookingclass_old

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type FarePricePNRWithBookingClassOld struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_12_4_1A Fare_PricePNRWithBookingClass"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`  // minOccurs="0"

	// Passenger segment association of the transaction. Several passengers and/or segments can be priced in the same transaction.
	PaxSegReference *ReferenceInformationTypeI_94605S `xml:"paxSegReference,omitempty"`  // minOccurs="0"

	// Information on the pricing override request
	OverrideInformation *CodedAttributeType `xml:"overrideInformation"`

	// Booking date override information.
	DateOverride *StructuredDateTimeInformationType `xml:"dateOverride,omitempty"`  // minOccurs="0"

	// Validating carrier override information.
	ValidatingCarrier *TransportIdentifierType `xml:"validatingCarrier,omitempty"`  // minOccurs="0"

	// Point of sale and point of ticketing override.
	CityOverride *PointOfSaleInformationTypeU `xml:"cityOverride,omitempty"`  // minOccurs="0"

	// Currency override information. 1st C661 : currency to use for pricing. 2nd C661 : fare selection currency.
	CurrencyOverride *ConversionRateTypeI `xml:"currencyOverride,omitempty"`  // minOccurs="0"

	// Details about taxes specified in the request. The number of repetition is 60, as we can have 60 taxes.
	TaxDetails []*DutyTaxFeeDetailsTypeU `xml:"taxDetails,omitempty"`  // minOccurs="0" maxOccurs="60"

	DiscountInformation []*DiscountInformation `xml:"discountInformation,omitempty"`  // minOccurs="0" maxOccurs="99"

	PricingFareBase []*PricingFareBase `xml:"pricingFareBase,omitempty"`  // minOccurs="0" maxOccurs="1584"

	FlightInformation []*FlightInformation `xml:"flightInformation,omitempty"`  // minOccurs="0" maxOccurs="16"

	OpenSegmentsInformation []*OpenSegmentsInformation `xml:"openSegmentsInformation,omitempty"`  // minOccurs="0" maxOccurs="4"

	BookingClassSelection []*BookingClassSelection `xml:"bookingClassSelection,omitempty"`  // minOccurs="0" maxOccurs="16"

	FopInformation *FopInformation `xml:"fopInformation,omitempty"`  // minOccurs="0"

	// Commercial Agreements with specific Carrier, Ex: Award publishing Carrier.
	CarrierAgreements *CommercialAgreementsTypeI `xml:"carrierAgreements,omitempty"`  // minOccurs="0"

	FrequentFlyerInformation []*FrequentFlyerInformation `xml:"frequentFlyerInformation,omitempty"`  // minOccurs="0" maxOccurs="99"

	InstantPricingOption *MonetaryInformationType `xml:"instantPricingOption,omitempty"`  // minOccurs="0"
}

type DiscountInformation struct {

	// Penalty/discount/Fees details specified in the request.
	PenDisInformation *DiscountAndPenaltyInformationTypeI `xml:"penDisInformation"`

	// Passenger or segments qualifiers
	ReferenceQualifier *ReferenceInformationTypeI_94606S `xml:"referenceQualifier,omitempty"`  // minOccurs="0"
}

type PricingFareBase struct {

	// Segment details for pricing by fare basis request.
	FareBasisOptions *FareQualifierDetailsTypeI `xml:"fareBasisOptions"`

	// Reference of segments associated to a fare basis.  S- indicator for segment
	FareBasisSegReference *ReferenceInformationTypeI_94606S `xml:"fareBasisSegReference,omitempty"`  // minOccurs="0"

	// provides the dates for fare basis : Not Valid Before and Not Valid After dates.
	FareBasisDates []*StructuredDateTimeInformationType_94602S `xml:"fareBasisDates,omitempty"`  // minOccurs="0" maxOccurs="2"

	// In this segment you can specify the passengers
	PassengerTattoo *ReferenceInformationType_94604S `xml:"passengerTattoo,omitempty"`  // minOccurs="0"
}

type FlightInformation struct {

	// Override options for itinerary specified in the request (breakpoint, turnaround, ...).
	ItineraryOptions *TravelItineraryInformationTypeI `xml:"itineraryOptions"`

	// Reference of segments (tattoo number) associated to the override option : point of turnaround, connexion... 'S' is the only reference qualifier relevant here.
	ItinerarySegReference *ReferenceInformationTypeI_94610S `xml:"itinerarySegReference,omitempty"`  // minOccurs="0"
}

type OpenSegmentsInformation struct {

	// Extra segment specified in the request. This segment is not part of the PNR.
	ExtendedItinerary *TravelProductInformationTypeI `xml:"extendedItinerary"`

	// Reference of segments (tattoo number) associated to extended itinerary option.  'S' is the only reference qualifier relevant here.
	ExtendedItinerarySegReference *ReferenceInformationTypeI_94610S `xml:"extendedItinerarySegReference,omitempty"`  // minOccurs="0"
}

type BookingClassSelection struct {

	// Contains the list of booking classes requested per segment.  This option is valid only for Best buy messages.
	BookingClassList *ProductInformationTypeI `xml:"bookingClassList"`

	// The REF segment is used to indicate the segment associated to the booking codes option selected.
	BookCodeSegAsso *ReferenceInformationTypeI `xml:"bookCodeSegAsso"`
}

type FopInformation struct {

	// Form of payment
	Fop *FormOfPaymentTypeI `xml:"fop"`
}

type FrequentFlyerInformation struct {

	// Used for tier level override
	FlequentFlyerdetails *FrequentTravellerIdentificationCodeType `xml:"flequentFlyerdetails"`

	// Pax
	PassengerReference *ReferenceInformationType `xml:"passengerReference,omitempty"`  // minOccurs="0"

	// Segment
	FlightReference *ReferenceInfoType `xml:"flightReference,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type AdditionalFareQualifierDetailsTypeI struct {

	// Primary code of the fare basis. This is not a codeset but a free flow text field.
	PrimaryCode formats.AlphaNumericString_Length1To3 `xml:"primaryCode"`

	// Fare basis code of the fare basis. This is not a codeset but a free flow text field.
	FareBasisCode formats.AlphaNumericString_Length1To6 `xml:"fareBasisCode,omitempty"`  // minOccurs="0"

	// Ticket designator of the fare basis
	TicketDesignator formats.AlphaNumericString_Length1To6 `xml:"ticketDesignator,omitempty"`  // minOccurs="0"

	// provides the routing Number.
	DiscTktDesignator formats.AlphaNumericString_Length1To4 `xml:"discTktDesignator,omitempty"`  // minOccurs="0"
}

type CodedAttributeInformationType struct {

	// provides the attribute type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType"`

	// provides a description for the attribute
	AttributeDescription formats.AlphaNumericString_Length1To100 `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeType struct {

	// provides details for the Attribute
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"`  // maxOccurs="30"
}

type CommercialAgreementsTypeI struct {

	// Company options
	CompanyDetails []*CompanyRoleIdentificationTypeI `xml:"companyDetails"`  // maxOccurs="99"
}

type CompanyIdentificationTypeI struct {

	// Code of the airline
	CarrierCode formats.AlphaNumericString_Length1To3 `xml:"carrierCode,omitempty"`  // minOccurs="0"
}

type CompanyRoleIdentificationTypeI struct {

	// Option qualifier.
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier"`

	// company identification
	Company formats.AlphaNumericString_Length2To3 `xml:"company,omitempty"`  // minOccurs="0"
}

type ConversionRateDetailsTypeI struct {

	// Currency of the rate
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`  // minOccurs="0"

	RateType formats.AlphaNumericString_Length1To3 `xml:"rateType,omitempty"`  // minOccurs="0"

	// Amount/percentage
	Amount *formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`  // minOccurs="0"
}

type ConversionRateTypeI struct {

	// Used for fare override, e.g. /FC-USD
	FirstRateDetail *ConversionRateDetailsTypeI `xml:"firstRateDetail"`

	// Used for fare selection, e.g. /FS-USD
	SecondRateDetail *ConversionRateDetailsTypeI `xml:"secondRateDetail,omitempty"`  // minOccurs="0"
}

type DiscountAndPenaltyInformationTypeI struct {

	// Qualify the type of information.  -Penalties are not passenger associated and are pure monetary information. Discount are passenger associated but only discount code is specified. -The type can be for Zapp off discount which is segment associated. -The type of Fee: OB are neither passenger associated nor segment associated.
	InfoQualifier formats.AlphaNumericString_Length1To3 `xml:"infoQualifier,omitempty"`  // minOccurs="0"

	// Used to specify penalty information
	PenDisData []*DiscountPenaltyMonetaryInformationTypeI `xml:"penDisData,omitempty"`  // minOccurs="0" maxOccurs="16"
}

type DiscountPenaltyMonetaryInformationTypeI struct {

	// Type of discount/penalty/fee
	PenaltyType formats.AlphaNumericString_Length1To3 `xml:"penaltyType,omitempty"`  // minOccurs="0"

	// The amount Type can be a percentage or an amount
	PenaltyQualifier formats.AlphaNumericString_Length1To3 `xml:"penaltyQualifier,omitempty"`  // minOccurs="0"

	// specify the value
	PenaltyAmount *formats.NumericDecimal_Length1To11 `xml:"penaltyAmount,omitempty"`  // minOccurs="0"

	// provides the discount or Fee code
	DiscountCode formats.AlphaNumericString_Length1To11 `xml:"discountCode,omitempty"`  // minOccurs="0"

	// penalty currency code
	PenaltyCurrency formats.AlphaNumericString_Length1To3 `xml:"penaltyCurrency,omitempty"`  // minOccurs="0"
}

type DutyTaxFeeAccountDetailTypeU struct {

	// Iso country of the tax
	IsoCountry formats.AlphaNumericString_Length1To3 `xml:"isoCountry"`
}

type DutyTaxFeeDetailTypeU struct {

	// Tax override value
	TaxRate *formats.NumericDecimal_Length1To9 `xml:"taxRate,omitempty"`  // minOccurs="0"

	// Tax value qualifier will be either amount or percentage.
	TaxValueQualifier formats.AlphaNumericString_Length1To1 `xml:"taxValueQualifier,omitempty"`  // minOccurs="0"
}

type DutyTaxFeeDetailsTypeU struct {

	// Tax data qualifier
	TaxQualifier formats.AlphaNumericString_Length1To3 `xml:"taxQualifier"`

	// Tax type identification
	TaxIdentification *DutyTaxFeeTypeDetailsTypeU `xml:"taxIdentification"`

	// Type of the tax
	TaxType *DutyTaxFeeAccountDetailTypeU `xml:"taxType,omitempty"`  // minOccurs="0"

	// Nature of the tax
	TaxNature formats.AlphaNumericString_Length1To3 `xml:"taxNature,omitempty"`  // minOccurs="0"

	// Tax fare data information
	TaxData *DutyTaxFeeDetailTypeU `xml:"taxData,omitempty"`  // minOccurs="0"
}

type DutyTaxFeeTypeDetailsTypeU struct {

	// Tax type identifier
	TaxIdentifier formats.AlphaNumericString_Length1To3 `xml:"taxIdentifier,omitempty"`  // minOccurs="0"
}

type FareDetailsTypeI struct {
}

type FareQualifierDetailsTypeI struct {

	// Type of movement for this segment to take into account by Fare Quote to calculate the fare.
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`  // minOccurs="0"

	// This composite contains the block ID called Pax-fare basis association
	BlockPaxSegmentDetails *FareDetailsTypeI `xml:"blockPaxSegmentDetails,omitempty"`  // minOccurs="0"

	// Fare basis detail
	FareBasisDetails *AdditionalFareQualifierDetailsTypeI `xml:"fareBasisDetails"`
}

type FormOfPaymentDetailsTypeI struct {

	// Type of Form Of Payment (Credit card, cash, check...).
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`  // minOccurs="0"

	// amount to be charged on this form
	ChargedAmount *formats.NumericDecimal_Length1To12 `xml:"chargedAmount,omitempty"`  // minOccurs="0"

	// It is mandatory if the Form of Payment at pricing was a credit card with a bin number. It may only occur for Credit Card Types.  It is composed of the first 6 bin numbers of the credit card.  Wildcards are not possible.
	CreditCardNumber formats.AlphaNumericString_Length1To20 `xml:"creditCardNumber,omitempty"`  // minOccurs="0"
}

type FormOfPaymentTypeI struct {

	// FORM OF PAYMENT DETAILS
	FormOfPayment []*FormOfPaymentDetailsTypeI `xml:"formOfPayment,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type FrequentTravellerIdentificationCodeType struct {

	// Frequent Traveller Info
	FrequentTravellerDetails *FrequentTravellerIdentificationType `xml:"frequentTravellerDetails"`
}

type FrequentTravellerIdentificationType struct {

	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"`  // minOccurs="0"

	// To specify a Tier level
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`  // minOccurs="0"
}

type LocationTypeI struct {

	// code of the city
	CityCode formats.AlphaNumericString_Length1To3 `xml:"cityCode,omitempty"`  // minOccurs="0"
}

type LocationTypeU struct {

	// Code of the city
	CityCode formats.AlphaNumericString_Length1To3 `xml:"cityCode,omitempty"`  // minOccurs="0"

	// Qualifier of the city
	CityQualifier formats.AlphaNumericString_Length1To3 `xml:"cityQualifier,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsType struct {

	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount *formats.NumericDecimal_Length1To35 `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationType struct {

	// Indicate  - Action type: IPA - Amount - Currency
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`
}

type PointOfSaleInformationTypeU struct {

	// Detail on city to override
	CityDetail []*LocationTypeU `xml:"cityDetail,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type ProductDetailsTypeI struct {

	Designator formats.AlphaNumericString_Length1To17 `xml:"designator"`
}

type ProductIdentificationDetailsTypeI struct {

	// OPEN or AIR are the two identifications accepted. OPEN means the segment described here is an open segment. AIR means that it is a valid AIR segment
	Identification formats.AlphaNumericString_Length1To6 `xml:"identification"`

	// Class of service to use in order to price the extra segment.
	ClassOfService formats.AlphaString_Length1To1 `xml:"classOfService,omitempty"`  // minOccurs="0"
}

type ProductInformationTypeI struct {

	BookingClassDetails []*ProductDetailsTypeI `xml:"bookingClassDetails,omitempty"`  // minOccurs="0" maxOccurs="26"
}

type ProductTypeDetailsTypeI struct {

	// Different types of flight
	FlightType formats.AlphaNumericString_Length1To6 `xml:"flightType"`

	// Other types of flights as combination might happen
	OtherFlightTypes []formats.AlphaNumericString_Length1To6 `xml:"otherFlightTypes,omitempty"`  // minOccurs="0" maxOccurs="8"
}

type ReferenceInfoType struct {

	ReferenceDetails []*ReferencingDetailsType_191491C `xml:"referenceDetails,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type ReferenceInformationType struct {

	// Used to convey the passenger tatoo or display number.
	PassengerReference []*ReferencingDetailsType `xml:"passengerReference"`  // maxOccurs="99"
}

type ReferenceInformationTypeI struct {

	// The segment reference details
	ReferenceDetails *ReferencingDetailsTypeI_47428C `xml:"referenceDetails"`
}

type ReferenceInformationTypeI_94605S struct {

	// Passenger/segment/TST/fare reference details
	RefDetails []*ReferencingDetailsTypeI_142222C `xml:"refDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferenceInformationTypeI_94606S struct {

	// Passenger/segment/TST/fare reference details
	RefDetails []*ReferencingDetailsTypeI_142223C `xml:"refDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferenceInformationTypeI_94610S struct {

	// segment reference details
	RefDetails *ReferencingDetailsTypeI_142223C `xml:"refDetails,omitempty"`  // minOccurs="0"
}

type ReferenceInformationType_94604S struct {

	// Used to convey the passenger tatoo or display number.
	PassengerReference []*ReferencingDetailsTypeI `xml:"passengerReference"`  // maxOccurs="99"
}

type ReferencingDetailsType struct {

	Type formats.AlphaNumericString_Length1To10 `xml:"type"`

	Value formats.AlphaNumericString_Length1To60 `xml:"value"`
}

type ReferencingDetailsTypeI struct {

	// Passenger tattoo type P.  It is possible to choose Adult passenger or Infant Passenger
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// Passenger tattoo
	Value formats.NumericInteger_Length1To5 `xml:"value"`
}

type ReferencingDetailsTypeI_142222C struct {

	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber *formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`  // minOccurs="0"
}

type ReferencingDetailsTypeI_142223C struct {

	// Qualifyer of the reference (Pax/Seg)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier"`

	// Passenger/segment reference number
	RefNumber formats.NumericInteger_Length1To5 `xml:"refNumber"`
}

type ReferencingDetailsTypeI_47428C struct {

	// Only segment reference type selection is accepted.
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// segment tattoo number
	Value formats.AlphaNumericString_Length1To35 `xml:"value"`
}

type ReferencingDetailsType_191491C struct {

	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`  // minOccurs="0"

	Value formats.AlphaNumericString_Length1To60 `xml:"value,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI struct {

	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {

	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}

type StructuredDateTimeInformationType struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeInformationType_94602S struct {

	// provides the qualifier for the fare basis dates : Not Valid Before and Not Valid After.
	FareBasisDateQualifier formats.AlphaNumericString_Length1To3 `xml:"fareBasisDateQualifier"`

	// provides the Not Valid Before or the Not Valid After Date.
	FareBasisDate *StructuredDateTimeType `xml:"fareBasisDate"`
}

type StructuredDateTimeType struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.NumericInteger_Length1To6 `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month formats.NumericInteger_Length1To2 `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day formats.NumericInteger_Length1To2 `xml:"day"`
}

type TransportIdentifierType struct {

	// Information related to validating carrier.
	CarrierInformation *CompanyIdentificationTypeI `xml:"carrierInformation,omitempty"`  // minOccurs="0"
}

type TravelItineraryInformationTypeI struct {

	// Global routing data for the type of movement.
	GlobalRoute formats.AlphaNumericString_Length1To3 `xml:"globalRoute,omitempty"`  // minOccurs="0"

	// Booking class override.
	BookingClass formats.AlphaString_Length1To1 `xml:"bookingClass,omitempty"`  // minOccurs="0"

	// Extra information about the flight
	FlightDetails *ProductTypeDetailsTypeI `xml:"flightDetails,omitempty"`  // minOccurs="0"
}

type TravelProductInformationTypeI struct {

	// departure city code
	DepartureCity *LocationTypeI `xml:"departureCity,omitempty"`  // minOccurs="0"

	// arrival city code
	ArrivalCity *LocationTypeI `xml:"arrivalCity,omitempty"`  // minOccurs="0"

	// Airline detail information of the extra segment
	AirlineDetail *CompanyIdentificationTypeI `xml:"airlineDetail,omitempty"`  // minOccurs="0"

	// Segment detail information
	SegmentDetail *ProductIdentificationDetailsTypeI `xml:"segmentDetail,omitempty"`  // minOccurs="0"

	// Ticketing status for this segment. Relevant only in case of reply
	TicketingStatus formats.AlphaNumericString_Length1To2 `xml:"ticketingStatus,omitempty"`  // minOccurs="0"
}
