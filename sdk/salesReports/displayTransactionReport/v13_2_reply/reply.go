package SalesReports_DisplayTransactionReportReply_v13_2 // tsrtrr132

//import "encoding/xml"

type SalesReportsDisplayTransactionReportReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TSRTRR_13_2_1A SalesReports_DisplayTransactionReportReply"`

	// Conveys a potential error and its associated Free flow text.
	ErrorGroup *ErrorGroupType `xml:"errorGroup,omitempty"` // minOccurs="0"

	// Conveys all the details associated to a transaction.
	TransactionReportDataDetails *TransactionReportDataDetails `xml:"transactionReportDataDetails,omitempty"` // minOccurs="0"
}

type TransactionReportDataDetails struct {
	// Conveys agency details:  - IATA number - Office ID.
	AgencyDetails *AdditionalBusinessSourceInformationTypeI `xml:"agencyDetails"`

	// Conveys details related to the agent associated to the requested document.
	AgentDetails *TicketAgentInfoType `xml:"agentDetails"`

	// Conveys the issuance date.
	DateDetails *StructuredDateTimeInformationType `xml:"dateDetails"`

	// Logical document number with airline code and related details.
	DocumentNumber *TicketNumberTypeI_51826S `xml:"documentNumber"`

	// Transaction details:  - transaction code - transaction type
	Transaction *TransactionInformationForTicketingType `xml:"transaction"`

	// This conveys the sequence number of the document within the sales report.
	DocumentIdentification *ItemNumberTypeI `xml:"documentIdentification"`

	// Conveys various document amounts:  - fees amount - commission amount - total taxes amount
	MonetaryInfo *MonetaryInformationTypeI `xml:"monetaryInfo"`

	// Reservation information: RLOC of the PNR.
	ReservationInformation *ReservationControlInformationTypeI `xml:"reservationInformation,omitempty"` // minOccurs="0"

	// Conveys the passenger name.
	PassengerName *TravellerInformationTypeI `xml:"passengerName,omitempty"` // minOccurs="0"

	EnhancedPaxInfoCriteria *EnhancedTravellerInformationType `xml:"enhancedPaxInfoCriteria,omitempty"` // minOccurs="0"

	// Conveys the information about the tour code associated to the document.
	TourCodeDetails *TourInformationTypeI `xml:"tourCodeDetails,omitempty"` // minOccurs="0"

	// Conveys the invoice details.
	InvoiceDetails *DocumentInformationDetailsTypeI `xml:"invoiceDetails,omitempty"` // minOccurs="0"

	// This group conveys data related to the FOP.
	FopGroup []*FOPRepresentationType `xml:"fopGroup"` // maxOccurs="9"

	Tax []*TaxType `xml:"tax,omitempty"` // minOccurs="0" maxOccurs="3"

	// Conveys the fare calculation details in all cases.   In specific cases (market/airline specifics) it conveys also: - the revenue attribution number/pseudo IATA number (entered in FZRA) - the payment on demand number generated at issuance time
	FareCalcDetails *CodedAttributeType `xml:"fareCalcDetails,omitempty"` // minOccurs="0"

	// Conveys the pricing indicator.
	PricingIndicator *PricingTicketingDetailsTypeI `xml:"pricingIndicator,omitempty"` // minOccurs="0"

	// Conveys the Settlement and authorisation code if present.
	ReferenceData *ReferenceInfoType `xml:"referenceData,omitempty"` // minOccurs="0"

	// Conveys the fees data if present
	FeeGroup *FeeGroup `xml:"feeGroup,omitempty"` // minOccurs="0"

	// Details about exchanged tickets
	ExchangedGroup []*ExchangedGroup `xml:"exchangedGroup,omitempty"` // minOccurs="0" maxOccurs="99"
}

type FeeGroup struct {
	// Conveys the fee type.
	FeeType *SelectionDetailsTypeI `xml:"feeType"`

	// Conveys details on the fee:  - fee subcode - fee amount - tax on fee code - tax on fee amount
	FeeDetails []*FeeDetails `xml:"feeDetails"` // maxOccurs="99"
}

type FeeDetails struct {
	// Fee subtype
	FeeSubType *SpecificDataInformationTypeI `xml:"feeSubType"`

	// fee monetary Info
	FeeMonetaryInfo *MonetaryInformationTypeI_52231S `xml:"feeMonetaryInfo"`

	FeeTaxInfo []*TaxType `xml:"feeTaxInfo,omitempty"` // minOccurs="0" maxOccurs="99"
}

