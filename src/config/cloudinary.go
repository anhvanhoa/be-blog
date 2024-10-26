package config

import "github.com/cloudinary/cloudinary-go/v2"

var Cloudinary *cloudinary.Cloudinary

func NewCloudinary() (err error) {
	Cloudinary, err = cloudinary.NewFromParams("dlzbq5oho", "973672683344266", "GI9yaxUH8_uAYFvVcwG9A4eVZ-U")
	return err
}
