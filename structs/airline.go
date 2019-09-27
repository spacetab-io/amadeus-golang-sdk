package structs

// Airline structure
type Airline struct {
	CodeEng string `json:"code_eng" bson:"codeEng"`
	CodeRus string `json:"code_rus,omitempty" bson:"codeRus,omitempty"`
	NameEng string `json:"name_eng,omitempty" bson:"nameEng,omitempty"`
	NameRus string `json:"name_rus,omitempty" bson:"nameRus,omitempty"`
}

// Aircraft structure
type Aircraft struct {
	Code *string           `json:"code" db:"code" bson:"code"`
	Name map[string]string `json:"name,omitempty" db:"name" bson:"name,omitempty"`
}
