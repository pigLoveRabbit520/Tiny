package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/salamander-mh/SalamanderBlog/db"
	"github.com/salamander-mh/SalamanderBlog/models"
	"github.com/gin-gonic/gin"
	"github.com/salamander-mh/SalamanderBlog/controllers"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/salamander-mh/SalamanderBlog/config"
	"fmt"
)

func main() {
	conf := initConfig()
	db.ConnectAndInit(conf, new(models.User), new(models.Content))

	router := gin.Default()

	adminRouter := router.Group("/admin")
	{
		adminRouter.GET("/login", controllers.Login)

		adminRouter.GET("/contents", controllers.GetContents)
	}

	router.Run(fmt.Sprintf(":%d", conf.Port))
}

func initConfig() config.Config {
	var conf config.Config
	source, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}