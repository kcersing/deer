namespace go base

    /**
     * @description 短信服务相关信息
     */
    struct Sms{
        /**通知短信数量*/
        1: optional i64 noticeCount=0  (api.raw = "noticeCount")
        /**已用通知*/
        2: optional i64 usedNotice=0  (api.raw = "usedNotice")
        /**创建时间*/
        16: optional string createdAt = ""  (api.raw = "createdAt")
        /**更新时间*/
        17: optional string updatedAt = "" (api.raw = "updatedAt")
   }

    /**
     * @description 短信发送记录
     */
    struct SmsSend{
        /**创建时间*/
         1: optional string createdAt = ""  (api.raw = "createdAt")
         /**状态*/
         2: optional i64 status=0  (api.raw = "status")
         /**手机号*/
         3: optional string mobile="" (api.raw = "mobile")
         /**验证码*/
         5: optional string code="" (api.raw = "code")
         /**业务ID*/
         7: optional string bizId = ""  (api.raw = "bizId")
         /**通知类型[1会员;2员工]*/
         8: optional i64 userType = 1  (api.raw = "userType")
         /**内容*/
         9: optional string content = ""  (api.raw = "content")
         /**模板*/
         10: optional string templates = ""  (api.raw = "templates")
         11: optional i64 noticeCount=0  (api.raw = "noticeCount")
     }

     /**
      * @description 消息发送记录
      */
     struct MessagesSend{
        /**创建时间*/
        1: optional string createdAt = ""  (api.raw = "createdAt")
        /**状态*/
        2: optional i64 status=0  (api.raw = "status")
        /**接收时间*/
        4: optional string receivedAt = ""  (api.raw = "receivedAt")
        /**阅读时间*/
        5: optional string readAt = ""  (api.raw = "readAt")
        /**id*/
        6: optional i64 id=0  (api.raw = "id")
        /**类型*/
        8: optional string type = ""  (api.raw = "type")
        /**内容*/
        9: optional string content = ""  (api.raw = "content")
        /**消息ID*/
        10: optional i64 messagesId = 0  (api.raw = "messagesId")
        /**发送用户ID*/
        11: optional i64 fromUserId=0  (api.raw = "fromUserId")
        /**发送用户名*/
        12: optional string fromUserName = ""  (api.raw = "fromUserName")
     }

     /**
      * @description 消息主体
      */
     struct Messages{
         /**创建时间*/
         1: optional string createdAt = ""  (api.raw = "createdAt")
         /**状态*/
         2: optional i64 status=0  (api.raw = "status")
         /**发送用户ID*/
         3: optional i64 createdId=0  (api.raw = "createdId")
         5: optional string createdName=""  (api.raw = "createdName")
         /**id*/
         4: optional i64 id=0  (api.raw = "id")
         /**类型*/
         8: optional string type="" (api.raw = "type")
         /**内容*/
         9: optional string content = ""  (api.raw = "content")
         10: optional string title="" (api.raw = "title")
     }
