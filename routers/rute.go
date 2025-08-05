package routers

import (
	"go-download/controllers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Routerx() {
	godotenv.Load(".env")
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	router.GET("/download/:namafile", controllers.Download)
	router.GET("/", controllers.Index)
	router.POST("/kirim", controllers.Kirim)
	router.DELETE("/hapus/:id", controllers.Hapus)
	router.GET("/users", controllers.User)
	router.Run(":" + port)

}
