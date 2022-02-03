package blogRepository

import "github.com/graphql-go/graphql"

type Blog struct {
	ID     int
	Title  string
	Author string
}

type BlogRepository interface {
	GetAll() ([]Blog, error)
	GetByID(id int) (*Blog, error)
}

var BlogType = graphql.NewObject(graphql.ObjectConfig{
	Name: "blog",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
})
