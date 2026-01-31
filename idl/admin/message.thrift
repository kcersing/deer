namespace go message
include "../base/base.thrift"
include "../base/message.thrift"

    struct SendSmsReq{
        1: optional string mobile=""
        2: optional string type=""
        3: optional string content = ""
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
    struct SendMemberMessagesReq {
        1: optional i64 memberId=0
        2: optional string type=""
        3: optional string content = ""
        4: optional string title = ""
        5: optional i64 createdId=0
        6:optional list<i64> tagId=[]
        7:optional i64 status=0
    }
    struct SendUserMessagesReq {
        1: optional i64 userId=0
        2: optional string type=""
        3: optional string content = ""
        4: optional string title = ""
        5: optional i64 createdId=0
        6:optional list<i64> tagId=[]
        7:optional i64 status=0

    }

    service MessageService  {
        //短信条数信息
        base.NilResponse Sms(1: base.IdReq req)(api.post = "/service/message/sms")
        //短信发送记录
        base.NilResponse SmsSendList(1: SmsSendListReq req)(api.post = "/service/message/sms-send-list")
        //发送记录
        base.NilResponse MessagesList(1: MessagesListReq req)(api.post = "/service/message/messages-list")
        //发送信息
        base.NilResponse SendMemberMessages(1: SendMemberMessagesReq req)(api.post = "/service/message/send-member-messages")

        base.NilResponse SendUserMessages(1: SendUserMessagesReq req)(api.post = "/service/message/send-user-messages")
        //发送记录
        base.NilResponse MessagesSendList(1: MessagesListReq req)(api.post = "/service/message/messages-send-list")

        base.NilResponse MessagesTypes()(api.post = "/service/message/types")

    }