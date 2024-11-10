package tag_service

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/models"
	"time"

	"github.com/google/uuid"
)

func GetTags() ([]entities.Tag, error) {
	query := `
	SELECT t.*, COUNT(b.id) AS count
	FROM tags t
	LEFT JOIN blogs b ON t.id = ANY(b.tags)
	GROUP BY t.id
	ORDER BY t.id;
	`
	var tags []entities.Tag
	_, err := config.DB.Query(&tags, query)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func CreateTag(tag models.TagReq) error {
	status := true
	newTag := entities.TagCreate{
		ID:        uuid.New().String(),
		Status:    &status,
		Name:      tag.Name,
		Slug:      tag.Slug,
		Variables: tag.Variables,
	}
	_, err := config.DB.Model(&newTag).Insert()
	if err != nil {
		return err
	}
	return nil
}

func GetTag(id string) (*entities.Tag, error) {
	tag := &entities.Tag{ID: id}
	err := config.DB.Model(tag).WherePK().Group("id").Select()
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func UpdateTag(id string, tag models.TagReq) error {
	newTag := entities.TagUpdate{
		ID:        tag.ID,
		Name:      tag.Name,
		Slug:      tag.Slug,
		Variables: tag.Variables,
		UpdatedAt: time.Now(),
	}
	_, err := config.DB.Model(&newTag).Where("id = ?", id).Update("name", "slug", "variables", "updated_at")
	if err != nil {
		return err
	}
	return nil
}

func DeleteTag(id string) error {
	query := `
	SELECT t.*, COUNT(b.id) AS count
	FROM tags t
	LEFT JOIN blogs b ON t.id = ANY(b.tags)
	WHERE t.id = ?
	GROUP BY t.id
	ORDER BY t.id;
	`
	var tags entities.Tag
	_, err := config.DB.Query(&tags, query, id)
	if err != nil {
		return err
	}
	if tags.Count > 0 {
		err := errors.NewErrorBadRequest("Tag này đang được sử dụng")
		return err
	}
	tag := &entities.Tag{ID: id}
	_, err = config.DB.Model(tag).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}
