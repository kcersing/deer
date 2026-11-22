package service

import (
	"context"
	order "gen/kitex_gen/order"
	"order/biz/dal/db/ent"
	"order/biz/infras"

	"github.com/dtm-labs/client/dtmcli"
	extamqp "github.com/kcersing/eventbus/ext/amqp"
)

type CreateOrderService struct {
	ctx       context.Context
	publisher *extamqp.Publisher
}

// NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{
		ctx:       ctx,
		publisher: infras.GetManager().Publisher,
	}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderReq) (resp *order.OrderResp, err error) {

	return nil, err
}

func TransferFunds(ctx context.Context, client *ent.Client, fromID, toID int, amount float64) error {
	return WithTx(ctx, client, func(tx *ent.Tx) error {
		return nil
	})
}
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			panic(v)
		}
	}()
	err = fn(tx)
	if err != nil {
		return err
	}

	// 使用 DTM 管理分布式事务示例
	err = SubmitAndDo(ctx, "", func(bb *dtmcli.Saga) error {
		return client.OrderPay.UpdateOneID(1).AddPay(-100).Exec(ctx)
	})
	if err != nil {
		return err
	}

	// 执行事务操作并处理提交/回滚
	return tx.Commit() // 提交事务
}

//func TransferFunds(ctx context.Context, client *ent.Client, fromID, toID int, amount float64) error {
//	return ent.WithTx(ctx, client, func(tx *ent.Tx) error {
//		// 扣减转出方余额
//		if err := tx.Account.UpdateOneID(fromID).AddBalance(-amount).Exec(ctx); err != nil {
//			return err
//		}
//
//		// 增加转入方余额
//		if err := tx.Account.UpdateOneID(toID).AddBalance(amount).Exec(ctx); err != nil {
//			return err
//		}
//
//		// 记录交易流水
//		if _, err := tx.Transaction.Create().SetFromID(fromID).SetToID(toID).SetAmount(amount).Save(ctx); err != nil {
//			return err
//		}
//
//		return nil
//	})
//}

// 自定义封装的 SubmitAndDo 示例
func SubmitAndDo(ctx context.Context, dtmServer string, fn func(*dtmcli.Saga) error) error {
	gid := dtmcli.MustGenGid(dtmServer)
	saga := dtmcli.NewSaga(dtmServer, gid)
	if err := fn(saga); err != nil {
		return err
	}
	return saga.Submit()
}

//// Saga 模式分布式事务示例
//func TransferSaga(ctx context.Context, entClient *ent.Client, dtmServer string, fromID, toID int, amount int) error {
//	// 创建Saga实例
//	saga := dtmcli.NewSaga(dtmServer, dtmcli.MustGenGid(dtmServer))

//	// 添加转出分支：正向操作+补偿操作
//	saga.Add(
//		// 正向操作：转出扣减
//		"http://your-service/trans-out",
//		// 补偿操作：转出回滚
//		"http://your-service/trans-out-compensate",
//		map[string]interface{}{
//			"user_id": fromID,
//			"amount":  amount,
//		},
//	)
//
//	// 添加转入分支
//	saga.Add(
//		"http://your-service/trans-in",
//		"http://your-service/trans-in-compensate",
//		map[string]interface{}{
//			"user_id": toID,
//			"amount":  amount,
//		},
//	)
//
//	// 提交Saga事务
//	return saga.Submit(ctx)
//}
