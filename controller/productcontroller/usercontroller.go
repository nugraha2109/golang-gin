package productcontroller

import (
	"crud/models"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GenerateAPIKey() (string, error) {
	// Buat byte array dengan panjang yang sesuai
	keyLength := 32 // Panjang key dalam byte
	keyBytes := make([]byte, keyLength)

	// Baca randomness dari crypto/rand package
	_, err := rand.Read(keyBytes)
	if err != nil {
		return "", err
	}

	// Konversi byte array menjadi string heksadesimal
	apiKey := hex.EncodeToString(keyBytes)
	return apiKey, nil
}

func RegisterController(c *gin.Context) {
	// Terima input dari form

	var formData struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek kondisi apakah harus menghasilkan API key baru atau tidak
	generateAPIKey := true
	// Misalnya, jika username mengandung kata "admin", maka tidak perlu membuat API key
	var existingUser models.Users
	if err := models.DB.Where("name = ? OR email = ?", formData.Name, formData.Email).First(&existingUser).Error; err == nil {
		// Jika ditemukan pengguna dengan nama atau email yang sama, kirimkan pesan kesalahan
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username atau Email sudah digunakan"})
		return
	}

	// Buat instance user baru
	user := models.Users{
		Name:     formData.Name,
		Email:    formData.Email,
		Password: formData.Password,
	}

	// Jika generateAPIKey true, maka buat dan set API key
	if generateAPIKey {
		apiKey, err := GenerateAPIKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat API key"})
			return
		}
		user.APIkey = apiKey
	}

	// Simpan pengguna ke database
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan pengguna ke database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil didaftarkan", "user": user})
}

func Getdata(c *gin.Context) {
	var user []models.Users

	if err := models.DB.Find(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": user})
}
