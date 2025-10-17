package service

import (
	"context"
	system "gen/kitex_gen/system"
	"testing"
)

func TestUpdateDictht_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateDicthtService(ctx)
	// init req and assert value

	req := &system.Dictht{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
