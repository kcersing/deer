package mw

import (
	"context"
	system "gen/kitex_gen/system"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func LogMw() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		c.Next(ctx)
		var log system.CreateLogReq

		log.Type = "Interface"
		log.Method = string(c.Request.Method())
		log.Api = string(c.Request.Path())
		log.UserAgent = string(c.Request.Header.UserAgent())
		log.Ip = c.ClientIP()

		reqBodyStr := string(c.Request.Body())
		if len(reqBodyStr) > 200 {
			reqBodyStr = reqBodyStr[:200]
		}
		log.ReqContent = reqBodyStr

		respBodyStr := string(c.Request.Body())
		if len(respBodyStr) > 200 {
			respBodyStr = respBodyStr[:200]
		}

		if c.Response.Header.StatusCode() == 200 {
			log.Success = true
		}

		costTime := time.Since(start).Milliseconds()
		log.Time = int64(int32(costTime))
		//var username = "Anonymous"

		//userIn, exist := c.Get("user_id")
		//if !(exist || userIn == nil) {
		//	userId := toInt(userIn)
		//	userInfo, _ := service.NewUser(ctx, c).Info(userId)
		//	if userInfo != nil {
		//		username = userInfo.Name
		//	}
		//	logs.Operatorsr = username
		//	logs.Identity = 2
		//}

		//memberIn, exist := c.Get("member_id")
		//if !(exist || userIn == nil) {
		//	memberId := toInt(memberIn)
		//	//userInfo, _ := service.NewMember(ctx, c).MemberInfo(memberId)
		//	if userInfo != nil {
		//		username = userInfo.Name
		//	}
		//	logs.Operatorsr = username
		//	logs.Identity = 1
		//}
		//err := service.NewLogs(ctx, c).Create(&logs)
		hlog.Info(log)
		//if err != nil {
		//	hlog.Error(err)
		//}

	}
}
func toInt(idIn interface{}) int64 {
	var idStr string
	var ok bool
	idStr, ok = idIn.(string)
	if !ok {
		idStr = "0"
	}
	id, _ := strconv.Atoi(idStr)
	return int64(id)
}
