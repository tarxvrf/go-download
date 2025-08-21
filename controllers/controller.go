package controllers

import (
	"go-download/databases"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ini halaman index",
	})

}

func Download(c *gin.Context) {
	paramd := c.Param("namafile")
	filepat := "./files/" + paramd
	filename := filepath.Base(filepat) // ambil nama file
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filepat)

}

func User(c *gin.Context) {
	var users []databases.User
	err := databases.DB.Find(&users).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": users,
	})

}

func Kirim(c *gin.Context) {
	var user databases.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}
	databases.DB.Save(&user)
	c.JSON(200, gin.H{
		"message": "data berhasil disimpan",
	})

}

func Hapus(c *gin.Context) {
	var user databases.User
	id := c.Param("id")
	ceking := databases.DB.First(&user, id).Error
	if ceking != nil {
		c.JSON(500, gin.H{
			"message": "data tidak ditemukan",
		})
	}
	databases.DB.Delete(&user, id)
	c.JSON(200, gin.H{
		"message": "data berhasil dihapus",
	})

}

// file: controllers/user.go
func Edit(c *gin.Context) {
	var user databases.User // struct dari DB
	id := c.Param("id")

	// 1. Cari data berdasarkan ID
	if err := databases.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	// 2. Ambil input dari body JSON
	var input struct {
		Name       string `json:"name"`
		Lokasi     string `json:"lokasi"`
		Telpon     string `json:"telepon"`
		Picgedung  string `json:"picgedung"`
		Tanggal    string `json:"tanggal"`
		Status     string `json:"status"`
		Foto1      string `json:"foto1"`
		Foto2      string `json:"foto2"`
		Foto3      string `json:"foto3"`
		Keterangan string `json:"keterangan"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"message": "Input tidak valid"})
		return
	}

	// 3. Update data
	if err := databases.DB.Model(&user).Updates(input).Error; err != nil {
		c.JSON(500, gin.H{"message": "Gagal update data"})
		return
	}

	// 4. Beri response
	c.JSON(200, gin.H{"message": "Data berhasil diupdate", "data": user})
}
