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
		"message": users,
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
