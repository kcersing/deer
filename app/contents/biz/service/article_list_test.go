package service

import (
	"context"
	contents "gen/kitex_gen/contents"
	"testing"
)

func TestArticleList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewArticleListService(ctx)
	// init req and assert value

	req := &contents.ArticleListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
