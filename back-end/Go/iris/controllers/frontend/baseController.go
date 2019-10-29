package frontend

import (
	"blog/config"
	"blog/database"
	"blog/helper"
	"blog/models"
	"blog/service"
	template "blog/views"
	"fmt"
	"github.com/kataras/iris/v12/sessions"
	"sort"
	"time"
)

var (
	CONFIG_KEY             = "blog_config"
	ROTATION_KEY           = "blog_rotation"
	PHOTO_KEY              = "blog_photo"
	CATEGORY_KEY           = "blog_category"
	SOCIAL_KEY             = "blog_social"
	TAG_KEY                = "blog_tag"
	ARTICLE_KEY            = "blog_article"
	COMMENT_KEY            = "blog_comment"
	LINK_KEY               = "blog_link"
	FOOTER_KEY             = "blog_footer"
	cookieNameForSessionID = "3f904b9b9b6606cd5df65e8b0b1e4b53"
	Sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
	Response               = service.Response{}
)

/**
初始化数据
*/
func InitData() {
	template.Config = GetWebConfigContent()
	template.Category = getCategoryData()
	template.SocialData = getSocialData()
	template.Tag = getTagData()
	template.HotArticle = getHotArticle()
	template.RecentComment = getRecentComments()
	template.Link = getLinkData()
	template.FooterData = getFooterData()
	template.AppUrl = config.GetConfig("app.url").(string)
}

/**
获取配置内容
*/
func GetWebConfigContent() map[string]string {
	s, e := service.Redis.Get(CONFIG_KEY).Result()
	if e != nil {
		configs := []models.Config{}
		database.Db.Select("name,val").Where("type = ?", 3).Find(&configs)
		if len(configs) > 0 {
			data := make(map[string]string)
			for _, value := range configs {
				data[*value.Name] = *value.Val
			}
			res, _ := service.FastJson.Marshal(&data)
			s = string(res)
			service.Redis.Set(CONFIG_KEY, s, 0)
			return data
		}
		return map[string]string{}
	} else {
		config_data := make(map[string]string)
		service.FastJson.Unmarshal([]byte(s), &config_data)
		return config_data
	}
}

/**
获取分类数据
*/
func getCategoryData() []*models.Category {
	s, e := service.Redis.Get(CATEGORY_KEY).Result()
	if e != nil {
		all_data := make([]*models.Category, 0)
		category_list := []models.Category{}
		database.Db.Find(&category_list)
		if len(category_list) > 0 {
			build_data := make(map[int64]*models.Category)
			for _, value := range category_list {
				category := value
				build_data[*value.ID] = &category
			}
			for _, value := range build_data {
				if build_data[*value.Pid] != nil {
					build_data[*value.Pid].Children = append(build_data[*value.Pid].Children, value)
				} else {
					all_data = append(all_data, value)
				}
			}
			sort.Slice(all_data, func(i, j int) bool {
				return all_data[i].CreatedAt < all_data[j].CreatedAt
			})
			res, _ := service.FastJson.Marshal(&all_data)
			s = string(res)
			service.Redis.Set(CATEGORY_KEY, s, 0)
			return all_data
		}
		return all_data
	} else {
		category_data := make([]*models.Category, 0)
		service.FastJson.Unmarshal([]byte(s), &category_data)
		return category_data
	}
}

/**
获取社交地址
*/
func getSocialData() []models.Config {
	s, e := service.Redis.Get(SOCIAL_KEY).Result()
	if e != nil {
		configs := []models.Config{}
		database.Db.Where("type = ?", 1).Find(&configs)
		if len(configs) > 0 {
			res, _ := service.FastJson.Marshal(&configs)
			s = string(res)
			service.Redis.Set(SOCIAL_KEY, s, 0)
		}
		return configs
	} else {
		config_data := make([]models.Config, 0)
		service.FastJson.Unmarshal([]byte(s), &config_data)
		return config_data
	}
}

/**
获取tag
*/
func getTagData() []models.Tag {
	s, e := service.Redis.Get(TAG_KEY).Result()
	if e != nil {
		tag := []models.Tag{}
		database.Db.Find(&tag)
		if len(tag) > 0 {
			res, _ := service.FastJson.Marshal(&tag)
			s = string(res)
			service.Redis.Set(TAG_KEY, s, 0)
		}
		return tag
	} else {
		tag := make([]models.Tag, 0)
		service.FastJson.Unmarshal([]byte(s), &tag)
		return tag
	}
}

