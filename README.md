# kcers-order

### DDD和事件溯源模式

#### RPC 调用的流程
一次 rpc 调用包括以下基本流程，分为客户端和服务端两个部分：

（客户端）构造请求参数，发起调用
（客户端）通过服务发现、负载均衡等得到服务端实例地址，并建立连接
（客户端）请求参数序列化成二进制数据
（客户端）通过网络将数据发送给服务端
（服务端）服务端接收数据
（服务端）反序列化出请求参数
（服务端）handler 处理请求并返回响应结果
（服务端）将响应结果序列化成二进制数据
（服务端）通过网络将数据返回给客户端
（客户端）接收数据
（客户端）反序列化出结果
（客户端）得到调用的结果


┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│    领域模型层    │     │    基础设施层    │     │    事件处理层    │
│  order.go       │────▶│  orderRepository.go │────▶│  observers.go   │
│  orderStateMachine.go │  subscription_service.go │  orderEvent.go  │
└─────────────────┘     └─────────────────┘     └─────────────────┘

1. 创建订单 → NewOrder()
   ├─ 初始化状态机
   └─ 生成OrderCreatedEvent

2. 支付订单 → order.Pay()
   ├─ 生成OrderPaidEvent
   ├─ 状态机验证状态转换
   └─ 更新订单状态为"paid"

3. 保存订单 → repository.Save()
   ├─ 事务保存订单数据
   ├─ 事务保存事件记录
   └─ 提交后通知订阅者

4. 事件处理 → subscriptionService.ProcessEvent()
   ├─ 查询订阅者
   ├─ 通知外部系统
   └─ 更新订阅状态


biz/infras/
├── common/           # 通用组件：聚合根、事件、错误定义
├── order/            # 订单领域相关实现
│   ├── aggregate/    # 订单聚合根
│   ├── events/       # 订单事件定义
│   └── repo/         # 订单仓储
├── status/           # 状态定义
├── observers.go      # 事件分发器
├── subscription_service.go # 订阅服务
└── order_test.go     # 测试代码



cwgo server --type RPC --idl idl/order.thrift --server_name order --module kcers-order --hex 




go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./biz/dal/db/mysql/ent/schema
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./biz/dal/db/mysql/ent/schema
