package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"github.com/yockii/inkruins/internal/cache"
	"github.com/yockii/inkruins/internal/config"
	"github.com/yockii/inkruins/internal/controller"
	"github.com/yockii/inkruins/internal/data"
	"github.com/yockii/inkruins/internal/database"
	"github.com/yockii/inkruins/pkg/util"
)

func init() {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.nodeId", 1)
}

func main() {
	// Initialize configuration
	config.InitLogger()
	util.InitSnowflake(viper.GetUint64("server.nodeId"))

	database.Initialize()
	defer database.Close()

	cache.InitializeRedis()
	defer cache.CloseRedis()

	database.AutoMigrate()

	if err := data.Initialize(); err != nil {
		fmt.Printf("初始化数据失败: %v\n", err)
	}

	app := fiber.New(fiber.Config{
		AppName:      "墨墟 InkRuins v1.0.0",
		ServerHeader: "InkRuins Server",
	})
	controller.InitializeRouter(app)
	app.Listen(fmt.Sprintf(":%d", viper.GetInt("server.port")), fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}
