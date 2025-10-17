package service

import (
	"context"
	system "gen/kitex_gen/system"
	"testing"
)

func TestDicthtList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDicthtListService(ctx)
	// init req and assert value

	req := &system.DicthtListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
