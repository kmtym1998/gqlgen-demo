package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/kmtym1998/gqlgen-demo/graph"
	"github.com/kmtym1998/gqlgen-demo/graph/generated"
)

type Sample struct {
	Id   int `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

const defaultPort = "8081"

func main() {
	readEnvErr := godotenv.Load()
	if readEnvErr != nil {
		print(".envが読めなかった", readEnvErr)
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
