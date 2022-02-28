package application

import (
	"github.com/Ad3bay0c/graphqlTesting/application/handler"
	"github.com/Ad3bay0c/graphqlTesting/domain/blogRepository"
	service "github.com/Ad3bay0c/graphqlTesting/service/blogService"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func defineRoutes(r *gin.Engine, db *sqlx.DB) {
	b := &handler.BlogHandler{Service: service.NewBlogService(blogRepository.NewBlogDB(db))}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.POST("/graphql", b.GraphQuery)
}
func SetupRouter() *gin.Engine {
	router := gin.New()
	db, err := sqlx.Connect("sqlite3", "./blog.db")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err.Error())
	}
	log.Printf("Database Connected Succesfully\n")

	defineRoutes(router, db)

	return router
}
