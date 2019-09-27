package configuration

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/microparts/configuration-golang"
	"github.com/microparts/errors-go"
	"github.com/microparts/logs-go"
	"gopkg.in/yaml.v2"
)

type connectionConfig struct {
	URL         string `yaml:"url"`
	WSAP        string `yaml:"wsap"`
	Originator  string `yaml:"originator"`
	Password    string `yaml:"password"`
	PasswordRaw string
	PinCode     string `yaml:"pin_code"`
	VatURL      string `yaml:"vat_url"`
	VatSOAPURL  string `yaml:"vat_soap_url"`
	VatOfficeID string `yaml:"vat_office_id"`
	VatSign     string `yaml:"vat_sign"`
}

type amadeusConfig struct {
	Connection         connectionConfig `yaml:"connection"`
	MaxRecommendations int              `yaml:"max_recommendations"`
}

type httpConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type formatsConfig struct {
	Time    string `yaml:"time"`
	Date    string `yaml:"date"`
	XMLDate string `yaml:"xml_date"`
}

// CfgType service config structure
type CfgType struct {
	Amadeus       amadeusConfig       `yaml:"amadeus"`
	HTTP          httpConfig          `yaml:"http"`
	Formats       formatsConfig       `yaml:"formats"`
	Queue         queueConfig         `yaml:"queue"`
	Log           *logs.Config        `yaml:"log"`
	Settings      SettingsConfig      `yaml:"settings"`
	Notifications NotificationsConfig `yaml:"notifications"`
}

type queueConfig struct {
	PubSub PubSubConfig `yaml:"pubsub"`
}

// PubSubConfig  PubSub Config
type PubSubConfig struct {
	Host         string `yaml:"host"`
	ProjectID    string `yaml:"project_id"`
	Topic        string `yaml:"topic"`
	Subscription string `yaml:"subscription"`
	Key          string `yaml:"key"`
}

// SettingsConfig Settings Config
type SettingsConfig struct {
	FoidRequreAirlines        []string `yaml:"foid_require_airlines"`
	RemoveDuplicateAirlines   bool     `yaml:"remove_duplicate_airlines"`
	MaxAttemptsCancelVoid     int      `yaml:"max_attempts_cancel_void"`
	MaxAttemptsDocIssuance    int      `yaml:"max_attempts_doc_issuance"`
	MaxAttemptsPNRRET         int      `yaml:"max_attempts_pnr_ret"`
	MaxAttemptsPNRADD         int      `yaml:"max_attempts_pnr_add"`
	IssueExpire               int      `yaml:"issue_expire"`
	BookingRequestsDelay      int      `yaml:"booking_requests_delay"`
	IssueRequestsDelay        int      `yaml:"issue_requests_delay"`
	FareRulesParagraphsToShow []string `yaml:"fare_rules_paragraphs_to_show"`
	FareQualifierList         []string `yaml:"fare_qualifier_list"`
}

type NotificationsConfig struct {
	ErrorMessageEmail string       `yaml:"error_message_email"`
	MailFrom          string       `yaml:"mail_from"`
	Queue             PubSubConfig `yaml:"pubsub"`
}

var (
	// Provider is current service
	Provider = "amadeus"
)

var Config *CfgType

// InitConfig initialize config data
func InitConfig() error {
	configPath := config.GetEnv("CONFIG_PATH", "")
	configBytes, err := config.ReadConfigs(configPath)
	if errors.HasErrors(err) {
		log.Printf("[config] read error: %+v", err)
		return err
	}

	err = yaml.Unmarshal(configBytes, &Config)
	if errors.HasErrors(err) {
		log.Printf("[config] unmarshal error for bytes: %+v", configBytes)
		return err
	}
	//
	// structs.TimeFormat = func() string {
	// 	return Config.Formats.Time
	// }
	//
	// structs.DateFormat = func() string {
	// 	return Config.Formats.Date
	// }

	if pwdBytes, err := base64.StdEncoding.DecodeString(Config.Amadeus.Connection.Password); err == nil {
		Config.Amadeus.Connection.PasswordRaw = string(pwdBytes)
	} else {
		log.Println("Error decoding password. Is it not encrypted in bass64?")
		os.Exit(-1)
	}

	return nil
}
