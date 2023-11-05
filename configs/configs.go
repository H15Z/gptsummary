package configs

import (
	"log"

	"github.com/spf13/viper"
)

var CfgFile string

func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName(".cmd")
	}

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
