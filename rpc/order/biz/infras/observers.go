package infras

import (
	"context"
	"deer/rpc/order/biz/infras/common"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"sync"
	"time"
)

// EventHandler 事件处理器接口
type EventHandler interface {
	Handle(ctx context.Context, event common.Event) error
}

// EventDispatcher 事件分发器 EventBus
type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

// RegisterHandler 注册处理器
func (d *EventDispatcher) RegisterHandler(eventType string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

// Dispatch 并发分发事件
func (d *EventDispatcher) Dispatch(ctx context.Context, event common.Event) error {

	d.mu.RLock()
	defer d.mu.RUnlock()

	handlers, ok := d.handlers[event.GetType()]
	if !ok || len(handlers) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(handlers))

	for _, handler := range handlers {
		klog.Infof("开始处理事件: %s", event.GetType())

		wg.Go(func() {
			func(h EventHandler) {
				// 带超时和重试的事件处理
				if err := withRetry(func() error {
					ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
					defer cancel()
					return h.Handle(ctx, event)
				}, 3); err != nil {
					errCh <- err
				}
			}(handler)
		})
		//wg.Add(1)
		//go func(h EventHandler) {
		//	defer wg.Done()
		//	// 带超时和重试的事件处理
		//	if err := withRetry(func() error {
		//		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		//		defer cancel()
		//		return h.Handle(ctx, event)
		//	}, 3); err != nil {
		//		errCh <- err
		//	}
		//}(handler)
	}

	// 等待所有处理器完成
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
	for i := 0; i < maxRetries; i++ {
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

func initEventHandlers() *EventDispatcher {
	// 使用sync.Once确保初始化代码只执行一次
	once.Do(func() {
		dispatcherInstance = NewEventDispatcher()
		inventoryHandler := &InventoryHandler{}
		dispatcherInstance.RegisterHandler("created", inventoryHandler)
		dispatcherInstance.RegisterHandler("cancelled", inventoryHandler)
	})
	return dispatcherInstance
}
