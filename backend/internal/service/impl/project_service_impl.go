package impl

import (
	"errors"

	"github.com/yockii/yunjin/internal/dao"
	"github.com/yockii/yunjin/internal/database"
	"github.com/yockii/yunjin/internal/model"
	"gorm.io/gorm"
)

type ProjectServiceImpl struct{}

func NewProjectService() *ProjectServiceImpl {
	return &ProjectServiceImpl{}
}

func (s *ProjectServiceImpl) CreateProject(project *model.Project) error {
	if project.Title == "" {
		return errors.New("title is required")
	}
	if project.UserID == 0 {
		return errors.New("user_id is required")
	}
	return database.DB.Create(project).Error
}

func (s *ProjectServiceImpl) GetProjectByID(id uint64) (*model.Project, error) {
	var project model.Project
	if err := database.DB.Where(dao.BaseModel.ID.Eq(id)).First(&project).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}

func (s *ProjectServiceImpl) GetProjectList(condition *model.Project, page, pageSize int) ([]*model.Project, int64, error) {
	var projects []*model.Project
	var total int64
	tx := database.DB.Model(&model.Project{}).Order(dao.BaseModel.CreatedAt.Desc())
	if condition != nil {
		if condition.Title != "" {
			tx = tx.Where(dao.Project.Title.Like("%" + condition.Title + "%"))
		}
	}
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return projects, 0, nil
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if err := tx.Offset((page - 1) * pageSize).Limit(pageSize).Find(&projects).Error; err != nil {
		return nil, 0, err
	}
	return projects, total, nil
}

func (s *ProjectServiceImpl) GetProjectCount(condition *model.Project) (int64, error) {
	var total int64
	tx := database.DB.Model(&model.Project{})
	if condition != nil {
		if condition.Title != "" {
			tx = tx.Where(dao.Project.Title.Like("%" + condition.Title + "%"))
		}
	}
	return total, nil
}

func (s *ProjectServiceImpl) UpdateProject(project *model.Project) error {
	return database.DB.Save(project).Error
}

func (s *ProjectServiceImpl) DeleteProject(id uint64) error {
	return database.DB.Delete(&model.Project{}, id).Error
}
