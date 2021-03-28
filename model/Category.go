package model

import (
	"fmt"
	"goblog/utils/errmsg"
	"strconv"
)

type Category struct {
	Id         int64  `json:"id"`
	CreateTime string `xorm:"created" json:"createTime"`
	UpdateTime string `xorm:"updated" json:"updateTime"`
	DeleteTime string `xorm:"deleted" json:"deleteTime"`
	Name       string `xorm:"varchar(20) notnull" json:"name"`
}

type Data struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// 检查分类
func CheckCategory(name string) int {
	var category Category
	_, _ = Db.Select("id").Where("name = ?", name).Get(&category)
	if category.Id > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCSE
}

// 新增分类
func CreateCategory(data *Category) int {
	category := Category{
		Name: data.Name,
	}
	_, err := Db.Insert(&category)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 模糊查询指定分类
func FindCategory(pageSize int, pageNum int, name string) ([]Category, int64) {
	var category []Category
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	sql := fmt.Sprintf("name Like %v%v%v ", "'%", name, "%'")
	err := Db.Where(sql).Limit(pageSize, offset).Find(&category)
	total, _ := Db.Where(sql).Count(&Category{})
	if err != nil {
		return nil, total
	}
	return category, total
}

// 获取分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int64) {
	var category []Category
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	err := Db.Select("id,name,create_time,update_time").Limit(pageSize, offset).Find(&category)
	total, _ := Db.Count(&Category{})
	if err != nil {
		return nil, -1
	}
	return category, total
}

// 获取所有分类
func GetAllCategory() []Data {
	var category []Category
	err := Db.Select("id,name,create_time,update_time").Find(&category)
	if err != nil {
		return nil
	}
	var data []Data
	for _, c := range category {
		data = append(data, Data{
			Label: c.Name,
			Value: strconv.FormatInt(c.Id, 10),
		})
	}
	return data
}

//删除分类
func DeleteCategory(id int) int {
	var category Category
	_, err := Db.ID(id).Delete(&category)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//编辑分类
func EditCategory(id int, data *Category) int {
	category := Category{
		Name: data.Name,
	}
	_, err := Db.ID(id).Update(&category)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
