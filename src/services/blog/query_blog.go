package blog_service

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/models"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

func GetBlogs(status bool) ([]models.Blog, error) {
	var blogs []models.Blog = []models.Blog{}
	query := `
SELECT
    b.id,
    b.title,
    b.description,
    b.slug,
	b.thumbnail,
	c.slug as slug_category,
	c.name as category,
    b.updated_at,
	b.status,
	(
        SELECT ARRAY_AGG(json_build_object('id', t.id, 'slug', t.slug, 'name', t.name, 'variables', t.variables))
        FROM UNNEST(b.tags) AS tag_id
        JOIN tags t ON t.id = tag_id
        WHERE t.status = TRUE
    ) AS tags
FROM blogs b
JOIN categories c ON b.category_id = c.id WHERE b.status = ?;`

	_, err := config.DB.Query(&blogs, query, status)
	return blogs, err
}

func GetManagerBlogs() ([]models.Blog, error) {
	var blogs []models.Blog = []models.Blog{}
	query := `
SELECT
    b.id,
    b.title,
    b.description,
    b.slug,
	b.thumbnail,
	c.slug as slug_category,
	c.name as category,
    b.updated_at,
	b.status,
	(
        SELECT ARRAY_AGG(json_build_object('id', t.id, 'slug', t.slug, 'name', t.name, 'variables', t.variables))
        FROM UNNEST(b.tags) AS tag_id
        JOIN tags t ON t.id = tag_id
        WHERE t.status = TRUE
    ) AS tags
FROM blogs b
JOIN categories c ON b.category_id = c.id`

	_, err := config.DB.Query(&blogs, query)
	return blogs, err
}

