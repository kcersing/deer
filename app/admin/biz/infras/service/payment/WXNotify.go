package payment

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/notify/request"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"system/biz/dal/wechat"
)

func Notify(ctx context.Context, c *app.RequestContext) {

	req, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
		hlog.Error(err)
		return
	}

	res, err := wechat.PaymentWechatApp.HandlePaidNotify(req, func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{} {

		// 看下支付通知事件状态
		// 这里可能是微信支付失败的通知，所以可能需要在数据库做一些记录，然后告诉微信我处理完成了。
		if message.EventType != "TRANSACTION.SUCCESS" {
			hlog.Errorf("支付失败：%s", message.EventType)
			return true
		}
		hlog.Infof("支付信息：%s", transaction)

		if transaction.OutTradeNo != "" {

			tr, _ := sonic.Marshal(transaction)

			// 这里对照自有数据库里面的订单做查询以及支付状态改变
			hlog.Infof("订单号：%s 支付成功，支付信息：%s", transaction.OutTradeNo, tr)
		} else {
			// 因为微信这个回调不存在订单号，所以可以告诉微信我还没处理成功，等会它会重新发起通知
			// 如果不需要，直接返回true即可
			fail("payment fail")
			return nil
		}
		return true
	},
	)

	// 这里可能是因为不是微信官方调用的，无法正常解析出transaction和message，所以直接抛错。
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
