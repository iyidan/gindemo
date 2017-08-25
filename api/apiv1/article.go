package apiv1

import (
	"strings"

	"github.com/iyidan/goutils/mise"

	"github.com/gin-gonic/gin"

	"github.com/iyidan/gindemo/models"
	"github.com/iyidan/gindemo/util"
)

// RegisterArticle register the article api
func RegisterArticle(group *gin.RouterGroup) {

	// get article detail
	group.GET("/article/:id", func(c *gin.Context) {
		id, err := mise.ParseInt(c.Param("id"))
		if id <= 0 || err != nil {
			util.APIParamError(c, "id param error")
			return
		}
		article := &models.Article{}
		if err := models.DB.First(article, id).Error; err != nil {
			if models.IsNotExists(err) {
				util.APIParamError(c, err.Error())
			} else {
				util.APISystemError(c, err.Error())
			}
			return
		}
		util.APISucc(c, article)
	})

	// set article
	group.POST("/article/:id", func(c *gin.Context) {
		id, err := mise.ParseInt(c.Param("id"))
		if id <= 0 || err != nil {
			util.APIParamError(c, "id param error")
			return
		}
		content := strings.TrimSpace(c.PostForm("content"))
		title := strings.TrimSpace(c.PostForm("title"))
		if len(content) == 0 || len(title) == 0 {
			util.APIParamError(c, "title/content param empty")
			return
		}

		updateArt := models.Article{ID: id, Title: title, Content: content}
		md := models.DB.Select("title", "content", "updated_at").Save(&updateArt)
		rowsAffected := md.RowsAffected
		err = md.Error
		if err != nil {
			util.APISystemError(c, err.Error())
			return
		}
		if rowsAffected == 0 {
			util.APIParamError(c, "article not exists")
			return
		}
		util.APISucc(c, "success")
	})

	// create a new article
	group.PUT("/article", func(c *gin.Context) {
		article := &models.Article{}
		err := c.BindJSON(article)
		if err != nil {
			util.APIParamError(c, err.Error())
			return
		}
		article.ID = 0
		article.Status = 1
		if err := models.DB.Save(article).Error; err != nil {
			util.APISystemError(c, err.Error())
			return
		}
		util.APISucc(c, article.ID)
	})
}
