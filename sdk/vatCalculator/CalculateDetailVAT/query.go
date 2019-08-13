package calculatedetailvat

import "encoding/xml"

// CalculateDetailVAT was auto-generated from WSDL.
type CalculateDetailVAT struct {
	XMLName  xml.Name `xml:"http://amadeus.ru/ CalculateDetailVAT"`
	PINCode  *string  `xml:"PINCode,omitempty" json:"PINCode,omitempty" yaml:"PINCode,omitempty"`
	TSTMasks *string  `xml:"TSTMasks,omitempty" json:"TSTMasks,omitempty" yaml:"TSTMasks,omitempty"`
}
