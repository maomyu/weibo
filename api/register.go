package api

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/weibo/models"
	"github.com/yuwe1/weibo/pkg"
)

func RegisterGet(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(pkg.SESSION_KEY) != nil {
		c.Redirect(http.StatusSeeOther, "/index")
	}
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "微博注册",
	})
}

func RegisterPost(c *gin.Context) {
	var user models.User
	errorMessage := ""

	user.Username = c.PostForm("username")
	if models.FindUserDuplicate(user.Username) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"flag":         false,
			"errormessage": "用户名重复",
		})
		return
	}

	user.Password = c.PostForm("password")
	user.Sex, _ = strconv.Atoi(c.PostForm("sex"))
	user.Age, _ = strconv.Atoi(c.PostForm("age"))
	flag := models.AddUser(user)

	c.JSON(http.StatusOK, gin.H{
		"flag":         flag,
		"errormessage": errorMessage,
	})
}
