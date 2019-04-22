package service

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/amadeus-golang-sdk/logger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/nilLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/stdoutLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk"
)

var (
	url               string
	originator        string
	passwordRaw       string
	officeId          string
	stdOutLog, nilLog logger.LogWriter
)

func tearUp() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url = os.Getenv("URL")
	originator = os.Getenv("ORIGINATOR")
	passwordRaw = os.Getenv("PASSWORD_RAW")
	officeId = os.Getenv("OFFICE_ID")
	nilLog = nilLogger.Init()
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
		client := sdk.CreateAmadeusClient(url, originator, passwordRaw, officeId, nilLog)

		amadeusSDK := NewSKD(client, GetLatestMethodsMap())

		_, err := amadeusSDK.CommandCryptic("AN20MAYMOWLED/ALH")
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
	t.Run("methods versions test", func(t *testing.T) {
		client := sdk.CreateAmadeusClient(url, originator, passwordRaw, officeId, nilLog)

		var CommandCrypticV071 = 123
		mm := GetLatestMethodsMap()
		mm.CommandCryptic = CommandCrypticV071
		amadeusSDK := NewSKD(client, mm)

		_, err := amadeusSDK.CommandCryptic("AN20MAYMOWLED/ALH")
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}
