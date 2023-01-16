package main

import (
	"Quiz-3/controllers"
	"Quiz-3/database"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed load file environment")
	} else {
		fmt.Println("Success read file environment")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Database Connection Failed")
		panic(err)
	} else {
		fmt.Println("Database Connected")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()

	// Route Bangun Datar
	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.Segitiga)
	router.GET("/bangun-datar/jajar-genjang", controllers.JajarGenjang)
	router.GET("/bangun-datar/persegi", controllers.Persegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.Lingkaran)

	// category
	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", controllers.InsertCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)

	router.Run("localhost:8090")
}
