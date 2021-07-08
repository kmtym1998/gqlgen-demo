package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Open() (*gorm.DB) {
	// TODO: 環境変数を使う
	dsn := "host=localhost user=hasura password=secret dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	return db
}

func Close() {
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
}
