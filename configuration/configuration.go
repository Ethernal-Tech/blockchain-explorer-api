package configuration

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Configuration struct {
	DbUser               string
	DbPassword           string
	DbHost               string
	DbPort               string
	DbName               string
	CallTimeoutInSeconds uint
}

// Load returns an initialized Configuration instance.
func Load() *Configuration {
	configFile, err := filepath.Abs(".env")

	if err != nil {
		os.Exit(1)
	}

	err = read(configFile)

	if err != nil {
		os.Exit(1)
	}

	configuration := Configuration{
		DbUser:               viper.GetString("DB_USER"),
		DbPassword:           viper.GetString("DB_PASSWORD"),
		DbHost:               viper.GetString("DB_HOST"),
		DbPort:               viper.GetString("DB_PORT"),
		DbName:               viper.GetString("DB_NAME"),
		CallTimeoutInSeconds: viper.GetUint("CALL_TIMEOUT_IN_SECONDS"),
	}

	return &configuration
}

// Read reads the contents of the file during application startup.
func read(file string) error {
	viper.SetConfigFile(file)
	return viper.ReadInConfig()
}
