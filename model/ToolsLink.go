package model

import "goblog/utils/errmsg"

type ToolsLink struct {
	Id         int64  `json:"id"`
	CreateTime string `xorm:"created" json:"createTime"`
	UpdateTime string `xorm:"updated" json:"updateTime"`
	DeleteTime string `xorm:"deleted" json:"deleteTime"`
	IconImg    string `xorm:"varchar(255) notnull" json:"icon_img"`
	Link       string `xorm:"varchar(255) notnull" json:"link"`
	Title      string `xorm:"varchar(255) notnull" json:"title"`
	Introduce  string `xorm:"varchar(255) notnull" json:"introduce"`
}

func GetOneToolsLink(id int) (ToolsLink, int) {
	var toolsLink ToolsLink
	var code int
	_, err := Db.ID(id).Get(&toolsLink)
	if err != nil {
		code = errmsg.ERROR
	}
	code = errmsg.SUCCSE
	return toolsLink, code
}

func CreateToolsLink(data *ToolsLink) int {
	var toolsLink = ToolsLink{
		IconImg:   data.IconImg,
		Link:      data.Link,
		Title:     data.Title,
		Introduce: data.Introduce,
	}
	_, err := Db.Insert(&toolsLink)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetToolsLinkList() ([]ToolsLink, int) {
	var toolsLink []ToolsLink
	err := Db.Limit(10, 0).Find(&toolsLink)
	if err != nil {
		return toolsLink, errmsg.ERROR
	}
	return toolsLink, errmsg.SUCCSE
}

func DeleteToolsLink(id int) int {
	var toolsLink ToolsLink
	_, err := Db.ID(id).Delete(&toolsLink)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func EditToolsLink(id int, data *ToolsLink) int {
	var toolsLink = ToolsLink{
		IconImg:   data.IconImg,
		Link:      data.Link,
		Title:     data.Title,
		Introduce: data.Introduce,
	}
	_, err := Db.ID(id).Update(&toolsLink)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
