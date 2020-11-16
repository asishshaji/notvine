package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/asishshaji/notvine/app"
	"github.com/asishshaji/notvine/app/controller"
	"github.com/asishshaji/notvine/app/repository"
	"github.com/asishshaji/notvine/app/usecase"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return
	}

	log.Println("Loaded .env file")
}

func main() {
	port := os.Getenv("SERVER_PORT")
	dbName := os.Getenv("DB_NAME")
	mongodbURL := os.Getenv("MONGODB_URL")
	storageBucket := os.Getenv("STORAGE_BUCKET")
	credentialFilePath := os.Getenv("CRED_FILE")

	db := initDB(mongodbURL)
	bucket := initStorage(storageBucket, credentialFilePath)

	repo := repository.NewMongoRepo(db, dbName)
	usecase := usecase.NewAppUsecase(*repo)
	controller := controller.NewAppController(*usecase, bucket)

	app := app.NewApp(port, *controller)

	app.RunServer()

}

func initStorage(storageBucket, credentialFilePath string) *storage.BucketHandle {
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

func initDB(mongodbURL string) *mongo.Database {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")

	return client.Database("DB")
}
