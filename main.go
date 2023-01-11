package main

import (
	"fmt"
	"log"

	"github.com/Lermaa2/go-db/pkg/product"
	"github.com/Lermaa2/go-db/storage"
)

func main() {
	driver := storage.Postgres

	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	serviceProduct := product.NewService(myStorage)

	m, err := serviceProduct.GetByID(4)
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(m)
}

// // Crear instancia que maneja Producto
// PsqlProduct := database.NewPsqlProduct(DB)
// ProductHandler := product.NewDBHandler(PsqlProduct)

// // // Migrate
// // if err := ProductHandler.MigrateTable(); err != nil {
// // 	log.Fatalf("ProductHandler.MigrateTable: %v", err)
// // }

// // // Crear un registro de Producto
// // m := &product.Product{
// // 	Name:         "Tigre",
// // 	Price:        2,
// // 	Observations: "Una 🪢🪢 random",
// // }
// // if err := ProductHandler.Create(m); err != nil {
// // 	log.Fatalf("ProductHandler.Create(): %+v", err)
// // }
// // fmt.Printf("m:	\n%+v\n", m)

// // GetAll
// ms, err := ProductHandler.GetAll()
// if err != nil {
// 	log.Fatalf("ProductHandler.GetAll(): %v", err)
// }
// fmt.Println("\nms:")
// fmt.Println(ms)

// // GetbyID
// m, err := ProductHandler.GetByID(6)
// switch {
// case errors.Is(err, sql.ErrNoRows):
// 	fmt.Println("No hay product con ese ID")
// case err != nil:
// 	log.Fatalf("ProductHandler.GetByID(): %v", err)
// default:
// 	fmt.Println(m)
// }

// // // Update
// // mod := &product.Product{
// // 	ID:           6,
// // 	Name:         "Perico",
// // 	Price:        69,
// // 	Observations: "esto es un loro 🦜",
// // }

// // err = ProductHandler.Update(mod)
// // if err != nil {
// // 	log.Fatalf("ProductHandler.Update(): %v", err)
// // }

// // // Delete
// // err = ProductHandler.Delete(1)
// // if err != nil {
// // 	log.Fatalf("ProductHandler.Delete(): %v", err)
// // }

// //-
// //-
// //-
// //-
// //-
// // // Crear instancia que maneja InvoiceHeader
// // PsqlInvoiceHeade := database.NewPsqlInvoiceHeader(DB)
// // InvoiceHeaderHandler := invoiceheader.NewDBHandler(PsqlInvoiceHeade)
// // if err := InvoiceHeaderHandler.MigrateTable(); err != nil {
// // 	log.Fatalf("InvoiceHeaderHandler.MigrateTable: %v", err)
// // }

// // // Crear instancia que maneja InvoiceItem
// // PsqlInvoiceItem := database.NewPsqlInvoiceItem(DB)
// // InvoiceItemHandler := invoiceitem.NewDBHandler(PsqlInvoiceItem)
// // if err := InvoiceItemHandler.MigrateTable(); err != nil {
// // 	log.Fatalf("InvoiceItemHandler.MigrateTable: %v", err)
// // }

// -
//	Tranxanctions

// storageHeader := database.NewPsqlInvoiceHeader(DB)
// storageItems := database.NewPsqlInvoiceItem(DB)

// storageInvoice := database.NewPsqlInvoice(
// 	DB,
// 	storageHeader,
// 	storageItems,
// )

// m := &invoice.Invoice{
// 	Header: &invoiceheader.Invoiceheader{
// 		Client: "Jaime",
// 	},
// 	Items: invoiceitem.InvoiceitemList{
// 		&invoiceitem.Invoiceitem{ProductID: 5},
// 		&invoiceitem.Invoiceitem{ProductID: 2},
// 		&invoiceitem.Invoiceitem{ProductID: 99},
// 	},

// 	serviceInvoice := invoice.NewDBHandler(storageInvoice)
// if err := serviceInvoice.Create(m); err != nil {
// 	log.Fatalf("invoice.Create: %v", err)
// }
