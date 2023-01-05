package product

import "time"

// Product representa un producto
type Product struct {
	ID           uint // PK, incrementador
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// ProductList es una lista de productos
type ProductList []*Product

// ProductStorage es una interfaz que define la funcionalidad necesaria para almacenar y gestionar productos en una base de datos
type DBKeeper interface {
	// Migrate crea la tabla de productos en la base de datos
	CreateTable() error

	//Create(*Model) error
	//Update(*Model) error
	//Delete(*Model) error

	//GetByID(uint) error
	//GetAll() (*Models, error)
}

// DataManager proporciona una capa de abstracción para interactuar con la base de datos de productos
type DBHandler struct {
	storage DBKeeper
}

// NewDBHandler crea una nueva instancia de DBHandler
func NewDBHandler(s DBKeeper) *DBHandler {
	return &DBHandler{s}
}

// CreateTable ejecuta la creacion de la tabla productos en la base de datos
func (s *DBHandler) CreateTable() error {
	return s.storage.CreateTable()
}
