package impl

import (
	"errors"

	"github.com/yockii/inkruins/internal/dao"
	"github.com/yockii/inkruins/internal/database"
	"github.com/yockii/inkruins/internal/model"
)

type StoryEventServiceImpl struct {
}

func NewStoryEventServiceImpl() *StoryEventServiceImpl {
	return &StoryEventServiceImpl{}
}

func (s *StoryEventServiceImpl) CreateStoryEvent(storyEvent *model.StoryEvent) error {
	if storyEvent == nil {
		return errors.New("storyEvent is required")
	}
	if storyEvent.ProjectID == 0 {
		return errors.New("projectID is required")
	}
	if storyEvent.Title == "" {
		return errors.New("title is required")
	}
	if err := database.DB.Create(storyEvent).Error; err != nil {
		return err
	}
	return nil
}

func (s *StoryEventServiceImpl) GetStoryEventByID(id uint64) (*model.StoryEvent, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}
	var storyEvent model.StoryEvent
	if err := database.DB.Where(dao.BaseModel.ID.Eq(id)).First(&storyEvent).Error; err != nil {
		return nil, err
	}
	return &storyEvent, nil
}

func (s *StoryEventServiceImpl) GetStoryEventsByProjectID(projectID uint64, page, pageSize int) ([]*model.StoryEvent, int64, error) {
	if projectID == 0 {
		return nil, 0, errors.New("projectID is required")
	}
	var total int64
	tx := database.DB.Model(&model.StoryEvent{}).Where(dao.StoryEvent.ProjectID.Eq(projectID))

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return nil, total, nil
	}

	var storyEvents []*model.StoryEvent
	if err := tx.Order(dao.StoryEvent.Seq.Asc()).Offset((page - 1) * pageSize).Limit(pageSize).Find(&storyEvents).Error; err != nil {
		return nil, 0, err
	}
	return storyEvents, total, nil
}

func (s *StoryEventServiceImpl) UpdateStoryEvent(storyEvent *model.StoryEvent) error {
	if storyEvent == nil {
		return errors.New("storyEvent is required")
	}
	if storyEvent.ID == 0 {
		return errors.New("id is required")
	}
	if err := database.DB.Save(storyEvent).Error; err != nil {
		return err
	}
	return nil
}

func (s *StoryEventServiceImpl) DeleteStoryEvent(id uint64) error {
	if id == 0 {
		return errors.New("id is required")
	}
	return database.DB.Delete(&model.StoryEvent{}, id).Error
}
