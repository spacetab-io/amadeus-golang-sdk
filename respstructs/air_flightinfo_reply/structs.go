package air_flightinfo_reply

//import "encoding/xml"

type AirFlightInfoReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FLIRES_05_1_1A Air_FlightInfoReply"`

	MessageActionDetails *MessageActionDetails `xml:"messageActionDetails"`

	InteractiveFreeText *InteractiveFreeText `xml:"interactiveFreeText,omitempty"`  // minOccurs="0"

	ResponseError *ResponseError `xml:"responseError,omitempty"`  // minOccurs="0"

	FlightScheduleDetails *FlightScheduleDetails `xml:"flightScheduleDetails,omitempty"`  // minOccurs="0"
}

type MessageActionDetails struct {

	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails"`
}

type MessageFunctionDetails struct {

	// Format limitations: an..3
	BusinessFunction string `xml:"businessFunction,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	MessageFunction string `xml:"messageFunction"`

	// Format limitations: an..3
	ResponsibleAgency string `xml:"responsibleAgency"`

	// Format limitations: an..3
	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type InteractiveFreeText struct {

	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText string `xml:"freeText,omitempty"`  // minOccurs="0"
}

type FreeTextQualification struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..4
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Language string `xml:"language,omitempty"`  // minOccurs="0"
}

type ResponseError struct {

	ErrorInfo *ErrorInfo `xml:"errorInfo"`

	InteractiveFreeText *InteractiveFreeText `xml:"interactiveFreeText,omitempty"`  // minOccurs="0"
}

type ErrorInfo struct {

	ErrorDetails *ErrorDetails `xml:"errorDetails"`
}

type ErrorDetails struct {

	// Format limitations: an..3
	ErrorCode string `xml:"errorCode"`

	// Format limitations: an..3
	ErrorCategory string `xml:"errorCategory,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`  // minOccurs="0"
}

type FlightScheduleDetails struct {

	DummySegment *DummySegment `xml:"dummySegment"`

	GeneralFlightInfo *GeneralFlightInfo `xml:"generalFlightInfo"`

	AdditionalProductDetails *AdditionalProductDetails `xml:"additionalProductDetails,omitempty"`  // minOccurs="0"

	ProductInfo *ProductInfo `xml:"productInfo,omitempty"`  // minOccurs="0"

	InteractiveFreeText []*InteractiveFreeText `xml:"interactiveFreeText,omitempty"`  // minOccurs="0" maxOccurs="99"

	ErrorResponseOrWarningData *ErrorResponseOrWarningData `xml:"errorResponseOrWarningData,omitempty"`  // minOccurs="0"

	BoardPointAndOffPointDetails []*BoardPointAndOffPointDetails `xml:"boardPointAndOffPointDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type DummySegment struct {
}

type GeneralFlightInfo struct {

	FlightDate *FlightDate `xml:"flightDate"`

	BoardPointDetails *BoardPointDetails `xml:"boardPointDetails,omitempty"`  // minOccurs="0"

	OffPointDetails *OffPointDetails `xml:"offPointDetails,omitempty"`  // minOccurs="0"

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`  // minOccurs="0"

	ProductIdDetails *ProductIdDetails `xml:"productIdDetails,omitempty"`  // minOccurs="0"
}

type FlightDate struct {

	// Format limitations: n6
	DepartureDate string `xml:"departureDate,omitempty"`  // minOccurs="0"

	// Format limitations: n4
	DepartureTime string `xml:"departureTime,omitempty"`  // minOccurs="0"

	// Format limitations: n6
	ArrivalDate string `xml:"arrivalDate,omitempty"`  // minOccurs="0"

	// Format limitations: n4
	ArrivalTime string `xml:"arrivalTime,omitempty"`  // minOccurs="0"
}

type BoardPointDetails struct {

	// Format limitations: a3
	TrueLocationId string `xml:"trueLocationId"`
}

type OffPointDetails struct {

	// Format limitations: a3
	TrueLocationId string `xml:"trueLocationId"`
}

type CompanyDetails struct {

	// Format limitations: an..3
	MarketingCompany string `xml:"marketingCompany"`

	// Format limitations: an..3
	OperatingCompany string `xml:"operatingCompany,omitempty"`  // minOccurs="0"
}

