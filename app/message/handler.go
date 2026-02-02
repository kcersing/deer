package main

import (
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"message/biz/service"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// Sms implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) Sms(ctx context.Context, req *base.IdReq) (resp *message.SmsResp, err error) {
	resp, err = service.NewSmsService(ctx).Run(req)

	return resp, err
}

// SmsSendList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SmsSendList(ctx context.Context, req *message.SmsSendListReq) (resp *message.SmsSendListResp, err error) {
	resp, err = service.NewSmsSendListService(ctx).Run(req)

	return resp, err
}

// SendSms implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendSms(ctx context.Context, req *message.SendSmsReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewSendSmsService(ctx).Run(req)

	return resp, err
}

// MessagesList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessagesList(ctx context.Context, req *message.MessagesListReq) (resp *message.MessagesListResp, err error) {
	resp, err = service.NewMessagesListService(ctx).Run(req)

	return resp, err
}

// SendMemberMessages implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendMemberMessages(ctx context.Context, req *message.SendMemberMessagesReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewSendMemberMessagesService(ctx).Run(req)

	return resp, err
}

// SendUserMessages implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendUserMessages(ctx context.Context, req *message.SendUserMessagesReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewSendUserMessagesService(ctx).Run(req)

	return resp, err
}

// MessagesSendList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessagesSendList(ctx context.Context, req *message.MessagesListReq) (resp *message.MessagesSendListResp, err error) {
	resp, err = service.NewMessagesSendListService(ctx).Run(req)

	return resp, err
}

// DeleteMessages implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) DeleteMessages(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	// TODO: Your code here...
	return
}
