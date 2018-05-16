package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/salamander-mh/SalamanderBlog/config"
	"github.com/salamander-mh/SalamanderBlog/models"
	"github.com/salamander-mh/SalamanderBlog/utils"
	"time"
)

const (
	TOKEN_EXPIRE_TIME = 7 * 86400 // 七天有效
)

type LoginCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginCmd LoginCommand
	err := c.BindJSON(&loginCmd)
	if err != nil {
		c.JSON(200, ApiRes{
			Errcode: 1,
			Errmsg:  "登录数据格式不正确！",
		})
		return
	}

	if len(loginCmd.Username) <= 0 || len(loginCmd.Password) <= 0 {
		c.JSON(200, ApiRes{
			Errcode: 1,
			Errmsg:  "账号或者密码不能为空！",
		})
		return
	}
	user, err := models.LoginUser(loginCmd.Username, loginCmd.Password)
	if err != nil {
		c.JSON(200, ApiRes{
			Errcode: 1,
			Errmsg:  "账号或者密码错误！",
		})
		return
	}
	token, err := utils.Encrypt(fmt.Sprintf("%d:%d", user.Uid, time.Now().Unix()+TOKEN_EXPIRE_TIME),
		[]byte(config.Conf.EncryptKey))
	c.JSON(200, ApiRes{
		Errcode: 0,
		Errmsg:  "登录成功",
		Data: gin.H{
			"token": token,
		},
	})
}

func GetUserInfo(c *gin.Context) {
	uidStr, _ := c.Get(USER_ID_KEY)
	uid := uidStr.(int)
	uinfo := models.GetUser(uid)
	c.JSON(200, ApiRes{
		Errcode: 0,
		Errmsg:  "查询成功",
		Data:    uinfo,
	})
}
