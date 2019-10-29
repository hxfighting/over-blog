package models

import (
	"blog/database"
	"blog/helper"
	"blog/service"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"strings"
	"time"
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

func (this *Article) AfterFind() {
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取文章列表
*/
func (this Article) GetArticleList(pageNum, pageSize, category_id int64, search string) map[string]interface{} {
	article_category := []simpleCategory{}
	database.Db.Table("category").
		Select("id,title").
		Where("type = ?", 1).
		Find(&article_category)
	tags := []simpleTag{}
	database.Db.Table("tag").
		Select("id,name").
		Find(&tags)
	db := database.Db.Table("article")
	if search != "" {
		db = db.Where("title like ?", "%"+search+"%")
	}
	if category_id != 0 {
		db = db.Where("category_id = ?", category_id)
	}
	articles, total := GetArticles(db, pageNum, pageSize, false, false)
	data := make(map[string]interface{})
	data["category"] = article_category
	data["tag"] = tags
	data["list"] = articles
	data["total"] = total
	return data
}

/**
获取文章数据
*/
func GetArticles(db *gorm.DB, pageNum, pageSize int64, is_top, is_show bool) ([]Article, int64) {
	articles := []Article{}
	var total int64 = 0
	if is_show {
		db = db.Where("is_show = ?", 1)
	}
	db.Count(&total)
	if is_top {
		db = db.Order("is_top desc,created_at desc")
	} else {
		db = db.Order("created_at desc")
	}
	db.Select("article.*,(select count(*) from article_comment where article_comment.article_id = article.id) as comments_count").
		Offset((pageNum - 1) * pageSize).Limit(pageSize).
		Find(&articles)
	if total > 0 {
		article_ids := []int64{}
		category_ids := []int64{}
		for _, value := range articles {
			article_ids = append(article_ids, *value.ID)
			category_ids = append(category_ids, *value.CategoryID)
		}
		article_tags := []simpleTag{}
		database.Db.Table("tag").Select("tag.id,tag.name,article_tag.article_id").
			Where("article_tag.article_id in (?)", article_ids).
			Joins("left join article_tag on article_tag.tag_id = tag.id").Find(&article_tags)
		if len(article_tags) > 0 {
			for k, value := range articles {
				for _, val := range article_tags {
					if val.ArticleID == *value.ID {
						value.Tags = append(value.Tags, val)
						articles[k] = value
					}
				}
			}
		}
		article_categorys := []simpleCategory{}
		database.Db.Table("category").Select("id,title").
			Where("id in (?)", category_ids).Find(&article_categorys)
		if len(article_categorys) > 0 {
			for kk, value := range articles {
				for _, val := range article_categorys {
					if *val.ID == *value.CategoryID {
						value.Category = val
						articles[kk] = value
					}
				}
			}
		}
	}
	return articles, total
}

/**
根据文章id获取文章
*/
func GetArticleById(id int64) (Article, error) {
	article := Article{}
	database.Db.Where("is_show = ? and id = ?", 1, id).
		Select("article.*,(select count(*) from article_comment where article_comment.article_id = article.id) as comments_count").
		First(&article)
	if article.ID == nil {
		return article, errors.New("文章不存在！")
	}
	article_tag := []simpleTag{}
	article_category := simpleCategory{}
	database.Db.Table("tag").Select("tag.id,tag.name,article_tag.article_id").
		Where("article_tag.article_id = ?", *article.ID).
		Joins("left join article_tag on article_tag.tag_id = tag.id").Find(&article_tag)
	if len(article_tag) > 0 {
		article.Tags = article_tag
	}
	database.Db.Table("category").Select("id,title").
		Where("id = ?", *article.CategoryID).Find(&article_category)
	article.Category = article_category
	return article, nil
}

/**
获取上一篇或下一篇文章
flag true 上一篇文章
flag false 下一篇文章
*/
func GetPreOrNextArticle(id int64, flag bool) (Article, error) {
	article := Article{}
	db := database.Db.Table("article").Select("id,title")
	if flag {
		db = db.Order("id desc").Where("id < ? and is_show = ?", id, 1)
	} else {
		db = db.Order("id asc").Where("id > ? and is_show = ?", id, 1)
	}
	db.First(&article)
	if article.ID != nil {
		return article, nil
	}
	return article, errors.New("暂无文章数据")
}

/**
获取随机文章
num 文章数
*/
func GerRandArticle(num int64) []SimpleArticle {
	articles := []SimpleArticle{}
	database.Db.Table("article").
		Where("is_show = ?", 1).
		Order("RAND()").Limit(num).Find(&articles)
	return articles
}

/**
删除文章
*/
func (this Article) DeleteArticle() bool {
	exist_article := Article{}
	database.Db.First(&exist_article, *this.ID)
	if exist_article.ID == nil {
		return false
	}
	tx := database.Db.Begin()
	res := tx.Delete(&this)
	if res.Error != nil {
		service.Log.Error(res.Error.Error())
		tx.Rollback()
		return false
	}
	res = database.Db.Table("article_tag").Where("article_id = ?", *exist_article.ID).Delete(Article{})
	if res.Error != nil {
		service.Log.Error(res.Error.Error())
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

/**
添加文章
*/
func (this Article) AddArticle(tag_ids []int64) bool {
	tx := database.Db.Begin()
	this.CreatedUnix = time.Now().Unix()
	this.UpdatedUnix = time.Now().Unix()
	if this.Description == "" {
		this.Description = getDescription(this.ContentHtml)
	}
	res := tx.Create(&this)
	if e := res.Error; e != nil {
		service.Log.Error(e.Error())
		tx.Rollback()
		return false
	}
	tag_sql := "INSERT INTO article_tag (`article_id`,`tag_id`) VALUES "
	for k, value := range tag_ids {
		if len(tag_ids)-1 == k {
			tag_sql += fmt.Sprintf("(%d,%d);", *this.ID, value)
		} else {
			tag_sql += fmt.Sprintf("(%d,%d),", *this.ID, value)
		}
	}
	if e := tx.Exec(tag_sql).Error; e != nil {
		service.Log.Error(e.Error())
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

/**
获取前正文前200个字符
*/
func getDescription(content string) string {
	da := bluemonday.StrictPolicy().Sanitize(content)
	da_bytes := []rune(da)
	if len(da_bytes) > 200 {
		return string(da_bytes[:200])
	} else {
		return string(da_bytes[:len(da_bytes)-1])
	}
}

/**
修改文章内容
*/
func (this Article) UpdateArticle(tag_ids []int64) bool {
	tx := database.Db.Begin()
	if this.Description == "" {
		this.Description = getDescription(this.ContentHtml)
	}
	data := map[string]interface{}{
		"is_show":      *this.IsShow,
		"is_top":       *this.IsTop,
		"is_original":  *this.IsOriginal,
		"title":        this.Title,
		"description":  this.Description,
		"author":       this.Author,
		"thumb":        this.Thumb,
		"updated_at":   time.Now().Unix(),
		"category_id":  *this.CategoryID,
		"content_html": this.ContentHtml,
		"content_md":   this.ContentMd,
		"keywords":     this.Keywords,
	}
	result := database.Db.Model(&this).Updates(data)
	if result.Error != nil {
		service.Log.Error(result.Error.Error())
		tx.Rollback()
		return false
	}
	res := tx.Table("article_tag").Where("article_id = ?", *this.ID).Delete(Article{})
	if res.Error != nil {
		service.Log.Error(res.Error.Error())
		tx.Rollback()
		return false
	}
	tag_sql := "INSERT INTO article_tag (`article_id`,`tag_id`) VALUES "
	for k, value := range tag_ids {
		if len(tag_ids)-1 == k {
			tag_sql += fmt.Sprintf("(%d,%d);", *this.ID, value)
		} else {
			tag_sql += fmt.Sprintf("(%d,%d),", *this.ID, value)
		}
	}
	if e := tx.Exec(tag_sql).Error; e != nil {
		service.Log.Error(e.Error())
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

/**
缓存文章浏览数
*/
func CacheArticleView(id int64) {
	field := "article_" + fmt.Sprintf("%d", id)
	by := service.Redis.HIncrBy(ARTICLE_VIEW, field, 1)
	if by.Err() != nil {
		service.Log.Error(by.Err().Error())
	}
}

/**
增加文章每日浏览数
*/
func (this *Article) IncrementArticleView() {
	log.Println(fmt.Sprintf("%x", md5.Sum([]byte("测试你哦"))))
	all := service.Redis.HGetAll(ARTICLE_VIEW)
	if all.Err() != nil {
		return
	}
	res := all.Val()
	if len(res) <= 0 {
		return
	}
	sql := "UPDATE `article` SET `click` = CASE id"
	id_str := ""
	for key, val := range res {
		id := key[8:]
		id_str += id + ","
		sql += " WHEN " + id + " THEN `click`+" + val + " "
	}
	id_str = strings.Trim(id_str, ",")
	sql += " END WHERE id in (" + id_str + ")"
	exec := database.Db.Exec(sql)
	if exec.Error != nil {
		service.Log.Error(exec.Error.Error())
		return
	}
	service.Redis.Del(ARTICLE_VIEW)
	service.Log.Info("文章浏览数更新成功！")
}
