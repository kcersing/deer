# eventbus

## 事件总线（EventBus）
#### 是一种实现了发布/订阅（Publish/Subscribe）模式的架构模式。它充当生产者（发布者）和消费者（订阅者）之间的中间件，允许系统中的各个组件进行通信，而无需它们之间建立直接依赖关系，从而实现了解耦。在 Go 语言中，EventBus 的实现通常利用 Go 的核心特性，如 goroutine 和 channel，来实现高效、并发的消息传递。

### EventBus 的核心概念

#### 事件（Event/Message）: 表示系统中发生的事情，通常是一个包含相关数据结构的载体。
#### 发布者（Publisher）: 生成事件并将其发送到事件总线（特定主题）的组件。
#### 订阅者（Subscriber/Handler）: 注册（订阅）特定类型事件，并在事件发生时执行相应处理逻辑的组件或函数。
#### 主题（Topic）: 用于对事件进行分类的标识符，订阅者通过主题来指定感兴趣的事件类型。


## RabbitMQ 核心概念
#### 1：Connection：与 RabbitMQ 服务器的 TCP 连接。
#### 2：Channel：在连接内部建立的轻量级通信通道，大多数 API 操作都在其上进行。
#### 3：Exchange：消息到达 RabbitMQ 后的第一站。它根据路由键（Routing Key）将消息路由到一个或多个队列。
#### 4：Queue：存储消息的命名缓冲区。
#### 5：Binding：将 Exchange 与 Queue 连接起来的规则（基于路由键）。
