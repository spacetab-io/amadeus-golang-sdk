package Fare_MasterPricerTravelBoardSearchResponse_v16_3

import (
	"encoding/json"
	"fmt"
	structsCommon "github.com/tmconsulting/amadeus-golang-sdk/v2/structs"
	search "github.com/tmconsulting/amadeus-golang-sdk/v2/structs/fare/masterPricerTravelBoardSearch"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/utils"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/utils/convert"

	"sort"
	"strconv"
	"strings"
)

type InternalGroupOfFlights struct {
	ItineraryID       int
	GroupOfSegmentsID int
	Flights           []*structsCommon.Flight
	MajorityCarrier   string
	Duration          int
}

func (r *Response) CheckErrorReply() error {

	if r.ErrorMessage != nil {
		return fmt.Errorf(strings.Join(r.ErrorMessage.ErrorMessageText.Description, "\n"))
	}
	return nil
}

func (reply *Response) ToCommon(request *search.SearchRequest) (*search.SearchResponse, error) {

	return ParseReply(request, reply)
}

// ParseReply Parse Reply from FareMasterPricerTravelBoardSearch
func ParseReply(request *search.SearchRequest, reply *Response) (*search.SearchResponse, error) {

	var requestItineraries []*search.Itinerary
	for i := 1; i <= len(request.Itineraries); i++ {
		var itinerary = request.Itineraries[i]
		requestItineraries = append(requestItineraries, itinerary)
	}

	//var recommendationIds []int
	var groupsRef = make(map[string][][]int)
	for _, recReply := range reply.Recommendation {
		var recommendationID = recReply.ItemNumber.ItemNumberId.Number
		var segmentsRef [][]int
		for _, segmentFlightRef := range recReply.SegmentFlightRef {
			var groupOfSegmentsIds []int
			for _, referencingDetail := range segmentFlightRef.ReferencingDetail {
				if referencingDetail.RefQualifier == "S" {
					groupOfSegmentsIds = append(groupOfSegmentsIds, int(referencingDetail.RefNumber))
				}
			}
			segmentsRef = append(segmentsRef, groupOfSegmentsIds)
		}

		groupsRef[recommendationID] = segmentsRef
	}

	var itineraries, err = ParseListOfFlights(reply.FlightIndex, groupsRef, request.BaseClass[0])
	if err != nil {
		return nil, err
	}
	if itineraries == nil || len(itineraries) == 0 {
		return nil, nil
	}

	var currencyDefault string
	//	Specifies the currency used for pricing
	if reply.ConversionRate != nil {
		if len(reply.ConversionRate.ConversionRateDetail) == 1 {
			currencyDefault = reply.ConversionRate.ConversionRateDetail[0].Currency
		}
	}

	var majorClass string
	if len(request.BaseClass) == 1 {
		majorClass = request.BaseClass[0]
	}

	//	Service fee information per passenger
	var _, freeBagAllowances = parseServiceFeesGrp(reply.ServiceFeesGrp)
	var responseSegmentIds []string
	var recommendations []*search.Recommendation
	var segments = make(map[string]*structsCommon.Flight)

	//	Recommendation details
	for _, recReply := range reply.Recommendation {
		var recommendationID = recReply.ItemNumber.ItemNumberId.Number
		var variantRef []map[int][]int
		var flightsRef = make(map[int]map[int]*InternalGroupOfFlights)
		var baggageRef = make(map[int]int)
		for variantIndex, segmentFlightRef := range recReply.SegmentFlightRef {
			var itineraryID, baggageID = 1, 0
			var groupOfSegmentsIds = make(map[int][]int)
			groupOfSegmentsIds[itineraryID] = []int{}
			for _, referencingDetail := range segmentFlightRef.ReferencingDetail {
				var refNumber = int(referencingDetail.RefNumber)
				switch referencingDetail.RefQualifier {
				case "S":
					//flightsRef[itineraryID] = append(flightsRef[itineraryID], refNumber)
					groupOfSegmentsIds[itineraryID] = append(groupOfSegmentsIds[itineraryID], refNumber)

					group := itineraries[itineraryID][refNumber]
					if flightsRef[itineraryID] == nil {
						flightsRef[itineraryID] = make(map[int]*InternalGroupOfFlights)
					}
					flightsRef[itineraryID][refNumber] = group
					itineraryID++
				case "B":
					baggageID = refNumber
				}
			}
			variantRef = append(variantRef, groupOfSegmentsIds)
			if baggageID != 0 {
				baggageRef[variantIndex] = baggageID
			}
		}

		if len(variantRef) == 0 {
			continue
		}

		var recClass []string

		var validatingAirlines []string
		var validatingAirlineOverride string
		var mapMajorityCarrier = make(map[int]string)
		var matchValidatingAirline = true

		var refundableSlice []string
		var lastTicketSlice []string

		var perPaxPricing = make(structsCommon.PerPaxType)

		var passengerTypes []string
		var adtFareGroups = make(map[int]map[int]*search.FlightDetail)
		var otherFareGroups = make(map[int]map[int]map[string]*search.FlightDetail)

		// Passenger fare product details
		for _, paxFareProduct := range recReply.PaxFareProduct {

			var passengerType string
			if len(paxFareProduct.PaxReference) == 1 {
				for _, paxReference := range paxFareProduct.PaxReference {
					if len(paxReference.Ptc) == 1 {
						for _, ptc := range paxReference.Ptc {
							switch ptc {
							case structsCommon.TravellerTypeCH:
								passengerType = structsCommon.TravellerTypeCHD
							case structsCommon.TravellerTypeIN:
								passengerType = structsCommon.TravellerTypeINF
							default:
								passengerType = ptc
							}
							break
						}
					}
				}
			}
			if passengerType != "" {
				if utils.InArrayString(passengerType, passengerTypes) {
					continue
				}
			}

			for _, fareDetail := range paxFareProduct.FareDetails {
				var itineraryID = int(fareDetail.SegmentRef.SegRef)

				adtFareDetails, exists := adtFareGroups[itineraryID]
				if !exists {
					adtFareDetails = make(map[int]*search.FlightDetail)
				}
				otherFareDetails, exists := otherFareGroups[itineraryID]
				if !exists {
					otherFareDetails = make(map[int]map[string]*search.FlightDetail)
				}

				for flightIndex, fare := range fareDetail.GroupOfFares {

					var passengerType string
					switch fare.ProductInformation.FareProductDetail.PassengerType {
					case structsCommon.TravellerTypeCH:
						passengerType = structsCommon.TravellerTypeCHD
					case structsCommon.TravellerTypeIN:
						passengerType = structsCommon.TravellerTypeINF
					default:
						passengerType = fare.ProductInformation.FareProductDetail.PassengerType
					}

					availability, err := strconv.Atoi(fare.ProductInformation.CabinProduct.AvlStatus)
					if err != nil {
						//log.Println("[strconvError]", "Value is not a number (fare.ProductInformation.CabinProduct.AvlStatus):", err)
						//continue
						availability = 9
					}

					var flightDetail = &search.FlightDetail{
						// Reservation booking designator - BookingClass.
						BookingClass:  fare.ProductInformation.CabinProduct.Rbd,
						Cabin:         fare.ProductInformation.CabinProduct.Cabin,
						FareBasisCode: fare.ProductInformation.FareProductDetail.FareBasis,
						//FareType:      fare.ProductInformation.FareProductDetail.FareType,
						Availability: availability,
					}

					// https://www.biletik.aero/handbook/pomoshch/perelet/klassy-obsluzhivaniya-v-samoletakh-klassy-aviabiletov/

					// ABCDEFGHIJKLMNOPQRSTUVWXYZ
					//  *  + **  ***** * **+****
					//   **    **               *
					// *    *         * +
					switch flightDetail.Cabin {
					case "A", "F", "P", "R":
						flightDetail.FareClass = "F" // 3 = first
					case "C", "D", "I", "J", "Z":
						flightDetail.FareClass = "C" // 2 = business
					default: // "B", "E", "G", "H", "K", "L", "M", "N", "O", "Q", "S", "T", "U", "V", "W", "X", "Y"
						flightDetail.FareClass = "Y" // 1 = economic
					}

					//if majorClass != "" {
					//	if !utils.InArrayString(majorClass, recClass) {
					//		recClass = append(recClass, majorClass)
					//	}
					//} else {
					if !utils.InArrayString(flightDetail.FareClass, recClass) {
						recClass = append(recClass, flightDetail.FareClass)
					}
					//}

					if passengerType == structsCommon.TravellerTypeADT {
						adtFareDetails[flightIndex] = flightDetail
					} else {
						perPaxFareDetails, exists := otherFareDetails[flightIndex]
						if !exists {
							perPaxFareDetails = make(map[string]*search.FlightDetail)
						}
						perPaxFareDetails[passengerType] = flightDetail
						otherFareDetails[flightIndex] = perPaxFareDetails
					}
				}
				adtFareGroups[itineraryID] = adtFareDetails
				otherFareGroups[itineraryID] = otherFareDetails
			}

			for _, paxReference := range paxFareProduct.PaxReference {
				for _, travellerType := range paxReference.Ptc {

					// Company role identification
					for _, codeShareDetails := range paxFareProduct.PaxFareDetail.CodeShareDetails {
						if codeShareDetails.Company == "" {
							continue
						}

						//	Codeset for Transport stage qualifier (Ref: 8051 1A 07.2.10)
						// V - Possible validating carrier – used for price calculation
						// X - Excluded Validating carrier
						// default: Other Possible Validating carrier
						switch codeShareDetails.TransportStageQualifier {
						case "V":
							if matchValidatingAirline {
								if validatingAirlineOverride == "" {
									validatingAirlineOverride = codeShareDetails.Company
								} else if validatingAirlineOverride != codeShareDetails.Company {
									validatingAirlineOverride = ""
									matchValidatingAirline = false

									if exists, _ := utils.InArray(codeShareDetails.Company, validatingAirlines); !exists {
										validatingAirlines = append(validatingAirlines, codeShareDetails.Company)
									}
								}
							} else {
								if exists, _ := utils.InArray(codeShareDetails.Company, validatingAirlines); !exists {
									validatingAirlines = append(validatingAirlines, codeShareDetails.Company)
								}
							}
						case "X":
						default:
							if exists, _ := utils.InArray(codeShareDetails.Company, validatingAirlines); !exists {
								validatingAirlines = append(validatingAirlines, codeShareDetails.Company)
							}
						}
					}

					var tax = 0.0
					if paxFareProduct.PaxFareDetail.TotalTaxAmount != nil {
						tax = *paxFareProduct.PaxFareDetail.TotalTaxAmount
					}
					var pricing = &structsCommon.Pricing{
						Currency: currencyDefault,
						Total:    paxFareProduct.PaxFareDetail.TotalFareAmount,
						Base:     paxFareProduct.PaxFareDetail.TotalFareAmount - tax,
						Tax:      tax,
					}
					for _, monetaryDetail := range paxFareProduct.PaxFareDetail.MonetaryDetails {

						currency := ""
						if monetaryDetail.Currency != currencyDefault || currencyDefault == "" {
							currency = monetaryDetail.Currency
						}

						pricing.AddTax(&structsCommon.Tax{
							//	Codeset for Monetary Amount Type Qualifier (Ref: 5025 1A 12.4.2)
							// A   - Total Additional Collection
							// B   - The Tax (or surcharge in WQ) total Difference in currency of reissue selling
							// BFA - Base fare Amount
							// BGT - Grand Total Balance per Requested Segment
							// BTA - Tax Balance per Requested Segment
							// C   - MCO Residual Value
							// CR  - Converted Recommendation amount used for information and sorting. Conversion rate not guaranteed
							// CT  - Converted Tax amount used for information and sorting. Conversion rate not guaranteed
							// D   - Grand Total Difference amount (includes penalty. May be with or without tax and surcharges depending on the request) in currency of reissue selling
							// F   - Fuel surcharge amount ( New total YQ/YR amount )
							// G   - Total amount to be paid in cash
							// H   - Total of taxes & surcharges to be paid in cash
							// I   - Total amount to be paid in miles
							// J   - Minimum amount to be paid in miles
							// K   - Base Fare Amount in Miles
							// M   - Grand Total amount (includes penalty. May be with or without tax and surcharges depending on the request) in currency of reissue selling
							// MAC - Mileage accrual
							// MIL - Miles
							// N   - New Tax (or surcharge in WQ) total in currency of reissue selling
							// NF  - No Show fee amount
							// OB  - Ticketing fees total
							// OR  - Recommendation amount in original fare currency
							// OT  - Recommendation Tax amount in original fare currency
							// P   - Penalty amount in currency of reissue selling
							// PNB - Penalty Before Netting per Requested Segment
							// PND - PTC non discounted amount
							// PO  - total amount of polled services
							// R   - Real Recommendation amount in published currency
							// RES - Residual Value per Requested Segment
							// S   - Sum penalty
							// T   - Real tax amount in published currency
							// TAC - Total Additional Collection per Requested Segment
							// TP  - Total Penalty amount
							// U   - US taxes amount
							// X   - Max penalty
							// XOB - Total amount without OB fees
							// YQ  - YQ amounts included in taxes
							// YR  - YR amounts included in taxes
							Code:     monetaryDetail.AmountType,
							Amount:   monetaryDetail.Amount,
							Currency: currency,
						})
					}

					switch travellerType {
					case structsCommon.TravellerTypeCH:
						travellerType = structsCommon.TravellerTypeCHD
					case structsCommon.TravellerTypeIN:
						travellerType = structsCommon.TravellerTypeINF
					}

					if pp, exists := perPaxPricing[travellerType]; exists {
						pp.AddUp(pricing)
					} else {
						perPaxPricing[travellerType] = pricing
					}

					for _, fare := range paxFareProduct.Fare {
						//	Codeset for Text subject qualifier
						// 1   - Coded free text
						// 3   - Literal text
						// 4   - Coded and literal text
						// APM - Appended message
						// IFC - Internal fare calc
						// LTD - Last Date to Ticket
						// PEN - Penalties Message
						// SUR - Surcharges
						// SYS - SYSTEM CHECKS
						// TFC - Ticketing fare calc
						// WRN - Warning message
						switch fare.PricingMessage.FreeTextQualification.TextSubjectQualifier {
						case "PEN":
							for _, description := range fare.PricingMessage.Description {
								switch description {
								case "TICKETS ARE NON-REFUNDABLE",
									"TICKETS ARE NON REFUNDABLE",
									"TICKETS ARE NON REFUNDABLE BEFORE DEPARTURE":
									refundableSlice = append(refundableSlice, "No")
								case "TICKETS ARE NON REFUNDABLE AFTER DEPARTURE":
									refundableSlice = append(refundableSlice, "Yes")
								case "PENALTY APPLIES",
									"PENALTY APPLIES - CHECK RULES",
									"PERCENT PENALTY APPLIES",
									"SUBJ TO CANCELLATION/CHANGE PENALTY":
									refundableSlice = append(refundableSlice, "Penalty")
								default:
									//logs.Log.Info("Error Refundable Value:", description)
									// TODO log
								}
							}
						case "LTD":
							for _, description := range fare.PricingMessage.Description {
								if len(description) == 7 {
									if lastTicketDate, err := convert.AmadeusDateConvert(description); err == nil {
										lastTicketStr := lastTicketDate.Format(structsCommon.TimeFormat)
										if exists, _ := utils.InArray(lastTicketStr, lastTicketSlice); !exists {
											lastTicketSlice = append(lastTicketSlice, lastTicketStr)
										}
									}
								}
							}
						}
					}
				}
			}
		}

		if len(adtFareGroups) == 0 {
			for itineraryID, otherFareDetails := range otherFareGroups {
				if len(otherFareDetails) > 0 {
					var adtFareDetails = make(map[int]*search.FlightDetail)
					for flightIndex, perPaxFareDetails := range otherFareDetails {
						if chdFareDetails, exists := perPaxFareDetails[structsCommon.TravellerTypeCHD]; exists {
							adtFareDetails[flightIndex] = chdFareDetails
						}
					}
					if len(adtFareDetails) > 0 {
						adtFareGroups[itineraryID] = adtFareDetails
					}
				}
			}
		}

		//test
		if len(adtFareGroups) == 0 {
			continue
		}

		for itineraryID, adtFareDetails := range adtFareGroups {
			var otherFareDetails = otherFareGroups[itineraryID]
			if len(otherFareDetails) == 0 {
				continue
			}
			for flightIndex, flightDetail := range adtFareDetails {
				var perPaxFareDetails = otherFareDetails[flightIndex]
				var removeKeys []string
				for passengerType, testFareDetails := range perPaxFareDetails {
					if flightDetail.Equal(testFareDetails) {
						removeKeys = append(removeKeys, passengerType)
					}
				}
				if len(removeKeys) > 0 {
					for _, removeKey := range removeKeys {
						delete(perPaxFareDetails, removeKey)
					}
					if len(perPaxFareDetails) == 0 {
						delete(otherFareDetails, flightIndex)
					}
				}
			}
		}

		var interline = false // Рекомендация носит признак interline тогда, когда на одном или более сегментах OA отличается от VA.

		var validatingAirlinesCount = len(validatingAirlines)
		if validatingAirlineOverride == "" && (validatingAirlinesCount == 0 || validatingAirlinesCount > 1) {
			matchValidatingAirline = true
			var airlines []string

			for _, majorityCarrier := range mapMajorityCarrier {
				if exists, _ := utils.InArray(majorityCarrier, airlines); !exists {
					airlines = append(airlines, majorityCarrier)
				}
			}
			if len(airlines) == 1 {
				validatingAirlineOverride = airlines[0]
			}
			for _, validatingAirline := range airlines {
				if exists, _ := utils.InArray(validatingAirline, validatingAirlines); !exists {
					validatingAirlines = append(validatingAirlines, validatingAirline)
				}
			}
		}
		if len(validatingAirlines) == 1 {
			if validatingAirlineOverride == validatingAirlines[0] {
				validatingAirlines = []string{}
			} else if validatingAirlineOverride == "" {
				validatingAirlineOverride = validatingAirlines[0]
				validatingAirlines = []string{}
			}
		}

		var recSegmentIds []string

		var itineraryRef = make(map[int]map[int][]int)
		for variantIndex, groupOfSegmentsIds := range variantRef {
			var variantID = variantIndex + 1
			for itineraryID, groupOfSegmentsID := range groupOfSegmentsIds {
				for _, n := range groupOfSegmentsID {
					var variants, exists = itineraryRef[itineraryID]
					if !exists {
						variants = make(map[int][]int)
					}
					variants[variantID] = append(variants[variantID], n)
					if !exists {
						itineraryRef[itineraryID] = variants
					}
				}
			}
		}

		var routesSegments = make(search.RoutesSegments)
		for itineraryID, itineraryVariants := range itineraryRef {
			var itinerary, existsVariants = routesSegments[itineraryID]
			if !existsVariants {
				itinerary = make(map[int]*search.RouteSegments)
			}
			var adtFareDetails = adtFareGroups[itineraryID]
			var otherFareDetails = otherFareGroups[itineraryID]
			for variantID, groupOfSegmentsIds := range itineraryVariants {
				for _, groupOfSegmentsID := range groupOfSegmentsIds {
					var group = flightsRef[itineraryID][groupOfSegmentsID]
					var variantSegmentID []string
					var itinerarySegments []*search.ItinerarySegment
					for flightIndex, flight := range group.Flights {
						var segmentID = flight.SegmentID
						variantSegmentID = append(variantSegmentID, segmentID)

						var flightDetail = adtFareDetails[flightIndex]
						if flightDetail == nil {
							continue
						}
						var baggage *structsCommon.BaggageType
						if flightDetail.Baggage == nil {
							if num, exists := baggageRef[variantID-1]; exists {
								if freeBagGroups, exists := freeBagAllowances[num]; exists {
									if freeBagLegs, exists := freeBagGroups[itineraryID]; exists {
										baggage = freeBagLegs[flightIndex+1]
									}
								}
							}
							if baggage == nil {
								baggage = &structsCommon.BaggageType{
									Unit: "PC",
								}
							}
						}

						if _, exists := segments[segmentID]; !exists {
							segments[segmentID] = flight
						}

						if !utils.InArrayString(segmentID, recSegmentIds) {
							recSegmentIds = append(recSegmentIds, segmentID)
						}

						if !interline {
							if flight.OperatingAirline.CodeEng != validatingAirlineOverride {
								interline = true
							}
						}

						var paxDetails = otherFareDetails[flightIndex]
						itinerarySegments = append(itinerarySegments, &search.ItinerarySegment{
							SegmentID: segmentID,
							FlightDetail: &search.FlightDetail{
								Class:         flightDetail.Class, // @TODO Deprecate it!
								Cabin:         flightDetail.Cabin,
								FareClass:     flightDetail.FareClass, // @TODO rename it to Subclass !
								BookingClass:  flightDetail.BookingClass,
								FareBasisCode: flightDetail.FareBasisCode,
								Availability:  flightDetail.Availability,
								Baggage:       baggage,
								//FareType:      flightDetail.FareType,
							},
							PaxDetails: paxDetails,
						})
					}

					itinerary[variantID] = &search.RouteSegments{
						GroupOfSegmentsID: groupOfSegmentsID,
						ItinerarySegments: itinerarySegments,
						TravelTime:        group.Duration,
					}
				}
			}
			if !existsVariants && len(itinerary) > 0 {
				routesSegments[itineraryID] = itinerary
			}
		}

		for _, segmentID := range recSegmentIds {
			if !utils.InArrayString(segmentID, responseSegmentIds) {
				responseSegmentIds = append(responseSegmentIds, segmentID)
			}
		}

		if interline {
			for _, variants := range routesSegments {
				for _, variant := range variants {
					for _, itinerarySegment := range variant.ItinerarySegments {
						var flight = segments[itinerarySegment.SegmentID]
						itinerarySegment.CodeShare = flight.CodeShare()
					}
				}
			}
		}

		var lastTicketDate string
		if l := len(lastTicketSlice); l > 0 {
			sort.Strings(lastTicketSlice)
			lastTicketDate = lastTicketSlice[l-1]
		}

		var refundable = "N"
		if exists, _ := utils.InArray("No", refundableSlice); !exists {
			exists, _ := utils.InArray("Penalty", refundableSlice)
			if exists {
				refundable = "P"
			} else {
				exists, _ := utils.InArray("Yes", refundableSlice)
				if exists {
					refundable = "Y"
				}
			}
		}

		if len(recClass) == 0 && majorClass != "" {
			recClass = []string{majorClass}
		}

		var recommendation = &search.Recommendation{
			ID:                recommendationID,
			Provider:          "amadeus",
			Class:             recClass,
			ItinerarySegments: routesSegments,
			ValidatingAirline: structsCommon.Airline{},
			//ValidatingAirlines: validatingAirlines,
			Interline:      interline,
			LastTicketDate: lastTicketDate,
			Refundable:     refundable, //TODO bool?
		}
		recommendation.ValidatingAirline.CodeEng = validatingAirlineOverride

		// To describe type of recReply
		//for _, warningMessage := range recReply.warningMessage {
		//	warning := &errs.warningMessage{}
		//	warning.Parse(warningMessage)
		//	recommendation.WarningMessages = append(recommendation.WarningMessages, warning)
		//	log.Println("Recommendation.warningMessage:", warning.String())
		//}

		// Recommendation Pricing and Taxes
		for num, monetaryDetail := range recReply.RecPriceInfo.MonetaryDetail {
			switch num {
			case 0:
				recommendation.Price.Total = monetaryDetail.Amount
				if recommendation.Price.Total < 2000 {
					if _, err := json.Marshal(&request.Itineraries); err == nil {
						//logs.Log.Warnf("Price too low. Price: %v, Request: %s", recommendation.Price.Total, string(a)) //TODO log
					}
				}
			case 1:
				recommendation.Price.Tax = monetaryDetail.Amount
			default:
				currency := ""
				if monetaryDetail.Currency != currencyDefault || currencyDefault == "" {
					currency = monetaryDetail.Currency
				}
				recommendation.Price.AddTax(&structsCommon.Tax{
					//	Codeset for Monetary Amount Type Qualifier (Ref: 5025 1A 14.2.1)
					// A   - Total Additional Collection
					// B   - The Tax (or surcharge in WQ) total Difference in currency of reissue selling
					// BFA - Base fare Amount
					// BGT - Grand Total Balance per Requested Segment
					// BTA - Tax Balance per Requested Segment
					// C   - MCO Residual Value
					// CR  - Converted Recommendation amount used for information and sorting. Conversion rate not guaranteed
					// CT  - Converted Tax amount used for information and sorting. Conversion rate not guaranteed
					// D   - Grand Total Difference amount (includes penalty. May be with or without tax and surcharges depending on the request) in currency of reissue selling
					// F   - Fuel surcharge amount ( New total YQ/YR amount )
					// FCN - Minimum fee amount for credit card
					// FCX - Maximum fee amount for credit card
					// FDN - Minimum fee amount for debit card
					// FDX - Maximum fee amount for debit card
					// FON - Minimum fee amount for any card
					// FOX - Maximum fee amount for any card
					// G   - Total amount to be paid in cash
					// H   - Total of taxes & surcharges to be paid in cash
					// I   - Total amount to be paid in miles
					// J   - Minimum amount to be paid in miles
					// K   - Base Fare Amount in Miles
					// M   - Grand Total amount (includes penalty. May be with or without tax and surcharges depending on the request) in currency of reissue selling
					// MAC - Mileage accrual
					// MIL - Miles
					// N   - New Tax (or surcharge in WQ) total in currency of reissue selling
					// NF  - No Show fee amount
					// OB  - Ticketing fees total
					// OR  - Recommendation amount in original fare currency
					// OT  - Recommendation Tax amount in original fare currency
					// P   - Penalty amount in currency of reissue selling
					// PNB - Penalty Before Netting per Requested Segment
					// PND - PTC non discounted amount
					// PO  - total amount of polled services
					// R   - Real Recommendation amount in published currency
					// RES - Residual Value per Requested Segment
					// S   - Sum penalty
					// T   - Real tax amount in published currency
					// TAC - Total Additional Collection per Requested Segment
					// TP  - Total Penalty amount
					// U   - US taxes amount
					// X   - Max penalty
					// XOB - Total amount without OB fees
					// YQ  - YQ amounts included in taxes
					// YR  - YR amounts included in taxes
					Code:     monetaryDetail.AmountType,
					Amount:   monetaryDetail.Amount,
					Currency: currency,
				})
			}
		}
		recommendation.Price.Currency = currencyDefault
		recommendation.Price.PerPax = &perPaxPricing
		recommendation.Price.RoundUp()
		recommendations = append(recommendations, recommendation)
	}

	if len(recommendations) == 0 || len(segments) == 0 {
		return nil, nil
	}

	var result = &search.SearchResponse{
		Segments: make(map[string]*structsCommon.Flight),
	}

	for segmentID, segment := range segments {
		if utils.InArrayString(segmentID, responseSegmentIds) {
			result.Segments[segmentID] = segment
		}
	}

	for _, recommendation := range recommendations {
		result.Recommendations = append(result.Recommendations, recommendation)
	}

	recommendations = cleanRecomendationsFromDuplicatedItinerarySegments(recommendations)

	return result, nil
}

