package util

import(
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"encoding/json"
)

//ak : LTAIB0aokppPwwxW
//sk : ZKZCmKYXAKOxvM6xQbsKx2Ft1iEqT6
func AliSendSMS(phone string,captcha string)error{
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAIB0aokppPwwxW", "ZKZCmKYXAKOxvM6xQbsKx2Ft1iEqT6")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = "蜜传消息"
	request.TemplateCode = "SMS_149417711"
	param := map[string]string{
		"code" : captcha,
	}
	tp,_ := json.Marshal(param)
	log.Println(string(tp))
	request.TemplateParam = string(tp)

	response, err := client.SendSms(request)
	log.Println(response)
	return err
}