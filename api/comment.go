package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/weibo/models"
	"github.com/yuwe1/weibo/pkg"
)

// 发评论
func CommentPost(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(pkg.SESSION_KEY) == nil {
		c.Redirect(http.StatusSeeOther, "/login")
	}
	var comment models.Comment
	comment.CommentText = c.PostForm("text")
	if comment.CommentText == "" {
		c.JSON(http.StatusOK, gin.H{
			"flag":         false,
			"errormessage": "评论不可为空",
		})
		return
	}
	uid := session.Get(pkg.SESSION_KEY).(int)
	user := models.FindUserByID(uid)
	comment.Userid = uid
	comment.Username = user.Username
	comment.Weiboid, _ = strconv.Atoi(c.PostForm("weiboid"))
	comment.CreatedAt = time.Now()

	fmt.Println(comment)
	flag := models.AddComment(comment)

	c.JSON(http.StatusOK, gin.H{
		"flag": flag,
	})
}
