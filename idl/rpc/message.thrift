namespace go message
include "../base/base.thrift"
include "../base/message.thrift"
   struct SendSmsListReq {
        1: optional i64 page=1 (api.raw = "page")
        2: optional i64 pageSize=100 (api.raw = "pageSize")
        3: optional string mobile="" (api.raw = "mobile")
   }
    struct SmsResp{
        1:optional message.Sms data={}
        255:optional base.BaseResp baseResp={}
    }
    struct SendSmsListResp{
        1:optional list<message.SmsSend> data={}
        255:optional base.BaseResp baseResp={}
    }
    struct SendSmsReq {
        1: optional string mobile="" (api.raw = "mobile")
        2: optional string msg="" (api.raw = "msg")
    }

    service MessageService  {
        //获取验证码
        base.NilResponse ImgCaptcha() (api.post = "/service/captcha/img")
        /**短信信息*/
        SmsResp Sms(1: base.IdReq req) (api.post = "/service/sms")
        /**发送记录*/
        SendSmsListResp SmsList(1: SendSmsListReq req) (api.post = "/service/sms/send-list")
        //发送信息
        base.NilResponse SendSms(1: SendSmsReq req) (api.post = "/service/sms/send")
    }