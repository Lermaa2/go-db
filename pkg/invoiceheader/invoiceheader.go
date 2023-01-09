package invoiceheader

import "time"

// Modelo de invoiceheader
type Invoiceheader struct {
	ID        uint // 	PK, incrementador
	Client    string
	CreatedAt time.Time
	UpdatedAt []time.Time
}

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
