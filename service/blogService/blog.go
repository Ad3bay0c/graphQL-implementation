package blogService

import (
	"errors"
	domain "github.com/Ad3bay0c/graphqlTesting/domain/blogRepository"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

type Blog interface {
	Query(query string) (interface{}, error)
}

type DefaultBlogService struct {
	repo domain.BlogRepository
}

func NewBlogService(repo domain.BlogRepository) *DefaultBlogService {
	return &DefaultBlogService{
		repo: repo,
	}
}
func (s *DefaultBlogService) mutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "Blog Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type: domain.BlogType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					blog := domain.Blog{
						ID:     uuid.NewString(),
						Title:  p.Args["title"].(string),
						Author: p.Args["author"].(string),
					}
					return s.repo.CreateBlog(blog)
				},
			},
			"update": &graphql.Field{
				Type: domain.BlogType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					blog := domain.Blog{
						ID:     p.Args["id"].(string),
						Title:  p.Args["title"].(string),
						Author: p.Args["author"].(string),
					}
					return s.repo.UpdateBlog(blog)
				},
			},
		},
	})
}

func (s *DefaultBlogService) rootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"blogs": &graphql.Field{
				Type: graphql.NewList(domain.BlogType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return s.repo.GetAll()
				},
			},
			"blog": &graphql.Field{
				Type: domain.BlogType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return s.repo.GetByID(p.Args["id"].(string))
				},
			},
		},
	})
}
func (s *DefaultBlogService) Query(query string) (interface{}, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    s.rootQuery(),
		Mutation: s.mutation(),
	})
	if err != nil {
		return nil, errors.New("error creating schema")
	}
	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	result := graphql.Do(params)

	return result, err
}
