package model

import (
	"fmt"
	"goblog/utils/errmsg"
)

type User struct {
	Id         int64
	CreateTime string `xorm:"created"`
	UpdateTime string `xorm:"updated"`
	Username   string `xorm:"varchar(20) notnull" json:"username"`
	Password   string `xorm:"varchar(20) notnull" json:"password"`
	Role       int    `xorm:"int" json:"role"`
	//Avatar string
}

func CheckUser(username string) int {
	var user User
	isBool, _ := Db.Select("id").Where("username = ?", username).Get(&user)
	if user.Id > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	fmt.Println(isBool)
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	user := User{
		Username: data.Username,
		Password: data.Password,
		Role:     data.Role,
	}
	affected, err := Db.Insert(&user)
	fmt.Println(affected, "插入成功返回的状态")
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	err := Db.Limit(pageSize, offset).Find(&users)
	if err != nil {
		return nil
	}
	return users
}
