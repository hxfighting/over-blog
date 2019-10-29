package sms

import (
	"blog/service"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type Ali struct {
	AppID        string
	Secret       string
	SignName     string
	TemplateCode string
	Data         map[string]interface{}
	PhoneNumber  string
}

/**
发送短信
*/
func (this Ali) Send() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", this.AppID, this.Secret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = this.PhoneNumber
	request.SignName = this.SignName
	request.TemplateCode = this.TemplateCode
	bytes, _ := service.FastJson.Marshal(this.Data)
	request.TemplateParam = string(bytes)
	response, err := client.SendSms(request)
	if err != nil {
		service.Log.Error(err.Error())
		return
	}
	if response.Message != "OK" && response.Code != "OK" {
		service.Log.Error(response.Message)
		return
	}
	service.Log.Info("短信发送成功")
}
