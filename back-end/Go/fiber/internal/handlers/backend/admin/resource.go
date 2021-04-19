package admin

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/ohdata/blog/internal/middlewares/jwt"
	"github.com/ohdata/blog/internal/models"
	"github.com/ohdata/blog/tools"
	"github.com/ohdata/blog/tools/util"
)

func Route(app *fiber.App, handlers ...fiber.Handler) {

	g := app.Group("/api/admin", handlers...)
	// 获取个人信息
	g.Get("", getUserInfo)
	// 修改个人信息
	g.Put("", update)
	// 退出登录
	g.Post("/logout", logout)
	// 修改密码
	g.Put("/password", resetPassword)
}

func getUserInfo(ctx *fiber.Ctx) error {
	uid, err := jwt.GetUID(ctx)
	if err != nil {
		return tools.Error(ctx, err.Error())
	}
	admin, err := models.GetUserInfo(ctx, uid)
	if err != nil {
		return tools.SingleRecordResponse(err, "用户不存在")
	}
	avatarName := admin.Avatar[strings.LastIndex(admin.Avatar, "/")+1:]
	user := fiber.Map{
		"avatarName": avatarName,
		"user_id":    admin.ID,
		"name":       admin.Name,
		"access":     []string{"super_admin"},
		"avatar":     admin.Avatar,
		"email":      admin.Email,
		"mobile":     admin.Mobile,
	}
	return tools.Success(ctx, "获取数据成功", user)
}

type updateReq struct {
	Avatar string `json:"avatar" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	Phone  string `json:"phone" validate:"required"`
}

func update(ctx *fiber.Ctx) error {
	uid, err := jwt.GetUID(ctx)
	if err != nil {
		return tools.Error(ctx, err.Error())
	}
	req := new(updateReq)
	if err := tools.ValidateStruct(ctx, req); err != nil {
		return tools.Error(ctx, "参数错误")
	}
	if !util.ValidatePhone(req.Phone) {
		return tools.Error(ctx, "请输入正确的电话")
	}
	if err = models.UpdateInfo(ctx, uid, req.Name, req.Avatar, req.Email, req.Phone); err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	return tools.Success(ctx, "修改个人信息成功！")
}

func logout(ctx *fiber.Ctx) error {
	if err := jwt.BlockToken(ctx); err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	return tools.Success(ctx, "退出登录成功！")
}

type resetPassReq struct {
	Password string `json:"password" validate:"required"`
}

func resetPassword(ctx *fiber.Ctx) error {
	uid, err := jwt.GetUID(ctx)
	if err != nil {
		return tools.Error(ctx, err.Error())
	}
	req := new(resetPassReq)
	if err := tools.ValidateStruct(ctx, req); err != nil {
		return tools.Error(ctx, "参数错误")
	}
	pass, err := util.PasswordHash(req.Password)
	if err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	if err = models.ResetPassword(ctx, uid, pass); err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	return tools.Success(ctx, "修改密码成功！")
}
