package order

import (
	"context"
	"errors"
	"kcers-order/biz/dal/db/mysql/ent"
	"kcers-order/biz/dal/db/mysql/ent/eventsubscriptions"
	"time"
)

// 订单状态变更 → 生成事件 → 仓储保存事件（事务）→ 事务提交 → 订阅服务通知订阅者 → 更新订阅状态

// SubscriptionService 订阅服务接口
type SubscriptionService interface {
	// ProcessEvent 处理事件并通知订阅者
	ProcessEvent(ctx context.Context, event Event) error
}

// subscriptionService 订阅服务实现
type subscriptionService struct {
	client *ent.Client // ent客户端
}

// NewSubscriptionService 创建订阅服务
func NewSubscriptionService(client *ent.Client) SubscriptionService {
	return &subscriptionService{
		client: client,
	}
}

// ProcessEvent 处理事件并更新订阅状态
func (s *subscriptionService) ProcessEvent(ctx context.Context, event Event) error {
	// 1. 查询事件类型的活跃订阅者
	subscribers, err := s.client.EventSubscriptions.
		Query().
		Where(
			eventsubscriptions.EventType(event.GetType()),
			eventsubscriptions.IsActive(1), // 仅活跃订阅
		).
		All(ctx)
	if err != nil {
		return errors.New("查询订阅者失败: " + err.Error())
	}

	// 2. 逐个通知订阅者
	for _, sub := range subscribers {
		// 模拟事件处理（实际应调用订阅者的回调/API）
		err := s.handleSubscriberEvent(ctx, sub, event)

		// 3. 更新订阅状态（无论成功/失败）
		update := s.client.EventSubscriptions.
			UpdateOneID(sub.ID).
			SetLastProcessedAt(time.Now())

		if err != nil {
			// 失败：更新错误信息和错误计数
			update.
				SetErrorCount(sub.ErrorCount + 1).
				SetLastError(err.Error())

			// 连续失败5次自动禁用订阅
			if sub.ErrorCount+1 >= 5 {
				update.SetIsActive(0)
			}
		} else {
			// 成功：更新最后处理事件ID和版本，重置错误计数
			update.
				SetLastProcessedID(event.GetID()).
				SetLastProcessedVersion(event.GetVersion()).
				SetErrorCount(0).
				SetLastError("")
		}

		// 执行更新
		if _, err := update.Save(ctx); err != nil {
			return errors.New("更新订阅状态失败: " + err.Error())
		}
	}
	return nil
}

// 模拟处理订阅者事件（实际应替换为HTTP调用/消息队列发送等）
func (s *subscriptionService) handleSubscriberEvent(ctx context.Context, sub *ent.EventSubscriptions, event Event) error {
	// 示例：打印事件处理日志（实际项目中替换为业务逻辑）
	// 如：调用库存服务的HTTP接口处理OrderCreated事件
	// 或：发送消息到Kafka队列供订阅者消费
	return nil
}
