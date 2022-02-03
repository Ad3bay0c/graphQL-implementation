package blog

import domain "github.com/Ad3bay0c/graphqlTesting/domain/blog"

type Blog interface {
	List() []domain.Blog
}

type DefaultBlogService struct {
	repo domain.BlogRepository
}

func NewBlogService(repo domain.BlogRepository) *DefaultBlogService {
	return &DefaultBlogService{
		repo: repo,
	}
}

func (s *DefaultBlogService) List() []domain.Blog {
	return s.repo.GetAll()
}
