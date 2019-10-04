package Fare_MasterPricerTravelBoardSearchRequest_v16_3

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs"
	search "github.com/tmconsulting/amadeus-golang-sdk/v2/structs/fare/masterPricerTravelBoardSearch"
	"testing"
)

func TestMakeQuery(t *testing.T) {
	t.Run("Add Travelers", func(t *testing.T) {
		request := search.SearchRequest{
			Passengers: structs.Travellers{
				ADT: 1,
				CHD: 0,
				INF: 0,
			},
		}
		paxReference := addTravelers(&request)
		paxReferenceXML, err := xml.Marshal(paxReference)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, "<TravellerReferenceInformationType><ptc>ADT</ptc><traveller><ref>1</ref></traveller></TravellerReferenceInformationType>", string(paxReferenceXML))
	})

	t.Run("Add Travelers", func(t *testing.T) {
		request := search.SearchRequest{
			Passengers: structs.Travellers{
				ADT: 1,
				CHD: 1,
				INF: 0,
			},
		}
		paxReference := addTravelers(&request)
		paxReferenceXML, err := xml.Marshal(paxReference)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, "<TravellerReferenceInformationType><ptc>ADT</ptc><traveller><ref>1</ref></traveller></TravellerReferenceInformationType><TravellerReferenceInformationType><ptc>CHD</ptc><traveller><ref>2</ref></traveller></TravellerReferenceInformationType>", string(paxReferenceXML))
	})

	t.Run("Add Travelers", func(t *testing.T) {
		request := search.SearchRequest{
			Passengers: structs.Travellers{
				ADT: 1,
				CHD: 1,
				INF: 1,
			},
		}
		paxReference := addTravelers(&request)
		paxReferenceXML, err := xml.Marshal(paxReference)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, "<TravellerReferenceInformationType><ptc>ADT</ptc><traveller><ref>1</ref></traveller></TravellerReferenceInformationType><TravellerReferenceInformationType><ptc>CHD</ptc><traveller><ref>2</ref></traveller></TravellerReferenceInformationType><TravellerReferenceInformationType><ptc>INF</ptc><traveller><ref>1</ref><infantIndicator>1</infantIndicator></traveller></TravellerReferenceInformationType>", string(paxReferenceXML))
	})
}
