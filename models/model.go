package models

import (
	"fmt"
	"time"

	"github.com/yuwe1/basic/client/dbpool"
	"github.com/yuwe1/basic/logger"
)

type User struct {
	Userid   int
	Username string
	Password string
	Sex      int
	Age      int
}

type Comment struct {
	Commentid   int
	Weiboid     int
	Userid      int
	Username    string
	CreatedAt   time.Time
	CommentText string
}

type Weibo struct {
	Weiboid      int
	Userid       int
	Username     string
	CreatedAt    time.Time
	Text         string
	Like         int
	CommentCount int
}

type Weibos struct {
	Weibo
	Comment []Comment
}

type Follow struct {
	Userid   int
	Followid int
}

type SearchUser struct {
	User
	Relation int
	// Relation 0:无关系 1:被关注 2:关注 3:互相关注
}

// 根据name查询出用户
func FindUserByName(userid string) (user User) {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	q := "select * from user where username = ?"
	s.QueryRow(q, userid).Scan(
		&user.Userid,
		&user.Username,
		&user.Password,
		&user.Sex,
		&user.Age,
	)
	return user
}
func FindUserDuplicate(userid string) int {
	var count int
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	fmt.Println("username:", userid)
	q := "select count(*) from user where username = ?"
	s.QueryRow(q, userid).Scan(&count)
	fmt.Println("count:", count)
	return count
}

// 根据ID查询用户
func FindUserByID(userid int) (user User) {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	q := "select * from user where userid = ?"
	s.QueryRow(q, userid).Scan(
		&user.Userid,
		&user.Username,
		&user.Password,
		&user.Sex,
		&user.Age,
	)
	return user
}

// 添加用户
func AddUser(user User) bool {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	s.Begin()
	q1 := "SELECT max(userid) FROM user"
	result := 0
	s.QueryRow(q1).Scan(&result)
	user.Userid = result + 1
	q2 := "insert into user values(?,?,?,?,?)"
	if _, err = s.Exec(q2, user.Userid, user.Username, user.Password, user.Sex, user.Age); err != nil {
		s.Rollback()
		return false
	}
	s.Commit()
	return true
}

// 根据用户自己查找微博
func FindAllWeiboByUserID(userid int) []Weibos {
	var query []Weibo
	var ans []Weibos
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	q1 := "select *from weibo where userid = ? order by created_at desc"
	result, _ := s.Query(q1, userid)
	// 赋值
	for result.Next() {
		var weibo Weibo
		result.Scan(&weibo.Weiboid, &weibo.Userid,
			&weibo.Username, &weibo.CreatedAt, &weibo.Text, &weibo.Like, &weibo.CommentCount,
		)
		query = append(query, weibo)
	}
	// 查找评论
	for _, weibo := range query {
		var this Weibos
		this.Weibo = weibo
		q2 := "select * from comment where weiboid=? order by created_at desc"
		wr, _ := s.Query(q2, weibo.Weiboid)
		for wr.Next() {
			var c Comment
			wr.Scan(&c.Commentid, &c.Weiboid, &c.Userid, &c.Username, &c.CreatedAt, &c.CommentText)
			this.Comment = append(this.Comment, c)
		}
		ans = append(ans, this)
	}
	return ans
}

// 根据用户自己，查找微博的数量
func FindWeiboCountByUserID(userid int) int {
	var count int
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	q := "select count(*) from weibo where userid = ?"
	s.QueryRow(q, userid).Scan(&count)
	return count
}

// 根据用户自己和关注查找微博
func FindAllIndexWeibo(userid int) []Weibos {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	var query []Weibo
	var ans []Weibos
	q := "SELECT * FROM weibo WHERE userid = ?  or userid in " +
		"(SELECT followid FROM follow WHERE userid = ?) order by created_at desc"
	result, _ := s.Query(q, userid, userid)
	for result.Next() {
		var weibo Weibo
		result.Scan(&weibo.Weiboid, &weibo.Userid,
			&weibo.Username, &weibo.CreatedAt, &weibo.Text, &weibo.Like, &weibo.CommentCount,
		)
		query = append(query, weibo)
	}
	fmt.Println("weibo:", query)
	// 查找评论
	for _, weibo := range query {
		var this Weibos
		this.Weibo = weibo
		q2 := "select * from comment where weiboid=? order by created_at desc"
		wr, _ := s.Query(q2, weibo.Weiboid)
		for wr.Next() {
			var c Comment
			wr.Scan(&c.Commentid, &c.Weiboid, &c.Userid, &c.Username, &c.CreatedAt, &c.CommentText)
			this.Comment = append(this.Comment, c)
		}
		ans = append(ans, this)
	}
	fmt.Println(ans)
	return ans
}

