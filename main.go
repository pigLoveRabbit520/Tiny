package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/salamander-mh/SalamanderBlog/config"
	"github.com/salamander-mh/SalamanderBlog/controllers"
	"github.com/salamander-mh/SalamanderBlog/db"
	"github.com/salamander-mh/SalamanderBlog/models"
)

func main() {
	config.IniConfig()
	db.ConnectAndInit(config.Conf, new(models.User), new(models.Content))

	router := gin.Default()
	router.POST("/admin/login", controllers.Login)

	adminRouter := router.Group("/admin")
	{
		adminRouter.GET("/contents", controllers.GetContents)
	}

	router.Run(fmt.Sprintf(":%d", config.Conf.Port))
}
