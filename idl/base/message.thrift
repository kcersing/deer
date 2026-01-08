namespace go base

    struct Sms{
        /**通知短信数量*/
        1: optional i64 noticeCount=0  (api.raw = "noticeCount")
        /**已用通知*/
        2: optional i64 usedNotice=0  (api.raw = "usedNotice")
        16: optional string createdAt = ""  (api.raw = "createdAt")
        17: optional string updatedAt = "" (api.raw = "updatedAt")
   }


    struct SmsSend{
         1: optional string createdAt = ""  (api.raw = "createdAt")
         2: optional i64 status=0  (api.raw = "status")
         3: optional string mobile="" (api.raw = "mobile")
         5: optional string code="" (api.raw = "code")
         7: optional string bizId = ""  (api.raw = "bizId")
         /**通知类型[1会员;2员工]*/
         8: optional i64 userType = 1  (api.raw = "userType")
         9: optional string content = ""  (api.raw = "content")
         10: optional string templates = ""  (api.raw = "templates")
     }
     struct MessagesSend{
        1: optional string createdAt = ""  (api.raw = "createdAt")
        2: optional i64 status=0  (api.raw = "status")
        4: optional string receivedAt = ""  (api.raw = "receivedAt")
        5: optional string readAt = ""  (api.raw = "readAt")
        6: optional i64 id=0  (api.raw = "id")
        8: optional i64 type = 1  (api.raw = "type")
        9: optional string content = ""  (api.raw = "content")
        10: optional i64 messagesId = 0  (api.raw = "messagesId")
        11: optional i64 fromUserId=0  (api.raw = "fromUserId")
        12: optional string fromUserName = ""  (api.raw = "fromUserName")
     }
     struct Messages{
         1: optional string createdAt = ""  (api.raw = "createdAt")
         2: optional i64 status=0  (api.raw = "status")
         3: optional i64 fromUserId=0  (api.raw = "fromUserId")
         4: optional i64 id=0  (api.raw = "id")
         8: optional i64 type = 1  (api.raw = "type")
         9: optional string content = ""  (api.raw = "content")
     }