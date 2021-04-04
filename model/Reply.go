package model

import (
	"fmt"
	"goblog/utils/errmsg"
)

type Reply struct {
	Id               int64  `json:"id"`
	CreateTime       string `xorm:"created" json:"createTime"`
	UpdateTime       string `xorm:"updated" json:"updateTime"`
	DeleteTime       string `xorm:"deleted" json:"deleteTime"`
	Content          string `xorm:"varchar(200) notnull" json:"content"`
	ArticleId        int64  `xorm:"int notnull" json:"articleId"`
	ReplierId        int64  `xorm:"int notnull" json:"replierId"`
	ReplyToCommentId int64  `xorm:"int notnull" json:"replyToCommentId"`
	Guestbook        string `xorm:"varchar(50)" json:"guestbook"`
}

type ReplyWithUser struct {
	Reply     `xorm:"extends"`
	Username1 string `json:"username1"`
	Username2 string `json:"username2"`
}

// 新增回复
func CreateReply(data *Reply) int {
	var reply = Reply{
		ArticleId:        data.ArticleId,
		ReplierId:        data.ReplierId,
		Content:          data.Content,
		Guestbook:        data.Guestbook,
		ReplyToCommentId: data.ReplyToCommentId,
	}
	_, err := Db.Insert(&reply)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 模糊查询回复
func FindAllReply(pageSize int, pageNum int,
	Content string, ArticleId string,
	ReplierId string, ReplyToCommentId string,
	Guestbook string,
) ([]ReplyWithUser, int64) {
	var replyWithUser []ReplyWithUser
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	sql := fmt.Sprintf("content Like %v%v%v ", "'%", Content, "%'")
	sql += fmt.Sprintf("AND article_id Like %v%v%v ", "'%", ArticleId, "%'")
	sql += fmt.Sprintf("AND replier_id Like %v%v%v ", "'%", ReplierId, "%'")
	sql += fmt.Sprintf("AND reply_to_comment_id Like %v%v%v ", "'%", ReplyToCommentId, "%'")
	sql += fmt.Sprintf("AND guestbook Like %v%v%v ", "'%", Guestbook, "%'")
	fmt.Println(sql)
	err := Db.Table("reply").
		Select(
			"reply.create_time,"+
				"reply.id,"+
				"reply.content,"+
				"reply.article_id,"+
				"reply.replier_id,"+
				"reply.reply_to_comment_id,"+
				"u1.username as username1,"+
				"reply.guestbook").
		Join("INNER", "user as u1", "reply.replier_id = u1.id").
		Where(sql).Desc("reply.create_time").Limit(pageSize, offset).Find(&replyWithUser)
	total, _ := Db.Where(sql).Count(&Reply{})
	if err != nil {
		return nil, total
	}
	return replyWithUser, total
}

//删除回复
func DeleteReply(id int) int {
	var reply Reply
	_, err := Db.ID(id).Delete(&reply)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
