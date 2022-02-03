package blogRepository

import (
	"github.com/jmoiron/sqlx"
)

type blogDB struct {
	DB *sqlx.DB
}

func NewBlogDB(db *sqlx.DB) *blogDB {
	return &blogDB{DB: db}
}

func (db *blogDB) GetAll() ([]Blog, error) {
	var blog []Blog
	err := db.DB.Select(&blog, "SELECT * FROM blog")
	return blog, err
}

func (db *blogDB) GetByID(id string) (*Blog, error) {
	var blog Blog
	err := db.DB.Get(&blog, "SELECT * FROM blog WHERE id=?", id)
	return &blog, err
}

func (db *blogDB) CreateBlog(blog Blog) (*Blog, error) {
	_, err := db.DB.Exec("INSERT INTO blog(id, title, author) VALUES($1, $2, $3)", blog.ID, blog.Title, blog.Author)
	return &blog, err
}
