# EventBus 架构重构总结

## 🎯 核心改进

### 旧架构问题
```
❌ 双向拦截中间件 → 自动循环
❌ 职责不清晰 → 难以维护
❌ 死循环风险 → 生产事故
❌ 开发体验差 → 容易出错
```

### 新架构优势
```
✅ 单一职责原则 → 清晰的数据流
✅ 显式发布选择 → 避免循环
✅ 三种发布模式 → 满足不同场景
✅ 完全向后兼容 → 平滑迁移
```

---

## 📊 架构对比

### 旧设计（有问题）
```
┌─────────────────┐
│ EventBus.Publish│
│   (统一入口)     │
└────────┬────────┘
         │
         ├─→ 中间件链处理
         │   └─→ AMQPPublishingMiddleware ⚠️ (自动拦截)
         │       └─→ 发送到MQ
         │       └─→ 发送到内存
         │
         └─→ 内存订阅者处理

问题：从MQ接收回的事件也会被中间件拦截，再次发送到MQ
结果：事件无限循环！
```

### 新设计（解决方案）
```
┌────────────────────────────────┐
│  EventPublisher (发布管理器)    │
└────┬───────────┬───────────┬────┘
     │           │           │
PublishLocal   PublishDistributed   PublishToMQOnly
   │               │                    │
   │          ┌────┴────┐              │
   │          │          │              │
   ▼         ▼          ▼              ▼
内存总线    MQ+内存    MQ仅
 │          │          │
 ├─→本服务  ├─→本服务  └─→其他服务
 │          │
 │          └─→其他服务
 │
从MQ接收（AMQPListener）
 │
 ├─→标记Source="amqp"
 │
 └─→转发到内存总线
    (处理完，不会再发回MQ)
    
好处：清晰的单向流动，不会循环
```

---

## 🔄 数据流对比

### 场景1：本地事件处理

**旧方式**：
```
用户发送消息
  ├─→ eventBus.Publish()
  ├─→ 中间件拦截
  ├─→ 检查是否是amqp来源
  ├─→ 决定是否发到MQ
  └─→ 转发到内存

每次 Publish 都要经过判断逻辑 ⚠️
```

**新方式**：
```
用户发送消息
  └─→ publisher.PublishLocal()
      └─→ 直接发送到内存
          （快速、简洁）
```

### 场景2：跨服务事件

**旧方式**：
```
订单创建
  └─→ eventBus.Publish()
      ├─→ 检查中间件
      ├─→ 发送到MQ (异步)
      └─→ 发送到内存

本服务处理
  └─→ 从MQ接收（起另一个监听器）
      ├─→ 再次Publish()到内存
      └─→ 中间件再次拦截 ⚠️ 死循环!
```

**新方式**：
```
订单创建
  └─→ publisher.PublishDistributed()
      ├─→ 异步发送到MQ
      └─→ 立即发送到内存 ✅

AMQPListener 从MQ接收
  └─→ 标记 Source="amqp"
      └─→ 转发到内存
          （不会触发MQ发送逻辑）
```

---

## 📦 新增组件

### 1. EventPublisher (publisher.go)
```go
type EventPublisher struct {
    memoryBus *EventBus        // 内存总线
    amqpPub   *amqpclt.Publish // MQ发布（可选）
}

// 三种发布方法
PublishLocal()         // 仅内存
PublishDistributed()   // 内存+MQ
PublishToMQOnly()      // 仅MQ
```

### 2. AMQPListener (amqp_bridge.go 重构)
```go
type AMQPListener struct {
    eventBus   *EventBus          // 内存总线
    subscriber *amqpclt.Subscribe // MQ订阅
}

// 职责
StartListener()  // 从MQ监听 → 转发到内存
```

---

## 🎓 使用指南

### 1️⃣ 本地异步处理
```go
pub.PublishLocal(ctx, "send_message", data)

场景：
- Message服务发送用户消息
- Order服务内部状态更新
- Cache更新

特点：
- 高速 (<1ms)
- 不持久化
- 不跨服务
```

### 2️⃣ 跨服务通信
```go
pub.PublishDistributed(ctx, "order.created", data)

场景：
- 订单创建通知库存、支付、通知服务
- 商品更新通知其他服务
- 用户认证通知审计服务

特点：
- 本服务立即处理
- MQ异步发送给其他服务
- 具有持久化
```

### 3️⃣ 触发外部服务
```go
pub.PublishToMQOnly(ctx, "send_email", data)

场景：
- 触发邮件服务
- 触发第三方系统同步
- 发起后台任务

特点：
- 本服务不处理
- 仅通过MQ转发
- 其他服务异步消费
```

---

## ✅ 迁移清单

- [ ] 移除 `eventBus.Use(bridge.AMQPPublishingMiddleware())`
- [ ] 创建 `EventPublisher` 实例
- [ ] 将 `bridge` 替换为 `AMQPListener`
- [ ] 审查所有 `eventBus.Publish()` 调用，改为：
  - [ ] `pub.PublishLocal()` - 本服务处理
  - [ ] `pub.PublishDistributed()` - 跨服务
  - [ ] `pub.PublishToMQOnly()` - 外部触发
- [ ] 测试事件流，确认无循环
- [ ] 验证日志中 `eventId` 不重复

---

## 🔍 验证清单

| 检查项 | 旧架构 | 新架构 |
|--------|--------|--------|
| 死循环风险 | ⚠️ 高 | ✅ 无 |
| 职责清晰 | ❌ 模糊 | ✅ 清晰 |
| 易于理解 | ❌ 复杂 | ✅ 简洁 |
| 性能 | ⚠️ 需判断 | ✅ 高效 |
| 可维护性 | ❌ 差 | ✅ 好 |
| 向后兼容 | - | ✅ 有 |

---

## 📝 文件清单

新增/修改的文件：

- ✨ `publisher.go` - **新增** EventPublisher（发布管理器）
- 📝 `amqp_bridge.go` - **重构** 改为单向监听
- 📝 `event.go` - **修改** 新建事件时设置Source="service"
- 📄 `ARCHITECTURE.md` - **新增** 详细的架构文档
- 💡 `examples.go` - **新增** 实际使用示例

---

## 🚀 下一步

1. 在 Message 服务中改用 `publisher.PublishLocal()`
2. 在 Order 服务中改用 `publisher.PublishDistributed()`
3. 测试并验证事件不再循环
4. 更新其他服务的事件处理代码
5. 监控日志，确认一切正常

---

## 📚 参考

- 详细文档：[ARCHITECTURE.md](ARCHITECTURE.md)
- 使用示例：[examples.go](examples.go)
- 核心代码：
  - [publisher.go](publisher.go) - 发布管理
  - [amqp_bridge.go](amqp_bridge.go) - MQ监听
  - [eventbus.go](eventbus.go) - 内存总线
