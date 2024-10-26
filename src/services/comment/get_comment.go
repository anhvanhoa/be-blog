package comment

import (
	"be-blog/src/config"
	"be-blog/src/models"
)

func GetComments(slug string, status bool) ([]models.Comment, error) {
	var comments []models.Comment
	query := `
SELECT
    c.id,
    c.content,
    c.created_at,
    u.username,
    u.avatar,
    u.full_name,
	b.id as blog_id,
    COALESCE(
        (
            SELECT json_agg(
                json_build_object(
                    'id', r.id,
                    'content', r.content,
                    'createdAt', r.created_at,
                    'username', ru.username,
                    'avatar', ru.avatar,
                    'fullName', ru.full_name,
					'blogId', b.id 
                )
            )
            FROM comments r
            JOIN public.users ru ON r.author_id = ru.id
            WHERE r.parent_comment_id = c.id
				AND r.status = true
        ), '[]'::json
    ) AS replies
FROM comments c
JOIN public.blogs b ON c.blog_id = b.id
JOIN public.users u ON c.author_id = u.id
WHERE b.slug = ?
	AND c.status = ?
	AND c.parent_comment_id IS NULL
    ORDER BY c.created_at DESC;
	`
	_, err := config.DB.Query(&comments, query, slug, status)
	if err != nil {
		return comments, err
	}
	return comments, nil
}
