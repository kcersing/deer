# EventBus 架构指南 - 责任单一化

## 架构设计

重新设计了 EventBus 系统，使其职责清晰、易于维护：

```
┌─────────────────────────────────────┐
│     Event Publishing Manager        │
│    (EventPublisher)                 │
├─────────────────────────────────────┤
│  PublishLocal() ────→ 内存总线       │  ← 仅本服务处理
│  PublishDistributed() → MQ + 内存    │  ← 跨服务分布式
│  PublishToMQOnly() ──→ MQ           │  ← 其他服务处理
└─────────────────────────────────────┘

┌─────────────────────────────────────┐
│      Memory EventBus                │  内存事件总线（高速）
│  - Subscribe/SubscribeAsync         │
│  - SubscribeWithPool (消费者池)      │
│  - Middleware chain (中间件链)      │
└─────────────────────────────────────┘

┌─────────────────────────────────────┐
│       AMQP Listener                 │  单向监听MQ
│  MQ消息 → 转换 → 内存总线           │
│  (Source="amqp")                    │
└─────────────────────────────────────┘
```

---

## 三种发布模式

### 1️⃣ 本地内存事件 (PublishLocal)

**职责**：单个服务内部的异步处理

**特点**：
- 🚀 高速、低延迟
- 💾 不持久化
- 🏠 不跨服务
- 📍 内存中处理

**使用场景**：
```go
// 发送用户消息通知（本服务处理）
publisher.PublishLocal(ctx, "send_user_messages", payload)

// 订单状态变化（本服务内部处理）
publisher.PublishLocal(ctx, "order_status_updated", orderData)
```

**数据流**：
```
Publisher
    ↓
PublishLocal()
    ↓
内存总线
    ↓
订阅者 (同个进程内)
```

---

### 2️⃣ 分布式事件 (PublishDistributed)

**职责**：跨服务通信、需要持久化的事件

**特点**：
- 🔄 双向发布：MQ + 内存
- 💾 MQ持久化
- 🌐 跨服务通信
- ⚡ 本服务立即处理，其他服务异步处理

**使用场景**：
```go
// 订单创建 → 需要库存服务、支付服务处理
publisher.PublishDistributed(ctx, "order.created", orderData)

// 商品更新 → 需要其他服务同步
publisher.PublishDistributed(ctx, "product.updated", productData)

// 用户认证 → 需要日志审计
publisher.PublishDistributed(ctx, "user.authenticated", userData)
```

**数据流**：
```
Publisher
    ↓
PublishDistributed()
    ├─→ RabbitMQ (异步，不阻塞)
    │     ↓
    │   其他服务 (MQ消费者)
    │
    └─→ 内存总线 (同步)
          ↓
        本服务订阅者
```

**时序图**：
```
时间 → 

事件发布
  ↓
MQ发送 (异步)     内存处理
  │                ↓
  │             本服务立即处理
  ↓
其他服务接收 (稍后)
```

---

### 3️⃣ 仅MQ事件 (PublishToMQOnly)

**职责**：触发其他服务处理，本服务不处理

**特点**：
- 📤 只发送到MQ
- 🚫 本服务不处理
- 🌐 仅跨服务
- ⏱️ 异步处理

**使用场景**：
```go
// 通知邮件服务发送邮件
publisher.PublishToMQOnly(ctx, "notification.send_email", emailData)

// 触发数据分析服务
publisher.PublishToMQOnly(ctx, "analytics.track_event", eventData)

// 触发第三方系统同步
publisher.PublishToMQOnly(ctx, "external.sync_data", syncData)
```

**数据流**：
```
Publisher
    ↓
PublishToMQOnly()
    ↓
RabbitMQ
    ↓
其他服务
```

---

## 使用示例

### 示例1：Message 服务 - 发送用户消息

```go
// 初始化
eventBus := eventbus.NewEventBus()
publisher := eventbus.NewEventPublisher(eventBus, amqpPublisher)

// 处理器1：本地处理（如数据库记录）
eventBus.SubscribeAsync("send_user_messages", 
    func(ctx context.Context, event *eventbus.Event) error {
        // 保存消息到数据库
        return db.SaveMessage(event.Payload)
    }, 5)

// 处理器2：本地处理（如推送通知）
eventBus.SubscribeAsync("send_user_messages",
    func(ctx context.Context, event *eventbus.Event) error {
        // 发送推送通知
        return pushService.Send(event.Payload)
    }, 10)

// 发送消息 - 只在本服务处理，不需要其他服务
publisher.PublishLocal(ctx, "send_user_messages", messageData)
```

---

### 示例2：Order 服务 - 订单创建

```go
// 初始化
eventBus := eventbus.NewEventBus()
publisher := eventbus.NewEventPublisher(eventBus, amqpPublisher)

// 本服务处理：创建订单快照
eventBus.SubscribeWithPool("order.created", 
    &CreateOrderSnapshotHandler{}, 10)

// 发送订单创建事件 - 同时通知其他服务
publisher.PublishDistributed(ctx, "order.created", orderData)

// 说明：
// 1. 本服务立即处理 → 创建快照
// 2. 库存服务从MQ消费 → 扣减库存
// 3. 支付服务从MQ消费 → 发起支付
// 4. 通知服务从MQ消费 → 发送通知
```

---

### 示例3：Notification 服务 - 发送邮件

```go
// 仅从MQ消费邮件发送请求，不主动发布到MQ
amqpListener := eventbus.NewAMQPListener(eventBus, amqpSubscriber)
amqpListener.StartListener(ctx)

// 订阅MQ中的邮件事件
eventBus.SubscribeWithPool("notification.send_email",
    &SendEmailHandler{}, 5)
```

