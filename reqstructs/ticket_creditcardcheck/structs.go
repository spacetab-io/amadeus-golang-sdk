package ticket_creditcardcheck

import "encoding/xml"

type TicketCreditCardCheck struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/CCVRQT_06_1_1A Ticket_CreditCardCheck"`

	CommonCcData *CommonCcData `xml:"commonCcData"`

	HiddenData *HiddenData `xml:"hiddenData,omitempty"`  // minOccurs="0"

	Country *Country `xml:"country,omitempty"`  // minOccurs="0"
}

type CommonCcData struct {

	CcInfo *CcInfo `xml:"ccInfo"`

	PriceInfo *PriceInfo `xml:"priceInfo,omitempty"`  // minOccurs="0"

	AvsInfo *AvsInfo `xml:"avsInfo,omitempty"`  // minOccurs="0"
}

type CcInfo struct {

	VendorCode string `xml:"vendorCode"`

	CardNumber string `xml:"cardNumber"`

	ExpiryDate string `xml:"expiryDate"`

	ValidatingCarrier string `xml:"validatingCarrier,omitempty"`  // minOccurs="0"
}

type PriceInfo struct {

	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	Amount float64 `xml:"amount"`
}

type AvsInfo struct {

	AddressCode string `xml:"addressCode"`

	ZipCode string `xml:"zipCode"`
}

type HiddenData struct {

	SecurityId string `xml:"securityId,omitempty"`  // minOccurs="0"

	AuthenticationData *AuthenticationData `xml:"authenticationData,omitempty"`  // minOccurs="0"
}

type AuthenticationData struct {

	TransactionIdentifier string `xml:"transactionIdentifier,omitempty"`  // minOccurs="0"

	Veres string `xml:"veres,omitempty"`  // minOccurs="0"

	Pares string `xml:"pares,omitempty"`  // minOccurs="0"

	CreditCardCompany string `xml:"creditCardCompany"`

	AuthenticationIndicator string `xml:"authenticationIndicator,omitempty"`  // minOccurs="0"

	AuthenticationCode string `xml:"authenticationCode,omitempty"`  // minOccurs="0"
}

type Country struct {

	CountryCode string `xml:"countryCode"`
}
