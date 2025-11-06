package service

import (
	"context"
	product "gen/kitex_gen/product"
	"testing"
)

func TestUpdateItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateItemService(ctx)
	// init req and assert value

	req := &product.UpdateItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
