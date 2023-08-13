package services

import (
	"context"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImageToCloudinary(r *http.Request) (*uploader.UploadResult, error) {
	r.ParseMultipartForm(32 << 20)

	cloudinaryString := os.Getenv("CLOUDINARY_URL")

	file, _, err := r.FormFile("file")
	
	if err != nil {
		return nil, err
	}

	cld, err := cloudinary.NewFromURL(cloudinaryString)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		return nil, err
	}

	return resp, nil

}