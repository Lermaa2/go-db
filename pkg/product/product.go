package product

import "time"

// Product representa un producto con sus atributos
type Product struct {
	ID           uint        // ID es la clave primaria del producto y se incrementa automáticamente
	Name         string      // Name es el nombre del producto
	Observations string      // Observations es un campo opcional para almacenar observaciones adicionales sobre el producto
	Price        int         // Price es el precio del producto en unidades monetarias
	CreatedAt    time.Time   // CreatedAt es la fecha de creación del producto
	UpdatedAt    []time.Time // UpdatedAt es un slice de fechas que registra cada vez que se actualiza el producto
}

// ProductList es un slice de punteros a Product
type ProductList []*Product

// DBKeeper es una interface que describe las operaciones que se pueden realizar sobre una tabla de productos en una base de datos
type DBKeeper interface {
	// CreateTable crea la tabla de productos en la base de datos si aún no existe
	MigrateTable() error
}

// DBHandler es una estructura que mantiene una referencia a una implementación de DBKeeper
type DBHandler struct {
	dbKeeper DBKeeper
}

func NewDBHandler(s DBKeeper) *DBHandler {
	return &DBHandler{s}
}

func (s *DBHandler) MigrateTable() error {
	return s.dbKeeper.MigrateTable()
}
