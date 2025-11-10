package service

import (
	"context"
	base "gen/kitex_gen/base"
	"testing"
)

func TestUpdateDict_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateDictService(ctx)
	// init req and assert value

	req := &base.Dict{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
