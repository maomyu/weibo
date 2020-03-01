package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/weibo/models"
	"github.com/yuwe1/weibo/pkg"
)

func NavbarGet(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get(pkg.SESSION_KEY).(int)
	user := models.FindUserByID(userid)
	c.HTML(http.StatusOK, "navbar.html", gin.H{
		"user": user,
	})
}
