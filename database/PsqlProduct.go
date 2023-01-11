package database

import (
	"database/sql"
	"fmt"

	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/product"
)

// Best Practice: usar Const para definir las Queries a usar
const (
	// psqlMigrateProduct es una constante que contiene una consulta SQL para crear la tabla "products"
	// en la base de datos
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products( 
		id SERIAL NOT NULL, 
		name VARCHAR(25) NOT NULL, 
		observations VARCHAR(100), 
		price INT NOT NULL, 
		created_at TIMESTAMP NOT NULL DEFAULT now(), 
		updated_at TIMESTAMP, 

		CONSTRAINT products_id_pk PRIMARY KEY(id) 
		);`
	// Marcadores de posicion en postgreSQL: 	$1 $2 $3
	psqlCreateProduct = `INSERT INTO products(
		name, observations, price, created_at) 
		VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct = `SELECT id, name, observations, price, created_at, updated_at
	FROM products`
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2,
	price = $3, updated_at = $4 WHERE id = $5`
	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
) // `para tabular` "No tabula"

// -
// ProductTable es una estructura que se utiliza para interactuar con la tabla "products" en la base de datos
type PsqlProduct struct {
	// db es un puntero a la conexión a la base de datos
	db *sql.DB
}

// NewProductTable crea y devuelve una estructura ProductTable
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// -
// Implementa interfaz de DBKeeper con consulta psqlCreateProduct
// MigrateTable ejecuta la consulta psqlMigrateProduct para crear la tabla "products" en la base de datos
func (p *PsqlProduct) MigrateTable() error {
	// db.Prepare se utiliza para preparar una consulta SQL para su ejecución posterior
	stmt, err := p.db.Prepare(psqlMigrateProduct)
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

	fmt.Println("	- Migración de products ejecutada correctamente")
	return nil
}

// Implementa interfaz de DBKeeper con consulta psqlCreateProduct
func (p *PsqlProduct) Create(m *product.Model) error {

	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("\n	- Se crear Producto correctamente")
	return nil
}

// - Helpers
type scanner interface {
	Scan(dest ...any) error
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	// Nulls handling
	observationNull := sql.NullString{}
	updateAtNull := sql.NullTime{}

	err := s.Scan(

		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updateAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updateAtNull.Time

	return m, nil
}

// Implementa interfaz de DBKeeper con SQL para retornar todo
func (p *PsqlProduct) GetAll() (product.ProductList, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.ProductList, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// Implementa interfaz de DBKeeper con SQL para retornar todo
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {

	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Implementa interfaz de DBKeeper con SQL para Actualizar registro
func (p *PsqlProduct) Update(m *product.Model) error {
	stm, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stm.Close()

	res, err := stm.Exec(
		&m.Name,
		stringToNull(m.Observations),
		&m.Price,
		timeToNull(m.UpdatedAt),
		&m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("\nno existe el prodcuto con ID %d", m.ID)
	}

	fmt.Println("\nRegistro actualizado")
	return nil
}

// Implementa interfaz de DBKeeper con SQL para Eliminar registro
func (p *PsqlProduct) Delete(id uint) error {
	stm, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stm.Close()

	res, err := stm.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("\nno existe el producto con ID %d", id)
	} else if rowsAffected > 1 {
		return fmt.Errorf("\nSe eliminaron varios registros con ID %d", id)
	}

	fmt.Println("\nRegistro ELIMINADO ❌")
	return nil
}
