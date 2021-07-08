package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kmtym1998/gqlgen-demo/graph"
	"github.com/kmtym1998/gqlgen-demo/graph/generated"
)

type Sample struct {
	Id   int `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

const defaultPort = "8081"

func gormConnect() *gorm.DB {
	dsn := "host=localhost user=hasura password=secret dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("接続失敗", err)
	}

	return db
}

func main() {
	db := gormConnect()
	// defer db.Close()
	var samples []Sample
	err := db.Where("id", 1).Find(&samples).Error
	if err != nil {
		log.Fatalln("取得失敗", err)
	}

    fmt.Println(samples)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
