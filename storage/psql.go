package storage

import (
	"database/sql"
	"fmt"
)

// Best Practice: usar Const para definir las Queries a usar
const (
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

// PsqlsProduct usado para trabajar con postgresql y product
type PsqlsProduct struct {
	db *sql.DB
}

// NewPsqlsProduct Retorna un nuevo puntero de PsqlsProdcut
func NewPsqlsProduct(db *sql.DB) *PsqlsProduct {
	return &PsqlsProduct{db}
}

// Migrate implemente la interfaz product.Storage
func (p *PsqlsProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("	- Migración de producto ejecutada correctamente")
	return nil
}