type ExchangedGroup struct {
	// exchanged Ticket Number
	TicketExchanged *TicketNumberTypeI `xml:"ticketExchanged"`

	// Coupon numbers and SAC number
	CouponAssociated *CouponInformationType `xml:"couponAssociated,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type AdditionalBusinessSourceInformationTypeI struct {
	// Source type
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	// details of the sales report originator.
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails"`
}

type ApplicationErrorDetailType struct {
	// Code identifying the data validation error condition.
	// xmlType: AlphaNumericString_Length1To5
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	// xmlType: AlphaNumericString_Length1To3
	ErrorCategory string `xml:"errorCategory,omitempty"` // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	// xmlType: AlphaNumericString_Length1To3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"` // minOccurs="0"
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type CodedAttributeInformationType struct {
	// provides the attribute Type
	// xmlType: AlphaNumericString_Length1To5
	AttributeType string `xml:"attributeType"`

	// provides a description for the attribute
	// xmlType: AlphaNumericString_Length1To256
	AttributeDescription string `xml:"attributeDescription,omitempty"` // minOccurs="0"
}

type CodedAttributeType struct {
	// provides details for the Attribute
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"` // maxOccurs="99"
}

type CouponInformationDetailsType struct {
	// SAC number
	// xmlType: AlphaNumericString_Length1To35
	SettlementAuthorization string `xml:"settlementAuthorization,omitempty"` // minOccurs="0"

	// coupon number
	// xmlType: AlphaNumericString_Length1To60
	CpnReferenceNumber string `xml:"cpnReferenceNumber,omitempty"` // minOccurs="0"
}

type CouponInformationType struct {
	// Coupon number and SAC number
	CouponDetails1 *CouponInformationDetailsType `xml:"couponDetails1,omitempty"` // minOccurs="0"

	// Coupon number and SAC number
	CouponDetails2 *CouponInformationDetailsType `xml:"couponDetails2,omitempty"` // minOccurs="0"

	// Coupon number and SAC number
	CouponDetails3 *CouponInformationDetailsType `xml:"couponDetails3,omitempty"` // minOccurs="0"

	// Coupon number and SAC number
	CouponDetails4 *CouponInformationDetailsType `xml:"couponDetails4,omitempty"` // minOccurs="0"
}

type DataTypeInformationTypeI struct {
	// Conveys the fee subcode.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type"`
}

type DocumentDetailsTypeI struct {
	// Conveys the number of the document
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"` // minOccurs="0"
}

type DocumentInformationDetailsTypeI struct {
	// Invoice details
	DocumentDetails *DocumentDetailsTypeI `xml:"documentDetails"`
}

type EnhancedTravellerInformationType struct {
	// Name attributes unique for one passenger.
	TravellerNameInfo *TravellerNameInfoType `xml:"travellerNameInfo,omitempty"` // minOccurs="0"

	// 5 possible types of names, for one passenger (adult OR infant).
	OtherPaxNamesDetails []*TravellerNameDetailsType `xml:"otherPaxNamesDetails"` // maxOccurs="5"
}

type ErrorGroupType struct {
	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType `xml:"errorOrWarningCodeDetails"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"` // minOccurs="0"
}

type FOPRepresentationType struct {
	// Form of payment details.
	FopDescription *FormOfPaymentTypeI `xml:"fopDescription"`

	// To convey monetary amounts, currency...
	MonetaryInfo *MonetaryInformationTypeI `xml:"monetaryInfo"`
}

type FormOfPaymentDetailsTypeI struct {
	// Form of payment code.
	// xmlType: AlphaNumericString_Length1To10
	Type string `xml:"type"`

	// Form of payment indicator
	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"` // minOccurs="0"

	// Amount paid by the form of payment.
	// xmlType: NumericDecimal_Length1To18
	Amount *float64 `xml:"amount,omitempty"` // minOccurs="0"

	// Credit card company identification (VI, CA, AX...)
	// xmlType: AlphaNumericString_Length1To35
	VendorCode string `xml:"vendorCode,omitempty"` // minOccurs="0"

	// Credit card number
	// xmlType: AlphaNumericString_Length1To35
	CreditCardNumber string `xml:"creditCardNumber,omitempty"` // minOccurs="0"

	// Expiry date of the credit card.
	// xmlType: AlphaNumericString_Length1To35
	ExpiryDate string `xml:"expiryDate,omitempty"` // minOccurs="0"

	// Approval code of the credit card.
	// xmlType: AlphaNumericString_Length1To17
	ApprovalCode string `xml:"approvalCode,omitempty"` // minOccurs="0"

	// Source of the approval code: manual or automatic.
	// xmlType: AlphaNumericString_Length1To3
	SourceOfApproval string `xml:"sourceOfApproval,omitempty"` // minOccurs="0"

	// Extended payment indicator in case of CC Form of payment. It is composed of the sales channel indicator and the instalement value.  Example:  FP CCVI4012122222222226/1012/N123456/E01  Extended payment = E01
	// xmlType: AlphaNumericString_Length1To3
	ExtendedPayment string `xml:"extendedPayment,omitempty"` // minOccurs="0"
}

type FormOfPaymentTypeI struct {
	// Form of payment details
	FormOfPayment *FormOfPaymentDetailsTypeI `xml:"formOfPayment"`
}

type FreeTextDetailsType struct {
	// Qualifier of the subject
	// xmlType: AlphaNumericString_Length1To3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Type of information
	// xmlType: AlphaNumericString_Length1To4
	InformationType string `xml:"informationType,omitempty"` // minOccurs="0"

	// Status
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"` // minOccurs="0"

	// Company ID
	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId,omitempty"` // minOccurs="0"

	// Language
	// xmlType: AlphaNumericString_Length1To3
	Language string `xml:"language,omitempty"` // minOccurs="0"

	// Source of the free text
	// xmlType: AlphaNumericString_Length1To3
	Source string `xml:"source"`

	// Encoding type
	// xmlType: AlphaNumericString_Length1To3
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {
	// Error free text details
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	// xmlType: AlphaNumericString_Length1To199
	FreeText []string `xml:"freeText"` // maxOccurs="99"
}

type InternalIDDetailsType struct {
	// Conveys the agent's sine and duty code.
	// xmlType: AlphaNumericString_Length1To9
	InhouseId string `xml:"inhouseId"`

	// type of the inHouse ID.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"
}

type ItemNumberIdentificationTypeI struct {
	// Conveys a document sequence number within the sales report.
	// xmlType: AlphaNumericString_Length1To6
	Number string `xml:"number,omitempty"` // minOccurs="0"

	// Type of the number.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"
}

type ItemNumberTypeI struct {
	// Conveys the sequence number of a transaction within a sales report display.
	ItemNumberDetails *ItemNumberIdentificationTypeI `xml:"itemNumberDetails"`
}

type MonetaryInformationDetailsTypeI struct {
	// Type of the amount
	// xmlType: AlphaNumericString_Length1To3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount.
	// xmlType: AlphaNumericString_Length1To35
	Amount string `xml:"amount,omitempty"` // minOccurs="0"

	// Currency of the amount.
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_85122C struct {
	// amount type
	// xmlType: AlphaNumericString_Length1To3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount value
	// xmlType: AlphaNumericString_Length1To35
	Amount string `xml:"amount"`

	// Currency
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"` // minOccurs="0"
}

type MonetaryInformationTypeI struct {
	// Details on the amounts
	MonetaryDetails *MonetaryInformationDetailsTypeI `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"otherMonetaryDetails,omitempty"` // minOccurs="0" maxOccurs="19"
}

type MonetaryInformationTypeI_52231S struct {
	// Amount details
	MonetaryDetails *MonetaryInformationDetailsTypeI_85122C `xml:"monetaryDetails"`

	// Other monetary details
	OtherMonetaryDetails []*MonetaryInformationDetailsTypeI_85122C `xml:"otherMonetaryDetails,omitempty"` // minOccurs="0" maxOccurs="19"
}

type OriginatorIdentificationDetailsTypeI struct {
	// IATA number of the agency.
	OriginatorId int32 `xml:"originatorId"`

	// Office ID of the agency.
	// xmlType: AlphaNumericString_Length1To9
	InHouseIdentification1 string `xml:"inHouseIdentification1,omitempty"` // minOccurs="0"
}

type PricingTicketingDetailsTypeI struct {
	// Conveys details on the pricing of the ticket.
	PriceTicketDetails *PricingTicketingInformationTypeI `xml:"priceTicketDetails,omitempty"` // minOccurs="0"
}

type PricingTicketingInformationTypeI struct {
	// Conveys the pricing indicator.
	// xmlType: AlphaNumericString_Length1To3
	Indicators []string `xml:"indicators,omitempty"` // minOccurs="0" maxOccurs="20"
}

type ReferenceInfoType struct {
	// REFERENCING DETAILS
	ReferenceDetails *ReferencingDetailsTypeI `xml:"referenceDetails,omitempty"` // minOccurs="0"
}

type ReferencingDetailsTypeI struct {
	// Conveys the type of the reference.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"

	// Value of the number
	// xmlType: AlphaNumericString_Length1To35
	Value string `xml:"value,omitempty"` // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI struct {
	// Reservation control number : RLoc of the PNR.
	// xmlType: AlphaNumericString_Length1To20
	ControlNumber string `xml:"controlNumber,omitempty"` // minOccurs="0"
}

type ReservationControlInformationTypeI struct {
	// Reservation details
	Reservation []*ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"` // minOccurs="0" maxOccurs="9"
}

type SelectionDetailsInformationTypeI struct {
	// To specify the type of carrier fee.
	// xmlType: AlphaNumericString_Length1To3
	Option string `xml:"option"`
}

type SelectionDetailsTypeI struct {
	// Selection details
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails"`
}

type SourceTypeDetailsTypeI struct {
	// Source qualifier: reporting office.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier1 string `xml:"sourceQualifier1"`

	// Indicates that data are associated to all agencies sharing the same IATA number.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier2 string `xml:"sourceQualifier2,omitempty"` // minOccurs="0"
}

type SpecificDataInformationTypeI struct {
	// Conveys fees subcode and information
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`
}

type StructuredDateTimeInformationType struct {
	// This data element can be used to provide the semantic of the information provided. Examples : - Current date - Requested specific date
	// xmlType: AlphaNumericString_Length1To3
	BusinessSemantic string `xml:"businessSemantic,omitempty"` // minOccurs="0"

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"` // minOccurs="0"
}

type StructuredDateTimeType struct {
	// Year number.
	//formats: Year_YYYY
	Year string `xml:"year,omitempty"` // minOccurs="0"

	// Month number in the year ( begins to 1 )
	//formats: Month_mM
	Month string `xml:"month,omitempty"` // minOccurs="0"

	// Day number in the month ( begins to 1 )
	//formats: Day_nN
	Day string `xml:"day,omitempty"` // minOccurs="0"
}

type TaxDetailsType struct {
	// Tax/fee/charge amount
	// xmlType: AlphaNumericString_Length1To17
	Rate string `xml:"rate,omitempty"` // minOccurs="0"

	// Country code
	// xmlType: AlphaNumericString_Length1To3
	CountryCode string `xml:"countryCode,omitempty"` // minOccurs="0"

	// ISO code identifying the currency
	// xmlType: AlphaNumericString_Length1To3
	CurrencyCode string `xml:"currencyCode,omitempty"` // minOccurs="0"

	// xmlType: AlphaNumericString_Length1To3
	Type []string `xml:"type,omitempty"` // minOccurs="0" maxOccurs="2"
}

type TaxType struct {
	// Tax category
	TaxCategory *int32 `xml:"taxCategory,omitempty"` // minOccurs="0"

	// Refund tax details
	TaxDetails *TaxDetailsType `xml:"taxDetails,omitempty"` // minOccurs="0"
}

type TicketAgentInfoType struct {
	// Company number
	// xmlType: AlphaNumericString_Length1To15
	CompanyIdNumber string `xml:"companyIdNumber,omitempty"` // minOccurs="0"

	// Conveys the agent's details.
	InternalIdDetails []*InternalIDDetailsType `xml:"internalIdDetails,omitempty"` // minOccurs="0" maxOccurs="5"
}

type TicketNumberDetailsTypeI struct {
	// document number
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number"`
}

type TicketNumberDetailsTypeI_84557C struct {
	// Document number.
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"` // minOccurs="0"

	// In case of conjunctiv, indicates the total number of tickets.
	NumberOfBooklets *int32 `xml:"numberOfBooklets,omitempty"` // minOccurs="0"
}

type TicketNumberTypeI struct {
	// document identifier
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`

	// status
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"` // minOccurs="0"
}

type TicketNumberTypeI_51826S struct {
	// Details on the document transaction.
	DocumentDetails *TicketNumberDetailsTypeI_84557C `xml:"documentDetails"`

	// Provide status on the document in the sales report: it can be confirmed or not.
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"` // minOccurs="0"
}

type TourDetailsTypeI struct {
	// Tour code.
	// xmlType: AlphaNumericString_Length1To35
	TourCode string `xml:"tourCode,omitempty"` // minOccurs="0"
}

type TourInformationTypeI struct {
	// Details on the tour.
	TourInformationDetails *TourDetailsTypeI `xml:"tourInformationDetails,omitempty"` // minOccurs="0"
}

type TransactionInformationForTicketingType struct {
	// Reporting transaction details
	TransactionDetails *TransactionInformationsType `xml:"transactionDetails"`
}

type TransactionInformationsType struct {
	// Transaction Code, coded : CANR  MCOA  MCOM  MDnn MPnn PTAM TKTA  TKTB  TKTM  TKTT  TORM  XSBA  XSBM  ACMR   RENM  RFND  ACMA  SSAC  TAAD  ADMA  RCSM  SSAD  BPAS  CANX CANN  PSCN  VSCN  RSCN RENA TASF
	// xmlType: AlphaNumericString_Length1To4
	Code string `xml:"code,omitempty"` // minOccurs="0"

	// Transaction Type, coded :     SALE     INVT     REFD      ADJA     ADJP     AUTS     CCAS     CCCS     MANS     VOID
	// xmlType: AlphaNumericString_Length1To4
	Type string `xml:"type,omitempty"` // minOccurs="0"

	// 'O' --) Original transaction code 'C' --) Current transaction code
	// xmlType: AlphaNumericString_Length1To1
	IssueIndicator string `xml:"issueIndicator,omitempty"` // minOccurs="0"
}

type TravellerInformationTypeI struct {
	// Conveys details on the passenger.
	PaxDetails *TravellerSurnameInformationTypeI `xml:"paxDetails"`
}

type TravellerNameDetailsType struct {
	// Name type. Allows to identify the type of name, Native name, Universal name...
	// xmlType: AlphaNumericString_Length1To5
	NameType string `xml:"nameType,omitempty"` // minOccurs="0"

	// Reference name.
	// xmlType: AlphaNumericString_Length1To1
	ReferenceName string `xml:"referenceName,omitempty"` // minOccurs="0"

	// Displayed name.
	// xmlType: AlphaNumericString_Length1To1
	DisplayedName string `xml:"displayedName,omitempty"` // minOccurs="0"

	// Romanization method.
	// xmlType: AlphaNumericString_Length1To4
	RomanizationMethod string `xml:"romanizationMethod,omitempty"` // minOccurs="0"

	// Passenger surname
	// xmlType: AlphaNumericString_Length1To70
	Surname string `xml:"surname"`

	// Passenger firstname
	// xmlType: AlphaNumericString_Length1To70
	GivenName string `xml:"givenName,omitempty"` // minOccurs="0"

	// Title, when separated from the firstname field.
	// xmlType: AlphaNumericString_Length1To70
	Title []string `xml:"title,omitempty"` // minOccurs="0" maxOccurs="2"
}

type TravellerNameInfoType struct {
	// PAX = PAX IN = Infant
	// xmlType: AlphaNumericString_Length1To3
	Qualifier string `xml:"qualifier,omitempty"` // minOccurs="0"

	// Quantity.
	Quantity *int32 `xml:"quantity,omitempty"` // minOccurs="0"

	// Passenger type (PTC).
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"

	// Passenger type (PTC).
	// xmlType: AlphaNumericString_Length1To3
	OtherType string `xml:"otherType,omitempty"` // minOccurs="0"

	// This field contains the tattoo number of the passenger that will be updated. The tattoo number is retrieved from a PNR structured image.
	UniqueCustomerIdentifier *int32 `xml:"uniqueCustomerIdentifier,omitempty"` // minOccurs="0"

	// 1. Infant (INF) No more info in Edifact. 2. Infant given name only (INF/BILL) Infant given name will be placed in another occurence of C325 of this (adult) passenger ETI. This other C324/6353 element will contain INF. 3. Infant given and last name (INFGATES/BILL) Infant is treated as a separate ETI following immediately this (adult) passenger ETI. This following ETI element will contain INF.
	// xmlType: AlphaNumericString_Length1To1
	InfantIndicator string `xml:"infantIndicator,omitempty"` // minOccurs="0"
}

type TravellerSurnameInformationTypeI struct {
	// Passenger surname followed by the name.
	// xmlType: AlphaNumericString_Length1To70
	Surname string `xml:"surname"`

	// Qualifier
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"` // minOccurs="0"
}
