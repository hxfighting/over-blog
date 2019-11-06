package models

import (
	"blog/config"
	"blog/database"
	"blog/helper"
	"blog/service"
	"fmt"
	"sort"
	"strings"
	"time"
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
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

func (this *ArticleComment) AfterFind() {
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取评论列表
*/
func GetCommentList(pageNum, pageSize, article_id int64) map[string]interface{} {
	var data = make(map[string]interface{})
	articles := []SimpleArticle{}
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
	database.Db.First(comment)
	if *comment.ID == 0 {
		return false
	}
	tx := database.Db.Begin()
	if *comment.Pid == 0 {
		if err := tx.Where("pid = ?", *comment.UserID).Delete(Comment{}).Error; err != nil {
			tx.Rollback()
			return false
		}
	}
	if err := tx.Delete(comment).Error; err != nil {
		tx.Rollback()
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
	user := User{}
	database.Db.Where("is_admin = ?", 1).First(&user)
	if *user.ID != 0 {
		user_id = *user.ID
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
	new_comment.CreatedUnix = time.Now().Unix()
	new_comment.UpdatedUnix = time.Now().Unix()
	res := database.Db.Create(&new_comment)
	if res.Error != nil {
		return false
	}
	service.EmailChan <- *new_comment.ID
	return true
}

/**
添加评论
*/
func (this Comment) AddComment() bool {
	tx := database.Db.Begin()
	this.CreatedUnix = time.Now().Unix()
	this.UpdatedUnix = time.Now().Unix()
	commentRes := tx.Create(&this)
	if commentRes.Error != nil {
		service.Log.Error(commentRes.Error.Error())
		tx.Rollback()
		return false
	}
	userRes := tx.Table("user").Where("id = ?", *this.UserID).Update("email", this.Email)
	if userRes.Error != nil {
		service.Log.Error(commentRes.Error.Error())
		tx.Rollback()
		return false
	}
	tx.Commit()
	if *this.ReplyID != 0 {
		service.EmailChan <- *this.ID
	} else {
		service.CommentEmailChan <- *this.ID
	}
	return true
}

/**
发送回复邮件
*/
func HandleEmailQueue(id int64) {
	comment := Comment{}
	database.Db.Where("id = ?", id).First(&comment)
	if comment.ID == nil {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的评论发送邮件失败")
		return
	}
	users := []User{}
	database.Db.Table("user").Where("id in (?)", []int64{*comment.ReplyID, *comment.UserID}).Find(&users)
	if len(users) != 2 {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的评论发送邮件失败,缺少评论用户")
		return
	}
	reply_user := User{}
	comment_user := User{}
	for _, value := range users {
		if *value.ID == *comment.ReplyID {
			reply_user = value
		} else {
			comment_user = value
		}
	}
	article := Article{}
	database.Db.Where("id = ?", *comment.ArticleID).First(&article)
	if article.ID == nil {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的评论发送邮件失败")
		return
	}
	url := config.GetConfig("app.url").(string)
	blog_name := config.GetConfig("app.name").(string)
	if url == "" || blog_name == "" {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的评论发送邮件失败,配置文件app项为空")
		return
	}
	year := time.Now().Year()
	article_url := strings.TrimRight(url, "/") + "/article/" + fmt.Sprintf("%d", *comment.ArticleID)
	html_raw := `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
</head>
<body style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; background-color: #f8fafc; color: #74787e; height: 100%; hyphens: auto; line-height: 1.4; margin: 0; -moz-hyphens: auto; -ms-word-break: break-all; width: 100% !important; -webkit-hyphens: auto; -webkit-text-size-adjust: none; word-break: break-word;">
    <style>
        @media  only screen and (max-width: 600px) {
            .inner-body {
                width: 100% !important;
            }

            .footer {
                width: 100% !important;
            }
        }

        @media  only screen and (max-width: 500px) {
            .button {
                width: 100% !important;
            }
        }
    </style>

    <table class="wrapper" width="100%" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; background-color: #f8fafc; margin: 0; padding: 0; width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 100%;">
        <tr>
            <td align="center" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                <table class="content" width="100%" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; margin: 0; padding: 0; width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 100%;">
                    <tr>
    <td class="header" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; padding: 25px 0; text-align: center;">
        <a href="` + url + `" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #bbbfc3; font-size: 19px; font-weight: bold; text-decoration: none; text-shadow: 0 1px 0 white;">
            ` + blog_name + `
        </a>
    </td>
</tr>

                    <!-- Email Body -->
                    <tr>
                        <td class="body" width="100%" cellpadding="0" cellspacing="0" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; background-color: #ffffff; border-bottom: 1px solid #edeff2; border-top: 1px solid #edeff2; margin: 0; padding: 0; width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 100%;">
                            <table class="inner-body" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; background-color: #ffffff; margin: 0 auto; padding: 0; width: 570px; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 570px;">
                                <!-- Body content -->
                                <tr>
                                    <td class="content-cell" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; padding: 35px;">
                                        <h1 style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #3d4852; font-size: 19px; font-weight: bold; margin-top: 0; text-align: left;">` + reply_user.Name + ` 你好</h1>
<pre style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;"><code>` + comment_user.Name + `在文章&lt;&lt;` + article.Title + `&gt;&gt;中回复了你:

` + *comment.Content + `</code></pre>
<table class="action" align="center" width="100%" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; margin: 30px auto; padding: 0; text-align: center; width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 100%;">
    <tr>
        <td align="center" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
            <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                <tr>
                    <td align="center" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                        <table border="0" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                            <tr>
                                <td style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                                    <a href="` + article_url + `" class="button button-primary" target="_blank" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; border-radius: 3px; box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16); color: #fff; display: inline-block; text-decoration: none; -webkit-text-size-adjust: none; background-color: #3490dc; border-top: 10px solid #3490dc; border-right: 18px solid #3490dc; border-bottom: 10px solid #3490dc; border-left: 18px solid #3490dc;">` + article.Title + `</a>
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
            </table>
        </td>
    </tr>
</table>
<p style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #3d4852; font-size: 16px; line-height: 1.5em; margin-top: 0; text-align: left;">Thanks,<br>
<a href="` + url + `" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #3869d4;">` + blog_name + `</a></p>

                                        
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>

                    <tr>
    <td style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
        <table class="footer" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; margin: 0 auto; padding: 0; text-align: center; width: 570px; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 570px;">
            <tr>
                <td class="content-cell" align="center" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; padding: 35px;">
                    <p style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; line-height: 1.5em; margin-top: 0; color: #aeaeae; font-size: 12px; text-align: center;">© ` + fmt.Sprintf("%d", year) + blog_name + `. 版本所有。</p>
                </td>
            </tr>
        </table>
    </td>
</tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`
	email := service.Email{
		To:          reply_user.Email,
		ContentType: "text/html",
		Body:        html_raw,
	}
	res := email.SendEmail()
	if !res {
		service.Log.Error("发送邮件失败！")
	} else {
		service.Log.Info("发送邮件成功！")
	}
}

/**
发送评论邮件
*/
func HandleCommentEmailQueue(id int64) {
	comment := Comment{}
	database.Db.Where("id = ?", id).First(&comment)
	if comment.ID == nil {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的评论发送邮件失败，评论不存在")
		return
	}
	user := User{}
	database.Db.Where("id = ?", *comment.UserID).First(&user)
	if user.ID == nil {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的评论发送邮件失败,用户不存在")
		return
	}
	article := Article{}
	database.Db.Where("id = ?", *comment.ArticleID).First(&article)
	if article.ID == nil {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的评论发送邮件失败，文章不存在")
		return
	}
	url := config.GetConfig("app.url").(string)
	blog_name := config.GetConfig("app.name").(string)
	blog_email := config.GetConfig("email.from").(string)
	if url == "" || blog_name == "" || blog_email == "" {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的评论发送邮件失败,配置文件必要参数为空")
		return
	}
	article_url := strings.TrimRight(url, "/") + "/article/" + fmt.Sprintf("%d", *comment.ArticleID)
	year := time.Now().Year()
	html_raw := `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
</head>
<body style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; background-color: #f8fafc; color: #74787e; height: 100%; hyphens: auto; line-height: 1.4; margin: 0; -moz-hyphens: auto; -ms-word-break: break-all; width: 100% !important; -webkit-hyphens: auto; -webkit-text-size-adjust: none; word-break: break-word;">
    <style>
        @media  only screen and (max-width: 600px) {
            .inner-body {
                width: 100% !important;
            }

            .footer {
                width: 100% !important;
            }
        }

        @media  only screen and (max-width: 500px) {
            .button {
                width: 100% !important;
            }
        }
    </style>

    <table class="wrapper" width="100%" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; background-color: #f8fafc; margin: 0; padding: 0; width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 100%;">
        <tr>
            <td align="center" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                <table class="content" width="100%" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; margin: 0; padding: 0; width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 100%;">
                    <tr>
    <td class="header" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; padding: 25px 0; text-align: center;">
        <a href="` + url + `" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #bbbfc3; font-size: 19px; font-weight: bold; text-decoration: none; text-shadow: 0 1px 0 white;">
            ` + blog_name + `
        </a>
    </td>
</tr>

                    <!-- Email Body -->
                    <tr>
                        <td class="body" width="100%" cellpadding="0" cellspacing="0" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; background-color: #ffffff; border-bottom: 1px solid #edeff2; border-top: 1px solid #edeff2; margin: 0; padding: 0; width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 100%;">
                            <table class="inner-body" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; background-color: #ffffff; margin: 0 auto; padding: 0; width: 570px; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 570px;">
                                <!-- Body content -->
                                <tr>
                                    <td class="content-cell" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; padding: 35px;">
                                        <p style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #3d4852; font-size: 16px; line-height: 1.5em; margin-top: 0; text-align: left;">` + user.Name + `在文章&lt;&lt;` + article.Title + `&gt;&gt;中评论了。<br>
评论内容：` + *comment.Content + `</p>
<table class="action" align="center" width="100%" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; margin: 30px auto; padding: 0; text-align: center; width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 100%;">
    <tr>
        <td align="center" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
            <table width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                <tr>
                    <td align="center" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                        <table border="0" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                            <tr>
                                <td style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
                                    <a href="` + article_url + `" class="button button-primary" target="_blank" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; border-radius: 3px; box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16); color: #fff; display: inline-block; text-decoration: none; -webkit-text-size-adjust: none; background-color: #3490dc; border-top: 10px solid #3490dc; border-right: 18px solid #3490dc; border-bottom: 10px solid #3490dc; border-left: 18px solid #3490dc;">` + article.Title + `</a>
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
            </table>
        </td>
    </tr>
</table>
<p style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #3d4852; font-size: 16px; line-height: 1.5em; margin-top: 0; text-align: left;">Thanks,<br>
<a href="` + url + `" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #3869d4;">` + blog_name + `</a></p>

                                        
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>

                    <tr>
    <td style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box;">
        <table class="footer" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; margin: 0 auto; padding: 0; text-align: center; width: 570px; -premailer-cellpadding: 0; -premailer-cellspacing: 0; -premailer-width: 570px;">
            <tr>
                <td class="content-cell" align="center" style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; padding: 35px;">
                    <p style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; line-height: 1.5em; margin-top: 0; color: #aeaeae; font-size: 12px; text-align: center;">©` + fmt.Sprintf("%d", year) + blog_name + `. 版本所有。</p>
                </td>
            </tr>
        </table>
    </td>
</tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`
	email := service.Email{
		To:          blog_email,
		ContentType: "text/html",
		Body:        html_raw,
	}
	res := email.SendEmail()
	if !res {
		service.Log.Error("发送评论邮件失败！")
	} else {
		service.Log.Info("发送评论邮件成功！")
	}
}

/**
获取文章评论数据
article_id 文章ID
*/
func GetArticleComment(article_id int64) []*ArticleComment {
	all_comments := []ArticleComment{}
	new_all_comments := make([]*ArticleComment, 0)
	database.Db.Table("article_comment").
		Joins("left join `user` as `a` on `article_comment`.reply_id = `a`.id left join `user` as `b` on `b`.id = `article_comment`.user_id").
		Select("`article_comment`.*, `a`.`avatar` as `reply_avatar`, `a`.`name` as `reply_name`, `b`.`avatar` as `user_avatar`, `b`.`name` as `username`").
		Where("article_comment.article_id = ?", article_id).Find(&all_comments)
	if len(all_comments) > 0 {
		comments := make(map[int64]*ArticleComment)
		for _, value := range all_comments {
			comment := value
			comments[*value.ID] = &comment
		}
		for _, value := range comments {
			if comments[*value.Pid] != nil {
				comments[*value.Pid].Children = append(comments[*value.Pid].Children, value)
			} else {
				new_all_comments = append(new_all_comments, value)
			}
		}
		sort.Slice(new_all_comments, func(i, j int) bool {
			return new_all_comments[i].CreatedUnix > new_all_comments[j].CreatedUnix
		})
	}
	return new_all_comments
}
