package pnr_reply

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type PNRReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PNR_Reply"`

	PnrHeader *PnrHeader `xml:"pnrHeader,omitempty"`

	// specify the amadeus PNR record locator security information for different pnr elements .
	SecurityInformation *ReservationSecurityInformationType `xml:"securityInformation,omitempty"`

	// specify a queue , that is  an office , a queue number , a category number ,and a date range number or a date .This segment can be used in an output message.
	QueueInformations *QueueType `xml:"queueInformations,omitempty"`

	// specify the number of units required
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits,omitempty"`

	GeneralErrorInfo *GeneralErrorInfo `xml:"generalErrorInfo,omitempty"`

	// Conveys PNR special types
	PnrType *CodedAttributeType `xml:"pnrType,omitempty"`

	// provide free form or coded long text information
	FreetextData *LongFreeTextType `xml:"freetextData,omitempty"`

	// Conveys the Header tags of the PNR.
	PnrHeaderTag *StatusType `xml:"pnrHeaderTag,omitempty"`

	// provide free from or coded text information
	FreeFormText *InteractiveFreeTextTypeI_136698S `xml:"freeFormText,omitempty"`

	// specify an amadeus PNR history data
	HistoryData *PnrHistoryDataType_6022S `xml:"historyData,omitempty"`

	// Point of Sale Information at SBR level. Owner of the PNR.
	SbrPOSDetails *POSGroupType `xml:"sbrPOSDetails,omitempty"`

	// Point of Sale Information at SBR level. Creator of the PNR.
	SbrCreationPosDetails *POSGroupType `xml:"sbrCreationPosDetails,omitempty"`

	// Point of Sale Information at SBR level. Updator of the PNR.
	SbrUpdatorPosDetails *POSGroupType `xml:"sbrUpdatorPosDetails,omitempty"`

	TechnicalData *TechnicalData `xml:"technicalData,omitempty"`

	TravellerInfo *TravellerInfo `xml:"travellerInfo,omitempty"`

	OriginDestinationDetails *OriginDestinationDetails `xml:"originDestinationDetails,omitempty"`

	// specify the segment marriages and connections
	SegmentGroupingInfo *SegmentGroupingInformationType `xml:"segmentGroupingInfo,omitempty"`

	DataElementsMaster DataElementsMaster `xml:"dataElementsMaster,omitempty"`

	TstData *TstData `xml:"tstData,omitempty"`

	DcdData *DcdData `xml:"dcdData,omitempty"`
}

type TstData struct {

	// TST general information
	TstGeneralInformation *TstGeneralInformationType `xml:"tstGeneralInformation,omitempty"`

	// provide free form or coded long text information
	TstFreetext *LongFreeTextType `xml:"tstFreetext,omitempty"`

	// describe fare basis information
	FareBasisInfo *FareBasisCodesLineType `xml:"fareBasisInfo,omitempty"`

	// fare data
	FareData *FareDataType `xml:"fareData,omitempty"`

	// selection details
	SegmentAssociation *SelectionDetailsTypeI_2067S `xml:"segmentAssociation,omitempty"`

	// provide specific reference identification
	ReferenceForTstData *ReferenceInfoType `xml:"referenceForTstData,omitempty"`
}

type DcdData struct {

	// This is used as DUM segment
	MarkerPax *PassengerFlightDetailsTypeI `xml:"markerPax,omitempty"`

	// This is used as a Dum segment.
	MarkerSegment *PassengerFlightDetailsTypeI `xml:"markerSegment,omitempty"`

	SegmentSection *SegmentSection `xml:"segmentSection,omitempty"`

	// This is used as a Dum segement.
	MarkerLeg *PassengerFlightDetailsTypeI `xml:"markerLeg,omitempty"`

	LegSection *LegSection `xml:"legSection,omitempty"`
}

type LegSection struct {

	// Specify structured elements references
	ElementManagementStructData *ElementManagementSegmentType_127983S `xml:"elementManagementStructData,omitempty"`

	// provide specific reference identification
	ReferenceForStructDataElement *ReferenceInfoType `xml:"referenceForStructDataElement,omitempty"`

	DcsLegInfo *DcsLegInfo `xml:"dcsLegInfo,omitempty"`
}

type SegmentSection struct {

	// Specify structured elements references
	ElementManagementStructData *ElementManagementSegmentType_127983S `xml:"elementManagementStructData,omitempty"`

	// provide specific reference identification
	ReferenceForStructDataElement *ReferenceInfoType `xml:"referenceForStructDataElement,omitempty"`

	DcsSegmentInfo *DcsSegmentInfo `xml:"dcsSegmentInfo,omitempty"`
}

type DcsSegmentInfo struct {

	// Booking information. This is not required by the process it self, but can be used to easily track problems. This segment is required also to solve ambiguity problems. It can be empty if there is not need to convey information (more meaningful than a DUM)
	Booking *TravelProductInformationTypeI_127288S `xml:"booking,omitempty"`

	// Determines if the DCS Data apply to the adult or to the infant (in case there is one). By default, it applies to the adult.
	PaxType *ReferenceInformationTypeI `xml:"paxType,omitempty"`

	// Gives the compensation type. * attributeDetails/attributeType = - DBA - DBN - DBO - DBV - DBM - DBD - DBW
	TypeOfCOP *CodedAttributeType_127282S `xml:"typeOfCOP,omitempty"`
}

type DcsLegInfo struct {

	// Contains the leg position inside the booking
	LegPosition *TravelItineraryInformationTypeI `xml:"legPosition,omitempty"`

	// Indenties uniquely a leg inside a multi-leg booking
	Leg *OriginAndDestinationDetailsTypeI `xml:"leg,omitempty"`

	// Determines if the DCS Data apply to the adult or to the infant (in case there is one). By default, it applies to the adult.
	PaxType *ReferenceInformationTypeI `xml:"paxType,omitempty"`

	// Contains information on the seat delivered by the DCS
	SeatDelivery *SpecialRequirementsDetailsType `xml:"seatDelivery,omitempty"`

	// Third data element provide the category of attribute: NOREC information, acceptance status... First data element contains the value of the attribute : the NOREC flag, the acceptance status, the boarding status and the cabin regrade type, Check Bags indicator, Waitlist status.
	PaxStatus *StatusTypeI `xml:"paxStatus,omitempty"`

	AccregReason *AccregReason `xml:"accregReason,omitempty"`

	// Regrade cabin code
	RegradeCabin *SegmentCabinIdentificationType `xml:"regradeCabin,omitempty"`

	AcceptanceChannel *AcceptanceChannel `xml:"acceptanceChannel,omitempty"`

	// Provides information on the compensation offered to passengers with valid tickets, airline turned down at check-in/boarding. - coded form of payment (NGDCS only) - currency code (NGDCS only) - amount (NGDCS only) - compensation type (voluntary/involuntary) - free text (both valid for PFS clients and NGDCS)
	CompensationData *CompensationType `xml:"compensationData,omitempty"`
}

type AcceptanceChannel struct {

	// Qualifies originator of the acceptance actions: - A for Check-in Agent - D for Direct consumer  - S for automated Devices - E for internet/web based application
	AcceptanceOrigin *UserIdentificationType_127265S `xml:"acceptanceOrigin,omitempty"`

	// contains the application used to perform check-in operations: - cryptic application - Java front End application - SMS application - Telephone - web application - External departure control application
	ApplicationType *ApplicationType `xml:"applicationType,omitempty"`
}

type AccregReason struct {

	// Reason code for: - Acceptation/Cancellation - Regrade
	Reasons *CodedAttributeType_127279S `xml:"reasons,omitempty"`

	// Contains Acceptance Reason, Regrade Reason description.
	DeliveryInformation *InteractiveFreeTextTypeI `xml:"deliveryInformation,omitempty"`
}

type DataElementsMaster struct {

	// marker
	Marker2 *DummySegmentTypeI `xml:"marker2,omitempty"`

	DataElementsIndiv *DataElementsIndiv `xml:"dataElementsIndiv,omitempty"`
}

type DataElementsIndiv struct {

	// specify the PNR segments/elements references and action to apply
	ElementManagementData *ElementManagementSegmentType `xml:"elementManagementData,omitempty"`

	// specify the amadeus PNR individual security element
	PnrSecurity *IndividualPnrSecurityInformationType `xml:"pnrSecurity,omitempty"`

	// Specify the amadeus accounting information
	Accounting *AccountingInformationElementType `xml:"accounting,omitempty"`

	// specify miscellaneous, confidential, quality control and invoice remarks
	MiscellaneousRemarks *MiscellaneousRemarksType_211S `xml:"miscellaneousRemarks,omitempty"`

	// specify special request or services information relating to a traveller
	ServiceRequest *SpecialRequirementsDetailsTypeI `xml:"serviceRequest,omitempty"`

	SeatPaxInfo *SeatPaxInfo `xml:"seatPaxInfo,omitempty"`

	// To Convey the Reason for Issuance Code (RFIC) and Reason For Issuance Sub code (RFISC)
	ReasonForIssuanceCode *PricingOrTicketingSubsequentType `xml:"reasonForIssuanceCode,omitempty"`

	// Rail Seat Preferences
	RailSeatPreferences *RailSeatPreferencesType `xml:"railSeatPreferences,omitempty"`

	CityPair *CityPair `xml:"cityPair,omitempty"`

	RailSeatDetails *RailSeatDetails `xml:"railSeatDetails,omitempty"`

	// provide date and time details relative to flight movements
	DateAndTimeInformation *DateAndTimeInformationTypeI `xml:"dateAndTimeInformation,omitempty"`

	// Details of SSR FQT content (Frequent Flyer Data)
	FrequentFlyerInformationGroup *FrequentFlyerInformationGroupType `xml:"frequentFlyerInformationGroup,omitempty"`

	// specify an amadeus PNR ticket element
	TicketElement *TicketElementType `xml:"ticketElement,omitempty"`

	ReferencedRecord *ReferencedRecord `xml:"referencedRecord,omitempty"`

	// option element
	OptionElement *OptionElementType `xml:"optionElement,omitempty"`

	// provide free form or coded long text information
	OtherDataFreetext *LongFreeTextType `xml:"otherDataFreetext,omitempty"`

	// specify the way data are mapped for the structured addresses
	StructuredAddress *StructuredAddressType `xml:"structuredAddress,omitempty"`

	// To specify the monetary information
	MonetaryInformation *MonetaryInformationTypeI_1689S `xml:"monetaryInformation,omitempty"`

	ElementErrorInformation *ElementErrorInformation `xml:"elementErrorInformation,omitempty"`

	McoRecord *McoRecord `xml:"mcoRecord,omitempty"`

	// Group Total Price
	TotalPrice *TotalPriceType `xml:"totalPrice,omitempty"`

	// Indicators at element level
	ElementsIndicators *StatusTypeI_33257S `xml:"elementsIndicators,omitempty"`

	// provide specific reference identification
	ReferenceForDataElement *ReferenceInfoType `xml:"referenceForDataElement,omitempty"`

	// Carries a Form of Payment in structured way.
	StructuredFop *FOPRepresentationType `xml:"structuredFop,omitempty"`
}

type McoRecord struct {

	// specify that a MCO element is present in the PNR - this is a visual trigger of the MCO
	McoType *MiscellaneousChargeOrderType `xml:"mcoType,omitempty"`

	// Contains the data relative to the MCO element itself
	McoInformation *FreeTextInformationType_9865S `xml:"mcoInformation,omitempty"`

	GroupOfFareElements *GroupOfFareElements `xml:"groupOfFareElements,omitempty"`
}

type GroupOfFareElements struct {

	// Sequence Number for a Fare element
	SequenceNumber *SequenceDetailsTypeU `xml:"sequenceNumber,omitempty"`

	// Contains the Fare Element data
	FareElementData *FreeTextInformationType_9865S `xml:"fareElementData,omitempty"`
}

type ElementErrorInformation struct {

	// identify the type of application error  within a message
	ErrorInformation *ApplicationErrorInformationType `xml:"errorInformation,omitempty"`

	// provide free from or coded text information
	ElementErrorText *InteractiveFreeTextTypeI_136698S `xml:"elementErrorText,omitempty"`
}

type ReferencedRecord struct {

	// specify a reference to a reservation
	ReferencedReservationInfo *ReservationControlInformationTypeI_87792S `xml:"referencedReservationInfo,omitempty"`

	// specify the amadeus PNR record locator security information for different pnr elements .
	SecurityInformation *ReservationSecurityInformationType_167761S `xml:"securityInformation,omitempty"`
}

type RailSeatDetails struct {

	// Used to convey specific seat details relative to Train for a specific request or the "near-to" seat details for a "next-to" request.
	RailSeatReferenceInformation *RailSeatReferenceInformationType `xml:"railSeatReferenceInformation,omitempty"`

	// Rail Seat Denomination
	RailSeatDenomination *FreeTextInformationType_29860S `xml:"railSeatDenomination,omitempty"`
}

type CityPair struct {

	// Departure station location
	DepLocation *PlaceLocationIdentificationTypeU_35293S `xml:"depLocation,omitempty"`

	// Arrival station location
	ArrLocation *PlaceLocationIdentificationTypeU_35293S `xml:"arrLocation,omitempty"`
}

type SeatPaxInfo struct {

	// details of the seat at pax level
	SeatPaxDetails *SeatRequestParametersTypeI `xml:"seatPaxDetails,omitempty"`

	// seat indicator at pax level
	SeatPaxIndicator *StatusTypeI_33257S `xml:"seatPaxIndicator,omitempty"`

	// ref to pax tattoo
	CrossRef *ReferenceInfoType_6074S `xml:"crossRef,omitempty"`
}

type PnrHeader struct {
	// specify a reference to a reservation
	ReservationInfo *ReservationControlInformationTypeI_87792S `xml:"reservationInfo,omitempty"`

	// provide specific reference identification
	ReferenceForRecordLocator *ReferenceInfoType `xml:"referenceForRecordLocator,omitempty"`
}

type GeneralErrorInfo struct {
	// identify the type of application error  within a message
	MessageErrorInformation *ApplicationErrorInformationType `xml:"messageErrorInformation,omitempty"`

	// provide free from or coded text information
	MessageErrorText *InteractiveFreeTextTypeI_136698S `xml:"messageErrorText,omitempty"`
}

type TechnicalData struct {
	// Contains the enveloppe number of the PNR, issue at last EOT.
	EnveloppeNumberData *SequenceDetailsTypeU `xml:"enveloppeNumberData,omitempty"`

	// CS assumption on last transmitted envelope number
	LastTransmittedEnvelopeNumber *PnrHistoryDataType `xml:"lastTransmittedEnvelopeNumber,omitempty"`

	// Contains the Purge Date of the PNR at the time of the retrieval
	PurgeDateData *StructuredDateTimeInformationType_27086S `xml:"purgeDateData,omitempty"`

	// Contains general information relative to the state of the PNR
	GeneralPNRInformation *StatusTypeI_32775S `xml:"generalPNRInformation,omitempty"`
}

type TravellerInfo struct {
	// specify the PNR segments/elements references and action to apply
	ElementManagementPassenger *ElementManagementSegmentType `xml:"elementManagementPassenger,omitempty"`

	PassengerData *PassengerData `xml:"passengerData,omitempty"`

	NameError *NameError `xml:"nameError,omitempty"`
}

type NameError struct {

	// identify the type of application error  within a message
	NameErrorInformation *ApplicationErrorInformationType `xml:"nameErrorInformation,omitempty"`

	// provide free form or error coded text
	NameErrorFreeText *InteractiveFreeTextTypeI_136698S `xml:"nameErrorFreeText,omitempty"`
}

type OriginDestinationDetails struct {

	// convey origin and destination of a journey
	OriginDestination *OriginAndDestinationDetailsTypeI_3061S `xml:"originDestination,omitempty"`

	ItineraryInfo *ItineraryInfo `xml:"itineraryInfo,omitempty"`
}

type ItineraryInfo struct {

	// specify the PNR segments/elements references and action to apply
	ElementManagementItinerary *ElementManagementSegmentType `xml:"elementManagementItinerary,omitempty"`

	// specify details related to a product
	TravelProduct *TravelProductInformationTypeI `xml:"travelProduct,omitempty"`

	// specify the message type and business function
	ItineraryMessageAction *MessageActionDetailsTypeI `xml:"itineraryMessageAction,omitempty"`

	// specify a reference to a reservation
	ItineraryReservationInfo *ReservationControlInformationTypeI_115879S `xml:"itineraryReservationInfo,omitempty"`

	// indicate quantity and action required in relation to a product
	RelatedProduct *RelatedProductInformationTypeI `xml:"relatedProduct,omitempty"`

	// Indicator at element level: - CGB (chargeable indicator)
	ElementsIndicators *StatusType_99946S `xml:"elementsIndicators,omitempty"`

	// To Convey the Reason for Issuance Code (RFIC) and Reason For Issuance Sub code (RFISC)
	ReasonForIssuanceCode *PricingOrTicketingSubsequentType `xml:"reasonForIssuanceCode,omitempty"`

	// convey additional information concerning an airline flight
	FlightDetail *AdditionalProductDetailsTypeI `xml:"flightDetail,omitempty"`

	// Contain the cabin class details
	CabinDetails *CabinDetailsType `xml:"cabinDetails,omitempty"`

	// specify details for making a selection
	SelectionDetails *SelectionDetailsTypeI_2067S `xml:"selectionDetails,omitempty"`

	// provide free from or coded text information
	ItineraryfreeFormText *InteractiveFreeTextTypeI_136698S `xml:"itineraryfreeFormText,omitempty"`

	// provide free form or coded long text information
	ItineraryFreetext *LongFreeTextType `xml:"itineraryFreetext,omitempty"`

	// specify the details for hotel transaction
	HotelProduct *HotelProductInformationType `xml:"hotelProduct,omitempty"`

	// specify the Rate information
	RateInformations *RateInformationType `xml:"rateInformations,omitempty"`

	// specify one option
	GeneralOption *GeneralOptionType `xml:"generalOption,omitempty"`

	// provide the ISO code of a country
	Country *CountryCodeListType `xml:"country,omitempty"`

	// To specify details relating to tax(es).
	TaxInformation *TaxTypeI `xml:"taxInformation,omitempty"`

	// specify operating flight additional information
	CustomerTransactionData *CustomerTransactionDataType `xml:"customerTransactionData,omitempty"`

	YieldGroup *YieldGroup `xml:"yieldGroup,omitempty"`

	LegInfo *LegInfo `xml:"legInfo,omitempty"`

	// Group used to carry FLIX information associated to the segment
	FlixInfo *FLIXType `xml:"flixInfo,omitempty"`

	// Provide Date and Time Details relative to the Itinerary
	DateTimeDetails *DateAndTimeInformationTypeI `xml:"dateTimeDetails,omitempty"`

	LccTypicalData *LccTypicalData `xml:"lccTypicalData,omitempty"`

	InsuranceInformation *InsuranceInformation `xml:"insuranceInformation,omitempty"`

	// This group handles all Insurance structured details.
	InsuranceDetails *InsuranceBusinessDataType `xml:"insuranceDetails,omitempty"`

	HotelReservationInfo *HotelReservationInfo `xml:"hotelReservationInfo,omitempty"`

	TypicalCarData *TypicalCarData `xml:"typicalCarData,omitempty"`

	// Dedicated to convey the specific cruise sailing information.  It is in fact the reflection of the cruise segment PNR content.  An example of Amadeus PNR with 2 Cruise segments, one confirmed the other cancelled. RP/NYCP02001/NYCP02001 AA/SU 19AUG03/1423Z YYAA63 NYCP02001/0001AA/19AUG03 1.TEST/AAA MR 2 CRU PCL HK1 SJUSJU 18JAN2004-7 DAWN PRINCESS/CF-9VHJ2K/NM-TEST/AAA MR 3 CRU PCL HX1 XYZPKG 18MAY2004-7 CORAL PRINCESS/GP-ABCDEF/CX-654321/NM-TEST/AAA MR  4 AP 0494503434 5 TK OK19AUG/NYCP02001
	TypicalCruiseData *CruiseBusinessDataType `xml:"typicalCruiseData,omitempty"`

	// Information pertaining to a rail segment
	RailInfo *TrainInformationType `xml:"railInfo,omitempty"`

	// Marks separation between Rail Group and Tour Group, avoid Ambiguity with CPY segment.
	MarkerRailTour *DummySegmentTypeI `xml:"markerRailTour,omitempty"`

	// Dedicated to convey the specific Tour information.
	TourInfo *TourInformationType `xml:"tourInfo,omitempty"`

	// Information pertaining to a Ferry segment
	FerryLegInformation *FerryBookingDescriptionType `xml:"ferryLegInformation,omitempty"`

	ErrorInfo *ErrorInfo `xml:"errorInfo,omitempty"`

	// provide specific reference identification
	ReferenceForSegment *ReferenceInfoType `xml:"referenceForSegment,omitempty"`
}

type ErrorInfo struct {

	// identify the type of application error  within a message
	ErrorInformation *ApplicationErrorInformationType `xml:"errorInformation,omitempty"`

	// provide free from or coded text information
	ErrorfreeFormText *InteractiveFreeTextTypeI_136698S `xml:"errorfreeFormText,omitempty"`
}

type YieldGroup struct {

	// It contains some specific air segment's indicators data, not present in yieldDataGroup
	YieldData *ODKeyPerformanceDataType `xml:"yieldData,omitempty"`

	// Details of the yield data.
	YieldDataGroup *ONDType `xml:"yieldDataGroup,omitempty"`
}

type LegInfo struct {

	// Marks separation to avoid Ambiguity between TVL segments.
	MarkerLegInfo *FlightSegmentDetailsTypeI `xml:"markerLegInfo,omitempty"`

	LegTravelProduct *TravelProductInformationTypeI_99362S `xml:"legTravelProduct,omitempty"`

	InteractiveFreeText *InteractiveFreeTextTypeI_99363S `xml:"interactiveFreeText,omitempty"`
}

type LccTypicalData struct {

	// Fare data obtained from D/A availabilty (mapped under PRXP20LCC structure)
	LccFareData *TariffInformationTypeI_28460S `xml:"lccFareData,omitempty"`

	// Connection key tattoo if any
	LccConnectionData *ItemReferencesAndVersionsType_6550S `xml:"lccConnectionData,omitempty"`
}

type InsuranceInformation struct {

	// contain data related to each passenger
	InsuranceName *InsuranceNameType `xml:"insuranceName,omitempty"`

	// To specify monetary information
	InsuranceMonetaryInformation *MonetaryInformationTypeI_1689S `xml:"insuranceMonetaryInformation,omitempty"`

	// Specify an Amadeus PNR Ticket element
	InsurancePremiumInfo *TravellerInsuranceInformationType `xml:"insurancePremiumInfo,omitempty"`

	// provide traveller document information
	InsuranceDocInfo *TravellerDocumentInformationTypeU `xml:"insuranceDocInfo,omitempty"`
}

type HotelReservationInfo struct {

	// This segment is used to convey the hotel property information.
	HotelPropertyInfo *HotelPropertyType `xml:"hotelPropertyInfo,omitempty"`

	// This segment is used to convey the hotel chain code and name.
	CompanyIdentification *CompanyInformationType `xml:"companyIdentification,omitempty"`

	// This segment is used to convey the dates.
	RequestedDates *StructuredPeriodInformationType_11026S `xml:"requestedDates,omitempty"`

	RoomRateDetails *RoomRateDetails `xml:"roomRateDetails,omitempty"`

	// This segment is used to convey the confirmation number or the cancellation number. control type (9958) is: - 2 for Confirmation reference - X for cancellation reference - O for on request reference
	CancelOrConfirmNbr *ReservationControlInformationTypeI_115879S `xml:"cancelOrConfirmNbr,omitempty"`

	// indicate the roomstay index in case of groups
	RoomstayIndex *ItemNumberTypeU_33258S `xml:"roomstayIndex,omitempty"`

	// This segment is used to convey the booking source.
	BookingSource *UserIdentificationType_21014S `xml:"bookingSource,omitempty"`

	// This segment is used to convey the billable information
	BillableInfo *BillableInformationTypeU `xml:"billableInfo,omitempty"`

	// This segment is used to convey the customer reference number
	CustomerInfo *ConsumerReferenceInformationTypeI `xml:"customerInfo,omitempty"`

	// This segment is used to convey the frequent traveler number.
	FrequentTravellerInfo *FrequentTravellerIdentificationCodeType_38226S `xml:"frequentTravellerInfo,omitempty"`

	GuaranteeOrDeposit *GuaranteeOrDeposit `xml:"guaranteeOrDeposit,omitempty"`

	// This segment is used to convey additional information which are entered by the agent and stored on the hotel booking.
	TextOptions *MiscellaneousRemarksType_664S `xml:"textOptions,omitempty"`

	// This segment is used to convey the saving amount/percentage information.
	SavingAmountInfo *MonetaryInformationTypeI_1689S `xml:"savingAmountInfo,omitempty"`

	// This segment is used to convey the fax or the E-Mail to receive the written confirmation
	WrittenConfirmationContact *ContactInformationTypeU `xml:"writtenConfirmationContact,omitempty"`

	// This segment is used to convey the name and address of the party to recieve the written confirmation
	WrittenConfirmationInfo *NameAndAddressBatchTypeU `xml:"writtenConfirmationInfo,omitempty"`

	// documentInformationDetails
	DocumentInformationDetails *DocumentInformationDetailsTypeI_9936S `xml:"documentInformationDetails,omitempty"`

	ArrivalFlightDetails *ArrivalFlightDetails `xml:"arrivalFlightDetails,omitempty"`

	// This segment is used to convey specific hotel booking indicators
	BookingIndicator *StatusType_99582S `xml:"bookingIndicator,omitempty"`
}

type RoomRateDetails struct {

	// This segment is used to convey the room information.
	RoomInformation *HotelRoomType `xml:"roomInformation,omitempty"`

	// This group is used to conveys list of children
	Children *ChildrenGroupType `xml:"children,omitempty"`

	// This segment is used to convey the tariff/rate details.
	TariffDetails *TariffInformationTypeI `xml:"tariffDetails,omitempty"`

	// This segment is used to convey the rate code indicator.
	RateCodeIndicator *RuleInformationTypeU `xml:"rateCodeIndicator,omitempty"`
}

type GuaranteeOrDeposit struct {

	// This segment is used to convey the guarantee or deposit information
	PaymentInfo *PaymentInformationTypeI `xml:"paymentInfo,omitempty"`

	// This segment is used to convey the credit card information.
	CreditCardInfo *FormOfPaymentTypeI_29553S `xml:"creditCardInfo,omitempty"`
}

type ArrivalFlightDetails struct {

	// Travel Product Information
	TravelProductInformation *TravelProductInformationTypeI_186189S `xml:"travelProductInformation,omitempty"`

	// Additional Transport Details
	AdditionalTransportDetails *AdditionalTransportDetailsTypeU `xml:"additionalTransportDetails,omitempty"`
}

type TypicalCarData struct {

	// Vehicle information - vehicle type (SIPP code), - vehicle special equipments - vehicle details
	VehicleInformation *VehicleInformationType `xml:"vehicleInformation,omitempty"`

	// Additional vehicle info
	AdditionalInfo *FreeTextInformationType_136708S `xml:"additionalInfo,omitempty"`

	// Voucher Print Acknowledgement.
	VoucherPrintAck *ReferenceInformationTypeI_136704S `xml:"voucherPrintAck,omitempty"`

	// CAR provider code
	CompanyIdentification *CompanyInformationType `xml:"companyIdentification,omitempty"`

	// Car AVL pickup and dropoff location parameters (for Amadeus and Provider locations). Used as well to transport the Collection and Delivery place information.
	LocationInfo *PlaceLocationIdentificationTypeU `xml:"locationInfo,omitempty"`

	DeliveryAndCollection *DeliveryAndCollection `xml:"deliveryAndCollection,omitempty"`

	// Pickup and dropoff dates and times.
	PickupDropoffTimes *StructuredPeriodInformationType_136705S `xml:"pickupDropoffTimes,omitempty"`

	// Cancellation or Confirmation number.
	CancelOrConfirmNbr *ReservationControlInformationTypeI_136703S `xml:"cancelOrConfirmNbr,omitempty"`

	RateCodeGroup *RateCodeGroup `xml:"rateCodeGroup,omitempty"`

	// Frequent flyer number.
	FFlyerNbr *FrequentTravellerIdentificationCodeType `xml:"fFlyerNbr,omitempty"`

	// Customer information ID and CD numbers
	CustomerInfo *ConsumerReferenceInformationTypeI `xml:"customerInfo,omitempty"`

	// This segments is used to convey: 1)General Rate information (identifier, plan, category) and Unstructured RB/RQ/RG 2) Structured rate quoted (RQ) or guaranted (RG) 3) Structured base rate (RB) 4) Estimated total information 5) Drop amount data 6) Voucher coupon print references (VV) 7) Rate Override (RO) 8) Modification fee indicator 9) Cancellation fee indicator 10) prepayment
	RateInfo *TariffInformationTypeI_136706S `xml:"rateInfo,omitempty"`

	ErrorWarning *ErrorWarning `xml:"errorWarning,omitempty"`

	RulesPoliciesGroup *RulesPoliciesGroup `xml:"rulesPoliciesGroup,omitempty"`

	// - Form of payment (FP) - Form of guarantee (G)
	Payment *FormOfPaymentTypeI `xml:"payment,omitempty"`

	// - Billing reference number (contains conpany data to be built) - Billing mumber (number included in the billing) - Agency account
	BillingData *BillableInformationTypeU `xml:"billingData,omitempty"`

	// The booking source
	BookingSource *AdditionalBusinessSourceInformationType `xml:"bookingSource,omitempty"`

	// Tour code
	InclusiveTour *TourInformationTypeI `xml:"inclusiveTour,omitempty"`

	// Contains: 1) up to 6 lines of marketing text sent by the car provider. 2) up to 3 lines of Other services messages (advertisments).
	MarketingInfo *InteractiveFreeTextTypeI_136698S `xml:"marketingInfo,omitempty"`

	// This segment is used to convey the supplementary  informations (SI). e.g: "Customer arriving after agency closure hour. Car keys waiting at the hotel reception located next to the agency".
	SupleInfo *MiscellaneousRemarksType_136700S `xml:"supleInfo,omitempty"`

	// This segment is used to convey distances.  1) Intercity distance. Distance between the Pickup and the Dropoff cities. Information returned by the Car provider for customer notification purpose.  2) Estimated distance Distance that is going to be runned during the rental period.
	EstimatedDistance *QuantityTypeI `xml:"estimatedDistance,omitempty"`

	// Booking agent name
	AgentInformation *NameTypeU_136701S `xml:"agentInformation,omitempty"`

	// Tracking Option (TK)
	TrackingOpt *AgreementIdentificationTypeU `xml:"trackingOpt,omitempty"`

	// Electronic Voucher Number
	ElectronicVoucherNumber *TicketNumberTypeI `xml:"electronicVoucherNumber,omitempty"`

	// E-mail
	CustomerEmail *CommunicationContactTypeU `xml:"customerEmail,omitempty"`

	// This mandatory segment marks the end of the CAR data group. It specifies also if the booking is leisure or not.
	Attribute *AttributeType `xml:"attribute,omitempty"`
}

type DeliveryAndCollection struct {

	// This Segment is used to Delivery and Collection information:  Format 1- (Home Collection): - Address - City - State - Country - Zip Code  Format 2- (Site Collection): - Site Ref Id - Site Name
	AddressDeliveryCollection *AddressTypeU_136710S `xml:"addressDeliveryCollection,omitempty"`

	// This segment is used to carry phone number associated to a Delivery / Collection address
	PhoneNumber *PhoneAndEmailAddressType `xml:"phoneNumber,omitempty"`
}

type RateCodeGroup struct {

	// Rate code
	RateCodeInfo *FareQualifierDetailsTypeI `xml:"rateCodeInfo,omitempty"`

	// Additional Rate Code Information
	AdditionalInfo *FreeTextInformationType_136708S `xml:"additionalInfo,omitempty"`
}

type ErrorWarning struct {

	// Error/warning
	ApplicationError *ApplicationErrorInformationType_136725S `xml:"applicationError,omitempty"`

	// Error or Warning freetext
	ErrorFreeText *FreeTextInformationType_136708S `xml:"errorFreeText,omitempty"`
}

type RulesPoliciesGroup struct {

	// Dummy segment to mark the beginning of the group
	Dummy1 *DummySegmentTypeI `xml:"dummy1,omitempty"`

	// Present only if information is linked to a seamless availability
	SourceLevel *SelectionDetailsTypeI `xml:"sourceLevel,omitempty"`

	// Used to convey remarks corresponding to rule information.
	Remarks *FreeTextInformationType_136708S `xml:"remarks,omitempty"`

	TaxCovSurchargeGroup *TaxCovSurchargeGroup `xml:"taxCovSurchargeGroup,omitempty"`

	OtherRulesGroup *OtherRulesGroup `xml:"otherRulesGroup,omitempty"`

	PickupDropoffLocation *PickupDropoffLocation `xml:"pickupDropoffLocation,omitempty"`

	SpecialEquipmentDetails *SpecialEquipmentDetails `xml:"specialEquipmentDetails,omitempty"`
}

type TaxCovSurchargeGroup struct {

	// This segment is used to convey Tax, Coverage, Coupon, Surcharge or Delivery and collection information  (If period associated to the surcharge, tariff and period definition conveyed in group 6)
	TaxSurchargeCoverageInfo *TariffInformationTypeI_136714S `xml:"taxSurchargeCoverageInfo,omitempty"`

	// Additional information for Tax, Surcharge or Coverage section
	AdditionalInfo *FreeTextInformationType_136708S `xml:"additionalInfo,omitempty"`

	SurchargePeriods *SurchargePeriods `xml:"surchargePeriods,omitempty"`
}

type SurchargePeriods struct {

	// tariff period/distance validity in number of days, weeks, months, km, miles.
	Period *RangeDetailsTypeI `xml:"period,omitempty"`

	// This segment is used to convey Tax, Coverage, Coupon, Surcharge or Delivery and collection information
	SurchargePeriodTariff *TariffInformationTypeI_136719S `xml:"surchargePeriodTariff,omitempty"`

	// This segment conveys the Unit Qualifier for maximum range of associated RNG.
	MaximumUnitQualifier *MeasurementsBatchTypeU `xml:"maximumUnitQualifier,omitempty"`
}

type OtherRulesGroup struct {

	// Used to convey the following type of information: - Pickup Information - Advance Payment Information - Policy Information - Deposit Information - Advance Booking Information - Guarantee Information - One Way Information
	OtherRules *RuleInformationTypeU_136720S `xml:"otherRules,omitempty"`

	// Used to convey date/time Information (only used for Pickup and Guarantee rules)
	DateTimeInfo *StructuredPeriodInformationType_136705S `xml:"dateTimeInfo,omitempty"`
}

type PickupDropoffLocation struct {

	// Car AVL pickup and dropoff location parameters (for Amadeus and Provider locations). Used as well to transport the Collection and Delivery place information.
	LocationInfo *PlaceLocationIdentificationTypeU_136722S `xml:"locationInfo,omitempty"`

	// Location Address
	Address *AddressTypeU_136721S `xml:"address,omitempty"`

	// Location opening hours
	OpeningHours *StructuredPeriodInformationType_136724S `xml:"openingHours,omitempty"`

	// Phone / Fax number
	Phone *PhoneAndEmailAddressType_136723S `xml:"phone,omitempty"`
}

type PassengerData struct {
	// specify traveler and personal details relating to a traveler
	TravellerInformation *TravellerInformationTypeI_6097S `xml:"travellerInformation,omitempty"`

	// In case of group, contains the group counters (Booked, Canceled, Split)
	GroupCounters *NumberOfUnitsType_76106S `xml:"groupCounters,omitempty"`

	// Inf/Child date of birth (DDMMYYYY). For instance 01122007
	DateOfBirth *DateAndTimeInformationType `xml:"dateOfBirth,omitempty"`
}

type SpecialEquipmentDetails struct {

	// DUM used to remove ambiguity between TFFs
	Dummy2 *DummySegmentTypeI `xml:"dummy2,omitempty"`

	RangePeriod *RangePeriod `xml:"rangePeriod,omitempty"`

	// Additional special equipment information
	AdditionalInfo *FreeTextInformationType_136715S `xml:"additionalInfo,omitempty"`

	// First TFF occurence convey main data: - 013 spec. equipment code - Qualifier (Included / Optional) - Spec. equipment name  Up to 5 next occurences convey tarrif periods. - 013 spec. equipment code - converted indicator - amount/currency - period (/day, /weekend, /week, /month, /rental) - max amount / currency
	SpecialEquipmentTariff *TariffInformationTypeI_136714S `xml:"specialEquipmentTariff,omitempty"`
}

type RangePeriod struct {

	// define age period validity associted to the special equipment
	AgePeriod *RangeDetailsTypeI `xml:"agePeriod,omitempty"`

	// This segment conveys the Unit Qualifier for maximum range of associated RNG.
	MaximumUnitQualifier *MeasurementsBatchTypeU `xml:"maximumUnitQualifier,omitempty"`
}

type AccommodationAllocationInformationDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AccommodationAllocationInformationDetailsTypeU"`

	// Accommodation (room/compartment) number
	ReferenceId formats.AlphaNumericString_Length1To4 `xml:"referenceId,omitempty"`

	// Accommodation (room/compartment) code
	Code formats.AlphaNumericString_Length1To2 `xml:"code,omitempty"`
}

type AccommodationAllocationInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AccommodationAllocationInformationTypeU"`

	// Allocated accommodation
	AccommAllocation *AccommodationAllocationInformationDetailsTypeU `xml:"accommAllocation,omitempty"`
}

type AccountingElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AccountingElementType"`

	// Account number
	Number formats.AlphaNumericString_Length1To10 `xml:"number,omitempty"`

	// Cost Number
	CostNumber formats.AlphaNumericString_Length1To50 `xml:"costNumber,omitempty"`

	// IATA company number
	CompanyNumber formats.AlphaNumericString_Length1To12 `xml:"companyNumber,omitempty"`

	// Client Reference Number
	ClientReference formats.AlphaNumericString_Length1To50 `xml:"clientReference,omitempty"`

	GSTTaxDetails formats.AlphaNumericString_Length1To109 `xml:"gSTTaxDetails,omitempty"`
}

type AccountingInformationElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AccountingInformationElementType"`

	// One of these 4 data elements is mandatory , but non in particular
	Account *AccountingElementType `xml:"account,omitempty"`

	// Number of units qualifier
	AccountNumberOfUnits formats.AlphaNumericString_Length1To3 `xml:"accountNumberOfUnits,omitempty"`
}

type ActionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ActionDetailsTypeI"`

	// Contains the details about the product knowledge
	NumberOfItemsDetails *ProcessingInformationTypeI `xml:"numberOfItemsDetails,omitempty"`
}

type AdditionalBusinessSourceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AdditionalBusinessSourceInformationType"`

	// ORIGINATOR DETAILS
	OriginatorDetails *OriginatorIdentificationDetailsTypeI_198179C `xml:"originatorDetails,omitempty"`
}

type AdditionalProductDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AdditionalProductDetailsTypeI"`

	ProductDetails *AdditionalProductTypeI `xml:"productDetails,omitempty"`

	DepartureInformation *StationInformationTypeI `xml:"departureInformation,omitempty"`

	ArrivalStationInfo *StationInformationTypeI_119771C `xml:"arrivalStationInfo,omitempty"`

	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`

	TimeDetail *TravellerTimeDetailsTypeI `xml:"timeDetail,omitempty"`

	Facilities *ProductFacilitiesTypeI `xml:"facilities,omitempty"`
}

type AdditionalProductDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AdditionalProductDetailsTypeU"`

	// Conveys the product area (TOU)
	ProductArea formats.AlphaNumericString_Length1To3 `xml:"productArea,omitempty"`

	// The general product description
	ProductDetails *ProductDataInformationTypeU `xml:"productDetails,omitempty"`
}

type AdditionalProductTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AdditionalProductTypeI"`

	Equipment formats.AlphaNumericString_Length1To3 `xml:"equipment,omitempty"`

	NumOfStops formats.NumericInteger_Length1To2 `xml:"numOfStops,omitempty"`

	Duration formats.Time24_HHMM `xml:"duration,omitempty"`

	WeekDay formats.NumericInteger_Length1To1 `xml:"weekDay,omitempty"`
}

type AdditionalTransportDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AdditionalTransportDetailsTypeU"`

	// Terminal Information
	TerminalInformation *TerminalInformationTypeU `xml:"terminalInformation,omitempty"`
}

type AddressDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressDetailsTypeU"`

	// Address Format . Will be 5 unstructured
	Format formats.AlphaNumericString_Length1To3 `xml:"format,omitempty"`

	// Address Text. Any of the following address lines may start with a tag: Door number- Street- ExternalNumber- InternalNumber- County- Neighbourhood- State-
	Line1 formats.AlphaNumericString_Length1To70 `xml:"line1,omitempty"`

	Line2 formats.AlphaNumericString_Length1To70 `xml:"line2,omitempty"`

	Line3 formats.AlphaNumericString_Length1To70 `xml:"line3,omitempty"`

	Line4 formats.AlphaNumericString_Length1To70 `xml:"line4,omitempty"`

	Line5 formats.AlphaNumericString_Length1To70 `xml:"line5,omitempty"`

	Line6 formats.AlphaNumericString_Length1To70 `xml:"line6,omitempty"`
}

type AddressDetailsTypeU_17987C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressDetailsTypeU_17987C"`

	// To specify what kind of address we have
	Format formats.AlphaNumericString_Length1To3 `xml:"format,omitempty"`

	// Address
	Line1 formats.AlphaNumericString_Length1To50 `xml:"line1,omitempty"`

	// Address
	Line2 formats.AlphaNumericString_Length1To50 `xml:"line2,omitempty"`

	// PO Box
	Line4 formats.AlphaNumericString_Length1To8 `xml:"line4,omitempty"`
}

type AddressDetailsTypeU_198210C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressDetailsTypeU_198210C"`

	// Address format
	Format formats.AlphaNumericString_Length1To3 `xml:"format,omitempty"`

	// Address Field in free flow text
	Line1 formats.AlphaNumericString_Length1To60 `xml:"line1,omitempty"`
}

type AddressDetailsTypeU_198226C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressDetailsTypeU_198226C"`

	// - 5 for unstructured address
	Format formats.AlphaNumericString_Length1To3 `xml:"format,omitempty"`

	// address line 1
	Line1 formats.AlphaNumericString_Length1To70 `xml:"line1,omitempty"`

	// address line 2
	Line2 formats.AlphaNumericString_Length1To70 `xml:"line2,omitempty"`
}

type AddressTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressTypeU"`

	// to specify the information of the address
	AddressDetails *AddressDetailsTypeU_17987C `xml:"addressDetails,omitempty"`

	// city name of the given address
	City formats.AlphaNumericString_Length1To30 `xml:"city,omitempty"`

	// zip code of the given address
	ZipCode formats.AlphaNumericString_Length1To17 `xml:"zipCode,omitempty"`

	// To convey a sub-entity within a country : region, states..
	RegionDetails *CountrySubEntityDetailsTypeU `xml:"regionDetails,omitempty"`

	// to specify the countryname
	LocationDetails *LocationIdentificationTypeU `xml:"locationDetails,omitempty"`
}

type AddressTypeU_136710S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressTypeU_136710S"`

	// Address Type
	AddressUsageDetails *AddressUsageTypeU `xml:"addressUsageDetails,omitempty"`

	// Format 1 - Home Delivery/Collection
	AddressDetails *AddressDetailsTypeU_198210C `xml:"addressDetails,omitempty"`

	// City name
	City formats.AlphaNumericString_Length1To30 `xml:"city,omitempty"`

	// Postal Code
	ZipCode formats.AlphaNumericString_Length1To10 `xml:"zipCode,omitempty"`

	// Country code
	CountryCode formats.AlphaNumericString_Length1To2 `xml:"countryCode,omitempty"`

	// To convey a sub-entity within a country : region, states..
	RegionDetails *CountrySubEntityDetailsTypeU_198213C `xml:"regionDetails,omitempty"`

	// Format 2 - Site Delivery/Collection
	LocationDetails *LocationIdentificationTypeU_198211C `xml:"locationDetails,omitempty"`
}

type AddressTypeU_136721S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressTypeU_136721S"`

	// Location address
	AddressDetails *AddressDetailsTypeU_198226C `xml:"addressDetails,omitempty"`

	// City Name
	City formats.AlphaNumericString_Length1To35 `xml:"city,omitempty"`

	// Postal Code
	ZipCode formats.AlphaNumericString_Length1To17 `xml:"zipCode,omitempty"`

	// Country code
	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"`

	// To convey a sub-entity within a country : region, states..
	RegionDetails *CountrySubEntityDetailsTypeU_198229C `xml:"regionDetails,omitempty"`
}

type AddressType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressType"`

	// will convey the adress text
	AddressDetails *AddressDetailsTypeU `xml:"addressDetails,omitempty"`

	// City name.
	City formats.AlphaNumericString_Length1To35 `xml:"city,omitempty"`

	// postal identification code.
	ZipCode formats.AlphaNumericString_Length1To17 `xml:"zipCode,omitempty"`

	// Country code. ISO 3166 code for the country
	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"`
}

type AddressUsageTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AddressUsageTypeU"`

	// Address Type: - DEL for Delivery - COL for Collection
	Purpose formats.AlphaNumericString_Length1To3 `xml:"purpose,omitempty"`
}

type AgreementIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AgreementIdentificationTypeU"`

	// Agreement identification
	AgreementDetails *AgreementTypeIdentificationTypeU `xml:"agreementDetails,omitempty"`
}

type AgreementTypeIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AgreementTypeIdentificationTypeU"`

	// - TK for Tracking option
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`

	// Agreement description
	Description formats.AlphaNumericString_Length1To25 `xml:"description,omitempty"`
}

type ApplicationErrorDetailTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ApplicationErrorDetailTypeI"`

	// Message number or "ZZZ" if no number
	ErrorCode formats.AlphaNumericString_Length1To5 `xml:"errorCode,omitempty"`

	// EC for Error codes  WEC for Warning code  INF for Information code
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// 3 for IATA   UN for UN  1A for AMADEUS
	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`
}

type ApplicationErrorDetailType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ApplicationErrorDetailType"`

	// Code identifying the data validation error condition.
	ErrorCode formats.AlphaNumericString_Length1To5 `xml:"errorCode,omitempty"`

	// Identification of a code list.
	ErrorCategory formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`
}

type ApplicationErrorDetailType_198235C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ApplicationErrorDetailType_198235C"`

	// Code identifying the data validation error condition.
	ErrorCode formats.AlphaNumericString_Length1To5 `xml:"errorCode,omitempty"`

	// Identification of a code list.
	ErrorCategory formats.AlphaNumericString_Length1To3 `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner formats.AlphaNumericString_Length1To3 `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ApplicationErrorInformationType"`

	// Detail the error type
	ErrorDetail *ApplicationErrorDetailTypeI `xml:"errorDetail,omitempty"`
}

type ApplicationErrorInformationType_136725S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ApplicationErrorInformationType_136725S"`

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType_198235C `xml:"errorDetails,omitempty"`
}

type ApplicationErrorInformationType_94519S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ApplicationErrorInformationType_94519S"`

	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails,omitempty"`
}

type ApplicationIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ApplicationIdentificationType"`

	// application internal identifier
	InternalId formats.AlphaNumericString_Length1To35 `xml:"internalId,omitempty"`

	// Item Version Number
	VersionNumber formats.AlphaNumericString_Length1To35 `xml:"versionNumber,omitempty"`
}

type ApplicationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ApplicationType"`

	// provides information on application identification
	ApplicationDetails *ApplicationIdentificationType `xml:"applicationDetails,omitempty"`
}

type AssociatedChargesInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AssociatedChargesInformationTypeI"`

	// This data element is used to identify the type of charge entered in the other fields.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// This data element is used to convey the amount of the supplementary charge.
	Amount formats.NumericDecimal_Length1To12 `xml:"amount,omitempty"`

	// To qualify the amount, can be - UNL (for unlimited mileage) when used for free mileage - 3 (for included in base rate) - 4 (for not included in base rate)
	Description formats.AlphaNumericString_Length1To20 `xml:"description,omitempty"`

	// This data element is used to specify the number of charge needed.
	NumberInParty formats.NumericInteger_Length1To4 `xml:"numberInParty,omitempty"`

	// This data element is used to convey the currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// This data element is used to convey the voucher text (in case of voucher).
	Comment formats.AlphaNumericString_Length1To70 `xml:"comment,omitempty"`
}

type AssociatedChargesInformationTypeI_198205C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AssociatedChargesInformationTypeI_198205C"`

	// This data element is used to identify the type of charge entered in the other fields.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Mileage charge amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// To qualify the amount, can be - UNL (for unlimited mileage) when used for free mileage - 3 (for included in base rate) - 4 (for not included in base rate)
	Description formats.AlphaNumericString_Length1To35 `xml:"description,omitempty"`

	// Quantity of free mileage
	NumberInParty formats.NumericInteger_Length1To15 `xml:"numberInParty,omitempty"`

	// Unit: - K Kilometer - M Miles
	PeriodType formats.AlphaNumericString_Length1To3 `xml:"periodType,omitempty"`

	// the currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Unstructured RG,RG and RQ rates.
	Comment formats.AlphaNumericString_Length1To70 `xml:"comment,omitempty"`
}

type AssociatedChargesInformationTypeI_198218C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AssociatedChargesInformationTypeI_198218C"`

	// - 045 Tax - 108 Surchage - COV Coverage - CPN Coupon
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Policy amount (coupon amount)
	Amount formats.NumericDecimal_Length1To9 `xml:"amount,omitempty"`

	// Qualifier: The possible values are: - IES included in Estimated Total - IBR included in Base Rate - OPT Optional - MAN Mandatory - NBR Not Included in Base Rate - ITX Policy amount Includes Tax - NTX Policy amount Not Includes Tax
	Description formats.AlphaNumericString_Length1To3 `xml:"description,omitempty"`

	// Maximum days
	NumberInParty formats.NumericInteger_Length1To15 `xml:"numberInParty,omitempty"`

	// 001 per day 002 per week 003 per month 004 per rental 012 tax percentage 013 no coupon value available
	PeriodType formats.AlphaNumericString_Length1To3 `xml:"periodType,omitempty"`

	// Policy amount currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Policy name
	Comment formats.AlphaNumericString_Length1To70 `xml:"comment,omitempty"`
}

type AssociatedChargesInformationTypeI_39535C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AssociatedChargesInformationTypeI_39535C"`

	// To specify the type of tax, the type of converted amount. It is coded on our side if not specifued by provider.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// to specify the tax in a foreign currency.
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// Tax name
	Description formats.AlphaNumericString_Length1To3 `xml:"description,omitempty"`

	// foreign currency.
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type AssociatedChargesInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AssociatedChargesInformationTypeU"`

	// Qualify the associated charge. For Tour, only "employee" is used to define a commission
	ChargeUnitCode formats.AlphaNumericString_Length1To3 `xml:"chargeUnitCode,omitempty"`

	// Value of the associated charge
	Amount formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`

	// Commission's percentage
	Percentage formats.NumericDecimal_Length1To10 `xml:"percentage,omitempty"`
}

type AttributeInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeInformationTypeU"`

	// Type of the authorization data.  Some of the possible types are:  25: (AUT) Context (Credit Mutuel) 26: (ATN) Customer instruction (Barclays) 27: (ATN) Cryptogram computation method (Credit Mutuel) 28: (AUT) Modified securisation mode (Credit Mutuel) 29: (ATN) Electronic commerce transaction type (Credit Mutuel) E: (ATN) Result of the secured payment VADS (Credit Mutuel)  MID: (AUT) Merchant ID
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType,omitempty"`

	// value of the data
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type AttributeInformationTypeU_198185C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeInformationTypeU_198185C"`

	// the attribute type LEI for leisure booking CLP for clip booking
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType,omitempty"`

	// Not Used
	AttributeDescription formats.AlphaNumericString_Length1To3 `xml:"attributeDescription,omitempty"`
}

type AttributeInformationTypeU_36633C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeInformationTypeU_36633C"`

	// This element is used to convey the service code of the service group of the ferry booking. The list of possible values depends of the Ferry provider.
	AttributeType formats.AlphaNumericString_Length1To7 `xml:"attributeType,omitempty"`
}

type AttributeInformationTypeU_45068C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeInformationTypeU_45068C"`

	// The list of possible values is: ADT Adult CHD Child FDC Diplomatic corps FEU Disabled FFM Family FFR Free FIR Inter rail FJO Journalist FSL School pupil INF Infant MIL Military NAT Nato official REC Child resident RES Resident SRC Senior citizen STU Student YTH Young person
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType,omitempty"`
}

type AttributeTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeTypeU"`

	// Describes the service type.
	AttributeFunction formats.AlphaNumericString_Length1To1 `xml:"attributeFunction,omitempty"`

	// Service details.
	AttributeDetails *AttributeInformationTypeU_36633C `xml:"attributeDetails,omitempty"`
}

type AttributeTypeU_24552S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeTypeU_24552S"`

	// provides the function of the attribute
	AttributeFunction formats.AlphaNumericString_Length1To3 `xml:"attributeFunction,omitempty"`

	// provides details for the Attribute
	AttributeDetails *AttributeInformationTypeU_45068C `xml:"attributeDetails,omitempty"`
}

type AttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeType"`

	// Specify which attribute is described in E003. BAT for booking attribute
	CriteriaSetType formats.AlphaNumericString_Length1To3 `xml:"criteriaSetType,omitempty"`

	// Details for the attribute type. LEI:Y for leisure booking CLP:Y for clip booking
	CriteriaDetails *AttributeInformationTypeU_198185C `xml:"criteriaDetails,omitempty"`
}

type AttributeType_94514S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeType_94514S"`

	// Determines if the set of criteria corresponds to the message identification criteria or to normal criteria.
	CriteriaSetType formats.AMA_EDICodesetType_Length1to3 `xml:"criteriaSetType,omitempty"`

	// List of attributes and status linked to credit card process. Most of them are link dependant.
	CriteriaDetails *AttributeInformationTypeU `xml:"criteriaDetails,omitempty"`
}

type AttributeType_94553S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeType_94553S"`

	// Type of Data Exple :  SAL sale indicator EXT for extended payment PAY payment type
	CriteriaSetType formats.AlphaNumericString_Length1To3 `xml:"criteriaSetType,omitempty"`

	// Details for the message criteria (name, value).
	CriteriaDetails *AttributeInformationTypeU `xml:"criteriaDetails,omitempty"`
}

type AttributeType_94576S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AttributeType_94576S"`

	// Type of information: - is this a switch? - is this a structured data?
	CriteriaSetType formats.AlphaNumericString_Length1To3 `xml:"criteriaSetType,omitempty"`

	// Details for the message criteria (name, value).
	CriteriaDetails *AttributeInformationTypeU `xml:"criteriaDetails,omitempty"`
}

type AuthenticationDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AuthenticationDataType"`

	// VERes status (enrollment) Values : Y : authentication available N : cardholder not participating U : Unable to authenticate E : error message
	Veres formats.AlphaString_Length1To1 `xml:"veres,omitempty"`

	// PARes status (authentication). Values : Y : authentication successful N : authentication failed U : authentication could not be performed A : attempts processing performed
	Pares formats.AlphaString_Length1To1 `xml:"pares,omitempty"`

	// CC Directory Server performing the enrollment process: VISA, MasterCard
	CreditCardCompany formats.AlphaNumericString_Length4To4 `xml:"creditCardCompany,omitempty"`

	// To indicate whether the transaction was successful, different indicators for Visa/MasterCard. - ECI for VISA - UCAF collection indicator for Matercard
	AuthenticationIndicator formats.AlphaNumericString_Length2To2 `xml:"authenticationIndicator,omitempty"`

	// Indicates the algorithm used to generate the Cardholder Authentication Verification Value (CAAV = authentication code)
	CaavAlgorithm formats.NumericInteger_Length1To1 `xml:"caavAlgorithm,omitempty"`
}

type AuthorizationApprovalDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A AuthorizationApprovalDataType"`

	// will convey the value of the approval code of the payment authorisation
	ApprovalCode formats.AlphaNumericString_Length1To12 `xml:"approvalCode,omitempty"`

	// Source of approval for the payment authorisation. A Automatically obtained by the system. M Manually entered by an agent.  F: Credit card automatic approval code of a settlement authorization transaction B: Credit card manual approval code of a settlement transaction.
	SourceOfApproval formats.AlphaNumericString_Length1To3 `xml:"sourceOfApproval,omitempty"`
}

type BillableInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A BillableInformationTypeU"`

	// This composite is used to convey the billable information.
	BillingInfo *DiagnosisTypeU `xml:"billingInfo,omitempty"`
}

type BinaryDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A BinaryDataType"`

	// Length of the BLB
	DataLength formats.NumericInteger_Length1To15 `xml:"dataLength,omitempty"`

	// type of the data
	DataType formats.AlphaNumericString_Length1To1 `xml:"dataType,omitempty"`

	// used to store binary data
	BinaryData formats.AlphaNumericString_Length1To99999 `xml:"binaryData,omitempty"`
}

type BrowserInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A BrowserInformationType"`

	// Indicates the type of cardholder device.
	DeviceCategory formats.NumericInteger_Length1To3 `xml:"deviceCategory,omitempty"`
}

type CabinClassDesignationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CabinClassDesignationType"`

	// Designates the class of service on the means of transport  in which the passenger will travel:  - M for Economy - W for Economy Premium - C for Business (Club) - F for First  - Y for Economy All
	ClassDesignator formats.AlphaString_Length1To1 `xml:"classDesignator,omitempty"`
}

type CabinDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CabinDetailsType"`

	CabinDetails *CabinClassDesignationType `xml:"cabinDetails,omitempty"`
}

type CardValidityType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CardValidityType"`

	// Type of the compensation, ie voluntary or involuntary
	Type formats.AlphaNumericString_Length1To1 `xml:"type,omitempty"`

	// Form of the payment of the compensation
	Form formats.AlphaNumericString_Length1To1 `xml:"form,omitempty"`

	// Amount of the compensation
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// Currency used for the compensation.
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Any comment related to the compensation
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type ChildrenGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ChildrenGroupType"`

	// This segment is used to convey age for a child.
	Age *QuantityTypeI_65488S `xml:"age,omitempty"`

	// This segment is used to convey the passenger association
	ReferenceForPassenger *ReferenceInformationType_65487S `xml:"referenceForPassenger,omitempty"`
}

type ClassConfigurationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ClassConfigurationDetailsType"`

	// Class Details -Class Group -Sub Class -number of seats
	ClassDetails *ClassDetailsType `xml:"classDetails,omitempty"`
}

type ClassDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ClassDetailsType"`

	// Class Group : A-First Class, Seat B-Second Class, Seat C-First Class, Berth D-Second Class, Berth F-Binded Seat V-First Class, Sleeping-car W-Second Class, Sleeping-car
	Code formats.AlphaNumericString_Length1To1 `xml:"code,omitempty"`

	// Rail class code.
	BookingClass formats.AlphaNumericString_Length2To2 `xml:"bookingClass,omitempty"`

	// Number of Free Seats
	NumberOfSeats formats.NumericInteger_Length1To3 `xml:"numberOfSeats,omitempty"`
}

type ClassDetailsType_52782C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ClassDetailsType_52782C"`

	// For the booking class code.
	Code formats.AlphaNumericString_Length1To1 `xml:"code,omitempty"`

	BookingClass formats.AlphaNumericString_Length2To2 `xml:"bookingClass,omitempty"`
}

type CoachProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CoachProductInformationType"`

	// Coach Product Details
	CoachDetails *ReferencingDetailsTypeI_36941C `xml:"coachDetails,omitempty"`

	// Coach Equipment Qualifier
	EquipmentCode formats.AlphaNumericString_Length1To1 `xml:"equipmentCode,omitempty"`
}

type CodedAttributeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CodedAttributeInformationType"`

	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`
}

type CodedAttributeInformationType_142109C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CodedAttributeInformationType_142109C"`

	// will convey the following QF data:  ONOD onoData     Order Number(Qantas specific)  GWTD gwtData     Government Warrant number(Qantas specific)  HOLDN ccHolderName    Conveys the CN (company name) (Qantas specific).This has sense only in case of automatic creation of attribute record (cards is a bets card). In the other cases this information cannot be filled.  ONOR onoRequired    This ONO indicator indicates whether or not ONO data is mandatory.(Information provided by Qantas IGW link) This has sense only in case of automatic creation of attribute record (cards is a bets card)  GWTR gwtRequired     This GWT indicator indicates whether or not GWT data is mandatory.(Information provided by Qantas IGW link) This has sense only in case of automatic creation of attribute record (cards is a bets card)  CIND cind     Conveys CIND indicator : - MANU - AUTO This indicates whether attributes records have been created manually (no bets card) or automatically (bets card).  BFAREC bestFareCandidate     Conveys best fare indicator: - Yes - No This indicates whether or not the card is best fare candidate. (this implies card is a bets card)
	AttributeType formats.AlphaNumericString_Length1To5 `xml:"attributeType,omitempty"`

	// onoData   Order Number(Qantas specific)  gwtData   Government Warrant number  ccHolderName  Conveys the CN   onoRequired  YES - NO  gwtRequired   YES - NO  cind    MANU - AUTO  bestFareCandidate   YES - NO
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeInformationType_185753C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CodedAttributeInformationType_185753C"`

	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType,omitempty"`

	// provides a description for the attribute: If 950K set to ARC, value can be: AS: Airline Staff, BA: Baggage not Accepted, CB: Customer failed to board, CR: Customer Request, CU: Customer Unwell, DB: Denied Boarding, FA: Flight Alternative Offered and accepted by the customer, FD: Flight Delayed, FO: Flight Oversold, FC: Flight Cancelled, MC: Missed Connection, MR: Medical Reasons, NO: NOSHOW (can only be used with Target Customer Acceptance Status set to 'Rejected'), OT: Other, RR: Regulatory Requirement not met, SR: Security Reasons, TC:  Travel in different cabin through another booking, TD: Travel Documentation Incomplete, TI: Travel Industry Staff, UC: User Error Correction. If 950K set to RRC, value can be: OC: Cabin/Flight oversold (current flight), OO: Cabin/Flight oversold (other flight), MC: Misidentification of customer at check-in, PM: Previously mishandled, SO: Special occasion, AC: Aircraft change, CF: Cabin configuration change, RQ: Request from a special requestor, CO: Compassionate, MI: Marketing Initiative e.g. please try Club Class for free, DC: Disruption on current flight, DO: Disruption on other flight, CC: Crew level change, TR: Technical reason, CS: Catering shortfall, CI: Check-in error, IA: Inadmissible, ST: Staff, SR: Service Recovery Entitlement, AB: Authorized by, US: Unsuitable.
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CodedAttributeType"`

	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`
}

type CodedAttributeType_127279S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CodedAttributeType_127279S"`

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType_185753C `xml:"attributeDetails,omitempty"`
}

type CodedAttributeType_127282S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CodedAttributeType_127282S"`

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails,omitempty"`
}

type CodedAttributeType_94497S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CodedAttributeType_94497S"`

	// Usage of this element will be the transport of the DescriptiveBilingInformation first value will be QF
	AttributeFunction formats.AlphaNumericString_Length1To3 `xml:"attributeFunction,omitempty"`

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType_142109C `xml:"attributeDetails,omitempty"`
}

type CodeshareFlightDataTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CodeshareFlightDataTypeI"`

	// Company identification
	Airline formats.AlphaNumericString_Length1To3 `xml:"airline,omitempty"`

	// Product identification
	FlightNumber formats.NumericInteger_Length0To4 `xml:"flightNumber,omitempty"`

	// general indicator
	Inventory formats.AlphaNumericString_Length0To3 `xml:"inventory,omitempty"`

	// Characteristic identification
	SellingClass formats.AlphaString_Length1To1 `xml:"sellingClass,omitempty"`

	// Item characteristic
	Type formats.AlphaNumericString_Length1To2 `xml:"type,omitempty"`

	// Product identification characteristic
	Suffix formats.AlphaString_Length1To1 `xml:"suffix,omitempty"`

	// 1 :  codeshare cascading is successful 0 : codeshare cascading unsuccessful blank: n/s
	CascadingIndicator formats.NumericInteger_Length1To1 `xml:"cascadingIndicator,omitempty"`
}

type CommissionDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CommissionDetailsType"`

	// Commission type : 'NEW' --) New commission 'OLD' --) Old Commission 'XLP' --) Commission on cancellation Penalty 'FMA' --) Airline Commission A 'FMB' --) Airline Commission B
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Commission amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Commission percentage
	Rate formats.NumericInteger_Length1To8 `xml:"rate,omitempty"`

	// Deal number
	DealNumber formats.NumericInteger_Length1To8 `xml:"dealNumber,omitempty"`
}

type CommissionInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CommissionInformationType"`

	// Commission details
	CommissionDetails *CommissionDetailsType `xml:"commissionDetails,omitempty"`
}

type CommunicationContactDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CommunicationContactDetailsTypeU"`

	// Email
	Email formats.AlphaNumericString_Length1To70 `xml:"email,omitempty"`

	// Contact qualifier. EM for Electronic mail
	ContactQualifier formats.AlphaNumericString_Length1To3 `xml:"contactQualifier,omitempty"`
}

type CommunicationContactDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CommunicationContactDetailsType"`

	// The communication address: an Url.
	UrlAddress formats.AlphaNumericString_Length1To512 `xml:"urlAddress,omitempty"`

	// will be AH for World Wide Web
	UrlType formats.AlphaNumericString_Length1To3 `xml:"urlType,omitempty"`
}

type CommunicationContactTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CommunicationContactTypeU"`

	// Customer contact
	Contact *CommunicationContactDetailsTypeU `xml:"contact,omitempty"`
}

type CommunicationContactType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CommunicationContactType"`

	// Communication channel
	Communication *CommunicationContactDetailsType `xml:"communication,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyIdentificationTypeI"`

	OperatingCompany formats.AlphaNumericString_Length1To3 `xml:"operatingCompany,omitempty"`
}

type CompanyIdentificationTypeI_2785C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyIdentificationTypeI_2785C"`

	// Company code
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification,omitempty"`

	// AIR segment : second airline code for joint flight number. Train Amtrack segment : system provider UIC code. Train SNCF segment : system provider UIC code. Tour segment : tour provider code.
	SecondIdentification formats.AlphaNumericString_Length1To4 `xml:"secondIdentification,omitempty"`

	// Tour segment : source code.
	SourceCode formats.AlphaNumericString_Length1To4 `xml:"sourceCode,omitempty"`
}

type CompanyIdentificationTypeI_46335C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyIdentificationTypeI_46335C"`

	// carrier details
	MarketingCompany formats.AlphaNumericString_Length1To24 `xml:"marketingCompany,omitempty"`
}

type CompanyIdentificationTypeI_46351C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyIdentificationTypeI_46351C"`

	// Targeted provider system code
	OperatingCompany formats.AlphaNumericString_Length1To2 `xml:"operatingCompany,omitempty"`
}

type CompanyIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyIdentificationTypeU"`

	// Conveys the provider name
	ProviderName formats.AlphaNumericString_Length1To20 `xml:"providerName,omitempty"`
}

type CompanyInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyInformationType"`

	// Qualify the company code, to identify the industry business it belongs.
	TravelSector formats.AlphaNumericString_Length1To3 `xml:"travelSector,omitempty"`

	// This data element is used to convey the context in which the code applies. The hotel chain code are managed by Amadeus.
	CompanyCodeContext formats.AlphaNumericString_Length1To3 `xml:"companyCodeContext,omitempty"`

	// This data element is used to convey the company code of a non-air company
	CompanyCode formats.AlphaNumericString_Length1To3 `xml:"companyCode,omitempty"`

	// This data element is used to convey the company name of a non-air company
	CompanyName formats.AlphaNumericString_Length1To35 `xml:"companyName,omitempty"`

	// This data element is used to convey the type of access the non-air company has with Amadeus.
	AccessLevel formats.AlphaNumericString_Length1To3 `xml:"accessLevel,omitempty"`
}

type CompanyInformationType_19450S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyInformationType_19450S"`

	// This data element is used to convey the company code
	CompanyCode formats.AlphaNumericString_Length2To3 `xml:"companyCode,omitempty"`

	// This data element is used to convey the UIC code
	CompanyNumericCode formats.NumericInteger_Length2To4 `xml:"companyNumericCode,omitempty"`
}

type CompanyInformationType_20151S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyInformationType_20151S"`

	// This element is used to qualify the company code, to identify the industry business it belongs. For Ferry, the codes are mutually agreed between Amadeus and the Ferry providers and are only valid within the Amadeus Ferry application.
	TravelSector formats.AlphaNumericString_Length1To3 `xml:"travelSector,omitempty"`

	// This data element is used to convey the company code of a company. For Ferry, the list of providers is not fixed. The providers implemented actually are: Baleria                 BAL Brittany ferries        BRI Color Lines             COL Comarit                 COM Corsica ferries         CSF Grandi Navi Veloci      GNV Hoverspeed              HOV Moby Lines              MBL Seafrance               SEA Smyril Line             SMY SNCM                    SNC Stena Line              STE TT Line                 TTL
	CompanyCode formats.AlphaNumericString_Length3To3 `xml:"companyCode,omitempty"`

	// This data element is used to convey the company name of a company
	CompanyName formats.AlphaNumericString_Length1To35 `xml:"companyName,omitempty"`
}

type CompanyInformationType_25420S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyInformationType_25420S"`

	// This data element is used to qualify the company code, to identify the industry business it belongs.
	TravelSector formats.AlphaNumericString_Length1To3 `xml:"travelSector,omitempty"`

	// This data element is used to convey the company code of a company
	CompanyCode formats.AlphaNumericString_Length1To4 `xml:"companyCode,omitempty"`

	// This data element is used to convey the company name of a company
	CompanyName formats.AlphaNumericString_Length1To32 `xml:"companyName,omitempty"`
}

type CompanyInformationType_26258S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyInformationType_26258S"`

	// This data element is used to qualify the company code, to identify the industry business it belongs.
	TravelSector formats.AlphaNumericString_Length1To3 `xml:"travelSector,omitempty"`

	// This data element is used to convey the context in which the code applies
	CompanyCodeContext formats.AlphaNumericString_Length1To3 `xml:"companyCodeContext,omitempty"`

	// This data element is used to convey the company code of a company
	CompanyCode formats.AlphaNumericString_Length1To3 `xml:"companyCode,omitempty"`

	// This data element is used to convey the company name of a company
	CompanyName formats.AlphaNumericString_Length1To35 `xml:"companyName,omitempty"`
}

type CompanyInformationType_83550S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyInformationType_83550S"`

	// This data element is used to qualify the company code, to identify the industry business it belongs.
	TravelSector formats.AlphaNumericString_Length3To3 `xml:"travelSector,omitempty"`

	// This data element is used to convey the company code of a company
	CompanyCode formats.AlphaNumericString_Length3To3 `xml:"companyCode,omitempty"`
}

type CompanyInformationType_8953S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyInformationType_8953S"`

	// This data element is used to qualify the company code, to identify the industry business it belongs.
	TravelSector formats.AlphaNumericString_Length1To3 `xml:"travelSector,omitempty"`

	// This data element is used to convey the context in which the code applies
	CompanyCodeContext formats.AlphaNumericString_Length1To2 `xml:"companyCodeContext,omitempty"`

	// This data element is used to convey the company code of a non-air company
	CompanyCode formats.AlphaNumericString_Length1To3 `xml:"companyCode,omitempty"`

	// This data element is used to convey the company name of a non-air company
	CompanyName formats.AlphaNumericString_Length1To35 `xml:"companyName,omitempty"`
}

type CompanyInformationType_94554S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompanyInformationType_94554S"`

	// This data element is used to convey the company code of a company  Ex:  AF for Air France MIL for millenium foundation
	CompanyCode formats.AlphaNumericString_Length1To35 `xml:"companyCode,omitempty"`

	// This data element is used to convey the numeric merchant ID.
	CompanyNumericCode formats.AlphaNumericString_Length1To15 `xml:"companyNumericCode,omitempty"`
}

type CompensationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CompensationType"`

	// Compensation details
	CompensationDetails *CardValidityType `xml:"compensationDetails,omitempty"`
}

type ConsumerReferenceIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ConsumerReferenceIdentificationTypeI"`

	// Indicator - 1  for ID (customer number) - CD for CD (customer discount number)
	ReferenceQualifier formats.AlphaNumericString_Length1To3 `xml:"referenceQualifier,omitempty"`

	// Either the ID or CD number
	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber,omitempty"`
}

type ConsumerReferenceInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ConsumerReferenceInformationTypeI"`

	// Consumer reference information
	CustomerReferences *ConsumerReferenceIdentificationTypeI `xml:"customerReferences,omitempty"`
}

type ContactInformationDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ContactInformationDetailsTypeU"`

	// W for party to receive Written confirmation
	PartyQualifier formats.AlphaNumericString_Length1To3 `xml:"partyQualifier,omitempty"`

	// FAX number or E-Mail address
	ComAddress formats.AlphaNumericString_Length1To57 `xml:"comAddress,omitempty"`

	// type of medium
	ComChannelQualifier formats.AlphaNumericString_Length1To3 `xml:"comChannelQualifier,omitempty"`
}

type ContactInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ContactInformationTypeU"`

	// This composite is used to convey the E-mail address or FAX number to be used when a written confirmation is needed.
	ContactInformation *ContactInformationDetailsTypeU `xml:"contactInformation,omitempty"`
}

type CountryCodeListType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CountryCodeListType"`

	// ISO country code of the DESTINATION of the trip.
	DestinationCountryCode formats.AlphaNumericString_Length1To2 `xml:"destinationCountryCode,omitempty"`
}

type CountrySubEntityDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CountrySubEntityDetailsTypeU"`

	// 84: state
	Qualifier formats.AlphaNumericString_Length1To2 `xml:"qualifier,omitempty"`

	// Region or State of the given address
	Name formats.AlphaNumericString_Length1To25 `xml:"name,omitempty"`
}

type CountrySubEntityDetailsTypeU_198213C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CountrySubEntityDetailsTypeU_198213C"`

	// State Code. Mandatory if CountryCode is US, CA, AU.
	Code formats.AlphaNumericString_Length1To2 `xml:"code,omitempty"`
}

