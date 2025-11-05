package service

import (
	"context"
	product "gen/kitex_gen/product"
	"testing"
)

func TestDecrStockRevert_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDecrStockRevertService(ctx)
	// init req and assert value

	req := &product.DecrStockReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
