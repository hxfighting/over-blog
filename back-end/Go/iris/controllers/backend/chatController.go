package backend

import (
	"blog/models"
	"blog/service"
	"errors"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

/**
获取chat model
*/
func getChatModel(ctx iris.Context, validates []service.BlogValidate) (models.Chat, error) {
	chat := models.Chat{}
	requestData, err := getRequestData(ctx, validates)
	if err != nil {
		return chat, err
	}
	err = mapstructure.Decode(requestData, &chat)
	if err != nil {
		return chat, errors.New("参数错误！")
	}
	return chat, nil
}

/**
获取说说列表
*/
func GetChatList(ctx iris.Context) {
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	list := models.GetChatList(pageNum, pageSize)
	if len(list) > 0 {
		response.RenderSuccess(ctx, "获取说说列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无说说列表数据", nil)
}

/**
添加说说
*/
func AddChat(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"is_show", "uint", "required,myString,oneof=0 1", "说说是否显示值错误！"},
		{"content", "string", "required,myString,gte=2,lte=255", "说说内容在2到255个字符之间！"},
	}
	chat, err := getChatModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.AddChat(&chat)
	if !res {
		response.RenderError(ctx, "添加说说失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "添加说说成功！", nil)
}

/**
修改说说
*/
func UpdateChat(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"id", "int64", "required,myString,gt=0", "说说ID错误！"},
		{"is_show", "uint", "required,myString,oneof=0 1", "说说是否显示值错误！"},
		{"content", "string", "required,myString,gte=2,lte=255", "说说内容在2到255个字符之间！"},
	}
	chat, err := getChatModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.UpdateChat(&chat)
	if !res {
		response.RenderError(ctx, "修改说说失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "修改说说成功！", nil)
}

/**
删除说说
*/
func DeleteChat(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"id", "int64", "required,myString,gt=0", "说说ID错误！"},
	}
	chat, err := getChatModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.DeleteChat(&chat)
	if !res {
		response.RenderError(ctx, "删除说说失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除说说成功！", nil)
}
