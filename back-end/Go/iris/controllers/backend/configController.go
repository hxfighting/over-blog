package backend

import (
	"blog/controllers"
	"blog/controllers/frontend"
	"blog/models"
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
	fields := []string{"type", "name", "title", "val"}
	validateFields := []string{"Type", "Name", "Title", "Val"}
	_, err := controllers.GetRightModel(ctx, &config, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.AddConfig(&config)
	if !res["flag"].(bool) {
		response.RenderError(ctx, res["msg"].(string), nil)
		return
	}
	removeFrontendCache(frontend.CONFIG_KEY, frontend.FOOTER_KEY, frontend.SOCIAL_KEY)
	response.RenderSuccess(ctx, res["msg"].(string), nil)
}

/**
修改配置
*/
func UpdateConfig(ctx iris.Context) {
	config := models.Config{}
	fields := []string{"id", "type", "name", "title", "val"}
	validateFields := []string{"ID", "Type", "Name", "Title", "Val"}
	_, err := controllers.GetRightModel(ctx, &config, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := config.UpdateConfig()
	if !res["flag"].(bool) {
		response.RenderError(ctx, res["msg"].(string), nil)
		return
	}
	removeFrontendCache(frontend.CONFIG_KEY, frontend.FOOTER_KEY, frontend.SOCIAL_KEY)
	response.RenderSuccess(ctx, res["msg"].(string), nil)
}

/**
删除配置
*/
func DeleteConfig(ctx iris.Context) {
	config := models.Config{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	_, err := controllers.GetRightModel(ctx, &config, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := config.DeleteConfig()
	if !res {
		response.RenderError(ctx, "删除配置失败，请稍后再试", nil)
		return
	}
	removeFrontendCache(frontend.CONFIG_KEY, frontend.FOOTER_KEY, frontend.SOCIAL_KEY)
	response.RenderSuccess(ctx, "删除配置成功！", nil)
}
