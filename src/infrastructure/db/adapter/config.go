package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/juheth/Go-Clean-Arquitecture/src/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
	)

	db, err := (gorm.Open(mysql.Open(dsn), &gorm.Config{}))
	if err != nil {
		log.Fatalf(" Error al abrir conexión con MySQL: %v", err)
	}

	log.Println(" Conexión exitosa a MySQL")
	return db, nil
}
