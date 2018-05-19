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
	db.ConnectAndInit(
		config.Conf,
		new(models.User),
		new(models.Content),
		new(models.Comment),
		new(models.Meta),
		new(models.Option),
		new(models.Relationship),
	)

	router := gin.Default()
	// recovers from any panics and writes a 500 if there was one
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/admin/login", controllers.Login)

	adminRouter := router.Group("/admin")
	adminRouter.Use(controllers.UserAuth())
	{
		adminRouter.GET("/posts", controllers.GetContents)
		adminRouter.GET("/user/info", controllers.GetUserInfo)
	}

	router.Run(fmt.Sprintf(":%d", config.Conf.Port))
}
