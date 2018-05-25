package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/salamander-mh/SalamanderBlog/models"
	"strconv"
)

func GetComments(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	comments, count, err := models.GetComments(page)
	if err != nil {
		respond(c, DB_ERR_CODE, "数据库错误："+err.Error(), nil)
	} else {
		respond(c, 0, "success", gin.H{
			"comments": comments,
			"total":    count,
		})
	}
}