package payment

import (
	"context"
	"system/biz/dal/wechat"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/notify/request"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func RefundNotify(ctx context.Context, c *app.RequestContext) {

	req, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
		hlog.Error(err)
		return
	}

	res, err := wechat.PaymentWechatApp.HandleRefundedNotify(req, func(message *request.RequestNotify, transaction *models.Refund, fail func(message string)) interface{} {
		if message.EventType != "REFUND.SUCCESS" {
			hlog.Errorf("退款失败：%s", message.EventType)
			return true
		}
		hlog.Infof("支付信息：%s", transaction)
		if transaction.OutTradeNo != "" {

			tr, _ := sonic.Marshal(transaction)

			hlog.Infof("订单号：%s 退款成功,退款信息：%s", transaction.OutTradeNo, tr)
		} else {
			fail("payment fail")
			return nil
		}
		return true
	})
	if err != nil {
		res.StatusCode = 502
		err = res.Write(c.Request.BodyWriter())
		return
		//panic(err)
	}

	// 这里根据之前返回的是true或者fail，框架这边自动会帮你回复微信
	err = res.Write(c.Request.BodyWriter())
	if err != nil {
		panic(err)
	}
}
