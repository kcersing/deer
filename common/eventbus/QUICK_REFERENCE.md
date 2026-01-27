# EventBus 快速参考

## 三种发布方式对比

```
┌─────────────────────┬──────────────┬──────────────┬──────────────┐
│ 发布方式             │ PublishLocal │ Distributed  │ PublishToMQ  │
├─────────────────────┼──────────────┼──────────────┼──────────────┤
│ 发送到内存总线       │      ✅      │      ✅      │      ❌      │
│ 发送到RabbitMQ      │      ❌      │      ✅      │      ✅      │
│ 本服务处理           │      ✅      │      ✅      │      ❌      │
│ 其他服务处理        │      ❌      │      ✅      │      ✅      │
│ 延迟                 │     <1ms     │     <2ms     │     <2ms     │
│ 持久化               │      ❌      │      ✅      │      ✅      │
│ 跨服务               │      ❌      │      ✅      │      ✅      │
└─────────────────────┴──────────────┴──────────────┴──────────────┘
```

## 快速选择

```
    需要其他服务处理？
    │
    ├─ 是 ─→ 本服务也要处理？
    │        │
    │        ├─ 是 ─→ PublishDistributed() ⭐
    │        │
    │        └─ 否 ─→ PublishToMQOnly()
    │
    └─ 否 ─→ PublishLocal()
```

## 代码示例

### PublishLocal - 本地内存
```go
pub := NewEventPublisher(eventBus, nil)
pub.PublishLocal(ctx, "send_user_messages", payload)
```

### PublishDistributed - 同时发送
```go
pub := NewEventPublisher(eventBus, amqpPublisher)
pub.PublishDistributed(ctx, "order.created", orderData)
```

### PublishToMQOnly - 仅MQ
```go
pub := NewEventPublisher(eventBus, amqpPublisher)
pub.PublishToMQOnly(ctx, "notification.send_email", emailData)
```

## 事件处理

### 订阅方式

```go
// 方式1：简单订阅（同步）
ch := eventBus.Subscribe("topic_name")
for event := range ch {
    process(event)
}

// 方式2：异步处理（回调）
eventBus.SubscribeAsync("topic_name", handler, concurrency)

// 方式3：消费者池（高吞吐）
eventBus.SubscribeWithPool("topic_name", handler, workerNum)

// 方式4：消费者注册表（集中管理）
registry := NewConsumerRegistry()
registry.RegisterHandler("name", handler)
registry.RegisterConsumer("topic", "name", workerNum)
registry.StartAll(eventBus)
```

## 监听MQ事件

```go
// 创建监听器
listener := NewAMQPListener(eventBus, amqpSubscriber)

// 启动监听
listener.StartListener(ctx)

// 订阅来自MQ的事件（会自动标记Source="amqp"）
eventBus.SubscribeWithPool("topic_name", handler, workers)

// 优雅关闭
listener.Stop()
```

## 事件结构

```go
type Event struct {
    Id        string      // 唯一ID (UUID)
    Topic     string      // 事件主题
    Payload   any         // 事件数据
    Source    string      // "service" | "amqp" | "local"
    Version   int64       // 版本号
    Timestamp time.Time   // 时间戳
    Priority  int64       // 优先级
}

// 创建事件（Source自动设置为"service"）
event := NewEvent("topic", payload)
```

## Source 标记

| Source | 含义 | 说明 |
|--------|------|------|
| `service` | 本服务发布 | PublishDistributed() |
| `amqp` | 来自MQ | AMQPListener转发 |
| `local` | 本地内存 | PublishLocal() |

## 中间件

```go
// 添加中间件
eventBus.Use(LoggingPlugin())      // 日志记录
eventBus.Use(RecoverPlugin())      // Panic恢复
eventBus.Use(FilterPlugin("spam")) // 事件过滤
eventBus.Use(TransformPlugin())    // 事件转换

// 自定义中间件
func CustomPlugin() Middleware {
    return func(next Handler) Handler {
        return EventHandlerFunc(func(ctx context.Context, event *Event) error {
            // 前置处理
            err := next.Handle(ctx, event)
            // 后置处理
            return err
        })
    }
}
```

