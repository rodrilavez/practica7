package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	// Configurar la conexión a la base de datos MySQL
	DB, err = sql.Open("mysql", "root:practica@tcp(localhost:3306)/practica7db")
	if err != nil {
		log.Fatal(err)
	}

	// Verificar la conexión
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Conexión exitosa a la base de datos MySQL.")

	// Crear la tabla si no existe
	createTable := `CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE
    );`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
