package apiv1

import "github.com/gin-gonic/gin"

// Register api
func Register(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	RegisterArticle(v1)
}
