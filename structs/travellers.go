package structs

// TravellerTypes Ptc constants
const (
	TravellerTypeADT = "ADT"
	TravellerTypeCH  = "CH"
	TravellerTypeCHD = "CHD"
	TravellerTypeINF = "INF"
	TravellerTypeIN  = "IN"
)

// Travellers types structure
type Travellers struct {
	ADT int `json:"ADT,omitempty"`
	CHD int `json:"CHD,omitempty"`
	INF int `json:"INF,omitempty"`
}
