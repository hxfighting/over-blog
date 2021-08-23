package category

import (
	"html"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/ohdata/blog/internal/models"
	"github.com/ohdata/blog/internal/pkg/event"
	"github.com/ohdata/blog/tools"
)

func Route(app *fiber.App, handlers ...fiber.Handler) {
	g := app.Group("/api/admin/category", handlers...)
	g.Get("", list)
	g.Post("", save)
	g.Put("", update)
	g.Delete("", deleteHandler)
}

func list(ctx *fiber.Ctx) error {
	data, err := models.GetCategoryList(ctx)
	if err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	return tools.Success(ctx, "获取数据成功", data)
}

type saveReq struct {
	PID   int64  `json:"pid" validate:"required,gt=0"`
	Title string `json:"title" validate:"required,lte=50"`
}

func save(ctx *fiber.Ctx) error {
	req := new(saveReq)
	if err := tools.ValidateStruct(ctx, req); err != nil {
		return tools.Error(ctx, "参数错误")
	}
	exist, err := models.CountCategoryByTitle(ctx, req.Title)
	if err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	if exist > 0 {
		return tools.Error(ctx, "分类名称已存在，请换一个！")
	}
	req.Title = html.EscapeString(req.Title)
	now := time.Now().Unix()
	c := models.Category{
		Title:     &req.Title,
		Pid:       &req.PID,
		Type:      models.CategoryArticleType.UInt(),
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err = models.CreateCategory(ctx, &c); err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	event.EventBus.Publish(event.CategoryCreateEvent.String())
	return tools.Success(ctx, "分类添加成功")
}

type updateReq struct {
	ID    int64  `json:"id" validate:"required,gt=0"`
	Title string `json:"title" validate:"required,lte=50"`
}

func update(ctx *fiber.Ctx) error {
	req := new(updateReq)
	if err := tools.ValidateStruct(ctx, req); err != nil {
		return tools.Error(ctx, "参数错误")
	}
	exist, err := models.CountCategoryByTitle(ctx, req.Title)
	if err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	if exist > 0 {
		return tools.Error(ctx, "分类名称已存在，请换一个！")
	}
	updateData := map[string]interface{}{
		"title":      req.Title,
		"updated_at": time.Now().Unix(),
	}
	if err = models.UpdateCategory(ctx, updateData); err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	event.EventBus.Publish(event.CategoryCreateEvent.String())
	return tools.Success(ctx, "分类修改成功")
}

type deleteReq struct {
	ID int64 `json:"id" validate:"required,gt=0"`
}

func deleteHandler(ctx *fiber.Ctx) error {
	req := new(deleteReq)
	if err := tools.ValidateStruct(ctx, req); err != nil {
		return tools.Error(ctx, "参数错误")
	}

	return nil
}
