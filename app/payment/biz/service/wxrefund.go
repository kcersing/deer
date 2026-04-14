package service

import (
	"context"
	"errors"
	"payment/conf"
	"system/biz/dal/wechat"

	base "gen/kitex_gen/base"
	payment "gen/kitex_gen/payment"

	"github.com/cloudwego/kitex/pkg/klog"

	request3 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/refund/request"
)

type WXRefundService struct {
	ctx context.Context
}

// NewWXRefundService new WXRefundService
func NewWXRefundService(ctx context.Context) *WXRefundService {
	return &WXRefundService{ctx: ctx}
}

// Run create note info
func (s *WXRefundService) Run(req *payment.WXRefundReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	options := &request3.RequestRefund{
		TransactionID: req.TransactionId,
		OutRefundNo:   req.OutRefundNo,
		Reason:        "退款",
		NotifyUrl:     conf.GetConf().Wechat.RefundNotifyUrl,
		//FundsAccount: "",
		Amount: &request3.RefundAmount{
			Refund:   int(req.Fee * 100),             // 退款金额，单位：分
			Total:    int(req.Total * 100),           // 订单总金额，单位：分
			From:     []*request3.RefundAmountFrom{}, // 退款出资账户及金额。不传仍然需要这个空数组防止微信报错
			Currency: "CNY",
		},
	}

	rs, err := wechat.PaymentWechatApp.Refund.Refund(s.ctx, options)
	if err != nil {
		klog.Errorf("退款 error: %s", err)
		return nil, err
	}

	if rs.Code != "" {
		if rs.Code == "NOT_ENOUGH" {
			klog.Errorf("退款失败：%s", rs.Message)
			return nil, errors.New("平台账户余额不足，请联系管理员")
		}
		return nil, errors.New(rs.Message)
	}
	klog.Infof("rs: %s", rs)
	return nil, nil
}
