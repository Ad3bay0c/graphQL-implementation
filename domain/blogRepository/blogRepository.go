package blogRepository

import "database/sql"

type blogDB struct {
	DB *sql.DB
}

func NewBlogDB(db *sql.DB) *blogDB {
	return &blogDB{DB: db}
}

func (db *blogDB) GetAll() ([]Blog, error) {
	return []Blog{}, nil
}
