# aliyun 阿里云统一服务
### 微服务内核访问演示
```
{
    "service": "aliyun",
    "method": "Aliyun.ProcessCommonRequest",
    "request": {
        "Method": "POST",
        "Scheme": "https",
        "Domain": "dysmsapi.aliyuncs.com",
        "Version": "2017-05-25",
        "ApiName": "SendSms",
        "QueryParams": {
        	"PhoneNumbers":"13954386521",
        	"SignName":"阿里云",
        	"TemplateCode":"SMS_135275049",
        	"TemplateParam":"{code: '562374'}"
        }
    }
}
```
- 具体参数参考阿里云开发文档
- https://api.aliyun.com/
