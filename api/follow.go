package api

import (
	"net/http"
	"strconv"

	"github.com/yuwe1/weibo/models"

	"github.com/gin-gonic/gin"
)

//增加关注
func Follow(c *gin.Context) {
	userid, _ := strconv.Atoi(c.PostForm("userid"))
	followid, _ := strconv.Atoi(c.PostForm("followid"))
	c.JSON(http.StatusOK, gin.H{
		"flag": models.AddFollow(userid, followid),
	})
}

//取消关注
func Unfollow(c *gin.Context) {
	userid, _ := strconv.Atoi(c.PostForm("userid"))
	followid, _ := strconv.Atoi(c.PostForm("followid"))
	c.JSON(http.StatusOK, gin.H{
		"flag": models.DeleteFollow(userid, followid),
	})
}

func CountFollow(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Query("userid"))
	cf, cfd := models.CountFollow(userid)
	c.JSON(http.StatusOK, gin.H{
		"CountFollow":   cf,
		"CountFollowed": cfd,
	})
}
