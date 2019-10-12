package backend

import (
	"blog/config"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/kataras/iris"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io/ioutil"
	"net/http"
)

//5MB
const maxSize = 5 << 20

/**
上传图片
*/
func Upload(ctx iris.Context) {
	accessKey := config.GetConfig("qiniu.accessKey").(string)
	secretKey := config.GetConfig("qiniu.secretKey").(string)
	bucket := config.GetConfig("qiniu.bucket").(string)
	url := config.GetConfig("qiniu.url").(string)
	if accessKey == "" || secretKey == "" || bucket == "" || url == "" {
		response.RenderError(ctx, "请完善七牛配置参数！", nil)
		return
	}
	file, info, e := ctx.FormFile("file")
	if e != nil {
		response.RenderError(ctx, "上传失败："+e.Error(), nil)
		return
	}
	defer file.Close()
	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		response.RenderError(ctx, "上传失败：文件读取失败", nil)
		return
	}
	contentType := http.DetectContentType(bytes)
	if info.Size > maxSize {
		response.RenderError(ctx, "上传失败：文件最大5MB", nil)
		return
	}
	if contentType != "image/jpeg" && contentType != "image/jpg" &&
		contentType != "image/png" && contentType != "image/gif" {
		response.RenderError(ctx, "上传图片允许的格式只能是:gif、png、jpg、jpeg！", nil)
		return
	}
	h := md5.New()
	h.Write([]byte(info.Filename))
	key := hex.EncodeToString(h.Sum(nil))
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
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
	e = formUploader.Put(context.Background(), &ret, upToken, key, file, info.Size, &putExtra)
	if e != nil {
		response.RenderError(ctx, "上传失败："+e.Error(), nil)
		return
	}
	response.RenderSuccess(ctx, "上传成功！", url+ret.Key)
}
