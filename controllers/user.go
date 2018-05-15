package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/salamander-mh/SalamanderBlog/models"
	"fmt"
	"time"
	"github.com/salamander-mh/SalamanderBlog/utils"
	"github.com/salamander-mh/SalamanderBlog/config"
)

const (
	TOKEN_EXPIRE_TIME = 7 * 86400 // 七天有效
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if len(username) <= 0 || len(password) <= 0 {
		c.JSON(200, ApiRes{
			Errcode: 1,
			Errmsg:  "账号或者密码不能为空！",
		})
		return
	}
	user, err := models.LoginUser(username, password)
	if err != nil {
		c.JSON(200, ApiRes{
			Errcode: 1,
			Errmsg:  "账号或者密码错误！",
		})
		return
	}
	token, err := utils.Encrypt(fmt.Sprintf("%d:%s", user.Uid, time.Now().Unix() + TOKEN_EXPIRE_TIME),
		[]byte(config.Conf.Key))
	c.JSON(200, ApiRes{
		Errcode: 0,
		Errmsg:  "登录成功",
		Data:  gin.H{
			"token": token,
		},
	})
}
