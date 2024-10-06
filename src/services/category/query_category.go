package category_service

import (
	"be-blog/src/config"
	"be-blog/src/models"
)

func GetCategories(status bool) ([]models.Category, error) {
	var categories []models.Category = []models.Category{}
	err := config.DB.Model(&categories).Where("status = ?", status).Select()
	return categories, err
}

func GetCategoryById(id int) (models.Category, error) {
	var category models.Category
	err := config.DB.Model(&category).Where("id = ?", id).Select()
	return category, err
}
