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

func (db *blogDB) GetByID(id int) (*Blog, error) {
	var blog Blog
	err := db.DB.Get(&blog, "SELECT * FROM blog WHERE id=?", id)
	return &blog, err
}
