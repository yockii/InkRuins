package service

import "github.com/yockii/inkruins/internal/model"

type IProjectService interface {
	CreateProject(project *model.Project) error
	GetProjectByID(id uint64) (*model.Project, error)
	GetProjectList(condition *model.Project, page, pageSize int) ([]*model.Project, int64, error)
	GetProjectCount(condition *model.Project) (int64, error)
	UpdateProject(project *model.Project) error
	DeleteProject(id uint64) error
}
