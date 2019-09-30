package service

import (
	"encoding/json"
	"fmt"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs"
	search "github.com/tmconsulting/amadeus-golang-sdk/v2/structs/fare/masterPricerTravelBoardSearch"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs/pnr/retrieve"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/amadeus-golang-sdk/v2/client"
	l "github.com/tmconsulting/amadeus-golang-sdk/v2/logger"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/logger/nilLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/logger/stdoutLogger"
)

var (
	url               string
	originator        string
	passwordRaw       string
	officeId          string
	stdOutLog, logger l.LogWriter
)

func tearUp() {
	url = os.Getenv("URL")
	originator = os.Getenv("ORIGINATOR")
	passwordRaw = os.Getenv("PASSWORD_RAW")
	officeId = os.Getenv("OFFICE_ID")
	logger = nilLogger.Init()
	stdOutLog = stdoutLogger.Init()

	log.Printf("url: %s\noriginator: %s\npasswordRaw: %s\nofficeId: %s", url, originator, passwordRaw, officeId)
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestNewSKD(t *testing.T) {
	t.Run("initiating test", func(t *testing.T) {

		tearUp()

		cl := client.New(client.SetURL(url), client.SetUser(originator), client.SetPassword(passwordRaw), client.SetAgent(officeId), client.SetLogger(logger))

		amadeusSDK := New(cl)

		_, err := amadeusSDK.CommandCryptic("AN20AUGMOWLED/ALH")
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
	t.Run("methods versions test", func(t *testing.T) {
		cl := client.New(client.SetURL(url), client.SetUser(originator), client.SetPassword(passwordRaw), client.SetAgent(officeId), client.SetLogger(logger))

		amadeusSDK := New(cl, SetMethodVersion(CommandCryptic, MethodVersion(123)))

		_, err := amadeusSDK.CommandCryptic("AN20MAYMOWLED/ALH")
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})

	t.Run("PNR_Retrieve test", func(t *testing.T) {

		cl := client.New(client.SetURL(url), client.SetUser(originator), client.SetPassword(passwordRaw), client.SetAgent(officeId), client.SetLogger(logger))

		amadeusSDK := New(cl)
		//amadeusSDK := New(cl, SetMethodVersion(PNRAddMultiElements, MethodVersion(PNRRetrieveV113)))

		request := PNR_Information.Request{
			PNR: "RR3QXE",
		}

		response, _, err := amadeusSDK.PNRRetrieve(&request)
		t.Logf("%+v\n", response)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})

	t.Run("Search test", func(t *testing.T) {

		tearUp()

		cl := client.New(client.SetURL(url), client.SetUser(originator), client.SetPassword(passwordRaw), client.SetAgent(officeId), client.SetLogger(stdOutLog))

		amadeusSDK := New(cl)
		//amadeusSDK := New(cl, SetMethodVersion(PNRAddMultiElements, MethodVersion(PNRRetrieveV113)))

		itinerary := structs.Itinerary{
			DepartureLocation: structs.Location{
				AirportCode: "SVO",
				CityCode:    "MOW",
				CountryCode: "RU",
				Type:        "city",
			},
			ArrivalLocation: structs.Location{
				AirportCode: "LED",
				CityCode:    "LED",
				CountryCode: "RU",
				Type:        "city",
			},
		}
		request := search.SearchRequest{
			ClientData: structs.ClientInfo{
				OfficeID: officeId,
			},
			BaseClass: []string{
				"E",
			},
			Changes:     false,
			WithBaggage: true,
			TravelType:  "OW",
			Itineraries: map[int]*search.Itinerary{
				1: {
					Itinerary: &itinerary,
					DepartureDate: structs.FlightDateTime{
						Date: time.Now().Add(10 * 24 * time.Hour), // add 10 days
					},
					DepartureDateTill: structs.FlightDateTimeUniversal{
						DateStr: time.Now().Add(11 * 24 * time.Hour).String(), // add 10 days
					},
				},
			},
			Currency:   "RUB",
			Passengers: structs.Travellers{ADT: 1, CHD: 0, INF: 0},
			Airlines: []string{
				"SU",
			},
		}

		response, _, err := amadeusSDK.FareMasterPricerTravelBoardSearch(&request)
		t.Logf("response: %+v\n", response)

		aa, _ := json.Marshal(response)
		fmt.Println("Search response: ", string(aa))

		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
}
