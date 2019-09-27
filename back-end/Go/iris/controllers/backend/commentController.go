package backend

import (
	"blog/models"
	"blog/service"
	"errors"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

/**
获取comment model
*/
func getCommentModel(ctx iris.Context, validates []service.BlogValidate) (models.Comment, error) {
	comment := models.Comment{}
	requestData, err := getRequestData(ctx, validates)
	if err != nil {
		return comment, err
	}
	err = mapstructure.Decode(requestData, &comment)
	if err != nil {
		return comment, errors.New("参数错误！")
	}
	return comment, nil
}

/**
获取评论列表
*/
func GetCommentList(ctx iris.Context) {
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	article_id := ctx.URLParamInt64Default("article_id", 0)
	list := models.GetCommentList(pageNum, pageSize, article_id)
	if len(list) > 0 {
		response.RenderSuccess(ctx, "获取评论列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无评论列表数据", nil)
}

/**
删除评论
*/
func DeleteComment(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"id", "required,gt=0", "评论ID错误！"},
	}
	comment, err := getCommentModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.DeleteComment(&comment)
	if !res {
		response.RenderError(ctx, "删除评论失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除评论成功！", nil)
}

/**
回复评论
*/
func ReplyComment(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"id", "required,gt=0", "评论ID错误！"},
		{"reply_content", "required,gte=2,lte=255", "回复内容在2到255个字符之间错误！"},
	}
	comment, err := getCommentModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.ReplyComment(&comment)
	if !res {
		response.RenderError(ctx, "回复评论失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "回复评论成功！", nil)
}
