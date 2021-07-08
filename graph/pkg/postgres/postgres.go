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

// hirameのコードを拝借
// https://github.com/buysell-technologies/hirame_server/blob/c5f7eb0d1545cb07ab3951ec460b14f366eceaa9/lambda/pkg/postgres/postgres.go
func Open() (*gorm.DB) {
	// TODO: 環境変数を使う
	dsn := "host=localhost user=hasura password=secret dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	return db
}

// hirameのコードを拝借
// https://github.com/buysell-technologies/hirame_server/blob/c5f7eb0d1545cb07ab3951ec460b14f366eceaa9/lambda/pkg/postgres/postgres.go
func Close() {
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
}
