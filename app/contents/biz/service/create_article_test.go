package service

import (
	"context"
	contents "gen/kitex_gen/contents"
	"testing"
)

func TestCreateArticle_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateArticleService(ctx)
	// init req and assert value

	req := &contents.CreateArticleReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
