package handler

import (
	blogService "github.com/Ad3bay0c/graphqlTesting/service/blogService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogHandler struct {
	Service blogService.Blog
}

func (blog *BlogHandler) GraphQuery(c *gin.Context) {
	input := &struct {
		Query string `json:"query"`
	}{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, _ := blog.Service.Query(input.Query)
	c.JSON(200, response)
}
