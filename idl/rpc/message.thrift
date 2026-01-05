namespace go message
include "../base/base.thrift"
include "../base/message.thrift"
   struct SmsListReq {
        1: optional i64 page=1
        2: optional i64 pageSize=100
        3: optional string mobile=""
   }
    struct SmsResp{
        1:optional message.Sms data={}
        255:optional base.BaseResp baseResp={}
    }
    struct SendSmsReq{
        1: optional string mobile=""
    }
    struct SmsSendListResp{
        1:optional list<message.SmsSend> data={}
        255:optional base.BaseResp baseResp={}
    }
    struct SmsSendListReq {
        1: optional string mobile=""
        2: optional string msg=""
    }
    struct MessagesListReq{
        1:optional list<message.SmsSend> data={}

        255:optional base.BaseResp baseResp={}
    }
    struct SendMessagesReq {
        1: optional string mobile=""
        2: optional string msg=""
    }


    struct MessagesListResp{
        1:optional list<message.Messages> data={}
        255:optional base.BaseResp baseResp={}
    }
    struct MessagesSendListResp{
        1:optional list<message.MessagesSend> data={}
        255:optional base.BaseResp baseResp={}
    }

    service MessageService  {
        //获取验证码
        base.NilResponse ImgCaptcha()
        /**短信信息*/
        SmsResp Sms(1: base.IdReq req)
        /**发送记录*/
        SmsSendListResp SmsSendList(1: SmsSendListReq req)
        //发送信息
        base.NilResponse SendSms(1: SendSmsReq req)

        /**发送记录*/
        MessagesListResp MessagesList(1: MessagesListReq req)
        //发送信息
        base.NilResponse SendMessages(1: SendMessagesReq req)

        MessagesSendListResp MessagesSendList(1: MessagesListReq req)
    }