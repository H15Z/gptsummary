package configs

import (
	"log"

	"github.com/spf13/viper"
)

var CfgFile string

// initConfig reads in config file and ENV variables if set
// this is used to load configs for general settings such as db creds etc..
func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
