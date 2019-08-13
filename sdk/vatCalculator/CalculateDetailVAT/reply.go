package calculatedetailvat

// CalculateDetailVATResponse was auto-generated from WSDL.
type CalculateDetailVATResponse struct {
	CalculateDetailVATResult *VATDetailCalculationResult `xml:"CalculateDetailVATResult,omitempty" json:"CalculateDetailVATResult,omitempty" yaml:"CalculateDetailVATResult,omitempty"`
}

// VATDetailCalculationResult was auto-generated from WSDL.
type VATDetailCalculationResult struct {
	VATDetailCalculationData []TSTVATDetailCalculationResult `xml:"VATDetailCalculationData,omitempty" json:"VATDetailCalculationData,omitempty" yaml:"VATDetailCalculationData,omitempty"`
	VATCalculationMessage    *string                         `xml:"VATCalculationMessage,omitempty" json:"VATCalculationMessage,omitempty" yaml:"VATCalculationMessage,omitempty"`
}

//type TSTVATDetailCalculationResults struct {
//	TSTVATDetailCalculationResult TSTVATDetailCalculationResult `xml:"TSTVATDetailCalculationResult,omitempty" json:"TSTVATDetailCalculationResult,omitempty" yaml:"TSTVATDetailCalculationResult,omitempty"`
//}

// TSTVATDetailCalculationResult was auto-generated from WSDL.
type TSTVATDetailCalculationResult struct {
	VATTSTsID                  string                           `xml:"VATTSTsID,omitempty" json:"VATTSTsID,omitempty" yaml:"VATTSTsID,omitempty"`
	VATPAXexID                 string                           `xml:"VATPAXexID,omitempty" json:"VATPAXexID,omitempty" yaml:"VATPAXexID,omitempty"`
	VATSegmentsID              string                           `xml:"VATSegmentsID,omitempty" json:"VATSegmentsID,omitempty" yaml:"VATSegmentsID,omitempty"`
	VATDetailCalculationResult []VATItemDetailCalculationResult `xml:"VATDetailCalculationResult,omitempty" json:"VATDetailCalculationResult,omitempty" yaml:"VATDetailCalculationResult,omitempty"`
}

//type VATItemDetailCalculationResults struct {
//	VATItemDetailCalculationResult VATItemDetailCalculationResult `xml:"VATItemDetailCalculationResult,omitempty" json:"VATItemDetailCalculationResult,omitempty" yaml:"VATItemDetailCalculationResult,omitempty"`
//}

// VATItemDetailCalculationResult was auto-generated from WSDL.
type VATItemDetailCalculationResult struct {
	VATBaseType string  `xml:"VATBaseType,omitempty" json:"VATBaseType,omitempty" yaml:"VATBaseType,omitempty"`
	VATBaseCode string  `xml:"VATBaseCode,omitempty" json:"VATBaseCode,omitempty" yaml:"VATBaseCode,omitempty"`
	VATCurrency string  `xml:"VATCurrency,omitempty" json:"VATCurrency,omitempty" yaml:"VATCurrency,omitempty"`
	VATBase     float64 `xml:"VATBase" json:"VATBase" yaml:"VATBase"`
	VATRate     float64 `xml:"VATRate" json:"VATRate" yaml:"VATRate"`
	VATValue    float64 `xml:"VATValue" json:"VATValue" yaml:"VATValue"`
}
