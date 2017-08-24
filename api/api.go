package api

import (
	"github.com/iyidan/gindemo/api/apiv1"
	"github.com/iyidan/gindemo/middleware"

	"github.com/gin-gonic/gin"
)

// Register apis
func Register(router *gin.Engine) {
	// api sign
	group := router.Group("/api", middleware.APISign())
	apiv1.Register(group)
}
