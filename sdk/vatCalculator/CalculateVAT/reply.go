package calculatevat

// CalculateVATResponse was auto-generated from WSDL.
type CalculateVATResponse struct {
	CalculateVATResult *VATCalculationResult `xml:"CalculateVATResult,omitempty" json:"CalculateVATResult,omitempty" yaml:"CalculateVATResult,omitempty"`
}

// VATCalculationResult was auto-generated from WSDL.
type VATCalculationResult struct {
	VATCalculationData    []TSTVATCalculationResults `xml:"VATCalculationData,omitempty" json:"VATCalculationData,omitempty" yaml:"VATCalculationData,omitempty"`
	VATCalculationMessage *string                    `xml:"VATCalculationMessage,omitempty" json:"VATCalculationMessage,omitempty" yaml:"VATCalculationMessage,omitempty"`
}

type TSTVATCalculationResults struct {
	TSTVATCalculationResult TSTVATCalculationResult `xml:"TSTVATCalculationResult"`
}

// TSTVATCalculationResult was auto-generated from WSDL.
type TSTVATCalculationResult struct {
	VATTSTsID     string  `xml:"VATTSTsID" json:"VATTSTsID,omitempty" yaml:"VATTSTsID,omitempty"`
	VATPAXexID    string  `xml:"VATPAXexID,omitempty" json:"VATPAXexID,omitempty" yaml:"VATPAXexID,omitempty"`
	VATSegmentsID string  `xml:"VATSegmentsID,omitempty" json:"VATSegmentsID,omitempty" yaml:"VATSegmentsID,omitempty"`
	VATCurrency   string  `xml:"VATCurrency,omitempty" json:"VATCurrency,omitempty" yaml:"VATCurrency,omitempty"`
	VATValue      float64 `xml:"VATValue" json:"VATValue" yaml:"VATValue"`
	VATFareValue  float64 `xml:"VATFareValue" json:"VATFareValue" yaml:"VATFareValue"`
	VATTaxValue   float64 `xml:"VATTaxValue" json:"VATTaxValue" yaml:"VATTaxValue"`
	SVATValue     string  `xml:"sVATValue,omitempty" json:"sVATValue,omitempty" yaml:"sVATValue,omitempty"`
	SVATFareValue string  `xml:"sVATFareValue,omitempty" json:"sVATFareValue,omitempty" yaml:"sVATFareValue,omitempty"`
	SVATTaxValue  string  `xml:"sVATTaxValue,omitempty" json:"sVATTaxValue,omitempty" yaml:"sVATTaxValue,omitempty"`
}
