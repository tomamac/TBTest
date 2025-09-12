package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tomamac/tbtest/backend/internal/approval"
	"github.com/tomamac/tbtest/backend/internal/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New((cors.Config{
		AllowOrigins: "http://localhost:5173",
	})))

	group := app.Group("/api")

	approval.RegisterHandlers(group.Group("/approval"), approval.NewService(approval.NewRepository(db)))

	app.Listen(":8001")
}
