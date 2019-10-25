package frontend

import (
	"blog/controllers"
	"blog/models"
	"blog/service"
	"github.com/kataras/iris"
)

/**
添加评论
*/
func AddComment(ctx iris.Context) {
	comment := models.Comment{}
	fields := []string{"email", "content", "user_id", "reply_id", "article_id", "pid"}
	validateFields := []string{"Email", "Content", "UserID", "ReplyID", "ArticleID", "Pid"}
	_, err := controllers.GetRightModel(ctx, &comment, fields, validateFields)
	if err != nil {
		Response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := comment.AddComment()
	if !res {
		Response.RenderError(ctx, "申请友联失败,请稍后再试!", nil)
		return
	}
	cacheRes := service.Redis.Del(COMMENT_KEY)
	if cacheRes.Err() == nil {
		InitData()
	}
	Response.RenderSuccess(ctx, "申请友联成功！", nil)
}