// ParseListOfFlights Parse List Of Flights
func ParseListOfFlights(listOfFlights []*FlightIndex, recommendationGroupsRef map[string][][]int, baseclass string) (map[int]map[int]*InternalGroupOfFlights, error) {

	// (itineraryId) int = flightIndex[].RequestedSegmentRef.SegRef
	// (groupOfSegmentsId) int = flightIndex[].GroupOfFlights[].PropFlightGrDetail.FlightProposal[].Ref
	var itineraries = make(map[int]map[int]*InternalGroupOfFlights)

	for _, flightIndex := range listOfFlights {
		var flightNum = 1
		var itineraryID = int(flightIndex.RequestedSegmentRef.SegRef)

		// List of proposed segments per requested segment
		for _, groupOfFlights := range flightIndex.GroupOfFlights {
			var groupOfSegmentsID = -1
			var variant = &InternalGroupOfFlights{
				ItineraryID: itineraryID,
			}

			// Parameters for proposed flight group
			for _, flightProposal := range groupOfFlights.PropFlightGrDetail.FlightProposal {
				//	Codeset for Type of Units Qualifier (Ref: 6353 1A 09.1.7)
				// CDC - Cross-over date combi
				// DUR - Duration for Rail journey
				// EFT - Elapse Flying Time
				// MCX - Majority carrier
				// NAV - No Availability
				// NFA - No fare
				// NIT - No Itinerary
				// NJO - No Journey
				// RJS - Rank in Journey Server
				switch flightProposal.UnitQualifier {
				case "EFT":
					variant.Duration = convert.AmadeusTimeToMinutes(flightProposal.Ref)
				case "MCX":
					variant.MajorityCarrier = flightProposal.Ref
				case "":
					if ref, err := strconv.Atoi(flightProposal.Ref); err == nil {
						groupOfSegmentsID = ref
					} else {
						//logs.Log.WithError(err).Error("Value is not a number (flightProposal.Ref)")
						// TODO log
					}
				}
			}
			if groupOfSegmentsID < 0 {
				//errorFlightCount++
				continue
			}
			variant.GroupOfSegmentsID = groupOfSegmentsID

			// List of flight per proposed segment
			for _, flightDetail := range groupOfFlights.FlightDetails {
				var flight = &structsCommon.Flight{
					SegmentID: fmt.Sprintf("amadeus_%d_%d_%s", itineraryID, flightNum, baseclass),
				}

				// Location of departure and arrival
				for locationID, location := range flightDetail.FlightInformation.Location {
					if locationID == 0 {
						flight.DepartureLocation.AirportCode = location.LocationId
						flight.DepartureLocation.Terminal = location.Terminal
					} else {
						flight.ArrivalLocation.AirportCode = location.LocationId
						flight.ArrivalLocation.Terminal = location.Terminal
					}
				}

				// Date and time of departure and arrival
				var dateOfDeparture = flightDetail.FlightInformation.ProductDateTime.DateOfDeparture
				var timeOfDeparture = flightDetail.FlightInformation.ProductDateTime.TimeOfDeparture
				flight.DepartureDate = convert.AmadeusDateTimeConvert(dateOfDeparture, timeOfDeparture)
				var dateOfArrival = flightDetail.FlightInformation.ProductDateTime.DateOfArrival
				var timeOfArrival = flightDetail.FlightInformation.ProductDateTime.TimeOfArrival
				flight.ArrivalDate = convert.AmadeusDateTimeConvert(dateOfArrival, timeOfArrival)
				//flight.TravelTime    = int(flight.ArrivalDate.Sub(flight.DepartureDate) / time.Minute)

				// Company identification
				flight.OperatingAirline = &structsCommon.Airline{}
				flight.OperatingAirline.CodeEng = flightDetail.FlightInformation.CompanyId.OperatingCarrier

				flight.MarketingAirline = &structsCommon.Airline{}
				flight.MarketingAirline.CodeEng = flightDetail.FlightInformation.CompanyId.MarketingCarrier

				flight.FlightNumber = flightDetail.FlightInformation.FlightOrtrainNumber
				flight.Aircraft = &structsCommon.Aircraft{} // Type of aircraft
				flight.Aircraft.Code = &flightDetail.FlightInformation.ProductDetail.EquipmentType

				variant.Flights = append(variant.Flights, flight)
				flightNum++
			}

			if len(variant.Flights) == 0 {
				continue
			}

			if dp, ok := itineraries[itineraryID]; ok {
				dp[groupOfSegmentsID] = variant
			} else {
				dp = make(map[int]*InternalGroupOfFlights)
				dp[groupOfSegmentsID] = variant
				itineraries[itineraryID] = dp
			}
		}
		if _, ok := itineraries[itineraryID]; !ok {
			return nil, nil
		}
	}

	return itineraries, nil
}

