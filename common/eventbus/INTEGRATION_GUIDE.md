# EventBus + AMQP 集成指南

## 概述

`common/amqpclt` 与 `common/eventbus` 的集成，通过 `AMQPBridge` 实现**内存事件总线**与**消息队列**的双向同步，支持以下场景：

- ✅ 内存事件自动同步到 RabbitMQ（持久化）
- ✅ RabbitMQ 消息自动转发到内存事件总线
- ✅ 批量发布和异步发布
- ✅ 中间件链式处理
- ✅ 分布式系统间的事件协作

---

## 架构图

```
┌─────────────────────────────────────────────────────────┐
│                    业务服务层                              │
│  (用户服务、订单服务、消息服务等)                          │
└────────┬──────────────────────────┬────────────────────┘
         │ Publish Event            │ Handle Event
         ▼                          ▲
┌────────────────────────────────────────────────────────┐
│           内存事件总线 (EventBus)                        │
│  • 本地快速处理                                          │
│  • 支持中间件链                                          │
│  • 支持同步/异步订阅                                      │
└────┬────────────────────────────────────┬──────────────┘
     │ 中间件拦截：                       │ 后台监听消费：
     │ - 日志记录                         │ - 反序列化
     │ - 消息过滤                         │ - 事件转换
     │ - 发布到 AMQP                      │ - 转发到总线
     ▼                                    ▲
┌────────────────────────────────────────────────────────┐
│         AMQPBridge（桥接层）                             │
│  • 事件 <-> AMQP 消息转换                                │
│  • 批量操作                                              │
│  • 错误处理与重试                                        │
└────┬────────────────────────────────────┬──────────────┘
     │ Publish                            │ Subscribe
     ▼                                    ▲
┌────────────────────────────────────────────────────────┐
│        RabbitMQ（分布式消息队列）                        │
│  Exchange: events (fanout 广播)                        │
│  Queue: 动态创建（每个消费者一个）                       │
│  特性：持久化、可靠投递、消息确认                        │
└────────────────────────────────────────────────────────┘
         │
         ├─► 其他微服务消费
         ├─► 外部系统集成
         └─► 审计日志记录
```

---

## 核心组件

### 1. **Event（事件）**

```go
type Event struct {
    Id        string    // 唯一标识
    Topic     string    // 事件主题
    Payload   any       // 事件负载
    Source    string    // 事件来源
    Version   int64     // 版本号
    Timestamp time.Time // 时间戳
    Priority  int64     // 优先级
}
```

### 2. **AMQPBridge（桥接器）**

```go
type AMQPBridge struct {
    eventBus    *EventBus
    publisher   *amqpclt.Publish
    subscriber  *amqpclt.Subscribe
    ctx         context.Context
    cancel      context.CancelFunc
    done        chan struct{}
}
```

**核心方法**：
- `NewAMQPBridge()` - 创建桥接器
- `AMQPPublishingMiddleware()` - 中间件：发布到 AMQP
- `StartListener()` - 启动后台监听
- `PublishBatchToAMQP()` - 批量发布
- `PublishAsyncBatchToAMQP()` - 异步批量发布
- `Stop()` - 优雅关闭

---

## 使用方式

### 方式 1：基础集成（推荐用于初学）

```go
// 1. 初始化连接
mq.InitMQ()
publisher, _ := amqpclt.NewPublisher(mq.Client, "events")
subscriber, _ := amqpclt.NewSubscribe(mq.Client, "events")

// 2. 创建事件总线和桥接器
eb := eventbus.NewEventBus()
bridge := eventbus.NewAMQPBridge(eb, publisher, subscriber)

// 3. 添加中间件：拦截内存事件，同时发布到 AMQP
eb.Use(bridge.AMQPPublishingMiddleware())

// 4. 启动后台监听：从 AMQP 消费消息，转发到内存总线
ctx := context.Background()
bridge.StartListener(ctx)

// 5. 订阅内存事件
ch := eb.Subscribe("user_registered")
go func() {
    for event := range ch {
        fmt.Printf("Event: %v\n", event)
    }
}()

// 6. 发布事件（自动同步到 AMQP）
event := eventbus.NewEvent("user_registered", map[string]interface{}{
    "user_id": 123,
    "name": "John",
})
eb.Publish(ctx, event)

// 7. 优雅关闭
defer bridge.Stop()
```

### 方式 2：群发消息（批量场景）

```go
// 创建多个事件
recipients := []string{"user_1", "user_2", "user_3"}
events := make([]*eventbus.Event, len(recipients))

for i, recipientID := range recipients {
    events[i] = &eventbus.Event{
        Id:    fmt.Sprintf("msg-%d", i),
        Topic: "member_message",
        Payload: map[string]interface{}{
            "recipient_id": recipientID,
            "title":        "群发消息",
            "content":      "这是一条群发内容",
        },
        Timestamp: time.Now(),
    }
}

// 异步批量发布
resultCh := bridge.PublishAsyncBatchToAMQP(ctx, events)

// 收集结果
for result := range resultCh {
    if result.Error != nil {
        fmt.Printf("Message %d failed: %v\n", result.Index, result.Error)
    }
}
```

