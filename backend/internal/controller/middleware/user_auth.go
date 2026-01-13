package middleware

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
	"github.com/yockii/inkruins/internal/cache"
	"github.com/yockii/inkruins/internal/constant"
	"github.com/yockii/inkruins/internal/domain"
)

func UserAuth(ctx fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		token = ctx.Get("token")
	}
	if token == "" {
		token = ctx.Query("token")
	}
	if token == "" {
		token = ctx.Cookies("token")
	}
	if token != "" {
		token = strings.TrimPrefix(token, "Bearer ")

		uid, err := cache.GetRedisClient().Digest(ctx, constant.CacheKeyUserToken+token).Result()
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				slog.Error("user auth failed", "err", err)
			}
			return ctx.Status(fiber.StatusUnauthorized).JSON(domain.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized"))
		}
		ctx.Locals(constant.LocalUserID, uid)
		return ctx.Next()
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(domain.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized"))
}
