package main

import (
	"crud/controller/productcontroller"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	models.DB.AutoMigrate()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.GET("/api/ambil", productcontroller.Ambil)
	r.GET("/api/barang", productcontroller.Barang)
	r.POST("/api/tambah", productcontroller.Tambah)
	r.POST("/api/cihuy", productcontroller.Cihuy)
	r.GET("/api/cuy/:id", productcontroller.Cuy)
	r.PUT("/api/lampu/:id", productcontroller.Lampu)
	r.POST("/api/user", productcontroller.RegisterController)
	r.GET("/api/getdata", productcontroller.Getdata)

	// http.HandleFunc("/post", productcontroller.Lihat)

	// log.Print("Server starter on: http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	r.Run()
}