/**
获取热门文章
*/
func getHotArticle() []map[string]string {
	s, e := service.Redis.Get(ARTICLE_KEY).Result()
	if e != nil {
		articles := []models.Article{}
		database.Db.Table("article").
			Where("is_show = ?", 1).
			Select("article.*,(select count(*) from article_comment where article_comment.article_id = article.id) as comments_count").
			Order("click desc,created_at desc").
			Limit(10).Find(&articles)
		data := make([]map[string]string, 0)
		if len(articles) > 0 {
			end_time, err := helper.GetUnixTimeFromDate(time.Now().Format(helper.YMD)+" 23:59:59", helper.YMDHIS)
			if err == nil {
				for _, value := range articles {
					da := map[string]string{
						"id":            fmt.Sprintf("%d", *value.ID),
						"comment_count": fmt.Sprintf("%d", value.CommentsCount),
						"title":         value.Title,
						"click":         fmt.Sprintf("%d", value.Click),
						"created_at":    value.CreatedAt[:10],
						"created_unix":  fmt.Sprintf("%d", value.CreatedUnix),
					}
					data = append(data, da)
				}
				res, _ := service.FastJson.Marshal(&data)
				s = string(res)
				expire := time.Unix(end_time, 0).Sub(time.Now())
				service.Redis.Set(ARTICLE_KEY, s, expire)
			}
		}
		return data
	} else {
		article := make([]map[string]string, 0)
		service.FastJson.Unmarshal([]byte(s), &article)
		return article
	}
}

/**
获取最新评论
*/
func getRecentComments() []models.RecentComment {
	s, e := service.Redis.Get(COMMENT_KEY).Result()
	if e != nil {
		comments := []models.RecentComment{}
		database.Db.Table("article_comment").
			Joins("left join user on user.id = article_comment.user_id").
			Select("article_comment.article_id,article_comment.content,article_comment.created_at,user.avatar,user.name").
			Limit(10).Order("created_at desc").Find(&comments)
		if len(comments) > 0 {
			res, _ := service.FastJson.Marshal(&comments)
			s = string(res)
			service.Redis.Set(COMMENT_KEY, s, 0)
		}
		return comments
	} else {
		comments := make([]models.RecentComment, 0)
		service.FastJson.Unmarshal([]byte(s), &comments)
		return comments
	}
}

/**
获取友联
*/
func getLinkData() []models.SimpleLink {
	s, e := service.Redis.Get(LINK_KEY).Result()
	if e != nil {
		links := []models.SimpleLink{}
		database.Db.Table("link").
			Where("is_show = ?", 1).
			Select("url,name,description").
			Order("`order` desc,created_at asc").Find(&links)
		if len(links) > 0 {
			res, _ := service.FastJson.Marshal(&links)
			s = string(res)
			service.Redis.Set(LINK_KEY, s, 0)
		}
		return links
	} else {
		links := make([]models.SimpleLink, 0)
		service.FastJson.Unmarshal([]byte(s), &links)
		return links
	}
}

/**
获取footer内容
*/
func getFooterData() map[string]map[int64]models.SimpleConfig {
	s, e := service.Redis.Get(FOOTER_KEY).Result()
	if e != nil {
		data := make(map[string]map[int64]models.SimpleConfig)
		configs := []models.Config{}
		database.Db.Where("type = ?", 2).Find(&configs)
		if len(configs) > 0 {
			for _, value := range configs {
				if data[*value.Name] == nil {
					da := map[int64]models.SimpleConfig{}
					da[*value.ID] = models.SimpleConfig{
						Title: *value.Title,
						Val:   *value.Val,
						Name:  *value.Name,
					}
					data[*value.Name] = da
				} else {
					data[*value.Name][*value.ID] = models.SimpleConfig{
						Title: *value.Title,
						Val:   *value.Val,
						Name:  *value.Name,
					}
				}
			}
			res, _ := service.FastJson.Marshal(&data)
			s = string(res)
			service.Redis.Set(FOOTER_KEY, s, 0)
		}
		return data
	} else {
		config_data := make(map[string]map[int64]models.SimpleConfig)
		service.FastJson.Unmarshal([]byte(s), &config_data)
		return config_data
	}
}
