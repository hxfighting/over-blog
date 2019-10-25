package backend

import (
	"blog/controllers"
	"blog/models"
	"github.com/kataras/iris"
)

/**
获取会员列表
*/
func GetUserList(ctx iris.Context) {
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	search := ctx.URLParamTrim("search")
	user := models.User{}
	list := user.GetUserList(pageNum, pageSize, search)
	if list["total"].(int64) > 0 {
		response.RenderSuccess(ctx, "获取用户列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无用户列表数据", nil)
}

/**
删除会员
*/
func DeleteUser(ctx iris.Context) {
	user := models.User{}
	fields := []string{"id"}
	validateFields := []string{"ID"}
	_, err := controllers.GetRightModel(ctx, &user, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := user.DeleteUser()
	if !res {
		response.RenderError(ctx, "删除会员失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除会员成功！", nil)
}

/**
修改会员信息
*/
func UpdateUser(ctx iris.Context) {
	user := models.User{}
	fields := []string{"id", "is_admin"}
	validateFields := []string{"ID", "IsAdmin"}
	_, err := controllers.GetRightModel(ctx, &user, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := user.UpdateUser()
	if !res {
		response.RenderError(ctx, "修改会员信息失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "修改会员信息成功！", nil)
}
