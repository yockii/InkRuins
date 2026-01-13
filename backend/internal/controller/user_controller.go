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

// Register 注册用户
// @Summary 注册用户
// @Description 注册用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param req body domain.RegisterUserReq true "注册用户请求"
// @Success 200 {object} domain.Response{data=nil}
// @Failure 400 {object} domain.Response{data=nil}
// @Failure 500 {object} domain.Response{data=nil}
// @Router /api/v1/register [post]
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

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param req body domain.LoginReq true "用户登录请求"
// @Success 200 {object} domain.Response{data=domain.LoginResp}
// @Failure 400 {object} domain.Response{data=nil}
// @Failure 500 {object} domain.Response{data=nil}
// @Router /api/v1/login [post]
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

// GetMyInfo 获取当前登录用户信息
// @Summary 获取当前登录用用户信息
// @Description 获取当前登录用户的信息
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response{data=model.User}
// @Failure 400 {object} domain.Response{data=nil}
// @Failure 500 {object} domain.Response{data=nil}
// @Router /api/v1/user/my-info [get]
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
