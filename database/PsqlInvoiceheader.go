package database

import (
	"database/sql"
	"fmt"

	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceheader"
)

// Best Practice: usar Const para definir las Queries a usar
const (
	// psqlMigrateInvoiceHeader es una constante que contiene una consulta SQL para crear la tabla "invoice_header"
	// en la base de datos
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers ( 
		id SERIAL NOT NULL, 
		client VARCHAR(100) NOT NULL, 
		created_at TIMESTAMP NOT NULL DEFAULT now(), 
		updated_at TIMESTAMP, 
		
		CONSTRAINT invoice_header_id_pk PRIMARY KEY(id) 
		);`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at`

// Solo acepta “
)

type PsqlInvoiceHeader struct {
	db *sql.DB
}

func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

func (p *PsqlInvoiceHeader) MigrateTable() error {
	// db.Prepare se utiliza para preparar una consulta SQL para su ejecución posterior
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("	- Migración de InvoiceHeader ejecutada correctamente")
	return nil
}

func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Invoiceheader) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
