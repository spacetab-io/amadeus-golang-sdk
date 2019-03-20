package Security_SignOut_v04_1 // vlssoq041

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/sdk/formats"
)

type SecuritySignOut struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOQ_04_1_1A Security_SignOut"`

	// It contains conversation properties between the SI and the JFE.
	ConversationClt *ConversationIDType `xml:"conversationClt,omitempty"` // minOccurs="0"
}

//
// Complex structs
//

type ConversationIDType struct {
	// Sender identification
	SenderIdentification formats.AlphaNumericString_Length1To35 `xml:"senderIdentification"`

	// Recipient identification
	RecipientIdentification formats.AlphaNumericString_Length1To35 `xml:"recipientIdentification"`

	// Sender's interchange control reference
	SenderInterchangeControlReference formats.AlphaNumericString_Length1To14 `xml:"senderInterchangeControlReference"`

	// Recipient's interchange control reference
	RecipientInterchangeControlReference formats.AlphaNumericString_Length1To14 `xml:"recipientInterchangeControlReference"`
}
