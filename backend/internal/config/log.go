package config

import (
	"log/slog"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	viper.SetDefault("log.filename", "yunjin.log")
	viper.SetDefault("log.maxSizeInMb", 10)
	viper.SetDefault("log.maxAgeInDays", 30)
	viper.SetDefault("log.maxBackups", 10)
	viper.SetDefault("log.compress", false)
}

func InitLogger() {
	r := &lumberjack.Logger{
		Filename:   viper.GetString("log.filename"),
		LocalTime:  true,
		MaxSize:    viper.GetInt("log.maxSizeInMb"),
		MaxAge:     viper.GetInt("log.maxAgeInDays"),
		MaxBackups: viper.GetInt("log.maxBackups"),
		Compress:   viper.GetBool("log.compress"),
	}
	logger := slog.New(slog.NewJSONHandler(r, nil))
	slog.SetDefault(logger)
}
