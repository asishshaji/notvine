package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
)

func UploadVideo(file *multipart.FileHeader, bucket *storage.BucketHandle) (string, error) {

	//TODO Delete files in temp folder here

	src, err := file.Open()

	if err != nil {
		return "", fmt.Errorf("Failed to run os.Open:  %v", err)
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join("temp", file.Filename))
	if err != nil {
		return "", fmt.Errorf("Error creating file : %v", err)
	}

	// TODO Preprocess the video file here

	defer dst.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	wc := bucket.Object(file.Filename).NewWriter(ctx)

	if _, err = io.Copy(wc, src); err != nil {
		return "", fmt.Errorf("Error copying %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}
	return wc.MediaLink, nil
}
