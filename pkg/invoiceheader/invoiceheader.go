package invoiceheader

import "time"

// Modelo de invoiceheader
type Model struct {
	ID        uint // 	PK, incrementador
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
