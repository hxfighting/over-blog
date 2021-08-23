package base

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	stdJson "encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	"github.com/ohdata/blog/configs"
	"github.com/ohdata/blog/internal/middlewares/jwt"
	"github.com/ohdata/blog/internal/pkg/database"
	"github.com/ohdata/blog/tools"
	"github.com/ohdata/blog/tools/captcha"
	"github.com/ohdata/blog/tools/log"
)

// 5MB
const maxSize = 5 << 20

func Route(app *fiber.App, handlers ...fiber.Handler) {

	app.Get("/api/captcha", getCaptcha)
	g := app.Group("/api/admin", handlers...)
	g.Post("/upload", upload)
	g.Get("/count", count)
	g.Get("/token", refreshToken)
}

// 生成验证码
func getCaptcha(ctx *fiber.Ctx) error {
	key, captchaBase64 := captcha.GenerateCaptcha()
	data := make(map[string]string)
	data["img"] = captchaBase64
	data["key"] = key
	return tools.Success(ctx, "登录成功！", data)
}

// 上传图片
func upload(ctx *fiber.Ctx) error {
	c := configs.Config.QiNiu
	if c.AccessKey == "" || c.SecretKey == "" || c.Bucket == "" {
		return tools.Error(ctx, "请完善七牛配置参数！")
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Log.Err(err).Send()
		return tools.Error(ctx, "上传失败")
	}
	if file.Size > maxSize {
		return tools.Error(ctx, "上传失败：文件最大5MB")
	}
	f, err := file.Open()
	if err != nil {
		log.Log.Err(err).Send()
		return tools.Error(ctx, "上传失败")
	}
	defer f.Close()
	fileByte, err := ioutil.ReadAll(f)
	if err != nil {
		return tools.Error(ctx, "上传失败：文件读取失败")
	}
	contentType := http.DetectContentType(fileByte)
	if contentType != "image/jpeg" && contentType != "image/jpg" &&
		contentType != "image/png" && contentType != "image/gif" {
		return tools.Error(ctx, "上传图片允许的格式只能是:gif、png、jpg、jpeg！")
	}
	h := md5.New()
	h.Write(fileByte)
	key := hex.EncodeToString(h.Sum(nil))
	putPolicy := storage.PutPolicy{
		Scope: c.Bucket,
	}
	new_file := ioutil.NopCloser(bytes.NewReader(fileByte))
	mac := qbox.NewMac(c.AccessKey, c.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err = formUploader.Put(ctx.Context(), &ret, upToken, key, new_file, file.Size, &putExtra)
	if err != nil {
		log.Log.Err(err).Send()
		return tools.Error(ctx, "上传失败")
	}
	img := strings.Replace(ret.Key, "http://image.ohdata.top", "https://pic.ohdata.top", -1)
	return tools.Success(ctx, "上传成功！", img)
}

func count(ctx *fiber.Ctx) error {
	data := make([]map[string]interface{}, 8)
	user, link, article, comment, tag, category, contact, chat := int64(0), int64(0), int64(0),
		int64(0), int64(0), int64(0), int64(0), int64(0)
	database.DB.Table("user").Count(&user)
	database.DB.Table("link").Count(&link)
	database.DB.Table("article").Count(&article)
	database.DB.Table("article_comment").Count(&comment)
	database.DB.Table("tag").Count(&tag)
	database.DB.Table("category").Count(&category)
	database.DB.Table("contact").Count(&contact)
	database.DB.Table("chat").Count(&chat)
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
	return tools.Success(ctx, "获取统计成功", data)
}

func refreshToken(ctx *fiber.Ctx) error {
	info, err := jwt.GetUserInfo(ctx)
	if err != nil {
		return err
	}
	uidJSON := info["subject"].(stdJson.Number)
	uid, err := uidJSON.Int64()
	if err != nil {
		return tools.ErrServer
	}
	if err = jwt.BlockToken(ctx); err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	nameJSON := info["username"].(stdJson.Number)
	name := nameJSON.String()
	data, err := jwt.GenerateToken(uid, name, "admin")
	if err != nil {
		return tools.ServerErrResponse(ctx, err)
	}
	return tools.Success(ctx, "token刷新成功", data)
}
