package postgres

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	dbConnectionErr error
)

func Open() (*gorm.DB) {
	readEnvErr := godotenv.Load()
	if readEnvErr != nil {
		print(".envが読めなかった", readEnvErr)
	}
	DB_HOST := "host=" + os.Getenv("DB_HOST")
	DB_USER := " user=" + os.Getenv("DB_USER")
	DB_PASSWORD := " password=" + os.Getenv("DB_PASSWORD")
	DB_NAME := " dbname=" + os.Getenv("DB_NAME")
	DB_PORT := " port=" + os.Getenv("DB_PORT")
	DB_SSLMODE := " sslmode=" + os.Getenv("DB_SSLMODE")
	DB_TIMEZONE := " Timezone=" + os.Getenv("DB_TIMEZONE")

	dsn := DB_HOST + DB_USER + DB_PASSWORD + DB_NAME + DB_PORT + DB_SSLMODE + DB_TIMEZONE
	db, dbConnectionErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbConnectionErr != nil {
		print("DB接続失敗", dbConnectionErr)
	}
	return db
}

func Close() {
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
}
