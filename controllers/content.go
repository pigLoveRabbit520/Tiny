package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/salamander-mh/SalamanderBlog/models"
)

func GetContents(c *gin.Context) {
	page := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	posts, count, err := models.GetPosts(page)
	if err != nil {
		respond(c, DB_ERR_CODE, "数据库错误："+err.Error(), nil)
	} else {
		respond(c, 0, "success", gin.H{
			"posts":      posts,
			"total": count,
		})
	}
}
