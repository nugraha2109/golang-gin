package productcontroller

import (
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Cihuy(c *gin.Context) {
	var barangs models.Barangs
	// var categories models.Categories

	if err := c.ShouldBindJSON(&barangs); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// if err := models.DB.SetupJoinTable(&models.Barangs{}, "category_id", &models.Categories{}); err != nil {
	// 	println(err.Error())
	// 	panic("Failed to setup join table")
	// 	return
	// }

	models.DB.Select("categories.*", "barangs.*").Create(&barangs)
	c.JSON(http.StatusOK, gin.H{"Product": barangs})
}

func Tambah(c *gin.Context) {
	//Handler untuk menambahkan data pengguna dan alamat
	var requestData struct {
		ProductName   string `json:"product_name" binding:"required"`
		Unit          string `json:"unit" binding:"required"`
		Price         string `json:"price" binding:"required"`
		Category_name string `json:"category_name" binding:"required"`
		Description   string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//Validasi dan tambahkan data ke dalam database
	barangs := models.Barangs{ProductName: requestData.ProductName, Unit: requestData.Unit, Price: requestData.Price}
	categories := models.Categories{Category_name: requestData.Category_name, Description: requestData.Description}

	// Mulai transaksi
	tx := models.DB.Begin()
	if tx.Error != nil {
		c.JSON(500, gin.H{"error": "failed to begin transaction"})
		return
	}

	// Tambahkan data barang
	if err := tx.Create(&barangs).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "failed to create"})
		return
	}

	// Tambahkan data categories yang berelasi
	categories.Category_id = barangs.Id
	if err := tx.Create(&categories).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "failed to create"})
		return
	}

	// Commit transaksi
	if err := tx.Commit().Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to commit transaction"})
		return
	}
	c.JSON(200, gin.H{"message": "data added successfully"})

}

func Ambil(c *gin.Context) {
	var barangs []models.Barangs
	models.DB.Preload("Barangs").Find(&barangs)

	// models.DB.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"products": barangs})
}

func Barang(c *gin.Context) {
	// Dapatkan ID pengguna dari URL
	var barangs models.Barangs

	// Dapatkan data pengguna dari database
	if err := models.DB.Preload("Barangs").Find(&barangs).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": barangs})

}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengubah data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diubah"})
}
func Cuy(c *gin.Context) {

	var barangs models.Barangs
	id := c.Param("id")
	if err := models.DB.Preload("Barangs").First(&barangs, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": barangs})
}

func Lampu(c *gin.Context) {
	var barangs models.Barangs
	id := c.Param("id")

	// Temukan barang yang akan diperbarui
	if err := models.DB.Preload("Barangs").First(&barangs, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Barang tidak ditemukan"})
		return
	}

	// Bind data JSON ke struct barang
	if err := c.ShouldBindJSON(&barangs); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Mulai transaksi
	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Perbarui data barang
	if err := tx.Save(&barangs).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui barang"})
		return
	}
	// Perbarui data relasi
	if err := tx.Save(&barangs.Barangs).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui data relasi"})
		return
	}

	//commit transaksi
	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan perubahan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diubah"})
}
