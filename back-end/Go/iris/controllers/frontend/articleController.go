package frontend

import (
	"blog/database"
	"blog/models"
	"blog/service"
	template "blog/views"
	"bytes"
	"github.com/kataras/iris"
	"strings"
)

/**
获取文章数据
pgeSize 分页数量
is_search 是否搜索
is_top 文章是否置顶排序
is_show 文章是否显示
tag_id 标签ID
category_id 分类ID
*/
func getArticleData(ctx iris.Context, pageSize, tag_id, category_id int64, is_search, is_top, is_show bool) ([]models.Article,
	int64, int64, string) {
	pageNum := ctx.URLParamInt64Default("page", 1)
	var db = database.Db.Table("article")
	var search string
	if is_search {
		search = strings.Trim(ctx.URLParamEscape("search"), "")
		if search != "" {
			db = db.Where("title like ?", "%"+search+"%")
		}
	} else {
		search = ""
	}
	if tag_id > 0 {
		db = db.Where("exists ?", database.Db.
			Select("tag.*").
			Table("article_tag").
			Joins("inner join tag on tag.id = article_tag.tag_id").
			Where("`article`.`id` = `article_tag`.`article_id` and tag.id = ?", tag_id).SubQuery())
	}
	if category_id > 0 {
		db = db.Where("category_id = ?", category_id)
	}
	articles, total := models.GetArticles(db, pageNum, pageSize, is_top, is_show)
	return articles, total, pageNum, search
}

/**
文章详情页
*/
func GetArticleDetail(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	article, e := models.GetArticleById(id)
	if e != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("404")
		return
	}
	pre_article, _ := models.GetPreOrNextArticle(id, true)
	after_article, _ := models.GetPreOrNextArticle(id, false)
	rand_articles := models.GerRandArticle(10)
	comments := models.GetArticleComment(id)
	models.CacheArticleView(id)
	buffer := new(bytes.Buffer)
	template.ArticleInfo(article, pre_article, after_article, rand_articles, comments, buffer)
	_, err := ctx.Write(buffer.Bytes())
	if err != nil {
		service.Log.Error(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
	}
}

/**
文章搜索
*/
func SearchArticle(ctx iris.Context) {
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	var pageSize int64 = 6
	articles, total, pageNum, search := getArticleData(ctx, pageSize, 0, 0,
		true, false, true)
	if search != "" {
		search_rune := []rune(search)
		if len(search_rune) > 20 {
			search = string(search_rune[0:20]) + "..."
		} else {
			search = string(search_rune[:])
		}
	}
	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage = total/pageSize + 1
	}
	buffer := new(bytes.Buffer)
	template.ArticleByCategory(articles, totalPage, pageNum, search, "", buffer)
	_, err := ctx.Write(buffer.Bytes())
	if err != nil {
		service.Log.Error(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
	}
}

/**
根据标签获取文章列表
*/
func GetArticleByTag(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	var pageSize int64 = 6
	articles, total, pageNum, _ := getArticleData(ctx, pageSize, id, 0,
		false, false, true)
	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage = total/pageSize + 1
	}
	tag := template.Tag
	var tag_name string
	for _, value := range tag {
		if *value.ID == id {
			tag_name = value.Name
		}
	}
	buffer := new(bytes.Buffer)
	template.ArticleByCategory(articles, totalPage, pageNum, "", tag_name, buffer)
	_, err := ctx.Write(buffer.Bytes())
	if err != nil {
		service.Log.Error(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
	}
}

/**
根据分类获取文章
*/
func GetArticleByCategory(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	var pageSize int64 = 6
	articles, total, pageNum, _ := getArticleData(ctx, pageSize, 0, id,
		false, false, true)
	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage = total/pageSize + 1
	}
	buffer := new(bytes.Buffer)
	template.ArticleByCategory(articles, totalPage, pageNum, "", "", buffer)
	_, err := ctx.Write(buffer.Bytes())
	if err != nil {
		service.Log.Error(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
	}
}