type ProductIdDetails struct {

	// Format limitations: an..4
	FlightNumber string `xml:"flightNumber"`

	// Format limitations: an1
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`  // minOccurs="0"

	// Format limitations: an..7
	Modifier []string `xml:"modifier,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type AdditionalProductDetails struct {

	LegDetails *LegDetails `xml:"legDetails,omitempty"`  // minOccurs="0"

	DepartureStationInfo *DepartureStationInfo `xml:"departureStationInfo,omitempty"`  // minOccurs="0"

	ArrivalStationInfo *ArrivalStationInfo `xml:"arrivalStationInfo,omitempty"`  // minOccurs="0"

	FlightLegMileag *FlightLegMileag `xml:"flightLegMileag,omitempty"`  // minOccurs="0"

	TravellerTimeDetails *TravellerTimeDetails `xml:"travellerTimeDetails,omitempty"`  // minOccurs="0"

	FacilitiesInformation []*FacilitiesInformation `xml:"facilitiesInformation,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type LegDetails struct {

	// Format limitations: an..3
	Equipment string `xml:"equipment,omitempty"`  // minOccurs="0"

	// Format limitations: n..2
	NumberOfStops *float64 `xml:"numberOfStops,omitempty"`  // minOccurs="0"

	// Format limitations: n4
	Duration string `xml:"duration,omitempty"`  // minOccurs="0"

	// Format limitations: n..7
	DaysOfOperation *float64 `xml:"daysOfOperation,omitempty"`  // minOccurs="0"

	// Format limitations: an1
	ComplexingFlightIndicator string `xml:"complexingFlightIndicator,omitempty"`  // minOccurs="0"

	// Format limitations: a3
	Location1 string `xml:"location1,omitempty"`  // minOccurs="0"

	// Format limitations: a3
	Location2 string `xml:"location2,omitempty"`  // minOccurs="0"

	// Format limitations: a3
	Location3 string `xml:"location3,omitempty"`  // minOccurs="0"
}

type DepartureStationInfo struct {

	// Format limitations: an..6
	GateDescription string `xml:"gateDescription,omitempty"`  // minOccurs="0"

	// Format limitations: an..2
	Terminal string `xml:"terminal,omitempty"`  // minOccurs="0"

	// Format limitations: an..2
	Concourse string `xml:"concourse,omitempty"`  // minOccurs="0"
}

type ArrivalStationInfo struct {

	// Format limitations: an..6
	GateDescription string `xml:"gateDescription,omitempty"`  // minOccurs="0"

	// Format limitations: an..2
	Terminal string `xml:"terminal,omitempty"`  // minOccurs="0"

	// Format limitations: an..2
	Concourse string `xml:"concourse,omitempty"`  // minOccurs="0"
}

type FlightLegMileag struct {

	// Format limitations: n..8
	FlightLegMileage *float64 `xml:"flightLegMileage,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	UnitQualifier string `xml:"unitQualifier,omitempty"`  // minOccurs="0"
}

type TravellerTimeDetails struct {

	// Format limitations: n4
	DepartureTime string `xml:"departureTime,omitempty"`  // minOccurs="0"

	// Format limitations: n4
	ArrivalTime string `xml:"arrivalTime,omitempty"`  // minOccurs="0"

	// Format limitations: n10
	CheckInDetails string `xml:"checkInDetails,omitempty"`  // minOccurs="0"
}

type FacilitiesInformation struct {

	// Format limitations: n4
	Description string `xml:"description,omitempty"`  // minOccurs="0"
}

type ProductInfo struct {

	BookingClassDetails []*BookingClassDetails `xml:"bookingClassDetails"`  // maxOccurs="26"
}

type BookingClassDetails struct {

	// Format limitations: an1
	Designator string `xml:"designator"`

	// Format limitations: an..3
	AvailabilityStatus string `xml:"availabilityStatus,omitempty"`  // minOccurs="0"

	// Format limitations: an..2
	SpecialService string `xml:"specialService,omitempty"`  // minOccurs="0"

	// Format limitations: an1
	Option []string `xml:"option,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type ErrorResponseOrWarningData struct {

	ErrorInfo *ErrorInfo `xml:"errorInfo"`

	InteractiveFreeText *InteractiveFreeText `xml:"interactiveFreeText,omitempty"`  // minOccurs="0"
}

type BoardPointAndOffPointDetails struct {

	GeneralFlightInfo *GeneralFlightInfo `xml:"generalFlightInfo"`

	AdditionalProductDetails *AdditionalProductDetails `xml:"additionalProductDetails,omitempty"`  // minOccurs="0"

	ProductInfo *ProductInfo `xml:"productInfo,omitempty"`  // minOccurs="0"

	EquipmentInfo *EquipmentInfo `xml:"equipmentInfo,omitempty"`  // minOccurs="0"

	InteractiveFreeText []*InteractiveFreeText `xml:"interactiveFreeText,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type EquipmentInfo struct {

	CabinClassDetails []*CabinClassDetails `xml:"cabinClassDetails,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Format limitations: an..3
	IataAircraftTypeCode string `xml:"iataAircraftTypeCode,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	ConfigVersionDescription string `xml:"configVersionDescription,omitempty"`  // minOccurs="0"
}

type CabinClassDetails struct {

	// Format limitations: an1
	ClassDesignator string `xml:"classDesignator"`

	// Format limitations: n..3
	NumberOfSeats float64 `xml:"numberOfSeats"`
}
