package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/juheth/Go-Clean-Arquitecture/src/common/config"
)

type DBConnection struct {
	*sql.DB
}

func NewDBConnection(cfg *config.Config) *DBConnection {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
	)

	result, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf(" Error al abrir conexión con MySQL: %v", err)
	}

	if err := result.Ping(); err != nil {
		log.Fatalf(" No se pudo conectar con MySQL: %v", err)
	}

	log.Println(" Conexión exitosa a MySQL")
	return &DBConnection{result}
}
