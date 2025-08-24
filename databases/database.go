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
	ID           int    `gorm:"primaryKey" json:"id"`
	Marketing    string `gorm:"unique" json:"marketing"`
	Namalokasi   string `gorm:"unique" json:"namalokasi"`
	Alamatlokasi string `json:"alamatlokasi"`
	Picgedung    string `json:"picgedung"`
	Tanggal      string `json:"tanggal"`
	Status       string `json:"status"`
	Operator     string `json:"operator"`
	Sistemparkir string `json:"sistemparkir"`
	Pk           string `json:"pk"`
	Pm           string `json:"pm"`
	Fu1          string `json:"fu1"`
	Fu2          string `json:"fu2"`
	Fu3          string `json:"fu3"`
	Kondisi      string `json:"kondisi"`
	Kontrak      string `json:"kontrak"`
	Telpon       string `json:"telepon"`
	Foto1        string `json:"foto1"`
	Foto2        string `json:"foto2"`
	Foto3        string `json:"foto3"`
	Keterangan   string `json:"keterangan"`
}

func Init() {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}

}

func Konekdb() {
	//Init()

	//dsn := os.Getenv("DATABASE_URL")
	dsn := "host=localhost user=postgres password=123456 dbname=toko port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gagal koneksi ke database")
		return
	}
	fmt.Println("sukses konek ke DB")
	db.AutoMigrate(&User{})
	DB = db

}
