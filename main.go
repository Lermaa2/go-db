package main

// go mod init github.com/Lermaa2/github.com/Lermaa2/go-db
// go build main.go
// ./main

import (
	"fmt"
	"log"

	"github.com/Lermaa2/github.com/Lermaa2/go-db/database"
	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceheader"
	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceitem"
	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/product"
)

// main es la funcion principal en Go
func main() {
	fmt.Println("\n	--	GO DATABASE	--")

	// ConnectToDB se encarga de establecer una conexión a la base de datos
	database.ConnectToDB()
	DB := database.DBPointer()

	// Crear instancia que maneja Producto
	PsqlProduct := database.NewPsqlProduct(DB)
	ProductHandler := product.NewDBHandler(PsqlProduct)
	if err := ProductHandler.MigrateTable(); err != nil {
		log.Fatalf("ProductHandler.MigrateTable: %v", err)
	}

	// Crear instancia que maneja InvoiceHeader
	InvoiceHeaderTb := database.NewPsqlInvoiceHeader(DB)
	InvoiceHeaderHandler := invoiceheader.NewDBHandler(InvoiceHeaderTb)
	if err := InvoiceHeaderHandler.MigrateTable(); err != nil {
		log.Fatalf("InvoiceHeaderHandler.MigrateTable: %v", err)
	}

	// Crear instancia que maneja InvoiceItem
	InvoiceItemTb := database.NewPsqlInvoiceItem(DB)
	InvoiceItemHandler := invoiceitem.NewDBHandler(InvoiceItemTb)
	if err := InvoiceItemHandler.MigrateTable(); err != nil {
		log.Fatalf("InvoiceItemHandler.MigrateTable: %v", err)
	}

}
