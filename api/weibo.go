package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuwe1/weibo/models"
)

// 发微博
func WeiboPost(c *gin.Context) {
	var weibo models.Weibo
	weibo.Text = c.PostForm("text")
	weibo.Userid, _ = strconv.Atoi(c.PostForm("userid"))
	weibo.Username = c.PostForm("username")
	fmt.Println(c.PostForm("userid"))
	fmt.Println(c.PostForm("text"))
	weibo.CreatedAt = time.Now()
	weibo.Like = 0
	weibo.CommentCount = 0
	flag := models.AddWeibo(weibo)

	c.JSON(http.StatusOK, gin.H{
		"flag": flag,
	})
}
