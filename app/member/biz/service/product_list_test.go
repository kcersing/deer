package service

import (
	"context"

	"gen/kitex_gen/member"
	"testing"
)

func TestProductList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewProductListService(ctx)
	// init req and assert value

	req := &member.ProductListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
