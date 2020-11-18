package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

func UploadVideo(file *multipart.FileHeader, bucket *storage.BucketHandle) (string, error) {

	//TODO Delete files in temp folder here

	src, err := file.Open()

	if err != nil {
		return "", fmt.Errorf("Failed to run os.Open:  %v", err)
	}
	defer src.Close()

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

// InitStorage initializes the google cloud storageG
func InitStorage(storageBucket, credentialFilePath string) *storage.BucketHandle {
	config := &firebase.Config{
		StorageBucket: storageBucket,
	}

	opt := option.WithCredentialsFile(credentialFilePath)

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.Bucket("videos")
	if err != nil {
		log.Fatalln(err)
	}

	return bucket

}

// InitDB creates a connection to MongoDB instance
func InitDB(mongodbURL, dbName string) *mongo.Database {
	log.Printf("Starting connection to MongoDB at : %v", mongodbURL)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB. Make sure mongodb instance is running.")

	}

	log.Println("Connected to MongoDB")

	return client.Database(dbName)
}
