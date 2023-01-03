package product

import "time"

// Modelo de producto
type Model struct {
	ID           uint // 	PK, incrementador
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Slices de Model (registros)
type Models []*Model

// 			----	Interfaces		-----
type Storage interface {
	Create(*Model) error
	Update(*Model) error
	Delete(*Model) error

	GetByID(uint) error
	GetAll() (*Models, error)
}
