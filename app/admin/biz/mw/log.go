package mw

import (
	"admin/biz/infras/utils"
	"admin/rpc/client"
	"context"
	system "gen/kitex_gen/system"
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

		log.Success = int64(c.Response.Header.StatusCode())

		costTime := time.Since(start).Milliseconds()
		log.Time = int64(int32(costTime))

		log.Identity = utils.GetTokenId(ctx, c)
		_, err := client.SystemClient.CreateLog(ctx, &log)
		hlog.Info(log)
		if err != nil {
			hlog.Error(err)
		}

	}
}
