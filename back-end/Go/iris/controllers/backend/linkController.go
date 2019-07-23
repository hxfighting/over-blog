package backend

import (
	"blog/models"
	"blog/service"
	"github.com/kataras/iris"
)

/**
获取友联列表
*/
func GetLinkList(ctx iris.Context) {
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	name := ctx.URLParamTrim("name")
	list := models.GetLinkList(pageNum, pageSize, name)
	if len(list) > 0 {
		response.RenderSuccess(ctx, "获取友联列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无友联列表数据", nil)
}

/**
删除友联
*/
func DeleteLink(ctx iris.Context) {
	link := models.Link{}
	err := ctx.ReadJSON(&link)
	if err != nil {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	validates := []service.BlogValidate{
		{link.ID, "required,gt=0", "友联ID错误！"},
	}
	err = service.ValidateField(validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.DeleteLink(&link)
	if !res {
		response.RenderError(ctx, "删除友联失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除友联成功！", nil)
}

/**
添加友联
*/
func AddLink(ctx iris.Context) {
	link := models.Link{}
	err := ctx.ReadJSON(&link)
	if err != nil {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	validates := []service.BlogValidate{
		{link.Name, "required,gte=2,lte=30", "友联描述在2到30个字符之间！"},
		{link.Order, "required,gte=0", "友联排序值错误！"},
		{link.IsShow, "required,oneof=0 1", "友联是否显示值错误！"},
		{link.Url, "required,url", "友联URL错误！"},
		{link.Description, "required,gte=2,lte=50", "友联描述在2到50个字符之间！"},
	}
	err = service.ValidateField(validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.AddLink(&link)
	if !res {
		response.RenderError(ctx, "添加友联失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "添加友联成功！", nil)
}

/**
修改友联
*/
func UpdateLink(ctx iris.Context) {
	link := models.Link{}
	err := ctx.ReadJSON(&link)
	if err != nil {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	validates := []service.BlogValidate{
		{link.ID, "required,gt=0", "友联ID错误！"},
		{link.Name, "required,gte=2,lte=30", "友联描述在2到30个字符之间！"},
		{link.Order, "required,gte=0", "友联排序值错误！"},
		{link.IsShow, "required,oneof=0 1", "友联是否显示值错误！"},
		{link.Url, "required,url", "友联URL错误！"},
		{link.Description, "required,gte=2,lte=50", "友联描述在2到50个字符之间！"},
	}
	err = service.ValidateField(validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.UpdateLink(&link)
	if !res {
		response.RenderError(ctx, "修改友联失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "修改友联成功！", nil)
}
