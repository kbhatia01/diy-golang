package connection

import (
	"diy/src/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl := "postgres://postgres:karan@localhost:5432/test?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Product{})
	return db
}
