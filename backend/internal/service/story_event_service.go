package service

import "github.com/yockii/inkruins/internal/model"

type StoryEventService interface {
	CreateStoryEvent(storyEvent *model.StoryEvent) error
	GetStoryEventByID(id uint64) (*model.StoryEvent, error)
	GetStoryEventByProjectID(projectID uint64, page, pageSize int) ([]*model.StoryEvent, int64, error)
	UpdateStoryEvent(storyEvent *model.StoryEvent) error
	DeleteStoryEvent(id uint64) error
}
