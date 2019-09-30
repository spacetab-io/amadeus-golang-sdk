package structs

import (
	"encoding/json"
	"math"
)

// Price structure
type Price struct {
	Currency string  `json:"currency,omitempty"`
	Total    float64 `json:"total_amount"`
	Base     float64 `json:"fare_amount"`
	Tax      float64 `json:"tax_amount"`
	Taxes    []*Tax  `json:"taxes,omitempty"`
	//Fee		float64		`json:"fee_amount"`
	//Vats		[]*Vat		`json:"vats"`
	PerPax *PerPaxType `json:"pax_fare,omitempty"`
}

// Amount structure
type Amount struct {
	Value    int64  `json:"amount"`
	Currency string `json:"currency"`
}

// MarshalJSON overrides MarshalJSON
func (price *Price) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Total  Amount      `json:"total_amount"`
		Base   Amount      `json:"fare_amount"`
		Tax    Amount      `json:"tax_amount"`
		Taxes  []*Tax      `json:"taxes,omitempty"`
		PerPax *PerPaxType `json:"pax_fare,omitempty"`
	}{
		Total: Amount{
			Currency: price.Currency,
			Value:    int64(price.Total * 100),
		},
		Base: Amount{
			Currency: price.Currency,
			Value:    int64(price.Base * 100),
		},
		Tax: Amount{
			Currency: price.Currency,
			Value:    int64(price.Tax * 100),
		},
		Taxes:  price.Taxes,
		PerPax: price.PerPax,
	})
}

// UnmarshalJSON overrides UnmarshalJSON
func (price *Price) UnmarshalJSON(data []byte) error {
	return nil
}

// RoundUp rounds price to greater value
func (price *Price) RoundUp() {
	price.Base = math.Ceil(price.Total - price.Tax)
	price.Tax = math.Ceil(price.Tax)

	if totalDiff := math.Abs((price.Base + price.Tax) - price.Total); totalDiff != 0.0 { // totalDiff < 2.0
		price.Total = price.Base + price.Tax
	}
}

// AddTax Add Tax to Price.Taxes
func (price *Price) AddTax(tax *Tax) {
	price.Taxes = append(price.Taxes, tax)
}

// Equal compare price.Total
func (price *Price) Equal(price2 *Price) bool {
	if price.Total == price2.Total {
		return true
	}
	return false
}

// PerPaxType map [Ptc] of Pricing
type PerPaxType map[string]*Pricing

// Pricing structure
type Pricing struct {
	Currency           string
	Total              float64
	Base               float64
	Tax                float64
	Taxes              []*Tax
	NumberOfPassengers int
	//Type				string
}

// MarshalJSON overrides MarshalJSON
func (pricing *Pricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Total Amount `json:"total_amount"`
		Base  Amount `json:"fare_amount"`
		Tax   Amount `json:"tax_amount"`
		Taxes []*Tax `json:"taxes,omitempty"`
	}{
		Total: Amount{
			Currency: pricing.Currency,
			Value:    int64(pricing.Total * 100),
		},
		Base: Amount{
			Currency: pricing.Currency,
			Value:    int64(pricing.Base * 100),
		},
		Tax: Amount{
			Currency: pricing.Currency,
			Value:    int64(pricing.Tax * 100),
		},
		Taxes: pricing.Taxes,
	})
}

// UnmarshalJSON overrides UnmarshalJSON
func (pricing *Pricing) UnmarshalJSON(data []byte) error {
	return nil
}

// AddUp add pricing.Taxes values rounded to greater
func (pricing *Pricing) AddUp(plus *Pricing) {
	pricing.Base = math.Ceil(pricing.Base + plus.Base)
	pricing.Tax = math.Ceil(pricing.Tax + plus.Tax)

	pricing.Total = pricing.Base + pricing.Tax

	for _, tax := range plus.Taxes {
		pricing.Taxes = append(pricing.Taxes, tax)
	}
}

// AddTax add pricing.Taxes
func (pricing *Pricing) AddTax(tax *Tax) {
	pricing.Taxes = append(pricing.Taxes, tax)
}

// Tax structure
type Tax struct {
	Amount   float64 `json:"amount"`
	Code     string  `json:"code"`
	Country  string  `json:"country,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// MarshalJSON overrides MarshalJSON
func (tax *Tax) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Amount `json:"costs"`
		Code   string `json:"code"`
	}{
		Amount: Amount{
			Currency: tax.Currency,
			Value:    int64(tax.Amount * 100),
		},
		Code: tax.Code,
	})
}

// UnmarshalJSON overrides UnmarshalJSON
func (tax *Tax) UnmarshalJSON(data []byte) error {
	return nil
}

// PriceV2 ver 2
type PriceV2 struct {
	Total Amount   `json:"total_amount"`
	Base  Amount   `json:"fare_amount"`
	Tax   Amount   `json:"tax_amount"`
	Taxes []*TaxV2 `json:"taxes,omitempty"`
	//Fee		Amount			`json:"fee_amount"`
	//Vats	[]*VatV2		`json:"vats"`
	PerPax *PerPaxTypeV2 `json:"pax_fare,omitempty"`
}

// TaxV2 Tax ver 2
type TaxV2 struct {
	Amount `json:"costs"`
	Code   string `json:"code"`   // taxType.isoCountry
	Nature string `json:"nature"` // taxNature
}

// PerPaxTypeV2 map of [Ptc]  PricingV2
type PerPaxTypeV2 map[string]*pricingV2

type pricingV2 struct {
	Total Amount   `json:"total_amount"`
	Base  Amount   `json:"fare_amount"`
	Tax   Amount   `json:"tax_amount"`
	Taxes []*TaxV2 `json:"taxes,omitempty"`
	Vats  []*VatV2 `json:"vats,omitempty"`
}

//VatV2 Новая структура НДС
type VatV2 struct {
	Rate   uint   `json:"rate"`
	Amount int64  `json:"amount"`
	Base   string `json:"base"`
}
