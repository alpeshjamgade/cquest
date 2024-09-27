package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	// *core
	HttpPort = "8080"

	//*database
	DbHost     = "localhost"
	DbPort     = "3306"
	DbUsername = "root"
	DbPassword = "root"
	DbSSLMode  = "disable"
	DbName     = "cquest"

	// *logger
	LogLevel    = "info"
	LogEncoding = "console"
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

	// database
	DbHost = viper.GetString("database.host")
	DbPort = viper.GetString("database.port")
	DbUsername = viper.GetString("database.username")
	DbPassword = viper.GetString("database.password")
	DbSSLMode = viper.GetString("database.ssl_mode")
	DbName = viper.GetString("database.name")

	// logger
	LogLevel = viper.GetString("logger.level")
	LogEncoding = viper.GetString("logger.encoding")

	return nil
}
