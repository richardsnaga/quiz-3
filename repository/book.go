package repository

import (
	"Quiz-3/structs"
	"database/sql"
)

func GetAllBook(db *sql.DB) (err error, results []structs.Book) {
	sql := `SELECT * FROM books`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.UpdatedAt, &book.CategoryId)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := `INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	errs := db.QueryRow(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryId)

	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := `UPDATE books SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8, updated_at = NOW()::timestamp WHERE id = $9`

	errs := db.QueryRow(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryId, book.ID)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := `DELETE FROM books WHERE id = $1`

	errs := db.QueryRow(sql, book.ID)

	return errs.Err()
}

func GetBookCategory(db *sql.DB, id int) (err error, results []structs.Book) {
	sql := `SELECT * FROM books WHERE category_id = $1`

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.UpdatedAt, &book.CategoryId)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}
