package soap2_0

import "encoding/xml"

type Session struct {
	//XMLName xml.Name `xml:"http://xml.amadeus.com/ws/2009/01/WBS_Session-2.0.xsd Session"`
	XMLName xml.Name `xml:"Session"`

	// This element defines the identifier part of the SessionId.
	SessionId string `xml:"SessionId"` //,omitempty

	// This element defines the sequence number of the SessionId.
	SequenceNumber string `xml:"SequenceNumber"` //,omitempty

	// This element defines the SecurityToken of the SessionId.
	SecurityToken string `xml:"SecurityToken"` //,omitempty
}
