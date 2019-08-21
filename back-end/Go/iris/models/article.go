package models

type Article struct {
	ID          *int64 `json:"id"`
	Type        uint   `json:"type"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Thumb       string `json:"thumb"`
	ContentHtml string `json:"content_html"`
	ContentMd   string `json:"content_md" gorm:"column:content_md"`
	IsShow      *int8  `json:"is_show"`
	IsTop       *int8  `json:"is_top"`
	IsOriginal  *int8  `json:"is_original"`
	Click       *int64 `json:"click"`
	CategoryID  *int64 `json:"category_id"`
	CreatedUnix int64  `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64  `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string `json:"created_at" gorm:"-"`
	UpdatedAt   string `json:"updated_at" gorm:"-"`
}
