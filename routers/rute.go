package routers

import (
	"go-download/controllers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Routerx() {
	godotenv.Load()
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.GET("/download/:namafile", controllers.Download)
	router.GET("/", controllers.Index)
	router.POST("/kirim", controllers.Kirim)
	router.GET("/users", controllers.User)
	router.Run(":" + port)

}
