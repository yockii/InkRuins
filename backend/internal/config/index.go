package config

import "github.com/spf13/viper"

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		viper.SafeWriteConfig()
	}
}
