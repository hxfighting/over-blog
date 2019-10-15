package backend

import (
	"blog/models"
	"github.com/kataras/iris"
)

/**
获取留言列表
*/
func GetContactList(ctx iris.Context) {
	search := ctx.URLParamTrim("search")
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	contacts := models.Contact{}.GetContactList(search, pageNum, pageSize)
	if contacts["total"].(int64) > 0 {
		response.RenderSuccess(ctx, "获取留言列表成功", contacts)
		return
	}
	response.RenderError(ctx, "暂无留言列表数据", nil)
}

/**
删除留言
*/
func DeleteContact(ctx iris.Context) {
	contact := models.Contact{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	_, err := getRightModel(ctx, &contact, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := contact.DeleteContact()
	if !res {
		response.RenderError(ctx, "删除留言失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除留言成功！", nil)
}

/**
回复留言
*/
func ReplyContact(ctx iris.Context) {
	contact := models.Contact{}
	fields := []string{"id", "reply_content"}
	validateFields := []string{"ID", "ReplyContent"}
	_, err := getRightModel(ctx, &contact, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := contact.ReplyContact()
	if !res {
		response.RenderError(ctx, "回复留言失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "回复留言成功！", nil)
}
