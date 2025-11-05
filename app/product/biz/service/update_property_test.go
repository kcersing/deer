package service

import (
	"context"
	product "gen/kitex_gen/product"
	"testing"
)

func TestUpdateProperty_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdatePropertyService(ctx)
	// init req and assert value

	req := &product.UpdatePropertyReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
