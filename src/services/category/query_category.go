package category_service

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/models"
	"time"

	"github.com/google/uuid"
)

func GetCategories(status bool) ([]models.Category, error) {
	var categories []models.Category = []models.Category{}
	err := config.DB.Model(&categories).Where("status = ?", status).Select()
	return categories, err
}

func GetAll() ([]models.CategoryManager, error) {
	var categories []models.CategoryManager = []models.CategoryManager{}
	query := `SELECT *, (select count(b.id) from blogs b WHERE b.category_id = c.id) as count_blog FROM categories c;`
	_, err := config.DB.Query(&categories, query)
	return categories, err
}

func GetCategoryById(id string) (models.Category, error) {
	var category models.Category
	err := config.DB.Model(&category).Where("id = ?", id).Select()
	if err != nil {
		return category, err
	}
	return category, err
}

func CreateCategory(body models.CategoryReq) error {
	status := true
	category := entities.Category{
		ID:        uuid.New().String(),
		Name:      body.Name,
		Slug:      body.Slug,
		OrderC:    body.OrderC,
		Parent:    body.Parent,
		Thumbnail: body.Thumbnail,
		Status:    &status,
	}
	_, err := config.DB.Model(&category).Insert()
	return err
}

func UpdateCategory(id string, body models.CategoryReq) error {
	category := entities.CategoryUpdate{
		ID:        id,
		Name:      body.Name,
		Slug:      body.Slug,
		OrderC:    body.OrderC,
		Parent:    body.Parent,
		Thumbnail: body.Thumbnail,
		UpdatedAt: time.Now(),
	}
	_, err := config.DB.Model(&category).WherePK().Update()
	return err
}

func DeleteCategory(id string) error {
	category := entities.Category{
		ID: id,
	}
	_, err := config.DB.Model(&category).WherePK().Delete()
	return err
}
