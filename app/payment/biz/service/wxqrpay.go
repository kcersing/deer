package service

import (
	"context"
	"system/biz/dal/wechat"

	base "gen/kitex_gen/base"
	payment "gen/kitex_gen/payment"

	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/order/request"
	"github.com/cloudwego/kitex/pkg/klog"
)

type WXQRPayService struct {
	ctx context.Context
}

// NewWXQRPayService new WXQRPayService
func NewWXQRPayService(ctx context.Context) *WXQRPayService {
	return &WXQRPayService{ctx: ctx}
}

// Run create note info
func (s *WXQRPayService) Run(req *payment.WXQRPayReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	description := "Image形象店-深圳腾大-QQ公仔"
	options := &request2.RequestNativePrepay{
		Amount: &request2.NativeAmount{
			//Total:    int(req.Total * 100),
			Currency: "CNY",
		},
		Description: description,
		//OutTradeNo:  req.OrderSn, // 这里是商户订单号，不能重复提交给微信
	}

	response2, err := wechat.PaymentWechatApp.Order.TransactionNative(s.ctx, options)

	if err != nil {
		klog.Infof("error: %s", err)
		//return response2, err
	}
	klog.Infof("response2: %s", response2)
	return nil, nil
}
