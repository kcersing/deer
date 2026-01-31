namespace go message
include "../base/base.thrift"
include "../base/message.thrift"

    struct SmsResp{
        1:optional message.Sms data={}
        255:optional base.BaseResp baseResp={}
    }
    struct SendSmsReq{
        1: optional string mobile=""
        2: optional string type=""
        3: optional string content = ""
    }
    struct SmsSendListResp{
        1:optional list<message.SmsSend> data={}
        255:optional base.BaseResp baseResp={}
    }
    struct SmsSendListReq {
        1: optional i64 page=1
        2: optional i64 pageSize=100
        3: optional string mobile=""
    }
    struct MessagesListReq{
       1: optional i64 page=1
       2: optional i64 pageSize=100

    }
//    struct SendMemberMessagesReq {
//        1: optional i64 memberId=0
//        2: optional string type=""
//        3: optional string content = ""
//        4: optional string title = ""
//        5: optional i64 createdId=0
//        6:optional list<i64> tagId=[]
//        7:optional i64 status=0
//    }
    struct SendUserMessagesReq {
        1: optional i64 userId=0
        2: optional string type=""
        3: optional string content = ""
        4: optional string title = ""
        5: optional i64 createdId=0
        6:optional list<i64> tagId=[]
        7:optional i64 status=0
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
        //短信条数信息
        SmsResp Sms(1: base.IdReq req)
        //发送记录
        SmsSendListResp SmsSendList(1: SmsSendListReq req)
        //发送信息
        base.NilResponse SendSms(1: SendSmsReq req)
        //发送记录
        MessagesListResp MessagesList(1: MessagesListReq req)
        //发送信息
        base.NilResponse SendMemberMessages(1: SendMemberMessagesReq req)
        base.NilResponse SendUserMessages(1: SendUserMessagesReq req)
        //发送记录
        MessagesSendListResp MessagesSendList(1: MessagesListReq req)

    }