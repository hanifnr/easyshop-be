package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
)

const (
	projectID  = "easy-shop-364408"
	bucketName = "images-es-bucket"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var cl *storage.Client

func init() {
	// os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/hanifnr/.config/gcloud/application_default_credentials.json")
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	cl = client
}

func GetImageUploader() *ClientUploader {
	return &ClientUploader{
		cl:         cl,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "images/",
	}
}

func GetImageFile(w http.ResponseWriter, r *http.Request) (multipart.File, StatusReturn) {
	_, file, err := r.FormFile("image_file")
	if err != nil {
		return nil, StatusReturn{ErrCode: ErrIO, Message: err.Error()}
	}

	blobFile, err := file.Open()
	if err != nil {
		return nil, StatusReturn{ErrCode: ErrIO, Message: err.Error()}
	}

	return blobFile, StatusReturnOK()
}

func (c *ClientUploader) UploadFile(file multipart.File, fileName string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func GenerateSignedUrl(objName string) string {
	url, err := cl.Bucket(bucketName).SignedURL(objName, &storage.SignedURLOptions{
		GoogleAccessID: "hanif.nr11@gmail.com",
	})
	if err != nil {
		// TODO: Handle error.
	}
	return url
}
