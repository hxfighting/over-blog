package backend

import (
	"blog/models"
	"blog/service"
	"github.com/kataras/iris"
)

/**
获取配置列表
*/
func GetConfigList(ctx iris.Context) {
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	type_id := ctx.URLParamInt64Default("type", 0)
	list := models.GetConfigList(pageNum, pageSize, type_id)
	if list["total"].(int64) == 0 {
		response.RenderError(ctx, "暂无配置列表数据", nil)
		return
	}
	response.RenderSuccess(ctx, "获取配置成功", list)
}

/**
添加配置
*/
func AddConfig(ctx iris.Context) {
	config := models.Config{}
	err := ctx.ReadJSON(&config)
	if err != nil {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	validates := []service.BlogValidate{
		{config.Type, "required,oneof=1 2 3", "配置类型值错误！"},
		{config.Name, "required,gte=2,lte=200", "配置key在2到200个字符之间！"},
		{config.Title, "required,gte=2,lte=200", "配置名称在2到200个字符之间！"},
		{config.Val, "required,gte=2,lte=65535", "配置值在2到65535个字符之间！"},
	}
	err = service.ValidateField(validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	//res := models.AddConfig(&config)
	//if !res["flag"].(bool) {
	//	response.RenderError(ctx, res["msg"].(string), nil)
	//	return
	//}
	//response.RenderSuccess(ctx, res["msg"].(string), nil)
}
