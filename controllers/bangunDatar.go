package controllers

import (
	"Quiz-3/library"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var BangunDatar library.HitungBangunDatar

func Segitiga(c *gin.Context) {
	var result int

	if c.Request.Method == "GET" {

		hitung := c.Query("hitung")
		alas, _ := strconv.Atoi(c.Query("alas"))
		tinggi, _ := strconv.Atoi(c.Query("tinggi"))

		BangunDatar = library.Segitiga{alas, tinggi}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		c.JSON(http.StatusOK, gin.H{

			"message": "Success",
			"result":  result,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{

		"message": "Method Salah",
	})
}

func PersegiPanjang(c *gin.Context) {
	var result int

	if c.Request.Method == "GET" {

		hitung := c.Query("hitung")
		panjang, _ := strconv.Atoi(c.Query("panjang"))
		lebar, _ := strconv.Atoi(c.Query("lebar"))

		BangunDatar = library.PersegiPanjang{panjang, lebar}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		c.JSON(http.StatusOK, gin.H{

			"message": "Success",
			"result":  result,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{

		"message": "Method Salah",
	})
}

func Persegi(c *gin.Context) {
	var result int

	if c.Request.Method == "GET" {

		hitung := c.Query("hitung")
		sisi, _ := strconv.Atoi(c.Query("sisi"))

		BangunDatar = library.Persegi{sisi}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		c.JSON(http.StatusOK, gin.H{

			"message": "Success",
			"result":  result,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{

		"message": "Method Salah",
	})
}

func JajarGenjang(c *gin.Context) {
	var result int

	if c.Request.Method == "GET" {

		hitung := c.Query("hitung")
		sisi, _ := strconv.Atoi(c.Query("sisi"))
		alas, _ := strconv.Atoi(c.Query("alas"))
		tinggi, _ := strconv.Atoi(c.Query("tinggi"))

		BangunDatar = library.JajarGenjang{float64(sisi), float64(alas), float64(tinggi)}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		c.JSON(http.StatusOK, gin.H{

			"message": "Success",
			"result":  result,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{

		"message": "Method Salah",
	})
}

func Lingkaran(c *gin.Context) {
	var result int

	if c.Request.Method == "GET" {

		hitung := c.Query("hitung")
		jari, _ := strconv.Atoi(c.Query("jari"))

		BangunDatar = library.Lingkaran{float64(jari)}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		c.JSON(http.StatusOK, gin.H{

			"message": "Success",
			"result":  result,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{

		"message": "Method Salah",
	})
}
