package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	// cegah koneksi berulang
	if DB != nil {
		return
	}

	// load .env
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println("⚠️ .env tidak ditemukan di current maupun parent directory")
		}
	}

	// ambil DSN
	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan. Pastikan .env berisi DSN Supabase")
	}

	// konfigurasi logger (BIAR TIDAK BERISIK 🔥)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Silent, // 🔥 ini yang bikin bersih
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	// koneksi DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Gagal konek ke database: %v", err)
	}

	DB = db
	fmt.Println("✅ Koneksi ke PostgreSQL (Supabase) BERHASIL")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("DB belum diinisialisasi. Panggil config.InitDB() lebih dulu.")
	}
	return DB
}
