package databases

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size=60" json:"name"`
}

func Konekdb() {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal load file .env")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gagal koneksi ke database")
		return
	}
	fmt.Println("sukses konek ke DB")
	db.AutoMigrate(&User{})
	DB = db

}
