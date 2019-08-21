package models

import (
	"blog/database"
	"blog/helper"
	"blog/service"
)

type Comment struct {
	ID           *int64      `json:"id"`
	Pid          *int64      `json:"pid" gorm:"column:pid"`
	ReplyID      *int64      `json:"reply_id"`
	UserID       *int64      `json:"user_id"`
	ArticleID    *int64      `json:"article_id"`
	Content      *string     `json:"content"`
	CreatedUnix  int64       `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix  int64       `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt    string      `json:"created_at" gorm:"-"`
	UpdatedAt    string      `json:"updated_at" gorm:"-"`
	User         *simpleUser `json:"user"`
	Replier      *simpleUser `json:"replier"`
	ReplyContent *string     `gorm:"-" json:"reply_content"`
}

type simpleArticle struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type simpleUser struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (Comment) TableName() string {
	return "article_comment"
}

func (this *Comment) AfterFind() {
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取评论列表
*/
func GetCommentList(pageNum, pageSize, article_id int64) map[string]interface{} {
	var data = make(map[string]interface{})
	articles := []simpleArticle{}
	database.Db.Table("article").Select("id,title").Find(&articles)
	comments := []Comment{}
	total := 0
	var db = database.Db.Table("article_comment")
	if article_id != 0 {
		db = db.Where("article_id = ?", article_id)
	}
	db.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&comments)
	if len(comments) > 0 {
		user_ids, replier_ids := []int64{}, []int64{}
		users := []User{}
		repliers := []User{}
		for _, value := range comments {
			if *value.UserID != 0 {
				user_ids = append(user_ids, *value.UserID)
			}
			if *value.ReplyID != 0 {
				replier_ids = append(replier_ids, *value.ReplyID)
			}
		}
		database.Db.Table("user").Where("id in (?)", user_ids).Find(&users)
		database.Db.Table("user").Where("id in (?)", replier_ids).Find(&repliers)
		users_map := make(map[int64]User)
		repliers_map := make(map[int64]User)
		if len(users) > 0 {
			for _, value := range users {
				users_map[*value.ID] = value
			}
		}
		if len(repliers) > 0 {
			for _, value := range repliers {
				repliers_map[*value.ID] = value
			}
		}
		for k, value := range comments {
			if val, ok := users_map[*value.UserID]; ok {
				value.User = &simpleUser{*val.ID, val.Name}
			}
			if va, o := repliers_map[*value.ReplyID]; o {
				value.Replier = &simpleUser{*va.ID, va.Name}
			}
			comments[k] = value
		}
	}
	db.Count(&total)
	data["article"] = articles
	data["list"] = comments
	data["total"] = total
	return data
}

/**
删除评论
*/
func DeleteComment(comment *Comment) bool {
	database.Db.Where("id = ?", *comment.ID).First(comment)
	if *comment.ID == 0 {
		return false
	}
	tx := database.Db.Begin();
	if *comment.Pid == 0 {
		if err := tx.Where("pid = ?", *comment.UserID).Delete(Comment{}).Error; err != nil {
			tx.Rollback();
			return false
		}
	}
	if err := tx.Delete(comment).Error; err != nil {
		tx.Rollback();
		return false
	}
	tx.Commit()
	return true
}

/**
回复评论
 */
func ReplyComment(comment *Comment) bool {
	database.Db.First(comment)
	if *comment.ID == 0 {
		return false
	}
	var user_id int64
	if *comment.ReplyID == 0 {
		user := User{}
		database.Db.Where("is_admin = ?", 1).First(&user)
		if *user.ID != 0 {
			user_id = *user.ID
		}
	}
	if user_id == 0 {
		return false
	}
	var pid int64 = *comment.Pid
	if pid == 0 {
		pid = *comment.UserID
	}
	new_comment := Comment{}
	new_comment.Pid = &pid
	new_comment.Content = comment.ReplyContent
	new_comment.UserID = &user_id
	new_comment.ReplyID = comment.UserID
	new_comment.ArticleID = comment.ArticleID
	res := database.Db.Create(&new_comment)
	if res.Error != nil {
		return false
	}
	push := service.Redis.LPush("email_queue", *new_comment.ID)
	if push.Err() != nil {
		return false
	}
	return true
}
