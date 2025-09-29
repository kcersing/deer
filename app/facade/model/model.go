package model

import (
	"common/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var (
	UserAuthMiddleware *jwt.HertzJWTMiddleware
	DeerAuthMiddleware *jwt.HertzJWTMiddleware
)

type Response struct {
	Time      string      `json:"time"`
	Code      int64       `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Total     int64       `json:"total"`
	CacheTime string      `json:"cacheTime"`
}

func SendResponse(c *app.RequestContext, err error, data interface{}, Total int64, cacheTime string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:      Err.ErrCode,
		Message:   Err.ErrMsg,
		Data:      data,
		Total:     Total,
		Time:      time.Now().Format(time.DateTime),
		CacheTime: cacheTime,
	})
}
