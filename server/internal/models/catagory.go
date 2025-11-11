package models

import "database/sql"

type CategoryModel struct {
	DB *sql.DB
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
