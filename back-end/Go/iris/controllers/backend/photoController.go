package backend

import (
	"blog/controllers"
	"blog/controllers/frontend"
	"blog/models"
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
	fields := []string{"image_url"}
	validateFields := []string{"ImageUrl"}
	_, err := controllers.GetRightModel(ctx, &photo, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := photo.AddPhoto()
	if !res {
		response.RenderError(ctx, "添加照片失败，请稍后再试！", nil)
		return
	}
	removeFrontendCache(frontend.PHOTO_KEY)
	response.RenderSuccess(ctx, "添加照片成功！", nil)
}

/**
修改照片
*/
func UpdatePhoto(ctx iris.Context) {
	photo := models.Photo{}
	fields := []string{"id", "image_url"}
	validateFields := []string{"ID", "ImageUrl"}
	_, err := controllers.GetRightModel(ctx, &photo, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := photo.UpdatePhoto()
	if !res {
		response.RenderError(ctx, "修改照片失败，请稍后再试！", nil)
		return
	}
	removeFrontendCache(frontend.PHOTO_KEY)
	response.RenderSuccess(ctx, "修改照片成功！", nil)
}

/**
删除照片
*/
func DeletePhoto(ctx iris.Context) {
	photo := models.Photo{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	_, err := controllers.GetRightModel(ctx, &photo, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := photo.DeletePhoto()
	if !res {
		response.RenderError(ctx, "删除照片失败，请稍后再试！", nil)
		return
	}
	removeFrontendCache(frontend.PHOTO_KEY)
	response.RenderSuccess(ctx, "删除照片成功！", nil)
}
