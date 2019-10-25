package frontend

import (
	"blog/config"
	"blog/database"
	"blog/models"
	"blog/service"
	template "blog/views"
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"unicode/utf8"
)

/**
博客首页
*/
func Index(ctx iris.Context) {
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	var pageSize int64 = 5
	articles, total, pageNum, _ := getArticleData(ctx, pageSize, 0, 0,
		false, true, true)
	rotation := getRotation()
	photo := getPhoto()
	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage = total/pageSize + 1
	}
	buffer := new(bytes.Buffer)
	template.Index(articles, totalPage, pageNum, rotation, photo, buffer)
	_, err := ctx.Write(buffer.Bytes())
	if err != nil {
		service.Log.Error(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
	}
}

/**
获取照片
*/
func getPhoto() []map[string]string {
	utf8.RuneCountInString("世界")
	s, e := service.Redis.Get(PHOTO_KEY).Result()
	if e != nil {
		ph := models.Photo{}
		photo := ph.GetPhotoList()
		if len(photo) > 0 {
			all_data := []map[string]string{}
			for _, value := range photo {
				data := map[string]string{
					"id":        fmt.Sprintf("%d", *value.ID),
					"image_id":  fmt.Sprintf("%d", *value.Image_id),
					"image_url": *value.Image_url,
				}
				all_data = append(all_data, data)
			}
			res, _ := service.FastJson.Marshal(all_data)
			service.Redis.Set(PHOTO_KEY, string(res), 0)
			return all_data
		}
		return []map[string]string{}
	} else {
		photo_data := make([]map[string]string, 0)
		service.FastJson.Unmarshal([]byte(s), &photo_data)
		return photo_data
	}
}

/**
获取轮播图
*/
func getRotation() []map[string]string {
	s, e := service.Redis.Get(ROTATION_KEY).Result()
	if e != nil {
		ro := models.Rotation{}
		rotation := ro.GetRotationList()
		if len(rotation) > 0 {
			all_data := []map[string]string{}
			for _, value := range rotation {
				i := value.Image.(models.Rotation)
				data := map[string]string{
					"id":        fmt.Sprintf("%d", *value.ID),
					"image_id":  fmt.Sprintf("%d", *value.Image_id),
					"image_url": *value.Image_url,
					"words":     i.Words,
				}
				all_data = append(all_data, data)
			}
			res, _ := service.FastJson.Marshal(all_data)
			service.Redis.Set(ROTATION_KEY, string(res), 0)
			return all_data
		}
		return []map[string]string{}
	} else {
		rotation_data := make([]map[string]string, 0)
		service.FastJson.Unmarshal([]byte(s), &rotation_data)
		return rotation_data
	}
}

/**
从百度统计获取博客每日UV
*/
func GetBlogCount(ctx iris.Context) {
	url := config.GetConfig("tongji.url").(string)
	username := config.GetConfig("tongji.username").(string)
	password := config.GetConfig("tongji.password").(string)
	token := config.GetConfig("tongji.token").(string)
	siteID := config.GetConfig("tongji.siteId").(string)
	if url == "" || username == "" || password == "" || token == "" || siteID == "" {
		service.Log.Error("缺少统计URL配置")
		return
	}
	data := make(map[string]map[string]string)
	data["header"] = make(map[string]string)
	data["body"] = make(map[string]string)
	data["header"]["username"] = username
	data["header"]["password"] = password
	data["header"]["token"] = token
	data["header"]["account_type"] = "1"
	data["body"]["site_id"] = siteID
	data["body"]["method"] = "overview/getOutline"
	json_byte, e := service.FastJson.Marshal(&data)
	if e != nil {
		service.Log.Error(e.Error())
		return
	}
	client := &http.Client{}
	request, e := http.NewRequest("POST", url, bytes.NewReader(json_byte))
	if e != nil {
		service.Log.Error(e.Error())
		return
	}
	response, e := client.Do(request)
	if e != nil {
		service.Log.Error(e.Error())
		return
	}
	respBody, e := ioutil.ReadAll(response.Body)
	if e != nil {
		service.Log.Error(e.Error())
		return
	}
	result := string(respBody)
	uv := gjson.Get(result, "body.data.0.result.items.1.3").Int()
	if uv > 0 {
		res := database.Db.Table("web_config").
			Where("name = ?", "blog_view_count").
			Update("val", gorm.Expr("val + ?", uv))
		if res.Error != nil {
			service.Log.Error(res.Error.Error())
			return
		}
	}
}
