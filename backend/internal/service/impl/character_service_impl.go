package impl

import (
	"fmt"

	"github.com/yockii/inkruins/internal/dao"
	"github.com/yockii/inkruins/internal/database"
	"github.com/yockii/inkruins/internal/model"
)

type CharacterServiceImpl struct {
}

func (s *CharacterServiceImpl) CreateCharacter(character *model.Character) error {
	if character == nil {
		return fmt.Errorf("character is nil")
	}
	if character.ProjectID == 0 {
		return fmt.Errorf("project id is required")
	}
	if character.Name == "" {
		return fmt.Errorf("character name is required")
	}
	if err := database.DB.Create(character).Error; err != nil {
		return err
	}
	return nil
}

func (s *CharacterServiceImpl) GetCharacterByID(id uint64) (*model.Character, error) {
	if id == 0 {
		return nil, fmt.Errorf("character id is required")
	}
	var character model.Character
	if err := database.DB.First(&character, id).Error; err != nil {
		return nil, err
	}
	return &character, nil
}

func (s *CharacterServiceImpl) GetCharacterList(projectID uint64, page, pageSize int) ([]*model.Character, int64, error) {
	if projectID == 0 {
		return nil, 0, fmt.Errorf("project id is required")
	}
	var total int64
	tx := database.DB.Where(dao.Character.ProjectID.Eq(projectID))
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}

	var characters []*model.Character
	if err := tx.Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&characters).Error; err != nil {
		return nil, 0, err
	}
	return characters, total, nil
}

func (s *CharacterServiceImpl) UpdateCharacter(character *model.Character) error {
	if character == nil {
		return fmt.Errorf("character is nil")
	}
	if character.ID == 0 {
		return fmt.Errorf("character id is required")
	}
	if err := database.DB.Save(character).Error; err != nil {
		return err
	}
	return nil
}

func (s *CharacterServiceImpl) DeleteCharacter(id uint64) error {
	if id == 0 {
		return fmt.Errorf("character id is required")
	}
	if err := database.DB.Delete(&model.Character{}, id).Error; err != nil {
		return err
	}
	return nil
}