// 添加微博
func AddWeibo(w Weibo) bool {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	s.Begin()
	insert := "insert into weibo values(?,?,?,?,?,?,?)"
	if _, err = s.Exec(insert, w.Weiboid, w.Userid, w.Username, w.CreatedAt, w.Text, w.Like, w.CommentCount); err != nil {
		s.Rollback()
		return false
	}
	s.Commit()
	return true
}

// 添加微博的评论
func AddComment(com Comment) bool {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	s.Begin()
	insert := "insert into comment values(?,?,?,?,?,?)"
	if _, err = s.Exec(insert, com.Commentid,
		com.Weiboid, com.Userid, com.Username,
		com.CreatedAt, com.CommentText); err != nil {
		s.Rollback()
		return false
	}
	// 增加对应微博的数量
	update := "update weibo set comment_count = comment_count +1 where weiboid = ?"
	if _, err = s.Exec(update, com.Weiboid); err != nil {
		s.Rollback()
		return false
	}
	s.Commit()
	return true
}

// 查找一个用户和其他用户的关系
// Relation属性取值 0:无关系 1:被关注 2:关注 3:互相关注
func FindRelation(userid int, users []User) []SearchUser {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	var result []SearchUser
	for _, user := range users {
		var this SearchUser
		this.User = user
		this.Relation = 0
		var count int
		q1 := "select count(*) from follow where userid = ? and followid = ?"
		s.QueryRow(q1, userid, user.Userid).Scan(&count)
		if count > 0 {
			this.Relation = this.Relation + 2
		}
		q2 := "select count(*) from follow where userid = ? and followid = ?"
		s.QueryRow(q2, user.Userid, userid).Scan(&count)
		if count > 0 {
			this.Relation = this.Relation + 1
		}
		result = append(result, this)
	}
	fmt.Println(result)
	return result
}

// 根据用户名查找用户和关系
func SearchUserByName(name string, userid int) []SearchUser {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	fmt.Println("***************")
	fmt.Println(name, userid)
	fmt.Println("*************")
	var users []User
	q := "select * from user where username LIKE ?"
	result, _ := s.Query(q, "%"+name+"%")
	for result.Next() {
		var user User
		result.Scan(&user.Userid, &user.Username, &user.Password, &user.Sex, &user.Age)
		users = append(users, user)
	}
	return FindRelation(userid, users)
}

// 查找我关注的用户
func SearchFollowUser(userid int) []SearchUser {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	var users []User
	q := "SELECT * FROM user WHERE userid in " +
		"(SELECT followid FROM follow WHERE userid = ?)"
	result, _ := s.Query(q, userid)
	for result.Next() {
		var user User
		result.Scan(&user.Userid, &user.Username, &user.Password, &user.Sex, &user.Age)
		users = append(users, user)
	}
	return FindRelation(userid, users)
}

//查找关注我的用户
func SearchFollowedUser(userid int) []SearchUser {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	var users []User
	q := "SELECT * FROM user WHERE userid in " +
		"(SELECT userid FROM follow WHERE followid = ?)"
	result, _ := s.Query(q, userid)
	for result.Next() {
		var user User
		result.Scan(&user.Userid, &user.Username, &user.Password, &user.Sex, &user.Age)
		users = append(users, user)
	}
	return FindRelation(userid, users)
}

//添加关注关系
func AddFollow(userid int, followid int) bool {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	s.Begin()
	insert := "INSERT INTO follow (`userid`, `followid`) VALUES (?, ?)"
	if _, err = s.Exec(insert, userid, followid); err != nil {
		s.Rollback()
		return false
	}
	s.Commit()
	return true
}

// 删除关注关系
func DeleteFollow(userid int, followid int) bool {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	s.Begin()
	de := "delete from follow where userid = ? and followid = ?"
	if _, err = s.Exec(de, userid, followid); err != nil {
		s.Rollback()
		return false
	}
	s.Commit()
	return true
}

//查找关注和被关注的总数
func CountFollow(userid int) (ct int, cfd int) {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	q1 := "select count(*) from follow where userid = ?"
	q2 := "select count(*) from follow where followid = ?"
	s.QueryRow(q1, userid).Scan(&ct)
	s.QueryRow(q2, userid).Scan(&cfd)
	return
}
