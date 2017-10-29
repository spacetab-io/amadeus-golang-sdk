package security_signout

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type SecuritySignOut struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOQ_04_1_1A Security_SignOut"`

	// It contains conversation properties between the SI and the JFE.
	ConversationClt *ConversationIDType `xml:"conversationClt,omitempty"`
}

type ConversationIDType struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOQ_04_1_1A ConversationIDType"`

	// Sender identification
	SenderIdentification formats.AlphaNumericString_Length1To35 `xml:"senderIdentification,omitempty"`

	// Recipient identification
	RecipientIdentification formats.AlphaNumericString_Length1To35 `xml:"recipientIdentification,omitempty"`

	// Sender's interchange control reference
	SenderInterchangeControlReference formats.AlphaNumericString_Length1To14 `xml:"senderInterchangeControlReference,omitempty"`

	// Recipient's interchange control reference
	RecipientInterchangeControlReference formats.AlphaNumericString_Length1To14 `xml:"recipientInterchangeControlReference,omitempty"`
}
