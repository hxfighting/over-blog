package models

import (
	"github.com/ohdata/blog/tools/util"
)

const ARTICLE_VIEW = "blog_article_view"

type Article struct {
	ID            *int64         `json:"id" validate:"gt=0"`
	Title         string         `json:"title" validate:"gte=2,lte=100"`
	Author        string         `json:"author" validate:"gte=2,lte=20"`
	Keywords      string         `json:"keywords" validate:"gte=2,lte=255"`
	Description   string         `json:"description" validate:"lte=255"`
	Thumb         string         `json:"thumb" validate:"url"`
	ContentHtml   string         `json:"content_html" validate:"gte=2" mapstructure:"content_html"`
	ContentMd     string         `json:"content_md" gorm:"column:content_md" validate:"gte=2" mapstructure:"content_md"`
	IsShow        *int8          `json:"is_show" validate:"oneof=0 1" mapstructure:"is_show"`
	IsTop         *int8          `json:"is_top" validate:"oneof=0 1" mapstructure:"is_top"`
	IsOriginal    *int8          `json:"is_original" validate:"oneof=0 1" mapstructure:"is_original"`
	Click         int64          `json:"click"`
	CategoryID    *int64         `json:"category_id" validate:"gt=0,numeric" mapstructure:"category_id"`
	CreatedUnix   int64          `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix   int64          `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt     string         `json:"created_at" gorm:"-"`
	UpdatedAt     string         `json:"updated_at" gorm:"-"`
	Category      simpleCategory `json:"category" gorm:"-"`
	Tags          []simpleTag    `json:"tags" gorm:"-"`
	CommentsCount int64          `json:"comments_count" gorm:"-"`
}

type SimpleArticle struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Thumb string `json:"thumb"`
}

type simpleCategory struct {
	ID    *int64  `json:"id"`
	Title *string `json:"title"`
}

type simpleTag struct {
	ID        *int64 `json:"id"`
	Name      string `json:"name"`
	ArticleID int64  `json:"-"`
}

func (a *Article) AfterFind() {
	a.CreatedAt = util.GetDateTime(a.CreatedUnix, util.YMDHIS)
	a.UpdatedAt = util.GetDateTime(a.UpdatedUnix, util.YMDHIS)
}
