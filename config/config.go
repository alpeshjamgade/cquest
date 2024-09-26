package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	// *core
	HttpPort = "8080"
)

func LoadConf() error {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cquest")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./etc/cquest")
	viper.AddConfigPath("config")
	viper.AddConfigPath("$HOME/.cquest")

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		return err
	}

	// core
	HttpPort = viper.GetString("core.http_port")

	return nil
}
