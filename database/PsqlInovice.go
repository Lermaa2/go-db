package database

import (
	"database/sql"

	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceheader"
	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceitem"
)

type PsqlInovice struct {
	db             *sql.DB
	DBKeeperHeader invoiceheader.DBKeeper
	DBKeeperitem   invoiceitem.DBKeeper
}

func NewPsqlInvoice(db *sql.DB, h invoiceheader.DBKeeper, itm invoiceitem.DBKeeper) *PsqlInovice {
	return &PsqlInovice{db, h, itm}
}

func (p *PsqlInovice) Create(m *invoice.Invoice) error {

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

}
