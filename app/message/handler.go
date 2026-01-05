package main

import (
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"message/biz/service"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// ImgCaptcha implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) ImgCaptcha(ctx context.Context) (resp *base.NilResponse, err error) {
	resp, err = service.NewImgCaptchaService(ctx).Run()

	return resp, err
}

// Sms implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) Sms(ctx context.Context, req *base.IdReq) (resp *message.SmsResp, err error) {
	resp, err = service.NewSmsService(ctx).Run(req)

	return resp, err
}

// SmsList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SmsList(ctx context.Context, req *message.SendSmsListReq) (resp *message.SendSmsListResp, err error) {
	resp, err = service.NewSmsListService(ctx).Run(req)

	return resp, err
}

// SendSms implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendSms(ctx context.Context, req *message.SendSmsReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewSendSmsService(ctx).Run(req)

	return resp, err
}
