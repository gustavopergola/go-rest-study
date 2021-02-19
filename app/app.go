package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavopergola/go-rest-study/entity"
	"github.com/gustavopergola/go-rest-study/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=pipoca port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database, error: %s", err.Error())
	}
	return db
}

func StartServer(db *gorm.DB){
	app := fiber.New()

	SetupRoutes(app, db)
	err := app.Listen(":8000")
	if err != nil {
		log.Fatalf("Error binding port! Maybe already in use?")
	}
}