package service

import (
	"common/amqpclt"
	"common/eventbus"
	"context"
	"fmt"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"message/biz/dal/mq"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// 全局事件总线和桥接器
var (
	globalEventBus *eventbus.EventBus
	globalBridge   *eventbus.AMQPBridge
	once           sync.Once
)

type SendMemberMessagesService struct {
	ctx context.Context
}

// NewSendMemberMessagesService new SendMemberMessagesService
func NewSendMemberMessagesService(ctx context.Context) *SendMemberMessagesService {
	return &SendMemberMessagesService{ctx: ctx}
}

// InitGlobalEventBus 全局初始化（应在应用启动时调用）
func InitGlobalEventBus() error {
	var err error
	once.Do(func() {
		mq.InitMQ()

		publisher, err := amqpclt.NewPublisher(mq.Client, "events")
		if err != nil {
			klog.Errorf("failed to create publisher: %v", err)
			return
		}

		subscriber, err := amqpclt.NewSubscribe(mq.Client, "events")
		if err != nil {
			klog.Errorf("failed to create subscriber: %v", err)
			return
		}

		globalEventBus = eventbus.NewEventBus()
		globalBridge = eventbus.NewAMQPBridge(globalEventBus, publisher, subscriber)

		// 注册中间件：日志 + AMQP 发布
		globalEventBus.Use(eventbus.LoggingPlugin())
		globalEventBus.Use(globalBridge.AMQPPublishingMiddleware())

		// 启动后台监听
		ctx := context.Background()
		globalBridge.StartListener(ctx)

		klog.Infof("[MessageService] Global event bus initialized")
	})

	return err
}

// GetGlobalAMQPBridge 获取全局桥接器
func GetGlobalAMQPBridge() *eventbus.AMQPBridge {
	return globalBridge
}

// Run 批量发送消息给多个接收者（通过事件总线 + AMQP）
// 特点：
// - 内存事件自动同步到 RabbitMQ
// - 支持大规模接收者
// - 异步处理，不阻塞业务流程
func (s *SendMemberMessagesService) Run(req *message.SendMemberMessagesReq) (resp *base.NilResponse, err error) {
	// 确保全局事件总线已初始化
	if globalBridge == nil {
		if err := InitGlobalEventBus(); err != nil {
			klog.Errorf("failed to initialize event bus: %v", err)
			return nil, fmt.Errorf("event bus initialization failed: %w", err)
		}
	}

	// 获取接收者列表
	recipients := s.getRecipients(req)
	if len(recipients) == 0 {
		klog.Warn("[SendMemberMessagesService] no recipients to send messages")
		return &base.NilResponse{}, nil
	}

	// 方案 A: 通过事件总线发送（推荐）
	// 特点：事件自动同步到 AMQP，支持中间件处理
	return s.sendViaEventBus(req, recipients)

	// 方案 B: 若需要直接 AMQP 发送，可使用以下方法（备用）
	// return s.sendViaDirectAMQP(req, recipients)
}

// sendViaEventBus 通过事件总线发送（推荐方案）
func (s *SendMemberMessagesService) sendViaEventBus(req *message.SendMemberMessagesReq, recipients []string) (*base.NilResponse, error) {
	ctx := context.Background()

	// 创建事件
	events := s.prepareEventBusEvents(req, recipients)
	klog.Infof("[SendMemberMessagesService] prepared %d events for %d recipients", len(events), len(recipients))

	// 异步批量发布到事件总线（会自动转发到 AMQP）
	resultCh := globalBridge.PublishAsyncBatchToAMQP(ctx, events)

	// 异步收集结果
	go func() {
		successCount := 0
		failedCount := 0

		for result := range resultCh {
			if result.Error != nil {
				failedCount++
				klog.Errorf("[SendMemberMessagesService] event %d failed: %v", result.Index, result.Error)
			} else {
				successCount++
			}
		}

		if failedCount > 0 {
			klog.Warnf("[SendMemberMessagesService] batch publish: %d success, %d failed", successCount, failedCount)
		} else {
			klog.Infof("[SendMemberMessagesService] batch publish: all %d events published successfully", successCount)
		}
	}()

	return &base.NilResponse{}, nil
}

// sendViaDirectAMQP 直接通过 AMQP 发送（备用方案）
func (s *SendMemberMessagesService) sendViaDirectAMQP(req *message.SendMemberMessagesReq, recipients []string) (*base.NilResponse, error) {
	// 创建发布者
	publisher, err := amqpclt.NewPublisher(mq.Client, "messages")
	if err != nil {
		klog.Errorf("[SendMemberMessagesService] cannot create publisher: %v", err)
		return nil, fmt.Errorf("publisher creation failed: %w", err)
	}
	defer publisher.Close()

	// 准备批量消息
	batchMessages := s.prepareBatchMessages(req, recipients)

	// 异步发布
	resultCh := make(chan amqpclt.PublishResult, len(batchMessages))
	publisher.PublishAsync(s.ctx, batchMessages, resultCh)

	// 收集结果
	successCount := 0
	failedCount := 0

	for result := range resultCh {
		if result.Error != nil {
			failedCount++
			klog.Errorf("[SendMemberMessagesService] message %d failed: %v", result.Index, result.Error)
		} else {
			successCount++
		}
	}

	if failedCount > 0 {
		klog.Warnf("[SendMemberMessagesService] batch publish: %d success, %d failed", successCount, failedCount)
	} else {
		klog.Infof("[SendMemberMessagesService] batch publish: all %d messages published successfully", successCount)
	}

	return &base.NilResponse{}, nil
}

// getRecipients 获取消息接收者列表（从请求或数据库读取）
func (s *SendMemberMessagesService) getRecipients(req *message.SendMemberMessagesReq) []string {
	// TODO: 实现从请求或数据库获取接收者 ID
	// 示例实现：模拟 10 个接收者
	recipients := make([]string, 10)
	for i := 0; i < 10; i++ {
		recipients[i] = fmt.Sprintf("user_%d", i+1)
	}
	return recipients
}

// prepareEventBusEvents 为事件总线准备事件
func (s *SendMemberMessagesService) prepareEventBusEvents(req *message.SendMemberMessagesReq, recipients []string) []*eventbus.Event {
	events := make([]*eventbus.Event, len(recipients))

	// 消息内容（从请求中获取或生成）
	messageContent := map[string]interface{}{
		"title":   "群发消息标题",
		"content": "这是一条群发消息内容",
		"sendAt":  time.Now().Unix(),
	}

	for i, recipientID := range recipients {
		events[i] = &eventbus.Event{
			Id:    fmt.Sprintf("msg-%s-%d", recipientID, time.Now().UnixNano()),
			Topic: "member_message_send",
			Payload: map[string]interface{}{
				"recipient_id": recipientID,
				"title":        messageContent["title"],
				"content":      messageContent["content"],
				"sendAt":       messageContent["sendAt"],
			},
			Timestamp: time.Now(),
			Source:    "service",
			Version:   1,
		}
	}

	return events
}

// prepareBatchMessages 为 AMQP 直接发送准备批量消息
func (s *SendMemberMessagesService) prepareBatchMessages(req *message.SendMemberMessagesReq, recipients []string) []amqpclt.BatchMessage {
	batchMessages := make([]amqpclt.BatchMessage, len(recipients))

	messageContent := map[string]interface{}{
		"title":   "群发消息标题",
		"content": "这是一条群发消息内容",
		"sendAt":  time.Now().Unix(),
	}

	for i, recipientID := range recipients {
		batchMessages[i] = amqpclt.BatchMessage{
			RoutingKey:    recipientID,
			CorrelationId: fmt.Sprintf("msg-%s-%d", recipientID, time.Now().UnixNano()),
			Message: amqpclt.Message{
				Event:     "member_message_send",
				Payload:   messageContent,
				Timestamp: time.Now(),
			},
		}
	}

	return batchMessages
}