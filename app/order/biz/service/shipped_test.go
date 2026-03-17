package service

import (
	"context"
	base "gen/kitex_gen/base"
	"order/biz/dal/db"
	"order/biz/dal/mq"
	"order/biz/dal/redis"
	"testing"
)

func init() {
	redis.Init()
	db.InitDB()
	mq.InitMQ()
}
func TestShipped_Run(t *testing.T) {

	ctx := context.Background()
	s := NewShippedService(ctx)
	// init req and assert value

	req := &base.IdReq{Id: 1}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
