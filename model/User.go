package model

import (
	"encoding/base64"
	"fmt"
	"goblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	Id         int64
	CreateTime string `xorm:"created"`
	UpdateTime string `xorm:"updated"`
	DeleteTime string `xorm:"deleted"`
	Username   string `xorm:"varchar(20) notnull" json:"username"`
	Password   string `xorm:"varchar(20) notnull" json:"password"`
	Role       int    `xorm:"int" json:"role"`
	//Avatar string
}

func CheckUser(username string) int {
	var user User
	_, _ = Db.Select("id").Where("username = ?", username).Get(&user)
	if user.Id > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	user := User{
		Username: data.Username,
		Password: data.Password,
		Role:     data.Role,
	}
	before := func(bean interface{}) {
		user.Password = ScryptPw(user.Password)
	}
	affected, err := Db.Before(before).Insert(&user)
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

//删除用户
func DeleteUser(id int) int {
	var user User
	_, err := Db.ID(id).Delete(&user)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//编辑用户
func EditUser(id int, data *User) int {
	user := User{
		Username: data.Username,
		Role:     data.Role,
	}
	_, err := Db.ID(id).Update(&user)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{
		12, 11, 33, 44, 55, 11, 22, 33,
	}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