---

## 与 AMQP 的交互

### 新架构的变化

| 组件 | 旧设计 | 新设计 |
|------|-------|-------|
| **AMQPBridge** | 双向拦截 + 监听 | 单向监听 (AMQPListener) |
| **发布方式** | 统一 Publish() | 三种模式：Local/Distributed/MQOnly |
| **中间件** | AMQPPublishingMiddleware | ❌ 移除（改用显式发布） |
| **死循环风险** | ⚠️ 高 | ✅ 完全消除 |
| **职责** | 不清晰 | 清晰单一 |

### 消除死循环的方式

**旧方式**（已移除）：
```go
// ❌ 自动拦截所有事件，导致循环
eventBus.Use(bridge.AMQPPublishingMiddleware())
eventBus.Publish(event)  // 自动发送到MQ
  ↓ 从MQ接收
  ↓ 再次发送到MQ（死循环）
```

**新方式**（明确控制）：
```go
// ✅ 用户明确选择发布模式
publisher.PublishLocal(event)  // 仅内存
publisher.PublishDistributed(event)  // MQ + 内存
publisher.PublishToMQOnly(event)  // 仅MQ
```

---

## 迁移指南

### 对现有代码的影响

**移除的 API**：
```go
// ❌ 不再支持
bridge.AMQPPublishingMiddleware()  // 移除了
bridge.PublishBatchToAMQP()        // 移除了
bridge.PublishAsyncBatchToAMQP()   // 移除了
```

**新 API**：
```go
// ✅ 使用新的发布管理器
publisher := NewEventPublisher(eventBus, amqpPub)
publisher.PublishLocal(ctx, topic, payload)
publisher.PublishDistributed(ctx, topic, payload)
publisher.PublishToMQOnly(ctx, topic, payload)
```

### 迁移步骤

1. **替换监听器初始化**：
   ```go
   // 旧
   bridge := NewAMQPBridge(eventBus, publisher, subscriber)
   bridge.StartListener(ctx)
   
   // 新
   listener := NewAMQPListener(eventBus, subscriber)
   listener.StartListener(ctx)
   ```

2. **删除中间件注册**：
   ```go
   // ❌ 移除这行
   eventBus.Use(bridge.AMQPPublishingMiddleware())
   ```

3. **用发布管理器替换发布调用**：
   ```go
   // 旧
   eventBus.Publish(ctx, event)
   
   // 新 - 根据需要选择
   publisher.PublishLocal(ctx, topic, payload)      // 本服务处理
   publisher.PublishDistributed(ctx, topic, payload) // 跨服务
   publisher.PublishToMQOnly(ctx, topic, payload)    // 其他服务处理
   ```

---

## 最佳实践

### ✅ DO - 正确做法

```go
// 1. 分清场景
if needCrossService {
    publisher.PublishDistributed(ctx, topic, data)  // 跨服务
} else {
    publisher.PublishLocal(ctx, topic, data)        // 本服务
}

// 2. 使用消费者池处理高吞吐事件
eventBus.SubscribeWithPool("high_volume_topic", 
    handler, 20)  // 20个worker

// 3. 使用Registry集中管理消费者
registry := NewConsumerRegistry()
registry.RegisterHandler("email_handler", emailHandler)
registry.RegisterConsumer("send_email", "email_handler", 5)
registry.StartAll(eventBus)
```

### ❌ DON'T - 错误做法

```go
// 1. 不要混淆发布模式
publisher.PublishDistributed(ctx, "internal_only", data)  // ❌ 不需要跨服务却发MQ

// 2. 不要添加拦截所有事件的中间件
eventBus.Use(someUniversalMiddleware)  // ❌ 可能导致各种问题

// 3. 不要在消费者中再次发布相同事件
func handler(ctx context.Context, event *Event) {
    // ❌ 不要这样做
    publisher.PublishDistributed(ctx, event.Topic, event.Payload)
}
```

---

## 故障排查

### Q: 为什么某个事件没有被处理？

**A**: 检查发布模式：
```go
// 本服务订阅者看不到？
// → 检查是否用了 PublishToMQOnly()（只发到MQ）
// → 改用 PublishLocal() 或 PublishDistributed()

// 其他服务看不到？
// → 检查是否用了 PublishLocal()（只在内存）
// → 改用 PublishDistributed() 或 PublishToMQOnly()
```

### Q: 如何确认事件是否在循环？

**A**: 检查日志中的 `eventId`：
```log
[AMQPListener] event forwarded from MQ to memory bus, 
  topic=order.created, eventId=uuid-xxx
[HandleOrderSnapshot] processing, eventId=uuid-xxx
```

相同 `eventId` 重复出现 = 事件循环 ⚠️

---

## 性能对比

| 操作 | 内存 | 内存+MQ | 仅MQ |
|------|------|--------|------|
| 发布延迟 | <1ms | <2ms | <2ms |
| 消费延迟 | <1ms | <1ms | 100-500ms |
| 持久化 | ❌ | ✅ | ✅ |
| 跨服务 | ❌ | ✅ | ✅ |
| 单服务 | ✅ | ✅ | ❌ |

---

## 总结

新的 EventBus 架构提供了：

1. **清晰的职责分离** - 发布、内存处理、MQ监听各司其职
2. **三种明确的发布模式** - 根据需求选择
3. **完全消除死循环** - 不再自动拦截
4. **易于维护和测试** - 行为可预测
5. **向后兼容** - NewAMQPBridge 仍然可用
