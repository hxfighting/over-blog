package backend

import (
	"blog/models"
	"blog/service"
	"errors"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

/**
获取photo model
*/
func getPhotoModel(ctx iris.Context, validates []service.BlogValidate) (models.Photo, error) {
	photo := models.Photo{}
	requestData, err := getRequestData(ctx, validates)
	if err != nil {
		return photo, err
	}
	err = mapstructure.Decode(requestData, &photo)
	if err != nil {
		return photo, errors.New("参数错误！")
	}
	return photo, nil
}

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
	validates := []service.BlogValidate{
		{"image_url", "required,url", "照片连接错误！"},
	}
	photo, err := getPhotoModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := photo.AddPhoto()
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
	validates := []service.BlogValidate{
		{"id", "required,gt=0", "照片ID错误！"},
		{"image_url", "required,url", "照片连接错误！"},
	}
	photo, err := getPhotoModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	image.ID = photo.ID
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
	validates := []service.BlogValidate{
		{"id", "required,gt=0", "照片ID错误！"},
	}
	photo, err := getPhotoModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := photo.DeletePhoto()
	if !res {
		response.RenderError(ctx, "删除照片失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除照片成功！", nil)
}
