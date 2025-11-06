package service

import (
	"context"
	product "gen/kitex_gen/product"
	"testing"
)

func TestItemList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewItemListService(ctx)
	// init req and assert value

	req := &product.ItemListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
