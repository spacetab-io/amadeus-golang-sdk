package command_cryptic_reply

import "encoding/xml"

type CommandCrypticReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/HSFRES_07_3_1A Command_CrypticReply"`

	MessageActionDetails *MessageActionDetails `xml:"messageActionDetails,omitempty"`

	LongTextString *LongTextString `xml:"longTextString,omitempty"`
}

type MessageActionDetails struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`

	ResponseType string `xml:"responseType,omitempty"`
}

type MessageFunctionDetails struct {
	BusinessFunction string `xml:"businessFunction,omitempty"`

	MessageFunction string `xml:"messageFunction,omitempty"`
}

type LongTextString struct {
	TextStringDetails string `xml:"textStringDetails,omitempty"`
}
