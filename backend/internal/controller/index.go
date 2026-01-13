package controller

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/yockii/inkruins/internal/cache"
	"github.com/yockii/inkruins/internal/constant"
)

type Controller interface {
	InitializeRouter(router fiber.Router)
}

var controllers []Controller

func InitializeRouter(app *fiber.App) {
	router := app.Group("/api/v1")
	// 初始化所有需要认证的controller
	for _, c := range controllers {
		c.InitializeRouter(router)
	}
}

// GetUserIDFromContext 从请求上下文中获取用户ID
// 从Authorization header中获取token，然后从Redis中获取userID
func GetUserIDFromContext(c fiber.Ctx) (uint64, error) {
	// 从Authorization header获取token
	token := c.Get("Authorization")
	if token == "" {
		// 尝试从query参数获取
		token = c.Query("token")
	}
	if token == "" {
		// 尝试从cookie获取
		token = c.Cookies("token")
	}

	if token == "" {
		return 0, fmt.Errorf("未提供认证token")
	}

	// 从Redis获取userID
	userIDStr, err := cache.Get(constant.CacheKeyUserToken + token)
	if err != nil {
		slog.Error("获取用户信息失败", "err", err, "token", token)
		return 0, fmt.Errorf("无效的token")
	}

	// 转换为uint64
	var userID uint64
	_, err = fmt.Sscanf(userIDStr, "%d", &userID)
	if err != nil {
		return 0, fmt.Errorf("无效的用户ID格式")
	}

	return userID, nil
}
