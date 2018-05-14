package controllers

import (
	"github.com/salamander-mh/SalamanderBlog/models"
	"github.com/salamander-mh/SalamanderBlog/db"
	"github.com/gin-gonic/gin"
	"net/http"
)


func GetContents(c *gin.Context) {
	var contents []models.Content
	db.DB.Select("title, `order`").Order("created desc").Find(&contents)
	c.JSON(http.StatusOK, contents)
}