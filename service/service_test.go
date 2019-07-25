package service

import (
	search "github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/amadeus-golang-sdk/client"
	l "github.com/tmconsulting/amadeus-golang-sdk/logger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/nilLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/stdoutLogger"
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

		cl := client.New(client.SetURL(url), client.SetUser(originator), client.SetPassword(passwordRaw), client.SetAgent(officeId), client.SetLogger(stdOutLog))

		amadeusSDK := New(cl)
		//amadeusSDK := New(cl, SetMethodVersion(PNRAddMultiElements, MethodVersion(PNRRetrieveV113)))

		request := search.Request{}

		response, _, err := amadeusSDK.FareMasterPricerTravelBoardSearch(&request)
		t.Logf("%+v\n", response)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
}
