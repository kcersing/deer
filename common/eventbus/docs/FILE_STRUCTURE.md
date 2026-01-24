# 文件组织说明

## 📁 EventBus 模块文件清单

### 核心文件（5 个）

| 文件 | 行数 | 功能 | 说明 |
|------|------|------|------|
| **event.go** | ~40 | 事件定义 | Event 结构体、NewEvent() |
| **handler.go** | ~40 | 处理器接口 | Handler 接口、EventHandlerFunc、TypedHandler |
| **eventbus.go** | ~163 | 事件总线核心 | 发布、订阅、中间件链、SubscribeAsync |
| **middleware.go** | ~60 | 中间件机制 | 日志、过滤、转换插件 |
| **consumer_pool.go** | ~71 | 消费者池 | 并发处理、工作线程池 |
| **registry.go** | ~123 | 消费者注册表 | 集中管理、统一启动 |

**总计：~497 行 核心代码**

### AMQP 集成（3 个）

| 文件 | 功能 | 说明 |
|------|------|------|
| **amqp.go** | AMQP 连接 | 连接管理 |
| **amqp_bridge.go** | 双向桥接 | 内存↔AMQP 同步 |
| **store.go** | 事件存储 | 可选的持久化 |

### 测试文件（3 个）

| 文件 | 功能 |
|------|------|
| **event_test.go** | 事件系统测试 |
| **amqp_test.go** | AMQP 集成测试 |
| **middleware_test.go** | 中间件测试 |

### 文档文件（4 个）

| 文件 | 位置 | 功能 |
|------|------|------|
| **README.md** | 根目录 | 主入口文档 |
| **INTEGRATION_GUIDE.md** | docs/ | 详细集成指南 |
| **INTEGRATION_SUMMARY.md** | docs/ | 快速参考 |
| **FILE_STRUCTURE.md** | docs/ | 本文件说明 |

---

## 🏗️ 推荐的使用流程

### 1️⃣ 了解阶段
- 阅读 [README.md](./README.md) - 了解整体结构
- 查看 [event.go](./event.go) - 理解 Event 结构

### 2️⃣ 集成阶段
- 参考 [INTEGRATION_GUIDE.md](./docs/INTEGRATION_GUIDE.md) - 详细步骤
- 查看 [app/message/events.go](../../app/message/events.go) - 参考实现

### 3️⃣ 开发阶段
- 定义事件：`app/{service}/events.go`
- 实现处理器：`app/{service}/biz/dal/eventbus/handler.go`
- 注册消费者：`app/{service}/biz/dal/eventbus/events.go`

### 4️⃣ 运维阶段
- 监控消费者状态
- 调整工作线程数
- 查看性能指标

---

## 📊 核心类依赖关系

```
Event
  ├─ Handler (消费事件)
  │   ├─ EventHandlerFunc (函数适配)
  │   ├─ TypedHandler (泛型处理)
  │   └─ Middleware (中间件链)
  │
EventBus (事件总线)
  ├─ Publish (发布事件)
  ├─ Subscribe (简单订阅)
  ├─ SubscribeAsync (异步订阅)
  ├─ SubscribeWithPool (消费者池)
  └─ Use (注册中间件)

ConsumerPool (消费者池)
  ├─ Start (启动工作线程)
  ├─ Consume (添加事件)
  └─ Stop (优雅关闭)

ConsumerRegistry (注册表)
  ├─ RegisterHandler (注册处理器)
  ├─ RegisterConsumer (注册消费者)
  └─ StartAll (启动所有消费者)
```

---

## 🎯 快速参考

### 定义事件
```go
// app/message/events.go
const EventSendUserMessages = "send_user_messages"
```

### 实现处理器
```go
// app/message/biz/dal/eventbus/handler.go
func HandleSendUserMessages(ctx context.Context, event *eventbus.Event) error {
    // 处理逻辑
    return nil
}
```

### 注册消费者
```go
// app/message/biz/dal/eventbus/events.go
func InitMessageConsumers() error {
    registry := eventbus.NewConsumerRegistry()
    registry.RegisterHandler("send_user_messages", 
        eventbus.EventHandlerFunc(HandleSendUserMessages))
    registry.RegisterConsumer("send_user_messages", "send_user_messages", 10)
    return nil
}
```

### 启动消费者
```go
// app/message/main.go
func main() {
    eventbus.InitGlobalEventBus()
    eventbus.InitMessageConsumers()
    eventbus.StartMessageConsumers()
}
```

### 发布事件
```go
// app/message/biz/service/send_user_messages.go
func (s *SendUserMessagesService) Run(req *message.SendUserMessagesReq) error {
    eb := eventbus.GetGlobalEventBus()
    event := eventbus.NewEvent(EventSendUserMessages, req)
    eb.Publish(s.ctx, event)
    return nil
}
```

---

## 📈 性能建议

### 工作线程数

| 业务类型 | 建议值 | 队列大小 |
|---------|------|---------|
| 关键业务 | 15-20 | 5000 |
| 普通业务 | 8-10 | 2000 |
| 后台任务 | 3-5 | 500 |

### 事件发布

- 发布事件本身很快（<1ms）
- 不要等待处理完成
- 使用 fire-and-forget 模式

---

## 📚 文档导航

```
common/eventbus/
├── README.md                 ◄─ 从这里开始
│
└── docs/
    ├── README.md             ◄─ 核心概念
    ├── INTEGRATION_GUIDE.md  ◄─ 集成步骤
    ├── INTEGRATION_SUMMARY.md ◄─ 快速参考
    └── FILE_STRUCTURE.md     ◄─ 本文件
```

---

## ✨ 总结

- **核心文件** 6 个：eventbus 系统的完整实现
- **集成文件** 3 个：AMQP 集成和存储
- **测试文件** 3 个：完整的测试覆盖
- **文档文件** 4 个：清晰的使用指南

整个 eventbus 模块设计清晰、逻辑完整、易于使用！