## 常见错误

```go
// ❌ 错误：混淆发布方式
pub.PublishDistributed(ctx, "internal_only", data)  // 不需要MQ还发

// ❌ 错误：在消费者中重新发布
func handler(ctx context.Context, event *Event) {
    pub.PublishDistributed(ctx, event.Topic, event.Payload)
}

// ❌ 错误：使用旧的中间件方式
eventBus.Use(bridge.AMQPPublishingMiddleware())  // 已移除

// ✅ 正确：明确选择发布方式
pub.PublishLocal(ctx, topic, data)        // 本服务
pub.PublishDistributed(ctx, topic, data)  // 跨服务
pub.PublishToMQOnly(ctx, topic, data)     // 其他服务
```

## 性能建议

| 场景 | 建议 | 原因 |
|------|------|------|
| 高吞吐 | 使用消费者池 | 固定线程数，CPU友好 |
| 大批量 | 使用PublishLocal | 避免不必要的MQ |
| 低延迟 | 避免PublishToMQOnly | MQ延迟100-500ms |
| 跨服务 | 使用PublishDistributed | 本服务+MQ同时处理 |

## 故障排查

| 问题 | 检查项 |
|------|--------|
| 事件未处理 | 检查Source和发布方式是否匹配 |
| 事件重复 | 检查日志中eventId是否相同（表示循环） |
| MQ消费失败 | 检查AMQPListener是否启动 |
| 高CPU占用 | 检查消费者池workerNum是否过多 |
| 内存泄漏 | 检查是否正确调用Stop()释放资源 |

## 逐步迁移

```go
// 步骤1：初始化发布管理器
pub := NewEventPublisher(eventBus, amqpPub)

// 步骤2：移除旧的中间件
// ❌ 删除这行
// eventBus.Use(bridge.AMQPPublishingMiddleware())

// 步骤3：替换所有 eventBus.Publish() 调用
// 旧
eventBus.Publish(ctx, event)

// 新 - 根据需要选择
pub.PublishLocal(ctx, event.Topic, event.Payload)
pub.PublishDistributed(ctx, event.Topic, event.Payload)
pub.PublishToMQOnly(ctx, event.Topic, event.Payload)

// 步骤4：使用AMQPListener替换AMQPBridge
// 旧
bridge := NewAMQPBridge(eventBus, pub, sub)
bridge.StartListener(ctx)

// 新
listener := NewAMQPListener(eventBus, sub)
listener.StartListener(ctx)
```

## 配置参考

```go
// 消费者池大小建议
CPU核心数=4   → workers=10-20
CPU核心数=8   → workers=20-50
CPU核心数=16  → workers=50-100

// 队列大小建议
轻负载 → 100-500
中负载 → 500-2000
重负载 → 2000+
```

## 常用代码片段

### 发布并订阅
```go
eventBus := NewEventBus()
pub := NewEventPublisher(eventBus, nil)

// 订阅
eventBus.SubscribeWithPool("my_event", myHandler, 10)

// 发布
pub.PublishLocal(ctx, "my_event", data)
```

### 跨服务通信
```go
// 发送端
pub.PublishDistributed(ctx, "user.created", userData)

// 接收端（其他服务）
listener := NewAMQPListener(eventBus, sub)
listener.StartListener(ctx)
eventBus.SubscribeAsync("user.created", handler, 5)
```

### 错误处理
```go
type MyHandler struct {
    logger *log.Logger
}

func (h *MyHandler) Handle(ctx context.Context, event *Event) error {
    defer func() {
        if r := recover(); r != nil {
            h.logger.Errorf("Panic in event handler: %v", r)
        }
    }()
    
    // 处理事件
    return nil
}
```

## 最佳实践

```
✅ DO:
- 根据场景选择发布方式
- 使用消费者池处理高吞吐事件
- 添加适当的中间件（日志、recovery）
- 在服务关闭时优雅停止监听
- 在日志中记录eventId便于追踪

❌ DON'T:
- 不要混淆三种发布方式
- 不要添加全局拦截中间件
- 不要在消费者中重新发布同一事件
- 不要忽视错误处理
- 不要忘记释放资源（Stop、close等）
```
