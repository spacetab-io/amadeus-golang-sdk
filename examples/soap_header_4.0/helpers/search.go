package helpers

import (
	"github.com/tmconsulting/amadeus-golang-sdk/formats"
	fmptbq143 "github.com/tmconsulting/amadeus-golang-sdk/reqstructs/fare_masterpricertravelboardsearch"
)

func MakeSearchQuery(cabinClass, departureLocationId, arrivalLocationId, date string) *fmptbq143.FareMasterPricerTravelBoardSearch {

	return &fmptbq143.FareMasterPricerTravelBoardSearch{
		NumberOfUnit: &fmptbq143.NumberOfUnitsType{
			UnitNumberDetail: []*fmptbq143.NumberOfUnitDetailsType_260583C{
				{
					NumberOfUnits: formats.NumericInteger_Length1To6(1),
					TypeOfUnit: formats.AlphaNumericString_Length1To3("PX"),
				},
				{
					NumberOfUnits: formats.NumericInteger_Length1To6(200),
					TypeOfUnit: formats.AlphaNumericString_Length1To3("RC"),
				},
			},
		},
		PaxReference: []*fmptbq143.TravellerReferenceInformationType{
			{
				Ptc: []formats.AlphaNumericString_Length1To6{
					formats.AlphaNumericString_Length1To6("ADT"),
				},
				Traveller: []*fmptbq143.TravellerDetailsType{
					{
						Ref: formats.NumericInteger_Length1To3(1),
					},
				},
			},
		},
		FareOptions: &fmptbq143.FareOptions{
			PricingTickInfo: &fmptbq143.PricingTicketingDetailsType{
				PricingTicketing: &fmptbq143.PricingTicketingInformationType{
					PriceType: []formats.AlphaNumericString_Length0To3{
						formats.AlphaNumericString_Length0To3("RU"),  // Unifares
						formats.AlphaNumericString_Length0To3("ET"),
						formats.AlphaNumericString_Length0To3("NAD"),  // No Airline Distribution
						formats.AlphaNumericString_Length0To3("TAC"),  // Ticket ability check
						formats.AlphaNumericString_Length0To3("RP"),
					},
				},
			},
		},
		TravelFlightInfo: &fmptbq143.TravelFlightInformationType_185853S{
			CabinId: &fmptbq143.CabinIdentificationType_233500C{
				//CabinQualifier: formats.AlphaNumericString_Length1To2("MD"),
				Cabin: []formats.AlphaString_Length0To1{
					formats.AlphaString_Length0To1(cabinClass),
				},
			},
		},
		Itinerary: []*fmptbq143.Itinerary{
			{
				RequestedSegmentRef: &fmptbq143.OriginAndDestinationRequestType{
					SegRef: formats.NumericInteger_Length1To2(1),
				},
				DepartureLocalization: &fmptbq143.DepartureLocationType{
					DepMultiCity: []*fmptbq143.MultiCityOptionType{
						{
							LocationId: formats.AlphaString_Length3To5(departureLocationId),
						},
					},
				},
				ArrivalLocalization: &fmptbq143.ArrivalLocalizationType{
					ArrivalMultiCity: []*fmptbq143.MultiCityOptionType{
						{
							LocationId: formats.AlphaString_Length3To5(arrivalLocationId),
						},
					},
				},
				TimeDetails: &fmptbq143.DateAndTimeInformationType_181295S{
					FirstDateTimeDetail: &fmptbq143.DateAndTimeDetailsTypeI{
						Date: formats.Date_DDMMYY(date),
					},
				},
			},
		},
	}
}
