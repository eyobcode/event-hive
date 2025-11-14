package models

import "database/sql"

type CategoryModel struct {
	DB *sql.DB
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (m *CategoryModel) CategoryExist(id int) (bool, error) {
	var exists bool
	err := m.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id=$1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
