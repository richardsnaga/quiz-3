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

	port, err := strconv.Atoi(os.Getenv("PGPORT"))
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		port,
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

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

	authMiddleware := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}))

	// Route Bangun Datar
	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.Segitiga)
	router.GET("/bangun-datar/jajar-genjang", controllers.JajarGenjang)
	router.GET("/bangun-datar/persegi", controllers.Persegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.Lingkaran)

	// category
	router.GET("/categories", controllers.GetAllCategory)
	authMiddleware.POST("/categories", controllers.InsertCategory)
	authMiddleware.PUT("/categories/:id", controllers.UpdateCategory)
	authMiddleware.DELETE("/categories/:id", controllers.DeleteCategory)

	// book
	router.GET("/books", controllers.GetAllBook)
	authMiddleware.POST("/books", controllers.InsertBook)
	authMiddleware.PUT("/books/:id", controllers.UpdateBook)
	authMiddleware.DELETE("/books/:id", controllers.DeleteBook)
	router.GET("/categories/:id/books", controllers.GetBookCategory)

	router.Run(":" + os.Getenv("PORT"))
}
