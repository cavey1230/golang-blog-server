package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goblog/utils"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var Db *xorm.Engine
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	Db, err = xorm.NewEngine(utils.Db, dsn)
	if err != nil {
		fmt.Println("数据库链接失败", err)
	} else {
		Db.ShowSQL(true)
		Db.SetConnMaxLifetime(time.Second * 10)
		Db.SetMaxIdleConns(200)
		Db.SetMaxOpenConns(200)
		Db.SetMapper(names.GonicMapper{})
		err := Db.Sync2(new(User), new(Article), new(Category), new(Info), new(ToolsLink))
		if err != nil {
			fmt.Println("自动更新失败", err)
		}
	}

}
