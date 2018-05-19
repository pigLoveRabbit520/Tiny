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
	cates, errMeta := models.GetAllCategories()
	if err != nil {
		respond(c, DB_ERR_CODE, "数据库错误："+err.Error(), nil)
	} else if errMeta != nil {
		respond(c, DB_ERR_CODE, "数据库错误："+errMeta.Error(), nil)
	} else {
		respond(c, 0, "success", gin.H{
			"posts": posts,
			"total": count,
			"cates": cates,
		})
	}
}

func GetPages(c *gin.Context) {
	page := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	pages, count, err := models.GetPages(page)
	if err != nil {
		respond(c, DB_ERR_CODE, "数据库错误："+err.Error(), nil)
	} else {
		respond(c, 0, "success", gin.H{
			"pages": pages,
			"total": count,
		})
	}
}
