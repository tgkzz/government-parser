package repository

import (
	"database/sql"
	"parser/internal/models"
)

type Repository struct {
	db *sql.DB
}

type IRepository interface {
	UploadOne(lot models.Result) error
}

func NewRepository(DB *sql.DB) *Repository {
	return &Repository{db: DB}
}
