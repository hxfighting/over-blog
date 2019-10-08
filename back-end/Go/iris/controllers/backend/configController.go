package backend

import (
	"blog/models"
	"blog/service"
	"errors"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

/**
获取config model
*/
func getConfigModel(ctx iris.Context, validates []service.BlogValidate) (models.Config, error) {
	config := models.Config{}
	requestData, err := getRequestData(ctx, validates)
	if err != nil {
		return config, err
	}
	err = mapstructure.Decode(requestData, &config)
	if err != nil {
		return config, errors.New("参数错误！")
	}
	return config, nil
}

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
	validates := []service.BlogValidate{
		{"type", "int64", "required,myString,oneof=1 2 3", "配置类型值错误！"},
		{"name", "string", "required,myString,gte=2,lte=200", "配置key在2到200个字符之间！"},
		{"title", "string", "required,myString,gte=2,lte=200", "配置名称在2到200个字符之间！"},
		{"val", "string", "required,myString,gte=2,lte=65535", "配置值在2到65535个字符之间！"},
	}
	config, err := getConfigModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.AddConfig(&config)
	if !res["flag"].(bool) {
		response.RenderError(ctx, res["msg"].(string), nil)
		return
	}
	response.RenderSuccess(ctx, res["msg"].(string), nil)
}

/**
修改配置
*/
func UpdateConfig(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"id", "int64", "required,myString,gt=0", "配置ID错误！"},
		{"type", "int64", "required,myString,oneof=1 2 3", "配置类型值错误！"},
		{"name", "string", "required,myString,gte=2,lte=200", "配置key在2到200个字符之间！"},
		{"title", "string", "required,myString,gte=2,lte=200", "配置名称在2到200个字符之间！"},
		{"val", "string", "required,myString,gte=2,lte=65535", "配置值在2到65535个字符之间！"},
	}
	config, err := getConfigModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := config.UpdateConfig()
	if !res["flag"].(bool) {
		response.RenderError(ctx, res["msg"].(string), nil)
		return
	}
	response.RenderSuccess(ctx, res["msg"].(string), nil)
}

/**
删除配置
*/
func DeleteConfig(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"id", "int64", "required,myString,gt=0", "配置ID错误！"},
	}
	config, err := getConfigModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := config.DeleteConfig()
	if !res {
		response.RenderError(ctx, "删除配置失败，请稍后再试", nil)
		return
	}
	response.RenderSuccess(ctx, "删除配置成功！", nil)
}
