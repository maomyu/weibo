package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/weibo/internal/weibo"
	"github.com/yuwe1/weibo/models"
	"github.com/yuwe1/weibo/pkg"
)

func MyGet(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(pkg.SESSION_KEY) == nil {
		c.Redirect(http.StatusSeeOther, "/login")
	}
	uid := session.Get(pkg.SESSION_KEY).(int)
	user := models.FindUserByID(uid)
	fmt.Println(uid)
	c.HTML(http.StatusOK, "my.html", gin.H{
		"userid":     uid,
		"username":   user.Username,
		"user":       user,
		"weibo":      weibo.GetAllWeibo(user.Userid),
		"weibocount": models.FindWeiboCountByUserID(user.Userid),
	})
}
func MyAlternativeGet(c *gin.Context) {
	username := c.Param("username")
	user := models.FindUserByName(username)
	user.Password = "***"
	wb := weibo.GetAllWeibo(user.Userid)
	c.HTML(http.StatusOK, "my.html", gin.H{
		"userid":     user.Userid,
		"username":   username,
		"user":       user,
		"weibo":      wb,
		"weibocount": models.FindWeiboCountByUserID(user.Userid),
	})
}
