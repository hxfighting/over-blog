package frontend

import (
	"blog/service"
	template "blog/views"
	"bytes"
	"github.com/kataras/iris"
)

/**
联系我页面
 */
func GetContactPage(ctx iris.Context) {
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	buffer := new(bytes.Buffer)
	template.ContactPage(buffer)
	_, err := ctx.Write(buffer.Bytes())
	if err != nil {
		service.Log.Error(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
	}
}