### 方式 3：异步订阅处理

```go
// 定义处理器
handler := eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
    fmt.Printf("[Handler] Processing: %v\n", event.Payload)
    return nil
})

// 异步订阅（并发数为 3）
eb.SubscribeAsync("order_created", handler, 3)

// 发布事件
event := eventbus.NewEvent("order_created", map[string]interface{}{
    "order_id": 12345,
    "total": 299.99,
})
eb.Publish(ctx, event)
```

### 方式 4：全局初始化（推荐用于生产）

```go
// app/main.go

// 全局变量
var (
    globalEventBus *eventbus.EventBus
    globalBridge   *eventbus.AMQPBridge
)

// 初始化函数
func init() {
    mq.InitMQ()
    publisher, _ := amqpclt.NewPublisher(mq.Client, "events")
    subscriber, _ := amqpclt.NewSubscribe(mq.Client, "events")
    
    globalEventBus = eventbus.NewEventBus()
    globalBridge = eventbus.NewAMQPBridge(globalEventBus, publisher, subscriber)
    
    // 注册中间件
    globalEventBus.Use(eventbus.LoggingPlugin())
    globalEventBus.Use(globalBridge.AMQPPublishingMiddleware())
    
    // 启动监听
    ctx := context.Background()
    globalBridge.StartListener(ctx)
}

// 获取全局事件总线
func GetEventBus() *eventbus.EventBus {
    return globalEventBus
}

// 在服务中使用
func (s *UserService) RegisterUser(user *User) error {
    // 业务逻辑...
    
    // 发布事件
    event := eventbus.NewEvent("user_registered", user)
    GetEventBus().Publish(context.Background(), event)
    
    return nil
}
```

---

## 数据流示例

### 场景：群发消息

```
1. 服务调用
   SendMemberMessagesService.Run()
   │
   ├─► 获取接收者列表：[user_1, user_2, user_3]
   │
   ├─► 创建事件：
   │   Event{Topic: "member_message", Payload: {recipient: user_1, ...}}
   │   Event{Topic: "member_message", Payload: {recipient: user_2, ...}}
   │   Event{Topic: "member_message", Payload: {recipient: user_3, ...}}
   │
   └─► 批量异步发布：bridge.PublishAsyncBatchToAMQP(events)
       │
       ├─► 转换为 AMQP BatchMessage
       │   BatchMessage{RoutingKey: "member_message", Payload: {...}}
       │
       └─► publisher.PublishAsync()
           │
           ├─► 逐条发送到 RabbitMQ Exchange "events"
           │
           └─► 返回结果通道
               ├─► result 1: {Index: 0, Error: nil}
               ├─► result 2: {Index: 1, Error: nil}
               └─► result 3: {Index: 2, Error: nil}

2. 消息在 RabbitMQ 中
   Exchange: events (fanout)
   │
   ├─► Queue 1 (本服务订阅)
   ├─► Queue 2 (其他微服务订阅)
   └─► Queue 3 (外部系统订阅)

3. 消费端处理
   bridge.StartListener()
   │
   └─► subscriber.Subscribe()
       │
       ├─► 从 Queue 1 消费消息
       │
       ├─► 反序列化为 Message
       │   Message{Event: "member_message", Payload: {...}}
       │
       ├─► 转换为 Event
       │   Event{Topic: "member_message", Payload: {...}}
       │
       └─► eventBus.Publish(event)
           │
           └─► 转发到内存总线
               │
               └─► 本地订阅者处理
                   fmt.Printf("Event received: %v", event)
```

---

## 中间件链执行顺序

```go
eb.Use(eventbus.LoggingPlugin())              // 1. 先执行
eb.Use(bridge.AMQPPublishingMiddleware())     // 2. 再执行
eb.Use(eventbus.TransformPlugin())            // 3. 最后执行
eb.Publish(ctx, event)

// 执行链：
Publish()
  └─► TransformPlugin
      └─► AMQPPublishingMiddleware
          └─► LoggingPlugin
              └─► dispatch (内存分发)
                  └─► 内存订阅者处理

// 同时：
AMQPPublishingMiddleware (go func)
  └─► publisher.Publish()
      └─► RabbitMQ Exchange
```

---

## 关键概念

### 1. **单向同步 vs 双向同步**

```go
// 单向：仅内存 → AMQP
eb.Use(bridge.AMQPPublishingMiddleware())

// 双向：内存 ↔ AMQP
eb.Use(bridge.AMQPPublishingMiddleware())
bridge.StartListener(ctx)
```

### 2. **事件来源标记**

区分事件来源，避免重复处理：

```go
// 内存发送
event.Source = "service"

// AMQP 消费
event.Source = "amqp"
```

### 3. **事件ID追踪**

```go
event.Id = fmt.Sprintf("msg-%d-%s", time.Now().Unix(), recipientID)
// 通过 CorrelationId 追踪整个处理链
```

---

## 错误处理

