package resolvers

import (
	"redis"

	"github.com/graphql-go/graphql"
)

// User represents a user in the system
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RootResolver represents the root GraphQL resolver
var RootResolver = graphql.Fields{
	"getUser": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id":    &graphql.Field{Type: graphql.String},
				"name":  &graphql.Field{Type: graphql.String},
				"email": &graphql.Field{Type: graphql.String},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(string)
			user, err := redis.GetUser(id)
			if err != nil {
				return nil, err
			}
			return user, nil
		},
	},
	// Implement other resolvers (getUsers, createUser, updateUser, deleteUser)
}

// Implement other resolvers for CRUD operations
