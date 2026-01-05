package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type ImgCaptchaService struct {
	ctx context.Context
}

// NewImgCaptchaService new ImgCaptchaService
func NewImgCaptchaService(ctx context.Context) *ImgCaptchaService {
	return &ImgCaptchaService{ctx: ctx}
}

// Run create note info
func (s *ImgCaptchaService) Run() (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
