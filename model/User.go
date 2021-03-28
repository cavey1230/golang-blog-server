package model

import (
	"encoding/base64"
	"fmt"
	"goblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	Id         int64  `json:"id"`
	CreateTime string `xorm:"created" json:"createTime"`
	UpdateTime string `xorm:"updated" json:"updateTime"`
	DeleteTime string `xorm:"deleted" json:"deleteTime"`
	Username   string `xorm:"varchar(20) notnull" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password   string `xorm:"varchar(20) notnull" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role       int    `xorm:"int default 2" json:"role" validate:"required" label:"角色名"`
	Avatar     string `xorm:"varchar(100)" json:"avatar" validate:"max=100"  label:"头像地址"`
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
		//Avatar:   data.Avatar,
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
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	err := Db.Select("id,username,create_time,update_time,role").Limit(pageSize, offset).Find(&users)
	total, _ := Db.Count(&User{})
	if err != nil {
		return nil, -1
	}
	return users, total
}

// 查询指定用户
func FindUsers(pageSize int, pageNum int, username string, role int) ([]User, int64) {
	var users []User
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	sql := fmt.Sprintf("username Like %v%v%v ", "'%", username, "%'")
	sql = fmt.Sprintf("%v And role Like %v%v%v ", sql, "'%", role, "%'")
	err := Db.Where(sql).Limit(pageSize, offset).Find(&users)
	total, _ := Db.Where(sql).Count(&User{})
	if err != nil {
		return nil, total
	}
	return users, total
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

//登录验证
func CheckLogin(username string, password string) int {
	var user User
	_, _ = Db.Where("username=?", username).Get(&user)
	if user.Id == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	fmt.Println(ScryptPw(password))
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 2 {
		return errmsg.ERROR_USERS_ROLE_ERROR
	}
	return errmsg.SUCCSE
}

//后台管理员认证
func CheckAdminLogin(username string, password string) int {
	var user User
	_, _ = Db.Where("username=? And role = 3", username).Get(&user)
	if user.Id == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 3 {
		return errmsg.ERROR_USERS_ROLE_ERROR
	}
	return errmsg.SUCCSE
}
