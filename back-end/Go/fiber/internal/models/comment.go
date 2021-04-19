package models

import (
	"github.com/ohdata/blog/tools/util"
)

type Comment struct {
	ID           *int64      `json:"id" validate:"gt=0"`
	Pid          *int64      `json:"pid" gorm:"column:pid" mapstructure:"pid" validate:"gte=0"`
	ReplyID      *int64      `json:"reply_id" validate:"gte=0" mapstructure:"reply_id"`
	UserID       *int64      `json:"user_id" validate:"gt=0" mapstructure:"user_id"`
	ArticleID    *int64      `json:"article_id" validate:"gt=0" mapstructure:"article_id"`
	Content      *string     `json:"content" validate:"gte=2,lte=255"`
	CreatedUnix  int64       `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix  int64       `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt    string      `json:"created_at" gorm:"-"`
	UpdatedAt    string      `json:"updated_at" gorm:"-"`
	User         *simpleUser `json:"user"`
	Replier      *simpleUser `json:"replier"`
	ReplyContent *string     `gorm:"-" json:"reply_content" mapstructure:"reply_content" validate:"gte=2,lte=255"`
	Email        string      `gorm:"-" json:"email" validate:"email"`
}

type ArticleComment struct {
	ID          *int64            `json:"id" validate:"gt=0"`
	Pid         *int64            `json:"pid" gorm:"column:pid"`
	ReplyID     *int64            `json:"reply_id"`
	UserID      *int64            `json:"user_id"`
	ArticleID   *int64            `json:"article_id"`
	Content     *string           `json:"content"`
	CreatedUnix int64             `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64             `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string            `json:"created_at" gorm:"-"`
	UpdatedAt   string            `json:"updated_at" gorm:"-"`
	ReplyAvatar string            `json:"reply_avatar"`
	ReplyName   string            `json:"reply_name"`
	Username    string            `json:"username"`
	UserAvatar  string            `json:"user_avatar"`
	Children    []*ArticleComment `json:"children"`
}

type RecentComment struct {
	ArticleID int64  `json:"article_id"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
}

type simpleUser struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (Comment) TableName() string {
	return "article_comment"
}

func (this *Comment) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}

func (this *ArticleComment) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}
