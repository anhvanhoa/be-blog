package statistic

import (
	"be-blog/src/config"
	"be-blog/src/models"

	"github.com/google/uuid"
)

func Statistic() ([]models.StatisticRes, error) {
	query := `
	  	SELECT
    (SELECT COUNT(*) FROM users) AS total_users,
    (SELECT COUNT(*) FROM users WHERE created_at BETWEEN date_trunc('month', NOW()) AND NOW()) AS value_users,
    (SELECT COUNT(*) FROM blogs) AS total_blogs,
    (SELECT COUNT(*) FROM blogs WHERE created_at BETWEEN date_trunc('month', NOW()) AND NOW()) AS value_blogs,
    (SELECT COUNT(*) FROM comments) AS total_comments,
    (SELECT COUNT(*) FROM comments WHERE created_at BETWEEN date_trunc('month', NOW()) AND NOW()) AS value_comments;
	`
	result := models.Statistic{}
	_, err := config.DB.Query(&result, query)
	if err != nil {
		return nil, err
	}
	res := []models.StatisticRes{
		{
			Id:    uuid.NewString(),
			Total: result.TotalUsers,
			Value: result.ValueUsers,
			Title: "Người dùng mới",
		},
		{
			Id:    uuid.NewString(),
			Total: result.TotalBlogs,
			Value: result.ValueBlogs,
			Title: "Bài viết mới",
		},
		{
			Id:    uuid.NewString(),
			Total: result.TotalComments,
			Value: result.ValueComments,
			Title: "Bình luận mới",
		},
	}
	return res, nil
}
