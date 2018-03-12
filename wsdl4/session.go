package wsdl4

import (
	"encoding/xml"
)

var TransactionStatusCode = [...]string{
	"Start",		// This is the first message within a transaction.
	"End",			// This is the last message within a transaction.
	"Rollback",		// This indicates that all messages within the current transaction must be ignored.
	"InSeries",		// This is any message that is not the first or last message within a transaction.
	"Continuation",	// Specifies that this is a followup request asking for more of what was requested in the previous request.
	"Subsequent",	// This request message is a subsequent request based on the previous message sent in this transaction.
}

type Session_v3 struct {

	// This attributes defines the status code of the session in a stateful flow.
	TransactionStatusCode string   `xml:"TransactionStatusCode,attr,omitempty"`

	// This element defines the identifier part of the SessionId.
	SessionId             string   `xml:"SessionId,omitempty"`

	// This element defines the sequence number of the SessionId.
	SequenceNumber        string   `xml:"SequenceNumber,omitempty"`

	// This element defines the SecurityToken of the SessionId.
	SecurityToken         string   `xml:"SecurityToken,omitempty"`
}

type RequestSession_v3 struct {
	XMLName               xml.Name `xml:"http://xml.amadeus.com/2010/06/Session_v3 Session"`

	Session_v3
}
