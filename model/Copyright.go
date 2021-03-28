package model

import "goblog/utils/errmsg"

type Copyright struct {
	Id         int64  `json:"id"`
	CreateTime string `xorm:"created" json:"createTime"`
	UpdateTime string `xorm:"updated" json:"updateTime"`
	DeleteTime string `xorm:"deleted" json:"deleteTime"`
	Content    string `xorm:"varchar(255) notnull" json:"content"`
	Title      string `xorm:"varchar(255) notnull" json:"title"`
}

func GetOneCopyright(id int) (Copyright, int) {
	var copyright Copyright
	var code int
	_, err := Db.ID(id).Get(&copyright)
	if err != nil {
		code = errmsg.ERROR
	}
	code = errmsg.SUCCSE
	return copyright, code
}

func CreateCopyright(data *Copyright) int {
	var copyright = Copyright{
		Content: data.Content,
		Title:   data.Title,
	}
	_, err := Db.Insert(&copyright)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetCopyright() ([]Copyright, int) {
	var copyright []Copyright
	err := Db.Limit(10, 0).Find(&copyright)
	if err != nil {
		return copyright, errmsg.ERROR
	}
	return copyright, errmsg.SUCCSE
}

func DeleteCopyright(id int) int {
	var copyright Copyright
	_, err := Db.ID(id).Delete(&copyright)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func EditCopyright(id int, data *Copyright) int {
	var copyright = Copyright{
		Content: data.Content,
		Title:   data.Title,
	}
	_, err := Db.ID(id).Update(&copyright)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
