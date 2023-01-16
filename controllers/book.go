package controllers

import (
	"Quiz-3/database"
	"Quiz-3/library"
	"Quiz-3/repository"
	"Quiz-3/structs"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	book, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book
	var errorBook []library.Error

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	if book.ReleaseYear < 1980 {
		result := library.Error{"Tahun Rilis Minimal 1980"}.Validate()
		errorBook = append(errorBook, result)
	} else if book.ReleaseYear > 2021 {
		result := library.Error{"Tahun Rilis Maximal 2021"}.Validate()
		errorBook = append(errorBook, result)
	}

	if _, err := url.ParseRequestURI(book.ImageURL); err != nil {
		result := library.Error{"Gambar Harus Berupa URL"}.Validate()
		errorBook = append(errorBook, result)
	}

	if len(errorBook) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Error":  errorBook,
		})
		return
	}

	if book.TotalPage <= 100 {
		book.Thickness = "tipis"
	} else if book.TotalPage < 200 && book.TotalPage > 100 {
		book.Thickness = "sedang"
	} else if book.TotalPage >= 201 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tidak ada"
	}

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert book",
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book
	var errorBook []library.Error

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = int(id)

	if book.ReleaseYear < 1980 {
		result := library.Error{"Tahun Rilis Minimal 1980"}.Validate()
		errorBook = append(errorBook, result)
	} else if book.ReleaseYear > 2021 {
		result := library.Error{"Tahun Rilis Maximal 2021"}.Validate()
		errorBook = append(errorBook, result)
	}

	if _, err := url.ParseRequestURI(book.ImageURL); err != nil {
		result := library.Error{"Gambar Harus Berupa URL"}.Validate()
		errorBook = append(errorBook, result)
	}

	if len(errorBook) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"Error":  errorBook,
		})
		return
	}

	if book.TotalPage <= 100 {
		book.Thickness = "tipis"
	} else if book.TotalPage < 200 && book.TotalPage > 100 {
		book.Thickness = "sedang"
	} else if book.TotalPage >= 201 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tidak ada"
	}

	err = repository.UpdateBook(database.DbConnection, book)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update book",
	})
}

func DeleteBook(c *gin.Context) {
	var book structs.Book

	id, err := strconv.Atoi(c.Param("id"))

	book.ID = int(id)

	err = repository.DeleteBook(database.DbConnection, book)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete book",
	})
}
