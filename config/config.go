package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Warningf("%v", err)
	}
	log.Info("Using config file: ", viper.ConfigFileUsed())
}

func Env() string {
	return viper.GetString("env")
}

func Port() string {
	if !viper.IsSet("ports") {
		return "8080"
	}
	return viper.GetString("ports")
}

func DBHost() string {
	return viper.GetString("database.host")
}

func DBDatabase() string {
	return viper.GetString("database.database")
}

func DBUser() string {
	return viper.GetString("database.username")
}

func DBPassword() string {
	return viper.GetString("database.password")
}

func DBDSN() string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", DBUser(), DBPassword(), DBHost(), DBDatabase())
}
