package backend

import (
	"blog/database"
	"blog/models"
	"blog/service"
	"errors"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

/**
获取category model
*/
func getCategory(ctx iris.Context, validates []service.BlogValidate) (models.Category, error) {
	category := models.Category{}
	requestData, err := getRequestData(ctx, validates)
	if err != nil {
		return category, err
	}
	err = mapstructure.Decode(requestData, &category)
	if err != nil {
		return category, errors.New("参数错误！")
	}
	return category, nil
}

/**
获取分类列表
*/
func GetCategoryList(ctx iris.Context) {
	categories := models.GetCategoryList()
	if len(categories) > 0 {
		response.RenderSuccess(ctx, "获取分类列表成功", categories)
	}
	response.RenderError(ctx, "暂无分类列表数据", nil)
}

/**
添加分类
*/
func AddCategory(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"pid", "required,gte=0", "分类pid错误！"},
		{"title", "required,gte=2,lte=20", "分类名称在2到20个字符之间！"},
	}
	category, err := getCategory(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	if !checkExistTitle(category, true) {
		response.RenderError(ctx, "分类名称已存在，请换一个！", nil)
		return
	}
	res := models.AddCategory(&category)
	if !res {
		response.RenderError(ctx, "分类添加失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "分类添加成功", nil)
}

/**
检查分类名称是否存在
*/
func checkExistTitle(cate models.Category, flag bool) bool {
	category := models.Category{}
	database.Db.Where("title = ?", cate.Title).First(&category)
	if category.Url != "" {
		if flag {
			return false
		} else {
			return *category.ID == *cate.ID
		}
	}
	return true
}

/**
修改分类
*/
func UpdateCategory(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"id", "required,gt=0", "分类id错误！"},
		{"title", "required,gte=2,lte=20", "分类名称在2到20个字符之间！"},
	}
	category, err := getCategory(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	if !checkExistTitle(category, false) {
		response.RenderError(ctx, "分类名称已存在，请换一个！", nil)
		return
	}
	res := models.UpdateCategory(&category)
	if !res {
		response.RenderError(ctx, "分类修改失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "分类修改成功", nil)
}

/**
删除分类
*/
func DeleteCategory(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"id", "required,gt=0", "分类id错误！"},
	}
	category, err := getCategory(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.DeleteCategory(&category)
	if !res {
		response.RenderError(ctx, "分类删除失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "分类删除成功", nil)
}
