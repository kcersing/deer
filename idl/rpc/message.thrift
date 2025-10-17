namespace go message
include "../base/base.thrift"

    struct SendSmsListReq {
         1: optional i64 page=1 (api.raw = "page")
         2: optional i64 pageSize=100 (api.raw = "pageSize")
         3: optional string mobile="" (api.raw = "mobile")
     }

     struct SmsSend{
         1: optional string createdAt = ""  (api.raw = "createdAt")
         2: optional i64 status=0  (api.raw = "status")
         3: optional string mobile="" (api.raw = "mobile")
         5: optional string code="" (api.raw = "code")
         7: optional string bizId = ""  (api.raw = "bizId")
         /**通知类型[1会员;2员工]*/
         8: optional i64 notifyType = 1  (api.raw = "notifyType")
         9: optional string content = ""  (api.raw = "content")
         10: optional string templates = ""  (api.raw = "templates")
     }

    struct Sms{
        /**通知短信数量*/
        1: optional i64 noticeCount=0  (api.raw = "noticeCount")
        /**已用通知*/
        2: optional i64 usedNotice=0  (api.raw = "usedNotice")

        16: optional string createdAt = ""  (api.raw = "createdAt")
        17: optional string updatedAt = "" (api.raw = "updatedAt")
        }
    struct SmsResp{
        1:optional Sms data={}
        255:optional base.BaseResp baseResp={}
    }
    struct SendSmsListResp{
        1:optional list<SmsSend> data={}
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