package infras

import (
	"context"
	"fmt"
	"log"
	"order/biz/dal/db/ent"
	"order/biz/dal/db/ent/ordereventsubscriptions"
	"order/biz/infras/common"
	"sync"
	"time"
)

// 订单状态变更 → 生成事件 → 仓储保存事件（事务）→ 事务提交 → 订阅服务通知订阅者 → 更新订阅状态

const maxErrorCount = 5 // 最大失败次数，超过则禁用

// SubscriptionService 订阅服务接口
type SubscriptionService interface {
	// ProcessEvent 处理事件并通知订阅者
	ProcessEvent(ctx context.Context, event common.Event) error
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
func (s *subscriptionService) ProcessEvent(ctx context.Context, event common.Event) error {
	// 1. 查询事件类型的活跃订阅者
	subscribers, err := s.client.OrderEventSubscriptions.
		Query().
		Where(
			ordereventsubscriptions.EventType(event.GetType()),
			ordereventsubscriptions.IsActive(1), // 仅活跃订阅
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("查询事件 %s 的订阅者失败: %w", event.GetType(), err)
	}

	if len(subscribers) == 0 {
		return nil
	}

	// 2. 并发通知所有订阅者
	var wg sync.WaitGroup
	for _, sub := range subscribers {
		wg.Add(1)
		go func(sub *ent.OrderEventSubscriptions) {
			defer wg.Done()

			// 模拟事件处理（实际应调用订阅者的回调/API）
			err := s.handleSubscriberEvent(ctx, sub, event)

			// 3. 更新订阅状态（无论成功/失败）
			update := s.client.OrderEventSubscriptions.
				UpdateOneID(sub.ID).
				SetLastProcessedAt(time.Now())

			if err != nil {
				log.Printf("处理订阅者 %d 的事件 %s 失败: %v", sub.ID, event.GetType(), err)
				// 失败：更新错误信息和错误计数
				newErrorCount := sub.ErrorCount + 1
				update.SetErrorCount(newErrorCount).SetLastError(err.Error())

				// 连续失败达到阈值则自动禁用订阅
				if newErrorCount >= maxErrorCount {
					update.SetIsActive(0)
					log.Printf("订阅者 %d 因连续失败 %d 次被自动禁用", sub.ID, newErrorCount)
				}
			} else {
				// 成功：更新最后处理事件ID和版本，重置错误计数
				update.SetLastProcessedID(event.GetId()).
					SetLastProcessedVersion(event.GetVersion()).
					SetErrorCount(0).
					SetLastError("")
			}

			// 执行更新
			if _, updateErr := update.Save(ctx); updateErr != nil {
				// 此处只记录日志，不中断其他订阅者的处理
				log.Printf("更新订阅者 %d 的状态失败: %v", sub.ID, updateErr)
			}
		}(sub)
	}
	wg.Wait() // 等待所有通知处理完成

	return nil
}

// 模拟处理订阅者事件（实际应替换为HTTP调用/消息队列发送等）
func (s *subscriptionService) handleSubscriberEvent(ctx context.Context, sub *ent.OrderEventSubscriptions, event common.Event) error {
	// 示例：打印事件处理日志（实际项目中替换为业务逻辑）
	// 如：调用库存服务的HTTP接口处理OrderCreated事件
	// 或：发送消息到Kafka队列供订阅者消费
	return nil
}
