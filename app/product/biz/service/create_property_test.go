package service

import (
	"context"
	product "gen/kitex_gen/product"
	"testing"
)

func TestCreateProperty_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreatePropertyService(ctx)
	// init req and assert value

	req := &product.CreatePropertyReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
