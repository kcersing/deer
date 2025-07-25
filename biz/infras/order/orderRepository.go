package order

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"kcers-order/biz/dal/db/mysql/ent"
)

type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
	Update(ctx context.Context, order *Order) error
	FindById(ctx context.Context, Id int64) (*Order, error)
}
type OrderRepositoryImpl struct {
	db           *ent.Client
	snapshotFreq int
}

func NewOrderRepository(db *ent.Client, snapshotFreq int) OrderRepository {
	return &OrderRepositoryImpl{db: db, snapshotFreq: snapshotFreq}
}

func (o OrderRepositoryImpl) Save(ctx context.Context, order *Order) error {
	events := order.GetUncommittedEvents()
	if len(events) == 0 {
		return nil
	}
	tx, err := o.db.Tx(ctx)

	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	defer func() {
		if r := recover(); r != nil {
			err = tx.Rollback()
			if err != nil {
				klog.Errorf("rollback transaction failed: %v", err)
			}
			panic(r)
		} else if err != nil {
			err = tx.Rollback()
			if err != nil {
				klog.Errorf("rollback transaction failed: %v", err)
			}
		}
		err = tx.Commit()
		if err != nil {
			err = errors.Wrap(err, "committing transaction:")
			klog.Error("commit transaction failed", err)
		}
	}()

	for _, event := range events {
		eventData, err := json.Marshal(event)
		if err != nil {
			return errors.Wrap(err, "marshal event failed")
		}
		klog.Info("eventData", eventData)
		// 插入数据 event_store
	}
	// 插入数据 order_views

	// 插入数据 order_item

	//插入数据 order_snapshots

	return nil
}

func (o OrderRepositoryImpl) Update(ctx context.Context, order *Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepositoryImpl) FindById(ctx context.Context, Id int64) (*Order, error) {
	//TODO implement me
	panic("implement me")
}
