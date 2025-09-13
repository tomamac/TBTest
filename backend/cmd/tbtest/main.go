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
		log.Println("warning: .env not loaded, using system env")
	}

	app := fiber.New()

	app.Use(cors.New((cors.Config{
		AllowOrigins: "http://localhost:5173",
	})))

	group := app.Group("/api")

	dsn := os.Getenv("DSN")

	if dsn != "" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect to the database")
		}

		db.AutoMigrate(&entity.Approval{})
		approval.RegisterHandlers(group.Group("/approval"), approval.NewService(approval.NewRepository(db)))
	}

	approval.RegisterHandlers(group.Group("/mock/approval"), approval.NewService(approval.NewMockRepository()))

	app.Listen(":8001")
}