type CountrySubEntityDetailsTypeU_198229C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CountrySubEntityDetailsTypeU_198229C"`

	// State code
	Code formats.AlphaNumericString_Length1To9 `xml:"code,omitempty"`
}

type CountrydescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CountrydescriptionType"`

	// To specify the destination zone.
	GeographicalZone formats.AlphaNumericString_Length1To35 `xml:"geographicalZone,omitempty"`

	// To specify the countries but in a coded way. up to 198 repetitions as we can have 99 segments in the PNR
	CountryCode formats.AlphaNumericString_Length1To2 `xml:"countryCode,omitempty"`
}

type CreditCardDataGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CreditCardDataGroupType"`

	// will convey all the data related to the credit card
	CreditCardDetails *CreditCardDataType `xml:"creditCardDetails,omitempty"`

	// will convey both the CVV and the Credit card number Ids stored in the fortknox Database  it could also be used to store identifiers from external Tokenization Service Provider (TSP).
	FortknoxIds *ReferenceInformationTypeI_94503S `xml:"fortknoxIds,omitempty"`

	// Contains card holder's address information.
	CardHolderAddress *AddressType `xml:"cardHolderAddress,omitempty"`
}

type CreditCardDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CreditCardDataType"`

	// Credit Card information
	CcInfo *CreditCardInformationType `xml:"ccInfo,omitempty"`
}

type CreditCardInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CreditCardInformationTypeU"`

	// Credit card name
	Name formats.AlphaNumericString_Length1To2 `xml:"name,omitempty"`

	// Credit card number
	CardNumber formats.NumericInteger_Length1To20 `xml:"cardNumber,omitempty"`

	// Credit card Expire date
	ExpireDate formats.AlphaNumericString_Length1To6 `xml:"expireDate,omitempty"`
}

type CreditCardInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CreditCardInformationType"`

	// Vendor code (VI,CA,AX.)
	VendorCode formats.AlphaNumericString_Length2To2 `xml:"vendorCode,omitempty"`

	// may contain CC sub Types. eg: Maestro or Solo cards
	VendorCodeSubType formats.AlphaNumericString_Length1To25 `xml:"vendorCodeSubType,omitempty"`

	// Card number  Card number
	CardNumber formats.AlphaNumericString_Length1To19 `xml:"cardNumber,omitempty"`

	// Conveys the security ID of the Credit Card (CVV,CVV2), 3-4 digits stored on the back of the card
	SecurityId formats.AlphaNumericString_Length1To4 `xml:"securityId,omitempty"`

	// Expiry date :  format    MMYY
	ExpiryDate formats.NumericInteger_Length4To4 `xml:"expiryDate,omitempty"`

	// This field indicates the date the Credit Card was issued. This data is present in case of (UK) maestro cards.
	StartDate formats.NumericInteger_Length4To4 `xml:"startDate,omitempty"`

	// This field indicates the date the Credit Card will not be valid anymore This data is present in case of (UK) maestro cards. May be different from the expiry date
	EndDate formats.NumericInteger_Length4To4 `xml:"endDate,omitempty"`

	// Conveys Credit card holder's name, as written on the card
	CcHolderName formats.AlphaNumericString_Length1To99 `xml:"ccHolderName,omitempty"`

	// will contain the code of the bank that issued the credit card
	IssuingBankName formats.AlphaNumericString_Length2To3 `xml:"issuingBankName,omitempty"`

	// CC country of issuance details
	CardCountryOfIssuance formats.AlphaNumericString_Length1To3 `xml:"cardCountryOfIssuance,omitempty"`

	// This is the Credit Card Issue number. This represents the number of time a card has been issued.  1 is for the first time then in case of card renewal or card loss this issue number will be increased Today this is applicable to maestro cards.
	IssueNumber formats.NumericInteger_Length1To3 `xml:"issueNumber,omitempty"`

	// Will convey the full name of the institution that issued he credit card
	IssuingBankLongName formats.AlphaNumericString_Length1To64 `xml:"issuingBankLongName,omitempty"`

	// Stores the CC track 1 information (base64 encoded)
	Track1 formats.AlphaNumericString_Length0To108 `xml:"track1,omitempty"`

	// Stores the CC track 2 information (base64 encoded)
	Track2 formats.AlphaNumericString_Length0To56 `xml:"track2,omitempty"`

	// Stores the CC track 3 information (base64 encoded)
	Track3 formats.AlphaNumericString_Length0To144 `xml:"track3,omitempty"`

	// Stores the CC pin code information
	PinCode formats.AlphaNumericString_Length1To100 `xml:"pinCode,omitempty"`

	// All the tracks of a swipe credit card are contained here as one block.
	RawTrackData formats.AlphaNumericString_Length1To400 `xml:"rawTrackData,omitempty"`
}

type CreditCardSecurityType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CreditCardSecurityType"`

	// Conveys all data of authentication process. Only used today for "Verified by Visa" process
	AuthenticationDataDetails *AuthenticationDataType `xml:"authenticationDataDetails,omitempty"`
}

type CreditCardStatusGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CreditCardStatusGroupType"`

	// This segment is used to store specific data of links following ISO8583 standard.
	AuthorisationSupplementaryData *SpecificVisaLinkCreditCardInformationType `xml:"authorisationSupplementaryData,omitempty"`

	// will convey the approval code/source
	ApprovalDetails *GenericAuthorisationResultType `xml:"approvalDetails,omitempty"`

	// This segment conveys date and time information.  You can specify the time mode used (GMT, UTC or Local)and what for it refers.  - Transmission date and time This contains the date and time the request was submitted to the link (Visa, Nedbank...field 7). GMT can be used.  - Local transaction date and time Date and time when Amadeus builds the authorization message (local according to the point of sale)(Visa, Nedbank...field 12/13)  - Transaction receipt date and time date and time when amadeus receives the authorization message.
	LocalDateTime *StructuredDateTimeInformationType_94516S `xml:"localDateTime,omitempty"`

	// Transaction Information: - type of authorization message submit for the given FOP - bulk, superbulk, no bulk process - STAN number (identifying a pair of Credit Card authorization request/response).
	AuthorisationInformation *TransactionInformationForTicketingType `xml:"authorisationInformation,omitempty"`

	BrowserData *BrowserData `xml:"browserData,omitempty"`

	// this group will convey all the 3DS related data
	TdsInformation *ThreeDomainSecureGroupType `xml:"tdsInformation,omitempty"`

	// This will allow the transmission of credit card data.
	CardSupplementaryData *AttributeType_94514S `xml:"cardSupplementaryData,omitempty"`

	// will convey the various sub status that can be associated to a credit card payment CVV, AVS, AUT, ATN....
	TransactionStatus *ErrorGroupType `xml:"transactionStatus,omitempty"`
}

type BrowserData struct {

	// This segment contains data about the customer's browser :  0 PC (HTML) 1 Mobile Internet Device (WML)
	BrowserProperties *BrowserInformationType `xml:"browserProperties,omitempty"`

	// Contains in freeflow format data about the customer's browser. - userAgent - acceptHeaders This entities are independantly optional.
	FreeFlowBrowserData *FreeTextInformationType_94526S `xml:"freeFlowBrowserData,omitempty"`
}

type CreditCardType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CreditCardType"`

	// credit card company code
	CreditCardCompany formats.AlphaNumericString_Length1To3 `xml:"creditCardCompany,omitempty"`

	// credit card number
	CreditCardNumber formats.AlphaNumericString_Length1To20 `xml:"creditCardNumber,omitempty"`

	// expiration date
	ExpirationDate formats.NumericInteger_Length4To4 `xml:"expirationDate,omitempty"`
}

type CruiseBusinessDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CruiseBusinessDataType"`

	// Details of sailing ship for the sailing trip.  Each cruise provider has a ship name table in the Amadeus system. This table is used for converting ship codes in ship names and vice-versa. Since both information are stored in the Cruise segment of the PNR, no DB access is  necessary for the PNRACC processing.
	SailingShipInformation *ShipIdentificationType_8952S `xml:"sailingShipInformation,omitempty"`

	// Details of the cruise line provider for the sailing trip.
	SailingProviderInformation *CompanyInformationType_8953S `xml:"sailingProviderInformation,omitempty"`

	// Details of embarkation and disembarkation ports for the sailing trip.  The codes sent by the cruise providers can be non-Iata codes.
	SailingPortsInformation *PlaceLocationIdentificationTypeU_8954S `xml:"sailingPortsInformation,omitempty"`

	// Details of the departure and arrival dates of the sailing trip.  The cruise segment in the PNR actually stores the departure date and the duration length in days. For the PNRACC 4.1 process, the arrival date is re-calculated.
	SailingDateInformation *StructuredPeriodInformationType_8955S `xml:"sailingDateInformation,omitempty"`

	// Details of passengers for the sailing trip.  For a cruise booking, the passenger names elements from the PNR can be different from the passengers in the cruise segment. They are identical at booking creation time. But the cruise providers allow adding passenger name(s) to an existing booking. That is not possible in an Amadeus PNR. Therefore, the name information had to be stored in the cruise segment itself.
	PassengerInfo *TravellerInformationTypeI_8956S `xml:"passengerInfo,omitempty"`

	BookingDetails *BookingDetails `xml:"bookingDetails,omitempty"`

	// Booking Date.
	BookingDate *StructuredDateTimeInformationType_20645S `xml:"bookingDate,omitempty"`

	// Details of the sailing group code for the sailing trip.
	SailingGroupInformation *ItemReferencesAndVersionsType_9271S `xml:"sailingGroupInformation,omitempty"`
}

type BookingDetails struct {

	// Details of the booking references for the sailing trip.  These references are returned by the cruise provider at booking creation time or at booking cancellation time.  Note that as re-instate of a cruise booking is possible even several days after cancellation. Therefore, when a cruise booking is cancelled, the segment is kept in the PNR and the status updated to HX.
	CruiseBookingReferenceInfo *ReservationControlInformationTypeI_8957S `xml:"cruiseBookingReferenceInfo,omitempty"`

	// Company in which the booking is created: Amadeus or external.
	BookingCompany *CompanyInformationType_26258S `xml:"bookingCompany,omitempty"`
}

type CustomerTransactionDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A CustomerTransactionDataType"`

	// Point of sell details
	Pos *PointOfSaleDataTypeI `xml:"pos,omitempty"`

	// flight suplementary data
	Flight *OtherSegmentDataTypeI `xml:"flight,omitempty"`

	// CONNECTION NUMBER
	Connection formats.NumericInteger_Length1To2 `xml:"connection,omitempty"`

	// Codeshare flight details
	CodeShare *CodeshareFlightDataTypeI `xml:"codeShare,omitempty"`
}

type DataInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DataInformationTypeI"`

	// Animal type.
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Number of animals of the specified category.
	Value formats.NumericInteger_Length1To2 `xml:"value,omitempty"`
}

type DataTypeInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DataTypeInformationTypeI"`

	// Type of data.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type DateAndTimeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DateAndTimeDetailsTypeI"`

	// Seat SSR : Date of change of gauge. Group seat SSR  : Date of change of gauge. MCO element : Date.
	FirstDate formats.Date_DDMMYY `xml:"firstDate,omitempty"`

	// MCO element : ARC carrier code code.
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`

	// MCO element : ARC city code.
	LocationIdentification formats.AlphaNumericString_Length1To3 `xml:"locationIdentification,omitempty"`
}

type DateAndTimeDetailsTypeI_56946C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DateAndTimeDetailsTypeI_56946C"`

	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Inf/Child date of birth
	Date formats.AlphaNumericString_Length1To8 `xml:"date,omitempty"`
}

type DateAndTimeInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DateAndTimeInformationTypeI"`

	// Date and Time details for flight movements
	DateAndTime *DateAndTimeDetailsTypeI `xml:"dateAndTime,omitempty"`
}

type DateAndTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DateAndTimeInformationType"`

	// DATE AND TIME DETAILS
	DateAndTimeDetails *DateAndTimeDetailsTypeI_56946C `xml:"dateAndTimeDetails,omitempty"`
}

type DateRangeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DateRangeType"`

	// In range [1-4]
	DateRangeNum formats.NumericInteger_Length1To1 `xml:"dateRangeNum,omitempty"`
}

type DetailedPaymentDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DetailedPaymentDataType"`

	// This segment will convey the type of the FOP. Exple : CC credit card CA cash CH cheque WW web
	FopInformation *FormOfPaymentType `xml:"fopInformation,omitempty"`

	// will allow the usage of FOP segment as trigger for GASS and GIVR groups
	Dummy *DummySegmentTypeI `xml:"dummy,omitempty"`

	// This group will convey the detailed status of the credit card payment
	CreditCardDetailedData *CreditCardStatusGroupType `xml:"creditCardDetailedData,omitempty"`
}

type DeviceControlDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DeviceControlDetailsType"`

	// Stores the identification of the device.
	DeviceIdentification *IdentificationNumberTypeI `xml:"deviceIdentification,omitempty"`
}

type DiagnosisTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DiagnosisTypeU"`

	// This data element can convey either an agency accounting or a billing number.
	BillingDetails formats.AlphaNumericString_Length1To25 `xml:"billingDetails,omitempty"`

	// This data element is used to specify the type of billable information that could be found in this segment
	BillingQualifier formats.AlphaNumericString_Length1To3 `xml:"billingQualifier,omitempty"`
}

type DiningIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DiningIdentificationType"`

	// meal plan information (HALFBOARD, BREAKFAST ...)
	DiningDescription formats.AlphaNumericString_Length1To16 `xml:"diningDescription,omitempty"`
}

type DiningInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DiningInformationType"`

	// Conveys dining information
	DiningIdentification *DiningIdentificationType `xml:"diningIdentification,omitempty"`
}

type DiscountInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DiscountInformationDetailsType"`

	// Promotion code used to define redemption/upgrade price in miles
	DiscountCode formats.AlphaNumericString_Length1To6 `xml:"discountCode,omitempty"`
}

type DiscountInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DiscountInformationType"`

	// Contains the discount code
	DiscountDetails *DiscountInformationDetailsType `xml:"discountDetails,omitempty"`
}

type DistributionChannelType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DistributionChannelType"`

	// This field is used to indicate the type of channel used for authorization process: e-commerce (web / Internet), MOTO (Mail Order / telephone Order), Face to face ...) Example:  05 for API 0=MOTO (Mail Order / Telephone Order) 1=e-Commerce (Internet)
	DistributionChannelField formats.NumericInteger_Length1To3 `xml:"distributionChannelField,omitempty"`

	// Subgroup field.
	SubGroup formats.NumericInteger_Length1To3 `xml:"subGroup,omitempty"`

	// Access Type.
	AccessType formats.NumericInteger_Length1To3 `xml:"accessType,omitempty"`
}

type DocumentDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DocumentDetailsTypeI"`

	// To convey the document number
	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`

	// To convey if the document has been printed
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// To convey the date of the impression.
	Date formats.AlphaNumericString_Length1To8 `xml:"date,omitempty"`
}

type DocumentDetailsTypeI_19732C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DocumentDetailsTypeI_19732C"`

	// documentNumber
	Number formats.NumericInteger_Length10To10 `xml:"number,omitempty"`

	// Status Code
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`
}

type DocumentDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DocumentDetailsType"`

	// Document type: PT for passport, VS for visa.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Passport number.
	Number formats.AlphaNumericString_Length1To20 `xml:"number,omitempty"`

	// Country where the document has been issued.
	CountryOfIssue formats.AlphaNumericString_Length1To3 `xml:"countryOfIssue,omitempty"`

	// Expiry date of the document. YYYYMMDD
	ExpiryDate formats.AlphaNumericString_Length1To8 `xml:"expiryDate,omitempty"`

	// Date of issue of the document. YYYYMMDD
	IssueDate formats.AlphaNumericString_Length1To8 `xml:"issueDate,omitempty"`
}

type DocumentInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DocumentInformationDetailsTypeI"`

	// To convey the printing results.
	DocumentDetails *DocumentDetailsTypeI `xml:"documentDetails,omitempty"`
}

type DocumentInformationDetailsTypeI_9936S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DocumentInformationDetailsTypeI_9936S"`

	// documentDetails
	DocumentDetails *DocumentDetailsTypeI_19732C `xml:"documentDetails,omitempty"`
}

type DocumentInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DocumentInformationTypeU"`

	// Document type being provided: PP: Passport DL: Driving License NI: National Id. card ID: Local Id. Document
	TypeOfDocument formats.AlphaNumericString_Length1To3 `xml:"typeOfDocument,omitempty"`

	// Document Number
	DocumentNumber formats.AlphaNumericString_Length1To20 `xml:"documentNumber,omitempty"`

	// Country code where document has been issued
	CountryOfIssue formats.AlphaNumericString_Length1To3 `xml:"countryOfIssue,omitempty"`
}

type DummySegmentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A DummySegmentTypeI"`
}

type ElementManagementSegmentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ElementManagementSegmentType"`

	// Action to perform (When a PNR segment/element is transmitted) .  IF for Information only (Value by default, Code used in a Server response)
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Reference details
	Reference *ReferencingDetailsType_127526C `xml:"reference,omitempty"`

	// PNR segment or element name
	SegmentName formats.AlphaNumericString_Length1To3 `xml:"segmentName,omitempty"`

	// PNR segment/element 'line' number attributed by the Server
	LineNumber formats.NumericInteger_Length1To3 `xml:"lineNumber,omitempty"`
}

type ElementManagementSegmentType_127983S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ElementManagementSegmentType_127983S"`

	// reference of the element
	ElementReference *ReferencingDetailsType_127526C `xml:"elementReference,omitempty"`

	// PNR segment or element name
	SegmentName formats.AlphaNumericString_Length1To3 `xml:"segmentName,omitempty"`

	// PNR segment/element 'line' number.
	LineNumber formats.NumericInteger_Length1To3 `xml:"lineNumber,omitempty"`
}

type ElementManagementSegmentType_83559S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ElementManagementSegmentType_83559S"`

	// Reference details
	Reference *ReferencingDetailsType_127526C `xml:"reference,omitempty"`
}

type EquipmentDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A EquipmentDetailsTypeU"`

	// equipment type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// equipment details
	SizeTypeDetails *EquipmentTypeAndSizeTypeU `xml:"sizeTypeDetails,omitempty"`
}

type EquipmentTypeAndSizeTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A EquipmentTypeAndSizeTypeU"`

	// equipment description
	Description formats.AlphaNumericString_Length1To35 `xml:"description,omitempty"`
}

type ErrorGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ErrorGroupType"`

	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType_94519S `xml:"errorOrWarningCodeDetails,omitempty"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType_94495S `xml:"errorWarningDescription,omitempty"`
}

type FLIXType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FLIXType"`

	// provides the Flix and Source Types. The Codes FX, LX or FD can be used to specify the Flix-Type. The codes USR or GUI can be used to specify the Data Source
	FlixAndSourceTypes *ItemDescriptionType `xml:"flixAndSourceTypes,omitempty"`

	FlixComment *FreeTextInformationType `xml:"flixComment,omitempty"`

	AirportGroup *AirportGroup `xml:"airportGroup,omitempty"`
}

type AirportGroup struct {

	// Only used for Flix-LX or Flix-Disruption: provides the code of the Impacted Airport
	ImpactedAirport *TerminalTimeInformationTypeS `xml:"impactedAirport,omitempty"`
}

type FOPRepresentationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FOPRepresentationType"`

	// will convey all the data related to the various codes used by the FOP package, billing, ETS...
	FopPNRDetails *TicketingFormOfPaymentType `xml:"fopPNRDetails,omitempty"`

	// Conveys the sequence number of the Form of Payment in the FP Line. It must be set to 1 if there is only 1 FOP in the FOP  Old FOP are referenced with sequence number: 0
	FopSequenceNumber *SequenceDetailsTypeU_94494S `xml:"fopSequenceNumber,omitempty"`

	// This segment conveys Form of Payment FreeText.   Old FOP(s) are considered as one freeflow text even if there is more than one old form of payment.  e.g.: FP O/CA+CCVI+/CH CA and CCVI are considered as freeflow text.
	FopFreeflow *FreeTextInformationType_94495S `xml:"fopFreeflow,omitempty"`

	// will convey the switches and data associated to the FOP table
	PnrSupplementaryData *PNRSupplementaryDataType `xml:"pnrSupplementaryData,omitempty"`

	// will contain all the data related to the payment transaction
	PaymentModule *PaymentGroupType `xml:"paymentModule,omitempty"`
}

type FareBasisCodesLineType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FareBasisCodesLineType"`

	// Fare element information
	FareElement *FareElementType `xml:"fareElement,omitempty"`
}

type FareCategoryCodesTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FareCategoryCodesTypeI"`

	// Rate Code (code set list not used)
	FareType formats.AlphaNumericString_Length1To6 `xml:"fareType,omitempty"`
}

type FareDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FareDataType"`

	// Issue identifier
	IssueIdentifier formats.AlphaNumericString_Length1To1 `xml:"issueIdentifier,omitempty"`

	// To specify the type of monetary amount, the amount and the currency code
	MonetaryInfo *MonetaryInformationDetailsTypeI_8308C `xml:"monetaryInfo,omitempty"`

	// Tax fields
	TaxFields *TaxFieldsType `xml:"taxFields,omitempty"`
}

type FareElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FareElementType"`

	// Contains primary code of the fare element
	PrimaryCode formats.AlphaNumericString_Length1To3 `xml:"primaryCode,omitempty"`

	// Connection indicator
	Connection formats.AlphaNumericString_Length1To3 `xml:"connection,omitempty"`

	// Not valid before
	NotValidBefore formats.AlphaNumericString_Length6To6 `xml:"notValidBefore,omitempty"`

	// Not valid after
	NotValidAfter formats.AlphaNumericString_Length6To6 `xml:"notValidAfter,omitempty"`

	// Baggage allowance
	BaggageAllowance formats.AlphaNumericString_Length2To3 `xml:"baggageAllowance,omitempty"`

	// Fare basis
	FareBasis formats.AlphaNumericString_Length1To6 `xml:"fareBasis,omitempty"`

	// Ticket designator
	TicketDesignator formats.AlphaNumericString_Length1To6 `xml:"ticketDesignator,omitempty"`
}

type FareQualifierDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FareQualifierDetailsTypeI"`

	// Rate Code Information
	FareCategories *FareCategoryCodesTypeI `xml:"fareCategories,omitempty"`
}

type FerryAccomodationPackageDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FerryAccomodationPackageDescriptionType"`

	// This segment conveys the package code.
	PackageCode *ProductInformationTypeI `xml:"packageCode,omitempty"`

	// This segment conveys the hotel code the of the accomodation package to which it is attached.
	HotelInformation *HotelPropertyType_26378S `xml:"hotelInformation,omitempty"`

	// This segment conveys the check-in date and time for the accomodation package to which it is attached.
	HotelCheckInInformation *StructuredDateTimeInformationType_24436S `xml:"hotelCheckInInformation,omitempty"`

	// This segment is used to convey the hotel area code information.
	AreaCodeInfo *PlaceLocationIdentificationTypeU_24573S `xml:"areaCodeInfo,omitempty"`

	// This segment is used to give the number of nights spent in the accomodation package to which it is attached.
	NumberOfNights *NumberOfUnitsType `xml:"numberOfNights,omitempty"`

	// This segment is used to convey the price of the accomodation to which it is attached.
	HotelItemPrice *TariffInformationTypeU `xml:"hotelItemPrice,omitempty"`

	RoomInfoGroup *RoomInfoGroup `xml:"roomInfoGroup,omitempty"`
}

type RoomInfoGroup struct {

	// This segment is used to describe the room to which it is attached.
	RoomDetailsInformation *HotelRoomType_20177S `xml:"roomDetailsInformation,omitempty"`

	// This segment is used to convey the number of instances of the room to which it is attached.
	NumberOfRooms *NumberOfUnitsType `xml:"numberOfRooms,omitempty"`
}

type FerryBookingDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FerryBookingDescriptionType"`

	// This segment is used to identify the ferry target provider for the message and is leading the description group for the ferry provider booking.
	FerryProviderInformation *CompanyInformationType_20151S `xml:"ferryProviderInformation,omitempty"`

	// This group describes the ferry booking itinerary. It contains the ferry sailing leg information.
	ItineraryInfoGroup *FerryLegDescriptionType `xml:"itineraryInfoGroup,omitempty"`

	// This group describes the accomodation (hotel) package attached to the booking.
	AccomodationPackageInfoGroup *FerryAccomodationPackageDescriptionType `xml:"accomodationPackageInfoGroup,omitempty"`

	// This segment conveys the ferry booking number information.
	BookingNumberInformation *ReservationControlInformationTypeI_20153S `xml:"bookingNumberInformation,omitempty"`
}

type FerryLegDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FerryLegDescriptionType"`

	// Conveys the sailing details for an itinerary leg.
	SailingDetails *TravelProductInformationTypeU `xml:"sailingDetails,omitempty"`

	// Conveys the ship code and ship name.
	ShipDescription *ShipIdentificationType `xml:"shipDescription,omitempty"`

	// This segment conveys the check-in time for the ferry sailing leg to which it is attached.
	SailingLegCheckInInformation *StructuredDateTimeInformationType_21109S `xml:"sailingLegCheckInInformation,omitempty"`

	// Conveys the list of passengers associated to the ferry leg.
	PassengerAssociation *ReferenceInformationTypeI_25132S `xml:"passengerAssociation,omitempty"`

	PriceInfoGroup *PriceInfoGroup `xml:"priceInfoGroup,omitempty"`

	VehicleInfoGroup *VehicleInfoGroup `xml:"vehicleInfoGroup,omitempty"`

	ServiceInfoGroup *ServiceInfoGroup `xml:"serviceInfoGroup,omitempty"`

	AnimalInfoGroup *AnimalInfoGroup `xml:"animalInfoGroup,omitempty"`
}

type AnimalInfoGroup struct {

	// This segment conveys the type of animal.
	AnimalInformation *SpecificDataInformationTypeI `xml:"animalInformation,omitempty"`

	// This segment conveys the price per animal of the same type.
	AnimalRoutePrice *TariffInformationTypeU `xml:"animalRoutePrice,omitempty"`
}

type ServiceInfoGroup struct {

	// This segment describes the on-board service.
	ServiceInformation *AttributeTypeU `xml:"serviceInformation,omitempty"`

	// This segment conveys the number of services of the attached service.
	NumberOfServices *NumberOfUnitsType `xml:"numberOfServices,omitempty"`

	// This segment conveys the price per unit of the attached service.
	ServiceRoutePrice *TariffInformationTypeU `xml:"serviceRoutePrice,omitempty"`
}

type VehicleInfoGroup struct {

	// This segment conveys the description of a vehicle.
	VehicleInformation *VehicleTypeU `xml:"vehicleInformation,omitempty"`

	// This segment is used to convey the number of bicycles associated to a ferry booking. Note: this segment is ignored if the vehicule description is not "bicycle".
	NumberOfBicycles *NumberOfUnitsType `xml:"numberOfBicycles,omitempty"`

	// This segment holds the price per vehicle.
	VehicleRoutePrice *TariffInformationTypeU `xml:"vehicleRoutePrice,omitempty"`
}

type PriceInfoGroup struct {

	// This segment conveys the route price information for the passenger it is linked to.
	RoutePriceInformation *TariffInformationTypeU `xml:"routePriceInformation,omitempty"`

	// This segment describes the passenger category type.
	PassengerCategoryType *AttributeTypeU_24552S `xml:"passengerCategoryType,omitempty"`

	// This segment is used to convey the number of passengers to which the price applies.
	NumberOfPassengers *NumberOfUnitsType `xml:"numberOfPassengers,omitempty"`
}

type FlightSegmentDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FlightSegmentDetailsTypeI"`
}

type FormOfPaymentDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentDetailsTypeI"`

	// Fop type (Cash, Credit card...)
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	// Credit card vendor code
	VendorCode formats.AlphaNumericString_Length1To3 `xml:"vendorCode,omitempty"`

	// Credit card number
	CreditCardNumber formats.AlphaNumericString_Length1To35 `xml:"creditCardNumber,omitempty"`

	// expiry date (MMYY)
	ExpiryDate formats.AlphaNumericString_Length1To4 `xml:"expiryDate,omitempty"`

	// FOP purpose
	ExtendedPayment formats.AlphaNumericString_Length1To3 `xml:"extendedPayment,omitempty"`

	// Unstructured fop layout (used for Voucher print purpose or guarantee details).
	FopFreeText formats.AlphaNumericString_Length1To70 `xml:"fopFreeText,omitempty"`
}

type FormOfPaymentDetailsTypeI_20667C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentDetailsTypeI_20667C"`

	// Reporting code
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	// Currency code per form of payment
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Form of payment amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// Vendor code of the credit card. ex: VI
	VendorCode formats.AlphaNumericString_Length1To35 `xml:"vendorCode,omitempty"`

	// Account number
	CreditCardNumber formats.AlphaNumericString_Length1To35 `xml:"creditCardNumber,omitempty"`

	// Expiration date
	ExpiryDate formats.AlphaNumericString_Length1To4 `xml:"expiryDate,omitempty"`

	// Approval code
	ApprovalCode formats.AlphaNumericString_Length1To17 `xml:"approvalCode,omitempty"`

	// Source of approval code
	SourceOfApproval formats.AlphaNumericString_Length1To3 `xml:"sourceOfApproval,omitempty"`

	// Authorised amount
	AuthorisedAmount formats.NumericDecimal_Length1To18 `xml:"authorisedAmount,omitempty"`

	// Address verification code
	AddressVerification formats.AlphaNumericString_Length1To3 `xml:"addressVerification,omitempty"`

	// Customer file reference
	CustomerAccount formats.AlphaNumericString_Length1To35 `xml:"customerAccount,omitempty"`

	// Extended payment code
	ExtendedPayment formats.AlphaNumericString_Length1To3 `xml:"extendedPayment,omitempty"`

	// not used
	FopFreeText formats.AlphaNumericString_Length1To70 `xml:"fopFreeText,omitempty"`

	// Credit card corporate contract
	MembershipStatus formats.AlphaNumericString_Length1To3 `xml:"membershipStatus,omitempty"`

	// Credit card transaction information
	TransactionInfo formats.AlphaNumericString_Length1To35 `xml:"transactionInfo,omitempty"`
}

type FormOfPaymentDetailsTypeI_52343C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentDetailsTypeI_52343C"`

	// Fop type (Cash, Credit card...)
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Credit card vendor code
	VendorCode formats.AlphaNumericString_Length1To3 `xml:"vendorCode,omitempty"`

	// Credit card number
	CreditCardNumber formats.AlphaNumericString_Length1To35 `xml:"creditCardNumber,omitempty"`

	// expiry date (MMYY)
	ExpiryDate formats.AlphaNumericString_Length1To4 `xml:"expiryDate,omitempty"`

	// FOP purpose
	ExtendedPayment formats.AlphaNumericString_Length1To3 `xml:"extendedPayment,omitempty"`

	// Unstructured fop layout (used for Voucher print purpose or guarantee details).
	FopFreeText formats.AlphaNumericString_Length1To70 `xml:"fopFreeText,omitempty"`
}

type FormOfPaymentDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentDetailsType"`

	// Generic type of the Mean of Payment used : CC credit Card CA cash CH cheque WW web payment... INV invoice
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`
}

type FormOfPaymentInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentInformationType"`

	// Format key that identify the FOP within a FOP table. (CCVI, ...)
	FopCode formats.AlphaNumericString_Length1To20 `xml:"fopCode,omitempty"`

	// Name of the FOP map table used in order to validate the FP element.
	FopMapTable formats.AlphaNumericString_Length1To20 `xml:"fopMapTable,omitempty"`

	// This corresponds to the fop billing code (CASH CA / Credit CC). This is only used in case of a MS reporting code. (it corresponds to XX of @FPMSXX tag of TPF tables)
	FopBillingCode formats.AlphaNumericString_Length1To20 `xml:"fopBillingCode,omitempty"`

	// Fop is a old / new fop.
	FopStatus formats.AlphaNumericString_Length1To3 `xml:"fopStatus,omitempty"`

	// Corresponds to the EDIFACT code.  This enables to identify the type of FOP that will be added in case of a structured EDIFACT (i.e. via PNRADD) addition of the FOP. (it corresponds to the @EDI tag of TPF tables) Here is an example: Customer is eager to add a structured cash FOP using an EDIFACT message. The fopEdiCode will be filled with CA which means cash. Then in the FOP table in charge of validating free flow and generating FOP free flow, the system will try to look for the FOP map having CA as fop EDI code. If we are in an Air France (AF) ATO/CTO: the system will get FP CA.... If we are in an Iberia (IB) ATO/CTO: the system will get FP CASH,.... If we are in an United Airline (UA) ATO/CTO: the system will get FP S.... ...  (@EDI value)
	FopEdiCode formats.AlphaNumericString_Length1To20 `xml:"fopEdiCode,omitempty"`

	// This corresponds to the fop code which is used on reporting side.  (XX value of @FPXXxx)
	FopReportingCode formats.AlphaNumericString_Length1To20 `xml:"fopReportingCode,omitempty"`

	// This is the FOP printed code  (@PR value)
	FopPrintedCode formats.AlphaNumericString_Length1To20 `xml:"fopPrintedCode,omitempty"`

	// This is the FOP electronic ticketing code. This is used to classify any FOP from the FOP table and also to determine how the FOP should be transmitted to the airline concerned. Based on this, the absence of the switch would make the FOP disallowed for ETKT, for National System Ticketing Server Travel Agency locations and all Central Ticketing offices  (@ET value)
	FopElecTicketingCode formats.AlphaNumericString_Length1To20 `xml:"fopElecTicketingCode,omitempty"`
}

type FormOfPaymentTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentTypeI"`

	// Fop details
	FormOfPayment *FormOfPaymentDetailsTypeI `xml:"formOfPayment,omitempty"`
}

type FormOfPaymentTypeI_16862S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentTypeI_16862S"`

	// Description of the form of paiement
	FormOfPayment *FormOfPaymentDetailsTypeI_20667C `xml:"formOfPayment,omitempty"`

	OtherFormOfPayment *FormOfPaymentDetailsTypeI_20667C `xml:"otherFormOfPayment,omitempty"`
}

type FormOfPaymentTypeI_29553S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentTypeI_29553S"`

	// FOP details
	FormOfPayment *FormOfPaymentDetailsTypeI_52343C `xml:"formOfPayment,omitempty"`
}

type FormOfPaymentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FormOfPaymentType"`

	// Generic status(new/old) and type(cash, cheque, card...) of the MOP
	FormOfPayment *FormOfPaymentDetailsType `xml:"formOfPayment,omitempty"`
}

type FraudScreeningGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FraudScreeningGroupType"`

	// This data element is used to indicate if risk management must be performed at authorization time: - Y means risk management data will be appended to author; - N means risk management data will not be appended;
	FraudScreening *StatusType_94568S `xml:"fraudScreening,omitempty"`

	// this segment contains the IP address used in RMM (risk management module or fraud screening)
	IpAdress *DeviceControlDetailsType `xml:"ipAdress,omitempty"`

	// Merchant's website URL.
	MerchantURL *CommunicationContactType `xml:"merchantURL,omitempty"`

	// will convey either the phone or the email adress of the payer
	PayerPhoneOrEmail *PhoneAndEmailAddressType_94565S `xml:"payerPhoneOrEmail,omitempty"`

	// this segment contains the shopper session used in RMM (risk management module)
	ShopperSession *SystemDetailsInfoType_94569S `xml:"shopperSession,omitempty"`

	// conveys information about payer LastName (surName) and FirstName (givenName)
	PayerName *TravellerInformationType_94570S `xml:"payerName,omitempty"`

	// stores the payer date of birth
	PayerDateOfBirth *StructuredDateTimeInformationType_94567S `xml:"payerDateOfBirth,omitempty"`

	// Information about the billing address (can be extracted from the AB PNR element)
	BillingAddress *AddressType `xml:"billingAddress,omitempty"`

	// Used to store reference information on the payer for fraud screening purpose: social security number driving license information frequent flyer information
	FormOfIdDetails *ReferenceInfoType_94566S `xml:"formOfIdDetails,omitempty"`
}

type FreeTextDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextDetailsType"`

	// text subject qualifier
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// information type
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// status
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// company id
	CompanyId formats.AlphaNumericString_Length1To35 `xml:"companyId,omitempty"`

	// Language, coded
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`

	// source, coded
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// encoding
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextDetailsType_142107C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextDetailsType_142107C"`

	// text subject qualifier
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Manual : M
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// encoding
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextDetailsType_142141C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextDetailsType_142141C"`

	// mutually defined ZZZ
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// AH Browser Accept headers UA Browser User Agent
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// Manual : M
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// ZZZ mutually agreed
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextDetailsType_187698C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextDetailsType_187698C"`

	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`

	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextDetailsType_198207C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextDetailsType_198207C"`

	// Text qualifier - 3 for literal text
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Information type
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// 1A for Amadeus
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// ISO language code
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`

	// Text source Manual or System
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// Character set
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextDetailsType_46357C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextDetailsType_46357C"`

	// booking description subject qualifier
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// booking information type
	InformationType formats.AlphaNumericString_Length1To3 `xml:"informationType,omitempty"`

	// booking desscription source
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// booking description encoding information
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type FreeTextInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType"`

	FreeTextDetails *FreeTextDetailsType_187698C `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText formats.AlphaNumericString_Length1To320 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_136708S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_136708S"`

	// Free text type
	FreeTextDetails *FreeTextDetailsType_198207C `xml:"freeTextDetails,omitempty"`

	// Free text
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_136715S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_136715S"`

	// Free text type
	FreeTextDetails *FreeTextDetailsType_198207C `xml:"freeTextDetails,omitempty"`

	// 1 or 2 lines of free text
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_20551S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_20551S"`

	// Text attributes
	FreeTextDetails *FreeTextDetailsType_142107C `xml:"freeTextDetails,omitempty"`

	// SVCs / Service Information (Amtrak). 63 characters maximum length, and a maximum of 5 lines per train segment.
	FreeText formats.AlphaNumericString_Length1To63 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_25445S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_25445S"`

	// booking description details
	FreeTextDetails *FreeTextDetailsType_46357C `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_29860S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_29860S"`

	// Free text information.
	FreeTextDetails *FreeTextDetailsType_187698C `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText formats.AlphaNumericString_Length1To5 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_6235S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_6235S"`

	// To convey the type of the freeflow text.
	FreeTextDetails *FreeTextDetailsType_187698C `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText formats.AlphaNumericString_Length1To64 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_94495S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_94495S"`

	// will contain the FOP free flow text
	FreeTextDetails *FreeTextDetailsType_142107C `xml:"freeTextDetails,omitempty"`

	// FOP freeflow
	FreeText formats.AlphaNumericString_Length1To199 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_94526S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_94526S"`

	// will contain the browser information
	FreeTextDetails *FreeTextDetailsType_142141C `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText formats.AlphaNumericString_Length1To199 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_94561S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_94561S"`

	// will describe the purchase
	FreeTextDetails *FreeTextDetailsType_142107C `xml:"freeTextDetails,omitempty"`

	// Purchase free text description
	FreeText formats.AlphaNumericString_Length1To199 `xml:"freeText,omitempty"`
}

type FreeTextInformationType_9865S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextInformationType_9865S"`

	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText formats.AlphaNumericString_Length1To500 `xml:"freeText,omitempty"`
}

type FreeTextQualificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextQualificationTypeI"`

	// Identifies whether the free text is coded or not coded  3 for Literal text
	SubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"subjectQualifier,omitempty"`

	// Coded text, or specifies type of  info   Surface segment : 2 for Address  or 5 for Telephone nature un known  Cruise segment : P30 for Ship Name
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// Company code
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// ISO code for language of free text
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`
}

type FreeTextQualificationTypeI_148295C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextQualificationTypeI_148295C"`

	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`
}

type FreeTextQualificationTypeI_185754C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextQualificationTypeI_185754C"`

	// Categorise the format of the text (free text, coded,...)
	TextSubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"textSubjectQualifier,omitempty"`

	// Provides a code identifying the information (phone, OSI, etc...)
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// The airline code that may be associated to this information
	CompanyId formats.AlphaNumericString_Length1To35 `xml:"companyId,omitempty"`
}

type FreeTextQualificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FreeTextQualificationType"`

	// Identifies whether the free text is coded or not coded :  3 for Literal text
	SubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"subjectQualifier,omitempty"`

	// Information type, coded. see code list
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// Transmittable/non-transmittable indicator (S or X).
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Airline or system code.
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`
}

type FrequencyDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequencyDetailsTypeU"`

	// Indicates number of instalments for the payment
	InstalmentsNumber formats.NumericInteger_Length1To9 `xml:"instalmentsNumber,omitempty"`

	// Indicates frequency of instalments for the payment D daily M monthly W weekly
	InstalmentsFrequency formats.AlphaNumericString_Length1To3 `xml:"instalmentsFrequency,omitempty"`

	// Indicates when first instalment should take place
	InstalmentsStartDate formats.AlphaNumericString_Length1To35 `xml:"instalmentsStartDate,omitempty"`

	// indicates extended payment start date format
	InstalmentsDatrDateFormat formats.AlphaNumericString_Length1To3 `xml:"instalmentsDatrDateFormat,omitempty"`
}

type FrequencyTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequencyTypeU"`

	// extended payment characteristics
	ExtendedPaymentDetails *FrequencyDetailsTypeU `xml:"extendedPaymentDetails,omitempty"`
}

type FrequencyType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequencyType"`

	// Indicate if the sequence number represents days of the week or days of the month.
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Used to represent days of the week or days of the month. For week : 1 is monday and 7 is sunday. For month : 1 is the first day of the month.
	Value formats.AlphaNumericString_Length1To1 `xml:"value,omitempty"`
}

type FrequentFlyerInformationGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequentFlyerInformationGroupType"`

	// To specify frequent traveller information
	FrequentTravellerInfo *FrequentTravellerIdentificationCodeType_74327S `xml:"frequentTravellerInfo,omitempty"`

	// Promotion code used to compute redemption/upgrade price in miles, when applicable
	DiscountInformation *DiscountInformationType `xml:"discountInformation,omitempty"`

	// Original booking class
	BookingClassInformation *ProductInformationTypeI_73824S `xml:"bookingClassInformation,omitempty"`
}

type FrequentTravellerIdentificationCodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequentTravellerIdentificationCodeType"`

	// Frequent Traveller Info
	AirlineFrequentTraveler *FrequentTravellerIdentificationType_198190C `xml:"airlineFrequentTraveler,omitempty"`
}

type FrequentTravellerIdentificationCodeType_38226S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequentTravellerIdentificationCodeType_38226S"`

	// Airline Frequent Traveller Info
	AirlineFrequentTraveler *FrequentTravellerIdentificationType `xml:"airlineFrequentTraveler,omitempty"`

	// Alliance Frequent Traveller Info
	AllianceFrequentTraveler *FrequentTravellerIdentificationType_64816C `xml:"allianceFrequentTraveler,omitempty"`
}

type FrequentTravellerIdentificationCodeType_74327S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequentTravellerIdentificationCodeType_74327S"`

	// FREQUENT TRAVELER IDENTIFICATION
	FrequentTraveler *FrequentTravellerIdentificationTypeI `xml:"frequentTraveler,omitempty"`

	// PRIORITY DETAILS
	PriorityDetails *PriorityDetailsType `xml:"priorityDetails,omitempty"`

	// Specify the redemption information
	RedemptionInformation *ProductAccountDetailsTypeI `xml:"redemptionInformation,omitempty"`
}

type FrequentTravellerIdentificationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequentTravellerIdentificationTypeI"`

	// Airline code
	Company formats.AlphaNumericString_Length2To3 `xml:"company,omitempty"`

	// Frequent traveler number
	MembershipNumber formats.AlphaNumericString_Length1To27 `xml:"membershipNumber,omitempty"`

	// Provide airline customer value of the frequent traveller.
	CustomerValue formats.NumericInteger_Length1To4 `xml:"customerValue,omitempty"`
}

type FrequentTravellerIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequentTravellerIdentificationType"`

	// Carrier where the FQTV is registered.
	Company formats.AlphaNumericString_Length2To3 `xml:"company,omitempty"`

	// Frequent Traveller Reference Number
	MembershipNumber formats.AlphaNumericString_Length1To27 `xml:"membershipNumber,omitempty"`

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length4To4 `xml:"tierLevel,omitempty"`

	// To specify the Priority of the FQTV.
	PriorityCode formats.AlphaNumericString_Length1To1 `xml:"priorityCode,omitempty"`

	// Full Text Tier description ex: EMERALD, SAPPHIRE
	TierDescription formats.AlphaNumericString_Length1To35 `xml:"tierDescription,omitempty"`
}

type FrequentTravellerIdentificationType_198190C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequentTravellerIdentificationType_198190C"`

	// Carrier where the FQTV is registered.
	Company formats.AlphaNumericString_Length1To3 `xml:"company,omitempty"`

	// Frequent Traveller Reference Number
	MembershipNumber formats.AlphaNumericString_Length1To28 `xml:"membershipNumber,omitempty"`
}

type FrequentTravellerIdentificationType_64816C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A FrequentTravellerIdentificationType_64816C"`

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length1To4 `xml:"tierLevel,omitempty"`

	// To specify the Priority of the FQTV.
	PriorityCode formats.AlphaNumericString_Length1To1 `xml:"priorityCode,omitempty"`

	// Full Text Tier description ex: EMERALD, SAPPHIRE
	TierDescription formats.AlphaNumericString_Length1To35 `xml:"tierDescription,omitempty"`

	// Alliance name
	CompanyCode formats.AlphaNumericString_Length2To3 `xml:"companyCode,omitempty"`
}

type GategoryType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A GategoryType"`

	// A (first) category number
	CategoryNum1 formats.NumericInteger_Length1To3 `xml:"categoryNum1,omitempty"`

	// [2-16] characters for Special [2-10] characters for Dual
	CategoryName formats.AlphaNumericString_Length1To16 `xml:"categoryName,omitempty"`
}

type GeneralOptionInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A GeneralOptionInformationType"`

	// Option type.  hotel/car/cruise/train/insurance options type.  CAR : ACD : Action Code for Display BCS : Billing Control System CWB : Car Warning Banner MKT : Marketing informations OS : Other Services informations  HOTEL : ACD : Action code for Display BCS : Billing Control System TXT : Marketing text NAM : Hotel property name CNM : Hotel chain name MKT : Marketing informations CXL : Cancellation policy DES : Rate description PRI : Pricing information  CRUISE : GPN : Group code. CXN : Booking cancellation number. CFN : Booking confirmation number. NME : Passenger's last name, first name and title.
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// 1. Hotel segment: for Update option indicator  Y for yes  N for no
	UpdateIndicator formats.AlphaString_Length1To1 `xml:"updateIndicator,omitempty"`

	// Free text
	Freetext formats.AlphaNumericString_Length1To255 `xml:"freetext,omitempty"`
}

type GeneralOptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A GeneralOptionType"`

	// ONLY ONE OCCURRENCE  Each option is one segment
	OptionDetail *GeneralOptionInformationType `xml:"optionDetail,omitempty"`
}

type GenericAuthorisationResultType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A GenericAuthorisationResultType"`

	// transaction authorization approval data
	ApprovalCodeData *AuthorizationApprovalDataType `xml:"approvalCodeData,omitempty"`
}

type GenericDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A GenericDetailsTypeI"`

	// Seat Characteristic
	SeatCharacteristic formats.AlphaNumericString_Length1To2 `xml:"seatCharacteristic,omitempty"`
}

type HotelProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelProductInformationType"`

	// Property header details
	Property *PropertyHeaderDetailsType `xml:"property,omitempty"`

	// Room details
	HotelRoom *RoomDetailsType `xml:"hotelRoom,omitempty"`

	// Rate code
	Negotiated *RateCodeRestrictedType `xml:"negotiated,omitempty"`

	// Other information
	OtherHotelInfo *OtherHotelInformationType `xml:"otherHotelInfo,omitempty"`
}

type HotelPropertyType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelPropertyType"`

	// This composite is used to convey the hotel identifier.
	HotelReference *HotelUniqueIdType `xml:"hotelReference,omitempty"`

	// This composite is used to convey the hotel name
	HotelName formats.AlphaNumericString_Length1To40 `xml:"hotelName,omitempty"`

	// This data element is used to indicates if the hotel is compliant with the fire safety rules.
	FireSafetyIndicator formats.AlphaNumericString_Length1To1 `xml:"fireSafetyIndicator,omitempty"`
}

type HotelPropertyType_26378S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelPropertyType_26378S"`

	// This composite is used to convey the hotel identification details.
	HotelReference *HotelUniqueIdType_47769C `xml:"hotelReference,omitempty"`
}

type HotelRoomRateInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelRoomRateInformationType"`

	// This data element is used to convey the room type
	RoomType formats.AlphaNumericString_Length3To3 `xml:"roomType,omitempty"`

	// This data element is used to convey the hotel rate code
	RatePlanCode formats.AlphaNumericString_Length3To3 `xml:"ratePlanCode,omitempty"`

	// This data element is used to convey the hotel rate category code
	RateCategoryCode formats.AlphaNumericString_Length1To35 `xml:"rateCategoryCode,omitempty"`

	// This data element is used to indicate if the rate code is a qualified rate code or not.
	RateQualifiedIndic formats.AlphaNumericString_Length1To1 `xml:"rateQualifiedIndic,omitempty"`
}

type HotelRoomRateInformationType_46329C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelRoomRateInformationType_46329C"`

	// This data element is used to convey the room type.
	RoomType formats.AlphaNumericString_Length1To3 `xml:"roomType,omitempty"`
}

type HotelRoomType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelRoomType"`

	// This composite is used to convey the room rate identifier.
	RoomRateIdentifier *HotelRoomRateInformationType `xml:"roomRateIdentifier,omitempty"`

	// This data element is used to convey the booking code.
	BookingCode formats.AlphaNumericString_Length1To10 `xml:"bookingCode,omitempty"`

	// This composite is used to convey the occupancy level of the hotel room.
	GuestCountDetails *NumberOfUnitDetailsTypeI_18670C `xml:"guestCountDetails,omitempty"`

	// This data element is used to convey the override room type (non-Amadeus room types).
	RoomTypeOverride formats.AlphaNumericString_Length1To8 `xml:"roomTypeOverride,omitempty"`
}

type HotelRoomType_20177S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelRoomType_20177S"`

	// This data element is used to convey the override room type (non-Amadeus room types).  NOTE: WAS AN..35 IN STANDRD.
	RoomTypeOverride formats.AlphaNumericString_Length1To10 `xml:"roomTypeOverride,omitempty"`
}

type HotelRoomType_25429S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelRoomType_25429S"`

	// This composite is used to convey the room type
	RoomRateIdentifier *HotelRoomRateInformationType_46329C `xml:"roomRateIdentifier,omitempty"`

	// This data element is used to convey the booking code.
	BookingCode formats.AlphaNumericString_Length1To35 `xml:"bookingCode,omitempty"`

	// This composite is used to convey the occupancy level of the hotel room.
	GuestCountDetails *NumberOfUnitDetailsTypeI_46330C `xml:"guestCountDetails,omitempty"`

	// This data element is used to convey the override room type (non-Amadeus room types).
	RoomTypeOverride formats.AlphaNumericString_Length1To16 `xml:"roomTypeOverride,omitempty"`
}

type HotelUniqueIdType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelUniqueIdType"`

	// To convey the chain code in the property ID
	ChainCode formats.AlphaNumericString_Length2To3 `xml:"chainCode,omitempty"`

	// To convey the city code in the hotel Id
	CityCode formats.AlphaNumericString_Length3To3 `xml:"cityCode,omitempty"`

	// To convey the property code in the Hotel Id
	HotelCode formats.AlphaNumericString_Length3To3 `xml:"hotelCode,omitempty"`
}

type HotelUniqueIdType_47769C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A HotelUniqueIdType_47769C"`

	// This element is used to convey the hotel code. The list of possible values is different for each Ferry provider.
	HotelCode formats.AlphaNumericString_Length1To10 `xml:"hotelCode,omitempty"`
}

type IdentificationNumberTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A IdentificationNumberTypeI"`

	// will contain the IP adress of the shopper
	Address formats.AlphaNumericString_Length1To35 `xml:"address,omitempty"`

	// will contain IP for IP adress
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`
}

type IndividualPnrSecurityInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A IndividualPnrSecurityInformationType"`

	// Returned before End of transaction when retrieving a PNR security element
	Security *IndividualSecurityType_3194C `xml:"security,omitempty"`

	// Returned when retrieving a PNR
	SecurityInfo *SecurityInformationType `xml:"securityInfo,omitempty"`

	// Code as in the display:  G for Amadeus Global Core Office Identification  I for IATA number  P for Pseudo-Office Identification Default is G.
	Indicator formats.AlphaNumericString_Length1To1 `xml:"indicator,omitempty"`
}

type IndividualSecurityType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A IndividualSecurityType"`

	// office Id
	Office formats.AlphaNumericString_Length1To9 `xml:"office,omitempty"`

	// R for Read  B for Both read and write  N for None
	AccessMode formats.AlphaString_Length1To1 `xml:"accessMode,omitempty"`

	// - F for Family identifier
	OfficeIdentifier formats.AlphaNumericString_Length1To1 `xml:"officeIdentifier,omitempty"`
}

type IndividualSecurityType_3194C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A IndividualSecurityType_3194C"`

	// Type of receiver  G: Amadeus Global Core Office Id with possible wild card chars '*' to mask some part(s) of it// I:IATA nb or '*' for all, no wild card char //P:Pseudo-Office Id or '*' for all, no wild card char.
	Identification formats.AlphaNumericString_Length1To14 `xml:"identification,omitempty"`

	// R for Read  B for Both read and write  N for None
	AccessMode formats.AlphaString_Length1To1 `xml:"accessMode,omitempty"`
}

type InsuranceBusinessDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsuranceBusinessDataType"`

	// To convey the provider code and country. Can be empty in case preferences have been set up.
	ProviderProductDetails *InsuranceProductDetailsType `xml:"providerProductDetails,omitempty"`

	// provides details about the substitute name or the nanny name.
	SubstiteName *TravellerInformationTypeI `xml:"substiteName,omitempty"`

	// Amount that is added to the total premium in case an extrareference is specified.
	ExtraPremium *MonetaryInformationTypeI `xml:"extraPremium,omitempty"`

	ProductSection *ProductSection `xml:"productSection,omitempty"`

	// contains the different amounts (net premium/taxes/total premium)
	PlanCostInfo *TariffInformationTypeI_22057S `xml:"planCostInfo,omitempty"`

	PlanTypeDetails *PlanTypeDetails `xml:"planTypeDetails,omitempty"`

	ContactDetails *ContactDetails `xml:"contactDetails,omitempty"`

	SubscriberAddressSection *SubscriberAddressSection `xml:"subscriberAddressSection,omitempty"`

	CoverageDetails *CoverageDetails `xml:"coverageDetails,omitempty"`

	// to specify a commission.
	ComissionAmount *CommissionInformationType `xml:"comissionAmount,omitempty"`

	InsuranceFopSection *InsuranceFopSection `xml:"insuranceFopSection,omitempty"`

	// To specify a number which means that the insurance is in a confirmed status.
	ConfirmationNumber *ReservationControlInformationTypeI `xml:"confirmationNumber,omitempty"`

	// Used to specify the necesary data for pricing
	ProductKnowledge *ActionDetailsTypeI `xml:"productKnowledge,omitempty"`

	PassengerDetails *PassengerDetails `xml:"passengerDetails,omitempty"`

	// To convey information if the document has been printed or not.
	PrintInformation *DocumentInformationDetailsTypeI `xml:"printInformation,omitempty"`
}

type ContactDetails struct {

	// data to add some comments on the insurance element
	Miscelaneous *MiscellaneousRemarksType_12240S `xml:"miscelaneous,omitempty"`

	// Used to specify a phone number as an emergency contact
	PhoneNumber *PhoneAndEmailAddressType_32298S `xml:"phoneNumber,omitempty"`

	// to specify the name of a person in case of an emergeny
	ContactName *TravellerInformationTypeI `xml:"contactName,omitempty"`
}

type SubscriberAddressSection struct {

	// This segment is used to convey the contact name
	NameDetails *NameTypeU `xml:"nameDetails,omitempty"`

	// to specify the address of the subscriber
	AddressInfo *AddressTypeU `xml:"addressInfo,omitempty"`

	// Used to specify a phone number
	PhoneNumber *PhoneAndEmailAddressType_32298S `xml:"phoneNumber,omitempty"`
}

type CoverageDetails struct {

	// To specify the details of the insurance policy.
	PolicyDetails *InsurancePolicyType `xml:"policyDetails,omitempty"`

	CoverageInfo *CoverageInfo `xml:"coverageInfo,omitempty"`

	// To specifie the covered persons: here it conveys the NB/NM and ON options
	CoveredPassenger *TravellerInformationTypeI_15923S `xml:"coveredPassenger,omitempty"`

	// starting date and end date
	CoverageDates *StructuredPeriodInformationType `xml:"coverageDates,omitempty"`

	// Details of the subscription: date and time.
	SubscriptionDetails *StructuredDateTimeInformationType_20644S `xml:"subscriptionDetails,omitempty"`

	// To convey the details of the insurance seller.
	AgentReferenceDetails *UserIdentificationType_9456S `xml:"agentReferenceDetails,omitempty"`
}

type CoverageInfo struct {

	// For codelist 415Z, only values CP, CV, CM may apply here
	Coverage *InsuranceCoverageType `xml:"coverage,omitempty"`

	// Values and currency of the different coverages amounts.
	CoverageValues *MonetaryInformationTypeI `xml:"coverageValues,omitempty"`
}

type InsuranceFopSection struct {

	// To convey the form of payment
	FormOfPaymentSection *FormOfPaymentTypeI_16862S `xml:"formOfPaymentSection,omitempty"`

	// To provide form of payment extended data
	FopExtendedData *StatusTypeI_13270S `xml:"fopExtendedData,omitempty"`
}

type PassengerDetails struct {

	// to specify to which passenger the insurance is associated: if omitted then it's for all the names in the PNR.
	PassengerAssociation *ReferenceInformationType `xml:"passengerAssociation,omitempty"`

	// product knowledge indicator
	PerPaxProdKnowledge *ActionDetailsTypeI `xml:"perPaxProdKnowledge,omitempty"`

	// To specify the birthdate of the insuree.
	DateOfBirthInfo *StructuredDateTimeInformationType `xml:"dateOfBirthInfo,omitempty"`

	// to specify the name /age of the insuree
	PassengerFeatures *TravellerInformationType `xml:"passengerFeatures,omitempty"`

	// to specify a remark for the insuree
	InsureeRemark *MiscellaneousRemarksType `xml:"insureeRemark,omitempty"`

	// To specify the details concerning the documentation and the age of the insuree.
	TravelerDocInfo *PassengerDocumentDetailsType `xml:"travelerDocInfo,omitempty"`

	// fare discount code used per Pax
	PolicyDetails *InsurancePolicyType `xml:"policyDetails,omitempty"`

	TravelerValueDetails *TravelerValueDetails `xml:"travelerValueDetails,omitempty"`

	PremiumPerTariffPerPax *PremiumPerTariffPerPax `xml:"premiumPerTariffPerPax,omitempty"`

	// To convey the premium perpax
	PremiumPerpaxInfo *TariffInformationTypeI_22057S `xml:"premiumPerpaxInfo,omitempty"`

	// The Individual passenger reservation information
	VoucherNumber *ReservationControlInformationTypeU_31804S `xml:"voucherNumber,omitempty"`
}

type PremiumPerTariffPerPax struct {

	// To convey the tariffcode information per passenger.
	TariffCodeInfo *InsuranceProductDetailsType_20775S `xml:"tariffCodeInfo,omitempty"`

	// to specify the amount/currency  per passenger and tariff code.
	TariffCodePerPaxAmount *MonetaryInformationTypeI `xml:"tariffCodePerPaxAmount,omitempty"`
}

type TravelerValueDetails struct {

	// For codelist 415Z, only TV value may apply here
	TravelCost *InsuranceCoverageType `xml:"travelCost,omitempty"`

	// to specify the amount/currency  per insuree
	TravelAmount *MonetaryInformationTypeI `xml:"travelAmount,omitempty"`
}

type PlanTypeDetails struct {

	// Provides information about the type of plan being quoted/booked
	PlanType *InsuranceProviderAndProductsType `xml:"planType,omitempty"`

	// to specify the value of the trip.
	TravelValue *MonetaryInformationTypeI `xml:"travelValue,omitempty"`
}

type ProductSection struct {

	// To convey the products or the tariffcodes together with description and amounts.
	ProductCode *InsuranceProductDetailsType_20774S `xml:"productCode,omitempty"`

	// To convey the information the provider estimates important on a given product.
	InformationLines *FreeTextInformationType_6235S `xml:"informationLines,omitempty"`
}

type InsuranceCoverageType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsuranceCoverageType"`

	// Indicate type of amount (eg. Medical Coverage, Trip Value, etc)
	CoverageIndicator formats.AlphaNumericString_Length0To2 `xml:"coverageIndicator,omitempty"`
}

type InsuranceCoverageType_25483S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsuranceCoverageType_25483S"`

	// to indicate which coverage we are talking about.
	CoverageIndicator formats.AlphaNumericString_Length1To2 `xml:"coverageIndicator,omitempty"`
}

type InsuranceNameType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsuranceNameType"`

	// insurance traveller details
	InsuranceTravelerDetails *SpecificTravellerDetailsType `xml:"insuranceTravelerDetails,omitempty"`

	// travelerperpax details
	TravelerPerpaxDetails *TravelerPerpaxDetailsType `xml:"travelerPerpaxDetails,omitempty"`
}

type InsurancePolicyType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsurancePolicyType"`

	// to specify a discount for the insuree like if it's a family etc..
	FareDiscount formats.AlphaNumericString_Length1To3 `xml:"fareDiscount,omitempty"`
}

type InsuranceProductDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsuranceProductDetailsType"`

	// This data element is used to convey the company code of a non-air company
	CompanyCode formats.AlphaNumericString_Length1To3 `xml:"companyCode,omitempty"`

	// To identify the countrycode from the provider.
	CountryCode formats.AlphaNumericString_Length1To2 `xml:"countryCode,omitempty"`

	// Authorization number provided by ht insurance company
	ExtraReference formats.AlphaNumericString_Length1To60 `xml:"extraReference,omitempty"`
}

type InsuranceProductDetailsType_20774S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsuranceProductDetailsType_20774S"`

	// This data element is used to convey the company code of a non-air company
	CompanyCode formats.AlphaNumericString_Length1To3 `xml:"companyCode,omitempty"`

	// To identify the countrycode from the provider.
	CountryCode formats.AlphaNumericString_Length1To2 `xml:"countryCode,omitempty"`

	// This composite contains the code of the insurance elements.
	ProductDetails *ProviderInformationType `xml:"productDetails,omitempty"`

	// contains the extensions for the main insurance product
	ExtensionIdentification *ProviderInformationType `xml:"extensionIdentification,omitempty"`

	// tariff code info. tariff code and tariff familly code.
	TariffCodeDetails *TariffcodeType `xml:"tariffCodeDetails,omitempty"`
}

type InsuranceProductDetailsType_20775S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsuranceProductDetailsType_20775S"`

	// tariff code info. tariff code and tariff familly code.
	TariffCodeDetails *TariffcodeType `xml:"tariffCodeDetails,omitempty"`
}

type InsuranceProviderAndProductsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InsuranceProviderAndProductsType"`

	// Type of trip (package. leisure etc...)
	TripType formats.AlphaNumericString_Length1To3 `xml:"tripType,omitempty"`

	// Code of the operator who provides the TOUR.
	TourOperator formats.AlphaNumericString_Length1To4 `xml:"tourOperator,omitempty"`

	// To specify the countries involved in the Travel assistance element.
	CountryInfo *CountrydescriptionType `xml:"countryInfo,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InteractiveFreeTextTypeI"`

	// Provides information on the text conveyed in the IFT: language, type...
	FreeTextQualification *FreeTextQualificationTypeI_185754C `xml:"freeTextQualification,omitempty"`

	// The information itself
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type InteractiveFreeTextTypeI_136698S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InteractiveFreeTextTypeI_136698S"`

	// Describes free text type
	FreetextDetail *FreeTextQualificationTypeI `xml:"freetextDetail,omitempty"`

	// One occurrence is supposed to represent a logical entity of free text (e.g. one line of text).
	Text formats.AlphaNumericString_Length1To70 `xml:"text,omitempty"`
}

type InteractiveFreeTextTypeI_99363S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A InteractiveFreeTextTypeI_99363S"`

	FreeTextQualification *FreeTextQualificationTypeI_148295C `xml:"freeTextQualification,omitempty"`

	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type ItemDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemDescriptionType"`

	// Qualify the item being described
	ItemCharacteristic formats.AlphaNumericString_Length1To3 `xml:"itemCharacteristic,omitempty"`
}

type ItemNumberIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemNumberIdentificationTypeU"`

	// leg number
	Number formats.AlphaNumericString_Length1To2 `xml:"number,omitempty"`
}

type ItemNumberIdentificationTypeU_46320C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemNumberIdentificationTypeU_46320C"`

	// The place of the product in the Tour booking.
	ItemID formats.NumericInteger_Length1To2 `xml:"itemID,omitempty"`

	// It qualifies the item ID type.
	ItemIDQualifier formats.AlphaNumericString_Length1To3 `xml:"itemIDQualifier,omitempty"`
}

type ItemNumberTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemNumberTypeU"`

	// Provides information about the product place in the tour booking. It locally identifies the product in the booking.
	ItemIdentification *ItemNumberIdentificationTypeU_46320C `xml:"itemIdentification,omitempty"`
}

type ItemNumberTypeU_33258S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemNumberTypeU_33258S"`

	// leg number - idicate with leg is the first one, the second one, etc.
	ItemNumberDetails *ItemNumberIdentificationTypeU `xml:"itemNumberDetails,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemReferencesAndVersionsType"`

	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To1 `xml:"referenceType,omitempty"`

	// The value of the reference
	UniqueReference formats.AlphaNumericString_Length1To3 `xml:"uniqueReference,omitempty"`
}

type ItemReferencesAndVersionsType_6550S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemReferencesAndVersionsType_6550S"`

	// ID details
	IDSection *UniqueIdDescriptionType `xml:"iDSection,omitempty"`
}

type ItemReferencesAndVersionsType_9271S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemReferencesAndVersionsType_9271S"`

	// Defines the type of reference used: GPN : group code
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// The value of the reference.
	UniqueReference formats.AlphaNumericString_Length1To10 `xml:"uniqueReference,omitempty"`
}

type ItemReferencesAndVersionsType_94556S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ItemReferencesAndVersionsType_94556S"`

	// qualifies the type of the reference used. Here it will be: PRI Payment Record Id APP Application Correlator Id EXT Third party Record Id
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`

	// The value of the payment record/correlator Id
	UniqueReference formats.AlphaNumericString_Length1To35 `xml:"uniqueReference,omitempty"`
}

type LocationIdentificationBatchTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationIdentificationBatchTypeU"`

	// Set to: IATA to indicate IATA location code 1A to indicate a 1A location CPY to indicate a Car provider location
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`

	// Location extended name for  - Amadeus location type  - Provider location type (followed by an *) - Free text for collection option.  - Free text for delivery option.
	Name formats.AlphaNumericString_Length1To60 `xml:"name,omitempty"`
}

type LocationIdentificationBatchTypeU_198230C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationIdentificationBatchTypeU_198230C"`

	// Location Code "1A" to indicate Amadeus location type "CPY" to indicate a Provider location type
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`

	// Location extended name for  Amadeus location type and  Provider location type (followed by an *)
	Name formats.AlphaNumericString_Length1To11 `xml:"name,omitempty"`
}

type LocationIdentificationBatchTypeU_46344C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationIdentificationBatchTypeU_46344C"`

	// location code
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`

	// location qualifier (city, country ...)
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// location name
	Name formats.AlphaNumericString_Length1To70 `xml:"name,omitempty"`
}

type LocationIdentificationBatchTypeU_56454C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationIdentificationBatchTypeU_56454C"`

	// Railway station location code
	Code formats.AlphaNumericString_Length1To5 `xml:"code,omitempty"`

	// Code type
	Qualifier formats.AlphaNumericString_Length1To2 `xml:"qualifier,omitempty"`

	// Location name
	Name formats.AlphaNumericString_Length1To60 `xml:"name,omitempty"`
}

type LocationIdentificationBatchTypeU_60738C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationIdentificationBatchTypeU_60738C"`

	// Railway station location code
	Code formats.AlphaNumericString_Length1To5 `xml:"code,omitempty"`

	// Code type
	Qualifier formats.AlphaNumericString_Length1To2 `xml:"qualifier,omitempty"`
}

type LocationIdentificationTypeS struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationIdentificationTypeS"`

	// Station. See IATA Airline Coding directory. IATA 3 letter city/aircode code
	CityCode formats.AlphaNumericString_Length1To25 `xml:"cityCode,omitempty"`
}

type LocationIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationIdentificationTypeU"`

	// 162: country
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Country name
	Name formats.AlphaNumericString_Length1To25 `xml:"name,omitempty"`
}

type LocationIdentificationTypeU_198211C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationIdentificationTypeU_198211C"`

	// Identification of the site
	Code formats.AlphaNumericString_Length1To10 `xml:"code,omitempty"`

	// Site name
	Name formats.AlphaNumericString_Length1To30 `xml:"name,omitempty"`
}

type LocationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationTypeI"`

	TrueLocationId formats.AlphaString_Length3To3 `xml:"trueLocationId,omitempty"`
}

type LocationTypeI_2784C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationTypeI_2784C"`

	// AIR segment : boarding point ATX segment : boarding point CAR segment : pick-up point city CCR segment : pick-up point city HHL segment : city code HTL segment : check-in city MIS segment : city code SUR segment : city/airport code Trn Amtrak sgt: board point city code Trn SNCF sgt: board point city code (RESARAIL code) TTO segment: departure location TUR segment: tour start city CRU segment: sailing departure port
	CityCode formats.AlphaNumericString_Length1To5 `xml:"cityCode,omitempty"`

	// TRN SNCF segment : board point city name.
	CityName formats.AlphaNumericString_Length1To17 `xml:"cityName,omitempty"`
}

type LocationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationTypeU"`

	// Port code.
	Code formats.AlphaNumericString_Length1To5 `xml:"code,omitempty"`

	// Port name.
	Name formats.AlphaNumericString_Length1To70 `xml:"name,omitempty"`

	// Port codes are non-standard and specific to the Amadeus ferry business.
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`
}

type LocationTypeU_46324C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LocationTypeU_46324C"`

	// city code
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`

	// city name
	Name formats.AlphaNumericString_Length1To70 `xml:"name,omitempty"`

	// country code
	Country formats.AlphaNumericString_Length1To2 `xml:"country,omitempty"`

	// location qualifier for the repetition (departure location, arrival location ...)
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`
}

type LongFreeTextType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A LongFreeTextType"`

	// To specify the type of freetext
	FreetextDetail *FreeTextQualificationType `xml:"freetextDetail,omitempty"`

	// Long free text information.
	LongFreetext formats.AlphaNumericString_Length1To999 `xml:"longFreetext,omitempty"`
}

type MeanOfPaymentDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MeanOfPaymentDataType"`

	// This segment will convey the type of the FOP. Exple : CC credit card CA cash CH cheque SWI swipe card WA web account WB web bank(fund tranfer)
	FopInformation *FormOfPaymentType `xml:"fopInformation,omitempty"`

	// will allow the usage of FOP segment as trigger for GASY and GINV groups
	Dummy *DummySegmentTypeI `xml:"dummy,omitempty"`

	// will convey all credit card data needed for the payment
	CreditCardData *CreditCardDataGroupType `xml:"creditCardData,omitempty"`
}

type MeasurementsBatchTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MeasurementsBatchTypeU"`

	// Measurement qualifier (maximum unit qualifier).
	MeasurementQualifier formats.AlphaNumericString_Length1To3 `xml:"measurementQualifier,omitempty"`

	// Unit Qualifer for the range value.
	ValueRange *ValueRangeTypeU `xml:"valueRange,omitempty"`
}

type MessageActionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MessageActionDetailsTypeI"`

	// MESSAGE FUNCTION OR BUSINESS DETAILS
	Business *MessageFunctionBusinessDetailsTypeI `xml:"business,omitempty"`
}

type MessageFunctionBusinessDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MessageFunctionBusinessDetailsTypeI"`

	// - Message function: REG for regular feed - Itinerary type: see codeset list: air/car/hotel/taxi/train/tour/surface...
	Function formats.AlphaNumericString_Length1To3 `xml:"function,omitempty"`
}

type MessageReferenceType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MessageReferenceType"`

	// This number is used to identify and track ALL messages related to a given cardholder transaction (author, retry, reversal ...). It is usually composed of: - the date when the message was formatted followed by - the message number   Field 37  Official definition of Retrieval Reference Number from ISO8583: Field 37 contains a number used with other key data elements to identify and track all messages related to a given cardholder transaction (referred to as a transaction set). It is usually assigned by the acquirer, but it may be assigned by a merchant or by an individual electronic terminal. V.I.P. will also generate the retrieval reference number for transactions it initiates. This field contains two parts. The first four digits are usually a yddd date (Julian date format). The date is defined to be the same day as the date in Field 7_Transmission Date and Time, of the original request. The last eight digits are a numeric transaction identification number. The value in field 37 can be based on the content of fields 7 and 11 in the original request or advice as shown in the recommendation below: . Positions 1_4: the yddd equivalent of the field 7 date . Positions 5_6: the hours from the time in field 7 . Positions 7_12: the value from field 11
	RetrievalReferenceNumber formats.AlphaNumericString_Length1To12 `xml:"retrievalReferenceNumber,omitempty"`

	// Authorization characteristics indicator   Field 62.1 Possible values: A C E F K M S U V W R I P N T
	AuthorCharacteristicIndicator formats.AlphaString_Length1To1 `xml:"authorCharacteristicIndicator,omitempty"`

	// Authorization response code   Field 39
	AuthorResponseCode formats.AlphaNumericString_Length2To2 `xml:"authorResponseCode,omitempty"`

	// Card Level Result (Product Identification value)  Field 62.23
	CardLevelResult formats.AlphaNumericString_Length2To2 `xml:"cardLevelResult,omitempty"`

	// Additional POS Information - Terminal Type  Field 60.1 - Position 1  CAT (Cardholder-Activated Terminal indicator) or UAT (Unattended Acceptance Terminal)
	TerminalType formats.AlphaNumericString_Length1To1 `xml:"terminalType,omitempty"`
}

type MileageTimeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MileageTimeDetailsTypeI"`

	FlightLegMileage formats.NumericInteger_Length1To18 `xml:"flightLegMileage,omitempty"`

	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`
}

type MiscellaneousChargeOrderType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousChargeOrderType"`

	// Type of service
	Type formats.AlphaNumericString_Length1To2 `xml:"type,omitempty"`
}

type MiscellaneousRemarkType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarkType"`

	// This data element is used to convey the type of the remark. (see data mapping to view the codes)
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	// Free text and message sequence numbers of the remarks.
	Freetext formats.AlphaNumericString_Length1To120 `xml:"freetext,omitempty"`

	// This data element is used to convey the business function
	BusinessFunction formats.AlphaNumericString_Length1To3 `xml:"businessFunction,omitempty"`

	// language used for the free text.
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`

	// Indicates if it has been manually entered by an agent or system generated.
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// Coded identification of the character encoding used in the interchange
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type MiscellaneousRemarkType_151C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarkType_151C"`

	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// This is the 3rd character (x) of the remark title RIx or RMx, or 2 letter code for RMxx, conditional for RM, not applicable for RC and RQ
	Category formats.AlphaNumericString_Length1To2 `xml:"category,omitempty"`

	// Free text and message sequence numbers of the remarks.
	Freetext formats.AlphaNumericString_Length1To127 `xml:"freetext,omitempty"`

	// Provider type (element RIA):  1 for Air provider   2 for Car provider (CCR)  3 for Hotel Provider (HHL)  M for Miscellaneous
	ProviderType formats.AlphaNumericString_Length1To3 `xml:"providerType,omitempty"`
}

type MiscellaneousRemarkType_18076C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarkType_18076C"`

	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark . ACC for Acceptance . BGG for Baggage . BPP for Boarding Pass Printing . GT for Gate . GNL for General
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	// Free text and message sequence numbers of the remarks.
	Freetext formats.AlphaNumericString_Length1To70 `xml:"freetext,omitempty"`
}

type MiscellaneousRemarkType_198195C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarkType_198195C"`

	// This data element is used to convey the type of the remark. (see data mapping to view the codes)
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	// Free text and message sequence numbers of the remarks.
	Freetext formats.AlphaNumericString_Length1To90 `xml:"freetext,omitempty"`

	// This data element is used to convey the business function
	BusinessFunction formats.AlphaNumericString_Length1To3 `xml:"businessFunction,omitempty"`

	// language used for the free text.
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`

	// Indicates if it has been manually entered by an agent or system generated.
	Source formats.AlphaNumericString_Length1To3 `xml:"source,omitempty"`

	// Coded identification of the character encoding used in the interchange
	Encoding formats.AlphaNumericString_Length1To3 `xml:"encoding,omitempty"`
}

type MiscellaneousRemarkType_861C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarkType_861C"`

	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark . ACC for Acceptance . BGG for Baggage . BPP for Boarding Pass Printing . GT for Gate . GNL for General
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	// Free text and message sequence numbers of the remarks.
	Freetext formats.AlphaNumericString_Length1To35 `xml:"freetext,omitempty"`
}

type MiscellaneousRemarksType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarksType"`

	// miscellaneous remarks
	RemarkDetails *MiscellaneousRemarkType_861C `xml:"remarkDetails,omitempty"`
}

type MiscellaneousRemarksType_12240S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarksType_12240S"`

	// miscellaneous remarks
	RemarkDetails *MiscellaneousRemarkType_18076C `xml:"remarkDetails,omitempty"`
}

type MiscellaneousRemarksType_136700S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarksType_136700S"`

	// miscellaneous remarks
	RemarkDetails *MiscellaneousRemarkType_198195C `xml:"remarkDetails,omitempty"`
}

type MiscellaneousRemarksType_211S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarksType_211S"`

	// Miscellaneous remqrks
	Remarks *MiscellaneousRemarkType_151C `xml:"remarks,omitempty"`

	// For confidential remark RC
	IndividualSecurity *IndividualSecurityType `xml:"individualSecurity,omitempty"`
}

type MiscellaneousRemarksType_664S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MiscellaneousRemarksType_664S"`

	// miscellaneous remarks
	RemarkDetails *MiscellaneousRemarkType `xml:"remarkDetails,omitempty"`
}

type MonetaryInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationDetailsTypeI"`

	// Yield type
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Amount of the Yield
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`
}

type MonetaryInformationDetailsTypeI_17849C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationDetailsTypeI_17849C"`

	// Indicates amount is Fare amount
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier,omitempty"`

	// Used to specify an amount of money
	Amount formats.AlphaNumericString_Length1To18 `xml:"amount,omitempty"`

	// currency in which the amount is expressed
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsTypeI_4220C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationDetailsTypeI_4220C"`

	// Monetary amount qualifier : - NP : Net Premium - PR : Premium - CV : Coverage - TV : Travel Value - SAV : Saving Amount
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Amount
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// Eg: USD,FRF,EUR...
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`
}

type MonetaryInformationDetailsTypeI_8308C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationDetailsTypeI_8308C"`

	// . F for Fare basis . E for Equivalent  . T for Total
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Amount
	Amount formats.AlphaNumericString_Length1To18 `xml:"amount,omitempty"`

	// Eg: USD,FRF,EUR...
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`
}

type MonetaryInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationDetailsType"`

	// Here is the list and the purpose of each amount today stored in the FP: I Transaction total amount  Total amount authorized in authorization transaction  IPC Transaction total amount in PNR currency Total amount authorized is also stored in PNR currency. Indeed, reversal must be done with the rate of exchange valid at time of authorization and therefore this avoids storing the rate of exchange and performing amount conversion at reversal time.  IT Initial TST total amount  Amount of TST multiplied by the number of passengers associated to the TST  ITC Initial TST total amount in PNR currency   IT amount in PNR currency for same reason as IPC amount  R Total amount / Remaining amount  Current authorized amount. Originally it is the total amount authorized and then this amount  may decrease in case of total/partial reversal.  T Initial Tst Individual amount  Amount of TST  TPC Initial Tst Individual amount in PNR currency  Amount of TST in PNR currency for same reason as IPC amount  AUT Authorized Amount Maybe different from the one given in input(for exple, if on input we have 2pax and the amount per pax. In case of bulk, we will authorize the sum of both amounts) It can also be used for:  Total Fare Amount 712 or additional collection amount A
	TypeQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"typeQualifier,omitempty"`

	// Value of the amount.  This is conveyed as a 'string' and therefore several strings can stand for the  same amount (eg. 14 , 1400, 14.00... could potentially stand for 14.00 EUR). This  means that sender/receiver of this message will need to come to an agreement  concerning the way the amount is transferred in this segment.
	Amount formats.NumericDecimal_Length1To35 `xml:"amount,omitempty"`

	// IATA alphabetic currency code.  Eg: USD,GBP,EUR...
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationTypeI"`

	// Total Trip value in a given currency
	MonetaryDetails *MonetaryInformationDetailsTypeI_17849C `xml:"monetaryDetails,omitempty"`

	// Base Trip value in a given currency
	OtherMonetaryDetails *MonetaryInformationDetailsTypeI_17849C `xml:"otherMonetaryDetails,omitempty"`
}

type MonetaryInformationTypeI_1689S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationTypeI_1689S"`

	// To specify monetary information
	Information *MonetaryInformationDetailsTypeI_4220C `xml:"information,omitempty"`
}

type MonetaryInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationType"`

	// Yield info
	MonetaryDetails *MonetaryInformationDetailsTypeI `xml:"monetaryDetails,omitempty"`

	OtherMonetaryDetails *MonetaryInformationDetailsTypeI `xml:"otherMonetaryDetails,omitempty"`
}

type MonetaryInformationType_94557S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A MonetaryInformationType_94557S"`

	// Contains the currencies and the various amounts
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails,omitempty"`

	OtherMonetaryDetails *MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"`
}

type NameAndAddressBatchTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NameAndAddressBatchTypeU"`

	// W for party to revieve written confirmation
	PartyQualifier formats.AlphaNumericString_Length1To3 `xml:"partyQualifier,omitempty"`

	// This composite is used to convey the address
	AddressDetails *NameAndAddressDetailsTypeU `xml:"addressDetails,omitempty"`

	// This composite is used to convey the party name
	PartyNameDetails *PartyNameBatchTypeU `xml:"partyNameDetails,omitempty"`
}

type NameAndAddressDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NameAndAddressDetailsTypeU"`

	// Address line 1
	Line1 formats.AlphaNumericString_Length1To25 `xml:"line1,omitempty"`

	// address line 2
	Line2 formats.AlphaNumericString_Length1To25 `xml:"line2,omitempty"`
}

type NameInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NameInformationTypeU"`

	// name qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// name
	Name formats.AlphaNumericString_Length1To12 `xml:"name,omitempty"`
}

type NameInformationTypeU_9747C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NameInformationTypeU_9747C"`

	// to convey to who the address applies
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Company name
	Name formats.AlphaNumericString_Length1To30 `xml:"name,omitempty"`

	// Insuree name
	Identifier formats.AlphaNumericString_Length1To30 `xml:"identifier,omitempty"`
}

type NameTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NameTypeU"`

	// Used to specify the name field in the address field.
	NameInformation *NameInformationTypeU_9747C `xml:"nameInformation,omitempty"`
}

type NameTypeU_136701S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NameTypeU_136701S"`

	// Name information
	NameInformation *NameInformationTypeU `xml:"nameInformation,omitempty"`
}

type NumberOfUnitDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NumberOfUnitDetailsTypeI"`

	// Group counter corresponding to passengers, and so value from 0 to 99.
	NumberOfUnit formats.NumericInteger_Length1To2 `xml:"numberOfUnit,omitempty"`

	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`
}

type NumberOfUnitDetailsTypeI_18670C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NumberOfUnitDetailsTypeI_18670C"`

	// This data element is used to convey the occupancy level of the room
	NumberOfUnit formats.NumericInteger_Length1To1 `xml:"numberOfUnit,omitempty"`

	// This data element is used to indicate the occupancy is the number of Adults in the room.
	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`
}

type NumberOfUnitDetailsTypeI_2755C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NumberOfUnitDetailsTypeI_2755C"`

	// PNR Header / Queue header / number of remaining items in Queue
	Number formats.NumericInteger_Length1To5 `xml:"number,omitempty"`

	// PNR for PNR
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`
}

type NumberOfUnitDetailsTypeI_35712C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NumberOfUnitDetailsTypeI_35712C"`

	// Number of units.
	NumberOfUnit formats.NumericInteger_Length1To2 `xml:"numberOfUnit,omitempty"`
}

type NumberOfUnitDetailsTypeI_46330C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NumberOfUnitDetailsTypeI_46330C"`

	// occupation of the room
	NumberOfUnit formats.NumericInteger_Length1To2 `xml:"numberOfUnit,omitempty"`

	// unit qualifier
	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`
}

type NumberOfUnitsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NumberOfUnitsTypeI"`

	// Number of Units detail
	NumberDetail *NumberOfUnitDetailsTypeI_2755C `xml:"numberDetail,omitempty"`
}

type NumberOfUnitsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NumberOfUnitsType"`

	// Number of Unit Details
	QuantityDetails *NumberOfUnitDetailsTypeI_35712C `xml:"quantityDetails,omitempty"`
}

type NumberOfUnitsType_76106S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A NumberOfUnitsType_76106S"`

	// Number of Unit Details
	QuantityDetails *NumberOfUnitDetailsTypeI `xml:"quantityDetails,omitempty"`
}

type ODKeyPerformanceDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ODKeyPerformanceDataType"`

	// schedule change indicator -'C' or void
	ScheduleChange formats.AlphaString_Length1To1 `xml:"scheduleChange,omitempty"`

	// oversale data
	Oversale *OversaleDataType `xml:"oversale,omitempty"`
}

type ONDType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ONDType"`

	// Yield informations:  Adjusted Yield Segment Bid Price Effective Yield Revenue Loss OND Yield
	YieldInformations *MonetaryInformationType `xml:"yieldInformations,omitempty"`

	// Class code as defined in yield retrieved / Class combinaison of the yield retrieved
	ClassCombinaison *ProductInformationTypeI_76271S `xml:"classCombinaison,omitempty"`

	// Origin and Destination of the Yield
	Ondyield *OriginAndDestinationDetailsTypeI_76268S `xml:"ondyield,omitempty"`

	// Origin And Destination of the Trip
	TripOnD *OriginAndDestinationDetailsTypeI_76268S `xml:"tripOnD,omitempty"`
}

type OptionElementInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OptionElementInformationType"`

	// Option element office id
	MainOffice formats.AlphaNumericString_Length1To9 `xml:"mainOffice,omitempty"`

	// Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// Queue number
	Queue formats.NumericInteger_Length1To3 `xml:"queue,omitempty"`

	// Category number
	Category formats.NumericInteger_Length1To3 `xml:"category,omitempty"`

	Freetext formats.AlphaNumericString_Length1To61 `xml:"freetext,omitempty"`

	// queuing or cancellation time
	Time formats.Time24_HHMM `xml:"time,omitempty"`
}

type OptionElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OptionElementType"`

	OptionElementInfo *OptionElementInformationType `xml:"optionElementInfo,omitempty"`

	// Individual Security for OPQ/OPX elements
	IndividualSecurity *IndividualSecurityType `xml:"individualSecurity,omitempty"`
}

type OriginAndDestinationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OriginAndDestinationDetailsTypeI"`

	// City pair to indentify uniquely a leg in a multi-leg booking
	Origin formats.AlphaNumericString_Length1To3 `xml:"origin,omitempty"`

	// City pair to indentify uniquely a leg in a multi-leg booking
	Destination formats.AlphaNumericString_Length1To3 `xml:"destination,omitempty"`
}

type OriginAndDestinationDetailsTypeI_3061S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OriginAndDestinationDetailsTypeI_3061S"`

	// Airport/city code of  Origin In a Client request message, a non-blank ODI is used in an air sell request to advise that the following segments (TVL etc...) are connected. There is a maximum of 6 TVLs following a non-blank ODI.
	Origin formats.AlphaString_Length3To3 `xml:"origin,omitempty"`

	// Airport/city code of  Destination
	Destination formats.AlphaString_Length3To3 `xml:"destination,omitempty"`
}

type OriginAndDestinationDetailsTypeI_76268S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OriginAndDestinationDetailsTypeI_76268S"`

	// Departure's city code:3 character ATA/IATA airport/city code
	Origin formats.AlphaString_Length1To3 `xml:"origin,omitempty"`

	// Arrival's city code:3 character ATA/IATA airport/city code
	Destination formats.AlphaString_Length1To3 `xml:"destination,omitempty"`
}

type OriginatorDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OriginatorDetailsTypeI"`

	// Country code
	CodedCountry formats.AlphaNumericString_Length1To3 `xml:"codedCountry,omitempty"`

	// Currency code
	CodedCurrency formats.AlphaNumericString_Length1To3 `xml:"codedCurrency,omitempty"`

	// Language code
	CodedLanguage formats.AlphaNumericString_Length1To3 `xml:"codedLanguage,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OriginatorIdentificationDetailsTypeI"`

	// IATA code
	OriginatorId formats.NumericInteger_Length1To9 `xml:"originatorId,omitempty"`

	// Office ID of the PNR owner.
	InHouseIdentification1 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification1,omitempty"`

	// Amid of the owner of the SBR.
	InHouseIdentification2 formats.NumericInteger_Length1To9 `xml:"inHouseIdentification2,omitempty"`
}

type OriginatorIdentificationDetailsTypeI_198179C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OriginatorIdentificationDetailsTypeI_198179C"`

	// Agency Iata code
	OriginatorId formats.AlphaNumericString_Length8To8 `xml:"originatorId,omitempty"`
}

type OriginatorIdentificationDetailsTypeI_37406C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OriginatorIdentificationDetailsTypeI_37406C"`

	// This data element is used to convey the bouking source.
	OriginatorId formats.NumericInteger_Length5To8 `xml:"originatorId,omitempty"`
}

type OriginatorIdentificationDetailsTypeI_46358C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OriginatorIdentificationDetailsTypeI_46358C"`

	// Origin OficeID
	InHouseIdentification1 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification1,omitempty"`

	// Target OfficeID
	InHouseIdentification2 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification2,omitempty"`
}

type OtherHotelInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OtherHotelInformationType"`

	// Currency Code at Property 1. For AY Direct Access segment: Currency code
	CurrencyCode formats.AlphaString_Length3To3 `xml:"currencyCode,omitempty"`
}

type OtherInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OtherInformationType"`

	// Queue cycle complete indicator that may appear in Queue working response message.  QCC for Queue cycle complete
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Indicates the type of Queue in a Queue working response message.  PNR for PNR  MSG for Message
	QueueType formats.AlphaNumericString_Length1To3 `xml:"queueType,omitempty"`
}

type OtherSegmentDataTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OtherSegmentDataTypeI"`

	// Cabin Code
	Cabin formats.AlphaNumericString_Length1To2 `xml:"cabin,omitempty"`

	// Sub class number
	Subclass formats.NumericInteger_Length1To1 `xml:"subclass,omitempty"`

	// Flight type : - D for Domestic - I for International - L for Longhaul - S for Shorthaul
	FlightType formats.AlphaNumericString_Length1To2 `xml:"flightType,omitempty"`

	// Overbooking indicator
	Overbooking formats.AlphaString_Length2To2 `xml:"overbooking,omitempty"`
}

type OversaleDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A OversaleDataType"`

	// Bid price oversale number
	OversaleNumber formats.NumericDecimal_Length1To3 `xml:"oversaleNumber,omitempty"`

	// Oversale indicator F  for Bid-Price Feed Oversale O for Bid-Price Oversale P  for Pushed Minimum Oversale
	OversaleIndicator formats.AlphaString_Length1To1 `xml:"oversaleIndicator,omitempty"`
}

type PNRSupplementaryDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PNRSupplementaryDataType"`

	// will convey the values of the FOP data and switch maps
	DataAndSwitchMap *AttributeType_94576S `xml:"dataAndSwitchMap,omitempty"`
}

type POSGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A POSGroupType"`

	// - Office ID owner of the SBR. - IATA Code - Agent type
	SbrUserIdentificationOwn *UserIdentificationType `xml:"sbrUserIdentificationOwn,omitempty"`

	// - Corporate Code - City Code
	SbrSystemDetails *SystemDetailsInfoType_33158S `xml:"sbrSystemDetails,omitempty"`

	// Preferences - Country - Language - Currency
	SbrPreferences *UserPreferencesType `xml:"sbrPreferences,omitempty"`
}

type PackageDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PackageDescriptionType"`

	// Inclusive package type: I
	PackageType formats.AlphaNumericString_Length1To1 `xml:"packageType,omitempty"`

	// List of inclusive package
	PackageDetails *PackageIdentificationType `xml:"packageDetails,omitempty"`
}

type PackageIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PackageIdentificationType"`

	// Description of a package
	PackageDesc formats.AlphaNumericString_Length1To40 `xml:"packageDesc,omitempty"`
}

type PartyNameBatchTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PartyNameBatchTypeU"`

	// name
	Name1 formats.AlphaNumericString_Length1To35 `xml:"name1,omitempty"`
}

type PassengerDocumentDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PassengerDocumentDetailsType"`

	// Used to convey the age of the insuree
	BirthDate formats.AlphaNumericString_Length1To4 `xml:"birthDate,omitempty"`

	// Details on the document (visa, passport...)
	DocumentDetails *DocumentDetailsType `xml:"documentDetails,omitempty"`
}

type PassengerFlightDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PassengerFlightDetailsTypeI"`
}

type PaymentDataGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PaymentDataGroupType"`

	// Contains merchant information (Entity selling a product/service for wich payment is requested: airline, insurance provider...).
	MerchantInformation *CompanyInformationType_94554S `xml:"merchantInformation,omitempty"`

	// will convey all the monetary informations related to the payment : amount, currency, sub-amounts
	MonetaryInformation *MonetaryInformationType_94557S `xml:"monetaryInformation,omitempty"`

	// Conveys Payment Record ID (used by Payment Manager) to identify payment in a unique manner.  May convey also a "correlator Id" used by the calling application to reconciliate its payment data.  And also the "transaction Id" generated by the third party system (bank/PSP/PAyPAL...)
	PaymentId *ItemReferencesAndVersionsType_94556S `xml:"paymentId,omitempty"`

	// It will describe the content of the extended payment : when it will start, the frequency and how many times it should occur
	ExtendedPaymentInfo *FrequencyTypeU `xml:"extendedPaymentInfo,omitempty"`

	// The segment conveys the date/time of the transaction
	TransactionDateTime *StructuredDateTimeInformationType_94559S `xml:"transactionDateTime,omitempty"`

	// Will show the duration of validity of the payment request, mesured from receipt by the issuer. The customer has to agree to the payment within this period. Expressed in seconds.
	ExpirationPeriod *QuantityType_94558S `xml:"expirationPeriod,omitempty"`

	// Distribution Channel information
	DistributionChannelInformation *TerminalIdentificationDescriptionType `xml:"distributionChannelInformation,omitempty"`

	// will convey in free text the description of the purchase
	PurchaseDescription *FreeTextInformationType_94561S `xml:"purchaseDescription,omitempty"`

	// will convey all information needed to perform the checks requested by the banks/PSPs regarding the prevention of fraud.
	FraudScreeningData *FraudScreeningGroupType `xml:"fraudScreeningData,omitempty"`

	// Will be used to convey information dedicated to the Payment.
	PaymentDataMap *AttributeType_94553S `xml:"paymentDataMap,omitempty"`
}

type PaymentDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PaymentDetailsTypeI"`

	// To convey the guarantee /deposit form
	FormOfPaymentCode formats.AlphaNumericString_Length1To4 `xml:"formOfPaymentCode,omitempty"`

	// This data element is used to idicates if it is a guarantee or a deposit
	PaymentType formats.AlphaNumericString_Length1To4 `xml:"paymentType,omitempty"`

	// This data element is used to identify the type of service to be paid, in our case it will always be 3 for hotel
	ServiceToPay formats.AlphaNumericString_Length1To3 `xml:"serviceToPay,omitempty"`

	// This data element is used to convey the guarantee or the deposit reference.
	ReferenceNumber formats.AlphaNumericString_Length1To31 `xml:"referenceNumber,omitempty"`
}

type PaymentDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PaymentDetailsTypeU"`

	// Identify the mode of payment:  - CASH  - CC for credit card
	MethodCode formats.AlphaNumericString_Length1To4 `xml:"methodCode,omitempty"`

	// Purpose of the payment:  - DEPO for deposit  - FINA for final payment
	PurposeCode formats.AlphaNumericString_Length1To4 `xml:"purposeCode,omitempty"`

	// Amount paid
	Amount formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`

	// Currency used for the payment
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// date of the payment
	Date formats.AlphaNumericString_Length1To11 `xml:"date,omitempty"`
}

type PaymentGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PaymentGroupType"`

	// Used to describe the element on which the action is performed : FP/FC/PAY and in which context integrated/non integrated
	GroupUsage *CodedAttributeType_127282S `xml:"groupUsage,omitempty"`

	// will convey all data necessary for the paiment and not dependant from the Mean Of Payment
	PaymentData *PaymentDataGroupType `xml:"paymentData,omitempty"`

	// it will convey the Descriptive Billing Information: ONO, GWT, best Fare indicator....
	PaymentSupplementaryData *CodedAttributeType_94497S `xml:"paymentSupplementaryData,omitempty"`

	// will convey all the specificities of the Mean of Payment
	MopInformation *MeanOfPaymentDataType `xml:"mopInformation,omitempty"`

	// will allow the usage of FOP segment as trigger for MOPD and MOPS groups
	Dummy *DummySegmentTypeI `xml:"dummy,omitempty"`

	// will convey the result of the payment and related to the detailed Mean Of Payment
	MopDetailedData *DetailedPaymentDataType `xml:"mopDetailedData,omitempty"`
}

type PaymentInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PaymentInformationTypeI"`

	// This composite is used to convey the payment information
	PaymentDetails *PaymentDetailsTypeI `xml:"paymentDetails,omitempty"`
}

type PaymentInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PaymentInformationTypeU"`

	// Tour deposit details
	PaymentDetails *PaymentDetailsTypeU `xml:"paymentDetails,omitempty"`

	// Credit card name, number and exp. date
	CreditCardInformation *CreditCardInformationTypeU `xml:"creditCardInformation,omitempty"`
}

type PhoneAndEmailAddressType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PhoneAndEmailAddressType"`

	// Phone or Email contact type
	PhoneOrEmailType formats.AlphaNumericString_Length1To3 `xml:"phoneOrEmailType,omitempty"`

	// Structured telephone number
	TelephoneNumberDetails *StructuredTelephoneNumberType_198214C `xml:"telephoneNumberDetails,omitempty"`
}

type PhoneAndEmailAddressType_136723S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PhoneAndEmailAddressType_136723S"`

	// - PHO phone number - FAX fax number - MAI
	PhoneOrEmailType formats.AlphaNumericString_Length1To3 `xml:"phoneOrEmailType,omitempty"`

	// Email address
	EmailAddress formats.AlphaNumericString_Length1To70 `xml:"emailAddress,omitempty"`
}

type PhoneAndEmailAddressType_32298S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PhoneAndEmailAddressType_32298S"`

	// Phone or Email contact type
	PhoneOrEmailType formats.AlphaNumericString_Length1To4 `xml:"phoneOrEmailType,omitempty"`

	// Structured telephone number
	TelephoneNumber *StructuredTelephoneNumberType_48448C `xml:"telephoneNumber,omitempty"`

	// Email address
	EmailAddress formats.AlphaNumericString_Length1To90 `xml:"emailAddress,omitempty"`
}

type PhoneAndEmailAddressType_94565S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PhoneAndEmailAddressType_94565S"`

	// Phone or Email contact type
	PhoneOrEmailType formats.AlphaNumericString_Length1To4 `xml:"phoneOrEmailType,omitempty"`

	// Structured telephone number
	TelephoneNumberDetails *StructuredTelephoneNumberType `xml:"telephoneNumberDetails,omitempty"`

	// Email address
	EmailAddress formats.AlphaNumericString_Length1To70 `xml:"emailAddress,omitempty"`
}

type PlaceLocationIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PlaceLocationIdentificationTypeU"`

	// location code qualifier
	LocationType formats.AlphaNumericString_Length1To3 `xml:"locationType,omitempty"`

	// location text
	LocationDescription *LocationIdentificationBatchTypeU `xml:"locationDescription,omitempty"`

	// Associated airport/City code. Present if the pickup location is not an airport/city code.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_198193C `xml:"firstLocationDetails,omitempty"`
}

type PlaceLocationIdentificationTypeU_136722S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PlaceLocationIdentificationTypeU_136722S"`

	// Used to differenciate the pickup location (176) from the Dropoff location (DOL)
	LocationType formats.AlphaNumericString_Length1To3 `xml:"locationType,omitempty"`

	// Pickup or dropoff location details
	LocationDescription *LocationIdentificationBatchTypeU_198230C `xml:"locationDescription,omitempty"`
}

type PlaceLocationIdentificationTypeU_24573S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PlaceLocationIdentificationTypeU_24573S"`

	// Location type qualifier (ZZZ-Mutually defined for Ferry).
	LocationType formats.AlphaNumericString_Length1To3 `xml:"locationType,omitempty"`

	// Hotel location details.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_45087C `xml:"firstLocationDetails,omitempty"`
}

type PlaceLocationIdentificationTypeU_25436S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PlaceLocationIdentificationTypeU_25436S"`

	// location type (place of arrival, place of departure or staying)
	LocationType formats.AlphaNumericString_Length1To3 `xml:"locationType,omitempty"`

	// city information
	LocationDescription *LocationIdentificationBatchTypeU_46344C `xml:"locationDescription,omitempty"`

	// country description
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_46345C `xml:"firstLocationDetails,omitempty"`
}

type PlaceLocationIdentificationTypeU_32347S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PlaceLocationIdentificationTypeU_32347S"`

	// Type of location
	LocationType formats.AlphaNumericString_Length1To2 `xml:"locationType,omitempty"`

	// Railway station location details.
	LocationDescription *LocationIdentificationBatchTypeU_56454C `xml:"locationDescription,omitempty"`

	// Railway station country details.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_56455C `xml:"firstLocationDetails,omitempty"`
}

type PlaceLocationIdentificationTypeU_35293S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PlaceLocationIdentificationTypeU_35293S"`

	// Type of location
	LocationType formats.AlphaNumericString_Length1To2 `xml:"locationType,omitempty"`

	// Railway station location details.
	LocationDescription *LocationIdentificationBatchTypeU_60738C `xml:"locationDescription,omitempty"`

	// Railway station country details.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU_56455C `xml:"firstLocationDetails,omitempty"`
}

type PlaceLocationIdentificationTypeU_8954S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PlaceLocationIdentificationTypeU_8954S"`

	// Details of the embarkation port.
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU `xml:"firstLocationDetails,omitempty"`

	// Details of the disembarkation port.
	SecondLocationDetails *RelatedLocationTwoIdentificationTypeU `xml:"secondLocationDetails,omitempty"`
}

type PnrHistoryDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PnrHistoryDataType"`

	// Contains the last EOTed envelop number.
	CurrentRecord formats.NumericInteger_Length1To5 `xml:"currentRecord,omitempty"`
}

type PnrHistoryDataType_6022S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PnrHistoryDataType_6022S"`

	// Reference to previous envelop It may not exist when we are on element creation case.
	PreviousRecord formats.NumericInteger_Length1To3 `xml:"previousRecord,omitempty"`

	// Current envelop
	CurrentRecord formats.NumericInteger_Length1To3 `xml:"currentRecord,omitempty"`

	// History element name  ON, AS, RF... First char for type of action done, followed by a letter related to the element concerned.
	ElementType formats.AlphaString_Length1To2 `xml:"elementType,omitempty"`

	// Free flow text (history data element not detailed) Max length put from 254 to 255 for  the case of long history
	ElementData formats.AlphaNumericString_Length1To255 `xml:"elementData,omitempty"`
}

type PointOfSaleDataTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PointOfSaleDataTypeI"`

	// POSINV Classification: - C for Country - R for CRS
	Classification formats.AlphaString_Length1To1 `xml:"classification,omitempty"`

	// Point of Sale CRS
	Crs formats.AlphaNumericString_Length1To3 `xml:"crs,omitempty"`

	// Point of Sale Country Code
	PointOfSaleCountry formats.AlphaString_Length2To2 `xml:"pointOfSaleCountry,omitempty"`
}

type PricingOrTicketingSubsequentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PricingOrTicketingSubsequentType"`

	// Reason for issuance code.
	SpecialCondition formats.AlphaNumericString_Length1To3 `xml:"specialCondition,omitempty"`

	// Reason for Issuance Sub code
	OtherSpecialCondition formats.AlphaNumericString_Length1To3 `xml:"otherSpecialCondition,omitempty"`
}

type PriorityDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PriorityDetailsType"`

	// 1 : airline 2 : alliance
	Qualifier formats.AlphaNumericString_Length1To2 `xml:"qualifier,omitempty"`

	// Priority code
	PriorityCode formats.AlphaNumericString_Length1To1 `xml:"priorityCode,omitempty"`

	// Tier level
	TierLevel formats.AlphaNumericString_Length1To4 `xml:"tierLevel,omitempty"`

	// Tier description
	TierDescription formats.AlphaNumericString_Length1To70 `xml:"tierDescription,omitempty"`
}

type ProcessingInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProcessingInformationTypeI"`

	// Identifies the element we are talking about
	ActionQualifier formats.AlphaNumericString_Length1To3 `xml:"actionQualifier,omitempty"`

	// Used to qualifie the element with an indicator.
	ReferenceQualifier formats.AlphaNumericString_Length1To3 `xml:"referenceQualifier,omitempty"`
}

type ProductAccountDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductAccountDetailsTypeI"`

	// the award code returned by loyalty system in booking time and send to loyalty system in ticketing time.
	Category formats.AlphaNumericString_Length1To6 `xml:"category,omitempty"`

	// Contains the old class of the segment before the upgrade.
	SequenceNumber formats.AlphaNumericString_Length1To6 `xml:"sequenceNumber,omitempty"`

	// certificate number
	VersionNumber formats.AlphaNumericString_Length1To11 `xml:"versionNumber,omitempty"`

	// Fake Tier level received by TTY in.
	RateClass formats.AlphaNumericString_Length1To35 `xml:"rateClass,omitempty"`

	// stock control number
	ApprovalCode formats.AlphaNumericString_Length1To14 `xml:"approvalCode,omitempty"`
}

type ProductDataInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDataInformationTypeU"`

	// Tour product category (StandAlone, Package, Supplementary service ...)
	ProductCategory formats.AlphaNumericString_Length1To3 `xml:"productCategory,omitempty"`

	// Conveys the product code
	ProductCode formats.AlphaNumericString_Length1To14 `xml:"productCode,omitempty"`

	// Set to 1 if the product is an addOn.
	AddOnIndicator formats.NumericInteger_Length1To1 `xml:"addOnIndicator,omitempty"`

	// The product description
	ProductDescription formats.AlphaNumericString_Length1To188 `xml:"productDescription,omitempty"`
}

type ProductDateAndTimeTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDateAndTimeTypeU"`

	// Convey the begin date of a period. Format is ddmmyyyy.
	DepartureDate formats.AlphaNumericString_Length1To8 `xml:"departureDate,omitempty"`

	// Convey the begin time of a period. Format is hhmm.
	DepartureTime formats.Time24_HHMM `xml:"departureTime,omitempty"`

	// Convey the end date of a period. Format is ddmmyyyy.
	ArrivalDate formats.AlphaNumericString_Length1To8 `xml:"arrivalDate,omitempty"`

	// Convey the end time of a period. Format is hhmm.
	ArrivalTime formats.Time24_HHMM `xml:"arrivalTime,omitempty"`
}

type ProductDateAndTimeTypeU_46325C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDateAndTimeTypeU_46325C"`

	// Conveys departure date
	DepartureDate formats.AlphaNumericString_Length1To7 `xml:"departureDate,omitempty"`

	// Conveys departure time
	DepartureTime formats.Time24_HHMM `xml:"departureTime,omitempty"`

	// Conveys arrival date
	ArrivalDate formats.AlphaNumericString_Length1To7 `xml:"arrivalDate,omitempty"`

	// Conveys arrival time
	ArrivalTime formats.Time24_HHMM `xml:"arrivalTime,omitempty"`
}

type ProductDateTimeTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDateTimeTypeI"`

	DepartureDate formats.Date_DDMMYY `xml:"departureDate,omitempty"`

	DepartureTime formats.Time24_HHMM `xml:"departureTime,omitempty"`

	ArrivalDate formats.Date_DDMMYY `xml:"arrivalDate,omitempty"`

	ArrivalTime formats.Time24_HHMM `xml:"arrivalTime,omitempty"`
}

type ProductDateTimeTypeI_171495C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDateTimeTypeI_171495C"`

	// AIR segment : departure date ATX segment : requested date CAR segment : pick-up date CCR segment : pick-up date HHL segment : check-in date HTL segment : check-in date MIS segment : date for service requested SUR segment : date Trn Amtrak sgt: departure date Trn SNCF sgt: departure date TTO segment: departure date of the tour TUR segment: tour departure date INS element: departure date CRU segment: sailing departure date
	DepDate formats.Date_DDMMYY `xml:"depDate,omitempty"`

	// AIR segment : departure time SUR segment : pick-up time Trn Amtrak sgt: departure time Trn SNCF sgt: departure time
	DepTime formats.Time24_HHMM `xml:"depTime,omitempty"`

	// AIR segment : arrival date CAR segment : drop-off date CCR segment : return date HHL segment : check-out date HTL segment : check-out date TTO segment: return date of the tour INS element: return date
	ArrDate formats.Date_DDMMYY `xml:"arrDate,omitempty"`

	// AIR segment : arrival time Trn Amtrak sgt: arrival time Trn SNCF sgt: arrival time
	ArrTime formats.Time24_HHMM `xml:"arrTime,omitempty"`

	// AIR segment: day change indicator (1,2,-1) TRN Amtrak sgt: day change indicator (1,2,-1) TRN SNCF sgt: day change indicator (1,2,-1)
	DayChangeIndicator formats.NumericInteger_Length1To1 `xml:"dayChangeIndicator,omitempty"`
}

type ProductDateTimeTypeI_260882C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDateTimeTypeI_260882C"`

	// AIR segment : departure date ATX segment : requested date CAR segment : pick-up date CCR segment : pick-up date HHL segment : check-in date HTL segment : check-in date MIS segment : date for service requested SUR segment : date Trn Amtrak sgt: departure date Trn SNCF sgt: departure date TTO segment: departure date of the tour TUR segment: tour departure date INS element: departure date CRU segment: sailing departure date
	DepDate formats.Date_DDMMYY `xml:"depDate,omitempty"`

	// AIR segment : departure time SUR segment : pick-up time Trn Amtrak sgt: departure time Trn SNCF sgt: departure time
	DepTime formats.Time24_HHMM `xml:"depTime,omitempty"`

	// AIR segment : arrival date CAR segment : drop-off date CCR segment : return date HHL segment : check-out date HTL segment : check-out date TTO segment: return date of the tour INS element: return date
	ArrDate formats.Date_DDMMYY `xml:"arrDate,omitempty"`

	// AIR segment : arrival time Trn Amtrak sgt: arrival time Trn SNCF sgt: arrival time
	ArrTime formats.Time24_HHMM `xml:"arrTime,omitempty"`

	// AIR segment: day change indicator (1,2,-1) TRN Amtrak sgt: day change indicator (1,2,-1) TRN SNCF sgt: day change indicator (1,2,-1)
	DayChangeIndicator formats.NumericInteger_Length1To1 `xml:"dayChangeIndicator,omitempty"`
}

