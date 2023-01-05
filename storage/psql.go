package storage

import (
	"database/sql"
	"fmt"
)

// Best Practice: usar Const para definir las Queries a usar
const (
	// psqlCreateProduct es una constante que contiene una consulta SQL para crear la tabla "products"
	// en la base de datos
	psqlCreateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY(id)
	);` // Solo acepta ``
)

// PsqlsProduct es una estructura que se utiliza para interactuar con la tabla "products" en la base de datos
type PsqlsProduct struct {
	// db es un puntero a la conexión a la base de datos
	db *sql.DB
}

// NewPsqlsProduct crea y devuelve una estructura PsqlsProduct
func NewPsqlsProduct(db *sql.DB) *PsqlsProduct {
	return &PsqlsProduct{db}
}

// CreateTable ejecuta la consulta psqlCreateProduct para crear la tabla "products" en la base de datos
func (p *PsqlsProduct) CreateTable() error {
	// db.Prepare se utiliza para preparar una consulta SQL para su ejecución posterior
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	// stmt.Close cierra el objeto stmt después de que se haya ejecutado
	defer stmt.Close()

	// stmt.Exec ejecuta la consulta SQL preparada
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("	- Migración de producto ejecutada correctamente")
	return nil
}
