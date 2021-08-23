package admin

import (
	"html"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/ohdata/blog/internal/middlewares/jwt"
	"github.com/ohdata/blog/internal/models"
	"github.com/ohdata/blog/tools"
	"github.com/ohdata/blog/tools/captcha"
	"github.com/ohdata/blog/tools/util"
)

func Route(app *fiber.App, handlers ...fiber.Handler) {
	prefix := "/api/admin"
	auth := app.Group(prefix, handlers...)
	// 获取个人信息
	auth.Get("", getUserInfo)
	// 修改个人信息
	auth.Put("", update)
	// 退出登录
	auth.Post("/logout", logout)
	// 修改密码
	auth.Put("/password", resetPassword)

	noAuth := app.Group(prefix)
	// 登录
	noAuth.Post("/login", login)
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

type loginReq struct {
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Captcha  string `json:"captcha" validate:"required"`
	Key      string `json:"key" validate:"required"`
}

func login(ctx *fiber.Ctx) error {
	req := new(loginReq)
	if err := tools.ValidateStruct(ctx, req); err != nil {
		return tools.Error(ctx, "参数错误")
	}
	req.Password = html.EscapeString(req.Password)
	req.Name = html.EscapeString(req.Name)
	if captcha.ValidateCaptcha(req.Key, req.Captcha) {
		user, err := models.GetUserByName(ctx, req.Name)
		if err != nil {
			return tools.SingleRecordResponse(err, "参数错误")
		}
		if user.ID <= 0 {
			return tools.Error(ctx, "参数错误")
		}
		data, err := jwt.GenerateToken(user.ID, user.Name, "admin")
		if err != nil {
			return tools.ServerErrResponse(ctx, err)
		}
		return tools.Success(ctx, "登录成功！", data)
	}
	return tools.Error(ctx, "验证码错误")
}
