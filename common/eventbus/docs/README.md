# 📚 EventBus 文档

欢迎阅读 EventBus 事件总线的完整文档！

## 📖 文档导航

### 1. **快速入门** 🚀
- [FILE_STRUCTURE.md](./FILE_STRUCTURE.md) - 了解文件结构和快速参考

### 2. **详细指南** 📘
- [INTEGRATION_GUIDE.md](./INTEGRATION_GUIDE.md) - 完整的集成指南和最佳实践
- [INTEGRATION_SUMMARY.md](./INTEGRATION_SUMMARY.md) - 快速查询和参考

### 3. **核心概念** 💡
- Event 事件：消息单元，包含主题、载体、时间戳等
- Handler 处理器：事件消费的业务逻辑
- Middleware 中间件：事件处理链中的拦截器
- ConsumerPool 消费者池：高性能并发处理
- ConsumerRegistry 消费者注册表：集中管理消费者

---

## 🎯 按场景快速找文档

| 场景 | 推荐文档 | 重点内容 |
|------|---------|---------|
| 我是初学者 | FILE_STRUCTURE.md | 文件结构、快速参考 |
| 我要集成 EventBus | INTEGRATION_GUIDE.md | 4 种集成方式、完整示例 |
| 我要快速查询 | INTEGRATION_SUMMARY.md | API 参考、常见问题 |
| 我要看性能优化 | INTEGRATION_GUIDE.md | 性能建议、最佳实践 |
| 我要看参考实现 | FILE_STRUCTURE.md | app/message 服务示例 |

---

## 🏗️ EventBus 架构简览

```
业务服务 (app/{service})
    ↓ 发布事件
EventBus 内存总线
    ├─ 本地快速处理 (microseconds)
    ├─ 中间件链
    └─ 消费者池 (并发处理)
        ↓
    RabbitMQ 分布式消息队列
        └─ 其他微服务消费
        └─ 外部系统集成
```

---

## ⚡ 核心特性

✅ **高性能** - 内存事件总线，微秒级处理  
✅ **高吞吐** - 消费者池，支持 1000+ 并发  
✅ **分布式** - AMQP 桥接，支持微服务协作  
✅ **可靠性** - 事件持久化，错误重试  
✅ **易扩展** - 中间件链，灵活定制  

---

## 📚 更多资源

- [README.md](../README.md) - 主入口文档
- [源码位置](../) - common/eventbus 源代码

---

## 💬 常见问题速查

**Q: 事件发布是同步还是异步？**  
A: 发布本身是同步的，但中间件和消费者处理可以异步。

**Q: 如何处理大量并发事件？**  
A: 使用 ConsumerPool 和配置多个工作线程。

**Q: 事件会丢失吗？**  
A: 内存事件可能丢失（重启丢失），但通过 AMQP 桥接可以持久化。

**Q: 多个消费者如何处理同一个事件？**  
A: 每个消费者独立处理，互不影响。

**更详细的问题查看** [INTEGRATION_GUIDE.md](./INTEGRATION_GUIDE.md#常见问题) 

---

## 🔧 快速命令

```bash
# 启动 EventBus（在 main.go 中）
eventbus.InitGlobalEventBus()
eventbus.InitMessageConsumers()
eventbus.StartMessageConsumers()

# 发布事件（在服务中）
eb := eventbus.GetGlobalEventBus()
event := eventbus.NewEvent("event_topic", data)
eb.Publish(ctx, event)

# 优雅关闭
eventbus.StopMessageConsumers()
```

---

最后更新：2025年  
版本：1.0.0
