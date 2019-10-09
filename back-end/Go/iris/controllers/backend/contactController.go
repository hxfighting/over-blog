package backend

import (
	"blog/models"
	"blog/service"
	"errors"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

func getContactModel(ctx iris.Context, validates []service.BlogValidate) (models.Contact, error) {
	contact := models.Contact{}
	requestData, err := getRequestData(ctx, validates)
	if err != nil {
		return contact, err
	}
	err = mapstructure.Decode(requestData, &contact)
	if err != nil {
		return contact, errors.New("参数错误！")
	}
	return contact, nil
}

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
	validates := []service.BlogValidate{
		{"id", "int64", "required,myString,gt=0", "留言ID错误！"},
	}
	contact, err := getContactModel(ctx, validates)
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
	validates := []service.BlogValidate{
		{"id", "int64", "required,myString,gt=0", "留言ID错误！"},
		{"reply_content", "string", "required,myString,gte=2,lte=255", "留言内容在2到255个字符之间！"},
	}
	contact, err := getContactModel(ctx, validates)
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
