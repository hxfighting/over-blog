package backend

import (
	"blog/database"
	"github.com/kataras/iris"
)

/**
后台首页统计
*/
func GetTotalCount(ctx iris.Context) {
	data := make([]map[string]interface{}, 8)
	user, link, article, comment, tag, category, contact, chat := 0, 0, 0, 0, 0, 0, 0, 0
	database.Db.Table("user").Count(&user)
	database.Db.Table("link").Count(&link)
	database.Db.Table("article").Count(&article)
	database.Db.Table("article_comment").Count(&comment)
	database.Db.Table("tag").Count(&tag)
	database.Db.Table("category").Count(&category)
	database.Db.Table("contact").Count(&contact)
	database.Db.Table("chat").Count(&chat)
	data[0] = map[string]interface{}{
		"title": "用户统计",
		"icon":  "md-people",
		"count": user,
		"color": "#2d8cf0",
	}
	data[1] = map[string]interface{}{
		"title": "友联统计",
		"icon":  "ios-link",
		"count": link,
		"color": "#19be6b",
	}
	data[2] = map[string]interface{}{
		"title": "文章统计",
		"icon":  "ios-book",
		"count": article,
		"color": "#ff9900",
	}
	data[3] = map[string]interface{}{
		"title": "评论统计",
		"icon":  "ios-chatboxes",
		"count": comment,
		"color": "#ed3f14",
	}
	data[4] = map[string]interface{}{
		"title": "标签统计",
		"icon":  "md-pricetags",
		"count": tag,
		"color": "#E46CBB",
	}
	data[5] = map[string]interface{}{
		"title": "分类统计",
		"icon":  "md-list",
		"count": category,
		"color": "#9A66E4",
	}
	data[6] = map[string]interface{}{
		"title": "留言统计",
		"icon":  "md-mail",
		"count": contact,
		"color": "#FF99CC",
	}
	data[7] = map[string]interface{}{
		"title": "说说统计",
		"icon":  "ios-chatbubbles",
		"count": chat,
		"color": "#FFFF00",
	}
	response.RenderSuccess(ctx, "获取统计成功", data)
}
