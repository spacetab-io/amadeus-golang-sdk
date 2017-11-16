package pnr_reply

//import "encoding/xml"

type PNRReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PNR_Reply"`

	PnrHeader []*PnrHeader `xml:"pnrHeader,omitempty"`  // minOccurs="0" maxOccurs="198"

	// specify the amadeus PNR record locator security information for different pnr elements .
	SecurityInformation *ReservationSecurityInformationType `xml:"securityInformation,omitempty"`  // minOccurs="0"

	// specify a queue , that is  an office , a queue number , a category number ,and a date range number or a date .This segment can be used in an output message.
	QueueInformations *QueueType `xml:"queueInformations,omitempty"`  // minOccurs="0"

	// specify the number of units required
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits,omitempty"`  // minOccurs="0"

	GeneralErrorInfo []*GeneralErrorInfo `xml:"generalErrorInfo,omitempty"`  // minOccurs="0" maxOccurs="99"

	// Conveys PNR special types
	PnrType *CodedAttributeType `xml:"pnrType,omitempty"`  // minOccurs="0"

	// provide free form or coded long text information
	FreetextData []*LongFreeTextType `xml:"freetextData,omitempty"`  // minOccurs="0" maxOccurs="3"

	// Conveys the Header tags of the PNR.
	PnrHeaderTag *StatusType `xml:"pnrHeaderTag,omitempty"`  // minOccurs="0"

	// provide free from or coded text information
	FreeFormText []*InteractiveFreeTextTypeI_136698S `xml:"freeFormText,omitempty"`  // minOccurs="0" maxOccurs="99"

	// specify an amadeus PNR history data
	HistoryData []*PnrHistoryDataType_6022S `xml:"historyData,omitempty"`  // minOccurs="0" maxOccurs="999"

	// Point of Sale Information at SBR level. Owner of the PNR.
	SbrPOSDetails *POSGroupType `xml:"sbrPOSDetails"`

	// Point of Sale Information at SBR level. Creator of the PNR.
	SbrCreationPosDetails *POSGroupType `xml:"sbrCreationPosDetails"`

	// Point of Sale Information at SBR level. Updator of the PNR.
	SbrUpdatorPosDetails *POSGroupType `xml:"sbrUpdatorPosDetails"`

	TechnicalData *TechnicalData `xml:"technicalData,omitempty"`  // minOccurs="0"

	TravellerInfo []*TravellerInfo `xml:"travellerInfo,omitempty"`  // minOccurs="0" maxOccurs="100"

	OriginDestinationDetails []*OriginDestinationDetails `xml:"originDestinationDetails,omitempty"`  // minOccurs="0" maxOccurs="50"

	// specify the segment marriages and connections
	SegmentGroupingInfo []*SegmentGroupingInformationType `xml:"segmentGroupingInfo,omitempty"`  // minOccurs="0" maxOccurs="50"

	DataElementsMaster *DataElementsMaster `xml:"dataElementsMaster,omitempty"`  // minOccurs="0"

	TstData []*TstData `xml:"tstData,omitempty"`  // minOccurs="0" maxOccurs="100"

	DcdData *DcdData `xml:"dcdData,omitempty"`  // minOccurs="0"
}

type PnrHeader struct {

	// specify a reference to a reservation
	ReservationInfo *ReservationControlInformationTypeI_87792S `xml:"reservationInfo"`

	// provide specific reference identification
	ReferenceForRecordLocator *ReferenceInfoType `xml:"referenceForRecordLocator,omitempty"`  // minOccurs="0"
}

type GeneralErrorInfo struct {

	// identify the type of application error  within a message
	MessageErrorInformation *ApplicationErrorInformationType `xml:"messageErrorInformation"`

	// provide free from or coded text information
	MessageErrorText *InteractiveFreeTextTypeI_136698S `xml:"messageErrorText"`
}

type TechnicalData struct {

	// Contains the enveloppe number of the PNR, issue at last EOT.
	EnveloppeNumberData *SequenceDetailsTypeU `xml:"enveloppeNumberData"`

	// CS assumption on last transmitted envelope number
	LastTransmittedEnvelopeNumber *PnrHistoryDataType `xml:"lastTransmittedEnvelopeNumber,omitempty"`  // minOccurs="0"

	// Contains the Purge Date of the PNR at the time of the retrieval
	PurgeDateData *StructuredDateTimeInformationType_27086S `xml:"purgeDateData,omitempty"`  // minOccurs="0"

	// Contains general information relative to the state of the PNR
	GeneralPNRInformation *StatusTypeI_32775S `xml:"generalPNRInformation,omitempty"`  // minOccurs="0"
}

type TravellerInfo struct {

	// specify the PNR segments/elements references and action to apply
	ElementManagementPassenger *ElementManagementSegmentType `xml:"elementManagementPassenger"`

	PassengerData []*PassengerData `xml:"passengerData,omitempty"`  // minOccurs="0" maxOccurs="2"

	NameError *NameError `xml:"nameError,omitempty"`  // minOccurs="0"
}

type PassengerData struct {

	// specify traveler and personal details relating to a traveler
	TravellerInformation *TravellerInformationTypeI_6097S `xml:"travellerInformation"`

	// In case of group, contains the group counters (Booked, Canceled, Split)
	GroupCounters *NumberOfUnitsType_76106S `xml:"groupCounters,omitempty"`  // minOccurs="0"

	// Inf/Child date of birth (DDMMYYYY). For instance 01122007
	DateOfBirth *DateAndTimeInformationType `xml:"dateOfBirth,omitempty"`  // minOccurs="0"
}

type NameError struct {

	// identify the type of application error  within a message
	NameErrorInformation *ApplicationErrorInformationType `xml:"nameErrorInformation"`

	// provide free form or error coded text
	NameErrorFreeText *InteractiveFreeTextTypeI_136698S `xml:"nameErrorFreeText,omitempty"`  // minOccurs="0"
}

type OriginDestinationDetails struct {

	// convey origin and destination of a journey
	OriginDestination *OriginAndDestinationDetailsTypeI_3061S `xml:"originDestination"`

	ItineraryInfo []*ItineraryInfo `xml:"itineraryInfo,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ItineraryInfo struct {

	// specify the PNR segments/elements references and action to apply
	ElementManagementItinerary *ElementManagementSegmentType `xml:"elementManagementItinerary"`

	// specify details related to a product
	TravelProduct *TravelProductInformationTypeI `xml:"travelProduct,omitempty"`  // minOccurs="0"

	// specify the message type and business function
	ItineraryMessageAction *MessageActionDetailsTypeI `xml:"itineraryMessageAction,omitempty"`  // minOccurs="0"

	// specify a reference to a reservation
	ItineraryReservationInfo *ReservationControlInformationTypeI_115879S `xml:"itineraryReservationInfo,omitempty"`  // minOccurs="0"

	// indicate quantity and action required in relation to a product
	RelatedProduct *RelatedProductInformationTypeI `xml:"relatedProduct,omitempty"`  // minOccurs="0"

	// Indicator at element level: - CGB (chargeable indicator)
	ElementsIndicators *StatusType_99946S `xml:"elementsIndicators,omitempty"`  // minOccurs="0"

	// To Convey the Reason for Issuance Code (RFIC) and Reason For Issuance Sub code (RFISC)
	ReasonForIssuanceCode *PricingOrTicketingSubsequentType `xml:"reasonForIssuanceCode,omitempty"`  // minOccurs="0"

	// convey additional information concerning an airline flight
	FlightDetail *AdditionalProductDetailsTypeI `xml:"flightDetail,omitempty"`  // minOccurs="0"

	// Contain the cabin class details
	CabinDetails *CabinDetailsType `xml:"cabinDetails,omitempty"`  // minOccurs="0"

	// specify details for making a selection
	SelectionDetails *SelectionDetailsTypeI_2067S `xml:"selectionDetails,omitempty"`  // minOccurs="0"

	// provide free from or coded text information
	ItineraryfreeFormText []*InteractiveFreeTextTypeI_136698S `xml:"itineraryfreeFormText,omitempty"`  // minOccurs="0" maxOccurs="2"

	// provide free form or coded long text information
	ItineraryFreetext *LongFreeTextType `xml:"itineraryFreetext,omitempty"`  // minOccurs="0"

	// specify the details for hotel transaction
	HotelProduct *HotelProductInformationType `xml:"hotelProduct,omitempty"`  // minOccurs="0"

	// specify the Rate information
	RateInformations *RateInformationType `xml:"rateInformations,omitempty"`  // minOccurs="0"

	// specify one option
	GeneralOption []*GeneralOptionType `xml:"generalOption,omitempty"`  // minOccurs="0" maxOccurs="199"

	// provide the ISO code of a country
	Country *CountryCodeListType `xml:"country,omitempty"`  // minOccurs="0"

	// To specify details relating to tax(es).
	TaxInformation []*TaxTypeI `xml:"taxInformation,omitempty"`  // minOccurs="0" maxOccurs="3"

	// specify operating flight additional information
	CustomerTransactionData *CustomerTransactionDataType `xml:"customerTransactionData,omitempty"`  // minOccurs="0"

	YieldGroup *YieldGroup `xml:"yieldGroup,omitempty"`  // minOccurs="0"

	LegInfo []*LegInfo `xml:"legInfo,omitempty"`  // minOccurs="0" maxOccurs="99"

	// Group used to carry FLIX information associated to the segment
	FlixInfo []*FLIXType `xml:"flixInfo,omitempty"`  // minOccurs="0" maxOccurs="4"

	// Provide Date and Time Details relative to the Itinerary
	DateTimeDetails *DateAndTimeInformationTypeI `xml:"dateTimeDetails,omitempty"`  // minOccurs="0"

	LccTypicalData *LccTypicalData `xml:"lccTypicalData,omitempty"`  // minOccurs="0"

	InsuranceInformation []*InsuranceInformation `xml:"insuranceInformation,omitempty"`  // minOccurs="0" maxOccurs="198"

	// This group handles all Insurance structured details.
	InsuranceDetails *InsuranceBusinessDataType `xml:"insuranceDetails,omitempty"`  // minOccurs="0"

	HotelReservationInfo *HotelReservationInfo `xml:"hotelReservationInfo,omitempty"`  // minOccurs="0"

	TypicalCarData *TypicalCarData `xml:"typicalCarData,omitempty"`  // minOccurs="0"

	// Dedicated to convey the specific cruise sailing information.  It is in fact the reflection of the cruise segment PNR content.  An example of Amadeus PNR with 2 Cruise segments, one confirmed the other cancelled. RP/NYCP02001/NYCP02001 AA/SU 19AUG03/1423Z YYAA63 NYCP02001/0001AA/19AUG03 1.TEST/AAA MR 2 CRU PCL HK1 SJUSJU 18JAN2004-7 DAWN PRINCESS/CF-9VHJ2K/NM-TEST/AAA MR 3 CRU PCL HX1 XYZPKG 18MAY2004-7 CORAL PRINCESS/GP-ABCDEF/CX-654321/NM-TEST/AAA MR  4 AP 0494503434 5 TK OK19AUG/NYCP02001
	TypicalCruiseData *CruiseBusinessDataType `xml:"typicalCruiseData,omitempty"`  // minOccurs="0"

	// Information pertaining to a rail segment
	RailInfo *TrainInformationType `xml:"railInfo,omitempty"`  // minOccurs="0"

	// Marks separation between Rail Group and Tour Group, avoid Ambiguity with CPY segment.
	MarkerRailTour *DummySegmentTypeI `xml:"markerRailTour"`

	// Dedicated to convey the specific Tour information.
	TourInfo *TourInformationType `xml:"tourInfo,omitempty"`  // minOccurs="0"

	// Information pertaining to a Ferry segment
	FerryLegInformation *FerryBookingDescriptionType `xml:"ferryLegInformation,omitempty"`  // minOccurs="0"

	ErrorInfo *ErrorInfo `xml:"errorInfo,omitempty"`  // minOccurs="0"

	// provide specific reference identification
	ReferenceForSegment *ReferenceInfoType `xml:"referenceForSegment,omitempty"`  // minOccurs="0"
}

type YieldGroup struct {

	// It contains some specific air segment's indicators data, not present in yieldDataGroup
	YieldData *ODKeyPerformanceDataType `xml:"yieldData"`

	// Details of the yield data.
	YieldDataGroup *ONDType `xml:"yieldDataGroup,omitempty"`  // minOccurs="0"
}

type LegInfo struct {

	// Marks separation to avoid Ambiguity between TVL segments.
	MarkerLegInfo *FlightSegmentDetailsTypeI `xml:"markerLegInfo"`

	LegTravelProduct *TravelProductInformationTypeI_99362S `xml:"legTravelProduct"`

	InteractiveFreeText []*InteractiveFreeTextTypeI_99363S `xml:"interactiveFreeText,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type LccTypicalData struct {

	// Fare data obtained from D/A availabilty (mapped under PRXP20LCC structure)
	LccFareData *TariffInformationTypeI_28460S `xml:"lccFareData"`

	// Connection key tattoo if any
	LccConnectionData *ItemReferencesAndVersionsType_6550S `xml:"lccConnectionData,omitempty"`  // minOccurs="0"
}

type InsuranceInformation struct {

	// contain data related to each passenger
	InsuranceName *InsuranceNameType `xml:"insuranceName"`

	// To specify monetary information
	InsuranceMonetaryInformation *MonetaryInformationTypeI_1689S `xml:"insuranceMonetaryInformation,omitempty"`  // minOccurs="0"

	// Specify an Amadeus PNR Ticket element
	InsurancePremiumInfo *TravellerInsuranceInformationType `xml:"insurancePremiumInfo,omitempty"`  // minOccurs="0"

	// provide traveller document information
	InsuranceDocInfo *TravellerDocumentInformationTypeU `xml:"insuranceDocInfo,omitempty"`  // minOccurs="0"
}

type HotelReservationInfo struct {

	// This segment is used to convey the hotel property information.
	HotelPropertyInfo *HotelPropertyType `xml:"hotelPropertyInfo"`

	// This segment is used to convey the hotel chain code and name.
	CompanyIdentification *CompanyInformationType `xml:"companyIdentification"`

	// This segment is used to convey the dates.
	RequestedDates *StructuredPeriodInformationType_11026S `xml:"requestedDates"`

	RoomRateDetails *RoomRateDetails `xml:"roomRateDetails,omitempty"`  // minOccurs="0"

	// This segment is used to convey the confirmation number or the cancellation number. control type (9958) is: - 2 for Confirmation reference - X for cancellation reference - O for on request reference
	CancelOrConfirmNbr []*ReservationControlInformationTypeI_115879S `xml:"cancelOrConfirmNbr"`  // maxOccurs="3"

	// indicate the roomstay index in case of groups
	RoomstayIndex *ItemNumberTypeU_33258S `xml:"roomstayIndex,omitempty"`  // minOccurs="0"

	// This segment is used to convey the booking source.
	BookingSource *UserIdentificationType_21014S `xml:"bookingSource"`

	// This segment is used to convey the billable information
	BillableInfo *BillableInformationTypeU `xml:"billableInfo,omitempty"`  // minOccurs="0"

	// This segment is used to convey the customer reference number
	CustomerInfo *ConsumerReferenceInformationTypeI `xml:"customerInfo,omitempty"`  // minOccurs="0"

	// This segment is used to convey the frequent traveler number.
	FrequentTravellerInfo *FrequentTravellerIdentificationCodeType_38226S `xml:"frequentTravellerInfo,omitempty"`  // minOccurs="0"

	GuaranteeOrDeposit *GuaranteeOrDeposit `xml:"guaranteeOrDeposit,omitempty"`  // minOccurs="0"

	// This segment is used to convey additional information which are entered by the agent and stored on the hotel booking.
	TextOptions []*MiscellaneousRemarksType_664S `xml:"textOptions,omitempty"`  // minOccurs="0" maxOccurs="5"

	// This segment is used to convey the saving amount/percentage information.
	SavingAmountInfo *MonetaryInformationTypeI_1689S `xml:"savingAmountInfo,omitempty"`  // minOccurs="0"

	// This segment is used to convey the fax or the E-Mail to receive the written confirmation
	WrittenConfirmationContact *ContactInformationTypeU `xml:"writtenConfirmationContact,omitempty"`  // minOccurs="0"

	// This segment is used to convey the name and address of the party to recieve the written confirmation
	WrittenConfirmationInfo *NameAndAddressBatchTypeU `xml:"writtenConfirmationInfo,omitempty"`  // minOccurs="0"

	// documentInformationDetails
	DocumentInformationDetails *DocumentInformationDetailsTypeI_9936S `xml:"documentInformationDetails,omitempty"`  // minOccurs="0"

	ArrivalFlightDetails *ArrivalFlightDetails `xml:"arrivalFlightDetails,omitempty"`  // minOccurs="0"

	// This segment is used to convey specific hotel booking indicators
	BookingIndicator *StatusType_99582S `xml:"bookingIndicator,omitempty"`  // minOccurs="0"
}

type RoomRateDetails struct {

	// This segment is used to convey the room information.
	RoomInformation *HotelRoomType `xml:"roomInformation"`

	// This group is used to conveys list of children
	Children []*ChildrenGroupType `xml:"children,omitempty"`  // minOccurs="0" maxOccurs="99"

	// This segment is used to convey the tariff/rate details.
	TariffDetails *TariffInformationTypeI `xml:"tariffDetails"`

	// This segment is used to convey the rate code indicator.
	RateCodeIndicator *RuleInformationTypeU `xml:"rateCodeIndicator,omitempty"`  // minOccurs="0"
}

type GuaranteeOrDeposit struct {

	// This segment is used to convey the guarantee or deposit information
	PaymentInfo *PaymentInformationTypeI `xml:"paymentInfo"`

	// This segment is used to convey the credit card information.
	CreditCardInfo *FormOfPaymentTypeI_29553S `xml:"creditCardInfo,omitempty"`  // minOccurs="0"
}

type ArrivalFlightDetails struct {

	// Travel Product Information
	TravelProductInformation *TravelProductInformationTypeI_186189S `xml:"travelProductInformation"`

	// Additional Transport Details
	AdditionalTransportDetails *AdditionalTransportDetailsTypeU `xml:"additionalTransportDetails"`
}

type TypicalCarData struct {

	// Vehicle information - vehicle type (SIPP code), - vehicle special equipments - vehicle details
	VehicleInformation *VehicleInformationType `xml:"vehicleInformation"`

	// Additional vehicle info
	AdditionalInfo *FreeTextInformationType_136708S `xml:"additionalInfo,omitempty"`  // minOccurs="0"

	// Voucher Print Acknowledgement.
	VoucherPrintAck *ReferenceInformationTypeI_136704S `xml:"voucherPrintAck,omitempty"`  // minOccurs="0"

	// CAR provider code
	CompanyIdentification *CompanyInformationType `xml:"companyIdentification"`

	// Car AVL pickup and dropoff location parameters (for Amadeus and Provider locations). Used as well to transport the Collection and Delivery place information.
	LocationInfo []*PlaceLocationIdentificationTypeU `xml:"locationInfo"`  // maxOccurs="6"

	DeliveryAndCollection []*DeliveryAndCollection `xml:"deliveryAndCollection,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Pickup and dropoff dates and times.
	PickupDropoffTimes *StructuredPeriodInformationType_136705S `xml:"pickupDropoffTimes"`

	// Cancellation or Confirmation number.
	CancelOrConfirmNbr []*ReservationControlInformationTypeI_136703S `xml:"cancelOrConfirmNbr,omitempty"`  // minOccurs="0" maxOccurs="2"

	RateCodeGroup *RateCodeGroup `xml:"rateCodeGroup,omitempty"`  // minOccurs="0"

	// Frequent flyer number.
	FFlyerNbr *FrequentTravellerIdentificationCodeType `xml:"fFlyerNbr,omitempty"`  // minOccurs="0"

	// Customer information ID and CD numbers
	CustomerInfo *ConsumerReferenceInformationTypeI `xml:"customerInfo,omitempty"`  // minOccurs="0"

	// This segments is used to convey: 1)General Rate information (identifier, plan, category) and Unstructured RB/RQ/RG 2) Structured rate quoted (RQ) or guaranted (RG) 3) Structured base rate (RB) 4) Estimated total information 5) Drop amount data 6) Voucher coupon print references (VV) 7) Rate Override (RO) 8) Modification fee indicator 9) Cancellation fee indicator 10) prepayment
	RateInfo []*TariffInformationTypeI_136706S `xml:"rateInfo,omitempty"`  // minOccurs="0" maxOccurs="10"

	ErrorWarning *ErrorWarning `xml:"errorWarning,omitempty"`  // minOccurs="0"

	RulesPoliciesGroup *RulesPoliciesGroup `xml:"rulesPoliciesGroup,omitempty"`  // minOccurs="0"

	// - Form of payment (FP) - Form of guarantee (G)
	Payment *FormOfPaymentTypeI `xml:"payment,omitempty"`  // minOccurs="0"

	// - Billing reference number (contains conpany data to be built) - Billing mumber (number included in the billing) - Agency account
	BillingData *BillableInformationTypeU `xml:"billingData,omitempty"`  // minOccurs="0"

	// The booking source
	BookingSource *AdditionalBusinessSourceInformationType `xml:"bookingSource,omitempty"`  // minOccurs="0"

	// Tour code
	InclusiveTour *TourInformationTypeI `xml:"inclusiveTour,omitempty"`  // minOccurs="0"

	// Contains: 1) up to 6 lines of marketing text sent by the car provider. 2) up to 3 lines of Other services messages (advertisments).
	MarketingInfo []*InteractiveFreeTextTypeI_136698S `xml:"marketingInfo,omitempty"`  // minOccurs="0" maxOccurs="2"

	// This segment is used to convey the supplementary  informations (SI). e.g: "Customer arriving after agency closure hour. Car keys waiting at the hotel reception located next to the agency".
	SupleInfo []*MiscellaneousRemarksType_136700S `xml:"supleInfo,omitempty"`  // minOccurs="0" maxOccurs="6"

	// This segment is used to convey distances.  1) Intercity distance. Distance between the Pickup and the Dropoff cities. Information returned by the Car provider for customer notification purpose.  2) Estimated distance Distance that is going to be runned during the rental period.
	EstimatedDistance []*QuantityTypeI `xml:"estimatedDistance,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Booking agent name
	AgentInformation *NameTypeU_136701S `xml:"agentInformation,omitempty"`  // minOccurs="0"

	// Tracking Option (TK)
	TrackingOpt *AgreementIdentificationTypeU `xml:"trackingOpt,omitempty"`  // minOccurs="0"

	// Electronic Voucher Number
	ElectronicVoucherNumber *TicketNumberTypeI `xml:"electronicVoucherNumber,omitempty"`  // minOccurs="0"

	// E-mail
	CustomerEmail *CommunicationContactTypeU `xml:"customerEmail,omitempty"`  // minOccurs="0"

	// This mandatory segment marks the end of the CAR data group. It specifies also if the booking is leisure or not.
	Attribute *AttributeType `xml:"attribute"`
}

type DeliveryAndCollection struct {

	// This Segment is used to Delivery and Collection information:  Format 1- (Home Collection): - Address - City - State - Country - Zip Code  Format 2- (Site Collection): - Site Ref Id - Site Name
	AddressDeliveryCollection *AddressTypeU_136710S `xml:"addressDeliveryCollection"`

	// This segment is used to carry phone number associated to a Delivery / Collection address
	PhoneNumber *PhoneAndEmailAddressType `xml:"phoneNumber"`
}

type RateCodeGroup struct {

	// Rate code
	RateCodeInfo *FareQualifierDetailsTypeI `xml:"rateCodeInfo"`

	// Additional Rate Code Information
	AdditionalInfo *FreeTextInformationType_136708S `xml:"additionalInfo,omitempty"`  // minOccurs="0"
}

type ErrorWarning struct {

	// Error/warning
	ApplicationError *ApplicationErrorInformationType_136725S `xml:"applicationError"`

	// Error or Warning freetext
	ErrorFreeText *FreeTextInformationType_136708S `xml:"errorFreeText,omitempty"`  // minOccurs="0"
}

type RulesPoliciesGroup struct {

	// Dummy segment to mark the beginning of the group
	Dummy1 *DummySegmentTypeI `xml:"dummy1"`

	// Present only if information is linked to a seamless availability
	SourceLevel *SelectionDetailsTypeI `xml:"sourceLevel,omitempty"`  // minOccurs="0"

	// Used to convey remarks corresponding to rule information.
	Remarks *FreeTextInformationType_136708S `xml:"remarks,omitempty"`  // minOccurs="0"

	TaxCovSurchargeGroup []*TaxCovSurchargeGroup `xml:"taxCovSurchargeGroup,omitempty"`  // minOccurs="0" maxOccurs="98"

	OtherRulesGroup []*OtherRulesGroup `xml:"otherRulesGroup,omitempty"`  // minOccurs="0" maxOccurs="13"

	PickupDropoffLocation []*PickupDropoffLocation `xml:"pickupDropoffLocation,omitempty"`  // minOccurs="0" maxOccurs="2"

	SpecialEquipmentDetails []*SpecialEquipmentDetails `xml:"specialEquipmentDetails,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type TaxCovSurchargeGroup struct {

	// This segment is used to convey Tax, Coverage, Coupon, Surcharge or Delivery and collection information  (If period associated to the surcharge, tariff and period definition conveyed in group 6)
	TaxSurchargeCoverageInfo *TariffInformationTypeI_136714S `xml:"taxSurchargeCoverageInfo"`

	// Additional information for Tax, Surcharge or Coverage section
	AdditionalInfo *FreeTextInformationType_136708S `xml:"additionalInfo,omitempty"`  // minOccurs="0"

	SurchargePeriods []*SurchargePeriods `xml:"surchargePeriods,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type SurchargePeriods struct {

	// tariff period/distance validity in number of days, weeks, months, km, miles.
	Period *RangeDetailsTypeI `xml:"period"`

	// This segment is used to convey Tax, Coverage, Coupon, Surcharge or Delivery and collection information
	SurchargePeriodTariff *TariffInformationTypeI_136719S `xml:"surchargePeriodTariff"`

	// This segment conveys the Unit Qualifier for maximum range of associated RNG.
	MaximumUnitQualifier *MeasurementsBatchTypeU `xml:"maximumUnitQualifier,omitempty"`  // minOccurs="0"
}

type OtherRulesGroup struct {

	// Used to convey the following type of information: - Pickup Information - Advance Payment Information - Policy Information - Deposit Information - Advance Booking Information - Guarantee Information - One Way Information
	OtherRules *RuleInformationTypeU_136720S `xml:"otherRules"`

	// Used to convey date/time Information (only used for Pickup and Guarantee rules)
	DateTimeInfo []*StructuredPeriodInformationType_136705S `xml:"dateTimeInfo,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type PickupDropoffLocation struct {

	// Car AVL pickup and dropoff location parameters (for Amadeus and Provider locations). Used as well to transport the Collection and Delivery place information.
	LocationInfo *PlaceLocationIdentificationTypeU_136722S `xml:"locationInfo"`

	// Location Address
	Address *AddressTypeU_136721S `xml:"address,omitempty"`  // minOccurs="0"

	// Location opening hours
	OpeningHours []*StructuredPeriodInformationType_136724S `xml:"openingHours,omitempty"`  // minOccurs="0" maxOccurs="10"

	// Phone / Fax number
	Phone []*PhoneAndEmailAddressType_136723S `xml:"phone,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type SpecialEquipmentDetails struct {

	// DUM used to remove ambiguity between TFFs
	Dummy2 *DummySegmentTypeI `xml:"dummy2"`

	RangePeriod []*RangePeriod `xml:"rangePeriod,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Additional special equipment information
	AdditionalInfo *FreeTextInformationType_136715S `xml:"additionalInfo,omitempty"`  // minOccurs="0"

	// First TFF occurence convey main data: - 013 spec. equipment code - Qualifier (Included / Optional) - Spec. equipment name  Up to 5 next occurences convey tarrif periods. - 013 spec. equipment code - converted indicator - amount/currency - period (/day, /weekend, /week, /month, /rental) - max amount / currency
	SpecialEquipmentTariff []*TariffInformationTypeI_136714S `xml:"specialEquipmentTariff"`  // maxOccurs="6"
}

type RangePeriod struct {

	// define age period validity associted to the special equipment
	AgePeriod *RangeDetailsTypeI `xml:"agePeriod"`

	// This segment conveys the Unit Qualifier for maximum range of associated RNG.
	MaximumUnitQualifier *MeasurementsBatchTypeU `xml:"maximumUnitQualifier,omitempty"`  // minOccurs="0"
}

type ErrorInfo struct {

	// identify the type of application error  within a message
	ErrorInformation *ApplicationErrorInformationType `xml:"errorInformation"`

	// provide free from or coded text information
	ErrorfreeFormText *InteractiveFreeTextTypeI_136698S `xml:"errorfreeFormText,omitempty"`  // minOccurs="0"
}

type DataElementsMaster struct {

	// marker
	Marker2 *DummySegmentTypeI `xml:"marker2"`

	DataElementsIndiv []*DataElementsIndiv `xml:"dataElementsIndiv,omitempty"`  // minOccurs="0" maxOccurs="999"
}

type DataElementsIndiv struct {

	// specify the PNR segments/elements references and action to apply
	ElementManagementData *ElementManagementSegmentType `xml:"elementManagementData"`

	// specify the amadeus PNR individual security element
	PnrSecurity *IndividualPnrSecurityInformationType `xml:"pnrSecurity,omitempty"`  // minOccurs="0"

	// Specify the amadeus accounting information
	Accounting *AccountingInformationElementType `xml:"accounting,omitempty"`  // minOccurs="0"

	// specify miscellaneous, confidential, quality control and invoice remarks
	MiscellaneousRemarks *MiscellaneousRemarksType_211S `xml:"miscellaneousRemarks,omitempty"`  // minOccurs="0"

	// specify special request or services information relating to a traveller
	ServiceRequest *SpecialRequirementsDetailsTypeI `xml:"serviceRequest,omitempty"`  // minOccurs="0"

	SeatPaxInfo []*SeatPaxInfo `xml:"seatPaxInfo,omitempty"`  // minOccurs="0" maxOccurs="9"

	// To Convey the Reason for Issuance Code (RFIC) and Reason For Issuance Sub code (RFISC)
	ReasonForIssuanceCode *PricingOrTicketingSubsequentType `xml:"reasonForIssuanceCode,omitempty"`  // minOccurs="0"

	// Rail Seat Preferences
	RailSeatPreferences *RailSeatPreferencesType `xml:"railSeatPreferences,omitempty"`  // minOccurs="0"

	CityPair *CityPair `xml:"cityPair,omitempty"`  // minOccurs="0"

	RailSeatDetails []*RailSeatDetails `xml:"railSeatDetails,omitempty"`  // minOccurs="0" maxOccurs="9"

	// provide date and time details relative to flight movements
	DateAndTimeInformation *DateAndTimeInformationTypeI `xml:"dateAndTimeInformation,omitempty"`  // minOccurs="0"

	// Details of SSR FQT content (Frequent Flyer Data)
	FrequentFlyerInformationGroup *FrequentFlyerInformationGroupType `xml:"frequentFlyerInformationGroup,omitempty"`  // minOccurs="0"

	// specify an amadeus PNR ticket element
	TicketElement *TicketElementType `xml:"ticketElement,omitempty"`  // minOccurs="0"

	ReferencedRecord *ReferencedRecord `xml:"referencedRecord,omitempty"`  // minOccurs="0"

	// option element
	OptionElement *OptionElementType `xml:"optionElement,omitempty"`  // minOccurs="0"

	// provide free form or coded long text information
	OtherDataFreetext []*LongFreeTextType `xml:"otherDataFreetext,omitempty"`  // minOccurs="0" maxOccurs="2"

	// specify the way data are mapped for the structured addresses
	StructuredAddress *StructuredAddressType `xml:"structuredAddress,omitempty"`  // minOccurs="0"

	// To specify the monetary information
	MonetaryInformation []*MonetaryInformationTypeI_1689S `xml:"monetaryInformation,omitempty"`  // minOccurs="0" maxOccurs="2"

	ElementErrorInformation *ElementErrorInformation `xml:"elementErrorInformation,omitempty"`  // minOccurs="0"

	McoRecord *McoRecord `xml:"mcoRecord,omitempty"`  // minOccurs="0"

	// Group Total Price
	TotalPrice *TotalPriceType `xml:"totalPrice,omitempty"`  // minOccurs="0"

	// Indicators at element level
	ElementsIndicators []*StatusTypeI_33257S `xml:"elementsIndicators,omitempty"`  // minOccurs="0" maxOccurs="5"

	// provide specific reference identification
	ReferenceForDataElement *ReferenceInfoType `xml:"referenceForDataElement,omitempty"`  // minOccurs="0"

	// Carries a Form of Payment in structured way.
	StructuredFop []*FOPRepresentationType `xml:"structuredFop,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type SeatPaxInfo struct {

	// details of the seat at pax level
	SeatPaxDetails *SeatRequestParametersTypeI `xml:"seatPaxDetails"`

	// seat indicator at pax level
	SeatPaxIndicator *StatusTypeI_33257S `xml:"seatPaxIndicator,omitempty"`  // minOccurs="0"

	// ref to pax tattoo
	CrossRef *ReferenceInfoType_6074S `xml:"crossRef"`
}

type CityPair struct {

	// Departure station location
	DepLocation *PlaceLocationIdentificationTypeU_35293S `xml:"depLocation"`

	// Arrival station location
	ArrLocation *PlaceLocationIdentificationTypeU_35293S `xml:"arrLocation"`
}

type RailSeatDetails struct {

	// Used to convey specific seat details relative to Train for a specific request or the "near-to" seat details for a "next-to" request.
	RailSeatReferenceInformation *RailSeatReferenceInformationType `xml:"railSeatReferenceInformation"`

	// Rail Seat Denomination
	RailSeatDenomination *FreeTextInformationType_29860S `xml:"railSeatDenomination,omitempty"`  // minOccurs="0"
}

type ReferencedRecord struct {

	// specify a reference to a reservation
	ReferencedReservationInfo *ReservationControlInformationTypeI_87792S `xml:"referencedReservationInfo"`

	// specify the amadeus PNR record locator security information for different pnr elements .
	SecurityInformation *ReservationSecurityInformationType_167761S `xml:"securityInformation"`
}

type ElementErrorInformation struct {

	// identify the type of application error  within a message
	ErrorInformation *ApplicationErrorInformationType `xml:"errorInformation"`

	// provide free from or coded text information
	ElementErrorText *InteractiveFreeTextTypeI_136698S `xml:"elementErrorText,omitempty"`  // minOccurs="0"
}

