package main

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/paitpat003/goredis/handlers"
	"github.com/paitpat003/goredis/repositories"
	"github.com/paitpat003/goredis/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	productService := services.NewCatalogService(productRepo)
	productHandler := handlers.NewCatalogServiceHandler(productService)

	app := fiber.New()

	app.Get("/products", productHandler.GetProducts)

	app.Listen(":3000")
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:pass1234@tcp(localhost:3306)/mydatabase")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}