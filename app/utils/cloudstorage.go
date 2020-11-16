package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
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

	// dst, err := os.Create(filepath.Join("temp", file.Filename))
	// if err != nil {
	// 	return "", fmt.Errorf("Error creating file : %v", err)
	// }

	// if _, err = io.Copy(dst, src); err != nil {
	// 	return "", fmt.Errorf("Error copying %v", err)
	// }

	// TODO Preprocess the video file here

	// defer dst.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	obj := bucket.Object(file.Filename)
	acl := obj.ACL()

	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		log.Println(err)

		return "", fmt.Errorf("ACLHandle.Set: %v", err)
	}

	w := obj.NewWriter(ctx)

	objAtrrs, _ := obj.Attrs(ctx)

	if _, err = io.Copy(w, src); err != nil {
		log.Println("Error copying to storage : ", err)
		return "", fmt.Errorf("Error copying %v", err)
	}

	if err := w.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	log.Println(w.Bucket)

	return objAtrrs.MediaLink, nil
}
