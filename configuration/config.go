package configuration

var Config ConfigType

type ConfigType struct {
	Formats  formatsConfig
	Provider string
}

type formatsConfig struct {
	Time    string
	Date    string
	XMLDate string
}
