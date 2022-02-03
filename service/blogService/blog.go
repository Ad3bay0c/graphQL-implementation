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
			"list": &graphql.Field{
				Type: graphql.NewList(domain.BlogType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return []domain.Blog{{ID: 1, Title: "First result"}, {ID: 2, Title: "Second Result"}}, nil
				},
			},
		},
	})
}
func (s *DefaultBlogService) Query(query string) (interface{}, error) {
	rootQuery := s.rootQuery()
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
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
