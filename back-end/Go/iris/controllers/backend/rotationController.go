package backend

import (
	"blog/controllers"
	"blog/controllers/frontend"
	"blog/models"
	"github.com/kataras/iris"
)

/**
获取轮播图列表
*/
func GetRotationList(ctx iris.Context) {
	rotation := models.Rotation{}
	list := rotation.GetRotationList()
	if len(list) > 0 {
		response.RenderSuccess(ctx, "获取轮播图列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无轮播图列表数据", nil)
}

/**
添加轮播图
*/
func AddRotation(ctx iris.Context) {
	rotation := models.Rotation{}
	fields := []string{"image_url", "words"}
	validateFields := []string{"ImageUrl", "Words"}
	_, err := controllers.GetRightModel(ctx, &rotation, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := rotation.AddRotation()
	if !res {
		response.RenderError(ctx, "添加轮播图失败，请稍后再试！", nil)
		return
	}
	removeFrontendCache(frontend.ROTATION_KEY)
	response.RenderSuccess(ctx, "添加轮播图成功！", nil)
}

/**
修改轮播图信息
*/
func UpdateRotation(ctx iris.Context) {
	rotation := models.Rotation{}
	fields := []string{"id", "image_url", "words"}
	validateFields := []string{"ID", "ImageUrl", "Words"}
	_, err := controllers.GetRightModel(ctx, &rotation, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := rotation.UpdateRotation()
	if !res {
		response.RenderError(ctx, "修改轮播图失败，请稍后再试！", nil)
		return
	}
	removeFrontendCache(frontend.ROTATION_KEY)
	response.RenderSuccess(ctx, "修改轮播图成功！", nil)
}

/**
删除轮播图
*/
func DeleteRotation(ctx iris.Context) {
	rotation := models.Rotation{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	_, err := controllers.GetRightModel(ctx, &rotation, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := rotation.DeleteRotation()
	if !res {
		response.RenderError(ctx, "删除轮播图失败，请稍后再试！", nil)
		return
	}
	removeFrontendCache(frontend.ROTATION_KEY)
	response.RenderSuccess(ctx, "删除轮播图成功！", nil)
}
