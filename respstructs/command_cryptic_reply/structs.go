package command_cryptic_reply

//import "encoding/xml"

type CommandCrypticReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/HSFRES_07_3_1A Command_CrypticReply"`

	MessageActionDetails *MessageActionDetails `xml:"messageActionDetails,omitempty"`  // minOccurs="0"

	LongTextString *LongTextString `xml:"longTextString"`
}

type MessageActionDetails struct {

	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails"`

	// Format limitations: an..3
	ResponseType string `xml:"responseType,omitempty"`  // minOccurs="0"
}

type MessageFunctionDetails struct {

	// Format limitations: an..3
	BusinessFunction string `xml:"businessFunction,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	MessageFunction string `xml:"messageFunction"`
}

type LongTextString struct {

	// Format limitations: an..9999
	TextStringDetails string `xml:"textStringDetails"`
}
