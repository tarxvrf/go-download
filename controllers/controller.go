package controllers

import (
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
