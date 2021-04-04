package model

import (
	"fmt"
	"goblog/utils/errmsg"
)

type Timeline struct {
	Id         int64  `json:"id"`
	CreateTime string `xorm:"created" json:"createTime"`
	UpdateTime string `xorm:"updated" json:"updateTime"`
	DeleteTime string `xorm:"deleted" json:"deleteTime"`
	Title      string `xorm:"varchar(200) notnull" json:"title"`
	Color      string `xorm:"varchar(50)" json:"color"`
	Com        string `xorm:"varchar(50)" json:"com"`
}

// 新增时间轴
func CreateTimeline(data *Timeline) int {
	var timeline = Timeline{
		Color: data.Color,
		Title: data.Title,
		Com:   data.Com,
	}
	_, err := Db.Insert(&timeline)
	if err != nil {
		return errmsg.ERROR_TIMELINE_CREATE_ERROR
	}
	return errmsg.SUCCSE
}

// 模糊查询时间轴
func FindAllTimeline(pageSize int, pageNum int,
	Title string, Color string, Com string,
) ([]Timeline, int64) {
	var timeline []Timeline
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	sql := fmt.Sprintf("title Like %v%v%v ", "'%", Title, "%'")
	sql += fmt.Sprintf("AND color Like %v%v%v ", "'%", Color, "%'")
	sql += fmt.Sprintf("AND com Like %v%v%v ", "'%", Com, "%'")
	fmt.Println(sql)
	err := Db.Where(sql).Limit(pageSize, offset).Find(&timeline)
	total, _ := Db.Where(sql).Count(&Timeline{})
	if err != nil {
		return nil, total
	}
	return timeline, total
}

//删除时间轴
func DeleteTimeline(id int) int {
	var timeline Timeline
	_, err := Db.ID(id).Delete(&timeline)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//编辑时间轴
func EditTimeline(id int, data *Timeline) int {
	var timeline = Timeline{
		Color: data.Color,
		Title: data.Title,
		Com:   data.Com,
	}
	_, err := Db.ID(id).Update(&timeline)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
