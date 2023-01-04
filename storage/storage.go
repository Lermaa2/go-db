package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" // paquete del driver para PostgreSQL
)

var (
	db   *sql.DB
	once sync.Once
)

// Modelo de producto
func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:Barranquilla2812@localhost:5432/go-db?sslmode=disable")
		if err != nil {
			log.Fatalf("No se puedo Conectar con Base de Datos:	%v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo hacer PING:	%v", err)
		}
		fmt.Println("	***	Conectado a PostgeSQL	***")
	})
}

// Pool retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}
