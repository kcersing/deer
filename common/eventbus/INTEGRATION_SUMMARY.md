# EventBus + AMQP 集成总结

## 已完成的工作

### 1. 核心集成模块

**文件**: `common/eventbus/amqp_bridge.go`

**功能**:
- ✅ 创建 `AMQPBridge` 类，实现内存事件总线与 AMQP 的双向桥接
- ✅ 中间件：`AMQPPublishingMiddleware()` - 自动发布内存事件到 AMQP
- ✅ 后台监听：`StartListener()` - 消费 AMQP 消息转发到内存总线
- ✅ 批量操作：`PublishBatchToAMQP()` 和 `PublishAsyncBatchToAMQP()`
- ✅ 优雅关闭：`Stop()` 方法

### 2. 使用示例

**文件**: `app/message/biz/service/eventbus_integration_example.go`

**包含 7 个完整示例**:
1. 基础集成 - 内存事件 + AMQP 双向同步
2. 群发消息场景 - 批量发布
3. 异步群发 - 监听发布结果
4. 订阅处理 - 从 AMQP 接收并处理
5. 整合到实际服务 - 群发消息
6. 中间件链 - 多个中间件组合
7. 全局初始化 - 推荐用于生产环境

### 3. 详细文档

**文件**: `common/eventbus/INTEGRATION_GUIDE.md`

**内容**:
- 完整架构图
- 核心组件说明
- 4 种使用方式
- 数据流示例
- 中间件链执行
- 错误处理
- 性能优化建议
- 最佳实践
- 常见问题解答

### 4. 改进实现

**文件**: `app/message/biz/service/send_member_messages.go`

**改进点**:
- ✅ 集成全局事件总线和桥接器
- ✅ 初始化函数 `InitGlobalEventBus()`
- ✅ 两种发送方案：
  - **方案 A**: 通过事件总线（推荐）- 自动中间件处理、日志记录
  - **方案 B**: 直接 AMQP 发送（备用）- 轻量级，无中间件开销
- ✅ 异步发布，不阻塞业务流程
- ✅ 详细的日志和结果追踪

---

## 数据流对比

### 旧方案 vs 新方案

```
旧方案（直接 AMQP）:
Service → Publisher.PublishAsync()
          └─► RabbitMQ
          └─► 返回结果

新方案（事件总线 + AMQP）:
Service → EventBus.Publish(event)
          ├─► 内存总线分发
          │   └─► 本地订阅者处理
          │
          └─► 中间件链
              ├─► LoggingPlugin
              ├─► AMQPPublishingMiddleware
              │   └─► Publisher.PublishAsync()
              │       └─► RabbitMQ
              │           └─► 返回结果
              └─► 其他中间件
```

### 核心优势

| 特性 | 直接 AMQP | 事件总线 + AMQP |
|------|----------|----------------|
| 内存处理 | ❌ | ✅ 快速本地处理 |
| 中间件支持 | ❌ | ✅ 支持中间件链 |
| 日志记录 | 🟡 需手动 | ✅ 自动记录 |
| 消息持久化 | ✅ | ✅ 自动同步 |
| 分布式架构 | ✅ | ✅ 完整支持 |
| 监听消费 | 📝 需额外代码 | ✅ 自动反向转发 |
| 复杂度 | 较低 | 中等 |
| 可维护性 | 🟡 | ✅ 易于维护 |

---

## 关键代码段

### 1. 初始化（应用启动时调用）

```go
// main.go 或 init.go
func init() {
    if err := InitGlobalEventBus(); err != nil {
        log.Fatal(err)
    }
}
```

### 2. 发送消息（推荐方式）

```go
// 方案 A: 通过事件总线（推荐）
events := []*eventbus.Event{
    {
        Id: "msg-1",
        Topic: "member_message_send",
        Payload: map[string]interface{}{
            "recipient_id": "user_1",
            "title": "标题",
        },
    },
}

resultCh := globalBridge.PublishAsyncBatchToAMQP(ctx, events)
for result := range resultCh {
    if result.Error != nil {
        // 处理失败
    }
}
```

### 3. 订阅处理

```go
// 定义处理器
handler := eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
    fmt.Printf("Event: %v\n", event.Payload)
    return nil
})

// 注册异步订阅（3 个并发）
eb := GetEventBus()
eb.SubscribeAsync("member_message_send", handler, 3)
```

### 4. 中间件链

```go
eb := eventbus.NewEventBus()

// 依次添加中间件
eb.Use(eventbus.LoggingPlugin())              // 日志
eb.Use(bridge.AMQPPublishingMiddleware())     // AMQP 发布
eb.Use(eventbus.TransformPlugin())            // 消息转换
```

---

## 生产环境建议

### 1. 初始化位置

