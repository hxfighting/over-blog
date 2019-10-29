package backend

import (
	"blog/controllers"
	"blog/models"
	"github.com/kataras/iris/v12"
)

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
	chat := models.Chat{}
	fields := []string{"is_show", "content"}
	validateFields := []string{"IsShow", "Content"}
	_, err := controllers.GetRightModel(ctx, &chat, fields, validateFields)
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
	chat := models.Chat{}
	fields := []string{"id", "is_show", "content"}
	validateFields := []string{"ID", "IsShow", "Content"}
	_, err := controllers.GetRightModel(ctx, &chat, fields, validateFields)
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
	chat := models.Chat{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	_, err := controllers.GetRightModel(ctx, &chat, fields, validateFields)
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
