package invoiceitem

import "time"

// Modelo de invoiceitem
type Model struct {
	ID              uint // 	PK, incrementador
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