type McoRecord struct {

	// specify that a MCO element is present in the PNR - this is a visual trigger of the MCO
	McoType *MiscellaneousChargeOrderType `xml:"mcoType"`

	// Contains the data relative to the MCO element itself
	McoInformation *FreeTextInformationType_9865S `xml:"mcoInformation"`

	GroupOfFareElements []*GroupOfFareElements `xml:"groupOfFareElements,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type GroupOfFareElements struct {

	// Sequence Number for a Fare element
	SequenceNumber *SequenceDetailsTypeU `xml:"sequenceNumber"`

	// Contains the Fare Element data
	FareElementData *FreeTextInformationType_9865S `xml:"fareElementData"`
}

type TstData struct {

	// TST general information
	TstGeneralInformation *TstGeneralInformationType `xml:"tstGeneralInformation"`

	// provide free form or coded long text information
	TstFreetext []*LongFreeTextType `xml:"tstFreetext,omitempty"`  // minOccurs="0" maxOccurs="2"

	// describe fare basis information
	FareBasisInfo *FareBasisCodesLineType `xml:"fareBasisInfo,omitempty"`  // minOccurs="0"

	// fare data
	FareData *FareDataType `xml:"fareData,omitempty"`  // minOccurs="0"

	// selection details
	SegmentAssociation *SelectionDetailsTypeI_2067S `xml:"segmentAssociation,omitempty"`  // minOccurs="0"

	// provide specific reference identification
	ReferenceForTstData *ReferenceInfoType `xml:"referenceForTstData,omitempty"`  // minOccurs="0"
}

type DcdData struct {

	// This is used as DUM segment
	MarkerPax *PassengerFlightDetailsTypeI `xml:"markerPax"`

	// This is used as a Dum segment.
	MarkerSegment *PassengerFlightDetailsTypeI `xml:"markerSegment"`

	SegmentSection []*SegmentSection `xml:"segmentSection,omitempty"`  // minOccurs="0" maxOccurs="19602"

	// This is used as a Dum segement.
	MarkerLeg *PassengerFlightDetailsTypeI `xml:"markerLeg"`

	LegSection []*LegSection `xml:"legSection,omitempty"`  // minOccurs="0" maxOccurs="176418"
}

type SegmentSection struct {

	// Specify structured elements references
	ElementManagementStructData *ElementManagementSegmentType_127983S `xml:"elementManagementStructData"`

	// provide specific reference identification
	ReferenceForStructDataElement *ReferenceInfoType `xml:"referenceForStructDataElement,omitempty"`  // minOccurs="0"

	DcsSegmentInfo *DcsSegmentInfo `xml:"dcsSegmentInfo"`
}

type DcsSegmentInfo struct {

	// Booking information. This is not required by the process it self, but can be used to easily track problems. This segment is required also to solve ambiguity problems. It can be empty if there is not need to convey information (more meaningful than a DUM)
	Booking *TravelProductInformationTypeI_127288S `xml:"booking"`

	// Determines if the DCS Data apply to the adult or to the infant (in case there is one). By default, it applies to the adult.
	PaxType *ReferenceInformationTypeI `xml:"paxType"`

	// Gives the compensation type. * attributeDetails/attributeType = - DBA - DBN - DBO - DBV - DBM - DBD - DBW
	TypeOfCOP *CodedAttributeType_127282S `xml:"typeOfCOP,omitempty"`  // minOccurs="0"
}

type LegSection struct {

	// Specify structured elements references
	ElementManagementStructData *ElementManagementSegmentType_127983S `xml:"elementManagementStructData"`

	// provide specific reference identification
	ReferenceForStructDataElement *ReferenceInfoType `xml:"referenceForStructDataElement,omitempty"`  // minOccurs="0"

	DcsLegInfo *DcsLegInfo `xml:"dcsLegInfo"`
}

type DcsLegInfo struct {

	// Contains the leg position inside the booking
	LegPosition *TravelItineraryInformationTypeI `xml:"legPosition"`

	// Indenties uniquely a leg inside a multi-leg booking
	Leg *OriginAndDestinationDetailsTypeI `xml:"leg"`

	// Determines if the DCS Data apply to the adult or to the infant (in case there is one). By default, it applies to the adult.
	PaxType *ReferenceInformationTypeI `xml:"paxType"`

	// Contains information on the seat delivered by the DCS
	SeatDelivery *SpecialRequirementsDetailsType `xml:"seatDelivery,omitempty"`  // minOccurs="0"

	// Third data element provide the category of attribute: NOREC information, acceptance status... First data element contains the value of the attribute : the NOREC flag, the acceptance status, the boarding status and the cabin regrade type, Check Bags indicator, Waitlist status.
	PaxStatus *StatusTypeI `xml:"paxStatus,omitempty"`  // minOccurs="0"

	AccregReason []*AccregReason `xml:"accregReason,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Regrade cabin code
	RegradeCabin *SegmentCabinIdentificationType `xml:"regradeCabin,omitempty"`  // minOccurs="0"

	AcceptanceChannel *AcceptanceChannel `xml:"acceptanceChannel,omitempty"`  // minOccurs="0"

	// Provides information on the compensation offered to passengers with valid tickets, airline turned down at check-in/boarding. - coded form of payment (NGDCS only) - currency code (NGDCS only) - amount (NGDCS only) - compensation type (voluntary/involuntary) - free text (both valid for PFS clients and NGDCS)
	CompensationData *CompensationType `xml:"compensationData,omitempty"`  // minOccurs="0"
}

type AccregReason struct {

	// Reason code for: - Acceptation/Cancellation - Regrade
	Reasons *CodedAttributeType_127279S `xml:"reasons"`

	// Contains Acceptance Reason, Regrade Reason description.
	DeliveryInformation *InteractiveFreeTextTypeI `xml:"deliveryInformation,omitempty"`  // minOccurs="0"
}

type AcceptanceChannel struct {

	// Qualifies originator of the acceptance actions: - A for Check-in Agent - D for Direct consumer  - S for automated Devices - E for internet/web based application
	AcceptanceOrigin *UserIdentificationType_127265S `xml:"acceptanceOrigin"`

	// contains the application used to perform check-in operations: - cryptic application - Java front End application - SMS application - Telephone - web application - External departure control application
	ApplicationType *ApplicationType `xml:"applicationType"`
}

//
// Complex structs
//

type AccommodationAllocationInformationDetailsTypeU struct {

	// Accommodation (room/compartment) number
	ReferenceId string `xml:"referenceId"`

	// Accommodation (room/compartment) code
	Code string `xml:"code,omitempty"`  // minOccurs="0"
}

type AccommodationAllocationInformationTypeU struct {

	// Allocated accommodation
	AccommAllocation *AccommodationAllocationInformationDetailsTypeU `xml:"accommAllocation"`
}

type AccountingElementType struct {

	// Account number
	Number string `xml:"number,omitempty"`  // minOccurs="0"

	// Cost Number
	CostNumber string `xml:"costNumber,omitempty"`  // minOccurs="0"

	// IATA company number
	CompanyNumber string `xml:"companyNumber,omitempty"`  // minOccurs="0"

	// Client Reference Number
	ClientReference string `xml:"clientReference,omitempty"`  // minOccurs="0"

	// Format limitations: an..109
	GSTTaxDetails string `xml:"gSTTaxDetails,omitempty"`  // minOccurs="0"
}

type AccountingInformationElementType struct {

	// One of these 4 data elements is mandatory , but non in particular
	Account *AccountingElementType `xml:"account,omitempty"`  // minOccurs="0"

	// Number of units qualifier
	AccountNumberOfUnits string `xml:"accountNumberOfUnits,omitempty"`  // minOccurs="0"
}

type ActionDetailsTypeI struct {

	// Contains the details about the product knowledge
	NumberOfItemsDetails *ProcessingInformationTypeI `xml:"numberOfItemsDetails,omitempty"`  // minOccurs="0"
}

type AdditionalBusinessSourceInformationType struct {

	// ORIGINATOR DETAILS
	OriginatorDetails *OriginatorIdentificationDetailsTypeI_198179C `xml:"originatorDetails"`
}

type AdditionalProductDetailsTypeI struct {

	ProductDetails *AdditionalProductTypeI `xml:"productDetails,omitempty"`  // minOccurs="0"

	DepartureInformation *StationInformationTypeI `xml:"departureInformation,omitempty"`  // minOccurs="0"

	ArrivalStationInfo *StationInformationTypeI_119771C `xml:"arrivalStationInfo,omitempty"`  // minOccurs="0"

	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`  // minOccurs="0"

	TimeDetail *TravellerTimeDetailsTypeI `xml:"timeDetail,omitempty"`  // minOccurs="0"

	Facilities []*ProductFacilitiesTypeI `xml:"facilities,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type AdditionalProductDetailsTypeU struct {

	// Conveys the product area (TOU)
	ProductArea string `xml:"productArea"`

	// The general product description
	ProductDetails *ProductDataInformationTypeU `xml:"productDetails"`
}

type AdditionalProductTypeI struct {

	// Format limitations: an..3
	Equipment string `xml:"equipment,omitempty"`  // minOccurs="0"

	// Format limitations: n..2
	NumOfStops *int32 `xml:"numOfStops,omitempty"`  // minOccurs="0"

	// Time format: 24H. All digits are mandatory . Example: from 0000 to 2359
	Duration string `xml:"duration,omitempty"`  // minOccurs="0"

	// Format limitations: n1
	WeekDay *int32 `xml:"weekDay,omitempty"`  // minOccurs="0"
}

type AdditionalTransportDetailsTypeU struct {

	// Terminal Information
	TerminalInformation []*TerminalInformationTypeU `xml:"terminalInformation"`  // maxOccurs="2"
}

type AddressDetailsTypeU struct {

	// Address Format . Will be 5 unstructured
	Format string `xml:"format"`

	// Address Text. Any of the following address lines may start with a tag: Door number- Street- ExternalNumber- InternalNumber- County- Neighbourhood- State-
	Line1 string `xml:"line1"`

	// Format limitations: an..70
	Line2 string `xml:"line2,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	Line3 string `xml:"line3,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	Line4 string `xml:"line4,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	Line5 string `xml:"line5,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	Line6 string `xml:"line6,omitempty"`  // minOccurs="0"
}

type AddressDetailsTypeU_17987C struct {

	// To specify what kind of address we have
	Format string `xml:"format"`

	// Address
	Line1 string `xml:"line1"`

	// Address
	Line2 string `xml:"line2,omitempty"`  // minOccurs="0"

	// PO Box
	Line4 string `xml:"line4,omitempty"`  // minOccurs="0"
}

type AddressDetailsTypeU_198210C struct {

	// Address format
	Format string `xml:"format"`

	// Address Field in free flow text
	Line1 string `xml:"line1"`
}

type AddressDetailsTypeU_198226C struct {

	// - 5 for unstructured address
	Format string `xml:"format"`

	// address line 1
	Line1 string `xml:"line1"`

	// address line 2
	Line2 string `xml:"line2,omitempty"`  // minOccurs="0"
}

type AddressType struct {

	// will convey the adress text
	AddressDetails *AddressDetailsTypeU `xml:"addressDetails,omitempty"`  // minOccurs="0"

	// City name.
	City string `xml:"city,omitempty"`  // minOccurs="0"

	// postal identification code.
	ZipCode string `xml:"zipCode,omitempty"`  // minOccurs="0"

	// Country code. ISO 3166 code for the country
	CountryCode string `xml:"countryCode,omitempty"`  // minOccurs="0"
}

type AddressTypeU struct {

	// to specify the information of the address
	AddressDetails *AddressDetailsTypeU_17987C `xml:"addressDetails,omitempty"`  // minOccurs="0"

	// city name of the given address
	City string `xml:"city,omitempty"`  // minOccurs="0"

	// zip code of the given address
	ZipCode string `xml:"zipCode,omitempty"`  // minOccurs="0"

	// To convey a sub-entity within a country : region, states..
	RegionDetails *CountrySubEntityDetailsTypeU `xml:"regionDetails,omitempty"`  // minOccurs="0"

	// to specify the countryname
	LocationDetails *LocationIdentificationTypeU `xml:"locationDetails,omitempty"`  // minOccurs="0"
}

type AddressTypeU_136710S struct {

	// Address Type
	AddressUsageDetails *AddressUsageTypeU `xml:"addressUsageDetails"`

	// Format 1 - Home Delivery/Collection
	AddressDetails *AddressDetailsTypeU_198210C `xml:"addressDetails,omitempty"`  // minOccurs="0"

	// City name
	City string `xml:"city,omitempty"`  // minOccurs="0"

	// Postal Code
	ZipCode string `xml:"zipCode,omitempty"`  // minOccurs="0"

	// Country code
	CountryCode string `xml:"countryCode,omitempty"`  // minOccurs="0"

	// To convey a sub-entity within a country : region, states..
	RegionDetails *CountrySubEntityDetailsTypeU_198213C `xml:"regionDetails,omitempty"`  // minOccurs="0"

	// Format 2 - Site Delivery/Collection
	LocationDetails *LocationIdentificationTypeU_198211C `xml:"locationDetails,omitempty"`  // minOccurs="0"
}

type AddressTypeU_136721S struct {

	// Location address
	AddressDetails *AddressDetailsTypeU_198226C `xml:"addressDetails"`

	// City Name
	City string `xml:"city,omitempty"`  // minOccurs="0"

	// Postal Code
	ZipCode string `xml:"zipCode,omitempty"`  // minOccurs="0"

	// Country code
	CountryCode string `xml:"countryCode,omitempty"`  // minOccurs="0"

	// To convey a sub-entity within a country : region, states..
	RegionDetails *CountrySubEntityDetailsTypeU_198229C `xml:"regionDetails,omitempty"`  // minOccurs="0"
}

type AddressUsageTypeU struct {

	// Address Type: - DEL for Delivery - COL for Collection
	Purpose string `xml:"purpose"`
}

type AgreementIdentificationTypeU struct {

	// Agreement identification
	AgreementDetails *AgreementTypeIdentificationTypeU `xml:"agreementDetails,omitempty"`  // minOccurs="0"
}

type AgreementTypeIdentificationTypeU struct {

	// - TK for Tracking option
	Code string `xml:"code"`

	// Agreement description
	Description string `xml:"description"`
}

type ApplicationErrorDetailType struct {

	// Code identifying the data validation error condition.
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	ErrorCategory string `xml:"errorCategory"`
}

type ApplicationErrorDetailTypeI struct {

	// Message number or "ZZZ" if no number
	ErrorCode string `xml:"errorCode"`

	// EC for Error codes  WEC for Warning code  INF for Information code
	Qualifier string `xml:"qualifier"`

	// 3 for IATA   UN for UN  1A for AMADEUS
	ResponsibleAgency string `xml:"responsibleAgency"`
}

type ApplicationErrorDetailType_198235C struct {

	// Code identifying the data validation error condition.
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	ErrorCategory string `xml:"errorCategory,omitempty"`  // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`  // minOccurs="0"
}

type ApplicationErrorInformationType struct {

	// Detail the error type
	ErrorDetail *ApplicationErrorDetailTypeI `xml:"errorDetail"`
}

type ApplicationErrorInformationType_136725S struct {

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType_198235C `xml:"errorDetails"`
}

