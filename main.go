package main

// go mod init github.com/Lermaa2/github.com/Lermaa2/go-db
// go build main.go
// ./main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Lermaa2/github.com/Lermaa2/go-db/database"
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

	// // Migrate
	// if err := ProductHandler.MigrateTable(); err != nil {
	// 	log.Fatalf("ProductHandler.MigrateTable: %v", err)
	// }

	// // Crear un registro de Producto
	// m := &product.Product{
	// 	Name:         "Tigre",
	// 	Price:        2,
	// 	Observations: "Una 🪢🪢 random",
	// }
	// if err := ProductHandler.Create(m); err != nil {
	// 	log.Fatalf("ProductHandler.Create(): %+v", err)
	// }
	// fmt.Printf("m:	\n%+v\n", m)

	// GetAll
	ms, err := ProductHandler.GetAll()
	if err != nil {
		log.Fatalf("ProductHandler.GetAll(): %v", err)
	}
	fmt.Println("\nms:")
	fmt.Println(ms)

	// GetbyID
	m, err := ProductHandler.GetByID(6)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay product con ese ID")
	case err != nil:
		log.Fatalf("ProductHandler.GetByID(): %v", err)
	default:
		fmt.Println(m)
	}

	// // Update
	// mod := &product.Product{
	// 	ID:           6,
	// 	Name:         "Perico",
	// 	Price:        69,
	// 	Observations: "esto es un loro 🦜",
	// }

	// err = ProductHandler.Update(mod)
	// if err != nil {
	// 	log.Fatalf("ProductHandler.Update(): %v", err)
	// }

	// // Delete
	// err = ProductHandler.Delete(1)
	// if err != nil {
	// 	log.Fatalf("ProductHandler.Delete(): %v", err)
	// }

	//-
	//-
	//-
	//-
	//-
	// // Crear instancia que maneja InvoiceHeader
	// PsqlInvoiceHeade := database.NewPsqlInvoiceHeader(DB)
	// InvoiceHeaderHandler := invoiceheader.NewDBHandler(PsqlInvoiceHeade)
	// if err := InvoiceHeaderHandler.MigrateTable(); err != nil {
	// 	log.Fatalf("InvoiceHeaderHandler.MigrateTable: %v", err)
	// }

	// // Crear instancia que maneja InvoiceItem
	// PsqlInvoiceItem := database.NewPsqlInvoiceItem(DB)
	// InvoiceItemHandler := invoiceitem.NewDBHandler(PsqlInvoiceItem)
	// if err := InvoiceItemHandler.MigrateTable(); err != nil {
	// 	log.Fatalf("InvoiceItemHandler.MigrateTable: %v", err)
	// }

}
