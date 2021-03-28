package model

import (
	"fmt"
	"goblog/utils/errmsg"
)

type Article struct {
	Category `xorm:"extends"`
	Title    string `xorm:"varchar(50) notnull" json:"title"`
	Cid      int64  `xorm:"int notnull" json:"cid"`
	Synopsis string `xorm:"varchar(200) notnull" json:"synopsis"`
	Content  string `xorm:"longtext notnull" json:"content"`
	Img      string `xorm:"varchar(200)" json:"img"`
	Boutique string `xorm:"varchar(1) default(0)" json:"boutique"`
}

// 检查文章
func CheckArticle(title string) int {
	var article Article
	_, _ = Db.Select("id").Where("title = ?", title).Get(&article)
	if article.Id > 0 {
		return errmsg.ERROR_ARTICLE_TITLE_USED
	}
	return errmsg.SUCCSE
}

// 新增文章
func CreateArticle(data *Article) int {
	var article = Article{
		Category: data.Category,
		Title:    data.Title,
		Cid:      data.Cid,
		Synopsis: data.Synopsis,
		Content:  data.Content,
		Img:      data.Img,
		Boutique: data.Boutique,
	}
	_, err := Db.Insert(&article)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 模糊查询文章列表
func FindAllArticles(pageSize int, pageNum int,
	Title string, Cid string, Synopsis string,
	Content string, Img string, Boutique string) ([]Article, int64) {
	var articles []Article
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	sql := fmt.Sprintf("title Like %v%v%v ", "'%", Title, "%'")
	sql += fmt.Sprintf("AND cid Like %v%v%v ", "'%", Cid, "%'")
	sql += fmt.Sprintf("AND synopsis Like %v%v%v ", "'%", Synopsis, "%'")
	sql += fmt.Sprintf("AND content Like %v%v%v ", "'%", Content, "%'")
	sql += fmt.Sprintf("AND img Like %v%v%v ", "'%", Img, "%'")
	sql += fmt.Sprintf("AND boutique Like %v%v%v ", "'%", Boutique, "%'")
	fmt.Println(sql)
	err := Db.Table("article").
		Select(
			"article.`id`,"+
				"category.`name`,"+
				"article.`synopsis`,"+
				"article.`img`,"+
				"cid,"+
				"title,"+
				"article.`content`,"+
				"article.`create_time`,"+
				"article.`update_time`,"+
				"article.`boutique`").
		Join("INNER", "category", "article.cid = category.id").
		Where(sql).Limit(pageSize, offset).Find(&articles)
	total, _ := Db.Where(sql).Count(&Article{})
	if err != nil {
		return nil, total
	}
	return articles, total
}

// 获取精品文章列表
func GetAllBoutiqueArticle(pageSize int, pageNum int) ([]Article, int64) {
	var articles []Article
	var whereSql = "boutique='1'"
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	err := Db.Table("article").
		Select(
			"article.`id`,"+
				"category.`name`,"+
				"article.`synopsis`,"+
				"article.`img`,"+
				"cid,"+
				"title,"+
				"article.`content`,"+
				"article.`create_time`,"+
				"article.`update_time`,"+
				"article.`boutique`").
		Join("INNER", "category", "article.cid = category.id").
		Where(whereSql).Limit(pageSize, offset).Find(&articles)
	total, _ := Db.Where(whereSql).Count(&Article{})
	if err != nil {
		return nil, -1
	}
	return articles, total
}

// 获取文章列表
func GetAllArticle(pageSize int, pageNum int, cid int) ([]Article, int64) {
	var articles []Article
	var whereSql = ""
	if pageSize == 0 {
		pageSize = 10
	}
	if cid != 0 {
		whereSql = fmt.Sprintf("%v%v", "cid=", cid)
	}
	offset := (pageNum - 1) * pageSize
	err := Db.Table("article").
		Select(
			"article.`id`,"+
				"category.`name`,"+
				"article.`synopsis`,"+
				"article.`img`,"+
				"cid,"+
				"title,"+
				"article.`content`,"+
				"article.`create_time`,"+
				"article.`update_time`,"+
				"article.`boutique`").
		Join("INNER", "category", "article.cid = category.id").
		Where(whereSql).Desc("article.`create_time`").
		Limit(pageSize, offset).
		Find(&articles)
	total, _ := Db.Where(whereSql).Count(&Article{})
	if err != nil {
		return nil, -1
	}
	return articles, total
}

//获取单篇文章
func GetOneArticle(id int) interface{} {
	var article Article
	complete, _ := Db.Select(
		"article.`id`,"+
			"category.`name`,"+
			"article.`synopsis`,"+
			"article.`img`,"+
			"cid,"+
			"title,"+
			"article.`content`,"+
			"article.`create_time`,"+
			"article.`update_time`,"+
			"article.`boutique`").
		ID(id).
		Join("INNER", "category", "article.cid = category.id").
		Get(&article)
	if complete {
		return article
	} else {
		return errmsg.ERROR_ARTICLE_NOT_DEFINE
	}
}

//删除文章
func DeleteArticle(id int) int {
	var article Article
	_, err := Db.ID(id).Delete(&article)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//编辑文章
func EditArticle(id int, data *Article) int {
	var article = Article{
		Category: data.Category,
		Title:    data.Title,
		Cid:      data.Cid,
		Synopsis: data.Synopsis,
		Content:  data.Content,
		Img:      data.Img,
		Boutique: data.Boutique,
	}
	_, err := Db.ID(id).Update(&article)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
