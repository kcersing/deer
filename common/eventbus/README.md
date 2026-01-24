# 🚌 EventBus 事件总线框架

一个轻量级、高效的事件总线实现，支持内存事件、异步处理、消费者池和 AMQP 桥接。

## 📁 模块组织

### 核心框架层 (6 个)
```
├── event.go               # 事件定义：Event 结构体
├── handler.go             # 处理器接口：Handler, EventHandlerFunc, TypedHandler
├── eventbus.go            # 事件总线核心：Publish/Subscribe/分发
├── middleware.go          # 中间件机制：日志、过滤、转换
├── consumer_pool.go       # 消费者池：并发处理、工作线程管理
└── registry.go            # 消费者注册表：集中管理和统一启动
```
**总计 ~500 行** - 完整的事件处理框架

### AMQP 集成层 (3 个)
```
├── amqp.go                # AMQP 连接和管理
├── amqp_bridge.go         # 内存事件 ↔ AMQP 双向桥接
└── store.go               # 可选的事件存储
```
**支持分布式消息** - RabbitMQ 集成

### 测试覆盖 (3 个)
```
├── event_test.go          # 事件系统测试
├── amqp_test.go           # AMQP 集成测试
└── middleware_test.go     # 中间件链测试
```

### 📚 文档目录
```
docs/
├── README.md              # 文档导航和快速参考
├── FILE_STRUCTURE.md      # 详细文件说明（你在这里）
├── INTEGRATION_GUIDE.md   # 完整集成指南
└── INTEGRATION_SUMMARY.md # 快速查询表
```
**从 [docs/README.md](./docs/README.md) 开始阅读！**

## � 获取帮助

👉 **首先阅读：[docs/README.md](./docs/README.md)** - 完整文档导航  

其他资源：
- [FILE_STRUCTURE.md](./docs/FILE_STRUCTURE.md) - 核心文件说明和快速参考
- [INTEGRATION_GUIDE.md](./docs/INTEGRATION_GUIDE.md) - 详细集成步骤
- [INTEGRATION_SUMMARY.md](./docs/INTEGRATION_SUMMARY.md) - API 快速查询

## 🚀 快速开始（30秒）

### 1️⃣ 初始化全局事件总线
```go
// app/main.go
func init() {
    eventbus.InitGlobalEventBus()
}
```

### 2️⃣ 定义服务事件
```go
// app/message/events.go
const (
    EventSendUserMessages = "send_user_messages"
    EventMessageSent      = "message_sent"
)
```

### 3️⃣ 实现事件处理器
```go
// app/message/biz/dal/eventbus/handler.go
func HandleSendUserMessages(ctx context.Context, event *eventbus.Event) error {
    // 处理事件
    return nil
}
```

### 4️⃣ 注册并启动消费者
```go
// app/message/biz/dal/eventbus/events.go
func InitMessageConsumers() error {
    registry := eventbus.NewConsumerRegistry()
    registry.RegisterHandler("send_user_messages", 
        eventbus.EventHandlerFunc(HandleSendUserMessages))
    registry.RegisterConsumer("send_user_messages", "send_user_messages", 10)
    return nil
}

func StartMessageConsumers() error {
    registry := eventbus.GetConsumerRegistry()
    return registry.StartAll(eventbus.GetGlobalEventBus())
}
```

### 5️⃣ 发布事件（在服务中）
```go
// app/message/biz/service/send_user_messages.go
func (s *SendUserMessagesService) Run(req *Request) error {
    eb := eventbus.GetGlobalEventBus()
    event := eventbus.NewEvent(messageevent.EventSendUserMessages, req)
    return eb.Publish(s.ctx, event)
}
```

## 🏗️ 架构层次

```
业务层 (app/{service}/)
  ├─ events.go                    # 事件常量定义
  └─ biz/dal/eventbus/
      ├─ handler.go                # 处理器实现
      └─ events.go                 # 消费者初始化和启动
          ↓
框架层 (common/eventbus/)          ← 你在这里
  ├─ eventbus.go                  # 发布/订阅/分发
  ├─ registry.go                  # 消费者注册表
  ├─ consumer_pool.go             # 并发处理
  ├─ amqp_bridge.go               # AMQP 集成
  └─ middleware.go                # 中间件链
```

## 💡 三种订阅方式对比

| 方式 | 代码 | 场景 | 性能 |
|------|------|------|------|
| 简单订阅 | `eb.Subscribe(topic)` | 简单测试 | ⭐ 低 |
| 异步订阅 | `eb.SubscribeAsync(topic, handler, 3)` | 中等场景 | ⭐⭐ 中 |
| 消费者池 | `registry.RegisterConsumer(...)` | 高吞吐 | ⭐⭐⭐⭐⭐ 极高 |

## ✨ 核心特性

✅ **高性能** - 内存事件总线，微秒级延迟  
✅ **高吞吐** - 消费者池支持并发处理，可配置工作线程  
✅ **分布式** - AMQP 桥接，支持微服务通信  
✅ **易扩展** - 中间件链，灵活定制  
✅ **可靠性** - 消费者注册表集中管理，统一启动/停止  
✅ 消费者池支持  
✅ 中间件扩展机制  
✅ AMQP 双向桥接  
✅ 完整错误处理  

## 🔗 相关代码

- 消息服务示例：`app/message/events.go`
- 事件初始化：`app/message/biz/dal/eventbus/events.go`
- 处理器实现：`app/message/biz/dal/eventbus/handler.go`