type ProductDateTimeTypeI_46338C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDateTimeTypeI_46338C"`

	// flight departure date
	DepartureDate formats.AlphaNumericString_Length1To7 `xml:"departureDate,omitempty"`

	// flight departure time
	DepartureTime formats.Time24_HHMM `xml:"departureTime,omitempty"`

	// flight arrival date
	ArrivalDate formats.AlphaNumericString_Length1To7 `xml:"arrivalDate,omitempty"`

	// flight arrival time
	ArrivalTime formats.Time24_HHMM `xml:"arrivalTime,omitempty"`
}

type ProductDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDetailsTypeI"`

	// booking class
	Designator formats.AlphaString_Length1To1 `xml:"designator,omitempty"`
}

type ProductDetailsTypeI_118108C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDetailsTypeI_118108C"`

	// Class combination
	Designator formats.AlphaString_Length1To1 `xml:"designator,omitempty"`

	// indicate availability status . coded or numeric
	AvailabilityStatus formats.AlphaNumericString_Length1To3 `xml:"availabilityStatus,omitempty"`
}

type ProductDetailsTypeI_36664C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductDetailsTypeI_36664C"`

	// Conveys the package code.
	Designator formats.AlphaNumericString_Length1To10 `xml:"designator,omitempty"`
}

type ProductFacilitiesTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductFacilitiesTypeI"`

	Entertainement formats.AlphaNumericString_Length1To3 `xml:"entertainement,omitempty"`

	// For meal, the meal codes follow the IATA meal code standard
	EntertainementDescription formats.AlphaNumericString_Length1To2 `xml:"entertainementDescription,omitempty"`

	ProductQualifier formats.AlphaNumericString_Length1To2 `xml:"productQualifier,omitempty"`

	ProductExtensionCode formats.AlphaNumericString_Length1To4 `xml:"productExtensionCode,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductIdentificationDetailsTypeI"`

	FlightNumber formats.AlphaNumericString_Length1To4 `xml:"flightNumber,omitempty"`
}

type ProductIdentificationDetailsTypeI_2786C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductIdentificationDetailsTypeI_2786C"`

	// Flight number or OPEN - ARNK, car type, transportation type  (refer to VGTVD transaction), train number, insurance provider
	Identification formats.AlphaNumericString_Length1To6 `xml:"identification,omitempty"`

	// AIR segment : class of service TRN Amtrack segment : class of service (1 or 2 chars long). TRN SNCF segment : class of service.
	ClassOfService formats.AlphaNumericString_Length1To2 `xml:"classOfService,omitempty"`

	// AIR segment : flight number alpha suffix : A, B, C, D, E. SUR segment : departure code : A or D.
	Subtype formats.AlphaString_Length1To1 `xml:"subtype,omitempty"`

	// AIR segment :  N for Night class
	Description formats.AlphaNumericString_Length1To1 `xml:"description,omitempty"`
}

type ProductIdentificationDetailsTypeI_46336C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductIdentificationDetailsTypeI_46336C"`

	// flight number or transportation code
	FlightNumber formats.AlphaNumericString_Length1To20 `xml:"flightNumber,omitempty"`

	// booking class
	BookingClass formats.AlphaNumericString_Length1To4 `xml:"bookingClass,omitempty"`
}

type ProductIdentificationDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductIdentificationDetailsTypeU"`

	// Product code
	Number formats.AlphaNumericString_Length1To16 `xml:"number,omitempty"`

	// Product Name
	Name formats.AlphaNumericString_Length1To35 `xml:"name,omitempty"`
}

type ProductIdentificationDetailsTypeU_46327C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductIdentificationDetailsTypeU_46327C"`

	// Conveys the product code
	Code formats.AlphaNumericString_Length1To32 `xml:"code,omitempty"`

	// Conveys the product type (accomodation, vehicule, transportation, cruise ...)
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Conveys the subType of a product (Chalet or Villa for accomodation, Transfert or ticket for supplementary services ...)
	SubType formats.AlphaNumericString_Length1To3 `xml:"subType,omitempty"`

	// Conveys the product description
	Description formats.AlphaNumericString_Length1To32 `xml:"description,omitempty"`
}

type ProductIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductIdentificationTypeU"`

	// product name and code to which prices data apply
	ProductData *ProductIdentificationDetailsTypeU `xml:"productData,omitempty"`
}

type ProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductInformationTypeI"`

	// Conveys the package details.
	BookingClassDetails *ProductDetailsTypeI_36664C `xml:"bookingClassDetails,omitempty"`
}

type ProductInformationTypeI_73824S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductInformationTypeI_73824S"`

	// Booking class
	BookingClassDetails *ProductDetailsTypeI `xml:"bookingClassDetails,omitempty"`
}

type ProductInformationTypeI_76271S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductInformationTypeI_76271S"`

	// Booking Class Details
	BookingClassDetails *ProductDetailsTypeI_118108C `xml:"bookingClassDetails,omitempty"`
}

type ProductTypeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductTypeDetailsTypeI"`

	// AIR segment : Electronic ticketing indicator : ET for Electronic ticket candidate SUR segment : transportation zone number Amtrack segment : Equipement code  SNCF segment : train type (3 chars code)
	Detail formats.AlphaNumericString_Length1To3 `xml:"detail,omitempty"`
}

type ProductTypeDetailsTypeI_46337C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProductTypeDetailsTypeI_46337C"`

	// sequence indicator for connection
	FlightIndicator formats.AlphaNumericString_Length1To6 `xml:"flightIndicator,omitempty"`
}

type PropertyHeaderDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A PropertyHeaderDetailsType"`

	// 1. hotel Provider name
	ProviderName formats.AlphaNumericString_Length1To35 `xml:"providerName,omitempty"`

	// 1. HHL segment:hotel Property Code (or ID) 2. HTL AY Direct Access segment: Property location
	Code formats.AlphaNumericString_Length3To3 `xml:"code,omitempty"`

	// 1. HHL segment:hotel Property name. 2. HTL AY Direct Access segment: Hotel name. Alphanumeric type due to possible numeric values in the names.
	Name formats.AlphaNumericString_Length1To40 `xml:"name,omitempty"`
}

type ProviderInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ProviderInformationType"`

	// productcode
	Code formats.AlphaNumericString_Length1To4 `xml:"code,omitempty"`

	// Product name
	Name formats.AlphaNumericString_Length1To50 `xml:"name,omitempty"`

	// Product Family Code
	ProductFamilyCode formats.AlphaNumericString_Length1To5 `xml:"productFamilyCode,omitempty"`
}

type QuantityAndActionDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityAndActionDetailsTypeU"`

	// Quantity information
	Quantity formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`

	// Conveys the status code (HK, GK ...) of a booking, a product or a ticket
	StatusCode formats.AlphaNumericString_Length1To3 `xml:"statusCode,omitempty"`
}

type QuantityAndActionDetailsTypeU_56796C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityAndActionDetailsTypeU_56796C"`

	// accommodation reservation mandatoty, optionnal, advised, not possible
	StatusCode formats.AlphaNumericString_Length1To3 `xml:"statusCode,omitempty"`
}

type QuantityAndActionTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityAndActionTypeU"`

	// Conveys quantity and status information
	QuantityActionDetails *QuantityAndActionDetailsTypeU `xml:"quantityActionDetails,omitempty"`
}

type QuantityAndActionTypeU_32609S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityAndActionTypeU_32609S"`

	// accommodation status
	AccoStatus *QuantityAndActionDetailsTypeU_56796C `xml:"accoStatus,omitempty"`
}

type QuantityDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityDetailsTypeI"`

	// A for age
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Age = number of years(default) or monthes.
	Value formats.NumericInteger_Length1To3 `xml:"value,omitempty"`
}

type QuantityDetailsTypeI_142179C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityDetailsTypeI_142179C"`

	// it will be L for Life time period
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// duration expressed in Seconds during the consumer has to do the payment
	Value formats.NumericInteger_Length1To15 `xml:"value,omitempty"`

	// SEC for duration in seconds
	Unit formats.AlphaNumericString_Length1To3 `xml:"unit,omitempty"`
}

type QuantityDetailsTypeI_198209C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityDetailsTypeI_198209C"`

	// -NOD Number of Doors -MOD Maximum number of Doors -NOS Number of Seats -MOD Number of Seats -NOB Number of Bags -VOB Volume of Boots
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Value number corresponding to the qualifier type
	Value formats.NumericInteger_Length1To15 `xml:"value,omitempty"`

	// DM3 or FT3 if applicable
	Unit formats.AlphaNumericString_Length1To3 `xml:"unit,omitempty"`
}

type QuantityDetailsTypeI_46334C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityDetailsTypeI_46334C"`

	// Quantity qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Quantity value
	Value formats.NumericInteger_Length1To3 `xml:"value,omitempty"`

	// Quantity unit
	Unit formats.AlphaNumericString_Length1To3 `xml:"unit,omitempty"`
}

type QuantityTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityTypeI"`

	// Estinated distance details
	QuantityDetails *QuantityDetailsTypeI_142179C `xml:"quantityDetails,omitempty"`
}

type QuantityTypeI_65488S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityTypeI_65488S"`

	// This composite is used to convey the quantity details
	QuantityDetails *QuantityDetailsTypeI `xml:"quantityDetails,omitempty"`
}

type QuantityType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityType"`

	// To specify an appropriate quantity.
	QuantityDetails *QuantityDetailsTypeI_46334C `xml:"quantityDetails,omitempty"`
}

type QuantityType_94558S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QuantityType_94558S"`

	// To specify an appropriate quantity.
	QuantityDetails *QuantityDetailsTypeI_142179C `xml:"quantityDetails,omitempty"`
}

type QueueDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QueueDetailsType"`

	// A (first) queue number
	QueueNum1 formats.NumericInteger_Length1To3 `xml:"queueNum1,omitempty"`

	// [2-7] characters
	QueueName formats.AlphaNumericString_Length1To7 `xml:"queueName,omitempty"`
}

type QueueType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A QueueType"`

	// Queue detail
	QueueDetail *QueueDetailsType `xml:"queueDetail,omitempty"`

	// Queue category detail
	CategoryDetail *GategoryType `xml:"categoryDetail,omitempty"`

	// date range
	DateRange *DateRangeType `xml:"dateRange,omitempty"`

	// Other queue information
	Informations *OtherInformationType `xml:"informations,omitempty"`
}

type RailLegDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RailLegDataType"`

	// Information pertaining to the train product
	TrainProductInfo *TrainProductInformationType_32331S `xml:"trainProductInfo,omitempty"`

	// Reservation Mandatory, Advised, Possible, Not Possible
	ReservableStatus *QuantityAndActionTypeU_32609S `xml:"reservableStatus,omitempty"`

	// Leg departure and arrival dates and times
	LegDateTime *StructuredDateTimeInformationType_32362S `xml:"legDateTime,omitempty"`

	// Departure station location
	DepLocation *PlaceLocationIdentificationTypeU_32347S `xml:"depLocation,omitempty"`

	// Arrival station location
	ArrLocation *PlaceLocationIdentificationTypeU_32347S `xml:"arrLocation,omitempty"`

	// leg reference: leg order within the itinerary
	LegReference *ItemNumberTypeU_33258S `xml:"legReference,omitempty"`
}

type RailSeatConfigurationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RailSeatConfigurationType"`

	// Seat space.
	SeatSpace formats.AlphaNumericString_Length2To2 `xml:"seatSpace,omitempty"`

	// Coach type.
	CoachType formats.AlphaNumericString_Length2To2 `xml:"coachType,omitempty"`

	// Seat equipment.
	SeatEquipment formats.AlphaNumericString_Length2To2 `xml:"seatEquipment,omitempty"`

	// Seat position.
	SeatPosition formats.AlphaNumericString_Length1To1 `xml:"seatPosition,omitempty"`

	// Seat direction.
	SeatDirection formats.AlphaNumericString_Length1To1 `xml:"seatDirection,omitempty"`

	// Seat deck.
	SeatDeck formats.AlphaNumericString_Length1To1 `xml:"seatDeck,omitempty"`

	// Special passenger information.
	SpecialPassengerType formats.AlphaNumericString_Length1To1 `xml:"specialPassengerType,omitempty"`
}

type RailSeatPreferencesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RailSeatPreferencesType"`

	// Selection of the type of seat request.
	SeatRequestFunction formats.AlphaNumericString_Length1To1 `xml:"seatRequestFunction,omitempty"`

	// Seat smoking zone indicator.
	SmokingIndicator formats.AlphaString_Length1To1 `xml:"smokingIndicator,omitempty"`

	// Seat class details.
	ClassDetails *ClassDetailsType_52782C `xml:"classDetails,omitempty"`

	// Seat configuration details.
	SeatConfiguration *RailSeatConfigurationType `xml:"seatConfiguration,omitempty"`

	SleeperDescription *RailSleeperDescriptionType `xml:"sleeperDescription,omitempty"`
}

type RailSeatReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RailSeatReferenceInformationType"`

	// Rail seat reference information.
	RailSeatReferenceDetails *SeatReferenceInformationType `xml:"railSeatReferenceDetails,omitempty"`
}

type RailSleeperDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RailSleeperDescriptionType"`

	// Berth deck
	BerthDeck formats.AlphaNumericString_Length2To2 `xml:"berthDeck,omitempty"`

	// Cabin position
	CabinPosition formats.AlphaNumericString_Length2To2 `xml:"cabinPosition,omitempty"`

	// Cabin share type
	CabinShareType formats.AlphaNumericString_Length2To2 `xml:"cabinShareType,omitempty"`

	// Cabin occupancy
	CabinOccupancy formats.AlphaNumericString_Length2To2 `xml:"cabinOccupancy,omitempty"`
}

type RangeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RangeDetailsTypeI"`

	// 701 for range definition
	RangeQualifier formats.AlphaNumericString_Length1To3 `xml:"rangeQualifier,omitempty"`

	// Range definition
	RangeDetails *RangeTypeI `xml:"rangeDetails,omitempty"`
}

type RangeDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RangeDetailsTypeU"`

	// Range qualifier
	RangeQualifier formats.AlphaNumericString_Length1To3 `xml:"rangeQualifier,omitempty"`

	// Range details
	RangeDetails *RangeTypeU `xml:"rangeDetails,omitempty"`
}

type RangeTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RangeTypeI"`

	// Duration qualifier: - DAY Duration in days - WE  Duration in weeks - MTH Duration in months - G Kilometers - M Mileage - A Age
	DataType formats.AlphaNumericString_Length1To3 `xml:"dataType,omitempty"`

	// Base of the Range
	Min formats.NumericInteger_Length1To3 `xml:"min,omitempty"`

	// Top of the Range
	Max formats.NumericInteger_Length1To3 `xml:"max,omitempty"`
}

type RangeTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RangeTypeU"`

	// Range data type
	DataType formats.AlphaNumericString_Length1To3 `xml:"dataType,omitempty"`

	// min Occupancy
	MinOccupancy formats.NumericInteger_Length1To2 `xml:"minOccupancy,omitempty"`

	// Occupancy maximum
	MaxOccupancy formats.NumericInteger_Length1To2 `xml:"maxOccupancy,omitempty"`
}

type RateCodeRestrictedType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RateCodeRestrictedType"`

	// 1. HHL segment: hotel  Rate code (an3) 2. For AY Direct Access segment: Rate type =  MINR, MODR, MAXR, ADVR, DAYR, SRTE
	RateCode formats.AlphaNumericString_Length1To4 `xml:"rateCode,omitempty"`
}

type RateIndicatorsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RateIndicatorsType"`

	// 1. HTL AY Direct Access segment:  Y for Yes  (rate change)
	RateChangeIndicator formats.AlphaString_Length1To1 `xml:"rateChangeIndicator,omitempty"`
}

type RateInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RateInformationDetailsType"`

	// 1. Hotel segment:  Total or daily indicator
	RatePlan formats.AlphaNumericString_Length2To2 `xml:"ratePlan,omitempty"`
}

type RateInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RateInformationTypeI"`

	// Rate Category.
	Category formats.AlphaNumericString_Length1To9 `xml:"category,omitempty"`
}

type RateInformationTypeI_198204C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RateInformationTypeI_198204C"`

	// Rate Category 002 Inclusive 006 Convention 007 Corporate 009 Government 011 Package 019 Association 020 Business 021 Consortium 022 Credential 023 Industry 024 Standard G General
	Category formats.AlphaNumericString_Length1To3 `xml:"category,omitempty"`
}

type RateInformationTypeI_50732C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RateInformationTypeI_50732C"`

	// Fare Group
	FareGroup formats.AlphaNumericString_Length1To9 `xml:"fareGroup,omitempty"`
}

type RateInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RateInformationType"`

	// Rate Price
	RatePrice *RatePriceType `xml:"ratePrice,omitempty"`

	// Rate information
	RateInfo *RateInformationDetailsType `xml:"rateInfo,omitempty"`

	// Rate indicator
	RateIndicator *RateIndicatorsType `xml:"rateIndicator,omitempty"`
}

type RatePriceType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RatePriceType"`

	// 1. Hotel segment:  Rate value 2. Hotel AY Direct Access segment: Room rate (imbedded decimal point)
	RateAmount formats.NumericDecimal_Length1To11 `xml:"rateAmount,omitempty"`
}

type RateTypesTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RateTypesTypeU"`

	// This element holds the rate code that applies to the Ferry booking.
	RateCode formats.AlphaNumericString_Length1To15 `xml:"rateCode,omitempty"`
}

type ReferenceInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInfoType"`

	// This composite is used to transmit association information
	Reference *ReferencingDetailsType_111975C `xml:"reference,omitempty"`
}

type ReferenceInfoType_25422S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInfoType_25422S"`

	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsTypeI_46317C `xml:"referenceDetails,omitempty"`
}

type ReferenceInfoType_6074S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInfoType_6074S"`

	// This composite is used to transmit association information
	Reference *ReferencingDetailsType `xml:"reference,omitempty"`
}

type ReferenceInfoType_94524S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInfoType_94524S"`

	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsType_142140C `xml:"referenceDetails,omitempty"`
}

type ReferenceInfoType_94566S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInfoType_94566S"`

	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsType_142187C `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInformationTypeI"`

	// Details of the referencing
	ReferenceDetails *ReferencingDetailsTypeI_185716C `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationTypeI_136704S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInformationTypeI_136704S"`

	// Use to convey the reference details
	ReferenceDetails *ReferencingDetailsTypeI_198199C `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationTypeI_25132S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInformationTypeI_25132S"`

	// Conveys the passenger reference.
	ReferenceDetails *ReferencingDetailsTypeI_45901C `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationTypeI_83551S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInformationTypeI_83551S"`

	// Reference details
	ReferenceDetails *ReferencingDetailsTypeI_127514C `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationTypeI_94503S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInformationTypeI_94503S"`

	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsTypeI `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInformationType"`

	// Used to specify the passenger association and the data per passanger.
	ReferenceDetails *ReferencingDetailsTypeI_17164C `xml:"referenceDetails,omitempty"`
}

type ReferenceInformationType_65487S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferenceInformationType_65487S"`

	// Used to convey the passenger tatoo or display number.
	PassengerReference *ReferencingDetailsTypeI `xml:"passengerReference,omitempty"`
}

type ReferencingDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsTypeI"`

	// Qualifier of the type of reference.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Value of the association reference
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type ReferencingDetailsTypeI_127514C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsTypeI_127514C"`

	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	Value formats.AlphaNumericString_Length1To10 `xml:"value,omitempty"`
}

type ReferencingDetailsTypeI_17164C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsTypeI_17164C"`

	// to specify the segment association
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// used to identify the segment(s) from which we want to extract the data.
	Value formats.AlphaNumericString_Length1To5 `xml:"value,omitempty"`
}

type ReferencingDetailsTypeI_185716C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsTypeI_185716C"`

	// A code which identifies the type of identifier that is used.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type ReferencingDetailsTypeI_198199C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsTypeI_198199C"`

	// Reference qualifier Amadeus codes :  OT for Other element(non name, non segment) Tatoo   PT for Passenger Tatoo   ST for Segment Tatoo   SS for Segment Tatoo+SubTatoo
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Reference number Number attributed by the Server to reference the PNR segment/element Limited to the time the PNR is worked (First retrieve - End of Transaction)
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type ReferencingDetailsTypeI_36941C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsTypeI_36941C"`

	// Coach Number
	Value formats.AlphaNumericString_Length1To5 `xml:"value,omitempty"`
}

type ReferencingDetailsTypeI_45901C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsTypeI_45901C"`

	// Qualifies the type of reference used.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Conveys the passenger sequence number.
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type ReferencingDetailsTypeI_46317C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsTypeI_46317C"`

	// Qualify the type of reference: passenger or product
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// Passenger tatoo or Product sequence number
	Value formats.AlphaNumericString_Length1To2 `xml:"value,omitempty"`
}

type ReferencingDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsType"`

	// Amadeus codes are used here.  PT for Passenger Tatoo // ST for Segment Tatoo //OT for Other element Tatoo //SS for Segment Tatoo+SubTatoo
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// reference number refers to a PNR segment/element that has this number in its related element reference segment in the same message (qualifier PT, SS, ST).
	Number formats.AlphaNumericString_Length1To5 `xml:"number,omitempty"`
}

type ReferencingDetailsType_111975C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsType_111975C"`

	// Amadeus codes are used here.  PT for Passenger Tatoo // ST for Segment Tatoo //OT for Other element Tatoo //SS for Segment Tatoo+SubTatoo
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// reference number refers to a PNR segment/element that has this number in its related element reference segment in the same message (qualifier PT, SS, ST).
	Number formats.AlphaNumericString_Length1To9 `xml:"number,omitempty"`
}

type ReferencingDetailsType_127526C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsType_127526C"`

	// Shopping Basket codes : CDS Shopping Basket Distribution record CRR Shopping Basket Reservation RecordCST Shopping Basket customer DOC Shopping Basket document FAR Shopping Basket fares and fees information FFY Shopping Basket frequent flyer information FOP Shopping Basket form of payment PRD Shopping Basket product RMK Shopping Basket remark SBK Shopping Basket (used in search results)
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Number attributed by the Server to reference the shopping basket item.
	Number formats.NumericInteger_Length1To5 `xml:"number,omitempty"`
}

type ReferencingDetailsType_142140C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsType_142140C"`

	// will have the following values: XID  Transaction identifier of the 3DS process CAAV authentication verification code for Visa AAV  authentication verification code for MasterCard PAREQ authentication message PARES authentication response message
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type ReferencingDetailsType_142187C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsType_142187C"`

	// FOID document type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// FOID document number
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"`
}

type ReferencingDetailsType_2780C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReferencingDetailsType_2780C"`

	// Amadeus codes are used here.  D for Dominant segment in a marriage  N for Non dominant segment in a marriage
	MarriageQualifier formats.AlphaNumericString_Length1To3 `xml:"marriageQualifier,omitempty"`

	// Tatoo number of the segment
	TatooNum formats.AlphaNumericString_Length1To5 `xml:"tatooNum,omitempty"`
}

type RelatedLocationOneIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RelatedLocationOneIdentificationTypeU"`

	// Conveys the embarkation port code.
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`
}

type RelatedLocationOneIdentificationTypeU_198193C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RelatedLocationOneIdentificationTypeU_198193C"`

	// Assiciated airport code.
	Code formats.AlphaNumericString_Length1To25 `xml:"code,omitempty"`

	// Associated airport code qualifier.
	Qualifier formats.AlphaNumericString_Length1To17 `xml:"qualifier,omitempty"`

	// Set to IA to indicate that the associated location code is a IATA airport or city code.
	Agency formats.AlphaNumericString_Length1To3 `xml:"agency,omitempty"`
}

type RelatedLocationOneIdentificationTypeU_45087C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RelatedLocationOneIdentificationTypeU_45087C"`

	// Area code.
	Code formats.AlphaNumericString_Length1To10 `xml:"code,omitempty"`
}

type RelatedLocationOneIdentificationTypeU_46345C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RelatedLocationOneIdentificationTypeU_46345C"`

	// location code
	Code formats.AlphaNumericString_Length1To2 `xml:"code,omitempty"`

	// location qualifier
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`
}

type RelatedLocationOneIdentificationTypeU_56455C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RelatedLocationOneIdentificationTypeU_56455C"`

	// Railway station country code
	Code formats.AlphaNumericString_Length1To2 `xml:"code,omitempty"`

	// Code type
	Qualifier formats.AlphaNumericString_Length1To2 `xml:"qualifier,omitempty"`
}

type RelatedLocationTwoIdentificationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RelatedLocationTwoIdentificationTypeU"`

	// Conveys the disembarkation port code.
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`
}

type RelatedProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RelatedProductInformationTypeI"`

	// No quantity returned
	Quantity formats.NumericInteger_Length1To3 `xml:"quantity,omitempty"`

	// see code list
	Status formats.AlphaString_Length1To2 `xml:"status,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationDetailsTypeI"`

	// 1A or Other airline record locator information  Passive segment airline code
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// 1. Record - 1A record locator or - OA record locator
	ControlNumber formats.AlphaNumericString_Length1To19 `xml:"controlNumber,omitempty"`

	// 1. Profile record locator information: Customer type:  C for Corporate  T for Traveler     F for Frequent Flyer
	ControlType formats.AlphaNumericString_Length1To1 `xml:"controlType,omitempty"`

	// 1. RR element: Date 2. SP element: Date 3. PNR header/RP line: Date of last End of transaction
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// 1. PNR header/RP line: time of last End of transaction
	Time formats.Time24_HHMM `xml:"time,omitempty"`
}

type ReservationControlInformationDetailsTypeI_16352C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationDetailsTypeI_16352C"`

	// Conveys the booking number.
	ControlNumber formats.AlphaNumericString_Length1To19 `xml:"controlNumber,omitempty"`

	// Conveys the booking number qualifier.
	ControlType formats.AlphaNumericString_Length1To1 `xml:"controlType,omitempty"`
}

type ReservationControlInformationDetailsTypeI_170724C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationDetailsTypeI_170724C"`

	// 1A or Other airline record locator information  Passive segment airline code
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// 1. Record - 1A record locator or - OA record locator  Record locator is truncated to 7 characters.
	ControlNumber formats.AlphaNumericString_Length1To19 `xml:"controlNumber,omitempty"`

	// 1. Profile record locator information: Customer type:  C for Corporate  T for Traveler     F for Frequent Flyer
	ControlType formats.AlphaNumericString_Length1To1 `xml:"controlType,omitempty"`

	// 1. RR element: Date 2. SP element: Date 3. PNR header/RP line: Date of last End of transaction
	Date formats.NumericInteger_Length6To6 `xml:"date,omitempty"`

	// 1. PNR header/RP line: time of last End of transaction
	Time formats.Time24_HHMM `xml:"time,omitempty"`
}

type ReservationControlInformationDetailsTypeI_18446C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationDetailsTypeI_18446C"`

	// this is used to specify the confirmation number meaning that the booking was succesfull.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationDetailsTypeI_198198C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationDetailsTypeI_198198C"`

	// - 1A or Other airline record locator information - Passive segment airline code
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// - 1A record locator or - OA record locator
	ControlNumber formats.AlphaNumericString_Length1To19 `xml:"controlNumber,omitempty"`

	// PNR split type.
	ControlType formats.AlphaNumericString_Length1To1 `xml:"controlType,omitempty"`

	// 1. RR element: Date 2. SP element: Date 3. PNR header/RP line: Date of lastest End of transaction
	Date formats.NumericInteger_Length6To6 `xml:"date,omitempty"`

	// 1. PNR header/RP line: time of lastest End of transaction
	Time formats.NumericInteger_Length4To4 `xml:"time,omitempty"`
}

type ReservationControlInformationDetailsTypeI_35709C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationDetailsTypeI_35709C"`

	// This element conveys the booking number which is used as a booking reference by the Ferry provider.
	ControlNumber formats.AlphaNumericString_Length1To17 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationDetailsTypeU"`

	// Conveys the tour operator code
	TourOperatorCode formats.AlphaNumericString_Length1To4 `xml:"tourOperatorCode,omitempty"`

	// Conveys the reservation control number qualifier
	ReservationControlNumberQual formats.AlphaNumericString_Length1To3 `xml:"reservationControlNumberQual,omitempty"`

	// Conveys the reservation control number. Can have up to 32 chars. for Tour Server
	ReservationControlNumber formats.AlphaNumericString_Length1To32 `xml:"reservationControlNumber,omitempty"`
}

type ReservationControlInformationDetailsTypeU_55378C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationDetailsTypeU_55378C"`

	// The individual Passenger confirmation number in the Provider database.
	Value formats.AlphaNumericString_Length1To22 `xml:"value,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationTypeI"`

	// To specify the confirmation number in case the booking was succesfull
	Reservation *ReservationControlInformationDetailsTypeI_18446C `xml:"reservation,omitempty"`
}

type ReservationControlInformationTypeI_115879S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationTypeI_115879S"`

	// Reservation Information
	Reservation *ReservationControlInformationDetailsTypeI_170724C `xml:"reservation,omitempty"`
}

type ReservationControlInformationTypeI_136703S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationTypeI_136703S"`

	// Reservation Information
	Reservation *ReservationControlInformationDetailsTypeI_198198C `xml:"reservation,omitempty"`
}

type ReservationControlInformationTypeI_20153S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationTypeI_20153S"`

	// Provides details of the Ferry booking number. The booking number is a unique reference per provider per booking in the provider system. As such, it is stored in the PNR in all the legs of the same booking and it is used in the Ferry PNR indexing.
	Reservation *ReservationControlInformationDetailsTypeI_35709C `xml:"reservation,omitempty"`
}

type ReservationControlInformationTypeI_87792S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationTypeI_87792S"`

	// Reservation Information
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}

type ReservationControlInformationTypeI_8957S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationTypeI_8957S"`

	// Cruise booking reference.
	Reservation *ReservationControlInformationDetailsTypeI_16352C `xml:"reservation,omitempty"`
}

type ReservationControlInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationTypeU"`

	// Conveys the reservation control Id
	ReservationControlId *ReservationControlInformationDetailsTypeU `xml:"reservationControlId,omitempty"`
}

type ReservationControlInformationTypeU_31804S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationControlInformationTypeU_31804S"`

	// The reference to the Provider Database
	ReferenceDetails *ReservationControlInformationDetailsTypeU_55378C `xml:"referenceDetails,omitempty"`
}

type ReservationSecurityInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationSecurityInformationType"`

	// Responsibility Information
	ResponsibilityInformation *ResponsibilityInformationType `xml:"responsibilityInformation,omitempty"`

	// Ticket Information
	QueueingInformation *TicketInformationType_5120C `xml:"queueingInformation,omitempty"`

	// 1. PNR Header: Pseudo City Code (not in the CRT display) AGY for Travel agency EHD for First level Help Desk DAP for Data processing center / Amadeus Help Desk Nice SEC for Security administrator WZ for AIS security administrator
	CityCode formats.AlphaString_Length3To3 `xml:"cityCode,omitempty"`

	// Second RP line information
	SecondRpInformation *SecondRpLineInformationType `xml:"secondRpInformation,omitempty"`
}

type ReservationSecurityInformationType_167761S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ReservationSecurityInformationType_167761S"`

	// Responsibility Information
	ResponsibilityInformation *ResponsibilityInformationType_6835C `xml:"responsibilityInformation,omitempty"`

	// Ticket Information
	QueueingInformation *TicketInformationType_5120C `xml:"queueingInformation,omitempty"`

	// 1. PNR Header: Pseudo City Code (not in the CRT display) AGY for Travel agency EHD for First level Help Desk DAP for Data processing center / Amadeus Help Desk Nice SEC for Security administrator WZ for AIS security administrator
	CityCode formats.AlphaString_Length3To3 `xml:"cityCode,omitempty"`

	// Second RP line information
	SecondRpInformation *SecondRpLineInformationType_237255C `xml:"secondRpInformation,omitempty"`
}

type ResponseIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ResponseIdentificationType"`

	// Transaction identifier   Field 62.2  Official definition: Visa-generated identifier that is unique for each original transaction. The transaction identifier (TID) is a key element that links original authorization requests to subsequent messages, such as reversals.
	TransacIdentifier formats.AlphaNumericString_Length1To15 `xml:"transacIdentifier,omitempty"`

	// Validation code    Field 62.3
	ValidationCode formats.AlphaNumericString_Length1To4 `xml:"validationCode,omitempty"`

	// Gateway Transaction Identifier - Banknet reference number   Field 62.17 - Position 8-13
	BanknetRefNumber formats.AlphaNumericString_Length6To9 `xml:"banknetRefNumber,omitempty"`

	// Gateway Transaction Identifier - Banknet date in mmdd format  Field 62.17 - Position 1-4
	BanknetDate formats.AlphaNumericString_Length4To4 `xml:"banknetDate,omitempty"`
}

type ResponsibilityInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ResponsibilityInformationType"`

	// Type of PNR element: - RR for Associated Cross Reference Record - SP for Split Party - RP for PNR Header line
	TypeOfPnrElement formats.AlphaNumericString_Length2To2 `xml:"typeOfPnrElement,omitempty"`

	// 1. RR element: 2. SP element: 3. PNR Header:Agent initials and duty code (eg: AASU)
	AgentId formats.AlphaNumericString_Length4To4 `xml:"agentId,omitempty"`

	// 1. RR element office that copied the PNR 2. SP element: office that split the PNR 3. PNR Header: office responsibility or - OA office (City code + OA code)  which is 5 chars long
	OfficeId formats.AlphaNumericString_Length1To9 `xml:"officeId,omitempty"`

	// ATA/IATA reference number assigned to a travel agent
	IataCode formats.NumericInteger_Length8To8 `xml:"iataCode,omitempty"`
}

type ResponsibilityInformationType_6835C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ResponsibilityInformationType_6835C"`

	// Type of PNR element: - RR for Associated Cross Reference Record - SP for Split Party - RP for PNR Header line
	TypeOfPnrElement formats.AlphaNumericString_Length2To2 `xml:"typeOfPnrElement,omitempty"`

	// 1. RR element: 2. SP element: 3. PNR Header:Agent initials and duty code (eg: AASU)
	AgentId formats.AlphaNumericString_Length4To4 `xml:"agentId,omitempty"`

	// 1. RR element office that copied the PNR 2. SP element: office that split the PNR 3. PNR Header: office responsibility or - OA office (City code + OA code)  which is 5 chars long
	OfficeId formats.AlphaNumericString_Length1To9 `xml:"officeId,omitempty"`

	// ATA/IATA reference number assigned to a travel agent
	IataCode formats.NumericInteger_Length1To9 `xml:"iataCode,omitempty"`
}

type RoomDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RoomDetailsType"`

	// 1. room Occupancy
	Occupancy formats.NumericInteger_Length1To1 `xml:"occupancy,omitempty"`

	// room Type
	TypeCode formats.AlphaNumericString_Length1To4 `xml:"typeCode,omitempty"`
}

type RuleDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RuleDetailsTypeU"`

	// This data element is used to identify the type of rule (see data mapping).
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// This data element is used to convey the afternoon time which is the maximum time to check-in. ex. if equal to 1, it means that the room is kept until 1PM.
	Quantity formats.NumericInteger_Length1To1 `xml:"quantity,omitempty"`

	// This data element is used to specify that the Maximim check-in time is given in hours.
	QuantityUnit formats.AlphaNumericString_Length1To3 `xml:"quantityUnit,omitempty"`
}

type RuleDetailsTypeU_198224C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RuleDetailsTypeU_198224C"`

	// Coded rule type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// quantity (if applicable)
	Quantity formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"`

	// DAY for Day HOR for Hour (if applicable)
	QuantityUnit formats.AlphaNumericString_Length1To3 `xml:"quantityUnit,omitempty"`

	// Deposit Information: - BRE Before Rental - AFT After Booking Pickup Information - MAX Maximum Days Rental - MIN Minimum Days Rental One Way Information: - 009 for One Way Allowed - 005 for One Way not Allowed - 006 for Restricted One Way Allowed
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Day of the week (Monday=1, Sunday=7)
	DaysOfOperation formats.AlphaNumericString_Length1To7 `xml:"daysOfOperation,omitempty"`

	// Rule amount (if applicable)
	Amount formats.NumericInteger_Length1To35 `xml:"amount,omitempty"`

	// Rule currency amount (if applicable)
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type RuleInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RuleInformationTypeU"`

	// This composite is used to convey the rules details.
	RuleDetails *RuleDetailsTypeU `xml:"ruleDetails,omitempty"`

	// This composite is used to indicate that the rule have been entered manually by the agent.
	RuleStatusDetails *RuleStatusTypeU `xml:"ruleStatusDetails,omitempty"`
}

