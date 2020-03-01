package weibo

import "github.com/yuwe1/weibo/models"

// 获取一个用户自己的微博
func GetAllWeibo(userid int) []models.Weibos {
	var weibo []models.Weibos
	weibo = models.FindAllWeiboByUserID(userid)
	return weibo
}

func GetAllIndexWeibo(userid int) []models.Weibos {
	var weibo []models.Weibos
	weibo = models.FindAllIndexWeibo(userid)
	//	fmt.Println(weibo)
	//	fmt.Println(query.Username)
	return weibo
}
