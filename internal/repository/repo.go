package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

type IRepository interface {
}

func NewRepository(DB *sql.DB) *Repository {
	return &Repository{db: DB}
}
