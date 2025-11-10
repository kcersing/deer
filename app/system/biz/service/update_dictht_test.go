package service

import (
	"context"
	"gen/kitex_gen/base"
	"testing"
)

func TestUpdateDictht_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateDicthtService(ctx)
	// init req and assert value

	req := &base.Dictht{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
