package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tomamac/tbtest/backend/internal/approval"
	"github.com/tomamac/tbtest/backend/internal/entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}

	app := fiber.New()

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&entity.Approval{})

	group := app.Group("/api")

	approval.RegisterHandlers(group.Group("/approval"), approval.NewService(approval.NewRepository(db)))

	app.Listen(":8001")
}
