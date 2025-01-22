package config

import "github.com/spf13/viper"

func InitViper() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