### 发布失败

```go
resultCh := bridge.PublishAsyncBatchToAMQP(ctx, events)

for result := range resultCh {
    if result.Error != nil {
        // 可实现重试机制
        retryQueue.Enqueue(events[result.Index])
        klog.Errorf("publish failed, will retry: %v", result.Error)
    }
}
```

### 消费失败

```go
// 在 AMQP 层面
// 反序列化失败 → msg.Nack(false, false) （不重新入队）
// 处理失败 → msg.Nack(false, true) （重新入队）
```

---

## 性能优化建议

### 1. 批量操作

```go
// ✅ 推荐：批量发布 1000 条
events := make([]*eventbus.Event, 1000)
bridge.PublishAsyncBatchToAMQP(ctx, events)

// ❌ 避免：逐条发布 1000 次
for _, event := range events {
    eb.Publish(ctx, event)
}
```

### 2. 异步处理

```go
// ✅ 推荐：异步发布，不阻塞业务
resultCh := bridge.PublishAsyncBatchToAMQP(ctx, events)
go func() {
    for result := range resultCh {
        // 异步处理结果
    }
}()

// ❌ 避免：同步等待所有结果
for result := range resultCh {
    // 同步等待（阻塞业务逻辑）
}
```

### 3. 通道缓冲

```go
// ✅ 推荐：设置合理的缓冲大小
resultCh := make(chan amqpclt.PublishResult, len(batchMessages))
publisher.PublishAsync(ctx, batchMessages, resultCh)

// ❌ 避免：无缓冲通道
resultCh := make(chan amqpclt.PublishResult)
```

### 4. 并发消费

```go
// ✅ 推荐：多个 goroutine 并发处理
eb.SubscribeAsync("topic", handler, 10)  // 10 个并发处理

// ❌ 避免：单个 goroutine 处理
ch := eb.Subscribe("topic")
go func() {
    for event := range ch {
        handler(event)  // 顺序处理，效率低
    }
}()
```

---

## 最佳实践

### 1. 统一事件命名

```go
// ✅ 推荐：动宾结构
"user_registered"
"order_created"
"payment_completed"

// ❌ 避免：过于复杂
"user_on_registered_event"
"order_has_been_created"
```

### 2. 事件载体设计

```go
// ✅ 推荐：包含必要信息，避免冗余
Payload: map[string]interface{}{
    "user_id": 123,
    "email": "user@example.com",
}

// ❌ 避免：包含整个对象
Payload: *User{  // 太大，序列化效率低
    // ... 100+ 字段
}
```

### 3. 错误恢复

```go
// ✅ 推荐：记录失败的消息，异步重试
failedMessages := []amqpclt.BatchMessage{}
for result := range resultCh {
    if result.Error != nil {
        failedMessages = append(failedMessages, messages[result.Index])
    }
}

// 定时重试（如使用定时任务）
go func() {
    time.Sleep(5 * time.Second)
    bridge.PublishBatchToAMQP(ctx, failedMessages)
}()
```

### 4. 监控和日志

```go
eb.Use(eventbus.LoggingPlugin())  // 所有事件都有日志

// 自定义中间件：统计发布指标
eb.Use(func(next eventbus.Handler) eventbus.Handler {
    return eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
        defer func(start time.Time) {
            metrics.RecordEventProcessTime(event.Topic, time.Since(start))
        }(time.Now())
        return next.Handle(ctx, event)
    })
})
```

---

## 常见问题

### Q1: 内存事件和 AMQP 消息重复吗？

**A**: 不会。通过 `Source` 字段区分：
- `Source="service"` - 本地服务发送
- `Source="amqp"` - 从 RabbitMQ 消费

### Q2: 如何处理顺序性要求？

**A**: 使用路由键确保同一用户的消息进入同一队列：
```go
RoutingKey: userID  // 相同 userID 的消息按顺序处理
```

### Q3: 如何监控发布成功率？

**A**: 使用异步结果通道统计：
```go
successCount := 0
for result := range resultCh {
    if result.Error == nil {
        successCount++
    }
}
rate := float64(successCount) / float64(totalMessages) * 100
fmt.Printf("Success rate: %.2f%%\n", rate)
```

### Q4: 内存总线和消息队列如何选择？

**A**: 
- **内存总线**：本地微服务内部事件，快速处理
- **消息队列**：跨微服务通信，需要持久化

### Q5: 性能瓶颈在哪里？

**A**: 通常是 RabbitMQ 网络 IO，优化建议：
- 批量发送（减少网络往返）
- 异步处理（不阻塞业务线程）
- 增加消费并发（提升吞吐量）

---

## 总结

| 特性 | 内存事件总线 | AMQP 消息队列 |
|------|----------|-------------|
| 速度 | 🚀 极快 | ⚡ 较快 |
| 持久化 | ❌ 否 | ✅ 是 |
| 分布式 | ❌ 否 | ✅ 是 |
| 使用场景 | 本地快速事件 | 跨系统可靠通信 |

**通过 AMQPBridge，两者完美结合** ✨
