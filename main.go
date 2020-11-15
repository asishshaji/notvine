package main

import (
	"context"
	"log"
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
	db := initDB()
	repo := repository.NewMongoRepo(db, "repo")

	usecase := usecase.NewAppUsecase(*repo)
	controller := controller.NewAppController(*usecase)
	app := app.NewApp(":9090", *controller)

	app.RunServer()

}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
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

	log.Println("Connected to MongoDB ")

	return client.Database("DB")
}
