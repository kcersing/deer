package service

import (
	"context"
	Base "gen/kitex_gen/base"
	"testing"
)

func TestCreateDictht_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateDicthtService(ctx)
	// init req and assert value

	req := &Base.Dictht{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
