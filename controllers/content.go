package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/salamander-mh/SalamanderBlog/db"
	"github.com/salamander-mh/SalamanderBlog/models"
	"net/http"
)

func GetContents(c *gin.Context) {
	var contents []models.Content
	db.DB.Select("title, `order`").Order("created desc").Find(&contents)
	c.JSON(http.StatusOK, contents)
}
