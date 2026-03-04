package infras

import (
	"context"
	"errors"
	"fmt"
	"order/biz/infras/common"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// EventHandler 事件处理器接口
type EventHandler interface {
	Handle(ctx context.Context, event common.Event) error
}

type Dispatcher interface {
	RegisterHandler(eventType string, handler EventHandler)
	Unregister(eventType string, handler EventHandler)
	Dispatch(ctx context.Context, event common.Event) error
}

// EventDispatcher 分发 eventbus
type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

// RegisterHandler 注册
func (d *EventDispatcher) RegisterHandler(eventType string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

// Unregister 移除
func (d *EventDispatcher) Unregister(eventType string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for i, obs := range d.handlers[eventType] {
		if obs == handler {
			d.handlers[eventType] = append(d.handlers[eventType][:i], d.handlers[eventType][i+1:]...)
			break
		}
	}
}

// Dispatch 并发分发事件
func (d *EventDispatcher) Dispatch(ctx context.Context, event common.Event) error {
	d.mu.RLock()
	handlers, ok := d.handlers[event.GetType()]
	d.mu.RUnlock()
	if !ok || len(handlers) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(handlers))

	for _, handler := range handlers {
		klog.Infof("开始处理事件: %s", event.GetType())
		wg.Add(1)
		go func(h EventHandler) {
			defer wg.Done()
			if err := withRetry(func() error {
				// 为每个处理器创建一个独立的、带超时的上下文
				ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
				defer cancel()
				return h.Handle(ctx2, event)
			}, 3); err != nil {
				// 包装错误，提供更多上下文
				errCh <- fmt.Errorf("handler %T failed: %w", h, err)
			}
		}(handler)
	}

	// 等待所有处理器完成并关闭错误通道
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// 收集错误
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	// 优化：使用更详细的方式聚合错误
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// withRetry 重试机制
func withRetry(fn func() error, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		if err = fn(); err == nil {
			return nil
		}
		klog.Errorf("Attempt %d failed: %v. Retrying...", i+1, err)
		// 指数退避
		time.Sleep(time.Millisecond * 100 * time.Duration(i+1))
	}
	return fmt.Errorf("failed after %d attempts: %w", maxRetries, err)
}

// 添加单例实例和同步控制变量
var (
	dispatcherInstance *EventDispatcher
	once               sync.Once
)

func InitEventHandlers() *EventDispatcher {
	// 使用sync.Once确保初始化代码只执行一次
	once.Do(func() {
		dispatcherInstance = NewEventDispatcher()
		inventoryHandler := &InventoryHandler{}
		// 订阅更具体的事件类型
		dispatcherInstance.RegisterHandler(string(common.Created), inventoryHandler)
		dispatcherInstance.RegisterHandler(string(common.Cancelled), inventoryHandler)
	})
	return dispatcherInstance
}
