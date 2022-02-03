package blogService

import (
	"errors"
	domain "github.com/Ad3bay0c/graphqlTesting/domain/blogRepository"
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
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return s.repo.GetByID(p.Args["id"].(int))
				},
			},
		},
	})
}
func (s *DefaultBlogService) Query(query string) (interface{}, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: s.rootQuery(),
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
