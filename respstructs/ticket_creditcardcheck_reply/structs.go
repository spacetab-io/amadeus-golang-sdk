package ticket_creditcardcheck_reply

//import "encoding/xml"

type TicketCreditCardCheckReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/CCVRSP_06_1_1A Ticket_CreditCardCheckReply"`

	ResultInfo *ResultInfo `xml:"resultInfo,omitempty"`  // minOccurs="0"

	CreditCardData *CreditCardData `xml:"creditCardData"`
}

type ResultInfo struct {

	ErrorWarningInfo *ErrorWarningInfo `xml:"errorWarningInfo"`

	ErrorWarningText *ErrorWarningText `xml:"errorWarningText,omitempty"`  // minOccurs="0"
}

type ErrorWarningInfo struct {

	ApplicationDetails *ApplicationDetails `xml:"applicationDetails"`
}

type ApplicationDetails struct {

	// Format limitations: an..5
	Number string `xml:"number"`

	// Format limitations: an..3
	Qualifier string `xml:"qualifier"`
}

type ErrorWarningText struct {

	FreeTextDetails *FreeTextDetails `xml:"freeTextDetails,omitempty"`  // minOccurs="0"

	// Format limitations: an..70
	FreeText string `xml:"freeText"`
}

type FreeTextDetails struct {

	// Format limitations: an..3
	Qualifier string `xml:"qualifier"`
}

type CreditCardData struct {

	CommonCcData *CommonCcData `xml:"commonCcData"`

	VisaInfo *VisaInfo `xml:"visaInfo,omitempty"`  // minOccurs="0"

	CannedMessageReference *CannedMessageReference `xml:"cannedMessageReference,omitempty"`  // minOccurs="0"

	SecurityIdResponse *SecurityIdResponse `xml:"securityIdResponse,omitempty"`  // minOccurs="0"
}

type CommonCcData struct {

	CcInfo *CcInfo `xml:"ccInfo"`

	PriceInfo *PriceInfo `xml:"priceInfo,omitempty"`  // minOccurs="0"

	AvsInfo *AvsInfo `xml:"avsInfo,omitempty"`  // minOccurs="0"

	MerchantInfo *MerchantInfo `xml:"merchantInfo,omitempty"`  // minOccurs="0"

	// Format limitations: an..5
	RejectCode string `xml:"rejectCode,omitempty"`  // minOccurs="0"
}

type CcInfo struct {

	// Format limitations: an2
	VendorCode string `xml:"vendorCode"`

	// Format limitations: an..19
	CardNumber string `xml:"cardNumber"`

	// Format limitations: n4
	ExpiryDate string `xml:"expiryDate"`

	// Format limitations: an..8
	AuthorizationCode string `xml:"authorizationCode,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	ValidatingCarrier string `xml:"validatingCarrier,omitempty"`  // minOccurs="0"
}

type PriceInfo struct {

	// Format limitations: n3
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Format limitations: n..12
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"
}

type AvsInfo struct {

	// Format limitations: an..20
	AddressCode string `xml:"addressCode,omitempty"`  // minOccurs="0"

	// Format limitations: an..9
	ZipCode string `xml:"zipCode,omitempty"`  // minOccurs="0"

	// Format limitations: an1
	ResultCode string `xml:"resultCode,omitempty"`  // minOccurs="0"
}

type MerchantInfo struct {

	// Format limitations: an..4
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// Format limitations: an..15
	Id string `xml:"id,omitempty"`  // minOccurs="0"
}

type VisaInfo struct {

	MsgRef *MsgRef `xml:"msgRef,omitempty"`  // minOccurs="0"

	RespIdentification *RespIdentification `xml:"respIdentification,omitempty"`  // minOccurs="0"

	CardBillingInfo *CardBillingInfo `xml:"cardBillingInfo,omitempty"`  // minOccurs="0"

	ReversalInfo *ReversalInfo `xml:"reversalInfo,omitempty"`  // minOccurs="0"
}

type MsgRef struct {

	// Format limitations: an..12
	Number string `xml:"number,omitempty"`  // minOccurs="0"

	// Format limitations: an1
	AuthorizIndic string `xml:"authorizIndic,omitempty"`  // minOccurs="0"

	// Format limitations: an2
	PointOfService string `xml:"pointOfService,omitempty"`  // minOccurs="0"
}

type RespIdentification struct {

	// Format limitations: an..15
	TransacIdentifier string `xml:"transacIdentifier,omitempty"`  // minOccurs="0"

	// Format limitations: an1
	StipReason string `xml:"stipReason,omitempty"`  // minOccurs="0"

	// Format limitations: an..4
	ValidationCode string `xml:"validationCode,omitempty"`  // minOccurs="0"
}

type CardBillingInfo struct {

	// Format limitations: n3
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Format limitations: n..12
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"

	// Format limitations: n..8
	Rate *float64 `xml:"rate,omitempty"`  // minOccurs="0"
}

type ReversalInfo struct {

	// Format limitations: an..42
	OriginalData string `xml:"originalData,omitempty"`  // minOccurs="0"

	// Format limitations: an..12
	ReplaceAmount string `xml:"replaceAmount,omitempty"`  // minOccurs="0"
}

type CannedMessageReference struct {

	ReferenceDetails *ReferenceDetails `xml:"referenceDetails"`
}

type ReferenceDetails struct {

	// Format limitations: an..3
	Type string `xml:"type"`

	// Format limitations: an..5
	Value string `xml:"value"`
}

type SecurityIdResponse struct {

	FreeTextDetails *FreeTextDetails1 `xml:"freeTextDetails"`

	// Format limitations: an..199
	FreeText string `xml:"freeText"`
}

type FreeTextDetails1 struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..3
	Source string `xml:"source"`

	// Format limitations: an..3
	Encoding string `xml:"encoding"`
}
