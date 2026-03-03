package infras

import (
	"common/eventbus"
	"context"
	"encoding/json"
	"log"
	"order/biz/dal/db"
	"order/biz/dal/db/ent"
	"order/biz/dal/db/ent/ordereventsubscriptions"
	"order/biz/infras/common"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/cloudwego/kitex/pkg/klog"
)

const maxErrorCount = 5 // 最大失败次数，超过则禁用

// SubscriptionListener 监听事件并通知订阅者
type SubscriptionListener struct {
	client   *ent.Client
	bus      eventbus.EventBus
	subCache *sync.Map // 优化1: 订阅者缓存
}

// NewSubscriptionListener 创建一个新的订阅监听器
func NewSubscriptionListener(bus eventbus.EventBus) *SubscriptionListener {
	listener := &SubscriptionListener{
		client:   db.Client,
		bus:      bus,
		subCache: &sync.Map{},
	}
	// 启动后台goroutine来刷新缓存
	go listener.runCacheRefresher(context.Background(), 5*time.Minute)
	return listener
}

// runCacheRefresher 定期从数据库加载所有活跃的订阅者并更新缓存
func (s *SubscriptionListener) runCacheRefresher(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 立即执行一次以填充初始缓存
	s.refreshCache(ctx)

	for {
		select {
		case <-ticker.C:
			s.refreshCache(ctx)
		case <-ctx.Done():
			klog.Info("Cache refresher stopped.")
			return
		}
	}
}

// refreshCache 执行一次缓存刷新操作
func (s *SubscriptionListener) refreshCache(ctx context.Context) {
	klog.Info("Refreshing subscription cache...")
	allSubs, err := s.client.OrderEventSubscriptions.Query().Where(ordereventsubscriptions.IsActive(1)).All(ctx)
	if err != nil {
		klog.Errorf("Failed to refresh subscription cache: %v", err)
		return
	}

	// 按 event_type 分组
	groupedSubs := make(map[string][]*ent.OrderEventSubscriptions)
	for _, sub := range allSubs {
		groupedSubs[sub.EventType] = append(groupedSubs[sub.EventType], sub)
	}

	// 更新缓存
	for topic, subs := range groupedSubs {
		s.subCache.Store(topic, subs)
	}
	klog.Infof("Subscription cache refreshed. Loaded %d subscribers for %d topics.", len(allSubs), len(groupedSubs))
}

// Start 开始监听事件
func (s *SubscriptionListener) Start() {
	// 订阅订单创建事件
	topic := string(common.Created)
	s.bus.SubscribeAsync(
		topic,
		eventbus.EventHandlerFunc(s.handleEvent),
		1, // 消费者数量
	)
	klog.Infof("SubscriptionListener started, waiting for topic: %s", topic)
}

// handleEvent 是通用的事件处理器
func (s *SubscriptionListener) handleEvent(ctx context.Context, event *eventbus.Event) error {
	klog.Infof("Received event from topic: %s", event.Topic)

	// 优化1: 从缓存中获取订阅者，而不是查询数据库
	cachedSubs, ok := s.subCache.Load(event.Topic)
	if !ok {
		klog.Infof("No active subscribers in cache for event: %s", event.Topic)
		return nil
	}

	subscribers, ok := cachedSubs.([]*ent.OrderEventSubscriptions)
	if !ok || len(subscribers) == 0 {
		klog.Infof("No active subscribers in cache for event: %s", event.Topic)
		return nil
	}

	// 并发通知所有订阅者
	var wg sync.WaitGroup
	for _, sub := range subscribers {
		wg.Add(1)
		go func(sub *ent.OrderEventSubscriptions) {
			defer wg.Done()

			// 优化2: 调用带有重试逻辑的通知方法
			notificationErr := s.handleSubscriberNotification(ctx, sub, event)

			// 更新订阅状态
			s.updateSubscriptionState(ctx, sub, event, notificationErr)
		}(sub)
	}
	wg.Wait()

	return nil
}

