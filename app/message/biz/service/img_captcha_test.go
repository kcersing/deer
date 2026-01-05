package service

import (
	"context"
	"testing"
)

func TestImgCaptcha_Run(t *testing.T) {
	ctx := context.Background()
	s := NewImgCaptchaService(ctx)
	// init req and assert value
	resp, err := s.Run()
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
