package main

import (
	"fmt"
	"net/http"
	"os"

	"go_crud/resolvers"

	// "your-project/graphql"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: resolvers.RootResolver}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: resolvers.RootResolver}),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
		return
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
