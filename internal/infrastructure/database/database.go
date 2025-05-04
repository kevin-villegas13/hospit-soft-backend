package database

import (
	"hospit-soft-backend/internal/domain/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// Obtener la URL de la base de datos desde las variables de entorno
	databaseURL := os.Getenv("DATABASE_URL")

	// Conectar a la base de datos
	db, _ := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	// Migración de modelos
	db.AutoMigrate(&entity.User{})

	return db
}
