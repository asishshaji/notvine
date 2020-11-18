package main

import (
	"log"
	"os"

	"github.com/asishshaji/notvine/app"
	"github.com/asishshaji/notvine/app/controller"
	"github.com/asishshaji/notvine/app/repository"
	"github.com/asishshaji/notvine/app/usecase"
	"github.com/asishshaji/notvine/app/utils"
	"github.com/joho/godotenv"
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
	collectionName := os.Getenv("COLLECTION_NAME")

	db := utils.InitDB(mongodbURL, dbName)
	bucket := utils.InitStorage(storageBucket, credentialFilePath)

	repo := repository.NewMongoRepo(db, collectionName)
	usecase := usecase.NewAppUsecase(repo)
	controller := controller.NewAppController(*usecase, bucket)

	app := app.NewApp(port, controller)

	app.RunServer()

}