func GetBlogBySlug(slug string) (models.BlogBySlug, error) {
	var blog models.BlogBySlug
	query := `
    SELECT
        b.id,
        b.title,
        b.description,
        b.slug,
        b.content_md,
        b.content_html,
        b.security,
        b.thumbnail,
        c.slug as slug_category,
        b.updated_at,
        (
            SELECT ARRAY_AGG(json_build_object('id', t.id, 'slug', t.slug, 'name', t.name, 'variables', t.variables))
            FROM UNNEST(b.tags) AS tag_id
            JOIN tags t ON t.id = tag_id
            WHERE t.status = TRUE
        ) AS tags
    FROM blogs b
    JOIN categories c ON b.category_id = c.id WHERE b.slug = ? AND b.status = TRUE;`
	_, err := config.DB.Query(&blog, query, slug)
	if err != nil {
		return blog, err
	}
	if blog.ID == "" {
		return blog, errors.NewErrorNotFound("Không tìm thấy bài viết")
	}
	queryRelated := `
        SELECT
            b.id,
            b.title,
            b.description,
            b.slug,
            b.thumbnail,
            c.slug as slug_category,
            c.name as category,
            b.updated_at,
            (
                SELECT ARRAY_AGG(json_build_object('id', t.id, 'slug', t.slug, 'name', t.name, 'variables', t.variables))
                FROM UNNEST(b.tags) AS tag_id
                JOIN tags t ON t.id = tag_id
                WHERE t.status = TRUE
            ) AS tags
        FROM blogs b
        JOIN categories c ON b.category_id = c.id
        WHERE b.id != ? AND b.status = TRUE;`
	_, err = config.DB.Query(&blog.BlogsRelated, queryRelated, blog.ID)
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func GetBlogById(id string) (models.BlogByID, error) {
	var blog models.BlogByID
	query := `
	SELECT
		b.id,
		b.title,
		b.description,
		b.slug,
		b.content_md,
		b.content_html,
		b.security,
		b.thumbnail,
		b.category_id,
		c.name as category,
		c.slug as slug_category,
		b.updated_at,
		b.status,
		COALESCE(
		(
			SELECT ARRAY_AGG(json_build_object('id', t.id, 'slug', t.slug, 'name', t.name, 'variables', t.variables))
			FROM UNNEST(b.tags) AS tag_id
			JOIN tags t ON t.id = tag_id
			WHERE t.status = TRUE
		)
		, '{}') as tags
	FROM blogs b
	JOIN categories c ON b.category_id = c.id WHERE b.id = ?;`
	_, err := config.DB.Query(&blog, query, id)
	if err != nil {
		return blog, err
	}
	if blog.ID == "" {
		return blog, errors.NewErrorNotFound("Không tìm thấy bài viết")
	}
	return blog, nil
}

func GetBlogsByCategory(slug string) (models.BlogCategory, error) {
	var category models.BlogCategory = models.BlogCategory{}
	err := config.Transaction(func(tx *pg.Tx) error {
		queryCategory := `
            SELECT
                c.id,
                c.name,
                c.slug,
                c.thumbnail
            FROM categories c
            WHERE c.slug = ? AND c.status = TRUE;`
		_, err := tx.Query(&category, queryCategory, slug)

		if err != nil {
			return errors.NewError(err)
		}
		if category.ID == "" {
			return errors.NewErrorNotFound("Không tìm thấy danh mục")
		}

		queryBlogs := `
            SELECT
                b.id,
                b.title,
                b.description,
                b.slug,
                b.thumbnail,
                c.slug as slug_category,
                b.updated_at,
                (
                    SELECT ARRAY_AGG(json_build_object('id', t.id, 'slug', t.slug, 'name', t.name, 'variables', t.variables))
                    FROM UNNEST(b.tags) AS tag_id
                    JOIN tags t ON t.id = tag_id
                    WHERE t.status = TRUE
                ) AS tags
            FROM blogs b
            JOIN categories c ON b.category_id = c.id WHERE c.id = ?;`
		_, err = tx.Query(&category.Blogs, queryBlogs, category.ID)

		if err != nil {
			return errors.NewError(err)
		}
		return nil
	})
	return category, err
}

func CreateBlog(body models.BlogReq) error {
	blog := entities.Blog{
		ID:          uuid.New().String(),
		Title:       body.Title,
		ContentMd:   body.ContentMd,
		ContentHtml: body.ContentHtml,
		Description: body.Description,
		AuthorId:    body.AuthorId,
		CategoryId:  body.CategoryId,
		Slug:        body.Slug,
		Thumbnail:   body.Thumbnail,
		Status:      body.Status,
	}
	if body.Tags != nil {
		for _, tag := range body.Tags {
			blog.Tags = append(blog.Tags, tag.ID)
		}
	}
	isSlugExist, err := config.DB.Model(&blog).Where("slug = ?", blog.Slug).Exists()
	if err != nil {
		return err
	}
	if isSlugExist {
		return errors.NewErrorBadRequest("Slug đã tồn tại")
	}

	_, err = config.DB.Model(&blog).Insert()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func UpdateBlog(id string, body models.BlogReq) error {
	blog := entities.Blog{
		ID:          id,
		Title:       body.Title,
		ContentMd:   body.ContentMd,
		ContentHtml: body.ContentHtml,
		Description: body.Description,
		AuthorId:    body.AuthorId,
		CategoryId:  body.CategoryId,
		Slug:        body.Slug,
		Thumbnail:   body.Thumbnail,
		Status:      body.Status,
		UpdatedAt:   time.Now(),
	}
	if body.Tags != nil {
		for _, tag := range body.Tags {
			blog.Tags = append(blog.Tags, tag.ID)
		}
	}
	isSlugExist, err := config.DB.Model(&blog).Where("slug = ?", blog.Slug).Where("id != ?", id).Exists()
	if err != nil {
		return err
	}
	if isSlugExist {
		return errors.NewErrorBadRequest("Slug đã tồn tại")
	}
	_, err = config.DB.Model(&blog).Where("id = ?", id).Update()
	if err != nil {
		return err
	}
	return nil
}

func DeleteBlog(id string) error {
	blog := entities.Blog{
		ID: id,
	}
	_, err := config.DB.Model(&blog).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}
