package model

import (
	"fmt"
	"goblog/utils/errmsg"
)

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

// 查询单一工具链接
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

// 创建工具链接
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

// 模糊查询工具链接
func FindToolsLink(pageSize int, pageNum int,
	iconImg string, link string,
	title string, introduce string) ([]ToolsLink, int64) {
	var toolsLink []ToolsLink
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	sql := fmt.Sprintf("icon_img Like %v%v%v ", "'%", iconImg, "%'")
	sql += fmt.Sprintf("AND link  Like %v%v%v ", "'%", link, "%'")
	sql += fmt.Sprintf("AND title  Like %v%v%v ", "'%", title, "%'")
	sql += fmt.Sprintf("AND introduce  Like %v%v%v ", "'%", introduce, "%'")
	fmt.Println(sql)
	err := Db.Where(sql).Limit(pageSize, offset).Find(&toolsLink)
	total, _ := Db.Where(sql).Count(&ToolsLink{})
	if err != nil {
		return nil, total
	}
	return toolsLink, total
}

// 查询所有工具链接
func GetToolsLinkList(pageSize int, pageNum int) ([]ToolsLink, int64) {
	var toolsLink []ToolsLink
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	err := Db.Limit(pageSize, offset).Find(&toolsLink)
	total, _ := Db.Count(&ToolsLink{})
	if err != nil {
		return toolsLink, total
	}
	return toolsLink, total
}

// 删除工具链接
func DeleteToolsLink(id int) int {
	var toolsLink ToolsLink
	_, err := Db.ID(id).Delete(&toolsLink)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 编辑工具链接
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
