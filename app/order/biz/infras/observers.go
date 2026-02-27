package infras

import (
	"context"
	"order/biz/infras/common"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
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
				ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
				defer cancel()
				return h.Handle(ctx2, event)
			}, 3); err != nil {
				errCh <- err
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

	if len(errs) > 0 {
		return errors.Errorf("共%d个事件处理器执行失败", len(errs))
	}
	return nil
}

// withRetry 重试机制
func withRetry(fn func() error, maxRetries int) error {
	var err error
	for i := range maxRetries {
		if err = fn(); err == nil {
			return nil
		}
		klog.Errorf("重试第%d次失败: %v", i+1, err)
		time.Sleep(time.Millisecond * 100 * time.Duration(i+1)) // 指数退避
	}
	return err
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
		dispatcherInstance.RegisterHandler("created", inventoryHandler)
		dispatcherInstance.RegisterHandler("cancelled", inventoryHandler)
	})
	return dispatcherInstance
}
