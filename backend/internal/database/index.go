package database

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/yockii/yunjin/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "yunjin")
	viper.SetDefault("database.name", "yunjin")
	viper.SetDefault("database.password", "yunjin")
}

func Initialize() {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
				viper.GetString("database.host"),
				viper.GetInt("database.port"),
				viper.GetString("database.user"),
				viper.GetString("database.name"),
				viper.GetString("database.password"),
			),
		),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		panic(err)
	}
	DB = db
}

func Close() {
	if DB != nil {
		db, err := DB.DB()
		if err == nil {
			err = db.Close()
			if err != nil {
				panic(err)
			}
		}
	}
}

func AutoMigrate() {
	if DB != nil {
		err := DB.AutoMigrate(
			model.Models...,
		)
		if err != nil {
			panic(err)
		}
	}
}
