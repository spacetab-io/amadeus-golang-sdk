package structs

import (
	"encoding/json"
)

// ClientInfo customer info structure
type ClientInfo struct {
	OfficeID string `json:"office_id"`
	Timezone string `json:"timezone"`
}

// UnmarshalJSON overrides UnmarshalJSON
func (c *ClientInfo) UnmarshalJSON(data []byte) error {
	type ClientInfoInternal struct {
		Var1 string `json:"amadeus_office_id"`
		Var2 string `json:"office_id"`
	}
	var clientInfo ClientInfoInternal
	if err := json.Unmarshal(data, &clientInfo); err != nil {
		//logs.Log.WithError(err).Error("Error unmarshal json")
		return err
	}
	var OfficeID = ""
	if clientInfo.Var1 != "" {
		OfficeID = clientInfo.Var1
	} else {
		OfficeID = clientInfo.Var2
	}
	c.OfficeID = OfficeID
	return nil
}
