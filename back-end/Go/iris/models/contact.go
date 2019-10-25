package models

import (
	"blog/config"
	"blog/database"
	"blog/helper"
	"blog/service"
	"fmt"
	"time"
)

type Contact struct {
	ID           *int64  `json:"id" validate:"gt=0"`
	Content      *string `json:"content" validate:"gte=2,lte=255"`
	Name         *string `json:"name" validate:"gte=2,lte=20"`
	Email        *string `json:"email" validate:"email"`
	IsReply      *int64  `json:"is_reply" mapstructure:"is_reply"`
	CreatedUnix  int64   `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix  int64   `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt    string  `json:"created_at" gorm:"-"`
	UpdatedAt    string  `json:"updated_at" gorm:"-"`
	ReplyContent *string `json:"reply_content" mapstructure:"reply_content" validate:"gte=2,lte=255"`
}

func (Contact) TableName() string {
	return "contact"
}

func (this *Contact) AfterFind() {
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取留言列表
*/
func (this Contact) GetContactList(search string, pageNum, pageSize int64) map[string]interface{} {
	contact := []Contact{}
	var db = database.Db.Table("contact")
	if search != "" {
		db = db.Where("name like ?", "%"+search+"%").Or("email like ?", "%"+search+"%")
	}
	var total int64 = 0
	db.Count(&total)
	db.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&contact)
	data := make(map[string]interface{})
	data["total"] = total
	data["list"] = contact
	return data
}

/**
删除留言
*/
func (this Contact) DeleteContact() bool {
	res := database.Db.Delete(&this)
	if res.Error != nil {
		return false
	}
	return true
}

/**
回复留言
*/
func (this Contact) ReplyContact() bool {
	exist_contact := Contact{}
	database.Db.First(&exist_contact, *this.ID)
	if exist_contact.ID == nil {
		return false
	}
	data := map[string]interface{}{
		"is_reply":      1,
		"reply_content": *this.ReplyContent,
		"updated_at":    time.Now().Unix(),
	}
	result := database.Db.Model(&this).Updates(data)
	if result.Error != nil {
		return false
	}
	service.ReplyEmailChan <- *exist_contact.ID
	return true
}

/**
添加留言
*/
func (this Contact) AddContact() bool {
	this.CreatedUnix = time.Now().Unix()
	this.UpdatedUnix = time.Now().Unix()
	res := database.Db.Create(&this)
	if res.Error != nil {
		return false
	}
	service.ReplyEmailChan <- *this.ID
	return true
}

/**
发送留言回复邮件
*/
func HandleReplyContactEmail(id int64) {
	blog_name := config.GetConfig("app.name").(string)
	url := config.GetConfig("app.url").(string)
	if url == "" {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的留言回复发送邮件失败")
		return
	}
	contact := Contact{}
	database.Db.First(&contact, id)
	if contact.ID == nil {
		service.Log.Error("ID 为" + fmt.Sprintf("%d", id) + "的留言回复发送邮件失败,留言不存在！")
		return
	}
	var content string
	if contact.ReplyContent == nil {
		content = "非常感谢你的留言,我会尽快回复你的."
	} else {
		content = *contact.ReplyContent
	}
	year := time.Now().Year()
	html := `<html xmlns="http://www.w3.org/1999/xhtml">
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
                                        <h1 style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #3d4852; font-size: 19px; font-weight: bold; margin-top: 0; text-align: left;">` + *contact.Name + ` 你好</h1>
<p style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol'; box-sizing: border-box; color: #3d4852; font-size: 16px; line-height: 1.5em; margin-top: 0; text-align: left;">` + content + `</p>
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
		To:          *contact.Email,
		ContentType: "text/html",
		Body:        html,
	}
	res := email.SendEmail()
	if !res {
		service.Log.Error("发送留言回复邮件失败！")
	} else {
		service.Log.Info("发送留言回复邮件成功！")
	}
}
