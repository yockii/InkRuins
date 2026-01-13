package controller

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/yockii/inkruins/internal/constant"
	"github.com/yockii/inkruins/internal/controller/middleware"
	"github.com/yockii/inkruins/internal/domain"
	"github.com/yockii/inkruins/internal/model"
	"github.com/yockii/inkruins/internal/service"
)

type userController struct{}

func init() {
	controllers = append(controllers, &userController{})
}

func (c *userController) InitializeRouter(router fiber.Router) {
	router.Post("/register", c.Register)
	router.Post("/login", c.Login)

	r := router.Group("/user", middleware.UserAuth)
	r.Get("/my-info", c.GetMyInfo)
}

func (*userController) Register(ctx fiber.Ctx) error {
	var req domain.RegisterUserReq
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "invalid request body"))
	}
	if err := service.UserService.CreateUser(&model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		slog.Error("create user failed", "err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, "create user failed"))
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.NewSuccessResponse(nil))
}

func (c *userController) Login(ctx fiber.Ctx) error {
	var req domain.LoginReq
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "invalid request body"))
	}
	token, user, err := service.UserService.Login(req.Username, req.Password)
	if err != nil {
		slog.Error("login failed", "err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, "login failed"))
	}
	user.Password = ""
	return ctx.Status(fiber.StatusOK).JSON(domain.NewSuccessResponse(domain.LoginResp{
		Token: token,
		User:  user,
	}))
}

func (c *userController) GetMyInfo(ctx fiber.Ctx) error {
	userID := fiber.Locals(ctx, constant.LocalUserID, uint64(0))
	if userID == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(domain.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized"))
	}
	user, err := service.UserService.GetUserByID(userID)
	if err != nil {
		slog.Error("get user failed", "err", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, "get user failed"))
	}
	user.Password = ""
	return ctx.Status(fiber.StatusOK).JSON(domain.NewSuccessResponse(user))
}
