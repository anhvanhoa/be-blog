package blog_service

import (
	"be-blog/src/config"
	"be-blog/src/libs/errors"
	"be-blog/src/models"

	"github.com/go-pg/pg/v10"
)

func GetBlogs() ([]models.Blog, error) {
	var blogs []models.Blog = []models.Blog{}
	query := `
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
JOIN categories c ON b.category_id = c.id;`

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
    JOIN categories c ON b.category_id = c.id WHERE b.slug = ?;`
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
            b.updated_at,
            (
                SELECT ARRAY_AGG(json_build_object('id', t.id, 'slug', t.slug, 'name', t.name, 'variables', t.variables))
                FROM UNNEST(b.tags) AS tag_id
                JOIN tags t ON t.id = tag_id
                WHERE t.status = TRUE
            ) AS tags
        FROM blogs b
        JOIN categories c ON b.category_id = c.id
        WHERE b.id != ?;`
	_, err = config.DB.Query(&blog.BlogsRelated, queryRelated, blog.ID)
	if err != nil {
		return blog, err
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