type ApplicationErrorInformationType_94519S struct {

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type ApplicationIdentificationType struct {

	// application internal identifier
	InternalId string `xml:"internalId"`

	// Item Version Number
	VersionNumber string `xml:"versionNumber,omitempty"`  // minOccurs="0"
}

type ApplicationType struct {

	// provides information on application identification
	ApplicationDetails *ApplicationIdentificationType `xml:"applicationDetails"`
}

type AssociatedChargesInformationTypeI struct {

	// This data element is used to identify the type of charge entered in the other fields.
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// This data element is used to convey the amount of the supplementary charge.
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// To qualify the amount, can be - UNL (for unlimited mileage) when used for free mileage - 3 (for included in base rate) - 4 (for not included in base rate)
	Description string `xml:"description,omitempty"`  // minOccurs="0"

	// This data element is used to specify the number of charge needed.
	NumberInParty *int32 `xml:"numberInParty,omitempty"`  // minOccurs="0"

	// This data element is used to convey the currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// This data element is used to convey the voucher text (in case of voucher).
	Comment string `xml:"comment,omitempty"`  // minOccurs="0"
}

type AssociatedChargesInformationTypeI_198205C struct {

	// This data element is used to identify the type of charge entered in the other fields.
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Mileage charge amount
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// To qualify the amount, can be - UNL (for unlimited mileage) when used for free mileage - 3 (for included in base rate) - 4 (for not included in base rate)
	Description string `xml:"description,omitempty"`  // minOccurs="0"

	// Quantity of free mileage
	NumberInParty *int32 `xml:"numberInParty,omitempty"`  // minOccurs="0"

	// Unit: - K Kilometer - M Miles
	PeriodType string `xml:"periodType,omitempty"`  // minOccurs="0"

	// the currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Unstructured RG,RG and RQ rates.
	Comment string `xml:"comment,omitempty"`  // minOccurs="0"
}

type AssociatedChargesInformationTypeI_198218C struct {

	// - 045 Tax - 108 Surchage - COV Coverage - CPN Coupon
	Type string `xml:"type"`

	// Policy amount (coupon amount)
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Qualifier: The possible values are: - IES included in Estimated Total - IBR included in Base Rate - OPT Optional - MAN Mandatory - NBR Not Included in Base Rate - ITX Policy amount Includes Tax - NTX Policy amount Not Includes Tax
	Description string `xml:"description,omitempty"`  // minOccurs="0"

	// Maximum days
	NumberInParty *int32 `xml:"numberInParty,omitempty"`  // minOccurs="0"

	// 001 per day 002 per week 003 per month 004 per rental 012 tax percentage 013 no coupon value available
	PeriodType string `xml:"periodType,omitempty"`  // minOccurs="0"

	// Policy amount currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Policy name
	Comment string `xml:"comment,omitempty"`  // minOccurs="0"
}

type AssociatedChargesInformationTypeI_39535C struct {

	// To specify the type of tax, the type of converted amount. It is coded on our side if not specifued by provider.
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// to specify the tax in a foreign currency.
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Tax name
	Description string `xml:"description,omitempty"`  // minOccurs="0"

	// foreign currency.
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type AssociatedChargesInformationTypeU struct {

	// Qualify the associated charge. For Tour, only "employee" is used to define a commission
	ChargeUnitCode string `xml:"chargeUnitCode"`

	// Value of the associated charge
	Amount float64 `xml:"amount"`

	// Commission's percentage
	Percentage float64 `xml:"percentage"`
}

type AttributeInformationTypeU struct {

	// Type of the authorization data.  Some of the possible types are:  25: (AUT) Context (Credit Mutuel) 26: (ATN) Customer instruction (Barclays) 27: (ATN) Cryptogram computation method (Credit Mutuel) 28: (AUT) Modified securisation mode (Credit Mutuel) 29: (ATN) Electronic commerce transaction type (Credit Mutuel) E: (ATN) Result of the secured payment VADS (Credit Mutuel)  MID: (AUT) Merchant ID
	AttributeType string `xml:"attributeType"`

	// value of the data
	AttributeDescription string `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type AttributeInformationTypeU_198185C struct {

	// the attribute type LEI for leisure booking CLP for clip booking
	AttributeType string `xml:"attributeType"`

	// Not Used
	AttributeDescription string `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type AttributeInformationTypeU_36633C struct {

	// This element is used to convey the service code of the service group of the ferry booking. The list of possible values depends of the Ferry provider.
	AttributeType string `xml:"attributeType"`
}

type AttributeInformationTypeU_45068C struct {

	// The list of possible values is: ADT Adult CHD Child FDC Diplomatic corps FEU Disabled FFM Family FFR Free FIR Inter rail FJO Journalist FSL School pupil INF Infant MIL Military NAT Nato official REC Child resident RES Resident SRC Senior citizen STU Student YTH Young person
	AttributeType string `xml:"attributeType"`
}

type AttributeType struct {

	// Specify which attribute is described in E003. BAT for booking attribute
	CriteriaSetType string `xml:"criteriaSetType"`

	// Details for the attribute type. LEI:Y for leisure booking CLP:Y for clip booking
	CriteriaDetails []*AttributeInformationTypeU_198185C `xml:"criteriaDetails"`  // maxOccurs="2"
}

type AttributeTypeU struct {

	// Describes the service type.
	AttributeFunction string `xml:"attributeFunction"`

	// Service details.
	AttributeDetails *AttributeInformationTypeU_36633C `xml:"attributeDetails"`
}

type AttributeTypeU_24552S struct {

	// provides the function of the attribute
	AttributeFunction string `xml:"attributeFunction"`

	// provides details for the Attribute
	AttributeDetails *AttributeInformationTypeU_45068C `xml:"attributeDetails"`
}

type AttributeType_94514S struct {

	// Determines if the set of criteria corresponds to the message identification criteria or to normal criteria.
	CriteriaSetType string `xml:"criteriaSetType,omitempty"`  // minOccurs="0"

	// List of attributes and status linked to credit card process. Most of them are link dependant.
	CriteriaDetails *AttributeInformationTypeU `xml:"criteriaDetails"`
}

type AttributeType_94553S struct {

	// Type of Data Exple :  SAL sale indicator EXT for extended payment PAY payment type
	CriteriaSetType string `xml:"criteriaSetType"`

	// Details for the message criteria (name, value).
	CriteriaDetails *AttributeInformationTypeU `xml:"criteriaDetails"`
}

type AttributeType_94576S struct {

	// Type of information: - is this a switch? - is this a structured data?
	CriteriaSetType string `xml:"criteriaSetType"`

	// Details for the message criteria (name, value).
	CriteriaDetails []*AttributeInformationTypeU `xml:"criteriaDetails"`  // maxOccurs="999"
}

type AuthenticationDataType struct {

	// VERes status (enrollment) Values : Y : authentication available N : cardholder not participating U : Unable to authenticate E : error message
	Veres string `xml:"veres,omitempty"`  // minOccurs="0"

	// PARes status (authentication). Values : Y : authentication successful N : authentication failed U : authentication could not be performed A : attempts processing performed
	Pares string `xml:"pares,omitempty"`  // minOccurs="0"

	// CC Directory Server performing the enrollment process: VISA, MasterCard
	CreditCardCompany string `xml:"creditCardCompany"`

	// To indicate whether the transaction was successful, different indicators for Visa/MasterCard. - ECI for VISA - UCAF collection indicator for Matercard
	AuthenticationIndicator string `xml:"authenticationIndicator,omitempty"`  // minOccurs="0"

	// Indicates the algorithm used to generate the Cardholder Authentication Verification Value (CAAV = authentication code)
	CaavAlgorithm *int32 `xml:"caavAlgorithm,omitempty"`  // minOccurs="0"
}

type AuthorizationApprovalDataType struct {

	// will convey the value of the approval code of the payment authorisation
	ApprovalCode string `xml:"approvalCode"`

	// Source of approval for the payment authorisation. A Automatically obtained by the system. M Manually entered by an agent.  F: Credit card automatic approval code of a settlement authorization transaction B: Credit card manual approval code of a settlement transaction.
	SourceOfApproval string `xml:"sourceOfApproval,omitempty"`  // minOccurs="0"
}

type BillableInformationTypeU struct {

	// This composite is used to convey the billable information.
	BillingInfo []*DiagnosisTypeU `xml:"billingInfo"`  // maxOccurs="3"
}

type BinaryDataType struct {

	// Length of the BLB
	DataLength int32 `xml:"dataLength"`

	// type of the data
	DataType string `xml:"dataType,omitempty"`  // minOccurs="0"

	// used to store binary data
	BinaryData string `xml:"binaryData"`
}

type BrowserInformationType struct {

	// Indicates the type of cardholder device.
	DeviceCategory int32 `xml:"deviceCategory"`
}

type CabinClassDesignationType struct {

	// Designates the class of service on the means of transport  in which the passenger will travel:  - M for Economy - W for Economy Premium - C for Business (Club) - F for First  - Y for Economy All
	ClassDesignator string `xml:"classDesignator,omitempty"`  // minOccurs="0"
}

type CabinDetailsType struct {

	CabinDetails *CabinClassDesignationType `xml:"cabinDetails,omitempty"`  // minOccurs="0"
}

type CardValidityType struct {

	// Type of the compensation, ie voluntary or involuntary
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Form of the payment of the compensation
	Form string `xml:"form,omitempty"`  // minOccurs="0"

	// Amount of the compensation
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency used for the compensation.
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Any comment related to the compensation
	FreeText string `xml:"freeText,omitempty"`  // minOccurs="0"
}

type ChildrenGroupType struct {

	// This segment is used to convey age for a child.
	Age *QuantityTypeI_65488S `xml:"age"`

	// This segment is used to convey the passenger association
	ReferenceForPassenger *ReferenceInformationType_65487S `xml:"referenceForPassenger,omitempty"`  // minOccurs="0"
}

type ClassConfigurationDetailsType struct {

	// Class Details -Class Group -Sub Class -number of seats
	ClassDetails *ClassDetailsType `xml:"classDetails"`
}

type ClassDetailsType struct {

	// Class Group : A-First Class, Seat B-Second Class, Seat C-First Class, Berth D-Second Class, Berth F-Binded Seat V-First Class, Sleeping-car W-Second Class, Sleeping-car
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Rail class code.
	BookingClass string `xml:"bookingClass,omitempty"`  // minOccurs="0"

	// Number of Free Seats
	NumberOfSeats int32 `xml:"numberOfSeats"`
}

type ClassDetailsType_52782C struct {

	// For the booking class code.
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Format limitations: an2
	BookingClass string `xml:"bookingClass,omitempty"`  // minOccurs="0"
}

type CoachProductInformationType struct {

	// Coach Product Details
	CoachDetails *ReferencingDetailsTypeI_36941C `xml:"coachDetails,omitempty"`  // minOccurs="0"

	// Coach Equipment Qualifier
	EquipmentCode string `xml:"equipmentCode,omitempty"`  // minOccurs="0"
}

type CodedAttributeInformationType struct {

	// provides the attribute Type
	AttributeType string `xml:"attributeType"`
}

type CodedAttributeInformationType_142109C struct {

	// will convey the following QF data:  ONOD onoData     Order Number(Qantas specific)  GWTD gwtData     Government Warrant number(Qantas specific)  HOLDN ccHolderName    Conveys the CN (company name) (Qantas specific).This has sense only in case of automatic creation of attribute record (cards is a bets card). In the other cases this information cannot be filled.  ONOR onoRequired    This ONO indicator indicates whether or not ONO data is mandatory.(Information provided by Qantas IGW link) This has sense only in case of automatic creation of attribute record (cards is a bets card)  GWTR gwtRequired     This GWT indicator indicates whether or not GWT data is mandatory.(Information provided by Qantas IGW link) This has sense only in case of automatic creation of attribute record (cards is a bets card)  CIND cind     Conveys CIND indicator : - MANU - AUTO This indicates whether attributes records have been created manually (no bets card) or automatically (bets card).  BFAREC bestFareCandidate     Conveys best fare indicator: - Yes - No This indicates whether or not the card is best fare candidate. (this implies card is a bets card)
	AttributeType string `xml:"attributeType"`

	// onoData   Order Number(Qantas specific)  gwtData   Government Warrant number  ccHolderName  Conveys the CN   onoRequired  YES - NO  gwtRequired   YES - NO  cind    MANU - AUTO  bestFareCandidate   YES - NO
	AttributeDescription string `xml:"attributeDescription"`
}

type CodedAttributeInformationType_185753C struct {

	// provides the attribute Type
	AttributeType string `xml:"attributeType"`

	// provides a description for the attribute: If 950K set to ARC, value can be: AS: Airline Staff, BA: Baggage not Accepted, CB: Customer failed to board, CR: Customer Request, CU: Customer Unwell, DB: Denied Boarding, FA: Flight Alternative Offered and accepted by the customer, FD: Flight Delayed, FO: Flight Oversold, FC: Flight Cancelled, MC: Missed Connection, MR: Medical Reasons, NO: NOSHOW (can only be used with Target Customer Acceptance Status set to 'Rejected'), OT: Other, RR: Regulatory Requirement not met, SR: Security Reasons, TC:  Travel in different cabin through another booking, TD: Travel Documentation Incomplete, TI: Travel Industry Staff, UC: User Error Correction. If 950K set to RRC, value can be: OC: Cabin/Flight oversold (current flight), OO: Cabin/Flight oversold (other flight), MC: Misidentification of customer at check-in, PM: Previously mishandled, SO: Special occasion, AC: Aircraft change, CF: Cabin configuration change, RQ: Request from a special requestor, CO: Compassionate, MI: Marketing Initiative e.g. please try Club Class for free, DC: Disruption on current flight, DO: Disruption on other flight, CC: Crew level change, TR: Technical reason, CS: Catering shortfall, CI: Check-in error, IA: Inadmissible, ST: Staff, SR: Service Recovery Entitlement, AB: Authorized by, US: Unsuitable.
	AttributeDescription string `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeType struct {

	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"`  // maxOccurs="99"
}

type CodedAttributeType_127279S struct {

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType_185753C `xml:"attributeDetails"`
}

type CodedAttributeType_127282S struct {

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails"`
}

type CodedAttributeType_94497S struct {

	// Usage of this element will be the transport of the DescriptiveBilingInformation first value will be QF
	AttributeFunction string `xml:"attributeFunction"`

	// provides details for the Attribute
	AttributeDetails []*CodedAttributeInformationType_142109C `xml:"attributeDetails"`  // maxOccurs="99"
}

type CodeshareFlightDataTypeI struct {

	// Company identification
	Airline string `xml:"airline"`

	// Product identification
	FlightNumber int32 `xml:"flightNumber"`

	// general indicator
	Inventory string `xml:"inventory"`

	// Characteristic identification
	SellingClass string `xml:"sellingClass"`

	// Item characteristic
	Type string `xml:"type"`

	// Product identification characteristic
	Suffix string `xml:"suffix,omitempty"`  // minOccurs="0"

	// 1 :  codeshare cascading is successful 0 : codeshare cascading unsuccessful blank: n/s
	CascadingIndicator *int32 `xml:"cascadingIndicator,omitempty"`  // minOccurs="0"
}

type CommissionDetailsType struct {

	// Commission type : 'NEW' --) New commission 'OLD' --) Old Commission 'XLP' --) Commission on cancellation Penalty 'FMA' --) Airline Commission A 'FMB' --) Airline Commission B
	Type string `xml:"type"`

	// Commission amount
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Commission percentage
	Rate *int32 `xml:"rate,omitempty"`  // minOccurs="0"

	// Deal number
	DealNumber *int32 `xml:"dealNumber,omitempty"`  // minOccurs="0"
}

type CommissionInformationType struct {

	// Commission details
	CommissionDetails *CommissionDetailsType `xml:"commissionDetails"`
}

type CommunicationContactDetailsType struct {

	// The communication address: an Url.
	UrlAddress string `xml:"urlAddress"`

	// will be AH for World Wide Web
	UrlType string `xml:"urlType"`
}

type CommunicationContactDetailsTypeU struct {

	// Email
	Email string `xml:"email"`

	// Contact qualifier. EM for Electronic mail
	ContactQualifier string `xml:"contactQualifier"`
}

type CommunicationContactType struct {

	// Communication channel
	Communication *CommunicationContactDetailsType `xml:"communication"`
}

type CommunicationContactTypeU struct {

	// Customer contact
	Contact *CommunicationContactDetailsTypeU `xml:"contact"`
}

type CompanyIdentificationTypeI struct {

	// Format limitations: an..3
	OperatingCompany string `xml:"operatingCompany,omitempty"`  // minOccurs="0"
}

type CompanyIdentificationTypeI_2785C struct {

	// Company code
	Identification string `xml:"identification"`

	// AIR segment : second airline code for joint flight number. Train Amtrack segment : system provider UIC code. Train SNCF segment : system provider UIC code. Tour segment : tour provider code.
	SecondIdentification string `xml:"secondIdentification,omitempty"`  // minOccurs="0"

	// Tour segment : source code.
	SourceCode string `xml:"sourceCode,omitempty"`  // minOccurs="0"
}

type CompanyIdentificationTypeI_46335C struct {

	// carrier details
	MarketingCompany string `xml:"marketingCompany"`
}

type CompanyIdentificationTypeI_46351C struct {

	// Targeted provider system code
	OperatingCompany string `xml:"operatingCompany"`
}

type CompanyIdentificationTypeU struct {

	// Conveys the provider name
	ProviderName string `xml:"providerName"`
}

type CompanyInformationType struct {

	// Qualify the company code, to identify the industry business it belongs.
	TravelSector string `xml:"travelSector"`

	// This data element is used to convey the context in which the code applies. The hotel chain code are managed by Amadeus.
	CompanyCodeContext string `xml:"companyCodeContext,omitempty"`  // minOccurs="0"

	// This data element is used to convey the company code of a non-air company
	CompanyCode string `xml:"companyCode"`

	// This data element is used to convey the company name of a non-air company
	CompanyName string `xml:"companyName,omitempty"`  // minOccurs="0"

	// This data element is used to convey the type of access the non-air company has with Amadeus.
	AccessLevel string `xml:"accessLevel,omitempty"`  // minOccurs="0"
}

type CompanyInformationType_19450S struct {

	// This data element is used to convey the company code
	CompanyCode string `xml:"companyCode"`

	// This data element is used to convey the UIC code
	CompanyNumericCode *int32 `xml:"companyNumericCode,omitempty"`  // minOccurs="0"
}

type CompanyInformationType_20151S struct {

	// This element is used to qualify the company code, to identify the industry business it belongs. For Ferry, the codes are mutually agreed between Amadeus and the Ferry providers and are only valid within the Amadeus Ferry application.
	TravelSector string `xml:"travelSector,omitempty"`  // minOccurs="0"

	// This data element is used to convey the company code of a company. For Ferry, the list of providers is not fixed. The providers implemented actually are: Baleria                 BAL Brittany ferries        BRI Color Lines             COL Comarit                 COM Corsica ferries         CSF Grandi Navi Veloci      GNV Hoverspeed              HOV Moby Lines              MBL Seafrance               SEA Smyril Line             SMY SNCM                    SNC Stena Line              STE TT Line                 TTL
	CompanyCode string `xml:"companyCode"`

	// This data element is used to convey the company name of a company
	CompanyName string `xml:"companyName,omitempty"`  // minOccurs="0"
}

type CompanyInformationType_25420S struct {

	// This data element is used to qualify the company code, to identify the industry business it belongs.
	TravelSector string `xml:"travelSector"`

	// This data element is used to convey the company code of a company
	CompanyCode string `xml:"companyCode"`

	// This data element is used to convey the company name of a company
	CompanyName string `xml:"companyName,omitempty"`  // minOccurs="0"
}

type CompanyInformationType_26258S struct {

	// This data element is used to qualify the company code, to identify the industry business it belongs.
	TravelSector string `xml:"travelSector,omitempty"`  // minOccurs="0"

	// This data element is used to convey the context in which the code applies
	CompanyCodeContext string `xml:"companyCodeContext,omitempty"`  // minOccurs="0"

	// This data element is used to convey the company code of a company
	CompanyCode string `xml:"companyCode"`

	// This data element is used to convey the company name of a company
	CompanyName string `xml:"companyName,omitempty"`  // minOccurs="0"
}

type CompanyInformationType_83550S struct {

	// This data element is used to qualify the company code, to identify the industry business it belongs.
	TravelSector string `xml:"travelSector"`

	// This data element is used to convey the company code of a company
	CompanyCode string `xml:"companyCode"`
}

type CompanyInformationType_8953S struct {

	// This data element is used to qualify the company code, to identify the industry business it belongs.
	TravelSector string `xml:"travelSector"`

	// This data element is used to convey the context in which the code applies
	CompanyCodeContext string `xml:"companyCodeContext"`

	// This data element is used to convey the company code of a non-air company
	CompanyCode string `xml:"companyCode"`

	// This data element is used to convey the company name of a non-air company
	CompanyName string `xml:"companyName"`
}

type CompanyInformationType_94554S struct {

	// This data element is used to convey the company code of a company  Ex:  AF for Air France MIL for millenium foundation
	CompanyCode string `xml:"companyCode,omitempty"`  // minOccurs="0"

	// This data element is used to convey the numeric merchant ID.
	CompanyNumericCode string `xml:"companyNumericCode,omitempty"`  // minOccurs="0"
}

type CompensationType struct {

	// Compensation details
	CompensationDetails *CardValidityType `xml:"compensationDetails"`
}

type ConsumerReferenceIdentificationTypeI struct {

	// Indicator - 1  for ID (customer number) - CD for CD (customer discount number)
	ReferenceQualifier string `xml:"referenceQualifier"`

	// Either the ID or CD number
	ReferenceNumber string `xml:"referenceNumber"`
}

type ConsumerReferenceInformationTypeI struct {

	// Consumer reference information
	CustomerReferences []*ConsumerReferenceIdentificationTypeI `xml:"customerReferences"`  // maxOccurs="3"
}

type ContactInformationDetailsTypeU struct {

	// W for party to receive Written confirmation
	PartyQualifier string `xml:"partyQualifier"`

	// FAX number or E-Mail address
	ComAddress string `xml:"comAddress"`

	// type of medium
	ComChannelQualifier string `xml:"comChannelQualifier"`
}

type ContactInformationTypeU struct {

	// This composite is used to convey the E-mail address or FAX number to be used when a written confirmation is needed.
	ContactInformation []*ContactInformationDetailsTypeU `xml:"contactInformation,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type CountryCodeListType struct {

	// ISO country code of the DESTINATION of the trip.
	DestinationCountryCode []string `xml:"destinationCountryCode"`  // maxOccurs="198"
}

type CountrySubEntityDetailsTypeU struct {

	// 84: state
	Qualifier string `xml:"qualifier"`

	// Region or State of the given address
	Name string `xml:"name"`
}

type CountrySubEntityDetailsTypeU_198213C struct {

	// State Code. Mandatory if CountryCode is US, CA, AU.
	Code string `xml:"code"`
}

type CountrySubEntityDetailsTypeU_198229C struct {

	// State code
	Code string `xml:"code,omitempty"`  // minOccurs="0"
}

type CountrydescriptionType struct {

	// To specify the destination zone.
	GeographicalZone string `xml:"geographicalZone,omitempty"`  // minOccurs="0"

	// To specify the countries but in a coded way. up to 198 repetitions as we can have 99 segments in the PNR
	CountryCode []string `xml:"countryCode,omitempty"`  // minOccurs="0" maxOccurs="198"
}

type CreditCardDataGroupType struct {

	// will convey all the data related to the credit card
	CreditCardDetails *CreditCardDataType `xml:"creditCardDetails"`

	// will convey both the CVV and the Credit card number Ids stored in the fortknox Database  it could also be used to store identifiers from external Tokenization Service Provider (TSP).
	FortknoxIds []*ReferenceInformationTypeI_94503S `xml:"fortknoxIds,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Contains card holder's address information.
	CardHolderAddress *AddressType `xml:"cardHolderAddress,omitempty"`  // minOccurs="0"
}

type CreditCardDataType struct {

	// Credit Card information
	CcInfo *CreditCardInformationType `xml:"ccInfo,omitempty"`  // minOccurs="0"
}

type CreditCardInformationType struct {

	// Vendor code (VI,CA,AX.)
	VendorCode string `xml:"vendorCode,omitempty"`  // minOccurs="0"

	// may contain CC sub Types. eg: Maestro or Solo cards
	VendorCodeSubType string `xml:"vendorCodeSubType,omitempty"`  // minOccurs="0"

	// Card number  Card number
	CardNumber string `xml:"cardNumber,omitempty"`  // minOccurs="0"

	// Conveys the security ID of the Credit Card (CVV,CVV2), 3-4 digits stored on the back of the card
	SecurityId string `xml:"securityId,omitempty"`  // minOccurs="0"

	// Expiry date :  format    MMYY
	ExpiryDate *int32 `xml:"expiryDate,omitempty"`  // minOccurs="0"

	// This field indicates the date the Credit Card was issued. This data is present in case of (UK) maestro cards.
	StartDate *int32 `xml:"startDate,omitempty"`  // minOccurs="0"

	// This field indicates the date the Credit Card will not be valid anymore This data is present in case of (UK) maestro cards. May be different from the expiry date
	EndDate *int32 `xml:"endDate,omitempty"`  // minOccurs="0"

	// Conveys Credit card holder's name, as written on the card
	CcHolderName string `xml:"ccHolderName,omitempty"`  // minOccurs="0"

	// will contain the code of the bank that issued the credit card
	IssuingBankName string `xml:"issuingBankName,omitempty"`  // minOccurs="0"

	// CC country of issuance details
	CardCountryOfIssuance string `xml:"cardCountryOfIssuance,omitempty"`  // minOccurs="0"

	// This is the Credit Card Issue number. This represents the number of time a card has been issued.  1 is for the first time then in case of card renewal or card loss this issue number will be increased Today this is applicable to maestro cards.
	IssueNumber *int32 `xml:"issueNumber,omitempty"`  // minOccurs="0"

	// Will convey the full name of the institution that issued he credit card
	IssuingBankLongName string `xml:"issuingBankLongName,omitempty"`  // minOccurs="0"

	// Stores the CC track 1 information (base64 encoded)
	Track1 string `xml:"track1,omitempty"`  // minOccurs="0"

	// Stores the CC track 2 information (base64 encoded)
	Track2 string `xml:"track2,omitempty"`  // minOccurs="0"

	// Stores the CC track 3 information (base64 encoded)
	Track3 string `xml:"track3,omitempty"`  // minOccurs="0"

	// Stores the CC pin code information
	PinCode string `xml:"pinCode,omitempty"`  // minOccurs="0"

	// All the tracks of a swipe credit card are contained here as one block.
	RawTrackData string `xml:"rawTrackData,omitempty"`  // minOccurs="0"
}

type CreditCardInformationTypeU struct {

	// Credit card name
	Name string `xml:"name"`

	// Credit card number
	CardNumber int32 `xml:"cardNumber"`

	// Credit card Expire date
	ExpireDate string `xml:"expireDate"`
}

type CreditCardSecurityType struct {

	// Conveys all data of authentication process. Only used today for "Verified by Visa" process
	AuthenticationDataDetails *AuthenticationDataType `xml:"authenticationDataDetails,omitempty"`  // minOccurs="0"
}

type CreditCardStatusGroupType struct {

	// This segment is used to store specific data of links following ISO8583 standard.
	AuthorisationSupplementaryData *SpecificVisaLinkCreditCardInformationType `xml:"authorisationSupplementaryData"`

	// will convey the approval code/source
	ApprovalDetails *GenericAuthorisationResultType `xml:"approvalDetails,omitempty"`  // minOccurs="0"

	// This segment conveys date and time information.  You can specify the time mode used (GMT, UTC or Local)and what for it refers.  - Transmission date and time This contains the date and time the request was submitted to the link (Visa, Nedbank...field 7). GMT can be used.  - Local transaction date and time Date and time when Amadeus builds the authorization message (local according to the point of sale)(Visa, Nedbank...field 12/13)  - Transaction receipt date and time date and time when amadeus receives the authorization message.
	LocalDateTime []*StructuredDateTimeInformationType_94516S `xml:"localDateTime,omitempty"`  // minOccurs="0" maxOccurs="3"

	// Transaction Information: - type of authorization message submit for the given FOP - bulk, superbulk, no bulk process - STAN number (identifying a pair of Credit Card authorization request/response).
	AuthorisationInformation *TransactionInformationForTicketingType `xml:"authorisationInformation,omitempty"`  // minOccurs="0"

	// This group contains all data about the customer's browser.
	BrowserData *BrowserData `xml:"browserData,omitempty"`  // minOccurs="0"

	// this group will convey all the 3DS related data
	TdsInformation *ThreeDomainSecureGroupType `xml:"tdsInformation,omitempty"`  // minOccurs="0"

	// This will allow the transmission of credit card data.
	CardSupplementaryData []*AttributeType_94514S `xml:"cardSupplementaryData,omitempty"`  // minOccurs="0" maxOccurs="99"

	// will convey the various sub status that can be associated to a credit card payment CVV, AVS, AUT, ATN....
	TransactionStatus []*ErrorGroupType `xml:"transactionStatus,omitempty"`  // minOccurs="0" maxOccurs="7"
}

type BrowserData struct {

	// This segment contains data about the customer's browser :  0 PC (HTML) 1 Mobile Internet Device (WML)
	BrowserProperties *BrowserInformationType `xml:"browserProperties"`

	// Contains in freeflow format data about the customer's browser. - userAgent - acceptHeaders This entities are independantly optional.
	FreeFlowBrowserData []*FreeTextInformationType_94526S `xml:"freeFlowBrowserData,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type CreditCardType struct {

	// credit card company code
	CreditCardCompany string `xml:"creditCardCompany"`

	// credit card number
	CreditCardNumber string `xml:"creditCardNumber,omitempty"`  // minOccurs="0"

	// expiration date
	ExpirationDate *int32 `xml:"expirationDate,omitempty"`  // minOccurs="0"
}

type CruiseBusinessDataType struct {

	// Details of sailing ship for the sailing trip.  Each cruise provider has a ship name table in the Amadeus system. This table is used for converting ship codes in ship names and vice-versa. Since both information are stored in the Cruise segment of the PNR, no DB access is  necessary for the PNRACC processing.
	SailingShipInformation *ShipIdentificationType_8952S `xml:"sailingShipInformation"`

	// Details of the cruise line provider for the sailing trip.
	SailingProviderInformation *CompanyInformationType_8953S `xml:"sailingProviderInformation"`

	// Details of embarkation and disembarkation ports for the sailing trip.  The codes sent by the cruise providers can be non-Iata codes.
	SailingPortsInformation *PlaceLocationIdentificationTypeU_8954S `xml:"sailingPortsInformation"`

	// Details of the departure and arrival dates of the sailing trip.  The cruise segment in the PNR actually stores the departure date and the duration length in days. For the PNRACC 4.1 process, the arrival date is re-calculated.
	SailingDateInformation *StructuredPeriodInformationType_8955S `xml:"sailingDateInformation"`

	// Details of passengers for the sailing trip.  For a cruise booking, the passenger names elements from the PNR can be different from the passengers in the cruise segment. They are identical at booking creation time. But the cruise providers allow adding passenger name(s) to an existing booking. That is not possible in an Amadeus PNR. Therefore, the name information had to be stored in the cruise segment itself.
	PassengerInfo []*TravellerInformationTypeI_8956S `xml:"passengerInfo"`  // maxOccurs="9"

	// Booking information, including confirmation and cancellation number, and a flag telling where the booking has been originally created.
	BookingDetails *BookingDetails `xml:"bookingDetails,omitempty"`  // minOccurs="0"

	// Booking Date.
	BookingDate *StructuredDateTimeInformationType_20645S `xml:"bookingDate"`

	// Details of the sailing group code for the sailing trip.
	SailingGroupInformation *ItemReferencesAndVersionsType_9271S `xml:"sailingGroupInformation,omitempty"`  // minOccurs="0"
}

type BookingDetails struct {

	// Details of the booking references for the sailing trip.  These references are returned by the cruise provider at booking creation time or at booking cancellation time.  Note that as re-instate of a cruise booking is possible even several days after cancellation. Therefore, when a cruise booking is cancelled, the segment is kept in the PNR and the status updated to HX.
	CruiseBookingReferenceInfo *ReservationControlInformationTypeI_8957S `xml:"cruiseBookingReferenceInfo"`

	// Company in which the booking is created: Amadeus or external.
	BookingCompany *CompanyInformationType_26258S `xml:"bookingCompany,omitempty"`  // minOccurs="0"
}

type CustomerTransactionDataType struct {

	// Point of sell details
	Pos *PointOfSaleDataTypeI `xml:"pos"`

	// flight suplementary data
	Flight *OtherSegmentDataTypeI `xml:"flight"`

	// CONNECTION NUMBER
	Connection *int32 `xml:"connection,omitempty"`  // minOccurs="0"

	// Codeshare flight details
	CodeShare *CodeshareFlightDataTypeI `xml:"codeShare,omitempty"`  // minOccurs="0"
}

type DataInformationTypeI struct {

	// Animal type.
	Indicator string `xml:"indicator"`

	// Number of animals of the specified category.
	Value *int32 `xml:"value,omitempty"`  // minOccurs="0"
}

type DataTypeInformationTypeI struct {

	// Type of data.
	Type string `xml:"type"`
}

type DateAndTimeDetailsTypeI struct {

	// Seat SSR : Date of change of gauge. Group seat SSR  : Date of change of gauge. MCO element : Date.
	FirstDate string `xml:"firstDate"`

	// MCO element : ARC carrier code code.
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	// MCO element : ARC city code.
	LocationIdentification string `xml:"locationIdentification,omitempty"`  // minOccurs="0"
}

type DateAndTimeDetailsTypeI_56946C struct {

	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Inf/Child date of birth
	Date string `xml:"date,omitempty"`  // minOccurs="0"
}

type DateAndTimeInformationType struct {

	// DATE AND TIME DETAILS
	DateAndTimeDetails *DateAndTimeDetailsTypeI_56946C `xml:"dateAndTimeDetails,omitempty"`  // minOccurs="0"
}

type DateAndTimeInformationTypeI struct {

	// Date and Time details for flight movements
	DateAndTime *DateAndTimeDetailsTypeI `xml:"dateAndTime"`
}

type DateRangeType struct {

	// In range [1-4]
	DateRangeNum *int32 `xml:"dateRangeNum,omitempty"`  // minOccurs="0"
}

type DetailedPaymentDataType struct {

	// This segment will convey the type of the FOP. Exple : CC credit card CA cash CH cheque WW web
	FopInformation *FormOfPaymentType `xml:"fopInformation"`

	// will allow the usage of FOP segment as trigger for GASS and GIVR groups
	Dummy *DummySegmentTypeI `xml:"dummy"`

	// This group will convey the detailed status of the credit card payment
	CreditCardDetailedData *CreditCardStatusGroupType `xml:"creditCardDetailedData,omitempty"`  // minOccurs="0"
}

type DeviceControlDetailsType struct {

	// Stores the identification of the device.
	DeviceIdentification *IdentificationNumberTypeI `xml:"deviceIdentification,omitempty"`  // minOccurs="0"
}

type DiagnosisTypeU struct {

	// This data element can convey either an agency accounting or a billing number.
	BillingDetails string `xml:"billingDetails"`

	// This data element is used to specify the type of billable information that could be found in this segment
	BillingQualifier string `xml:"billingQualifier"`
}

type DiningIdentificationType struct {

	// meal plan information (HALFBOARD, BREAKFAST ...)
	DiningDescription string `xml:"diningDescription"`
}

type DiningInformationType struct {

	// Conveys dining information
	DiningIdentification *DiningIdentificationType `xml:"diningIdentification"`
}

type DiscountInformationDetailsType struct {

	// Promotion code used to define redemption/upgrade price in miles
	DiscountCode string `xml:"discountCode"`
}

type DiscountInformationType struct {

	// Contains the discount code
	DiscountDetails *DiscountInformationDetailsType `xml:"discountDetails"`
}

type DistributionChannelType struct {

	// This field is used to indicate the type of channel used for authorization process: e-commerce (web / Internet), MOTO (Mail Order / telephone Order), Face to face ...) Example:  05 for API 0=MOTO (Mail Order / Telephone Order) 1=e-Commerce (Internet)
	DistributionChannelField int32 `xml:"distributionChannelField"`

	// Subgroup field.
	SubGroup *int32 `xml:"subGroup,omitempty"`  // minOccurs="0"

	// Access Type.
	AccessType *int32 `xml:"accessType,omitempty"`  // minOccurs="0"
}

type DocumentDetailsType struct {

	// Document type: PT for passport, VS for visa.
	Type string `xml:"type"`

	// Passport number.
	Number string `xml:"number"`

	// Country where the document has been issued.
	CountryOfIssue string `xml:"countryOfIssue,omitempty"`  // minOccurs="0"

	// Expiry date of the document. YYYYMMDD
	ExpiryDate string `xml:"expiryDate,omitempty"`  // minOccurs="0"

	// Date of issue of the document. YYYYMMDD
	IssueDate string `xml:"issueDate,omitempty"`  // minOccurs="0"
}

type DocumentDetailsTypeI struct {

	// To convey the document number
	Number string `xml:"number,omitempty"`  // minOccurs="0"

	// To convey if the document has been printed
	Status string `xml:"status"`

	// To convey the date of the impression.
	Date string `xml:"date,omitempty"`  // minOccurs="0"
}

type DocumentDetailsTypeI_19732C struct {

	// documentNumber
	Number *int32 `xml:"number,omitempty"`  // minOccurs="0"

	// Status Code
	Status string `xml:"status"`
}

type DocumentInformationDetailsTypeI struct {

	// To convey the printing results.
	DocumentDetails *DocumentDetailsTypeI `xml:"documentDetails"`
}

type DocumentInformationDetailsTypeI_9936S struct {

	// documentDetails
	DocumentDetails *DocumentDetailsTypeI_19732C `xml:"documentDetails"`
}

type DocumentInformationTypeU struct {

	// Document type being provided: PP: Passport DL: Driving License NI: National Id. card ID: Local Id. Document
	TypeOfDocument string `xml:"typeOfDocument"`

	// Document Number
	DocumentNumber string `xml:"documentNumber"`

	// Country code where document has been issued
	CountryOfIssue string `xml:"countryOfIssue,omitempty"`  // minOccurs="0"
}

type DummySegmentTypeI struct {
}

type ElementManagementSegmentType struct {

	// Action to perform (When a PNR segment/element is transmitted) .  IF for Information only (Value by default, Code used in a Server response)
	Status string `xml:"status,omitempty"`  // minOccurs="0"

	// Reference details
	Reference *ReferencingDetailsType_127526C `xml:"reference,omitempty"`  // minOccurs="0"

	// PNR segment or element name
	SegmentName string `xml:"segmentName,omitempty"`  // minOccurs="0"

	// PNR segment/element 'line' number attributed by the Server
	LineNumber *int32 `xml:"lineNumber,omitempty"`  // minOccurs="0"
}

type ElementManagementSegmentType_127983S struct {

	// reference of the element
	ElementReference *ReferencingDetailsType_127526C `xml:"elementReference,omitempty"`  // minOccurs="0"

	// PNR segment or element name
	SegmentName string `xml:"segmentName,omitempty"`  // minOccurs="0"

	// PNR segment/element 'line' number.
	LineNumber *int32 `xml:"lineNumber,omitempty"`  // minOccurs="0"
}

type ElementManagementSegmentType_83559S struct {

	// Reference details
	Reference *ReferencingDetailsType_127526C `xml:"reference"`
}

type EquipmentDetailsTypeU struct {

	// equipment type
	Type string `xml:"type"`

	// equipment details
	SizeTypeDetails *EquipmentTypeAndSizeTypeU `xml:"sizeTypeDetails"`
}

type EquipmentTypeAndSizeTypeU struct {

	// equipment description
	Description string `xml:"description"`
}

type ErrorGroupType struct {

	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType_94519S `xml:"errorOrWarningCodeDetails"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType_94495S `xml:"errorWarningDescription,omitempty"`  // minOccurs="0"
}

type FLIXType struct {

	// provides the Flix and Source Types. The Codes FX, LX or FD can be used to specify the Flix-Type. The codes USR or GUI can be used to specify the Data Source
	FlixAndSourceTypes *ItemDescriptionType `xml:"flixAndSourceTypes"`

	FlixComment *FreeTextInformationType `xml:"flixComment,omitempty"`  // minOccurs="0"

	AirportGroup *AirportGroup `xml:"airportGroup,omitempty"`  // minOccurs="0"
}

type AirportGroup struct {

	// Only used for Flix-LX or Flix-Disruption: provides the code of the Impacted Airport
	ImpactedAirport *TerminalTimeInformationTypeS `xml:"impactedAirport"`
}

type FOPRepresentationType struct {

	// will convey all the data related to the various codes used by the FOP package, billing, ETS...
	FopPNRDetails *TicketingFormOfPaymentType `xml:"fopPNRDetails"`

	// Conveys the sequence number of the Form of Payment in the FP Line. It must be set to 1 if there is only 1 FOP in the FOP  Old FOP are referenced with sequence number: 0
	FopSequenceNumber *SequenceDetailsTypeU_94494S `xml:"fopSequenceNumber,omitempty"`  // minOccurs="0"

	// This segment conveys Form of Payment FreeText.   Old FOP(s) are considered as one freeflow text even if there is more than one old form of payment.  e.g.: FP O/CA+CCVI+/CH CA and CCVI are considered as freeflow text.
	FopFreeflow *FreeTextInformationType_94495S `xml:"fopFreeflow,omitempty"`  // minOccurs="0"

	// will convey the switches and data associated to the FOP table
	PnrSupplementaryData []*PNRSupplementaryDataType `xml:"pnrSupplementaryData,omitempty"`  // minOccurs="0" maxOccurs="2"

	// will contain all the data related to the payment transaction
	PaymentModule *PaymentGroupType `xml:"paymentModule,omitempty"`  // minOccurs="0"
}

type FareBasisCodesLineType struct {

	// Fare element information
	FareElement []*FareElementType `xml:"fareElement"`  // maxOccurs="28"
}

type FareCategoryCodesTypeI struct {

	// Rate Code (code set list not used)
	FareType string `xml:"fareType"`
}

type FareDataType struct {

	// Issue identifier
	IssueIdentifier string `xml:"issueIdentifier"`

	// To specify the type of monetary amount, the amount and the currency code
	MonetaryInfo []*MonetaryInformationDetailsTypeI_8308C `xml:"monetaryInfo"`  // maxOccurs="3"

	// Tax fields
	TaxFields []*TaxFieldsType `xml:"taxFields,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FareElementType struct {

	// Contains primary code of the fare element
	PrimaryCode string `xml:"primaryCode"`

	// Connection indicator
	Connection string `xml:"connection,omitempty"`  // minOccurs="0"

	// Not valid before
	NotValidBefore string `xml:"notValidBefore,omitempty"`  // minOccurs="0"

	// Not valid after
	NotValidAfter string `xml:"notValidAfter,omitempty"`  // minOccurs="0"

	// Baggage allowance
	BaggageAllowance string `xml:"baggageAllowance,omitempty"`  // minOccurs="0"

	// Fare basis
	FareBasis string `xml:"fareBasis,omitempty"`  // minOccurs="0"

	// Ticket designator
	TicketDesignator string `xml:"ticketDesignator,omitempty"`  // minOccurs="0"
}

type FareQualifierDetailsTypeI struct {

	// Rate Code Information
	FareCategories *FareCategoryCodesTypeI `xml:"fareCategories"`
}

type FerryAccomodationPackageDescriptionType struct {

	// This segment conveys the package code.
	PackageCode *ProductInformationTypeI `xml:"packageCode"`

	// This segment conveys the hotel code the of the accomodation package to which it is attached.
	HotelInformation *HotelPropertyType_26378S `xml:"hotelInformation,omitempty"`  // minOccurs="0"

	// This segment conveys the check-in date and time for the accomodation package to which it is attached.
	HotelCheckInInformation *StructuredDateTimeInformationType_24436S `xml:"hotelCheckInInformation,omitempty"`  // minOccurs="0"

	// This segment is used to convey the hotel area code information.
	AreaCodeInfo *PlaceLocationIdentificationTypeU_24573S `xml:"areaCodeInfo,omitempty"`  // minOccurs="0"

	// This segment is used to give the number of nights spent in the accomodation package to which it is attached.
	NumberOfNights *NumberOfUnitsType `xml:"numberOfNights,omitempty"`  // minOccurs="0"

	// This segment is used to convey the price of the accomodation to which it is attached.
	HotelItemPrice *TariffInformationTypeU `xml:"hotelItemPrice,omitempty"`  // minOccurs="0"

	// This segment is used to give details about the rooms (if any) lined to the accomodation package.
	RoomInfoGroup []*RoomInfoGroup `xml:"roomInfoGroup,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type RoomInfoGroup struct {

	// This segment is used to describe the room to which it is attached.
	RoomDetailsInformation *HotelRoomType_20177S `xml:"roomDetailsInformation"`

	// This segment is used to convey the number of instances of the room to which it is attached.
	NumberOfRooms *NumberOfUnitsType `xml:"numberOfRooms,omitempty"`  // minOccurs="0"
}

type FerryBookingDescriptionType struct {

	// This segment is used to identify the ferry target provider for the message and is leading the description group for the ferry provider booking.
	FerryProviderInformation *CompanyInformationType_20151S `xml:"ferryProviderInformation"`

	// This group describes the ferry booking itinerary. It contains the ferry sailing leg information.
	ItineraryInfoGroup *FerryLegDescriptionType `xml:"itineraryInfoGroup"`

	// This group describes the accomodation (hotel) package attached to the booking.
	AccomodationPackageInfoGroup []*FerryAccomodationPackageDescriptionType `xml:"accomodationPackageInfoGroup,omitempty"`  // minOccurs="0" maxOccurs="4"

	// This segment conveys the ferry booking number information.
	BookingNumberInformation *ReservationControlInformationTypeI_20153S `xml:"bookingNumberInformation"`
}

type FerryLegDescriptionType struct {

	// Conveys the sailing details for an itinerary leg.
	SailingDetails *TravelProductInformationTypeU `xml:"sailingDetails"`

	// Conveys the ship code and ship name.
	ShipDescription *ShipIdentificationType `xml:"shipDescription,omitempty"`  // minOccurs="0"

	// This segment conveys the check-in time for the ferry sailing leg to which it is attached.
	SailingLegCheckInInformation *StructuredDateTimeInformationType_21109S `xml:"sailingLegCheckInInformation,omitempty"`  // minOccurs="0"

	// Conveys the list of passengers associated to the ferry leg.
	PassengerAssociation *ReferenceInformationTypeI_25132S `xml:"passengerAssociation,omitempty"`  // minOccurs="0"

	// Conveys the price information per leg per passenger.
	PriceInfoGroup []*PriceInfoGroup `xml:"priceInfoGroup,omitempty"`  // minOccurs="0" maxOccurs="9"

	// This group describes the list of vehicles attached to the linked sailing leg.
	VehicleInfoGroup []*VehicleInfoGroup `xml:"vehicleInfoGroup,omitempty"`  // minOccurs="0" maxOccurs="5"

	// This segment describes the on-board service(s) in the linked sailing leg.
	ServiceInfoGroup []*ServiceInfoGroup `xml:"serviceInfoGroup,omitempty"`  // minOccurs="0" maxOccurs="18"

	// This group is used to describe the animals linked to the ferry booking.
	AnimalInfoGroup []*AnimalInfoGroup `xml:"animalInfoGroup,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type AnimalInfoGroup struct {

	// This segment conveys the type of animal.
	AnimalInformation *SpecificDataInformationTypeI `xml:"animalInformation"`

	// This segment conveys the price per animal of the same type.
	AnimalRoutePrice *TariffInformationTypeU `xml:"animalRoutePrice,omitempty"`  // minOccurs="0"
}

type PriceInfoGroup struct {

	// This segment conveys the route price information for the passenger it is linked to.
	RoutePriceInformation *TariffInformationTypeU `xml:"routePriceInformation"`

	// This segment describes the passenger category type.
	PassengerCategoryType *AttributeTypeU_24552S `xml:"passengerCategoryType"`

	// This segment is used to convey the number of passengers to which the price applies.
	NumberOfPassengers *NumberOfUnitsType `xml:"numberOfPassengers"`
}

type ServiceInfoGroup struct {

	// This segment describes the on-board service.
	ServiceInformation *AttributeTypeU `xml:"serviceInformation"`

	// This segment conveys the number of services of the attached service.
	NumberOfServices *NumberOfUnitsType `xml:"numberOfServices,omitempty"`  // minOccurs="0"

	// This segment conveys the price per unit of the attached service.
	ServiceRoutePrice *TariffInformationTypeU `xml:"serviceRoutePrice,omitempty"`  // minOccurs="0"
}

type VehicleInfoGroup struct {

	// This segment conveys the description of a vehicle.
	VehicleInformation *VehicleTypeU `xml:"vehicleInformation"`

	// This segment is used to convey the number of bicycles associated to a ferry booking. Note: this segment is ignored if the vehicule description is not "bicycle".
	NumberOfBicycles *NumberOfUnitsType `xml:"numberOfBicycles,omitempty"`  // minOccurs="0"

	// This segment holds the price per vehicle.
	VehicleRoutePrice *TariffInformationTypeU `xml:"vehicleRoutePrice,omitempty"`  // minOccurs="0"
}

type FlightSegmentDetailsTypeI struct {
}

type FormOfPaymentDetailsType struct {

	// Generic type of the Mean of Payment used : CC credit Card CA cash CH cheque WW web payment... INV invoice
	Type string `xml:"type"`
}

type FormOfPaymentDetailsTypeI struct {

	// Fop type (Cash, Credit card...)
	Type string `xml:"type"`

	// Credit card vendor code
	VendorCode string `xml:"vendorCode,omitempty"`  // minOccurs="0"

	// Credit card number
	CreditCardNumber string `xml:"creditCardNumber,omitempty"`  // minOccurs="0"

	// expiry date (MMYY)
	ExpiryDate string `xml:"expiryDate,omitempty"`  // minOccurs="0"

	// FOP purpose
	ExtendedPayment string `xml:"extendedPayment,omitempty"`  // minOccurs="0"

	// Unstructured fop layout (used for Voucher print purpose or guarantee details).
	FopFreeText string `xml:"fopFreeText,omitempty"`  // minOccurs="0"
}

type FormOfPaymentDetailsTypeI_20667C struct {

	// Reporting code
	Type string `xml:"type"`

	// Currency code per form of payment
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"

	// Form of payment amount
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Vendor code of the credit card. ex: VI
	VendorCode string `xml:"vendorCode,omitempty"`  // minOccurs="0"

	// Account number
	CreditCardNumber string `xml:"creditCardNumber,omitempty"`  // minOccurs="0"

	// Expiration date
	ExpiryDate string `xml:"expiryDate,omitempty"`  // minOccurs="0"

	// Approval code
	ApprovalCode string `xml:"approvalCode,omitempty"`  // minOccurs="0"

	// Source of approval code
	SourceOfApproval string `xml:"sourceOfApproval,omitempty"`  // minOccurs="0"

	// Authorised amount
	AuthorisedAmount *float64 `xml:"authorisedAmount,omitempty"`  // minOccurs="0"

	// Address verification code
	AddressVerification string `xml:"addressVerification,omitempty"`  // minOccurs="0"

	// Customer file reference
	CustomerAccount string `xml:"customerAccount,omitempty"`  // minOccurs="0"

	// Extended payment code
	ExtendedPayment string `xml:"extendedPayment,omitempty"`  // minOccurs="0"

	// not used
	FopFreeText string `xml:"fopFreeText,omitempty"`  // minOccurs="0"

	// Credit card corporate contract
	MembershipStatus string `xml:"membershipStatus,omitempty"`  // minOccurs="0"

	// Credit card transaction information
	TransactionInfo string `xml:"transactionInfo,omitempty"`  // minOccurs="0"
}

type FormOfPaymentDetailsTypeI_52343C struct {

	// Fop type (Cash, Credit card...)
	Type string `xml:"type"`

	// Format limitations: an..3
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"

	// Credit card vendor code
	VendorCode string `xml:"vendorCode,omitempty"`  // minOccurs="0"

	// Credit card number
	CreditCardNumber string `xml:"creditCardNumber,omitempty"`  // minOccurs="0"

	// expiry date (MMYY)
	ExpiryDate string `xml:"expiryDate,omitempty"`  // minOccurs="0"

	// FOP purpose
	ExtendedPayment string `xml:"extendedPayment,omitempty"`  // minOccurs="0"

	// Unstructured fop layout (used for Voucher print purpose or guarantee details).
	FopFreeText string `xml:"fopFreeText,omitempty"`  // minOccurs="0"
}

type FormOfPaymentInformationType struct {

	// Format key that identify the FOP within a FOP table. (CCVI, ...)
	FopCode string `xml:"fopCode,omitempty"`  // minOccurs="0"

	// Name of the FOP map table used in order to validate the FP element.
	FopMapTable string `xml:"fopMapTable,omitempty"`  // minOccurs="0"

	// This corresponds to the fop billing code (CASH CA / Credit CC). This is only used in case of a MS reporting code. (it corresponds to XX of @FPMSXX tag of TPF tables)
	FopBillingCode string `xml:"fopBillingCode,omitempty"`  // minOccurs="0"

	// Fop is a old / new fop.
	FopStatus string `xml:"fopStatus,omitempty"`  // minOccurs="0"

	// Corresponds to the EDIFACT code.  This enables to identify the type of FOP that will be added in case of a structured EDIFACT (i.e. via PNRADD) addition of the FOP. (it corresponds to the @EDI tag of TPF tables) Here is an example: Customer is eager to add a structured cash FOP using an EDIFACT message. The fopEdiCode will be filled with CA which means cash. Then in the FOP table in charge of validating free flow and generating FOP free flow, the system will try to look for the FOP map having CA as fop EDI code. If we are in an Air France (AF) ATO/CTO: the system will get FP CA.... If we are in an Iberia (IB) ATO/CTO: the system will get FP CASH,.... If we are in an United Airline (UA) ATO/CTO: the system will get FP S.... ...  (@EDI value)
	FopEdiCode string `xml:"fopEdiCode,omitempty"`  // minOccurs="0"

	// This corresponds to the fop code which is used on reporting side.  (XX value of @FPXXxx)
	FopReportingCode string `xml:"fopReportingCode,omitempty"`  // minOccurs="0"

	// This is the FOP printed code  (@PR value)
	FopPrintedCode string `xml:"fopPrintedCode,omitempty"`  // minOccurs="0"

	// This is the FOP electronic ticketing code. This is used to classify any FOP from the FOP table and also to determine how the FOP should be transmitted to the airline concerned. Based on this, the absence of the switch would make the FOP disallowed for ETKT, for National System Ticketing Server Travel Agency locations and all Central Ticketing offices  (@ET value)
	FopElecTicketingCode string `xml:"fopElecTicketingCode,omitempty"`  // minOccurs="0"
}

type FormOfPaymentType struct {

	// Generic status(new/old) and type(cash, cheque, card...) of the MOP
	FormOfPayment *FormOfPaymentDetailsType `xml:"formOfPayment"`
}

type FormOfPaymentTypeI struct {

	// Fop details
	FormOfPayment []*FormOfPaymentDetailsTypeI `xml:"formOfPayment"`  // maxOccurs="2"
}

type FormOfPaymentTypeI_16862S struct {

	// Description of the form of paiement
	FormOfPayment *FormOfPaymentDetailsTypeI_20667C `xml:"formOfPayment,omitempty"`  // minOccurs="0"

	OtherFormOfPayment []*FormOfPaymentDetailsTypeI_20667C `xml:"otherFormOfPayment,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type FormOfPaymentTypeI_29553S struct {

	// FOP details
	FormOfPayment []*FormOfPaymentDetailsTypeI_52343C `xml:"formOfPayment"`  // maxOccurs="2"
}

type FraudScreeningGroupType struct {

	// This data element is used to indicate if risk management must be performed at authorization time: - Y means risk management data will be appended to author; - N means risk management data will not be appended;
	FraudScreening *StatusType_94568S `xml:"fraudScreening"`

	// this segment contains the IP address used in RMM (risk management module or fraud screening)
	IpAdress *DeviceControlDetailsType `xml:"ipAdress,omitempty"`  // minOccurs="0"

	// Merchant's website URL.
	MerchantURL *CommunicationContactType `xml:"merchantURL,omitempty"`  // minOccurs="0"

	// will convey either the phone or the email adress of the payer
	PayerPhoneOrEmail []*PhoneAndEmailAddressType_94565S `xml:"payerPhoneOrEmail,omitempty"`  // minOccurs="0" maxOccurs="2"

	// this segment contains the shopper session used in RMM (risk management module)
	ShopperSession *SystemDetailsInfoType_94569S `xml:"shopperSession,omitempty"`  // minOccurs="0"

	// conveys information about payer LastName (surName) and FirstName (givenName)
	PayerName *TravellerInformationType_94570S `xml:"payerName,omitempty"`  // minOccurs="0"

	// stores the payer date of birth
	PayerDateOfBirth *StructuredDateTimeInformationType_94567S `xml:"payerDateOfBirth,omitempty"`  // minOccurs="0"

	// Information about the billing address (can be extracted from the AB PNR element)
	BillingAddress *AddressType `xml:"billingAddress,omitempty"`  // minOccurs="0"

	// Used to store reference information on the payer for fraud screening purpose: social security number driving license information frequent flyer information
	FormOfIdDetails []*ReferenceInfoType_94566S `xml:"formOfIdDetails,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type FreeTextDetailsType struct {

	// text subject qualifier
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// information type
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// status
	Status string `xml:"status,omitempty"`  // minOccurs="0"

	// company id
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	// Language, coded
	Language string `xml:"language,omitempty"`  // minOccurs="0"

	// source, coded
	Source string `xml:"source"`

	// encoding
	Encoding string `xml:"encoding"`
}

type FreeTextDetailsType_142107C struct {

	// text subject qualifier
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Manual : M
	Source string `xml:"source"`

	// encoding
	Encoding string `xml:"encoding"`
}

type FreeTextDetailsType_142141C struct {

	// mutually defined ZZZ
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// AH Browser Accept headers UA Browser User Agent
	InformationType string `xml:"informationType"`

	// Manual : M
	Source string `xml:"source"`

	// ZZZ mutually agreed
	Encoding string `xml:"encoding"`
}

type FreeTextDetailsType_187698C struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..3
	Language string `xml:"language,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Source string `xml:"source"`

	// Format limitations: an..3
	Encoding string `xml:"encoding"`
}

type FreeTextDetailsType_198207C struct {

	// Text qualifier - 3 for literal text
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Information type
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// 1A for Amadeus
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	// ISO language code
	Language string `xml:"language,omitempty"`  // minOccurs="0"

	// Text source Manual or System
	Source string `xml:"source"`

	// Character set
	Encoding string `xml:"encoding"`
}

type FreeTextDetailsType_46357C struct {

	// booking description subject qualifier
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// booking information type
	InformationType string `xml:"informationType"`

	// booking desscription source
	Source string `xml:"source"`

	// booking description encoding information
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {

	FreeTextDetails *FreeTextDetailsType_187698C `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText []string `xml:"freeText"`  // maxOccurs="99"
}

type FreeTextInformationType_136708S struct {

	// Free text type
	FreeTextDetails *FreeTextDetailsType_198207C `xml:"freeTextDetails"`

	// Free text
	FreeText []string `xml:"freeText"`  // maxOccurs="24"
}

type FreeTextInformationType_136715S struct {

	// Free text type
	FreeTextDetails *FreeTextDetailsType_198207C `xml:"freeTextDetails"`

	// 1 or 2 lines of free text
	FreeText []string `xml:"freeText"`  // maxOccurs="2"
}

type FreeTextInformationType_20551S struct {

	// Text attributes
	FreeTextDetails *FreeTextDetailsType_142107C `xml:"freeTextDetails"`

	// SVCs / Service Information (Amtrak). 63 characters maximum length, and a maximum of 5 lines per train segment.
	FreeText []string `xml:"freeText"`  // maxOccurs="5"
}

type FreeTextInformationType_25445S struct {

	// booking description details
	FreeTextDetails *FreeTextDetailsType_46357C `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText []string `xml:"freeText"`  // maxOccurs="2"
}

type FreeTextInformationType_29860S struct {

	// Free text information.
	FreeTextDetails *FreeTextDetailsType_187698C `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText string `xml:"freeText"`
}

type FreeTextInformationType_6235S struct {

	// To convey the type of the freeflow text.
	FreeTextDetails *FreeTextDetailsType_187698C `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText []string `xml:"freeText"`  // maxOccurs="5"
}

type FreeTextInformationType_94495S struct {

	// will contain the FOP free flow text
	FreeTextDetails *FreeTextDetailsType_142107C `xml:"freeTextDetails"`

	// FOP freeflow
	FreeText string `xml:"freeText"`
}

type FreeTextInformationType_94526S struct {

	// will contain the browser information
	FreeTextDetails *FreeTextDetailsType_142141C `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText []string `xml:"freeText"`  // maxOccurs="99"
}

type FreeTextInformationType_94561S struct {

	// will describe the purchase
	FreeTextDetails *FreeTextDetailsType_142107C `xml:"freeTextDetails"`

	// Purchase free text description
	FreeText []string `xml:"freeText"`  // maxOccurs="99"
}

type FreeTextInformationType_9865S struct {

	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	FreeText string `xml:"freeText"`
}

type FreeTextQualificationType struct {

	// Identifies whether the free text is coded or not coded :  3 for Literal text
	SubjectQualifier string `xml:"subjectQualifier"`

	// Information type, coded. see code list
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Transmittable/non-transmittable indicator (S or X).
	Status string `xml:"status,omitempty"`  // minOccurs="0"

	// Airline or system code.
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"
}

type FreeTextQualificationTypeI struct {

	// Identifies whether the free text is coded or not coded  3 for Literal text
	SubjectQualifier string `xml:"subjectQualifier"`

	// Coded text, or specifies type of  info   Surface segment : 2 for Address  or 5 for Telephone nature un known  Cruise segment : P30 for Ship Name
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Company code
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	// ISO code for language of free text
	Language string `xml:"language,omitempty"`  // minOccurs="0"
}

type FreeTextQualificationTypeI_148295C struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..4
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Language string `xml:"language,omitempty"`  // minOccurs="0"
}

type FreeTextQualificationTypeI_185754C struct {

	// Categorise the format of the text (free text, coded,...)
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Provides a code identifying the information (phone, OSI, etc...)
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// The airline code that may be associated to this information
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"
}

type FrequencyDetailsTypeU struct {

	// Indicates number of instalments for the payment
	InstalmentsNumber int32 `xml:"instalmentsNumber"`

	// Indicates frequency of instalments for the payment D daily M monthly W weekly
	InstalmentsFrequency string `xml:"instalmentsFrequency,omitempty"`  // minOccurs="0"

	// Indicates when first instalment should take place
	InstalmentsStartDate string `xml:"instalmentsStartDate,omitempty"`  // minOccurs="0"

	// indicates extended payment start date format
	InstalmentsDatrDateFormat string `xml:"instalmentsDatrDateFormat,omitempty"`  // minOccurs="0"
}

type FrequencyType struct {

	// Indicate if the sequence number represents days of the week or days of the month.
	Qualifier string `xml:"qualifier"`

	// Used to represent days of the week or days of the month. For week : 1 is monday and 7 is sunday. For month : 1 is the first day of the month.
	Value string `xml:"value"`
}

type FrequencyTypeU struct {

	// extended payment characteristics
	ExtendedPaymentDetails *FrequencyDetailsTypeU `xml:"extendedPaymentDetails"`
}

type FrequentFlyerInformationGroupType struct {

	// To specify frequent traveller information
	FrequentTravellerInfo *FrequentTravellerIdentificationCodeType_74327S `xml:"frequentTravellerInfo"`

	// Promotion code used to compute redemption/upgrade price in miles, when applicable
	DiscountInformation *DiscountInformationType `xml:"discountInformation,omitempty"`  // minOccurs="0"

	// Original booking class
	BookingClassInformation *ProductInformationTypeI_73824S `xml:"bookingClassInformation,omitempty"`  // minOccurs="0"
}

type FrequentTravellerIdentificationCodeType struct {

	// Frequent Traveller Info
	AirlineFrequentTraveler *FrequentTravellerIdentificationType_198190C `xml:"airlineFrequentTraveler"`
}

type FrequentTravellerIdentificationCodeType_38226S struct {

	// Airline Frequent Traveller Info
	AirlineFrequentTraveler *FrequentTravellerIdentificationType `xml:"airlineFrequentTraveler"`

	// Alliance Frequent Traveller Info
	AllianceFrequentTraveler *FrequentTravellerIdentificationType_64816C `xml:"allianceFrequentTraveler,omitempty"`  // minOccurs="0"
}

type FrequentTravellerIdentificationCodeType_74327S struct {

	// FREQUENT TRAVELER IDENTIFICATION
	FrequentTraveler *FrequentTravellerIdentificationTypeI `xml:"frequentTraveler"`

	// PRIORITY DETAILS
	PriorityDetails []*PriorityDetailsType `xml:"priorityDetails,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Specify the redemption information
	RedemptionInformation *ProductAccountDetailsTypeI `xml:"redemptionInformation,omitempty"`  // minOccurs="0"
}

type FrequentTravellerIdentificationType struct {

	// Carrier where the FQTV is registered.
	Company string `xml:"company"`

	// Frequent Traveller Reference Number
	MembershipNumber string `xml:"membershipNumber"`

	// To specify a Tier linked to the FQTV
	TierLevel string `xml:"tierLevel,omitempty"`  // minOccurs="0"

	// To specify the Priority of the FQTV.
	PriorityCode string `xml:"priorityCode,omitempty"`  // minOccurs="0"

	// Full Text Tier description ex: EMERALD, SAPPHIRE
	TierDescription string `xml:"tierDescription,omitempty"`  // minOccurs="0"
}

type FrequentTravellerIdentificationTypeI struct {

	// Airline code
	Company string `xml:"company"`

	// Frequent traveler number
	MembershipNumber string `xml:"membershipNumber"`

	// Provide airline customer value of the frequent traveller.
	CustomerValue *int32 `xml:"customerValue,omitempty"`  // minOccurs="0"
}

type FrequentTravellerIdentificationType_198190C struct {

	// Carrier where the FQTV is registered.
	Company string `xml:"company"`

	// Frequent Traveller Reference Number
	MembershipNumber string `xml:"membershipNumber"`
}

type FrequentTravellerIdentificationType_64816C struct {

	// To specify a Tier linked to the FQTV
	TierLevel string `xml:"tierLevel,omitempty"`  // minOccurs="0"

	// To specify the Priority of the FQTV.
	PriorityCode string `xml:"priorityCode,omitempty"`  // minOccurs="0"

	// Full Text Tier description ex: EMERALD, SAPPHIRE
	TierDescription string `xml:"tierDescription,omitempty"`  // minOccurs="0"

	// Alliance name
	CompanyCode string `xml:"companyCode"`
}

type GategoryType struct {

	// A (first) category number
	CategoryNum1 *int32 `xml:"categoryNum1,omitempty"`  // minOccurs="0"

	// [2-16] characters for Special [2-10] characters for Dual
	CategoryName string `xml:"categoryName,omitempty"`  // minOccurs="0"
}

type GeneralOptionInformationType struct {

	// Option type.  hotel/car/cruise/train/insurance options type.  CAR : ACD : Action Code for Display BCS : Billing Control System CWB : Car Warning Banner MKT : Marketing informations OS : Other Services informations  HOTEL : ACD : Action code for Display BCS : Billing Control System TXT : Marketing text NAM : Hotel property name CNM : Hotel chain name MKT : Marketing informations CXL : Cancellation policy DES : Rate description PRI : Pricing information  CRUISE : GPN : Group code. CXN : Booking cancellation number. CFN : Booking confirmation number. NME : Passenger's last name, first name and title.
	Type string `xml:"type"`

	// 1. Hotel segment: for Update option indicator  Y for yes  N for no
	UpdateIndicator string `xml:"updateIndicator,omitempty"`  // minOccurs="0"

	// Free text
	Freetext []string `xml:"freetext,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type GeneralOptionType struct {

	// ONLY ONE OCCURRENCE  Each option is one segment
	OptionDetail *GeneralOptionInformationType `xml:"optionDetail"`
}

type GenericAuthorisationResultType struct {

	// transaction authorization approval data
	ApprovalCodeData *AuthorizationApprovalDataType `xml:"approvalCodeData"`
}

type GenericDetailsTypeI struct {

	// Seat Characteristic
	SeatCharacteristic []string `xml:"seatCharacteristic,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type HotelProductInformationType struct {

	// Property header details
	Property *PropertyHeaderDetailsType `xml:"property,omitempty"`  // minOccurs="0"

	// Room details
	HotelRoom *RoomDetailsType `xml:"hotelRoom,omitempty"`  // minOccurs="0"

	// Rate code
	Negotiated []*RateCodeRestrictedType `xml:"negotiated,omitempty"`  // minOccurs="0" maxOccurs="8"

	// Other information
	OtherHotelInfo *OtherHotelInformationType `xml:"otherHotelInfo,omitempty"`  // minOccurs="0"
}

type HotelPropertyType struct {

	// This composite is used to convey the hotel identifier.
	HotelReference *HotelUniqueIdType `xml:"hotelReference"`

	// This composite is used to convey the hotel name
	HotelName string `xml:"hotelName"`

	// This data element is used to indicates if the hotel is compliant with the fire safety rules.
	FireSafetyIndicator string `xml:"fireSafetyIndicator,omitempty"`  // minOccurs="0"
}

type HotelPropertyType_26378S struct {

	// This composite is used to convey the hotel identification details.
	HotelReference *HotelUniqueIdType_47769C `xml:"hotelReference"`
}

type HotelRoomRateInformationType struct {

	// This data element is used to convey the room type
	RoomType string `xml:"roomType,omitempty"`  // minOccurs="0"

	// This data element is used to convey the hotel rate code
	RatePlanCode string `xml:"ratePlanCode,omitempty"`  // minOccurs="0"

	// This data element is used to convey the hotel rate category code
	RateCategoryCode string `xml:"rateCategoryCode,omitempty"`  // minOccurs="0"

	// This data element is used to indicate if the rate code is a qualified rate code or not.
	RateQualifiedIndic string `xml:"rateQualifiedIndic,omitempty"`  // minOccurs="0"
}

type HotelRoomRateInformationType_46329C struct {

	// This data element is used to convey the room type.
	RoomType string `xml:"roomType"`
}

type HotelRoomType struct {

	// This composite is used to convey the room rate identifier.
	RoomRateIdentifier *HotelRoomRateInformationType `xml:"roomRateIdentifier,omitempty"`  // minOccurs="0"

	// This data element is used to convey the booking code.
	BookingCode string `xml:"bookingCode,omitempty"`  // minOccurs="0"

	// This composite is used to convey the occupancy level of the hotel room.
	GuestCountDetails *NumberOfUnitDetailsTypeI_18670C `xml:"guestCountDetails"`

	// This data element is used to convey the override room type (non-Amadeus room types).
	RoomTypeOverride string `xml:"roomTypeOverride,omitempty"`  // minOccurs="0"
}

type HotelRoomType_20177S struct {

	// This data element is used to convey the override room type (non-Amadeus room types).  NOTE: WAS AN..35 IN STANDRD.
	RoomTypeOverride string `xml:"roomTypeOverride"`
}

type HotelRoomType_25429S struct {

	// This composite is used to convey the room type
	RoomRateIdentifier *HotelRoomRateInformationType_46329C `xml:"roomRateIdentifier"`

	// This data element is used to convey the booking code.
	BookingCode string `xml:"bookingCode,omitempty"`  // minOccurs="0"

	// This composite is used to convey the occupancy level of the hotel room.
	GuestCountDetails *NumberOfUnitDetailsTypeI_46330C `xml:"guestCountDetails,omitempty"`  // minOccurs="0"

	// This data element is used to convey the override room type (non-Amadeus room types).
	RoomTypeOverride string `xml:"roomTypeOverride,omitempty"`  // minOccurs="0"
}

type HotelUniqueIdType struct {

	// To convey the chain code in the property ID
	ChainCode string `xml:"chainCode"`

	// To convey the city code in the hotel Id
	CityCode string `xml:"cityCode"`

	// To convey the property code in the Hotel Id
	HotelCode string `xml:"hotelCode"`
}

type HotelUniqueIdType_47769C struct {

	// This element is used to convey the hotel code. The list of possible values is different for each Ferry provider.
	HotelCode string `xml:"hotelCode"`
}

type IdentificationNumberTypeI struct {

	// will contain the IP adress of the shopper
	Address string `xml:"address"`

	// will contain IP for IP adress
	Qualifier string `xml:"qualifier"`
}

type IndividualPnrSecurityInformationType struct {

	// Returned before End of transaction when retrieving a PNR security element
	Security []*IndividualSecurityType_3194C `xml:"security,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Returned when retrieving a PNR
	SecurityInfo *SecurityInformationType `xml:"securityInfo,omitempty"`  // minOccurs="0"

	// Code as in the display:  G for Amadeus Global Core Office Identification  I for IATA number  P for Pseudo-Office Identification Default is G.
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"
}

type IndividualSecurityType struct {

	// office Id
	Office string `xml:"office"`

	// R for Read  B for Both read and write  N for None
	AccessMode string `xml:"accessMode"`

	// - F for Family identifier
	OfficeIdentifier string `xml:"officeIdentifier,omitempty"`  // minOccurs="0"
}

type IndividualSecurityType_3194C struct {

	// Type of receiver  G: Amadeus Global Core Office Id with possible wild card chars '*' to mask some part(s) of it// I:IATA nb or '*' for all, no wild card char //P:Pseudo-Office Id or '*' for all, no wild card char.
	Identification string `xml:"identification"`

	// R for Read  B for Both read and write  N for None
	AccessMode string `xml:"accessMode"`
}

type InsuranceBusinessDataType struct {

	// To convey the provider code and country. Can be empty in case preferences have been set up.
	ProviderProductDetails *InsuranceProductDetailsType `xml:"providerProductDetails"`

	// provides details about the substitute name or the nanny name.
	SubstiteName []*TravellerInformationTypeI `xml:"substiteName,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Amount that is added to the total premium in case an extrareference is specified.
	ExtraPremium *MonetaryInformationTypeI `xml:"extraPremium,omitempty"`  // minOccurs="0"

	// To convey the products and it's directly related data.
	ProductSection []*ProductSection `xml:"productSection,omitempty"`  // minOccurs="0" maxOccurs="48"

	// contains the different amounts (net premium/taxes/total premium)
	PlanCostInfo *TariffInformationTypeI_22057S `xml:"planCostInfo,omitempty"`  // minOccurs="0"

	// Provides details about the type of plan beeing booked.
	PlanTypeDetails *PlanTypeDetails `xml:"planTypeDetails,omitempty"`  // minOccurs="0"

	// To specify remarks and an emergency contact (phone or name)
	ContactDetails *ContactDetails `xml:"contactDetails,omitempty"`  // minOccurs="0"

	// To specify the address of the subscriber.
	SubscriberAddressSection *SubscriberAddressSection `xml:"subscriberAddressSection,omitempty"`  // minOccurs="0"

	// This is used to convey the different coverages and it's values.
	CoverageDetails *CoverageDetails `xml:"coverageDetails,omitempty"`  // minOccurs="0"

	// to specify a commission.
	ComissionAmount *CommissionInformationType `xml:"comissionAmount,omitempty"`  // minOccurs="0"

	// To convey a structered FOP element.
	InsuranceFopSection *InsuranceFopSection `xml:"insuranceFopSection,omitempty"`  // minOccurs="0"

	// To specify a number which means that the insurance is in a confirmed status.
	ConfirmationNumber *ReservationControlInformationTypeI `xml:"confirmationNumber,omitempty"`  // minOccurs="0"

	// Used to specify the necesary data for pricing
	ProductKnowledge []*ActionDetailsTypeI `xml:"productKnowledge,omitempty"`  // minOccurs="0" maxOccurs="20"

	// to specify to which passenger the insurance is associated: if omitted then it's for all the names in the PNR. The repetition factor is 198 because we can have 99 passengers in a PNR each of them an infant.
	PassengerDetails []*PassengerDetails `xml:"passengerDetails,omitempty"`  // minOccurs="0" maxOccurs="198"

	// To convey information if the document has been printed or not.
	PrintInformation *DocumentInformationDetailsTypeI `xml:"printInformation,omitempty"`  // minOccurs="0"
}

type ContactDetails struct {

	// data to add some comments on the insurance element
	Miscelaneous *MiscellaneousRemarksType_12240S `xml:"miscelaneous"`

	// Used to specify a phone number as an emergency contact
	PhoneNumber *PhoneAndEmailAddressType_32298S `xml:"phoneNumber,omitempty"`  // minOccurs="0"

	// to specify the name of a person in case of an emergeny
	ContactName *TravellerInformationTypeI `xml:"contactName,omitempty"`  // minOccurs="0"
}

type CoverageDetails struct {

	// To specify the details of the insurance policy.
	PolicyDetails *InsurancePolicyType `xml:"policyDetails"`

	// This group is used to describe the coverage conditions details.
	CoverageInfo []*CoverageInfo `xml:"coverageInfo,omitempty"`  // minOccurs="0" maxOccurs="3"

	// To specifie the covered persons: here it conveys the NB/NM and ON options
	CoveredPassenger []*TravellerInformationTypeI_15923S `xml:"coveredPassenger,omitempty"`  // minOccurs="0" maxOccurs="3"

	// starting date and end date
	CoverageDates *StructuredPeriodInformationType `xml:"coverageDates,omitempty"`  // minOccurs="0"

	// Details of the subscription: date and time.
	SubscriptionDetails *StructuredDateTimeInformationType_20644S `xml:"subscriptionDetails,omitempty"`  // minOccurs="0"

	// To convey the details of the insurance seller.
	AgentReferenceDetails *UserIdentificationType_9456S `xml:"agentReferenceDetails,omitempty"`  // minOccurs="0"
}

type InsuranceFopSection struct {

	// To convey the form of payment
	FormOfPaymentSection *FormOfPaymentTypeI_16862S `xml:"formOfPaymentSection"`

	// To provide form of payment extended data
	FopExtendedData []*StatusTypeI_13270S `xml:"fopExtendedData,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type PassengerDetails struct {

	// to specify to which passenger the insurance is associated: if omitted then it's for all the names in the PNR.
	PassengerAssociation *ReferenceInformationType `xml:"passengerAssociation"`

	// product knowledge indicator
	PerPaxProdKnowledge []*ActionDetailsTypeI `xml:"perPaxProdKnowledge,omitempty"`  // minOccurs="0" maxOccurs="9"

	// To specify the birthdate of the insuree.
	DateOfBirthInfo *StructuredDateTimeInformationType `xml:"dateOfBirthInfo,omitempty"`  // minOccurs="0"

	// to specify the name /age of the insuree
	PassengerFeatures []*TravellerInformationType `xml:"passengerFeatures,omitempty"`  // minOccurs="0" maxOccurs="2"

	// to specify a remark for the insuree
	InsureeRemark *MiscellaneousRemarksType `xml:"insureeRemark,omitempty"`  // minOccurs="0"

	// To specify the details concerning the documentation and the age of the insuree.
	TravelerDocInfo *PassengerDocumentDetailsType `xml:"travelerDocInfo,omitempty"`  // minOccurs="0"

	// fare discount code used per Pax
	PolicyDetails *InsurancePolicyType `xml:"policyDetails,omitempty"`  // minOccurs="0"

	// Details per insuree of the travel cost
	TravelerValueDetails *TravelerValueDetails `xml:"travelerValueDetails,omitempty"`  // minOccurs="0"

	// to convey for each tariff code and passenger the premium for this tariff.
	PremiumPerTariffPerPax []*PremiumPerTariffPerPax `xml:"premiumPerTariffPerPax,omitempty"`  // minOccurs="0" maxOccurs="4"

	// To convey the premium perpax
	PremiumPerpaxInfo *TariffInformationTypeI_22057S `xml:"premiumPerpaxInfo,omitempty"`  // minOccurs="0"

	// The Individual passenger reservation information
	VoucherNumber *ReservationControlInformationTypeU_31804S `xml:"voucherNumber,omitempty"`  // minOccurs="0"
}

type PlanTypeDetails struct {

	// Provides information about the type of plan being quoted/booked
	PlanType *InsuranceProviderAndProductsType `xml:"planType"`

	// to specify the value of the trip.
	TravelValue *MonetaryInformationTypeI `xml:"travelValue,omitempty"`  // minOccurs="0"
}

type ProductSection struct {

	// To convey the products or the tariffcodes together with description and amounts.
	ProductCode *InsuranceProductDetailsType_20774S `xml:"productCode"`

	// To convey the information the provider estimates important on a given product.
	InformationLines *FreeTextInformationType_6235S `xml:"informationLines,omitempty"`  // minOccurs="0"
}

type SubscriberAddressSection struct {

	// This segment is used to convey the contact name
	NameDetails *NameTypeU `xml:"nameDetails"`

	// to specify the address of the subscriber
	AddressInfo *AddressTypeU `xml:"addressInfo,omitempty"`  // minOccurs="0"

	// Used to specify a phone number
	PhoneNumber *PhoneAndEmailAddressType_32298S `xml:"phoneNumber,omitempty"`  // minOccurs="0"
}

type InsuranceCoverageType struct {

	// Indicate type of amount (eg. Medical Coverage, Trip Value, etc)
	CoverageIndicator []string `xml:"coverageIndicator,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type InsuranceCoverageType_25483S struct {

	// to indicate which coverage we are talking about.
	CoverageIndicator string `xml:"coverageIndicator"`
}

type InsuranceNameType struct {

	// insurance traveller details
	InsuranceTravelerDetails *SpecificTravellerDetailsType `xml:"insuranceTravelerDetails,omitempty"`  // minOccurs="0"

	// travelerperpax details
	TravelerPerpaxDetails []*TravelerPerpaxDetailsType `xml:"travelerPerpaxDetails,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type InsurancePolicyType struct {

	// to specify a discount for the insuree like if it's a family etc..
	FareDiscount string `xml:"fareDiscount,omitempty"`  // minOccurs="0"
}

type InsuranceProductDetailsType struct {

	// This data element is used to convey the company code of a non-air company
	CompanyCode string `xml:"companyCode,omitempty"`  // minOccurs="0"

	// To identify the countrycode from the provider.
	CountryCode string `xml:"countryCode,omitempty"`  // minOccurs="0"

	// Authorization number provided by ht insurance company
	ExtraReference []string `xml:"extraReference,omitempty"`  // minOccurs="0" maxOccurs="4"
}

type InsuranceProductDetailsType_20774S struct {

	// This data element is used to convey the company code of a non-air company
	CompanyCode string `xml:"companyCode,omitempty"`  // minOccurs="0"

	// To identify the countrycode from the provider.
	CountryCode string `xml:"countryCode,omitempty"`  // minOccurs="0"

	// This composite contains the code of the insurance elements.
	ProductDetails *ProviderInformationType `xml:"productDetails,omitempty"`  // minOccurs="0"

	// contains the extensions for the main insurance product
	ExtensionIdentification []*ProviderInformationType `xml:"extensionIdentification,omitempty"`  // minOccurs="0" maxOccurs="7"

	// tariff code info. tariff code and tariff familly code.
	TariffCodeDetails []*TariffcodeType `xml:"tariffCodeDetails,omitempty"`  // minOccurs="0" maxOccurs="48"
}

type InsuranceProductDetailsType_20775S struct {

	// tariff code info. tariff code and tariff familly code.
	TariffCodeDetails []*TariffcodeType `xml:"tariffCodeDetails,omitempty"`  // minOccurs="0" maxOccurs="48"
}

type InsuranceProviderAndProductsType struct {

	// Type of trip (package. leisure etc...)
	TripType string `xml:"tripType,omitempty"`  // minOccurs="0"

	// Code of the operator who provides the TOUR.
	TourOperator string `xml:"tourOperator,omitempty"`  // minOccurs="0"

	// To specify the countries involved in the Travel assistance element.
	CountryInfo *CountrydescriptionType `xml:"countryInfo,omitempty"`  // minOccurs="0"
}

type InteractiveFreeTextTypeI struct {

	// Provides information on the text conveyed in the IFT: language, type...
	FreeTextQualification *FreeTextQualificationTypeI_185754C `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// The information itself
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type InteractiveFreeTextTypeI_136698S struct {

	// Describes free text type
	FreetextDetail *FreeTextQualificationTypeI `xml:"freetextDetail,omitempty"`  // minOccurs="0"

	// One occurrence is supposed to represent a logical entity of free text (e.g. one line of text).
	Text []string `xml:"text,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type InteractiveFreeTextTypeI_99363S struct {

	FreeTextQualification *FreeTextQualificationTypeI_148295C `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText string `xml:"freeText,omitempty"`  // minOccurs="0"
}

type ItemDescriptionType struct {

	// Qualify the item being described
	ItemCharacteristic string `xml:"itemCharacteristic,omitempty"`  // minOccurs="0"
}

type ItemNumberIdentificationTypeU struct {

	// leg number
	Number string `xml:"number"`
}

type ItemNumberIdentificationTypeU_46320C struct {

	// The place of the product in the Tour booking.
	ItemID int32 `xml:"itemID"`

	// It qualifies the item ID type.
	ItemIDQualifier string `xml:"itemIDQualifier"`
}

type ItemNumberTypeU struct {

	// Provides information about the product place in the tour booking. It locally identifies the product in the booking.
	ItemIdentification *ItemNumberIdentificationTypeU_46320C `xml:"itemIdentification"`
}

type ItemNumberTypeU_33258S struct {

	// leg number - idicate with leg is the first one, the second one, etc.
	ItemNumberDetails *ItemNumberIdentificationTypeU `xml:"itemNumberDetails"`
}

type ItemReferencesAndVersionsType struct {

	// qualifies the type of the reference used. Code set to define
	ReferenceType string `xml:"referenceType"`

	// The value of the reference
	UniqueReference string `xml:"uniqueReference"`
}

type ItemReferencesAndVersionsType_6550S struct {

	// ID details
	IDSection *UniqueIdDescriptionType `xml:"iDSection"`
}

type ItemReferencesAndVersionsType_9271S struct {

	// Defines the type of reference used: GPN : group code
	ReferenceType string `xml:"referenceType"`

	// The value of the reference.
	UniqueReference string `xml:"uniqueReference"`
}

type ItemReferencesAndVersionsType_94556S struct {

	// qualifies the type of the reference used. Here it will be: PRI Payment Record Id APP Application Correlator Id EXT Third party Record Id
	ReferenceType string `xml:"referenceType"`

	// The value of the payment record/correlator Id
	UniqueReference string `xml:"uniqueReference"`
}

type LocationIdentificationBatchTypeU struct {

	// Set to: IATA to indicate IATA location code 1A to indicate a 1A location CPY to indicate a Car provider location
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Location extended name for  - Amadeus location type  - Provider location type (followed by an *) - Free text for collection option.  - Free text for delivery option.
	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type LocationIdentificationBatchTypeU_198230C struct {

	// Location Code "1A" to indicate Amadeus location type "CPY" to indicate a Provider location type
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Location extended name for  Amadeus location type and  Provider location type (followed by an *)
	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type LocationIdentificationBatchTypeU_46344C struct {

	// location code
	Code string `xml:"code"`

	// location qualifier (city, country ...)
	Qualifier string `xml:"qualifier"`

	// location name
	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type LocationIdentificationBatchTypeU_56454C struct {

	// Railway station location code
	Code string `xml:"code"`

	// Code type
	Qualifier string `xml:"qualifier"`

	// Location name
	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type LocationIdentificationBatchTypeU_60738C struct {

	// Railway station location code
	Code string `xml:"code"`

	// Code type
	Qualifier string `xml:"qualifier"`
}

type LocationIdentificationTypeS struct {

	// Station. See IATA Airline Coding directory. IATA 3 letter city/aircode code
	CityCode string `xml:"cityCode"`
}

type LocationIdentificationTypeU struct {

	// 162: country
	Qualifier string `xml:"qualifier"`

	// Country name
	Name string `xml:"name"`
}

type LocationIdentificationTypeU_198211C struct {

	// Identification of the site
	Code string `xml:"code"`

	// Site name
	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type LocationTypeI struct {

	// Format limitations: a3
	TrueLocationId string `xml:"trueLocationId,omitempty"`  // minOccurs="0"
}

type LocationTypeI_2784C struct {

	// AIR segment : boarding point ATX segment : boarding point CAR segment : pick-up point city CCR segment : pick-up point city HHL segment : city code HTL segment : check-in city MIS segment : city code SUR segment : city/airport code Trn Amtrak sgt: board point city code Trn SNCF sgt: board point city code (RESARAIL code) TTO segment: departure location TUR segment: tour start city CRU segment: sailing departure port
	CityCode string `xml:"cityCode"`

	// TRN SNCF segment : board point city name.
	CityName string `xml:"cityName,omitempty"`  // minOccurs="0"
}

type LocationTypeU struct {

	// Port code.
	Code string `xml:"code"`

	// Port name.
	Name string `xml:"name"`

	// Port codes are non-standard and specific to the Amadeus ferry business.
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"
}

type LocationTypeU_46324C struct {

	// city code
	Code string `xml:"code"`

	// city name
	Name string `xml:"name,omitempty"`  // minOccurs="0"

	// country code
	Country string `xml:"country,omitempty"`  // minOccurs="0"

	// location qualifier for the repetition (departure location, arrival location ...)
	Qualifier string `xml:"qualifier"`
}

type LongFreeTextType struct {

	// To specify the type of freetext
	FreetextDetail *FreeTextQualificationType `xml:"freetextDetail,omitempty"`  // minOccurs="0"

	// Long free text information.
	LongFreetext string `xml:"longFreetext,omitempty"`  // minOccurs="0"
}

type MeanOfPaymentDataType struct {

	// This segment will convey the type of the FOP. Exple : CC credit card CA cash CH cheque SWI swipe card WA web account WB web bank(fund tranfer)
	FopInformation *FormOfPaymentType `xml:"fopInformation"`

	// will allow the usage of FOP segment as trigger for GASY and GINV groups
	Dummy *DummySegmentTypeI `xml:"dummy"`

	// will convey all credit card data needed for the payment
	CreditCardData *CreditCardDataGroupType `xml:"creditCardData,omitempty"`  // minOccurs="0"
}

type MeasurementsBatchTypeU struct {

	// Measurement qualifier (maximum unit qualifier).
	MeasurementQualifier string `xml:"measurementQualifier"`

	// Unit Qualifer for the range value.
	ValueRange *ValueRangeTypeU `xml:"valueRange"`
}

type MessageActionDetailsTypeI struct {

	// MESSAGE FUNCTION OR BUSINESS DETAILS
	Business *MessageFunctionBusinessDetailsTypeI `xml:"business"`
}

type MessageFunctionBusinessDetailsTypeI struct {

	// - Message function: REG for regular feed - Itinerary type: see codeset list: air/car/hotel/taxi/train/tour/surface...
	Function string `xml:"function"`
}

type MessageReferenceType struct {

	// This number is used to identify and track ALL messages related to a given cardholder transaction (author, retry, reversal ...). It is usually composed of: - the date when the message was formatted followed by - the message number   Field 37  Official definition of Retrieval Reference Number from ISO8583: Field 37 contains a number used with other key data elements to identify and track all messages related to a given cardholder transaction (referred to as a transaction set). It is usually assigned by the acquirer, but it may be assigned by a merchant or by an individual electronic terminal. V.I.P. will also generate the retrieval reference number for transactions it initiates. This field contains two parts. The first four digits are usually a yddd date (Julian date format). The date is defined to be the same day as the date in Field 7_Transmission Date and Time, of the original request. The last eight digits are a numeric transaction identification number. The value in field 37 can be based on the content of fields 7 and 11 in the original request or advice as shown in the recommendation below: . Positions 1_4: the yddd equivalent of the field 7 date . Positions 5_6: the hours from the time in field 7 . Positions 7_12: the value from field 11
	RetrievalReferenceNumber string `xml:"retrievalReferenceNumber,omitempty"`  // minOccurs="0"

	// Authorization characteristics indicator   Field 62.1 Possible values: A C E F K M S U V W R I P N T
	AuthorCharacteristicIndicator string `xml:"authorCharacteristicIndicator,omitempty"`  // minOccurs="0"

	// Authorization response code   Field 39
	AuthorResponseCode string `xml:"authorResponseCode,omitempty"`  // minOccurs="0"

	// Card Level Result (Product Identification value)  Field 62.23
	CardLevelResult string `xml:"cardLevelResult,omitempty"`  // minOccurs="0"

	// Additional POS Information - Terminal Type  Field 60.1 - Position 1  CAT (Cardholder-Activated Terminal indicator) or UAT (Unattended Acceptance Terminal)
	TerminalType string `xml:"terminalType,omitempty"`  // minOccurs="0"
}

type MileageTimeDetailsTypeI struct {

	// Format limitations: n..18
	FlightLegMileage *int32 `xml:"flightLegMileage,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	UnitQualifier string `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type MiscellaneousChargeOrderType struct {

	// Type of service
	Type string `xml:"type"`
}

type MiscellaneousRemarkType struct {

	// This data element is used to convey the type of the remark. (see data mapping to view the codes)
	Type string `xml:"type"`

	// Free text and message sequence numbers of the remarks.
	Freetext string `xml:"freetext,omitempty"`  // minOccurs="0"

	// This data element is used to convey the business function
	BusinessFunction string `xml:"businessFunction,omitempty"`  // minOccurs="0"

	// language used for the free text.
	Language string `xml:"language,omitempty"`  // minOccurs="0"

	// Indicates if it has been manually entered by an agent or system generated.
	Source string `xml:"source,omitempty"`  // minOccurs="0"

	// Coded identification of the character encoding used in the interchange
	Encoding string `xml:"encoding,omitempty"`  // minOccurs="0"
}

type MiscellaneousRemarkType_151C struct {

	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark
	Type string `xml:"type"`

	// This is the 3rd character (x) of the remark title RIx or RMx, or 2 letter code for RMxx, conditional for RM, not applicable for RC and RQ
	Category string `xml:"category,omitempty"`  // minOccurs="0"

	// Free text and message sequence numbers of the remarks.
	Freetext string `xml:"freetext,omitempty"`  // minOccurs="0"

	// Provider type (element RIA):  1 for Air provider   2 for Car provider (CCR)  3 for Hotel Provider (HHL)  M for Miscellaneous
	ProviderType string `xml:"providerType,omitempty"`  // minOccurs="0"
}

type MiscellaneousRemarkType_18076C struct {

	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark . ACC for Acceptance . BGG for Baggage . BPP for Boarding Pass Printing . GT for Gate . GNL for General
	Type string `xml:"type"`

	// Free text and message sequence numbers of the remarks.
	Freetext string `xml:"freetext,omitempty"`  // minOccurs="0"
}

type MiscellaneousRemarkType_198195C struct {

	// This data element is used to convey the type of the remark. (see data mapping to view the codes)
	Type string `xml:"type"`

	// Free text and message sequence numbers of the remarks.
	Freetext string `xml:"freetext,omitempty"`  // minOccurs="0"

	// This data element is used to convey the business function
	BusinessFunction string `xml:"businessFunction,omitempty"`  // minOccurs="0"

	// language used for the free text.
	Language string `xml:"language,omitempty"`  // minOccurs="0"

	// Indicates if it has been manually entered by an agent or system generated.
	Source string `xml:"source,omitempty"`  // minOccurs="0"

	// Coded identification of the character encoding used in the interchange
	Encoding string `xml:"encoding,omitempty"`  // minOccurs="0"
}

type MiscellaneousRemarkType_861C struct {

	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark . ACC for Acceptance . BGG for Baggage . BPP for Boarding Pass Printing . GT for Gate . GNL for General
	Type string `xml:"type"`

	// Free text and message sequence numbers of the remarks.
	Freetext string `xml:"freetext,omitempty"`  // minOccurs="0"
}

type MiscellaneousRemarksType struct {

	// miscellaneous remarks
	RemarkDetails *MiscellaneousRemarkType_861C `xml:"remarkDetails,omitempty"`  // minOccurs="0"
}

type MiscellaneousRemarksType_12240S struct {

	// miscellaneous remarks
	RemarkDetails *MiscellaneousRemarkType_18076C `xml:"remarkDetails,omitempty"`  // minOccurs="0"
}

type MiscellaneousRemarksType_136700S struct {

	// miscellaneous remarks
	RemarkDetails *MiscellaneousRemarkType_198195C `xml:"remarkDetails,omitempty"`  // minOccurs="0"
}

type MiscellaneousRemarksType_211S struct {

	// Miscellaneous remqrks
	Remarks *MiscellaneousRemarkType_151C `xml:"remarks,omitempty"`  // minOccurs="0"

	// For confidential remark RC
	IndividualSecurity []*IndividualSecurityType `xml:"individualSecurity,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type MiscellaneousRemarksType_664S struct {

	// miscellaneous remarks
	RemarkDetails *MiscellaneousRemarkType `xml:"remarkDetails,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsType struct {

	// Here is the list and the purpose of each amount today stored in the FP: I Transaction total amount  Total amount authorized in authorization transaction  IPC Transaction total amount in PNR currency Total amount authorized is also stored in PNR currency. Indeed, reversal must be done with the rate of exchange valid at time of authorization and therefore this avoids storing the rate of exchange and performing amount conversion at reversal time.  IT Initial TST total amount  Amount of TST multiplied by the number of passengers associated to the TST  ITC Initial TST total amount in PNR currency   IT amount in PNR currency for same reason as IPC amount  R Total amount / Remaining amount  Current authorized amount. Originally it is the total amount authorized and then this amount  may decrease in case of total/partial reversal.  T Initial Tst Individual amount  Amount of TST  TPC Initial Tst Individual amount in PNR currency  Amount of TST in PNR currency for same reason as IPC amount  AUT Authorized Amount Maybe different from the one given in input(for exple, if on input we have 2pax and the amount per pax. In case of bulk, we will authorize the sum of both amounts) It can also be used for:  Total Fare Amount 712 or additional collection amount A
	TypeQualifier string `xml:"typeQualifier"`

	// Value of the amount.  This is conveyed as a 'string' and therefore several strings can stand for the  same amount (eg. 14 , 1400, 14.00... could potentially stand for 14.00 EUR). This  means that sender/receiver of this message will need to come to an agreement  concerning the way the amount is transferred in this segment.
	Amount float64 `xml:"amount"`

	// IATA alphabetic currency code.  Eg: USD,GBP,EUR...
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI struct {

	// Yield type
	TypeQualifier string `xml:"typeQualifier"`

	// Amount of the Yield
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_17849C struct {

	// Indicates amount is Fare amount
	TypeQualifier string `xml:"typeQualifier"`

	// Used to specify an amount of money
	Amount string `xml:"amount,omitempty"`  // minOccurs="0"

	// currency in which the amount is expressed
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_4220C struct {

	// Monetary amount qualifier : - NP : Net Premium - PR : Premium - CV : Coverage - TV : Travel Value - SAV : Saving Amount
	Qualifier string `xml:"qualifier"`

	// Amount
	Amount float64 `xml:"amount"`

	// Eg: USD,FRF,EUR...
	CurrencyCode string `xml:"currencyCode"`
}

type MonetaryInformationDetailsTypeI_8308C struct {

	// . F for Fare basis . E for Equivalent  . T for Total
	Qualifier string `xml:"qualifier"`

	// Amount
	Amount string `xml:"amount"`

	// Eg: USD,FRF,EUR...
	CurrencyCode string `xml:"currencyCode"`
}

type MonetaryInformationType struct {

	// Yield info
	MonetaryDetails *MonetaryInformationDetailsTypeI `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"otherMonetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="4"
}

type MonetaryInformationTypeI struct {

	// Total Trip value in a given currency
	MonetaryDetails *MonetaryInformationDetailsTypeI_17849C `xml:"monetaryDetails"`

	// Base Trip value in a given currency
	OtherMonetaryDetails *MonetaryInformationDetailsTypeI_17849C `xml:"otherMonetaryDetails,omitempty"`  // minOccurs="0"
}

type MonetaryInformationTypeI_1689S struct {

	// To specify monetary information
	Information *MonetaryInformationDetailsTypeI_4220C `xml:"information,omitempty"`  // minOccurs="0"
}

type MonetaryInformationType_94557S struct {

	// Contains the currencies and the various amounts
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="7"
}

type NameAndAddressBatchTypeU struct {

	// W for party to revieve written confirmation
	PartyQualifier string `xml:"partyQualifier"`

	// This composite is used to convey the address
	AddressDetails *NameAndAddressDetailsTypeU `xml:"addressDetails,omitempty"`  // minOccurs="0"

	// This composite is used to convey the party name
	PartyNameDetails *PartyNameBatchTypeU `xml:"partyNameDetails,omitempty"`  // minOccurs="0"
}

type NameAndAddressDetailsTypeU struct {

	// Address line 1
	Line1 string `xml:"line1"`

	// address line 2
	Line2 string `xml:"line2,omitempty"`  // minOccurs="0"
}

type NameInformationTypeU struct {

	// name qualifier
	Qualifier string `xml:"qualifier"`

	// name
	Name string `xml:"name"`
}

type NameInformationTypeU_9747C struct {

	// to convey to who the address applies
	Qualifier string `xml:"qualifier"`

	// Company name
	Name string `xml:"name,omitempty"`  // minOccurs="0"

	// Insuree name
	Identifier string `xml:"identifier,omitempty"`  // minOccurs="0"
}

type NameTypeU struct {

	// Used to specify the name field in the address field.
	NameInformation *NameInformationTypeU_9747C `xml:"nameInformation"`
}

type NameTypeU_136701S struct {

	// Name information
	NameInformation *NameInformationTypeU `xml:"nameInformation"`
}

type NumberOfUnitDetailsTypeI struct {

	// Group counter corresponding to passengers, and so value from 0 to 99.
	NumberOfUnit *int32 `xml:"numberOfUnit,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	UnitQualifier string `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type NumberOfUnitDetailsTypeI_18670C struct {

	// This data element is used to convey the occupancy level of the room
	NumberOfUnit int32 `xml:"numberOfUnit"`

	// This data element is used to indicate the occupancy is the number of Adults in the room.
	UnitQualifier string `xml:"unitQualifier"`
}

type NumberOfUnitDetailsTypeI_2755C struct {

	// PNR Header / Queue header / number of remaining items in Queue
	Number *int32 `xml:"number,omitempty"`  // minOccurs="0"

	// PNR for PNR
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"
}

type NumberOfUnitDetailsTypeI_35712C struct {

	// Number of units.
	NumberOfUnit int32 `xml:"numberOfUnit"`
}

type NumberOfUnitDetailsTypeI_46330C struct {

	// occupation of the room
	NumberOfUnit int32 `xml:"numberOfUnit"`

	// unit qualifier
	UnitQualifier string `xml:"unitQualifier"`
}

type NumberOfUnitsType struct {

	// Number of Unit Details
	QuantityDetails *NumberOfUnitDetailsTypeI_35712C `xml:"quantityDetails"`
}

type NumberOfUnitsTypeI struct {

	// Number of Units detail
	NumberDetail *NumberOfUnitDetailsTypeI_2755C `xml:"numberDetail"`
}

type NumberOfUnitsType_76106S struct {

	// Number of Unit Details
	QuantityDetails []*NumberOfUnitDetailsTypeI `xml:"quantityDetails"`  // maxOccurs="3"
}

type ODKeyPerformanceDataType struct {

	// schedule change indicator -'C' or void
	ScheduleChange string `xml:"scheduleChange,omitempty"`  // minOccurs="0"

	// oversale data
	Oversale *OversaleDataType `xml:"oversale,omitempty"`  // minOccurs="0"
}

type ONDType struct {

	// Yield informations:  Adjusted Yield Segment Bid Price Effective Yield Revenue Loss OND Yield
	YieldInformations *MonetaryInformationType `xml:"yieldInformations"`

	// Class code as defined in yield retrieved / Class combinaison of the yield retrieved
	ClassCombinaison *ProductInformationTypeI_76271S `xml:"classCombinaison,omitempty"`  // minOccurs="0"

	// Origin and Destination of the Yield
	Ondyield *OriginAndDestinationDetailsTypeI_76268S `xml:"ondyield"`

	// Origin And Destination of the Trip
	TripOnD *OriginAndDestinationDetailsTypeI_76268S `xml:"tripOnD,omitempty"`  // minOccurs="0"
}

type OptionElementInformationType struct {

	// Option element office id
	MainOffice string `xml:"mainOffice"`

	// Date
	Date string `xml:"date,omitempty"`  // minOccurs="0"

	// Queue number
	Queue *int32 `xml:"queue,omitempty"`  // minOccurs="0"

	// Category number
	Category *int32 `xml:"category,omitempty"`  // minOccurs="0"

	// Format limitations: an..61
	Freetext string `xml:"freetext,omitempty"`  // minOccurs="0"

	// queuing or cancellation time
	Time string `xml:"time,omitempty"`  // minOccurs="0"
}

type OptionElementType struct {

	OptionElementInfo *OptionElementInformationType `xml:"optionElementInfo,omitempty"`  // minOccurs="0"

	// Individual Security for OPQ/OPX elements
	IndividualSecurity []*IndividualSecurityType `xml:"individualSecurity,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type OriginAndDestinationDetailsTypeI struct {

	// City pair to indentify uniquely a leg in a multi-leg booking
	Origin string `xml:"origin"`

	// City pair to indentify uniquely a leg in a multi-leg booking
	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type OriginAndDestinationDetailsTypeI_3061S struct {

	// Airport/city code of  Origin In a Client request message, a non-blank ODI is used in an air sell request to advise that the following segments (TVL etc...) are connected. There is a maximum of 6 TVLs following a non-blank ODI.
	Origin string `xml:"origin,omitempty"`  // minOccurs="0"

	// Airport/city code of  Destination
	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type OriginAndDestinationDetailsTypeI_76268S struct {

	// Departure's city code:3 character ATA/IATA airport/city code
	Origin string `xml:"origin,omitempty"`  // minOccurs="0"

	// Arrival's city code:3 character ATA/IATA airport/city code
	Destination string `xml:"destination,omitempty"`  // minOccurs="0"
}

type OriginatorDetailsTypeI struct {

	// Country code
	CodedCountry string `xml:"codedCountry"`

	// Currency code
	CodedCurrency string `xml:"codedCurrency,omitempty"`  // minOccurs="0"

	// Language code
	CodedLanguage string `xml:"codedLanguage,omitempty"`  // minOccurs="0"
}

type OriginatorIdentificationDetailsTypeI struct {

	// IATA code
	OriginatorId *int32 `xml:"originatorId,omitempty"`  // minOccurs="0"

	// Office ID of the PNR owner.
	InHouseIdentification1 string `xml:"inHouseIdentification1"`

	// Amid of the owner of the SBR.
	InHouseIdentification2 *int32 `xml:"inHouseIdentification2,omitempty"`  // minOccurs="0"
}

type OriginatorIdentificationDetailsTypeI_198179C struct {

	// Agency Iata code
	OriginatorId string `xml:"originatorId"`
}

type OriginatorIdentificationDetailsTypeI_37406C struct {

	// This data element is used to convey the bouking source.
	OriginatorId int32 `xml:"originatorId"`
}

type OriginatorIdentificationDetailsTypeI_46358C struct {

	// Origin OficeID
	InHouseIdentification1 string `xml:"inHouseIdentification1,omitempty"`  // minOccurs="0"

	// Target OfficeID
	InHouseIdentification2 string `xml:"inHouseIdentification2,omitempty"`  // minOccurs="0"
}

type OtherHotelInformationType struct {

	// Currency Code at Property 1. For AY Direct Access segment: Currency code
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"
}

type OtherInformationType struct {

	// Queue cycle complete indicator that may appear in Queue working response message.  QCC for Queue cycle complete
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"

	// Indicates the type of Queue in a Queue working response message.  PNR for PNR  MSG for Message
	QueueType string `xml:"queueType,omitempty"`  // minOccurs="0"
}

type OtherSegmentDataTypeI struct {

	// Cabin Code
	Cabin string `xml:"cabin"`

	// Sub class number
	Subclass *int32 `xml:"subclass,omitempty"`  // minOccurs="0"

	// Flight type : - D for Domestic - I for International - L for Longhaul - S for Shorthaul
	FlightType string `xml:"flightType,omitempty"`  // minOccurs="0"

	// Overbooking indicator
	Overbooking string `xml:"overbooking,omitempty"`  // minOccurs="0"
}

type OversaleDataType struct {

	// Bid price oversale number
	OversaleNumber *float64 `xml:"oversaleNumber,omitempty"`  // minOccurs="0"

	// Oversale indicator F  for Bid-Price Feed Oversale O for Bid-Price Oversale P  for Pushed Minimum Oversale
	OversaleIndicator []string `xml:"oversaleIndicator,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type PNRSupplementaryDataType struct {

	// will convey the values of the FOP data and switch maps
	DataAndSwitchMap *AttributeType_94576S `xml:"dataAndSwitchMap"`
}

type POSGroupType struct {

	// - Office ID owner of the SBR. - IATA Code - Agent type
	SbrUserIdentificationOwn *UserIdentificationType `xml:"sbrUserIdentificationOwn"`

	// - Corporate Code - City Code
	SbrSystemDetails *SystemDetailsInfoType_33158S `xml:"sbrSystemDetails,omitempty"`  // minOccurs="0"

	// Preferences - Country - Language - Currency
	SbrPreferences *UserPreferencesType `xml:"sbrPreferences,omitempty"`  // minOccurs="0"
}

type PackageDescriptionType struct {

	// Inclusive package type: I
	PackageType string `xml:"packageType"`

	// List of inclusive package
	PackageDetails *PackageIdentificationType `xml:"packageDetails,omitempty"`  // minOccurs="0"
}

type PackageIdentificationType struct {

	// Description of a package
	PackageDesc string `xml:"packageDesc"`
}

type PartyNameBatchTypeU struct {

	// name
	Name1 string `xml:"name1"`
}

type PassengerDocumentDetailsType struct {

	// Used to convey the age of the insuree
	BirthDate string `xml:"birthDate,omitempty"`  // minOccurs="0"

	// Details on the document (visa, passport...)
	DocumentDetails *DocumentDetailsType `xml:"documentDetails,omitempty"`  // minOccurs="0"
}

type PassengerFlightDetailsTypeI struct {
}

type PaymentDataGroupType struct {

	// Contains merchant information (Entity selling a product/service for wich payment is requested: airline, insurance provider...).
	MerchantInformation *CompanyInformationType_94554S `xml:"merchantInformation"`

	// will convey all the monetary informations related to the payment : amount, currency, sub-amounts
	MonetaryInformation []*MonetaryInformationType_94557S `xml:"monetaryInformation,omitempty"`  // minOccurs="0" maxOccurs="999"

	// Conveys Payment Record ID (used by Payment Manager) to identify payment in a unique manner.  May convey also a "correlator Id" used by the calling application to reconciliate its payment data.  And also the "transaction Id" generated by the third party system (bank/PSP/PAyPAL...)
	PaymentId []*ItemReferencesAndVersionsType_94556S `xml:"paymentId,omitempty"`  // minOccurs="0" maxOccurs="3"

	// It will describe the content of the extended payment : when it will start, the frequency and how many times it should occur
	ExtendedPaymentInfo *FrequencyTypeU `xml:"extendedPaymentInfo,omitempty"`  // minOccurs="0"

	// The segment conveys the date/time of the transaction
	TransactionDateTime *StructuredDateTimeInformationType_94559S `xml:"transactionDateTime,omitempty"`  // minOccurs="0"

	// Will show the duration of validity of the payment request, mesured from receipt by the issuer. The customer has to agree to the payment within this period. Expressed in seconds.
	ExpirationPeriod *QuantityType_94558S `xml:"expirationPeriod,omitempty"`  // minOccurs="0"

	// Distribution Channel information
	DistributionChannelInformation *TerminalIdentificationDescriptionType `xml:"distributionChannelInformation,omitempty"`  // minOccurs="0"

	// will convey in free text the description of the purchase
	PurchaseDescription *FreeTextInformationType_94561S `xml:"purchaseDescription,omitempty"`  // minOccurs="0"

	// will convey all information needed to perform the checks requested by the banks/PSPs regarding the prevention of fraud.
	FraudScreeningData *FraudScreeningGroupType `xml:"fraudScreeningData,omitempty"`  // minOccurs="0"

	// Will be used to convey information dedicated to the Payment.
	PaymentDataMap []*AttributeType_94553S `xml:"paymentDataMap,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type PaymentDetailsTypeI struct {

	// To convey the guarantee /deposit form
	FormOfPaymentCode string `xml:"formOfPaymentCode"`

	// This data element is used to idicates if it is a guarantee or a deposit
	PaymentType string `xml:"paymentType"`

	// This data element is used to identify the type of service to be paid, in our case it will always be 3 for hotel
	ServiceToPay string `xml:"serviceToPay"`

	// This data element is used to convey the guarantee or the deposit reference.
	ReferenceNumber string `xml:"referenceNumber,omitempty"`  // minOccurs="0"
}

type PaymentDetailsTypeU struct {

	// Identify the mode of payment:  - CASH  - CC for credit card
	MethodCode string `xml:"methodCode,omitempty"`  // minOccurs="0"

	// Purpose of the payment:  - DEPO for deposit  - FINA for final payment
	PurposeCode string `xml:"purposeCode"`

	// Amount paid
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency used for the payment
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// date of the payment
	Date string `xml:"date,omitempty"`  // minOccurs="0"
}

type PaymentGroupType struct {

	// Used to describe the element on which the action is performed : FP/FC/PAY and in which context integrated/non integrated
	GroupUsage *CodedAttributeType_127282S `xml:"groupUsage"`

	// will convey all data necessary for the paiment and not dependant from the Mean Of Payment
	PaymentData *PaymentDataGroupType `xml:"paymentData,omitempty"`  // minOccurs="0"

	// it will convey the Descriptive Billing Information: ONO, GWT, best Fare indicator....
	PaymentSupplementaryData []*CodedAttributeType_94497S `xml:"paymentSupplementaryData,omitempty"`  // minOccurs="0" maxOccurs="99"

	// will convey all the specificities of the Mean of Payment
	MopInformation *MeanOfPaymentDataType `xml:"mopInformation,omitempty"`  // minOccurs="0"

	// will allow the usage of FOP segment as trigger for MOPD and MOPS groups
	Dummy *DummySegmentTypeI `xml:"dummy"`

	// will convey the result of the payment and related to the detailed Mean Of Payment
	MopDetailedData *DetailedPaymentDataType `xml:"mopDetailedData,omitempty"`  // minOccurs="0"
}

type PaymentInformationTypeI struct {

	// This composite is used to convey the payment information
	PaymentDetails *PaymentDetailsTypeI `xml:"paymentDetails"`
}

type PaymentInformationTypeU struct {

	// Tour deposit details
	PaymentDetails *PaymentDetailsTypeU `xml:"paymentDetails"`

	// Credit card name, number and exp. date
	CreditCardInformation *CreditCardInformationTypeU `xml:"creditCardInformation,omitempty"`  // minOccurs="0"
}

type PhoneAndEmailAddressType struct {

	// Phone or Email contact type
	PhoneOrEmailType string `xml:"phoneOrEmailType"`

	// Structured telephone number
	TelephoneNumberDetails *StructuredTelephoneNumberType_198214C `xml:"telephoneNumberDetails"`
}

type PhoneAndEmailAddressType_136723S struct {

	// - PHO phone number - FAX fax number - MAI
	PhoneOrEmailType string `xml:"phoneOrEmailType"`

	// Email address
	EmailAddress string `xml:"emailAddress"`
}

type PhoneAndEmailAddressType_32298S struct {

	// Phone or Email contact type
	PhoneOrEmailType string `xml:"phoneOrEmailType"`

	// Structured telephone number
	TelephoneNumber *StructuredTelephoneNumberType_48448C `xml:"telephoneNumber,omitempty"`  // minOccurs="0"

	// Email address
	EmailAddress string `xml:"emailAddress,omitempty"`  // minOccurs="0"
}

type PhoneAndEmailAddressType_94565S struct {

	// Phone or Email contact type
	PhoneOrEmailType string `xml:"phoneOrEmailType"`

	// Structured telephone number
	TelephoneNumberDetails *StructuredTelephoneNumberType `xml:"telephoneNumberDetails,omitempty"`  // minOccurs="0"

	// Email address
	EmailAddress string `xml:"emailAddress,omitempty"`  // minOccurs="0"
}

type PlaceLocationIdentificationTypeU struct {

	// location code qualifier
	LocationType string `xml:"locationType"`

	// location text
	LocationDescription *LocationIdentificationBatchTypeU `xml:"locationDescription"`

	// Associated airport/City code. Present if the pickup location is not an airport/city code.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_198193C `xml:"firstLocationDetails,omitempty"`  // minOccurs="0"
}

type PlaceLocationIdentificationTypeU_136722S struct {

	// Used to differenciate the pickup location (176) from the Dropoff location (DOL)
	LocationType string `xml:"locationType"`

	// Pickup or dropoff location details
	LocationDescription *LocationIdentificationBatchTypeU_198230C `xml:"locationDescription,omitempty"`  // minOccurs="0"
}

type PlaceLocationIdentificationTypeU_24573S struct {

	// Location type qualifier (ZZZ-Mutually defined for Ferry).
	LocationType string `xml:"locationType"`

	// Hotel location details.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_45087C `xml:"firstLocationDetails"`
}

type PlaceLocationIdentificationTypeU_25436S struct {

	// location type (place of arrival, place of departure or staying)
	LocationType string `xml:"locationType"`

	// city information
	LocationDescription *LocationIdentificationBatchTypeU_46344C `xml:"locationDescription"`

	// country description
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_46345C `xml:"firstLocationDetails,omitempty"`  // minOccurs="0"
}

type PlaceLocationIdentificationTypeU_32347S struct {

	// Type of location
	LocationType string `xml:"locationType"`

	// Railway station location details.
	LocationDescription *LocationIdentificationBatchTypeU_56454C `xml:"locationDescription"`

	// Railway station country details.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_56455C `xml:"firstLocationDetails,omitempty"`  // minOccurs="0"
}

type PlaceLocationIdentificationTypeU_35293S struct {

	// Type of location
	LocationType string `xml:"locationType"`

	// Railway station location details.
	LocationDescription *LocationIdentificationBatchTypeU_60738C `xml:"locationDescription"`

	// Railway station country details.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_56455C `xml:"firstLocationDetails,omitempty"`  // minOccurs="0"
}

type PlaceLocationIdentificationTypeU_8954S struct {

	// Details of the embarkation port.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU `xml:"firstLocationDetails"`

	// Details of the disembarkation port.
	SecondLocationDetails *RelatedLocationTwoIdentificationTypeU `xml:"secondLocationDetails"`
}

type PnrHistoryDataType struct {

	// Contains the last EOTed envelop number.
	CurrentRecord *int32 `xml:"currentRecord,omitempty"`  // minOccurs="0"
}

type PnrHistoryDataType_6022S struct {

	// Reference to previous envelop It may not exist when we are on element creation case.
	PreviousRecord *int32 `xml:"previousRecord,omitempty"`  // minOccurs="0"

	// Current envelop
	CurrentRecord *int32 `xml:"currentRecord,omitempty"`  // minOccurs="0"

	// History element name  ON, AS, RF... First char for type of action done, followed by a letter related to the element concerned.
	ElementType string `xml:"elementType,omitempty"`  // minOccurs="0"

	// Free flow text (history data element not detailed) Max length put from 254 to 255 for  the case of long history
	ElementData string `xml:"elementData"`
}

type PointOfSaleDataTypeI struct {

	// POSINV Classification: - C for Country - R for CRS
	Classification string `xml:"classification"`

	// Point of Sale CRS
	Crs string `xml:"crs,omitempty"`  // minOccurs="0"

	// Point of Sale Country Code
	PointOfSaleCountry string `xml:"pointOfSaleCountry,omitempty"`  // minOccurs="0"
}

type PricingOrTicketingSubsequentType struct {

	// Reason for issuance code.
	SpecialCondition string `xml:"specialCondition,omitempty"`  // minOccurs="0"

	// Reason for Issuance Sub code
	OtherSpecialCondition string `xml:"otherSpecialCondition,omitempty"`  // minOccurs="0"
}

type PriorityDetailsType struct {

	// 1 : airline 2 : alliance
	Qualifier string `xml:"qualifier"`

	// Priority code
	PriorityCode string `xml:"priorityCode,omitempty"`  // minOccurs="0"

	// Tier level
	TierLevel string `xml:"tierLevel,omitempty"`  // minOccurs="0"

	// Tier description
	TierDescription string `xml:"tierDescription,omitempty"`  // minOccurs="0"
}

type ProcessingInformationTypeI struct {

	// Identifies the element we are talking about
	ActionQualifier string `xml:"actionQualifier,omitempty"`  // minOccurs="0"

	// Used to qualifie the element with an indicator.
	ReferenceQualifier string `xml:"referenceQualifier,omitempty"`  // minOccurs="0"
}

type ProductAccountDetailsTypeI struct {

	// the award code returned by loyalty system in booking time and send to loyalty system in ticketing time.
	Category string `xml:"category,omitempty"`  // minOccurs="0"

	// Contains the old class of the segment before the upgrade.
	SequenceNumber string `xml:"sequenceNumber,omitempty"`  // minOccurs="0"

	// certificate number
	VersionNumber string `xml:"versionNumber,omitempty"`  // minOccurs="0"

	// Fake Tier level received by TTY in.
	RateClass string `xml:"rateClass,omitempty"`  // minOccurs="0"

	// stock control number
	ApprovalCode string `xml:"approvalCode,omitempty"`  // minOccurs="0"
}

type ProductDataInformationTypeU struct {

	// Tour product category (StandAlone, Package, Supplementary service ...)
	ProductCategory string `xml:"productCategory"`

	// Conveys the product code
	ProductCode string `xml:"productCode,omitempty"`  // minOccurs="0"

	// Set to 1 if the product is an addOn.
	AddOnIndicator *int32 `xml:"addOnIndicator,omitempty"`  // minOccurs="0"

	// The product description
	ProductDescription string `xml:"productDescription,omitempty"`  // minOccurs="0"
}

type ProductDateAndTimeTypeU struct {

	// Convey the begin date of a period. Format is ddmmyyyy.
	DepartureDate string `xml:"departureDate"`

	// Convey the begin time of a period. Format is hhmm.
	DepartureTime string `xml:"departureTime,omitempty"`  // minOccurs="0"

	// Convey the end date of a period. Format is ddmmyyyy.
	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"

	// Convey the end time of a period. Format is hhmm.
	ArrivalTime string `xml:"arrivalTime,omitempty"`  // minOccurs="0"
}

type ProductDateAndTimeTypeU_46325C struct {

	// Conveys departure date
	DepartureDate string `xml:"departureDate"`

	// Conveys departure time
	DepartureTime string `xml:"departureTime,omitempty"`  // minOccurs="0"

	// Conveys arrival date
	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"

	// Conveys arrival time
	ArrivalTime string `xml:"arrivalTime,omitempty"`  // minOccurs="0"
}

type ProductDateTimeTypeI struct {

	// Date format: DDMMYY
	DepartureDate string `xml:"departureDate,omitempty"`  // minOccurs="0"

	// Time format: 24H. All digits are mandatory . Example: from 0000 to 2359
	DepartureTime string `xml:"departureTime,omitempty"`  // minOccurs="0"

	// Date format: DDMMYY
	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"

	// Time format: 24H. All digits are mandatory . Example: from 0000 to 2359
	ArrivalTime string `xml:"arrivalTime,omitempty"`  // minOccurs="0"
}

type ProductDateTimeTypeI_171495C struct {

	// AIR segment : departure date ATX segment : requested date CAR segment : pick-up date CCR segment : pick-up date HHL segment : check-in date HTL segment : check-in date MIS segment : date for service requested SUR segment : date Trn Amtrak sgt: departure date Trn SNCF sgt: departure date TTO segment: departure date of the tour TUR segment: tour departure date INS element: departure date CRU segment: sailing departure date
	DepDate string `xml:"depDate"`

	// AIR segment : departure time SUR segment : pick-up time Trn Amtrak sgt: departure time Trn SNCF sgt: departure time
	DepTime string `xml:"depTime,omitempty"`  // minOccurs="0"

	// AIR segment : arrival date CAR segment : drop-off date CCR segment : return date HHL segment : check-out date HTL segment : check-out date TTO segment: return date of the tour INS element: return date
	ArrDate string `xml:"arrDate,omitempty"`  // minOccurs="0"

	// AIR segment : arrival time Trn Amtrak sgt: arrival time Trn SNCF sgt: arrival time
	ArrTime string `xml:"arrTime,omitempty"`  // minOccurs="0"

	// AIR segment: day change indicator (1,2,-1) TRN Amtrak sgt: day change indicator (1,2,-1) TRN SNCF sgt: day change indicator (1,2,-1)
	DayChangeIndicator *int32 `xml:"dayChangeIndicator,omitempty"`  // minOccurs="0"
}

type ProductDateTimeTypeI_260882C struct {

	// AIR segment : departure date ATX segment : requested date CAR segment : pick-up date CCR segment : pick-up date HHL segment : check-in date HTL segment : check-in date MIS segment : date for service requested SUR segment : date Trn Amtrak sgt: departure date Trn SNCF sgt: departure date TTO segment: departure date of the tour TUR segment: tour departure date INS element: departure date CRU segment: sailing departure date
	DepDate string `xml:"depDate,omitempty"`  // minOccurs="0"

	// AIR segment : departure time SUR segment : pick-up time Trn Amtrak sgt: departure time Trn SNCF sgt: departure time
	DepTime string `xml:"depTime,omitempty"`  // minOccurs="0"

	// AIR segment : arrival date CAR segment : drop-off date CCR segment : return date HHL segment : check-out date HTL segment : check-out date TTO segment: return date of the tour INS element: return date
	ArrDate string `xml:"arrDate,omitempty"`  // minOccurs="0"

	// AIR segment : arrival time Trn Amtrak sgt: arrival time Trn SNCF sgt: arrival time
	ArrTime string `xml:"arrTime"`

	// AIR segment: day change indicator (1,2,-1) TRN Amtrak sgt: day change indicator (1,2,-1) TRN SNCF sgt: day change indicator (1,2,-1)
	DayChangeIndicator *int32 `xml:"dayChangeIndicator,omitempty"`  // minOccurs="0"
}

type ProductDateTimeTypeI_46338C struct {

	// flight departure date
	DepartureDate string `xml:"departureDate"`

	// flight departure time
	DepartureTime string `xml:"departureTime"`

	// flight arrival date
	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"

	// flight arrival time
	ArrivalTime string `xml:"arrivalTime"`
}

type ProductDetailsTypeI struct {

	// booking class
	Designator string `xml:"designator"`
}

type ProductDetailsTypeI_118108C struct {

	// Class combination
	Designator string `xml:"designator"`

	// indicate availability status . coded or numeric
	AvailabilityStatus string `xml:"availabilityStatus,omitempty"`  // minOccurs="0"
}

type ProductDetailsTypeI_36664C struct {

	// Conveys the package code.
	Designator string `xml:"designator"`
}

type ProductFacilitiesTypeI struct {

	// Format limitations: an..3
	Entertainement string `xml:"entertainement,omitempty"`  // minOccurs="0"

	// For meal, the meal codes follow the IATA meal code standard
	EntertainementDescription string `xml:"entertainementDescription,omitempty"`  // minOccurs="0"

	// Format limitations: an..2
	ProductQualifier string `xml:"productQualifier,omitempty"`  // minOccurs="0"

	// Format limitations: an..4
	ProductExtensionCode []string `xml:"productExtensionCode,omitempty"`  // minOccurs="0" maxOccurs="26"
}

type ProductIdentificationDetailsTypeI struct {

	// Format limitations: an..4
	FlightNumber string `xml:"flightNumber"`
}

type ProductIdentificationDetailsTypeI_2786C struct {

	// Flight number or OPEN - ARNK, car type, transportation type  (refer to VGTVD transaction), train number, insurance provider
	Identification string `xml:"identification"`

	// AIR segment : class of service TRN Amtrack segment : class of service (1 or 2 chars long). TRN SNCF segment : class of service.
	ClassOfService string `xml:"classOfService,omitempty"`  // minOccurs="0"

	// AIR segment : flight number alpha suffix : A, B, C, D, E. SUR segment : departure code : A or D.
	Subtype string `xml:"subtype,omitempty"`  // minOccurs="0"

	// AIR segment :  N for Night class
	Description string `xml:"description,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetailsTypeI_46336C struct {

	// flight number or transportation code
	FlightNumber string `xml:"flightNumber"`

	// booking class
	BookingClass string `xml:"bookingClass"`
}

type ProductIdentificationDetailsTypeU struct {

	// Product code
	Number string `xml:"number"`

	// Product Name
	Name string `xml:"name"`
}

type ProductIdentificationDetailsTypeU_46327C struct {

	// Conveys the product code
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Conveys the product type (accomodation, vehicule, transportation, cruise ...)
	Type string `xml:"type"`

	// Conveys the subType of a product (Chalet or Villa for accomodation, Transfert or ticket for supplementary services ...)
	SubType string `xml:"subType,omitempty"`  // minOccurs="0"

	// Conveys the product description
	Description string `xml:"description,omitempty"`  // minOccurs="0"
}

type ProductIdentificationTypeU struct {

	// product name and code to which prices data apply
	ProductData *ProductIdentificationDetailsTypeU `xml:"productData"`
}

type ProductInformationTypeI struct {

	// Conveys the package details.
	BookingClassDetails *ProductDetailsTypeI_36664C `xml:"bookingClassDetails"`
}

type ProductInformationTypeI_73824S struct {

	// Booking class
	BookingClassDetails *ProductDetailsTypeI `xml:"bookingClassDetails"`
}

type ProductInformationTypeI_76271S struct {

	// Booking Class Details
	BookingClassDetails []*ProductDetailsTypeI_118108C `xml:"bookingClassDetails,omitempty"`  // minOccurs="0" maxOccurs="26"
}

type ProductTypeDetailsTypeI struct {

	// AIR segment : Electronic ticketing indicator : ET for Electronic ticket candidate SUR segment : transportation zone number Amtrack segment : Equipement code  SNCF segment : train type (3 chars code)
	Detail string `xml:"detail,omitempty"`  // minOccurs="0"
}

type ProductTypeDetailsTypeI_46337C struct {

	// sequence indicator for connection
	FlightIndicator string `xml:"flightIndicator"`
}

type PropertyHeaderDetailsType struct {

	// 1. hotel Provider name
	ProviderName string `xml:"providerName,omitempty"`  // minOccurs="0"

	// 1. HHL segment:hotel Property Code (or ID) 2. HTL AY Direct Access segment: Property location
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// 1. HHL segment:hotel Property name. 2. HTL AY Direct Access segment: Hotel name. Alphanumeric type due to possible numeric values in the names.
	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type ProviderInformationType struct {

	// productcode
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Product name
	Name string `xml:"name,omitempty"`  // minOccurs="0"

	// Product Family Code
	ProductFamilyCode string `xml:"productFamilyCode,omitempty"`  // minOccurs="0"
}

type QuantityAndActionDetailsTypeU struct {

	// Quantity information
	Quantity *int32 `xml:"quantity,omitempty"`  // minOccurs="0"

	// Conveys the status code (HK, GK ...) of a booking, a product or a ticket
	StatusCode string `xml:"statusCode"`
}

type QuantityAndActionDetailsTypeU_56796C struct {

	// accommodation reservation mandatoty, optionnal, advised, not possible
	StatusCode string `xml:"statusCode"`
}

type QuantityAndActionTypeU struct {

	// Conveys quantity and status information
	QuantityActionDetails []*QuantityAndActionDetailsTypeU `xml:"quantityActionDetails"`  // maxOccurs="2"
}

type QuantityAndActionTypeU_32609S struct {

	// accommodation status
	AccoStatus *QuantityAndActionDetailsTypeU_56796C `xml:"accoStatus"`
}

type QuantityDetailsTypeI struct {

	// A for age
	Qualifier string `xml:"qualifier"`

	// Age = number of years(default) or monthes.
	Value int32 `xml:"value"`
}

type QuantityDetailsTypeI_142179C struct {

	// it will be L for Life time period
	Qualifier string `xml:"qualifier"`

	// duration expressed in Seconds during the consumer has to do the payment
	Value int32 `xml:"value"`

	// SEC for duration in seconds
	Unit string `xml:"unit"`
}

type QuantityDetailsTypeI_198209C struct {

	// -NOD Number of Doors -MOD Maximum number of Doors -NOS Number of Seats -MOD Number of Seats -NOB Number of Bags -VOB Volume of Boots
	Qualifier string `xml:"qualifier"`

	// Value number corresponding to the qualifier type
	Value int32 `xml:"value"`

	// DM3 or FT3 if applicable
	Unit string `xml:"unit,omitempty"`  // minOccurs="0"
}

type QuantityDetailsTypeI_46334C struct {

	// Quantity qualifier
	Qualifier string `xml:"qualifier"`

	// Quantity value
	Value int32 `xml:"value"`

	// Quantity unit
	Unit string `xml:"unit"`
}

type QuantityType struct {

	// To specify an appropriate quantity.
	QuantityDetails *QuantityDetailsTypeI_46334C `xml:"quantityDetails"`
}

type QuantityTypeI struct {

	// Estinated distance details
	QuantityDetails *QuantityDetailsTypeI_142179C `xml:"quantityDetails"`
}

type QuantityTypeI_65488S struct {

	// This composite is used to convey the quantity details
	QuantityDetails *QuantityDetailsTypeI `xml:"quantityDetails"`
}

type QuantityType_94558S struct {

	// To specify an appropriate quantity.
	QuantityDetails []*QuantityDetailsTypeI_142179C `xml:"quantityDetails"`  // maxOccurs="20"
}

type QueueDetailsType struct {

	// A (first) queue number
	QueueNum1 *int32 `xml:"queueNum1,omitempty"`  // minOccurs="0"

	// [2-7] characters
	QueueName string `xml:"queueName,omitempty"`  // minOccurs="0"
}

type QueueType struct {

	// Queue detail
	QueueDetail *QueueDetailsType `xml:"queueDetail,omitempty"`  // minOccurs="0"

	// Queue category detail
	CategoryDetail *GategoryType `xml:"categoryDetail,omitempty"`  // minOccurs="0"

	// date range
	DateRange *DateRangeType `xml:"dateRange,omitempty"`  // minOccurs="0"

	// Other queue information
	Informations *OtherInformationType `xml:"informations,omitempty"`  // minOccurs="0"
}

type RailLegDataType struct {

	// Information pertaining to the train product
	TrainProductInfo *TrainProductInformationType_32331S `xml:"trainProductInfo"`

	// Reservation Mandatory, Advised, Possible, Not Possible
	ReservableStatus *QuantityAndActionTypeU_32609S `xml:"reservableStatus,omitempty"`  // minOccurs="0"

	// Leg departure and arrival dates and times
	LegDateTime []*StructuredDateTimeInformationType_32362S `xml:"legDateTime"`  // maxOccurs="2"

	// Departure station location
	DepLocation *PlaceLocationIdentificationTypeU_32347S `xml:"depLocation"`

	// Arrival station location
	ArrLocation *PlaceLocationIdentificationTypeU_32347S `xml:"arrLocation"`

	// leg reference: leg order within the itinerary
	LegReference *ItemNumberTypeU_33258S `xml:"legReference"`
}

type RailSeatConfigurationType struct {

	// Seat space.
	SeatSpace string `xml:"seatSpace,omitempty"`  // minOccurs="0"

	// Coach type.
	CoachType string `xml:"coachType,omitempty"`  // minOccurs="0"

	// Seat equipment.
	SeatEquipment string `xml:"seatEquipment,omitempty"`  // minOccurs="0"

	// Seat position.
	SeatPosition string `xml:"seatPosition,omitempty"`  // minOccurs="0"

	// Seat direction.
	SeatDirection string `xml:"seatDirection,omitempty"`  // minOccurs="0"

	// Seat deck.
	SeatDeck string `xml:"seatDeck,omitempty"`  // minOccurs="0"

	// Special passenger information.
	SpecialPassengerType []string `xml:"specialPassengerType,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type RailSeatPreferencesType struct {

	// Selection of the type of seat request.
	SeatRequestFunction string `xml:"seatRequestFunction,omitempty"`  // minOccurs="0"

	// Seat smoking zone indicator.
	SmokingIndicator string `xml:"smokingIndicator,omitempty"`  // minOccurs="0"

	// Seat class details.
	ClassDetails *ClassDetailsType_52782C `xml:"classDetails,omitempty"`  // minOccurs="0"

	// Seat configuration details.
	SeatConfiguration *RailSeatConfigurationType `xml:"seatConfiguration,omitempty"`  // minOccurs="0"

	SleeperDescription *RailSleeperDescriptionType `xml:"sleeperDescription,omitempty"`  // minOccurs="0"
}

type RailSeatReferenceInformationType struct {

	// Rail seat reference information.
	RailSeatReferenceDetails *SeatReferenceInformationType `xml:"railSeatReferenceDetails,omitempty"`  // minOccurs="0"
}

type RailSleeperDescriptionType struct {

	// Berth deck
	BerthDeck string `xml:"berthDeck,omitempty"`  // minOccurs="0"

	// Cabin position
	CabinPosition string `xml:"cabinPosition,omitempty"`  // minOccurs="0"

	// Cabin share type
	CabinShareType string `xml:"cabinShareType,omitempty"`  // minOccurs="0"

	// Cabin occupancy
	CabinOccupancy string `xml:"cabinOccupancy,omitempty"`  // minOccurs="0"
}

type RangeDetailsTypeI struct {

	// 701 for range definition
	RangeQualifier string `xml:"rangeQualifier"`

	// Range definition
	RangeDetails *RangeTypeI `xml:"rangeDetails"`
}

type RangeDetailsTypeU struct {

	// Range qualifier
	RangeQualifier string `xml:"rangeQualifier"`

	// Range details
	RangeDetails *RangeTypeU `xml:"rangeDetails"`
}

type RangeTypeI struct {

	// Duration qualifier: - DAY Duration in days - WE  Duration in weeks - MTH Duration in months - G Kilometers - M Mileage - A Age
	DataType string `xml:"dataType"`

	// Base of the Range
	Min int32 `xml:"min"`

	// Top of the Range
	Max int32 `xml:"max"`
}

type RangeTypeU struct {

	// Range data type
	DataType string `xml:"dataType"`

	// min Occupancy
	MinOccupancy *int32 `xml:"minOccupancy,omitempty"`  // minOccurs="0"

	// Occupancy maximum
	MaxOccupancy int32 `xml:"maxOccupancy"`
}

type RateCodeRestrictedType struct {

	// 1. HHL segment: hotel  Rate code (an3) 2. For AY Direct Access segment: Rate type =  MINR, MODR, MAXR, ADVR, DAYR, SRTE
	RateCode string `xml:"rateCode,omitempty"`  // minOccurs="0"
}

type RateIndicatorsType struct {

	// 1. HTL AY Direct Access segment:  Y for Yes  (rate change)
	RateChangeIndicator string `xml:"rateChangeIndicator,omitempty"`  // minOccurs="0"
}

type RateInformationDetailsType struct {

	// 1. Hotel segment:  Total or daily indicator
	RatePlan string `xml:"ratePlan,omitempty"`  // minOccurs="0"
}

type RateInformationType struct {

	// Rate Price
	RatePrice *RatePriceType `xml:"ratePrice,omitempty"`  // minOccurs="0"

	// Rate information
	RateInfo *RateInformationDetailsType `xml:"rateInfo,omitempty"`  // minOccurs="0"

	// Rate indicator
	RateIndicator *RateIndicatorsType `xml:"rateIndicator,omitempty"`  // minOccurs="0"
}

type RateInformationTypeI struct {

	// Rate Category.
	Category string `xml:"category,omitempty"`  // minOccurs="0"
}

type RateInformationTypeI_198204C struct {

	// Rate Category 002 Inclusive 006 Convention 007 Corporate 009 Government 011 Package 019 Association 020 Business 021 Consortium 022 Credential 023 Industry 024 Standard G General
	Category string `xml:"category"`
}

type RateInformationTypeI_50732C struct {

	// Fare Group
	FareGroup string `xml:"fareGroup,omitempty"`  // minOccurs="0"
}

type RatePriceType struct {

	// 1. Hotel segment:  Rate value 2. Hotel AY Direct Access segment: Room rate (imbedded decimal point)
	RateAmount *float64 `xml:"rateAmount,omitempty"`  // minOccurs="0"
}

type RateTypesTypeU struct {

	// This element holds the rate code that applies to the Ferry booking.
	RateCode string `xml:"rateCode"`
}

type ReferenceInfoType struct {

	// This composite is used to transmit association information
	Reference []*ReferencingDetailsType_111975C `xml:"reference,omitempty"`  // minOccurs="0" maxOccurs="198"
}

type ReferenceInfoType_25422S struct {

	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsTypeI_46317C `xml:"referenceDetails"`
}

type ReferenceInfoType_6074S struct {

	// This composite is used to transmit association information
	Reference []*ReferencingDetailsType `xml:"reference,omitempty"`  // minOccurs="0" maxOccurs="198"
}

type ReferenceInfoType_94524S struct {

	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsType_142140C `xml:"referenceDetails"`
}

type ReferenceInfoType_94566S struct {

	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsType_142187C `xml:"referenceDetails,omitempty"`  // minOccurs="0"
}

type ReferenceInformationType struct {

	// Used to specify the passenger association and the data per passanger.
	ReferenceDetails *ReferencingDetailsTypeI_17164C `xml:"referenceDetails,omitempty"`  // minOccurs="0"
}

type ReferenceInformationTypeI struct {

	// Details of the referencing
	ReferenceDetails *ReferencingDetailsTypeI_185716C `xml:"referenceDetails"`
}

type ReferenceInformationTypeI_136704S struct {

	// Use to convey the reference details
	ReferenceDetails *ReferencingDetailsTypeI_198199C `xml:"referenceDetails"`
}

type ReferenceInformationTypeI_25132S struct {

	// Conveys the passenger reference.
	ReferenceDetails []*ReferencingDetailsTypeI_45901C `xml:"referenceDetails"`  // maxOccurs="9"
}

type ReferenceInformationTypeI_83551S struct {

	// Reference details
	ReferenceDetails *ReferencingDetailsTypeI_127514C `xml:"referenceDetails"`
}

type ReferenceInformationTypeI_94503S struct {

	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsTypeI `xml:"referenceDetails"`
}

type ReferenceInformationType_65487S struct {

	// Used to convey the passenger tatoo or display number.
	PassengerReference *ReferencingDetailsTypeI `xml:"passengerReference"`
}

type ReferencingDetailsType struct {

	// Amadeus codes are used here.  PT for Passenger Tatoo // ST for Segment Tatoo //OT for Other element Tatoo //SS for Segment Tatoo+SubTatoo
	Qualifier string `xml:"qualifier"`

	// reference number refers to a PNR segment/element that has this number in its related element reference segment in the same message (qualifier PT, SS, ST).
	Number string `xml:"number"`
}

type ReferencingDetailsTypeI struct {

	// Qualifier of the type of reference.
	Type string `xml:"type"`

	// Value of the association reference
	Value string `xml:"value"`
}

type ReferencingDetailsTypeI_127514C struct {

	// Format limitations: an..3
	Type string `xml:"type"`

	// Format limitations: an..10
	Value string `xml:"value"`
}

type ReferencingDetailsTypeI_17164C struct {

	// to specify the segment association
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// used to identify the segment(s) from which we want to extract the data.
	Value string `xml:"value"`
}

type ReferencingDetailsTypeI_185716C struct {

	// A code which identifies the type of identifier that is used.
	Type string `xml:"type"`
}

type ReferencingDetailsTypeI_198199C struct {

	// Reference qualifier Amadeus codes :  OT for Other element(non name, non segment) Tatoo   PT for Passenger Tatoo   ST for Segment Tatoo   SS for Segment Tatoo+SubTatoo
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Reference number Number attributed by the Server to reference the PNR segment/element Limited to the time the PNR is worked (First retrieve - End of Transaction)
	Value string `xml:"value,omitempty"`  // minOccurs="0"
}

type ReferencingDetailsTypeI_36941C struct {

	// Coach Number
	Value string `xml:"value"`
}

type ReferencingDetailsTypeI_45901C struct {

	// Qualifies the type of reference used.
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Conveys the passenger sequence number.
	Value string `xml:"value"`
}

type ReferencingDetailsTypeI_46317C struct {

	// Qualify the type of reference: passenger or product
	Type string `xml:"type"`

	// Passenger tatoo or Product sequence number
	Value string `xml:"value"`
}

type ReferencingDetailsType_111975C struct {

	// Amadeus codes are used here.  PT for Passenger Tatoo // ST for Segment Tatoo //OT for Other element Tatoo //SS for Segment Tatoo+SubTatoo
	Qualifier string `xml:"qualifier"`

	// reference number refers to a PNR segment/element that has this number in its related element reference segment in the same message (qualifier PT, SS, ST).
	Number string `xml:"number"`
}

type ReferencingDetailsType_127526C struct {

	// Shopping Basket codes : CDS Shopping Basket Distribution record CRR Shopping Basket Reservation RecordCST Shopping Basket customer DOC Shopping Basket document FAR Shopping Basket fares and fees information FFY Shopping Basket frequent flyer information FOP Shopping Basket form of payment PRD Shopping Basket product RMK Shopping Basket remark SBK Shopping Basket (used in search results)
	Qualifier string `xml:"qualifier"`

	// Number attributed by the Server to reference the shopping basket item.
	Number int32 `xml:"number"`
}

type ReferencingDetailsType_142140C struct {

	// will have the following values: XID  Transaction identifier of the 3DS process CAAV authentication verification code for Visa AAV  authentication verification code for MasterCard PAREQ authentication message PARES authentication response message
	Value string `xml:"value"`
}

type ReferencingDetailsType_142187C struct {

	// FOID document type
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// FOID document number
	Value string `xml:"value,omitempty"`  // minOccurs="0"
}

type ReferencingDetailsType_2780C struct {

	// Amadeus codes are used here.  D for Dominant segment in a marriage  N for Non dominant segment in a marriage
	MarriageQualifier string `xml:"marriageQualifier,omitempty"`  // minOccurs="0"

	// Tatoo number of the segment
	TatooNum string `xml:"tatooNum"`
}

type RelatedLocationOneIdentificationTypeU struct {

	// Conveys the embarkation port code.
	Code string `xml:"code"`
}

type RelatedLocationOneIdentificationTypeU_198193C struct {

	// Assiciated airport code.
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Associated airport code qualifier.
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Set to IA to indicate that the associated location code is a IATA airport or city code.
	Agency string `xml:"agency,omitempty"`  // minOccurs="0"
}

type RelatedLocationOneIdentificationTypeU_45087C struct {

	// Area code.
	Code string `xml:"code"`
}

type RelatedLocationOneIdentificationTypeU_46345C struct {

	// location code
	Code string `xml:"code"`

	// location qualifier
	Qualifier string `xml:"qualifier"`
}

type RelatedLocationOneIdentificationTypeU_56455C struct {

	// Railway station country code
	Code string `xml:"code"`

	// Code type
	Qualifier string `xml:"qualifier"`
}

type RelatedLocationTwoIdentificationTypeU struct {

	// Conveys the disembarkation port code.
	Code string `xml:"code"`
}

type RelatedProductInformationTypeI struct {

	// No quantity returned
	Quantity *int32 `xml:"quantity,omitempty"`  // minOccurs="0"

	// see code list
	Status []string `xml:"status"`  // maxOccurs="2"
}

type ReservationControlInformationDetailsTypeI struct {

	// 1A or Other airline record locator information  Passive segment airline code
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	// 1. Record - 1A record locator or - OA record locator
	ControlNumber string `xml:"controlNumber,omitempty"`  // minOccurs="0"

	// 1. Profile record locator information: Customer type:  C for Corporate  T for Traveler     F for Frequent Flyer
	ControlType string `xml:"controlType,omitempty"`  // minOccurs="0"

	// 1. RR element: Date 2. SP element: Date 3. PNR header/RP line: Date of last End of transaction
	Date string `xml:"date,omitempty"`  // minOccurs="0"

	// 1. PNR header/RP line: time of last End of transaction
	Time string `xml:"time,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI_16352C struct {

	// Conveys the booking number.
	ControlNumber string `xml:"controlNumber"`

	// Conveys the booking number qualifier.
	ControlType string `xml:"controlType"`
}

type ReservationControlInformationDetailsTypeI_170724C struct {

	// 1A or Other airline record locator information  Passive segment airline code
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	// 1. Record - 1A record locator or - OA record locator  Record locator is truncated to 7 characters.
	ControlNumber string `xml:"controlNumber,omitempty"`  // minOccurs="0"

	// 1. Profile record locator information: Customer type:  C for Corporate  T for Traveler     F for Frequent Flyer
	ControlType string `xml:"controlType,omitempty"`  // minOccurs="0"

	// 1. RR element: Date 2. SP element: Date 3. PNR header/RP line: Date of last End of transaction
	Date *int32 `xml:"date,omitempty"`  // minOccurs="0"

	// 1. PNR header/RP line: time of last End of transaction
	Time string `xml:"time,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI_18446C struct {

	// this is used to specify the confirmation number meaning that the booking was succesfull.
	ControlNumber string `xml:"controlNumber,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI_198198C struct {

	// - 1A or Other airline record locator information - Passive segment airline code
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	// - 1A record locator or - OA record locator
	ControlNumber string `xml:"controlNumber,omitempty"`  // minOccurs="0"

	// PNR split type.
	ControlType string `xml:"controlType,omitempty"`  // minOccurs="0"

	// 1. RR element: Date 2. SP element: Date 3. PNR header/RP line: Date of lastest End of transaction
	Date *int32 `xml:"date,omitempty"`  // minOccurs="0"

	// 1. PNR header/RP line: time of lastest End of transaction
	Time *int32 `xml:"time,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI_35709C struct {

	// This element conveys the booking number which is used as a booking reference by the Ferry provider.
	ControlNumber string `xml:"controlNumber"`
}

type ReservationControlInformationDetailsTypeU struct {

	// Conveys the tour operator code
	TourOperatorCode string `xml:"tourOperatorCode,omitempty"`  // minOccurs="0"

	// Conveys the reservation control number qualifier
	ReservationControlNumberQual string `xml:"reservationControlNumberQual"`

	// Conveys the reservation control number. Can have up to 32 chars. for Tour Server
	ReservationControlNumber string `xml:"reservationControlNumber"`
}

type ReservationControlInformationDetailsTypeU_55378C struct {

	// The individual Passenger confirmation number in the Provider database.
	Value string `xml:"value"`
}

type ReservationControlInformationTypeI struct {

	// To specify the confirmation number in case the booking was succesfull
	Reservation *ReservationControlInformationDetailsTypeI_18446C `xml:"reservation,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationTypeI_115879S struct {

	// Reservation Information
	Reservation *ReservationControlInformationDetailsTypeI_170724C `xml:"reservation,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationTypeI_136703S struct {

	// Reservation Information
	Reservation *ReservationControlInformationDetailsTypeI_198198C `xml:"reservation,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationTypeI_20153S struct {

	// Provides details of the Ferry booking number. The booking number is a unique reference per provider per booking in the provider system. As such, it is stored in the PNR in all the legs of the same booking and it is used in the Ferry PNR indexing.
	Reservation *ReservationControlInformationDetailsTypeI_35709C `xml:"reservation"`
}

type ReservationControlInformationTypeI_87792S struct {

	// Reservation Information
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationTypeI_8957S struct {

	// Cruise booking reference.
	Reservation []*ReservationControlInformationDetailsTypeI_16352C `xml:"reservation"`  // maxOccurs="2"
}

type ReservationControlInformationTypeU struct {

	// Conveys the reservation control Id
	ReservationControlId *ReservationControlInformationDetailsTypeU `xml:"reservationControlId"`
}

type ReservationControlInformationTypeU_31804S struct {

	// The reference to the Provider Database
	ReferenceDetails *ReservationControlInformationDetailsTypeU_55378C `xml:"referenceDetails"`
}

type ReservationSecurityInformationType struct {

	// Responsibility Information
	ResponsibilityInformation *ResponsibilityInformationType `xml:"responsibilityInformation,omitempty"`  // minOccurs="0"

	// Ticket Information
	QueueingInformation *TicketInformationType_5120C `xml:"queueingInformation,omitempty"`  // minOccurs="0"

	// 1. PNR Header: Pseudo City Code (not in the CRT display) AGY for Travel agency EHD for First level Help Desk DAP for Data processing center / Amadeus Help Desk Nice SEC for Security administrator WZ for AIS security administrator
	CityCode string `xml:"cityCode,omitempty"`  // minOccurs="0"

	// Second RP line information
	SecondRpInformation *SecondRpLineInformationType `xml:"secondRpInformation,omitempty"`  // minOccurs="0"
}

type ReservationSecurityInformationType_167761S struct {

	// Responsibility Information
	ResponsibilityInformation *ResponsibilityInformationType_6835C `xml:"responsibilityInformation,omitempty"`  // minOccurs="0"

	// Ticket Information
	QueueingInformation *TicketInformationType_5120C `xml:"queueingInformation,omitempty"`  // minOccurs="0"

	// 1. PNR Header: Pseudo City Code (not in the CRT display) AGY for Travel agency EHD for First level Help Desk DAP for Data processing center / Amadeus Help Desk Nice SEC for Security administrator WZ for AIS security administrator
	CityCode string `xml:"cityCode,omitempty"`  // minOccurs="0"

	// Second RP line information
	SecondRpInformation *SecondRpLineInformationType_237255C `xml:"secondRpInformation,omitempty"`  // minOccurs="0"
}

type ResponseIdentificationType struct {

	// Transaction identifier   Field 62.2  Official definition: Visa-generated identifier that is unique for each original transaction. The transaction identifier (TID) is a key element that links original authorization requests to subsequent messages, such as reversals.
	TransacIdentifier string `xml:"transacIdentifier,omitempty"`  // minOccurs="0"

	// Validation code    Field 62.3
	ValidationCode string `xml:"validationCode,omitempty"`  // minOccurs="0"

	// Gateway Transaction Identifier - Banknet reference number   Field 62.17 - Position 8-13
	BanknetRefNumber string `xml:"banknetRefNumber,omitempty"`  // minOccurs="0"

	// Gateway Transaction Identifier - Banknet date in mmdd format  Field 62.17 - Position 1-4
	BanknetDate string `xml:"banknetDate,omitempty"`  // minOccurs="0"
}

type ResponsibilityInformationType struct {

	// Type of PNR element: - RR for Associated Cross Reference Record - SP for Split Party - RP for PNR Header line
	TypeOfPnrElement string `xml:"typeOfPnrElement"`

	// 1. RR element: 2. SP element: 3. PNR Header:Agent initials and duty code (eg: AASU)
	AgentId string `xml:"agentId,omitempty"`  // minOccurs="0"

	// 1. RR element office that copied the PNR 2. SP element: office that split the PNR 3. PNR Header: office responsibility or - OA office (City code + OA code)  which is 5 chars long
	OfficeId string `xml:"officeId,omitempty"`  // minOccurs="0"

	// ATA/IATA reference number assigned to a travel agent
	IataCode *int32 `xml:"iataCode,omitempty"`  // minOccurs="0"
}

type ResponsibilityInformationType_6835C struct {

	// Type of PNR element: - RR for Associated Cross Reference Record - SP for Split Party - RP for PNR Header line
	TypeOfPnrElement string `xml:"typeOfPnrElement"`

	// 1. RR element: 2. SP element: 3. PNR Header:Agent initials and duty code (eg: AASU)
	AgentId string `xml:"agentId,omitempty"`  // minOccurs="0"

	// 1. RR element office that copied the PNR 2. SP element: office that split the PNR 3. PNR Header: office responsibility or - OA office (City code + OA code)  which is 5 chars long
	OfficeId string `xml:"officeId,omitempty"`  // minOccurs="0"

	// ATA/IATA reference number assigned to a travel agent
	IataCode *int32 `xml:"iataCode,omitempty"`  // minOccurs="0"
}

type RoomDetailsType struct {

	// 1. room Occupancy
	Occupancy *int32 `xml:"occupancy,omitempty"`  // minOccurs="0"

	// room Type
	TypeCode string `xml:"typeCode,omitempty"`  // minOccurs="0"
}

type RuleDetailsTypeU struct {

	// This data element is used to identify the type of rule (see data mapping).
	Type string `xml:"type"`

	// This data element is used to convey the afternoon time which is the maximum time to check-in. ex. if equal to 1, it means that the room is kept until 1PM.
	Quantity int32 `xml:"quantity"`

	// This data element is used to specify that the Maximim check-in time is given in hours.
	QuantityUnit string `xml:"quantityUnit"`
}

type RuleDetailsTypeU_198224C struct {

	// Coded rule type
	Type string `xml:"type"`

	// quantity (if applicable)
	Quantity *int32 `xml:"quantity,omitempty"`  // minOccurs="0"

	// DAY for Day HOR for Hour (if applicable)
	QuantityUnit string `xml:"quantityUnit,omitempty"`  // minOccurs="0"

	// Deposit Information: - BRE Before Rental - AFT After Booking Pickup Information - MAX Maximum Days Rental - MIN Minimum Days Rental One Way Information: - 009 for One Way Allowed - 005 for One Way not Allowed - 006 for Restricted One Way Allowed
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// Day of the week (Monday=1, Sunday=7)
	DaysOfOperation string `xml:"daysOfOperation,omitempty"`  // minOccurs="0"

	// Rule amount (if applicable)
	Amount *int32 `xml:"amount,omitempty"`  // minOccurs="0"

	// Rule currency amount (if applicable)
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type RuleInformationTypeU struct {

	// This composite is used to convey the rules details.
	RuleDetails *RuleDetailsTypeU `xml:"ruleDetails"`

	// This composite is used to indicate that the rule have been entered manually by the agent.
	RuleStatusDetails *RuleStatusTypeU `xml:"ruleStatusDetails,omitempty"`  // minOccurs="0"
}

type RuleInformationTypeU_136720S struct {

	// Rule details
	RuleDetails []*RuleDetailsTypeU_198224C `xml:"ruleDetails"`  // maxOccurs="5"

	// Associated Rule Text
	RuleText *RuleTextTypeU `xml:"ruleText,omitempty"`  // minOccurs="0"
}

type RuleStatusTypeU struct {

	// This data element specifies the rule concerned by this information.
	StatusType string `xml:"statusType"`

	// This data element is used to indicate that the condition have been entered manually by the travel agent and is not coming from the supplier data.
	ProcessIndicator string `xml:"processIndicator,omitempty"`  // minOccurs="0"
}

type RuleTextTypeU struct {

	// Coded rule type
	TextType string `xml:"textType"`

	// Rule Information
	FreeText []string `xml:"freeText"`  // maxOccurs="20"
}

type SeatReferenceInformationType struct {

	// Coach number.
	CoachNumber string `xml:"coachNumber,omitempty"`  // minOccurs="0"

	// Deck number.
	DeckNumber string `xml:"deckNumber,omitempty"`  // minOccurs="0"

	// Seat number.
	SeatNumber string `xml:"seatNumber,omitempty"`  // minOccurs="0"
}

type SeatRequestParametersTypeI struct {

	// Details of the seat
	GenericDetails *GenericDetailsTypeI `xml:"genericDetails,omitempty"`  // minOccurs="0"
}

type SecondRpLineInformationType struct {

	// Creation office
	CreationOfficeId string `xml:"creationOfficeId"`

	// Creation agent sine/queue category (eg: 1234AA)
	AgentSignature string `xml:"agentSignature,omitempty"`  // minOccurs="0"

	// PNR creation date
	CreationDate string `xml:"creationDate"`

	// ATA/IATA number assigned to a travel agent
	CreatorIataCode *int32 `xml:"creatorIataCode,omitempty"`  // minOccurs="0"

	// PNR creation time
	CreationTime string `xml:"creationTime,omitempty"`  // minOccurs="0"
}

type SecondRpLineInformationType_237255C struct {

	// Creation office
	CreationOfficeId string `xml:"creationOfficeId"`

	// Creation agent sine/queue category (eg: 1234AA)
	AgentSignature string `xml:"agentSignature,omitempty"`  // minOccurs="0"

	// PNR creation date
	CreationDate string `xml:"creationDate"`

	// ATA/IATA number assigned to a travel agent
	CreatorIataCode *int32 `xml:"creatorIataCode,omitempty"`  // minOccurs="0"

	// PNR creation time
	CreationTime string `xml:"creationTime,omitempty"`  // minOccurs="0"
}

type SecurityInformationType struct {

	// Date of creation
	CreationDate string `xml:"creationDate"`

	// Agent initials and duty code as in Originator informations  (eg: AASU)
	AgentCode string `xml:"agentCode"`

	// Office Id of creation/update
	OfficeId string `xml:"officeId,omitempty"`  // minOccurs="0"
}

type SegmentCabinIdentificationType struct {

	// Cabin class designator
	CabinCode string `xml:"cabinCode"`
}

type SegmentGroupingInformationType struct {

	// Type of segment grouping:  Mxx for Marriage (see codeset)  CNX for Connection
	GroupingCode string `xml:"groupingCode"`

	// transmit the list of segments participating in one marriage or segments that are connected.
	MarriageDetail []*ReferencingDetailsType_2780C `xml:"marriageDetail,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type SelectionDetailsInformationTypeI struct {

	// see code list
	Option string `xml:"option"`

	// CRU segment - occurrence 1 : Duration of the cruise (saling length)  expressed in days.
	OptionInformation string `xml:"optionInformation,omitempty"`  // minOccurs="0"
}

type SelectionDetailsInformationTypeI_198215C struct {

	// - P6 for seamless availability
	Option string `xml:"option"`
}

type SelectionDetailsTypeI struct {

	// Booking access type requested
	SelectionDetails *SelectionDetailsInformationTypeI_198215C `xml:"selectionDetails"`
}

type SelectionDetailsTypeI_2067S struct {

	// Only the first occurrence of the composite is mandatory. Up to 10 occurrences of the composite.
	Selection []*SelectionDetailsInformationTypeI `xml:"selection"`  // maxOccurs="10"
}

type SequenceDetailsTypeU struct {

	// Format limitations: an..3
	ActionRequest string `xml:"actionRequest,omitempty"`  // minOccurs="0"

	SequenceDetails *SequenceInformationTypeU_24073C `xml:"sequenceDetails"`
}

type SequenceDetailsTypeU_94494S struct {

	// Sequence Information
	SequenceDetails *SequenceInformationTypeU `xml:"sequenceDetails,omitempty"`  // minOccurs="0"
}

type SequenceInformationTypeU struct {

	// Sequence number of the Mean Of Payment in the FOP line. There can be up to 3 New MOP and 3 Old MOP in a FOP line.  Old Fops are considered as freeflow text.
	Number int32 `xml:"number"`
}

type SequenceInformationTypeU_24073C struct {

	// Format limitations: an..10
	Number string `xml:"number"`
}

type ShipIdentificationDetailsType struct {

	// Used to convey the ship code as in the Cruise specific database ship's table.
	Code string `xml:"code"`

	// Used to convey the ship name as in the Cruise specific database ship's table.
	Name string `xml:"name,omitempty"`  // minOccurs="0"

	// Used to convey the cruise line provider code for the sailing ship.
	CruiseLineCode string `xml:"cruiseLineCode"`
}

type ShipIdentificationDetailsType_45069C struct {

	// Used to convey the ship name.
	Name string `xml:"name,omitempty"`  // minOccurs="0"
}

type ShipIdentificationType struct {

	// Detailed information for the sailing ship.
	ShipDetails *ShipIdentificationDetailsType_45069C `xml:"shipDetails"`
}

type ShipIdentificationType_8952S struct {

	// Detailed information for the sailing ship.
	ShipDetails *ShipIdentificationDetailsType `xml:"shipDetails"`
}

type SpecialRequirementsDataDetailsType struct {

	// The seat number
	SeatNumber string `xml:"seatNumber,omitempty"`  // minOccurs="0"

	// type of the seat
	SeatCharacteristic []string `xml:"seatCharacteristic,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type SpecialRequirementsDataDetailsTypeI struct {

	// Seat number + row (seat SSR)  Number of seats (Group seat SSR)
	Data string `xml:"data,omitempty"`  // minOccurs="0"

	// Refers a Traveller / Reference number for association purpose
	CrossRef string `xml:"crossRef,omitempty"`  // minOccurs="0"

	// 3 occurrences may be used for in Amadeus seat SSR to indicate: 1. Smoking/no smoking  2. 1st area preference   3. 2nd area preference or passenger type
	SeatType []string `xml:"seatType,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type SpecialRequirementsDetailsType struct {

	// To specify the Seat Number.
	SeatDetails *SpecialRequirementsDataDetailsType `xml:"seatDetails"`
}

type SpecialRequirementsDetailsTypeI struct {

	// Special requirements type details
	Ssr *SpecialRequirementsTypeDetailsTypeI `xml:"ssr"`

	// Group seat SSR cannot ask for specific seats but only smoking and/or non-smoking (see Group seat SSR). the maximum repetitions here are 9 seats (1 per passenger of non-group PNR).
	Ssrb []*SpecialRequirementsDataDetailsTypeI `xml:"ssrb,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type SpecialRequirementsTypeDetailsTypeI struct {

	// ATA/IATA defined Special Service Requirement code.  (refer to IATA AIRIMP documentation)
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Use defined code or an ATA/IATA defined action code (See AIRIMP 7.1.2/7.1.3/7.1.4/8.14.1 (as bilaterally agreed), SIPP 105.170.1.1).
	Status string `xml:"status,omitempty"`  // minOccurs="0"

	// Number of services requested
	Quantity *int32 `xml:"quantity,omitempty"`  // minOccurs="0"

	// Airline code or YY
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	// Seat Special service request  or  Frequent Flyer SSR.
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"

	// 1. Seat SSR Processing indicator, coded  PS for Partial segment  indicator
	ProcessingIndicator string `xml:"processingIndicator,omitempty"`  // minOccurs="0"

	// Board point
	Boardpoint string `xml:"boardpoint,omitempty"`  // minOccurs="0"

	// Off point
	Offpoint string `xml:"offpoint,omitempty"`  // minOccurs="0"

	// Free flow of the SSR that can be up to 127 chars long, therefore split on two 4440 (70 + 57)
	FreeText []string `xml:"freeText,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type SpecificDataInformationTypeI struct {

	// This composite holds information about the element to which it applies.  For the Ferry business, this element specifies that the type of information conveyed is about the animal(s) attached to a Ferry booking.
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`

	// Details and description of the conveyed information.
	DataInformation *DataInformationTypeI `xml:"dataInformation"`
}

type SpecificTravellerDetailsType struct {

	// passenger type indicator
	PassengerType string `xml:"passengerType"`

	// Format limitations: a..60
	TravellerSurname string `xml:"travellerSurname,omitempty"`  // minOccurs="0"

	// Format limitations: a..60
	TravellerGivenName string `xml:"travellerGivenName,omitempty"`  // minOccurs="0"

	// TravellerReferenceNumber
	TravellerReferenceNumber string `xml:"travellerReferenceNumber,omitempty"`  // minOccurs="0"

	// birthdate or age of passenger
	PassengerBirthdate string `xml:"passengerBirthdate,omitempty"`  // minOccurs="0"
}

type SpecificVisaLinkCreditCardInformationType struct {

	// ISO8583 specific info
	MsgRef *MessageReferenceType `xml:"msgRef,omitempty"`  // minOccurs="0"

	// Response identification
	RespIdentification *ResponseIdentificationType `xml:"respIdentification,omitempty"`  // minOccurs="0"
}

type StationInformationTypeI struct {

	// Format limitations: an..2
	DepartTerminal string `xml:"departTerminal,omitempty"`  // minOccurs="0"
}

type StationInformationTypeI_119771C struct {

	// Format limitations: an..2
	Terminal string `xml:"terminal,omitempty"`  // minOccurs="0"
}

type StatusDetailsType struct {

	// FRA for fraud screening
	Indicator string `xml:"indicator"`

	// This data element is used to indicate if risk management must be performed at authorization time: - Y means risk management data will be appended to author; - N means risk management data will not be appended;
	Action string `xml:"action"`
}

type StatusDetailsTypeI struct {

	// Indicator name.
	Indicator string `xml:"indicator"`
}

type StatusDetailsTypeI_185722C struct {

	// Status of the entity
	Indicator string `xml:"indicator"`

	// Qualifies the status
	Type string `xml:"type"`
}

type StatusDetailsTypeI_20684C struct {

	// Coded identifying type of Information
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"

	// Data information upon qualifier.
	Description string `xml:"description,omitempty"`  // minOccurs="0"
}

type StatusDetailsTypeI_37285C struct {

	// Indicates of the reservation is modifiable directly in the 1A system
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"
}

type StatusDetailsTypeI_57035C struct {

	// Indicator name.
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"

	// Contains "MOD" if the PNR has been modifed since it has been retrieved
	IsPNRModifDuringTrans string `xml:"isPNRModifDuringTrans,omitempty"`  // minOccurs="0"
}

type StatusDetailsType_148479C struct {

	// list of status/qualifiers Either His for Historical or     Crt for Current
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"
}

type StatusDetailsType_148716C struct {

	// Indicator name.
	Indicator string `xml:"indicator"`
}

type StatusDetailsType_183347C struct {

	// list of status/qualifiers Either His for Historical or     Crt for Current
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"

	// Conveys any additional data necessary to qualify the indicator
	Description string `xml:"description,omitempty"`  // minOccurs="0"
}

type StatusType struct {

	// STATUS DETAILS
	StatusInformation []*StatusDetailsType_183347C `xml:"statusInformation"`  // maxOccurs="99"
}

type StatusTypeI struct {

	// Provides a set of coded characteristics of the customer.
	StatusDetails []*StatusDetailsTypeI_185722C `xml:"statusDetails"`  // maxOccurs="9"
}

type StatusTypeI_13270S struct {

	// Provides information on the type of fop.
	StatusDetails *StatusDetailsTypeI_20684C `xml:"statusDetails"`

	OtherStatusDetails []*StatusDetailsTypeI_20684C `xml:"otherStatusDetails,omitempty"`  // minOccurs="0" maxOccurs="98"
}

type StatusTypeI_20923S struct {

	// Status information
	StatusDetails *StatusDetailsTypeI_37285C `xml:"statusDetails"`
}

type StatusTypeI_32775S struct {

	// Contains general indicators relative to the state of the PNR
	StatusDetails *StatusDetailsTypeI_57035C `xml:"statusDetails"`
}

type StatusTypeI_33257S struct {

	// indicates an open segment
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails"`
}

type StatusType_94568S struct {

	// will we perform the fraud screening ?
	StatusInformation *StatusDetailsType `xml:"statusInformation"`
}

type StatusType_99582S struct {

	// STATUS DETAILS
	StatusInformation []*StatusDetailsType_148479C `xml:"statusInformation,omitempty"`  // minOccurs="0" maxOccurs="4"
}

type StatusType_99946S struct {

	// STATUS DETAILS
	StatusInformation *StatusDetailsType_148716C `xml:"statusInformation"`
}

type StructuredAddressInformationType struct {

	// Following values are : CY for Company NA for Name L1 for Address line 1 L2 for Address line 2 PO for P.O. BOX ZP for Postal code CI for City ST for State CO for Country
	Option string `xml:"option"`

	// Alphanumeric information related to the level code.  Each code has its own max length, an..50 corresponds to the max length among.
	OptionText string `xml:"optionText"`
}

type StructuredAddressType struct {

	// Information type, coded  2  for billing address  P08  for general mailing address  P19  for miscellaneous mailing address  P24  for home mailing address  P25  for delivery mailing address
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// Structured Address
	Address []*StructuredAddressInformationType `xml:"address,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type StructuredDateTimeInformationType struct {

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_17997C `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeInformationType_20644S struct {

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_36775C `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeInformationType_20645S struct {

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_36777C `xml:"dateTime"`
}

type StructuredDateTimeInformationType_21109S struct {

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_35730C `xml:"dateTime"`
}

type StructuredDateTimeInformationType_24436S struct {

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_44876C `xml:"dateTime"`
}

type StructuredDateTimeInformationType_25444S struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic string `xml:"businessSemantic"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_44876C `xml:"dateTime"`
}

type StructuredDateTimeInformationType_27086S struct {

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_16347C `xml:"dateTime"`
}

type StructuredDateTimeInformationType_32362S struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	// Departure or arrival date and time.
	DateTime *StructuredDateTimeType_56472C `xml:"dateTime"`
}

type StructuredDateTimeInformationType_83553S struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic string `xml:"businessSemantic"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime"`
}

type StructuredDateTimeInformationType_94516S struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - LT : date and time corresponding to Authorization message built - T : date and time corresponding to Authorization message sent - AR : date and time corresponding to Authorization message receipt
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	// Indicate if the time is expressed in UTC or in local time mode ( Codes U and L ). In the last case, the time zone information can be provided in the composite C89K.
	TimeMode string `xml:"timeMode,omitempty"`  // minOccurs="0"

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_142129C `xml:"dateTime"`

	// Reference : IATA SSIM Appendix F If it is not provided, the time is considered to be given in UTC.
	TimeZoneInfo *TimeZoneIinformationType `xml:"timeZoneInfo,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeInformationType_94559S struct {

	// This data element is used to provide the semantic of the date information provided.  Examples : - GMT Transaction date - Local Transaction date ... Default being L local date and time
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_142180C `xml:"dateTime"`
}

type StructuredDateTimeInformationType_94567S struct {

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_142188C `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType struct {

	// Year number.
	Year string `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month string `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day string `xml:"day"`

	// Hour between 0 and 23
	Hour string `xml:"hour,omitempty"`  // minOccurs="0"

	// Minutes between 0 and 59
	Minutes string `xml:"minutes,omitempty"`  // minOccurs="0"

	// Seconds between 0 and 59
	Seconds *int32 `xml:"seconds,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_142129C struct {

	// Year number.
	Year int32 `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month int32 `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day int32 `xml:"day"`

	// Hour between 0 and 23
	Hour *int32 `xml:"hour,omitempty"`  // minOccurs="0"

	// Minutes between 0 and 59
	Minutes *int32 `xml:"minutes,omitempty"`  // minOccurs="0"

	// Seconds between 0 and 59
	Seconds *int32 `xml:"seconds,omitempty"`  // minOccurs="0"

	// Milliseconds between 0 and 999.
	Milliseconds *int32 `xml:"milliseconds,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_142180C struct {

	// Year number.
	Year string `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month string `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day string `xml:"day"`

	// Hour between 0 and 23
	Hour string `xml:"hour,omitempty"`  // minOccurs="0"

	// Minutes between 0 and 59
	Minutes string `xml:"minutes,omitempty"`  // minOccurs="0"

	// Seconds between 0 and 59
	Seconds *int32 `xml:"seconds,omitempty"`  // minOccurs="0"

	// Milliseconds between 0 and 999.
	Milliseconds *int32 `xml:"milliseconds,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_142188C struct {

	// Year number.
	Year *int32 `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month *int32 `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day *int32 `xml:"day,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_16347C struct {

	// Year number.
	Year int32 `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month string `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day string `xml:"day"`
}

type StructuredDateTimeType_17997C struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year string `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month string `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day string `xml:"day,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_18725C struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year *int32 `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month string `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day string `xml:"day,omitempty"`  // minOccurs="0"

	// Hour between 0 and 23
	Hour *int32 `xml:"hour,omitempty"`  // minOccurs="0"

	// Minutes between 0 and 59
	Minutes *int32 `xml:"minutes,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_198200C struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year *int32 `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month *int32 `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day *int32 `xml:"day,omitempty"`  // minOccurs="0"

	// Hour between 0 and 23
	Hour *int32 `xml:"hour,omitempty"`  // minOccurs="0"

	// Minutes between 0 and 59
	Minutes *int32 `xml:"minutes,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_198234C struct {

	// Hour between 0 and 23
	Hour int32 `xml:"hour"`

	// Minutes between 0 and 59
	Minutes int32 `xml:"minutes"`
}

type StructuredDateTimeType_35730C struct {

	// Hour between 0 and 23
	Hour string `xml:"hour"`

	// Minutes between 0 and 59
	Minutes string `xml:"minutes"`
}

type StructuredDateTimeType_36775C struct {

	// Year number.
	Year string `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month string `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day string `xml:"day,omitempty"`  // minOccurs="0"

	// Hour between 0 and 23
	Hour string `xml:"hour,omitempty"`  // minOccurs="0"

	// Minutes between 0 and 59
	Minutes string `xml:"minutes,omitempty"`  // minOccurs="0"

	// Seconds between 0 and 59
	Seconds *int32 `xml:"seconds,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType_36777C struct {

	// Year number.
	Year int32 `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month string `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day string `xml:"day"`
}

type StructuredDateTimeType_44876C struct {

	// Year number.
	Year string `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month string `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day string `xml:"day"`
}

type StructuredDateTimeType_56472C struct {

	// Year number.
	Year string `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month string `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day string `xml:"day"`

	// Hour between 0 and 23
	Hour string `xml:"hour,omitempty"`  // minOccurs="0"

	// Minutes between 0 and 59
	Minutes string `xml:"minutes,omitempty"`  // minOccurs="0"
}

type StructuredPeriodInformationType struct {

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_17997C `xml:"beginDateTime,omitempty"`  // minOccurs="0"

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_17997C `xml:"endDateTime,omitempty"`  // minOccurs="0"
}

type StructuredPeriodInformationType_11026S struct {

	// This data element can be used to provide the semantic of the information provided.
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	// Indicate the time is expressed in local time mode.
	TimeMode string `xml:"timeMode,omitempty"`  // minOccurs="0"

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_18725C `xml:"beginDateTime"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_18725C `xml:"endDateTime"`
}

type StructuredPeriodInformationType_136705S struct {

	// DDT Drop-off Date and Time OCH Opening and Closing hours PDT Pickup Date and Time PKT Early and Late Pickup time RTT Early and Late Return time
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	// Indicate if the time is expressed in UTC or in local time mode ( Codes U and L ). In the last case, the time zone information can be provided in the composite C89K.
	TimeMode string `xml:"timeMode,omitempty"`  // minOccurs="0"

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_198200C `xml:"beginDateTime"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_198200C `xml:"endDateTime"`
}

type StructuredPeriodInformationType_136724S struct {

	// DDT Drop-off Date and Time OCH Opening and Closing hours PDT Pickup Date and Time PKT Early and Late Pickup time RTT Early and Late Return time
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	// Indicate if the time is expressed in UTC or in local time mode ( Codes U and L ). In the last case, the time zone information can be provided in the composite C89K.
	TimeMode string `xml:"timeMode,omitempty"`  // minOccurs="0"

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_198234C `xml:"beginDateTime"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_198234C `xml:"endDateTime"`

	// It is used with a period to give a restriction for days impacted. It permits for example to indicate on which days, a flight operates.
	Frequency *FrequencyType `xml:"frequency"`
}

type StructuredPeriodInformationType_8955S struct {

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_16347C `xml:"beginDateTime"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_16347C `xml:"endDateTime"`
}

type StructuredTelephoneNumberType struct {

	// Telephone number
	TelephoneNumber string `xml:"telephoneNumber,omitempty"`  // minOccurs="0"
}

type StructuredTelephoneNumberType_198214C struct {

	// Telephone number
	TelephoneNumber string `xml:"telephoneNumber"`
}

type StructuredTelephoneNumberType_48448C struct {

	// International dial code
	InternationalDialCode string `xml:"internationalDialCode,omitempty"`  // minOccurs="0"

	// Local prefix code
	LocalPrefixCode string `xml:"localPrefixCode,omitempty"`  // minOccurs="0"

	// Area code
	AreaCode string `xml:"areaCode,omitempty"`  // minOccurs="0"

	// Telephone number
	TelephoneNumber string `xml:"telephoneNumber"`
}

type SystemDetailsInfoType struct {

	// BCS distribution channel
	CascadingSystem *SystemDetailsTypeI_46415C `xml:"cascadingSystem"`
}

type SystemDetailsInfoType_33158S struct {

	// POS airline
	DeliveringSystem *SystemDetailsTypeI_57708C `xml:"deliveringSystem"`
}

type SystemDetailsInfoType_94569S struct {

	// LNIATA of the agent.
	WorkstationId string `xml:"workstationId,omitempty"`  // minOccurs="0"

	// System delivering the shopper session ID
	DeliveringSystem *SystemDetailsTypeI `xml:"deliveringSystem,omitempty"`  // minOccurs="0"
}

type SystemDetailsTypeI struct {

	// will convey the name of the company ex: OPODO
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"
}

type SystemDetailsTypeI_46415C struct {

	// contains the distribution channel data. It is the concatenation of "DCD" + [access type] + [product] + [sub-product] access type, product and sub-product are represented on 3 chars.
	CompanyId string `xml:"companyId"`
}

type SystemDetailsTypeI_57708C struct {

	// Corporate Code
	CompanyId string `xml:"companyId"`

	// Pseudo City Code
	LocationId string `xml:"locationId,omitempty"`  // minOccurs="0"
}

type TariffInformationDetailsTypeI struct {

	// A unique rate product identifier.
	RateType string `xml:"rateType,omitempty"`  // minOccurs="0"

	// This field is used to convey the amount.
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// This field is used to convey the currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// This data element is used to convey the rate plan (Daily or total indicator).
	RatePlanIndicator string `xml:"ratePlanIndicator,omitempty"`  // minOccurs="0"

	// This data element is used to convey the rate amount type.
	AmountType string `xml:"amountType,omitempty"`  // minOccurs="0"

	// This data element is used to specify the fact that a rate change occurs during the period of the stay. If the is a change the value is * (for YES)
	RateChangeIndicator string `xml:"rateChangeIndicator,omitempty"`  // minOccurs="0"
}

type TariffInformationDetailsTypeI_198216C struct {

	// CNV for converted Rate
	RateChangeIndicator string `xml:"rateChangeIndicator"`
}

type TariffInformationDetailsTypeI_39533C struct {

	// Net premium
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// currency of the total price and net premium i.e in EUR/ USD
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// general indicator
	AmountType string `xml:"amountType,omitempty"`  // minOccurs="0"

	// Total amount of the insurance element.
	TotalAmount *float64 `xml:"totalAmount,omitempty"`  // minOccurs="0"
}

type TariffInformationDetailsTypeI_50731C struct {

	// Fare Basis Code 
	FareBasisCode string `xml:"fareBasisCode,omitempty"`  // minOccurs="0"

	// Fare Base amount
	FareBaseAmount *float64 `xml:"fareBaseAmount,omitempty"`  // minOccurs="0"

	// This field is used to convey the currency
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"
}

type TariffInformationDetailsTypeU struct {

	// Prive value. The value conveyed equals 100 times the original value in order to avoid transporting decimal placement information.
	PriceAmount float64 `xml:"priceAmount"`

	// Currency code.
	CurrencyCode string `xml:"currencyCode"`

	// Gives the type of amount.
	PriceQualifier string `xml:"priceQualifier"`
}

type TariffInformationDetailsTypeU_127523C struct {

	// Prive value. The value conveyed equals 100 times the original value in order to avoid transporting decimal placement information.
	PriceAmount float64 `xml:"priceAmount"`

	// Gives the type of amount.
	PriceQualifier string `xml:"priceQualifier"`
}

type TariffInformationDetailsTypeU_45479C struct {

	// Prive value. The value conveyed equals 100 times the original value in order to avoid transporting decimal placement information.
	PriceAmount *float64 `xml:"priceAmount,omitempty"`  // minOccurs="0"

	// Conveys the currency code.
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// This qualifier specifies how the price information is to be used. PPI = price per item. There is a price value. INC = Inclusive. There is no price value. This inclusive qualifier specifies that the price to which it applies is already accounted for in the price for another item.
	PriceQualifier string `xml:"priceQualifier"`
}

type TariffInformationDetailsTypeU_46314C struct {

	// A unique rate product identifier:  - PRODUCT = tariff for a product  - TOUR    = price of the tour  - TAXFEE  = tax or fee  - REMAIN  = remaining amount
	RateIdentifier string `xml:"rateIdentifier"`

	// unitary amount for the tariff
	UnitaryAmount float64 `xml:"unitaryAmount"`

	// currency code used for the tariff
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// To qualify and get information on the tariff: cancellation charge, tax amount, total amount, no amount of insurance...
	TariffQualifier string `xml:"tariffQualifier"`

	// Total Amount for the tariff, set if quantity is present
	TotalAmount *float64 `xml:"totalAmount,omitempty"`  // minOccurs="0"

	// quantity for the tariff, when tariff is detailed with a quantity x unitaryAmount and totalPrice
	Quantity *int32 `xml:"quantity,omitempty"`  // minOccurs="0"

	// value is codeset 65 if amount is negative. For Tour, the remaining amount to pay can be negative if the price of the tour has changed
	TariffStatus string `xml:"tariffStatus,omitempty"`  // minOccurs="0"
}

type TariffInformationType struct {

	// This composite gives details about the monetary amounts and their usage.
	PriceDetails *TariffInformationDetailsTypeU `xml:"priceDetails"`
}

type TariffInformationTypeI struct {

	// This composite is used to convey the tariff information
	TariffInfo *TariffInformationDetailsTypeI `xml:"tariffInfo"`

	// Additional rate type information
	RateInformation *RateInformationTypeI `xml:"rateInformation,omitempty"`  // minOccurs="0"

	// This composite is used to convey all the extra charge information.
	ChargeDetails []*AssociatedChargesInformationTypeI `xml:"chargeDetails,omitempty"`  // minOccurs="0" maxOccurs="30"
}

type TariffInformationTypeI_136706S struct {

	// This composite is used to convey the tariff information
	TariffInfo *TariffInformationDetailsTypeI `xml:"tariffInfo"`

	// Additional rate type information
	RateInformation *RateInformationTypeI_198204C `xml:"rateInformation,omitempty"`  // minOccurs="0"

	// This composite is used to convey all the extra charge information.
	ChargeDetails []*AssociatedChargesInformationTypeI_198205C `xml:"chargeDetails,omitempty"`  // minOccurs="0" maxOccurs="30"
}

type TariffInformationTypeI_136714S struct {

	// tariff conversion indicator associated to the charge
	TariffInfo *TariffInformationDetailsTypeI_198216C `xml:"tariffInfo,omitempty"`  // minOccurs="0"

	// Tax, Surcharge, Coverage, Coupon details  Repetition are used to carry: - General Info (estimated + name...) - Tariff per day / Max - Tariff per weekend / Max - Tariff per week / Max - Tariff per month / Max - Tariff per rental / Max - Excess amount - Liability amount
	ChargeDetails []*AssociatedChargesInformationTypeI_198218C `xml:"chargeDetails"`  // maxOccurs="16"
}

type TariffInformationTypeI_136719S struct {

	// tariff conversion indicator associated to the charge
	TariffInfo *TariffInformationDetailsTypeI_198216C `xml:"tariffInfo,omitempty"`  // minOccurs="0"

	// Tax, Surcharge, Coverage, Coupon details
	ChargeDetails []*AssociatedChargesInformationTypeI_198218C `xml:"chargeDetails"`  // maxOccurs="4"
}

type TariffInformationTypeI_22057S struct {

	// total price and net premium
	TariffInfo *TariffInformationDetailsTypeI_39533C `xml:"tariffInfo,omitempty"`  // minOccurs="0"

	// to specify the taxes and their values and converted values into other currencies if specified.
	ChargeDetails []*AssociatedChargesInformationTypeI_39535C `xml:"chargeDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type TariffInformationTypeI_28460S struct {

	// This composite is used to convey the tariff information
	TariffInfo *TariffInformationDetailsTypeI_50731C `xml:"tariffInfo,omitempty"`  // minOccurs="0"

	// Additional rate type information
	RateInformation *RateInformationTypeI_50732C `xml:"rateInformation,omitempty"`  // minOccurs="0"
}

type TariffInformationTypeU struct {

	// This composite provides details for a price.
	PriceDetails *TariffInformationDetailsTypeU_45479C `xml:"priceDetails"`
}

type TariffInformationTypeU_25419S struct {

	// Information about the tariffs of a Tour: tour price, product price, additional price
	TariffInformation []*TariffInformationDetailsTypeU_46314C `xml:"tariffInformation"`  // maxOccurs="3"

	// This composite is used to describe the commissions on the tariff
	AssociatedChargesInformation *AssociatedChargesInformationTypeU `xml:"associatedChargesInformation,omitempty"`  // minOccurs="0"
}

type TariffInformationType_83558S struct {

	// This composite gives details about the monetary amounts and their usage.
	PriceDetails *TariffInformationDetailsTypeU_127523C `xml:"priceDetails"`
}

type TariffcodeType struct {

	// To convey the Tariff code.
	TariffCode string `xml:"tariffCode,omitempty"`  // minOccurs="0"

	// to convey a description of the type of tariff.
	TariffCodeType string `xml:"tariffCodeType,omitempty"`  // minOccurs="0"
}

type TaxDetailsTypeI struct {

	// Tax Amount
	TaxRate string `xml:"taxRate,omitempty"`  // minOccurs="0"

	// See ISO 4217 codes
	CurrCode string `xml:"currCode,omitempty"`  // minOccurs="0"

	// Type of the tax
	TaxType []string `xml:"taxType,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type TaxDetailsTypeU struct {

	// Tax qualifier. For Ferry, only one tax qualifier applies: Port taxes.
	Qualifier string `xml:"qualifier"`

	// Tax amount.
	Amount float64 `xml:"amount"`
}

type TaxFieldsType struct {

	// Tax indicator
	TaxIndicator string `xml:"taxIndicator"`

	// Tax currency
	TaxCurrency string `xml:"taxCurrency"`

	// Tax amount
	TaxAmount string `xml:"taxAmount"`

	// Tax country code
	TaxCountryCode string `xml:"taxCountryCode"`

	// Tax nature code
	TaxNatureCode string `xml:"taxNatureCode,omitempty"`  // minOccurs="0"
}

type TaxTypeI struct {

	// specify the tax details
	TaxDetails []*TaxDetailsTypeI `xml:"taxDetails,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type TaxesType struct {

	// Tax details description.
	AdditionnalCharge *TaxDetailsTypeU `xml:"additionnalCharge"`
}

type TerminalIdentificationDescriptionType struct {

	// Identification of the transaction initiator.
	TerminalID string `xml:"terminalID,omitempty"`  // minOccurs="0"

	// The distribution channel.
	DistributionChannel *DistributionChannelType `xml:"distributionChannel"`
}

type TerminalInformationTypeU struct {

	// Arrival Terminal
	ArrivalTerminal string `xml:"arrivalTerminal"`
}

type TerminalTimeInformationTypeS struct {

	// LOCATION IDENTIFICATION
	LocationDetails *LocationIdentificationTypeS `xml:"locationDetails,omitempty"`  // minOccurs="0"
}

type ThreeDomainSecureGroupType struct {

	// This segment conveys a set of data resulting from the 3DS authentication process
	AuthenticationData *CreditCardSecurityType `xml:"authenticationData"`

	// Access Control Server's URL (up to 2048 characters).
	AcsURL *CommunicationContactType `xml:"acsURL,omitempty"`  // minOccurs="0"

	// will convey the various messages/encrypted data used during the 3DS authentication processes
	TdsBlobData []*TdsBlobData `xml:"tdsBlobData,omitempty"`  // minOccurs="0" maxOccurs="4"
}

type TdsBlobData struct {

	// will identify the content of the BLB that follows
	TdsBlbIdentifier *ReferenceInfoType_94524S `xml:"tdsBlbIdentifier"`

	TdsBlbData *BinaryDataType `xml:"tdsBlbData"`
}

type TicketElementType struct {

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType string `xml:"passengerType,omitempty"`  // minOccurs="0"

	// Ticket information
	Ticket *TicketInformationType `xml:"ticket"`

	// Print options (//print options after double slash)
	PrintOptions string `xml:"printOptions,omitempty"`  // minOccurs="0"
}

type TicketInformationType struct {

	// Ticketing type  TL, OK, DO, IN, MA, TR, AT, PT, XL, ST, SS
	Indicator string `xml:"indicator"`

	// Ticketing date
	Date *int32 `xml:"date,omitempty"`  // minOccurs="0"

	// Ticketing time
	Time string `xml:"time,omitempty"`  // minOccurs="0"

	// Office Id
	OfficeId string `xml:"officeId,omitempty"`  // minOccurs="0"

	// Free flow text
	Freetext string `xml:"freetext,omitempty"`  // minOccurs="0"

	// Air France flag (e.g. //TELEPAYE for MINITEL)
	TransactionFlag string `xml:"transactionFlag,omitempty"`  // minOccurs="0"

	// Electronic ticketing flag + airline code (e.g. //ETLH)  ET for Electronic ticket candidate
	ElectronicTicketFlag string `xml:"electronicTicketFlag,omitempty"`  // minOccurs="0"

	// Airline code
	AirlineCode string `xml:"airlineCode,omitempty"`  // minOccurs="0"

	// Queue number
	QueueNumber string `xml:"queueNumber,omitempty"`  // minOccurs="0"

	// Category number
	QueueCategory string `xml:"queueCategory,omitempty"`  // minOccurs="0"

	// SITA addresses
	SitaAddress []string `xml:"sitaAddress,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type TicketInformationType_5120C struct {

	// 1. PNR Header: Amadeus Queuing Office Id
	QueueingOfficeId string `xml:"queueingOfficeId,omitempty"`  // minOccurs="0"

	// 1. PNR Header: OA city code
	Location string `xml:"location,omitempty"`  // minOccurs="0"
}

type TicketNumberDetailsTypeI struct {

	// eVoucher number
	Number string `xml:"number"`
}

type TicketNumberTypeI struct {

	// documentDetails
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`
}

type TicketingFormOfPaymentType struct {

	// Form of payment details
	FopDetails []*FormOfPaymentInformationType `xml:"fopDetails,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type TimeZoneIinformationType struct {

	// ISO country Code See SSIM appendix F
	CountryCode string `xml:"countryCode"`

	// Time zone code. See SSIM appendix F.
	Code *int32 `xml:"code,omitempty"`  // minOccurs="0"

	// Time zone suffix to complete the time zone code when necessary. See SSIM appendix F.
	Suffix string `xml:"suffix,omitempty"`  // minOccurs="0"
}

type TotalPriceType struct {

	// The provider code.
	ProviderCode *CompanyInformationType_83550S `xml:"providerCode"`

	// External Reference of the pricing
	ExternalRef *ReferenceInformationTypeI_83551S `xml:"externalRef,omitempty"`  // minOccurs="0"

	// method of delivery, such as e-mail, pick at station, id card, etc.
	MethodOfDelivery *MethodOfDelivery `xml:"methodOfDelivery,omitempty"`  // minOccurs="0"

	// This segment is used to convey the main price information (e.g. the net total price for non-cancelled bookings, the cancellation fee for cancelled bookings).  The currency code stands not only for this segment, but for the whole group: all prices have the same currency.
	MainPrice *TariffInformationType `xml:"mainPrice"`

	// The remaining price items are described here. In ferry business, there may be a maximum of 12 prices (+ main price and taxes).  The currency code is not applicable because it is the same as in the mainPriceInformation segment.
	OtherPrices []*TariffInformationType_83558S `xml:"otherPrices,omitempty"`  // minOccurs="0" maxOccurs="12"

	// product associated to the price item
	ProductDescription *ProductDescription `xml:"productDescription,omitempty"`  // minOccurs="0"

	// This segment conveys the tax amount information. The repetition factor equals the number of codesets for the qualifier, because each type of tax may occur once. The currency code is the same as in the mainPriceInformation segment.
	AdditionnalChargeInformation []*TaxesType `xml:"additionnalChargeInformation,omitempty"`  // minOccurs="0" maxOccurs="4"

	// This segment is used to convey the booking fare information.
	RateCodeInformation *RateTypesTypeU `xml:"rateCodeInformation,omitempty"`  // minOccurs="0"

	// This segment will transport the optional booking confirmation dead-line information. Note: if this segment transports a valid confirmation dead-line, then the booking is considered as optional.
	OptionalBooking *StructuredDateTimeInformationType_83553S `xml:"optionalBooking,omitempty"`  // minOccurs="0"
}

type MethodOfDelivery struct {

	// Identification and semantic attached to the reference description (E.g: a customer can have multiple roles: payer, traveller, insured...)
	ElementManagement *ElementManagementSegmentType_83559S `xml:"elementManagement"`

	// Describes the details around this mode of delivery
	DeliveryDetails *PackageDescriptionType `xml:"deliveryDetails"`
}

type ProductDescription struct {

	// product associated to the price item
	Product *ProductIdentificationTypeU `xml:"product"`

	// product restrictions and attributes: route code and description, crossLondon and advanced purchase.
	ProductRestriction []*TrafficRestrictionDetailsType `xml:"productRestriction,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type TourAccountDetailsType struct {

	// Total price of the Tours. The segment can be repeated in case the total price is written in more than one currency. This trigger is M20 and not M1, but there is no grammar problem with that. There would be a problem if the group TURP was repeated, but this is not and shall never be the case.
	TourTotalPrices []*TariffInformationTypeU_25419S `xml:"tourTotalPrices"`  // maxOccurs="20"

	// The remaining amount to pay, for each Tour Operator
	RemainingAmountsDetails []*RemainingAmountsDetails `xml:"remainingAmountsDetails,omitempty"`  // minOccurs="0" maxOccurs="10"

	// All tour products (accomodation, transport, insurance...) prices and additional price elements such as airport tax and fees. We can have 99 descriptions of prices about products. And we can have 2 additional prices' descriptions, about other fees and taxes not linked to a product. This gives a max of 101 prices' description.
	TourDetailedPriceInfo []*TourDetailedPriceInfo `xml:"tourDetailedPriceInfo,omitempty"`  // minOccurs="0" maxOccurs="101"

	// Payments done by Tour Operator
	PaymentInformation []*PaymentInformation `xml:"paymentInformation,omitempty"`  // minOccurs="0" maxOccurs="10"
}

type PaymentInformation struct {

	// Deposit details: amount, currency, date, purpose and payment method information
	Payment *PaymentInformationTypeU `xml:"payment"`

	// The Tour Operator code
	OperatorCode *CompanyInformationType_25420S `xml:"operatorCode,omitempty"`  // minOccurs="0"
}

type RemainingAmountsDetails struct {

	// The Tour Operator code
	ProviderCode *CompanyInformationType_25420S `xml:"providerCode"`

	// The remaining amount to pay
	RemainingAmount *TariffInformationTypeU_25419S `xml:"remainingAmount"`
}

type TourDetailedPriceInfo struct {

	// dummy segment
	MarkerSpecificRead *DummySegmentTypeI `xml:"markerSpecificRead"`

	// Identifier of the product
	ProductId *ReferenceInfoType_25422S `xml:"productId,omitempty"`  // minOccurs="0"

	// Price of a product or tax or fee to pay for a product
	ProductPrice *TariffInformationTypeU_25419S `xml:"productPrice"`
}

type TourDetailsTypeI struct {

	// Inclusive tour number
	TourCode string `xml:"tourCode"`
}

type TourInformationType struct {

	// Conveys summary information of the booking (such as departure/arrival location or date of the booking, the booking type ...). The providerName data of the composite E988 DOES NOT conveys the Tour Operator's name. Providers' name and code are stored in CPYs.
	BookingSummaryInfo *TravelProductInformationTypeU_25428S `xml:"bookingSummaryInfo"`

	// Conveys information about the booking duration
	BookingDurationInfo *QuantityType `xml:"bookingDurationInfo,omitempty"`  // minOccurs="0"

	// Conveys staying location information
	StayingInfo *PlaceLocationIdentificationTypeU_25436S `xml:"stayingInfo,omitempty"`  // minOccurs="0"

	// Conveys the tour description (name and description)
	TourDescriptionInfo *AdditionalProductDetailsTypeU `xml:"tourDescriptionInfo,omitempty"`  // minOccurs="0"

	// Conveys booking reference and unique key in the provider system
	BookingReferenceInfo []*ReservationControlInformationTypeU `xml:"bookingReferenceInfo"`  // maxOccurs="2"

	// Conveys the status of the booking or of the ticket and the number in party. The composite E958 is M2:  -One instance can be the booking's status  -The other can be the TKOK status
	StatusInfo *QuantityAndActionTypeU `xml:"statusInfo"`

	// Indicates whether an insurance is included in the tour booking.
	InsuranceIndication *InsuranceCoverageType_25483S `xml:"insuranceIndication,omitempty"`  // minOccurs="0"

	// Conveys passenger information when there is a desynchronization between the PNR passengers and tour passengers (Tour Server). A Tour Server booking can contain its own passenger names. Tour Server's specifications specify a maximum of 99 passengers. If a Tour booking conveys its passengers in TIF, then there is no Pax assoc (no REF used for Pax assoc). BUT, when there is no TIF (for example in Tour Distribution), there may be a Pax assoc between product and PNR pax, and/or booking and PNR's pax. Then REFs are used
	PassengerInfo []*TravellerInformationType_25441S `xml:"passengerInfo,omitempty"`  // minOccurs="0" maxOccurs="99"

	// Conveys the booking expiration information (only the expiration date is needed).
	ExpireInfo *StructuredDateTimeInformationType_25444S `xml:"expireInfo,omitempty"`  // minOccurs="0"

	// Conveys description information (Remark booking description). For the remarks we can have 2 lines of text that is why the 101C composite is repeted 2 times.
	BookingDescriptionInfo []*FreeTextInformationType_25445S `xml:"bookingDescriptionInfo,omitempty"`  // minOccurs="0" maxOccurs="4"

	// Conveys information about the targeted system provider (TS for Tour source, TG for Royal Orchid Holiday ...). The only information conveyd is the code of the provider whose system has the master Tour booking. Example: the Tour Operator ROH provides its products throw the TG (Thai Airways) system: TRA conveys TG, one CPY conveys ROH.
	SystemProviderInfo *TransportIdentifierType `xml:"systemProviderInfo,omitempty"`  // minOccurs="0"

	// Conveys information about the tour operator (name, code ...)
	TourOperatorInfo []*CompanyInformationType_25420S `xml:"tourOperatorInfo"`  // maxOccurs="10"

	// Bokking source (/BS)
	BookingSource *UserIdentificationType_25447S `xml:"bookingSource,omitempty"`  // minOccurs="0"

	// Conveys the passenger association for the booking
	PassengerAssocation []*ReferenceInfoType_25422S `xml:"passengerAssocation,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Conveys tour payment information such as the detailed price of the booking, the commisssion, the deposit information ...
	TourAccountDetails *TourAccountDetailsType `xml:"tourAccountDetails,omitempty"`  // minOccurs="0"

	// Conveys information about the booked products (arrival/departure information, product identification, meal plan information, occupation ...)
	TourProductDetails []*TourServiceDetailsType `xml:"tourProductDetails"`  // maxOccurs="99"
}

type TourInformationTypeI struct {

	// Tour code
	TourInformationDetails *TourDetailsTypeI `xml:"tourInformationDetails"`
}

type TourServiceDetailsType struct {

	// Conveys the product sequence number which is the product place in the booking. This information locally identifies the product in the Tour booking.
	SequenceNumberInfo *ItemNumberTypeU `xml:"sequenceNumberInfo"`

	// Conveys information about the product status and the product quantity (number in party or number of service)
	StatusQuantityInfo *QuantityAndActionTypeU `xml:"statusQuantityInfo,omitempty"`  // minOccurs="0"

	// Conveys general Tour product information.
	ProductInfo *AdditionalProductDetailsTypeU `xml:"productInfo"`

	// Conveys product confirmation number.
	ConfirmationInfo *ReservationControlInformationTypeU `xml:"confirmationInfo,omitempty"`  // minOccurs="0"

	// Passenger association at product (package / standalone) level.
	PassengerAssociation []*ReferenceInfoType_25422S `xml:"passengerAssociation,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Conveys the service details which composes a tour product. If the product is a package this group can be repeted.
	ServiceDetails []*ServiceDetails `xml:"serviceDetails"`  // maxOccurs="99"
}

type ServiceDetails struct {

	// Conveys general service information such as departure/arrival information, service code, service description, service type ...
	ServiceInfo *TravelProductInformationTypeU_25428S `xml:"serviceInfo"`

	// Conveys duration information (number of day, night ...)
	ServiceDurationInfo *QuantityType `xml:"serviceDurationInfo,omitempty"`  // minOccurs="0"

	// Conveys information about booked rooms
	AccomodationDetails []*AccomodationDetails `xml:"accomodationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"

	// Conveys vehicule details
	VehiculeDetails *VehiculeDetails `xml:"vehiculeDetails,omitempty"`  // minOccurs="0"

	// Conveys transportation details. We store in the repetitions the legs (or connections) or this transportation. Tour Server specifications specify a max of 99 legs for cruises and flight.
	TransportationDetails []*TransportationDetails `xml:"transportationDetails,omitempty"`  // minOccurs="0" maxOccurs="99"

	// Billing Collection Statistic at Tour Product level
	ProductBCSDetails *ProductBCSDetails `xml:"productBCSDetails,omitempty"`  // minOccurs="0"
}

type TrafficRestrictionDetailsType struct {

	// restriction details
	RestrictionDetails *TrafficRestrictionDetailsTypeU `xml:"restrictionDetails"`
}

type TrafficRestrictionDetailsTypeU struct {

	// restriction code.
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// restriction type
	Type string `xml:"type"`

	// traffic restriction description
	Description string `xml:"description,omitempty"`  // minOccurs="0"
}

type TrainDataType struct {

	// Information pertaining to the train product
	TrainProductInfo *TrainProductInformationType `xml:"trainProductInfo"`

	// Trip dates and times
	TripDateTime []*StructuredDateTimeInformationType_32362S `xml:"tripDateTime"`  // maxOccurs="2"

	// Departure station location
	DepLocation *PlaceLocationIdentificationTypeU_32347S `xml:"depLocation"`

	// Arrival station location
	ArrLocation *PlaceLocationIdentificationTypeU_32347S `xml:"arrLocation"`

	// Rail leg (train number, train provider, departure/arrival locations and dates, reservable status)
	RailLeg []*RailLegDataType `xml:"railLeg,omitempty"`  // minOccurs="0" maxOccurs="6"
}

type TrainDetailsType struct {

	// Train company code
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Train number
	Number string `xml:"number,omitempty"`  // minOccurs="0"
}

type TrainInformationType struct {

	// Information pertaining to the rail company
	CompanyInfo *CompanyInformationType_19450S `xml:"companyInfo"`

	// Indicates whether or not the reservation can be modified directly in the Amadeus system
	UpdatePermission *StatusTypeI_20923S `xml:"updatePermission,omitempty"`  // minOccurs="0"

	// train number, equipment code, departure and arrival dates and times.
	TripDetails *TrainDataType `xml:"tripDetails"`

	// indicate that the train segment is open.
	OpenSegment *StatusTypeI_33257S `xml:"openSegment,omitempty"`  // minOccurs="0"

	// Journey direction: outward, return, single
	JourneyDirection *TravelItineraryInformationTypeI_33275S `xml:"journeyDirection,omitempty"`  // minOccurs="0"

	// Rail provider segment tattoo reference
	ProviderTattoo *ItemReferencesAndVersionsType `xml:"providerTattoo,omitempty"`  // minOccurs="0"

	// SVC / Service information
	ServiceInfo *FreeTextInformationType_20551S `xml:"serviceInfo,omitempty"`  // minOccurs="0"

	// Information pertaining to the class of service including number of seats
	ClassInfo *ClassConfigurationDetailsType `xml:"classInfo"`

	// Accommodation (room/compartment) details.
	AccommodationInfo *AccommodationAllocationInformationTypeU `xml:"accommodationInfo,omitempty"`  // minOccurs="0"

	// Coach information
	CoachInfo *CoachProductInformationType `xml:"coachInfo,omitempty"`  // minOccurs="0"

	// Reservation Mandatory, Advised, Possible, Not Possible
	ReservableStatus *QuantityAndActionTypeU_32609S `xml:"reservableStatus,omitempty"`  // minOccurs="0"
}

type TrainProductInformationType struct {

	// Train Details
	TrainDetails *TrainDetailsType `xml:"trainDetails,omitempty"`  // minOccurs="0"

	// Transportation mode (BUS, SHIP, TRAIN, TGV etc)
	Type string `xml:"type,omitempty"`  // minOccurs="0"
}

type TrainProductInformationType_32331S struct {

	// Rail Company
	RailCompany string `xml:"railCompany"`

	// Train Details
	TrainDetails *TrainDetailsType `xml:"trainDetails,omitempty"`  // minOccurs="0"

	// Train Equipment Type  (TGV,TGD,TGN...)
	Type string `xml:"type"`
}

type TransactionInformationForTicketingType struct {

	// Authorisation transaction details
	TransactionDetails *TransactionInformationsType `xml:"transactionDetails"`
}

type TransactionInformationsType struct {

	// Authorization message type  Eg 110: author according standard ISO8583 210: settlement according standard ISO858 ...
	Code string `xml:"code,omitempty"`  // minOccurs="0"

	// Credit Card link used to perform authorization.
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Process indicator (bulkIndicator): - bulk - superbulk - no bulk.
	IssueIndicator string `xml:"issueIndicator,omitempty"`  // minOccurs="0"

	// This is a message number that uniquely identifies a cardholder transaction.  According to the link this info can have various names:  - STAN number(Systems Trace Audit Number) - ISO8583 (VISA,Nedbank, Credit Mutuel...)  - Message number - APACS70 (Barclays,Euroline...)  ...   Official definition: This is a number assigned by the message initiator that uniquely identifies a cardholder transaction and all the message types (also known as system transactions) that it comprises, according to individual program rules. The trace number remains unchanged for all messages throughout the life of the transaction. For example, the same trace number is used in an authorization request and response, and in a subsequent reversal request and response, and in any advices of authorization or reversal.
	TransmissionControlNumber string `xml:"transmissionControlNumber,omitempty"`  // minOccurs="0"
}

type TransportIdentifierType struct {

	// Targeted provider system information
	CompanyIdentification *CompanyIdentificationTypeI_46351C `xml:"companyIdentification"`
}

type TravelItineraryInformationTypeI struct {

	// The sequence number indentifying the position of a leg in a booking
	ItemNumber int32 `xml:"itemNumber"`
}

type TravelItineraryInformationTypeI_33275S struct {

	// direction of travel indicator (outward, return, single)
	MovementType string `xml:"movementType"`
}

type TravelProductInformationTypeI struct {

	// To specify dates and times of the product
	Product *ProductDateTimeTypeI_171495C `xml:"product,omitempty"`  // minOccurs="0"

	// Boarding point detail
	BoardpointDetail *LocationTypeI_2784C `xml:"boardpointDetail,omitempty"`  // minOccurs="0"

	// Off Point details
	OffpointDetail *LocationTypeI_2784C `xml:"offpointDetail,omitempty"`  // minOccurs="0"

	// Company identification
	CompanyDetail *CompanyIdentificationTypeI_2785C `xml:"companyDetail,omitempty"`  // minOccurs="0"

	// Product identifications details
	ProductDetails *ProductIdentificationDetailsTypeI_2786C `xml:"productDetails,omitempty"`  // minOccurs="0"

	// Product Type details
	TypeDetail *ProductTypeDetailsTypeI `xml:"typeDetail,omitempty"`  // minOccurs="0"

	// AIR segment : to indicate an Informational Air segment :  N for No action required.
	ProcessingIndicator string `xml:"processingIndicator,omitempty"`  // minOccurs="0"
}

type TravelProductInformationTypeI_127288S struct {
}

type TravelProductInformationTypeI_186189S struct {

	// To specify dates and times of the product
	Product *ProductDateTimeTypeI_260882C `xml:"product,omitempty"`  // minOccurs="0"

	// Boarding point detail
	BoardpointDetail *LocationTypeI_2784C `xml:"boardpointDetail,omitempty"`  // minOccurs="0"

	// Off Point details
	OffpointDetail *LocationTypeI_2784C `xml:"offpointDetail,omitempty"`  // minOccurs="0"

	// Company identification
	CompanyDetail *CompanyIdentificationTypeI_2785C `xml:"companyDetail,omitempty"`  // minOccurs="0"

	// Product identifications details
	ProductDetails *ProductIdentificationDetailsTypeI_2786C `xml:"productDetails,omitempty"`  // minOccurs="0"

	// Product Type details
	TypeDetail *ProductTypeDetailsTypeI `xml:"typeDetail,omitempty"`  // minOccurs="0"

	// AIR segment : to indicate an Informational Air segment :  N for No action required.
	ProcessingIndicator string `xml:"processingIndicator,omitempty"`  // minOccurs="0"
}

type TravelProductInformationTypeI_25434S struct {

	// flight date information
	FlightDate *ProductDateTimeTypeI_46338C `xml:"flightDate"`

	// carrier details
	CompanyDetails *CompanyIdentificationTypeI_46335C `xml:"companyDetails"`

	// flight information
	FlightIdentification *ProductIdentificationDetailsTypeI_46336C `xml:"flightIdentification"`

	// connection sequence information
	FlightTypeDetails *ProductTypeDetailsTypeI_46337C `xml:"flightTypeDetails,omitempty"`  // minOccurs="0"
}

type TravelProductInformationTypeI_99362S struct {

	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`  // minOccurs="0"

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`  // minOccurs="0"

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`  // minOccurs="0"

	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`  // minOccurs="0"
}

type TravelProductInformationTypeU struct {

	// Conveys the departure and arrival date time descriptions. If absent, then the leg status may be considered as open information.
	ItineraryDateTimeInfo *ProductDateAndTimeTypeU `xml:"itineraryDateTimeInfo,omitempty"`  // minOccurs="0"

	// Conveys and itinerary leg embarkation and the disembarkation ports descriptions
	BoardPortDetails []*LocationTypeU `xml:"boardPortDetails"`  // maxOccurs="2"

	// Internal reference for the leg.
	LineNumber string `xml:"lineNumber"`
}

type TravelProductInformationTypeU_25428S struct {

	// Conveys information about the departure/ arrival date and time.
	DateTimeInformation *ProductDateAndTimeTypeU_46325C `xml:"dateTimeInformation,omitempty"`  // minOccurs="0"

	// Conveys the departure/arrival/staying location information
	LocationInformation []*LocationTypeU_46324C `xml:"locationInformation,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Conveys information about the provider of the product
	CompanyInformation *CompanyIdentificationTypeU `xml:"companyInformation,omitempty"`  // minOccurs="0"

	// Conveys details about the product
	ProductDetails *ProductIdentificationDetailsTypeU_46327C `xml:"productDetails"`
}

type TravelerPerpaxDetailsType struct {

	// Format limitations: an2
	PerpaxMask string `xml:"perpaxMask"`

	// perpax mask indicator (optional/mandatory)
	PerpaxMaskIndicator string `xml:"perpaxMaskIndicator"`
}

type TravellerDetailsType struct {

	// passenger first name
	GivenName string `xml:"givenName,omitempty"`  // minOccurs="0"
}

type TravellerDetailsTypeI struct {

	// Traveler First Name
	FirstName string `xml:"firstName,omitempty"`  // minOccurs="0"

	// Traveler Type using  Amadeus codification.
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// 1 code is used to mention that the traveler is accompanied by an infant with no seat.
	InfantIndicator string `xml:"infantIndicator,omitempty"`  // minOccurs="0"

	// Identification code, 2 cases: ID<1 to 51 char free text) or CR<1 to 40 char free text)
	IdentificationCode string `xml:"identificationCode,omitempty"`  // minOccurs="0"
}

type TravellerDetailsTypeI_16351C struct {

	// Conveys passenger first name.
	GivenName string `xml:"givenName,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Title string `xml:"title,omitempty"`  // minOccurs="0"
}

type TravellerDetailsTypeI_18004C struct {

	// firstname
	GivenName string `xml:"givenName"`
}

type TravellerDetailsTypeI_27968C struct {

	// firstname
	GivenName string `xml:"givenName"`

	// Title (Mr, Mrs)
	Title string `xml:"title,omitempty"`  // minOccurs="0"
}

type TravellerDetailsTypeI_46354C struct {

	// Passenger lastName
	GivenName string `xml:"givenName"`

	// passenger title
	Title string `xml:"title,omitempty"`  // minOccurs="0"
}

type TravellerDocumentInformationTypeU struct {

	DocumentInformation *DocumentInformationTypeU `xml:"documentInformation"`

	DatesOfValidity *ValidityDatesTypeU `xml:"datesOfValidity,omitempty"`  // minOccurs="0"
}

type TravellerInformationType struct {

	// to specify the last name, age and gender.
	PaxDetails *TravellerSurnameInformationType_858C `xml:"paxDetails,omitempty"`  // minOccurs="0"

	// to provide the first name
	OtherPaxDetails *TravellerDetailsTypeI_18004C `xml:"otherPaxDetails,omitempty"`  // minOccurs="0"
}

type TravellerInformationTypeI struct {

	// to specify the last name and the type of person (if it's a nanny or a substitute)
	PaxDetails *TravellerSurnameInformationTypeI_18003C `xml:"paxDetails"`

	// Only used to put the firstname
	OtherPaxDetails *TravellerDetailsTypeI_18004C `xml:"otherPaxDetails,omitempty"`  // minOccurs="0"
}

type TravellerInformationTypeI_15923S struct {

	// to specify the last name and the type of person (if it's a nanny or a substitute)
	PaxDetails *TravellerSurnameInformationTypeI_18003C `xml:"paxDetails"`

	// Other name info
	OtherPaxDetails *TravellerDetailsTypeI_27968C `xml:"otherPaxDetails,omitempty"`  // minOccurs="0"
}

type TravellerInformationTypeI_6097S struct {

	// Traveller surname details
	Traveller *TravellerSurnameInformationTypeI `xml:"traveller,omitempty"`  // minOccurs="0"

	// Occurrence one relates to the traveler.  Occurrence 2 relates only to an infant accompanying the traveler for whom only the given name is present.
	Passenger []*TravellerDetailsTypeI `xml:"passenger,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type TravellerInformationTypeI_8956S struct {

	// Passenger last name details.
	PaxDetails *TravellerSurnameInformationTypeI_16350C `xml:"paxDetails"`

	// Passnger first name details.
	OtherPaxDetails *TravellerDetailsTypeI_16351C `xml:"otherPaxDetails,omitempty"`  // minOccurs="0"
}

type TravellerInformationType_25441S struct {

	// passenger details
	PaxDetails *TravellerSurnameInformationType_46353C `xml:"paxDetails"`

	// other passenger details
	OtherPaxDetails *TravellerDetailsTypeI_46354C `xml:"otherPaxDetails"`
}

type TravellerInformationType_94570S struct {

	// will convey the name of the credit card holder
	PaxDetails *TravellerSurnameInformationType `xml:"paxDetails"`

	// will convey the CC holder first name
	OtherPaxDetails *TravellerDetailsType `xml:"otherPaxDetails,omitempty"`  // minOccurs="0"
}

type TravellerInsuranceInformationType struct {

	// currency of manual premium
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// manual total premium for this traveller
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// supplementary info
	SupplementaryInformation string `xml:"supplementaryInformation,omitempty"`  // minOccurs="0"

	// gender - male or female
	SexCode string `xml:"sexCode,omitempty"`  // minOccurs="0"

	// Credit card details
	CreditCardDetails []*CreditCardType `xml:"creditCardDetails,omitempty"`  // minOccurs="0" maxOccurs="5"

	// currency of the total premium ,
	TotalPremiumCurrency string `xml:"totalPremiumCurrency,omitempty"`  // minOccurs="0"

	// calculated total premium , all taxes included for this traveller
	TotalPremium *float64 `xml:"totalPremium,omitempty"`  // minOccurs="0"

	// for future use
	FutureCurrency string `xml:"futureCurrency,omitempty"`  // minOccurs="0"

	// for future use
	FutureAmount *float64 `xml:"futureAmount,omitempty"`  // minOccurs="0"

	// Reduction Code
	FareType string `xml:"fareType,omitempty"`  // minOccurs="0"

	// Beneficiary Name
	TravelerName string `xml:"travelerName,omitempty"`  // minOccurs="0"
}

type TravellerSurnameInformationType struct {

	// CC holder name details
	Surname string `xml:"surname"`
}

type TravellerSurnameInformationTypeI struct {

	// Traveler Last Name  Group name
	Surname string `xml:"surname"`

	// G for group (The traveler type is in C324/6353)
	Qualifier string `xml:"qualifier,omitempty"`  // minOccurs="0"

	// 1 :one traveler with exceptions below.  2 :traveler accompanied by an infant for whom only the given name is present.  n : total number of passengers of the group (assigned + unassigned)
	Quantity *int32 `xml:"quantity,omitempty"`  // minOccurs="0"

	// Staff type
	StaffType string `xml:"staffType,omitempty"`  // minOccurs="0"
}

type TravellerSurnameInformationTypeI_16350C struct {

	// Conveys passenger last name.
	Surname string `xml:"surname"`
}

type TravellerSurnameInformationTypeI_18003C struct {

	// last name
	Surname string `xml:"surname"`

	// to specify the type of person
	Type string `xml:"type,omitempty"`  // minOccurs="0"
}

type TravellerSurnameInformationType_46353C struct {

	// Passenger name
	Surname string `xml:"surname"`
}

type TravellerSurnameInformationType_858C struct {

	// Passenger name
	Surname string `xml:"surname,omitempty"`  // minOccurs="0"

	// Customer type: A=adult C=child IN = infant
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// to indicate if it's a man or a female.
	Gender string `xml:"gender,omitempty"`  // minOccurs="0"
}

type TravellerTimeDetailsTypeI struct {

	// Time format: 24H. All digits are mandatory . Example: from 0000 to 2359
	CheckinTime string `xml:"checkinTime,omitempty"`  // minOccurs="0"
}

type TstGeneralInformationDetailsType struct {

	// TST reference number
	TstReferenceNumber string `xml:"tstReferenceNumber"`

	// TST creation date
	TstCreationDate string `xml:"tstCreationDate"`

	// Sales indicator
	SalesIndicator string `xml:"salesIndicator,omitempty"`  // minOccurs="0"
}

type TstGeneralInformationType struct {

	// General information
	GeneralInformation *TstGeneralInformationDetailsType `xml:"generalInformation"`
}

type UniqueIdDescriptionType struct {

	// ID sequence number : envelope number
	IDSequenceNumber int32 `xml:"iDSequenceNumber"`

	// ID qualifier: must be 'PNV' as PNR Version Number
	IDQualifier string `xml:"iDQualifier"`
}

type UserIdentificationType struct {

	// Originator Identification Details
	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification"`

	// Agent type (A, T, E)
	OriginatorTypeCode string `xml:"originatorTypeCode,omitempty"`  // minOccurs="0"
}

type UserIdentificationType_127265S struct {

	// 1 character code for airline agent (A), travel agent (T), etc...
	OriginatorTypeCode string `xml:"originatorTypeCode"`
}

type UserIdentificationType_21014S struct {

	// Originator Identification Details
	OriginIdentification *OriginatorIdentificationDetailsTypeI_37406C `xml:"originIdentification"`
}

type UserIdentificationType_25447S struct {

	// Originator Identification Details
	OriginIdentification *OriginatorIdentificationDetailsTypeI_46358C `xml:"originIdentification,omitempty"`  // minOccurs="0"

	// Booking source or  [agent numeric sign] + [agent initial] + [duty code]
	Originator string `xml:"originator"`
}

type UserIdentificationType_9456S struct {

	// contains the client reference/signature.
	Originator string `xml:"originator,omitempty"`  // minOccurs="0"
}

type UserPreferencesType struct {

	// This composite contains details on user preferences : _ Country code _ Language code _ Currency code
	UserPreferences *OriginatorDetailsTypeI `xml:"userPreferences"`
}

type ValidityDatesTypeU struct {

	// Date the document was issued (YYYYMMDD)
	IssueDate string `xml:"issueDate,omitempty"`  // minOccurs="0"

	// Date document expires (YYYYMMDD)
	ExpirationDate string `xml:"expirationDate,omitempty"`  // minOccurs="0"
}

type ValueRangeTypeU struct {

	// Unit Qualifier for maximum range gives in previous RNG:  DAY: Duration in days G:   Kilometers M:   Mileage MTH: Durarion in Months WE:  Duration in weeks
	MeasureUnitQualifier string `xml:"measureUnitQualifier"`
}

type VehicleInformationType struct {

	// This composite is used to convey the vehicle type
	VehicleCharacteristic *VehicleTypeOptionType `xml:"vehicleCharacteristic"`

	// This data element is used to convey the equipment codes.
	VehSpecialEquipment []string `xml:"vehSpecialEquipment,omitempty"`  // minOccurs="0" maxOccurs="99"

	// To indicate vehicle details: -Number of doors -Number of seats -Max Number of doors -Max Number of seats -Number of bags -Volume of the boots
	VehicleInfo []*QuantityDetailsTypeI_198209C `xml:"vehicleInfo,omitempty"`  // minOccurs="0" maxOccurs="6"

	// Free text type
	FreeTextDetails *FreeTextDetailsType_198207C `xml:"freeTextDetails,omitempty"`  // minOccurs="0"

	// Description or Example of the Car
	CarModel string `xml:"carModel,omitempty"`  // minOccurs="0"
}

type VehicleInformationTypeU struct {

	// Vehicle make and model.
	MakeAndModel string `xml:"makeAndModel,omitempty"`  // minOccurs="0"
}

type VehicleInformationTypeU_46439C struct {

	// Conveys the occupancy of a vehicule
	Occupancy int32 `xml:"occupancy"`
}

type VehicleTypeOptionType struct {

	// This data element is used to convey the owner of the type code.
	VehicleTypeOwner string `xml:"vehicleTypeOwner"`

	// This data element is used to convey the SIPP code(s) selection criteria.
	VehicleRentalPrefType []string `xml:"vehicleRentalPrefType"`  // maxOccurs="2"
}

type VehicleTypeU struct {

	// Conveys the type of vehicle.
	Category string `xml:"category"`

	// Describes the vehicle.
	VehicleDetails *VehicleInformationTypeU `xml:"vehicleDetails,omitempty"`  // minOccurs="0"
}

type VehicleTypeU_25502S struct {

	// Describe the vehicule
	VehiculeDescription *VehicleInformationTypeU_46439C `xml:"vehiculeDescription"`
}

type CoverageInfo struct {

	// For codelist 415Z, only values CP, CV, CM may apply here
	Coverage *InsuranceCoverageType `xml:"coverage"`

	// Values and currency of the different coverages amounts.
	CoverageValues *MonetaryInformationTypeI `xml:"coverageValues,omitempty"`  // minOccurs="0"
}

type PremiumPerTariffPerPax struct {

	// To convey the tariffcode information per passenger.
	TariffCodeInfo *InsuranceProductDetailsType_20775S `xml:"tariffCodeInfo"`

	// to specify the amount/currency  per passenger and tariff code.
	TariffCodePerPaxAmount *MonetaryInformationTypeI `xml:"tariffCodePerPaxAmount,omitempty"`  // minOccurs="0"
}

type TravelerValueDetails struct {

	// For codelist 415Z, only TV value may apply here
	TravelCost *InsuranceCoverageType `xml:"travelCost"`

	// to specify the amount/currency  per insuree
	TravelAmount *MonetaryInformationTypeI `xml:"travelAmount,omitempty"`  // minOccurs="0"
}

type AccomodationDetails struct {

	// Conveys room information
	RoomInfo *HotelRoomType_25429S `xml:"roomInfo"`

	// Passenger association at accomodation room level
	PassengerAssociation []*ReferenceInfoType_25422S `xml:"passengerAssociation,omitempty"`  // minOccurs="0" maxOccurs="9"

	// Conveys room meal plan information
	RoomMealPlanInfo *DiningInformationType `xml:"roomMealPlanInfo"`

	// Conveys room occupancy information (room min or max occupancy)
	OccupancynInfo *RangeDetailsTypeU `xml:"occupancynInfo,omitempty"`  // minOccurs="0"
}

type ProductBCSDetails struct {

	// BCS Agent sign, office and target office ids
	AgentIdentification *UserIdentificationType_25447S `xml:"agentIdentification"`

	// BCS distribution channel
	DistributionChannelData *SystemDetailsInfoType `xml:"distributionChannelData"`
}

type TransportationDetails struct {

	// Departure location information. The composite C517 conveys city information and the C519 the country information.
	DepartureInfo *PlaceLocationIdentificationTypeU_25436S `xml:"departureInfo"`

	// Arrival location information. The composite C517 conveys city information and the C519 the country information.
	ArrivalInfo *PlaceLocationIdentificationTypeU_25436S `xml:"arrivalInfo"`

	// Conveys transportation information
	TransportationInfo *TravelProductInformationTypeI_25434S `xml:"transportationInfo"`

	// Conveys duration information (number of day, night ...)
	TransportationDuration *QuantityType `xml:"transportationDuration,omitempty"`  // minOccurs="0"

	// Conveys transportation equipment information
	EquipmentInfo *EquipmentDetailsTypeU `xml:"equipmentInfo,omitempty"`  // minOccurs="0"

	// Conveys transportation meal plan information
	TransportationMealPlanInfo *DiningInformationType `xml:"transportationMealPlanInfo,omitempty"`  // minOccurs="0"
}

type VehiculeDetails struct {

	// Conveys vehicule information (such as the vehicule occupancy)
	VehiculeInfo *VehicleTypeU_25502S `xml:"vehiculeInfo"`
}
