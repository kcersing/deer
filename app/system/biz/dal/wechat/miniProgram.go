package wechat

import (
	"common/consts"
	"sync"
	"system/conf"

	"github.com/ArtisanCloud/PowerLibs/v3/logger/drivers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var MiniProgramApp *miniProgram.MiniProgram

var onceMiniProgramApp sync.Once

func InitMiniProgramApp() {
	onceMiniProgramApp.Do(func() {
		MiniProgramApp = NewMiniMiniProgramService()
	})
}

func NewMiniMiniProgramService() *miniProgram.MiniProgram {
	conf := conf.GetConf().Wechat
	var cache kernel.CacheInterface

	if conf.GetConf().Redis.Address != "" {
	cache = kernel.NewRedisClient(&kernel.UniversalOptions{
		Addrs:    []string{conf.GetConf().Redis.Address},
		Password: conf.GetConf().Redis.Password,
		DB:       6,
	})
	}
	wechatFilePath := consts.WechatFilePath

	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:        conf.Appid,
		Secret:       conf.AppSecret,
		ResponseType: response.TYPE_MAP,
		Http:         miniProgram.Http{},
		Log: miniProgram.Log{
			Driver: &drivers.SimpleLogger{},
			Level:  "debug",
			File:   wechatFilePath + "/mini_log.log",
			Stdout: false,
		},
		Cache:     cache,
		HttpDebug: false,
		Debug:     false,
	})

	if err != nil {
		hlog.Error(err)
	}

	return app
}
