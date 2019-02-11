package amadeus_sdk

import (
	"encoding/xml"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/security/signOut/v04_1_query"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk/security/signOut/v04_1_reply"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/amadeus-agent-go/support"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/amadeus-agent-go/utils"
	"gitlab.teamc.io/tm-consulting/tmc24/provider-logs/receiver.git"
	"strings"
)

const (
	Start = iota
	End
	Rollback
	InSeries
	Continuation
	Subsequent
)

var TransactionStatusCode = [...]string{
	"Start",        // This is the first message within a transaction.
	"End",          // This is the last message within a transaction.
	"Rollback",     // This indicates that all messages within the current transaction must be ignored.
	"InSeries",     // This is any message that is not the first or last message within a transaction.
	"Continuation", // Specifies that this is a followup request asking for more of what was requested in the previous request.
	"Subsequent",   // This request message is a subsequent request based on the previous message sent in this transaction.
}

type Session struct {
	// This attributes defines the status code of the session in a stateful flow.
	TransactionStatusCode string `xml:"TransactionStatusCode,attr,omitempty"`

	// This element defines the identifier part of the SessionId.
	SessionId string `xml:"SessionId,omitempty"`

	// This element defines the sequence number of the SessionId.
	SequenceNumber int `xml:"SequenceNumber,omitempty"`

	// This element defines the SecurityToken of the SessionId.
	SecurityToken string `xml:"SecurityToken,omitempty"`
}

type RequestSession struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/2010/06/Session_v3 Session"`

	*Session
}

const (
	ActSessionSecuritySignOut = "VLSSOQ_04_1_1A"
)

func (client *AmadeusClient) SecuritySignOutV041(attr *receiver.LogAttributes) (*Security_SignOutReply_v04_1.SecuritySignOutReply, *ResponseSOAP4Header, error) {
	var soapAction = ActSessionSecuritySignOut
	var query = new(Security_SignOut_v04_1.SecuritySignOut)
	var reply = new(Security_SignOutReply_v04_1.SecuritySignOutReply)
	var messageId = strings.ToUpper(utils.RandStringBytesMaskImprSrc(22))
	header, err := client.service.Call(soapUrl+soapAction, messageId, query, reply, client.session, support.AddMethod(attr, soapAction), client)
	if err != nil {
		return nil, header, err
	}
	return reply, header, nil
}

func CreateSession() *Session {
	return &Session{
		TransactionStatusCode: TransactionStatusCode[Start],
	}
}

func (client *AmadeusClient) GetSession() *Session {
	return client.session
}

func (client *AmadeusClient) IncSequenceNumber(header *ResponseSOAP4Header) {
	if header != nil {
		client.session = &header.Session
	}
	client.session.SequenceNumber++
}

func (client *AmadeusClient) SessionIsClosed() bool {
	if client == nil || client.session == nil || client.session.TransactionStatusCode == TransactionStatusCode[End] {
		return true
	}
	return false
}

func (client *AmadeusClient) SetSessionEndTransaction() bool {
	if client == nil || client.session == nil || client.session.TransactionStatusCode == TransactionStatusCode[End] {
		return false
	}
	client.session.TransactionStatusCode = TransactionStatusCode[End]
	return true
}

func (client *AmadeusClient) UpdateSession(session *Session) bool {
	if client == nil {
		return false
	}
	client.session = session
	return true
}

func (client *AmadeusClient) CloseSession(attr *receiver.LogAttributes) (reply *Security_SignOutReply_v04_1.SecuritySignOutReply, err error) {
	if client == nil || client.service == nil {
		return
	}
	if client.session != nil && client.session.TransactionStatusCode != TransactionStatusCode[End] {
		var header *ResponseSOAP4Header
		client.session.TransactionStatusCode = TransactionStatusCode[End]
		reply, header, err = client.SecuritySignOutV041(attr)
		if err == nil {
			client.session = nil
		} else {
			client.session = &header.Session
		}
	}
	return
}
