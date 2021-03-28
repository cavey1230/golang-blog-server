package model

import (
	"fmt"
	"goblog/utils/errmsg"
)

type Info struct {
	Id         int64  `json:"id"`
	CreateTime string `xorm:"created" json:"createTime"`
	UpdateTime string `xorm:"updated" json:"updateTime"`
	DeleteTime string `xorm:"deleted" json:"deleteTime"`
	FillString string `xorm:"varchar(255) notnull" json:"fill_string"`
	Image      string `xorm:"varchar(255) notnull" json:"image"`
	Name       string `xorm:"varchar(20) notnull" json:"name"`
	Wechat     string `xorm:"varchar(255) notnull" json:"wechat"`
	Address    string `xorm:"varchar(255) notnull" json:"address"`
	Checked    string `xorm:"varchar(20) " json:"checked"`
}

// 取得个人信息
func GetInfo() (Info, int) {
	var code int
	var info Info
	_, err := Db.Where("checked = 1").Get(&info)
	if err != nil {
		code = errmsg.ERROR
	}
	code = errmsg.SUCCSE
	fmt.Println(info)
	return info, code
}

// 模糊查询个人信息
func FindInfo(pageSize int, pageNum int,
	FillString string, Image string, Name string,
	Wechat string, Address string, Checked string) ([]Info, int64) {
	var info []Info
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	sql := fmt.Sprintf("fill_string Like %v%v%v ", "'%", FillString, "%'")
	sql += fmt.Sprintf("AND image  Like %v%v%v ", "'%", Image, "%'")
	sql += fmt.Sprintf("AND name  Like %v%v%v ", "'%", Name, "%'")
	sql += fmt.Sprintf("AND wechat  Like %v%v%v ", "'%", Wechat, "%'")
	sql += fmt.Sprintf("AND address  Like %v%v%v ", "'%", Address, "%'")
	sql += fmt.Sprintf("AND checked  Like %v%v%v ", "'%", Checked, "%'")
	fmt.Println(sql)
	err := Db.Where(sql).Limit(pageSize, offset).Find(&info)
	total, _ := Db.Where(sql).Count(&ToolsLink{})
	if err != nil {
		return nil, total
	}
	return info, total
}

// 新增个人信息
func CreateInfo(data *Info) int {
	info := Info{
		FillString: data.FillString,
		Name:       data.Name,
		Image:      data.Image,
		Wechat:     data.Wechat,
		Address:    data.Address,
		Checked:    data.Checked,
	}
	_, err := Db.Insert(&info)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取个人信息
func GetAllInfo(pageSize int, pageNum int) ([]Info, int64) {
	var info []Info
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	err := Db.Limit(pageSize, offset).Find(&info)
	total, _ := Db.Count(&Info{})
	if err != nil {
		return nil, -1
	}
	return info, total
}

//删除个人信息
func DeleteInfo(id int) int {
	var info Info
	_, err := Db.ID(id).Delete(&info)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//编辑个人信息
func EditInfo(id int, data *Info) int {
	info := Info{
		FillString: data.FillString,
		Name:       data.Name,
		Image:      data.Image,
		Wechat:     data.Wechat,
		Address:    data.Address,
		Checked:    data.Checked,
	}
	_, err := Db.ID(id).Update(&info)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
