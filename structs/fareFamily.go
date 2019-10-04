package structs

// FareFamilyService FareFamily Service structure
type FareFamilyService struct {
	Name string
	//Description		string
	Type           string
	Classification string
	Group          string
	Status         string //enum("Included","Charge","NotOffered")
}
