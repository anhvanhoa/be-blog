package statistic

import (
	"be-blog/src/config"
	"be-blog/src/entities"
)

func StatisticCommentNew() ([]entities.StatisticComment, error) {
	comments := []entities.StatisticComment{}
	err := config.DB.Model(&comments).Join("INNER JOIN blogs b ON b.id = statistic_comment.blog_id").
		Join("INNER JOIN users u ON u.id = statistic_comment.author_id").
		ColumnExpr("b.title, statistic_comment.id, statistic_comment.content, statistic_comment.created_at, u.full_name").
		Where("statistic_comment.created_at >= date_trunc('day', CURRENT_DATE) AND statistic_comment.created_at < date_trunc('day', CURRENT_DATE + INTERVAL '1 day')").Select()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
