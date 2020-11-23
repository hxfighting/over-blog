package backend

import (
	"blog/config"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/kataras/iris/v12"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io/ioutil"
	"net/http"
	"strings"
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
	if accessKey == "" || secretKey == "" || bucket == "" {
		response.RenderError(ctx, "请完善七牛配置参数！", nil)
		return
	}
	file, info, e := ctx.FormFile("file")
	if e != nil {
		response.RenderError(ctx, "上传失败："+e.Error(), nil)
		return
	}
	defer file.Close()
	file_byte, e := ioutil.ReadAll(file)
	if e != nil {
		response.RenderError(ctx, "上传失败：文件读取失败", nil)
		return
	}
	contentType := http.DetectContentType(file_byte)
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
	h.Write(file_byte)
	key := hex.EncodeToString(h.Sum(nil))
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	new_file := ioutil.NopCloser(bytes.NewReader(file_byte))
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
	e = formUploader.Put(context.Background(), &ret, upToken, key, new_file, info.Size, &putExtra)
	if e != nil {
		response.RenderError(ctx, "上传失败："+e.Error(), nil)
		return
	}
	img := strings.Replace(ret.Key, "http://image.ohdata.top", "https://pic.ohdata.top", -1)
	response.RenderSuccess(ctx, "上传成功！", img)
}
