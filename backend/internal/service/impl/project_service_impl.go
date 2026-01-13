package impl

import (
	"errors"

	"github.com/yockii/inkruins/internal/dao"
	"github.com/yockii/inkruins/internal/database"
	"github.com/yockii/inkruins/internal/model"
	"gorm.io/gorm"
)

type ProjectServiceImpl struct{}

func NewProjectService() *ProjectServiceImpl {
	return &ProjectServiceImpl{}
}

// CreateProject 创建项目时自动创建一个初始故事节点
func (s *ProjectServiceImpl) CreateProject(project *model.Project) error {
	if project == nil {
		return errors.New("project is required")
	}
	if project.Title == "" {
		return errors.New("title is required")
	}
	if project.UserID == 0 {
		return errors.New("user_id is required")
	}
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(project).Error; err != nil {
			return err
		}
		if err := tx.Create(&model.StoryEvent{
			ProjectID: project.ID,
			Seq:       0,
			Title:     "初始故事节点",
		}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
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
		if condition.UserID != 0 {
			tx = tx.Where(dao.Project.UserID.Eq(condition.UserID))
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
	if project == nil {
		return errors.New("project is required")
	}
	return database.DB.Save(project).Error
}

func (s *ProjectServiceImpl) DeleteProject(id uint64) error {
	if id == 0 {
		return errors.New("id is required")
	}
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.Project{}, id).Error; err != nil {
			return err
		}

		if err := tx.Where(dao.StoryEvent.ProjectID.Eq(id)).Delete(&model.StoryEvent{}).Error; err != nil {
			return err
		}

		if err := tx.Where(dao.Character.ProjectID.Eq(id)).Delete(&model.Character{}).Error; err != nil {
			return err
		}

		return nil
	})
}
