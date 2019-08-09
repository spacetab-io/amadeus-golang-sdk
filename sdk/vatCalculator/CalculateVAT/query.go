package calculatevat

import "encoding/xml"

// CalculateVAT was auto-generated from WSDL.
type CalculateVAT struct {
	XMLName  xml.Name `xml:"http://amadeus.ru/CalculateVAT CalculateVAT"`
	PINCode  *string  `xml:"PINCode,omitempty" json:"PINCode,omitempty" yaml:"PINCode,omitempty"`
	TSTMasks *string  `xml:"TSTMasks,omitempty" json:"TSTMasks,omitempty" yaml:"TSTMasks,omitempty"`
}
