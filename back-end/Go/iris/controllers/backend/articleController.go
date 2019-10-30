package backend

import (
	"blog/controllers"
	"blog/models"
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/tidwall/gjson"
)

/**
获取文章列表
*/
func GetArticleList(ctx iris.Context) {
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	category_id := ctx.URLParamInt64Default("category_id", 0)
	search := ctx.URLParamTrim("search")
	list := models.Article{}.GetArticleList(pageNum, pageSize, category_id, search)
	ctx.Gzip(true)
	if list["total"].(int64) > 0 {
		response.RenderSuccess(ctx, "获取文章列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无文章列表数据", nil)
}

/**
添加文章
*/
func AddArticle(ctx iris.Context) {
	article := models.Article{}
	fields := []string{"category_id", "title", "author", "keywords",
		"description", "thumb", "content_html", "content_md", "is_show", "is_top", "is_original"}
	validateFields := []string{"CategoryID", "Title", "Author", "Keywords", "Description",
		"Thumb", "ContentHtml", "ContentMd", "IsShow", "IsTop", "IsOriginal"}
	data, err := controllers.GetRightModel(ctx, &article, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	tag_ids, e := validateTagID(data)
	if e != nil {
		response.RenderError(ctx, e.Error(), nil)
		return
	}
	res := article.AddArticle(tag_ids)
	if !res {
		response.RenderError(ctx, "添加文章失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "添加文章成功！", nil)
}

/**
获取正确的标签ID
*/
func validateTagID(data string) ([]int64, error) {
	tag_ids := []int64{}
	tag_data := gjson.Get(data, "tags").Array()
	if len(tag_data) <= 0 {
		return tag_ids, errors.New("缺少标签字段")
	}
	for _, value := range tag_data {
		val := value.Int()
		if val <= 0 {
			return tag_ids, errors.New("标签参数错误")
		}
		tag_ids = append(tag_ids, val)
	}
	if len(tag_ids) <= 0 {
		return tag_ids, errors.New("标签参数错误")
	}
	return tag_ids, nil
}

/**
修改文章
*/
func UpdateArticle(ctx iris.Context) {
	article := models.Article{}
	fields := []string{"id", "category_id", "title", "author", "keywords",
		"description", "thumb", "content_html", "content_md", "is_show", "is_top", "is_original"}
	validateFields := []string{"ID", "CategoryID", "Title", "Author", "Keywords", "Description",
		"Thumb", "ContentHtml", "ContentMd", "IsShow", "IsTop", "IsOriginal"}
	data, err := controllers.GetRightModel(ctx, &article, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	tag_ids, e := validateTagID(data)
	if e != nil {
		response.RenderError(ctx, e.Error(), nil)
		return
	}
	res := article.UpdateArticle(tag_ids)
	if !res {
		response.RenderError(ctx, "修改文章失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "修改文章成功！", nil)
}

/**
删除文章
*/
func DeleteArticle(ctx iris.Context) {
	article := models.Article{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	_, err := controllers.GetRightModel(ctx, &article, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := article.DeleteArticle()
	if !res {
		response.RenderError(ctx, "删除文章失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除文章成功！", nil)
}

/**
增加文章每日浏览数
*/
func IncrementArticleView(ctx iris.Context) {
	article := models.Article{}
	article.IncrementArticleView()
}
