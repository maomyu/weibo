package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/basic/logger"
	"github.com/yuwe1/weibo/models"
	"github.com/yuwe1/weibo/pkg"
)

func LoginGet(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(pkg.SESSION_KEY) != nil {
		c.Redirect(http.StatusSeeOther, "/index")
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "微博登录",
	})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	logger.Sugar.Info(username)
	// 根据姓名查找用户
	var user models.User
	user = models.FindUserByName(username)
	var flag bool
	if user.Username == "" {
		flag = false
	} else if user.Password == password {
		flag = true
	} else {
		flag = false
	}
	//保存用户信息
	if flag {
		session := sessions.Default(c)
		v := session.Get(pkg.SESSION_KEY)
		if v == nil {
			session.Set(pkg.SESSION_KEY, user.Userid)
			session.Save()
		}
		fmt.Println("session:", session.Get(pkg.SESSION_KEY))
	}

	c.JSON(http.StatusOK, gin.H{
		"flag": flag,
	})
}
func LogoutGet(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}