func parseServiceFeesGrp(serviceFeesGrp []*ServiceFeesGrp) (services map[string]*structsCommon.FareFamilyService, serviceCoverageInfoGrp map[int]map[int]map[int]*structsCommon.BaggageType) {

	services = make(map[string]*structsCommon.FareFamilyService)
	//serviceCoverageInfoGrp := make(map[int]map[int]map[int]*structsCommon.BaggageType)

	for _, serviceFeesGroup := range serviceFeesGrp {

		//	Codeset for Option (Ref: 9750 1A 13.1.2)
		// FBA - Free baggage allowance
		// OA  - Booking fees
		// OB  - Ticketing fees
		// OC  - Service fees
		// SSR - Service request
		switch serviceFeesGroup.ServiceTypeInfo.CarrierFeeDetails.Type {
		case "OC":

			//	Description of applicable services
			for _, serviceDetail := range serviceFeesGroup.ServiceDetailsGrp {
				number := serviceDetail.FeeDescriptionGrp.ItemNumberInfo.ItemNumberDetails.Number

				name, classification, group := "", "", ""
				if serviceDetail.FeeDescriptionGrp != nil {

					//	Other service information (service description, ...)
					if serviceDetail.FeeDescriptionGrp.ServiceDescriptionInfo != nil {

						//	To specify the Service Classification of the Service Requirement.
						//	Codeset for Special requirement type (Ref: 9962 IA 08.3.21)
						// F   - flight related. Must be associated to a flight
						// M   - merchandise
						// R   - Rule busted. Must be associated to a fare component
						// T   - Ticket. Must be associated to a ticket
						// UNK - Unknown
						classification = serviceDetail.FeeDescriptionGrp.ServiceDescriptionInfo.ServiceRequirementsInfo.ServiceClassification

						//	Specify the Service group.
						group = serviceDetail.FeeDescriptionGrp.ServiceDescriptionInfo.ServiceRequirementsInfo.ServiceGroup
					}

					if serviceDetail.FeeDescriptionGrp.CommercialName != nil {
						name = serviceDetail.FeeDescriptionGrp.CommercialName.FreeText
					}
				}

				services[number] = &structsCommon.FareFamilyService{
					Name:           name,
					Type:           serviceDetail.ServiceOptionInfo.DataTypeInformation.SubType,
					Classification: classification,
					Group:          group,
				}
			}
		case "FBA":
			serviceCoverageInfoGrp = parseFreeBagAllowance(serviceFeesGroup)
		}
	}
	return services, serviceCoverageInfoGrp
}

