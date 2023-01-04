package main

// go mod init github.com/Lermaa2/github.com/Lermaa2/go-db
// go build main.go
// ./main

import (
	"fmt"
	"log"

	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/product"
	"github.com/Lermaa2/github.com/Lermaa2/go-db/storage"
)

// main es la funcion principal en Go
func main() {
	fmt.Println("\n	--	GO DATABASE	--")

	// Crear coneccion con base de datos PostgreSQL
	storage.NewPostgresDB()

	// Crear instancia que maneja Producto
	storageProduct := storage.NewPsqlsProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
