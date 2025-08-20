package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/paitpat003/goredis/repositories"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{
		productRepo: productRepo,
		redisClient: redisClient,
	}
}

func (s catalogServiceRedis) GetProducts() (products []Product, err error) {

	key := "service::GetProducts"

	// Redis GET
	if productsJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(productsJson), &products) == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	// Repository
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productsDB {

		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	// Redis SET
	if jsonData, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(jsonData), time.Second * 10)
	}

	fmt.Println("database")
	return products, nil
}