```go
// app/main.go
func main() {
    // 第一步：初始化全局事件总线
    if err := InitGlobalEventBus(); err != nil {
        klog.Fatal("event bus init failed", err)
    }
    
    // ... 其他初始化
    
    // 第二步：启动 RPC 服务
    server.Start()
}
```

### 2. 错误重试

```go
// 收集失败的事件
var failedEvents []*eventbus.Event

for result := range resultCh {
    if result.Error != nil {
        failedEvents = append(failedEvents, events[result.Index])
    }
}

// 异步重试（延迟 5 秒）
go func() {
    time.Sleep(5 * time.Second)
    globalBridge.PublishAsyncBatchToAMQP(ctx, failedEvents)
}()
```

### 3. 监控指标

```go
// 自定义中间件：统计性能
eb.Use(func(next eventbus.Handler) eventbus.Handler {
    return eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
        start := time.Now()
        err := next.Handle(ctx, event)
        
        // 记录性能指标
        duration := time.Since(start)
        metrics.RecordEventLatency(event.Topic, duration)
        
        return err
    })
})
```

### 4. 优雅关闭

```go
// app/main.go
func shutdownHook() {
    globalBridge.Stop()
    klog.Infof("event bus stopped")
}

// 在 HTTP/RPC 服务关闭前调用
```

---

## 与其他模块的协作

### EventBus + Middleware (mw)

```go
// 在 RPC 服务中自动发布事件
func RPCMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
    return func(ctx context.Context, req, resp interface{}) error {
        err := next(ctx, req, resp)
        
        // 调用成功，发布事件
        if err == nil {
            event := eventbus.NewEvent("rpc_completed", resp)
            GetEventBus().Publish(ctx, event)
        }
        
        return err
    }
}
```

### EventBus + Consts (常量)

```go
// common/consts/consts.go
const (
    EventMemberMessageSend = "member_message_send"
    EventOrderCreated = "order_created"
    EventUserRegistered = "user_registered"
)

// 使用
event := eventbus.NewEvent(consts.EventMemberMessageSend, payload)
```

### EventBus + RPC Client

```go
// 跨微服务通信示例
// 服务 A 发布事件
GetEventBus().Publish(ctx, eventbus.NewEvent("user_registered", userData))

// 服务 B 订阅事件（通过 RabbitMQ）
eb.SubscribeAsync("user_registered", handler, 1)
```

---

## 测试验证

### 单元测试示例

```go
func TestEventBusAMQPBridge(t *testing.T) {
    // 1. 初始化
    mq.InitMQ()
    publisher, _ := amqpclt.NewPublisher(mq.Client, "test")
    subscriber, _ := amqpclt.NewSubscribe(mq.Client, "test")
    
    eb := eventbus.NewEventBus()
    bridge := eventbus.NewAMQPBridge(eb, publisher, subscriber)
    eb.Use(bridge.AMQPPublishingMiddleware())
    
    // 2. 发布事件
    event := eventbus.NewEvent("test_topic", "test_payload")
    eb.Publish(context.Background(), event)
    
    // 3. 验证
    time.Sleep(100 * time.Millisecond)
    // 检查 AMQP 中是否收到消息
    
    bridge.Stop()
}
```

---

## 常见问题

### Q1: 为什么要用事件总线而不是直接 AMQP？

**A**: 
- 事件总线支持本地快速处理（内存）+ 远程可靠投递（AMQP）
- 支持中间件链，便于实现日志、监控、转换等功能
- 代码解耦，易于维护和测试
- 支持单机和分布式灵活切换

### Q2: 内存事件和 AMQP 消息会重复吗？

**A**: 不会，通过 `Source` 字段区分：
- `Source="service"` - 本地发送
- `Source="amqp"` - 从 RabbitMQ 消费

### Q3: 性能瓶颈在哪里？

**A**: 通常是网络 IO，优化建议：
- 批量发送（减少往返次数）
- 异步处理（非阻塞）
- 增加消费并发

### Q4: 如何处理大量事件？

**A**: 
- 使用异步批量发布
- 增加消费并发（`SubscribeAsync` 的并发参数）
- 考虑消息队列分片

---

## 后续可优化方向

1. **消息持久化存储**: 添加数据库持久化，支持消息查询和重放
2. **死信队列**: 处理失败消息的死信队列机制
3. **事件溯源**: 完整的事件版本控制和历史记录
4. **实时仪表板**: 可视化监控事件流
5. **消息去重**: 基于 ID 的幂等性处理
6. **限流熔断**: 保护下游服务

---

## 总结

通过 `AMQPBridge`，`EventBus` 与 `AMQP` 完美结合：

```
本地快速处理（内存） ↔ 分布式可靠投递（RabbitMQ）
         ↓
    中间件链（日志、监控、转换）
         ↓
    支持同步/异步、批量/单条、本地/远程
```

✨ **推荐在生产环境使用新的事件总线方案** ✨
