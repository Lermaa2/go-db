package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" // paquete del driver para PostgreSQL
)

var (
	// db es un puntero a una conexión a la base de datos
	db *sql.DB
	// once es una estructura sync.Once que se utiliza para garantizar que la conexión se establece solo una vez
	once sync.Once
)

// ConnectToPostgresDB se encarga de establecer una conexión a la base de datos
func ConnectToPostgresDB() {
	// once.Do ejecuta la función solo una vez
	once.Do(func() {
		var err error
		// sql.Open se utiliza para abrir una conexión a la base de datos utilizando el controlador "github.com/lib/pq" para PostgreSQL
		db, err = sql.Open("postgres", "postgres://postgres:Barranquilla2812@localhost:5432/go-db?sslmode=disable")
		if err != nil {
			// log.Fatalf imprime un mensaje de error y termina la aplicación
			log.Fatalf("No se puedo Conectar con Base de Datos:	%v", err)
		}

		// db.Ping hace una solicitud simple a la base de datos para verificar que la conexión está funcionando correctamente
		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo hacer PING:	%v", err)
		}
		fmt.Println("	***	Conectado a PostgeSQL	***")
	})
}

// GetDB devuelve un puntero a la conexión a la base de datos
func GetDB() *sql.DB {
	return db
}
