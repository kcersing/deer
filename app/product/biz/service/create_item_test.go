package service

import (
	"context"
	product "gen/kitex_gen/product"
	"testing"
)

func TestCreateItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateItemService(ctx)
	// init req and assert value

	req := &product.CreateItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
