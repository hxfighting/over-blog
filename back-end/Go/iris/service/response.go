package service

import "github.com/kataras/iris"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (this *Response) RenderSuccess(ctx iris.Context, msg string, data interface{}) {
	this.Code = iris.StatusOK
	this.Message = msg
	this.Data = data
	ctx.JSON(this)
}

func (this *Response) RenderError(ctx iris.Context, msg string, data interface{}) {
	this.Code = -1
	this.Message = msg
	this.Data = data
	ctx.JSON(this)
}
