package product

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// -
// Product representa un producto con sus atributos
type Product struct {
	ID           uint      // ID es la clave primaria del producto y se incrementa automáticamente
	Name         string    // Name es el nombre del producto
	Observations string    // Observations es un campo opcional para almacenar observaciones adicionales sobre el producto
	Price        int       // Price es el precio del producto en unidades monetarias
	CreatedAt    time.Time // CreatedAt es la fecha de creación del producto
	UpdatedAt    time.Time // UpdatedAt es un slice de fechas que registra cada vez que se actualiza el producto
}

// Solo para visualziacion
func (m *Product) String() string {
	return fmt.Sprintf("%02d | %-20s | %-60s | %5v | %10s | %10s",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

// ProductList es un slice de punteros a Product
type ProductList []*Product

func (m ProductList) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-60s | %5s | %10s | %10s\n",
		"id", "name", "observations", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

// -
// DBKeeper es una interface que describe las operaciones que se pueden realizar sobre una tabla de productos en una base de datos
type DBKeeper interface {
	// CreateTable crea la tabla de productos en la base de datos si aún no existe
	MigrateTable() error
	Create(*Product) error

	GetAll() (ProductList, error)
	GetByID(uint) (*Product, error)

	Update(*Product) error

	Delete(uint) error
}

// -
// DBHandler es una estructura que mantiene una referencia a una implementación de DBKeeper
type DBHandler struct {
	dbKeeper DBKeeper
}

func NewDBHandler(s DBKeeper) *DBHandler {
	return &DBHandler{s}
}

// Migrar tabla a PostgreSQL
func (s *DBHandler) MigrateTable() error {
	return s.dbKeeper.MigrateTable()
}

// Crear un producto
func (s *DBHandler) Create(m *Product) error {
	m.CreatedAt = time.Now()
	return s.dbKeeper.Create(m)
}

// -
func (s *DBHandler) GetAll() (ProductList, error) {
	return s.dbKeeper.GetAll()
}

// -
func (s *DBHandler) GetByID(id uint) (*Product, error) {
	return s.dbKeeper.GetByID(id)
}

// -
var ErrIDNonFound = errors.New("el producto no contiene un ID")

func (s *DBHandler) Update(m *Product) error {
	if m.ID == 0 {
		return ErrIDNonFound
	}
	m.UpdatedAt = time.Now()
	return s.dbKeeper.Update(m)
}

// -
func (s *DBHandler) Delete(id uint) error {
	if id == 0 {
		return ErrIDNonFound
	}
	return s.dbKeeper.Delete(id)
}
