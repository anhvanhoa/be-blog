package media_service

import (
	"be-blog/src/config"
	"be-blog/src/constants"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/models"
	"context"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func GetImages(id string) ([]entities.Photo, error) {
	var photos []entities.Photo = []entities.Photo{}
	err := config.DB.Model(&photos).Where("author_id = ?", id).Select()
	if err != nil {
		return photos, err
	}
	return photos, nil
}

func UploadImage(file multipart.File, info *multipart.FileHeader, userId, title string) error {
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	fileType := info.Header["Content-Type"][0]
	if constants.MineTypes[fileType] == "" {
		err := errors.NewErrorBadRequest("File type is not supported")
		return err
	}
	photo := entities.Photo{
		ID:       uuid.New().String(),
		AuthorId: userId,
		Title:    title,
	}
	ctx := context.Background()
	paramImage := uploader.UploadParams{
		Folder:      "AnhVanHoaCom",
		DisplayName: title,
	}
	if title == "" {
		paramImage.DisplayName = info.Filename
		photo.Title = info.Filename
	}
	result, err := config.Cloudinary.Upload.Upload(ctx, file, paramImage)
	if err != nil {
		return err
	}
	photo.Url = result.SecureURL
	photo.PublicID = result.PublicID
	_, err = config.DB.Model(&photo).Insert()
	if err != nil {
		return err
	}
	return nil
}

func DeleteImage(userId, id string) error {
	photo := entities.Photo{
		ID:       id,
		AuthorId: userId,
	}
	err := config.DB.Model(&photo).WherePK().Where("author_id = ?", photo.AuthorId).Select()
	if err != nil {
		return err
	}
	if photo.PublicID != "" {
		param := uploader.DestroyParams{
			PublicID: photo.PublicID,
		}
		_, err := config.Cloudinary.Upload.Destroy(context.Background(), param)
		if err != nil {
			return err
		}
	}
	_, err = config.DB.Model(&photo).WherePK().Where("author_id = ?", photo.AuthorId).Delete()
	if err != nil {
		return err
	}
	return nil
}

func UpdateImage(body models.PhotoReq) error {
	image := entities.PhotoUpdate{
		ID:        body.Id,
		AuthorId:  body.AuthorId,
		Title:     body.Title,
		UpdatedAt: time.Now(),
	}
	_, err := config.DB.Model(&image).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