// handleSubscriberNotification 包含重试和死信队列逻辑
func (s *SubscriptionListener) handleSubscriberNotification(ctx context.Context, sub *ent.OrderEventSubscriptions, event *eventbus.Event) error {
	// 定义一个指数退避重试策略
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 1 * time.Minute // 最多重试1分钟

	var notificationErr error

	//operation := func() error {
	//	payloadBytes, err := s.getPayloadAsBytes(event)
	//	if err != nil {
	//		// 这是一个不可重试的错误，因为事件本身有问题
	//		return backoff.Permanent(fmt.Errorf("failed to marshal event payload: %w", err))
	//	}
	//
	//	// 真正的通知逻辑
	//	resp, err := http.Post(sub.CallbackURL, "application/json", bytes.NewBuffer(payloadBytes))
	//	if err != nil {
	//		klog.Warnf("Attempt to notify %s failed: %v. Retrying...", sub.CallbackURL, err)
	//		return err // 返回错误，触发重试
	//	}
	//	defer resp.Body.Close()
	//
	//	if resp.StatusCode >= 500 { // 5xx 错误是典型的服务器端临时错误，应该重试
	//		err = fmt.Errorf("subscriber %s returned server error: %s", sub.CallbackURL, resp.Status)
	//		klog.Warn(err.Error())
	//		return err
	//	}
	//
	//	// 4xx 错误是客户端错误，不应该重试
	//	if resp.StatusCode >= 400 {
	//		return backoff.Permanent(fmt.Errorf("subscriber %s returned client error: %s", sub.CallbackURL, resp.Status))
	//	}
	//
	//	klog.Infof("Successfully notified subscriber %s for event %s", sub.CallbackURL, event.Topic)
	//	return nil // 成功
	//}
	//
	//notificationErr = backoff.Retry(operation, bo)
	//
	//if notificationErr != nil {
	//	klog.Errorf("Notification to %s failed after multiple retries: %v", sub.CallbackURL, notificationErr)
	//	// 将这个失败的任务发送到死信队列
	//	s.sendToDeadLetterQueue(sub, event)
	//}

	return notificationErr
}

// sendToDeadLetterQueue 是一个占位符，用于处理无法投递的事件
func (s *SubscriptionListener) sendToDeadLetterQueue(sub *ent.OrderEventSubscriptions, event *eventbus.Event) {
	// 在实际应用中，这里会将事件和订阅者信息发送到一个专门的RabbitMQ队列或记录到特定的日志/数据库表中
	//klog.Criticalf("DEAD LETTER: Failed to notify subscriber %d (%s) for event %s. Payload: %v", sub.ID, sub.CallbackURL, event.Topic, event.Payload)
}

// getPayloadAsBytes 确保事件的payload是[]byte类型
func (s *SubscriptionListener) getPayloadAsBytes(event *eventbus.Event) ([]byte, error) {
	if payloadBytes, ok := event.Payload.([]byte); ok {
		return payloadBytes, nil
	}
	return json.Marshal(event.Payload)
}

func (s *SubscriptionListener) updateSubscriptionState(ctx context.Context, sub *ent.OrderEventSubscriptions, event *eventbus.Event, notifErr error) {
	update := s.client.OrderEventSubscriptions.
		UpdateOneID(sub.ID).
		SetLastProcessedAt(time.Now())

	payloadBytes, _ := s.getPayloadAsBytes(event)
	var eventID string
	if payloadBytes != nil {
		var payloadMap map[string]interface{}
		if err := json.Unmarshal(payloadBytes, &payloadMap); err == nil {
			if id, ok := payloadMap["EventID"].(string); ok {
				eventID = id
			}
		}
	}

	if notifErr != nil {
		// 注意：这里的错误信息现在可能来自 backoff 库，更具描述性
		log.Printf("最终处理订阅者 %d 的事件 %s 失败: %v", sub.ID, event.Topic, notifErr)
		newErrorCount := sub.ErrorCount + 1
		update.SetErrorCount(newErrorCount).SetLastError(notifErr.Error())

		if newErrorCount >= maxErrorCount {
			update.SetIsActive(0)
			log.Printf("订阅者 %d 因连续失败 %d 次被自动禁用", sub.ID, newErrorCount)
		}
	} else {
		if eventID != "" {
			update.SetLastProcessedID(eventID)
		}
		update.SetErrorCount(0).SetLastError("")
	}

	if _, updateErr := update.Save(ctx); updateErr != nil {
		log.Printf("更新订阅者 %d 的状态失败: %v", sub.ID, updateErr)
	}
}
