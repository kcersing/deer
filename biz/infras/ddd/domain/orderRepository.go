package domain

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/pkg/errors"
	db2 "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
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
	return &OrderRepositoryImpl{db: db2.DB, snapshotFreq: snapshotFreq}
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
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		}
		err = tx.Commit()
	}()

	for _, event := range events {
		eventData, err := json.Marshal(event)
		if err != nil {
			return errors.Wrap(err, "marshal event failed")
		}
		hlog.Info("eventData", eventData)
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
