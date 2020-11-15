package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/asishshaji/notvine/app"
	"github.com/asishshaji/notvine/app/controller"
	"github.com/asishshaji/notvine/app/repository"
	"github.com/asishshaji/notvine/app/usecase"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	db := initDB(mongodbURL)
	repo := repository.NewMongoRepo(db, dbName)

	usecase := usecase.NewAppUsecase(*repo)
	controller := controller.NewAppController(*usecase)
	app := app.NewApp(port, *controller)

	app.RunServer()

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