func parseFreeBagAllowance(serviceFeesGroup *ServiceFeesGrp) map[int]map[int]map[int]*structsCommon.BaggageType {

	// int    = freeBagAllowanceGrp[].ItemNumberInfo.ItemNumberDetails[].Number
	// string = freeBaggageAllowance
	freeBagAllowanceGroup := make(map[int]*structsCommon.BaggageType)

	//	Free baggage allowance information group
	for _, bagAllowance := range serviceFeesGroup.FreeBagAllowanceGrp {
		var freeBagAllowance structsCommon.BaggageType
		freeBagAllowance.Value = int(*bagAllowance.FreeBagAllownceInfo.BaggageDetails.FreeAllowance)

		//	Codeset for Allowance or charge qualifier (Ref: 5463 1A 11.1.44)
		// N - Number of pieces
		// W - Weight
		switch bagAllowance.FreeBagAllownceInfo.BaggageDetails.QuantityCode {
		case "N":
			//if freeBagAllowance.Value != 0 {
			freeBagAllowance.Unit = "PC"
			//}
		case "W":
			//	Codeset for Measure unit qualifier (Ref: 6411 1A 11.1.327)
			// K - Kilograms
			// L - Pounds
			switch bagAllowance.FreeBagAllownceInfo.BaggageDetails.UnitQualifier {
			case "K":
				freeBagAllowance.Unit = "Kilograms"
			case "L":
				freeBagAllowance.Unit = "Pounds"
			}
		}

		for _, itemNumberDetails := range bagAllowance.ItemNumberInfo.ItemNumberDetails {
			freeBagAllowanceGroup[int(*itemNumberDetails.Number)] = &freeBagAllowance
		}

		// bagAllowance.FreeBagAllownceInfo.BagTagDetails - ignore...
	}

	// int    = serviceCoverageInfoGrp[].ItemNumberInfo.ItemNumber.Number
	// int    = serviceCoverageInfoGrp[].serviceCovInfoGrp[].coveragePerFlightsInfo[].numberOfItemsDetails.refNum
	// int    = serviceCoverageInfoGrp[].serviceCovInfoGrp[].coveragePerFlightsInfo[].lastItemsDetails[].refOfLeg
	// string = freeBaggageAllowance
	freeBagAllowances := make(map[int]map[int]map[int]*structsCommon.BaggageType)

	//	Service coverage information per passenger
	for _, coverageInfoGroup := range serviceFeesGroup.ServiceCoverageInfoGrp {

		//	Item reference number for service coverage details
		number, err := strconv.Atoi(coverageInfoGroup.ItemNumberInfo.ItemNumber.Number)
		if err != nil {
			//logs.Log.WithError(err).Error("Value is not a number (coverageInfoGroup.ItemNumberInfo.ItemNumber.Number)")
			// todo log
			continue
		}

		freeBagGroups, existsGroups := freeBagAllowances[number]
		if !existsGroups || freeBagGroups == nil {
			freeBagGroups = make(map[int]map[int]*structsCommon.BaggageType)
		}

		//	Service coverage information group
		for _, serviceCovInfoGrp := range coverageInfoGroup.ServiceCovInfoGrp {

			//	Service reference number
			for _, referencingDetail := range serviceCovInfoGrp.RefInfo.ReferencingDetail {

				//	Codeset list: 1153 1A 13.1.5
				// F - Free baggage allowance reference
				// S - Service reference number
				if referencingDetail.RefQualifier == "F" {
					freeBagAllowance, exists := freeBagAllowanceGroup[int(referencingDetail.RefNumber)]
					if !exists {
						continue
					}

					//	Service coverage information at flight level Matched seat characteristics
					for _, coveragePerFlightsInfo := range serviceCovInfoGrp.CoveragePerFlightsInfo {
						if coveragePerFlightsInfo.NumberOfItemsDetails.ReferenceQualifier == "RS" {

							itineraryID, err := strconv.Atoi(coveragePerFlightsInfo.NumberOfItemsDetails.RefNum)
							if err != nil {
								continue
							}

							var firstItem = 0
							var lastItems = 0
							var refOfLegs []int

							for _, lastItemsDetails := range coveragePerFlightsInfo.LastItemsDetails {
								if lastItemsDetails.FirstItemIdentifier != nil {
									firstItem = int(*lastItemsDetails.FirstItemIdentifier)
								}
								if lastItemsDetails.LastItemIdentifier != nil {
									lastItems = int(*lastItemsDetails.LastItemIdentifier)
								}
								if lastItemsDetails.RefOfLeg != "" {
									if refOfLeg, err := strconv.Atoi(lastItemsDetails.RefOfLeg); err == nil {
										refOfLegs = append(refOfLegs, refOfLeg)
									}
								}
							}

							if firstItem != 0 && lastItems != 0 {
								for i := firstItem; i <= lastItems; i++ {
									refOfLegs = append(refOfLegs, i)
								}
							}

							if len(refOfLegs) == 0 {
								refOfLegs = []int{1}
							} else {
								refOfLegs = utils.UniqueIntSlice(refOfLegs)
							}

							freeBagLegs, exists := freeBagGroups[itineraryID]
							if !exists || freeBagLegs == nil {
								freeBagLegs = make(map[int]*structsCommon.BaggageType)
							}

							for _, refOfLeg := range refOfLegs {
								freeBagLegs[refOfLeg] = freeBagAllowance
							}

							freeBagGroups[itineraryID] = freeBagLegs
						}
					}
				}
			}
		}

		if !existsGroups {
			freeBagAllowances[number] = freeBagGroups
		}
	}
	return freeBagAllowances
}

func cleanRecomendationsFromDuplicatedItinerarySegments(recommendations []*search.Recommendation) []*search.Recommendation {
	for _, recommendation := range recommendations {
		for i, routesSegment := range recommendation.ItinerarySegments {
			var keys []int
			for k := range recommendation.ItinerarySegments[i] {
				keys = append(keys, k)
			}
			sort.Ints(keys)

			for _, k := range keys {
				if routesSegment[k] == nil {
					continue
				}

				for m := k + 1; m < len(keys)+1; m++ {
					if routesSegment[m] == nil {
						continue
					}

					if len(routesSegment[k].ItinerarySegments) != len(routesSegment[m].ItinerarySegments) {
						continue
					}

					equal := true

					for j, s1 := range routesSegment[k].ItinerarySegments {
						if s1.SegmentID != routesSegment[m].ItinerarySegments[j].SegmentID { // тут можно добавить условий для сравнения
							equal = false
						}
					}

					if equal {
						delete(recommendation.ItinerarySegments[i], m)
					}
				}
			}
		}
	}

	return recommendations
}
