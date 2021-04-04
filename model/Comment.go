package model

import (
	"fmt"
	"goblog/utils/errmsg"
)

type Comment struct {
	Id          int64  `json:"id"`
	CreateTime  string `xorm:"created" json:"createTime"`
	UpdateTime  string `xorm:"updated" json:"updateTime"`
	DeleteTime  string `xorm:"deleted" json:"deleteTime"`
	Content     string `xorm:"varchar(200) notnull" json:"content"`
	ArticleId   int64  `xorm:"int notnull" json:"articleId"`
	CommenterId int64  `xorm:"int notnull" json:"commenterId"`
	Guestbook   string `xorm:"varchar(50)" json:"guestbook"`
}

type CommentWithUser struct {
	Comment  `xorm:"extends"`
	Username string `json:"username"`
}

// 新增评论
func CreateComment(data *Comment) int {
	var comment = Comment{
		ArticleId:   data.ArticleId,
		CommenterId: data.CommenterId,
		Content:     data.Content,
		Guestbook:   data.Guestbook,
	}
	_, err := Db.Insert(&comment)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 模糊查询评论
func FindAllComment(pageSize int, pageNum int,
	Content string, ArticleId string,
	CommenterId string, Guestbook string,
) ([]CommentWithUser, int64) {
	var commentWithUser []CommentWithUser
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	sql := fmt.Sprintf("content Like %v%v%v ", "'%", Content, "%'")
	sql += fmt.Sprintf("AND article_id Like %v%v%v ", "'%", ArticleId, "%'")
	sql += fmt.Sprintf("AND commenter_id Like %v%v%v ", "'%", CommenterId, "%'")
	sql += fmt.Sprintf("AND guestbook Like %v%v%v ", "'%", Guestbook, "%'")
	fmt.Println(sql)
	err := Db.Table("comment").
		Select(
			"comment.create_time,"+
				"comment.id,"+
				"comment.content,"+
				"comment.article_id,"+
				"comment.commenter_id,"+
				"user.username,"+
				"comment.guestbook").
		Join("INNER", "user", "comment.commenter_id = user.id").
		Where(sql).Desc("comment.create_time").Limit(pageSize, offset).Find(&commentWithUser)
	total, _ := Db.Where(sql).Count(&Comment{})
	if err != nil {
		return nil, total
	}
	return commentWithUser, total
}

//删除评论
func DeleteComment(id int) int {
	var comment Comment
	_, err := Db.ID(id).Delete(&comment)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
