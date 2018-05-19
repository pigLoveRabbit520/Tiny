package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/salamander-mh/SalamanderBlog/config"
	"github.com/salamander-mh/SalamanderBlog/models"
	"github.com/salamander-mh/SalamanderBlog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ApiRes struct {
	Errcode uint        `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

const (
	AUTH_HEADER_NAME = "X-Auth-Token"
	USER_ID_KEY      = "uid"
	USER_DATA_KEY    = "udata"
	DB_ERR_CODE      = 258
)

func respond(c *gin.Context, errCode uint, errmsg string, data interface{}) {
	c.JSON(http.StatusOK, ApiRes{
		Errcode: errCode,
		Errmsg:  errmsg,
		Data:    data,
	})
}

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(AUTH_HEADER_NAME)
		if len(token) > 0 {
			uData, err := utils.Decrypt(token, []byte(config.Conf.EncryptKey))
			if err == nil {
				parts := strings.Split(uData, ":")
				if len(parts) == 2 {
					expireTime, err := strconv.ParseInt(parts[1], 10, 64)
					uid, _ := strconv.Atoi(parts[0])
					if err == nil && expireTime > time.Now().Unix() && uid > 0 {
						uinfo := models.GetUser(uid)
						c.Set(USER_ID_KEY, uid)
						c.Set(USER_DATA_KEY, uinfo)
						return
					}
				}
			}
		}
		c.Abort()
		c.JSON(http.StatusUnauthorized, ApiRes{
			Errcode: 1,
			Errmsg:  "用户未登录！",
		})
		return
	}
}
