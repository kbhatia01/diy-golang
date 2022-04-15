package handlers

import "gorm.io/gorm"

type handler struct {
	Db *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{
		Db: db,
	}
}
