package service

import "github.com/kataras/iris/v12"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (this *Response) RenderSuccess(ctx iris.Context, msg string, data interface{}) {
	this.Code = iris.StatusOK
	this.Message = msg
	this.Data = data
	ctx.JSON(this)
}

func (this *Response) RenderError(ctx iris.Context, msg string, data interface{}) {
	if this.Code == 0 {
		this.Code = -1
	}
	this.Message = msg
	this.Data = data
	ctx.JSON(this)
}
