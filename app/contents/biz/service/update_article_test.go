package service

import (
	"context"
	contents "gen/kitex_gen/contents"
	"testing"
)

func TestUpdateArticle_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateArticleService(ctx)
	// init req and assert value

	req := &contents.UpdateArticleReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
