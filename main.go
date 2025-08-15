package main

import (
	"fmt"
	"log"

	"github.com/paitpat003/goredis/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()

	productRepo := repositories.NewProductRepositoryDB(db)

	products, err := productRepo.GetProducts()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(products)
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:pass1234@tcp(localhost:3306)/mydatabase")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}