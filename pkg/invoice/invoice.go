package invoice

import (
	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceheader"
	"github.com/Lermaa2/github.com/Lermaa2/go-db/pkg/invoiceitem"
)

type Invoice struct {
	Header *invoiceheader.Invoiceheader
	Items  invoiceitem.InvoiceitemList
}

type DBKeeper interface {
	Create(*Invoice) error
}

type DBHandler struct {
	dbKeeper DBKeeper
}

func NewDBHandler(s DBKeeper) *DBHandler {
	return &DBHandler{s}
}

func (s *DBHandler) Create(m *Invoice) error {
	return s.dbKeeper.Create(m)
}
