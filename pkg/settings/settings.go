package settings

import (
	"log"

	"github.com/spf13/viper"
)

type Settings struct {
	StoragePath  string `mapstructure:"STORGAE_PATH"`
	ResponsePath string `mapstructure:"RESPONSE_PATH"`
}

func NewSettings() *Settings {
	var settings Settings
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Println("No env file found", err)
	}
	err = viper.Unmarshal(&settings)

	if err != nil {
		log.Println("Error: while trying to unmarshal configuration", err)
	}
	return &settings
}
