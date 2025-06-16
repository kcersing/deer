package domain

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers/biz/dal/db/mysql/ent"
)

// Repository 仓库接口
type Repository interface {
	Save(order *Aggregate) error
	Load(id int64) (*Aggregate, error)
}

// EventStore 事件存储实现
type EventStore struct {
	client *ent.Client
}

func (es *EventStore) Save(a *Aggregate) error {
	for _, event := range a.Events {
		// 保存事件到数据库
		switch e := event.(type) {
		case *CreatedEvent:
			hlog.Infof("保存订单创建事件: %+v", e)
			_, err := es.client.Order.Create().Save(context.Background())

			if err != nil {
				return err
			}
		}
		a.Version++
	}
	a.Events = nil
	return nil
}
