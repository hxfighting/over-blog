package models

type Category struct {
	ID        *int64      `json:"id" validate:"gt=0"`
	Title     *string     `json:"title" validate:"gte=2,lte=20"`
	Pid       *int64      `json:"pid" validate:"gte=0"`
	Url       string      `json:"url"`
	Type      uint        `json:"type"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
	Children  []*Category `json:"children" gorm:"foreignkey:pid;PRELOAD:false"`
}
