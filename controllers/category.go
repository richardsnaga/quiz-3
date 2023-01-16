package controllers

import (
	"Quiz-3/database"
	"Quiz-3/repository"
	"Quiz-3/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	category, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.CreateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert category",
	})
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = int(id)

	err = repository.UpdateCategory(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update category",
	})
}
