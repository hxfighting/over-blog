package backend

import (
	"blog/models"
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
	if list["total"].(int64) > 0 {
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
	fields := []string{"id"}
	validateFields := []string{"ID"}
	err := getRightModel(ctx, &link, fields, validateFields)
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
	fields := []string{"name", "order", "is_show", "url", "description"}
	validateFields := []string{"Name", "Order", "IsShow", "Description", "Url"}
	err := getRightModel(ctx, &link, fields, validateFields)
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
	fields := []string{"id", "name", "order", "is_show", "url", "description"}
	validateFields := []string{"ID", "Name", "Order", "IsShow", "Description", "Url"}
	err := getRightModel(ctx, &link, fields, validateFields)
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
