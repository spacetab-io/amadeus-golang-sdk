package command_cryptic_reply

import "encoding/xml"

type CommandCrypticReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/HSFRES_07_3_1A Command_CrypticReply"`

	MessageActionDetails struct {
		MessageFunctionDetails struct {
			BusinessFunction string `xml:"businessFunction,omitempty"`

			MessageFunction string `xml:"messageFunction,omitempty"`
		} `xml:"messageFunctionDetails,omitempty"`

		ResponseType string `xml:"responseType,omitempty"`
	} `xml:"messageActionDetails,omitempty"`

	LongTextString struct {
		TextStringDetails string `xml:"textStringDetails,omitempty"`
	} `xml:"longTextString,omitempty"`
}
