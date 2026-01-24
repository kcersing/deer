# 🎯 EventBus 快速导航

> 📌 **新用户必读**：从 [README.md](../README.md) 开始 → 然后查看 [docs/README.md](README.md)

---

## 🚀 我想...

### 了解 EventBus 是什么？
👉 [docs/README.md](README.md)

### 30 秒快速开始？
👉 [README.md](../README.md) 中的 "快速开始" 部分

### 详细集成步骤？
👉 [docs/INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md)

### 看 API 参考？
👉 [docs/INTEGRATION_SUMMARY.md](INTEGRATION_SUMMARY.md)

### 了解文件结构？
👉 [docs/FILE_STRUCTURE.md](FILE_STRUCTURE.md)

### 看参考实现？
👉 [app/message/](../message/) 目录

### 性能优化建议？
👉 [docs/INTEGRATION_GUIDE.md#性能优化建议](INTEGRATION_GUIDE.md)

### 解决常见问题？
👉 [docs/INTEGRATION_GUIDE.md#常见问题](INTEGRATION_GUIDE.md)

---

## 📚 文档地图

```
common/eventbus/ (你在这里)
│
├─ 📖 主文档 (1 分钟)
│  └─ README.md ★ 从这里开始
│
├─ 📚 完整文档 (docs/)
│  ├─ README.md         ← 文档导航 (5 分钟)
│  ├─ FILE_STRUCTURE.md ← 文件说明 (10 分钟)
│  ├─ INTEGRATION_GUIDE.md ← 详细指南 (30 分钟)
│  ├─ INTEGRATION_SUMMARY.md ← 快速参考 (5 分钟)
│  └─ CHECKLIST.md ← 整理清单
│
├─ 🔧 源代码 (12 个 Go 文件)
│  ├─ event.go
│  ├─ handler.go
│  ├── ... (详见 FILE_STRUCTURE.md)
│
└─ 📝 说明文档
   ├─ README.md ← 主入口
   └─ ORGANIZATION_SUMMARY.md ← 整理说明
```

---

## ⏱️ 预计阅读时间

| 文档 | 时间 | 适合 |
|------|------|------|
| README.md | ⏱️ 1 分钟 | 快速了解 |
| docs/README.md | ⏱️ 5 分钟 | 核心概念 |
| FILE_STRUCTURE.md | ⏱️ 10 分钟 | 文件细节 |
| INTEGRATION_GUIDE.md | ⏱️ 30 分钟 | 完整实现 |
| INTEGRATION_SUMMARY.md | ⏱️ 5 分钟 | 快速查询 |

**总计：51 分钟** - 掌握全部内容

---

## 🎓 学习路径

### 初学者 (1-2 小时)
1. README.md (1 min)
2. docs/README.md (5 min)
3. FILE_STRUCTURE.md (10 min)
4. 快速开始示例 (15 min)
5. 参考 app/message/ (30 min)

### 中级开发者 (30 分钟)
1. 快速开始示例 (5 min)
2. INTEGRATION_GUIDE.md (25 min)

### 高级开发者 / 架构师 (15 分钟)
1. FILE_STRUCTURE.md (10 min)
2. 查看源代码 (5 min)

---

## 🔍 按功能快速找文件

### 发布/订阅相关
- **代码** → [eventbus.go](../eventbus.go#L1-L50)
- **文档** → [docs/INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md#使用方式)

### 消费者注册相关
- **代码** → [registry.go](../registry.go)
- **文档** → [docs/INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md#使用方式)

### 高性能并发处理
- **代码** → [consumer_pool.go](../consumer_pool.go)
- **文档** → [docs/FILE_STRUCTURE.md](FILE_STRUCTURE.md#性能建议)

### AMQP 集成
- **代码** → [amqp_bridge.go](../amqp_bridge.go)
- **文档** → [docs/INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md#架构图)

### 中间件扩展
- **代码** → [middleware.go](../middleware.go)
- **文档** → [docs/INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md#中间件链执行顺序)

---

## ✨ 关键概念速查

| 概念 | 说明 | 位置 |
|------|------|------|
| **Event** | 事件消息单元 | event.go / docs/ |
| **Handler** | 事件消费处理器 | handler.go / docs/ |
| **Middleware** | 事件处理链 | middleware.go / docs/ |
| **ConsumerPool** | 高并发处理 | consumer_pool.go / docs/ |
| **ConsumerRegistry** | 消费者管理 | registry.go / docs/ |
| **AMQPBridge** | 分布式集成 | amqp_bridge.go / docs/ |

---

## 💡 常见场景

### 场景 1：我需要快速发布事件
```
1. 读：README.md 快速开始
2. 查：docs/INTEGRATION_SUMMARY.md API
3. 测试：app/message/biz/service/
```

### 场景 2：我需要处理高并发事件
```
1. 读：docs/INTEGRATION_GUIDE.md 性能建议
2. 查：docs/FILE_STRUCTURE.md 消费者池
3. 参考：app/message/biz/dal/eventbus/events.go 配置
```

### 场景 3：我需要分布式集成
```
1. 读：docs/INTEGRATION_GUIDE.md AMQP 部分
2. 查：amqp_bridge.go 实现
3. 配置：docker-compose.yml RabbitMQ
```

### 场景 4：我需要理解架构
```
1. 读：docs/README.md 概念
2. 查：docs/FILE_STRUCTURE.md 设计
3. 看：docs/INTEGRATION_GUIDE.md 架构图
```

---

## 🆘 问题排查

| 问题 | 原因 | 解决 |
|------|------|------|
| 不知道怎么开始 | 没读主文档 | → 读 README.md |
| 事件没被消费 | 没启动消费者 | → 看 docs/INTEGRATION_GUIDE.md |
| 性能很慢 | 没用消费者池 | → 读 docs/FILE_STRUCTURE.md |
| AMQP 连不上 | 配置问题 | → 查 docker-compose.yml |
| 代码哪里写 | 不清楚文件用途 | → 看 FILE_STRUCTURE.md |

---

## 🎯 下一步

**👉 点击这里开始：[README.md](../README.md)**

1. 1 分钟了解整体
2. 查看快速开始
3. 阅读详细文档
4. 查看参考实现
5. 开始编码

---

**祝你使用愉快！** 🚀
