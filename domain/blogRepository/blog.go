package blog

type Blog struct {
	ID     int
	Title  string
	Author string
}

type BlogRepository interface {
	GetAll() []Blog
}
