package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/weibo/internal/weibo"
	"github.com/yuwe1/weibo/models"
	"github.com/yuwe1/weibo/pkg"
)

func IndexGet(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(pkg.SESSION_KEY) == nil {
		c.Redirect(http.StatusSeeOther, "/login")
	}
	uid := session.Get(pkg.SESSION_KEY).(int)
	user := models.FindUserByID(uid)
	wb := weibo.GetAllIndexWeibo(uid)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"userid":     uid,
		"username":   user.Username,
		"user":       user,
		"weibo":      wb,
		"weibocount": models.FindWeiboCountByUserID(user.Userid),
	})
}
