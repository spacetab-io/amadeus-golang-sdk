package respstructs

import "encoding/xml"

type QueueListReply struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRR_11_1_1A Queue_ListReply"`

	ErrorReturn *ErrorReturn `xml:"errorReturn,omitempty"`

	QueueView *QueueView `xml:"queueView,omitempty"`
}
