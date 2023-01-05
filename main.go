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

	// ConnectToPostgresDB se encarga de establecer una conexión a la base de datos
	storage.ConnectToPostgresDB()

	// Crear instancia que maneja Producto
	storageProduct := storage.NewPsqlsProduct(storage.GetDB())
	//
	serviceProduct := product.NewDBHandler(storageProduct)
	fmt.Printf("%T,%+v", serviceProduct, serviceProduct)

	if err := serviceProduct.CreateTable(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
