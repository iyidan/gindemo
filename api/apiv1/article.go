package apiv1

import (
	"github.com/iyidan/gindemo/util"

	"github.com/gin-gonic/gin"
)

// RegisterArticle register the article api
func RegisterArticle(group *gin.RouterGroup) {

	// get article detail
	group.GET("/article/:id", func(c *gin.Context) {
		util.APISucc(c, "hello, @todo:"+c.Param("id"))
	})

	// set article detail
	group.POST("/article/:id", func(c *gin.Context) {
		id := c.Param("id")
		content := c.PostForm("content")
		util.APISucc(c, "hello, @todo:"+id+" "+content)
	})
}
