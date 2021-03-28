package model

import "goblog/utils/errmsg"

type Blogroll struct {
	Id         int64  `json:"id"`
	CreateTime string `xorm:"created" json:"createTime"`
	UpdateTime string `xorm:"updated" json:"updateTime"`
	DeleteTime string `xorm:"deleted" json:"deleteTime"`
	Link       string `xorm:"varchar(255) notnull" json:"link"`
	Title      string `xorm:"varchar(255) notnull" json:"title"`
}

func GetOneBlogroll(id int) (Blogroll, int) {
	var blogroll Blogroll
	var code int
	_, err := Db.ID(id).Get(&blogroll)
	if err != nil {
		code = errmsg.ERROR
	}
	code = errmsg.SUCCSE
	return blogroll, code
}

func CreateBlogroll(data *Blogroll) int {
	var blogroll = Blogroll{
		Link:  data.Link,
		Title: data.Title,
	}
	_, err := Db.Insert(&blogroll)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetBlogroll() ([]Blogroll, int) {
	var blogroll []Blogroll
	err := Db.Limit(10, 0).Find(&blogroll)
	if err != nil {
		return blogroll, errmsg.ERROR
	}
	return blogroll, errmsg.SUCCSE
}

func DeleteBlogroll(id int) int {
	var blogroll Blogroll
	_, err := Db.ID(id).Delete(&blogroll)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func EditBlogroll(id int, data *Blogroll) int {
	var blogroll = Blogroll{
		Link:  data.Link,
		Title: data.Title,
	}
	_, err := Db.ID(id).Update(&blogroll)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
