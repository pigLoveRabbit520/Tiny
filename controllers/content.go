package controllers

import (
	"github.com/salamander-mh/SalamanderBlog/models"
	"github.com/salamander-mh/SalamanderBlog/db"
)

type ContentController struct {
	BaseController
}

// @router / [get]
func (this *ContentController) GetContents() {
	defer this.ServeJSON()
	var contents []models.Content
	db.DB.Select("title, order").Order("created desc").Find(&contents)
	this.Data["json"] = contents
}