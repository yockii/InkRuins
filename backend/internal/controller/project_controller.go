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

type projectController struct{}

func init() {
	controllers = append(controllers, &projectController{})
}

// InitializeRouter 初始化项目路由
func (c *projectController) InitializeRouter(router fiber.Router) {
	r := router.Group("/projects", middleware.UserAuth)
	r.Get("/list", c.GetMyProjects)
	r.Post("/create", c.CreateProject)
	r.Get("/instance/:id", c.GetProject)
	r.Put("/update/:id", c.UpdateProject)
	r.Delete("/delete/:id", c.DeleteProject)
}

// GetMyProjects 获取我的项目列表
// @Summary 获取我的项目列表
// @Description 获取我的项目列表
// @Tags 项目
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param title query string false "项目标题"
// @Success 200 {object} domain.PaginateResponse{data=[]model.Project}
// @Failure 400 {object} domain.Response{data=nil}
// @Failure 500 {object} domain.Response{data=nil}
// @Router /api/v1/projects/list [get]
func (c *projectController) GetMyProjects(ctx fiber.Ctx) error {
	userID := fiber.Locals(ctx, constant.LocalUserID, uint64(0))
	if userID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "user id is required"))
	}

	var req domain.ListProjectReq
	if err := ctx.Bind().Query(&req); err != nil {
		slog.Error("bind query error", "err", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	projects, count, err := service.ProjectService.GetProjectList(&model.Project{UserID: userID}, req.Page, req.PageSize)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.NewSuccessPaginateResponse(count, req.Page, req.PageSize, projects))
}

// CreateProject 创建项目
// @Summary 创建项目
// @Description 创建项目
// @Tags 项目
// @Accept json
// @Produce json
// @Param project body domain.CreateProjectReq true "项目信息"
// @Success 200 {object} domain.Response{data=model.Project}
// @Failure 400 {object} domain.Response{data=nil}
// @Failure 500 {object} domain.Response{data=nil}
// @Router /api/v1/projects/create [post]
func (c *projectController) CreateProject(ctx fiber.Ctx) error {
	userID := fiber.Locals(ctx, constant.LocalUserID, uint64(0))
	if userID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "user id is required"))
	}
	var req domain.CreateProjectReq
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}
	project := &model.Project{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Genre:       req.Genre,
	}
	if err := service.ProjectService.CreateProject(project); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.NewSuccessResponse(project))
}

// GetProject 获取项目详情
// @Summary 获取项目详情
// @Description 获取项目详情
// @Tags 项目
// @Accept json
// @Produce json
// @Param id path uint64 true "项目ID"
// @Success 200 {object} domain.Response{data=model.Project}
// @Failure 400 {object} domain.Response{data=nil}
// @Failure 500 {object} domain.Response{data=nil}
// @Router /api/v1/projects/instance/{id} [get]
func (c *projectController) GetProject(ctx fiber.Ctx) error {
	userID := fiber.Locals(ctx, constant.LocalUserID, uint64(0))
	if userID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "user id is required"))
	}
	projectID := fiber.Params(ctx, "id", uint64(0))
	if projectID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "project id is required"))
	}
	project, err := service.ProjectService.GetProjectByID(projectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	if project.UserID != userID {
		return ctx.Status(fiber.StatusForbidden).JSON(domain.NewErrorResponse(fiber.StatusForbidden, "you are not the owner of this project"))
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.NewSuccessResponse(project))
}

// UpdateProject 更新项目
// @Summary 更新项目
// @Description 更新项目
// @Tags 项目
// @Accept json
// @Produce json
// @Param id path uint64 true "项目ID"
// @Param project body model.Project true "项目信息"
// @Success 200 {object} domain.Response{data=nil}
// @Failure 400 {object} domain.Response{data=nil}
// @Failure 500 {object} domain.Response{data=nil}
// @Router /api/v1/projects/update/{id} [put]
func (c *projectController) UpdateProject(ctx fiber.Ctx) error {
	userID := fiber.Locals(ctx, constant.LocalUserID, uint64(0))
	if userID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "user id is required"))
	}
	projectID := fiber.Params(ctx, "id", uint64(0))
	if projectID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "project id is required"))
	}

	var req model.Project
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	project, err := service.ProjectService.GetProjectByID(projectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	if project.UserID != userID {
		return ctx.Status(fiber.StatusForbidden).JSON(domain.NewErrorResponse(fiber.StatusForbidden, "you are not the owner of this project"))
	}
	if err := service.ProjectService.UpdateProject(&req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.NewSuccessResponse(nil))
}

// DeleteProject 删除项目
// @Summary 删除项目
// @Description 删除项目
// @Tags 项目
// @Accept json
// @Produce json
// @Param id path uint64 true "项目ID"
// @Success 200 {object} domain.Response{data=nil}
// @Failure 400 {object} domain.Response{data=nil}
// @Failure 500 {object} domain.Response{data=nil}
// @Router /api/v1/projects/delete/{id} [delete]
func (c *projectController) DeleteProject(ctx fiber.Ctx) error {
	userID := fiber.Locals(ctx, constant.LocalUserID, uint64(0))
	if userID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "user id is required"))
	}
	projectID := fiber.Params(ctx, "id", uint64(0))
	if projectID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.NewErrorResponse(fiber.StatusBadRequest, "project id is required"))
	}
	project, err := service.ProjectService.GetProjectByID(projectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	if project.UserID != userID {
		return ctx.Status(fiber.StatusForbidden).JSON(domain.NewErrorResponse(fiber.StatusForbidden, "you are not the owner of this project"))
	}
	if err := service.ProjectService.DeleteProject(projectID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(domain.NewSuccessResponse(nil))
}
