package routers

import (
	"go-download/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

func Routerx() {
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.GET("/download/:namafile", controllers.Download)
	router.GET("/", controllers.Index)
	router.Run(":" + port)

}
