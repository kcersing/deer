package eventbus

import (
	"context"
	"fmt"
	"sync"
)

// ConsumerPool 消费者池 - 用于高吞吐场景
type ConsumerPool struct {
	name      string             // 消费者池的名字（用于日志/识别）
	handler   Handler            // 事件处理器（处理每个事件）
	workerNum int32              // 工作线程数量（并发度）
	queue     chan *Event        // 事件队列（缓冲通道）
	wg        sync.WaitGroup     // 等待组（确保优雅关闭）
	ctx       context.Context    // 上下文（用于控制）
	cancel    context.CancelFunc // 取消函数（停止所有 worker）
}

// NewConsumerPool 创建消费者池
func NewConsumerPool(name string, handler Handler, workerNum int32) *ConsumerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &ConsumerPool{
		name:      name,
		handler:   handler,
		workerNum: workerNum,
		queue:     make(chan *Event, DefaultConfig.QueueSize),
		ctx:       ctx,
		cancel:    cancel,
	}
}

// Start 启动消费者池
func (cp *ConsumerPool) Start() {
	for i := int32(0); i < cp.workerNum; i++ {
		cp.wg.Add(1)
		go cp.worker()
	}
}

func (cp *ConsumerPool) worker() {
	defer cp.wg.Done()
	for {
		select {
		case <-cp.ctx.Done():
			return
		case event := <-cp.queue:
			if event == nil {
				return
			}
			// 每个 handler 调用使用可选超时，并捕获 panic
			func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("[Pool Recover] pool=%s panic: %v\n", cp.name, r)
					}
				}()

				handlerCtx := cp.ctx
				if DefaultConfig.HandlerTimeout > 0 {
					var cancel context.CancelFunc
					handlerCtx, cancel = context.WithTimeout(cp.ctx, DefaultConfig.HandlerTimeout)
					defer cancel()
				}
				_ = cp.handler.Handle(handlerCtx, event)
			}()
		}
	}
}

// Consume 消费事件
func (cp *ConsumerPool) Consume(event *Event) {
	select {
	case cp.queue <- event:
	default:
		fmt.Printf("警告: 消费者池 %s 队列已满，丢弃事件\n", cp.name)
	}
}

// Stop 停止消费者池
func (cp *ConsumerPool) Stop() {
	cp.cancel()
	cp.wg.Wait()
	close(cp.queue)
}
