package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/client"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/cancel/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve/v11.3/response"
)

func (s *service) PNRCancel(query *PNR_Cancel_v11_3.Request) (*PNR_Retrieve_v11_3_response.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.PNRCancelV113(query)
}
