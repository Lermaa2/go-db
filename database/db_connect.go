package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" // paquete del driver para PostgreSQL
)

var (
	// DB es un puntero a una conexión a la base de datos
	DB *sql.DB
	// once es una estructura sync.Once que se utiliza para garantizar que la conexión se establece solo una vez
	once sync.Once
)

// ConnectToDB se encarga de establecer una conexión a la base de datos
func ConnectToDB() {
	// once.Do ejecuta la función solo una vez
	once.Do(func() {
		var err error
		// sql.Open se utiliza para abrir una conexión a la base de datos utilizando el controlador "github.com/lib/pq" para PostgreSQL
		DB, err = sql.Open("postgres", "postgres://postgres:Barranquilla2812@localhost:5432/go-db?sslmode=disable")
		if err != nil {
			// log.Fatalf imprime un mensaje de error y termina la aplicación
			log.Fatalf("No se puedo Conectar con Base de Datos:	%v", err)
		}

		// db.Ping hace una solicitud simple a la base de datos para verificar que la conexión está funcionando correctamente
		if err = DB.Ping(); err != nil {
			log.Fatalf("No se pudo hacer PING:	%v", err)
		}
		fmt.Println("	***	Conectado a PostgeSQL	***")
	})
}

// DBPointer devuelve un puntero a la conexión a la base de datos
func DBPointer() *sql.DB {
	return DB
}
