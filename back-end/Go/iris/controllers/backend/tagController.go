package backend

import (
	"blog/models"
	"github.com/kataras/iris"
)

/**
获取标签列表数据
*/
func GetTagList(ctx iris.Context) {
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	tag := models.Tag{}
	list := tag.GetTagList(pageNum, pageSize)
	if list["total"].(int64) > 0 {
		response.RenderSuccess(ctx, "获取标签列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无标签列表数据", nil)
}

/**
删除标签
*/
func DeleteTag(ctx iris.Context) {
	tag := models.Tag{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	err := getRightModel(ctx, &tag, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := tag.DeleteTag()
	if !res {
		response.RenderError(ctx, "删除标签失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除标签成功！", nil)
}

/**
添加标签
*/
func AddTag(ctx iris.Context) {
	tag := models.Tag{}
	fields := []string{"name"}
	validateFields := []string{"Name"}
	err := getRightModel(ctx, &tag, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := tag.AddTag()
	if !res {
		response.RenderError(ctx, "添加标签失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "添加标签成功！", nil)
}

/**
修改标签
*/
func UpdateTag(ctx iris.Context) {
	tag := models.Tag{}
	fields := []string{"id", "name"}
	validateFields := []string{"ID", "Name"}
	err := getRightModel(ctx, &tag, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := tag.UpdateTag()
	if !res {
		response.RenderError(ctx, "修改标签失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "修改标签成功！", nil)
}
