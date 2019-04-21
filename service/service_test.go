package service

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/amadeus-golang-sdk/logger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/stdoutLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk"
)

var (
	url         string
	originator  string
	passwordRaw string
	officeId    string
	l           logger.Service
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

	ls := stdoutLogger.Init()
	l = logger.NewLogger(ls)

}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestNewSKD(t *testing.T) {
	t.Run("initiating test", func(t *testing.T) {
		client := sdk.CreateAmadeusClient(url, originator, passwordRaw, officeId, stdoutLogger.Init())

		amadeusSDK := NewSKD(client)

		response, _, err := amadeusSDK.CommandCryptic("AN20MAYMOWLED/ALH")
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		fmt.Printf("response: %v\n", response)

		fmt.Printf("error: %v\n", err)
	})
}
