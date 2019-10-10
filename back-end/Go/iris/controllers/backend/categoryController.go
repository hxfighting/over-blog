package backend

import (
	"blog/database"
	"blog/models"
	"github.com/kataras/iris"
)

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
	category := models.Category{}
	fields := []string{"pid", "title"}
	validateFields := []string{"Pid", "Title"}
	err := getRightModel(ctx, &category, fields, validateFields)
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
	category := models.Category{}
	fields := []string{"id", "title"}
	validateFields := []string{"ID", "Title"}
	err := getRightModel(ctx, &category, fields, validateFields)
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
	category := models.Category{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	err := getRightModel(ctx, &category, fields, validateFields)
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
