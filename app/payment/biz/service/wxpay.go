package service

import (
	"context"
	"system/biz/dal/wechat"

	base "gen/kitex_gen/base"
	payment "gen/kitex_gen/payment"

	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/order/request"
	"github.com/cloudwego/kitex/pkg/klog"
)

type WXPayService struct {
	ctx context.Context
}

// NewWXPayService new WXPayService
func NewWXPayService(ctx context.Context) *WXPayService {
	return &WXPayService{ctx: ctx}
}

// Run create note info
func (s *WXPayService) Run(req *payment.WXPayReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	options := &request2.RequestJSAPIPrepay{
		Amount: &request2.JSAPIAmount{
			Total:    int(req.Total * 100),
			Currency: "CNY",
		},

		Description: req.ProductName,
		OutTradeNo:  req.OrderSn, // 这里是商户订单号，不能重复提交给微信
		Payer: &request2.JSAPIPayer{
			OpenID: req.OpenId, // 用户的openid， 记得也是动态的。
		},
	}

	klog.Infof("wechat: %s", wechat.PaymentWechatApp)

	response2, err := wechat.PaymentWechatApp.Order.JSAPITransaction(s.ctx, options)
	klog.Infof("response2: %s", response2)
	if err != nil {
		klog.Infof("response2: %s", response2)
		return nil, err
	}

	payConf, err := wechat.PaymentWechatApp.JSSDK.BridgeConfig(response2.PrepayID, false)
	if err != nil {
		panic(err)
	}
	klog.Infof("payConf: %s", payConf)
	return nil, nil

}
