package comment

import (
	"be-blog/src/config"
	"be-blog/src/constants"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/mail"
	"be-blog/src/models"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func CreateComment(body models.CommentReq, usernameReceiver string, authorComment models.AuthorComment) error {
	status := true
	var comment entities.Comment = entities.Comment{
		ID:              uuid.New().String(),
		Content:         body.Content,
		AuthorId:        body.AuthorId,
		BlogId:          body.BlogId,
		ParentCommentId: body.ParentCommentId,
		Status:          &status,
	}
	var blog entities.BlogComment
	queryBlogAuthor := `SELECT b.id, b.title, u.email FROM blogs b JOIN public.users u on b.author_id = u.id WHERE b.id = ?;`
	_, err := config.DB.Query(&blog, queryBlogAuthor, body.BlogId)
	if err != nil {
		return err
	}
	if usernameReceiver != "" {
		var userReceiver entities.Auth
		err = config.DB.Model(&userReceiver).Where("username = ? AND status = ?", strings.ToLower(usernameReceiver), constants.STATUS_ACTIVE).Select()
		if err != nil {
			if err == pg.ErrNoRows {
				return errors.NewErrorBadRequest("Tài khoản không chính xác")
			}
			return err
		}
		go mail.SendMail(mail.PayloadMail{
			Tos:      []string{userReceiver.Email},
			Template: "COMMENT_REPLY",
			Data:     map[string]interface{}{"fullName": authorComment.FullName, "content": body.Content, "link": viper.GetString("clientUrl") + "/blog/" + body.BlogId, "nameBlog": blog.Title},
			From:     viper.GetString("mailSender"),
		})
	}
	go mail.SendMail(mail.PayloadMail{
		Tos:      []string{blog.Email},
		Template: "COMMENT",
		Data:     map[string]interface{}{"fullName": authorComment.FullName, "content": body.Content, "link": viper.GetString("clientUrl") + "/blog/" + body.BlogId, "nameBlog": blog.Title},
		From:     viper.GetString("mailSender"),
	})
	_, err = config.DB.Model(&comment).Insert()
	if err != nil {
		return err
	}
	return nil
}
