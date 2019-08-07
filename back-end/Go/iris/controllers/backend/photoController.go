package backend

import (
	"blog/helper"
	"blog/models"
	"blog/service"
	"github.com/kataras/iris"
)

/**
获取照片列表数据
*/
func GetPhotoList(ctx iris.Context) {
	photo := models.Photo{}
	list := photo.GetPhotoList()
	if len(list) > 0 {
		response.RenderSuccess(ctx, "获取照片列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无照片列表数据", nil)
}

/**
添加照片
*/
func AddPhoto(ctx iris.Context) {
	photo := models.Photo{}
	data, e := helper.GetRequestData(ctx)
	if e != nil {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	if _, ok := data["image_url"]; !ok {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	validates := []service.BlogValidate{
		{data["image_url"], "required,url", "照片连接错误！"},
	}
	err := service.ValidateField(validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := photo.AddPhoto(data["image_url"].(string))
	if !res {
		response.RenderError(ctx, "添加照片失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "添加照片成功！", nil)
}

/**
修改照片
*/
func UpdatePhoto(ctx iris.Context) {
	image := models.Image{}
	err := ctx.ReadJSON(&image)
	if err != nil {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	validates := []service.BlogValidate{
		{image.ID, "required,gt=0", "照片ID错误！"},
		{image.Image_url, "required,url", "照片连接错误！"},
	}
	err = service.ValidateField(validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	photo := models.Photo{}
	res := photo.UpdatePhoto(image)
	if !res {
		response.RenderError(ctx, "修改照片失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "修改照片成功！", nil)
}

/**
删除照片
*/
func DeletePhoto(ctx iris.Context) {
	photo := models.Photo{}
	data, e := helper.GetRequestData(ctx)
	if e != nil {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	if _, ok := data["id"]; !ok {
		response.RenderError(ctx, "参数错误！", nil)
		return
	}
	validates := []service.BlogValidate{
		{data["id"], "required,gt=0", "照片ID错误！"},
	}
	err := service.ValidateField(validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := photo.DeletePhoto(int64(data["id"].(float64)))
	if !res {
		response.RenderError(ctx, "删除照片失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除照片成功！", nil)
}
