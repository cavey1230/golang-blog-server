package model

type Category struct {
	Id int64
	CreateTime string `xorm:"created"`
	UpdateTime string `xorm:"updated"`
	Name string `xorm:"varchar(20) notnull" json:"name"`
}