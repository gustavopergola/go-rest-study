package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gustavopergola/go-rest-study/entity"
	"github.com/gustavopergola/go-rest-study/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Mount Student routes
	studentHandler := handler.NewStudentHandler(db)
	studentHandler.MountRoutes(app)

	healthHandler := &handler.HealthHandler{}
	healthHandler.MountRoutes(app)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Student{})
}

func StartDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database, error: %s", err.Error())
	}
	return db
}

func StartServer(db *gorm.DB){
	app := fiber.New()

	SetupRoutes(app, db)
	err := app.Listen(":" + os.Getenv("PORT"))

	app.Use(favicon.New(favicon.Config{
		File: "./favicon.ico",
	}))

	if err != nil {
		log.Fatalf("Error binding port! Maybe already in use?")
	}
}