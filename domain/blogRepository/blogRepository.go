package blog

import "database/sql"

type blogDB struct {
	DB *sql.DB
}

func NewBlogDB(db *sql.DB) *blogDB {
	return &blogDB{DB: db}
}

func GetAll() []Blog {
	return []Blog{}
}
