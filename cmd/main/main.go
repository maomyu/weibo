package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/yuwe1/basic"
	"github.com/yuwe1/weibo/api"
)

func main() {
	basic.Init()

	fmt.Println("*********************")
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./views/**/*")

	setSessions(router)
	router.GET("/", api.IndexGet)
	router.GET("/index", api.IndexGet)
	router.GET("/my", api.MyGet)
	router.GET("/my/:username", api.MyAlternativeGet)
	router.GET("/login", api.LoginGet)
	router.GET("/register", api.RegisterGet)
	router.GET("/searchUser", api.SearchGet)
	router.GET("/navbar.html", api.NavbarGet)

	router.POST("/loginpost", api.LoginPost)
	router.POST("/registerpost", api.RegisterPost)
	router.POST("/weibopost", api.WeiboPost)
	router.POST("/commentpost", api.CommentPost)
	router.POST("/follow", api.Follow)
	router.POST("/unfollow", api.Unfollow)
	router.GET("/countfollow", api.CountFollow)
	router.GET("/logout", api.LogoutGet)

	router.Run(":9090")
}
func setSessions(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"}) //Also set Secure: true if using SSL, you should though
	router.Use(sessions.Sessions("gin-session", store))
}
