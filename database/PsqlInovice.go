package database

import (
	"database/sql"
	"fmt"

	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoice"
	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceheader"
	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceitem"
)

type PsqlInovice struct {
	db             *sql.DB
	DBKeeperHeader invoiceheader.DBKeeper
	DBKeeperItem   invoiceitem.DBKeeper
}

func NewPsqlInvoice(db *sql.DB, h invoiceheader.DBKeeper, itm invoiceitem.DBKeeper) *PsqlInovice {
	return &PsqlInovice{db, h, itm}
}

func (p *PsqlInovice) Create(m *invoice.Invoice) error {

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.DBKeeperHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		fmt.Println("Error, Rollback!")
		return err
	}

	fmt.Printf("Factura creada con ID:	%d\n", m.Header.ID)

	if err := p.DBKeeperItem.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		fmt.Println("Error, Rollback!")
		tx.Rollback()
		return err
	}

	fmt.Printf("Items creados:	%d\n", len(m.Items))

	return tx.Commit()
}
