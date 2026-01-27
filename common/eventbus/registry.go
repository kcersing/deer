package eventbus

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 消费者注册表 ============

// ConsumerInfo 消费者信息
type ConsumerInfo struct {
	Topic       string  // 事件主题（如 "send_user_messages"）
	HandlerName string  // 处理器名字（如 "send_user_handler"）
	WorkerNum   int32   // 工作线程数（如 10）
	Handler     Handler // 处理器对象（实现 Handle() 方法）
}

// ConsumerRegistry 消费者注册表 - 集中管理消费者
type ConsumerRegistry struct {
	mu        sync.RWMutex
	handlers  map[string]Handler         // handlerName → Handler
	consumers map[string][]*ConsumerInfo // topic → []ConsumerInfo
	pools     map[string]*ConsumerPool   // handlerName → ConsumerPool
}

// NewConsumerRegistry 创建消费者注册表
func NewConsumerRegistry() *ConsumerRegistry {
	return &ConsumerRegistry{
		handlers:  make(map[string]Handler),
		consumers: make(map[string][]*ConsumerInfo),
		pools:     make(map[string]*ConsumerPool),
	}
}

// RegisterHandler 注册处理器
func (cr *ConsumerRegistry) RegisterHandler(name string, handler Handler) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	if _, exists := cr.handlers[name]; exists {
		return fmt.Errorf("handler %s already exists", name)
	}

	cr.handlers[name] = handler
	klog.Infof("[Registry] Handler registered: %s\n", name)
	return nil
}

// RegisterConsumer 注册消费者
func (cr *ConsumerRegistry) RegisterConsumer(topic string, handlerName string, workerNum int32) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	handler, exists := cr.handlers[handlerName]
	if !exists {
		
		return fmt.Errorf("handler %s not found", handlerName)
	}

	info := &ConsumerInfo{
		Topic:       topic,
		HandlerName: handlerName,
		WorkerNum:   workerNum,
		Handler:     handler,
	}

	cr.consumers[topic] = append(cr.consumers[topic], info)
	klog.Infof("[Registry] Consumer registered: topic=%s, handler=%s, workers=%d\n",
		topic, handlerName, workerNum)
	return nil
}

// StartAll 启动所有消费者
func (cr *ConsumerRegistry) StartAll(eb *EventBus) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	// 遍历所有 topic
	for _, consumers := range cr.consumers {
		// 遍历该 topic 的所有消费者
		for _, info := range consumers {
			// 使用消费者池处理
			pool := eb.SubscribeWithPool(info.Topic, info.Handler, info.WorkerNum)
			cr.pools[info.HandlerName] = pool
			klog.Infof("[Registry] Consumer pool started: %s with %d workers\n",
				info.HandlerName, info.WorkerNum)
		}
	}
	return nil
}

// GetConsumers 获取主题的消费者列表
func (cr *ConsumerRegistry) GetConsumers(topic string) []*ConsumerInfo {
	cr.mu.RLock()
	defer cr.mu.RUnlock()

	if consumers, ok := cr.consumers[topic]; ok {
		return consumers
	}
	return []*ConsumerInfo{}
}

// GetAllConsumers 获取所有消费者
func (cr *ConsumerRegistry) GetAllConsumers() map[string][]*ConsumerInfo {
	cr.mu.RLock()
	defer cr.mu.RUnlock()

	result := make(map[string][]*ConsumerInfo)
	for topic, consumers := range cr.consumers {
		result[topic] = append([]*ConsumerInfo{}, consumers...)
	}
	return result
}

// Shutdown 关闭所有消费者
func (cr *ConsumerRegistry) Shutdown(ctx context.Context) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	for _, pool := range cr.pools {
		pool.Stop()
	}
	return nil
}
