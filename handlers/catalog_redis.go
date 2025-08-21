package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/paitpat003/goredis/services"
)

type catalogHandlerRedis struct {
	catalogSrv services.CatalogService
	redisClient *redis.Client
}

func NewCatalogServiceHandlerRedis(catalogSrv services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{
		catalogSrv: catalogSrv,
		redisClient: redisClient,
	}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {

	key := "handler::getProducts"

	// Redis Get
	productsJson, err := h.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		fmt.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(productsJson)
	}

	// Service
	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":  "ok",
		"products": products,
	}

	// Redis Set
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	err = h.redisClient.Set(context.Background(), key, string(data), time.Second * 10).Err()
	if err != nil {
		return err
	}

	fmt.Println("database")

	return c.JSON(response)
}