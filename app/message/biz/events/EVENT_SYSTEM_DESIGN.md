# 事件系统设计说明

## 🏗️ 架构

```
common/eventbus/                        ← 事件总线核心 + 消费者框架
├── eventbus.go                         # 事件总线 + ConsumerPool
├── registry.go                         # ConsumerRegistry（集中管理）
└── ...

app/{service}/                          ← 具体业务实现
├── events.go                           # 事件常量定义
└── biz/dal/eventbus/
    ├── handler.go                      # 处理器实现
    ├── events.go                       # 消费者注册和启动
    └── ...
```

## 🔄 使用流程

### 1. 定义事件

```go
// app/message/events.go
const (
    EventSendUserMessages = "send_user_messages"
    EventMessageSent      = "message_sent"
)
```

### 2. 实现处理器

```go
// app/message/biz/dal/eventbus/handler.go
func HandleSendUserMessages(ctx context.Context, event *eventbus.Event) error {
    req := event.Payload.(*message.SendUserMessagesReq)
    // 业务逻辑
    return nil
}
```

### 3. 注册消费者

```go
// app/message/biz/dal/eventbus/events.go
func InitMessageConsumers() error {
    consumerRegistry := eventbus.NewConsumerRegistry()
    
    // 注册处理器
    consumerRegistry.RegisterHandler("send_user_messages", 
        eventbus.EventHandlerFunc(HandleSendUserMessages))
    
    // 注册消费者
    consumerRegistry.RegisterConsumer(
        EventSendUserMessages,
        "send_user_messages",
        10,  // 10 个工作线程
    )
    
    return nil
}

func StartMessageConsumers() error {
    return consumerRegistry.StartAll(globalEventBus)
}
```

### 4. 启动应用

```go
// app/message/main.go
func init() {
    eventbus.InitGlobalEventBus()
    eventbus.InitMessageConsumers()
}

func main() {
    eventbus.StartMessageConsumers()
    // 消费者已启动
}
```

### 5. 发布事件

```go
// app/message/biz/service/send_user_messages.go
func (s *SendUserMessagesService) Run(req *message.SendUserMessagesReq) error {
    eb := eventbus.GetGlobalEventBus()
    
    // 只需发布，消费由已注册的消费者处理
    event := eventbus.NewEvent(EventSendUserMessages, req)
    eb.Publish(s.ctx, event)
    
    return nil
}
```

## 📊 核心类

### ConsumerPool（消费者池）

在 `common/eventbus/eventbus.go` 中：

```go
type ConsumerPool struct {
    name      string
    handler   Handler
    workerNum int32
    queue     chan *Event
    // ...
}

// 并发处理事件，工作线程数可配置
```

### ConsumerRegistry（消费者注册表）

在 `common/eventbus/registry.go` 中：

```go
type ConsumerRegistry struct {
    handlers  map[string]Handler          // 处理器
    consumers map[string][]*ConsumerInfo  // 消费者
    pools     map[string]*ConsumerPool    // 消费者池
}

// 集中管理所有处理器和消费者
```

## 🎯 优势

✅ **清晰分层** - 框架层与业务层分离  
✅ **集中管理** - 所有消费者在启动时注册  
✅ **高性能** - 支持并发处理、可配置工作线程  
✅ **易于扩展** - 添加新事件只需 5 步  
✅ **无文件夹冲突** - 所有代码都在现有结构中  

## 📁 文件清单

```
common/eventbus/
├── eventbus.go          ✅ 添加：ConsumerPool
├── registry.go          ✅ 新增：ConsumerRegistry
├── handler.go
├── event.go
└── ...

app/message/
├── events.go            ✅ 新增：事件定义
└── biz/dal/eventbus/
    ├── handler.go       ✅ 新增：处理器实现
    ├── events.go        ✅ 修改：添加消费者注册
    └── service/
        └── send_user_messages.go  ✅ 简化：只发布事件
```

## 💡 最佳实践

| 操作 | 位置 | 说明 |
|------|------|------|
| 定义事件 | `app/{service}/events.go` | 服务独立定义 |
| 实现处理 | `app/{service}/biz/dal/eventbus/handler.go` | 业务逻辑实现 |
| 注册消费者 | `app/{service}/biz/dal/eventbus/events.go` | 启动时注册 |
| 添加框架功能 | `common/eventbus/` | 在现有文件夹中扩展 |

## 🚀 总结

这个设计的核心特点：
- 所有改动都在现有的 `common/eventbus/` 目录下
- 不需要创建新文件夹
- 具体业务定义在各服务模块
- 框架代码集中在 eventbus 模块
- 简洁、清晰、易于维护