type RuleInformationTypeU_136720S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RuleInformationTypeU_136720S"`

	// Rule details
	RuleDetails *RuleDetailsTypeU_198224C `xml:"ruleDetails,omitempty"`

	// Associated Rule Text
	RuleText *RuleTextTypeU `xml:"ruleText,omitempty"`
}

type RuleStatusTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RuleStatusTypeU"`

	// This data element specifies the rule concerned by this information.
	StatusType formats.AlphaNumericString_Length1To3 `xml:"statusType,omitempty"`

	// This data element is used to indicate that the condition have been entered manually by the travel agent and is not coming from the supplier data.
	ProcessIndicator formats.AlphaNumericString_Length1To3 `xml:"processIndicator,omitempty"`
}

type RuleTextTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A RuleTextTypeU"`

	// Coded rule type
	TextType formats.AlphaNumericString_Length1To3 `xml:"textType,omitempty"`

	// Rule Information
	FreeText formats.AlphaNumericString_Length1To55 `xml:"freeText,omitempty"`
}

type SeatReferenceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SeatReferenceInformationType"`

	// Coach number.
	CoachNumber formats.AlphaNumericString_Length1To3 `xml:"coachNumber,omitempty"`

	// Deck number.
	DeckNumber formats.AlphaNumericString_Length1To3 `xml:"deckNumber,omitempty"`

	// Seat number.
	SeatNumber formats.AlphaNumericString_Length1To3 `xml:"seatNumber,omitempty"`
}

type SeatRequestParametersTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SeatRequestParametersTypeI"`

	// Details of the seat
	GenericDetails *GenericDetailsTypeI `xml:"genericDetails,omitempty"`
}

type SecondRpLineInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SecondRpLineInformationType"`

	// Creation office
	CreationOfficeId formats.AlphaNumericString_Length1To9 `xml:"creationOfficeId,omitempty"`

	// Creation agent sine/queue category (eg: 1234AA)
	AgentSignature formats.AlphaNumericString_Length1To6 `xml:"agentSignature,omitempty"`

	// PNR creation date
	CreationDate formats.Date_DDMMYY `xml:"creationDate,omitempty"`

	// ATA/IATA number assigned to a travel agent
	CreatorIataCode formats.NumericInteger_Length8To8 `xml:"creatorIataCode,omitempty"`

	// PNR creation time
	CreationTime formats.Time24_HHMM `xml:"creationTime,omitempty"`
}

type SecondRpLineInformationType_237255C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SecondRpLineInformationType_237255C"`

	// Creation office
	CreationOfficeId formats.AlphaNumericString_Length1To9 `xml:"creationOfficeId,omitempty"`

	// Creation agent sine/queue category (eg: 1234AA)
	AgentSignature formats.AlphaNumericString_Length1To6 `xml:"agentSignature,omitempty"`

	// PNR creation date
	CreationDate formats.Date_DDMMYY `xml:"creationDate,omitempty"`

	// ATA/IATA number assigned to a travel agent
	CreatorIataCode formats.NumericInteger_Length1To9 `xml:"creatorIataCode,omitempty"`

	// PNR creation time
	CreationTime formats.Time24_HHMM `xml:"creationTime,omitempty"`
}

type SecurityInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SecurityInformationType"`

	// Date of creation
	CreationDate formats.AlphaNumericString_Length6To6 `xml:"creationDate,omitempty"`

	// Agent initials and duty code as in Originator informations  (eg: AASU)
	AgentCode formats.AlphaNumericString_Length4To4 `xml:"agentCode,omitempty"`

	// Office Id of creation/update
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`
}

type SegmentCabinIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SegmentCabinIdentificationType"`

	// Cabin class designator
	CabinCode formats.AlphaString_Length1To1 `xml:"cabinCode,omitempty"`
}

type SegmentGroupingInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SegmentGroupingInformationType"`

	// Type of segment grouping:  Mxx for Marriage (see codeset)  CNX for Connection
	GroupingCode formats.AlphaNumericString_Length1To3 `xml:"groupingCode,omitempty"`

	// transmit the list of segments participating in one marriage or segments that are connected.
	MarriageDetail *ReferencingDetailsType_2780C `xml:"marriageDetail,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SelectionDetailsInformationTypeI"`

	// see code list
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`

	// CRU segment - occurrence 1 : Duration of the cruise (saling length)  expressed in days.
	OptionInformation formats.AlphaNumericString_Length1To35 `xml:"optionInformation,omitempty"`
}

type SelectionDetailsInformationTypeI_198215C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SelectionDetailsInformationTypeI_198215C"`

	// - P6 for seamless availability
	Option formats.AlphaNumericString_Length1To3 `xml:"option,omitempty"`
}

type SelectionDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SelectionDetailsTypeI"`

	// Booking access type requested
	SelectionDetails *SelectionDetailsInformationTypeI_198215C `xml:"selectionDetails,omitempty"`
}

type SelectionDetailsTypeI_2067S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SelectionDetailsTypeI_2067S"`

	// Only the first occurrence of the composite is mandatory. Up to 10 occurrences of the composite.
	Selection *SelectionDetailsInformationTypeI `xml:"selection,omitempty"`
}

type SequenceDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SequenceDetailsTypeU"`

	ActionRequest formats.AlphaNumericString_Length1To3 `xml:"actionRequest,omitempty"`

	SequenceDetails *SequenceInformationTypeU_24073C `xml:"sequenceDetails,omitempty"`
}

type SequenceDetailsTypeU_94494S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SequenceDetailsTypeU_94494S"`

	// Sequence Information
	SequenceDetails *SequenceInformationTypeU `xml:"sequenceDetails,omitempty"`
}

type SequenceInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SequenceInformationTypeU"`

	// Sequence number of the Mean Of Payment in the FOP line. There can be up to 3 New MOP and 3 Old MOP in a FOP line.  Old Fops are considered as freeflow text.
	Number formats.NumericInteger_Length1To10 `xml:"number,omitempty"`
}

type SequenceInformationTypeU_24073C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SequenceInformationTypeU_24073C"`

	Number formats.AlphaNumericString_Length1To10 `xml:"number,omitempty"`
}

type ShipIdentificationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ShipIdentificationDetailsType"`

	// Used to convey the ship code as in the Cruise specific database ship's table.
	Code formats.AlphaNumericString_Length1To2 `xml:"code,omitempty"`

	// Used to convey the ship name as in the Cruise specific database ship's table.
	Name formats.AlphaNumericString_Length1To30 `xml:"name,omitempty"`

	// Used to convey the cruise line provider code for the sailing ship.
	CruiseLineCode formats.AlphaNumericString_Length1To3 `xml:"cruiseLineCode,omitempty"`
}

type ShipIdentificationDetailsType_45069C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ShipIdentificationDetailsType_45069C"`

	// Used to convey the ship name.
	Name formats.AlphaNumericString_Length1To10 `xml:"name,omitempty"`
}

type ShipIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ShipIdentificationType"`

	// Detailed information for the sailing ship.
	ShipDetails *ShipIdentificationDetailsType_45069C `xml:"shipDetails,omitempty"`
}

type ShipIdentificationType_8952S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ShipIdentificationType_8952S"`

	// Detailed information for the sailing ship.
	ShipDetails *ShipIdentificationDetailsType `xml:"shipDetails,omitempty"`
}

type SpecialRequirementsDataDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SpecialRequirementsDataDetailsTypeI"`

	// Seat number + row (seat SSR)  Number of seats (Group seat SSR)
	Data formats.AlphaNumericString_Length1To4 `xml:"data,omitempty"`

	// Refers a Traveller / Reference number for association purpose
	CrossRef formats.AlphaNumericString_Length1To5 `xml:"crossRef,omitempty"`

	// 3 occurrences may be used for in Amadeus seat SSR to indicate: 1. Smoking/no smoking  2. 1st area preference   3. 2nd area preference or passenger type
	SeatType formats.AlphaNumericString_Length1To2 `xml:"seatType,omitempty"`
}

type SpecialRequirementsDataDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SpecialRequirementsDataDetailsType"`

	// The seat number
	SeatNumber formats.AlphaNumericString_Length1To10 `xml:"seatNumber,omitempty"`

	// type of the seat
	SeatCharacteristic formats.AlphaNumericString_Length1To2 `xml:"seatCharacteristic,omitempty"`
}

type SpecialRequirementsDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SpecialRequirementsDetailsTypeI"`

	// Special requirements type details
	Ssr *SpecialRequirementsTypeDetailsTypeI `xml:"ssr,omitempty"`

	// Group seat SSR cannot ask for specific seats but only smoking and/or non-smoking (see Group seat SSR). the maximum repetitions here are 9 seats (1 per passenger of non-group PNR).
	Ssrb *SpecialRequirementsDataDetailsTypeI `xml:"ssrb,omitempty"`
}

type SpecialRequirementsDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SpecialRequirementsDetailsType"`

	// To specify the Seat Number.
	SeatDetails *SpecialRequirementsDataDetailsType `xml:"seatDetails,omitempty"`
}

type SpecialRequirementsTypeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SpecialRequirementsTypeDetailsTypeI"`

	// ATA/IATA defined Special Service Requirement code.  (refer to IATA AIRIMP documentation)
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// Use defined code or an ATA/IATA defined action code (See AIRIMP 7.1.2/7.1.3/7.1.4/8.14.1 (as bilaterally agreed), SIPP 105.170.1.1).
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Number of services requested
	Quantity formats.NumericInteger_Length1To3 `xml:"quantity,omitempty"`

	// Airline code or YY
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Seat Special service request  or  Frequent Flyer SSR.
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// 1. Seat SSR Processing indicator, coded  PS for Partial segment  indicator
	ProcessingIndicator formats.AlphaNumericString_Length1To3 `xml:"processingIndicator,omitempty"`

	// Board point
	Boardpoint formats.AlphaString_Length3To3 `xml:"boardpoint,omitempty"`

	// Off point
	Offpoint formats.AlphaString_Length3To3 `xml:"offpoint,omitempty"`

	// Free flow of the SSR that can be up to 127 chars long, therefore split on two 4440 (70 + 57)
	FreeText formats.AlphaNumericString_Length1To70 `xml:"freeText,omitempty"`
}

type SpecificDataInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SpecificDataInformationTypeI"`

	// This composite holds information about the element to which it applies.  For the Ferry business, this element specifies that the type of information conveyed is about the animal(s) attached to a Ferry booking.
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation,omitempty"`

	// Details and description of the conveyed information.
	DataInformation *DataInformationTypeI `xml:"dataInformation,omitempty"`
}

type SpecificTravellerDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SpecificTravellerDetailsType"`

	// passenger type indicator
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	TravellerSurname formats.AlphaString_Length1To60 `xml:"travellerSurname,omitempty"`

	TravellerGivenName formats.AlphaString_Length1To60 `xml:"travellerGivenName,omitempty"`

	// TravellerReferenceNumber
	TravellerReferenceNumber formats.AlphaNumericString_Length1To10 `xml:"travellerReferenceNumber,omitempty"`

	// birthdate or age of passenger
	PassengerBirthdate formats.AlphaNumericString_Length1To8 `xml:"passengerBirthdate,omitempty"`
}

type SpecificVisaLinkCreditCardInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SpecificVisaLinkCreditCardInformationType"`

	// ISO8583 specific info
	MsgRef *MessageReferenceType `xml:"msgRef,omitempty"`

	// Response identification
	RespIdentification *ResponseIdentificationType `xml:"respIdentification,omitempty"`
}

type StationInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StationInformationTypeI"`

	DepartTerminal formats.AlphaNumericString_Length1To2 `xml:"departTerminal,omitempty"`
}

type StationInformationTypeI_119771C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StationInformationTypeI_119771C"`

	Terminal formats.AlphaNumericString_Length1To2 `xml:"terminal,omitempty"`
}

type StatusDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsTypeI"`

	// Indicator name.
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type StatusDetailsTypeI_185722C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsTypeI_185722C"`

	// Status of the entity
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Qualifies the status
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type StatusDetailsTypeI_20684C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsTypeI_20684C"`

	// Coded identifying type of Information
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Data information upon qualifier.
	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`
}

type StatusDetailsTypeI_37285C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsTypeI_37285C"`

	// Indicates of the reservation is modifiable directly in the 1A system
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type StatusDetailsTypeI_57035C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsTypeI_57035C"`

	// Indicator name.
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Contains "MOD" if the PNR has been modifed since it has been retrieved
	IsPNRModifDuringTrans formats.AlphaNumericString_Length1To3 `xml:"isPNRModifDuringTrans,omitempty"`
}

type StatusDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsType"`

	// FRA for fraud screening
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// This data element is used to indicate if risk management must be performed at authorization time: - Y means risk management data will be appended to author; - N means risk management data will not be appended;
	Action formats.AlphaNumericString_Length1To3 `xml:"action,omitempty"`
}

type StatusDetailsType_148479C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsType_148479C"`

	// list of status/qualifiers Either His for Historical or     Crt for Current
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type StatusDetailsType_148716C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsType_148716C"`

	// Indicator name.
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type StatusDetailsType_183347C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusDetailsType_183347C"`

	// list of status/qualifiers Either His for Historical or     Crt for Current
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Conveys any additional data necessary to qualify the indicator
	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`
}

type StatusTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusTypeI"`

	// Provides a set of coded characteristics of the customer.
	StatusDetails *StatusDetailsTypeI_185722C `xml:"statusDetails,omitempty"`
}

type StatusTypeI_13270S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusTypeI_13270S"`

	// Provides information on the type of fop.
	StatusDetails *StatusDetailsTypeI_20684C `xml:"statusDetails,omitempty"`

	OtherStatusDetails *StatusDetailsTypeI_20684C `xml:"otherStatusDetails,omitempty"`
}

type StatusTypeI_20923S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusTypeI_20923S"`

	// Status information
	StatusDetails *StatusDetailsTypeI_37285C `xml:"statusDetails,omitempty"`
}

type StatusTypeI_32775S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusTypeI_32775S"`

	// Contains general indicators relative to the state of the PNR
	StatusDetails *StatusDetailsTypeI_57035C `xml:"statusDetails,omitempty"`
}

type StatusTypeI_33257S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusTypeI_33257S"`

	// indicates an open segment
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails,omitempty"`
}

type StatusType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusType"`

	// STATUS DETAILS
	StatusInformation *StatusDetailsType_183347C `xml:"statusInformation,omitempty"`
}

type StatusType_94568S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusType_94568S"`

	// will we perform the fraud screening ?
	StatusInformation *StatusDetailsType `xml:"statusInformation,omitempty"`
}

type StatusType_99582S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusType_99582S"`

	// STATUS DETAILS
	StatusInformation *StatusDetailsType_148479C `xml:"statusInformation,omitempty"`
}

type StatusType_99946S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StatusType_99946S"`

	// STATUS DETAILS
	StatusInformation *StatusDetailsType_148716C `xml:"statusInformation,omitempty"`
}

type StructuredAddressInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredAddressInformationType"`

	// Following values are : CY for Company NA for Name L1 for Address line 1 L2 for Address line 2 PO for P.O. BOX ZP for Postal code CI for City ST for State CO for Country
	Option formats.AlphaNumericString_Length1To2 `xml:"option,omitempty"`

	// Alphanumeric information related to the level code.  Each code has its own max length, an..50 corresponds to the max length among.
	OptionText formats.AlphaNumericString_Length1To50 `xml:"optionText,omitempty"`
}

type StructuredAddressType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredAddressType"`

	// Information type, coded  2  for billing address  P08  for general mailing address  P19  for miscellaneous mailing address  P24  for home mailing address  P25  for delivery mailing address
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// Structured Address
	Address *StructuredAddressInformationType `xml:"address,omitempty"`
}

type StructuredDateTimeInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_17997C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_20644S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_20644S"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_36775C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_20645S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_20645S"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_36777C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_21109S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_21109S"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_35730C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_24436S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_24436S"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_44876C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_25444S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_25444S"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_44876C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_27086S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_27086S"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_16347C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_32362S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_32362S"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Departure or arrival date and time.
	DateTime *StructuredDateTimeType_56472C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_83553S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_83553S"`

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_94516S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_94516S"`

	// This data element can be used to provide the semantic of the information provided. Examples : - LT : date and time corresponding to Authorization message built - T : date and time corresponding to Authorization message sent - AR : date and time corresponding to Authorization message receipt
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Indicate if the time is expressed in UTC or in local time mode ( Codes U and L ). In the last case, the time zone information can be provided in the composite C89K.
	TimeMode formats.AlphaNumericString_Length1To3 `xml:"timeMode,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_142129C `xml:"dateTime,omitempty"`

	// Reference : IATA SSIM Appendix F If it is not provided, the time is considered to be given in UTC.
	TimeZoneInfo *TimeZoneIinformationType `xml:"timeZoneInfo,omitempty"`
}

type StructuredDateTimeInformationType_94559S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_94559S"`

	// This data element is used to provide the semantic of the date information provided.  Examples : - GMT Transaction date - Local Transaction date ... Default being L local date and time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_142180C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeInformationType_94567S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeInformationType_94567S"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType_142188C `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType"`

	// Year number.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour formats.Hour_hH `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.Minute_mM `xml:"minutes,omitempty"`

	// Seconds between 0 and 59
	Seconds formats.NumericInteger_Length1To2 `xml:"seconds,omitempty"`
}

type StructuredDateTimeType_142129C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_142129C"`

	// Year number.
	Year formats.NumericInteger_Length4To4 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.NumericInteger_Length1To2 `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.NumericInteger_Length1To2 `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour formats.NumericInteger_Length1To2 `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.NumericInteger_Length1To2 `xml:"minutes,omitempty"`

	// Seconds between 0 and 59
	Seconds formats.NumericInteger_Length1To2 `xml:"seconds,omitempty"`

	// Milliseconds between 0 and 999.
	Milliseconds formats.NumericInteger_Length1To3 `xml:"milliseconds,omitempty"`
}

type StructuredDateTimeType_142180C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_142180C"`

	// Year number.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour formats.Hour_hH `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.Minute_mM `xml:"minutes,omitempty"`

	// Seconds between 0 and 59
	Seconds formats.NumericInteger_Length1To2 `xml:"seconds,omitempty"`

	// Milliseconds between 0 and 999.
	Milliseconds formats.NumericInteger_Length1To3 `xml:"milliseconds,omitempty"`
}

type StructuredDateTimeType_142188C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_142188C"`

	// Year number.
	Year formats.NumericInteger_Length4To4 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.NumericInteger_Length1To2 `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.NumericInteger_Length1To2 `xml:"day,omitempty"`
}

type StructuredDateTimeType_16347C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_16347C"`

	// Year number.
	Year formats.NumericInteger_Length1To4 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type StructuredDateTimeType_17997C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_17997C"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type StructuredDateTimeType_18725C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_18725C"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.NumericInteger_Length1To6 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour formats.NumericInteger_Length1To6 `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.NumericInteger_Length1To6 `xml:"minutes,omitempty"`
}

type StructuredDateTimeType_198200C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_198200C"`

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.NumericInteger_Length4To4 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.NumericInteger_Length1To2 `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.NumericInteger_Length1To2 `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour formats.NumericInteger_Length1To2 `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.NumericInteger_Length1To2 `xml:"minutes,omitempty"`
}

type StructuredDateTimeType_198234C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_198234C"`

	// Hour between 0 and 23
	Hour formats.NumericInteger_Length1To2 `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.NumericInteger_Length1To2 `xml:"minutes,omitempty"`
}

type StructuredDateTimeType_35730C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_35730C"`

	// Hour between 0 and 23
	Hour formats.Hour_hH `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.Minute_mM `xml:"minutes,omitempty"`
}

type StructuredDateTimeType_36775C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_36775C"`

	// Year number.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour formats.Hour_hH `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.Minute_mM `xml:"minutes,omitempty"`

	// Seconds between 0 and 59
	Seconds formats.NumericInteger_Length1To2 `xml:"seconds,omitempty"`
}

type StructuredDateTimeType_36777C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_36777C"`

	// Year number.
	Year formats.NumericInteger_Length4To4 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type StructuredDateTimeType_44876C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_44876C"`

	// Year number.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type StructuredDateTimeType_56472C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredDateTimeType_56472C"`

	// Year number.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`

	// Hour between 0 and 23
	Hour formats.Hour_hH `xml:"hour,omitempty"`

	// Minutes between 0 and 59
	Minutes formats.Minute_mM `xml:"minutes,omitempty"`
}

type StructuredPeriodInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredPeriodInformationType"`

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_17997C `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_17997C `xml:"endDateTime,omitempty"`
}

type StructuredPeriodInformationType_11026S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredPeriodInformationType_11026S"`

	// This data element can be used to provide the semantic of the information provided.
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Indicate the time is expressed in local time mode.
	TimeMode formats.AlphaNumericString_Length1To3 `xml:"timeMode,omitempty"`

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_18725C `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_18725C `xml:"endDateTime,omitempty"`
}

type StructuredPeriodInformationType_136705S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredPeriodInformationType_136705S"`

	// DDT Drop-off Date and Time OCH Opening and Closing hours PDT Pickup Date and Time PKT Early and Late Pickup time RTT Early and Late Return time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Indicate if the time is expressed in UTC or in local time mode ( Codes U and L ). In the last case, the time zone information can be provided in the composite C89K.
	TimeMode formats.AlphaNumericString_Length1To3 `xml:"timeMode,omitempty"`

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_198200C `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_198200C `xml:"endDateTime,omitempty"`
}

type StructuredPeriodInformationType_136724S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredPeriodInformationType_136724S"`

	// DDT Drop-off Date and Time OCH Opening and Closing hours PDT Pickup Date and Time PKT Early and Late Pickup time RTT Early and Late Return time
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic,omitempty"`

	// Indicate if the time is expressed in UTC or in local time mode ( Codes U and L ). In the last case, the time zone information can be provided in the composite C89K.
	TimeMode formats.AlphaNumericString_Length1To3 `xml:"timeMode,omitempty"`

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_198234C `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_198234C `xml:"endDateTime,omitempty"`

	// It is used with a period to give a restriction for days impacted. It permits for example to indicate on which days, a flight operates.
	Frequency *FrequencyType `xml:"frequency,omitempty"`
}

type StructuredPeriodInformationType_8955S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredPeriodInformationType_8955S"`

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType_16347C `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType_16347C `xml:"endDateTime,omitempty"`
}

type StructuredTelephoneNumberType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredTelephoneNumberType"`

	// Telephone number
	TelephoneNumber formats.AlphaNumericString_Length1To32 `xml:"telephoneNumber,omitempty"`
}

type StructuredTelephoneNumberType_198214C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredTelephoneNumberType_198214C"`

	// Telephone number
	TelephoneNumber formats.AlphaNumericString_Length1To20 `xml:"telephoneNumber,omitempty"`
}

type StructuredTelephoneNumberType_48448C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A StructuredTelephoneNumberType_48448C"`

	// International dial code
	InternationalDialCode formats.AlphaNumericString_Length1To10 `xml:"internationalDialCode,omitempty"`

	// Local prefix code
	LocalPrefixCode formats.AlphaNumericString_Length1To10 `xml:"localPrefixCode,omitempty"`

	// Area code
	AreaCode formats.AlphaNumericString_Length1To25 `xml:"areaCode,omitempty"`

	// Telephone number
	TelephoneNumber formats.AlphaNumericString_Length1To25 `xml:"telephoneNumber,omitempty"`
}

type SystemDetailsInfoType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SystemDetailsInfoType"`

	// BCS distribution channel
	CascadingSystem *SystemDetailsTypeI_46415C `xml:"cascadingSystem,omitempty"`
}

type SystemDetailsInfoType_33158S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SystemDetailsInfoType_33158S"`

	// POS airline
	DeliveringSystem *SystemDetailsTypeI_57708C `xml:"deliveringSystem,omitempty"`
}

type SystemDetailsInfoType_94569S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SystemDetailsInfoType_94569S"`

	// LNIATA of the agent.
	WorkstationId formats.AlphaNumericString_Length1To25 `xml:"workstationId,omitempty"`

	// System delivering the shopper session ID
	DeliveringSystem *SystemDetailsTypeI `xml:"deliveringSystem,omitempty"`
}

type SystemDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SystemDetailsTypeI"`

	// will convey the name of the company ex: OPODO
	CompanyId formats.AlphaNumericString_Length1To35 `xml:"companyId,omitempty"`
}

type SystemDetailsTypeI_46415C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SystemDetailsTypeI_46415C"`

	// contains the distribution channel data. It is the concatenation of "DCD" + [access type] + [product] + [sub-product] access type, product and sub-product are represented on 3 chars.
	CompanyId formats.AlphaNumericString_Length1To12 `xml:"companyId,omitempty"`
}

type SystemDetailsTypeI_57708C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A SystemDetailsTypeI_57708C"`

	// Corporate Code
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Pseudo City Code
	LocationId formats.AlphaNumericString_Length1To3 `xml:"locationId,omitempty"`
}

type TariffInformationDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationDetailsTypeI"`

	// A unique rate product identifier.
	RateType formats.AlphaNumericString_Length1To20 `xml:"rateType,omitempty"`

	// This field is used to convey the amount.
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// This field is used to convey the currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// This data element is used to convey the rate plan (Daily or total indicator).
	RatePlanIndicator formats.AlphaNumericString_Length1To3 `xml:"ratePlanIndicator,omitempty"`

	// This data element is used to convey the rate amount type.
	AmountType formats.AlphaNumericString_Length1To3 `xml:"amountType,omitempty"`

	// This data element is used to specify the fact that a rate change occurs during the period of the stay. If the is a change the value is * (for YES)
	RateChangeIndicator formats.AlphaNumericString_Length1To3 `xml:"rateChangeIndicator,omitempty"`
}

type TariffInformationDetailsTypeI_198216C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationDetailsTypeI_198216C"`

	// CNV for converted Rate
	RateChangeIndicator formats.AlphaNumericString_Length1To3 `xml:"rateChangeIndicator,omitempty"`
}

type TariffInformationDetailsTypeI_39533C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationDetailsTypeI_39533C"`

	// Net premium
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// currency of the total price and net premium i.e in EUR/ USD
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// general indicator
	AmountType formats.AlphaNumericString_Length1To3 `xml:"amountType,omitempty"`

	// Total amount of the insurance element.
	TotalAmount formats.NumericDecimal_Length1To18 `xml:"totalAmount,omitempty"`
}

type TariffInformationDetailsTypeI_50731C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationDetailsTypeI_50731C"`

	// Fare Basis Code
	FareBasisCode formats.AlphaNumericString_Length1To8 `xml:"fareBasisCode,omitempty"`

	// Fare Base amount
	FareBaseAmount formats.NumericDecimal_Length1To18 `xml:"fareBaseAmount,omitempty"`

	// This field is used to convey the currency
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`
}

type TariffInformationDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationDetailsTypeU"`

	// Prive value. The value conveyed equals 100 times the original value in order to avoid transporting decimal placement information.
	PriceAmount formats.NumericDecimal_Length1To18 `xml:"priceAmount,omitempty"`

	// Currency code.
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// Gives the type of amount.
	PriceQualifier formats.AlphaNumericString_Length1To3 `xml:"priceQualifier,omitempty"`
}

type TariffInformationDetailsTypeU_127523C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationDetailsTypeU_127523C"`

	// Prive value. The value conveyed equals 100 times the original value in order to avoid transporting decimal placement information.
	PriceAmount formats.NumericDecimal_Length1To18 `xml:"priceAmount,omitempty"`

	// Gives the type of amount.
	PriceQualifier formats.AlphaNumericString_Length1To3 `xml:"priceQualifier,omitempty"`
}

type TariffInformationDetailsTypeU_45479C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationDetailsTypeU_45479C"`

	// Prive value. The value conveyed equals 100 times the original value in order to avoid transporting decimal placement information.
	PriceAmount formats.NumericDecimal_Length1To18 `xml:"priceAmount,omitempty"`

	// Conveys the currency code.
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// This qualifier specifies how the price information is to be used. PPI = price per item. There is a price value. INC = Inclusive. There is no price value. This inclusive qualifier specifies that the price to which it applies is already accounted for in the price for another item.
	PriceQualifier formats.AlphaNumericString_Length1To3 `xml:"priceQualifier,omitempty"`
}

type TariffInformationDetailsTypeU_46314C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationDetailsTypeU_46314C"`

	// A unique rate product identifier:  - PRODUCT = tariff for a product  - TOUR    = price of the tour  - TAXFEE  = tax or fee  - REMAIN  = remaining amount
	RateIdentifier formats.AlphaNumericString_Length1To7 `xml:"rateIdentifier,omitempty"`

	// unitary amount for the tariff
	UnitaryAmount formats.NumericDecimal_Length1To30 `xml:"unitaryAmount,omitempty"`

	// currency code used for the tariff
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	// To qualify and get information on the tariff: cancellation charge, tax amount, total amount, no amount of insurance...
	TariffQualifier formats.AlphaNumericString_Length1To3 `xml:"tariffQualifier,omitempty"`

	// Total Amount for the tariff, set if quantity is present
	TotalAmount formats.NumericDecimal_Length1To20 `xml:"totalAmount,omitempty"`

	// quantity for the tariff, when tariff is detailed with a quantity x unitaryAmount and totalPrice
	Quantity formats.NumericInteger_Length1To3 `xml:"quantity,omitempty"`

	// value is codeset 65 if amount is negative. For Tour, the remaining amount to pay can be negative if the price of the tour has changed
	TariffStatus formats.AlphaNumericString_Length1To3 `xml:"tariffStatus,omitempty"`
}

type TariffInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationTypeI"`

	// This composite is used to convey the tariff information
	TariffInfo *TariffInformationDetailsTypeI `xml:"tariffInfo,omitempty"`

	// Additional rate type information
	RateInformation *RateInformationTypeI `xml:"rateInformation,omitempty"`

	// This composite is used to convey all the extra charge information.
	ChargeDetails *AssociatedChargesInformationTypeI `xml:"chargeDetails,omitempty"`
}

type TariffInformationTypeI_136706S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationTypeI_136706S"`

	// This composite is used to convey the tariff information
	TariffInfo *TariffInformationDetailsTypeI `xml:"tariffInfo,omitempty"`

	// Additional rate type information
	RateInformation *RateInformationTypeI_198204C `xml:"rateInformation,omitempty"`

	// This composite is used to convey all the extra charge information.
	ChargeDetails *AssociatedChargesInformationTypeI_198205C `xml:"chargeDetails,omitempty"`
}

type TariffInformationTypeI_136714S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationTypeI_136714S"`

	// tariff conversion indicator associated to the charge
	TariffInfo *TariffInformationDetailsTypeI_198216C `xml:"tariffInfo,omitempty"`

	// Tax, Surcharge, Coverage, Coupon details  Repetition are used to carry: - General Info (estimated + name...) - Tariff per day / Max - Tariff per weekend / Max - Tariff per week / Max - Tariff per month / Max - Tariff per rental / Max - Excess amount - Liability amount
	ChargeDetails *AssociatedChargesInformationTypeI_198218C `xml:"chargeDetails,omitempty"`
}

type TariffInformationTypeI_136719S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationTypeI_136719S"`

	// tariff conversion indicator associated to the charge
	TariffInfo *TariffInformationDetailsTypeI_198216C `xml:"tariffInfo,omitempty"`

	// Tax, Surcharge, Coverage, Coupon details
	ChargeDetails *AssociatedChargesInformationTypeI_198218C `xml:"chargeDetails,omitempty"`
}

type TariffInformationTypeI_22057S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationTypeI_22057S"`

	// total price and net premium
	TariffInfo *TariffInformationDetailsTypeI_39533C `xml:"tariffInfo,omitempty"`

	// to specify the taxes and their values and converted values into other currencies if specified.
	ChargeDetails *AssociatedChargesInformationTypeI_39535C `xml:"chargeDetails,omitempty"`
}

type TariffInformationTypeI_28460S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationTypeI_28460S"`

	// This composite is used to convey the tariff information
	TariffInfo *TariffInformationDetailsTypeI_50731C `xml:"tariffInfo,omitempty"`

	// Additional rate type information
	RateInformation *RateInformationTypeI_50732C `xml:"rateInformation,omitempty"`
}

type TariffInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationTypeU"`

	// This composite provides details for a price.
	PriceDetails *TariffInformationDetailsTypeU_45479C `xml:"priceDetails,omitempty"`
}

type TariffInformationTypeU_25419S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationTypeU_25419S"`

	// Information about the tariffs of a Tour: tour price, product price, additional price
	TariffInformation *TariffInformationDetailsTypeU_46314C `xml:"tariffInformation,omitempty"`

	// This composite is used to describe the commissions on the tariff
	AssociatedChargesInformation *AssociatedChargesInformationTypeU `xml:"associatedChargesInformation,omitempty"`
}

type TariffInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationType"`

	// This composite gives details about the monetary amounts and their usage.
	PriceDetails *TariffInformationDetailsTypeU `xml:"priceDetails,omitempty"`
}

type TariffInformationType_83558S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffInformationType_83558S"`

	// This composite gives details about the monetary amounts and their usage.
	PriceDetails *TariffInformationDetailsTypeU_127523C `xml:"priceDetails,omitempty"`
}

type TariffcodeType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TariffcodeType"`

	// To convey the Tariff code.
	TariffCode formats.AlphaNumericString_Length1To35 `xml:"tariffCode,omitempty"`

	// to convey a description of the type of tariff.
	TariffCodeType formats.AlphaNumericString_Length1To35 `xml:"tariffCodeType,omitempty"`
}

type TaxDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TaxDetailsTypeI"`

	// Tax Amount
	TaxRate formats.AlphaNumericString_Length1To17 `xml:"taxRate,omitempty"`

	// See ISO 4217 codes
	CurrCode formats.AlphaNumericString_Length1To3 `xml:"currCode,omitempty"`

	// Type of the tax
	TaxType formats.AlphaNumericString_Length1To3 `xml:"taxType,omitempty"`
}

type TaxDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TaxDetailsTypeU"`

	// Tax qualifier. For Ferry, only one tax qualifier applies: Port taxes.
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Tax amount.
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`
}

type TaxFieldsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TaxFieldsType"`

	// Tax indicator
	TaxIndicator formats.AlphaNumericString_Length1To1 `xml:"taxIndicator,omitempty"`

	// Tax currency
	TaxCurrency formats.AlphaNumericString_Length1To3 `xml:"taxCurrency,omitempty"`

	// Tax amount
	TaxAmount formats.AlphaNumericString_Length1To12 `xml:"taxAmount,omitempty"`

	// Tax country code
	TaxCountryCode formats.AlphaNumericString_Length1To3 `xml:"taxCountryCode,omitempty"`

	// Tax nature code
	TaxNatureCode formats.AlphaNumericString_Length1To2 `xml:"taxNatureCode,omitempty"`
}

type TaxTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TaxTypeI"`

	// specify the tax details
	TaxDetails *TaxDetailsTypeI `xml:"taxDetails,omitempty"`

	DummyNET *DummyNET `xml:"Dummy.NET,omitempty"`
}

type DummyNET struct{}

type TaxesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TaxesType"`

	// Tax details description.
	AdditionnalCharge *TaxDetailsTypeU `xml:"additionnalCharge,omitempty"`
}

type TerminalIdentificationDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TerminalIdentificationDescriptionType"`

	// Identification of the transaction initiator.
	TerminalID formats.AlphaNumericString_Length8To8 `xml:"terminalID,omitempty"`

	// The distribution channel.
	DistributionChannel *DistributionChannelType `xml:"distributionChannel,omitempty"`
}

type TerminalInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TerminalInformationTypeU"`

	// Arrival Terminal
	ArrivalTerminal formats.AlphaNumericString_Length1To2 `xml:"arrivalTerminal,omitempty"`
}

type TerminalTimeInformationTypeS struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TerminalTimeInformationTypeS"`

	// LOCATION IDENTIFICATION
	LocationDetails *LocationIdentificationTypeS `xml:"locationDetails,omitempty"`
}

type ThreeDomainSecureGroupType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ThreeDomainSecureGroupType"`

	// This segment conveys a set of data resulting from the 3DS authentication process
	AuthenticationData *CreditCardSecurityType `xml:"authenticationData,omitempty"`

	// Access Control Server's URL (up to 2048 characters).
	AcsURL *CommunicationContactType `xml:"acsURL,omitempty"`

	TdsBlobData *TdsBlobData `xml:"tdsBlobData,omitempty"`
}

type TdsBlobData struct {

	// will identify the content of the BLB that follows
	TdsBlbIdentifier *ReferenceInfoType_94524S `xml:"tdsBlbIdentifier,omitempty"`

	TdsBlbData *BinaryDataType `xml:"tdsBlbData,omitempty"`
}

type TicketElementType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TicketElementType"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// Ticket information
	Ticket *TicketInformationType `xml:"ticket,omitempty"`

	// Print options (//print options after double slash)
	PrintOptions formats.AlphaNumericString_Length1To127 `xml:"printOptions,omitempty"`
}

type TicketInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TicketInformationType"`

	// Ticketing type  TL, OK, DO, IN, MA, TR, AT, PT, XL, ST, SS
	Indicator formats.AlphaString_Length2To2 `xml:"indicator,omitempty"`

	// Ticketing date
	Date formats.NumericInteger_Length6To6 `xml:"date,omitempty"`

	// Ticketing time
	Time formats.Time24_HHMM `xml:"time,omitempty"`

	// Office Id
	OfficeId formats.AlphaNumericString_Length1To9 `xml:"officeId,omitempty"`

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To15 `xml:"freetext,omitempty"`

	// Air France flag (e.g. //TELEPAYE for MINITEL)
	TransactionFlag formats.AlphaNumericString_Length1To9 `xml:"transactionFlag,omitempty"`

	// Electronic ticketing flag + airline code (e.g. //ETLH)  ET for Electronic ticket candidate
	ElectronicTicketFlag formats.AlphaNumericString_Length1To2 `xml:"electronicTicketFlag,omitempty"`

	// Airline code
	AirlineCode formats.AlphaNumericString_Length1To3 `xml:"airlineCode,omitempty"`

	// Queue number
	QueueNumber formats.AlphaNumericString_Length1To3 `xml:"queueNumber,omitempty"`

	// Category number
	QueueCategory formats.AlphaNumericString_Length1To3 `xml:"queueCategory,omitempty"`

	// SITA addresses
	SitaAddress formats.AlphaNumericString_Length7To7 `xml:"sitaAddress,omitempty"`
}

type TicketInformationType_5120C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TicketInformationType_5120C"`

	// 1. PNR Header: Amadeus Queuing Office Id
	QueueingOfficeId formats.AlphaNumericString_Length1To24 `xml:"queueingOfficeId,omitempty"`

	// 1. PNR Header: OA city code
	Location formats.AlphaString_Length3To3 `xml:"location,omitempty"`
}

type TicketNumberDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TicketNumberDetailsTypeI"`

	// eVoucher number
	Number formats.AlphaNumericString_Length1To35 `xml:"number,omitempty"`
}

type TicketNumberTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TicketNumberTypeI"`

	// documentDetails
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails,omitempty"`
}

type TicketingFormOfPaymentType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TicketingFormOfPaymentType"`

	// Form of payment details
	FopDetails *FormOfPaymentInformationType `xml:"fopDetails,omitempty"`
}

type TimeZoneIinformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TimeZoneIinformationType"`

	// ISO country Code See SSIM appendix F
	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"`

	// Time zone code. See SSIM appendix F.
	Code formats.NumericInteger_Length1To1 `xml:"code,omitempty"`

	// Time zone suffix to complete the time zone code when necessary. See SSIM appendix F.
	Suffix formats.AlphaString_Length1To1 `xml:"suffix,omitempty"`
}

type TotalPriceType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TotalPriceType"`

	// The provider code.
	ProviderCode *CompanyInformationType_83550S `xml:"providerCode,omitempty"`

	// External Reference of the pricing
	ExternalRef *ReferenceInformationTypeI_83551S `xml:"externalRef,omitempty"`

	MethodOfDelivery *MethodOfDelivery `xml:"methodOfDelivery,omitempty"`

	// This segment is used to convey the main price information (e.g. the net total price for non-cancelled bookings, the cancellation fee for cancelled bookings).  The currency code stands not only for this segment, but for the whole group: all prices have the same currency.
	MainPrice *TariffInformationType `xml:"mainPrice,omitempty"`

	// The remaining price items are described here. In ferry business, there may be a maximum of 12 prices (+ main price and taxes).  The currency code is not applicable because it is the same as in the mainPriceInformation segment.
	OtherPrices *TariffInformationType_83558S `xml:"otherPrices,omitempty"`

	ProductDescription *ProductDescription `xml:"productDescription,omitempty"`

	// This segment conveys the tax amount information. The repetition factor equals the number of codesets for the qualifier, because each type of tax may occur once. The currency code is the same as in the mainPriceInformation segment.
	AdditionnalChargeInformation *TaxesType `xml:"additionnalChargeInformation,omitempty"`

	// This segment is used to convey the booking fare information.
	RateCodeInformation *RateTypesTypeU `xml:"rateCodeInformation,omitempty"`

	// This segment will transport the optional booking confirmation dead-line information. Note: if this segment transports a valid confirmation dead-line, then the booking is considered as optional.
	OptionalBooking *StructuredDateTimeInformationType_83553S `xml:"optionalBooking,omitempty"`
}

type ProductDescription struct {

	// product associated to the price item
	Product *ProductIdentificationTypeU `xml:"product,omitempty"`

	// product restrictions and attributes: route code and description, crossLondon and advanced purchase.
	ProductRestriction *TrafficRestrictionDetailsType `xml:"productRestriction,omitempty"`
}

type MethodOfDelivery struct {

	// Identification and semantic attached to the reference description (E.g: a customer can have multiple roles: payer, traveller, insured...)
	ElementManagement *ElementManagementSegmentType_83559S `xml:"elementManagement,omitempty"`

	// Describes the details around this mode of delivery
	DeliveryDetails *PackageDescriptionType `xml:"deliveryDetails,omitempty"`
}

type TourAccountDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TourAccountDetailsType"`

	// Total price of the Tours. The segment can be repeated in case the total price is written in more than one currency. This trigger is M20 and not M1, but there is no grammar problem with that. There would be a problem if the group TURP was repeated, but this is not and shall never be the case.
	TourTotalPrices *TariffInformationTypeU_25419S `xml:"tourTotalPrices,omitempty"`

	RemainingAmountsDetails *RemainingAmountsDetails `xml:"remainingAmountsDetails,omitempty"`

	TourDetailedPriceInfo *TourDetailedPriceInfo `xml:"tourDetailedPriceInfo,omitempty"`

	PaymentInformation *PaymentInformation `xml:"paymentInformation,omitempty"`
}

type PaymentInformation struct {

	// Deposit details: amount, currency, date, purpose and payment method information
	Payment *PaymentInformationTypeU `xml:"payment,omitempty"`

	// The Tour Operator code
	OperatorCode *CompanyInformationType_25420S `xml:"operatorCode,omitempty"`
}

type TourDetailedPriceInfo struct {

	// dummy segment
	MarkerSpecificRead *DummySegmentTypeI `xml:"markerSpecificRead,omitempty"`

	// Identifier of the product
	ProductId *ReferenceInfoType_25422S `xml:"productId,omitempty"`

	// Price of a product or tax or fee to pay for a product
	ProductPrice *TariffInformationTypeU_25419S `xml:"productPrice,omitempty"`
}

type RemainingAmountsDetails struct {

	// The Tour Operator code
	ProviderCode *CompanyInformationType_25420S `xml:"providerCode,omitempty"`

	// The remaining amount to pay
	RemainingAmount *TariffInformationTypeU_25419S `xml:"remainingAmount,omitempty"`
}

type TourDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TourDetailsTypeI"`

	// Inclusive tour number
	TourCode formats.AlphaNumericString_Length1To20 `xml:"tourCode,omitempty"`
}

type TourInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TourInformationTypeI"`

	// Tour code
	TourInformationDetails *TourDetailsTypeI `xml:"tourInformationDetails,omitempty"`
}

type TourInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TourInformationType"`

	// Conveys summary information of the booking (such as departure/arrival location or date of the booking, the booking type ...). The providerName data of the composite E988 DOES NOT conveys the Tour Operator's name. Providers' name and code are stored in CPYs.
	BookingSummaryInfo *TravelProductInformationTypeU_25428S `xml:"bookingSummaryInfo,omitempty"`

	// Conveys information about the booking duration
	BookingDurationInfo *QuantityType `xml:"bookingDurationInfo,omitempty"`

	// Conveys staying location information
	StayingInfo *PlaceLocationIdentificationTypeU_25436S `xml:"stayingInfo,omitempty"`

	// Conveys the tour description (name and description)
	TourDescriptionInfo *AdditionalProductDetailsTypeU `xml:"tourDescriptionInfo,omitempty"`

	// Conveys booking reference and unique key in the provider system
	BookingReferenceInfo *ReservationControlInformationTypeU `xml:"bookingReferenceInfo,omitempty"`

	// Conveys the status of the booking or of the ticket and the number in party. The composite E958 is M2:  -One instance can be the booking's status  -The other can be the TKOK status
	StatusInfo *QuantityAndActionTypeU `xml:"statusInfo,omitempty"`

	// Indicates whether an insurance is included in the tour booking.
	InsuranceIndication *InsuranceCoverageType_25483S `xml:"insuranceIndication,omitempty"`

	// Conveys passenger information when there is a desynchronization between the PNR passengers and tour passengers (Tour Server). A Tour Server booking can contain its own passenger names. Tour Server's specifications specify a maximum of 99 passengers. If a Tour booking conveys its passengers in TIF, then there is no Pax assoc (no REF used for Pax assoc). BUT, when there is no TIF (for example in Tour Distribution), there may be a Pax assoc between product and PNR pax, and/or booking and PNR's pax. Then REFs are used
	PassengerInfo *TravellerInformationType_25441S `xml:"passengerInfo,omitempty"`

	// Conveys the booking expiration information (only the expiration date is needed).
	ExpireInfo *StructuredDateTimeInformationType_25444S `xml:"expireInfo,omitempty"`

	// Conveys description information (Remark booking description). For the remarks we can have 2 lines of text that is why the 101C composite is repeted 2 times.
	BookingDescriptionInfo *FreeTextInformationType_25445S `xml:"bookingDescriptionInfo,omitempty"`

	// Conveys information about the targeted system provider (TS for Tour source, TG for Royal Orchid Holiday ...). The only information conveyd is the code of the provider whose system has the master Tour booking. Example: the Tour Operator ROH provides its products throw the TG (Thai Airways) system: TRA conveys TG, one CPY conveys ROH.
	SystemProviderInfo *TransportIdentifierType `xml:"systemProviderInfo,omitempty"`

	// Conveys information about the tour operator (name, code ...)
	TourOperatorInfo *CompanyInformationType_25420S `xml:"tourOperatorInfo,omitempty"`

	// Bokking source (/BS)
	BookingSource *UserIdentificationType_25447S `xml:"bookingSource,omitempty"`

	// Conveys the passenger association for the booking
	PassengerAssocation *ReferenceInfoType_25422S `xml:"passengerAssocation,omitempty"`

	// Conveys tour payment information such as the detailed price of the booking, the commisssion, the deposit information ...
	TourAccountDetails *TourAccountDetailsType `xml:"tourAccountDetails,omitempty"`

	// Conveys information about the booked products (arrival/departure information, product identification, meal plan information, occupation ...)
	TourProductDetails *TourServiceDetailsType `xml:"tourProductDetails,omitempty"`
}

type TourServiceDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TourServiceDetailsType"`

	// Conveys the product sequence number which is the product place in the booking. This information locally identifies the product in the Tour booking.
	SequenceNumberInfo *ItemNumberTypeU `xml:"sequenceNumberInfo,omitempty"`

	// Conveys information about the product status and the product quantity (number in party or number of service)
	StatusQuantityInfo *QuantityAndActionTypeU `xml:"statusQuantityInfo,omitempty"`

	// Conveys general Tour product information.
	ProductInfo *AdditionalProductDetailsTypeU `xml:"productInfo,omitempty"`

	// Conveys product confirmation number.
	ConfirmationInfo *ReservationControlInformationTypeU `xml:"confirmationInfo,omitempty"`

	// Passenger association at product (package / standalone) level.
	PassengerAssociation *ReferenceInfoType_25422S `xml:"passengerAssociation,omitempty"`

	ServiceDetails *ServiceDetails `xml:"serviceDetails,omitempty"`
}

type ServiceDetails struct {

	// Conveys general service information such as departure/arrival information, service code, service description, service type ...
	ServiceInfo *TravelProductInformationTypeU_25428S `xml:"serviceInfo,omitempty"`

	// Conveys duration information (number of day, night ...)
	ServiceDurationInfo *QuantityType `xml:"serviceDurationInfo,omitempty"`

	AccomodationDetails *AccomodationDetails `xml:"accomodationDetails,omitempty"`

	VehiculeDetails *VehiculeDetails `xml:"vehiculeDetails,omitempty"`

	TransportationDetails *TransportationDetails `xml:"transportationDetails,omitempty"`

	ProductBCSDetails *ProductBCSDetails `xml:"productBCSDetails,omitempty"`
}

type ProductBCSDetails struct {

	// BCS Agent sign, office and target office ids
	AgentIdentification *UserIdentificationType_25447S `xml:"agentIdentification,omitempty"`

	// BCS distribution channel
	DistributionChannelData *SystemDetailsInfoType `xml:"distributionChannelData,omitempty"`
}

type TransportationDetails struct {

	// Departure location information. The composite C517 conveys city information and the C519 the country information.
	DepartureInfo *PlaceLocationIdentificationTypeU_25436S `xml:"departureInfo,omitempty"`

	// Arrival location information. The composite C517 conveys city information and the C519 the country information.
	ArrivalInfo *PlaceLocationIdentificationTypeU_25436S `xml:"arrivalInfo,omitempty"`

	// Conveys transportation information
	TransportationInfo *TravelProductInformationTypeI_25434S `xml:"transportationInfo,omitempty"`

	// Conveys duration information (number of day, night ...)
	TransportationDuration *QuantityType `xml:"transportationDuration,omitempty"`

	// Conveys transportation equipment information
	EquipmentInfo *EquipmentDetailsTypeU `xml:"equipmentInfo,omitempty"`

	// Conveys transportation meal plan information
	TransportationMealPlanInfo *DiningInformationType `xml:"transportationMealPlanInfo,omitempty"`
}

type VehiculeDetails struct {

	// Conveys vehicule information (such as the vehicule occupancy)
	VehiculeInfo *VehicleTypeU_25502S `xml:"vehiculeInfo,omitempty"`
}

type AccomodationDetails struct {

	// Conveys room information
	RoomInfo *HotelRoomType_25429S `xml:"roomInfo,omitempty"`

	// Passenger association at accomodation room level
	PassengerAssociation *ReferenceInfoType_25422S `xml:"passengerAssociation,omitempty"`

	// Conveys room meal plan information
	RoomMealPlanInfo *DiningInformationType `xml:"roomMealPlanInfo,omitempty"`

	// Conveys room occupancy information (room min or max occupancy)
	OccupancynInfo *RangeDetailsTypeU `xml:"occupancynInfo,omitempty"`
}

type TrafficRestrictionDetailsTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TrafficRestrictionDetailsTypeU"`

	// restriction code.
	Code formats.AlphaNumericString_Length1To5 `xml:"code,omitempty"`

	// restriction type
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// traffic restriction description
	Description formats.AlphaNumericString_Length1To16 `xml:"description,omitempty"`
}

type TrafficRestrictionDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TrafficRestrictionDetailsType"`

	// restriction details
	RestrictionDetails *TrafficRestrictionDetailsTypeU `xml:"restrictionDetails,omitempty"`
}

type TrainDataType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TrainDataType"`

	// Information pertaining to the train product
	TrainProductInfo *TrainProductInformationType `xml:"trainProductInfo,omitempty"`

	// Trip dates and times
	TripDateTime *StructuredDateTimeInformationType_32362S `xml:"tripDateTime,omitempty"`

	// Departure station location
	DepLocation *PlaceLocationIdentificationTypeU_32347S `xml:"depLocation,omitempty"`

	// Arrival station location
	ArrLocation *PlaceLocationIdentificationTypeU_32347S `xml:"arrLocation,omitempty"`

	// Rail leg (train number, train provider, departure/arrival locations and dates, reservable status)
	RailLeg *RailLegDataType `xml:"railLeg,omitempty"`
}

type TrainDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TrainDetailsType"`

	// Train company code
	Code formats.AlphaNumericString_Length1To2 `xml:"code,omitempty"`

	// Train number
	Number formats.AlphaNumericString_Length1To6 `xml:"number,omitempty"`
}

type TrainInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TrainInformationType"`

	// Information pertaining to the rail company
	CompanyInfo *CompanyInformationType_19450S `xml:"companyInfo,omitempty"`

	// Indicates whether or not the reservation can be modified directly in the Amadeus system
	UpdatePermission *StatusTypeI_20923S `xml:"updatePermission,omitempty"`

	// train number, equipment code, departure and arrival dates and times.
	TripDetails *TrainDataType `xml:"tripDetails,omitempty"`

	// indicate that the train segment is open.
	OpenSegment *StatusTypeI_33257S `xml:"openSegment,omitempty"`

	// Journey direction: outward, return, single
	JourneyDirection *TravelItineraryInformationTypeI_33275S `xml:"journeyDirection,omitempty"`

	// Rail provider segment tattoo reference
	ProviderTattoo *ItemReferencesAndVersionsType `xml:"providerTattoo,omitempty"`

	// SVC / Service information
	ServiceInfo *FreeTextInformationType_20551S `xml:"serviceInfo,omitempty"`

	// Information pertaining to the class of service including number of seats
	ClassInfo *ClassConfigurationDetailsType `xml:"classInfo,omitempty"`

	// Accommodation (room/compartment) details.
	AccommodationInfo *AccommodationAllocationInformationTypeU `xml:"accommodationInfo,omitempty"`

	// Coach information
	CoachInfo *CoachProductInformationType `xml:"coachInfo,omitempty"`

	// Reservation Mandatory, Advised, Possible, Not Possible
	ReservableStatus *QuantityAndActionTypeU_32609S `xml:"reservableStatus,omitempty"`
}

type TrainProductInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TrainProductInformationType"`

	// Train Details
	TrainDetails *TrainDetailsType `xml:"trainDetails,omitempty"`

	// Transportation mode (BUS, SHIP, TRAIN, TGV etc)
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type TrainProductInformationType_32331S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TrainProductInformationType_32331S"`

	// Rail Company
	RailCompany formats.AlphaNumericString_Length1To4 `xml:"railCompany,omitempty"`

	// Train Details
	TrainDetails *TrainDetailsType `xml:"trainDetails,omitempty"`

	// Train Equipment Type  (TGV,TGD,TGN...)
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type TransactionInformationForTicketingType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TransactionInformationForTicketingType"`

	// Authorisation transaction details
	TransactionDetails *TransactionInformationsType `xml:"transactionDetails,omitempty"`
}

type TransactionInformationsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TransactionInformationsType"`

	// Authorization message type  Eg 110: author according standard ISO8583 210: settlement according standard ISO858 ...
	Code formats.AlphaNumericString_Length1To4 `xml:"code,omitempty"`

	// Credit Card link used to perform authorization.
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// Process indicator (bulkIndicator): - bulk - superbulk - no bulk.
	IssueIndicator formats.AlphaNumericString_Length1To1 `xml:"issueIndicator,omitempty"`

	// This is a message number that uniquely identifies a cardholder transaction.  According to the link this info can have various names:  - STAN number(Systems Trace Audit Number) - ISO8583 (VISA,Nedbank, Credit Mutuel...)  - Message number - APACS70 (Barclays,Euroline...)  ...   Official definition: This is a number assigned by the message initiator that uniquely identifies a cardholder transaction and all the message types (also known as system transactions) that it comprises, according to individual program rules. The trace number remains unchanged for all messages throughout the life of the transaction. For example, the same trace number is used in an authorization request and response, and in a subsequent reversal request and response, and in any advices of authorization or reversal.
	TransmissionControlNumber formats.AlphaNumericString_Length1To25 `xml:"transmissionControlNumber,omitempty"`
}

type TransportIdentifierType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TransportIdentifierType"`

	// Targeted provider system information
	CompanyIdentification *CompanyIdentificationTypeI_46351C `xml:"companyIdentification,omitempty"`
}

type TravelItineraryInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelItineraryInformationTypeI"`

	// The sequence number indentifying the position of a leg in a booking
	ItemNumber formats.NumericInteger_Length1To1 `xml:"itemNumber,omitempty"`
}

type TravelItineraryInformationTypeI_33275S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelItineraryInformationTypeI_33275S"`

	// direction of travel indicator (outward, return, single)
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`
}

type TravelProductInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelProductInformationTypeI"`

	// To specify dates and times of the product
	Product *ProductDateTimeTypeI_171495C `xml:"product,omitempty"`

	// Boarding point detail
	BoardpointDetail *LocationTypeI_2784C `xml:"boardpointDetail,omitempty"`

	// Off Point details
	OffpointDetail *LocationTypeI_2784C `xml:"offpointDetail,omitempty"`

	// Company identification
	CompanyDetail *CompanyIdentificationTypeI_2785C `xml:"companyDetail,omitempty"`

	// Product identifications details
	ProductDetails *ProductIdentificationDetailsTypeI_2786C `xml:"productDetails,omitempty"`

	// Product Type details
	TypeDetail *ProductTypeDetailsTypeI `xml:"typeDetail,omitempty"`

	// AIR segment : to indicate an Informational Air segment :  N for No action required.
	ProcessingIndicator formats.AlphaNumericString_Length1To3 `xml:"processingIndicator,omitempty"`
}

type TravelProductInformationTypeI_127288S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelProductInformationTypeI_127288S"`
}

type TravelProductInformationTypeI_186189S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelProductInformationTypeI_186189S"`

	// To specify dates and times of the product
	Product *ProductDateTimeTypeI_260882C `xml:"product,omitempty"`

	// Boarding point detail
	BoardpointDetail *LocationTypeI_2784C `xml:"boardpointDetail,omitempty"`

	// Off Point details
	OffpointDetail *LocationTypeI_2784C `xml:"offpointDetail,omitempty"`

	// Company identification
	CompanyDetail *CompanyIdentificationTypeI_2785C `xml:"companyDetail,omitempty"`

	// Product identifications details
	ProductDetails *ProductIdentificationDetailsTypeI_2786C `xml:"productDetails,omitempty"`

	// Product Type details
	TypeDetail *ProductTypeDetailsTypeI `xml:"typeDetail,omitempty"`

	// AIR segment : to indicate an Informational Air segment :  N for No action required.
	ProcessingIndicator formats.AlphaNumericString_Length1To3 `xml:"processingIndicator,omitempty"`
}

type TravelProductInformationTypeI_25434S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelProductInformationTypeI_25434S"`

	// flight date information
	FlightDate *ProductDateTimeTypeI_46338C `xml:"flightDate,omitempty"`

	// carrier details
	CompanyDetails *CompanyIdentificationTypeI_46335C `xml:"companyDetails,omitempty"`

	// flight information
	FlightIdentification *ProductIdentificationDetailsTypeI_46336C `xml:"flightIdentification,omitempty"`

	// connection sequence information
	FlightTypeDetails *ProductTypeDetailsTypeI_46337C `xml:"flightTypeDetails,omitempty"`
}

type TravelProductInformationTypeI_99362S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelProductInformationTypeI_99362S"`

	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`

	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`
}

type TravelProductInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelProductInformationTypeU"`

	// Conveys the departure and arrival date time descriptions. If absent, then the leg status may be considered as open information.
	ItineraryDateTimeInfo *ProductDateAndTimeTypeU `xml:"itineraryDateTimeInfo,omitempty"`

	// Conveys and itinerary leg embarkation and the disembarkation ports descriptions
	BoardPortDetails *LocationTypeU `xml:"boardPortDetails,omitempty"`

	// Internal reference for the leg.
	LineNumber formats.AlphaNumericString_Length1To2 `xml:"lineNumber,omitempty"`
}

type TravelProductInformationTypeU_25428S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelProductInformationTypeU_25428S"`

	// Conveys information about the departure/ arrival date and time.
	DateTimeInformation *ProductDateAndTimeTypeU_46325C `xml:"dateTimeInformation,omitempty"`

	// Conveys the departure/arrival/staying location information
	LocationInformation *LocationTypeU_46324C `xml:"locationInformation,omitempty"`

	// Conveys information about the provider of the product
	CompanyInformation *CompanyIdentificationTypeU `xml:"companyInformation,omitempty"`

	// Conveys details about the product
	ProductDetails *ProductIdentificationDetailsTypeU_46327C `xml:"productDetails,omitempty"`
}

type TravelerPerpaxDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravelerPerpaxDetailsType"`

	PerpaxMask formats.AlphaNumericString_Length2To2 `xml:"perpaxMask,omitempty"`

	// perpax mask indicator (optional/mandatory)
	PerpaxMaskIndicator formats.AlphaNumericString_Length1To1 `xml:"perpaxMaskIndicator,omitempty"`
}

type TravellerDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerDetailsTypeI"`

	// Traveler First Name
	FirstName formats.AlphaNumericString_Length1To56 `xml:"firstName,omitempty"`

	// Traveler Type using  Amadeus codification.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// 1 code is used to mention that the traveler is accompanied by an infant with no seat.
	InfantIndicator formats.AlphaNumericString_Length1To1 `xml:"infantIndicator,omitempty"`

	// Identification code, 2 cases: ID<1 to 51 char free text) or CR<1 to 40 char free text)
	IdentificationCode formats.AlphaNumericString_Length1To70 `xml:"identificationCode,omitempty"`
}

type TravellerDetailsTypeI_16351C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerDetailsTypeI_16351C"`

	// Conveys passenger first name.
	GivenName formats.AlphaNumericString_Length1To30 `xml:"givenName,omitempty"`

	Title formats.AlphaNumericString_Length1To3 `xml:"title,omitempty"`
}

type TravellerDetailsTypeI_18004C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerDetailsTypeI_18004C"`

	// firstname
	GivenName formats.AlphaNumericString_Length1To30 `xml:"givenName,omitempty"`
}

type TravellerDetailsTypeI_27968C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerDetailsTypeI_27968C"`

	// firstname
	GivenName formats.AlphaNumericString_Length1To30 `xml:"givenName,omitempty"`

	// Title (Mr, Mrs)
	Title formats.AlphaNumericString_Length1To70 `xml:"title,omitempty"`
}

type TravellerDetailsTypeI_46354C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerDetailsTypeI_46354C"`

	// Passenger lastName
	GivenName formats.AlphaNumericString_Length1To40 `xml:"givenName,omitempty"`

	// passenger title
	Title formats.AlphaNumericString_Length1To3 `xml:"title,omitempty"`
}

type TravellerDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerDetailsType"`

	// passenger first name
	GivenName formats.AlphaNumericString_Length1To70 `xml:"givenName,omitempty"`
}

type TravellerDocumentInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerDocumentInformationTypeU"`

	DocumentInformation *DocumentInformationTypeU `xml:"documentInformation,omitempty"`

	DatesOfValidity *ValidityDatesTypeU `xml:"datesOfValidity,omitempty"`
}

type TravellerInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerInformationTypeI"`

	// to specify the last name and the type of person (if it's a nanny or a substitute)
	PaxDetails *TravellerSurnameInformationTypeI_18003C `xml:"paxDetails,omitempty"`

	// Only used to put the firstname
	OtherPaxDetails *TravellerDetailsTypeI_18004C `xml:"otherPaxDetails,omitempty"`
}

type TravellerInformationTypeI_15923S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerInformationTypeI_15923S"`

	// to specify the last name and the type of person (if it's a nanny or a substitute)
	PaxDetails *TravellerSurnameInformationTypeI_18003C `xml:"paxDetails,omitempty"`

	// Other name info
	OtherPaxDetails *TravellerDetailsTypeI_27968C `xml:"otherPaxDetails,omitempty"`
}

type TravellerInformationTypeI_6097S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerInformationTypeI_6097S"`

	// Traveller surname details
	Traveller *TravellerSurnameInformationTypeI `xml:"traveller,omitempty"`

	// Occurrence one relates to the traveler.  Occurrence 2 relates only to an infant accompanying the traveler for whom only the given name is present.
	Passenger *TravellerDetailsTypeI `xml:"passenger,omitempty"`
}

type TravellerInformationTypeI_8956S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerInformationTypeI_8956S"`

	// Passenger last name details.
	PaxDetails *TravellerSurnameInformationTypeI_16350C `xml:"paxDetails,omitempty"`

	// Passnger first name details.
	OtherPaxDetails *TravellerDetailsTypeI_16351C `xml:"otherPaxDetails,omitempty"`
}

type TravellerInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerInformationType"`

	// to specify the last name, age and gender.
	PaxDetails *TravellerSurnameInformationType_858C `xml:"paxDetails,omitempty"`

	// to provide the first name
	OtherPaxDetails *TravellerDetailsTypeI_18004C `xml:"otherPaxDetails,omitempty"`
}

type TravellerInformationType_25441S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerInformationType_25441S"`

	// passenger details
	PaxDetails *TravellerSurnameInformationType_46353C `xml:"paxDetails,omitempty"`

	// other passenger details
	OtherPaxDetails *TravellerDetailsTypeI_46354C `xml:"otherPaxDetails,omitempty"`
}

type TravellerInformationType_94570S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerInformationType_94570S"`

	// will convey the name of the credit card holder
	PaxDetails *TravellerSurnameInformationType `xml:"paxDetails,omitempty"`

	// will convey the CC holder first name
	OtherPaxDetails *TravellerDetailsType `xml:"otherPaxDetails,omitempty"`
}

type TravellerInsuranceInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerInsuranceInformationType"`

	// currency of manual premium
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// manual total premium for this traveller
	Amount formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	// supplementary info
	SupplementaryInformation formats.AlphaNumericString_Length1To70 `xml:"supplementaryInformation,omitempty"`

	// gender - male or female
	SexCode formats.AlphaNumericString_Length1To1 `xml:"sexCode,omitempty"`

	// Credit card details
	CreditCardDetails *CreditCardType `xml:"creditCardDetails,omitempty"`

	// currency of the total premium ,
	TotalPremiumCurrency formats.AlphaNumericString_Length1To3 `xml:"totalPremiumCurrency,omitempty"`

	// calculated total premium , all taxes included for this traveller
	TotalPremium formats.NumericDecimal_Length1To18 `xml:"totalPremium,omitempty"`

	// for future use
	FutureCurrency formats.AlphaNumericString_Length1To3 `xml:"futureCurrency,omitempty"`

	// for future use
	FutureAmount formats.NumericDecimal_Length1To18 `xml:"futureAmount,omitempty"`

	// Reduction Code
	FareType formats.AlphaNumericString_Length1To3 `xml:"fareType,omitempty"`

	// Beneficiary Name
	TravelerName formats.AlphaNumericString_Length1To70 `xml:"travelerName,omitempty"`
}

type TravellerSurnameInformationTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerSurnameInformationTypeI"`

	// Traveler Last Name  Group name
	Surname formats.AlphaNumericString_Length1To57 `xml:"surname,omitempty"`

	// G for group (The traveler type is in C324/6353)
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// 1 :one traveler with exceptions below.  2 :traveler accompanied by an infant for whom only the given name is present.  n : total number of passengers of the group (assigned + unassigned)
	Quantity formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`

	// Staff type
	StaffType formats.AlphaNumericString_Length1To3 `xml:"staffType,omitempty"`
}

type TravellerSurnameInformationTypeI_16350C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerSurnameInformationTypeI_16350C"`

	// Conveys passenger last name.
	Surname formats.AlphaNumericString_Length1To30 `xml:"surname,omitempty"`
}

type TravellerSurnameInformationTypeI_18003C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerSurnameInformationTypeI_18003C"`

	// last name
	Surname formats.AlphaNumericString_Length1To30 `xml:"surname,omitempty"`

	// to specify the type of person
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`
}

type TravellerSurnameInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerSurnameInformationType"`

	// CC holder name details
	Surname formats.AlphaNumericString_Length1To70 `xml:"surname,omitempty"`
}

type TravellerSurnameInformationType_46353C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerSurnameInformationType_46353C"`

	// Passenger name
	Surname formats.AlphaNumericString_Length1To40 `xml:"surname,omitempty"`
}

type TravellerSurnameInformationType_858C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerSurnameInformationType_858C"`

	// Passenger name
	Surname formats.AlphaNumericString_Length1To30 `xml:"surname,omitempty"`

	// Customer type: A=adult C=child IN = infant
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// to indicate if it's a man or a female.
	Gender formats.AlphaNumericString_Length1To3 `xml:"gender,omitempty"`
}

type TravellerTimeDetailsTypeI struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TravellerTimeDetailsTypeI"`

	CheckinTime formats.Time24_HHMM `xml:"checkinTime,omitempty"`
}

type TstGeneralInformationDetailsType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TstGeneralInformationDetailsType"`

	// TST reference number
	TstReferenceNumber formats.AlphaNumericString_Length1To3 `xml:"tstReferenceNumber,omitempty"`

	// TST creation date
	TstCreationDate formats.AlphaNumericString_Length6To6 `xml:"tstCreationDate,omitempty"`

	// Sales indicator
	SalesIndicator formats.AlphaNumericString_Length1To4 `xml:"salesIndicator,omitempty"`
}

type TstGeneralInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A TstGeneralInformationType"`

	// General information
	GeneralInformation *TstGeneralInformationDetailsType `xml:"generalInformation,omitempty"`
}

type UniqueIdDescriptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A UniqueIdDescriptionType"`

	// ID sequence number : envelope number
	IDSequenceNumber formats.NumericInteger_Length1To5 `xml:"iDSequenceNumber,omitempty"`

	// ID qualifier: must be 'PNV' as PNR Version Number
	IDQualifier formats.AlphaNumericString_Length1To3 `xml:"iDQualifier,omitempty"`
}

type UserIdentificationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A UserIdentificationType"`

	// Originator Identification Details
	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification,omitempty"`

	// Agent type (A, T, E)
	OriginatorTypeCode formats.AlphaNumericString_Length1To1 `xml:"originatorTypeCode,omitempty"`
}

type UserIdentificationType_127265S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A UserIdentificationType_127265S"`

	// 1 character code for airline agent (A), travel agent (T), etc...
	OriginatorTypeCode formats.AlphaNumericString_Length1To1 `xml:"originatorTypeCode,omitempty"`
}

type UserIdentificationType_21014S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A UserIdentificationType_21014S"`

	// Originator Identification Details
	OriginIdentification *OriginatorIdentificationDetailsTypeI_37406C `xml:"originIdentification,omitempty"`
}

type UserIdentificationType_25447S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A UserIdentificationType_25447S"`

	// Originator Identification Details
	OriginIdentification *OriginatorIdentificationDetailsTypeI_46358C `xml:"originIdentification,omitempty"`

	// Booking source or  [agent numeric sign] + [agent initial] + [duty code]
	Originator formats.AlphaNumericString_Length1To11 `xml:"originator,omitempty"`
}

type UserIdentificationType_9456S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A UserIdentificationType_9456S"`

	// contains the client reference/signature.
	Originator formats.AlphaNumericString_Length1To30 `xml:"originator,omitempty"`
}

type UserPreferencesType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A UserPreferencesType"`

	// This composite contains details on user preferences : _ Country code _ Language code _ Currency code
	UserPreferences *OriginatorDetailsTypeI `xml:"userPreferences,omitempty"`
}

type ValidityDatesTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ValidityDatesTypeU"`

	// Date the document was issued (YYYYMMDD)
	IssueDate formats.Date_YYYYMMDD `xml:"issueDate,omitempty"`

	// Date document expires (YYYYMMDD)
	ExpirationDate formats.Date_YYYYMMDD `xml:"expirationDate,omitempty"`
}

type ValueRangeTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A ValueRangeTypeU"`

	// Unit Qualifier for maximum range gives in previous RNG:  DAY: Duration in days G:   Kilometers M:   Mileage MTH: Durarion in Months WE:  Duration in weeks
	MeasureUnitQualifier formats.AlphaNumericString_Length1To3 `xml:"measureUnitQualifier,omitempty"`
}

type VehicleInformationTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A VehicleInformationTypeU"`

	// Vehicle make and model.
	MakeAndModel formats.AlphaNumericString_Length1To17 `xml:"makeAndModel,omitempty"`
}

type VehicleInformationTypeU_46439C struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A VehicleInformationTypeU_46439C"`

	// Conveys the occupancy of a vehicule
	Occupancy formats.NumericInteger_Length1To3 `xml:"occupancy,omitempty"`
}

type VehicleInformationType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A VehicleInformationType"`

	// This composite is used to convey the vehicle type
	VehicleCharacteristic *VehicleTypeOptionType `xml:"vehicleCharacteristic,omitempty"`

	// This data element is used to convey the equipment codes.
	VehSpecialEquipment formats.AlphaNumericString_Length1To3 `xml:"vehSpecialEquipment,omitempty"`

	// To indicate vehicle details: -Number of doors -Number of seats -Max Number of doors -Max Number of seats -Number of bags -Volume of the boots
	VehicleInfo *QuantityDetailsTypeI_198209C `xml:"vehicleInfo,omitempty"`

	// Free text type
	FreeTextDetails *FreeTextDetailsType_198207C `xml:"freeTextDetails,omitempty"`

	// Description or Example of the Car
	CarModel formats.AlphaNumericString_Length1To55 `xml:"carModel,omitempty"`
}

type VehicleTypeOptionType struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A VehicleTypeOptionType"`

	// This data element is used to convey the owner of the type code.
	VehicleTypeOwner formats.AlphaNumericString_Length1To3 `xml:"vehicleTypeOwner,omitempty"`

	// This data element is used to convey the SIPP code(s) selection criteria.
	VehicleRentalPrefType formats.AlphaNumericString_Length1To4 `xml:"vehicleRentalPrefType,omitempty"`
}

type VehicleTypeU struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A VehicleTypeU"`

	// Conveys the type of vehicle.
	Category formats.AlphaNumericString_Length1To3 `xml:"category,omitempty"`

	// Describes the vehicle.
	VehicleDetails *VehicleInformationTypeU `xml:"vehicleDetails,omitempty"`
}

type VehicleTypeU_25502S struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/PNRACC_11_3_1A VehicleTypeU_25502S"`

	// Describe the vehicule
	VehiculeDescription *VehicleInformationTypeU_46439C `xml:"vehiculeDescription,omitempty"`
}
