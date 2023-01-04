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
	Migrate() error

	//Create(*Model) error
	//Update(*Model) error
	//Delete(*Model) error

	//GetByID(uint) error
	//GetAll() (*Models, error)
}

// Servicios de product
type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s} // Retorna puntero de servicio.
}

// Migrate usado para migrar product...
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
