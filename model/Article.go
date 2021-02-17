package model


type Article struct {
	Category `xorm:"extends"`
	Title string `xorm:"varchar(50) notnull" json:"title"`
	Cid int `xorm:"int notnull" json:"cid"`
	Desc string `xorm:"varchar(200) notnull" json:"desc"`
	Content string `xorm:"longtext notnull" json:"content"`
	Img string `xorm:"varchar(20)" json:"img"`
}