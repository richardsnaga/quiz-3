package repository

import (
	"Quiz-3/structs"
	"database/sql"
)

func GetAllCategory(db *sql.DB) (err error, results []structs.Category) {
	sql := `SELECT * FROM categories`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}

	return
}

func CreateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := `INSERT INTO categories (name) VALUES ($1)`

	errs := db.QueryRow(sql, category.Name)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := `UPDATE categories SET name = $1, updated_at = NOW()::timestamp WHERE id = $2`

	errs := db.QueryRow(sql, category.Name, category.ID)

	return errs.Err()
}
